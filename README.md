# Key-Value Store

This project is a simple key-value store implemented in GoLang.

## Features

- Store and retrieve key-value pairs
- In-memory storage for fast access
- Simple and easy-to-use API

## Installation

To install the project, clone the repository and build the project:

```sh
git clone https://github.com/yourusername/key-value-store.git
cd key-value-store
go build
```

## Usage

To start the key-value store server, run the following command:

```sh
./key-value-store
```

You can interact with the server using HTTP requests. Here are some examples:

### Store a value

```sh
curl -X POST http://localhost:9000/set -d '{"key": "name", "value": "John"}'
```

### Retrieve a value

```sh
curl http://localhost:9000/get?key=name
```

## Original Author

This project was inspired by a tutorial from [Stephen Mwangi's blog](https://www.stephenmwangi.com/blog/in-memory-key-value-store/).
