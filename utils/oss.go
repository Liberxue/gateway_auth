package utils

import (
	"bytes"
	"io"

	"github.com/Liberxue/gateway_auth/conf"
	"github.com/Liberxue/gateway_auth/middleware/zap"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossClient *oss.Client

func InitAliyunOss(c *conf.ArgConfig) error {
	client, err := oss.New(c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
	if err != nil {
		zap.Log.Errorf("InitAliyunOss Error: %v", err)
		return err
	}
	ossClient = client
	return err
}

func UpdateOss(bucketName, objectName string, objectBytes []byte) error {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		zap.Log.Errorf("UpdateOss bucket Error: %v", err)
		return err
	}
	err = bucket.PutObject(objectName, bytes.NewReader([]byte(objectBytes)))
	if err != nil {
		zap.Log.Errorf("UpdateOss Error: %v", err)
		return err
	}
	return nil
}
func CacheOSS(bucketName, objectName string, data io.Reader) error {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		zap.Log.Errorf("UpdateOss bucket Error: %v", err)
		return err
	}
	err = bucket.PutObject(objectName, data)
	if err != nil {
		zap.Log.Errorf("UpdateOss Error: %v", err)
		return err
	}
	return nil
}

func GetOssObjectPath(bucketName, objectName string) (string, error) {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		zap.Log.Errorf("GetOssObjectPath bucket Error: %v", err)
		return "", err
	}
	//32400 9小时过期
	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 32400)
	if err != nil {
		zap.Log.Errorf("GetOssObjectPath signedURL Error: %v", err)
		return "", err
	}
	return signedURL, nil
}

func IsObjectExist(bucketName, objectName string) bool {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		zap.Log.Errorf("IsObjectExist bucket Error: %v", err)
		return false
	}
	// 判断文件是否存在。
	isExist, err := bucket.IsObjectExist(objectName)
	if err != nil && !isExist {
		zap.Log.Errorf("IsObjectExist objectName Error: %v", err)
		return false
	}
	return isExist
}

func GetOssVideoPath(bucketName, objectName string) (string, error) {
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	//32400 9小时过期
	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 32400)
	if err != nil {
		zap.Log.Errorf("GetOssVideoPath signedURL Error: %v", err)
		return "", err
	}
	return signedURL, nil
}
