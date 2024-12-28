package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TransparentButton struct {
	widget.BaseWidget
	onTapped func()
}

func NewTransparentButton(onTapped func()) *TransparentButton {
	btn := &TransparentButton{onTapped: onTapped}
	btn.ExtendBaseWidget(btn)
	return btn
}

func (btn *TransparentButton) Tapped(_ *fyne.PointEvent) {
	if btn.onTapped != nil {
		btn.onTapped()
	}
}

func (btn *TransparentButton) CreateRenderer() fyne.WidgetRenderer {
	return &transparentButtonRenderer{button: btn}
}

type transparentButtonRenderer struct {
	button *TransparentButton
}

func (r *transparentButtonRenderer) Layout(size fyne.Size) {}
func (r *transparentButtonRenderer) MinSize() fyne.Size    { return fyne.NewSize(100, 50) }
func (r *transparentButtonRenderer) Refresh()              {}
func (r *transparentButtonRenderer) Destroy()              {}
func (r *transparentButtonRenderer) Objects() []fyne.CanvasObject {
	return nil
}
