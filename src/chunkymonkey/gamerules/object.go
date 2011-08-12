// Defines non-block movable objects such as arrows in flight, boats and
// minecarts.

package gamerules

import (
	"io"
	"os"

	"chunkymonkey/physics"
	"chunkymonkey/proto"
	. "chunkymonkey/types"
	"nbt"
)

// TODO Object sub-types?

type Object struct {
	EntityId
	ObjTypeId
	physics.PointObject
	orientation OrientationBytes
}

func NewObject(objType ObjTypeId) (object *Object) {
	object = &Object{
		// TODO: proper orientation
		orientation: OrientationBytes{0, 0, 0},
	}
	object.ObjTypeId = objType
	return
}

func (object *Object) ReadNbt(tag nbt.ITag) (err os.Error) {
	if err = object.PointObject.ReadNbt(tag); err != nil {
		return
	}

	var typeName string
	if entityObjectId, ok := tag.Lookup("id").(*nbt.String); !ok {
		return os.NewError("missing object type id")
	} else {
		typeName = entityObjectId.Value
	}

	var ok bool
	if object.ObjTypeId, ok = ObjTypeByName[typeName]; !ok {
		return os.NewError("unknown object type id")
	}

	// TODO load orientation

	return
}

func (object *Object) WriteNbt() nbt.ITag {
	objTypeName, ok := ObjNameByType[object.ObjTypeId]
	if !ok {
		return nil
	}
	tag := &nbt.Compound{map[string]nbt.ITag{
		"id": &nbt.String{objTypeName},
		// TODO unknown fields
	}}
	object.PointObject.WriteIntoNbt(tag)
	return tag
}

func (object *Object) SendSpawn(writer io.Writer) (err os.Error) {
	// TODO: Send non-nil ObjectData (is there any?)
	err = proto.WriteObjectSpawn(writer, object.EntityId, object.ObjTypeId, &object.PointObject.LastSentPosition, nil)
	if err != nil {
		return
	}

	err = proto.WriteEntityVelocity(writer, object.EntityId, &object.PointObject.LastSentVelocity)
	return
}

func (object *Object) SendUpdate(writer io.Writer) (err os.Error) {
	if err = proto.WriteEntity(writer, object.EntityId); err != nil {
		return
	}

	// TODO: Should this be the Rotation information?
	err = object.PointObject.SendUpdate(writer, object.EntityId, &LookBytes{0, 0})

	return
}


func NewBoat() INonPlayerEntity {
	return NewObject(ObjTypeIdBoat)
}

func NewMinecart() INonPlayerEntity {
	return NewObject(ObjTypeIdMinecart)
}

func NewStorageCart() INonPlayerEntity {
	return NewObject(ObjTypeIdStorageCart)
}

func NewPoweredCart() INonPlayerEntity {
	return NewObject(ObjTypeIdPoweredCart)
}

func NewActivatedTnt() INonPlayerEntity {
	return NewObject(ObjTypeIdActivatedTnt)
}

func NewArrow() INonPlayerEntity {
	return NewObject(ObjTypeIdArrow)
}

func NewThrownSnowball() INonPlayerEntity {
	return NewObject(ObjTypeIdThrownSnowball)
}

func NewThrownEgg() INonPlayerEntity {
	return NewObject(ObjTypeIdThrownEgg)
}

func NewFallingSand() INonPlayerEntity {
	return NewObject(ObjTypeIdFallingSand)
}

func NewFallingGravel() INonPlayerEntity {
	return NewObject(ObjTypeIdFallingGravel)
}

func NewFishingFloat() INonPlayerEntity {
	return NewObject(ObjTypeIdFishingFloat)
}
