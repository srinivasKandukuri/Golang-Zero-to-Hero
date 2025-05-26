package main

import "fmt"

type parent struct {
	imageName string
	id        int
	registry  string
}

var imageParents = []parent{
	{
		imageName: "PHP",
		id:        3,
		registry:  "Adhoc",
	},
	{
		imageName: "ubuntu",
		id:        1,
		registry:  "docker.io",
	},
	{
		imageName: "nginx",
		id:        2,
		registry:  "docker.io",
	},
	{
		imageName: "postgres",
		id:        3,
		registry:  "docker.io",
	},
}

func main() {

	imps := filterBasedOnId(imageParents)

	fmt.Println(imps)
}

func filterBasedOnId(imageParents []parent) []parent {
	var imps []parent
	var unique = make(map[int]parent)
	for _, v := range imageParents {
		if existing, found := unique[v.id]; found {
			if existing.registry != "Adhoc" {
				continue
			}
			unique[v.id] = v
		} else {
			unique[v.id] = v
		}
	}
	for _, v := range unique {
		imps = append(imps, v)
	}
	return imps
}
