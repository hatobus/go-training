package arguments

import "golang.org/x/xerrors"

func GetIdentifier(args []string) (string, string, string, error) {
	if len(args) != 3 {
		return "", "", "", xerrors.Errorf("invalid arg length")
	}
	return args[0], args[1], args[2], nil
}
