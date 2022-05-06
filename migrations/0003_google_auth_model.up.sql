ALTER TABLE "users" DROP COLUMN "first_name";

ALTER TABLE "users" ALTER COLUMN "email" DROP NOT NULL;

CREATE TABLE "google_id_token" (
    "sub" varchar(255) PRIMARY KEY,
    "user" uuid NULL REFERENCES "users" ("id") ON DELETE CASCADE,
    "iat" integer NOT NULL,
    "exp" integer NOT NULL
);