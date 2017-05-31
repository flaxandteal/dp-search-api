export GOOS?=$(shell go env GOOS)
export GOARCH?=$(shell go env GOARCH)
thisOS:=$(shell uname -s)

MAIN=dp-search-query

HEALTHCHECK_ENDPOINT?=/healthcheck
DATA_CENTER?=dc1
DEV?=

CMD_DIR?=cmd
BUILD?=build
BUILD_ARCH=$(BUILD)/$(GOOS)-$(GOARCH)
HASH?=$(shell make hash)
DATE:=$(shell date '+%Y%m%d-%H%M%S')
TGZ_FILE=$(MAIN)-$(GOOS)-$(GOARCH)-$(DATE)-$(HASH).tar.gz

NOMAD?=
NOMAD_SRC_DIR?=nomad
NOMAD_PLAN_TARGET?=$(BUILD)
NOMAD_PLAN=$(NOMAD_PLAN_TARGET)/$(MAIN).nomad

ifdef DEV
HUMAN_LOG?=1
else
HUMAN_LOG?=
endif

ifeq ($(thisOS),Darwin)
SED?=gsed
else
SED?=sed
endif

build:
	@mkdir -p $(BUILD_ARCH)/bin
	go build -o $(BUILD_ARCH)/bin/$(MAIN) $(CMD_DIR)/$(MAIN)/main.go
	cp -r templates $(BUILD_ARCH)/templates

package: build
	tar -zcf $(TGZ_FILE) -C $(BUILD_ARCH) .

$(MAIN) run:
ifdef NOMAD
	@if [[ ! -f $(NOMAD_PLAN) ]]; then echo Cannot see $(NOMAD_PLAN); exit 1; fi; echo nomad run $(NOMAD_PLAN); nomad run $(NOMAD_PLAN)
else
	@main=$(CMD_DIR)/$@/main.go; if [[ ! -f $$main ]]; then echo Cannot see $$main; exit 1; fi; go run -race $$main
endif

nomad:
	@test -d $(NOMAD_PLAN_TARGET) || mkdir -p $(NOMAD_PLAN_TARGET)
	@driver=exec; [[ -n "$(DEV)" ]] && driver=raw_exec;	\
	$(SED) -r	\
		-e 's,\bDATA_CENTER\b,$(DATA_CENTER),g' \
		-e 's,\bS3_TAR_FILE\b,$(S3_TAR_FILE),g' \
		-e 's,\bELASTIC_SEARCH_URL\b,$(ELASTIC_URL),g' \
		-e 's,\bHEALTHCHECK_ENDPOINT\b,$(HEALTHCHECK_ENDPOINT),g' \
		-e 's,\bHUMAN_LOG_FLAG\b,$(HUMAN_LOG),g'		\
		-e 's,^(  *driver  *=  *)"exec",\1"'$$driver'",'	\
		< $(MAIN)-template.nomad > $(NOMAD_PLAN)
hash:
	@git rev-parse --short HEAD

debug: build
	HUMAN_LOG=$(HUMAN_LOG) ./$(BUILD_ARCH)/$(MAIN)

waitOnElastic:   export ELASTIC_URL = http://localhost:9999/
waitOnElastic:
	pause.sh


test: 	export BIND_ADDR = :10002
test:   export ELASTIC_URL = http://localhost:9999/
test:
	./waitForElastic.sh
	go test

bdd: startElastic test stopElastic


startElastic:
	docker run --name es-bdd  -d   -p 9999:9200 guidof/onswebsite-search:5.0.0

stopElastic:
	docker rm -f es-bdd

.PHONY: build debug test startElastic stopElastic package
