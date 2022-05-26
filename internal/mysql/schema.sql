CREATE DATABASE IF NOT EXISTS `visitor`;
USE `visitor`;

CREATE TABLE IF NOT EXISTS `visitor` (
	`id` int unsigned NOT NULL AUTO_INCREMENT,
	`nickname` varchar(30) NOT NULL,
	`visit_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE (`nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE IF NOT EXISTS `theme` (
	`id` int unsigned NOT NULL DEFAULT 0,
	`color` varchar(30) NOT NULL DEFAULT 'green'
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

INSERT INTO `theme` (`id`)
SELECT 0
WHERE NOT EXISTS (SELECT * FROM `theme`);
