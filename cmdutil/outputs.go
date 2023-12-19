package cmdutil

import (
	"io"
	"os"
)

func output(contain []byte) ([]byte, error) {
	for _, outputs := range Out {
		_, err := outputs.Write(contain)
		if err != nil {
			return []byte{}, err
		}
		outputs.Sync()
	}
	return contain, nil
}

func errput(contain []byte) ([]byte, error) {
	for _, outputs := range Err {
		_, err := outputs.Write(contain)
		if err != nil {
			return []byte{}, err
		}
		outputs.Sync()
	}
	return contain, nil
}

func outputsAfterRun() {
	defer resetBuffer()
	_, err := os.Stdout.Seek(0, io.SeekStart)
	if err != nil {
		println(err)
		// return err
	}
	contents, err := io.ReadAll(os.Stdout)
	if err != nil {
		println(err)
		// return err
	}
	_, err = output(contents)
	if err != nil {
		println(err)
		// return err
	}
	_, err = os.Stderr.Seek(0, io.SeekStart)
	if err != nil {
		println(err)
		// return err
	}
	contents, err = io.ReadAll(os.Stderr)
	if err != nil {
		println(err)
		// return err
	}
	os.Stderr.Close()
	_, err = errput(contents)
	if err != nil {
		println(err)
		// return err
	}
	// return nil
}
