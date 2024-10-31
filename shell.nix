{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShellNoCC {
  packages = with pkgs; [
    go_1_23
    nixfmt-rfc-style

    (callPackage ./default.nix { })
  ];
}
