.PHONY: deploy publish build test test-v

test:
	go test ./...

test-v:
	go test -v ./...

build:
	./scripts/version.sh >version.json
	go build

publish:
	./scripts/publish.sh

deploy:
	ssh ubuntu@builder.juhovuori.net bash -s < ./scripts/deploy.sh
