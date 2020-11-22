package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func (pi *interpreter) list(dst string) (int, error) {
	conn, err := pi.dataConnection()
	if err != nil {
		return StatusBadCommand, err
	}
	defer conn.Close()

	fi := new(fileInfo)

	if fi, err = pi.checkPathExist(dst); os.IsNotExist(err) {
		pi.printf("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	} else if err != nil {
		log.Println(err)
		pi.printf("%v: server error ", dst)
		return StatusFileUnavailable, nil
	}

	if fi.info.IsDir() {
		files, err := ioutil.ReadDir(fi.path)
		if err != nil {
			pi.printf("%v: No such file or directory ", dst)
			return StatusBadArguments, nil
		}

		pi.printf("%v\r\n", StatusAboutToSend)

		for _, f := range files {
			_, err := fmt.Fprint(conn, f.Name(), "\n")
			if err != nil {
				return StatusBadCommand, err
			}
		}
	} else {
		pi.printf("%v\r\n", StatusAboutToSend)
		_, err := fmt.Fprint(conn, fi.info.Name(), "\n")
		if err != nil {
			return StatusBadCommand, err
		}
	}

	return StatusClosingDataConnection, nil
}

func (pi *interpreter) delete(dst string) (int, error) {
	var err error
	fi := new(fileInfo)
	if fi, err = pi.checkPathExist(dst); os.IsNotExist(err) {
		pi.printf("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	} else if err != nil {
		log.Println(err)
		pi.printf("%v: server error ", dst)
		return StatusFileUnavailable, nil
	}

	if fi.info.IsDir() {
		if err := os.RemoveAll(fi.path); err != nil {
			return StatusActionAborted, nil
		}
	} else {
		if err := os.Remove(fi.path); err != nil {
			return StatusActionAborted, nil
		}
	}

	return StatusCommandOK, nil
}

func (pi *interpreter) retr(dst string) (int, error) {
	var err error
	fi := new(fileInfo)
	if fi, err = pi.checkPathExist(dst); os.IsNotExist(err) {
		pi.printf("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	} else if err != nil {
		log.Println(err)
		pi.printf("%v: server error ", dst)
		return StatusFileUnavailable, nil
	}

	var retrFile string

	if fi.info.IsDir() {
		pi.printf("%v is directory, please input files ", fi.path)
		return StatusFileUnavailable, nil
	} else {
		retrFile = fi.path
	}

	pi.printf("150 sending. \r\n")

	conn, err := pi.dataConnection()
	if err != nil {
		pi.printf("Data connection open failed ")
		return StatusFileUnavailable, nil
	}

	err = pi.send(conn, retrFile)
	if err != nil {
		log.Println(err)
		return StatusFileUnavailable, err
	}
	conn.Close()

	return StatusClosingDataConnection, nil
}

func (pi *interpreter) send(conn io.ReadWriteCloser, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if pi.binaryOption {
		_, err = io.Copy(conn, f)
		if err != nil {
			log.Println(err)
			pi.printf("file unavailable ")
			return fmt.Errorf("io copy failed")
		}
	} else {
		r := bufio.NewReader(f)
		w := bufio.NewWriter(conn)

		for {
			l, isPrefix, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println(err)
				pi.printf("file unavailable ")
				return fmt.Errorf("read line failed")
			}

			w.Write(l)
			if !isPrefix {
				w.Write([]byte("\r\n"))
			}
		}
	}
	return nil
}
