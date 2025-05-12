CREATE TABLE IF NOT EXISTS update_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    old_version VARCHAR(50) NOT NULL COMMENT '旧版本号',
    new_version VARCHAR(50) NOT NULL COMMENT '新版本号',
    status VARCHAR(20) NOT NULL COMMENT '更新状态：success/failed',
    error TEXT COMMENT '错误信息',
    INDEX idx_old_version (old_version),
    INDEX idx_new_version (new_version),
    INDEX idx_status (status),
    INDEX idx_updated_at (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; 