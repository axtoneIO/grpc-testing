build:	generate
	go build .

generate:
	protoc rocket/v1/*.proto --go_out=plugins=grpc:.

clean:
	rm rocket/v1/*.pb.go