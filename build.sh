#!/bin/bash
set -e

ROOT_DIR=$(pwd)
FRONTEND_DIR="$ROOT_DIR/frontend"
BACKEND_DIR="$ROOT_DIR/backend"

echo "===== 构建前端 ====="
cd "$FRONTEND_DIR"
npm install
npm run build

echo "===== 复制前端 dist 到 backend ====="
rm -rf "$BACKEND_DIR/frontend"
mkdir -p "$BACKEND_DIR/frontend"
cp -r "$FRONTEND_DIR/dist/"* "$BACKEND_DIR/frontend/"

echo "===== 构建后端 ====="
cd "$BACKEND_DIR"
go mod tidy
CGO_ENABLED=0 GOARCH=arm64 go build -o "$ROOT_DIR/server"

echo "===== 完成 ====="
echo "可执行文件位于： $ROOT_DIR/server"

