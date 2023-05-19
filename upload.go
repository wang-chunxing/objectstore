package objectstore

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

/**
实现文件上传功能：
	a. 前端分片上传
		1.前端分片完成，请求后端获取上传ID以及桶名称以及文件的sha256值
		2.前端分段请求上传(分段的id序号，和分段文件的sha256值，上传路径或者object名称)
	b. 依次将每个分片上传到Minio，并记录每个分片的信息（如分片序号、分片大小、上传时间等元信息）保存到数据库中
	c. 所有分片上传成功后，将分片的信息以及文件的元数据（如文件名称、大小、sha256等）保存到数据库中，返回文件ID给前端
*/

func InitMultipartUpload(ctx context.Context, minioClient *minio.Core, minIOBucketName, objectName string) (uploadID string) {
	// 分段上传创建一个新的Upload ID
	uploadID, err := minioClient.NewMultipartUpload(ctx, minIOBucketName, objectName, minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Error creating multipart upload:", err)
		return
	}
	fmt.Println("Created new multipart upload with ID:", uploadID)
	return uploadID
}

func UploadPart(ctx context.Context, minioClient *minio.Core, minIOBucketName, objectName string, data []byte, partNumber int, partSize int64, uploadID string) (minio.ObjectPart, error) {
	// fileName为文件名称, partNumber为分片数，例如1表示第一片; data为需要上传的数据
	part, err := minioClient.PutObjectPart(ctx, minIOBucketName, objectName, uploadID, partNumber, bytes.NewReader(data), partSize, "", "", nil)
	if err != nil {
		fmt.Println("Error uploading part", partNumber, ":", err)
		err = minioClient.AbortMultipartUpload(ctx, minIOBucketName, objectName, uploadID)
		if err != nil {
			fmt.Println("Error uploading part", partNumber, ":", err)
		}
	}
	fmt.Println("Uploaded part", partNumber, "successfully")
	return part, nil
}

func CompleteMultipartUpload(ctx context.Context, minioClient *minio.Core, minIOBucketName, object string, uploadID string, parts []minio.CompletePart) {
	_, err := minioClient.CompleteMultipartUpload(ctx, minIOBucketName, object, uploadID, parts, minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Error completing multipart upload:", err)
		err = minioClient.AbortMultipartUpload(ctx, minIOBucketName, object, uploadID)
		if err != nil {
			return
		}
		return
	}
	fmt.Println("Multipart upload completed successfully")
}
