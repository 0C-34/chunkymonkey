package gamerules

import (
	"os"

	"nbt"
)

const (
	chestInvWidth  = 9
	chestInvHeight = 3
)

type ChestInventory struct {
	Inventory
}

// InitChestInventory initializes inv as a 9x3 chest.
func NewChestInventory() (inv *ChestInventory) {
	inv = new(ChestInventory)
	inv.Inventory.Init(chestInvWidth * chestInvHeight)
	return inv
}

func (inv *ChestInventory) MarshalNbt(tag *nbt.Compound) (err os.Error) {
	tag.Set("id", &nbt.String{"Furnace"})
	return inv.Inventory.MarshalNbt(tag)
}
