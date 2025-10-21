.PHONY: user

user:
	@go build -o bin/user ./user
	@./bin/user