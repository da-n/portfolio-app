-- MySQL dump 10.13  Distrib 5.6.27, for Linux (x86_64)
--
-- Host: localhost    Database: portfolio_app
-- ------------------------------------------------------
-- Server version	5.6.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customers`
--
CREATE DATABASE portfolio_app;
USE portfolio_app;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `first_name` varchar(191) NOT NULL,
    `last_name` varchar(191) NOT NULL,
    `email` varchar(191) NOT NULL,
    `password` varchar(191) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `customers` WRITE;

/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES
(1,'Theia','Parker','theia@example.com','password123','2021-05-15 10:05:15'),
(2,'Kaine','Berger','kaine@example.com','password123','2021-05-16 11:45:17'),
(3,'Rami','Mejia','rami@example.com','password123','2021-06-05 08:04:25');

UNLOCK TABLES;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `account_type` varchar(10) NOT NULL DEFAULT 'portfolio',
    `balance` bigint(20) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `accounts_FK` (`customer_id`),
    CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `accounts` WRITE;

/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES
(1,1,'portfolio',20000000,'2021-05-15 10:15:25'),
(2,2,'portfolio',10000000,'2021-05-16 11:55:27'),
(3,2,'portfolio',3000000,'2021-05-18 12:35:02'),
(4,3,'portfolio',32000000,'2021-06-05 08:22:35');

UNLOCK TABLES;

DROP TABLE IF EXISTS `withdrawal_requests`;
CREATE TABLE `withdrawal_requests` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL,
    `amount` bigint(20) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `withdrawal_requests_FK` (`account_id`),
    CONSTRAINT `withdrawal_requests_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `withdrawal_requests` WRITE;

/*!40000 ALTER TABLE `withdrawal_requests` DISABLE KEYS */;
INSERT INTO `withdrawal_requests` VALUES
(1,1,2000000,'2021-05-18 12:05:15'),
(2,2,300000,'2021-05-19 11:25:35'),
(3,2,1400000,'2021-05-21 11:25:35');

UNLOCK TABLES;

DROP TABLE IF EXISTS `order_sheets`;
CREATE TABLE `order_sheets` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL,
    `withdrawal_request_id` int(11),
    `status` varchar(10) NOT NULL DEFAULT 'pending',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `order_sheets_accounts_FK` (`account_id`),
    CONSTRAINT `order_sheets_accounts_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE,
    KEY `order_sheets_withdrawal_requests_FK` (`withdrawal_request_id`),
    CONSTRAINT `order_sheets_withdrawal_requests_FK` FOREIGN KEY (`withdrawal_request_id`) REFERENCES `withdrawal_requests` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `order_sheets` WRITE;

/*!40000 ALTER TABLE `order_sheets` DISABLE KEYS */;
INSERT INTO `order_sheets` VALUES
(1,1,1,'complete','2021-05-18 12:05:18'),
(2,2,2,'complete','2021-05-19 11:25:37'),
(3,2,3,'complete','2021-05-21 11:25:36');

UNLOCK TABLES;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-31 10:25:19