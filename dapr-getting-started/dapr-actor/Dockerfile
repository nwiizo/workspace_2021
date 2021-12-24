FROM golang:1.17 as builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG SERVICE_PATH=actor
ARG CGO_ENABLED=0
#ARG GOOS=darwin
#ARG GOARCH=arm64
COPY api api
COPY domain domain
COPY $SERVICE_PATH $SERVICE_PATH
RUN go build -o /go/bin/app -ldflags '-s -w' $SERVICE_PATH/main.go

FROM scratch as runner

COPY --from=builder /go/bin/app /app

CMD ["/app"]
