![CI](https://github.com/loveyadav8478-blip/Golang/actions/workflows/ci.yml/badge.svg)

# Golang

A hands-on collection of Go programs I wrote while learning the language. Each folder focuses on one core concept, with small runnable examples you can read, run and tweak.

## Topics

| Folder | Concept |
| --- | --- |
| [`Basics`](./Basics) | Syntax, variables, constants, data types, control flow, functions |
| [`Slices`](./Slices) | Dynamic arrays, `append`, slicing, `len` / `cap` |
| [`Maps`](./Maps) | Key-value collections, lookups, insertion, deletion, iteration |
| [`struct`](./struct) | Custom types and grouping related data |
| [`methods`](./methods) | Attaching behaviour to types (value vs. pointer receivers) |
| [`pointers`](./pointers) | Memory addresses, dereferencing, pass-by-value vs. pass-by-reference |
| [`interface`](./interface) | Abstraction and polymorphism through interfaces |
| [`errorHandling`](./errorHandling) | Idiomatic `error` values and handling patterns |
| [`closers`](./closers) | Closures and functions as first-class values |
| [`variadicFunction`](./variadicFunction) | Functions that accept a variable number of arguments |

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed (see `go.mod` for the version used)

Check your install:

```bash
go version
```

### Clone the repo

```bash
git clone https://github.com/loveyadav8478-blip/Golang.git
cd Golang
```

### Run an example

Move into any topic folder and run the Go file(s) inside it:

```bash
cd Slices
go run .
```

Or run a single file directly:

```bash
go run Slices/main.go
```

> If a folder contains multiple `main` packages or files, run the specific file you want with `go run <file>.go`.

## Project Structure

```
Golang/
├── Basics/
├── Maps/
├── Slices/
├── closers/
├── errorHandling/
├── interface/
├── methods/
├── pointers/
├── struct/
├── variadicFunction/
└── go.mod
```

## Module

```
module my-go-app
```

## Purpose

This repo is a personal learning log. The code is written for clarity and practice rather than production use.

## Author

**Love Yadav** - [@loveyadav8478-blip](https://github.com/loveyadav8478-blip)