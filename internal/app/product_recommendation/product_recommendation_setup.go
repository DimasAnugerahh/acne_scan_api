package productrecommendation

import (
	handlerPKG "acne-scan-api/internal/app/product_recommendation/handlers"
	repositoryPKG "acne-scan-api/internal/app/product_recommendation/repository"
	routesPKG "acne-scan-api/internal/app/product_recommendation/routes"
	servicePKG "acne-scan-api/internal/app/product_recommendation/service"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"

	"database/sql"

	"github.com/go-playground/validator"
)

func ProductRecommendationSetup(db *sql.DB, validate *validator.Validate,bs cloudstorage.StorageBucketUploader) routesPKG.ProductRecommendationRoutes {

	productRecommendationRepository := repositoryPKG.NewProProductRecommendation(db)
	productRecommendationService := servicePKG.NewProductRecommendationService(productRecommendationRepository, validate,bs)
	productRecommendationHanlder := handlerPKG.NewProductRecommendationHandler(productRecommendationService)
	productRecommendationRoutes := routesPKG.NewProductRecommendationRoutes(productRecommendationHanlder)

	return productRecommendationRoutes

}
