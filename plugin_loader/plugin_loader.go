package plugin_loader

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)


func ReloadPlugins(pluginFolder string) {
	log.Println("loading plugins in plugin folder ", pluginFolder)
	src_files := findAllSourceFiles(pluginFolder)
	for _, filepath := range src_files {
		log.Println("found ", filepath)
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