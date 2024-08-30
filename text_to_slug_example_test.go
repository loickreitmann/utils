package utils_test

import (
	"fmt"

	"github.com/loickreitmann/utils"
)

func ExampleUtils_TextToSlug() {
	var u utils.Utils
	quotes := []string{
		"Il est grand temps de rallumer les étoiles.",             // — Guillaume Apollinaire
		"Cliché, but love conquers all.",                          // — Common English phrase
		"La vida es sueño, y los sueños, sueños son.",             // — Pedro Calderón de la Barca
		"Über den Wolken muss die Freiheit wohl grenzenlos sein.", // (Above the clouds, freedom must be boundless.) — Reinhard Mey
		"Człowiek bez marzeń jest jak ptak bez skrzydeł.",         // (A person without dreams is like a bird without wings.) — Polish proverb
	}

	for _, quote := range quotes {
		fmt.Println(u.TextToSlug(quote))
	}

	// Output:
	// il-est-grand-temps-de-rallumer-les-etoiles
	// cliche-but-love-conquers-all
	// la-vida-es-sueno-y-los-suenos-suenos-son
	// uber-den-wolken-muss-die-freiheit-wohl-grenzenlos-sein
	// cz-owiek-bez-marzen-jest-jak-ptak-bez-skrzyde
}
