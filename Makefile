# Include External Env and export system Env
export

DOCKER_IMG=go-health-id:oss
VERSION=v0.0.0+unknown

# CI automate version set
ifeq (${CI}, true)
	VERSION=${CI_COMMIT_REF_NAME}-build${CI_PIPELINE_ID}
	DOCKER_IMG=${CI_REGISTRY_IMAGE}:${CI_COMMIT_REF_NAME}
endif

.PHONY: version clean local compile docker push dev

version:
	echo ${VERSION}

clean:
	rm -rf ./out

local:
	docker build --target=server --build-arg="APP_VERSION=beta" -t ${DOCKER_IMG} .

compile:
	go build -ldflags "-X 'main.Version=${VERSION}'" -o ./out/server ./cmd/server/main.go
	./out/server version

docker:
	docker build --target=server --build-arg="APP_VERSION=${VERSION}" -t ${DOCKER_IMG} .

push:
	docker push ${DOCKER_IMG}

dev:
	go run -ldflags "-X 'main.Version=${VERSION}'" cmd/server/main.go start