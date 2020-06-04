package main

import (
    "github.com/server/route"
    "github.com/server/db"
)


func main() {
    db.SetupDB()
    _route := route.SetupRouter()

    _route.Run(":8000")
}
