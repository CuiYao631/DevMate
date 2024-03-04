<div align="center">
<a href="https://github.com/tiny-craft/tiny-rdm/"><img src="build/appicon.png" width="120"/></a>
</div>
<h1 align="center">DevMate</h1>
<!-- <h4 align="center"><strong>English</strong> | <a href="https://github.com/tiny-craft/tiny-rdm/blob/main/README_zh.md">
简体中文</a></h4> -->
<div align="center">

[![License](https://img.shields.io/github/v/release/CuiYao631/DevMate)](https://github.com/CuiYao631/DevMate/releases)
![GitHub All Releases](https://img.shields.io/github/downloads/CuiYao631/DevMate/total)
[![GitHub stars](https://img.shields.io/github/stars/CuiYao631/DevMate)](https://github.com/CuiYao631/DevMate/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/CuiYao631/DevMate)](https://github.com/CuiYao631/DevMate/fork)


</div>

<picture>

</picture>

## Feature
-  LAN Discovery Function
-  Chat function


## Roadmap
-  chatGPT
-  group chat
-  send file
-

## Installation

Available to download for free from [here](https://github.com/CuiYao631/DevMate/releases).

> If you can't open it after installation on macOS, exec the following command then reopen:
> ``` shell
>  sudo xattr -d com.apple.quarantine /Applications/DevMate.app
> ```

## Build Guidelines

### Prerequisites

* Go (latest version)
* Node.js >= 14
* NPM >= 9

### Install wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Clone the code

```bash
git clone https://github.com/CuiYao631/DevMate/edit
```

### Build frontend

```bash
npm install --prefix ./frontend
```

### Compile and run

```bash
wails dev
```

