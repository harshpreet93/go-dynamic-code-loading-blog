package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const plugin_source_folder_path  = "plugin_source"

func main()  {

	//not using https://github.com/spf13/cobra because this is just a POC
	pluginrepo := flag.String("pluginrepo", "git@github.com:harshpreet93/go-dynamic-code-loading-blog.git",
		"The Repo that contains the plugins to load")
	flag.Parse()

	fmt.Println("pluginrepo is "+*pluginrepo)

	cloneRepo(*pluginrepo)

}

func cloneRepo(repo string) {
	//make sure
	ensure_plugin_repo_folder_readiness()
}

func ensure_plugin_repo_folder_readiness()  {
	// clean up /var/plugin_loader/plugin_source/
	plugin_source_path := filepath.Join(".", plugin_source_folder_path)
	os.RemoveAll(plugin_source_path)
	err := os.MkdirAll(plugin_source_path, os.ModePerm)
	if err != nil {
		log.Println("error creating plugin source folder", err)
		return
	}
	
}