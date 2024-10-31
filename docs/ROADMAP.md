## Roadmap

### Project

- [ ] Makefile for building
- [ ] nixify build
  - [x] initial nix build in nixpkgs
  - [x] nix for dev (shell.nix + default.nix + direnv)
- [ ] gha for auto releases
  - pkgStatic/musl like miq (viperML/miq)
  - or use nix-portable as a bundler?
- [ ] Tests
  - all docker verisons
  - nixos vm tests
  - add test in nixpkgs too?

### Features

- [ ] Update dependencies and bring it up-to-date
- [ ] Auto refresh after any operations
- [ ] Creating distroboxes
  - [ ] prepopulated with distrobox official list
  - [ ] custom
- [ ] default arguments to pass to distrobox-enter
  - eg. distrobox-enter --clean-path alpine -- bash --norc
- [ ] config file XDG_DIR

