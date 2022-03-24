CREATE TABLE "universal_champion" (
  "id" bigserial PRIMARY KEY,
  "title_holder" varchar NOT NULL,
  "title_holder_number" int NOT NULL,
  "date_won" timestamp NOT NULL,
  "date_lost" timestamp,
  "show" varchar NOT NULL,
  "current_champion" boolean NOT NULL,
  "title_holder_order_number" int NOT NULL,
  "wrestler_id" bigint
);

CREATE TABLE "intercontinental_champion" (
  "id" bigserial PRIMARY KEY,
  "title_holder" varchar NOT NULL,
  "title_holder_number" int NOT NULL,
  "date_won" timestamp NOT NULL,
  "date_lost" timestamp,
  "show" varchar NOT NULL,
  "current_champion" boolean NOT NULL,
  "title_holder_order_number" int NOT NULL,
  "wrestler_id" bigint
);

CREATE INDEX ON "universal_champion" ("title_holder");

CREATE INDEX ON "universal_champion" ("date_won", "date_lost");

CREATE INDEX ON "champion" ("title_holder");

CREATE INDEX ON "champion" ("date_won", "date_lost");

CREATE INDEX ON "intercontinental_champion" ("title_holder");

CREATE INDEX ON "intercontinental_champion" ("date_won", "date_lost");
