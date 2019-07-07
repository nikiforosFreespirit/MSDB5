// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package catalog

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p := messageKeyToIndex[key]
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"it": &dictionary{index: itIndex, data: itData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"(%d of %s)": 13,
	"(Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)":                               14,
	"(Turn of: %s, Companion is: %s, Played cards: %s, Auction score: %d, Phase: %d)": 17,
	"Action %s not valid":         3,
	"Callers":                     15,
	"Coin":                        9,
	"Cudgel":                      12,
	"Cup":                         10,
	"Enter name and connect":      18,
	"Error: %+v\n":                0,
	"Expecting player %s to play": 1,
	"Game: %+v":                   6,
	"Others":                      16,
	"Phase is not %d but %d":      2,
	"Side deck section: %s\n":     5,
	"Side deck: %s\n":             4,
	"Sword":                       11,
	"The end - %s team has all briscola cards": 7,
	"The end - Callers: %d; Others: %d":        8,
}

var enIndex = []uint32{ // 20 elements
	0x00000000, 0x00000013, 0x00000032, 0x0000004f,
	0x00000066, 0x0000007c, 0x0000009a, 0x000000a7,
	0x000000d3, 0x000000fb, 0x00000100, 0x00000104,
	0x0000010a, 0x00000111, 0x00000122, 0x00000160,
	0x00000168, 0x0000016f, 0x000001ce, 0x000001e5,
} // Size: 104 bytes

const enData string = "" + // Size: 485 bytes
	"\x04\x00\x01\x0a\x0e\x02Error: %+[1]v\x02Expecting player %[1]s to play" +
	"\x02Phase is not %[1]d but %[2]d\x02Action %[1]s not valid\x04\x00\x01" +
	"\x0a\x11\x02Side deck: %[1]s\x04\x00\x01\x0a\x19\x02Side deck section: %" +
	"[1]s\x02Game: %+[1]v\x02The end - %[1]s team has all briscola cards\x02T" +
	"he end - Callers: %[1]d; Others: %[2]d\x02Coin\x02Cup\x02Sword\x02Cudgel" +
	"\x02(%[1]d of %[2]s)\x02(Name: %[1]s, Cards: %+[2]v, Pile: %+[3]v, Has f" +
	"olded? %[4]t)\x02Callers\x02Others\x02(Turn of: %[1]s, Companion is: %[2" +
	"]s, Played cards: %[3]s, Auction score: %[4]d, Phase: %[5]d)\x02Enter na" +
	"me and connect"

var itIndex = []uint32{ // 20 elements
	0x00000000, 0x00000014, 0x00000037, 0x00000055,
	0x0000006d, 0x0000007f, 0x0000009d, 0x000000ba,
	0x000000e1, 0x0000010f, 0x00000113, 0x00000119,
	0x0000011f, 0x00000127, 0x00000138, 0x0000017c,
	0x00000186, 0x00000194, 0x000001f3, 0x00000211,
} // Size: 104 bytes

const itData string = "" + // Size: 529 bytes
	"\x04\x00\x01\x0a\x0f\x02Errore: %+[1]v\x02Mi aspetto che sia %[1]s a gio" +
	"care\x02La fase non e' %[1]d ma %[2]d\x02Azione %[1]s non valida\x04\x00" +
	"\x01\x0a\x0d\x02Monte: %[1]s\x04\x00\x01\x0a\x19\x02Sezione del monte: %" +
	"[1]s\x02Informazioni di gioco %+[1]v\x02Fine - I %[1]s hanno tutte le br" +
	"iscole\x02Fine - Chiamanti: %[1]d; Non-chiamanti: %[2]d\x02Oro\x02Coppe" +
	"\x02Spade\x02Bastoni\x02(%[1]d di %[2]s)\x02(Nome: %[1]s, Mano: %+[2]v, " +
	"Carte prese: %+[3]v, Ha passato? %[4]t)\x02Chiamanti\x02Non chiamanti" +
	"\x02(Turno di: %[1]s, Compagno: %[2]s, Carte giocate: %[3]s, Punteggio d" +
	"'asta: %[4]d, Fase: %[5]d)\x02Inserisci il nome e collegati"

	// Total table size 1222 bytes (1KiB); checksum: 979D27C7