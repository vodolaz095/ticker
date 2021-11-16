# certs and shared libs
FROM alpine:3.6 AS alpine
RUN apk add -U --no-cache ca-certificates

# building frontend
FROM node:14.18.1 AS frontend
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json
WORKDIR /app
RUN npm ci

# adding vuejs source codes
COPY * /app
# building vuejs source code
RUN	npx vue-cli-service build

# preparing sane docker image to build app
FROM centos:8 AS build
RUN dnf upgrade -y && dnf install -y golang git make epel-release && dnf install -y upx && dnf clean all
RUN go get -u github.com/go-bindata/go-bindata/...
RUN /root/go/bin/go-bindata -version
RUN mv /root/go/bin/go-bindata /usr/bin/go-bindata
RUN go get -u golang.org/x/lint/golint
RUN mv /root/go/bin/golint /usr/bin/golint

# building app
RUN mkdir -p /opt/ticker
WORKDIR /opt/ticker
ADD . /opt/ticker/
# copy assets from nodejs image
COPY --from=node /app/public/ /opt/ticker/public
# copy frontend assets
RUN cp /opt/ticker/src/assets/favicon.ico /opt/ticker/public/
RUN cp /opt/ticker/src/assets/fix.go /opt/ticker/public/
RUN cp /opt/ticker/src/assets/robots.txt /opt/ticker/public/

# build production grade binary
RUN make build_prod
RUN ls -l /opt/ticker/build/

# result with size ~ 10mb
FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build /opt/ticker/build/ticker /bin/ticker

EXPOSE 3000
ENTRYPOINT ["/bin/ticker"]
