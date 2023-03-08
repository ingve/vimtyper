# vimtyper

> Use a normal editor for development, but pretend to be a Vim user in screencasts.

Unix filter to generate Vim keyboard activity in [vhs](https://github.com/charmbracelet/vhs#vhs) [tape file format](https://github.com/charmbracelet/vhs/blob/main/README.md#vhs-command-reference).

## Build

```
$ git clone https://github.com/ingve/vimtyper
$ cd vimtyper
$ go build .
```

## Usage

Make sure you have [vhs](https://github.com/charmbracelet/vhs#vhs) installed.

```
$ vimtyper < vimtyper.go > demo.tape
$ vhs < demo.tape
```

<img alt="Demo output from vimtyper-generated vhs tape" src="media/out.gif" width="600" />

## License

MIT © 2023 [Ingve Vormestrand](https://github.com/ingve)
