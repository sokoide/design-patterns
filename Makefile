DIRS = $(sort $(wildcard *-example))

.PHONY: run

run:
	@for dir in $(DIRS); do \
		echo "========================================"; \
		echo "Running $$dir..."; \
		echo "========================================"; \
		(cd $$dir && go run .); \
		echo ""; \
	done
