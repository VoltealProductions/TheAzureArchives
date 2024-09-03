CREATE TABLE IF NOT EXISTS characters (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_id` INT NOT NULL,
    `unique_id` VARCHAR(25) UNIQUE NOT NULL,
    `firstname` VARCHAR(255) NOT NULL,
    `lastname` VARCHAR(255) NOT NULL,
    `faction` VARCHAR(255) NOT NULL,
    `class` VARCHAR(255) NOT NULL,
    `species` VARCHAR(255) NOT NULL,
    `short_title` VARCHAR(255) DEFAULT '',
    `full_title` VARCHAR(255) DEFAULT '',
    `age` INT DEFAULT 0,
    `gender` VARCHAR(255) DEFAULT '',
    `pronouns` VARCHAR(255) DEFAULT '',
    `height` FLOAT DEFAULT 0.0,
    `weight` FLOAT DEFAULT 0.0,
    `birthplace` VARCHAR(255) DEFAULT '',
    `residence` VARCHAR(255) DEFAULT '',
    `about` TEXT,
    `history` TEXT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);