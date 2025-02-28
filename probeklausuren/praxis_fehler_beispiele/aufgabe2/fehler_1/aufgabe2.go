package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {

	ergebnis := []int{}
	for _, s := range list {
		if s > m && s < n {
			ergebnis = append(ergebnis, s)
		}

	}
	if len(ergebnis) == 0 || m > n {
		return []int{}
	}

	return ergebnis
}
