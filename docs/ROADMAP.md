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
  - all lts docker versions
    - podman (same)
    - lilipod (lts + unstable)
  - nixos vm tests + asciicapture
  - add test in nixpkgs too?
  - bubbletea test framework (it must exist)

### Features

- [x] Update dependencies and bring it up-to-date
- [x] Auto refresh after any operations
- [ ] lilipod support
  - [x] basic support
  - [ ] rm --volumes bug, filed a report 89luca89/lilipod#36
- [ ] Creating distroboxes
  - [ ] prepopulated with distrobox official list
  - [ ] custom
- [ ] default arguments to pass to distrobox-enter
  - eg. distrobox-enter --clean-path alpine -- bash --norc
- [ ] config file XDG_DIR
- [ ] macos support (no promises, I don't have a mac)
  - does distrobox work on macos at all? wait till some distrobox user asking
    for macos support in a github issue
  - [ ] conditional compilation for lilipod
    - lilipod doesn't build on macos
