// Package domain contains the business logic for handling products.
package domain

import "product/pb"

// ProductUsecase is an interface that defines the actions that can be performed on products.
type ProductUsecase interface {
	// Create creates a new product with the given details in the request.
	Create(req *pb.ProductCreateRequest) error
	Delete(req *pb.ProductDeleteRequest) (affected bool, err error)
}

// ProductRepository is an interface that defines the actions that can be performed on the repository
// for products.
type ProductRepository interface {
	// Save saves a product to the repository with the given details.
	Save(req *pb.ProductCreateRequest, generatedId string, createdTime int64) error
	Delete(req *pb.ProductDeleteRequest) (affected bool, err error)
}
