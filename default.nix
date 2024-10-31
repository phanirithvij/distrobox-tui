{
  lib,
  buildGoModule,
  fetchFromGitHub,
}:

buildGoModule rec {
  pname = "distrobox-tui";
  version = "dev";

  src = ./.;

  vendorHash = "sha256-zKJrhR/l2HPQXTVgrxYLGGz0pjA0e81jB6H4YrNfH94=";

  ldflags = [ "-s" ];

  meta = {
    description = "A TUI for DistroBox";
    changelog = "https://github.com/phanirithvij/distrobox-tui/releases/tag/v${version}";
    homepage = "https://github.com/phanirithvij/distrobox-tui";
    license = lib.licenses.gpl3Plus;
    maintainers = with lib.maintainers; [ phanirithvij ];
    mainProgram = "distrobox-tui";
  };
}
