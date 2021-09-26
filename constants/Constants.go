package constants

// Recombination recombine all constants and call ready function for each combination
func Recombination(constants *[]float64, ready func(constantsCombination *[]float64)) {
	ready(constants)
}
