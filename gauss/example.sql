
CREATE TABLE "public"."table1" (
                                     "id" serial NOT NULL,
                                     "name" varchar(255),
                                     "json" varchar(255),
                                     "ref" int8,
                                     PRIMARY KEY ("id")
);

CREATE TABLE "public"."table2" (
                                   "id" serial NOT NULL,
                                   "name" varchar(255),
                                   "age" int4,
                                   "ref" int8,
                                   PRIMARY KEY ("id")
);

CREATE TABLE "public"."table3" (
                                   "id" serial NOT NULL,
                                   "name" varchar(255),
                                   "info" varchar(255),
                                   "ref" int8,
                                   PRIMARY KEY ("id")
);
