CREATE DATABASE todo_list;

USE todo_list;

CREATE SCHEMA IF NOT EXISTS `todo_list` DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE IF NOT EXISTS `todo_list`
(
    `id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(128) GENERATED ALWAYS AS(JSON_UNQUOTE(JSON_EXTRACT(`content`, '$.user_id.value'))) VIRTUAL NOT NULL,
    `content` JSON NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,

    PRIMARY KEY(`id`),
    UNIQUE INDEX `id_UNIQUE`(`id` ASC),
    UNIQUE INDEX `user_id_UNIQUE`(`user_id` ASC)
) ENGINE = InnoDB;