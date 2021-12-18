# gocli

A boilerplate generator to create CLIs in Golang.

**WORK IN PROGRESS**

## Highlights

* Generate a new CLI project in one command.
* Generate automatically the commands:
* Create automatically environment variables for each flag
* Can configure every flag of your CLI using a configuration file

## Installing the CLI Creator

Download the latest release and put it wherever you want.

## Creating a New CLI

1. Generate a new CLI project with your project path and your Github username. For example: `./clicreator -p $HOME/myProject myGithubUsername`.
2. Go to your project's root and pull the latest dependencies with `go get -u`.

## Default Commands of Your New CLI

* `help` - Display the help. The flags `--help` or `-h` have the same effect.
* `version` - Display the current CLI version (need to be build with goreleaser before).
* `completion` - Generate a completion for the common \*nix shells.
* `example` - Example of a command; output a string.

## Adding New Commands to your CLI

1. Go to the directory `cmd`
2. Create a new file
3. You can get inspired from the example command (`cmd/example.go`) to create your own command.

## Adding New Flags to your CLI

## Commands

## Argument and Options available

The username of your favorite VCS (Version Control System) platform is the only mandatory argument.

--------                -------------------------------         -----------------
Option                  Description                             Default
--------                -------------------------------         -----------------
`-p`                    Path of your new CLI                    `./new`
`-v`                    Address of your VCS platform            `github.com`
--------                -------------------------------         -----------------

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
