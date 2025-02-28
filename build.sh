#!/usr/bin/env bash

cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./wasm_exec.js
GOOS=js GOARCH=wasm tinygo build -o wasm.wasm ./main.go
