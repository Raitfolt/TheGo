package main

import "fmt"

//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Herbs
	Acetaminophen = Paracetamol
)

func (p Pill) String2() string {
	switch p {
	case Placebo:
		return "Placebo"
	case Aspirin:
		return "Aspirin"
	case Ibuprofen:
		return "Ibuprofen"
	case Paracetamol: // == Acetaminophen
		return "Paracetamol"
	}
	return fmt.Sprintf("Pill(%d)", p)
}
