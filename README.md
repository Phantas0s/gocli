# gocli

**LOGO**

![travis CLI]() [![Go Report Card]()]() [![Codacy Badge]()]() [![Hits-of-Code]()]() [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

**TODO - automate this tweet badge**

[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Gocli%20-%20CLI%20:&url=https%3A%2F%2Fgithub.com%2Fphantas0s%2Fgocli)

A template to create CLI. This is an experiment with `go generate`, see if it's possible to create the whole boilerplate to create CLIs easily.



* It's great!
* Awesome

**SCREENSHOT**

**It's early version**

## Installation

### General
### Linux script

If you use a Linux-based OS, here's a simple way to download gocli and move it to `/usr/local/bin`. You can then call it wherever you want.

```shell
curl -L https://raw.githubusercontent.com/Phantas0s/<my_cli>/master/install/linux.sh | bash
```
### Manual installation

## Authorization

### Steps
### XDG Home Directory

The value of `$XDG_CONFIG_HOME` depends of your OS. Here are the defaults (if you didn't modify it):

* **Unix systems**: `~/.config`
* **macOS**: `~/Library/Application Support`
* **Windows**: `%LOCALAPPDATA%`

## Commands

Use the option `-h` for each command to output the help.

### Other

## Usage

If you choose to use the TUI, you can select a page and open it with your favorite browser using the `ENTER` key.

| Description      | Command        |
| ----             | ----           |
| Display nonsense | `gocli other`  |

## Option, Configuration, and Environment Variables

You can provide the different options to gocli using:

1. Command-line options 
2. Environment variables 
3. Configuration file 

If these options are defined multiple times, the priorities follow the order above (from higher priority to lower).

The names of the environment variables need to be uppercase and prefixed with `GO_CLI_`. Every hyphen `-` in the option's name needs to be replaced with an underscore `_`.

## TUI Keystrokes

### Navigation

<pre>
 <kbd>↑</kbd> or <kbd>k</kbd>: up
 <kbd>↓</kbd> or <kbd>j</kbd>: down
 <kbd>PgUp</kbd> or <kbd>CTRL</kbd>+<kbd>u</kbd>: One screen up
 <kbd>PgDn</kbd> or <kbd>CTRL</kbd>+<kbd>d</kbd>: One screen down
 <kbd>Home</kbd> or <kbd>g</kbd>: Top of the list
 <kbd>End</kbd> or <kbd>G</kbd>: Bottom of the list
</pre>

### Action

<pre>
 <kbd>ENTER</kbd>: Open the selected page with your favorite browser
 <kbd>d</kbd>: Delete Pocket entry
 <kbd>a</kbd>: Add (if list archive) or archive (if list unread)
</pre>

## Sponsorship

Consider [sponsoring my work](https://github.com/sponsors/Phantas0s) if you want to see new, fresh, and crunchy little CLIs (and TUIs) all over your system.

## Shameless Mouseless Plug

Switching between a keyboard and mouse costs cognitive energy. [My book will help you set up a Linux-based development environment](https://themouseless.dev) that keeps your hands on your keyboard. Take the brain power you've been using to juggle input devices and focus it where it belongs: on the things you create.

## Acknowledgements

* Thanks to [MariaLetta](https://github.com/MariaLetta/free-gophers-pack) for the awesome and beautiful Gopher pack! I used it for my logo on top.
* Thanks to [Lukasz Adam](https://lukaszadam.com/illustrations) for his free and amazing illustrations I use basically everywhere.

## Interesting Articles

[Go generate](https://blog.gopheracademy.com/advent-2015/reducing-boilerplate-with-go-generate/)
[man pages](https://www.golinuxcloud.com/create-man-page-template-linux-with-examples/)


## Licence

[Apache Licence 2.0](https://choosealicense.com/licenses/apache-2.0/)
