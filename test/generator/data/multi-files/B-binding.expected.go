// automatically generated by the ObjectBox, do not modify

package object

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type BBinding struct {
}

func (BBinding) AddToModel(model *objectbox.Model) {
	model.Entity("B", 2, 501233450539197794)
	model.Property("Id", objectbox.PropertyType_Long, 1, 3390393562759376202)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Name", objectbox.PropertyType_String, 2, 2669985732393126063)
	model.Property("Info", objectbox.PropertyType_String, 3, 1774932891286980153)
	model.EntityLastPropertyId(3, 1774932891286980153)
}

func asB(entity interface{}) (*B, error) {
	ent, ok := entity.(*B)
	if !ok {
		// Programming error, OK to panic
		// TODO don't panic here, handle in the caller if necessary to panic
		panic("Object has wrong type, expecting 'B'")
	}
	return ent, nil
}

func asBs(entities interface{}) ([]B, error) {
	ent, ok := entities.([]B)
	if !ok {
		// Programming error, OK to panic
		// TODO don't panic here, handle in the caller if necessary to panic
		panic("Object has wrong type, expecting 'B'")
	}
	return ent, nil
}

func (BBinding) GetId(entity interface{}) (uint64, error) {
	if ent, err := asB(entity); err != nil {
		return 0, err
	} else {
		return ent.Id, nil
	}
}

func (BBinding) Flatten(entity interface{}, fbb *flatbuffers.Builder, id uint64) {
	ent, err := asB(entity)
	if err != nil {
		// TODO return error and panic in the caller if really, really necessary
		panic(err)
	}

	// prepare the "offset" properties
	var offsetName = fbutils.CreateStringOffset(fbb, ent.Name)
	var offsetInfo = fbutils.CreateStringOffset(fbb, ent.Info)

	// build the FlatBuffers object
	fbb.StartObject(3)
	fbb.PrependUint64Slot(0, id, 0)
	fbb.PrependUOffsetTSlot(1, offsetName, 0)
	fbb.PrependUOffsetTSlot(2, offsetInfo, 0)
}

func (BBinding) ToObject(bytes []byte) interface{} {
	table := fbutils.GetRootAsTable(bytes, flatbuffers.UOffsetT(0))

	return &B{
		Id:   table.OffsetAsUint64(4),
		Name: table.OffsetAsString(6),
		Info: table.OffsetAsString(8),
	}
}

func (BBinding) MakeSlice(capacity int) interface{} {
	return make([]B, 0, capacity)
}

func (BBinding) AppendToSlice(slice interface{}, entity interface{}) interface{} {
	return append(slice.([]B), *entity.(*B))
}

type BBox struct {
	*objectbox.Box
}

func BoxForB(ob *objectbox.ObjectBox) *BBox {
	return &BBox{
		Box: ob.Box(2),
	}
}

func (box *BBox) Get(id uint64) (*B, error) {
	entity, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	}
	return asB(entity)
}

func (box *BBox) GetAll() ([]B, error) {
	entities, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return asBs(entities)
}

func (box *BBox) Remove(entity *B) (err error) {
	return box.Box.Remove(entity.Id)
}