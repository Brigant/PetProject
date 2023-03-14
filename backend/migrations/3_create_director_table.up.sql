CREATE TABLE "director" (
    "id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "birth_date" Timestamp Without Time Zone NOT NULL,
    "created" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
    "modified" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT "unique_director_id" UNIQUE("id")
);

CREATE TRIGGER update_director_modtime 
BEFORE UPDATE ON "director" 
FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();
