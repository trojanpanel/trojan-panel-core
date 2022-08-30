-- MySQL dump 10.13  Distrib 5.7.29, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: trojan_panel_db
-- ------------------------------------------------------
-- Server version	5.7.29

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
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '登录用户名',
  `pass` varchar(64) NOT NULL DEFAULT '' COMMENT '登录密码',
  `role_id` bigint(20) unsigned NOT NULL DEFAULT '3' COMMENT '角色id 1/系统管理员 3/普通用户',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `expire_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  `deleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用 0/正常 1/禁用',
  `quota` bigint(20) NOT NULL DEFAULT '0' COMMENT '配额 单位/byte',
  `download` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '下载 单位/byte',
  `upload` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '上传 单位/byte',
  `ip_limit` tinyint(2) unsigned NOT NULL DEFAULT '3' COMMENT '限制IP设备数',
  `upload_speed_limit` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '上传限速 单位/byte',
  `download_speed_limit` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '下载限速 单位/byte',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='账户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES (1,'sysadmin','eWTgjy8ZBnXL/U2lm08xLg==',3,'',4078656000000,0,-1,0,0,3,0,0,'2022-04-01 00:00:00','2022-04-01 00:00:00');
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `black_list`
--

DROP TABLE IF EXISTS `black_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `black_list` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='黑名单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `black_list`
--

LOCK TABLES `black_list` WRITE;
/*!40000 ALTER TABLE `black_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `black_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(32) NOT NULL DEFAULT '',
  `v0` varchar(255) NOT NULL DEFAULT '',
  `v1` varchar(255) NOT NULL DEFAULT '',
  `v2` varchar(255) NOT NULL DEFAULT '',
  `v3` varchar(255) NOT NULL DEFAULT '',
  `v4` varchar(255) NOT NULL DEFAULT '',
  `v5` varchar(255) NOT NULL DEFAULT '',
  KEY `idx_casbin_rule` (`p_type`,`v0`,`v1`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES ('p','sysadmin','/api/account/selectAccountById','GET','','',''),('p','sysadmin','/api/account/createAccount','POST','','',''),('p','sysadmin','/api/account/getAccountInfo','GET','','',''),('p','sysadmin','/api/account/selectAccountPage','GET','','',''),('p','sysadmin','/api/account/deleteAccountById','POST','','',''),('p','sysadmin','/api/account/updateAccountProfile','POST','','',''),('p','sysadmin','/api/account/updateAccountById','POST','','',''),('p','sysadmin','/api/account/logout','POST','','',''),('p','sysadmin','/api/role/selectRoleList','GET','','',''),('p','sysadmin','/api/node/selectNodeById','GET','','',''),('p','sysadmin','/api/node/createNode','POST','','',''),('p','sysadmin','/api/node/selectNodePage','GET','','',''),('p','sysadmin','/api/node/deleteNodeById','POST','','',''),('p','sysadmin','/api/node/updateNodeById','POST','','',''),('p','sysadmin','/api/node/nodeQRCode','POST','','',''),('p','sysadmin','/api/node/nodeURL','POST','','',''),('p','sysadmin','/api/nodeType/selectNodeTypeList','GET','','',''),('p','sysadmin','/api/dashboard/panelGroup','GET','','',''),('p','sysadmin','/api/dashboard/trafficRank','GET','','',''),('p','sysadmin','/api/system/selectSystemByName','GET','','',''),('p','sysadmin','/api/system/updateSystemById','POST','','',''),('p','sysadmin','/api/system/uploadWebFile','POST','','',''),('p','sysadmin','/api/blackList/selectBlackListPage','GET','','',''),('p','sysadmin','/api/blackList/deleteBlackListByIp','POST','','',''),('p','sysadmin','/api/blackList/createBlackList','POST','','',''),('p','sysadmin','/api/emailRecord/selectEmailRecordPage','GET','','',''),('p','user','/api/account/getAccountInfo','GET','','',''),('p','user','/api/account/updateAccountProfile','POST','','',''),('p','user','/api/account/logout','POST','','',''),('p','user','/api/node/selectNodePage','GET','','',''),('p','user','/api/node/nodeQRCode','POST','','',''),('p','user','/api/node/nodeURL','POST','','',''),('p','user','/api/nodeType/selectNodeTypeList','GET','','',''),('p','user','/api/dashboard/panelGroup','GET','','',''),('p','user','/api/dashboard/trafficRank','GET','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `email_record`
--

DROP TABLE IF EXISTS `email_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `email_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `to_email` varchar(64) NOT NULL DEFAULT '' COMMENT '收件人邮箱',
  `subject` varchar(64) NOT NULL DEFAULT '' COMMENT '主题',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '内容',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0/未发送 1/发送成功 -1/发送失败',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邮件发送记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `email_record`
--

LOCK TABLES `email_record` WRITE;
/*!40000 ALTER TABLE `email_record` DISABLE KEYS */;
/*!40000 ALTER TABLE `email_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node`
--

DROP TABLE IF EXISTS `node`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `node_sub_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '节点分表id',
  `node_type_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '节点类型id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
  `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `port` int(10) unsigned NOT NULL DEFAULT '443' COMMENT '端口',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='节点';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node`
--

LOCK TABLES `node` WRITE;
/*!40000 ALTER TABLE `node` DISABLE KEYS */;
/*!40000 ALTER TABLE `node` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node_hysteria`
--

