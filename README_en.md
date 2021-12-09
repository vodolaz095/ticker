Ticker
===============================
Tinkoff Investement Dashboard

Test task for usage of [VueJS 3x](https://v3.vuejs.org), [Gin](github.com/gin-gonic/gin) 
web framework and [official golang bindings](https://github.com/TinkoffCreditSystems/invest-openapi-go-sdk)
for [Tinkoff Open Investement API](https://tinkoffcreditsystems.github.io/invest-openapi/).
Simply display table with user's portfolio with positions and currencies


Warning!
================================
This application uses your Tinkoff Open Investement API token, at current moment (9 december 2021 year) it 
gives full access to all broker accounts, so, if attacker obtains your token he/she can buy/sell your positions, 
perform pump attacks and so on...
At first glance, application requires read only access to show state of portfolio, but it was not tested
by any security experts, and there can be risk your token can be stolen and misused.
So, by using this application you accept responsibility for all possible money loss you can encounter from using this app.


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
- 
- TITLE - page title, by default - "Инвестиционный портфель".
- ADDR - default value `127.0.0.1` - IP address of local network interface application binds to as HTTP server 
- PORT - default value `3000` - port, on which application starts HTTP server 
- TOKEN - no default value, can be obtained as described [here](https://tinkoffcreditsystems.github.io/invest-openapi/auth/)
- ACCOUNT_ID - account id, can be obtained from application console output

Docker swarm service definition
=====================================
See `contrib/ticket.stack.yml` for details

License
=====================================

MIT License

Copyright (c) 2021 Anatolij Ostroumov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


