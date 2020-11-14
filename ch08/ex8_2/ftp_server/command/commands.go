package command

const (
	CWD  = iota // change working directory
	DELE        // delete file or directory
	HELP        // show able to use commands
	LIST        // show file or directory
	PWD         // print working directory
	QUIT        // close FTP connection
	RETR        // Retrieve remote file
	USER        // user login
	PASS        // pass command
	ACCT        // acct command
	PORT        // port command
)

var CMD = map[string]int{
	"cd":   CWD,
	"rm":   DELE,
	"help": HELP,
	"ls":   LIST,
	"pwd":  PWD,
	"exit": QUIT,
	"cp":   RETR,
	"USER": USER,
	"PASS": PASS,
	"ACCT": ACCT,
	"PORT": PORT,
}
