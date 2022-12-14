package repository

import (
	"context"
	"product/domain"
	"product/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db       *mongo.Database
	products *mongo.Collection
}

func NewProductRepository(db *mongo.Database) domain.ProductRepository {
	return &ProductRepository{
		db:       db,
		products: db.Collection("products"),
	}
}

// Save saves a new product in the database using the provided request data.
func (pr *ProductRepository) Save(req *pb.ProductCreateRequest, generatedId string, createdTime int64) (affected bool, err error) {
	// Assume the operation will be successful
	affected = true

	// Create a context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create the product document to insert
	payload := bson.D{
		{Key: "product_id", Value: generatedId},
		{Key: "name", Value: req.Name},
		{Key: "price", Value: req.Price},
		{Key: "duration", Value: req.Duration},
		{Key: "description", Value: req.Description},
		{Key: "created_at", Value: createdTime},
	}

	// Insert the product document
	resp, err := pr.products.InsertOne(ctx, payload)
	if err != nil {
		return
	}

	// If the inserted ID is nil, then the insertion failed
	if resp.InsertedID == nil {
		affected = false
	}

	return
}
