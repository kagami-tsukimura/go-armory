#!/usr/bin/env bash
set -euo pipefail

REPO="kagami-tsukimura/go-armory"
# GitHubの "latest" URLからリダイレクト先のタグ名を取得する
echo "Fetching latest version tag..."
VERSION=$(basename $(curl -Ls -o /dev/null -w %{url_effective} "https://github.com/$REPO/releases/latest"))
echo "${VERSION} fetched."

# Install location under user's home
BIN_DIR="${HOME}/.local/bin"

# Create the directory if it does not exist
mkdir -p "$BIN_DIR"

install() {
  cmd=$1
  base="https://github.com/$REPO/releases/download/$VERSION"
  file="${cmd}-$(uname | tr '[:upper:]' '[:lower:]')-amd64"
  url="$base/$file"
  dest="$BIN_DIR/$cmd"

  echo "Downloading $cmd ..."
  curl -L -o "$dest" "$url"
  chmod +x "$dest"
}

if [ $# -eq 0 ]; then
  # Default: install all commands
  for cmd in $(ls cmd); do
    install "$cmd"
  done
else
  # Install only the specified command
  install "$1"
fi

echo "Done. Make sure $BIN_DIR is in your PATH:"
echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
