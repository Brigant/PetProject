CREATE TABLE "account" (
  "id" uuid DEFAULT gen_random_uuid() NOT NULL,
   "phone" VARCHAR(25) NOT NULL,
   "password" VARCHAR(255) NOT NULL,
   "age" INT NOT NULL,
   "role" VARCHAR(255) NOT NULL,
   "created" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
   "modified" Timestamp With Time Zone NOT NULL DEFAULT NOW(),
   PRIMARY KEY ("id"),
   CONSTRAINT "unique_account_id" UNIQUE("id"),
   CONSTRAINT "unique_account_phone" UNIQUE("phone")
);

CREATE FUNCTION update_modified_column()   
RETURNS TRIGGER AS $$
BEGIN
   IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
      NEW.modified = now(); 
      RETURN NEW;
   ELSE
      RETURN OLD;
   END IF;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_account_modtime 
BEFORE UPDATE ON "account" 
FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();
