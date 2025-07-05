gen:
	protoc \
  	--proto_path=proto \
  	--go_out=gen \
  	--go-grpc_out=gen \
  	proto/kvstore.proto