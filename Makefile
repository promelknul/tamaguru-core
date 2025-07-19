bootstrap: go-mod node-mod openapi gen build
go-mod:
	go mod tidy ./...
node-mod:
	cd web && npm ci --silent
openapi:
	npx @openapitools/openapi-generator-cli validate -i web/auth/spec/openapi.yaml
gen:
	docker run --rm -v $$PWD:/local openapitools/openapi-generator-cli:v7.14.0 generate -i /local/web/auth/spec/openapi.yaml -g go -o /local/core/vault/handler --package-name handler
	docker run --rm -v $$PWD:/local openapitools/openapi-generator-cli:v7.14.0 generate -i /local/web/auth/spec/openapi.yaml -g typescript-axios -o /local/web/src/api
	npx @openapitools/openapi-generator-cli generate -i web/auth/spec/openapi.yaml -g go -o core/vault/handler --package-name handler
	npx @openapitools/openapi-generator-cli generate -i web/auth/spec/openapi.yaml -g typescript-axios -o web/src/api
build:
	npx turbo run build
test:
	go test ./...
run:
	air & npm --prefix web run dev
