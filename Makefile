lint:
	gofmt  -w=true -s=true -l=true ./
	golint ./...
	go vet ./...

frontend:
	npx vue-cli-service build
	cp src/assets/favicon.ico public/
	cp src/assets/fix.go public/
	cp src/assets/robots.txt public/

check: lint
	go test -v ./...

deps:
	go mod download
	go mod verify
	go mod tidy

assets_dev:
	rm -f public/public.go
	go-bindata -debug -o public/public.go -prefix "public/" -pkg public -fs public/

assets_prod:
	rm -f public/public.go
	go-bindata -o public/public.go -prefix "public/" -pkg public -fs public/

start: assets_dev
	forego start
