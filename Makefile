.PHONY: lib fmt install build-examples build-x-examples clean verify

lib:
	rm -rf ./lib
	rm -rf ./sdk
	git clone https://github.com/Cycling74/max-sdk-base.git sdk
	mkdir -p lib/max lib/msp
	cp -r sdk/c74support/max-includes/ ./lib/max
	cp -r sdk/c74support/msp-includes/ ./lib/msp
	rm -rf sdk

fmt:
	go fmt ./...
	go vet ./...
	golint ./...
	clang-format  -style "{BasedOnStyle: Google, ColumnLimit: 120}" -i *.c *.h

install:
	go install ./cmd/maxgo


build-examples: install
	@make -C examples/example1 install
	@make -C examples/example2 install

build-x-examples: install
	cd example; maxgo -name maxgo -cross -install maxgo

clean:
	@make -C examples/example1 clean
	@make -C examples/example2 clean


verify:
	clang verify/verify.c -o verify/verify
	cd verify; ./verify
