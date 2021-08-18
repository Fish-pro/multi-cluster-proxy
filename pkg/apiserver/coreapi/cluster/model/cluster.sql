CREATE TABLE `cluster` (
    `id` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '集群ID',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '集群名称',
    `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '集群地址',
    `token` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '集群token',
    PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='集群信息表';