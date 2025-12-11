CREATE TABLE IF NOT EXISTS `url_table` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `short_code` VARCHAR(30) DEFAULT NULL,
  `actual_url` VARCHAR(500) NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `expires_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `short_code` (`short_code`)
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_0900_ai_ci;

