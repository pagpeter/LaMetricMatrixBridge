# Matrix -> LaMetric Time bridge

This small golang app notifies you about new messages on your LaMetric Time. This should be run on a Raspberry Pi or something similiar, but can also be run on your main PC.

It works on Windows, MacOS and Linux.

## Installation

You need to make sure that you have `go` installed. 

```sh
git clone https://github.com/wwhtrbbtt/LaMetricMatrixBridge
cd LaMetricMatrixBridge/
go build -o builds/bridge *.go
```

The binary can be run using
```sh
./builds/bridge
```

## TODO

    - [] Add e2ee support
    - [] Make message customizable
    - [] Add prebuilt binaries