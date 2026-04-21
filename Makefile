.PHONY: build install clean

BINARY_NAME=atg
# 默认安装到系统推荐的第三方二进制目录
INSTALL_DIR ?= /usr/local/bin
BUILD_DIR=./bin

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@mkdir -p $(INSTALL_DIR)
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Successfully installed $(BINARY_NAME) to $(INSTALL_DIR)/$(BINARY_NAME)"
	@echo "You can now run '$(BINARY_NAME)' from anywhere."

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
