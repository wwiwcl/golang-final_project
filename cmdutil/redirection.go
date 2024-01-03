package cmdutil

import (
	"fmt"
	"os"
	"os/exec"
)

func checkRedirection(c *exec.Cmd, mode int, args *[]string) (bool, []*os.File, error) {
	var returnFiles []*os.File
	if mode == 0 {
		rediretInFile := inSliceString([]string{"<"}, *args)
		for rediretInFile >= 0 {
			if len(*args) == rediretInFile+1 {
				println(fmt.Errorf("redirection: missing redirection target for stdin"))
				return false, []*os.File{}, nil
			}
			filePath, err := makePath(c, (*args)[rediretInFile+1])
			if err != nil {
				return false, []*os.File{}, err
			}
			fileout, err := os.Open(filePath)
			if err != nil {
				return false, []*os.File{}, err
			}
			*args = append((*args)[:rediretInFile], (*args)[rediretInFile+2:]...)
			returnFiles = append(returnFiles, fileout)
			rediretInFile = inSliceString([]string{"<"}, *args)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	if mode == 1 {
		rediretOutFile := inSliceString([]string{">", "1>"}, *args)
		for rediretOutFile >= 0 {
			if len(*args) == rediretOutFile+1 {
				println(fmt.Errorf("redirection: missing redirection target for stdout"))
				return false, []*os.File{}, nil
			}
			var fileout *os.File
			if (*args)[rediretOutFile+1] == "&1" {
				fileout = Stdout
			} else if (*args)[rediretOutFile+1] == "&2" {
				fileout = Stderr
			} else {
				filePath, err := makePath(c, (*args)[rediretOutFile+1])
				if err != nil {
					return false, []*os.File{}, err
				}
				fileout, err = os.Create(filePath)
				if err != nil {
					return false, []*os.File{}, err
				}
			}
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+2:]...)
			returnFiles = append(returnFiles, fileout)
			rediretOutFile = inSliceString([]string{">", "1>"}, *args)
		}
		rediretOutFile = inSliceString([]string{"1>&2"}, *args)
		if rediretOutFile >= 0 {
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+1:]...)
			returnFiles = append(returnFiles, Stderr)
		}
		rediretOutFile = inSliceString([]string{">&1", "1>&1"}, *args)
		if rediretOutFile >= 0 {
			*args = append((*args)[:rediretOutFile], (*args)[rediretOutFile+1:]...)
			returnFiles = append(returnFiles, Stdout)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	if mode == 2 {
		rediretErrFile := inSliceString([]string{"2>", ">>"}, *args)
		for rediretErrFile >= 0 {
			if len(*args) == rediretErrFile+1 {
				println(fmt.Errorf("redirection: missing redirection target for stderr"))
				return false, []*os.File{}, nil
			}
			var fileout *os.File
			if (*args)[rediretErrFile+1] == "&1" {
				fileout = Stdout
			} else if (*args)[rediretErrFile+1] == "&2" {
				fileout = Stderr
			} else {
				filePath, err := makePath(c, (*args)[rediretErrFile+1])
				if err != nil {
					return false, []*os.File{}, err
				}
				fileout, err = os.Create(filePath)
				if err != nil {
					return false, []*os.File{}, err
				}
			}
			*args = append((*args)[:rediretErrFile], (*args)[rediretErrFile+2:]...)
			returnFiles = append(returnFiles, fileout)
			rediretErrFile = inSliceString([]string{"2>", ">>"}, *args)
		}
		rediretErrFile = inSliceString([]string{"2>&1"}, *args)
		if rediretErrFile >= 0 {
			*args = append((*args)[:rediretErrFile], (*args)[rediretErrFile+1:]...)
			returnFiles = append(returnFiles, Stdout)
		}
		rediretErrFile = inSliceString([]string{"2>&2", ">>&2"}, *args)
		if rediretErrFile >= 0 {
			*args = append((*args)[:rediretErrFile], (*args)[rediretErrFile+1:]...)
			returnFiles = append(returnFiles, Stderr)
		}
		return (len(returnFiles) > 0), returnFiles, nil
	}
	return false, []*os.File{}, nil
}
