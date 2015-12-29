package alsa_test

import (
	"fmt"
	"testing"

	"github.com/s-urbaniak/alsa"
)

func TestCardNext(t *testing.T) {
	card, err := alsa.CardNext(-1)

	t.Log("list of", alsa.PcmStreamName(alsa.SND_PCM_STREAM_CAPTURE), "hardware devices")

	for card >= 0 {
		if err != nil {
			break
		}

		name := fmt.Sprintf("hw:%d", card)
		t.Log(name)

		ctl, err := alsa.CtlOpen(name, 0)
		if err != nil {
			break
		}

		if err = alsa.CtlClose(ctl); err != nil {
			break
		}
		card, err = alsa.CardNext(card)
	}

	if err != nil {
		t.Error(err)
	}

	if _, err = alsa.CtlOpen("unknown", 0); err == nil {
		t.Error("expected error, but got none")
	}

	t.Logf(
		"expected CtlOpen err %q alsa strerror %q",
		err,
		alsa.StrError(err.(alsa.Failure).Code()),
	)
}
