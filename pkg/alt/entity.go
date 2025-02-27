package alt

/*
#include <stdlib.h>
#include "capi.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/timo972/altv-go/internal/module"
)

type Entity struct {
	WorldObject
	id uint16
}

type IEntity interface {
	World
	AltEntity()
	Model() uint32
	IsPlayer() bool
	IsVehicle() bool
	Detach()
	AttachToEntity(entity IEntity, otherBoneIndex int16, myBoneIndex int16, position Vector3, rotation Vector3, collision bool, noFixedRotation bool)
	SetVisible(toggle bool)
	Visible() bool
	NetworkOwner() IPlayer
	ID() uint16
}

func newEntityArray(arr C.struct_array) []IEntity {
	size := int(arr.size)
	// FIXME: may causes a crash because it frees to early
	defer C.free(unsafe.Pointer(arr.array))

	cArray := (*[1 << 28]C.struct_entity)(arr.array)[:size:size]

	entities := make([]IEntity, size)

	if size == 0 {
		return entities
	}

	for i := 0; i < size; i++ {
		p := cArray[i]
		entities[i] = getEntity(p)
	}

	return entities
}

func newCEntity(e IEntity) C.struct_entity {
	return C.struct_entity{
		ptr: e.NativePointer(),
		typ: C.uchar(e.Type()),
	}
}

// AltEntity empty function to satisfy IEntity interface
func (e Entity) AltEntity() {}

func (e Entity) String() string {
	return fmt.Sprintf("Entity{}")
}

func (e Entity) IsPlayer() bool {
	return e.Type() == PlayerObject
}

func (e Entity) IsVehicle() bool {
	return e.Type() == VehicleObject
}

/*func (e Entity) AsPlayer() IPlayer {
	if e.Type() != PlayerObject {
		return nil
	}

	return e
}

func (e Entity) AsVehicle() IVehicle {
	if e.Type() != VehicleObject {
		return nil
	}

	return e
}*/

func (e Entity) Model() uint32 {
	if e.Type() == PlayerObject {
		return uint32(C.player_get_model(e.ptr))
	} else if e.Type() == VehicleObject {
		return uint32(C.vehicle_get_model(e.ptr))
	}
	return 0
}

func (e Entity) Detach() {
	if e.Type() == PlayerObject {
		C.player_detach(e.ptr)
	} else if e.Type() == VehicleObject {
		C.vehicle_detach(e.ptr)
	}
}

func (e Entity) AttachToEntity(entity IEntity, otherBoneIndex int16, myBoneIndex int16, position Vector3, rotation Vector3, collision bool, noFixedRotation bool) {
	if e.Type() == PlayerObject {
		C.player_attach_to_entity(e.ptr, newCEntity(entity), C.int(otherBoneIndex), C.int(myBoneIndex), newCPosition(position), newCRotation(rotation), C.int(module.Bool2int(collision)), C.int(module.Bool2int(noFixedRotation)))
	} else if e.Type() == VehicleObject {
		C.vehicle_attach_to_entity(e.ptr, newCEntity(entity), C.int(otherBoneIndex), C.int(myBoneIndex), newCPosition(position), newCRotation(rotation), C.int(module.Bool2int(collision)), C.int(module.Bool2int(noFixedRotation)))
	}
}

func (e Entity) SetVisible(toggle bool) {
	if e.Type() == PlayerObject {
		C.player_set_visible(e.ptr, C.int(module.Bool2int(toggle)))
	} else if e.Type() == VehicleObject {
		C.vehicle_set_visible(e.ptr, C.int(module.Bool2int(toggle)))
	}
}

func (e Entity) Visible() bool {
	if e.Type() == PlayerObject {
		return int(C.player_get_visible(e.ptr)) == 1
	} else if e.Type() == VehicleObject {
		return int(C.vehicle_get_visible(e.ptr)) == 1
	}
	return false
}

func (e Entity) ID() uint16 {
	return e.id
	/*if e.Type() == PlayerObject {
		return uint16(C.player_get_i_d(e.ptr))
	} else if e.Type() == VehicleObject {
		return uint16(C.vehicle_get_i_d(e.ptr))
	}
	return 0*/
}

func (e Entity) NetworkOwner() IPlayer {
	var cPtr C.struct_entity
	if e.Type() == PlayerObject {
		cPtr = C.player_get_network_owner(e.ptr)
	} else if e.Type() == VehicleObject {
		cPtr = C.vehicle_get_network_owner(e.ptr)
	}
	return getPlayer(cPtr)
}

func (e Entity) SetNetworkOwner(owner IPlayer, disableMigration bool) {
	if e.Type() == PlayerObject {
		C.player_set_network_owner(e.ptr, owner.NativePointer(), C.int(module.Bool2int(disableMigration)))
	} else if e.Type() == VehicleObject {
		C.vehicle_set_network_owner(e.ptr, owner.NativePointer(), C.int(module.Bool2int(disableMigration)))
	}
}

func (e Entity) Rotation() Vector3 {
	var cRot C.struct_rot
	if e.Type() == PlayerObject {
		cRot = C.player_get_rotation(e.ptr)
	} else if e.Type() == VehicleObject {
		cRot = C.vehicle_get_rotation(e.ptr)
	}
	return Vector3{X: float32(cRot.roll), Y: float32(cRot.pitch), Z: float32(cRot.yaw)}
}

