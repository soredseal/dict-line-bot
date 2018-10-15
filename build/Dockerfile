FROM golang:latest AS BUILDER
WORKDIR /go/src/github.com/sariakra/line-dictionary-bot/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dict-bot cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=BUILDER /go/src/github.com/sariakra/line-dictionary-bot/ .
CMD ["./dict-bot"]
