package rapid

import "time"

const secondsInYear = 31_536_000
const nanoSecondsInSeconds = 1_000_000_000

// generates a time in the range from unix 0 to a year into the future.
// max uint is 18_446_744_073_709_551_615
func Time() *Generator[time.Time] {
	return Custom[time.Time](func(t *T) time.Time {
		now := time.Now().Unix()
		secGen := Int64Range(0, now+int64(secondsInYear))
		nsecGen := Int64Range(0, int64(nanoSecondsInSeconds))
		stamp := time.Unix(secGen.Draw(t, ""), int64(nsecGen.Draw(t, "")))

		//todo implement drawing a timezone.
		return stamp
	})
}
