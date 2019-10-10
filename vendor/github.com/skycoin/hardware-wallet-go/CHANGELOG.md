# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Fixed

### Changed

### Removed

### Security

## v1.0.0

### Added

- Add travis CD instructions for github releases.
- Add BitEncodedFlags to encode/decode flags from/to a string.
- Add remove PIN code.
- Pretty(`json`) print features output.
- Add a progress bar, currently used in get entropy messages only.
- Add a `getRawEntropy` and `getMixedEntropy` to get entropy from the physical device.
- Use `protobuf` file definitions from a [`git submodule`](http://github.com/skycoin/hardware-wallet-protob.git).
- Mnemonic and recovery functions support `--wordCount` argument for seed sizes of `24` and `12` (default) .
- Add `-deviceType` flag and `DEVICE_TYPE` env var to set devicetype, options are `USB` or `EMULATOR`.
- Add autocomplete for cli
- Add `Devicer` and `DeviceDriver` interface for the hw wallet api to make it more testeable.
- Add mocks for `Devicer` and `DeviceDriver` interface.
- Add skycoin `v0.25.1` dependency.
- Support in apply settings command for configuring device label.
- Sign Skycoin transactions using `transactionSign` command.
- Add `SimulateButtonPress` function to simulate emulator button press.
- Define the device language through `-language` flag of `applySettings` command. Only `english` supported.
- Only connect to `USB` devices announcing Skycoin vendor and product through HID.
- Add `Available` function to check if a skycoin wallet is connected to the system.
- Added cli integration tests.

### Fixed

- Change protobuf messages for check signature to be consistent with [harware-wallet](https://github.com/skycoin/hardware-wallet/blob/2648cf384b5455c994ba54acf6a31cd1272c6f66/tiny-firmware/protob/messages.options#L21).
- CLI returns error during firmaware update if device is not in bootloader mode.

### Changed

- Change project structure to follow standard project layout.
- Replace `hardware-wallet-protob` submodule with a dep dependency.
- Updated usblib to fix issue on windows.
- Rename `device-wallet` package to `skywallet`.

### Removed

- Remove `protobuf` files from the project.
- Installation instructions for `protobuf` related tools, use this from `hardware-wallet-protob` submodule.
