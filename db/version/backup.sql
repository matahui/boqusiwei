-- MySQL dump 10.13  Distrib 8.0.39, for Win64 (x86_64)
--
-- Host: localhost    Database: homeschooledu
-- ------------------------------------------------------
-- Server version	8.0.39

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `cate` tinyint NOT NULL COMMENT '账号类型',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'admin','12345678',1,'2024-08-17 13:05:09','2024-08-17 13:05:51',0),(2,'zhangsan','12345678',2,'2024-08-19 12:43:22','2024-08-19 12:43:22',0);
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `activity_logs`
--

DROP TABLE IF EXISTS `activity_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `activity_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `resource_id` bigint NOT NULL,
  `activity_date` date NOT NULL,
  `points_award` int NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_student_resource_date` (`student_id`,`resource_id`,`activity_date`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity_logs`
--

LOCK TABLES `activity_logs` WRITE;
/*!40000 ALTER TABLE `activity_logs` DISABLE KEYS */;
INSERT INTO `activity_logs` VALUES (1,6,1,'2024-08-23',10,'2024-08-23 11:27:05','2024-08-23 11:27:05',0),(2,1000001,1,'2024-08-23',10,'2024-08-23 11:28:34','2024-08-23 11:28:34',0),(3,1000001,2,'2024-08-23',10,'2024-08-23 11:28:40','2024-08-23 11:28:40',0),(4,1000001,3,'2024-08-23',10,'2024-08-23 11:28:46','2024-08-23 11:28:46',0),(5,1000002,3,'2024-08-23',10,'2024-08-23 11:28:51','2024-08-23 11:28:51',0),(6,1000003,2,'2024-08-23',10,'2024-08-23 11:28:57','2024-08-23 11:28:57',0),(7,1000007,2,'2024-08-23',10,'2024-08-23 17:28:50','2024-08-23 17:28:50',0);
/*!40000 ALTER TABLE `activity_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `classes`
--

DROP TABLE IF EXISTS `classes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `classes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_name` varchar(255) NOT NULL COMMENT '班级名称',
  `school_id` bigint NOT NULL COMMENT '所属学校ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classes`
--

LOCK TABLES `classes` WRITE;
/*!40000 ALTER TABLE `classes` DISABLE KEYS */;
INSERT INTO `classes` VALUES (1,'班级01',1,'2024-08-19 15:57:54','2024-08-19 15:57:54',0),(2,'班级02',1,'2024-08-19 15:58:03','2024-08-19 15:58:03',0),(3,'小班3',2,'2024-08-19 15:58:24','2024-08-23 16:31:03',0),(4,'班级04',1,'2024-08-19 15:58:30','2024-08-19 15:58:30',0),(5,'班级05',1,'2024-08-19 15:58:37','2024-08-19 15:58:37',0),(6,'班级06',1,'2024-08-19 15:58:53','2024-08-19 15:58:53',0),(7,'班级01',2,'2024-08-19 15:59:04','2024-08-19 15:59:04',0),(8,'班级02',2,'2024-08-19 15:59:08','2024-08-19 15:59:08',0),(9,'班级03',2,'2024-08-19 15:59:17','2024-08-23 16:31:25',1),(10,'',1,'2024-08-23 16:33:55','2024-08-23 16:33:55',0);
/*!40000 ALTER TABLE `classes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `regions`
--

DROP TABLE IF EXISTS `regions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `regions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `regions`
--

LOCK TABLES `regions` WRITE;
/*!40000 ALTER TABLE `regions` DISABLE KEYS */;
INSERT INTO `regions` VALUES (1,'北京','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(2,'上海','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(3,'广州','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(4,'深圳','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(5,'成都','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(6,'武汉','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(7,'南京','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(8,'杭州','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(9,'苏州','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(10,'西安','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(11,'重庆','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(12,'天津','2024-08-18 13:50:41','2024-08-18 13:50:41',0),(13,'澳大利亚袋鼠小学','2024-08-18 14:00:41','2024-08-18 14:00:41',0),(14,'','2024-08-23 15:16:52','2024-08-23 15:16:52',0);
/*!40000 ALTER TABLE `regions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `resources`
--

DROP TABLE IF EXISTS `resources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `resources` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `resource_name` varchar(255) NOT NULL,
  `age_group` varchar(50) NOT NULL,
  `course` varchar(50) NOT NULL,
  `level_1` varchar(50) NOT NULL,
  `level_2` varchar(50) NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resources`
--

LOCK TABLES `resources` WRITE;
/*!40000 ALTER TABLE `resources` DISABLE KEYS */;
INSERT INTO `resources` VALUES (1,'资源1','小班','小龙人','阅读','闪卡','2024-08-20 15:30:50','2024-08-20 15:30:50',0),(2,'资源2','中班','小龙人','写作','作文','2024-08-20 15:30:50','2024-08-20 15:30:50',0),(3,'资源3','大班','小龙人','听说','一练','2024-08-20 15:30:50','2024-08-23 18:12:34',1),(4,'资源1','小班','小龙人','阅读','闪卡','2024-08-23 17:02:16','2024-08-23 17:02:16',0),(5,'资源2','中班','小龙人','写作','作文','2024-08-23 17:02:16','2024-08-23 17:02:16',0),(6,'资源3','大班','小龙人','听说','一练','2024-08-23 17:02:16','2024-08-23 17:02:16',0);
/*!40000 ALTER TABLE `resources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedules`
--

DROP TABLE IF EXISTS `schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedules` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `resource_id` bigint NOT NULL,
  `school_id` bigint NOT NULL,
  `class_id` bigint NOT NULL,
  `begin_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedules`
--

LOCK TABLES `schedules` WRITE;
/*!40000 ALTER TABLE `schedules` DISABLE KEYS */;
INSERT INTO `schedules` VALUES (4,1,1,4,'2024-08-20 23:04:05','2024-09-07 23:04:05','2024-08-21 12:25:39','2024-08-21 12:25:39',0),(5,1,1,5,'2024-08-20 23:04:05','2024-09-07 23:04:05','2024-08-21 12:25:39','2024-08-21 12:25:39',0),(6,1,1,6,'2024-08-20 23:04:05','2024-09-07 23:04:05','2024-08-21 12:25:39','2024-08-21 12:25:39',0),(7,3,1,4,'2024-08-20 23:04:05','2024-09-07 23:04:05','2024-08-23 17:20:45','2024-08-23 17:20:45',0),(8,3,1,5,'2024-08-20 23:04:05','2024-09-07 23:04:05','2024-08-23 17:20:45','2024-08-23 17:20:45',0),(9,3,1,6,'2024-08-20 23:04:05','2024-09-07 23:04:05','2024-08-23 17:20:45','2024-08-23 17:20:45',0);
/*!40000 ALTER TABLE `schedules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schools`
--

DROP TABLE IF EXISTS `schools`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schools` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `region` varchar(255) NOT NULL,
  `account` varchar(255) NOT NULL COMMENT '园长',
  `custom_id` varchar(255) DEFAULT NULL COMMENT '自定义id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_region_name` (`region`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schools`
--

LOCK TABLES `schools` WRITE;
/*!40000 ALTER TABLE `schools` DISABLE KEYS */;
INSERT INTO `schools` VALUES (1,'深圳北大附中','深圳','zhangsan','CUS1001','2024-08-18 13:50:46','2024-08-23 15:10:42',1),(2,'清华附中','北京','lisi','CUS1002','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(3,'上海中学','上海','wangwu','CUS1003','2024-08-18 13:50:46','2024-08-18 16:13:35',0),(4,'复旦附中','上海','zhaoliu','CUS1004','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(5,'广州实验中学','广州','sunqi','CUS1005','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(6,'深圳中学','深圳','zhouba','CUS1006','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(7,'成都七中','成都','wuma','CUS1007','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(8,'武汉外国语学校','武汉','zhengjiu','CUS1008','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(9,'南京外国语学校','南京','xushiyi','CUS1009','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(10,'杭州外国语学校','杭州','qianshisan','CUS1010','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(11,'苏州中学','苏州','zhaoshiwu','CUS1011','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(12,'西安中学','西安','fengshiliu','CUS1012','2024-08-18 13:50:46','2024-08-18 13:50:46',0),(13,'深圳市袋鼠小学','深圳','admin','','2024-08-18 14:01:13','2024-08-18 14:01:13',0),(14,'深圳市饿了么小学','深圳','admin','','2024-08-23 15:12:25','2024-08-23 15:12:25',0);
/*!40000 ALTER TABLE `schools` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student_points`
--

DROP TABLE IF EXISTS `student_points`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `student_points` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL COMMENT '学生id',
  `points` bigint NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_student` (`student_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student_points`
--

LOCK TABLES `student_points` WRITE;
/*!40000 ALTER TABLE `student_points` DISABLE KEYS */;
INSERT INTO `student_points` VALUES (1,6,10,'2024-08-23 11:27:05','2024-08-23 11:27:05',0),(2,1000001,30,'2024-08-23 11:28:34','2024-08-23 11:28:46',0),(3,1000002,10,'2024-08-23 11:28:51','2024-08-23 11:28:51',0),(4,1000003,10,'2024-08-23 11:28:57','2024-08-23 11:28:57',0),(5,1000007,10,'2024-08-23 17:28:50','2024-08-23 17:28:50',0);
/*!40000 ALTER TABLE `student_points` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `students` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `login_number` bigint NOT NULL COMMENT '登录账号',
  `student_name` varchar(255) NOT NULL COMMENT '学生姓名',
  `parent_name` varchar(255) DEFAULT NULL COMMENT '家长姓名',
  `phone_number` varchar(20) DEFAULT NULL COMMENT '电话号码',
  `class_id` bigint NOT NULL COMMENT '所属班级ID',
  `school_id` bigint NOT NULL COMMENT '所属学校ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  `password` varchar(255) NOT NULL DEFAULT '123456' COMMENT '登录密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_login_number` (`login_number`)
) ENGINE=InnoDB AUTO_INCREMENT=1000016 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students`
--

LOCK TABLES `students` WRITE;
/*!40000 ALTER TABLE `students` DISABLE KEYS */;
INSERT INTO `students` VALUES (1000001,19900008666,'詹姆斯','乔治亚','13489091234',11,2,'2024-08-19 12:14:05','2024-08-23 15:33:55',0,'123456'),(1000002,22222333333,'张小斐','张等等','888888888',123,1,'2024-08-19 12:14:05','2024-08-19 12:14:05',0,'123456'),(1000003,33333444444,'王晓红','王九九','999999999',123,1,'2024-08-19 12:14:05','2024-08-23 15:35:38',1,'123456'),(1000005,666777888,'詹姆斯','金正恩','123443211',7,1,'2024-08-20 13:56:42','2024-08-20 13:56:42',0,'123456'),(1000006,999000111,'杜兰特','习近平','888888888',7,1,'2024-08-20 13:56:42','2024-08-20 13:56:42',0,'123456'),(1000007,222555888,'特胖铺','邓光荣','999999999',6,1,'2024-08-20 13:56:42','2024-08-22 19:53:17',0,'123456'),(1000008,18811224567,'孙悟空','菩提祖师','13652348890',7,1,'2024-08-23 15:47:44','2024-08-23 15:47:44',0,''),(1000009,444444444,'面糊糊','北大荒','123443211',7,1,'2024-08-23 16:03:48','2024-08-23 16:03:48',0,''),(1000010,555555555,'刀削面','信天游','888888888',7,1,'2024-08-23 16:03:48','2024-08-23 16:03:48',0,''),(1000011,999999999,'牛肉面','老黄河','999999999',7,1,'2024-08-23 16:03:48','2024-08-23 16:03:48',0,''),(1000015,666666666,'猪八戒','猪刚鬣','123443211',7,1,'2024-08-23 18:04:45','2024-08-23 18:04:45',0,'');
/*!40000 ALTER TABLE `students` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teacher_class_assignments`
--

DROP TABLE IF EXISTS `teacher_class_assignments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teacher_class_assignments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `teacher_id` bigint NOT NULL,
  `class_id` bigint NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teacher_class_assignments`
--

LOCK TABLES `teacher_class_assignments` WRITE;
/*!40000 ALTER TABLE `teacher_class_assignments` DISABLE KEYS */;
INSERT INTO `teacher_class_assignments` VALUES (1,3,7,'2024-08-19 17:42:24','2024-08-20 10:43:12',1),(2,3,8,'2024-08-19 17:42:24','2024-08-19 18:03:50',1),(3,3,9,'2024-08-19 17:42:24','2024-08-19 18:03:50',1),(4,3,10,'2024-08-19 17:54:29','2024-08-19 18:03:50',1),(5,3,11,'2024-08-19 18:03:50','2024-08-20 10:20:09',1),(6,3,4,'2024-08-20 10:20:09','2024-08-20 10:43:12',1),(10,1,7,'2024-08-20 13:41:35','2024-08-20 13:41:35',0),(11,2,7,'2024-08-20 13:41:35','2024-08-20 13:41:35',0),(12,3,7,'2024-08-20 13:41:35','2024-08-23 18:11:20',1),(13,3,4,'2024-08-23 16:15:13','2024-08-23 18:11:20',1),(14,4,3,'2024-08-23 16:19:54','2024-08-23 16:19:54',0),(15,4,4,'2024-08-23 16:19:54','2024-08-23 16:19:54',0),(16,1,7,'2024-08-23 16:50:44','2024-08-23 16:50:44',0),(17,2,7,'2024-08-23 16:50:44','2024-08-23 16:50:44',0);
/*!40000 ALTER TABLE `teacher_class_assignments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teachers`
--

DROP TABLE IF EXISTS `teachers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teachers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `login_number` bigint NOT NULL COMMENT '登录账号',
  `teacher_name` varchar(50) NOT NULL COMMENT '老师姓名',
  `phone_number` varchar(20) DEFAULT NULL COMMENT '电话号码',
  `role` tinyint NOT NULL COMMENT '角色,园长1 教师2',
  `school_id` bigint NOT NULL COMMENT '所属学校',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_delete` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `u_login_number` (`login_number`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teachers`
--

LOCK TABLES `teachers` WRITE;
/*!40000 ALTER TABLE `teachers` DISABLE KEYS */;
INSERT INTO `teachers` VALUES (1,333000111,'王老师','110111112',1,1,'2024-08-19 15:09:20','2024-08-19 15:09:20',0),(2,444000222,'赵老师','220222333',1,1,'2024-08-19 15:09:20','2024-08-19 15:09:20',0),(3,4331111110000,'六老师','13489091234',1,1,'2024-08-19 15:09:20','2024-08-23 18:11:20',1),(4,333666999,'张钰琪','13985247862',2,1,'2024-08-23 16:19:54','2024-08-23 16:19:54',0),(5,999999999,'张老师','110111112',1,1,'2024-08-23 16:23:57','2024-08-23 16:23:57',0),(6,566666666,'杨老师','330333444',1,1,'2024-08-23 16:23:57','2024-08-23 16:23:57',0);
/*!40000 ALTER TABLE `teachers` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-24 17:26:13
