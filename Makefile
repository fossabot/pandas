IMPORTPATH = github.com/cloustone/pandas
# V := 1 # When V is set, print cmd and build progress.
Q := $(if $V,,@)

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

# Space separated patterns of packages to skip in list, test, format.
DOCKER_NAMESPACE := cloustone
DOCKERFILE_ADDRESS := docker/Dockerfile
API_DOCKERFILE_ADDRESS := cmd/apimachinery/Dockerfile

.PHONY: all
all: build

.PHONY: build
build: docker apimachinery  dmms  pms rulechain lbs headmast 

.PHONY: docker
docker: 
	@echo "Creating docker image (docker)..."
	$Q docker build -f ${DOCKERFILE_ADDRESS} -t ${DOCKER_NAMESPACE} .
	$Q docker build -f cmd/apimachinery/Dockerfile -t api .
	$Q docker build -f cmd/dmms/Dockerfile -t dmms .
	$Q docker build -f cmd/headmast/Dockerfile -t headmast .
	$Q docker build -f cmd/lbs/Dockerfile -t lbs .
	$Q docker build -f cmd/pms/Dockerfile -t pms .
	$Q docker build -f cmd/rulechain/Dockerfile -t rulechain .
#    docker run -it --name zhmm -P cloustone

.PHONY: apimachinery 
apimachinery: 
	@echo "building api server (apimachinery)..."
	$Q CGO_ENABLED=0 go build -v -o bin/apimachinery $(IMPORTPATH)/cmd/apimachinery

.PHONY: dmms 
dmms: cmd/dmms 
	@echo "building device management server (dmms)..."
	$Q CGO_ENABLED=0 go build -o bin/dmms $(IMPORTPATH)/cmd/dmms

.PHONY: pms 
pms: cmd/pms 
	@echo "building project management server (pms)..."
	$Q CGO_ENABLED=0 go build -o bin/pms $(IMPORTPATH)/cmd/pms

.PHONY: rulechain 
rulechain: cmd/rulechain
	@echo "building rulechain server (rulechain)..."
	$Q CGO_ENABLED=0 go build -o bin/rulechain $(IMPORTPATH)/cmd/rulechain

.PHONY: lbs 
lbs: cmd/lbs
	@echo "building location based service (lbs)..."
	$Q CGO_ENABLED=0 go build -o bin/lbs $(IMPORTPATH)/cmd/lbs

.PHONY: headmast 
headmast: cmd/headmast
	@echo "building headmast service (headmast)..."
	$Q CGO_ENABLED=0 go build -o bin/headmast $(IMPORTPATH)/cmd/headmast

.PHONY: test
test: 
	$Q go test  ./...




