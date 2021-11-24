
# Env & Vars --------------------------------------------------------

include .env
export $(shell sed 's/=.*//' .env)

go_path = PATH=${PATH}:~/go/bin GOPATH=~/go 
go_env = $(go_path) GO111MODULE=on

# Tasks -------------------------------------------------------------

## # Help task ------------------------------------------------------
##

## help		Print project tasks help
help: Makefile
	@echo "\n ws-pdf-publish project tasks:\n";
	@sed -n 's/^##/	/p' $<;
	@echo "\n";

##
## # Global tasks ---------------------------------------------------
##

## clean		Clean the 'wkhtmltopdf' docker image
clean:
	@echo "\n> Clean";
	#@rm -rf $(build_report_path)/tests.* $(build_report_path)/coverage.* $(build_bin_path)/ws-pdf-publish;
	@docker rmi wkhtmltopdf:ws-pdf-publish || true;

## test		Run the tests
.PHONY: test
test:
	@echo "\n> Run tests";
	@$(go_env) \
		go get -u github.com/onsi/ginkgo/ginkgo github.com/onsi/gomega/... \
			gotest.tools/gotestsum github.com/t-yuki/gocover-cobertura;
	@rm -rf $(build_report_path)/tests.* $(build_report_path)/coverage.*;
	@$(go_env) \
		gotestsum --format standard-verbose --junitfile $(build_report_path)/tests.xml -- \
		-failfast -count=1 -coverprofile=$(build_report_path)/coverage.out ./...;
	@$(go_env) \
		go tool cover -html=$(build_report_path)/coverage.out -o $(build_report_path)/coverage.html && \
		$(go_path) gocover-cobertura < $(build_report_path)/coverage.out > $(build_report_path)/coverage.xml;	

## lint		Execute lint task
lint:
	@echo "\n> Run lint";
	@$(go_env) \
		go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0;
	@$(go_env) \
		golangci-lint run ./...;

## build		Build the command
.PHONY: build
build:
	@echo "\n> Build";
	@rm -rf $(build_bin_path)/ws-pdf-publish;
	@$(go_env) \
		go build -ldflags="-X 'github.com/cclavero/ws-pdf-publish/cmd.Version=$(VERSION)'" -o $(build_bin_path)/ws-pdf-publish ./main.go && \
		ls -lah $(build_bin_path)/ws-pdf-publish;
	@echo "\n> Check builded version";
	@$(build_bin_path)/ws-pdf-publish -v;

	
## ci		Execute all the CI tasks
ci: test lint build		

##
## # Install task ---------------------------------------------------
##

## install	Install the 'ws-pdf-publish' command to the '~/go/bin' folder
install:
	@echo "\n> Install";
	@$(go_env) \
		go install -ldflags="-X 'github.com/cclavero/ws-pdf-publish/cmd.Version=$(VERSION)'" github.com/cclavero/ws-pdf-publish && \
		ls -lah ~/go/bin;
	@echo "\n> Check installed version";
	@$(go_path) \
		ws-pdf-publish -v;
