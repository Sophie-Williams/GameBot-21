package main

import (
	"errors"
	"strings"
)

type Hero struct {
	backpack  bool
	inventory []string
	position  Place
}

func (hero *Hero) init(position Place) {
	hero.position = position
}

func (hero *Hero) checkBackpack() bool {
	return hero.backpack
}

func (hero *Hero) setPosition(newPosition Place) {
	hero.position = newPosition
}

func (hero *Hero) getPosition() *Place {
	return &hero.position
}

func (hero *Hero) checkItemInInventory(item string) bool {
	return stringInSlice(item, hero.inventory)
}

func (hero *Hero) putItem(item string) string {
	hero.inventory = append(hero.inventory, item)
	return "предмет добавлен в инвентарь: " + item
}

func (hero *Hero) putOnItem(item string) (string, error) {
	switch {
	case strings.Compare(item, "рюкзак") == 0:
		hero.backpack = true
	default:
		return "", errors.New("такой предмет нельзя надеть :(")
	}

	hero.inventory = append(hero.inventory, item)
	return "вы надели: " + item, nil
}

func (hero *Hero) checkItem(item string) error {
	if !stringInSlice(item, hero.inventory) {
		return errors.New("нет предмета в инвентаре - " + item)
	}

	return nil
}

func (hero *Hero) getState() string {
	place := hero.getPosition()
	resultString := place.descriptionBefore

	if len(place.onTable) == 0 && len(place.onChair) == 0 && len(place.onFloor) == 0 {
		if place.name == "улица" {
			resultString += "пустая не комната"
		} else {
			resultString += "пустая комната"
		}
	} else {
		switch len(place.onTable) {
		case 0:
			break
		case 1:
			if resultString != "" {
				resultString += ", "
			}
			resultString += "на столе " + place.onTable[0]
		default:
			if resultString != "" {
				resultString += ", "
			}
			resultString += "на столе: " + strings.Join(place.onTable, ", ")
		}

		switch len(place.onChair) {
		case 0:
			break
		case 1:
			if resultString != "" {
				resultString += ", "
			}
			resultString += "на стуле - " + place.onChair[0]
		}

		switch len(place.onFloor) {
		case 0:
			break
		case 1:
			if resultString != "" {
				resultString += ", "
			}
			resultString += "ты разозлился и бросил на пол " + place.onFloor[0]
		default:
			if resultString != "" {
				resultString += ", "
			}
			resultString += "ты разозлился и бросил на пол: " + strings.Join(place.onFloor, ", ")
		}
	}

	if place.descriptionAfter != "" {
		resultString += ", "
		if hero.checkBackpack() {
			resultString += place.descriptionAfterIfHaveBackpack
		} else {
			resultString += place.descriptionAfter

		}
	}

	resultString += ". можно пройти - " + strings.Join(place.ways, ", ")

	if len(place.closeWays) != 0 {
		resultString += ", " + strings.Join(place.closeWays, ", ")
	}

	resultString += "."

	return resultString
}