func (e Entity) SetRotation(rotation Vector3) {
	if e.Type() == PlayerObject {
		C.player_set_rotation(e.ptr, C.float(rotation.X), C.float(rotation.Y), C.float(rotation.Z))
	} else if e.Type() == VehicleObject {
		C.vehicle_set_rotation(e.ptr, C.float(rotation.X), C.float(rotation.Y), C.float(rotation.Z))
	}
}

func (e Entity) HasSyncedMetaData(key string) bool {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	if e.Type() == PlayerObject {
		return int(C.player_has_synced_meta_data(e.ptr, cKey)) == 1
	}

	if e.Type() == VehicleObject {
		return int(C.vehicle_has_synced_meta_data(e.ptr, cKey)) == 1
	}

	return false
}

func (e Entity) SyncedMetaData(key string, val interface{}) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	/*var meta C.struct_array
	if e.Type() == PlayerObject {
		meta = C.player_get_synced_meta_data(e.ptr, cKey)
	} else if e.Type() == VehicleObject {
		meta = C.vehicle_get_synced_meta_data(e.ptr, cKey)
	}*/

	// err := decode(meta, val)

	return nil
}

func (e Entity) HasStreamSyncedMetaData(key string) bool {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	if e.Type() == PlayerObject {
		return int(C.player_has_stream_synced_meta_data(e.ptr, cKey)) == 1
	}

	if e.Type() == VehicleObject {
		return int(C.vehicle_has_stream_synced_meta_data(e.ptr, cKey)) == 1
	}

	return false
}

func (e Entity) StreamSyncedMetaData(key string, value interface{}) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	/*var meta C.struct_array
	if e.Type() == PlayerObject {
		meta = C.player_get_stream_synced_meta_data(e.ptr, cKey)
	} else if e.Type() == VehicleObject {
		meta = C.vehicle_get_stream_synced_meta_data(e.ptr, cKey)
	}

	err := decode(meta, value)*/

	return nil
}

func (e Entity) SetSyncedMetaData(key string, value interface{}) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	/*arr, err := encode(value)
	if err != nil {
		return false
	}
	defer C.free(arr.array)

	bytes := (*C.uchar)(arr.array)

	if e.Type() == PlayerObject {
		C.player_set_synced_meta_data(e.ptr, cKey, bytes, arr.size)
	} else if e.Type() == VehicleObject {
		C.vehicle_set_synced_meta_data(e.ptr, cKey, bytes, arr.size)
	}*/

	return nil
}

func (e Entity) DeleteSyncedMetaData(key string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	if e.Type() == PlayerObject {
		C.player_delete_synced_meta_data(e.ptr, cKey)
	}

	if e.Type() == VehicleObject {
		C.vehicle_delete_synced_meta_data(e.ptr, cKey)
	}
}

func (e Entity) SetStreamSyncedMetaData(key string, value interface{}) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	/*arr, err := encode(value)
	if err != nil {
		return false
	}
	defer C.free(arr.array)

	bytes := (*C.uchar)(arr.array)

	if e.Type() == PlayerObject {
		C.player_set_stream_synced_meta_data(e.ptr, cKey, bytes, arr.size)
	} else if e.Type() == VehicleObject {
		C.vehicle_set_stream_synced_meta_data(e.ptr, cKey, bytes, arr.size)
	}*/

	return nil
}

func (e Entity) DeleteStreamSyncedMetaData(key string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	if e.Type() == PlayerObject {
		C.player_delete_stream_synced_meta_data(e.ptr, cKey)
	}

	if e.Type() == VehicleObject {
		C.vehicle_delete_stream_synced_meta_data(e.ptr, cKey)
	}
}

func (e Entity) Streamed() bool {
	if e.Type() == PlayerObject {
		return int(C.player_get_streamed(e.ptr)) == 1
	} else if e.Type() == VehicleObject {
		return int(C.vehicle_get_streamed(e.ptr)) == 1
	}
	return false
}

func (e Entity) SetStreamed(toggle bool) {
	state := C.int(module.Bool2int(toggle))
	if e.Type() == PlayerObject {
		C.player_set_streamed(e.ptr, state)
	} else if e.Type() == VehicleObject {
		C.vehicle_set_streamed(e.ptr, state)
	}
}

func (e Entity) IsFrozen() bool {
	var cFrozen C.int
	if e.Type() == PlayerObject {
		cFrozen = C.player_is_frozen(e.ptr)
	} else if e.Type() == VehicleObject {
		cFrozen = C.vehicle_is_frozen(e.ptr)
	}

	return int(cFrozen) == 1
}

func (e Entity) SetFrozen(state bool) {
	s := C.int(module.Bool2int(state))
	if e.Type() == PlayerObject {
		C.player_set_frozen(e.ptr, s)
	} else if e.Type() == VehicleObject {
		C.vehicle_set_frozen(e.ptr, s)
	}
}

func (e Entity) HasCollision() bool {
	var cCol C.int
	if e.Type() == PlayerObject {
		cCol = C.player_has_collision(e.ptr)
	} else if e.Type() == VehicleObject {
		cCol = C.vehicle_has_collision(e.ptr)
	}

	return int(cCol) == 1
}

func (e Entity) SetCollision(state bool) {
	s := C.int(module.Bool2int(state))
	if e.Type() == PlayerObject {
		C.player_set_collision(e.ptr, s)
	} else if e.Type() == VehicleObject {
		C.vehicle_set_collision(e.ptr, s)
	}
}
