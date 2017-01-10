.PHONY: deploy publish

test:
	go test ./...

build:
	go build

publish:
	./scripts/publish.sh

deploy:
	ssh ubuntu@builder.juhovuori.net bash -s < ./scripts/deploy.sh
