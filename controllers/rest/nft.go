package restcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"opensea/internal/domain"
	"opensea/shared"
)

func (c *RestController) Create(ctx *gin.Context) {
	shared.WithUserRequestContext(ctx, 1) // 1 test creator user id

	// 1. Process the file upload
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Printf(err.Error())
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	// TODO: Generate an unique name for the filename
	filename := filepath.Base(file.Filename)
	if err := ctx.SaveUploadedFile(file, "./public/"+filename); err != nil {
		log.Printf(err.Error())
		ctx.AbortWithStatusJSON(500, gin.H{"upload file error": err.Error()})
		return
	}

	// 2. Process the movie data
	var newMovie domain.Movie
	newMovie.CreatorID = 1 // 1 test creator user id
	newMovie.File = filename
	newMovie.Description, _ = ctx.GetPostForm("description")
	if newMovie.Description == "" {
		log.Printf(err.Error())
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "Missing required field: description"})
	}

	p, _ := ctx.GetPostForm("price")
	newMovie.Price, _ = strconv.ParseFloat(p, 64)
	if newMovie.Price == 0 {
		log.Printf(err.Error())
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "Missing required field: price"})
	}

	//TODO: Validate te array, must be array of integers
	c_creators_str, _ := ctx.GetPostForm("co_creators")
	var c_creators []int
	err = json.Unmarshal([]byte(c_creators_str), &c_creators)
	if err != nil {
		log.Printf(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "There was a problem with this field: co_creators"})
	}

	var coCreators []*domain.StreamingPlatform
	for _, c_creator := range c_creators {
		coCreators = append(coCreators, &domain.StreamingPlatform{ID: int64(c_creator)})
	}

	newMovie.CoCreators = coCreators

	err = c.service.Create(ctx, &newMovie)
	if err != nil {
		log.Printf(err.Error())
		ctx.IndentedJSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusCreated, newMovie)
}

func (c *RestController) GetAll(ctx *gin.Context) {
	shared.WithUserRequestContext(ctx, 2) // 2 test buyer user id
	// just some validation of the page parameter
	// in case of anything we default to 0
	page, _ := strconv.Atoi(ctx.Param("page"))
	if page < 0 {
		page = 1
	}

	movies, err := c.service.GetAll(ctx, page)

	if err != nil {
		log.Printf(err.Error())
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusCreated, movies)
}

func (c *RestController) Get(ctx *gin.Context) {
	shared.WithUserRequestContext(ctx, 2) // 2 test buyer id

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	movie, err := c.service.Get(ctx, id)

	if err != nil {
		log.Printf(err.Error())
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, movie)
}

func (c *RestController) Buy(ctx *gin.Context) {
	shared.WithUserRequestContext(ctx, 2) // 2 test buyer user id

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := c.service.Buy(ctx, id)
	if err != nil {
		log.Printf(err.Error())
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "enjoy it!"})
}
