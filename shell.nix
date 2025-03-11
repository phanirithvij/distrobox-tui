{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShellNoCC {
  packages = with pkgs; [
    go_1_23
    nixfmt-rfc-style
    (lilipod.overrideAttrs (p: {
      version = "dev";
      src = fetchFromGitHub {
        inherit (p.src) owner repo;
        rev = "refs/heads/main";
        hash = "sha256-pSImeXLYZ7jQJWagvkgKVGgjdhd84FiCCozv6m5Ijqs=";
      };
      vendorHash = null;
      preBuild = ''
        cp ${pkgsStatic.busybox}/bin/busybox .
      '';
      ldflags = [ ];
      patches = [
        (writeText "busybox_unvendor.patch" ''
          diff --git a/Makefile b/Makefile
          index 3e7468c..92e1d82 100644
          --- a/Makefile
          +++ b/Makefile
          @@ -1,6 +1,6 @@
           .PHONY: all lilipod pty coverage

          -all: busybox pty lilipod
          +all: pty lilipod

           clean:
           	@rm -f lilipod
          @@ -19,12 +19,8 @@ coverage:
           	@rm -f pty.tar.gz
           	CGO_ENABLED=0 go build -mod vendor -gcflags=all="-l -B -C" -ldflags="-s -w" -o pty ptyagent/main.go ptyagent/pty.go
           	tar czfv pty.tar.gz pty
          -	@wget -c "https://busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox"
           	CGO_ENABLED=0 go build -mod vendor -cover -o coverage/lilipod main.go

          -busybox:
          -	@wget -c "https://busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox"
          -
           pty:
           	@rm -f pty
           	@rm -f pty.tar.gz
        '')
      ];
    }))

    (callPackage ./default.nix { })
  ];
}
