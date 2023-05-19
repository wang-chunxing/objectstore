package objectstore

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/aws/aws-sdk-go/aws"
	awscredentials "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/go-redis/redis"
)

func NewMinioClient() (minioClient *minio.Client, err error) {
	// 初始化MinIO客户端
	minioClient, err = minio.New(MinioCf.MinIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioCf.MinIOAccessKeyID, MinioCf.MinIOSecretAccessKey, ""),
		Secure: MinioCf.MinIOUseSSL,
	})
	if err != nil {
		log.Fatalln("Error creating MinIO client:", err)
		return
	}
	return
}

func NewS3Client() (s3Client *s3.S3, err error) {
	newSession, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(MinioCf.MinIOEndpoint),
		Region:      aws.String(MinioCf.MinIORegion),
		Credentials: awscredentials.NewStaticCredentials(MinioCf.MinIOAccessKeyID, MinioCf.MinIOSecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln("Error creating MinIOS3 client:", err)
	}
	s3Client = s3.New(newSession)
	return

}

func NewMysqlClient() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlCf.UserName,
		MysqlCf.Password,
		MysqlCf.Host,
		MysqlCf.Port,
		MysqlCf.Database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	return
	//}
	//user := User{Name: "Alice", Age: 18}
	//tx := db.Begin()
	//if tx.Error != nil {
	//	// 错误处理
	//}
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	} else if tx.Error != nil {
	//		tx.Rollback()
	//	} else {
	//		tx.Commit()
	//	}
	//}()
	//
	//// 在事务中执行数据库操作
	//tx.Create(&user)
	//tx.Model(&user).Update("age", 20)
}

func NewRedisClient() (rClient *redis.Client, err error) {

	opts, err := redis.ParseURL(RedisUrl)
	if err != nil {
		return
	}
	rClient = redis.NewClient(opts)
	pong, err := rClient.Ping().Result()
	if err != nil {
		return
	}
	log.Println(pong)
	return
}
