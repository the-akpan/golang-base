FROM golang:1.17.2-bullseye AS golang-base

WORKDIR /go/src/app
COPY . .
# RUN go build -o app.elf && chmod +x app.elf
RUN ./build.sh


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /deploy
COPY --from=golang-base /go/src/app/app.elf .
EXPOSE 8080

CMD ["./app.elf"]  

