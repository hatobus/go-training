package argument

import "os"

func ValidateIndexArguments(s []string) bool {
	return len(os.Args) == 3
}
