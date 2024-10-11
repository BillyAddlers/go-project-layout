
# NOTE: Welcome to Makefile, entry point for your Build Tools!
#
# This Makefile serves as a simple demonstration of how the build process works.
# In more complex projects, you'll want to encapsulate commands within this Makefile
# and utilize the `/scripts` directory to streamline your project's build and run operations.
#
# Be aware that some commands interact with `/tmp` directory.
# If your operating system does not support `/tmp`, lacks access to it,
# or requires a different path, please adjust the configuration accordingly.
# Remember to set `main_package_path` to your specific entry point.
# For example, `main_package_path = ./cmd/example` indicates that your entry point is located at `./cmd/example/main.go`.
main_package_path = .
binary_name = example

# For more detailed documentation, please visit https://www.gnu.org/software/make/manual/html_node/Introduction.html

# NOTE: Helpers command
# Used to display available subcommands for `make`

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)"

# NOTE: Quality control command
# Used to check for code errors and run test unit

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out


# NOTE: Development command
# Do some dependencies tidy and build project/run project live

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	# Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=/tmp/bin/${binary_name} ${main_package_path}

.PHONY: build/dev
## build/dev: run make build with DEBUG env set to true
build/dev:
	DEBUG=true make build

## run: run the  application
.PHONY: run
run: build
	/tmp/bin/${binary_name}

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/${binary_name}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


# NOTE: Operations command
# Do some version control and push commit to Git repository
# Also handle production deployment

## push: push changes to the remote Git repository
.PHONY: push
push: confirm audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: confirm audit no-dirty
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=/tmp/bin/linux_amd64/${binary_name} ${main_package_path}
	upx -5 /tmp/bin/linux_amd64/${binary_name}
	# Include additional deployment steps here...

# NOTE: Example command
# Command below is the sample of how Makefile should be
# Utilizing shell script inside `./scripts` directory

## hello-world: print hello world
.PHONY: hello-world
hello-world:
	./scripts/helloworld.sh
