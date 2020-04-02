CREATE TABLE "public"."content" (
  "id" int8 NOT NULL,
  "name" varchar(255) NOT NULL,
  "auth" varchar(255),
  "crawl_time" date NOT NULL,
  "content" varchar(255),
  "c_type" int8 NOT NULL,
  "width" float8,
  "height" float8,
  "remarks" varchar(255),

  PRIMARY KEY ("id")
)
;

COMMENT ON COLUMN "public"."Untitled"."id" IS '内容ID';

COMMENT ON COLUMN "public"."Untitled"."name" IS '作品名称';

COMMENT ON COLUMN "public"."Untitled"."auth" IS '主人公名称';

COMMENT ON COLUMN "public"."Untitled"."crawl_time" IS '收录时间戳';

COMMENT ON COLUMN "public"."Untitled"."content" IS '内容';

COMMENT ON COLUMN "public"."Untitled"."c_type" IS '内容类型';

COMMENT ON COLUMN "public"."Untitled"."width" IS '图片高度';

COMMENT ON COLUMN "public"."Untitled"."height" IS '图片宽度';

COMMENT ON COLUMN "public"."Untitled"."remarks" IS '备注';

