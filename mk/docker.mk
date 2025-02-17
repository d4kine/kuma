BUILD_DOCKER_IMAGES_DIR ?= $(BUILD_DIR)/docker-images-${GOARCH}
KUMA_VERSION ?= master

DOCKER_REGISTRY ?= kumahq
DOCKER_USERNAME ?=
DOCKER_API_KEY ?=

# If you want to build docker images with their own tags
IMAGE_ARCH_TAG_ENABLED ?=
define build_image
$(DOCKER_REGISTRY)/$(1):$(BUILD_INFO_VERSION)$(if $(IMAGE_ARCH_TAG_ENABLED),-${GOARCH},)
endef

export KUMA_BASE_IMAGE ?= kumahq/base-nossl-debian11:no-push
export KUMA_BASE_ROOT_IMAGE ?= kumahq/base-root-debian11:no-push
export KUMA_STATIC_IMAGE ?= kumahq/static-debian11:no-push
export KUMA_ENVOY_IMAGE ?= kumahq/envoy:no-push
export KUMA_CP_DOCKER_IMAGE ?= $(call build_image,kuma-cp)
export KUMA_DP_DOCKER_IMAGE ?= $(call build_image,kuma-dp)
export KUMACTL_DOCKER_IMAGE ?= $(call build_image,kumactl)
export KUMA_INIT_DOCKER_IMAGE ?= $(call build_image,kuma-init)
export KUMA_CNI_DOCKER_IMAGE ?= $(call build_image,kuma-cni)
export KUMA_UNIVERSAL_DOCKER_IMAGE ?= $(call build_image,kuma-universal)
KUMA_IMAGES ?= $(KUMA_BASE_ROOT_IMAGE) $(KUMA_BASE_IMAGE) $(KUMA_STATIC_IMAGE) $(KUMA_CP_DOCKER_IMAGE) $(KUMA_DP_DOCKER_IMAGE) $(KUMACTL_DOCKER_IMAGE) $(KUMA_INIT_DOCKER_IMAGE) $(KUMA_UNIVERSAL_DOCKER_IMAGE) $(KUMA_CNI_DOCKER_IMAGE)

IMAGES_TARGETS ?= images/release images/test
DOCKER_SAVE_TARGETS ?= docker/save/release docker/save/test
DOCKER_LOAD_TARGETS ?= docker/load/release docker/load/test

# Always use Docker BuildKit, see
# https://docs.docker.com/develop/develop-images/build_enhancements/
export DOCKER_BUILDKIT := 1

