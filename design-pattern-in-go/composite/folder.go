package main

import "fmt"

type folder struct {
    components []component
    name       string
}

func (f *folder) search(keyword string) {
    fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
    for _, composite := range f.components {
        composite.search(keyword)
    }
}

func (f *folder) printNode(indent int) {
    for i := 0; i < indent; i++ {
        fmt.Printf("\t")
    }
    fmt.Printf("+%s\n", f.name)
    for _, composite := range f.components {
        composite.printNode(indent + 1)
    }
}

func (f *folder) add(c component) {
    f.components = append(f.components, c)
}

