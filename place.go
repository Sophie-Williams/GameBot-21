package main

import (
	"errors"
	"strings"
)

type Place struct {
	name                           string
	whenEntering                   string
	descriptionBefore              string
	descriptionAfter               string
	descriptionAfterIfHaveBackpack string
	onTable                        []string
	onChair                        []string
	onFloor                        []string
	ways                           []string
	closeWays                      []string
}

func (place *Place) init(name, whenEntering, descriptionBefore, descriptionAfter, descriptionAfterIfHaveBackpack string, ways, closeWays []string) {
	place.name = name
	place.whenEntering = whenEntering
	place.descriptionBefore = descriptionBefore
	place.descriptionAfter = descriptionAfter
	place.descriptionAfterIfHaveBackpack = descriptionAfterIfHaveBackpack
	place.ways = ways
	place.closeWays = closeWays
}

func (place *Place) getName() string {
	return place.name
}

func (place *Place) setObjects(onTable, onChair []string) {
	place.onTable = onTable
	place.onChair = onChair
}

func (place *Place) getStateWhenEntering() string {
	resultString := place.whenEntering + ". можно пройти - " + strings.Join(place.ways, ", ")

	if len(place.closeWays) != 0 {
		resultString += ", " + strings.Join(place.closeWays, ", ")
	}

	resultString += "."

	return resultString
}

func (place *Place) getObj(item string) error {
	for i := range place.onTable {
		if strings.Compare(place.onTable[i], item) == 0 {
			place.onTable = append(place.onTable[:i], place.onTable[i+1:]...)
			return nil
		}
	}

	for i := range place.onChair {
		if strings.Compare(place.onChair[i], item) == 0 {
			place.onChair = append(place.onChair[:i], place.onChair[i+1:]...)
			return nil
		}
	}

	for i := range place.onFloor {
		if strings.Compare(place.onFloor[i], item) == 0 {
			place.onFloor = append(place.onFloor[:i], place.onFloor[i+1:]...)
			return nil
		}
	}

	return errors.New("нет такого")
}

func (place *Place) leave(direction string) error {
	if stringInSlice(direction, place.ways) {
		return nil
	}

	if stringInSlice(direction, place.closeWays) {
		return errors.New("дверь закрыта")
	}

	return errors.New("нет пути в " + direction)
}

func (place *Place) apply(item, object string) (string, error) {
	if object == "дверь" && len(place.closeWays) != 0 && item == "ключи" {
		place.ways = append(place.ways, place.closeWays...)
		place.closeWays = place.closeWays[:]
		return "дверь открыта", nil
	}

	return "", errors.New("не к чему применить")
}

func (place *Place) setObjOnFloor(item string) {
	place.onFloor = append(place.onFloor, item)
}
