package command

const (
	CWD  = iota // change working directory
	RMD         // delete file
	XRMD        // delete file or directory
	HELP        // show able to use commands
	LIST        // show file or directory
	PWD         // print working directory
	QUIT        // close FTP connection
	RETR        // Retrieve remote file
	USER        // user login
	PASS        // pass command
	PASV        // pasv command
	ACCT        // acct command
	PORT        // port command
	SYST        // syst command
	LPRT        // LPRT command
	TYPE        // TYPE command
)

var CMD = map[string]int{
	"CWD":  CWD,
	"RMD":  RMD,
	"XRMD": XRMD,
	"HELP": HELP,
	"LIST": LIST,
	"PWD":  PWD,
	"QUIT": QUIT,
	"RETR": RETR,
	"USER": USER,
	"PASS": PASS,
	"PASV": PASV,
	"ACCT": ACCT,
	"PORT": PORT,
	"SYST": SYST,
	"LPRT": LPRT,
	"TYPE": TYPE,
}
