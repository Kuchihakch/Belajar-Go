package main

import "fmt"

func main() {
	// kalau nggak dipakai variabel ini -> harusnya kan unused error, tapi kalau dengan const illegal file?
	/* const kuchi string = "Kuchi Trusted"
	const kazu = "Kazu Smeker" */
	const (
		kuchi string = "Kuchi Trusted"
		kazu  string = "Kazu Smeker"
	)
	// kazu = "Kazu Trusted" //Err, ilegal karena kazu emg skem (:v)
	fmt.Print(kuchi + "\n" + kazu)
}
