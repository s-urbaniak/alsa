package alsa_test

import (
	"fmt"
	"testing"

	"github.com/s-urbaniak/alsa"
)

func TestCardNext(t *testing.T) {
	t.Log("list of", alsa.PcmStreamName(alsa.SND_PCM_STREAM_CAPTURE), "hardware devices")

	card, err := alsa.CardNext(-1)
	if err != nil {
		t.Error(err)
	}

	for card >= 0 {
		name := fmt.Sprintf("hw:%d", card)
		t.Log(name)

		ctl, err := alsa.CtlOpen(name, 0)
		if err != nil {
			t.Error(err)
		}

		info, err := alsa.NewCtlCardInfo(ctl)
		if err != nil {
			t.Error(err)
		}
		t.Logf("info %p\n", info)

		dev, err := alsa.CtlPcmNextDevice(ctl, -1)
		if err != nil {
			t.Error(err)
		}

		for dev >= 0 {
			dev, err = alsa.CtlPcmNextDevice(ctl, dev)
			if err != nil {
				t.Error(err)
			}
		}

		if err = alsa.CtlClose(ctl); err != nil {
			t.Error(err)
		}

		card, err = alsa.CardNext(card)
		if err != nil {
			t.Error(err)
		}
	}
}
