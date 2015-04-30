# PyReload
PyReload is a simple Go script that reruns a Python file every time a change is done in a specific folder or file.

## Installation

Simply clone this repository and run `go install`.

## Usage
PyReload currently receives two flags:

+ `-f`: path of file that is going to be run
+ `-d`: folder/file being watched for changes

## Requirements
This tool relies in [go-fsnotify](https://github.com/go-fsnotify/fsnotify). It must be installed.

## To Do
+ Extensible to other languages
+ Specify which version of Python to use
+ Colors
+ Time duration of runtime