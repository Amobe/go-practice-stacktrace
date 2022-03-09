# Tracing error stack in practice
---

Go 1.13 supports a [error wrapping feature](https://go.dev/blog/go1.13-errors). This project aims to use a tool ([romanyx/stack](https://github.com/romanyx/stack)) that helping us trace the error stack. This tools is compatible with the functions in builtin package `errors`.

## Comparing the output error messaging 

In `cmd/original/main.go`, the error is wrapping by the builtin package `fmt`.
In `cmd/stacktrace/main.go`, the error is wrapping by the stacktrace tool.

### Output of original version

Executing command:
```bash
go run cmd/original/main.go
```

Output:
```
ERROR: handle hello failed: service world failed: hello debug world
```

### Output of stackstace version

Executing command:
```bash
go run cmd/stacktrace/main.go
```

Output:
```
ERROR: handle hello failed: service world failed: hello debug world

${workspace}/practice-stacktrace/pkg/service/service.go:25 practice-stacktrace/pkg/service.(*service).WorldStacktrace()
${workspace}/practice-stacktrace/pkg/handler/handler.go:34 practice-stacktrace/pkg/handler.(*handler).HelloStacktrace()
${workspace}/practice-stacktrace/cmd/stacktrace/main.go:30 main.mainWithError()
${workspace}/practice-stacktrace/cmd/stacktrace/main.go:12 main.main()
${GOROOT}/src/runtime/proc.go:255 runtime.main()
${GOROOT}/src/runtime/asm_amd64.s:1571 runtime.goexit()
```

> \${workspace} and \${GOROOT} would be different in differernt machine.
