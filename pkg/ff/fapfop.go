package ff

func IgnoreError[V any](value V, err error) V {
	return value
}
