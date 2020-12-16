# Clove (Command line ♥️)

Clove is a Golang module to build declarative description of a CLI application flags, commands and arguments, that can use variable values or fixed set of values. Clove than parses user provided application CLI parameters and returns a structured `Request`. Clove takes care of basic commands like `help` and `version` and makes sure the definitions you provide are not conflicting with each other (e.g. no two commands have the same abbreviation).

# Using clove in your app

TBD

## CLI args order expectations

App that uses clove would support the following calling convention:

`app <command> [<args>] [<flags>]`

When specifying arguments and flags `--` and `-` can be used interchangeably: `--verbose` is the same `-v` and `--v` and `-verbose`.

# Clove technical details

## Understanding common schema

Clove operates on a command line parameters array that excludes application executable name. Key entities in clove are commands, arguments, values and flags. When defining those entities you can use the following properties:

- `token` - (string) a single word that is mapped to this flag (`verbose`).
- `abbreviation` - (string, optional) short form of a command that can be used instead of the token (`v`).
- `hint` - (string, optional) _short_ description of the flag.
- `description` - (string, optional) full and helpful description of the flag.

Clove definitions file has additional properties that can be specified:

- `version` - (number) version of the definitions file format (1 is the latest right now)
- `env-prefix` - (string, optional) this prefix will be added to environment variables keys.

## Working with commands

Top level commands that allow users to control the application. Examples: `verify`.

### Commands schema extensions

- `arguments` - (optional) all argument tokens that apply to this command.
- `example`: (optional)
    - `argument` - argument used in this example (should be one of the tokens in the Arguments).
    - `value` - (optional) argument value we're using in the example.
    - `description`

## Working with arguments

Arguments are parameters that commands might need for operation. Examples: `path` for `verify` command. Argument value is separated from argument token with a space.

### Arguments schema extensions

Arguments allow passing parameters to commands. 
If arguments have `values` specified - they're expecting one or more value. If the `values` is an empty array - this argument would except any value from the user. If the `values` array is not empty, only those values would be considered valid.

If no `values` are specified argument will be considered binary, presence == `true` value and no argument == `false` value.

In addition to common schema arguments allow you to specify the following properties:

- `env` - (boolean, optional) read the value from environment variable if not specified in the command-line (CLI values are more specific and would take priority). Env key is a combination of ENV-PREFIX(when present)_COMMAND_ARGUMENT. For `env-prefix="CLV"`, `command="verify"` and `argument="path"` environment variable would be `CLV_VERIFY_PATH`.
- `default` - (boolean, optional) makes an argument default. Default argument doesn't require a token and can accept values right after the command - those values must be the first arguments for a command though. The first specified argument doesn't automatically become default - default argument can be specified in any position in the arguments list and there must be only one default argument for a command.
- `multiple` - specifies whether argument can be used multiple times. Unless this is explicitly set to `true`, argument (and a value, if applicable) is only expected once. If set to `true`, this argument can be present several times, or if argument supports values (fixed or variable) - argument token can be specified once and multiple values can be provided (fixed or variable).
- `required` - denotes a required argument that always must be present for a command.
- `values` - all the value tokens that can be used with this argument. The list is exclusive - if provided no other values would be supported. If the values are not provided, but the property is included (as an empty array), than this argument takes any arbitrary value(s) from a user.

## Working with values

When fixed values are defined for arguments, you can also specify the common schema values for hints and descriptions. Abbreviations are not supported for values - they'll be ignored if present.

## Working with flags

Flags allow to create application level hints or controls that apply to most or all commands, generically. All flags are optional. Examples: `verbose`.

# Clove app

Clove app can be used to validate definitions file:

`clove verify [--path] <path-to-definitions.json> [--verbose]`

Command details:

- `verify` - for a provided file, test all assumptions and provide feedback.

Argument details:

- `path` - (default, required) location of the definitions file to verify.

Flag details:

- `verbose` - display validation information as it happens.