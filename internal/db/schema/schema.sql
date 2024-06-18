CREATE TABLE "urls" (
  "id" serial PRIMARY KEY,
  "owner_id" uuid NOT NULL,
  "short_url" varchar(7) UNIQUE NOT NULL,
  "long_url" varchar(400) NOT NULL
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "urls" ("short_url");

CREATE INDEX ON "urls" ("long_url");

CREATE INDEX ON "users" ("username");

ALTER TABLE "urls" ADD CONSTRAINT "unique_long_url_per_owner" UNIQUE ("long_url", "owner_id");

ALTER TABLE "urls" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");
