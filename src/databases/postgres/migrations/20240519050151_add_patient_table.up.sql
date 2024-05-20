CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    identity_number BIGINT NOT NULL UNIQUE,
    phone_number VARCHAR(15) NOT NULL,
    name VARCHAR(30) NOT NULL,
    birth_date VARCHAR(255) NOT NULL,
    gender VARCHAR(6) NOT NULL CHECK(gender IN('male', 'female')),
    identity_card_scan_image VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)