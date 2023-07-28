CREATE TABLE IF NOT EXISTS
  `rlogs` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `read_url` VARCHAR(1024),
    `content` text,
    `state` BIGINT,
    `create_time` DATETIME,
    `update_time` DATETIME,
    PRIMARY KEY (`id`)
  ) 
--  如果需要节约空间可 使用ttl
  TTL = `create_time` + INTERVAL 6 MONTH;

CREATE TABLE
  `tags` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    `type` bigint(20) DEFAULT NULL,
    `emo` varchar(255) DEFAULT NULL,
    `state` bigint(20) DEFAULT NULL,
    `create_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
  );

  
CREATE TABLE `l2tags` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`tag` bigint(20) DEFAULT NULL,
`log` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) 
);