CREATE TABLE "user" (
    "id" uuid PRIMARY KEY,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "first_name" varchar(255) NULL,
    "username" varchar(255) NOT NULL,
    "email" varchar(320) NOT NULL
);

CREATE TABLE "post" (
    "id" uuid PRIMARY KEY,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "author" uuid REFERENCES "user" ("id"),
    "text" varchar(3000) NOT NULL
);