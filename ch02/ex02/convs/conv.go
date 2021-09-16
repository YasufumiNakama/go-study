package convs

// Celsius <-> Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Feet <-> Metre
func FeetToMetre(f Feet) Metre { return Metre(f / 3.2808) }
func MetreToFeet(m Metre) Feet { return Feet(m * 3.2808) }

// Pound <-> Kilogram
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p / 2.2046) }
func KilogramToPound(k Kilogram) Pound { return Pound(k * 2.2046) }
