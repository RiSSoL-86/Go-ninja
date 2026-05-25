-- Create "calculations" table
CREATE TABLE "public"."calculations" (
  "id" bigserial NOT NULL,
  "args" double precision[] NULL,
  "operator" text NULL,
  "result" double precision NULL,
  "note" text NULL,
  "created_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
