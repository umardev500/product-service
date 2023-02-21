// Package domain contains the business logic for handling products.
package domain

import "product/pb"

type Product struct {
	ProductId   string `bson:"product_id"`
	Name        string `bson:"name"`
	Price       int64  `bson:"price"`
	Duration    int64  `bson:"duration"`
	Description string `bson:"description"`
	CreatedAt   int64  `bson:"created_at"`
	UpdatedAt   int64  `bson:"updated_at"`
}

// ProductUsecase is an interface that defines the actions that can be performed on products.
type ProductUsecase interface {
	// Create creates a new product with the given details in the request.
	Create(req *pb.ProductCreateRequest) error
	Delete(req *pb.ProductDeleteRequest) (affected bool, err error)
	Update(req *pb.ProductUpdateRequest) (affected bool, err error)
	FindOne(req *pb.ProductFindOneRequest) (product *pb.ProductFindOneResponse, err error)
	FindAll(req *pb.ProductFindAllRequest) (products *pb.ProductFindAllResponse, err error)
}

// ProductRepository is an interface that defines the actions that can be performed on the repository
// for products.
type ProductRepository interface {
	// Save saves a product to the repository with the given details.
	Save(req *pb.ProductCreateRequest, generatedId string, createdTime int64) error
	Delete(req *pb.ProductDeleteRequest) (affected bool, err error)
	Update(req *pb.ProductUpdateRequest, updatedTime int64) (affected bool, err error)
	FindOne(req *pb.ProductFindOneRequest) (product *pb.ProductFindOneResponse, err error)
	FindAll(req *pb.ProductFindAllRequest) (products *pb.ProductFindAllResponse, err error)
}
