package cloudstorage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
)

func (si *StorageBucketUploaderImpl) Uploader(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	bucketName := si.Config.BucketName
	objectName := file.Filename

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	srcFile, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %s", err.Error())
	}
	
	defer srcFile.Close()

	if _, err := io.Copy(wc, srcFile); err != nil {
		return "", fmt.Errorf("failed to copy file to storage: %s", err.Error())
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close storage: %s", err.Error())
	}

	imageUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)

	return imageUrl, nil
}
