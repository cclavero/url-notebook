
# Env & Vars --------------------------------------------------------

version = 1.0

pwd = $(shell pwd)
userUID = $(shell id -u)
userGID = $(shell id -g)

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
	docker rmi url-notebook:$(version) || true;

## build			Build the docker image
build:
	@echo "\n> Build";
	docker build -f ./Dockerfile --tag url-notebook:$(version) .;

## run		Run the url-notebook command
run:
	@echo "\n> Run";
	go run main.go targetPath=$(pwd)/out userUID=$(userUID) userGID=$(userGID) dockerExtraParams="--network ni-net"
