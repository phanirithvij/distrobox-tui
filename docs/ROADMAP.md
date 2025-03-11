## Roadmap

### Project

- [ ] Makefile for building
- [x] nixify build
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

- [x] Update dependencies and bring it up-to-date
- [x] Auto refresh after any operations
- [ ] lilipod support
  - [ ] rm --volumes bug, filed a report 89luca89/lilipod#36
- [ ] Creating distroboxes
  - [ ] prepopulated with distrobox official list
  - [ ] custom
- [ ] default arguments to pass to distrobox-enter
  - eg. distrobox-enter --clean-path alpine -- bash --norc
- [ ] config file XDG_DIR
