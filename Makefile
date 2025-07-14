BIN=app

build_release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(BIN)

tailwind:
	tailwindcss -i ./static/input.tailwind.css -o ./static/styles.css --watch

alpine:
	curl https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js >> ./static/alpine.js
