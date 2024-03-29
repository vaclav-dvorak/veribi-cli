<!-- markdownlint-disable MD012 -->
# Changelog

## [Unreleased]

## [v0.1.12] - 2024-03-29

### 💥 Changed

- change go version to 1.22

### 🐞 Bug Fixes

- Fix detection of TH/s

## [v0.1.11] - 2023-07-04

### 🎉 Features

- Update completion scripts
- Use go-pretty table instead of uitable

### 🐞 Bug Fixes

- Fix regexp so ths is parsed correctly on offers

## [v0.1.10] - 2023-03-12

### ♻ Code Refactoring

- Do not use CGO_ENABLED anymore

## [v0.1.9] - 2023-03-12

### 📦 Added

- Add completion scripts into release package

## [v0.1.8] - 2023-03-09

### 🎉 Features

- Update homebrew formula automatically

### 📦 Added

- Add linux/arm to release

## [v0.1.7] - 2023-03-09

### ♻ Code Refactoring

- remove support for 32bit

## [v0.1.6] - 2023-03-09

### 🎉 Features

- New platforms added in release (darwin/arm64)

## [v0.1.5] - 2023-03-09

### 🎉 Features

- Make logo hidden with silent flag
- New completion scripts for bash, fish, zsh and powershell

### 🐞 Bug Fixes

- Fix logo missalignment caused by newline

### ♻ Code Refactoring

- Migrate from make to go-task

## [v0.1.4] - 2023-03-03

### 💥 Changed

- Change go version to 1.19

## [v0.1.3] - 2023-03-03

### 📦 Added

- Add welcome logo

### ♻ Code Refactoring

- Update all dependencies

## [v0.1.2] - 2022-10-10

### 🐞 Bug Fixes

- Fix auction parsing regexp

### ♻ Code Refactoring

- Make use of new info available on acution pages

## [v0.1.1] - 2022-09-26

### 📦 Added

- Add ability to calculate TH price on auctions
- Add health collumn to offers table

### 💥 Changed

- Remove unused `update` command.

### 🐞 Bug Fixes

- Fix some typos and error handlings

### ♻ Code Refactoring

- Migrate from `tabby` to `uitable`

## [v0.1.0] - 2022-08-22

### 🎉 Features

- Make `login` command functional
- New `update` command
- New curl based install script

### 📦 Added

- Added tabby for nicer table output
- `version` command checks for latest version and have JSON output

## [v0.0.4] - 2022-08-10

### 📦 Added

- Add sha256 checksums to release assets

### 💥 Changed

- make changelog nicer

## [v0.0.3] - 2022-08-10

### 🎉 Features

- add windows binaries to release

### 📦 Added

- Add changelog

### 💥 Changed

- Remove 386 arch from release

### 🐞 Bug Fixes

- Fix the executable extension on windows binaries
- Fix release name generation in pipeline

## [v0.0.2] - 2022-08-10

### 🎉 Features

- try to use single artifact
- different apporcha to automatic release
- try gh autogenerate notes
- build binaries for MacOs in release
- Add automatic release pipeline
- add offers command
- add "root" and "init" command

### 🐞 Bug Fixes

- job dependency
- try to point release by ref_name
- build binaries in one pipeline
- add release event to pipeline
- trigger build on release publish
- hopeful fix for pipeline
- add missing chglog config
- pipeline validation error

## [v0.0.1] - 2022-08-08


[Unreleased]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.12...HEAD
[v0.1.12]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.11...v0.1.12
[v0.1.11]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.10...v0.1.11
[v0.1.10]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.9...v0.1.10
[v0.1.9]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.8...v0.1.9
[v0.1.8]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.7...v0.1.8
[v0.1.7]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.6...v0.1.7
[v0.1.6]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.5...v0.1.6
[v0.1.5]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.4...v0.1.5
[v0.1.4]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.3...v0.1.4
[v0.1.3]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.0.4...v0.1.0
[v0.0.4]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.0.3...v0.0.4
[v0.0.3]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/vaclav-dvorak/veribi-cli/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/vaclav-dvorak/veribi-cli/releases/tag/v0.0.1
