package protein

import (
	"fmt"
	"unicode/utf8"
)

const testVersion = 1

type item struct {
	typ itemType
	val string
}

type itemType int

const (
	itemCodon itemType = iota
	itemStop
	itemError
	itemEOF
)

const eof rune = -1

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func (i item) String() string {
	switch i.typ {
	case itemError:
		return i.val
	case itemEOF:
		return "eof"
	case itemCodon:
		return fmt.Sprintf("codon (%s)", i.val)
	case itemStop:
		return fmt.Sprintf("stop (%s)", i.val)
	default:
		return ""
	}
}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*lexer) stateFn

// run lexes the input by executing state functions
// until the state is nil.
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

// lexer holds the state of the scanner.
type lexer struct {
	name  string    // only used for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read.
	items chan item // channel of scanned items.
}

func lex(name, input string) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l, l.items
}

// emit passes tokens back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func lexText(l *lexer) stateFn {
	for {
		if l.pos >= l.start+3 {
			c := FromCodon(l.input[l.start:l.pos])
			if c == "STOP" {
				l.emit(itemEOF)
				return nil
			}

			l.emit(itemCodon)
			return lexText
		}

		if l.next() == eof {
			break
		}

	}

	if l.pos > l.start {
		l.emit(itemCodon)
	}
	l.emit(itemEOF)
	return nil
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeLastRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func FromCodon(codon string) (p string) {
	p, _ = codons[codon]
	return p
}

func FromRNA(rna string) (p []string) {
	_, ch := lex("demo", rna)

	for {
		select {
		case s := <-ch:
			fmt.Printf("%s\n", s)
			p = append(p, codons[s.val])
			if s.typ == itemEOF {
				return p
			}
		}
	}

	return p
}
