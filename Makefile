install-tools:
	npm install -D

openapi-bundle:
	npm run spec:bundle

openapi-mock:
	npm run spec:mock

openapi-lint:
	npm run spec:lint

openapi-generate-ts:
	npm run gen:ts

openapi-generate-go:
	cd api && go tool oapi-codegen --config=../openapi/oapi-codegen/client.yaml ../openapi/dist/openapi.bundled.yaml
	cd api && go tool oapi-codegen --config=../openapi/oapi-codegen/models.yaml ../openapi/dist/openapi.bundled.yaml
	cd api && go tool oapi-codegen --config=../openapi/oapi-codegen/server.yaml ../openapi/dist/openapi.bundled.yaml
	cd api && go tool oapi-codegen --config=../openapi/oapi-codegen/server_urls.yaml ../openapi/dist/openapi.bundled.yaml

openapi-generate: openapi-bundle openapi-generate-ts openapi-generate-go

openapi-docs:
	npm run spec:docs

openapi-start:
	@$(MAKE) -j2 openapi-mock openapi-docs
