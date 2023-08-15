package oss

import (
	"context"
	"github.com/Tiktok-Lite/kotkit/pkg/oss"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMinioClientInit(t *testing.T) {
	exist, err := oss.Minio().BucketExists(context.Background(), "none-existing-bucket")
	assert.Equal(t, exist, false)
	assert.Equal(t, err, nil)
}

func TestCreateBucket(t *testing.T) {
	// 创建Minio客户端对象
	client := oss.Minio()

	// 测试创建存储桶
	bucketName := "test-bucket"
	err := oss.CreateBucket(bucketName)
	if err != nil {
		t.Fatalf("failed to create bucket '%s': %v", bucketName, err)
	}

	// 确认存储桶存在
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		t.Fatalf("failed to check bucket existence: %v", err)
	}
	if !exists {
		t.Fatalf("bucket '%s' does not exist after creation", bucketName)
	}

	// 清理测试数据
	err = client.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		t.Fatalf("failed to remove bucket '%s': %v", bucketName, err)
	}

}
