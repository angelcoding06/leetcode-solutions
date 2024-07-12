package solutions

func LongestCommonPrefix(strs []string) string {
	bandera := strs[0]
	for _, strg := range strs {
		i := 0
		for ; i < len(strg) && i < len(bandera) && bandera[i] == strg[i]; i++ {
		}
		bandera = bandera[:i]
	}
	return bandera
}
