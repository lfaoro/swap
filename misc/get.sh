#!/usr/bin/env bash
# SPDX-License-Identifier: AGPL-3.0-only
# 
# Copyright (c) 2024 Leonardo Faoro. All rights reserved.
# Use of this source code is governed by the AGPL-3.0 license
# found in the LICENSE file.

set -euo pipefail

# Cleanup function
cleanup() {
    if [ -n "${TEMP_FILE:-}" ]; then
        rm -f "$TEMP_FILE"
    fi
}
trap cleanup EXIT

# Error handler
error() {
    echo "Error: $1" >&2
    exit 1
}

check_permissions() {
    local dir="$1"
    TEMP_FILE=$(mktemp -t swap_install_XXXXXX) || error "Failed to create temp file"
    if ! mv "$TEMP_FILE" "$dir/" 2>/dev/null; then
        echo "Warning: No write permission in $dir"
        INSTALL_DIR="/tmp"
    fi
    rm -f "$dir/$(basename "$TEMP_FILE")"
}

check_path() {
    local dir="$1"
    if [[ ":$PATH:" != *":$dir:"* ]]; then
        echo "Warning: $dir is not in your PATH"
        case "$SHELL" in
            *bash) echo "Run: echo 'export PATH=\$PATH:$dir' >> ~/.bashrc" ;;
            *zsh)  echo "Run: echo 'export PATH=\$PATH:$dir' >> ~/.zshrc" ;;
            *)     echo "Add $dir to your PATH" ;;
        esac
    fi
}

# Configuration
APP_NAME=swap
REPO="lfaoro/swap"
LATEST_RELEASE_URL="https://github.com/${REPO}/releases/latest"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download"

# Detect system information
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Normalize architecture
case "${ARCH}" in
    x86_64|amd64) ARCH="x86_64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) error "Unsupported architecture: ${ARCH}" ;;
esac

# Set binary name and install directory based on OS
case "${OS}" in
    linux)
        BINARY_NAME="${APP_NAME}_linux_${ARCH}"
        INSTALL_DIR="$HOME/.local/bin"
        ;;
    darwin)
        BINARY_NAME="${APP_NAME}_darwin_${ARCH}"
        INSTALL_DIR="$HOME/.local/bin"
        ;;
    msys*|mingw*)
        OS="windows"
        BINARY_NAME="${APP_NAME}_windows_${ARCH}.exe"
        INSTALL_DIR="$HOME/bin"
        ;;
    *) error "Unsupported operating system: ${OS}" ;;
esac

# Get latest version
VERSION=$(curl -sSL "https://api.github.com/repos/${REPO}/releases/latest" | grep -o '"tag_name": "v[^"]*"' | sed 's/"tag_name": "//;s/"//')

# Create installation directory
mkdir -p "${INSTALL_DIR}" || error "Failed to create installation directory"
check_permissions "$INSTALL_DIR"

# Download and install binary
DOWNLOAD_BINARY_URL="${DOWNLOAD_URL}/${VERSION}/${BINARY_NAME}"
echo "Downloading ${APP_NAME} ${VERSION} for ${OS}/${ARCH}..."

if [ "${OS}" = "windows" ]; then
    curl -fsSL "${DOWNLOAD_BINARY_URL}" -o "${INSTALL_DIR}/${APP_NAME}.exe" --progress-bar || error "Download failed"
    chmod +x "${INSTALL_DIR}/${APP_NAME}.exe" || error "Failed to set executable permissions"
    BINARY_PATH="${INSTALL_DIR}/${APP_NAME}.exe"
else
    curl -fsSL "${DOWNLOAD_BINARY_URL}" -o "${INSTALL_DIR}/${APP_NAME}" --progress-bar || error "Download failed"
    chmod +x "${INSTALL_DIR}/${APP_NAME}" || error "Failed to set executable permissions"
    BINARY_PATH="${INSTALL_DIR}/${APP_NAME}"
fi

echo "Successfully installed ${APP_NAME} to: ${BINARY_PATH}"
check_path "${INSTALL_DIR}"

# Verify installation
"${BINARY_PATH}" --version || error "Failed to run ${APP_NAME}"
