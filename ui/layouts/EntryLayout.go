package layouts

import (
	"fyne.io/fyne/v2"
)

type EntryLayout struct{}

func (d *EntryLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(400, objects[0].MinSize().Height)
}

func (d *EntryLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	objects[0].Resize(objects[0].MinSize())
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(fyne.NewSize(400-objects[0].MinSize().Width, containerSize.Height))
	objects[1].Move(fyne.NewPos(objects[0].MinSize().Width, 0))
}
