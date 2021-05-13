module github.com/eteu-technologies/near-api-go

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/eteu-technologies/near-rpc-go v0.0.0-00010101000000-000000000000
	github.com/minio/sha256-simd v1.0.0
	github.com/mr-tron/base58 v1.2.0
	github.com/near/borsh-go v0.3.0
)

replace github.com/eteu-technologies/near-rpc-go => ./near-rpc-go
