# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.62-0+deb8u1)
# Database: user_balance
# Generation Time: 2022-10-24 14:39:42 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table accounting
# ------------------------------------------------------------

DROP TABLE IF EXISTS `accounting`;

CREATE TABLE `accounting` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `month` tinyint(2) unsigned NOT NULL,
  `year` mediumint(4) unsigned NOT NULL,
  `service_id` int(10) unsigned NOT NULL,
  `service_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `money` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni` (`month`,`year`,`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `accounting` WRITE;
/*!40000 ALTER TABLE `accounting` DISABLE KEYS */;

INSERT INTO `accounting` (`id`, `month`, `year`, `service_id`, `service_name`, `money`)
VALUES
	(1,10,2022,2,'турбопродажа',100000),
	(3,10,2022,3,'размещение объявления',40000);

/*!40000 ALTER TABLE `accounting` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table balance
# ------------------------------------------------------------

DROP TABLE IF EXISTS `balance`;

CREATE TABLE `balance` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `money` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `balance` WRITE;
/*!40000 ALTER TABLE `balance` DISABLE KEYS */;

INSERT INTO `balance` (`id`, `user_id`, `money`)
VALUES
	(1,1,10000),
	(9,2,0),
	(15,3,580000),
	(17,4,50000),
	(21,5,25000);

/*!40000 ALTER TABLE `balance` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table transactions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `service_id` int(10) unsigned NOT NULL,
  `service_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `order_id` int(10) unsigned NOT NULL,
  `type` tinyint(1) NOT NULL DEFAULT '0',
  `money` int(10) unsigned NOT NULL,
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `confirmed` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;

INSERT INTO `transactions` (`id`, `user_id`, `service_id`, `service_name`, `order_id`, `type`, `money`, `status`, `confirmed`, `created_at`)
VALUES
	(1,2,7,'Перевод от пользователя',124522,1,10000,1,1666182560,1666182550),
	(2,2,1,'Пополнение банковской картой',1245124,1,25000,1,1666182899,1666182874),
	(4,2,1,'Пополнение банковской картой',1245125,1,25000,1,1666183345,1666183088),
	(5,1,1,'Пополнение банковской картой',1245126,1,25000,1,1666205295,1666205295),
	(7,2,2,'турбопродажа',1245127,0,20000,1,1666205530,1666205507),
	(8,1,2,'турбопродажа',1245128,0,20000,1,1666205699,1666205619),
	(9,1,1,'Пополнение банковской картой',1245136,1,45000,1,1666207381,1666207381),
	(10,1,2,'турбопродажа',1245140,0,20000,2,1666207765,1666207636),
	(11,3,1,'Пополнение банковской картой',1245142,1,600000,1,1666209387,1666209387),
	(12,3,2,'турбопродажа',1245145,0,20000,1,1666281703,1666209414),
	(15,4,1,'Пополнение банковской картой',1245170,1,25000,1,1666279095,1666279095),
	(18,4,1,'Пополнение банковской картой',1245188,1,25000,1,1666281774,1666281774),
	(19,4,2,'турбопродажа',1245189,0,20000,2,1666283461,1666283385),
	(20,5,1,'Пополнение банковской картой',1245190,1,25000,1,1666286200,1666286200),
	(21,2,2,'турбопродажа',1245226,0,20000,1,1666363445,1666362680),
	(22,2,2,'турбопродажа',1245227,0,20000,1,1666363468,1666362740),
	(23,1,2,'турбопродажа',1245326,0,20000,2,1666365517,1666365086),
	(25,1,3,'размещение объявления',1245329,0,40000,1,1666365713,1666365636);

/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
