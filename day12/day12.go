package day12

import (
	"bufio"
	"github.com/skamoen/advent2021/util"
	"log"
	"os"
	"strings"
)

type d struct {
}

func Get() util.Entry {
	return &d{}
}

func (*d) Run() (int, int) {
	//file, err := os.Open("./day12/example.txt")
	file, err := os.Open("./day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var caves = make(map[string]*node)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		path := strings.Split(line, "-")
		f, t := path[0], path[1]

		from := getOrCreate(f, caves)
		to := getOrCreate(t, caves)

		from.connections = append(from.connections, to)
		to.connections = append(to.connections, from)
	}

	p := &path{
		path: []string{"start"},
	}

	q := p.findNext(caves)
	var partOne int
	var partTwo int

	for i := 0; len(q) > 0; i++ {
		currentPath := q[0]
		q = q[1:]
		if !currentPath.done {
			nextPaths := currentPath.findNext(caves)
			q = append(nextPaths, q...)
		} else {
			if currentPath.path[len(currentPath.path)-1] == "end" {
				if !currentPath.small {
					partOne++
				}
				partTwo++
			}
		}
	}

	return partOne, partTwo
}

type path struct {
	path  []string
	done  bool
	small bool
}

func (p *path) findNext(caves map[string]*node) []*path {
	currentNode := caves[p.path[len(p.path)-1]]
	var newPaths []*path
	if p.done {
		return nil
	}

	for _, connection := range currentNode.connections {
		if connection.name == "start" {
			continue
		}
		if p.small {
			if connection.isSmall && util.ArrayContains(p.path, connection.name) {
				continue
			}
		}

		pCopy := &path{path: make([]string, len(p.path)), done: false, small: false}
		if p.small {
			pCopy.small = true
		}
		copy(pCopy.path, p.path)
		pCopy.path = append(pCopy.path, connection.name)
		if connection.isSmall && util.ArrayContains(p.path, connection.name) {
			pCopy.small = true
		}

		if connection.name == "end" {
			pCopy.done = true
		}
		newPaths = append(newPaths, pCopy)
	}
	if newPaths == nil {
		p.done = true
	}
	return newPaths

}

func getOrCreate(name string, caves map[string]*node) *node {
	n, ok := caves[name]
	if ok {
		return n
	} else {
		isSmall := strings.ToUpper(name) != name
		nn := &node{isSmall: isSmall, name: name}
		caves[name] = nn
		return nn
	}
}

type node struct {
	isSmall     bool
	connections []*node
	name        string
}
