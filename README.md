<a href="https://t.me/fibercrypto"> <img src="resources/images/icons/github-related/telegram.svg" height=16 title="Follow us on Telegram"></a>
<a href="https://github.com/fibercrypto/fibercryptowallet"> <img src="resources/images/icons/appIcon/appIcon.png" height=64 align="right"></a>

# FiberCrypto Wallet

**Multi-coin cryptocurrency wallet**

[![License: GPL v3](resources/images/icons/github-related/license-gplv3.svg)](LICENSE.GPLv3 "GPL v3")
[![Built with Qt](resources/images/icons/github-related/built-with-qt.svg)](https://qt.io "The Qt Company")
[![Open-source](resources/images/icons/github-related/open-source.svg)](https://github.com "Open-source")
[![WIP](resources/images/icons/github-related/wip.svg)](https://github.com/fibercrypto/fibercryptowallet/issues "WIP: Check issues")
[![Contributions welcome](resources/images/icons/github-related/contributions-welcome.svg)](CONTRIBUTING.md "Contributions are welcome")

:star: Star us on GitHub — it really helps!


## Table of content

- [FiberCrypto Wallet](#fibercrypto-wallet)
  - [Table of content](#table-of-content)
  - [Features](#features)
  - [Releases](#releases)
    - [Release status by platform](#release-status-by-platform)
  - [Build](#build)
    - [Build System](#build-system)
    - [Requirements](#requirements)
  - [Statistics](#statistics)
    - [Build and release](#build-and-release)
    - [Activity](#activity)
    - [Downloads](#downloads)
    - [Analysis](#analysis)
    - [Issue tracking](#issue-tracking)
    - [Code coverage](#code-coverage)
    - [Size](#size)
  - [License](#license)


## Features

***FiberCrypto Wallet*** is a cryptocurrency software wallet aimed at:

* Provide easy-to-use interactions to users
* State
* Out-of-the-box support for every *SkyFiber* token in a single place
* Support other altcoins
* Facilitate exchange of crypto assets
* Buy and sell supported crypto assets using fiat (e.g. USD, GBP, EUR, ...)
* Integrations with trading tools
* Offer basic blockchain-specific tools


## Releases

[![Releases](resources/images/icons/github-related/releases.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/releases/ "Releases")

As of today, no releases has been published yet. However, a compilable tag was published. Releases are planned for the following platforms: *Windows*, *macOS*, *Linux*, *Android*, *iOS* and *Raspberry Pi*. In the long term, we also want to release for *watchOS*, *Asteroid* and more embedded devices.

### Release status by platform

**Desktop platforms:**

Release type | ![Win](resources/images/icons/github-related/windows.svg "Windows") | ![mac](resources/images/icons/github-related/macos.svg "macOS") | ![linux](resources/images/icons/github-related/linux.svg "Linux")
-------------------------------------------|--------------------|--------------------|--------------------
Source code *(compilable)*                 | Yes                | Yes                | Yes
Installer *(recommended)*                  | No                 | No                 | No
Portable<br> *(compressed folder)*         | No                 | No                 | No
Portable<br> *(static, single executable)* | No                 | No                 | No

**Mobile platforms:**

Release type | ![droid](resources/images/icons/github-related/android.svg "Android") | ![ios](resources/images/icons/github-related/ios.svg "iOS")
-------------------------------------------|--------------------|--------------------
Source code *(compilable)*                 | No                 | No
Default *(Store)*                          | No                 | No
Alternative                                | No                 | No

**Embedded platforms:**

Release type | ![raspi](resources/images/icons/github-related/raspberry-pi.svg "Raspberry Pi")
-------------------------------------------|--------------------
Source code *(compilable)*                 | No
Installer *(recommended)*                  | No
Portable<br> *(compressed folder)*         | No
Portable<br> *(static, single executable)* | No

## Build


### Build System

[![CMake](resources/images/icons/github-related/cmake.svg)](https://cmake.org/ "CMake")

The minimum required version is [CMake 3.16](https://cmake.org/files/v3.16/ "Download CMake 3.16"), but we always recommend the [latest version](https://cmake.org/download/ "Download CMake") available.


### Requirements

<table>
     <tr><td style="width:90px"><a href="https://qt.io/"><img src="resources/images/icons/qt_logo_green_rgb.svg" title="The Qt Company"></a></td><td>This project uses the <a href="https://www.qt.io/" title="The Qt Company">Qt framework</a> for the frontend. The minimum required version is <a href="https://download.qt.io/archive/qt/6.0/6.0.0/" title="Download Qt 6.0.0">Qt 6.0.0</a>, but we always recommend using the <a href="https://download.qt.io/archive/qt/" title="Download latest version">latest version</a> available.</td></tr>
     <tr><td style="width:90px"><a href="https://golang.org/"><img src="resources/images/icons/github-related/go.svg" title="The Go Programming Language"></a></td><td>This project uses <a href="https://golang.org/" title="The Go Programming Language">Go</a> for the backend. The minimum required version is <a href="https://blog.golang.org/go1.13/" title="Go 1.13 is released - The Go Blog">Go 1.13</a>, but we always recommend using the <a href="https://golang.org/dl/" title="Download latest version">latest version</a> available.</td></tr>
     <tr><td style="width:90px"><a href="https://imagemagick.org/"><img src="resources/images/icons/github-related/image-magick.svg" title="Image Magick"></a></td><td>Windows requires the command line tool <tt>convert</tt> (or <tt>magick convert</tt>, depending of the version used), that comes with the open-source <a href="https://imagemagick.org" title="Image Magick">ImageMagick</a> project in order to build the icon. A default icon is always provided, so this is not neccesary unless it's going to be replaced.</td></tr>
</table>


## Statistics

<!-- TODO: Add localization status -->
<!-- TODO: Add social networks status -->
<!-- TODO: Add funding status -->

### Build and release
[![Build Status](https://img.shields.io/travis/fibercrypto/fibercryptowallet/develop)](https://travis-ci.org/fibercrypto/fibercryptowallet "Build status")
[![GitHub release](https://img.shields.io/github/release/fibercrypto/fibercryptowallet.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/releases/ "Releases")
[![GitHub tag](https://img.shields.io/github/tag/fibercrypto/fibercryptowallet.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/tags/ "Tags")

### Activity
[![GitHub contributors](https://img.shields.io/github/contributors/fibercrypto/fibercryptowallet.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/commit/ "Contributors")
[![GitHub release date](https://img.shields.io/github/release-date/fibercrypto/fibercryptowallet.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/releases/ "Release date")
[![Milestone](https://img.shields.io/github/milestones/progress/fibercrypto/fibercryptowallet/2.svg)](https://github.com/fibercrypto/fibercryptowallet/milestones/2 "Progress of next release")
[![GitHub commits since last release](https://img.shields.io/github/commits-since/fibercrypto/fibercryptowallet/latest/develop.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/commit/ "Commits since last release")
[![GitHub last commit](https://img.shields.io/github/last-commit/fibercrypto/fibercryptowallet.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/commit/ "Last commit")

### Downloads
[![Github all downloads](https://img.shields.io/github/downloads/fibercrypto/fibercryptowallet/total.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/releases/ "All downloads")
[![Github latest version downloads](https://img.shields.io/github/downloads/fibercrypto/fibercryptowallet/latest/total.svg)](https://GitHub.com/fibercrypto/fibercryptowallet/releases/ "Latest version downloads")

### Analysis
![Github languages](https://img.shields.io/github/languages/count/fibercrypto/fibercryptowallet.svg "Languages count")
![Github top language](https://img.shields.io/github/languages/top/fibercrypto/fibercryptowallet.svg "Top language")
![Github languages](https://img.shields.io/scrutinizer/quality/g/fibercrypto/fibercryptowallet/develop.svg "Top language")

### Issue tracking
[![Github issues](https://img.shields.io/github/issues-raw/fibercrypto/fibercryptowallet.svg)](https://githib.com/fibercrypto/fibercryptowallet/issues "Open issues")
[![Pull requests](https://img.shields.io/github/issues-pr-raw/fibercrypto/fibercryptowallet.svg)](https://githib.com/fibercrypto/fibercryptowallet/pr "Open pull requests")

### Code coverage
[![Coverage Status](https://img.shields.io/coveralls/github/fibercrypto/FiberCryptoWallet/develop)](https://coveralls.io/github/fibercrypto/FiberCryptoWallet?branch=develop "Coverage status")

### Size
![Github code size](https://img.shields.io/github/languages/code-size/fibercrypto/fibercryptowallet.svg "Code size")
![Github repository size](https://img.shields.io/github/repo-size/fibercrypto/fibercryptowallet.svg "Repository size")
![Github file count](https://img.shields.io/github/directory-file-count/fibercrypto/fibercryptowallet.svg "File count")
![Github .go file count](https://img.shields.io/github/directory-file-count/fibercrypto/fibercryptowallet/*.svg?color=blue&extension=go&label=.go%20files.svg ".go file count")
![Github .qml file count](https://img.shields.io/github/directory-file-count/fibercrypto/fibercryptowallet/*.svg?color=blue&extension=qml&label=.qml%20files ".qml file count")


## License
[![License: GPL v3](resources/images/icons/github-related/license-gplv3.svg)](LICENSE.GPLv3 "GPL v3")

<h2>
WIP
</h2>

[![WIP](resources/images/icons/github-related/wip.svg)](https://github.com/fibercrypto/fibercryptowallet/issues "WIP: Check issues")
