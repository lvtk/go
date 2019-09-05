all:
	c-for-go lv2.yml 

clean:
	rm -f lv2/cgo_helpers.go lv2/cgo_helpers.h lv2/cgo_helpers.c
	rm -f lv2/const.go lv2/doc.go lv2/types.go
	rm -f lv2/lv2.go

test:
	cd lv2 && go build

install:
	cd lv2 && go install
