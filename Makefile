run:
	go run main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pb/*.proto

clean: rm pb/*.pb.go

# grpcurl execution
create:
	grpcurl --plaintext -d '{"name": "Trial 30 days", "price": 0, "duration": 30, "description": "This is the description"}' localhost:5010 ProductService.Create

delete:
	grpcurl --plaintext -d '{"product_id": "1671033950"}' localhost:5010 ProductService.Delete

update:
	grpcurl --plaintext -d '{"product_id": "1671034841", "detail": {"name": "The name of product updated", "price": 0, "duration": 30, "description": "The updated description"}}' localhost:5010 ProductService.Update

findone:
	grpcurl --plaintext -d '{"product_id": "1671034841"}' localhost:5010 ProductService.FindOne

findall:
	grpcurl --plaintext localhost:5010 ProductService.FindAll