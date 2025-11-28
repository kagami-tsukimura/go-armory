# go-armory

A collection of simple and useful CLI tools written in Go.

## Installation

There are several ways to install the tools.

### Using the `install.sh` script

You can install all available commands by cloning the repository and running the installation script.

The script installs the commands to **`$HOME/.local/bin`**.

```sh
git clone https://github.com/kagami-tsukimura/go-armory.git
cd go-armory
# Run the script (It automatically uses the latest version)
./install.sh
```

#### Option 1: From GitHub Releases

Binaries for each OS are automatically built and published on the [Releases page](https://github.com/kagami-tsukimura/go-armory/releases) whenever a new version is tagged.

#### Option 2: Manual Download

You can also download a specific binary directly from the [Releases page](https://github.com/kagami-tsukimura/go-armory/releases).

Choose the correct file for your OS and architecture, place it in a directory included in your `PATH` (e.g., `/usr/local/bin` or `~/bin`), and make it executable.

```sh
# Example: Install the 'br' command for linux (amd64) to /usr/local/bin
# Please replace VX.X.X with the latest version number.
VERSION="vX.X.X"
curl -L -o /usr/local/bin/br https://github.com/kagami-tsukimura/go-armory/releases/download/${VERSION}/br-linux-amd64
chmod +x /usr/local/bin/br
```

### From Source (For Go Developers)

If you have a Go development environment set up, you can install the commands directly using `go install`.

```sh
# To install the 'br' command
go install github.com/kagami-tsukimura/go-armory/cmd/br@latest
```

The binary will be installed in your `$GOPATH/bin` (or `$HOME/go/bin`). Ensure this directory is in your `PATH`.
