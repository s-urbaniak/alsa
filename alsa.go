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

const (
	SND_CTL_NONBLOCK = C.SND_CTL_NONBLOCK
	SND_CTL_ASYNC    = C.SND_CTL_ASYNC
)

type Ctl C.snd_ctl_t

type Failure interface {
	Code() int
}

type err struct {
	reason string
	code   int
}

func (e *err) Error() string {
	return fmt.Sprintf("%v, code %d", e.reason, e.code)
}

func (e *err) Code() int {
	return e.code
}

func CtlOpen(name string, mode int) (*Ctl, error) {
	var ctl *C.snd_ctl_t

	if ret := C.snd_ctl_open(
		&ctl,
		C.CString(name),
		C.int(mode),
	); ret < 0 {
		return nil, &err{"control open error", int(ret)}
	}

	return (*Ctl)(ctl), nil
}

func CtlClose(ctl *Ctl) error {
	if ret := C.snd_ctl_close((*C.snd_ctl_t)(ctl)); ret < 0 {
		return &err{"control close error", int(ret)}
	}

	return nil
}

func CardNext(card int) (int, error) {
	c := C.int(card)

	if ret := C.snd_card_next(&c); ret < 0 {
		return -1, &err{"card next error", int(ret)}
	}

	return int(c), nil
}

func PcmStreamName(stream int) string {
	return C.GoString(C.snd_pcm_stream_name(C.snd_pcm_stream_t(stream)))
}

func StrError(code int) string {
	return C.GoString(C.snd_strerror(C.int(code)))
}
