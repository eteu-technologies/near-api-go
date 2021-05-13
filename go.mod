module github.com/eteu-technologies/near-api-go

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/eteu-technologies/borsh-go v0.3.1 // indirect
	github.com/eteu-technologies/near-rpc-go v0.0.0-00010101000000-000000000000
	github.com/minio/sha256-simd v1.0.0
	github.com/mr-tron/base58 v1.2.0
	lukechampine.com/uint128 v1.1.1 // indirect
)

replace github.com/eteu-technologies/near-rpc-go => ./near-rpc-go
