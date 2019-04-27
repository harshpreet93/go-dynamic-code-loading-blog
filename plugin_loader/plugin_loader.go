package plugin_loader

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strings"
)

var loadedPluginHashes *hashset.Set
var plugins []*plugin.Plugin

func ReloadPlugins(pluginFolder string) []*plugin.Plugin {
	log.Println("loading plugins in plugin folder ", pluginFolder)
	srcFiles := findAllSourceFiles(pluginFolder)
	if loadedPluginHashes == nil {
		loadedPluginHashes = hashset.New()
	}
	for _, filepath := range srcFiles {
		log.Println("found ", filepath)
		builtPath, err := buildPlugin(filepath)
		if err == nil {
			//in order to not load the same file over and over again let's store the md5 hash of the plugin we just loaded
			hash, err := getMD5(filepath)

			if err == nil && loadedPluginHashes.Contains(hash) {
				log.Println("plugin at ", filepath, " with hash ", hash, " has already been loaded")
				continue
			}

			pluginLoaded, err := loadPlugin(builtPath)
			if err != nil {
				log.Println("unable to load plugin ", filepath)
				continue
			}
			plugins = append(plugins, pluginLoaded)
			loadedPluginHashes.Add(hash)
		}
		log.Println("number of loaded plugins ", len(plugins), " hash cache size ", loadedPluginHashes.Size())
	}
	return plugins
}

func getMD5(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	hashInBytes := h.Sum(nil)[:16]
	md5String := hex.EncodeToString(hashInBytes)
	log.Println("file ", filepath, " hash is ", md5String)
	return md5String, nil
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
	return strings.TrimSuffix(srcPath, ".go") + ".so", err
}

func loadPlugin(builtPath string) (*plugin.Plugin, error) {
	log.Println("loading ", builtPath)
	p, err := plugin.Open(builtPath)
	if err != nil {
		return nil, err
	}
	return p, nil
}