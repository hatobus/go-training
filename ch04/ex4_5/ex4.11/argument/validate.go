package arguments

func ValidateArgsRunning(args []string) bool {
	if len(args) < 2 {
		return false
	}
	return true
}

func ValidateSearchArguments(args []string) bool {
	if len(args) < 1 {
		return false
	}
	return true
}
