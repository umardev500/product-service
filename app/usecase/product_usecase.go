package usecase

import (
	"product/domain"
	"product/pb"
	"strconv"
	"time"
)

// ProductUsecase defines the use case for managing products.
type ProductUsecase struct {
	// repository is the underlying repository for storing products.
	repository domain.ProductRepository
}

// NewProductUsecase creates a new ProductUsecase with the given repository.
func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &ProductUsecase{
		repository: repo,
	}
}

// Create creates a new product with the given request and returns a boolean indicating if the product was saved successfully, and an error if one occurred.
func (pu *ProductUsecase) Create(req *pb.ProductCreateRequest) (affected bool, err error) {
	// Get the current time in UTC
	t := time.Now().UTC()

	generatedId := strconv.Itoa(int(t.Unix()))
	createdTime := t.Unix()

	affected, err = pu.repository.Save(req, generatedId, createdTime)

	return
}
