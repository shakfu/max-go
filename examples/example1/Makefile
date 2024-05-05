ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PACKAGE_NAME := maxgo-examples
PACKAGE := "$(HOME)/Documents/Max 8/Packages/$(PACKAGE_NAME)"
NAME := example1


.PHONY: all build install xbuild xinstall clean reset


all: build

build:
	@maxgo -name $(NAME) -out dist

install:
	@maxgo -name $(NAME) -out dist -install $(PACKAGE_NAME)
	@mkdir -p $(PACKAGE)/patchers
	#@cp -f $(ROOT_DIR)/*.maxpat $(PACKAGE)/patchers
	@if [ -f $(ROOT_DIR)/$(NAME).maxhelp ] ; then \
    	mkdir -p $(PACKAGE)/help; \
		cp -f $(ROOT_DIR)/$(NAME).maxhelp $(PACKAGE)/help; \
	fi;

xbuild:
	@maxgo -name $(NAME) -out dist -cross

xinstall:
	@maxgo -name $(NAME) -out dist -cross -install $(PACKAGE_NAME)
	@cp $(ROOT_DIR)/dist/$(NAME).mxe64 $(PACKAGE)/externals
	@mkdir -p $(PACKAGE)/patchers
	@cp -f $(ROOT_DIR)/*.maxpat $(PACKAGE)/patchers
	@if [ -f $(ROOT_DIR)/$(NAME).maxhelp ] ; then \
    	mkdir -p $(PACKAGE)/help; \
		cp -f $(ROOT_DIR)/$(NAME).maxhelp $(PACKAGE)/help; \
	fi;

clean:
	@rm -rf $(ROOT_DIR)/dist $(ROOT_DIR)/out

reset:
	@rm -rf $(ROOT_DIR)/dist $(ROOT_DIR)/out $(PACKAGE)

