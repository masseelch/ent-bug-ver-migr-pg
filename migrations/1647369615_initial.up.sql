CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "type" character varying NOT NULL DEFAULT '"PERSON"', PRIMARY KEY ("id"));