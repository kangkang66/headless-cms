/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : cms

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 21/11/2021 22:51:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for datas
-- ----------------------------
DROP TABLE IF EXISTS `datas`;
CREATE TABLE `datas` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `model_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `data` json NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `model_name` (`model_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of datas
-- ----------------------------
BEGIN;
INSERT INTO `datas` VALUES (1, 'source', '{\"name\": \"taobao\"}', 0, '2021-11-21 12:40:01', '2021-11-21 12:40:01', NULL);
INSERT INTO `datas` VALUES (2, 'like', '{\"name\": \"apple\", \"buy_source\": 1}', 0, '2021-11-21 12:40:01', '2021-11-21 12:40:01', NULL);
INSERT INTO `datas` VALUES (3, 'like', '{\"name\": \"banana\", \"buy_source\": 1}', 0, '2021-11-21 12:40:01', '2021-11-21 12:40:01', NULL);
INSERT INTO `datas` VALUES (4, 'user', '{\"name\": \"kang\", \"likes\": [2, 3], \"user_source\": 1}', 0, '2021-11-21 12:40:01', '2021-11-21 04:42:32', NULL);
COMMIT;

-- ----------------------------
-- Table structure for models
-- ----------------------------
DROP TABLE IF EXISTS `models`;
CREATE TABLE `models` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `desc` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `status` tinyint DEFAULT '0',
  `fields` json NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of models
-- ----------------------------
BEGIN;
INSERT INTO `models` VALUES (2, 'source', '购买渠道', 0, '[{\"desc\": \"渠道名称\", \"name\": \"name\", \"valid\": null, \"status\": 0, \"default\": \"\", \"is_must\": false, \"is_show\": true, \"is_unique\": false, \"nick_name\": \"name\", \"show_type\": \"string\", \"storage_type\": \"varchar\", \"link_model_name\": \"\"}]', '2021-11-21 11:13:35', '2021-11-21 11:13:35', NULL);
INSERT INTO `models` VALUES (4, 'like', '喜欢', 0, '[{\"desc\": \"名称\", \"name\": \"name\", \"valid\": [{\"rule\": 20, \"rule_name\": \"max_value\"}], \"status\": 0, \"default\": \"\", \"is_must\": false, \"is_show\": true, \"is_unique\": false, \"nick_name\": \"name\", \"show_type\": \"string\", \"storage_type\": \"varchar\", \"link_model_name\": \"\"}, {\"desc\": \"购买渠道\", \"name\": \"buy_source\", \"valid\": null, \"status\": 0, \"default\": \"\", \"is_must\": false, \"is_show\": true, \"is_unique\": false, \"nick_name\": \"buy_source\", \"show_type\": \"link_model_one_one\", \"storage_type\": \"link_model_one_one\", \"link_model_name\": \"source\"}]', '2021-11-21 11:34:00', '2021-11-21 03:57:07', NULL);
INSERT INTO `models` VALUES (7, 'user', '用户', 0, '[{\"desc\": \"姓名\", \"name\": \"name\", \"valid\": null, \"status\": 0, \"default\": \"\", \"is_must\": false, \"is_show\": true, \"is_unique\": false, \"nick_name\": \"name\", \"show_type\": \"string\", \"storage_type\": \"varchar\", \"link_model_name\": \"\"}, {\"desc\": \"爱好\", \"name\": \"likes\", \"valid\": null, \"status\": 0, \"default\": \"\", \"is_must\": false, \"is_show\": true, \"is_unique\": false, \"nick_name\": \"likes\", \"show_type\": \"link_model_one_many\", \"storage_type\": \"link_model_one_many\", \"link_model_name\": \"like\"}, {\"desc\": \"渠道\", \"name\": \"user_source\", \"valid\": null, \"status\": 0, \"default\": \"\", \"is_must\": false, \"is_show\": true, \"is_unique\": false, \"nick_name\": \"user_source\", \"show_type\": \"link_model_one_one\", \"storage_type\": \"link_model_one_one\", \"link_model_name\": \"source\"}]', '2021-11-21 12:10:03', '2021-11-21 04:29:11', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
