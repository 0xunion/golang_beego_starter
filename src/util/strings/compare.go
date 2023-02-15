package strings

import "github.com/antlabs/strsim"

/*
	if there is a long text we want to caculate the similarity
	the normal way is to use the Levenshtein distance, but it's too slow
	so the best choice is to use simhash algorithm
*/
func CompareSimilarityForComplex(a, b string) float64 {
	return strsim.Compare(a, b, strsim.Simhash(), strsim.IgnoreSpace(), strsim.IgnoreCase())
}

/*
	if there is a short text we want to caculate the similarity
	the best choice is to use Levenshtein distance
*/

func CompareSimilarityForSimple(a, b string) float64 {
	return strsim.Compare(a, b, strsim.IgnoreCase(), strsim.IgnoreSpace())
}

func CompareSimilarity(a, b string, filter ...func(string, string) bool) float64 {
	for _, f := range filter {
		if !f(a, b) {
			return 0
		}
	}
	if len(a) > 500 || len(b) > 500 {
		return CompareSimilarityForComplex(a, b)
	}
	return CompareSimilarityForSimple(a, b)
}
