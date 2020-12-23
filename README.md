# qc
![CI](https://github.com/simse/qc/workflows/CI/badge.svg)

A blazing fast, file format conversion CLI tool

**WARNING: qc is not yet released. There may be pre-releases available, but it may not work or break at any point.**


## Examples
**Convert all files to png**

`qc png`

which is an alias of `qc --in * --out png`


**Convert all jpg files to png**

`qc --in jpg --out png`


**Convert all ttf files to otf, including subdirectories**

`qc --in ttf --out otf --recursive`


## Install
### Windows

Use choco or something

### Linux/macOS

Use homebrew!


## Supported formats
You can also check with `qc formats`

### Images
- JPG
- PNG
- Webp
- TIFF
- BMP