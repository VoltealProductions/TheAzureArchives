CREATE TABLE IF NOT EXISTS users (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `username` VARCHAR(255) UNIQUE NOT NULL,
    `password` VARCHAR(75) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `public` BOOLEAN NOT NULL DEFAULT false,
    `banned` BOOLEAN NOT NULL DEFAULT false,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);