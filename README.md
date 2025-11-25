# go-armory

A collection of simple and useful CLI tools written in Go.

## Installation

There are two ways to install and use the tools.

### From Source (For Go Developers)

If you have a Go development environment set up, you can install the commands directly.

#### 1. Installing from remote repository

You can install the command directly from GitHub repository.

```sh
# Install the 'br' command
go install github.com/kagami-tsukimura/go-armory/cmd/br@latest
```

The binary will be installed in your `$GOPATH/bin` (or `$HOME/go/bin`). Make sure this directory is in your `PATH`.

## Available Commands

### br

A simple tool to list all branches in a Git repository.

#### Usage

- Run in the current directory:

  ```sh
  br
  ```
