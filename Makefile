.PHONY: run
run:
	go run cmd/main.go

go_files =$(shell find . -name '*.go')
results.md:$(go_files)
	go run cmd/main.go > $@

README.md: readme_template/README.md results.md
	cat $^ > $@
.PHONY: README.md