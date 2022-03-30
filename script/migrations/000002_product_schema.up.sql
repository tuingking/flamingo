CREATE TABLE IF NOT EXISTS `product` (
  `id`              bigint          NOT NULL AUTO_INCREMENT,
  `name`            varchar(250)    CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `price`           DECIMAL(18,4)   NOT NULL DEFAULT 0,
  `created_at`      timestamp(6)    NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at`      timestamp(6)    NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;