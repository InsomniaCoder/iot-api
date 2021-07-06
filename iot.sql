CREATE DATABASE  IF NOT EXISTS `iot-db` 
USE `iot-db`;

CREATE TABLE IF NOT EXISTS `sensors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `soil_moisture` FLOAT NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);