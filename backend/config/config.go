package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config 应用配置
type Config struct {
	Database     DBConfig `yaml:"database"`
	LogsDB       DBConfig `yaml:"logs_database"`
	OSS          OSSConfig `yaml:"oss"`
	Server       ServerConfig `yaml:"server"`
	JWT          JWTConfig `yaml:"jwt"`
	Email        EmailConfig `yaml:"email"`
	Redis        RedisConfig `yaml:"redis"`
	Log          LogConfig `yaml:"log"`
	Verification VerificationConfig `yaml:"verification"`
}

// DBConfig 数据库配置
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level    string `yaml:"level"`
	Filename string `yaml:"filename"`
}

// VerificationConfig 验证码配置
type VerificationConfig struct {
	CodeExpire  int `yaml:"code_expire"`
	MaxAttempts int `yaml:"max_attempts"`
}

var (
	cfg    *Config
	db     *gorm.DB
	logsDB *gorm.DB
)

// GetConfig 获取配置
func GetConfig() *Config {
	if cfg == nil {
		// 检查是否是本地运行环境
		isLocal := os.Getenv("LOCAL_DEV") == "true"
		
		// 根据环境设置数据库配置
		var dbHost, logsDbHost string
		if isLocal {
			dbHost = "https://www.redamancy.com.cn"
			logsDbHost = "https://www.redamancy.com.cn"
		} else {
			dbHost = getEnv("DB_HOST", "127.0.0.1")
			logsDbHost = getEnv("LOGS_DB_HOST", "127.0.0.1")
		}

		cfg = &Config{
			Database: DBConfig{
				Host:     dbHost,
				Port:     getEnv("DB_PORT", "3306"),
				User:     getEnv("DB_USER", "root"),
				Password: getEnv("DB_PASSWORD", "123456"),
				Name:     getEnv("DB_NAME", "sports_app"),
			},
			LogsDB: DBConfig{
				Host:     logsDbHost,
				Port:     getEnv("LOGS_DB_PORT", "3306"),
				User:     getEnv("LOGS_DB_USER", "root"),
				Password: getEnv("LOGS_DB_PASSWORD", "123456"),
				Name:     getEnv("LOGS_DB_NAME", "sports_app_logs"),
			},
			JWT: JWTConfig{
				PrivateKeyPath: getEnv("JWT_PRIVATE_KEY_PATH", filepath.Join("config", "keys", "private.pem")),
				PublicKeyPath:  getEnv("JWT_PUBLIC_KEY_PATH", filepath.Join("config", "keys", "public.pem")),
				ExpiresIn:      int(getEnvInt64("JWT_EXPIRE", 86400)), // 默认24小时
			},
			Server: ServerConfig{
				Port: getEnv("PORT", "8080"),
			},
		}
		log.Printf("主数据库配置: %+v", cfg.Database)
		log.Printf("日志数据库配置: %+v", cfg.LogsDB)
	}
	return cfg
}

// GetDB 获取主数据库连接
func GetDB() *gorm.DB {
	if db == nil {
		initDB()
	}
	return db
}

// GetLogsDB 获取日志数据库连接
func GetLogsDB() *gorm.DB {
	if logsDB == nil {
		initLogsDB()
	}
	return logsDB
}

// initDB 初始化主数据库连接
func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	// 配置数据库连接池
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            false,                                // 关闭预编译语句
		Logger:                 logger.Default.LogMode(logger.Error), // 只记录错误日志
		SkipDefaultTransaction: true,                                 // 跳过默认事务
		CreateBatchSize:        1000,                                 // 批量创建的大小
	})
	if err != nil {
		panic(fmt.Sprintf("连接主数据库失败: %v", err))
	}

	db = gormDB
}

// initLogsDB 初始化日志数据库连接
func initLogsDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.LogsDB.User,
		cfg.LogsDB.Password,
		cfg.LogsDB.Host,
		cfg.LogsDB.Port,
		cfg.LogsDB.Name,
	)

	// 配置数据库连接池
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            false,                                // 关闭预编译语句
		Logger:                 logger.Default.LogMode(logger.Error), // 只记录错误日志
		SkipDefaultTransaction: true,                                 // 跳过默认事务
		CreateBatchSize:        1000,                                 // 批量创建的大小
	})
	if err != nil {
		panic(fmt.Sprintf("连接日志数据库失败: %v", err))
	}

	logsDB = gormDB
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt64 获取环境变量并转换为int64，如果不存在则返回默认值
func getEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := time.ParseDuration(value); err == nil {
			return int64(intValue.Seconds())
		}
	}
	return defaultValue
}

// JWTConfig JWT 配置
type JWTConfig struct {
	SecretKey      string `yaml:"secret_key"`
	ExpiresIn      int    `yaml:"expires_in"`
	PrivateKeyPath string `yaml:"private_key_path"`
	PublicKeyPath  string `yaml:"public_key_path"`
}

// LoadConfig 从配置文件加载配置
func LoadConfig() (*Config, error) {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
