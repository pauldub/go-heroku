# go-heroku

A client for the Heroku HTTP API written in Go. This is a work in progress and is not yet in a usable state.

## TODO

 * Clean credentials.
 * Write documententation.
 * Write tests and a mocked API server.
 * Finish missing API methods.
 * Add examples.
 
## Installation

Currently this package is composed of 3 components: the client, the types and the api.

To use the client:

`go get github.com/PaulDub/go-heroku/client`

To use the types:

`go get github.com/PaulDub/go-heroku/types`

To use the api (essentially http.Request builders, client.HerokuClient role is mainly to authenticate and run those):

`go get github.com/PaulDub/go-heroku/api`

## Example: List applications

```
package main

import "fmt"
import heroku "github.com/PaulDub/go-heroku/client"

func main() {
	client := heroku.NewClient("Api key")
	
	apps, err := client.ListApplications()
	if err != nil {
		fmt.Println("An error occured:", err)
	} else if cap(apps) == 0 {
		fmt.Println("No applications.")
	} else {
		for _, app := range apps {
			fmt.Println("Application : ", app.Name)
		}
	}
}
```

## Author

Paul d'Hubert <paul@tymate.com>

## License

Copyright (c) 2013 Paul d'Hubert
All rights reserved.

Redistribution and use in source and binary forms are permitted
provided that the above copyright notice and this paragraph are
duplicated in all such forms and that any documentation,
advertising materials, and other materials related to such
distribution and use acknowledge that the software was developed
by Paul d'Hubert.  The name of the
<organization> may not be used to endorse or promote products derived
from this software without specific prior written permission.
THIS SOFTWARE IS PROVIDED ``AS IS'' AND WITHOUT ANY EXPRESS OR
IMPLIED WARRANTIES, INCLUDING, WITHOUT LIMITATION, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.
