package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/ini.v1"
)

func directorySet() string {
	dir, _ := os.UserHomeDir()
	dir += "\\AppData\\Local\\DeadByDaylight\\Saved\\Config\\EGS\\GameUserSettings.ini"
	fmt.Println(dir)
	return dir
}

func main() {
	a := app.New()
	w := a.NewWindow("dbd config changer")
	w.Resize(fyne.NewSize(800, 400))
	cfg, err := ini.Load(directorySet())
	if err != nil {
		log.Fatal(err)
	}

	start := widget.NewLabel("Dead By Daylight Config Changer")
	resLabel := widget.NewLabel("Resolution scale; current Value = not set")
	resSlider := widget.NewSlider(67, 100)
	resSlider.OnChanged = func(res float64) {
		resLabel.SetText("Resolution scale; current Value = " + fmt.Sprint(res))
		cfg.Section("ScalabilityGroups").Key("sg.ResolutionQuality").SetValue(fmt.Sprint(res))
	}
	widgets := []fyne.CanvasObject{
		start,
		widget.NewLabel("Graphical options"),
	}
	widgets = append(widgets, miscSlider(cfg, "Resolution scale;", "ScalabilityGroups", "MenuScaleFactor", 67, 100)...)
	widgets = append(widgets,
		newButtonsGraphical(cfg, "View Distance: ", "sg.ViewDistanceQuality"),
		newButtonsGraphical(cfg, "AntiAliasing: ", "sg.AntiAliasingQuality"),
		newButtonsGraphical(cfg, "Shadows: ", "sg.ShadowQuality"),
		newButtonsGraphical(cfg, "PostProccessing: ", "sg.PostProcessQuality"),
		newButtonsGraphical(cfg, "Texture Quality: ", "sg.TextureQuality"),
		newButtonsGraphical(cfg, "Effects: ", "sg.EffectsQuality"),
		newButtonsGraphical(cfg, "Foliage: ", "sg.FoliageQuality"),
		newButtonsGraphical(cfg, "Shading: ", "sg.ShadingQuality"),
		newButtonsGraphical(cfg, "Animations", "sg.AnimationQuality"),
	)
	w.SetContent(container.NewVBox(widgets...))
	w.SetOnClosed(func() {
		err := cfg.SaveTo(directorySet())
		if err != nil {
			log.Fatal(err)
		}
	}) //only BEFORE showandrun
	w.ShowAndRun() // anything after this func is not executed
}

func newButtonsGraphical(cfg *ini.File, label, key string) *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(label),
		widget.NewButton("minimum", func() {
			cfg.Section("ScalabilityGroups").Key(key).SetValue("0")
		}),
		widget.NewButton("medium", func() {
			cfg.Section("ScalabilityGroups").Key(key).SetValue("1")
		}),
		widget.NewButton("high", func() {
			cfg.Section("ScalabilityGroups").Key(key).SetValue("2")
		}),
		widget.NewButton("ultra", func() {
			cfg.Section("ScalabilityGroups").Key(key).SetValue("3")
		}),
	)
}
func miscSlider(cfg *ini.File, text, section, key string, min, max float64) []fyne.CanvasObject {
	label := widget.NewLabel(text)
	slider := widget.NewSlider(min, max)
	slider.OnChanged = func(f float64) {
		label.SetText(text + " current value: " + fmt.Sprint(f))
		cfg.Section(section).Key(key).SetValue(fmt.Sprint(f))
	}
	return []fyne.CanvasObject{label, slider}
}
