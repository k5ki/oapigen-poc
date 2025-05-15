.PHONY: gen-api
gen-api:
	mkdir -p ./.out/api
	@# bundle openapi spec
	docker run \
		-v ./api/definition:/spec \
		-v ./.out:/tmp/out \
		--rm \
		redocly/cli bundle --output /tmp/out/api/openapi_bundled.yaml openapi.yaml
	@# generate code
	oapi-codegen -config ./api/config/model.yaml ./.out/api/openapi_bundled.yaml
	oapi-codegen -config ./api/config/server.yaml ./.out/api/openapi_bundled.yaml

.PHONY: tools/install
tools/install:
	@# openapi
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
