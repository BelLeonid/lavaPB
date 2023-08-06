# lavaPB

# Autogenerate PB files:
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    osmosisTest/osmosis.proto

# Test Server with grpcurl

grpcurl -plaintext localhost:3333 list

grpcurl -plaintext localhost:3333 cosmos.base.tendermint.v1beta1.Service.GetLatestBlock

grpcurl -plaintext -d '{"height": 10308545}' localhost:3333 cosmos.base.tendermint.v1beta1.Service.GetBlockByHeight

# Test server's endpoint with grpcurl

grpcurl -plaintext grpc.osmosis.zone:9090 list

grpcurl -plaintext grpc.osmosis.zone:9090 cosmos.base.tendermint.v1beta1.Service.GetLatestBlock

grpcurl -plaintext -d '{"height": 10308545}' grpc.osmosis.zone:9090 cosmos.base.tendermint.v1beta1.Service.GetBlockByHeight 