.PHONY: build install clean

BINARY_NAME=amc
INSTALL_DIR=$(HOME)/bin
BUILD_DIR=./bin

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

install: build
	@mkdir -p $(INSTALL_DIR)
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)/$(BINARY_NAME)"
	@echo ''
	@echo 'IMPORTANT: Add ~/bin to your PATH by adding this line to ~/.zshrc or ~/.bashrc:'
	@echo '  export PATH="$$HOME/bin:$$PATH"'

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
