package oss

import (
	"bytes"
	"context"
	"github.com/Tiktok-Lite/kotkit/pkg/conf"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/tools"
	"github.com/Tiktok-Lite/kotkit/pkg/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"sync"
)

var (
	once            sync.Once
	minioClient     *minio.Client
	minioConf       = conf.LoadConfig(constant.DefaultMinioConfigName)
	endpoint        = minioConf.GetString("configure.endpoint")
	accessKeyID     = minioConf.GetString("configure.accessKeyID")
	secretAccessKey = minioConf.GetString("configure.secretAccessKey")
	useSSL          = minioConf.GetBool("configure.useSSL")
	expiryTime      = minioConf.GetDuration("expiryTime")
	videoBucketName = minioConf.GetString("name.VideoBucket")
	coverBucketName = minioConf.GetString("name.CoverBucket")
)

type MinioClient struct {
	*minio.Client
}

func newMinioClient() (*minio.Client, error) {
	logger := log.Logger()

	minioClient_, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logger.Errorf("minio client init error: %v", err)
		return nil, err
	}
	logger.Infof("minio client init success")

	return minioClient_, nil
}

func Minio() *minio.Client {
	once.Do(func() {
		minioClient, _ = newMinioClient()
	})

	return minioClient
}

func PublishVideo(data []byte, videoTitle, coverTitle string) error {
	logger := log.Logger()

	playURL, err := UploadVideo(data, videoTitle)
	if err != nil {
		logger.Errorf("upload video error: %v", err)
		return err
	}

	err = UploadCover(playURL, coverTitle)
	if err != nil {
		logger.Errorf("upload cover error: %v", err)
		return err
	}

	return nil
}

func UploadVideo(data []byte, videoTitle string) (string, error) {
	logger := log.Logger()
	if len(data) == 0 || len(videoTitle) == 0 {
		logger.Errorf("upload video to minio error: data or videoTitle is empty")
		return "", errors.New("data or videoTitle is empty")
	}

	videoBytes := bytes.NewReader(data)
	// upload video to minio
	uploadInfo, err := Minio().PutObject(context.Background(), videoBucketName, videoTitle, videoBytes, int64(len(data)), minio.PutObjectOptions{
		ContentType: "video/mp4",
	})
	if err != nil {
		logger.Errorf("upload video to minio error: %v", err)
		return "", err
	}
	logger.Infof("upload video to minio success: %v", uploadInfo)

	playURL, err := GetObjectURL(videoBucketName, videoTitle)
	if err != nil {
		logger.Errorf("get video url error: %v", err)
		return "", err
	}

	return playURL, nil
}

func GetObjectURL(bucketName, titleName string) (string, error) {
	logger := log.Logger()

	// get object url
	objectURL, err := Minio().PresignedGetObject(context.Background(), bucketName, titleName, expiryTime, nil)
	if err != nil {
		logger.Errorf("get object url error: %v", err)
		return "", err
	}

	logger.Infof("get object url success: %v", objectURL)

	return objectURL.String(), nil
}

func UploadCover(playURL, coverTitle string) error {
	logger := log.Logger()

	// 第一帧作为封面
	buf, err := tools.GetScreenshotBuffer(playURL, 1)
	if err != nil {
		logger.Errorf("Failed to get cover image. %v", err)
		return err
	}
	contentType := "image/jpg"

	uploadInfo, err := Minio().PutObject(context.Background(), coverBucketName, coverTitle, buf, int64(buf.Len()), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		logger.Errorf("Failed to put cover to OSS. %v", err)
		return err
	}

	coverURL, err := GetObjectURL(coverBucketName, coverTitle)
	if err != nil {
		logger.Errorf("Failed to get cover url from OSS. %v", err)
		return err
	}

	logger.Infof("Successfully upload cover to %v, size is %v", coverURL, uploadInfo.Size)

	return nil
}
