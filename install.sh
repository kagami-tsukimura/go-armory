#!/usr/bin/env bash
set -euo pipefail

REPO="kagami-tsukimura/go-armory"
VERSION="${1:-latest}"

BIN_DIR="/usr/local/bin"

# 書き込み可能かどうか判定
if [ -w "$BIN_DIR" ]; then
  SUDO=""
else
  SUDO="sudo"
fi

install() {
  cmd=$1
  base="https://github.com/$REPO/releases/download/$VERSION"
  file="${cmd}-$(uname | tr '[:upper:]' '[:lower:]')-amd64"
  url="$base/$file"
  dest="$BIN_DIR/$cmd"

  echo "Downloading $cmd ..."

  # sudo いる/いらないを自動切替
  $SUDO curl -L -o "$dest" "$url"
  $SUDO chmod +x "$dest"
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
