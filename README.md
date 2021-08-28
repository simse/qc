# qc
[![CI](https://github.com/simse/qc/actions/workflows/ci.yml/badge.svg)](https://github.com/simse/qc/actions/workflows/ci.yml)

qc, short for Quick Convert, is a tool for converting between file formats. It relies on existing libraries such as libvips to perform the conversions. The goal is to provide a single entrypoint, for performing any file conversion you can think of.

## Installation
_Windows is not currently supported_

**Using homebrew**
```bash
brew install simse/qc
```

**From source**
- Clone the repo
- Install `libvips` and `build-essentials`
- Run `go run build.go --enable-cgo`

## Usage
Go to any folder and type

```bash
qc jpg
```
or pick some other format and watch the magic happen.


## Contributing
Pull requests are welcome. For major changes, please create a discussion first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)

Do anything you want, **except redistributing closed-sourced versions**.