package helper

func BuildKey(prefix string, key string) string {
	return prefix + ":" + key
}
