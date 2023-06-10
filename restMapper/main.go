package main

import (
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	// hard coded resources
	// res := "pods"

	// allowing users to pass resource as input
	var res string

	flag.StringVar(&res, "res", "", "resource passed by user as input")
	flag.Parse()

	configFlag := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchVersionFlag := cmdutil.NewMatchVersionFlags(configFlag)
	m, err := cmdutil.NewFactory(matchVersionFlag).ToRESTMapper()
	if err != nil {
		fmt.Printf("getting rest mapper from newFactory %s", err.Error())
		return
	}
	gvr, err := m.ResourceFor(schema.GroupVersionResource{
		Resource: res,
	})
	if err != nil {
		fmt.Printf("gvr %s", err.Error())
		return
	}

	fmt.Printf("gruop %s, version %s, resource %s\n", gvr.Group, gvr.Version, gvr.Resource)
}
