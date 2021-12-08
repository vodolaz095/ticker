lint:
	gofmt  -w=true -s=true -l=true ./
	golint ./...
	go vet ./...
	npm run-script lint

frontend:
	npx vue-cli-service build
	cp src/assets/favicon.ico public/
	cp src/assets/fix.go public/
	cp src/assets/robots.txt public/

check: lint
	go test -v ./...

npm:
	npm ci

deps:
	go mod download
	go mod verify
	go mod tidy

build_prod: deps frontend assets_prod
	go build -o build/ticker main.go
	upx build/ticker
#	forego run build/ticker

assets_dev:
	rm -f public/public.go
	go-bindata -debug -o public/public.go -prefix "public/" -pkg public -fs public/

assets_prod:
	rm -f public/public.go
	go-bindata -o public/public.go -prefix "public/" -pkg public -fs public/css public/js public/

start: assets_dev
	forego start

podman_build:
	podman build -t localhost/ticker:latest .

podman_start: podman_build
	podman run --env-file=.env -p 3000:3000/tcp localhost/ticker:latest

docker_build:
	docker build -t localhost/ticker:latest .

docker_start: podman_build
	docker run --env-file=.env -p 3000:3000/tcp localhost/ticker:latest

