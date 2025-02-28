## Caution

This sample program consumes too much memory. Try carefully.
One pushing button will allocate 100mb memory.

## To reproduce error

Assuming you're using Chrome browser.

- Build and run HTTP server.

~~~sh
# Build
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" ./wasm_exec.js
GOOS=js GOARCH=wasm tinygo build -o wasm.wasm ./main.go

# Run HTTP Server
go run localserver/server.go
~~~

- Open http://localhost:8088 on your browser.
- Open `DevTools` and switch to `Console` tab.
- Push `Run` Button several times.
- After pushing button 22 times or so, an error message like the followings will appear:

~~~
wasm_exec.js:236 Uncaught RangeError: Start offset -2093352176 is outside the bounds of the buffer
    at new DataView (<anonymous>)
    at loadString (wasm_exec.js:236:27)
    at syscall/js.stringVal (wasm_exec.js:313:17)
    at main.syscall/js.ValueOf (wasm.wasm:0x12a21)
    at main.(:8088/syscall/js.Value).Set (http://localhost:8088/wasm.wasm)
    at main.(:8088/*main.hoge).incr$bound (http://localhost:8088/wasm.wasm)
    at main.runtime.resume$1 (wasm.wasm:0xe268)
    at main.runtime.resume$1$gowrapper (wasm.wasm:0x94af)
    at main.tinygo_launch (wasm.wasm:0x4e8)
    at main.runtime.scheduler (wasm.wasm:0x95b4)
~~~

Maybe the amount of memory exceeds upper limit of int32 (2_147_483_648 bytes).
