package main

import (
	"flag"
	"fmt"
)

func main()  {

	//not using https://github.com/spf13/cobra because this is just a POC
	pluginrepo := flag.String("pluginrepo", "git@github.com:harshpreet93/go-dynamic-code-loading-blog.git",
		"The Repo that contains the plugins to load")
	flag.Parse()

	fmt.Println("pluginrepo is "+*pluginrepo)

}


