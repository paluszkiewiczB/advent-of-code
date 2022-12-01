# [Advent of Code 2022](https://adventofcode.com/2022/)

## Prerequisites

- [Go 1.19.X](https://go.dev/dl/) installed
- `go` binary on the classpath

## Structure

Each directory stores task and its solution
```text
repository root
|- README.md
|- 01
| |- input.txt
| |- main.go
| |- go.mod
| |- go.sum
| |- task.md
|- 02
| |- input.txt
| |- main.go
| |- go.mod
| |- go.sum
| |- task.md
```
- task.md - description of the task
- main.go - implementation of the solution
- input.txt - input for the task
- go.mod, go.sum - files describing Go module's properties (not needed if you don't have dependencies), see [official reference](https://go.dev/doc/modules/gomod-ref)

## Run

First day:

```shell
cd 01
go run main.go
```

Answers are printed to stdout