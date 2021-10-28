
# Env & Vars --------------------------------------------------------

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

## clean			Clean the docker image
clean:
	@echo "\n> Clean";
	docker rmi wkhtmltopdf:notes-inxes || true;

## build			Build the url-notebook cammand
.PHONY: build
build:
	@echo "\n> Build";
	go build -o ./build/url-notebook ./cmd/main.go;

## run			Run the url-notebook command
# TEMPORAL: dockerExtraParams="--network ni-net";
run:
	@echo "\n> Run";
	go run ./cmd/main.go urlNotebookFile=./pdf/url-notebook-test-1.yaml targetPath=./pdf/test-1 dockerExtraParams="";

## install		Install the url-notebook command
install: build
	@echo "\n> Install";
	sudo cp ./build/url-notebook /usr/local/bin;
	ls -lah /usr/local/bin/url-notebook;
	url-notebook -v;
