{
  description = "gMountie";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    goflake.url = "github:sagikazarmark/go-flake";
    goflake.inputs.nixpkgs.follows = "nixpkgs";
    nixgl.url = "github:nix-community/nixGL";
    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
  };

  outputs = { self, nixpkgs, flake-utils, goflake, nixgl, gomod2nix, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ goflake.overlay nixgl.overlay];
        };
        buildDeps = with pkgs; [
          go-task
          git
          go
          gnumake
          gtk3
          gcc
          webkitgtk_4_0
          pkg-config
          nodejs
          upx
          nsis
          wails
        ];
        devDeps = with pkgs; buildDeps ++ [
          pkgs.nixgl.nixGLIntel
          protobuf
          protoc-gen-go
          protoc-gen-go-grpc
        ];
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = devDeps;
          shellHook = ''
            export LANG=en_US.UTF-8
            export LC_COLLATE=en_US.UTF-8
          '';
        };
      });
}
