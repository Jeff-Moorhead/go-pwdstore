USERBIN := ~/bin
PACKAGEBIN := ./bin
EXECUTABLE := go-pwdstore
MAIN := cmd/main.go

all: build

$(PACKAGEBIN):
	@echo "Creating output directory at ./bin/"
	@mkdir -p $(PACKAGEBIN)

$(USERBIN):
	@echo "Creating output directory at ~/bin/"
	@mkdir -p $(USERBIN)
	
	
build: $(PACKAGEBIN)
	@echo "Building go-pwdstore"
	@go build -o $(PACKAGEBIN)/$(EXECUTABLE) $(MAIN)

install-user: $(USERBIN)
	@echo "Building go-pwdstore"
	@go build -o $(USERBIN)/$(EXECUTABLE) $(MAIN)
	
