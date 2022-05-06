ALTER TABLE "users" ADD COLUMN "first_name" varchar(255) NULL;

ALTER TABLE "users" ALTER COLUMN "email" SET NOT NULL;

DROP TABLE "google_id_token";