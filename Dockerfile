FROM golang:1.8.0

RUN apt-get update && apt-get install lsb-release libgsf-1-dev -y && \
    curl -s https://raw.githubusercontent.com/h2non/bimg/master/preinstall.sh | bash -
