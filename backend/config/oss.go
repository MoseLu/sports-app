// backend/config/oss.go
package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// OSSConfig 存储阿里云 OSS 的配置信息
type OSSConfig struct {
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Endpoint        string `json:"endpoint"`
	Bucket          string `json:"bucket"`
}

// 全局 OSS 配置实例
var OSS *OSSConfig

func init() {
	var err error
	OSS, err = LoadOSSConfig()
	if err != nil {
		log.Fatalf("加载 OSS 配置失败：%v", err)
	}
}

// LoadOSSConfig 从环境变量或配置文件中加载 OSS 配置
func LoadOSSConfig() (*OSSConfig, error) {
	config := &OSSConfig{
		AccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
		Endpoint:        os.Getenv("OSS_ENDPOINT"),
		Bucket:          os.Getenv("OSS_BUCKET"),
	}

	// 如果环境变量未设置，尝试从配置文件加载
	if config.AccessKeyID == "" || config.AccessKeySecret == "" || config.Endpoint == "" || config.Bucket == "" {
		file, err := os.Open("config/oss.json")
		if err != nil {
			return nil, err
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(config); err != nil {
			return nil, err
		}
	}

	// 检查配置是否完整
	if config.AccessKeyID == "" || config.AccessKeySecret == "" || config.Endpoint == "" || config.Bucket == "" {
		return nil, fmt.Errorf("OSS 配置不完整，请检查环境变量或配置文件")
	}

	return config, nil
}

// GetOSSAccessKeyID 返回 AccessKeyID
func GetOSSAccessKeyID() string {
	return OSS.AccessKeyID
}

// GetOSSAccessKeySecret 返回 AccessKeySecret
func GetOSSAccessKeySecret() string {
	return OSS.AccessKeySecret
}

// GetOSSEndpoint 返回 Endpoint
func GetOSSEndpoint() string {
	return OSS.Endpoint
}

// GetOSSBucket 返回 Bucket
func GetOSSBucket() string {
	return OSS.Bucket
}
