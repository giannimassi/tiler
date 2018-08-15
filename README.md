# Tiler
![Build status](https://circleci.com/gh/giannimassi/tiler/tree/master.svg?style=shield&circle-token=4a4cf2c5038410b43d184e58792b971f58fefae5)

## Intro
Command line tool for splitting large images into tiles.

## Latest releases:
- [Linux 64 bit](https://7-144838954-gh.circle-artifacts.com/0/build/linux/tiler)
- [Mac 64 bit](https://7-144838954-gh.circle-artifacts.com/0/build/mac/tiler)
- [Windows 64 bit](https://7-144838954-gh.circle-artifacts.com/0/build/windows/tiler.exe)

## Usage:
`Tiler [<flags>] <input> [<output>]`

```
Flags:
      --help          Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose       Verbose mode.
      --tile=256      Tile size (default 256).
  -f, --format="png"  Image output format.

Args:
  <input>     Input image file path (accepted jpeg, png).
  [<output>]  Output directory.
  ````
