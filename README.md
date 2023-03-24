<h1 align='center'>
  Golang Backend API
</h1>

## 📜 Table of Contents
- [🧐 Business context](#-business-context)
- [🔨 Technology stack](#-technology-stack)
- [💻 Project setup](#-project-setup)
- [Clone this repo](#clone-this-repo)
- [Setup .env file](#setup-env-file)
- [Generate swagger](#generate-swagger)

---

# 🧐 Business context

The goal of the project is to create a golang codebase for backend apis
---

# 🔨 Technology stack

- Swagger
- Golang
- Gorm
- Gin

---


# 💻 Project setup


### Clone this repo

To install this project, first clone repo to your machine (use SSH) and run build.sh

### Install packages

Execute this command in your terminal:

```
go mod vendor
```

### Setup .env file

Create `.env` file based on `.env.sample` and fill in variables from a suitable source.

### Generate swagger
TLDR: running build.sh will generate swagger docs.

1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:

```sh
go get -u github.com/swaggo/swag/cmd/swag
```

Starting in Go 1.17, installing executables with `go get` is deprecated. `go install` may be used instead:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

3. Run the [Swag](https://github.com/swaggo/swag) at your Go project root path(for instance `~/root/golang-base`),
   [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`)
   at `~/root/golang-base/docs`.

```sh
swag init
```
