package layouts

import "fyne.io/fyne/v2"

type InputLayout struct{}

func (d *InputLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *InputLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var size float32 = 100
	objects[0].Resize(fyne.NewSize(containerSize.Width-size, containerSize.Height))
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(fyne.NewSize(size*0.95, containerSize.Height))
	objects[1].Move(fyne.NewPos(containerSize.Width-size*0.95, 0))
}
