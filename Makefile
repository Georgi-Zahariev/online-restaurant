.PHONY: build test run clean

BINARY_NAME = online-restaurant
BACKEND_DIR = backend

build:
    @echo "Building application..."
    @mkdir -p bin
    @go build -o bin/$(BINARY_NAME) $(BACKEND_DIR)

test:
    @echo "Running tests..."
    @go test -v $(BACKEND_DIR)/...

run: build
    @echo "Starting application locally..."
    @PORT=8080 ENV=development LOG_FORMAT=text LOG_SEVERITY=debug ./bin/$(BINARY_NAME)

clean:
    @echo "Cleaning build artifacts..."
    @rm -rf bin