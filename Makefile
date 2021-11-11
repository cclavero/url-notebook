
# Env & Vars --------------------------------------------------------

include .env
export $(shell sed 's/=.*//' .env)

# Tasks -------------------------------------------------------------

## # Help task --------------------------------------------------
##

## help			Print project tasks help
help: Makefile
	@echo "\n url-notebook project tasks:\n";
	@sed -n 's/^##/	/p' $<;
	@echo "\n";

##
## # Global tasks -----------------------------------------------
##

## clean			Clean the 'wkhtmltopdf' docker image
clean:
	@echo "\n> Clean";
	docker rmi wkhtmltopdf:notes-inxes || true;

## build			Build the url-notebook command
.PHONY: build
build:
	@echo "\n> Build";
	go build -ldflags="-X 'main.Version=$(VERSION)'" -o ./build/url-notebook ./cmd/main.go;

## run-test		Run the url-notebook command for test
run-test:
	@echo "\n> Run Test";
	go run -ldflags="-X 'main.Version=$(VERSION)'" ./cmd/main.go publishFile=./test/url-notebook-test.yaml targetPath=./test/pdf;


# TEMPORAL
## install		Install the url-notebook command
install: build
	@echo "\n> Install";
	sudo cp ./build/url-notebook /usr/local/bin;
	ls -lah /usr/local/bin/url-notebook;
#	#url-notebook -v;
