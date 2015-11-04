# shikaku
:mahjong: Shikaku game

[![Build Status](https://travis-ci.org/moul/shikaku.svg?branch=master)](https://travis-ci.org/moul/shikaku)
[![GoDoc](https://godoc.org/github.com/moul/shikaku?status.svg)](https://godoc.org/github.com/moul/shikaku)
[![Coverage Status](https://coveralls.io/repos/moul/shikaku/badge.svg?branch=master&service=github)](https://coveralls.io/github/moul/shikaku?branch=master)

Solution for [shikaku challenge](https://github.com/jeannedhack/programmingChallenges/tree/master/shikaku)

## Examples

```console
$ shikakugen -W 8 -H 4 -B 5
T 8 4
2 3 2
10 4 0
6 5 2
9 1 0
5 3 3
```

```console
$ shikakugen -W 4 -H 4 -B 5 --draw-solution --no-machine-output
+---+---+---+---+
|   |           |
+   +---+---+---+
|   |       |   |
+   +   +   +   +
|   |       |   |
+---+---+---+   +
|           |   |
+---+---+---+---+
```

```console
$ time shikakugen -W 4 -H 4 -B 5 --draw-map --draw-solution
T 4 4
3 3 1
3 2 0
4 2 1
3 0 1
3 0 3

+---+---+---+---+
|   |           |
+   +---+---+---+
|   |       |   |
+   +   +   +   +
|   |       |   |
+---+---+---+   +
|           |   |
+---+---+---+---+

+---+---+---+---+
|         3     |
+   +   +   +   +
| 3       4   3 |
+   +   +   +   +
|               |
+   +   +   +   +
| 3             |
+---+---+---+---+
shikakugen -W 4 -H 4 -B 5 --draw-map --draw-solution  0.00s user 0.00s system 70% cpu 0.005 total
```

```console
$ time shikakugen -W 20 -H 20 -B 42 --draw-map --draw-solution
T 20 20
12 18 5
5 13 11
15 10 2
12 3 12
8 2 7
8 8 12
6 8 0
8 8 4
4 6 14
6 11 8
10 14 0
21 15 17
15 15 8
9 16 11
21 13 18
5 3 16
12 3 1
8 10 19
4 8 9
9 2 10
4 16 3
10 2 18
6 6 12
10 3 2
10 9 17
4 2 9
5 1 17
5 12 7
12 18 2
8 6 16
12 6 8
4 8 6
10 4 2
8 1 5
9 1 14
9 12 9
15 13 2
10 7 2
15 8 16
10 5 10
6 0 0
20 19 17

+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
| 6                               6                       10                    |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|             12                                                                |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|             10  10          10          15          15                  12    |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                                                 4             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                 8                                             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|     8                                                                   12    |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                 4                                             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|         8                                       5                             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                         12                  6               15                |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|         4                       4               9                             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|         9           10                                                        |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                                     5           9             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|             12          6       8                                             |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                                                               |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|     9                   4                                                     |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                                                               |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|             5           8       15                                            |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|     5                               10                      21              20|
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|         10                                          21                        |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                                         8                                     |
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+

+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
|       |                       |           |                   |   |           |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|       |                       |           |                   |   |           |
+   +   +---+---+---+---+---+---+---+---+---+---+---+---+---+---+   +   +   +   +
|       |       |       |       |       |           |           |   |           |
+---+---+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|       |       |       |       |       |           |           |   |           |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +---+---+---+---+
|       |       |       |       |       |           |           |               |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|       |       |       |       |       |           |           |               |
+   +   +   +   +   +   +   +   +---+---+   +   +   +   +   +   +   +   +   +   +
|       |       |       |       |       |           |           |               |
+---+---+---+---+---+---+---+---+   +   +---+---+---+---+---+---+---+---+---+---+
|               |               |       |                   |                   |
+   +   +   +   +   +   +   +   +---+---+---+---+---+---+---+   +   +   +   +   +
|               |               |       |       |           |                   |
+---+---+---+---+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|               |               |       |       |           |                   |
+---+---+---+---+---+---+---+---+---+---+   +   +   +   +   +---+---+---+---+---+
|           |       |                   |       |           |           |       |
+   +   +   +   +   +   +   +   +   +   +---+---+---+---+---+   +   +   +   +   +
|           |       |                   |                   |           |       |
+   +   +   +   +   +---+---+---+---+---+---+---+---+---+---+   +   +   +   +   +
|           |       |           |               |           |           |       |
+---+---+---+   +   +   +   +   +   +   +   +   +   +   +   +---+---+---+   +   +
|           |       |           |               |           |           |       |
+   +   +   +   +   +---+---+---+---+---+---+---+   +   +   +   +   +   +   +   +
|           |       |       |                   |           |           |       |
+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|           |       |       |                   |           |           |       |
+---+---+---+---+---+---+---+   +   +   +   +   +   +   +   +   +   +   +   +   +
|                   |       |                   |           |           |       |
+---+---+---+---+---+   +   +---+---+---+---+---+   +   +   +   +   +   +   +   +
|                   |       |                   |           |           |       |
+---+---+---+---+---+   +   +   +   +   +   +   +   +   +   +   +   +   +   +   +
|                   |       |                   |           |           |       |
+   +   +   +   +   +   +   +---+---+---+---+---+---+---+---+   +   +   +   +   +
|                   |       |                               |           |       |
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
        0.52 real         0.58 user         0.03 sys
```

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
