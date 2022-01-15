# Matrix -> LaMetric Time bridge

This small golang app notifies you about new messages on your [LaMetric Time](https://lametric.com/en-US). This should be run on a Raspberry Pi or something similiar, but can also be run on your main PC.

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

## Setup

1. Rename the `config.example.yaml` file to `config.yaml`
2. Fill out the values
3. (Optionally) set up a whitelist or a blacklist, to exclude big rooms or to get notified about some people.

## TODO

    - [] Add e2ee support
    - [] Make message customizable
    - [] Add prebuilt binaries