# Gate Template

An template for creating your Minecraft proxy with Go.

> [Bug Report](https://github.com/minekube/gate/issues) · [Docs](https://gate.minekube.com/developers/)

## About The Project

This template repository bootstraps your [Minekube Gate](https://github.com/minekube/gate) project, a customizable
Minecraft proxy written in Go.

## Getting Started

1. Fork this repository on GitHub.
2. Clone forked repository (`git clone <your-forked-repo-url>`)
3. Open project in your favorite Go IDE.
4. Run the proxy: `go run .`
5. Start customizing Gate to your needs!

## Usage

To create a new Gate plugin, follow these steps:

1. Create and write your plugin code in a new `plugins/xyz/xyz.go` file.
2. Add your exported plugin to the `proxy.Plugins` slice in `gate.go`.
3. Build and run Gate with: `go run .`

Use the `-d` flag to run Gate in debug mode if you encounter issues. (`go run . -d`)
