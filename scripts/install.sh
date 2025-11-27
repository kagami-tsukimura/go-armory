#!/usr/bin/env bash
set -euo pipefail

REPO="kagami-tsukimura/go-armory"
VERSION="${1:-latest}"

install() {
  cmd=$1
  base="https://github.com/$REPO/releases/download/$VERSION"
  file="${cmd}-$(uname | tr '[:upper:]' '[:lower:]')-amd64"

  echo "Downloading $cmd ..."
  curl -L -o /usr/local/bin/$cmd "$base/$file"
  chmod +x /usr/local/bin/$cmd
}

if [ $# -eq 0 ]; then
  # デフォルト：全部インストール
  for cmd in $(ls cmd); do
    install "$cmd"
  done
else
  # 引数指定：そのコマンドのみ
  install "$1"
fi
