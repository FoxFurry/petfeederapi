

generate:
	docker run \
		-e GO_POST_PROCESS_FILE="/usr/local/go/bin/gofmt -w" --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
		-g go-server \
		-i /local/petfeederapi.yaml \
		--global-property=generateAliasAsModel \
		--additional-properties=outputAsLibrary=true,onlyInterfaces=true,router=chi,sourceFolder=internal/api \
		--git-repo-id="petfeederapi" \
		--git-user-id="FoxFurry" \
		-o /local/.

run:
	go run main.go serve

build:
	go build -o bin/petfeeder

test:
	go test ./...

migrate:
	./migrations/migrate.sh

totallines:
	find . -name '*.go' | xargs wc -l