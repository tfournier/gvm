# Introduction

GVM is a version manager for Golang, it is easy to use and compatible with all platforms and architecture
(provided that there is a version of Golang compatible with your system and architecture).

# Install

## For macOS ([Homebrew](https://brew.sh))

```bash
brew install tfournier/tap/gvm
```

## For linux ([Snapcraft](https://snapcraft.io/))
[![Installer à partir du Snap Store](https://snapcraft.io/static/images/badges/fr/snap-store-black.svg)](https://snapcraft.io/gvm)

```bash
snap install gvm
```

## For windows ([Scoop](https://scoop.sh/))

```bash
scoop install gvm
```

## Manually

Go to [releases](https://github.com/tfournier/gvm/releases) page

## Check installed

```bash
gvm
```

## Configure

In your shell configuration (example for ZSH: `~/.zshrc`)

```bash
echo 'eval $(gvm config show)' >> ~/.zshrc
```

# Quick Start

## Install Go version

You can check the versions available on [Golang's official website](https://golang.org/dl/)

```bash
gvm install 1.12.3
```

## Check installed versions

```bash
gvm list
```

## And use the previously installed version

````bash
gvm use 1.12.3
````

## Remove specific version

```bash
gvm remove 1.12.3
```

# Usage
[Full detailed commands](cmd/gvm.md)
