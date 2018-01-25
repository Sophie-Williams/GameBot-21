package main

import (
	"strings"
)

type AllWorlds struct {
	worlds []World
}

type World struct {
	places []Place
	hero   Hero
	chatId int64
}

func (allWorlds *AllWorlds) getWorld(chatId int64) *World {
	for i := range allWorlds.worlds {
		if allWorlds.worlds[i].getChatID() == chatId {
			return &allWorlds.worlds[i]
		}
	}
	world := World{}
	world.init(chatId)
	allWorlds.worlds = append(allWorlds.worlds, world)
	return &world
}

func (allWorlds *AllWorlds) reset(chatId int64) string {
	for i := range allWorlds.worlds {
		if allWorlds.worlds[i].getChatID() == chatId {
			allWorlds.worlds = append(allWorlds.worlds[:i], allWorlds.worlds[i+1:]...)
			return "состояние игры сброшено"
		}
	}
	return "а вы и так не играли :("
}

func (world *World) getChatID() int64 {
	return world.chatId
}

func (world *World) init(chatId int64) {
	world.chatId = chatId

	kitchen := Place{}
	kitchen.init("кухня",
		"кухня, ничего интересного",
		"ты находишься на кухне",
		"надо собрать рюкзак и идти в универ",
		"надо идти в универ",
		[]string{"коридор"},
		[]string{})

	kitchen.setObjects([]string{"чай"}, []string{})

	world.places = append(world.places, kitchen)

	quoridor := Place{}
	quoridor.init("коридор",
		"ничего интересного",
		"ничего интересного",
		"",
		"",
		[]string{"кухня", "комната"},
		[]string{"улица"})

	world.places = append(world.places, quoridor)

	room := Place{}
	room.init("комната",
		"ты в своей комнате",
		"",
		"",
		"",
		[]string{"коридор"},
		[]string{})

	room.setObjects([]string{"ключи", "конспекты"}, []string{"рюкзак"})

	world.places = append(world.places, room)

	street := Place{}
	street.init("улица",
		"на улице уже вовсю готовятся к новому году",
		"",
		"",
		"",
		[]string{"домой"},
		[]string{})

	world.places = append(world.places, street)

	world.hero.init(kitchen)
}

func (world *World) getState() string {
	return world.hero.getState()
}

func (world *World) changePosition(newPosition string) string {
	position := world.hero.getPosition()
	err := position.leave(newPosition)
	if err != nil {
		return err.Error()
	}

	for i := range world.places {
		if strings.Compare(world.places[i].getName(), newPosition) == 0 {
			world.hero.setPosition(world.places[i])
			return world.places[i].getStateWhenEntering()
		}
	}

	return "нет такого места"
}

func (world *World) putOn(item string) string {
	position := world.hero.getPosition()

	err := position.getObj(item)
	if err != nil {
		return err.Error()
	}

	resultString, err := world.hero.putOnItem(item)
	if err != nil {
		position.setObjOnFloor(item)
		return err.Error()
	}

	return resultString
}

func (world *World) getItem(item string) string {
	position := world.hero.getPosition()

	if !world.hero.checkBackpack() {
		return "некуда класть"
	}

	err := position.getObj(item)
	if err != nil {
		return err.Error()
	}
	var resultString string
	resultString = world.hero.putItem(item)

	return resultString
}

func (world *World) applyItem(itemAndObect string) string {
	itemAndObectArray := strings.Split(itemAndObect, " ")
	if len(itemAndObectArray) != 2 {
		return "Напишите правильно"
	}

	err := world.hero.checkItem(itemAndObectArray[0])
	if err != nil {
		return err.Error()
	}

	var resultString string
	resultString, err = world.hero.position.apply(itemAndObectArray[0], itemAndObectArray[1])
	if err != nil {
		return err.Error()
	}

	return resultString
}
