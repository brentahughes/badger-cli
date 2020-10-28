BOILERPLATE_FSPATH=../boot/boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/gitr.mk

run:
	go run main.go