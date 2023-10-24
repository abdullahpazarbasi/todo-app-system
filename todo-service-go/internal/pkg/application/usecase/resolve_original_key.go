package usecase

func resolveOriginalKey(keyForManipulation string) (originalKey string, removing bool) {
	if len(keyForManipulation) > 0 && keyForManipulation[0:1] == "-" {
		originalKey = keyForManipulation[1:]
		removing = true
	} else {
		originalKey = keyForManipulation
		removing = false
	}
	return
}
