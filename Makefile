.PHONY: build install clean

BINARY_NAME=amc
INSTALL_DIR=$(HOME)/bin

build:
	@mkdir -p $(INSTALL_DIR)
	go build -o $(INSTALL_DIR)/$(BINARY_NAME) .

clean:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

install: build
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)/$(BINARY_NAME)"
	@echo ''
	@echo 'IMPORTANT: Add ~/bin to your PATH by adding this line to ~/.zshrc or ~/.bashrc:'
	@echo '  export PATH="$$HOME/bin:$$PATH"'
