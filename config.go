package objectstore

import "sync"

type MinioConfig struct {
	MinIOEndpoint        string
	MinIOAccessKeyID     string
	MinIOSecretAccessKey string
	MinIOUseSSL          bool
	MinIOBucketName      string
	MinIORegion          string
}

type MysqlConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	Database string
}

var MinioCf *MinioConfig
var MysqlCf *MysqlConfig
var RedisUrl string

var once sync.Once

func init() {
	once.Do(func() {
		MinioCf = &MinioConfig{
			MinIOEndpoint:        "10.37.48.71:9000",
			MinIOAccessKeyID:     "E7zOykMaKaNE5dRd",
			MinIOSecretAccessKey: "bUSPBYCCKESYEHJmKlGMe8ya0tcUsrOC",
			MinIOUseSSL:          false,
			MinIOBucketName:      "wcx",
			MinIORegion:          "beijing-wcx-test",
		}
		MysqlCf = &MysqlConfig{
			UserName: "root",
			Password: "Password1",
			Host:     "localhost",
			Port:     "3306",
			Database: "test",
		}
		RedisUrl = "redis://localhost:6379/6"
	})
}
