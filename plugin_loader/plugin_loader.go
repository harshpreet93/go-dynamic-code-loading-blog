package plugin_loader

import (
	"github.com/emirpasic/gods/sets/hashset"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func ReloadPlugins(pluginFolder string) {
	log.Println("loading plugins in plugin folder ", pluginFolder)
	src_files := findAllSourceFiles(pluginFolder)
	loadedPlugins := hashset.Set{}

	for _, filepath := range src_files {
		log.Println("found ", filepath)
		if pluginNeedsToBeLoaded(filepath, loadedPlugins) {
			builtPath, _ := buildPlugin(filepath)
			loadPlugin(builtPath)
		}

	}
}

func findAllSourceFiles(pluginFolder string) []string {
	var files []string
	log.Println("trying to find source files in ", pluginFolder)
	filepath.Walk(pluginFolder, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func buildPlugin(srcPath string) (string, error) {
	os.Exec
}

func loadPlugin(builtPath string) error {
	return nil
}

func pluginNeedsToBeLoaded(filepath string, alreadyLoadedPlugins hashset.Set) bool {
	return false
}