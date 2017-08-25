# Golang Git Build Vars

This is a quick demo of how to set information about your current git repository and build time
as variables within your Go binary.

## How it works

1. Gather information and set `-ldflags` during build in the [`Makefile`](./Makefile)
2. Define variables in your Go code that will be set during build

## Variables

This example sets 3 variables:

1. Timestamp of when the build happens
2. The current git commit hash, including a `✗-` prefix if the working directory is dirty
3. The current git branch

## Run & example output

```
$ make
$ bin/vars
```
