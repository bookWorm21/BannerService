INSERT INTO banner_schema.banners VALUES(1, true, '{"title": "some_title", "text": "some_text", "url": "some_url"}');
INSERT INTO banner_schema.tags VALUES (1);
INSERT INTO banner_schema.tags VALUES (2);
INSERT INTO banner_schema.features VALUES(1);
INSERT INTO banner_schema.banner_feature VALUES(1, 1);
INSERT INTO banner_schema.banner_tag VALUES(1, 1);
INSERT INTO banner_schema.banner_tag VALUES(1, 2);

INSERT INTO banner_schema.banners VALUES(2, false, '{"text": "warning_banner"}');
INSERT INTO banner_schema.tags VALUES (3);
INSERT INTO banner_schema.banner_feature VALUES(2, 1);
INSERT INTO banner_schema.banner_tag VALUES(2, 3);