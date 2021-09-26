
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
	docker rmi url-notebook:local || true;

## build			Build the url-notebook cammand
.PHONY: build
build:
	@echo "\n> Build";
	go build -o ./build/url-notebook ./cmd/main.go;

## run			Run the url-notebook command
run:
	@echo "\n> Run";
	go run ./cmd/main.go targetPath=./out urlNotebookFile=./test/url-notebook-test1.yaml # dockerExtraParams="--network ni-net";

## install		Install the url-notebook command
install: build
	@echo "\n> Install";
	sudo cp ./build/url-notebook /usr/local/bin;
	ls -lah /usr/local/bin/url-notebook;
	url-notebook -v;
