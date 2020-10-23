<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

<h1 align="center">fgh</h1>

<p align="center">

<img alt="build" src="https://github.com/Matt-Gleich/fgh/workflows/build/badge.svg" />
<img alt="test" src="https://github.com/Matt-Gleich/fgh/workflows/test/badge.svg" />
<img alt="lint" src="https://github.com/Matt-Gleich/fgh/workflows/lint/badge.svg" />
<img alt="release" src="https://github.com/Matt-Gleich/fgh/workflows/release/badge.svg" />
<br />
<i>📁 Automate the organization of your cloned GitHub repositories</i>

</p>

<hr />

## 📜 Table of Contents

- [🚀 Install](#---install)

  - [🍎 macOS](#---macos)
  - [🐧 Linux and 🖥 Windows](#---linux-and----windows)

- [📖 Documentation](#---documentation)

- [📚 Usage](#---usage)

  - [⚙️ `fgh configure`](#----fgh-configure-)
  - [☁️ `fgh clone`](#----fgh-clone-)

- [💡 Example](#---example)

  - [☝️ `fgh update`](#----fgh-update-)
  - [🧼 `fgh clean`](#----fgh-clean-)
  - [🗑 `fgh remove`](#----fgh-remove-)

  * [🛣 Roadmap](#---roadmap)
  * [🙌 Contributing](#---contributing)
  * [👥 Contributors](#---contributors)

## 🚀 Install

### 🍎 macOS

```bash
brew tap Matt-Gleich/homebrew-taps
brew install fgh
```

### 🐧 Linux and 🖥 Windows

You can grab the binary from the [latest release](https://github.com/Matt-Gleich/fgh/releases/latest).

## 📖 Documentation

As you begin contributing to an increasing amount of GitHub repositories, you'll soon realize the effort it takes to clone and organize them on your machine. `fgh` aims to solve this issue through the use of a CLI (command line application) to manage this entire process, saving you time _and_ helping you scale!

## 📚 Usage

### ⚙️ `fgh configure`

Before using `fgh`, you'll need to set a GitHub PAT (personal access token) in order to make requests to GitHub's API. In order to do this:

- Head over to [https://github.com/settings/tokens/new](https://github.com/settings/tokens/new)
- Ccreate a new token with the `repo` scope.

### ☁️ `fgh clone`

To actually begin using `fgh`, you'll need to clone a repository, which you can do by running the following in a terminal window:

```bash
fgh clone <owner/name>
```

**OR**

```bash
fgh clone <name> # if the repo is under your account
```

All repositories are cloned into the following structure:

```
~
└─ github
    └─ OWNER
        └─ TYPE
            └─ MAIN LANGUAGE
                └─ NAME
```

These names correspond to the following:

- `OWNER` is the owner of the repository
- `TYPE` is the type of the repository; one of the following:
  - `public`
  - `private`
  - `template`
  - `archived`
  - `disabled`
  - `mirror`
  - `fork`
- `MAIN LANGUAGE` is The main language that the repository contains. If no language is detected, `fgh` will map it to `Other`
- `NAME` is the name of the repository

# 💡 Example

```bash
fgh clone Matt-Gleich/fgh
```

Would clone to `~/github/Matt-Gleich/public/Go/fgh/`, `~` being `$HOME`. Once cloned, this path will be copied to your clipboard automatically.

> NOTE: On Linux machines running the X Window System, this requires the `xclip` or `xsel` packages.

This structure can be somewhat difficult to navigate in the terminal using conventional methods such as the use of the `cd` command. I suggest TUI-based filesystem navigators such as [ranger](https://github.com/ranger/ranger) to help speed up the process.

### ☝️ `fgh update`

If any of a repository's fields are changed, such as its type, main language, owner, or name, the path to your local repository won't match.

Running `fgh update` will iterate over your local repositories and checks if any of them need updates. If they do, `fgh` will ask you if you want to move the entire repository to that new path.

For example: If I had this repository cloned and later decided to archive it, its path would change from `~/github/Matt-Gleich/public/Go/fgh/` to `~/github/Matt-Gleich/archived/Go/fgh/`.

### 🧼 `fgh clean`

When you run this subcommand, `fgh` will check for the following on each repository:

1. Has it modified locally in a certain amount of time?
   > By default, this "amount of time" is 2 months. However, it can be changed with a flag! See `fgh clean --help` for more info.
2. Has the repository been deleted permanently on GitHub?

If either of those conditions are met, `fgh` will ask you if you would like to remove the aforementioned repository. It'll additionally show you some information about the repository itself.

> NOTE: This only removes the repo locally!

### 🗑 `fgh remove`

Remove a selected repository cloned locally. Usage is as follows:

```bash
fgh remove <owner/name>
```

## 🛣 Roadmap

- Inform the user of a dirty working tree before removing a repository
- Add `pull` subcommand to pull the latest changes for each repository

## 🙌 Contributing

Thank you for considering contributing to `fgh`! Before contributing, make sure to read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/fgh/blob/master/CONTRIBUTING.md).

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
