package alt

/*
#cgo windows CFLAGS: -I../../lib
#cgo windows LDFLAGS: -Wl,--subsystem,windows,--kill-at

#cgo linux CFLAGS: -I../../lib


#ifndef GOLANG_APP
#define GOLANG_APP

#include <stdlib.h>
#include "capi.h"

#endif
*/
import "C"
import (
	"fmt"
	"github.com/timo972/altv-go/internal/module"
)

type VoiceChannel struct {
	BaseObject
}

func CreateVoiceChannel(spatial bool, maxDistance float32) IVoiceChannel {
	e := C.core_create_voice_channel(C.int(module.Bool2int(spatial)), C.float(maxDistance))
	return getVoiceChannel(e)
}

func (v VoiceChannel) String() string {
	return fmt.Sprintf("VoiceChannel{}")
}

func (v VoiceChannel) IsSpatial() bool {
	return int(C.voice_channel_is_spatial(v.ptr)) == 1
}

func (v VoiceChannel) MaxDistance() float32 {
	return float32(C.voice_channel_get_max_distance(v.ptr))
}

func (v VoiceChannel) HasPlayer(player IPlayer) bool {
	return int(C.voice_channel_has_player(v.ptr, player.NativePointer())) == 1
}

func (v VoiceChannel) AddPlayer(player IPlayer) {
	C.voice_channel_add_player(v.ptr, player.NativePointer())
}

func (v VoiceChannel) RemovePlayer(player IPlayer) {
	C.voice_channel_remove_player(v.ptr, player.NativePointer())
}

func (v VoiceChannel) IsPlayerMuted(player IPlayer) bool {
	return int(C.voice_channel_is_player_muted(v.ptr, player.NativePointer())) == 1
}

func (v VoiceChannel) MutePlayer(player IPlayer) {
	C.voice_channel_mute_player(v.ptr, player.NativePointer())
}

func (v VoiceChannel) UnmutePlayer(player IPlayer) {
	C.voice_channel_unmute_player(v.ptr, player.NativePointer())
}

func (v VoiceChannel) PlayerCount() uint64 {
	return uint64(C.voice_channel_get_player_count(v.ptr))
}

func (v VoiceChannel) Players() []IPlayer {
	arr := C.voice_channel_get_players(v.ptr)

	return newPlayerArray(arr)
}
