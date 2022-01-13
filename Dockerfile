FROM golang:1.17.1-alpine AS BUILDER

WORKDIR /app

ADD go.sum go.mod /app/
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download

ADD . .

RUN go build main.go

FROM alpine:latest

COPY --from=BUILDER /app/main /bin/main

RUN chmod +x /bin/main

EXPOSE 8080

CMD ["/bin/main"]
