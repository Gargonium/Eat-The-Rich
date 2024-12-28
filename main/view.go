package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image"
	"log"
	"os"
)

var (
	myWindow  fyne.Window
	myApp     fyne.App
	mapScreen *fyne.Container
)

func createMapWindow() {
	var buttons []*TransparentButton
	//var buttons []*widget.Button

	var currentButtonX float32 = 45
	var currentButtonY float32 = 50

	buttonIndex := 1

	for i := 0; i < 30; i++ {
		b := NewTransparentButton(spaceButton(buttonIndex))
		//b := widget.NewButton("", spaceButton(buttonIndex))
		b.Resize(fyne.NewSize(152, 90))
		b.Move(fyne.NewPos(currentButtonX, currentButtonY))

		if (i+1)%6 == 0 {
			currentButtonX = 45
			currentButtonY += 140
		} else {
			currentButtonX += 210
		}

		if i == 2 || i == 12 || i == 17 || i == 22 {
			continue
		}

		buttonIndex++
		buttons = append(buttons, b)
	}

	imgFile, err := os.Open("resources/map.png")
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	defer imgFile.Close()

	bgImage, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	background := canvas.NewImageFromImage(bgImage)
	background.FillMode = canvas.ImageFillContain
	background.Move(fyne.NewPos(0, 0))
	background.Resize(fyne.NewSize(1280, 720))

	mapScreen = container.NewWithoutLayout(background)
	for _, button := range buttons {
		mapScreen.Add(button)
	}
}

func spaceButton(id int) func() {
	return func() {
		fmt.Printf("Button %d pressed!\n", id)
	}
}

func createMainWindow() {
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

	myWindow.SetContent(paddedContent)

	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1280, 720))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

func playAloneButton() {
	fmt.Println("Starting a solo game...")
	//dialog.ShowInformation("Play Alone", "Starting a solo game...", myWindow)
	createMapWindow()
	myWindow.SetContent(mapScreen)
}

func connectButton() {
	fmt.Println("Connecting to server...")
	dialog.ShowInformation("Connect", "Connecting to server...", myWindow)

	//go getServerIp()
	//for {
	//	if serverIp == "" {
	//		dialog.ShowInformation("No server found", "Try again later or create your own server", myWindow)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println(serverIp)
}

func createServerButton() {
	fmt.Println("Creating a new server...")
	dialog.ShowInformation("Create Server", "Server is being created...", myWindow)
	//go sendMulticast()
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
