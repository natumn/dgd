package main

const usage = `
NAME:
   ddog - cli tool for cloud datastore entity and index documentation automatically generated from source code.

USAGE: 
   %s [options] <files>... 

   Files:
      Go files:  create documentation from source code in the specified Go file.

   Options: 
      --type <output file type>   specify document file type. For example, html, json etc...

      --help                      show this command usage.
			
      --version                   show ddog version.
`
