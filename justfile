# use with https://github.com/casey/just

_default:
    @just --list --unsorted --justfile {{ justfile() }}

# Clears the cache
clean-cache:
    go clean -cache

# Clears the test cache and runs the tests
test:
    go clean -testcache
    go test -v ./...
