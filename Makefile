all:
	c-for-go lv2.yml

clean:
	# rm -f lv2/cgo_helpers.go lv2/cgo_helpers.h lv2/cgo_helpers.c
	# rm -f lv2/const.go lv2/doc.go lv2/types.go
	# rm -f lv2/lv2.go

	rm -f pugl/cgo_helpers.go pugl/cgo_helpers.h pugl/cgo_helpers.c
	rm -f pugl/const.go pugl/doc.go pugl/types.go
	rm -f pugl/pugl.go

test:
	cd lv2 && go build
