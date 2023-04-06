
CREATE DATABASE "example";

use "example";

CREATE TABLE "table1" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" varchar(255),
    "json" varchar(255),
    "ref" integer
);

CREATE TABLE `table2`  (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" varchar(255),
    "age" INTEGER,
    "ref" INTEGER
);

CREATE TABLE `table3`  (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" varchar(255),
    "info" varchar(255),
    "ref" INTEGER
);
