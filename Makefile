PROJ=$(shell realpath $$PWD/../../../..)
ENV=env GOPATH=$(PROJ)
CLI=$(ENV) easyjson -all

JSON_SRC_FILES=\
	structs.go

JSON_OUT_FILES:=$(JSON_SRC_FILES:%.go=%_easyjson.go)

json: $(JSON_OUT_FILES)

%_easyjson.go: %.go
	$(CLI) $^

