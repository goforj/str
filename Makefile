.PHONY: help

HELP_FUN = %help; while (<>) { /^([A-Za-z0-9_-]+)\s*:.*\#\#(?:@([A-Za-z0-9_-]+))?\s(.*)$$/ or next; push @{$$help{$$2 || "other"}}, [$$1, $$3]; $$width = length($$1) if length($$1) > $$width } print "\n"; for $$category (sort keys %help) { print "\e[37m$$category\e[0m\n"; for $$entry (@{$$help{$$category}}) { printf "  \e[33m%-*s\e[0m  \e[32m%s\e[0m\n", $$width, $$entry->[0], $$entry->[1] } }

help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

##@quality
test: ##@quality Run the test suites.
	go test ./...
	go -C docs test ./...
	go -C examples test ./...

test-race: ##@quality Run the race-enabled test suite.
	go test -race ./...

vet: ##@quality Run Go vet for every module.
	go vet ./...
	go -C docs vet ./...
	go -C examples vet ./...

##@documentation
generate: ##@documentation Regenerate documentation examples and README content.
	go -C docs run ./examplegen
	go -C docs run ./readme
