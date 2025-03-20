.PHONY: build proto release clean audit sha tools

FLAGS=-trimpath -buildvcs=false -tags='netgo,osusergo,static_build'
LDFLAGS=-ldflags='-s -w -extldflags "-static"'

all: build

build:
	@go mod tidy
	CGO_ENABLED=0 go build ${FLAGS} ${LDFLAGS} -o ./bin/ ./cmd/swap

pre: audit
	go mod tidy
	go fmt ./... && go vet ./...
	
release: proto release-client
release-client:
	goreleaser release --clean --skip=announce,validate
release-dev:
	GORELEASER_CURRENT_TAG="v0.0.1" goreleaser release --clean --skip=announce,validate --snapshot --skip-publish

buildnix:
	nix flake init
	nix build

proto:
	rm -rf proto/go
	buf generate
	
clean:
	rm -rf gen/* bin/*

audit:
	gosec ./app ./cmd/swap

sha:
	shasum -a256 ./bin/swap | tee ./bin/swap.sum

update:
	go get -u ./cmd/swap

loc:
	find . -name "*.go" -not -path "*/src/*" -not -path "*/gen/*" -not -path "*/vendor/*" -not -path "*/test/*" | xargs wc -l

upgrade:
	go get -u ./cmd/swap

tools:
	go install github.com/air-verse/air@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

backup: clean
	tar -czvf ../swapcli-$(shell date +%Y%m%d).tgz --exclude='.git' .
