# Weather CLI

A simple command-line application to fetch a 3-day weather forecast for a given city.

## Prerequisites

- **Go**: Make sure you have Go installed. You can download it from the official [Go website](https://golang.org/dl/).
- **Make**: This project uses a Makefile to build and manage the installation. Make sure you have `make` installed (typically available by default on Linux/macOS).

## Features

- Fetches weather information from an external API.
- Provides a simple CLI to get weather forecasts.
- Supports easy installation and uninstallation.

## Installation

### Install the binary

To install the binary to `~/.local/bin`, run:

```bash
make install
```

This will copy the `weather` binary to your `~/.local/bin` directory (which is typically already in your PATH).

### Uninstall the binary

To uninstall the binary from `~/.local/bin`, run:

```bash
make uninstall
```

This will remove the `weather` binary from the installation directory.

## Usage

Once installed, you can run the weather CLI by using:

```bash
weather -c "City Name"
```

For example, to get the weather forecast for Kalmar:

```bash
weather -c "Kalmar"
```

You can also use the `-h` flag to display the help message:

```bash
weather -h
```
