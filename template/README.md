# gocli

A template to create CLI in Golang.

## Highlights

* Default commands (help, version).
* Possibility to configure each flag via config file.
* Automatic generation of environment variable.
* Easy way to create app releases.
* Script to automatically install your CLI on Linux.

## Installation

1. Clone this repository with the name of your new CLI (replace `greatcli` with something else):

```
git clone https://github.com/Phantas0s/gocli greatcli
```

2. Run `go build` in your new project. You've just compiled the example CLI!
3. Run `./gocli`: the CLI binary works.

## Usage

### Trying the Example

The example CLI has three commands by default:

* `version` - Display the CLI's version.
* `help` - Display the help of the command. For the general help: `./gocli help`. For a command's help: `./gocli <command> --help`.
* `example` - A dummy command.

If you try `./gocli version`, you'll notice that the version number is missing. It's because your CLI needs to be released to have a version. See the section "Releasing your Cli".

### Creating your CLI

First, let's edit main.go. Replace every instance of gocli with the name of your cli.


#### Defining the Command and the Flags

### Adding A Flag

## Releasing Your CLI

* Need a git tag
* Need to be in a clean state (git)

### Linux script

If you use a Linux-based OS, here's a simple way to download gocli and move it to `/usr/local/bin`. You can then call it wherever you want.

```shell
curl -L https://raw.githubusercontent.com/Phantas0s/<my_cli>/master/install/linux.sh | bash
```

## Navigation

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

## Building Your Mouseless Environment

Switching between a keyboard and mouse costs cognitive energy. [My book will help you set up a Linux-based development environment](https://themouseless.dev) that keeps your hands on your keyboard. Take the brain power you've been using to juggle input devices and focus it where it belongs: on the things you create.

## Licence

[Apache Licence 2.0](https://choosealicense.com/licenses/apache-2.0/)
