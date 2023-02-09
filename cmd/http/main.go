package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"

	"github.com/urfave/cli/v2"

	migrations "opensea/cmd/migrations"
	restcontroller "opensea/controllers/rest"
	openseaservice "opensea/internal/service"
	postgresrepository "opensea/repositories/postgres"
)

func main() {

	ctx := context.Background()
	// 192.168.99.100:5432
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr("postgres:5432"),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser("postgres"),
		pgdriver.WithPassword("password"),
		pgdriver.WithDatabase("opensea")))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())
	postgresrepository.RegisterModel(ctx, db)

	// Run migrations on every app run
	migrator := migrate.NewMigrator(db, migrations.Migrations)
	migrator.Init(ctx)
	_, err := migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}

	// This is just for accepting flags to database migration
	/*
		app := &cli.App{
			Commands: []*cli.Command{
				newDBCommand(migrate.NewMigrator(db, migrations.Migrations)),
			},
		}
		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
	*/
	repository := postgresrepository.NewPostgresRepository(db)
	service := openseaservice.NewService(repository)
	controller := restcontroller.NewController(service)

	router := gin.New()
	router.MaxMultipartMemory = 10 << 20 // 20 MiB

	router.POST("movie", controller.Create)
	router.GET("movie/:id", controller.Get)
	router.GET("movies/:page", controller.GetAll)
	router.POST("movie/buy/:id", controller.Buy)

	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8081")
}

func newDBCommand(migrator *migrate.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to run (database is up to date)\n")
						return nil
					}
					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}
					fmt.Printf("rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					return migrator.Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					return migrator.Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}
					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}
