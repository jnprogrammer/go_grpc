gen:
	protoc --proto_path=protocals protocals/*.proto --go_out=plugins=grpc:./pb

clean:
	rm pb/pb/*.go

run:
	go run main.go