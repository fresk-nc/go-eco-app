package main

import (
	"fmt"
	"time"
)

const (
	landfillsFilename       = "landfills.csv"
	recyclingPointsFilename = "recyclingPoints.csv"
)

func main() {
	clear()

	landfills := newLandfillsFromFile(landfillsFilename)
	recyclingPoints := newRecyclingPointsFromFile(recyclingPointsFilename)

	m := menu{
		items: []menuItem{
			menuItem{
				title: "Незаконные свалки",
				action: func() bool {
					subM := menu{
						items: []menuItem{
							menuItem{
								title: "Вывести список незаконных свалок",
								action: func() bool {
									if len(landfills) > 0 {
										landfills.print()
									} else {
										fmt.Println("Список пуст")
									}
									return true
								},
							},
							menuItem{
								title: "Добавить информацию о незаконной свалке",
								action: func() bool {
									address := readStringFromStdin("Введите адрес: ", func(text string) bool {
										return len(text) > 0
									})
									status := readStringFromStdin("Введите статус: ", func(text string) bool {
										return len(text) > 0
									})
									date := readStringFromStdin("Введите дату(dd.mm.yyyy): ", func(text string) bool {
										_, err := time.Parse("02.01.2006", text)

										return err == nil
									})
									landfill := landfill{
										address: address,
										status:  status,
										date:    date,
									}
									landfills = addLandfill(landfills, landfill)
									landfills.saveToFile(landfillsFilename)
									clear()
									fmt.Println("Запись успешно добавлена")

									return true
								},
							},
							menuItem{
								title: "Удалить информацию о незаконной свалке",
								action: func() bool {
									if len(landfills) > 0 {
										landfills.print()
										printSeparator()
										i := readIntFromStdin("Введите номер незаконной свалки: ", func(i int) bool {
											return i > 0 && i <= len(landfills)
										})
										landfills = removeLandfill(landfills, i-1)
										landfills.saveToFile(landfillsFilename)
										clear()
										fmt.Println("Запись успешно удалена")
									} else {
										fmt.Println("Список пуст")
									}

									return true
								},
							},
							menuItem{
								title: "Вернуться в главное меню",
								action: func() bool {
									return false
								},
							},
						},
					}

					subM.run()

					return true
				},
			},
			menuItem{
				title: "Пункты переработки отходов",
				action: func() bool {
					subM := menu{
						items: []menuItem{
							menuItem{
								title: "Вывести список пунктов переработки отходов",
								action: func() bool {
									if len(recyclingPoints) > 0 {
										recyclingPoints.print()
									} else {
										fmt.Println("Список пуст")
									}

									return true
								},
							},
							menuItem{
								title: "Добавить пункт переработки отходов",
								action: func() bool {
									address := readStringFromStdin("Введите адрес: ", func(text string) bool {
										return len(text) > 0
									})
									wasteType := readStringFromStdin("Введите тип отходов: ", func(text string) bool {
										return len(text) > 0
									})
									recyclingPoint := recyclingPoint{
										address:   address,
										wasteType: wasteType,
									}
									recyclingPoints = addRecyclingPoint(recyclingPoints, recyclingPoint)
									recyclingPoints.saveToFile(recyclingPointsFilename)
									clear()
									fmt.Println("Запись успешно добавлена")

									return true
								},
							},
							menuItem{
								title: "Удалить пункт переработки отходов",
								action: func() bool {
									if len(recyclingPoints) > 0 {
										recyclingPoints.print()
										printSeparator()
										i := readIntFromStdin("Введите номер пункта переработки отходов: ", func(i int) bool {
											return i > 0 && i <= len(recyclingPoints)
										})
										recyclingPoints = removeRecyclingPoint(recyclingPoints, i-1)
										recyclingPoints.saveToFile(recyclingPointsFilename)
										clear()
										fmt.Println("Запись успешно удалена")
									} else {
										fmt.Println("Список пуст")
									}

									return true
								},
							},
							menuItem{
								title: "Вернуться в главное меню",
								action: func() bool {
									return false
								},
							},
						},
					}

					subM.run()

					return true
				},
			},
			menuItem{
				title: "Статистика",
				action: func() bool {
					fmt.Println("Количество незаконных свалок:", len(landfills))
					fmt.Println("Количество пунктов переработки отходов:", len(recyclingPoints))
					return true
				},
			},
			menuItem{
				title: "Выход из программы",
				action: func() bool {
					return false
				},
			},
		},
	}

	m.run()
}
