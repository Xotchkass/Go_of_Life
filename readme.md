
# Go of Life

Simple implementation of Conway's Game of Life written in Go using the [raylib](https://www.raylib.com/) library.

## Requirements

- [Go](https://golang.org/doc/install)


## Installation


```shell
go install github.com/Xotchkass/Go_of_Life@latest
Go_of_Life
```

or

```shell
go clone https://github.com/Xotchkass/Go_of_Life
cd Go_of_Life
go build
./Go_of_Life
```


## Controls

- `Left mouse button` to paint a live cells
- `Right mouse button` to paint dead cells
- `R` key randomizes the grid
- `C` key clears the grid
- `N` key advances simulation a single generation if game is paused
- `Spacebar` pauses/resumes simulation
- `+`/`-` keys to increase/decrease simulation speed
- `Esc` key to quit

## Credits

- [raysan](https://github.com/raysan5) for [raylib](https://github.com/raysan5/raylib)
- [Milan Nikolic](https://github.com/gen2brain)  for [Raylib Go bindings](https://github.com/gen2brain/raylib-go)
