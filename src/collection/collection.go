package collection

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](origin []T, callback func(item T, index int) R) []R {
	result := make([]R, len(origin))

	for i, item := range origin {
		result[i] = callback(item, i)
	}

	return result
}

func ForEach[T any](origin []T, callback func(item T, idx int) error) error {
	for idx, c := range origin {
		if err := callback(c, idx); err != nil {
			return err
		}
	}
	return nil
}

func FirstWhen[T any](origin []T, callback func(item T) bool) *T {
	for _, c := range origin {
		if value := callback(c); value {
			return &c
		}
	}
	return nil
}
