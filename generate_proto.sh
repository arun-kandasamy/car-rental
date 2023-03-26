protoc --proto_path=proto \
    --go_out=proto/generated/policy \
    --go-grpc_out=proto/generated/policy  \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
   proto/policy.proto


protoc --proto_path=proto \
    --go_out=proto/generated/carrental \
    --go-grpc_out=proto/generated/carrental \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
   proto/carrental.proto