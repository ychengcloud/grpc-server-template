CREATE DATABASE IF NOT EXISTS `account` DEFAULT CHARACTER SET = `utf8mb4` COLLATE = utf8mb4_0900_ai_ci;

USE `account`;

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` char(36) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '0 未设置, 1 可用, 2 禁用, 3 注销',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_credential_identifier`;
CREATE TABLE `account_credential_identifier` (
  `id` char(36) NOT NULL,
  `identifier` varchar(255) NOT NULL DEFAULT '',
  `account_credential_id` char(36) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_credential_identifier_identifier_idx` (`identifier`),
  KEY `account_credential_id` (`account_credential_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_credential_type`;
CREATE TABLE `account_credential_type` (
  `id` char(36) NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_credential_type_name_idx` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_credential`;
CREATE TABLE `account_credential` (
  `id` char(36) NOT NULL,
  `config` json NOT NULL DEFAULT (JSON_OBJECT()),
  `account_credential_type_id` char(36) NOT NULL,
  `account_id` char(36) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `account_id` (`account_id`),
  KEY `account_credential_type_id` (`account_credential_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_group`;
CREATE TABLE `account_group` (
  `id` char(36) NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '名字',
  `description` varchar(32) NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT = '用户组';

DROP TABLE IF EXISTS `account_metadata`;
CREATE TABLE `account_metadata` (
  `id` char(36) NOT NULL,
  `account_id` char(36) NOT NULL COMMENT '帐号ID',
  `key` varchar(64) NOT NULL DEFAULT '' COMMENT 'key',
  `value` varchar(64) NOT NULL DEFAULT '' COMMENT 'value',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_group_relation`;
CREATE TABLE `account_group_relation` (
  `account_id` char(36) NOT NULL COMMENT '帐号ID',
  `group_id` char(36) NOT NULL COMMENT '用户组ID',
  PRIMARY KEY (`account_id`,`group_id`),
  KEY `account_group_relation_account_id_idx` (`account_id`),
  KEY `account_group_relation_group_id_idx` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='帐号组关联';

DROP TABLE IF EXISTS `account_recovery_address`;
CREATE TABLE `account_recovery_address` (
  `id` char(36) NOT NULL,
  `via` varchar(16) NOT NULL DEFAULT '',
  `value` varchar(400) NOT NULL DEFAULT '',
  `account_id` char(36) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_verifiable_address_value_via_uq_idx` (`via`,`value`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `account_verifiable_address`;
CREATE TABLE `account_verifiable_address` (
  `id` char(36) NOT NULL,
  `via` varchar(16) NOT NULL DEFAULT '',
  `verified` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0 未验证, 1 已验证',
  `value` varchar(400) NOT NULL DEFAULT '',
  `verified_at` datetime DEFAULT NULL,
  `account_id` char(36) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_verifiable_address_value_via_uq_idx` (`via`,`value`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;