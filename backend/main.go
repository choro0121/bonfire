package main

import (
    "bonfire/handler"
    "bonfire/model"
)

func main() {
    model.New()
    defer model.Close()

    handler.New()
    defer handler.Close()
}
