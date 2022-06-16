CREATE TABLE `dict`
(
    `k` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '键',
    `v` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '值',
    PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典表';