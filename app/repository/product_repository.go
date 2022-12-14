package repository

import (
	"context"
	"math"
	"product/domain"
	"product/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// func (pr *ProductRepository) {}

func (pr *ProductRepository) FindAll(req *pb.ProductFindAllRequest) (products *pb.ProductFindAllResponse, err error) {
	s := req.Search
	products = &pb.ProductFindAllResponse{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	findOpt := options.Find()

	if s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"product_id": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page := req.Page
	perPage := req.PerPage
	offset := page * perPage

	findOpt.SetSkip(offset)
	findOpt.SetLimit(perPage)

	if req.Sort == "desc" {
		findOpt.SetSort(bson.M{"product_id": -1})
	}

	cur, err := pr.products.Find(ctx, filter, findOpt)
	if err != nil {
		return
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var each domain.Product

		err = cur.Decode(&each)
		if err != nil {
			return
		}

		product := &pb.Product{
			ProductId:   each.ProductId,
			Name:        each.Name,
			Price:       each.Price,
			Duration:    each.Duration,
			Description: each.Description,
			CreatedAt:   each.CreatedAt,
			UpdatedAt:   each.UpdatedAt,
		}

		products.Products = append(products.Products, product)
	}

	rows, _ := pr.products.CountDocuments(ctx, filter)

	dataSize := int64(len(products.Products))
	products.Rows = rows
	products.Pages = int64(math.Ceil(float64(rows) / float64(perPage)))
	if dataSize < 1 {
		products.Pages = 0
	} else if perPage == 0 {
		products.Pages = 1
	}

	products.PerPage = perPage
	products.ActivePage = page + 1
	if dataSize < 1 {
		products.ActivePage = 0
	}
	products.Total = dataSize

	return
}

func (pr *ProductRepository) FindOne(req *pb.ProductFindOneRequest) (product *pb.Product, err error) {
	var result domain.Product

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"product_id": req.ProductId}

	err = pr.products.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return
	}

	product = &pb.Product{
		ProductId:   result.ProductId,
		Name:        result.Name,
		Price:       result.Price,
		Duration:    result.Duration,
		Description: result.Description,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}

	return
}

func (pr *ProductRepository) Update(req *pb.ProductUpdateRequest, updatedTime int64) (affected bool, err error) {
	affected = true

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"product_id": req.ProductId}
	detail := req.Detail

	payload := bson.M{
		"name":        detail.Name,
		"price":       detail.Price,
		"duration":    detail.Duration,
		"description": detail.Description,
		"updated_at":  updatedTime,
	}
	set := bson.M{"$set": payload}

	resp, err := pr.products.UpdateOne(ctx, filter, set)
	if resp.ModifiedCount < 1 {
		affected = false
	}

	return
}

func (pr *ProductRepository) Delete(req *pb.ProductDeleteRequest) (affected bool, err error) {
	affected = true

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"product_id": req.ProductId}

	resp, err := pr.products.DeleteOne(ctx, filter)
	if resp.DeletedCount < 1 {
		affected = false
	}

	return
}

// Save saves a new product in the database using the provided request data.
func (pr *ProductRepository) Save(req *pb.ProductCreateRequest, generatedId string, createdTime int64) (err error) {
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
	_, err = pr.products.InsertOne(ctx, payload)

	return
}
