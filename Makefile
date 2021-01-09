run:
	@go run . || true

watch:
	@reflex -r '\.go' -s -- sh -c "make build"

build:
	@go build -o sync.out

sync:
	@cd example && ../sync.out
