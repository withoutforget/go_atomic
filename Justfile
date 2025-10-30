[doc("Running golang application")]
@run:
    go run main.go ./src/asm.s
@rbuild:
    go build && ./go_atomic