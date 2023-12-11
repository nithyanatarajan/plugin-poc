# Plugin POC

[![Build Status](https://github.com/nithyanatarajan/plugin-poc/actions/workflows/ci.yml/badge.svg)](https://github.com/nithyanatarajan/plugin-poc/actions/workflows/ci.yml)

This project hosts idea of hosting multiple providers via plugin architecture

## Run

To run the application

```shell
make run
```

### Details

Compile the providers via:

```shell
go build -C providers -o ../out/providers
```

Compile the driver via:

```shell
go build -o ./out/driver ./cmd
```

Launch app via:

```shell
./out/driver
```
