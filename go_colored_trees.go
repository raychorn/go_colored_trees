package main

import (
	"fmt"
	"strings"
)

// Count the number of edges and apply colors that do not touch.
// colors are red, gree, blue
// Vertex node gets 2 edges.
// Last node on each side has 0 edges.
// All other nodes have 1 edge.

// State machine that builds 2 arrays, one array for left and one for right.
// Stop when number of edges is reached.
// Apply colors as the tree builds.

func main1() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

func main2() {
    var employee = map[string]int{"Mark": 10, "Sandy": 20,
        "Rocky": 30, "Rajiv": 40, "Kate": 50}
    for key, element := range employee {
        fmt.Println(key, "=>", element)
    }
}

func main3() {
	var vertex = map[int]int{}
    var left = map[int]int{}
	var right = map[int]int{}
	
	var j = 0
	for i:= 0; i < 10; i++ {
		if (j == 0) {
			vertex[0] = i
		} else {
			if (j % 2) == 1 {
				left[i] = i
			} else {
				right[i] = i
			}
		}
		j++
	}
	fmt.Println("vertex:")
    for key, element := range vertex {
        fmt.Println(key, "=>", element)
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("left:")
    for key, element := range left {
        fmt.Println(key, "=>", element)
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("right:")
    for key, element := range right{
        fmt.Println(key, "=>", element)
	}
	fmt.Println(strings.Repeat("-", 30))
}

func countEdges(nodes []node) int {
	n := len(nodes)
	edges := 0
    for i := range nodes {
		if (i < n-1) {
			edges++
		}
	}
	return edges
}

type node struct {
	value int
	color int
}

func nextColor(colors []string, ptr int) int {
	ptr++
	if (ptr > len(colors)-1) {
		ptr = 0
	}
	return ptr
}

type bucket struct {
	values []int
}

func (buck *bucket) prettyPrintBucket() string {
	output := "{"
	n := len(buck.values)
	for i, val := range buck.values {
		output += fmt.Sprintf("%d", val)
		if (i < n-1) {
			output += ", "
		}
	}
	output += "}"
	return output
}

func collectColorGroups(groups map[string]bucket, colors []string, ptr int, value int) {
	aBucket := groups[colors[ptr]]
	aBucket.values = append(aBucket.values, value)
	groups[colors[ptr]] = aBucket
}

func main() {
	var vertex = []node{}
    var left = []node{}
	var right = []node{}

	var colors = []string{"red", "green", "blue"}
	var colorPtr = 0

	var colorGroups = make(map[string]bucket)
	for _, aColor := range colors {
		colorGroups[aColor] = bucket{values: []int{}}
	}
	
	var j = 0
	var edgeCount = 0
	var i = 0  // starting value of vertex
	var maxEdges = 5
	var maxNodes = 100
	baseColor := -1
	leftArmNum := -1
	rightArmNum := -1
	//aBucket := bucket{values: []int{}}
	for {
		if (j == 0) {
			vertex = append(vertex, node{value:i, color: colorPtr})
			collectColorGroups(colorGroups, colors, colorPtr, i)
		} else {
			if (j % 2) == 1 {
				fmt.Printf("(1) edgeCount=%d\n", edgeCount)
				if (edgeCount+1 >= maxEdges) {
					break
				}
				leftArmNum = len(left)
				if (leftArmNum == 0) {
					baseColor = vertex[0].color
				} else {
					baseColor = left[leftArmNum-1].color
				}
				colorPtr = nextColor(colors, baseColor)
				left = append(left, node{value:i, color: colorPtr})
				collectColorGroups(colorGroups, colors, colorPtr, i)
				edgeCount += countEdges(left)
			} else {
				fmt.Printf("(2) edgeCount=%d\n", edgeCount)
				if (edgeCount+1 >= maxEdges) {
					break
				}
				rightArmNum = len(right)
				if (rightArmNum == 0) {
					baseColor = vertex[0].color
				} else {
					baseColor = right[rightArmNum-1].color
				}
				colorPtr = nextColor(colors, baseColor)
				right = append(right, node{value:i, color: colorPtr})
				collectColorGroups(colorGroups, colors, colorPtr, i)
				edgeCount += countEdges(right)
			}
		}
		fmt.Printf("(3) edgeCount=%d (%d)\n", edgeCount, maxEdges)
		if (i >= maxNodes) || (edgeCount >= maxEdges) {
			break
		}
		i++
		j++
	}
	fmt.Println("vertex:")
    for _, val := range vertex {
        fmt.Printf("%d : %s\n", val.value, colors[val.color])
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("left:")
    for _, val := range left {
        fmt.Printf("%d : %s\n", val.value, colors[val.color])
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("right:")
    for _, val := range right {
        fmt.Printf("%d : %s\n", val.value, colors[val.color])
	}
	fmt.Println(strings.Repeat("-", 30))
	numColorGroups := 0
    for key, val := range colorGroups {
		fmt.Printf("%s : %s\n", key, val.prettyPrintBucket())
		if (len(val.values) > 0) {
			numColorGroups++
		}
	}
	fmt.Printf("There are %d color groups.\n", numColorGroups)
	fmt.Println(strings.Repeat("-", 30))
}
