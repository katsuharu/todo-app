CREATE TABLE `todos` (
    `id` CHAR(36) PRIMARY KEY COMMENT 'ID',
    `title` VARCHAR(255) NOT NULL COMMENT 'タイトル',
    `body` VARCHAR(255) NOT NULL COMMENT '本文',
    `created_at` DATETIME(6) NOT NULL COMMENT '作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT '更新日時'
);
