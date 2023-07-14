package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorageType struct {
	Client *minio.Client
}

var MinioStorage MinioStorageType

func NewMinio() *MinioStorageType {
	accessKeyId := os.Getenv("MINIO_ROOT_USER")
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")
	minioClient, err := minio.New("minio-s3:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Println("Failed to connect to MinIO S3. Likely missing env variables")
		log.Fatalln(err)
	}
	MinioStorage = MinioStorageType{
		Client: minioClient,
	}
	fmt.Println("Created MinIO Client!")
	return &MinioStorage
}

func (minioStorage *MinioStorageType) CreateDefaultBuckets() {
	fmt.Println("TODO: Create default buckets")
}
