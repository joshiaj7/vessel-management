CREATE TABLE vessels (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `owner_id` BIGINT,
    `name` VARCHAR(255),
    `naccs_code` VARCHAR(20) UNIQUE,
    `created_at` DATETIME,
    `updated_at` DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX index_owner_id ON vessels (owner_id, type);
