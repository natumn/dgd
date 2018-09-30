package main

const usage = `
NAME:
   ddog - cli tool for cloud datastore entity and index documentation automatically generated from source code.

USAGE: 
   %s [files path]... [command options]...

   Files:
      *.go files:  create documentation from source code in the specified Go file.

      directory:   create documentation from source code in this package

  Options: 
      --type    <output file type>  document file type. ex html, json etc...

      --help                        show this command usage.

`
