package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	myWindow fyne.Window
	myApp    fyne.App
)

func main() {
	createWindow()
}

func createWindow() {
	// Создаем новое приложение
	myApp = app.New()
	myWindow = myApp.NewWindow("Eat The Rich")

	windowSize := myWindow.Canvas().Size()
	// Рассчитываем размеры кнопок и заголовка (20% от ширины и 10% от высоты экрана)
	elementWidth := windowSize.Width * 0.2
	elementHeight := windowSize.Height * 0.1

	// Заголовок
	title := widget.NewLabel("EAT THE RICH")
	title.Alignment = fyne.TextAlignCenter
	title.Resize(fyne.NewSize(elementWidth, elementHeight))

	// Кнопки
	playAlone := widget.NewButton("PLAY ALONE", playAloneButton)
	playAlone.Resize(fyne.NewSize(elementWidth, elementHeight))

	connect := widget.NewButton("CONNECT", connectButton)
	connect.Resize(fyne.NewSize(elementWidth, elementHeight))

	createServer := widget.NewButton("CREATE A SERVER", createServerButton)
	createServer.Resize(fyne.NewSize(elementWidth, elementHeight))

	rules := widget.NewButton("RULES", rulesButton)
	rules.Resize(fyne.NewSize(elementWidth, elementHeight))

	exit := widget.NewButton("EXIT", exitButton)
	exit.Resize(fyne.NewSize(elementWidth, elementHeight))

	// Создаем контейнер с двумя колонками для кнопок rules и exit
	rulesExitContainer := container.NewGridWithColumns(2, rules, exit)

	// Вертикальное расположение остальных кнопок и горизонтальной группы
	buttons := container.NewVBox(
		playAlone,
		connect,
		createServer,
		rulesExitContainer,
	)

	// Центрируем заголовок и кнопки по вертикали и горизонтали
	content := container.NewVBox(
		container.NewCenter(title),   // Центрируем заголовок
		widget.NewSeparator(),        // Разделитель
		container.NewCenter(buttons), // Центрируем кнопки
	)

	paddedContent := container.NewCenter(content)

	myWindow.SetFullScreen(true)
	myWindow.SetContent(paddedContent)
	myWindow.ShowAndRun()
}

func playAloneButton() {
	fmt.Println("Starting a solo game...")
	dialog.ShowInformation("Play Alone", "Starting a solo game...", myWindow)
}

func connectButton() {
	fmt.Println("Connecting to server...")
	dialog.ShowInformation("Connect", "Connecting to server...", myWindow)
}

func createServerButton() {
	fmt.Println("Creating a new server...")
	dialog.ShowInformation("Create Server", "Server is being created...", myWindow)
}

func rulesButton() {
	fmt.Println("Displaying game rules...")
	rulesText := "1. Rule One\n2. Rule Two\n3. Rule Three\n\nEnjoy the game!"
	dialog.ShowInformation("Game Rules", rulesText, myWindow)
}

func exitButton() {
	fmt.Println("Exiting the game...")
	myApp.Quit()
}
