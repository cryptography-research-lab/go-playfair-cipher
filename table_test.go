package playfair_cipher

import (
	"testing"
)

func TestMakeTable(t *testing.T) {
	options := &Options{
		BlackCharacter: 'Z',
	}
	table := MakeTable("play-fair", options)
	t.Log(table)
}
