alter table "public"."grades" alter column "test" drop not null;
alter table "public"."grades" add column "test" text;
