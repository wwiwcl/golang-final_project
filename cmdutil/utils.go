package cmdutil

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func resetBuffer(pipe ...bool) {
	os.Remove(InBufferFile.Name())
	os.Remove(ErrBufferFile.Name())
	InBufferFile, _ = os.CreateTemp("", ".inbuffer")
	if len(pipe) > 0 {
		err := pipelinePass()
		if err != nil {
			fmt.Fprintln(Stderr, err)
			// return err
		}
	}
	os.Remove(OutBufferFile.Name())
	OutBufferFile, _ = os.CreateTemp("", ".outbuffer")
	ErrBufferFile, _ = os.CreateTemp("", ".errbuffer")
	os.Stdin = InBufferFile
	os.Stdout = OutBufferFile
	os.Stderr = ErrBufferFile
	// return nil
}

func inSliceString(e []string, slice []string) int {
	i := 0
	for _, s := range slice {
		for _, v := range e {
			if v == s {
				return i
			}
		}
		i++
	}
	return -1
}

func keysOfStringMap(inputMap map[string]func(c *exec.Cmd, args ...string) error) []string {
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	return keys
}

func joinPath(pathA string, pathB string) string {
	return filepath.Join(pathA, pathB)
}

func isAbsPath(path string) bool {
	return filepath.IsAbs(path)
}

func pathExists(c *exec.Cmd, path string) bool {
	return NewCmd(c, exec.Command("ls", path)).Run() == nil
}

func makePath(c *exec.Cmd, path string) (string, error) {
	if isAbsPath(path) {
		return path, nil
	}
	root, err := Getcwd(c)
	if err != nil {
		return "", err
	}
	return joinPath(root, path), nil
}
