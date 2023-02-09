ALTER TABLE movie DROP CONSTRAINT "fk_creator_id";
ALTER TABLE movie_to_cocreator DROP CONSTRAINT "fk_user_id";
ALTER TABLE movie_to_cocreator DROP CONSTRAINT "fk_movie_id";
DROP TABLE IF EXISTS movie;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS movie_to_cocreator;