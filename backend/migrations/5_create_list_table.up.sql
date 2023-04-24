CREATE TABLE "list" (
    "id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "account_id" uuid NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "created" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
    "modified" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT "unique_list_id" UNIQUE("id"),
    CONSTRAINT "unique_account_id_type" UNIQUE("account_id", "type"),
    CONSTRAINT "movie_fk" FOREIGN KEY (movie_id) REFERENCES public."movie"(id),
    CONSTRAINT "account_fk" FOREIGN KEY (account_id) REFERENCES public."account"(id)
);

CREATE TRIGGER update_list_modtime 
BEFORE UPDATE ON "list" 
FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();
