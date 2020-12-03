CREATE TABLE `user_transactions`
(
    `id`              INT       NOT NULL AUTO_INCREMENT,
    `user_id`         INT       NOT NULL,
    `amount`          INT       NOT NULL,
    `current_balance` INT       NOT NULL,
    `description`     VARCHAR(150) CHARACTER SET utf8 COLLATE utf8_general_ci,
    `created_at`      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `user_id_index` (`user_id`) USING HASH,
    PRIMARY KEY (`id`)
);