#!/bin/bash

# 设置 GitHub 上二进制文件的下载 URL
BINARIES_URL="https://github.com/jeffcail/command-fanyi/releases/download/v1.0.2/command-fanyi-darwin-arm64"

# 设置下载目录（$HOME/command-fanyi）
INSTALL_DIR="$HOME/command-fanyi"
TARGET_DIR="/usr/local/bin"
BIN_NAME="fanyi"

# 创建下载目录（如果不存在）
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating directory $INSTALL_DIR..."
    mkdir -p "$INSTALL_DIR"
fi

# 检查目标目录下是否已有同名软链接或文件
if [ -e "$TARGET_DIR/$BIN_NAME" ] || [ -L "$TARGET_DIR/$BIN_NAME" ]; then
    echo "$BIN_NAME already exists in $TARGET_DIR. Removing existing file/link..."
    sudo rm -f "$TARGET_DIR/$BIN_NAME"
fi

# 下载二进制文件到指定目录
echo "Downloading binary from GitHub..."
curl -L -o "$INSTALL_DIR/$BIN_NAME" "$BINARIES_URL"

# 检查下载是否成功
if [ $? -eq 0 ]; then
    # 授予执行权限
    chmod +x "$INSTALL_DIR/$BIN_NAME"
    echo "Binary downloaded and made executable."
else
    echo "Failed to download the binary."
    exit 1
fi

# 创建软链接指向下载目录
echo "Creating symbolic link in /usr/local/bin..."
sudo ln -sf "$INSTALL_DIR/$BIN_NAME" "$TARGET_DIR/$BIN_NAME"

echo "Setup completed!"
