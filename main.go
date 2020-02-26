package main

import (
	"fmt"

	"github.com/hpromonex/go-ecs/component"
	"github.com/hpromonex/go-ecs/entitymgr"
)

func main() {

	fmt.Printf("Component count: %d\n", component.Count)

	var bm entitymgr.Bitmatrix

	bm.Set(0, component.A, true)
	bm.Set(0, component.B, true)
	bm.Set(0, component.F, true)
	bm.Set(0, component.L, true)

	bm.Set(1, component.A, true)
	bm.Set(1, component.L, true)

	bm.Set(2, component.A, true)
	bm.Set(2, component.B, true)

	fmt.Println(bm)

	fmt.Printf("Query A: %v\n", bm.Query(component.A))
	fmt.Printf("Query A B: %v\n", bm.Query(component.A, component.B))
	fmt.Printf("Query A L: %v\n", bm.Query(component.A, component.L))
	fmt.Printf("Query L: %v\n", bm.Query(component.L))
	fmt.Printf("Query B: %v\n", bm.Query(component.B))

}
