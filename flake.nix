{
  description = "A Nix-flake-based Go development environment";

  inputs.nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1"; # unstable Nixpkgs

  outputs = {self, ...} @ inputs: let
    goVersion = 25; # Change this to update the whole stack

    supportedSystems = [
      "x86_64-linux"
      "aarch64-linux"
      "aarch64-darwin"
    ];
    forEachSupportedSystem = f:
      inputs.nixpkgs.lib.genAttrs supportedSystems (
        system:
          f {
            inherit system;
            pkgs = import inputs.nixpkgs {
              inherit system;
              overlays = [inputs.self.overlays.default];
            };
          }
      );
  in {
    overlays.default = final: prev: {
      go = final."go_1_${toString goVersion}";
    };

    packages = forEachSupportedSystem (
      {
        pkgs,
        system,
      }: rec {
        default = pkgs.buildGoModule rec {
          pname = "bili-danmaku-tui";
          version = "v0.1.5";

          src = ./.;

          vendorHash = "sha256-Oj1QqfJHZsIriym+Xq6GYAe3RahM8LtPhhFOMYQNlNw=";

          ldflags = [
            "-X github.com/Youthdreamer/bili-danmaku-tui/cmd.Version=${version}"
            "-X github.com/Youthdreamer/bili-danmaku-tui/cmd.GitCommit=${
              if (self ? rev)
              then self.rev
              else "dirty"
            }"
          ];

          subPackags = ["."];

          meta = with pkgs.lib; {
            description = "Bilibili danmaku TUI client";
            license = licenses.mit;
            mainProgram = "bili-danmaku-tui";
          };
        };
      }
    );

    devShells = forEachSupportedSystem (
      {
        pkgs,
        system,
      }: {
        default = pkgs.mkShellNoCC {
          packages = with pkgs; [
            # go (version is specified by overlay)
            go
            # goimports, godoc, etc.
            gotools
            # https://github.com/golangci/golangci-lint
            golangci-lint
            gnumake
            self.formatter.${system}
          ];

          shellHook = ''
            echo "🐹 Go 1.${toString goVersion} Development Environment"
            echo "Run 'make build' or 'nix build' to compile."
          '';
        };
      }
    );

    formatter = forEachSupportedSystem ({pkgs, ...}: pkgs.nixfmt);
  };
}
