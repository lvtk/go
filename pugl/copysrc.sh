#!/bin/bash
# copies and renames source files to pkg directories so go build picks
# them up per platform
set -ex
rm -f *.c *.m
rm -f gl/*.c gl/*.m

cp -f src/pugl/detail/implementation.c ./implementation.c

cp -f src/pugl/detail/mac.m ./mac_darwin.m
cp -f src/pugl/detail/mac_gl.m ./gl/mac_gl_darwin.m

cp -f src/pugl/detail/x11.c ./x11_linux.c
cp -f src/pugl/detail/x11_gl.c ./gl/x11_gl_linux.c

cp -f src/pugl/detail/win.c ./win_windows.c
cp -f src/pugl/detail/win_gl.c ./gl/win_gl_windows.c
