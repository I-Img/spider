CREATE TABLE "public"."Untitled" (
  "id" int4 NOT NULL,
  "name" varchar(255) NOT NULL,
  "last_time" timestamp,
  "pid" int8,
  PRIMARY KEY ("id", "name")
)
;

COMMENT ON COLUMN "public"."Untitled"."name" IS '类型名称';

COMMENT ON COLUMN "public"."Untitled"."last_time" IS '收录时间';

COMMENT ON COLUMN "public"."Untitled"."pid" IS '图片ID';