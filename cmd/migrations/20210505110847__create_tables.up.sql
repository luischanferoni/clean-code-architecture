CREATE TABLE IF NOT EXISTS movie (
  id bigserial PRIMARY KEY,
  file varchar,
  description varchar(256),
  creator_id bigserial,
  price float,
  created_at bigint,
  created_by bigint,
  updated_at bigint,
  updated_by bigint
);

--bun:split

CREATE TABLE IF NOT EXISTS "user" (
  id bigserial PRIMARY KEY,
  "name" varchar(256),
  balance float
);

--bun:split

ALTER TABLE movie ADD CONSTRAINT "fk_creator_id" FOREIGN KEY (creator_id) REFERENCES "user"(id);

--bun:split

CREATE TABLE IF NOT EXISTS movie_to_cocreator (
  user_id bigserial,
  movie_id bigserial
);

ALTER TABLE movie_to_cocreator ADD CONSTRAINT "fk_user_id" FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE;
ALTER TABLE movie_to_cocreator ADD CONSTRAINT "fk_movie_id" FOREIGN KEY (movie_id) REFERENCES movie(id) ON UPDATE CASCADE;