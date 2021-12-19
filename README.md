# gocli

A boilerplate generator to create CLIs in Golang.

**WORK IN PROGRESS**

## CLI Creator

### Highlights

* Generate a new CLI project in one command.
* Default commands every CLI should have available by default.
* Create environment variables for each flag automatically.
* Possibility to configure every flag via options, via config file, or both.

### Installing the CLI Creator

Download the latest release and put it wherever you want.

### Creating a New CLI

1. Generate a new CLI project with your project path and your Github username. For example: `./clicreator -p $HOME/myProject myGithubUsername`.
2. Go to the project's root and pull the latest dependencies with `go get -u`.
3. You can test the default command and add the ones you want.

### Argument and Options available

The username of your favorite VCS (Version Control System) platform is the only mandatory argument.

--------                -------------------------------         -----------------
Option                  Description                             Default
--------                -------------------------------         -----------------
`-p`                    Path of your new CLI                    `./new`
`-v`                    Address of your VCS platform            `github.com`
--------                -------------------------------         -----------------

## Your New CLI Project

### Default Commands for Your New CLI

* `help` - Display the help. The flags `--help` or `-h` have the same effect.
* `version` - Display the current CLI version (need to be build with goreleaser before).
* `completion` - Generate a completion for the common \*nix shells.
* `example` - Simple example which output a string. The boolean option `--example` (or `-e`) output another string. Fantastic!

### Compiling Your CLI

1. Go to the root of your project.
2. Run `go get -u` to get the latest dependencies.
3. Run `go build` to build the executable.

You can call your new project in your favorite shell.

### General Structure of the Project

The CLI project will be created with these directories:

* `cmd` - The different commands for your CLI.
* `install` - You can put different scripts to install your CLI here.
* `internal` - It's where the logic of your CLI should go.
* `internal/platform` - It's where any third party (API, database...) should be called, to isolate them from your main logic (in `internal`).
* `

The following files are also important:

* `main.go` - The first function run by your CLI, `main`, is in this file

### Adding New Commands

Your new CLI use the `cobra` library to manage commands and flags. Open the file `cmd/example.go` and you'll see three functions:

* `exampleFlags` - All the flags for this command are declared in this struct.
* `exampleCmd` - Create and return the command. The flags are also created and attached to the command. Their values are set in the struct `exampleFlags`.
* `doSomething` - Function called when the command run. This function can be anything you want.

You can copy everything in a new file (in the folder `cmd`, for example `cmd/mySuperCommand.go`), rename what you need to rename (the command is not named `example` anymore!), and add whatever flag you want.

When you create a command, it needs to be added to the root command via the function `Execute` in the file `cmd/root.go`.

### Environment Variables

Every flag you'll add can be set via environment variables automatically. TODO

### Configuration File

TODO

### Releasing Your CLI

TODO 

* Need a git tag
* Need to be in a clean state (git)

### Linux script

If you use a Linux-based OS, here's a simple way to download gocli and move it to `/usr/local/bin`. You can then call it wherever you want.

```shell
curl -L https://raw.githubusercontent.com/Phantas0s/<my_cli>/master/install/linux.sh | bash
```

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

## Building Your Mouseless Environment

Switching between a keyboard and mouse costs cognitive energy. [My book will help you set up a Linux-based development environment](https://themouseless.dev) that keeps your hands on your keyboard. Take the brain power you've been using to juggle input devices and focus it where it belongs: on the things you create.

## Licence

[Apache Licence 2.0](https://choosealicense.com/licenses/apache-2.0/)
