Clo (Command line objects)
=======================

Clo is a Golang module to build declarative description of a CLI application flags, commands and
arguments. Clo parses user provided application CLI parameters and returns a structured `Request`.
Clo takes care of basic commands like `help` and `version` and makes sure the definitions you
provide are not conflicting with each other (e.g. no two commands have the same abbreviation).

# Using clo in your app

While clo has been designed with Go 1.16 embedding in mind, we'll provide directions on how to use
it today and will update once 1.16 is released.

## Using clo module

- Run `go get github.com/boggydigital/clo`
- In your app import `github.com/boggydgital/clo`
- Load definitions with `LoadDefinitions`, 

## Common clo use patterns

Apps that use clo might start with the following general sequence of actions:

- Add a definitions file named `clo.json` in the root of your project
- Parse `os.Args` (using default definitions) to get `Request` data with command, arguments, flags
- Dispatch request to your commands handlers and fallback to `clo.Dispatch` for unknown commands or nil `Request`

Here is an example of a `main.go` that implements this approach (NOTE: error handling is omitted for brevity)

```
package main

import (
	"fmt"
	"github.com/boggydigital/clo"
	"{your-app-module}/cmd"
	"io/ioutil"
	"os"
)

func main() {
    // Parse `os.Args` using definitions to get `Request` data
	req, _ := clo.Parse(os.Args[1:])

    // Dispatch request to command handlers
	cmd.Dispatch(req)
}
```

## Dispatching command requests and handling built-in commands

The recommended approach to handle `Request` commands, arguments and flags is to have
a `cmd/dispatch.go` single `Dispatch` method that routes arguments, flags data to command handlers.

In order to allow clo to handle built-in commands, in your `Dispatch` handler you need to
send `nil` and unknown commands to `clo.Dispatch`.

Example:

```
package cmd

import (
	"github.com/boggydigital/clo"
)

func Dispatch(request *clo.Request) error {
	
	// allow clo to handle nil requests (this will show help by default)
	if request == nil {
		return clo.Dispatch(nil)
	}

    switch request.Command {
    // case "yourCommand":
    // dispatch yourCommand here
    // ...
	default:
	    // allow clo to handle unknown commands
		return clo.Dispatch(request)
	}
  }
}
```

## CLI args order expectations

App that uses clo would support the following calling convention:

`app <command> [<args>] [<flags>]`

When specifying arguments and flags `--` and `-` can be used interchangeably: `--verbose` is the
same `-v` and `--v` and `-verbose`.

# Clo technical details

## Understanding common schema

Clo operates on a command line parameters array that excludes application executable name. Key
entities in clo are commands, arguments, values and flags. When defining those entities you can
use the following properties:

- `token` - (string) a single word that is mapped to this flag (`verbose`).
- `abbreviation` - (string, optional) short form of a command that can be used instead of the
  token (`v`).
- `hint` - (string, optional) _short_ description of the flag.
- `description` - (string, optional) full and helpful description of the flag.

Clo definitions file has additional properties that can be specified:

- `version` - (number) version of the definitions file format (1 is the latest right now)
- `env-prefix` - (string, optional) this prefix will be added to environment variables keys.
- `hint` - short application description
- `desc` - verbose application description

## Working with commands

Top level commands that allow users to control the application. Examples: `verify`.

### Commands schema extensions

- `arguments` - (optional) all argument tokens that apply to this command.
- `examples`: (optional)
    - `argumentsValues` - map of an arguments, each with a set of values that are used in this
      example. Argument key should be one of the values in the `arguments` definitions. If an
      argument has defined values - values for that key should be subset of defined values.
    - `description` - (optional) description of this example

### Built-in 'help' command

Clo provides a built-in 'help' command, unless one already exists, provided by you. Built-in '
help' commands uses certain conventions to avoid conflicts with user arguments and to support any
commands you might have created. Here is what you need to know:

- 'help' command is added right before parsing CLI args
- 'help' command is only added if there is not command with a token 'help' already
- 'help' command will attempt to add 'h' abbreviation if it's not used by any other command
- 'help' command will attempt to add 'help:command' argument if doesn't exist already
- 'help:command' argument uses special 'from:commands' value that will be expanded with all commands
  declared in clo.json (technically, you can use that value and if it's the only value, it'll be
  expanded)

Please see [Dispatching command requests and handling built-in commands](#dispatching-command-requests-and-handling-built-in-commands) to understand what needs to be
done to support 'help' command.

## Working with arguments

Arguments are parameters that commands might need for operation. Examples: `path` for `verify`
command. Argument value is separated from argument token with a space.

### Arguments schema extensions

Arguments allow passing parameters to commands. If arguments have `values` specified - they're
expecting one or more value. If the `values` is an empty array - this argument would except any
value from the user. If the `values` array is not empty, only those values would be considered
valid.

If no `values` are specified argument will be considered binary, presence == `true` value and no
argument == `false` value.

In addition to common schema arguments allow you to specify the following properties:

- `env` - (boolean, optional) read the value from environment variable if not specified in the
  command-line (CLI values are more specific and would take priority). Env key is a combination of
  ENV-PREFIX(when present)_COMMAND_ARGUMENT. For `env-prefix="CLV"`, `command="verify"`
  and `argument="path"` environment variable would be `CLV_VERIFY_PATH`.
- `default` - (boolean, optional) makes an argument default. Default argument doesn't require a
  token and can accept values right after the command - those values must be the first arguments for
  a command though. The first specified argument doesn't automatically become default - default
  argument can be specified in any position in the arguments list and there must be only one default
  argument for a command.
- `multiple` - specifies whether argument can be used multiple times. Unless this is explicitly set
  to `true`, argument (and a value, if applicable) is only expected once. If set to `true`, this
  argument can be present several times, or if argument supports values (fixed or variable) -
  argument token can be specified once and multiple values can be provided (fixed or variable).
- `required` - denotes a required argument that always must be present for a command.
- `values` - all the value tokens that can be used with this argument. The list is exclusive - if
  provided no other values would be supported. If the values are not provided, but the property is
  included (as an empty array), than this argument takes any arbitrary value(s) from a user.

## Working with values

When fixed values are defined for arguments, you can also specify the common schema values for hints
and descriptions. Abbreviations are not supported for values - they'll be ignored if present.

## Working with flags

Flags allow creating application level hints or controls that apply to most or all commands,
generically. All flags are optional. Examples: `verbose`.

# Clo app

clo app can be used to validate definitions file:

`clo verify [--path] <path-to-definitions.json> [--verbose]`

Command details:

- `verify` - for a provided file, test all assumptions and provide feedback.

Argument details:

- `path` - (default, required) location of the definitions file to verify.

Flag details:

- `verbose` - display validation information as it happens.