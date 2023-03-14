CREATE TABLE "movie" (
   "id" uuid DEFAULT gen_random_uuid() NOT NULL,
   "director_id" uuid NOT NULL,
   "title" VARCHAR(255) NOT NULL,
   "ganre" VARCHAR(255) NOT NULL,
   "rate" INT NOT NULL,
   "release_date" Timestamp Without Time Zone NOT NULL,
   "duration" INT NOT NULL DEFAULT 0,
   "created" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
   "modified" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
   PRIMARY KEY ("id"),
   CONSTRAINT "unique_movie_id" UNIQUE("id"),
   CONSTRAINT "diretor_fk" FOREIGN KEY (director_id) REFERENCES public."director"(id)
);

CREATE TRIGGER update_movie_modtime 
BEFORE UPDATE ON "movie" 
FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();
