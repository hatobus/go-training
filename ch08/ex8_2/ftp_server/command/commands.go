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
	SYST        // syst command
	LPRT        // LPRT command
)

var CMD = map[string]int{
	"CWD":  CWD,
	"DELE": DELE,
	"HELP": HELP,
	"LIST": LIST,
	"PWD":  PWD,
	"QUIT": QUIT,
	"RETR": RETR,
	"USER": USER,
	"PASS": PASS,
	"ACCT": ACCT,
	"PORT": PORT,
	"SYST": SYST,
	"LPRT": LPRT,
}
