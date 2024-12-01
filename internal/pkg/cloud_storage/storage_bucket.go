package cloudstorage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

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

	buffer := make([]byte, 512)
	_, err = srcFile.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error reading file: %s", err)
	}

	// Detect the MIME type
	filetype := http.DetectContentType(buffer)

	if (filetype != "image/jpeg") && (filetype != "image/png") {
		return "",fmt.Errorf("jenis file tidak valid")
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
