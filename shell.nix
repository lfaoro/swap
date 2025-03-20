{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.musl.dev
    pkgs.musl
    pkgs.binutils
    pkgs.gosec
    pkgs.upx
  ];

  shellHook = ''
    export CC=musl-gcc
    export CXX=musl-g++
    export CGO_ENABLED=0
    export GOFLAGS="-buildmode=pie -trimpath"
    export GOLDFLAGS="-s -w -linkmode external -extldflags '-static'"
  '';
}
