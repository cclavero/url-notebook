
- Info:

--- TEMPORAL

- https://hub.docker.com/r/openlabs/docker-wkhtmltopdf
- https://github.com/openlabs/docker-wkhtmltopdf
- https://github.com/pdfcpu/pdfcpu


$ docker run -u 1000:1000 -v $(pwd)/pdf:/pdf --network ni-net openlabs/docker-wkhtmltopdf http://notes-inxes:1313/docs/esquemes-generals/ /pdf/esquemes-generals.pdf


--- TEMPORAL: DOCKER IN DOCKER

$ docker build -f ./Dockerfile-publish-pdf --tag ni-publish-pdf:1.0 .
$ docker run -it --privileged -v /var/run/docker.sock:/var/run/docker.sock --name ni-publish-pdf --network ni-net --rm ni-publish-pdf:1.0 /bin/sh

--- TEMPORAL

$ go run main.go targetPath=$PWD/../out userUID=`id -u` userGID=`id -g` dockerExtraParams="--network ni-net"


- Makefile:

#docker build -f ./Dockerfile-publish-pdf --tag ni-publish-pdf:$(version) .;
#docker run -it -u $(shell id -u):$(shell id -g) --privileged -v /var/run/docker.sock:/var/run/docker.sock \
#	--name ni-publish-pdf --network ni-net --rm ni-publish-pdf:$(version);


- Dockerfile-publish-pdf:

FROM golang:1.16 AS build

WORKDIR /src

COPY ./publish-pdf .

RUN go build -o /out/publish-pdf .

# TEMPORAL:versio
FROM alpine

#RUN apk add docker && \
#    rc-update add docker boot;

#FROM docker:dind


COPY --from=build /out/publish-pdf /usr/local/bin

CMD ["publish-pdf"]


--- TEMPORAL:wkhtmltopdf

$ wkhtmltopdf --version
wkhtmltopdf 0.12.5 # qt unpatched
$ wkhtmltopdf --print-media-type http://localhost:1313/docs/do-met-gallec/ out.pdf

$ docker run -u 1000:1000 -v `pwd`:/out --network ni-net openlabs/docker-wkhtmltopdf:v0.12

$ docker run -u 1000:1000 -v `pwd`:/out --network ni-net openlabs/docker-wkhtmltopdf:v0.12 --print-media-type http://notes-inxes:1313/docs/do-met-gallec/ /out/out-docker.pdf


--- TEMPORAL

https://www.loginradius.com/blog/async/build-push-docker-images-golang/



---- TEMPORAL

$ cd build
$ ./url-notebook targetPath=../out urlNotebookFile=../test/url-notebook-test1.yaml