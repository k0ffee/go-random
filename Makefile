targets = jar shuffle map
suffix  = .go
tests   = tests/test-*

.PHONY: all
all: $(targets)

$(targets):
	go build $(@)$(suffix)

.PHONY: chmod test tests
chmod:
	chmod +x $(tests)

test: $(tests)

$(tests): $(targets) chmod
	$@ $(targets)

.PHONY: clean
clean:
	rm -f $(targets)
