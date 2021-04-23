[![github.com/iuthere/slowproverb](./doc/gobadge.svg)](https://pkg.go.dev/github.com/iuthere/slowproverb)

`slowproverb` is a silly Go proverb command line utility that emits a random Go Proverb (taken mostly from the Gopherfest SV 2015) in a manner that allows to test out different terminal proxies that accumulate outputs from different running in the background tasks due to internal delay and usage of `\r` for showing a progress and `\n` to split a proverb into several lines. The output lasts between 2 and 5 seconds initially, which can be adapted.

Additionally, `slowproverb` randomly chooses whether to use `Stdout` or `Stderr`.

## Install

`$ go install github.com/iuthere/slowproverb@latest`

## Run

`$ slowproverb`

## Options

Different command line options allow to modify the way the output flows:

```shell
-link
    include source web link (default true)
-max uint
    max seconds (default 5)
-min uint
    min seconds (default 2)
-width uint
    text width (default 20)
```

## Screenshot

![slowproverb](https://user-images.githubusercontent.com/8169082/115772765-7ee35800-a37d-11eb-9d6d-b2911b1f7d07.gif)

