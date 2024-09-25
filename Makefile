DIR := $(shell pwd)
SRC_DIR := $(DIR)/src
BIN_DIR := $(DIR)/bin

OUTPUT := main

# Компіляція проекту
all: $(BIN_DIR)/$(OUTPUT)

# Створюємо директорію bin, якщо її немає, і компілюємо програму
$(BIN_DIR)/$(OUTPUT): $(SRC_DIR)/*.go
	mkdir -p $(BIN_DIR)
	cd $(SRC_DIR) && go build -o $(BIN_DIR)/$(OUTPUT)

# Очищення згенерованих файлів
clean:
	rm -rf $(BIN_DIR)/$(OUTPUT)

# Повна очистка всіх скомпільованих файлів і каталогів
clean-all:
	rm -rf $(BIN_DIR)