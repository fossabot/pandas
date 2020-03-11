IMPORTPATH = github.com/cloustone/pandas
# V := 1 # When V is set, print cmd and build progress.
Q := $(if $V,,@)

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

# Space separated patterns of packages to skip in list, test, format.
DOCKER_NAMESPACE := cloustone
DOCKERFILE_ADDRESS := docker/Dockerfile

.PHONY: all
all: build

.PHONY: build
build: docker apimachinery dmms  pms rulechain lbs headmast shiro

.PHONY: docker
docker: 
	# @echo "Creating docker image (docker)..."
	# $Q docker build -f ${DOCKERFILE_ADDRESS} -t ${DOCKER_NAMESPACE} .
	# $Q docker build -f cmd/apimachinery/Dockerfile -t api .
	# $Q docker build -f cmd/dmms/Dockerfile -t dmms .
	# $Q docker build -f cmd/headmast/Dockerfile -t headmast .
	# $Q docker build -f cmd/lbs/Dockerfile -t lbs .
	# $Q docker build -f cmd/pms/Dockerfile -t pms .
	# $Q docker build -f cmd/rulechain/Dockerfile -t rulechain .
	# $Q docker build -f cmd/shiro/Dockerfile -t shiro .
#    docker run -it --name zhmm -P cloustone


.PHONY: apimachinery 
apimachinery: 
	@echo "building api server (apimachinery)..."
	$Q CGO_ENABLED=0 go build -v -o bin/pandas-apimachinery $(IMPORTPATH)/cmd/apimachinery

.PHONY: dmms 
dmms: cmd/dmms 
	@echo "building device management server (dmms)..."
	$Q CGO_ENABLED=0 go build -o bin/pandas-dmms $(IMPORTPATH)/cmd/dmms

.PHONY: pms 
pms: cmd/pms 
	@echo "building project management server (pms)..."
	$Q CGO_ENABLED=0 go build -o bin/pandas-pms $(IMPORTPATH)/cmd/pms

.PHONY: rulechain 
rulechain: cmd/rulechain
	@echo "building rulechain server (rulechain)..."
	$Q CGO_ENABLED=0 go build -o bin/pandas-rulechain $(IMPORTPATH)/cmd/rulechain

.PHONY: lbs 
lbs: cmd/lbs
	@echo "building location based service (lbs)..."
	$Q CGO_ENABLED=0 go build -o bin/pandas-lbs $(IMPORTPATH)/cmd/lbs

.PHONY: headmast 
headmast: cmd/headmast
	@echo "building headmast service (headmast)..."
	$Q CGO_ENABLED=0 go build -o bin/pandas-headmast $(IMPORTPATH)/cmd/headmast

.PHONY: shiro 
shiro: cmd/shiro
	@echo "building unified user manager center service (shiro)..."
	$Q CGO_ENABLED=0 go build -o bin/pandas-shiro $(IMPORTPATH)/cmd/shiro


.PHONY: test
test: 
	$Q go test  ./...




