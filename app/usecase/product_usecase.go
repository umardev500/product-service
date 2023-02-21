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
func (pu *ProductUsecase) Create(req *pb.ProductCreateRequest) error {
	// Get the current time in UTC
	t := time.Now().UTC()

	generatedId := strconv.Itoa(int(t.UnixNano()))
	createdTime := t.Unix()

	return pu.repository.Save(req, generatedId, createdTime)
}

func (pu *ProductUsecase) FindOne(req *pb.ProductFindOneRequest) (*pb.ProductFindOneResponse, error) {
	return pu.repository.FindOne(req)
}

func (pu *ProductUsecase) FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error) {
	res, err = pu.repository.FindAll(req)

	return
}

func (pu *ProductUsecase) Update(req *pb.ProductUpdateRequest) (bool, error) {
	updatedTime := time.Now().Unix()
	return pu.repository.Update(req, updatedTime)
}

func (pu *ProductUsecase) Delete(req *pb.ProductDeleteRequest) (bool, error) {
	return pu.repository.Delete(req)
}
