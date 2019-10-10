![hwgo-01](https://user-images.githubusercontent.com/8619106/57969063-e6255000-798f-11e9-973e-9fdf7a0bd5fc.png)

# Go bindings and CLI tool for the Skycoin hardware wallet

[![Build Status](https://travis-ci.com/skycoin/hardware-wallet-go.svg?branch=master)](https://travis-ci.com/skycoin/hardware-wallet-go)

## Table of contents

<!-- MarkdownTOC levels="1,2,3,4,5" autolink="true" bracket="round" -->
- [Installation](#installation)
- [Usage](#usage)
  - [Download source code](#download-source-code)
  - [Dependancies management](#dependancies-management)
  - [Run](#run)
- [Development guidelines](#development-guidelines)
  - [Versioning policies](#versioning-policies)
  - [Running tests](#running-tests)
  - [Releases](#releases)
    - [Update the version](#update-the-version)
    - [Pre-release testing](#pre-release-testing)
    - [Creating release builds](#creating-release-builds)
- [Wiki](#wiki)
<!-- /MarkdownTOC -->

## Installation

### Install golang

https://github.com/golang/go/wiki/Ubuntu

## Usage

### Download source code

```bash
$ go get github.com/skycoin/hardware-wallet-go
```

### Dependancies management

This project uses dep [dependancy manager](https://github.com/golang/dep).

Don't modify anything under vendor/ directory without using [dep commands](https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md).

Download dependencies using command:

```bash
$ make dep
```

### Run

```bash
$ go run cmd/cli/cli.go
```

See also [CLI README](https://github.com/skycoin/hardware-wallet-go/blob/master/cmd/cli/README.md) for information about the Command Line Interface.

# Development guidelines

Code added in this repository should comply to development guidelines documented in [Skycoin wiki](https://github.com/skycoin/skycoin/wiki).

The project has two branches: `master` and `develop`.

- `develop` is the default branch and will always have the latest code.
- `master` will always be equal to the current stable release on the website, and should correspond with the latest release tag.

# Versioning policies

This go client library follows the [versioning rules for SkyWallet client libraries](https://github.com/skycoin/hardware-wallet/tree/develop/README.md#versioning-libraries) .

# Running tests

Library test suite can be run by running the following command

```
make test
```

If a physical SkyWallet device is connected then the test suite will exchange test messages with it. Likewise, the emulator will reply to test messages if executed in advance as follows.

```
git clone https://github.com/skycoin/hardware-wallet /some/path/to/hardware-wallet
make clean -C /some/path/to/hardware-wallet && make -C /some/path/to/hardware-wallet run-emulator
```

If neither the emulator nor a physical device are connected then tests will be skipped silently.

# Releases

# Update the version

0. If the `master` branch has commits that are not in `develop` (e.g. due to a hotfix applied to `master`), merge `master` into `develop` (and fix any build or test failures)
0. Switch to a new release branch named `release-X.Y.Z` for preparing the release.
0. Run `make build` to make sure that the code base is up to date
0. Update `CHANGELOG.md`: move the "unreleased" changes to the version and add the date.
0. Follow the steps in [pre-release testing](#pre-release-testing)
0. Make a PR merging the release branch into `master`
0. Review the PR and merge it
0. Tag the `master` branch with the version number. Version tags start with `v`, e.g. `v0.20.0`. Sign the tag. If you have your GPG key in github, creating a release on the Github website will automatically tag the release. It can be tagged from the command line with `git tag -as v0.20.0 $COMMIT_ID`, but Github will not recognize it as a "release".
0. Tag the changeset of the `protob` submodule checkout with the same version number as above.
0. Release builds are created and uploaded by travis. To do it manually, checkout the master branch and follow the [create release builds instructions](#creating-release-builds).
0. Checkout `develop` branch and bump `tiny-firmware/VERSION` and `tiny-firmware/bootloader/VERSION` to next [`dev` version number](https://www.python.org/dev/peps/pep-0440/#developmental-releases).


# Pre-release testing

Pre-release testing procedure requires [skycoin-cli](https://github.com/skycoin/skycoin/tree/develop/cmd/cli) of version strictly greater than `0.26.0`. Please [install it](https://github.com/skycoin/skycoin/blob/develop/cmd/cli/README.md#install) if not available in your system. Some operations in the process require [running a Skycoin node](https://github.com/skycoin/skycoin/tree/master/INTEGRATION.md#running-the-skycoin-node). Also clone [Skywallet firmware repository](https://github.com/skycoin/hardware-wallet/) in advance.

The instructions that follow are meant to be followed for Skywallet devices flashed without memory protection. If your device memory is protected then some values might be different e.g. `firmwareFeatures`.

During the process beware of the fact that running an Skycoin node in the background can block the Skywallet from running.

Some values need to be known during the process. They are represented by the following variables:

- `WALLET1`, `WALLET2`, ... names of wallets created by `skycoin_cli`
- `ADDRESS1`, `ADDRESS2`, ... Skycoin addresses
- `TXN1_RAW`, `TXN2_RAW`, ... transactions data encoded in hex dump format
- `TXN1_JSON`, `TXN2_JSON`, ... transactions data encoded in JSON format, if numeric index value matches the one of another variable with `RAW` prefix then both refer to the same transaction
- `TXN1_ID`, `TXN2_ID`, ... Hash ID of transactions after broadcasting to the P2P network
- `AMOUNT` represents an arbitrary number of coins
- `ID1`, `ID2`, `ID3`, ... unique ID values , usually strings identifying hardware or software artifacts
- `AUTHTOKEN` is a CSRF token for the Skycoin REST API

Perform these actions before releasing:

**Note** : In all cases `skycoin-cli` would be equivalent to `go run cmd/cli/cli.go` if current working directory set to `$GOPATH/src/github.com/skycoin/skycoin`.

##### Run project test suite

- Open a terminal window and run Skywallet emulator. Wait for emulator UI to display.
- From a separate trminal window run the test suite as follows
```sh
npm run test
```
- Close emulator and plug Skywallet device. Run the same command another time.

##### Run transaction tests

- Create new wallet e.g. with `skycoin-cli` (or reuse existing wallet for testing purposes)
```sh
skycoin-cli walletCreate -f $WALLET1.wlt -l $WALLET1
```
- From command output take note of the seed `SEED1` and address `ADDRESS1`
- List wallet addresses and confirm that `ADDRESS1` is the only value in the list.
```sh
skycoin-cli listAddresses $WALLET1.wlt
```
- Transfer funds to `ADDRESS1` in new wallet in two transactions
- Check balance
```sh
skycoin-cli addressBalance $ADDRESS1
```
- List address for `WALLET1` and check that `head_outputs` in response includes to outputs with `address` set to `ADDRESS1`
```
skycoin-cli walletOutputs $WALLET1.wlt
```
- [Get device features](cmd/cli/README.md#device-features) and check that:
  * `vendor` is set to `Skycoin Foundation`
  * `deviceId` is a string of 24 chars, which we'll refer to as `ID1`
  * write down the value of `bootloaderHash` i.e. `ID2`
  * `model` is set to `'1'`
  * `fwMajor` is set to expected firmware major version number
  * `fwMinor` is set to expected firmware minor version number
  * `fwPatch` is set to expected firmware patch version number
  * `firmwareFeatures` is set to `0`
- Ensure device is seedless by [wiping it](cmd/cli/README.md#wipe-device). Check that device ends up in home screen with `NEEDS SEED!` message at the top.
- [Recover seed](cmd/cli/README.md#recovery-device) `SEED1` in Skywallet device (`dry-run=false`).
- [Get device features](cmd/cli/README.md#device-features) and check that:
  * `vendor` is set to `Skycoin Foundation`
  * `deviceId` is set to `ID1`
  * `pinProtection` is seto to `false`
  * `passphraseProtection` is set to `false`
  * `label` is set to `ID1`
  * `initialized` is set to `true`
  * `bootloaderHash` is set to `ID2`
  * `passphraseCached` is set to `false`
  * `needsBackup` is set to `false`
  * `model` is set to `'1'`
  * `fwMajor` is set to expected firmware major version number
  * `fwMinor` is set to expected firmware minor version number
  * `fwPatch` is set to expected firmware patch version number
  * `firmwareFeatures` is set to `0`
- [Set device label](cmd/cli/README.md#apply-settings) to a new value , say `ID3`. Specify `usePassphrase=false`.
- [Get device features](cmd/cli/README.md#device-features) and check that:
  * `label` is set to `ID3`
  * all other values did not change with respect to previous step, especially `deviceId`
  * as a result device label is displayed on top of Skycoin logo in device home screen
- Ensure you know at least two addresses for test wallet, if not, [generate some](cmd/cli/README.md#ask-device-to-generate-addresses). Choose the second and third in order, hereinafter referred to as `ADDRESS2`, `ADDRESS3` respectively
- Check that address sequence generated by SkyWallet matches the values generated by `skycoin-cli`
```sh
skycoin-cli walletAddAddresses -f $WALLET1.wlt -n 5
```
- Check once again with desktop wallet
- Create new transaction from `ADDRESS1` to `ADDRESS2` in test wallet (say `TXN_RAW1`) for an spendable amount higher than individual output's coins
```sh
export TXN1_RAW="$(skycoin-cli createRawTransaction -a $ADDRESS1 -f $WALLET1.wlt $ADDRESS2 $AMOUNT)"
echo $TXN1_RAW
```
- Display transaction details and confirm that it contains at least two inputs
```sh
export TXN1_JSON=$(skycoin-cli decodeRawTransaction $TXN1_RAW)
echo $TXN1_JSON
```
- [Sign transaction](cmd/cli/README.md#transaction-sign) with Skywallet by putting together a message using values resulting from previous step as follows.
  * For each hash in transaction `inputs` array and in same order there should be
    - an instance of `--inputHash` command line argument set to the very same hash 
    - an instance of `--inputIndex` set to `0`.
  * For each source item in transaction `outputs` array and in the same order there should be command-line arguments set as follows:
    - `--outputAddress` : source item's `dst`
    - `--coin` : source item's `coins * 1000000`
    - `--hour` : source item's `hours`
    - `--addressIndex` : set to `0` if source item `address` equals `ADDRESS1` or to `1` otherwise
- Check that `signatures` array returned by hardware wallet includes entries for each and every transaction input
- [Check signatures](cmd/cli/README.md#ask-device-to-check-signature) were signed by corresponding addresses
- Create transaction `TXN2_JSON` by replacing `TXN1_JSON` signatures with the array returned by SkyWallet
- Use `TXN2_JSON` to obtain encoded transaction `TXN2_RAW`
```sh
export $TXN2_RAW=$( echo "$TXN2_JSON" | skycoin-cli encodeJsonTransaction --fix - | grep '"rawtx"' | cut -d '"' -f4)
echo $TXN2_RAW
```
- Broadcast transaction. Refer to its id as `TXN2_ID`
```sh
export TXN2_ID=$(skycoin-cli broadcastTransaction $TXN2_RAW)
```
- After a a reasonable time check that balance changed.
```sh
skycoin-cli walletBalance $WALLET1.wlt
```
- Create a second wallet i.e. `WALLET2`
```sh
skycoin-cli walletCreate -f $WALLET2.wlt -l $WALLET2
```
- From command output take note of the seed `SEED2` and address `ADDRESS4`
- List `WALLET2` addresses and confirm that `ADDRESS4` is the only value in the list.
```sh
skycoin-cli lisAddresses $WALLET2.wlt
```
- Transfer funds to `WALLET2` (if not already done) and check balance
```sh
skycoin-cli addressBalance $ADDRESS4
```
- Request CSRF token (i.e. `AUTHTOKEN`) using Skycoin REST API.
```sh
curl http://127.0.0.1:6420/api/v1/csrf
```
- Use Skycoin REST API to create one transaction grabbing all funds from `ADDRESS1` (i.e. first address in `WALLET1` previously recovered in Skywallet device), `ADDRESS2` (i.e. second address in `WALLET1` previously recovered in Skywallet device) and `ADDRESS4` (i.e. first address in `WALLET2`) so as to transfer to the third address of `WALLET1` (i.e. `ADDRESS3`). Change should be sent back to `ADDRESS1`. In server response transaction JSON object (a.k.a `TNX3_JSON`) would be the object at `data.transaction` JSON path. If Skycoin node was started with default parameters this can be achieved as follows:
```sh
curl -X POST http://127.0.0.1:6420/api/v2/transaction -H 'content-type: application/json' -H "X-CSRF-Token: $AUTHTOKEN" -d "{
    \"hours_selection\": {
        \"type\": \"auto\",
        \"mode\": \"share\",
        \"share_factor\": \"0.5\"
    },
    \"addresses\": [\"$ADDRESS1\", \"$ADDRESS2\", \"$ADDRESS4\"],
    \"change_address\": \"$ADDRESS1\",
    \"to\": [{
        \"address\": \"$ADDRESS3\",
        \"coins\": \"$AMOUNT\"
    }],
    \"ignore_unconfirmed\": false
}"
```
- [Sign transaction](cmd/cli/README.md#transaction-sign) represented by `TXN3_JSON` for inputs owned by Skywallet (i.e. `WALLET1`)
  * For each object in transaction `inputs` array there should be:
    - an instance of `--inputHash` command line argument set to the value bound to object's `uxid` key
    - an instance of `--inputIndex` set as follows
      * empty string if input `address` is `ADDRESS4`
      * `0` if input `address` is `ADDRESS1`
      * `1` if input `address` is `ADDRESS2`
  * For each source item in transaction `outputs` array there should be command-line arguments set as follows:
    - `--outputAddress` : source item's `address`
    - `--coin` : source item's `coins * 1000000`
    - `--hour` : source item's `hours`
    - `--addressIndex` set to 2 (since destination address is `ADDRESS3`)
- Check that signatures array includes one entry for every input except the one associated to `ADDRESS4`, which should be an empty string
- [Recover seed](cmd/cli/README.md#recovery-device) `SEED2` in Skywallet device (`dry-run=false`).
- [Sign transaction](cmd/cli/README.md#transaction-sign) represented by `TXN3_JSON` for inputs owned by Skywallet (i.e. `WALLET2`)
  * For each hash in transaction `inputs` array and in same order there should be:
    - an instance of `--inputHash` command line argument set to the value bound to object's `uxid` key
    - an instance of `--inputIndex` set to a value as follows
      * not set if input `address` is `ADDRESS1`
      * not set if input `address` is `ADDRESS2`
      * `0` if input `address` is `ADDRESS4`
  * For each source item in transaction `outputs` array and in the same order there should be command-line arguments set as follows:
    - `--outputAddress` : source item's `dst`
    - `--coin` : source item's `coins * 1000000`
    - `--hour` : source item's `hours`
    - `--addressIndex` empty (since destination address `ADDRESS3` not owned by this wallet)
- Create a new transaction JSON object (a.k.a `TXN4_JSON`) from `TXN3_JSON` and the previous signatures like this
  * `type` same as in `TXN3_JSON`
  * `inner_hash` should be an empty string
  * `sigs` returned by SkyWallet in same order as corresponding input
  * `inputs` is an array of strings. For each item in `TXN3_JSON` `inputs` include the value of its `uxid` field in `TXN4_JSON` `inputs` array. Respect original order.
  * `outputs` is an array of objects constructed out of `TXN3_JSON` `outputs` items, in te same order, as follows
    - `dst` : source item's `address`
    - `coins` : source item's `coins * 1000000`
    - `hours` : source item's `hours` as integer
- Use `TXN4_JSON` to obtain encoded transaction `TXN4_RAW`
```sh
export $TXN4_RAW=$( echo "$TXN4_JSON" | skycoin-cli encodeJsonTransaction --fix - | grep '"rawtx"' | cut -d '"' -f4)
echo $TXN4_RAW
```
- Broadcast transaction. Refer to its id as `TXN4_ID`
```sh
export TXN4_ID=$(skycoin-cli broadcastTransaction $TXN4_RAW)
```
- After a a reasonable time check that wallets balance changed.
```sh
skycoin-cli walletBalance $WALLET1.wlt
skycoin-cli walletBalance $WALLET2.wlt
```

## Wiki

More information in [the wiki](https://github.com/skycoin/hardware-wallet-go/wiki)
