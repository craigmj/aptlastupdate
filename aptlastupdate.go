package aptlastupdate

import (
	"os"
	"time"
)

// modTime returns the last modification date of a
// given file, or an error if any occurs
func modTime(filename string) (time.Time, error) {
	s, err := os.Stat(filename)
	if nil != err {
		return time.Time{}, err
	}
	return s.ModTime(), nil
}

/**
 * Last returns the last time that
 * `apt-get update` succeeded on the system.
 * The idea for this is from: http://askubuntu.com/questions/410247/how-to-know-last-time-apt-get-was-executed
 */
func Last() (time.Time, error) {
	t, err := modTime(`/var/lib/apt/periodic/update-success-stamp`)
	if nil == err {
		return t, nil
	}
	t, err = modTime(`/var/cache/apt/pkgcache.bin`)
	return t, err
}

/**
 * Within returns true if apt-get update succeeded
 * within the given time-frame from the current time.
 */
func Within(d time.Duration) (bool, error) {
	t, err := Last()
	if nil != err {
		return false, err
	}
	// Add the duration (eg 24h) to the last update time.
	// This would be the point at which the update has
	// 'expired'. So if Now() is Before() the expiration
	// time, we're within the expiration time.
	return time.Now().Before(t.Add(d)), nil
}
