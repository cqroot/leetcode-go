.PHONY: test
test:
	@# Make sure these lines in the template are modified
	@sh -c "if grep -rn 'TestAdd(t \*testing.T)' solutions/; then exit 1; fi"
	@sh -c "if grep -rn '^// Name$$' solutions/; then exit 1; fi"
	go test -v ./...
