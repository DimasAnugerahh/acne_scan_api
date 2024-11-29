package cloudstorage

import (
	"acne-scan-api/configs"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type StorageBucketUploader interface{
	Uploader(c *fiber.Ctx, file *multipart.FileHeader) (string,error)
}

type StorageBucketUploaderImpl struct{
	Config *configs.StorageBucketConfig
}

func NewStorageBucketUploader(configs *configs.StorageBucketConfig)StorageBucketUploader{
	return &StorageBucketUploaderImpl{
		Config: configs,
	}
}
