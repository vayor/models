default: build

build:
	rm -rf migrator
	@go build -o migrator -ldflags "-s -w" -trimpath utils/migrate.go