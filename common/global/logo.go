package global

import "github.com/common-nighthawk/go-figure"

func SmartLogo() string {
	myFigure := figure.NewFigure("Hello, SMART", "", true)
	return myFigure.String()
}
