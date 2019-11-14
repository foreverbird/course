package util

import "io/ioutil"

func WriteResourceFile(data string, resource string) {
	writeFile(data, resource)
}

func readFile(filePath string) string {
	c, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(c)
}

func writeFile(data string, path string) {
	ioutil.WriteFile(path, []byte(data), 0666)
}
