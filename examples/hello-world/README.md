# hello-world

An example server for the `service-skeleton` project.

## Building

```shell
# substitute the version number as desired
go build -ldflags "-X main.Version=0.1.0
```

## Usage

```
Usage: hello-world [--version] [--help] <command> [<args>]

Available commands are:
    server     Server command
    version    Return the version of the binary
```