package File

import "strings"

/*
addZip := makeSuffix(".zip")
addTgz := makeSuffix(".tgz.gz")
fmt.PrintLn(addZip("go")) => go.zip
fmt.PrintLn(addTgz("go")) => go.tgz.gz
*/

func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
