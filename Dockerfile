FROM golang:1.20-alpine3.19
WORKDIR /go/src
COPY ./src .
RUN go mod download -x
RUN go build main.go &&  \
    mv main /go/bin/app && \
    rm -Rf /go/src
EXPOSE 8080
ENTRYPOINT /go/bin/app