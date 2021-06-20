package main

type component interface {
    search(string)
    printNode(indent int)
}
