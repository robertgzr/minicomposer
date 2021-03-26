
.DEFAULT_GOAL := test

.PHONY: test
test:
	@docker build -t minicomposer .
	@docker run --rm --privileged --mount=type=bind,src=$(PWD),dst=/src minicomposer
	@docker rmi minicomposer
