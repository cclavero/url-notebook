# WebSite PDF Publis (hws-pdf-publish)

@author Carles Clavero i Matas - carles.clavero@gmail.com

@date 2021-11-28

![Go Version](https://img.shields.io/badge/go-ver%3E=1.16-informational)
![Docker](https://img.shields.io/badge/docker-ver%3E=20-informational)
![License](https://img.shields.io/badge/license-MIT-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/cclavero/ws-pdf-publish)](https://goreportcard.com/report/github.com/cclavero/ws-pdf-publish)
![Tests](https://img.shields.io/badge/tests-passed-green)
![Coverage](https://img.shields.io/badge/coverage-80%25-green)

WebSite PDF Publish: Simple command line tool to publish a set of HTML pages to PDF.

Current version: 1.0-alpha

## Install & usage

### Requirements

- Golang env ver. >=1.16: https://go.dev, https://github.com/moovweb/gvm
- Docker runtime ver. >= 20: https://docker.com 

### Install 'ws-pdf-publish' command

To install the 'ws-pdf-publish' command to the '~/go/bin' folder from the source code, you only need to execute the Makefile `install` task:

```bash
$ make install
...
> Check installed version
ws-pdf-publish version 1.0-alpha
```

You can also install the 'ws-pdf-publish' command using the standard go install task, setting te concrete version:

```bash
$ go install -ldflags="-X 'github.com/cclavero/ws-pdf-publish/cmd.Version=1.0-alpha'" github.com/cclavero/ws-pdf-publish@latest
$ ws-pdf-publish -v
ws-pdf-publish version 1.0-alpha
```
### Command usage and flags

To get the basic help about the 'ws-pdf-publish' command and their flags, you can use the `-h` flag:

```bash
$ ws-pdf-publish -h
WebSite PDF Publish is a simple command line tool to publish some set of pages from a WebSite to PDF, using a 'ws-pub-pdf.yaml' configuration file.
Internally, uses the wkhtmltopdf utility.

Usage:
  ws-pdf-publish [flags]

Flags:
  -h, --help                 help for ws-pdf-publish
      --publishFile string   set the 'ws-pub-pdf.yaml' publish file, including absolute or relative path.
      --targetPath string    set the target path for publishing partial and final PDF files.
  -v, --version              version for ws-pdf-publish
```

Using the `-v` flag you can get the version of the command:

```bash
$ ws-pdf-publish -v
ws-pdf-publish version 1.0-alpha
```

To execute the command, you need to set the `publishFile` and `targetPath` flags, as a example:

```bash
$ ws-pdf-publish --publishFile ./build/test/ws-pub-pdf-example.yaml --targetPath /tmp/ws-pub-pdf
Starting ...
...
Done, full PDF file generated at: /tmp/ws-pub-pdf/go-refs.pdf
```

When the command execution finishes, you will find the PDF file for each URL defined under the `url` tag (partial PDFs, in the publish file under the `targetPath` 'url' subfolder) and the final resultant PDF file (in the `targetPath`):

```bash
$ ls -la /tmp/ws-pub-pdf/url
total 288
drwxrwxr-x 2 carolo carolo   4096 Nov 30 18:15 .
drwxrwxr-x 3 carolo carolo   4096 Nov 30 18:15 ..
-rw-r--r-- 1 carolo carolo  93276 Nov 30 18:15 go-cobra.pdf
-rw-r--r-- 1 carolo carolo 189171 Nov 30 18:15 go-viper.pdf
$ ls -la /tmp/ws-pub-pdf/go-refs.pdf
-rw-rw-r-- 1 carolo carolo 244153 Nov 30 18:15 /tmp/ws-pub-pdf/go-refs.pdf
```

### Format for the publish YAML file

The format of the publish YAML file is quite simple and self-explanatory. As an example:

```yaml
publish:
  # Name of the final PDF file (merged)
  file: go-refs.pdf
  # List of URLs to process (ordered)
  urls: 
    - url: https://github.com/spf13/cobra/blob/master/README.md
      file: go-cobra.pdf
    - url: https://github.com/spf13/viper/blob/master/README.md
      file: go-viper.pdf
  # Parameters for the docker execution
  dockerParams: 
  # Parameters for the wkhtmltopdf utility    
  wkhtmltopdfParams: --print-media-type --margin-top 20mm --margin-bottom 20mm 
```

### Configure the PDF generation with wkhtmltopdf

The wkhtmltopdf utility (https://wkhtmltopdf.org) supports a large number of parameters to configure the PDF documents generation: https://wkhtmltopdf.org/usage/wkhtmltopdf.txt

You can use the `wkhtmltopdfParams` field in the publish YAML file to set all those parameters.

## Develop

### Resources

- https://github.com/spf13/viper
- https://github.com/onsi/ginkgo
- https://github.com/pdfcpu/pdfcpu
- https://wkhtmltopdf.org


### Project layout

The project layout is a standard go project layout, with some special folders:

- 'Makefile' and '.env' files: Main Makefile for the develop tasks and env vars definition for the project.
- '/build' folder: Folder for the local environment developing, with some sub-folders:
  - '/build/bin': Folder for the command generated binary.
  - '/build/report': Folder for the CI generated reports (jUnit test files, Coverage files, including HTML report, etc).
  - '/build/test': Folder for test resource files and used as base folder for testing.
- '/test': go sourcecode only for testing (test functions and utilities).   

### Makefile tasks

All the develop tasks are defined in the Makefile.

1. Execute the `help` task to get the help for all the defined tasks:

```bash
$ make help

ws-pdf-publish project tasks:

	 # Help task ------------------------------------------------------

	 help		Print project tasks help
...   
```

2. Execute the `ci` task to run all the CI cycle (substasks `clean`, `test`, `lint` and `build`):

```bash
$ make ci
...
```

3. Execute the `run` task to compile and run the current source code for testing:

```bash
$ ARGS="--publishFile ./build/test/ws-pub-pdf-example.yaml --targetPath ./build/test/out" make run
...
```
