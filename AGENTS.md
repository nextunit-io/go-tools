# AGENTS.md

This file provides guidance to AI coding agents when working with code in this repository.

## Repository layout

Three independent Go modules (no root `go.mod`, `go.work` is gitignored):

- `interfaces/` — `github.com/nextunit-io/go-tools/interfaces`. Shared interfaces with no dependencies (e.g. `OsFileInterface`, mirroring `*os.File`).
- `tools/` — `github.com/nextunit-io/go-tools/tools`. Thin wrappers around stdlib (`os`, `os/exec`, `net/http`, `time`, env) plus generic helpers in `utils.go`.
- `toolsmock/` — `github.com/nextunit-io/go-tools/toolsmock`. Mocks for every `tools` interface, built on `github.com/nextunit-io/go-mock`.

`tools` depends on `interfaces` and (for tests) on `toolsmock`, both pinned as pseudo-versions in `tools/go.mod`. A change to `interfaces` or `toolsmock` is only visible to `tools` after it is pushed and the pseudo-version in `tools/go.mod` is bumped (or via a local `go.work` / `replace`, which must not be committed).

## Commands

Run commands inside the module directory; there is no root module.

```sh
cd tools && go test ./...
cd toolsmock && go test ./...
cd tools && go test -run TestOwnEnvClient ./...   # single test
```

CI (`.github/workflows/test.yml`, `test_mock.yml`) runs `go test ./...` in `tools/` and `toolsmock/` separately, triggered only by pushes to `main` that touch the matching path (no PR trigger). `interfaces/` has no tests and no CI.

## Architecture pattern

Each wrapper in `tools/` follows the same swappable-singleton shape so callers can inject mocks:

- An interface (`Env`, `OsInterface`, `Time`, `ExecInterface`, ...).
- An unexported `default<Name>Client` that embeds the interface and delegates to stdlib.
- A package-level instance with `Get<Name>Instance()` (lazily creates the default) and `Set<Name>Instance(client)` to override it (the env setter is named `SetEnvGetInstance`). Exception: `Http` uses a creator function (`GetHttpInstance(cfg)` / `SetHttpCreator`) because each client takes an `HttpConfig` (Bearer auth or AWS SigV4 signing via the default AWS config).
- Functions returning files return `interfaces.OsFileInterface`, not `*os.File`, so file mocks can be substituted.

Mocks in `toolsmock/` mirror this one-to-one:

- `Get<Name>Mock()` returns a struct whose `Mock` field holds one `*gomock.ToolMock[InputStruct, ResultType]` per method, created with `gomock.GetMock[...](fmt.Errorf("<Method> general error"))`.
- Each mock method records inputs with `AddInput(struct{...})` and returns `GetNextResult()`. Tests queue results with `AddReturnValue` and inspect calls with `HasBeenCalled()` / `GetInput(i)`.
- Methods that only return `error` use `bool` as the result type: `true` makes the mock return `fmt.Errorf("Error requested")`.
- Methods with no error return (e.g. `Getenv`) panic when no result is queued.

Adding a method therefore means touching: the interface and default client in `tools/`, the mock struct, constructor and method in `toolsmock/`, and tests in both modules (tests use external `_test` packages and `testify/assert`).
