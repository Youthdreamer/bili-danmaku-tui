{
  pkgs ? import <nixpkgs> { },
}:
pkgs.callPackage ./package.nix { useCnMirror = true; }
