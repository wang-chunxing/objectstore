package objectstore

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

func DownloadPart(minioClient *minio.Client, minIOBucketName, objectName string, fromByte int64, toByte int64) (int64, error) {
	options := minio.GetObjectOptions{}
	options.SetRange(fromByte, toByte)
	//minioClient.FGetObject()
	remoteReader, err := minioClient.GetObject(context.Background(), minIOBucketName, objectName, options)
	if err != nil {
		fmt.Println("Error getting remote object:", err)
		return 0, err
	}
	remoteReader.Close()
	return 0, err
}
