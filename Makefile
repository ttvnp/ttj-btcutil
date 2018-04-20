.PHONY: deps
deps: glide
	glide up

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif
