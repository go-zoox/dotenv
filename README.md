# DotEnv
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/dotenv)](https://pkg.go.dev/github.com/go-zoox/dotenv)
[![Build Status](https://github.com/go-zoox/dotenv/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/dotenv/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/dotenv)](https://goreportcard.com/report/github.com/go-zoox/dotenv)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/dotenv/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/dotenv?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/dotenv.svg)](https://github.com/go-zoox/dotenv/issues)
[![Release](https://img.shields.io/github/release/go-zoox/dotenv.svg?label=Release)](https://github.com/go-zoox/dotenv/releases)

Load application environment variables from a `.env` file into the current process.

```go
type Config struct {
  Host string `env:"HOST" default:"0.0.0.0"`
  Port int    `env:"PORT" default:"8080"`
  DatabaseUrl string `env:"DATABASE_URL", required:"true"`
  Email string `env:"EMAIL"`
}

var config Config
if err := dotenv.Load(&config); err != nil {
  panic(err)
}
```

## Inspired by
* [joho/godotenv](https://github.com/joho/godotenv) - dot env file
* [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) - struct tag
* [JeremyLoy/config](https://github.com/JeremyLoy/config) - ci