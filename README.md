# shikaku
:mahjong: Shikaku game

[![Build Status](https://travis-ci.org/moul/shikaku.svg?branch=master)](https://travis-ci.org/moul/shikaku)
[![GoDoc](https://godoc.org/github.com/moul/shikaku?status.svg)](https://godoc.org/github.com/moul/shikaku)
[![Coverage Status](https://coveralls.io/repos/moul/shikaku/badge.svg?branch=master&service=github)](https://coveralls.io/github/moul/shikaku?branch=master)

Solution for [shikaku challenge](https://github.com/jeannedhack/programmingChallenges/tree/master/shikaku)

## Usage

```console
$ shikakugen --help
Usage:
  shikakugen [OPTIONS]

Application Options:
  -W, --width=             Width of the grid
  -H, --height=            Height of the grid
  -B, --blocks=            Blocks in the grid
  -m, --draw-map           Draw the map in ascii-art
  -s, --draw-solution      Draw the solution in ascii-art
  -q, --no-machine-output  No machine output

Help Options:
  -h, --help               Show this help message

2015/11/04 23:16:45 Parsing error: Usage:
  shikakugen [OPTIONS]

Application Options:
  -W, --width=             Width of the grid
  -H, --height=            Height of the grid
  -B, --blocks=            Blocks in the grid
  -m, --draw-map           Draw the map in ascii-art
  -s, --draw-solution      Draw the solution in ascii-art
  -q, --no-machine-output  No machine output

Help Options:
  -h, --help               Show this help message
```

## Install

```console
$ go get github.com/moul/shikaku/cmd/shikakugen
```
## License

MIT
