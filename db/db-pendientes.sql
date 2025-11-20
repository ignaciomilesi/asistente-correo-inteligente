-- MySQL dump 10.13  Distrib 8.0.44, for Win64 (x86_64)
--
-- Host: localhost    Database: pendientes
-- ------------------------------------------------------
-- Server version	8.0.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `adjuntos`
--

DROP TABLE IF EXISTS `adjuntos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `adjuntos` (
  `id` int NOT NULL,
  `Descripcion` varchar(255) NOT NULL,
  `ubicacion_archivo` varchar(255) DEFAULT NULL,
  `Pediente_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Avances_Pediente_idx` (`Pediente_id`),
  CONSTRAINT `fk_Avances_Pediente0` FOREIGN KEY (`Pediente_id`) REFERENCES `pediente` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `adjuntos`
--

LOCK TABLES `adjuntos` WRITE;
/*!40000 ALTER TABLE `adjuntos` DISABLE KEYS */;
/*!40000 ALTER TABLE `adjuntos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `avances`
--

DROP TABLE IF EXISTS `avances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `avances` (
  `id` int NOT NULL,
  `Fecha_Avance` varchar(255) NOT NULL,
  `Descripcion` varchar(255) NOT NULL,
  `ubicacion_mail` varchar(255) DEFAULT NULL,
  `Pediente_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Avances_Pediente_idx` (`Pediente_id`),
  CONSTRAINT `fk_Avances_Pediente` FOREIGN KEY (`Pediente_id`) REFERENCES `pediente` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `avances`
--

LOCK TABLES `avances` WRITE;
/*!40000 ALTER TABLE `avances` DISABLE KEYS */;
/*!40000 ALTER TABLE `avances` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `etiquetas`
--

DROP TABLE IF EXISTS `etiquetas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `etiquetas` (
  `id` int NOT NULL,
  `nombre` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `etiquetas`
--

LOCK TABLES `etiquetas` WRITE;
/*!40000 ALTER TABLE `etiquetas` DISABLE KEYS */;
/*!40000 ALTER TABLE `etiquetas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pediente`
--

DROP TABLE IF EXISTS `pediente`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pediente` (
  `id` int NOT NULL,
  `Descripcion` text NOT NULL,
  `Finalizada` tinyint NOT NULL DEFAULT '0',
  `Fecha_iniciada` date NOT NULL,
  `Fecha_finalizada` date DEFAULT NULL,
  `Cierre` varchar(255) DEFAULT NULL,
  `asignado` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Pediente_usuarios1_idx` (`asignado`),
  CONSTRAINT `fk_Pediente_usuarios1` FOREIGN KEY (`asignado`) REFERENCES `usuarios` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pediente`
--

LOCK TABLES `pediente` WRITE;
/*!40000 ALTER TABLE `pediente` DISABLE KEYS */;
/*!40000 ALTER TABLE `pediente` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pediente_has_etiquetas`
--

DROP TABLE IF EXISTS `pediente_has_etiquetas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pediente_has_etiquetas` (
  `Pediente_id` int NOT NULL,
  `Etiquetas_id` int NOT NULL,
  PRIMARY KEY (`Pediente_id`,`Etiquetas_id`),
  KEY `fk_Pediente_has_Etiquetas_Etiquetas1_idx` (`Etiquetas_id`),
  KEY `fk_Pediente_has_Etiquetas_Pediente1_idx` (`Pediente_id`),
  CONSTRAINT `fk_Pediente_has_Etiquetas_Etiquetas1` FOREIGN KEY (`Etiquetas_id`) REFERENCES `etiquetas` (`id`),
  CONSTRAINT `fk_Pediente_has_Etiquetas_Pediente1` FOREIGN KEY (`Pediente_id`) REFERENCES `pediente` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pediente_has_etiquetas`
--

LOCK TABLES `pediente_has_etiquetas` WRITE;
/*!40000 ALTER TABLE `pediente_has_etiquetas` DISABLE KEYS */;
/*!40000 ALTER TABLE `pediente_has_etiquetas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pediente_has_sap`
--

DROP TABLE IF EXISTS `pediente_has_sap`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pediente_has_sap` (
  `Pediente_id` int NOT NULL,
  `Sap_id` int NOT NULL,
  PRIMARY KEY (`Pediente_id`,`Sap_id`),
  KEY `fk_Pediente_has_Sap_Sap1_idx` (`Sap_id`),
  KEY `fk_Pediente_has_Sap_Pediente1_idx` (`Pediente_id`),
  CONSTRAINT `fk_Pediente_has_Sap_Pediente1` FOREIGN KEY (`Pediente_id`) REFERENCES `pediente` (`id`),
  CONSTRAINT `fk_Pediente_has_Sap_Sap1` FOREIGN KEY (`Sap_id`) REFERENCES `sap` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pediente_has_sap`
--

LOCK TABLES `pediente_has_sap` WRITE;
/*!40000 ALTER TABLE `pediente_has_sap` DISABLE KEYS */;
/*!40000 ALTER TABLE `pediente_has_sap` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sap`
--

DROP TABLE IF EXISTS `sap`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sap` (
  `id` int NOT NULL,
  `SAP` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sap`
--

LOCK TABLES `sap` WRITE;
/*!40000 ALTER TABLE `sap` DISABLE KEYS */;
/*!40000 ALTER TABLE `sap` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuarios` (
  `id` int NOT NULL,
  `nombre` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'pendientes'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-11-07 11:15:14
