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
