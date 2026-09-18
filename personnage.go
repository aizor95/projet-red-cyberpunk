type character struct {
	nom string
	class string
	level int
	HP int
	MaxHP int
	inventaire []string
	argent int
	Equipment Equipment
}

type Equipment struct {
	head EquipmentItem
	torse EquipmentItem
	feet EquipmentItem
}

type EquipmentItem struct {
	Name string
	Slot string
	MaxHPbonus int
	Attack int
	Defense int
	Energy int
}
