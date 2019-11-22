# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- A `Bip44AddrIndex` type definition according to the bip 44 specification.
- A `Bip44AddrIndex` field to `SkycoinAddress`.
- A `Bip44AddrIndex` field to `SkycoinSignMessage`.
- A `Bip44AddrIndex` field to `SkycoinTransactionInput`.
- A `Bip44AddrIndex` field to `SkycoinTransactionOutput`.

### Fixed

### Changed

### Removed

### Fixed

### Security

## v1.0.0

### Added

- Add `RDP` level info in features msg.
- Add `fw_version_head` field to features msg.
- Add a `MessageType_GetMixedEntropy` rename `MessageType_GetEntropy` to `MessageType_GetRawEntropy`.
- Customize go import path using env var `GO_IMPORT`. Useful when project is a subfolder of a parent go project.
- Generate protobuffer classes for go lang with `gogofast` plugin for `protoc`
- Generate protobuffer classes for Javascript with `pbjs`
- Generate protobuffer classes for C and Python with `nanopb` and Python plugin for `protoc`
- Custom library generation paths by setting `OUT_PY`, `OUT_C`, `OUT_GO`, `OUT_JS`
- Add go code generated from proto files into the `go` folder.
- Make go folder a proper package that can be vendor'ed from another package
- Address `index` of `SkycoinTransactionInput` marked as optional.
