FROM golang:1.17.2-bullseye AS golang-base

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app.elf && chmod +x app.elf


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /deploy
COPY --from=golang-base /go/src/app/app.elf .
EXPOSE 8080

CMD ["./app.elf"]  

