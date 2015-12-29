#include <alsa/asoundlib.h>
#include <stdio.h>

static
void _snd_ctl_card_info_alloca(snd_ctl_card_info_t** info) {
  snd_ctl_card_info_alloca(info);
}

static
int _new_snd_ctl_card_info(snd_ctl_t *ctl,
                           snd_ctl_card_info_t **info) {
  snd_ctl_card_info_alloca(info);
  return snd_ctl_card_info(ctl, *info);
}
