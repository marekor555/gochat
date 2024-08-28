package layouts

import "fyne.io/fyne/v2"

type MainLayout struct{}

func (d *MainLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *MainLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	objects[0].Resize(fyne.NewSize(containerSize.Width, 40))
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(fyne.NewSize(containerSize.Width, containerSize.Height-80))
	objects[1].Move(fyne.NewPos(0, 40))

	objects[2].Resize(fyne.NewSize(containerSize.Width, 40))
	objects[2].Move(fyne.NewPos(0, containerSize.Height-40))
}
