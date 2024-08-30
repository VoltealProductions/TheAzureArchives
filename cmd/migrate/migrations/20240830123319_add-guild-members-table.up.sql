CREATE TABLE IF NOT EXISTS guildmembers (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `guild_slug` VARCHAR(255) NOT NULL,
    `character_id` INT UNSIGNED NOT NULL,
    `Rank` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT NOW(),
    `updated_at` TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (id),
    UNIQUE KEY (character_id),
    FOREIGN KEY (guild_slug) REFERENCES guilds (slug) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (character_id) REFERENCES characters (id) ON DELETE CASCADE ON UPDATE CASCADE
);