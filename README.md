<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

# fgh

📁 Automate your local GitHub workspace

![build](https://github.com/Matt-Gleich/fgh/workflows/build/badge.svg)
![test](https://github.com/Matt-Gleich/fgh/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/fgh/workflows/lint/badge.svg)
![release](https://github.com/Matt-Gleich/fgh/workflows/release/badge.svg)

- [🚀 Install](#---install)
  - [🍎 macOS](#---macos)
  - [🐧 Linux and 🖥 Windows](#---linux-and----windows)
- [📖 Documentation](#---documentation)
  - [⚙️ fgh configure](#----fgh-configure-)
  - [☁️ fgh clone](#----fgh-clone-)
  - [☝️ fgh update](#----fgh-update-)
  - [🧼 fgh clean](#----fgh-clean-)
- [🛣 Road Map](#---road-map)
- [🙌 Contributing](#---contributing)
- [👥 Contributors](#---contributors)

## 🚀 Install

### 🍎 macOS

```txt
brew tap Matt-Gleich/homebrew-taps && brew install fgh
```

### 🐧 Linux and 🖥 Windows

Download the binary from the [latest release](https://github.com/Matt-Gleich/fgh/releases/latest)

## 📖 Documentation

As you work on more and more git repos you need to clone and organize them on your machine. This can take time and usually doesn't scale. This CLI (command line application), called fgh, manages this all for by automating the entire process.

### ⚙️ `fgh configure`

Before using fgh you need to configure it by running `fgh configure`. When it asks you for the GitHub PAT (personal access token) just go to [https://github.com/settings/tokens/new](https://github.com/settings/tokens/new) and create a new token with the repo box check off.

### ☁️ `fgh clone`

It all starts with cloning a repo which is done when you run `fgh clone OWNER/NAME` or `fgh clone NAME` if the repo is under your account. All repos are cloned in the following structure:

```txt
~
└──github
   └── OWNER
       └── TYPE
           └── MAIN LANGUAGE
               └── NAME
```

- `OWNER`: The owner of the repo.
- `TYPE`: The type of the repo; one of the following:
  - `public`
  - `private`
  - `template`
  - `archived`
  - `disabled`
  - `mirror`
  - `fork`
- `MAIN LANGUAGE`: The main language for the repo. If no language is detected then will just be `Other`.
- `Name`: The name of the repo.

So if you were to clone this repo using `fgh clone Matt-Gleich/fgh` it would be cloned to the to `~/github/Matt-Gleich/public/Go/fgh/`.

Once you are done cloning the repo you can have the path copied to your clipboard automatically. If you are on Linux you need `xclip` or `xsel` for this to work.

This structure can be somewhat difficult to navigate in the terminal using conventional methods. I suggest navigators such as [ranger](https://github.com/ranger/ranger) to help speed up the process.

### ☝️ `fgh update`

If a repo changes its type, main language, owner, or name the path to your local repo won't match. Running `fgh update` will check every repo to see if it should have a new path. If it does it will ask you if you want to move the entire repo to that new path. So if I had this repo cloned and then I was to archive it the path would change from `~/github/Matt-Gleich/public/Go/fgh/` to `~/github/Matt-Gleich/archived/Go/fgh/`.

### 🧼 `fgh clean`

When you run this command fgh will check every single repo for two things:

1. If it hasn't been modified locally in a certain amount of time. The default amount of time is 2 months but this can be changed with flags. See `fgh clean --help` for more info.
2. If the repo has been deleted on GitHub.

If either of those conditions are met fgh will ask you if you would like to remove it and shows you some information about the repo. **This only removes the repo locally**.

## 🛣 Road Map

- Tell the user if they have a dirty working tree before removing a repo.
- Add pull command to pull the latest changes for every repo.

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/fgh/blob/master/CONTRIBUTING.md).

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
