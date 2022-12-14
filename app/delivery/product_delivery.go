package delivery

import (
	"context"
	"product/domain"
	"product/pb"
)

type ProductDelivery struct {
	usecase domain.ProductUsecase
	pb.UnimplementedProductServiceServer
}

func NewProductDelivery(usecase domain.ProductUsecase) *ProductDelivery {
	return &ProductDelivery{
		usecase: usecase,
	}
}

func (pd *ProductDelivery) Create(ctx context.Context, req *pb.ProductCreateRequest) (*pb.OperationResponse, error) {
	affected, err := pd.usecase.Create(req)

	return &pb.OperationResponse{IsAffected: affected}, err
}
