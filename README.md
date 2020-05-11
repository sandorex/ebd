# ebd
[![GitHub license](https://img.shields.io/github/license/sandorex/ebd)](https://github.com/sandorex/ebd/blob/master/LICENSE)

[`ebd`](https://github.com/sandorex/ebd) (from python version's name [`extract-browser-data.py`](https://github.com/sandorex/extract-browser-data.py)) is a library meant for data manipulation of browser profiles written in Go

> **WARNING** The library is very early in development

The library provides easy reading and writing of browser data, currently supported features are

| Browser        | State Detection [(1)]() | Bookmarks | Cookies      | History | Extension List [(2)]() | Account Info [(3)]() |
| -------------- | ----------------------- | --------- | ------------ | ------- | ---------------------- | -------------------- |
| Chromium-based | Done                    | TODO      | TODO [(4)]() | TODO    | TODO                   | TODO                 |
| Firefox-based  | Done                    | TODO      | TODO         | TODO    | TODO                   | TODO                 |

1. There are 3 states
   - Closed - the browser is not using it anymore
   - Running - the browser is using it currently
   - Unknown - the browser has crashed or the profile is corrupted

2. Returns list of extensions installed, with their info like
   - ID
   - Version
   - Author
   - Description
   - Link to extension page where it can be installed

3. Account information like email of the account *(can be used to check if there is an account signed in)*

4. Chromium encrypts the cookies, so this may be a pain

**NOTE:** There is more accessible data but some of it is browser-specific

This library is base of [`cbsync`](https://github.com/sandorex/cbsync.git)

## ebdutil
`ebdutil` is a CLI application that can do all the things the library can but with easy to use command-line interface that outputs JSON or YAML

## Browser Support
The library will support **Firefox** and **Chromium** based browsers

# License
Licensed under [Apache License 2.0](LICENSE)

# Credits
Huge thanks to the following projects

- [pycookiecheat](https://github.com/n8henrie/pycookiecheat) for Chromium cookie decryption code
- [Chromagnon](https://github.com/JRBANCEL/Chromagnon/wiki/Reverse-Engineering-SSNS-Format) for SSNS format reverse engineering
