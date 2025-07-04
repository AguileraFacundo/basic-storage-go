CREATE TABLE "sales" (
  "id" bigserial PRIMARY KEY,
  "balance" bigint NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "balance" bigint NOT NULL,
  "supplier_id" bigint NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "debts" (
  "id" bigserial PRIMARY KEY,
  "supplier_id" bigint NOT NULL,
  "balance" bigint NOT NULL,
  "paid" bool NOT NULL DEFAULT false,
  "date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "supplier" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "creation_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "payments" ("supplier_id");

CREATE INDEX ON "debts" ("supplier_id");

ALTER TABLE "payments" ADD FOREIGN KEY ("supplier_id") REFERENCES "supplier" ("id");

ALTER TABLE "debts" ADD FOREIGN KEY ("supplier_id") REFERENCES "supplier" ("id");