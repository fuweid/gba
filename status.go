package gba

import "github.com/fuweid/echo"

func Status() string {
	return echo.Version()
}
