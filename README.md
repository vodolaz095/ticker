Ticker
===============================
Tinkoff Investement Dashboard

Test task for usage of [VueJS 3x](https://v3.vuejs.org), [Gin](github.com/gin-gonic/gin) 
web framework and [official golang bindings](https://github.com/TinkoffCreditSystems/invest-openapi-go-sdk)
for [Tinkoff Open Investement API](https://tinkoffcreditsystems.github.io/invest-openapi/).
Simply display table with user's portfolio with positions and currencies

Screenshot
================================


Requirements
==================================
These packages are required in order to build application

1. Linux (tested on Centos 8 stream, Fedora 34+)
2. Golang (>=1.16.8)
3. NodeJS (>=14.18.1)
4. NPM (>=6.14.15)
5. GNU Make (>=4.3)
6. [Forego](https://github.com/ddollar/forego) - Optional dependency to start application loading environment from .env file 

How to build executable 
===================================

```shell
$ make build
```

Binary executable will be generated in `build/` subdirectory.
Исполняемый файл будет сгенерирован в поддиректории `build/`.

Building application using docker
===================================
```shell
# make docker_build
```

Building application using podman
===================================
```shell
# make podman_build 
```

Configuration parameters/настройки приложения
====================================
Application follows [twelve factor apps manifest](https://12factor.net/config), so, configuration parameters
are loaded from Posix environment.

This parameters are taken in account/учитываются эти параметры:

- ADDR - default value `127.0.0.1` - IP address of local network interface application binds to as HTTP server 
  / значение по умолчанию `127.0.0.1` - адрес локального сетевого интерфейса, на котором приложение служает соединения по протоколу HTTP 
- PORT - default value `3000` - port, on which application starts HTTP server 
  / значение по умолчанию - 3000, порт, на котором приложение запускает HTTP сервер
- TOKEN - nj
- ACCOUNT_ID, "")

Docker swarm service definition
=====================================