DROP TABLE IF EXISTS `node_hysteria`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node_hysteria` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `protocol` varchar(32) NOT NULL DEFAULT 'udp' COMMENT '协议名称 udp/faketcp',
  `up_mbps` int(3) NOT NULL DEFAULT '100' COMMENT '单客户端最大上传速度 单位:Mbps',
  `down_mbps` int(3) NOT NULL DEFAULT '100' COMMENT '单客户端最大下载速度 单位:Mbps',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Hyeteria节点';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node_hysteria`
--

LOCK TABLES `node_hysteria` WRITE;
/*!40000 ALTER TABLE `node_hysteria` DISABLE KEYS */;
/*!40000 ALTER TABLE `node_hysteria` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node_trojan_go`
--

DROP TABLE IF EXISTS `node_trojan_go`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node_trojan_go` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `sni` varchar(64) NOT NULL DEFAULT '' COMMENT 'sni',
  `mux_enable` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否开启多路复用 0/关闭 1/开启',
  `websocket_enable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否开启websocket 0/否 1/是',
  `websocket_path` varchar(64) NOT NULL DEFAULT 'trojan-panel-websocket-path' COMMENT 'websocket路径',
  `websocket_host` varchar(64) NOT NULL DEFAULT '' COMMENT 'websocket host',
  `ss_enable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否开启ss加密 0/否 1/是',
  `ss_method` varchar(32) NOT NULL DEFAULT 'AES-128-GCM' COMMENT 'ss加密方式',
  `ss_password` varchar(64) NOT NULL DEFAULT '' COMMENT 'ss密码',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='TrojanGO节点';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node_trojan_go`
--

LOCK TABLES `node_trojan_go` WRITE;
/*!40000 ALTER TABLE `node_trojan_go` DISABLE KEYS */;
/*!40000 ALTER TABLE `node_trojan_go` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node_type`
--

DROP TABLE IF EXISTS `node_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node_type` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '名称',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='节点类型';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node_type`
--

LOCK TABLES `node_type` WRITE;
/*!40000 ALTER TABLE `node_type` DISABLE KEYS */;
INSERT INTO `node_type` VALUES (1,'xray','2022-04-01 00:00:00','2022-04-01 00:00:00'),(2,'trojan-go','2022-04-01 00:00:00','2022-04-01 00:00:00'),(3,'hysteria','2022-04-01 00:00:00','2022-04-01 00:00:00');
/*!40000 ALTER TABLE `node_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `node_xray`
--

DROP TABLE IF EXISTS `node_xray`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `node_xray` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `protocol` varchar(32) NOT NULL DEFAULT '' COMMENT '协议名称',
  `settings` varchar(256) NOT NULL DEFAULT '' COMMENT 'settings',
  `stream_settings` varchar(256) NOT NULL DEFAULT '' COMMENT 'streamSettings',
  `tag` varchar(64) NOT NULL DEFAULT '' COMMENT 'tag',
  `sniffing` varchar(256) NOT NULL DEFAULT '' COMMENT 'sniffing',
  `allocate` varchar(256) NOT NULL DEFAULT '' COMMENT 'allocate',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Xray节点';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `node_xray`
--

LOCK TABLES `node_xray` WRITE;
/*!40000 ALTER TABLE `node_xray` DISABLE KEYS */;
/*!40000 ALTER TABLE `node_xray` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT '名称',
  `desc` varchar(16) NOT NULL DEFAULT '' COMMENT '描述',
  `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `path` varchar(128) NOT NULL DEFAULT '' COMMENT '路径',
  `level` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '等级',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `role_name_index` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'sysadmin','系统管理员',0,'',1,'2022-04-01 00:00:00','2022-04-01 00:00:00'),(2,'admin','管理员',1,'1-',2,'2022-04-01 00:00:00','2022-04-01 00:00:00'),(3,'user','普通用户',2,'1-2-',3,'2022-04-01 00:00:00','2022-04-01 00:00:00');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system`
--

DROP TABLE IF EXISTS `system`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT '系统名称',
  `open_register` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '开放注册 0/否 1/是',
  `register_quota` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '新默认流量 单位/MB',
  `register_expire_days` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '新用户默认过期天数 单位/天',
  `expire_warn_enable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否开启到期警告 0/否 1/是',
  `expire_warn_day` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '到期警告 单位/天',
  `email_enable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否开启邮箱功能 0/否 1/是',
  `email_host` varchar(64) NOT NULL DEFAULT '' COMMENT '系统邮箱设置-host',
  `email_port` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '系统邮箱设置-port',
  `email_username` varchar(32) NOT NULL DEFAULT '' COMMENT '系统邮箱设置-username',
  `email_password` varchar(32) NOT NULL DEFAULT '' COMMENT '系统邮箱设置-password',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统设置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system`
--

LOCK TABLES `system` WRITE;
/*!40000 ALTER TABLE `system` DISABLE KEYS */;
INSERT INTO `system` VALUES (1,'trojan-panel',1,0,0,0,0,0,'',0,'','','2022-04-01 00:00:00','2022-04-01 00:00:00');
/*!40000 ALTER TABLE `system` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-30 17:35:05
