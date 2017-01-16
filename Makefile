.PHONY: deploy publish build test test-v version

test:
	go test ./...

test-v:
	go test -v ./...

version:
	go generate ./version

build: version
	go build

publish: version
	./scripts/publish.sh

deploy:
	ssh ubuntu@builder.juhovuori.net bash -s < ./scripts/deploy.sh
