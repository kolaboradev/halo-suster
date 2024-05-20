CREATE TABLE IF NOT EXISTS records (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    identity_number BIGINT NOT NULL,
    symptoms TEXT NOT NULL,
    medications TEXT NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (identity_number) REFERENCES patients(identity_number) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
)