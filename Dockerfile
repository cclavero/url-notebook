FROM ubuntu:20.04

RUN apt-get update && \
    apt-get upgrade -y;
    
RUN apt-get install wget sudo -y && \
    wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.focal_amd64.deb && \
    sudo apt install ./wkhtmltox_0.12.6-1.focal_amd64.deb -y;

ENTRYPOINT ["wkhtmltopdf"]

CMD ["-h"]
