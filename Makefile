run:
	go run main.go

proto:
	protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	pb/*.proto

clean: rm pb/*.pb.go

# grpcurl execution
create:
	grpcurl --plaintext -d '{"name": "Trial 30 days", "price": 0, "duration": 30, "description": "This is the description"}' localhost:5010 ProductService.Create

delete:
	grpcurl --plaintext -d '{"product_id": "1671033950"}' localhost:5010 ProductService.Delete