package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	return nil
}
