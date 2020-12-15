# Clove (Command line ♥️)

Clove is a declarative way to define CLI application flags, commands and arguments, that can use variable values or fixed set of values. Clove takes care of basic commands like `help` and `version` and makes sure the definitions you provide are not conflicting with each other (e.g. no two commands have the same abbreviation).

## CLI args order expectations

App that uses clove would support the following calling convention: 

`app <command> [<args>] [<flags>]`

## Clove schema

Clove operates on an array of strings (args) that excludes application executable name.

All entries have the same base schema and potentially extensions defined in the relevant section:

- Token - a single word that is mapped to this flag (`verbose`).
- Abbreviation - short form of a command that can be used instead of the token (`v`).
- Hint - _short_ description of the flag.
- Description - full and helpful description of the flag.

## Prefixes for flags and arguments

- `--` - requires a full token to be specified.
- `-` - allows using abbreviations. Example: `--debug --verbose` is the same as `-d -v`. 

## Flags

Application level flags that apply to most or all commands, generically. Examples: `debug`, `verbose`.

## Commands

Top level commands that allow users to use the application. Examples: `clone`, `checkout`.

Note: Command abbreviation can be any length (not just a single character), since there is only one command in a given args set.

- Arguments - all argument tokens that apply to this command.
- Example:
    - Argument - argument used in this example (should be one of the tokens in the Arguments).
    - Value - (optional) argument value we're using in the example.
    - Description.

### Schema extensions

- Arguments - if present, lists all arguments that are applicable to the command.

## Arguments

Arguments are parameters that commands might need for operation. Examples: `<url>` for `clone` command, `<branch>` for `checkout` command. Arguments that have value(s) have token separated from the value with a (space) or `=`.

### Schema extensions

All arguments are considered to be binary (presence of the argument is `true` and absence is `false`), unless they have Values specified or 

- Default - denotes a default argument that doesn't require a token. Default argument values are the first arguments for a command. That doesn't mean that the first argument specified automatically becomes default! Default argument can be specified in any position in the arguments list and there can only be one default argument for a command.
- Multiple - specifies whether argument can be used multiple times. Unless this is explicitly set to `true`, argument (and it's value if applicable) is only expected once. If set to `true`, this argument can be present several times, or if argument supports values (fixed or variable) - argument token can be specified once and multiple values can be provided (fixed or variable).
- Required - denotes a required argument that always must be present for a command.
- Values - lists all the value tokens that can be used with this argument. The list is exclusive - if provided no other values would be expected (e.g. `windows`, `macos`, `linux` for an `operating_system` argument). If the values are not provided, but the property is included (as an empty array), than this argument takes any arbitrary value (or values if Multiple is also specified) from the user. 

## Values

(Optional) For values the common schema values can be provided to allow hints, descriptions. Abbreviations are not supported for values (they'll be ignored).

## Clove app

Clove app itself supports several commands:

- `verify` - for a provided file, test all assumptions and provide feedback.
- `embed` - generate code snippet that represents the definition file.
- `gencode` - generate code file that switches based on commands and provides entry points for handlers.

Arguments:

- `filepath` - (default, required), applies to all command.

## Open issues

- Provide a way to generate boilerplate code to handle all the commands?
- Outline the test procedure and expectations. Would there be fatal errors and warnings?
- Spec state machine:
    - Token types.
    - Next possible tokens expected for a given type (token funnel).
- Validate that with arguments that support multiple (unspecified) values we can break out of the sequence with a valid argument token. Example: `--id 1 2 3 4 5 --type music`