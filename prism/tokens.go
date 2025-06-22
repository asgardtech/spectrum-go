package prism

import (
	"fmt"
	"strings"
)

// SizeToken represents a Spectrum spatial token ("size-50", "size-100", …).
// Using a dedicated type instead of raw strings gives you auto-complete and compile-time checks.
// Only the most common values are listed here – extend as needed.
// See https://spectrum.adobe.com/page/size-tokens/ for the full list.
type SizeToken string

const (
	Size0    SizeToken = "size-0"
	Size25   SizeToken = "size-25"
	Size50   SizeToken = "size-50"
	Size75   SizeToken = "size-75"
	Size100  SizeToken = "size-100"
	Size200  SizeToken = "size-200"
	Size300  SizeToken = "size-300"
	Size400  SizeToken = "size-400"
	Size500  SizeToken = "size-500"
	Size600  SizeToken = "size-600"
	Size800  SizeToken = "size-800"
	Size900  SizeToken = "size-900"
	Size1000 SizeToken = "size-1000"
)

func (s SizeToken) String() string { return string(s) }

// Size is a convenience constructor that turns an integer into the corresponding
// size token (e.g. Size(200) == "size-200"). Use it when you want the
// flexibility of arbitrary token values instead of pre-declared constants.
func Size(n int) SizeToken { return SizeToken(fmt.Sprintf("size-%d", n)) }

// Fr units (fractional columns/rows in CSS Grid).
// We wrap it in its own type so that callers don't have to remember to write "1fr".
type FrUnit string

func Fr(n int) FrUnit { return FrUnit(fmt.Sprintf("%dfr", n)) }

func (f FrUnit) String() string { return string(f) }

// GridTemplate is a helper for the grid `Columns` method – you build it with
// `GridTemplate{Fr(1), Fr(1)}` instead of raw "1fr 1fr".
// It implements fmt.Stringer so underlying Spectrum component still receives a string.

type GridTemplate []FrUnit

func (g GridTemplate) String() string {
	parts := make([]string, len(g))
	for i, p := range g {
		parts[i] = p.String()
	}
	return strings.Join(parts, " ")
}

// ---------------------------------------------
// Icon names (subset). Keeping them in one place
// prevents typos and surfaces auto-completion.
// ---------------------------------------------

type Icon string

const (
	IconEdit   Icon = "edit"
	IconDelete Icon = "delete"
	IconPlus   Icon = "add"
	IconTable  Icon = "table"
)

func (i Icon) String() string { return string(i) }

// ---------------------------------------------
// Intent tokens – map to Spectrum semantic colors
// ---------------------------------------------

type Intent string

const (
	IntentNeutral  Intent = "neutral"
	IntentPositive Intent = "positive"
	IntentWarning  Intent = "warning"
	IntentDanger   Intent = "danger"
)

func (in Intent) String() string { return string(in) }
