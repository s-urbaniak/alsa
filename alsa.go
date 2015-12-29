package alsa

/*
#include <alsa/asoundlib.h>
#include "alsa.h"
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

type CtlCardInfo C.snd_ctl_card_info_t

func NewCtlCardInfo(ctl *Ctl) (*CtlCardInfo, error) {
	var info *C.snd_ctl_card_info_t

	if ret := C._new_snd_ctl_card_info(
		(*C.snd_ctl_t)(ctl),
		&info,
	); ret < 0 {
		return nil, &err{"control card info error", int(ret)}
	}

	return (*CtlCardInfo)(info), nil
}

type Ctl C.snd_ctl_t

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

func CtlPcmNextDevice(ctl *Ctl, dev int) (int, error) {
	cdev := C.int(dev)

	if ret := C.snd_ctl_pcm_next_device(
		(*C.snd_ctl_t)(ctl),
		&cdev,
	); ret < 0 {
		return -1, &err{"control pcm next device error", int(ret)}
	}

	return int(cdev), nil
}

func PcmStreamName(stream int) string {
	return C.GoString(C.snd_pcm_stream_name(C.snd_pcm_stream_t(stream)))
}

func StrError(code int) string {
	return C.GoString(C.snd_strerror(C.int(code)))
}
