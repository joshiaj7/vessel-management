CREATE TABLE shipments (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    vessel_id BIGINT,
    source VARCHAR(255),
    destination VARCHAR(255),
    current_location VARCHAR(255),
    state TINYINT,
    docked_at DATETIME,
    departed_at DATETIME,
    transited_at DATETIME,
    arrived_at DATETIME,
    created_at DATETIME,
    updated_at DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX index_vessel_id ON shipments (vessel_id, type);
