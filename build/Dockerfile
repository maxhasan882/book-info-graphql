FROM golang:alpine as build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN mkdir /ds_doc
WORKDIR /ds_doc
COPY go.mod go.sum /ds_doc/

RUN go mod download

COPY [".", "/ds_doc/"]
EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app cmd/app/main.go
FROM scratch
COPY --from=build-env /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]