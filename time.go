package rapid

import "time"

const secondsInYear = 31_536_000
const nanoSecondsInSeconds = 1_000_000_000

// generates a time in the range from unix 0 to a year into the future.
func Time() *Generator[time.Time] {
	return Custom(func(t *T) time.Time {
		now := time.Now().Unix()
		secGen := Int64Range(0, now+int64(secondsInYear))
		nsecGen := Int64Range(0, int64(nanoSecondsInSeconds))
		stamp := time.Unix(secGen.Draw(t, ""), int64(nsecGen.Draw(t, "")))

		//todo implement drawing a timezone.
		return stamp
	})
}

func TimeTo(to time.Time) *Generator[time.Time] {
	return Custom(func(t *T) time.Time {
		secGen := Int64Range(0, to.Unix())
		nsecGen := Int64Range(0, int64(1000))
		stamp := time.Unix(secGen.Draw(t, ""), int64(nsecGen.Draw(t, "")))

		//todo implement drawing a timezone.
		return stamp
	})
}

func TimeInterval(from, to time.Time) *Generator[time.Time] {
	return Custom(func(t *T) time.Time {
		secGen := Int64Range(from.Unix(), to.Unix())
		nsecGen := Int64Range(0, int64(1000))
		stamp := time.Unix(secGen.Draw(t, ""), int64(nsecGen.Draw(t, "")))

		//todo implement drawing a timezone.
		return stamp
	})
}
