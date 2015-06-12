# aptlastupdate
Simple golang package to return info on when apt-get update last succeeded.

Two functions are supported:

`Last() (time.Time,error)` returns the time.Time when apt-get update succeeded.

`Within(time.Duration) (bool,error)` returns TRUE if the last apt-get update succeeded within the last time.Duration.
