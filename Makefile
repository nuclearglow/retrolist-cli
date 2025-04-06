FORMATTED_DATE := $(shell echo `date +"%Y-%m-%d %T"`)

.PHONY: build

build: clean compile

list:
	@LC_ALL=C $(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | grep -E -v -e '^[^[:alnum:]]' -e '^$@$$'

clean:
	go clean

# production builds: consider ldflags -s -w
compile:
	go build -o=bin/pit -ldflags="-X 'retrocli.svenvowe.de/config.BuildVersion=0.0.1' -X 'retrocli.svenvowe.de/config.BuildTime=$(FORMATTED_DATE)'"
	echo -e "Built ${COLOR_LIGHT_GREEN}./bin/pit${COLOR_BLACK}"

run:
	./bin/pit

lint:
	watchexec --exts go --watch . --restart --clear=clear "golangci-lint run"

watch:
	watchexec --exts go,json --no-vcs-ignore --watch . --restart --clear=clear "go run ."

benchmark:
	hyperfine --warmup 3 --prepare 'make clean' 'make compile'

# TODO
#compile:
#	echo "Compiling for every OS and Platform"
# GOOS=darwin GOARCH=aarch? go build -o bin/main-darwin
#	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386
#	GOOS=linux GOARCH=386 go build -o bin/main-linux-386
#	GOOS=windows GOARCH=386 go build -o bin/main-windows-386
