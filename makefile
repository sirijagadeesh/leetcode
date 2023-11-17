test:
	@go test -count=1 -timeout 30m -p 1 ./...

vtest:
	@go test -count=1 -timeout 30m -p 1 ./... -v