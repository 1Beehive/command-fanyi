#!/bin/bash

BINARIES_URL="https://github.com/jeffcail/command-fanyi/releases/download/v1.0.2/command-fanyi-darwin-arm64"

INSTALL_DIR="$HOME/command-fanyi"
TARGET_DIR="/usr/local/bin"
BIN_NAME="fanyi"

if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating directory $INSTALL_DIR..."
    mkdir -p "$INSTALL_DIR"
fi

if [ -e "$TARGET_DIR/$BIN_NAME" ] || [ -L "$TARGET_DIR/$BIN_NAME" ]; then
    echo "$BIN_NAME already exists in $TARGET_DIR. Removing existing file/link..."
    sudo rm -f "$TARGET_DIR/$BIN_NAME"
fi

echo "Downloading binary from GitHub..."
curl -L -o "$INSTALL_DIR/$BIN_NAME" "$BINARIES_URL"

if [ $? -eq 0 ]; then
    chmod +x "$INSTALL_DIR/$BIN_NAME"
    echo "Binary downloaded and made executable."
else
    echo "Failed to download the binary."
    exit 1
fi

echo "Creating symbolic link in /usr/local/bin..."
sudo ln -sf "$INSTALL_DIR/$BIN_NAME" "$TARGET_DIR/$BIN_NAME"

echo "Setup completed!"
