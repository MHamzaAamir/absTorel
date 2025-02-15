package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Amount of Arguments")
		fmt.Println("Usage: go run main.go <current_directory> <relative_path>")
		return
	}

	curr_dir := os.Args[1]
	relative_path := os.Args[2]

	result := relativeToAbsolute(curr_dir, relative_path)
	fmt.Println(result)
}

func relativeToAbsolute(curr_dir string, relative_path string) string{

	path := strings.Split(curr_dir, "/")

	// handle start and end slashes if present

	if path[0] == ""{
		path = path[1:]
	}
	if path[len(path)-1] == ""{
		path = path[:len(path)-1]
	}

	i := len(path)-1  // pointer that points to current standing directory
	size := len(relative_path)
	dots := 0
	dir := ""

	for j := 0; j<size;j++{
		if relative_path[j] == '.'{
			if relative_path[j+1] != '.' && relative_path[j+1] != '/'{
				dir = dir + string(relative_path[j])
			}else{
				dots++
			}
			
		}else if relative_path[j] == '/'{
			if dots == 0{
				if dir != ""{
					path = path[:i+1]
					path = append(path, dir)
					dir = ""
					i = len(path) - 1
				}
			}else if dots == 1{
				dots = 0
			}else if dots == 2{
				dots = 0
				i--
			}
		}else{
			dir = dir + string(relative_path[j])
		}
	}
	if dir != ""{
		path = path[:i+1]
		path = append(path, dir)
	}

	absPath := strings.Join(path, "/")
	absPath = "/" + absPath
	
	return absPath

}