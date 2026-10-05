.PHONY: install-tools

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

openapi-docs:
	npm run spec:docs
	npx serve openapi/dist -l 8081

openapi-start:
	@$(MAKE) -j2 openapi-mock openapi-docs
