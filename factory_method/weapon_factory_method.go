package factory_method

func GetWeapon(weaponType string) Weapon {
	if weaponType == "sword" {
		return &Sword{}
	} else if weaponType == "bow" {
		return &Bow{
			arrows: 5,
		}
	}
	return &EmptyHanded{}
}
