# Tower

Tower defense in Golang

## Purpose

A little game implemented as a personnal exercise.

## Usage

1. clone the project locally
2. from the project folder, execute: `go run main.go`

## Features

### Map

- walls are placed loading map_0.txt
- borders, input and output are hardcoded
- enemy path is hardcoded

### Tower placement

Send a web request to `http://localhost:8080/tower/` to place a tower.

The tower coordinates are shown in the disaplyed grid : the top left corner is `a1` and the bottom right is `E30`.

#### Example

To place a tower in the middle of the grid, send the following request: `http://localhost:8080/tower/o16`


## Incoming

- avoid upgrading full-level tower
- increase price of tower upgrade
- manage counter of tower kills
- manage log display