export PATH := ./bin:$(PATH)
export GO111MODULE := on

test:
	go test ./... -coverprofile coverage.out

cover: test
	go tool cover -func coverage.out
	go tool cover -html coverage.out

lint:
	golangci-lint run

generete_doc:
	go run docs/generate.go

build-dev:
	GO111MODULE=on go build -gcflags "all=-N -l" -ldflags "-X github.com/tfournier/gvm/cmd.version=0.0.0-dev" -o /usr/local/bin/gvm

release:
	GO111MODULE=on goreleaser release --rm-dist --debug

dry-run:
	GO111MODULE=on goreleaser release --rm-dist --debug --skip-publish
