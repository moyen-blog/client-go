run:
	@cd moyen-cli && go run . || true

build:
	@cd moyen-cli && go build -o moyen-cli.out
