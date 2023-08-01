package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// MinIO server endpoint, access key, and secret key
	endpoint := "your-minio-endpoint"       // Replace with your MinIO server endpoint
	accessKeyID := "your-access-key-id"     // Replace with your access key ID
	secretAccessKey := "your-secret-key"    // Replace with your secret access key
	useSSL := false                         // Set to true if your MinIO server uses SSL

	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	// Create a new bucket (optional)
	bucketName := "your-bucket-name" // Replace with the desired bucket name
	location := "your-location"     // Replace with your desired bucket location
	ctx := context.Background()

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists.", bucketName)
		} else {
			log.Fatalf("Error creating bucket: %v", err)
		}
	} else {
		log.Printf("Bucket '%s' created successfully in location '%s'.", bucketName, location)
	}

	// Upload an object to the bucket
	objectName := "your-object-name"         // Replace with the desired object name
	filePath := "path/to/your/local/file"    // Replace with the local file path you want to upload

	// Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Get the file stats to get its size
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file stats: %v", err)
	}

	// Set object options with content type and size
	opts := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
		UserMetadata: map[string]string{
			"my-key": "my-value",
		},
		// The size can be determined from the file's stat
		// The content size must be specified for objects > 5MB to enable multipart upload.
		ContentLength: stat.Size(),
	}

	// Upload the object with the specified options
	n, err := minioClient.PutObject(ctx, bucketName, objectName, file, opts)
	if err != nil {
		log.Fatalf("Error uploading object: %v", err)
	}

	log.Printf("Uploaded %s of size %d successfully.", objectName, n)
}

