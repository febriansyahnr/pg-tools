run: build
	@./bin/pg-tools

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss postcss autoprefixer
	@npx tailwindcss init -p
	@npm install -D daisyui@latest

build:
	npx tailwindcss -i view/css/app.postcss -o public/styles.css
	@templ generate view
	@go build -o bin/pg-tools main.go