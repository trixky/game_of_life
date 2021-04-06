# game_of_life

An implementation of the [game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life), the cellular automaton devised by the British mathematician [John Horton Conway](https://en.wikipedia.org/wiki/John_Horton_Conway) in 1970, using [go](https://golang.org/) with the [sdl2](https://github.com/veandco/go-sdl2) package. __(linux)__

![Recordit GIF](https://github.com/trixky/game_of_life/blob/main/demo/demo.gif)

> This program generates patterns of 50 by 50.

## Rules

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, live or dead, (or populated and unpopulated, respectively). Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.


## Usage

> The file must contain 50 lines of 50 characters and end with a \n

```
go build -o game_of_life ./*.go
./game_of_life -pattern=default.pattern
```
