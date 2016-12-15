CREATE TABLE "public"."products" (
  "id" text,
  "name" text,
  "brand" text,
  "category" text,
  "variant" text,
  "code" text,
  "image" text,
  PRIMARY KEY ("id")
);

CREATE TABLE "public"."product_prices" (
  "id" serial,
  "product_id" text,
  "price" float,
  "sale" bool,
  "timestamp" timestamp default current_timestamp,
  PRIMARY KEY ("id")
);
