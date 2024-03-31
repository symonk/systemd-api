MODULE = $(shell go list -m)

.PHONY: generate build test lint docker compose compose-down

generate:
		go generate ./...

build:
		go build -a -o systemd-api $(MODULE)/cmd/server

test:
		go clean -testcache
		go test ./... -v

lint:
		gofmt -l .

docker:
		docker build -f cmd/server/Dockerfile -t systemd-api/ .

compose.%:
		$(eval CMD = ${subst compose.,,$(@)})
		tools/script/compose.sh $(CMD)

