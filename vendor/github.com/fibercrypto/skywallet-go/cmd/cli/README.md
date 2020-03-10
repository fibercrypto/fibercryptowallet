# CLI Documentation

Skycoin Hardware wallet command line interface

<!-- MarkdownTOC autolink="true" bracket="round" levels="1,2,3" -->

- [CLI Documentation](#cli-documentation)
  - [Install](#install)
  - [Usage](#usage)
    - [Apply settings](#apply-settings)
      - [Examples](#examples-apply-settings)
        - [Text output](#text-output-apply settings)
    - [Update firmware](#update-firmware)
    - [Ask device to generate addresses](#ask-device-to-generate-addresses)
      - [Examples](#examples-ask-device-to-generate-addresses)
        - [Text output](#text-output-ask-device-to-generate-addresses)
    - [Configure device mnemonic](#configure-device-mnemonic)
      - [Examples](#examples-configure-device-mnemonic)
        - [Text output](#text-output-configure-device-mnemonic)
    - [Ask device to generate mnemonic](#generate-mnemonic)
      - [Examples](#examples-ask-device to generate mnemonic)
        - [Text output](#text-output-ask-device-to-generate-mnemonic)
    - [Configure device PIN code](#configure-device-pin-code)
      - [Examples](#examples-configure-device-pin-code)
        - [Text output](#text-output-configure-device-pin-code)
    - [Ask device to sign message](#ask-device-to-sign-message)
      - [Examples](#examples-ask-device-to-sign-message)
        - [Text output](#text-output-ask-device-to-sign-message)
    - [Ask device to check signature](#ask-device-to-check-signature)
      - [Examples](#examples-ask-device-to-check-signature)
        - [Text output](#text-output-ask-device-to-check-signature)
      - [Note](#note)
    - [Wipe device](#wipe-device)
      - [Examples](#examples-wipe-device)
        - [Text output](#text-output-wipe-device)
    - [Ask the device to perform the seed backup procedure](#backup-device)
      - [Examples](#examples-ask-the-device-to-perform-the-seed-backup-procedure)
        - [Text output](#text-output-ask-the-device-to-perform-the-seed-backup-procedure)
    - [Ask the device to perform the seed recovery procedure](#recovery-device)
      - [Examples](#examples-ask-the-device-to-perform-the-seed-recovery-procedure)
        - [Text output](#text-output-ask-the-device-to-perform-the-seed-recovery-procedure)
    - [Ask the device Features](#device-features)
    - [Ask the device to cancel the ongoing procedure](#device-cancel)
    - [Ask the device to sign a transaction using the provided information](#transaction-sign)
    - [Ask the device to get internal raw entropy](#get-raw-entropy)
       - [Examples](#examples-ask-the-device-to-get-internal-raw-entropy)
        - [Text output](#text-output-ask-the-device-to-get-internal-raw-entropy)
    - [Ask the device to get internal mixed entropy](#get-mixed-entropy)
      - [Examples](#examples-ask-the-device-to-get-internal-mixed-entropy)
        - [Text output](#text-output-ask-the-device-to-get-internal-mixed-entropy)

<!-- /MarkdownTOC -->


## Install

```bash
$ cd $GOPATH/src/github.com/fibercrypto/skywallet-go/
$ ./install.sh
```

## Usage

After the installation, you can run `skycoin-hw-cli` to see the usage:

```
$ skycoin-hw-cli

NAME:
   skycoin-hw-cli - the skycoin hardware wallet command line interface

USAGE:
   skycoin-hw-cli [global options] command [command options] [arguments...]

VERSION:
   1.7.0

COMMANDS:
     applySettings          Apply settings.
     setMnemonic            Configure the device with a mnemonic.
     features               Ask the device Features.
     generateMnemonic       Ask the device to generate a mnemonic and configure itself with it.
     addressGen             Generate skycoin addresses using the firmware
     firmwareUpdate         Update device's firmware.
     signMessage            Ask the device to sign a message using the secret key at given index.
     checkMessageSignature  Check a message signature matches the given address.
     setPinCode             Configure a PIN code on a device.
     removePinCode          Remove a PIN code on a device.
     wipe                   Ask the device to wipe clean all the configuration it contains.
     backup                 Ask the device to perform the seed backup procedure.
     recovery               Ask the device to perform the seed recovery procedure.
     cancel                 Ask the device to cancel the ongoing procedure.
     transactionSign        Ask the device to sign a transaction using the provided information.
     getRawEntropy          Get device raw internal entropy and write it down to a file
     getMixedEntropy        Get device internal mixed entropy and write it down to a file
     getUsbDetails          Ask host usb about details for the hardware wallet
     help, h                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

All commands accept `--deviceType` option. Supported values are `USB` and `EMULATOR`.

### Internal entropy

There are two kinds of internal entropy, [`getRawEntropy`](#get-raw-entropy) and `getMixedEntropy`(#get-mixed-entropy). The difference between this two are that raw entropy comes from a random buffer function that uses a peripheral device under the hood, in the other hand the mixed entropy comes from a salted entropy source as described in [this FAQ](https://github.com/skycoin/hardware-wallet/blob/develop/FAQ.md#random-source).


### Apply settings

Configure device with settings such as: using passphrase, configuring label, setting device language


```
OPTIONS:
        --usePassphrase bool              Use this option if you want to activate passphrase on device
        --language      string            Configure device language
        --label         string            Configure label to identify the device
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli applySettings --usePassphrase
```
<details>
 <summary>View Output</summary>

```
2018/12/21 16:14:25 language:"" label:"" use_passphrase:true 
PinMatrixRequest response: 5959
2018/12/21 16:14:34 Setting pin: 5959
2018/12/21 16:14:34 Calling DecodeSuccessOrFailMsg on message kind 26
2018/12/21 16:14:34 MessagePinMatrixAck Answer is: 26 / 
Success with code:  Settings applied
```
</details>

Configure device with settings such as: change language


```bash
$ skycoin-hw-cli applySettings --language en
```
<details>
 <summary>View Output</summary>

```
[2019-03-24T17:02:36-04:00] INFO [device-wallet]: language:"en" label:"" use_passphrase:false 
Success with code:  Settings applied
```
</details>

### Update firmware

To update firmware from a usb message, the device needs to be in "bootloader mode". To turn on on "bootloader mode" unplug your device, hold both buttons at the same time and plug it back on.

The use this command:


```bash
$ skycoin-hw-cli firmwareUpdate --file=[your firmware .bin file]
```

```
OPTIONS:
        --file string            Path to your firmware file
```

### Ask device to generate addresses

Generate skycoin addresses using the firmware

```bash
$ skycoin-hw-cli addressGen [number of addresses] [start index]
```

```
OPTIONS:
        --addressN value    Number of addresses to generate. Assume 1 if not set. (default: 1)
        --startIndex value  Index where deterministic key generation will start from. Assume 0 if not set. (default: 0)
        --confirmAddress    If requesting one address it will be sent only if user confirms operation by pressing device's button.
        --deviceType value  Device type to send instructions to, hardware wallet (USB) or emulator. [$DEVICE_TYPE]
        --walletType value  Wallet type. Types are "deterministic" or "bip44"
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli addressGen --addressN=2 --startIndex=0
```
<details>
 <summary>View Output</summary>

```
MessageSkycoinAddress 117! array size is 2
MessageSkycoinAddress 117! Answer is: 2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw
MessageSkycoinAddress 117! Answer is: zC8GAQGQBfwk7vtTxVoRG7iMperHNuyYPs
```
</details>

### Configure device mnemonic

Configure the device with a mnemonic.

```bash
$ skycoin-hw-cli setMnemonic [mnemonic]
```

```
OPTIONS:
        --mnemonic value            Mnemonic that will be stored in the device to generate addresses.
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli setMnemonic --mnemonic="cloud flower upset remain green metal below cup stem infant art thank"
```
<details>
 <summary>View Output</summary>

```
MessageButtonAck Answer is: 2 / 
Ecloud flower upset remain green metal below cup stem infant art thank
```
</details>

### Generate mnemonic

Ask the device to generate a mnemonic and configure itself with it.

```bash
$ skycoin-hw-cli generateMnemonic
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli generateMnemonic
```
<details>
 <summary>View Output</summary>

```
2018/11/06 14:41:50 MessageButtonAck Answer is: 2 /
 Mnemonic successfully configured
```
</details>

### Configure device PIN code

Configure the device with a pin code.

```bash
$ skycoin-hw-cli setPinCode
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli setPinCode
```
<details>
 <summary>View Output</summary>

```
MessageButtonAck Answer is: 18 /

PinMatrixRequest response: 5757
Setting pin: 5757

MessagePinMatrixAck Answer is: 18 /

PinMatrixRequest response: 4343
Setting pin: 4343

MessagePinMatrixAck Answer is: 18 /

PinMatrixRequest response: 6262
Setting pin: 6262

MessagePinMatrixAck Answer is: 2 /

PIN changed
```

</details>

### Ask device to sign message

Ask the device to sign a message using the secret key at given index.

```bash
$ skycoin-hw-cli signMessage [address index] [message to sign]
```

```
OPTIONS:
        --addressIndex value  Index of the address that will issue the signature. Assume 0 if not set. (default: 0)
        --message value       The message that the signature claims to be signing.
        --deviceType value    Device type to send instructions to, hardware wallet (USB) or emulator. [$DEVICE_TYPE]
        --walletType value    Wallet type. Types are "deterministic" or "bip44"
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli signMessage  --addressN=2 --message="Hello World!"
```
<details>
 <summary>View Output</summary>

```
Success 2! address that issued the signature is: DEK8o3Dnnp8UfTZrZCcCPCA6oRLqDeuKKy85YoTmCjfR2xDcZCz1j6tC4nmaAxHH15wgff88R2xPatT4MRvGHz9nf
```
</details>

### Ask device to check signature

Check a message signature matches the given address.

```bash
$ skycoin-hw-cli checkMessageSignature [address] [signed message] [signature]
```

```
OPTIONS:
        --address value            Address that issued the signature.
        --message value            The message that the signature claims to be signing.
        --signature value          Signature of the message.
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli checkMessageSignature  --address=2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw --message="Hello World!" --signature=GvKS4S3CA2YTpEPFA47yFdC5CP3y3qB18jwiX1URXqWQTvMjokd3A4upPz4wyeAyKJEtRdRDGUvUgoGASpsTTUeMn
```
<details>
 <summary>View Output</summary>

```
Success 2! address that issued the signature is: 
#2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw

```
</details>

#### Note

The `[option]` in subcommand must be set before the rest of the values, otherwise the `option` won't
be parsed. For example:

If we want to specify a `change address` in `send` command, we can use `-c` option, if you run
the command in the following way:

```bash
$ skycoin-cli send $RECIPIENT_ADDRESS $AMOUNT -c $CHANGE_ADDRESS
```

The change coins won't go to the address as you wish, it will go to the
default `change address`, which can be by `from` address or the wallet's
coinbase address.

The right script should look like this:

```bash
$ skycoin-hw-cli send -c $CHANGE_ADDRESS $RECIPIENT_ADDRESS $AMOUNT
```

### Wipe device

Ask the device to generate a mnemonic and configure itself with it.

```bash
$ skycoin-hw-cli wipe
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli wipe
```
<details>
 <summary>View Output</summary>

```
2018/11/06 16:00:28 Wipe device 26! Answer is: 0806
2018/11/06 16:00:31 MessageButtonAck Answer is: 2 /

Device wiped
```
</details>

### Backup device

Ask the device to perform the seed backup procedure.

```bash
$ skycoin-hw-cli backup
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli backup
```
<details>
 <summary>View Output</summary>

```
2018/11/15 17:13:40 Backup device 26! Answer is:
2018/11/15 17:14:58 Success 2! Answer is: Seed successfully backed up
```
</details>

### Recovery device

Ask the device to perform the seed recovery procedure.

```bash
$ skycoin-hw-cli recovery
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli recovery
```
<details>
 <summary>View Output</summary>

```
2018/12/07 17:50:26 Recovery device 46! Answer is: 
Word: market
Word: gaze
Word: crouch
Word: enforce
Word: green
Word: art
Word: stem
Word: infant
Word: host
Word: metal
Word: flower
Word: cup
Word: exit
Word: thank
Word: upset
Word: cloud
Word: below
Word: body
Word: remain
Word: vocal
Word: team
Word: discover
Word: core
Word: abuse
Failed with code:  The seed is valid but does not match the one in the device
```
</details>

### Device features

Ask the device Features.

```bash
$ skycoin-hw-cli features
```

<details>
 <summary>View Output</summary>

```
2018/12/07 17:54:20 Vendor: Skycoin Foundation
MajorVersion: 1
MinorVersion: 6
PatchVersion: 1
BootloaderMode: false
DeviceId: 453543343446324545394145393446463443463634434445
PinProtection: false
PassphraseProtection: false
Language: 
Label: 
Initialized: true
BootloaderHash: 765b3ec3a9c5b2f70326d0afce869cef5d1081124b91e1440b5b96a41436b723
PinCached: false
PassphraseCached: false
FirmwarePresent: false
NeedsBackup: true
Model: 1
FwMajor: 0
FwMinor: 0
FwPatch: 0
FwVendor: 
FwVendorKeys: 
UnfinishedBackup: false
FirmwareFeatures: 0
```
</details>

- `FirmwareFeatures` is interpreted as a bits slice as follows
  * bit `0` (i.e. mask `0x1`) is active if user confirmation required prior to returning internal entropy
  * bit `1` (i.e. mask `0x2`) set if support for sending internal entropy back to the peer is enabled in firmware.
  * bit `2` (i.e. mask `0x4`) set if device is the emulator.

### Device cancel

Ask the device to cancel the ongoing procedure.

```bash
$ skycoin-hw-cli cancel
```

<details>
 <summary>View Output</summary>

```
2018/12/10 15:06:42 Action cancelled by user
```
</details>

### Transaction sign

Ask the device to sign a message using the secret key at given index.

```
OPTIONS:
        --inputHash value      Hash of the Input of the transaction we expect the device to sign
        --inputIndex value     Index of the input in the wallet
        --outputAddress value  Addresses of the output for the transaction
        --coin value           Amount of coins
        --hour value           Number of hours
        --addressIndex value   If the address is a return address tell its index in the wallet
        --deviceType value     Device type to send instructions to, hardware wallet (USB) or emulator. [$DEVICE_TYPE]
        --walletType value     Wallet type. Types are "deterministic" or "bip44"
```

```bash
$ skycoin-hw-cli transactionSign --inputHash a885343cc57aedaab56ad88d860f2bd436289b0248d1adc55bcfa0d9b9b807c3 --inputIndex=0 --outputAddress=zC8GAQGQBfwk7vtTxVoRG7iMperHNuyYPs --coin=1000000 --hour=1
```

<details>
 <summary>View Output</summary>

```
[a885343cc57aedaab56ad88d860f2bd436289b0248d1adc55bcfa0d9b9b807c3] [0]
[zC8GAQGQBfwk7vtTxVoRG7iMperHNuyYPs] [1000000] [1] []
```
</details>

### Get raw entropy

Ask the device to get internally generated [raw entropy](#internal-entropy).

```
OPTIONS:
        --entropyBytes value  Total number of how many bytes of raw entropy to read. (default: 1048576)
        --outFile value       File path to write out the raw entropy buffers, a "-" set the file to stdout. (default: "-")
        --deviceType value    Device type to send instructions to, hardware wallet (USB) or emulator. [$DEVICE_TYPE]
```

#### Examples
##### Text output
```bash
$ skycoin-hw-cli getRawEntropy --outFile - --entropyBytes 33
```

<details>
 <summary>View Output</summary>

```
INFO [skycoin-hw-cli]: Getting raw entropy from device
[23 239 184 152 31 15 62 85 216 241 180 64 251 108 122 204 241 116 35 96 112 154 122 162 53 243 178 209 28 43 99 174 79]
```
</details>

A real example about how to use this feature can be checked at the [TRNG validation](https://github.com/skycoin/hardware-wallet/tree/8edc2a28027875f464b68348c44fb188efb4dfbb#validate-the-trng) (please get noticed that the firmware should be build with this feature enabled trough `ENABLE_GETENTROPY`). The tool is use specifically [from here](https://github.com/skycoin/hardware-wallet/blob/8edc2a28027875f464b68348c44fb188efb4dfbb/trng-test/Makefile#L7-L8).

### Get mixed entropy

Ask the device to get internally generated [mixed entropy](#internal-entropy).

```
OPTIONS:
        --entropyBytes value  Total number of how many bytes of mixed entropy to read. (default: 1048576)
        --outFile value       File path to write out the mixed entropy buffers, a "-" set the file to stdout. (default: "-")
        --deviceType value    Device type to send instructions to, hardware wallet (USB) or emulator. [$DEVICE_TYPE]
```

#### Examples
##### Text output

```bash
$ skycoin-hw-cli getMixedEntropy --outFile - --entropyBytes 33
```

<details>
 <summary>View Output</summary>

```
INFO [skycoin-hw-cli]: Getting mixed entropy from device
[78 81 147 249 42 193 246 64 113 249 212 43 216 233 38 198 9 31 24 178 19 109 62 110 195 44 176 147 158 146 80 215 191]
```
</details>

A real example about how to use this feature can be checked at the [TRNG validation](https://github.com/skycoin/hardware-wallet/tree/8edc2a28027875f464b68348c44fb188efb4dfbb#validate-the-trng) (please get noticed that the firmware should be build with this feature enabled trough `ENABLE_GETENTROPY`). The tool is use specifically [from here](https://github.com/skycoin/hardware-wallet/blob/8edc2a28027875f464b68348c44fb188efb4dfbb/trng-test/Makefile#L7-L8).