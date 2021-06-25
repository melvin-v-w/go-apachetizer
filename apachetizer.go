package apachetizer

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// VHostConfDetector Function will detect files in a given path
//put all detected paths into one string array
//and return it.
func VHostConfDetector(cfgPath string) (files []string, err error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := f.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfo {
		files = append(files, cfgPath+file.Name())
	}
	return files, err
}

// VirtualHostConfig Type is the map that
//will be filled with information from the
//parser.
type VirtualHostConfig map[string]string

// VHostConfParser This is the parser function, you have to pass
//an io.Reader to to it, where it will read a
//apache virtualhost config and parse it to a
//map that is fully compatible with json.Marshal.
func VHostConfParser(file io.Reader) (VirtualHostConfig, error) {
	reader := bufio.NewReader(file)
	config := VirtualHostConfig{}
	for {
		line, err := reader.ReadString('\n')
		substring := " "
		if equal := strings.Index(line, substring); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+len(substring):])
				}
				key = strings.TrimPrefix(key, "\u003c/")
				key = strings.TrimSuffix(key, "\u003c/")
				value = strings.TrimSuffix(value, "\u003c/")
				value = strings.TrimPrefix(value, "\u003c/")

				key = strings.Replace(key, "\u003c", "", -1)
				key = strings.Replace(key, "\u003c/", "", -1)

				value = strings.Replace(value, "\u003e", "", -1)

				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
