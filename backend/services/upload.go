package services

import (
	"fmt"
	"mime/multipart"
	"path"
	"time"

	"sports-app/backend/config"
)

type UploadService struct{}

// UploadImage 上传图片到OSS
func (s *UploadService) UploadImage(file *multipart.FileHeader, userID int64) (string, error) {
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close()

	// 获取OSS Bucket
	bucket, err := config.GetOSSBucket()
	if err != nil {
		return "", fmt.Errorf("获取OSS Bucket失败: %v", err)
	}

	// 生成文件名
	ext := path.Ext(file.Filename)
	objectName := fmt.Sprintf("exercises/%d/%d%s", userID, time.Now().UnixNano(), ext)

	// 上传文件到OSS
	err = bucket.PutObject(objectName, src)
	if err != nil {
		return "", fmt.Errorf("上传文件到OSS失败: %v", err)
	}

	// 返回可访问的URL
	imageURL := config.OSS.BaseURL + objectName
	return imageURL, nil
}

// DeleteImage 从OSS删除图片
func (s *UploadService) DeleteImage(imageURL string) error {
	// 从URL中提取对象名称
	objectName := imageURL[len(config.OSS.BaseURL):]

	// 获取OSS Bucket
	bucket, err := config.GetOSSBucket()
	if err != nil {
		return fmt.Errorf("获取OSS Bucket失败: %v", err)
	}

	// 删除文件
	err = bucket.DeleteObject(objectName)
	if err != nil {
		return fmt.Errorf("删除OSS文件失败: %v", err)
	}

	return nil
} 