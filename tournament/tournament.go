// Package tournament summarises the results of a small football competition.
package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"sort"
)

type scoreCard struct {
	win  int
	draw int
	loss int
}

func (sc scoreCard) points() int {
	return (3 * sc.win) + sc.draw
}

func (sc scoreCard) played() int {
	return sc.win + sc.draw + sc.loss
}

// Tally reads a colon separated list of match scores and outputs
// a table summarising the overall tournament position.
func Tally(in io.Reader, out io.Writer) error {

	r := csv.NewReader(in)
	r.Comma = ';'
	r.Comment = '#'

	matches, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	sb := map[string]scoreCard{}

	// Read the scores of each match and build a scoreboard
	for l, m := range matches {
		if len(m) == 0 {
			continue
		}
		if len(m) != 3 {
			return fmt.Errorf("invalid score format on line: %d", l)
		}

		switch m[2] {
		case "win":
			wt := sb[m[0]]
			wt.win++
			sb[m[0]] = wt

			lt := sb[m[1]]
			lt.loss++
			sb[m[1]] = lt

		case "loss":
			lt := sb[m[0]]
			lt.loss++
			sb[m[0]] = lt

			wt := sb[m[1]]
			wt.win++
			sb[m[1]] = wt

		case "draw":
			t1 := sb[m[0]]
			t1.draw++
			sb[m[0]] = t1

			t2 := sb[m[1]]
			t2.draw++
			sb[m[1]] = t2
		default:
			return fmt.Errorf("unrecognised match outcome")
		}
	}

	// We can't sort maps so we create an slice of teams.
	ts := make([]string, 0, len(sb))
	for t := range sb {
		ts = append(ts, t)
	}

	sort.Slice(ts, func(i, j int) bool {

		// Sort numerically based on points if possible.
		if sb[ts[i]].points() > sb[ts[j]].points() {
			return true
		}

		// If equal on points, then sort alphabetically.
		if sb[ts[i]].points() == sb[ts[j]].points() {
			return ts[i] < ts[j]
		}

		return false
	})

	// Print the scores to the output
	fmt.Fprintln(out, "Team                           | MP |  W |  D |  L |  P")
	for _, t := range ts {
		fmt.Fprintf(out, "%-31s| %2d | %2d | %2d | %2d | %2d\n", t, sb[t].played(), sb[t].win, sb[t].draw, sb[t].loss, sb[t].points())
	}

	return nil
}
