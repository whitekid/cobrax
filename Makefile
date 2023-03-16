test:
	@go test ./...

# update modules & tidy
dep:
	@rm -f go.mod go.sum
	@go mod init github.com/whitekid/cobrax

	@$(MAKE) tidy

tidy:
	@go mod tidy -v
