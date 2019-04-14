# Introduction

GVM is a version manager for Golang, it is easy to use and compatible with all platforms and architecture
(provided that there is a version of Golang compatible with your system and architecture).

# Install

## Install the pre-compiled binary

**Homebrew** (for macOS):

```bash
brew install tfournier/tap/gvm
```

**Scoop** (for Windows):

```text
scoop install gvm
```

**Linux** manually:

```bash
wget -qO- https://github.com/tfournier/gvm/releases/latest/download/gvm_Linux_x86_64.tar.gz | tar xvz - -C /usr/local/bin gvm
```

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
