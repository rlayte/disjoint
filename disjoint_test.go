package disjoint

import "testing"

func TestMakeSet(t *testing.T) {
	bill := MakeSet("Bill")

	if bill.data != "Bill" {
		t.Error("Set should contain Bill but contained", bill.data)
	}

	if bill.parent != nil {
		t.Error("Sets shouldn't have a parent when created")
	}
}

func TestMergeSets(t *testing.T) {
	bill := MakeSet("Bill")
	ted := MakeSet("Ted")
	rufus := MakeSet("Rufus")
	plato := MakeSet("Plato")

	bill.Union(ted)
	rufus.Union(ted)

	if bill.Find() != ted.Find() {
		t.Error("Single sets should belong to themselves")
	}

	if rufus.Find() != bill.Find() {
		t.Error("Single sets should belong to themselves")
	}

	socrates := MakeSet("Socrates")
	death := MakeSet("Socrates")
	bill.Union(socrates)
	bill.Union(death)

	if socrates.Find() != death.Find() {
		t.Error("Single sets should belong to themselves")
	}

	if plato.Find() != plato && plato.Find() == bill.Find() {
		t.Error("Single sets should belong to themselves")
	}
}
