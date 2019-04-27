package main

import (
	"flag"
	"fmt"
	"github.com/harshpreet93/go-dynamic-code-loading-blog/plugin_loader"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"os"
	"path/filepath"
	"time"
)

const plugin_source_folder_path = "plugin_source"

func main() {

	//not using https://github.com/spf13/cobra because this is just a POC
	pluginrepo := flag.String("pluginrepo", "https://github.com/harshpreet93/go-dynamic-code-loading-blog.git",
		"The Repo that contains the plugins to load")
	flag.Parse()

	fmt.Println("pluginrepo is " + *pluginrepo)

	for range time.Tick(30 * time.Second) {
		createWorkspaceAndTest(*pluginrepo)
	}

}

func createWorkspaceAndTest(repo string) {
	//make sure
	ensurePluginRepoFolderReadiness(repo)
	loadedPlugins := plugin_loader.ReloadPlugins(plugin_source_folder_path + "/plugins")

	// since we have the plugins......let's go further and actually use them to prove that they are....usable
	for _, plugin := range loadedPlugins {
		v, err := plugin.Lookup("Run")
		if err != nil {
			log.Println("error finding Run symbol in plugin")
		}
		v.(func())()
	}

}

func ensurePluginRepoFolderReadiness(repo string) {

	const pluginSourceFolderPath = "plugin_source"

	plugin_source_path := filepath.Join(".", pluginSourceFolderPath)
	os.RemoveAll(plugin_source_path)
	err := os.MkdirAll(plugin_source_path, os.ModePerm)
	if err != nil {
		log.Println("error creating plugin source folder", err)
		return
	}

	r, err := git.PlainClone(plugin_source_path, false, &git.CloneOptions{
		URL: repo,
	})
	if err != nil {
		log.Println("error cloning repo ", err)
		return
	}
	fmt.Println("cloned ", r)

}
