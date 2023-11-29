
CREATE TABLE cities (
    guid VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255),
    country_id VARCHAR(36) REFERENCES countries(guid) ON DELETE CASCADE,
    city_code VARCHAR(255),
    latitude VARCHAR(255),
    longitude VARCHAR(255),
    "offset" VARCHAR(255),
    timezone_id VARCHAR(36),
    country_name VARCHAR(255)
);