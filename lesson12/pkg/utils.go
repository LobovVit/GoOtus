package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func ReadDir(dir string) (map[string]string, error) {

	var envMap = make(map[string]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf("%v Это каталог \n", f.Name())
		} else {
			fileContent, err := ioutil.ReadFile(filepath.Join(dir, f.Name()))
			if err != nil {
				return nil, err
			}
			envMap[f.Name()] = string(fileContent)
		}
	}
	//fmt.Printf("%v\n", envMap)

	return envMap, nil
}

func RunCmd(cmd string, paramCmd []string, env map[string]string) int {

	for param, _ := range env {
		os.Setenv(param, env[param])
	}

	c1 := exec.Command(cmd, paramCmd...)
	//c1.Env = append(os.Environ())
	//c1.Start()
	fmt.Println(c1.Env)
	out, _ := c1.Output()
	fmt.Printf("%v", string(out))
	//c1.Wait()
	return 0
}
