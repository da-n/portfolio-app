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

DROP TABLE IF EXISTS `portfolios`;
CREATE TABLE `portfolios` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(191) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `portfolios` WRITE;

/*!40000 ALTER TABLE `portfolios` DISABLE KEYS */;
INSERT INTO `portfolios` VALUES
(1,'Growth Portfolio');

UNLOCK TABLES;

DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `isin` varchar(12) NOT NULL,
    `name` varchar(191) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `assets` WRITE;

/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES
(1,'IE00B52L4369','BlackRock Institutional Cash Series Sterling Liquidity Agency Inc'),
(2,'GB00BQ1YHQ70','Threadneedle UK Property Authorised Investment Net GBP 1 Acc'),
(3,'GB00B3X7QG63','Vanguard FTSE U.K. All Share Index Unit Trust Accumulation'),
(4,'GB00BG0QP828','Legal & General Japan Index Trust C Class Accumulation'),
(5,'GB00BPN5P238','Vanguard US Equity Index Institutional Plus GBP Accumulation'),
(6,'IE00B1S74Q32','Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation');

UNLOCK TABLES;

DROP TABLE IF EXISTS `asset_portfolio`;
CREATE TABLE `asset_portfolio` (
    `asset_id` int(11) NOT NULL,
    `portfolio_id` int(11) NOT NULL,
    `percent` int(3) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `asset_portfolio` WRITE;

/*!40000 ALTER TABLE `asset_portfolio` DISABLE KEYS */;
INSERT INTO `asset_portfolio` VALUES
(1,1,20),
(2,1,20),
(3,1,10),
(4,1,5),
(5,1,15),
(6,1,30);

UNLOCK TABLES;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `portfolio_id` int(11) NOT NULL,
    `balance` bigint(20) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `accounts_customers_FK` (`customer_id`),
    CONSTRAINT `accounts_customers_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE,
    KEY `accounts_portfolios_FK` (`portfolio_id`),
    CONSTRAINT `accounts_portfolios_FK` FOREIGN KEY (`portfolio_id`) REFERENCES `portfolios` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `accounts` WRITE;

/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES
(1,1,1,20000000,'2021-05-15 10:15:25'),
(2,2,1,10000000,'2021-05-16 11:55:27'),
(3,2,1,3000000,'2021-05-18 12:35:02'),
(4,3,1,32000000,'2021-06-05 08:22:35');

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