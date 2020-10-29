
-- +migrate Up
CREATE TABLE IF NOT EXISTS `messages` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `body` TEXT,
  `created_at` DATETIME default current_timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS `messages`;