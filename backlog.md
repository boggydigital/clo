# clo backlog

## Changes

- ~~Roll back the version to v0.5.0-beta~~
- ~~deprecate examples~~
- ~~deprecate hint/desc - only leave help~~
- ~~deprecate flags~~
- ~~deprecate values~~
- ~~arguments should be "default" for a command, remove property on argument~~
- ~~arguments should be "required" by a command, remove property on argument~~
- ~~simplify value tokens to just one type "value"~~
- ~~parsing sequence should be expressed in groups, not individual tokens. This would make progression clear, as well as allow adding things like "default command" more easily~~

## New

- default command for an app (e.g. "glo 123" -> "glo convert -g 123")
- default argument values - when none are specified (can be generated as "arg=value")
- argument "excludes" property - list argument tokens that can't be included with this argument (e.g. username / username-file)
- arguments that are required for some arg-values? E.g. username/password for fetch-type = account-product

## Follow-up

- review and update unit tests based on the new logic
- update documentation and "generate"
