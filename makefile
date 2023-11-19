test:
	@go test -count=1 -timeout 30m ./...

vtest:
	@go test -count=1 -timeout 30m ./... -v