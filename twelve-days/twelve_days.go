package twelve

import "fmt"

const testVersion = 1

// Song returns all 12 verses with one verse
// per line
func Song() (song string) {
	for v := 1; v <= 12; v++ {
		song += fmt.Sprintf("%s\n", Verse(v))
	}
	return song
}

// Verse returns a single verse from the song
func Verse(v int) string {

	days := map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}

	gifts := map[int]string{
		1:  "a Partridge in a Pear Tree",
		2:  "two Turtle Doves",
		3:  "three French Hens",
		4:  "four Calling Birds",
		5:  "five Gold Rings",
		6:  "six Geese-a-Laying",
		7:  "seven Swans-a-Swimming",
		8:  "eight Maids-a-Milking",
		9:  "nine Ladies Dancing",
		10: "ten Lords-a-Leaping",
		11: "eleven Pipers Piping",
		12: "twelve Drummers Drumming",
	}

	gift := ""
	for d := v; d >= 1; d-- {

		if v != 1 && d == 1 {
			gift = fmt.Sprintf("%s, and %s", gift, gifts[d])
		} else {
			gift = fmt.Sprintf("%s, %s", gift, gifts[d])
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me%s.", days[v], gift)
}
