BIN=app

build_release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(BIN)

tailwind:
	tailwindcss -i ./static/input.tailwind.css -o ./static/styles.css --watch
