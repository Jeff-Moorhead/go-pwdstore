all: build

bin:
	@echo "Creating output directory at ./bin/"
	@mkdir -p bin
	
	
build: bin
	@echo "Build go-pwdstore"
	@go build -o bin/go-pwdstore cmd/main.go	

install-user:
	@mkdir -p ~/bin
	@go build -o ~/bin/go-pwdstore cmd/main.go	
	
