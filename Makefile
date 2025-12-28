DIRS = $(sort $(wildcard *-example))

.PHONY: run format test

run:
	@for dir in $(DIRS); do \
		echo "========================================"; \
		echo "Running $$dir..."; \
		echo "========================================"; \
		(cd $$dir && go run .); \
		echo ""; \
	done

format:
	@for dir in $(DIRS); do \
		echo "========================================"; \
		echo "Formatting $$dir..."; \
		echo "========================================"; \
		find "$$dir" -name '*.go' -not -path '*/vendor/*' -not -path '*/third_party/*' -print0 | xargs -0 gofmt -w; \
		echo ""; \
	done

test:
	@for dir in $(DIRS); do \
		echo "========================================"; \
		echo "Testing $$dir..."; \
		echo "========================================"; \
		(cd $$dir && go test -v ./...); \
		echo ""; \
	done
