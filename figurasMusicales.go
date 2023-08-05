package tocadiscos

type FiguraMusical float64

// Constants for FiguraMusical using bit-shifting.
const (
	R  FiguraMusical = 1 << iota // Whole Note (Semibreve) - Represents 1.
	B                            // Half Note (Minim) - Represents 2.
	N                            // Quarter Note (Crotchet) - Represents 4.
	C                            // Eighth Note (Quaver) - Represents 8.
	SC                           // Sixteenth Note (Semiquaver) - Represents 16.
	F                            // Thirty-second Note (Demisemiquaver) - Represents 32.
	SF                           // Sixty-fourth Note (Hemidemisemiquaver) - Represents 64.
	G                            // 128th Note (Semihemidemisemiquaver) - Represents 128.
	SG                           // 256th Note (Demisemihemidemisemiquaver) - Represents 256.
)

// String returns the string representation of FiguraMusical.
func (f FiguraMusical) String() string {
	// Map to store string representations of FiguraMusical values.
	figuraMusicalStrings := map[FiguraMusical]string{
		R:  "Whole Note",
		B:  "Half Note",
		N:  "Quarter Note",
		C:  "Eighth Note",
		SC: "Sixteenth Note",
		F:  "Thirty-second Note",
		SF: "Sixty-fourth Note",
		G:  "128th Note",
		SG: "256th Note",
	}

	// Check if the FiguraMusical value exists in the map, return its string representation.
	if str, ok := figuraMusicalStrings[f]; ok {
		return str
	}

	// If the FiguraMusical value is not found in the map, return "Unknown".
	return "Unknown"
}
