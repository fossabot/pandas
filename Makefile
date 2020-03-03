IMPORTPATH = github.com/cloustone/pandas
# V := 1 # When V is set, print cmd and build progress.
Q := $(if $V,,@)

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"'

# Space separated patterns of packages to skip in list, test, format.
DOCKER_NAMESPACE := cloustone

.PHONY: all
all: build

.PHONY: build
build: apimachinery  dmms  pms rulechain lbs headmast 

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

.PHONY: test
test: 
	$Q go test  ./...





