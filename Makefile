.PHONY: test
test:
	@# Make sure these lines in the template are modified
	@sh -c "if grep -rn 'TestAdd(t \*testing.T)' 0*/; then exit 1; fi"
	@sh -c "if grep -rn 'TestAdd(t \*testing.T)' 9*/; then exit 1; fi"
	@sh -c "if grep -rn '^// Name$$' 0*/; then exit 1; fi"
	@sh -c "if grep -rn '^// Name$$' 9*/; then exit 1; fi"
	go test -v ./...

.PHONY: check
check:
	golangci-lint run
	@echo
	if [ -n "$(gofumpt -l .)" ]; then gofumpt -l .; $(error need formatting); exit 1; fi

.PHONY: tree
tree:
	exa --only-dirs --long --header --git --icons --tree --level=4 --level 2
