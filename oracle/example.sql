
CREATE TABLE "SCOTT"."table1" (
    "id" NUMBER NOT NULL,
    "name" VARCHAR2(255),
    "json" VARCHAR2(255),
    "ref" NUMBER,
    PRIMARY KEY ("id")
);

CREATE TABLE "SCOTT"."table2" (
  "id" NUMBER NOT NULL,
  "name" VARCHAR2(255),
  "age" NUMBER,
  "ref" NUMBER,
  PRIMARY KEY ("id")
);

CREATE TABLE "SCOTT"."table3" (
  "id" NUMBER NOT NULL,
  "name" VARCHAR2(255),
  "info" VARCHAR2(255),
  "ref" NUMBER,
  PRIMARY KEY ("id")
);
