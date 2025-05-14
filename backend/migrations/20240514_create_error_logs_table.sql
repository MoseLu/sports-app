-- 创建错误日志表
CREATE TABLE IF NOT EXISTS `error_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `error_type` varchar(50) NOT NULL DEFAULT 'unknown' COMMENT '错误类型',
  `message` text NOT NULL COMMENT '错误消息',
  `status_code` int NOT NULL DEFAULT 500 COMMENT 'HTTP状态码',
  `path` varchar(255) NOT NULL COMMENT '请求路径',
  `method` varchar(10) NOT NULL COMMENT '请求方法',
  `request_body` text COMMENT '请求体',
  `response_body` text COMMENT '响应体',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_error_type` (`error_type`),
  KEY `idx_status_code` (`status_code`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; 