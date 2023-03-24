#!/bin/bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
go build -o app.elf && chmod +x app.elf