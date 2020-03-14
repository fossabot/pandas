IMPORT_PATH = github.com/cloustone/pandas
# V := 1 # When V is set, print cmd and build progress.
Q := $(if $V,,@)

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

# Space separated patterns of packages to skip in list, test, format.
DOCKER_NAMESPACE := cloustone

# Space separated patterns of packages to skip in list, test, format.
IGNORED_PACKAGES := /vendor/

UNAME = $(shell uname)
DOCKER_REPO = docker.io
IMAGES = apimachinery dmms pms rulechain headmast lbs  
IMAGE_NAME_PREFIX := pandas-
IMAGE_DIR := $(IMAGE_NAME)
ifeq ($(IMAGE_NAME),bridge)
    IMAGE_DIR := edge/$(IMAGE_NAME)
else ifneq (,$(filter $(IMAGE_NAME), apimachinery dmms pms rulechain headmast lbs shiro))
    IMAGE_DIR := cmd/$(IMAGE_NAME)
else ifeq ($(IMAGE_NAME),cabinet)
    IMAGE_DIR := security/$(IMAGE_NAME)
endif

GCFLAGS  := -gcflags="-N -l"

.PHONY: all
all: build

.PHONY: docker
docker: export GOOS=linux
docker: $(addprefix docker-build-, $(IMAGES)) 
	docker images | grep '<none>' | awk '{print $3}' | xargs docker rmi
	@echo "docker building completed!" 

# Docker build targets
$(addprefix docker-build-, $(IMAGES)): docker-build-%: %
	@IMAGE_NAME=$< make .docker-build

.docker-build:
	@echo building $(IMAGE_NAME_PREFIX)$(IMAGE_NAME) image ...
	@if [ ! -d "$(IMAGE_DIR)/bin/" ]; then mkdir $(IMAGE_DIR)/bin/ ; fi
	# @cp scripts/dockerize $(IMAGE_DIR)/bin/
#	@if [ "$(UNAME)" = "Linux" ]; then cp bin/$(IMAGE_NAME) $(IMAGE_DIR)/bin/main ; fi
#	@if [ "$(UNAME)" = "Darwin" ]; then cp bin/linux_amd64/$(IMAGE_NAME) $(IMAGE_DIR)/bin/main ; fi
	cp bin/$(IMAGE_NAME) $(IMAGE_DIR)/bin/main
	@full_img_name=$(IMAGE_NAME_PREFIX)$(IMAGE_NAME); \
		cd ./$(IMAGE_DIR)/ && \
			docker build -t $(DOCKER_REPO)/$(DOCKER_NAMESPACE)/$$full_img_name .
	@rm -rf $(IMAGE_DIR)/bin
	@"./scripts/push.sh" $(IMAGE_NAME)
	# @kubectl delete pod $$(kubectl get pod -n pandas | grep $(IMAGE_NAME) | awk '{print $$1}') -n pandas 

.PHONY: deploy
deploy:
	@cd deploy/helm && helm install .

.PHONY: upgrade
upgrade:
	@existing=$$(helm list | grep pandas | awk '{print $$1}' | head -n 1); \
		(if [ ! -z "$$existing" ]; then echo "Upgrade the stack via helm. This may take a while."; helm upgrade "$$existing"; echo "The stack has been upgraded."; fi) > /dev/null;

.PHONY: undeploy
undeploy:
	@existing=$$(helm list | grep pandas | awk '{print $$1}' | head -n 1); \
		(if [ ! -z "$$existing" ]; then echo "Undeploying the stack via helm. This may take a while."; helm del --purge "$$existing"; echo "The stack has been undeployed."; fi) > /dev/null;

.PHONY: all
all: build

.PHONY: build
build: apimachinery  dmms  pms rulechain lbs headmast  shiro

.PHONY: apimachinery 
apimachinery: 
	@echo "building api server (apimachinery)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/apimachinery 

.PHONY: dmms 
dmms: cmd/dmms 
	@echo "building device management server (dmms)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/dmms

.PHONY: pms 
pms: cmd/pms 
	@echo "building project management server (pms)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/pms

.PHONY: rulechain 
rulechain: cmd/rulechain
	@echo "building rulechain server (rulechain)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/rulechain

.PHONY: lbs 
lbs: cmd/lbs
	@echo "building location based service (lbs)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/lbs

.PHONY: headmast 
headmast: cmd/headmast
	@echo "building headmast service (headmast)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/headmast

.PHONY: shiro 
shiro: cmd/shiro
	@echo "building unified user manager center service (shiro)..."
	$Q CGO_ENABLED=0 go build -o bin/$@ $(GCFLAGS) $(if $V,-v) $(VERSION_FLAGS) $(IMPORT_PATH)/cmd/shiro


.PHONY: test
test: 
	$Q go test  ./...





