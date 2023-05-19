package main

// 模拟分段请求上传文件
/**
func main() {
	//初始化配置客户端
	config := objectstore.GetMinioConfig()
	client, _ := objectstore.NewMinioClient()

	ctx := context.Background()
	//创建项目名称<可以是文件目录形式，也可以是文件名>
	objectName := objectstore.RandString(60, rand.NewSource(time.Now().UnixNano()), "wcx")

	//创建上传ID
	uploadId := objectstore.InitMultipartUpload(ctx, &minio.Core{Client: client}, config.MinIOBucketName, objectName)

	// 打开本地文件
	file, _ := os.Open("/Users/bytedance/goProject/src/objectstore/file/python-v1.tar")
	defer file.Close()
	// 获取文件信息 根据文件大小模拟分段
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	partSize := int64(5 * 1024 * 1024) // 1M
	partsCount := int(math.Ceil(float64(fileSize) / float64(partSize)))
	var completeParts []minio.CompletePart

	for i := 0; i < partsCount; i++ {
		// Calculate the byte range for the current part.
		start := int64(i) * partSize
		end := int64(math.Min(float64(start+partSize), float64(fileSize)))

		partSize := end - start

		// 从文件读取指定大小的文件
		partData := make([]byte, partSize)
		file.ReadAt(partData, start)

		// 循环执行上传任务
		part, err := objectstore.UploadPart(ctx, &minio.Core{Client: client}, config.MinIOBucketName, objectName, partData, i+1, partSize, uploadId)
		if err != nil {
			fmt.Printf("err:%v", err)
		}
		completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
	}
	// 最后调用上传完全任务
	objectstore.CompleteMultipartUpload(ctx, &minio.Core{Client: client}, config.MinIOBucketName, objectName, uploadId, completeParts)
}
*/

func main() {

	//objectstore.NewMysqlClient()
}
