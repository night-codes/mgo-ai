# mgo-autoincrement (ai)
Package **ai** implements AutoIncrement methods for mgo(golang) 

## How To Install

```
go get github.com/night-codes/mgo-ai
```

## Getting Started

```go
package main

import (
    "github.com/night-codes/mgo-ai"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func main() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }

    // connect AutoIncrement to collection "counters"
    ai.Connect(session.DB("example-db").C("counters"))

    // ...

    // use AutoIncrement
    session.DB("example-db").C("users").Insert(bson.M{
        "_id":   ai.Next("users"),
        "login": "test",
        "age":   32,
    })
}

```

## License
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2016 Oleksiy Chechel <alex.mirrr@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
