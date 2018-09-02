# shake
Rhymes with "make". A work in progress.

## Install
Clone this repo and run `go install`. I don't have an install script written
yet.

## Usage
You must have a file called "Shakefile" in your current directory. It must be
a valid YAML file. A Shakefile may look like this:

```

---
# shake sets environment variables before running command
vars:
  filename: $main.go

targets:
  build:
    - go build $filename #shake will expand this variable

  test:
    - go test ./...
```


## Motivation
I found myself writing a lot of scripts to do a lot of different actions in
one go, and I found Makefiles somewhat unsuitable for this. Makefiles were
written for things like C, and they work very well for C, but not for building
a Go program and deploying to AWS in one command.

So I thought it would be cool to just specify a YAML file with build targets.
