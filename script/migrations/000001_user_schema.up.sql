CREATE TABLE IF NOT EXISTS `user` (
  `id`              varchar(36)     NOT NULL,
  `username`        varchar(32)     NOT NULL,
  `password`        varchar(60)     NOT NULL,
  `is_active`       tinyint(1)      NOT NULL DEFAULT 1,
  `is_superuser`    tinyint(1)      NOT NULL DEFAULT 0,
  `name`            varchar(250)    CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `phone`           varchar(250)    CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `email`           varchar(250)    CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at`      timestamp(6)    NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at`      timestamp(6)    NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `last_login`      timestamp(6)    NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `username_unique` UNIQUE (username),
  KEY `user_email_idx` (`email`)
) ENGINE=InnoDB;

-- group master
CREATE TABLE IF NOT EXISTS `group` (
  `id`        int             NOT NULL AUTO_INCREMENT,
  `name`      varchar(100)    NOT NULL,
  PRIMARY KEY (`id`),
  KEY `group_name_idx` (`name`)
) ENGINE=InnoDB;

-- permission master
CREATE TABLE IF NOT EXISTS `permission` (
  `id`        int           NOT NULL AUTO_INCREMENT,
  `name`      varchar(250)  NOT NULL,
  PRIMARY KEY (`id`),
  KEY `permission_name_idx` (`name`)
) ENGINE=InnoDB;

-- user group many to many relation
CREATE TABLE IF NOT EXISTS `user_groups` (
  `user_id`     varchar(36)     NOT NULL,
  `group_id`    int             NOT NULL,
  PRIMARY KEY (`user_id`, `group_id`),
  CONSTRAINT `user_groups_user_id_FK` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_groups_group_id_FK` FOREIGN KEY (`group_id`) REFERENCES `group` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `group_permissions` (
  `group_id`        int    NOT NULL,
  `permission_id`   int    NOT NULL,
  PRIMARY KEY (`group_id`, `permission_id`),
  CONSTRAINT `group_permissions_group_id_FK` FOREIGN KEY (`group_id`) REFERENCES `group` (`id`) ON DELETE CASCADE,
  CONSTRAINT `group_permissions_permission_id_FK` FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;