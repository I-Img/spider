-- ----------------------------
-- Table structure for px500pic
-- ----------------------------
DROP TABLE IF EXISTS "public"."px500pic";
CREATE TABLE "public"."px500pic" (
  "id" int8 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "auth" varchar(255) COLLATE "pg_catalog"."default",
  "width" int4,
  "height" int4,
  "url" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6) NOT NULL,
  "remark" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."px500pic" OWNER TO "iimg";
COMMENT ON COLUMN "public"."px500pic"."id" IS 'px500图片ID';
COMMENT ON COLUMN "public"."px500pic"."name" IS '图片名称';
COMMENT ON COLUMN "public"."px500pic"."auth" IS '图片作者';
COMMENT ON COLUMN "public"."px500pic"."width" IS '宽';
COMMENT ON COLUMN "public"."px500pic"."height" IS '高';
COMMENT ON COLUMN "public"."px500pic"."url" IS '访问地址';
COMMENT ON COLUMN "public"."px500pic"."create_time" IS '收入时间戳';
COMMENT ON COLUMN "public"."px500pic"."remark" IS '备注';

-- ----------------------------
-- Primary Key structure for table px500pic
-- ----------------------------
ALTER TABLE "public"."px500pic" ADD CONSTRAINT "px500pic_pkey" PRIMARY KEY ("id");
