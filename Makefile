.PHONY: deploy publish

test:
	go test ./...

test-v:
	go test -v ./...

build:
	go build

publish:
	./scripts/publish.sh

deploy:
	ssh ubuntu@builder.juhovuori.net bash -s < ./scripts/deploy.sh
