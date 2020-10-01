package command

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.11/github"

	"github.com/google/uuid"
)

var uu uuid.UUID

func init() {
	var err error
	uu, err = uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
}

func newIssueDataFromFilePointer(fname string) map[string]string {
	// 別のfpで開かないと作成時点でのデータが入ってきてしまうためここで新たに開いておく
	fp, _ := os.Open(fname)
	defer fp.Close()

	d := make(map[string]string)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		elems := strings.SplitN(line, ": ", 2)
		key, val := strings.ToLower(elems[0]), elems[1]
		d[key] = val
	}
	return d
}

func startUpEditor(editor string, fname string) error {
	// stdin, stdout, stdout を cmd と紐づけるとエディタを開くことができる
	cmd := exec.Command(editor, fname)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func EditIssue(owner, repo, number string) error {
	editor := os.Getenv("ISSUE_EDITOR")
	if editor == "" {
		// vi ならどこでも入っているはずなのでデフォルトは vi にする
		editor = "vi"
	}

	tmpFile, err := ioutil.TempFile("", fmt.Sprintf("editor_issue_%v", uu.String()))
	if err != nil {
		return err
	}
	defer func() {
		err = tmpFile.Close()
		if err != nil {
			log.Println(err)
		}
		err = os.Remove(tmpFile.Name())
		if err != nil {
			log.Println(err)
		}
	}()

	issue, err := github.ReadIssueFromIdentifier(owner, repo, number)
	if err != nil {
		return err
	}

	currentIssue := []string{
		"Title: " + issue.Title,
		"Body: " + issue.Body,
	}

	for _, line := range currentIssue {
		if _, err := tmpFile.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err = startUpEditor(editor, tmpFile.Name()); err != nil {
		return err
	}

	tmpFile.Seek(0, 0)
	newData := newIssueDataFromFilePointer(tmpFile.Name())

	newissue, err := github.EditIssue(owner, repo, number, newData)
	if err != nil {
		return err
	}

	log.Println(newissue)

	return nil
}
