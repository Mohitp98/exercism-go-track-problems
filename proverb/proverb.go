// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	collection := []string{}

	switch {
	case len(rhyme) == 0:
		return collection
	case len(rhyme) > 0:
		for i := 0; i < len(rhyme)-1; i++ {
			collection = append(collection, returnString(rhyme[i], rhyme[i+1]))
		}
		collection = append(collection, "And all for the want of a "+rhyme[0]+".")
	}
	return collection
}

// returnString: function to create a sentence from input
func returnString(str1, str2 string) string {
	return "For want of a " + str1 + " the " + str2 + " was lost."
}
