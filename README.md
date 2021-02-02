![Clo logo](clogo.png)

# Clo (Command line objectives)

Clo is a Golang module to build declarations of a command-line app objectives - commands and
arguments with values. Clo processes user provided command-line app input string (args) and returns
a structured `Request` object. Clo takes care of a `help` command (perhaps more in the future). Clo
can validate the definitions file and report any errors that would prevent normal operation.

## Using clo in your app

While clo has been designed with Go 1.16 embedding in mind, we'll provide directions on how to use
it today and will update once 1.16 is released.

### Using clo module

- Run `go get github.com/boggydigital/clo`
- Import `github.com/boggydgital/clo`
- Create declarations in `clo.json`
- Load definitions using `defs, _ := clo.LoadDefinitions("clo.json")`
- Call `clo.Parse(os.Args[1:], defs)` to get a `Request` struct
- Route to app code using `Request.Command` and pass `Request.Arguments`

Here is an example of a `main.go` that implements this approach (NOTE: error handling is omitted for
brevity)

```go
package main

import (
	"github.com/boggydigital/clo"
	"{your-app-module}/cmd"
	"os"
)

func main() {
	defs, _ := clo.LoadDefinitions("clo.json")

	// Parse `os.Args` using definitions to get `Request` data
	req, _ := clo.Parse(os.Args[1:], defs)

	// Route request to app command handlers
	cmd.Route(req, defs)
}
```

### Routing command requests and handling built-in commands

To route command-line objectives to app handlers you might add a `cmd/route.go` with a
single `Route` func that routes arguments data to command handlers.

NOTE: In order to allow clo to handle built-in commands, in your `Route` handler you need to
send `nil` and unknown commands to `clo.Route`.

Example:

```go
package cmd

import (
	"github.com/boggydigital/clo"
)

func Route(req *clo.Request, defs *clo.Definitions) error {

	// allow clo to handle nil requests (this will show help by default)
	if req == nil {
		return clo.Route(nil, defs)
	}

	switch req.Command {
	case "yourCommand":
		// route yourCommand here
		// ...
	default:
		// allow clo to handle unknown commands
		return clo.Route(req, defs)
	}

	return nil
}
```

### Getting values from a Request

`Request` provides few shortcuts to get values:

- `ArgVal(arg string)` - gets a single (first) value for an argument
- `ArgValues(arg string)` - gets all argument values specified in a `Request`
- `Flag(arg string)` - returns true is argument has been provided (with or without values)

Example:

```go
package cmd

import (
	"github.com/boggydigital/clo"
)

func Route(req *clo.Request, defs *clo.Definitions) error {
	if req == nil {
		return clo.Route(nil, defs)
	}
	switch req.Command {
	case "validate":
		return Validate(req.ArgVal("path"), req.Flag("verbose"))
	default:
		return clo.Route(req, defs)
	}
}
```

where `Validate` is defined as:

```go
package cmd

func Validate(path string, verbose bool) error { 
	
	// validate a file and display each test results if verbose was requested
	// ...
	
	return nil
}
```

## Command-line objectives calling convention

App that uses clo would support the following calling convention:

`app command [arguments [values]]`

- Commands don't have any prefix. Commands can be specified by a prefix - the first command to match
  provided prefix would be used.
- Arguments are specified with `-` or `--` prefixes. `--debug` is the same as `-debug` and, assuming
  there are no other arguments that start with `d`, it'll be the same as `--d` and `-d`
- Values are specified without any prefix.

## Creating clo.json definitions

Clo.json has the following top-level properties:

- `version` - version of the definitions file, currently the only supported version is `1`
- `cmd` - commands, arguments specified as a map:

```json
{
  "cmd": {
    "validate": [
      "path",
      "verbose"
    ]
  }
}
```

- `help` - help messages specified for topics (topic is a `:` delimited list of `command:argument`):

```json
{
  "help": {
    "clo": "command-line objectives",
    "validate": "validates that provided file doesn't have errors",
    "validate:path": "path to the file that should be validated",
    "validate:verbose": "print result of every validation test"
  }
}
```

### Specifying argument values

Command arguments can also specify values supported by that argument. If values are specified for an
argument, only those values would be considered valid and processing user input would stop when a
different value is specified. Values are specified following a `=` sign:

```json
{
  "cmd": {
    "command1": [
      "argument1=value1,value2",
      "argument2"
    ]
  }
}
```

In that declaration `argument1` can only be specified with either `value1` or `value2`,
while `argument2` can be specified with any arbitrary value.

### Objectives attributes

To provide more control to authors clo supports several inline attributes that can be used to define
constraints or set defaults:

- `_` - **default**. Applies to command, attributes, values:
    - When specified on a command, this command will be used if no other command has been a match.
      Example: `clo clo.json` is the same as `clo validate clo.json` if `validate_` command was
      specified as default.
    - When specified on an argument, this argument will be used if no other argument has been a
      match. Example: `clo validate clo.json` is the same as `clo validate --path clo.json`
      if `path_` was specified as default.
    - When specified on a value, this argument value pair will be added to `Request` if the user
      didn't provide another value for that pair.

- `!` - **required**. Applies to arguments. If set - argument value must be provided (can be in a form of
  default value). Example:

```json
{
  "cmd": {
    "validate_": [
      "path_!"
    ]
  }
}
```

- `...` - **multiple**. Applies to arguments. If set - argument can take more than one value. If not set - parsing multiple values for such argument would result in error.

- `$` - **env. variable**. Applies to arguments. If set - argument value can be read from env. variable, unless specified by the user.

## Clo app

clo app can be used to validate definitions file:

`clo validate [--path] <path-to-definitions.json> [--verbose]`

Command details:

- `validate` - validates that provided file doesn't have errors

Argument details:

- `path` - (default, required) path to the file that should be validated.
- `verbose` - print result of every validation test

As a reference here is how `clo.json` is defined for this app:

```json
{
  "version": 1,
  "cmd": {
    "validate_": [
      "path_$!",
      "verbose"
    ]
  },
  "help": {
    "clo": "command-line objectives",
    "validate": "validates that provided file doesn't have errors",
    "validate:path": "path to the file that should be validated",
    "validate:verbose": "print result of every validation test"
  }
}
```