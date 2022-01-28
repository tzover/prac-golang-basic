FROM golang:1.16.4-buster
RUN apt-get update && apt update
# CGO = C言語
ENV CGO_ENABLED=1 \ 
    GOOS=linux \
    GOARCH=amd64 \
    ROOT=/home/app \
    GO111MODULE=on
WORKDIR ${ROOT}
EXPOSE 8080
# CMD [app]