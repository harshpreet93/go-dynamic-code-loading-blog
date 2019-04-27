package plugin_loader

import (
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"log"
	"os"
	"os/exec"
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
	//go build -buildmode=plugin
	cmd := exec.Command("go", "generate", srcPath)
	wd, _ := os.Getwd()
	fmt.Println("wd is ", wd)
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("build output ", errOut.String(), stdOut.String())
	}
	return "", err
}

func loadPlugin(builtPath string) error {
	return nil
}

func pluginNeedsToBeLoaded(filepath string, alreadyLoadedPlugins hashset.Set) bool {
	return true
}