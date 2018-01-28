# golang-clicommand

clicommand provides Go CLI applications with subcommand/api-style interfaces and option/parameter handling

[![GoDoc](https://godoc.org/github.com/leehuk/go-clicommand?status.svg)](https://godoc.org/github.com/leehuk/go-clicommand)
[![Go Report Card](https://goreportcard.com/badge/github.com/leehuk/go-clicommand)](https://goreportcard.com/report/github.com/leehuk/go-clicommand)
[![Build Status](https://travis-ci.org/leehuk/go-clicommand.svg?branch=master)](https://travis-ci.org/leehuk/go-clicommand)

## Overview
The clicommand library makes the creation of Go CLI applications using a subcommand
interface easier.  The subcommand interface is structured as a parent/child tree so
the application can mimic an api, with edges of the tree running custom Handler
functions and the tree providing a structured way of grouping commands, attaching
option arguments and finding additional parameters.

## Command Tree

Command objects are chained together to build a tree which has arbitrary depth, providing
it follows the tree rules:

* Each parent command within the tree may have any number of children.
* Each child command object within the tree has a single parent.
* Every child Command object within the tree can have its own children, except when it has a Handler function.

## CLI Application Pseudo Example

This allows building a CLI application which can mimic an API, e.g.:
```
./clicommand                         // parent, has children
./clicommand http                    // child of clicommand, has children itself
./clicommand http get => Handler()   // child of clicommand->http, calls Handler() when run.
                                     // Cannot have children.
./clicommand http post => Handler()  // child of clicommand->http, calls Handler() when run.
                                     // Cannot have children.
```

## CLI Options

Options can be attached to the tree at any point and these are inherited along the
tree, so child commands also have options from their parent commands.  Options
are defined as either having or not having parameters, options with parameters use
double dashes and options without parameters use single dashes as selectors.

## CLI Parameters

Anything the parser doesnt recognise is stored as a parameter, alloowing applications to accept
things as a simple generic parameter, rather than requiring its specified as an option.

## Autogenerated Help

As each command and option is added to the tree with a name and description, the parser can
automatically construct help information and display it when the program is run without
parameters, or the 'help' command is used.  The following example uses the sample helloworld
program from https://git.io/vNDug

```
[golang@1394e13b1fac helloworld]$ ./helloworld help
helloworld
Sample hello world program

Available subcommands:
    hello        Hello saying options

For help information run:
  'helloworld help' .. 'helloworld <commands>* help' .. 'helloworld [commands]* help [subcommand]*'
[golang@1394e13b1fac helloworld]$
```

```
[golang@1394e13b1fac helloworld]$ ./helloworld hello say help
helloworld hello say
Says something

helloworld hello options:
    -u                    Uppercase output
    -lower                Lowercase output

helloworld hello say options:
    --say <arg>            Required: Thing to say

For help information run:
  'helloworld help' .. 'helloworld <commands>* help' .. 'helloworld [commands]* help [subcommand]*'
[golang@1394e13b1fac helloworld]$
```

## Sample Program

A sample [helloworld.go](examples/helloworld/helloworld.go) program can be found under examples.
