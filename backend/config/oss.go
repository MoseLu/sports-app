// backend/config/oss.go
package config

import (
    "fmt"
    "os"
)

// OSSConfig 存储阿里云 OSS 的配置信息
type OSSConfig struct {
    AccessKeyID     string
    AccessKeySecret string
    Endpoint        string
    Bucket          string
}

// LoadOSSConfig 只从环境变量加载，没有文件回退
func LoadOSSConfig() (*OSSConfig, error) {
    id := os.Getenv("OSS_ACCESS_KEY_ID")
    secret := os.Getenv("OSS_ACCESS_KEY_SECRET")
    endpoint := os.Getenv("OSS_ENDPOINT")
    bucket := os.Getenv("OSS_BUCKET")

    if id == "" || secret == "" || endpoint == "" || bucket == "" {
        return nil, fmt.Errorf("OSS 环境变量未正确设置，请检查 OSS_ACCESS_KEY_ID、OSS_ACCESS_KEY_SECRET、OSS_ENDPOINT、OSS_BUCKET")
    }

    return &OSSConfig{
        AccessKeyID:     id,
        AccessKeySecret: secret,
        Endpoint:        endpoint,
        Bucket:          bucket,
    }, nil
}
