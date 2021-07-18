package main

import (
	"fmt"

	"github.com/afriedrichsen/k8tering/flux"
	"github.com/afriedrichsen/k8tering/k3d"
)

func main() {
	fmt.Println("K8tering Cluster Maintenance")
	res, _ := k3d.K3D(true, true, "", "cluster", "create", "apartment0", "--servers", "3", "--agents", "3")
	fmt.Println(res)
	bootstrap, _ := flux.Flux(true, true, "", "bootstrap")
	fmt.Println(bootstrap)

}
