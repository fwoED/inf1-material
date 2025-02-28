package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	positionfirst := 100
	positionlast := 100
	for pos, val := range list {
		if val == first {
			positionfirst = pos
		}
		if val == last {
			positionlast = pos
		}
	}
	if positionfirst == 100 || positionfirst > positionlast {
		return []string{}
	} else {
		return append(list[:positionfirst], list[1+positionlast:]...)
	}
}
