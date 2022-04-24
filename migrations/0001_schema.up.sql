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
    "author" uuid NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    "text" varchar(3000) NOT NULL
);