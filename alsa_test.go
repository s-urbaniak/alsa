package alsa_test

import (
	"testing"

	"github.com/s-urbaniak/alsa"
)

func TestCardNext(t *testing.T) {
	var err error
	card := -1

	for {
		card, err = alsa.CardNext(card)
		if err != nil {
			t.Error(err)
		}

		if card < 0 {
			break
		}
	}

	t.Log(alsa.PcmStreamName(alsa.SND_PCM_STREAM_CAPTURE))
}
