# ws-pdf-publish: WebSite PDF Publish

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

## Requirements

- Golang env ver. >=1.16: https://go.dev, https://github.com/moovweb/gvm
- Docker runtime ver. >= 20: https://docker.com 

## Resources

- https://github.com/spf13/viper
- https://github.com/onsi/ginkgo
- https://github.com/pdfcpu/pdfcpu
- https://wkhtmltopdf.org

## Install & usage

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
$ ws-pdf-publish --publishFile ./build/test/ws-pub-pdf-test.yaml --targetPath /tmp/ws-pub-pdf
Starting ...
...
Done, full PDF file generated at: /tmp/ws-pub-pdf/test.pdf
```

When the command execution finishes, you will find the PDF file for each URL defined under the `url` tag (partial PDFs, in the publish file under the `targetPath` 'url' subfolder) and the final resultant PDF file (in the `targetPath`):

```bash
$ ls -la /tmp/ws-pub-pdf/url
total 2888
drwxr-xr-x  3 clavero  wheel       96 Nov 30 16:40 .
drwxr-xr-x  4 clavero  wheel      128 Nov 30 16:40 ..
-rw-r--r--  1 clavero  wheel  1478549 Nov 30 16:40 boe.pdf
...
$ ls -la /tmp/ws-pub-pdf/test.pdf
-rw-r--r--  1 clavero  wheel  1392835 Nov 30 16:40 /tmp/ws-pub-pdf/test.pdf
```

### Format for the publish YAML file

--- TODO: sintaxis i exemple de fitxer

```yaml
publish:
  file: notes-inxes.pdf # Name of the final PDF file
  urls: # List of URLs to process
    - url: http://notes-inxes:1313/
      file: portada.pdf
    - url: http://notes-inxes:1313/docs/esquemes-generals
      file: esquemes-generals.pdf
    - url: http://notes-inxes:1313/docs/do-met-gallec/
      file: do-met-gallec.pdf
    - url: http://notes-inxes:1313/docs/do-met-escoces/
      file: do-met-escoces.pdf
    - url: http://notes-inxes:1313/docs/sol-met-gallec/
      file: sol-met-gallec.pdf
  wkhtmltopdfParams: --print-media-type --margin-top 20mm --margin-bottom 20mm # Parameters for the wkhtmltopdf utility    
```

### Configure the PDF generation with wkhtmltopdf

The wkhtmltopdf utility (https://wkhtmltopdf.org) supports a large number of parameters to configure the PDF documents generation: https://wkhtmltopdf.org/usage/wkhtmltopdf.txt

You can use the `wkhtmltopdfParams` field in the publish YAML file to set all those parameters.

## Develop

### Project layout

--- TEMPORAL

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
$ ARGS="--publishFile ./build/test/ws-pub-pdf-test.yaml --targetPath ./build/test/out-cmd" make run
...
```
