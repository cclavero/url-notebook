
# Env & Vars --------------------------------------------------------

include .env
export $(shell sed 's/=.*//' .env)

# Tasks -------------------------------------------------------------

## # Help task --------------------------------------------------
##

## help			Print project tasks help
help: Makefile
	@echo "\n ws-pdf-publish project tasks:\n";
	@sed -n 's/^##/	/p' $<;
	@echo "\n";

##
## # Global tasks -----------------------------------------------
##

## clean			Clean the 'wkhtmltopdf' docker image
clean:
	@echo "\n> Clean";
	docker rmi wkhtmltopdf:ws-pdf-publish || true;

# TODO
## test			Run the tests
.PHONY: test
test:
	@echo "\n> Run Test";
	go run -ldflags="-X 'github.com/cclavero/ws-pdf-publish/cmd.Version=$(VERSION)'" ./main.go publishFile=./test/ws-pub-pdf-test.yaml targetPath=./test/pdf;

## build			Build the url-notebook command
.PHONY: build
build:
	@echo "\n> Build";
	go build -ldflags="-X 'github.com/cclavero/ws-pdf-publish/cmd.Version=$(VERSION)'" -o ./build/ws-pdf-publish ./main.go;

# TEMPORAL:REVISAR
## install		Install the url-notebook command
install: build
	@echo "\n> Install";
	sudo cp ./build/ws-pdf-publish /usr/local/bin;
	ls -lah /usr/local/bin/ws-pdf-publish;
	ws-pdf-publish -v;
