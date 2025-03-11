{
  lib,
  buildGoModule,
}:

buildGoModule rec {
  pname = "distrobox-tui";
  version = "dev";

  src = lib.cleanSource ./.;

  vendorHash = "sha256-y64KqlJsZ8aVK7oBcduEC8VvbutoRC15LMUeZdokPfY=";

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
