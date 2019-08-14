package utils

func CopyStringArray(s []string) []string {
	ns := make([]string, len(s))
	copy(ns, s)
	return ns
}
