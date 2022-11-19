package main

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	var result []string
	for i, c := range s {
		for _, p := range permutation(s[:i] + s[i+1:]) {
			result = append(result, string(c)+p)
		}
	}
	return result
}

func permutationIterative(s string) []string {
	var result []string
	result = append(result, "")
	for i := 0; i < len(s); i++ {
		var temp []string
		for _, r := range result {
			for j := 0; j <= len(r); j++ {
				temp = append(temp, r[:j]+string(s[i])+r[j:])
			}
		}
		result = temp
	}
	return result
}
