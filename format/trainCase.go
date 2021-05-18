package format

// TrainCase convert to train case
func TrainCase(t string) string {
	return KebabCase(t, true, false)
}
