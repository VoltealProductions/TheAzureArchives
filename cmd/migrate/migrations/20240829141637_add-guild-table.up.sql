CREATE TABLE IF NOT EXISTS guilds (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `owner_id` INT UNSIGNED NOT NULL,
    `slug` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `faction` VARCHAR(255) NOT NULL,
    `realm` VARCHAR(255) NOT NULL,
    `ranks` VARCHAR(255) default '',
    `recruiting` BOOLEAN NOT NULL DEFAULT false,
    `description` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT NOW(),
    `updated_at` TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (id),
    UNIQUE KEY (slug),
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);