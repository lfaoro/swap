.PHONY: build proto release clean audit sha tools

FLAGS=-trimpath -buildvcs=false -tags='netgo,osusergo,static_build'
LDFLAGS=-ldflags='-s -w -extldflags "-static" -X main.APIURL=api.swapcli.com:443'

all: proto build

build:
	@go mod tidy
	CGO_ENABLED=0 go build ${FLAGS} ${LDFLAGS} -o ./bin/ ./cmd/swap

pre:
	go mod tidy
	go fmt ./... && go vet ./...
	# gosec ./cmd/...
	# gosec ./app/...
	
deployapi: pre buildapi docker fly fly-deploy
buildapi:
	$(eval NAME:=swap_api)
	@go mod tidy
	CGO_ENABLED=0 go build ${FLAGS} ${LDFLAGS} -o bin/${NAME} ./cmd/${NAME}
fly: 
	flyctl auth docker
	docker push registry.fly.io/troca-api:latest
fly-deploy:
	flyctl deploy --detach --image registry.fly.io/troca-api:latest --config fly.toml
fly-secrets:
	flyctl secrets import < cmd/swap_api/.env
	flyctl secrets deploy
docker:
	docker build -t troca-api:latest .
	docker tag troca-api:latest registry.fly.io/troca-api:latest

release: proto secure-build sha
release-client:
	goreleaser release --clean --skip=announce,validate
secure-build:
	@go mod tidy
	GOARCH=amd64 GOOS=linux garble build ${FLAGS} ${LDFLAGS} -o ./bin/ ./cmd/swap
	strip ./bin/swap ||:
	upx -9 ./bin/swap ||:

buildnix:
	nix flake init
	nix build

proto:
	rm -rf proto/go
	buf generate
	
clean:
	rm -rf gen/* bin/*

audit:
	gosec ./cmd/...
	gosec ./app/...

sha:
	shasum -a256 ./bin/swap | tee ./bin/swap.sum

loc:
	find . -name "*.go" -not -path "*/src/*" -not -path "*/gen/*" -not -path "*/vendor/*" -not -path "*/test/*" | xargs wc -l

tools:
	go install github.com/air-verse/air@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install mvdan.cc/garble@latest
