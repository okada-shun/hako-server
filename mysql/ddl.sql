DROP DATABASE IF EXISTS `hako`;
CREATE DATABASE `hako`;
USE `hako`;

DROP TABLE IF EXISTS `hako`.`hako_infos`;
CREATE TABLE IF NOT EXISTS `hako`.`hako_infos`(
  `address` CHAR(64) PRIMARY KEY NOT NULL,
  `balance` INT NOT NULL,
  `total_supply` INT NOT NULL,
  `credit` INT NOT NULL,
  `debt` INT NOT NULL,
  `member_count` INT NOT NULL,
  `upper_limit` INT NOT NULL,
  `name` VARCHAR(32) NOT NULL,
  `symbol` VARCHAR(32) NOT NULL,
  `decimals` INT NOT NULL
);

/*
INSERT INTO hako_infos(address, balance, total_supply, credit, debt, member_count, upper_limit, name, symbol, decimals) 
VALUES ("", 0, 0, 0, 0, 0, 0, "", "", 0);
*/

DROP TABLE IF EXISTS `hako`.`user_infos`;
CREATE TABLE IF NOT EXISTS `hako`.`user_infos`(
  `address` CHAR(64) PRIMARY KEY NOT NULL,
  `balance` INT NOT NULL,
  `member` INT NOT NULL,
  `credit` INT NOT NULL,
  `debt` INT NOT NULL,
  `lending` INT NOT NULL,
  `borrowing` INT NOT NULL,
  `net_assets` INT NOT NULL,
  `value` INT NOT NULL,
  `duration` INT NOT NULL
);

/*
INSERT INTO user_infos(address, balance, member, credit, debt, lending, borrowing, net_assets, value, duration) 
VALUES ("", 0, 0, 0, 0, 0, 0, 0, 0, 0);
*/

DROP TABLE IF EXISTS `hako`.`owner_infos`;
CREATE TABLE IF NOT EXISTS `hako`.`owner_infos`(
  `hako_address` CHAR(64) PRIMARY KEY NOT NULL,
  `owner_address` CHAR(64) NOT NULL,
  `hako_balance` INT NOT NULL,
  `total_supply` INT NOT NULL,
  `hako_credit` INT NOT NULL,
  `hako_debt` INT NOT NULL,
  `owner_balance` INT NOT NULL,
  `lending_count` INT NOT NULL,
  `member_count` INT NOT NULL,
  `upper_limit` INT NOT NULL
);

/*
INSERT INTO owner_infos(hako_address, owner_address, hako_balance, total_supply, hako_credit, hako_debt, owner_balance, lending_count, member_count, upper_limit) 
VALUES ("", "", 0, 0, 0, 0, 0, 0, 0, 0);
*/

DROP TABLE IF EXISTS `hako`.`transfer_token_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`transfer_token_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `transfer_from` CHAR(64) NOT NULL,
  `transfer_to` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO transfer_token_histories(transfer_from, transfer_to, value, tx_hash, block_n) 
VALUES ("", "", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`transfer_credit_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`transfer_credit_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `transfer_from` CHAR(64) NOT NULL,
  `transfer_to` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO transfer_credit_histories(transfer_from, transfer_to, value, tx_hash, block_n) 
VALUES ("", "", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`join_hako_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`join_hako_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `new_member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO join_hako_histories(new_member, value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`leave_hako_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`leave_hako_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO leave_hako_histories(member, value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`deposit_token_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`deposit_token_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO deposit_token_histories(member, value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`withdraw_token_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`withdraw_token_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO withdraw_token_histories(member, value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`register_borrowing_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`register_borrowing_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `duration` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO register_borrowing_histories(member, value, duration, tx_hash, block_n) 
VALUES ("", 0, 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`lend_credit_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`lend_credit_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `lend_from` CHAR(64) NOT NULL,
  `lend_to` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `duration` INT NOT NULL,
  `lending_id` INT NOT NULL,
  `time` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO lend_credit_histories(lend_from, lend_to, value, duration, id, time, tx_hash, block_n) 
VALUES ("", "", 0, 0, 0, 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`collect_debt_from_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`collect_debt_from_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `creditor` CHAR(64) NOT NULL,
  `debtor` CHAR(64) NOT NULL,
  `lending_id` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO collect_debt_from_histories(creditor, debtor, id, tx_hash, block_n) 
VALUES ("", "", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`return_debt_to_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`return_debt_to_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `debtor` CHAR(64) NOT NULL,
  `creditor` CHAR(64) NOT NULL,
  `lending_id` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO return_debt_to_histories(debtor, creditor, id, tx_hash, block_n) 
VALUES ("", "", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`create_credit_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`create_credit_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO create_credit_histories(member, value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`reduce_debt_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`reduce_debt_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `member` CHAR(64) NOT NULL,
  `value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO reduce_debt_histories(member, value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`change_hako_owner_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`change_hako_owner_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `old_hako_owner` CHAR(64) NOT NULL,
  `new_hako_owner` CHAR(64) NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO change_hako_owner_histories(old_hako_owner, new_hako_owner, tx_hash, block_n) 
VALUES ("", "", "", 0);
*/

DROP TABLE IF EXISTS `hako`.`change_upper_limit_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`change_upper_limit_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `hako_owner` CHAR(64) NOT NULL,
  `new_upper_limit` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO change_upper_limit_histories(hako_owner, new_upper_limit, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/

DROP TABLE IF EXISTS `hako`.`get_reward_histories`;
CREATE TABLE IF NOT EXISTS `hako`.`get_reward_histories`(
  `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `hako_owner` CHAR(64) NOT NULL,
  `reward_value` INT NOT NULL,
  `tx_hash` CHAR(128) NOT NULL,
  `block_n` INT NOT NULL
);

/*
INSERT INTO get_reward_histories(hako_owner, reward_value, tx_hash, block_n) 
VALUES ("", 0, "", 0);
*/
