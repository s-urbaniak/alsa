package alsa

/*
#include <alsa/asoundlib.h>
#cgo pkg-config: alsa
*/
import "C"
import "fmt"

const (
	SND_PCM_STREAM_CAPTURE  = C.SND_PCM_STREAM_CAPTURE
	SND_PCM_STREAM_PLAYBACK = C.SND_PCM_STREAM_PLAYBACK
)

func CardNext(card int) (int, error) {
	c := C.int(card)

	if ret := int(C.snd_card_next(&c)); ret < 0 {
		return -1, fmt.Errorf("error code %d", ret)
	}

	return int(c), nil
}

func PcmStreamName(stream int) string {
	return C.GoString(C.snd_pcm_stream_name(C.snd_pcm_stream_t(stream)))
}
