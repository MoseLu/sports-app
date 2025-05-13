// backend/services/upload.go
package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sports-app/backend/config"
)

type UploadService struct{}

// UploadImage 上传图片到 OSS 并返回可访问 URL
func (s *UploadService) UploadImage(ctx context.Context, file *multipart.FileHeader, userID int64) (string, error) {
	// 打开文件流
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %w", err)
	}
	defer src.Close()

	// 加载 OSS 配置
	endpoint := config.GetOSSEndpoint()
	accessKeyID := config.GetOSSAccessKeyID()
	accessKeySecret := config.GetOSSAccessKeySecret()
	bucketName := config.GetOSSBucket()

	// 初始化 OSS 客户端
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return "", fmt.Errorf("初始化 OSS 客户端失败: %w", err)
	}

	// 获取 Bucket 实例
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("获取 OSS Bucket 实例失败: %w", err)
	}

	// 生成对象 Key
	ext := path.Ext(file.Filename)
	objectKey := fmt.Sprintf("exercises/%d/%d%s", userID, time.Now().UnixNano(), ext)

	// 上传到 OSS
	if err := bucket.PutObject(objectKey, src); err != nil {
		return "", fmt.Errorf("上传文件到 OSS 失败: %w", err)
	}

	// 拼接公有访问 URL
	baseURL := fmt.Sprintf("https://%s.%s/", bucketName, endpoint)
	return baseURL + objectKey, nil
}

// DeleteImage 从 OSS 删除图片
func (s *UploadService) DeleteImage(ctx context.Context, imageURL string) error {
	// 加载 OSS 配置
	endpoint := config.GetOSSEndpoint()
	accessKeyID := config.GetOSSAccessKeyID()
	accessKeySecret := config.GetOSSAccessKeySecret()
	bucketName := config.GetOSSBucket()

	// 解析 objectKey
	baseURL := fmt.Sprintf("https://%s.%s/", bucketName, endpoint)
	if !strings.HasPrefix(imageURL, baseURL) {
		return fmt.Errorf("无效的 imageURL: %s", imageURL)
	}
	objectKey := imageURL[len(baseURL):]

	// 初始化 OSS 客户端
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return fmt.Errorf("初始化 OSS 客户端失败: %w", err)
	}

	// 获取 Bucket 实例
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return fmt.Errorf("获取 OSS Bucket 实例失败: %w", err)
	}

	// 删除对象
	if err := bucket.DeleteObject(objectKey); err != nil {
		return fmt.Errorf("删除 OSS 对象失败: %w", err)
	}

	return nil
}
