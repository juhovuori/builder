.PHONY: deploy publish build test test-v version

test:
	go test ./...

test-v:
	go test -v ./...

version:
	./scripts/version.sh >version.json

build: version
	go build

publish: version
	./scripts/publish.sh

deploy:
	ssh ubuntu@builder.juhovuori.net bash -s < ./scripts/deploy.sh
