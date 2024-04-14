DROP SCHEMA IF EXISTS banner_schema CASCADE;

CREATE SCHEMA IF NOT EXISTS banner_schema;

CREATE TABLE IF NOT EXISTS banner_schema.banners (
    id bigserial primary key,
    enable boolean not null default true,
    content jsonb not null
);

CREATE TABLE IF NOT EXISTS banner_schema.tags (
    id bigserial primary key
);

CREATE TABLE IF NOT EXISTS banner_schema.features (
    id bigserial primary key
);

CREATE TABLE IF NOT EXISTS banner_schema.banner_tag (
    banner_id bigserial not null references banner_schema.banners(id) on delete cascade,
    tag_id bigserial not null references banner_schema.tags(id) on delete cascade,
    UNIQUE (banner_id, tag_id)
);

CREATE TABLE IF NOT EXISTS banner_schema.banner_feature (
    banner_id bigserial not null references banner_schema.banners (id) on delete cascade primary key,
    feature_id bigserial not null references banner_schema.features (id) on delete cascade
);