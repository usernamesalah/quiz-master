CREATE TABLE IF NOT EXISTS questions (
    id INT(12) PRIMARY KEY AUTO_INCREMENT,
    question VARCHAR(255) NOT NULL,
    answer VARCHAR(50) NOT NULL,
    created_at INT(12),
    updated_at INT(12),
    deleted_at INT(12),
    INDEX deleted_at_idx (deleted_at)
);
