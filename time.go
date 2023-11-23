package rapid

import "time"

// generates a time in the range from unix mili 0 to unix mili double of now.
// from unix returns a timestamp starting on jan 1, 1970, so thats our starting point.
// max uint is 18_446_744_073_709_551_615
func Time() *Generator[time.Time] {
	return Custom[time.Time](func(t *T) time.Time {
		now := time.Now().Unix()
		secGen := Int64Range(0, now*2)
		nsecGen := Int64Range(0, 1_000_000_000)
		stamp := time.Unix(secGen.Draw(t, ""), int64(nsecGen.Draw(t, "")))

		return stamp

	})
}
