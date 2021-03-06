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
	"(Undefined card)":           2,
	"Auction":                    10,
	"Auction score":              20,
	"Callers":                    21,
	"Card":                       13,
	"Cards":                      27,
	"Coin":                       3,
	"Companion":                  12,
	"Cudgel":                     6,
	"Cup":                        4,
	"End":                        14,
	"Enter name and connect":     0,
	"Error":                      7,
	"Exchange":                   11,
	"Expecting %s, but found %s": 8,
	"Game":                       15,
	"Has folded":                 29,
	"Join":                       9,
	"Last card":                  19,
	"Name":                       26,
	"No":                         31,
	"Others":                     22,
	"Phase":                      17,
	"Pile":                       28,
	"Played cards":               18,
	"Player":                     25,
	"Side deck":                  1,
	"Sword":                      5,
	"The end":                    23,
	"Turn of":                    16,
	"Yes":                        30,
	"have all briscola cards":    24,
}

var enIndex = []uint32{ // 33 elements
	// Entry 0 - 1F
	0x00000000, 0x00000017, 0x00000021, 0x00000032,
	0x00000037, 0x0000003b, 0x00000041, 0x00000048,
	0x0000004e, 0x0000006f, 0x00000074, 0x0000007c,
	0x00000085, 0x0000008f, 0x00000094, 0x00000098,
	0x0000009d, 0x000000a5, 0x000000ab, 0x000000b8,
	0x000000c2, 0x000000d0, 0x000000d8, 0x000000df,
	0x000000e7, 0x000000ff, 0x00000106, 0x0000010b,
	0x00000111, 0x00000116, 0x00000121, 0x00000125,
	// Entry 20 - 3F
	0x00000128,
} // Size: 156 bytes

const enData string = "" + // Size: 296 bytes
	"\x02Enter name and connect\x02Side deck\x02(Undefined card)\x02Coin\x02C" +
	"up\x02Sword\x02Cudgel\x02Error\x02Expecting %[1]s, but found %[2]s\x02Jo" +
	"in\x02Auction\x02Exchange\x02Companion\x02Card\x02End\x02Game\x02Turn of" +
	"\x02Phase\x02Played cards\x02Last card\x02Auction score\x02Callers\x02Ot" +
	"hers\x02The end\x02have all briscola cards\x02Player\x02Name\x02Cards" +
	"\x02Pile\x02Has folded\x02Yes\x02No"

var itIndex = []uint32{ // 33 elements
	// Entry 0 - 1F
	0x00000000, 0x0000001e, 0x00000024, 0x00000034,
	0x00000038, 0x0000003e, 0x00000044, 0x0000004c,
	0x00000053, 0x00000079, 0x00000087, 0x0000008c,
	0x0000009d, 0x000000ad, 0x000000b9, 0x000000be,
	0x000000c4, 0x000000cd, 0x000000d2, 0x000000e0,
	0x000000ed, 0x000000fe, 0x00000108, 0x00000116,
	0x0000011b, 0x00000133, 0x0000013d, 0x00000142,
	0x00000148, 0x00000154, 0x0000015f, 0x00000163,
	// Entry 20 - 3F
	0x00000166,
} // Size: 156 bytes

const itData string = "" + // Size: 358 bytes
	"\x02Inserisci il nome e collegati\x02Monte\x02(Carta assente)\x02Oro\x02" +
	"Coppe\x02Spade\x02Bastoni\x02Errore\x02Mi aspetto %[1]s, ma ho trovato %" +
	"[2]s\x02Registrazione\x02Asta\x02Cambio con monte\x02Scelta compagno\x02" +
	"Gioco carte\x02Fine\x02Gioco\x02Turno di\x02Fase\x02Carte giocate\x02Ult" +
	"ima carta\x02Punteggio d'asta\x02Chiamanti\x02Non chiamanti\x02Fine\x02h" +
	"anno tutte le briscole\x02Giocatore\x02Nome\x02Carte\x02Carte prese\x02H" +
	"a foldato\x02Yes\x02No"

	// Total table size 966 bytes (0KiB); checksum: C3C0D096
