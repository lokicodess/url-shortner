ALTER TABLE `url_table`
ADD CONSTRAINT `chk_actual_url_format`
    CHECK (actual_url LIKE 'http%' AND CHAR_LENGTH(actual_url) > 5),
ADD CONSTRAINT `chk_short_code_format`
    CHECK (short_code REGEXP '^[A-Za-z0-9]+$' OR short_code IS NULL);

