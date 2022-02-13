CREATE DATABASE todo_list;

USE todo_list;

CREATE SCHEMA IF NOT EXISTS `todo_list` DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE IF NOT EXISTS `todo_list`
(
    `id` VARCHAR(36) NOT NULL,
    `content` JSON NOT NULL,
    `user_id` VARCHAR(64) GENERATED ALWAYS AS(JSON_UNQUOTE(JSON_EXTRACT(`content`, '$.user_id'))) VIRTUAL NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,

    PRIMARY KEY(`id`),
    UNIQUE INDEX `id_UNIQUE`(`id` ASC),
) ENGINE = InnoDB;

INSERT INTO `todo_list`(`id`,`content`,`created`,`updated`) VALUE ('aaaa',JSON_OBJECT('user_id','riku0202','title','今日の晩御飯','description','シチュー'),NOW(),NOW());