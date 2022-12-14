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

// func (pu *ProductUsecase) {}

func (pu *ProductUsecase) Update(req *pb.ProductUpdateRequest) (affected bool, err error) {
	t := time.Now()
	updatedTime := t.Unix()

	affected, err = pu.repository.Update(req, updatedTime)

	return
}

func (pu *ProductUsecase) Delete(req *pb.ProductDeleteRequest) (affected bool, err error) {
	affected, err = pu.repository.Delete(req)

	return
}

// Create creates a new product with the given request and returns a boolean indicating if the product was saved successfully, and an error if one occurred.
func (pu *ProductUsecase) Create(req *pb.ProductCreateRequest) (err error) {
	// Get the current time in UTC
	t := time.Now().UTC()

	generatedId := strconv.Itoa(int(t.Unix()))
	createdTime := t.Unix()

	err = pu.repository.Save(req, generatedId, createdTime)

	return
}
