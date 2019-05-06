//go:generate docker run --rm -v $PWD/internal/protocol:/out -w /out znly/protoc --gogofaster_out=. -I=. gateway.proto login.proto
//go:generate docker run --rm -v $PWD/internal/protocol:/out -w /out znly/protoc --python_out=. -I=. gateway.proto login.proto

package protocol

func init() {}
