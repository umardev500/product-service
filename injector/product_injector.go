package injector

import (
	"product/app/delivery"
	"product/app/repository"
	"product/app/usecase"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductInjector(db *mongo.Database) *delivery.ProductDelivery {
	repo := repository.NewProductRepository(db)
	usecase := usecase.NewProductUsecase(repo)

	return delivery.NewProductDelivery(usecase)
}
