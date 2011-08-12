package chunkstore

import (
	"chunkymonkey/gamerules"
	. "chunkymonkey/types"
	"nbt"
)

func cloneByteArray(in []byte) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	return out
}

type nbtChunkWriter struct {
	loc ChunkXz

	// The NBT structure created.
	chunkTag *nbt.Compound
}

func newNbtChunkWriter() *nbtChunkWriter {
	return &nbtChunkWriter{
		chunkTag: &nbt.Compound{map[string]nbt.ITag{
			"Level": &nbt.Compound{map[string]nbt.ITag{
				"Entities":         &nbt.List{nbt.TagCompound, nil},
				"TileEntities":     &nbt.List{nbt.TagCompound, nil}, // TODO
				"Blocks":           &nbt.ByteArray{},
				"Data":             &nbt.ByteArray{},
				"HeightMap":        &nbt.ByteArray{},
				"SkyLight":         &nbt.ByteArray{},
				"BlockLight":       &nbt.ByteArray{},
				"LastUpdate":       &nbt.Long{0}, // TODO
				"TerrainPopulated": &nbt.Byte{1}, // TODO
				"xPos":             &nbt.Int{0},
				"zPos":             &nbt.Int{0},
			}},
		}},
	}
}

func (w *nbtChunkWriter) ChunkLoc() ChunkXz {
	return w.loc
}

func (w *nbtChunkWriter) SetChunkLoc(loc ChunkXz) {
	w.loc = loc
	w.chunkTag.Lookup("Level/xPos").(*nbt.Int).Value = int32(loc.X)
	w.chunkTag.Lookup("Level/zPos").(*nbt.Int).Value = int32(loc.Z)
}

func (w *nbtChunkWriter) SetBlocks(blocks []byte) {
	w.chunkTag.Lookup("Level/Blocks").(*nbt.ByteArray).Value = cloneByteArray(blocks)
}

func (w *nbtChunkWriter) SetBlockData(blockData []byte) {
	w.chunkTag.Lookup("Level/Data").(*nbt.ByteArray).Value = cloneByteArray(blockData)
}

func (w *nbtChunkWriter) SetBlockLight(blockLight []byte) {
	w.chunkTag.Lookup("Level/BlockLight").(*nbt.ByteArray).Value = cloneByteArray(blockLight)
}

func (w *nbtChunkWriter) SetSkyLight(skyLight []byte) {
	w.chunkTag.Lookup("Level/SkyLight").(*nbt.ByteArray).Value = cloneByteArray(skyLight)
}

func (w *nbtChunkWriter) SetHeightMap(heightMap []byte) {
	w.chunkTag.Lookup("Level/HeightMap").(*nbt.ByteArray).Value = cloneByteArray(heightMap)
}

func (w *nbtChunkWriter) SetEntities(entities map[EntityId]gamerules.INonPlayerEntity) {
	entitiesNbt := make([]nbt.ITag, 0, len(entities))
	for _, entity := range entities {
		nbtData := entity.WriteNbt()
		if nbtData != nil {
			entitiesNbt = append(entitiesNbt, nbtData)
		}
	}
	w.chunkTag.Lookup("Level/Entities").(*nbt.List).Value = entitiesNbt
}
