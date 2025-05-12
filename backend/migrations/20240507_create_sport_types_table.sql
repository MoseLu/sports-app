-- 创建运动类型表
CREATE TABLE IF NOT EXISTS `sport_types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '运动类型名称',
  `description` text COMMENT '运动类型描述',
  `icon` varchar(255) DEFAULT NULL COMMENT '运动类型图标',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入默认运动类型
INSERT INTO `sport_types` (`name`, `description`, `icon`) VALUES
('跑步', '户外或室内跑步运动', 'running'),
('游泳', '游泳运动', 'swimming'),
('骑行', '自行车骑行运动', 'cycling'),
('健身', '健身房锻炼', 'fitness'),
('瑜伽', '瑜伽练习', 'yoga'),
('篮球', '篮球运动', 'basketball'),
('足球', '足球运动', 'football'),
('网球', '网球运动', 'tennis'),
('羽毛球', '羽毛球运动', 'badminton'),
('乒乓球', '乒乓球运动', 'table-tennis'); 