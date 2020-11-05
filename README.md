<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich", "cjdenio", "safinsingh", "imgbot[bot]"]:end -->

<div align="center">
  <img alt="logo" src="./images/Entire%20Logo.png" height="250px">

  <h1>fgh</h1>

  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/Matt-Gleich/fgh">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/Matt-Gleich/fgh">
  <img alt="Golang report card" src ="https://goreportcard.com/badge/github.com/Matt-Gleich/fgh">
  <br>
  <img alt="build" src="https://github.com/Matt-Gleich/fgh/workflows/build/badge.svg" />
  <img alt="test" src="https://github.com/Matt-Gleich/fgh/workflows/test/badge.svg" />
  <img alt="lint" src="https://github.com/Matt-Gleich/fgh/workflows/lint/badge.svg" />
  <img alt="release" src="https://github.com/Matt-Gleich/fgh/workflows/release/badge.svg" />
  <br />
  <br />
  <i>📁 Automate the organization of your cloned GitHub repositories</i>
</div>

<hr />

## 📜 Table of Contents

- [📜 Table of Contents](#-table-of-contents)
- [👋 Getting started](#-getting-started)
- [📟 Commands](#-commands)
  - [🔒 `fgh login`](#-fgh-login)
  - [⚙️ `fgh configure`](#️-fgh-configure)
  - [☁️ `fgh clone`](#️-fgh-clone)
    - [🔠 Keywords](#-keywords)
  - [🚚 `fgh migrate`](#-fgh-migrate)
  - [⬆️ `fgh update`](#️-fgh-update)
  - [🧼 `fgh clean`](#-fgh-clean)
  - [🗑 `fgh remove`](#-fgh-remove)
  - [🧭 `fgh ls`](#-fgh-ls)
- [💡 Tips](#-tips)
  - [<owner/name> shorthand](#ownername-shorthand)
  - [`fgh ls` for `cd`](#fgh-ls-for-cd)
- [🗂 Custom Structures](#-custom-structures)
  - [📁 `structure_root`](#-structure_root)
  - [🗂 `structure`](#-structure)
  - [💡 Example Config](#-example-config)
  - [🚚 Moving Repos to New Structure](#-moving-repos-to-new-structure)
- [🚀 Install](#-install)
  - [🍎 macOS](#-macos)
  - [🐧 Linux and 🖥 Windows](#-linux-and--windows)
- [🙌 Contributing](#-contributing)
- [👥 Contributors](#-contributors)

## 👋 Getting started

As you begin contributing to an increasing amount of GitHub repositories, you'll soon realize the effort it takes to organize and maintain them on your machine. `fgh` aims to solve this issue through the use of a CLI (command line application) to automate the entire lifecycle of your cloned repos, saving you time _and_ helping you scale! Below is a list of the most useful automation commands:

- [`fgh clone`](#️-fgh-clone)
- [`fgh clean`](#-fgh-clean)
- [`fgh update`](#️-fgh-update)
- [`fgh ls`](#-fgh-ls)

See [the install section](#-install) to install `fgh`

## 📟 Commands

### 🔒 `fgh login`

Before using `fgh`, you'll need to give it access to your GitHub repos. Simply run `fgh login` to quickly get set up! fgh only uses this access to get metadata about the repo (e.g. main language, if private) and to clone the repo. fgh needs the full `repo` scope to access private repos.

If you need to use a GitHub custom access token, like a PAT, edit the secret configuration file. In Windows it is located in `~\.fgh\secrets.yml` and `~/.config/fgh/secrets.yml` in Linux and  macOS. You should change/add the `pat` as seen below:

```yaml
pat: <your token here>
```

### ⚙️ `fgh configure`

To configure other settings, run `fgh configure` for an interactive configuration experience.

### ☁️ `fgh clone`

To begin using `fgh`, you'll need to clone a repository, which you can do by running the following in a terminal window:

```bash
fgh clone <owner/name>
```

**OR**

```bash
fgh clone <name> # if the repo is under your account
```

All repositories are cloned into the following structure by default:

```
~
└─ github
   └─ OWNER
      └─ TYPE
         └─ MAIN_LANGUAGE
            └─ NAME
```

#### 🔠 Keywords

These names correspond to the following **keywords**:

- `OWNER` is the owner of the repository
- `TYPE` is the type of the repository; one of the following:
  - `public`
  - `private`
  - `template`
  - `archived`
  - `disabled`
  - `mirror`
  - `fork`
- `MAIN_LANGUAGE` is The main language that the repository contains. If no language is detected, `fgh` will just set it to `Other`
- `NAME` is the name of the repository

If you would like to use a custom structure see the [custom structures documentation](#-custom-structures). Usage of this command is as follows:

```bash
fgh clone <owner/name>
```

If we were to run `fgh clone Matt-Gleich/fgh` it would be cloned to `~/github/Matt-Gleich/public/Go/fgh/` by default (`~` being your home directory). Once cloned, this path will be copied to your clipboard automatically (this can be turned off with [`fgh configure`](#️-fgh-configure) or just by editing the config file directly).

> NOTE: On Linux machines running the X Window System, this program requires the `xclip` or `xsel` packages.

This structure can be somewhat difficult to navigate in the terminal using conventional methods such as the use of the `cd` command. Because of this, we created the [`fgh ls` command](#-fgh-ls) and a way to [use it with cd](#fgh-ls-for-cd).

### 🚚 `fgh migrate`

Would you like to add your existing repositories to the fgh structure? All you have to do run the following command and it will move every single git repo in that directory and all subdirectories into the structure:

```bash
fgh migrate <folder>
```

An example would be:

```
fgh migrate code
```

This would migrate all git repos in the `./code` folder into fgh's structure.

### ⬆️ `fgh update`

If any of a repository's fields are changed, such as its type, main language, owner, or name, the path to your local repository won't match.

Running `fgh update` will iterate over your local repositories and checks if any of them need updates. If they do, `fgh` will ask you if you want to move the entire repository to that new path.

For example, If I had this repository cloned and later decided to archive it, its path would change from `~/github/Matt-Gleich/public/Go/fgh/` to `~/github/Matt-Gleich/archived/Go/fgh/`.

### 🧼 `fgh clean`

When you run this subcommand, `fgh` will check for the following on each repository:

1. Has it modified locally in a certain amount of time?
   > By default, this "amount of time" is 2 months. However, it can be changed with a flag! See `fgh clean --help` for more info.
2. Has the repository been deleted on GitHub?

If either of those conditions are met, `fgh` will ask you if you would like to remove the aforementioned repository. It'll additionally show you some information about the repository itself.

> NOTE: This only removes the repo locally!

### 🗑 `fgh remove`

Remove a selected cloned repository. Usage is as follows:

```bash
fgh remove <owner/name>
```

### 🧭 `fgh ls`

Get the path of a cloned repository. Usage is as follows:

```bash
fgh ls <owner/name>
```

## 💡 Tips

### <owner/name> shorthand

Any command that takes `<owner/name>` as an argument allows you to leave off the `owner` if the repo is under your account. For example, I own this repo so I can just do

```bash
fgh clone fgh
```

instead of

```bash
fgh clone Matt-Gleich/fgh
```

### `fgh ls` for `cd`

> NOTE: This only works in macOS and Linux

If you would like to easily use the output of `fgh ls` for `cd` just add the following snippet to your `~/.zshrc` or `~/.bashrc`:

```bash
# cd with fgh (https://github.com/Matt-Gleich/fgh)
fcd() { cd "$(fgh ls "$@" 2>/dev/null)" || ( echo "Failed to find repository" && return 1; ) }
```

Once you add that and reload your terminal you can simply run `fcd <owner/name>` instead of `fgh ls <owner/name>`, copying the output to your clipboard, typing `cd`, and pasting the output. Much easier!

## 🗂 Custom Structures

Not a fan of the default structure used by fgh? Don't worry, you can change it without losing any of fgh's automation! Configuring custom structures takes place in the general configuration file. This file is located in `~/.config/fgh/config.yaml` on Linux or macOS and `~\.fgh\config.yaml` on Windows (`~` is your home directory). There are two parts to creating custom structures:

### 📁 `structure_root`

This is where the structure starts relative to your home folder. Make sure you use `\` if you are on Windows. By default, the `structure_root` is `github`. Below is an example of what you would put in the general config file:

```yaml
structure_root: 'Documents/code/'
```

If we were to run `fgh clone Matt-Gleich/fgh` with just the config shown above it would be cloned to `~/Documents/code/Matt-Gleich/public/Go/fgh`

### 🗂 `structure`

This is the structure used inside of the [`structure_root`](#-structure_root) If you use the [keywords shown in the clone structure](#-keywords) it will automatically be replaced by the value for the repo and add the name of the repo to the end. Below is an example of what you would put in the general config file:

```yaml
structure:
  - OWNER
  - repos
  - MAIN_LANGUAGE
```

If we were to run `fgh clone Matt-Gleich/fgh` with just the config shown above it would be cloned to `~/github/Matt-Gleich/repos/Go/fgh`.

### 💡 Example Config

Say we have the following config:

```yaml
structure_root: 'code'
structure:
  - OWNER
```

If we were to run `fgh clone Matt-Gleich/fgh` it would clone the repo to `~/code/Matt-Gleich/fgh`.

### 🚚 Moving Repos to New Structure

Just run:

```bash
fgh migrate <old project root>
```

## 🚀 Install

### 🍎 macOS

```bash
brew tap Matt-Gleich/homebrew-taps
brew install fgh
```

### 🐧 Linux and 🖥 Windows

You can grab the binary from the [latest release](https://github.com/Matt-Gleich/fgh/releases/latest).

## 🙌 Contributing

Thank you for considering contributing to `fgh`! Before contributing, make sure to read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/fgh/blob/master/CONTRIBUTING.md).

<!-- DO NOT REMOVE - contributor_list:start -->
## 👥 Contributors


- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

- **[@cjdenio](https://github.com/cjdenio)**

- **[@safinsingh](https://github.com/safinsingh)**

- **[@imgbot[bot]](https://github.com/apps/imgbot)**

<!-- DO NOT REMOVE - contributor_list:end -->
