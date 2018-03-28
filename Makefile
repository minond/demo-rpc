.PHONY: proto
proto:
	protoc -I src/ src/friends.proto --go_out=plugins=grpc:src/friends
