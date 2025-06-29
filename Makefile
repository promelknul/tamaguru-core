bootstrap: go-mod node-mod openapi gen build
go-mod:
	go mod tidy ./...
node-mod:
	cd web && npm ci --silent
openapi:
	npx @openapitools/openapi-generator-cli validate -i web/auth/spec/openapi.yaml
gen:
	npx @openapitools/openapi-generator-cli generate -i web/auth/spec/openapi.yaml -g go -o core/vault/handler --package-name handler
	npx @openapitools/openapi-generator-cli generate -i web/auth/spec/openapi.yaml -g typescript-axios -o web/src/api
build:
	npx turbo run build
test:
	go test ./...
run:
	air & npm --prefix web run dev
