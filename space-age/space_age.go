// Package space provides a function for computing ages on various plannets in the solar
// system. For example, if you were 1,000,000,000 seconds old, this would represent
// 31.69 Earth-years old.
//
// Results are rounded to the nearest 2 decimal places.
//
// Pluto has been excluded: http://www.youtube.com/watch?v=Z_2gbGXzFbs
package space

// Planet represents a planet in the solar system
type Planet string

const (
	//yearInSeconds is the number of seconds in a year on Earth
	yearInSeconds float64 = 31557600.0
)

var ages = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age takes a number of seconds and a planet and calculates an age in years
func Age(s float64, p Planet) float64 {
	age := (s / ages[p]) / yearInSeconds
	return round(age, 0.01)
}

// round takes a floating point number and rounds to a given precision
// source: https://stackoverflow.com/a/39544897
func round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}
