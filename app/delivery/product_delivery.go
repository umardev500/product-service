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

//func (pd *ProductDelivery) Delete(ctx context.Context, req *pb.) (res *pb., err error) {}

func (pd *ProductDelivery) Update(ctx context.Context, req *pb.ProductUpdateRequest) (res *pb.OperationResponse, err error) {
	affected, err := pd.usecase.Update(req)
	res = &pb.OperationResponse{IsAffected: affected}

	return
}

func (pd *ProductDelivery) Delete(ctx context.Context, req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error) {
	affected, err := pd.usecase.Delete(req)
	res = &pb.OperationResponse{IsAffected: affected}

	return
}

func (pd *ProductDelivery) Create(ctx context.Context, req *pb.ProductCreateRequest) (res *pb.Empty, err error) {
	err = pd.usecase.Create(req)
	res = &pb.Empty{}

	return
}