.PHONY: image/static
image/static: ## Dev: Rebuild `kuma-static` Docker image
	docker build -t $(KUMA_STATIC_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.static .

.PHONY: image/base
image/base: ## Dev: Rebuild `kuma-base` Docker image
	docker build -t $(KUMA_BASE_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.base .

.PHONY: image/base-root
image/base-root: ## Dev: Rebuild `kuma-base-root` Docker image
	docker build -t $(KUMA_BASE_ROOT_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.base-root .

.PHONY: image/envoy
image/envoy: build/artifacts-linux-${GOARCH}/envoy/envoy ## Dev: Rebuild `envoy` Docker image
	docker build -t $(KUMA_ENVOY_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} --build-arg ENVOY_VERSION=${ENVOY_VERSION} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.envoy .

.PHONY: images/supporting
images/supporting: image/static image/base image/base-root image/envoy

.PHONY: image/kuma-cp
image/kuma-cp: image/static build/kuma-cp/linux-${GOARCH} ## Dev: Rebuild `kuma-cp` Docker image
	docker build -t $(KUMA_CP_DOCKER_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.kuma-cp .

.PHONY: image/kuma-dp
image/kuma-dp: image/base image/envoy build/kuma-dp/linux-${GOARCH} build/coredns/linux-${GOARCH} build/artifacts-linux-${GOARCH}/envoy/envoy ## Dev: Rebuild `kuma-dp` Docker image
	docker build -t $(KUMA_DP_DOCKER_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.kuma-dp .

.PHONY: image/kumactl
image/kumactl: image/base build/kumactl/linux-${GOARCH} ## Dev: Rebuild `kumactl` Docker image
	docker build -t $(KUMACTL_DOCKER_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.kumactl .

.PHONY: image/kuma-init
image/kuma-init: build/kumactl/linux-${GOARCH} ## Dev: Rebuild `kuma-init` Docker image
	docker build -t $(KUMA_INIT_DOCKER_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.kuma-init .

.PHONY: image/kuma-cni
image/kuma-cni: image/base-root build/kuma-cni/linux-${GOARCH} build/install-cni/linux-${GOARCH}
	docker build -t $(KUMA_CNI_DOCKER_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --platform=linux/${GOARCH} -f $(TOOLS_DIR)/releases/dockerfiles/Dockerfile.kuma-cni .

.PHONY: image/kuma-universal
image/kuma-universal: build/kuma-cp/linux-${GOARCH} build/kuma-dp/linux-${GOARCH} build/artifacts-linux-${GOARCH}/envoy/envoy build/coredns/linux-${GOARCH} build/kumactl/linux-${GOARCH} build/test-server/linux-${GOARCH}
	docker build -t $(KUMA_UNIVERSAL_DOCKER_IMAGE) ${DOCKER_BUILD_ARGS} --build-arg ARCH=${GOARCH} --build-arg ENVOY_VERSION=${ENVOY_VERSION} --build-arg BASE_IMAGE_ARCH=${GOARCH} -f test/dockerfiles/Dockerfile.universal .

.PHONY: images
images: $(IMAGES_TARGETS) ## Dev: Rebuild release and test Docker images

.PHONY: images/release
images/release: image/kuma-cp image/kuma-dp image/kumactl image/kuma-init image/kuma-cni ## Dev: Rebuild release Docker images

.PHONY: images/test
images/test: image/kuma-universal ## Dev: Rebuild test Docker images

.PHONY: images/push
images/push: image/push/kuma-cp image/push/kuma-dp image/push/kumactl image/push/kuma-init image/kuma-cni

${BUILD_DOCKER_IMAGES_DIR}:
	mkdir -p ${BUILD_DOCKER_IMAGES_DIR}

.PHONY: docker/save
docker/save: $(DOCKER_SAVE_TARGETS)

.PHONY: docker/save/release
docker/save/release: docker/save/kuma-cp docker/save/kuma-dp docker/save/kumactl docker/save/kuma-init docker/save/kuma-cni

.PHONY: docker/save/test
docker/save/test: docker/save/kuma-universal

.PHONY: docker/save/kuma-cp
docker/save/kuma-cp: ${BUILD_DOCKER_IMAGES_DIR}
	docker save --output ${BUILD_DOCKER_IMAGES_DIR}/kuma-cp.tar $(KUMA_CP_DOCKER_IMAGE)

.PHONY: docker/save/kuma-dp
docker/save/kuma-dp: ${BUILD_DOCKER_IMAGES_DIR}
	docker save --output ${BUILD_DOCKER_IMAGES_DIR}/kuma-dp.tar $(KUMA_DP_DOCKER_IMAGE)

.PHONY: docker/save/kumactl
docker/save/kumactl: ${BUILD_DOCKER_IMAGES_DIR}
	docker save --output ${BUILD_DOCKER_IMAGES_DIR}/kumactl.tar $(KUMACTL_DOCKER_IMAGE)

.PHONY: docker/save/kuma-init
docker/save/kuma-init: ${BUILD_DOCKER_IMAGES_DIR}
	docker save --output ${BUILD_DOCKER_IMAGES_DIR}/kuma-init.tar $(KUMA_INIT_DOCKER_IMAGE)

.PHONY: docker/save/kuma-cni
docker/save/kuma-cni: ${BUILD_DOCKER_IMAGES_DIR}
	docker save --output ${BUILD_DOCKER_IMAGES_DIR}/kuma-cni.tar $(KUMA_CNI_DOCKER_IMAGE)

.PHONY: docker/save/kuma-universal
docker/save/kuma-universal: ${BUILD_DOCKER_IMAGES_DIR}
	docker save --output ${BUILD_DOCKER_IMAGES_DIR}/kuma-universal.tar $(KUMA_UNIVERSAL_DOCKER_IMAGE)

.PHONY: docker/load
docker/load: $(DOCKER_LOAD_TARGETS)

.PHONY: docker/load/release
docker/load/release: docker/load/kuma-cp docker/load/kuma-dp docker/load/kumactl docker/load/kuma-init docker/load/kuma-cni

.PHONY: docker/load/test
docker/load/test: docker/load/kuma-universal

.PHONY: docker/load/kuma-cp
docker/load/kuma-cp: ${BUILD_DOCKER_IMAGES_DIR}/kuma-cp.tar
	docker load --quiet --input ${BUILD_DOCKER_IMAGES_DIR}/kuma-cp.tar

.PHONY: docker/load/kuma-dp
docker/load/kuma-dp: ${BUILD_DOCKER_IMAGES_DIR}/kuma-dp.tar
	docker load --quiet --input ${BUILD_DOCKER_IMAGES_DIR}/kuma-dp.tar

.PHONY: docker/load/kumactl
docker/load/kumactl: ${BUILD_DOCKER_IMAGES_DIR}/kumactl.tar
	docker load --quiet --input ${BUILD_DOCKER_IMAGES_DIR}/kumactl.tar

.PHONY: docker/load/kuma-init
docker/load/kuma-init: ${BUILD_DOCKER_IMAGES_DIR}/kuma-init.tar
	docker load --quiet --input ${BUILD_DOCKER_IMAGES_DIR}/kuma-init.tar

.PHONY: docker/load/kuma-cni
docker/load/kuma-cni: ${BUILD_DOCKER_IMAGES_DIR}/kuma-cni.tar
	docker load --quiet --input ${BUILD_DOCKER_IMAGES_DIR}/kuma-cni.tar

.PHONY: docker/load/kuma-universal
docker/load/kuma-universal: ${BUILD_DOCKER_IMAGES_DIR}/kuma-universal.tar
	docker load --quiet --input ${BUILD_DOCKER_IMAGES_DIR}/kuma-universal.tar

.PHONY: docker/tag/kuma-cp
docker/tag/kuma-cp:
	docker tag $(KUMA_CP_DOCKER_IMAGE) $(DOCKER_REGISTRY)/kuma-cp:$(KUMA_VERSION)-${GOARCH}

.PHONY: docker/tag/kuma-dp
docker/tag/kuma-dp:
	docker tag $(KUMA_DP_DOCKER_IMAGE) $(DOCKER_REGISTRY)/kuma-dp:$(KUMA_VERSION)-${GOARCH}

.PHONY: docker/tag/kumactl
docker/tag/kumactl:
	docker tag $(KUMACTL_DOCKER_IMAGE) $(DOCKER_REGISTRY)/kumactl:$(KUMA_VERSION)-${GOARCH}

.PHONY: docker/tag/kuma-init
docker/tag/kuma-init:
	docker tag $(KUMA_INIT_DOCKER_IMAGE) $(DOCKER_REGISTRY)/kuma-init:$(KUMA_VERSION)-${GOARCH}

.PHONY: docker/tag/kuma-universal
docker/tag/kuma-universal:
	docker tag $(KUMA_UNIVERSAL_DOCKER_IMAGE) $(DOCKER_REGISTRY)/kuma-universal:$(KUMA_VERSION)-${GOARCH}

.PHONY: docker/purge
docker/purge: ## Dev: Remove all Docker containers, images, networks and volumes
	for c in `docker ps -q`; do docker kill $$c; done
	docker system prune --all --volumes --force

.PHONY: image/push/kuma-cp
image/push/kuma-cp: image/kuma-cp
	docker push $(KUMA_CP_DOCKER_IMAGE)

.PHONY: image/push/kuma-dp
image/push/kuma-dp: image/kuma-dp
	docker push $(KUMA_DP_DOCKER_IMAGE)

.PHONY: image/push/kumactl
image/push/kumactl: image/kumactl
	docker push $(KUMACTL_DOCKER_IMAGE)

.PHONY: image/push/kuma-init
image/push/kuma-init: image/kuma-init
	docker push $(KUMA_INIT_DOCKER_IMAGE)
