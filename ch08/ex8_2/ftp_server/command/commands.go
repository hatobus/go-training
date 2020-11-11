package command

const (
	CWD  = iota // change working directory
	DELE        // delete file or directory
	HELP        // show able to use commands
	LIST        // show file or directory
	PWD         // print working directory
	QUIT        // close FTP connection
	RETR        // Retrieve remote file
)
