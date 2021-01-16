# clo backlog

## Changes

- parsing sequence should be expressed in groups, not individual tokens. This would make progression clear, as well as allow to add things like "default command" more easily
- arguments should be "default" for a command, remove property on argument
- arguments should be "required" by a command, remove property on argument

## New

- default command for an app (e.g. "glo 123" -> "glo convert -g 123")
- default argument values - when none are specified (can be generated as "arg=value")
- argument "excludes" property - list argument tokens that can't be included with this argument (e.g. username / username-file)
- arguments that are required for some arg-values? E.g. username/password for fetch-type = account-product

## Follow-up

- update documentation and "generate"
