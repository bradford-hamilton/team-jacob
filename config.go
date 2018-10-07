package main

import (
	"bufio"
	"os"
	"strings"

	ui "github.com/gizak/termui"
)

func loadDepositAddresses() map[string]string {
	m := make(map[string]string)
	f, err := os.Open("addresses.cfg")
	if err != nil {
		Log.Println(err)
		return m
	}

	sc := bufio.NewScanner(f)
	if err != nil {
		Log.Println(err)
		return m
	}

	for sc.Scan() {
		arr := strings.Split(sc.Text(), " ")
		if len(arr) == 2 {
			m[arr[0]] = arr[1]
			Log.Println("Loaded destination address:", arr[0], arr[1])
		}
	}

	return m
}

// HeaderConfig is the configuration for the header which includes the logo and fox image
type HeaderConfig struct {
	FoxX, FoxY          int
	FoxHeight, FoxWidth int
	FoxTextFgColor      ui.Attribute

	LogoX, LogoY          int
	LogoHeight, LogoWidth int
	LogoTextFgColor       ui.Attribute
}

var DefaultHeaderConfig = &HeaderConfig{
	FoxX:            75,
	FoxY:            0,
	FoxHeight:       8,
	FoxWidth:        29,
	FoxTextFgColor:  ui.ColorCyan,
	LogoX:           10,
	LogoY:           1,
	LogoWidth:       70,
	LogoHeight:      7,
	LogoTextFgColor: ui.ColorCyan,
}

type ExchangeConfig struct {
	DepHeight, RecHeight, RetHeight int
	DepWidth, RecWidth, RetWidth    int
	DepX, RecX, RetX                int
	DepY, RecY, RetY                int
	DepColor, RecColor, RetColor    ui.Attribute
}

var DefaultExchangeConfig = &ExchangeConfig{
	DepHeight: 3,
	DepWidth:  46,
	DepX:      67,
	DepY:      15,
	DepColor:  ui.ColorRed,
	RecHeight: 3,
	RecWidth:  46,
	RecX:      67,
	RecY:      19,
	RecColor:  ui.ColorGreen,
	RetHeight: 3,
	RetWidth:  46,
	RetX:      67,
	RetY:      23,
	RetColor:  ui.ColorYellow,
}

type statsConfig struct {
	x, y                                    int
	minWidth, maxWidth, rateWidth, feeWidth int
}

var defaultStatsConfig = statsConfig{
	x:         1,
	y:         8,
	minWidth:  25,
	maxWidth:  25,
	rateWidth: 25,
	feeWidth:  25,
}

type LoadingConfig struct {
	X, Y          int
	Width, Height int
	TextFgColor   ui.Attribute
}

var DefaultLoadingConfig = &LoadingConfig{
	X:      30,
	Y:      10,
	Width:  10,
	Height: 3,
}

type setupConfig struct {
	entryX                int
	entryWidth            int
	entryHeight           int
	amtY, addrY, retY     int
	helpX, helpY          int
	helpHeight, helpWidth int
}

var setupQuickConfig = &setupConfig{
	entryX:      20,
	entryWidth:  60,
	entryHeight: 3,
	addrY:       13,
	retY:        20,
	helpX:       15,
	helpY:       26,
	helpHeight:  4,
	helpWidth:   71,
}

var setupPreciseConfig = &setupConfig{
	entryX:      20,
	entryWidth:  60,
	entryHeight: 3,
	amtY:        12,
	addrY:       17,
	retY:        22,
	helpX:       15,
	helpY:       26,
	helpHeight:  4,
	helpWidth:   71,
}

const FOX = `            ,^
           ;  ;
\'.,'/      ; ;
/_  _\'-----';

  \/' ,,,,,, ;
    )//     \))`

const SHAPESHIFT = "" +
	"  ____  _                      ____  _     _  __ _   \n" +
	" / ___|| |__   __ _ _ __   ___/ ___|| |__ (_)/ _| |_ \n" +
	" \\___ \\| '_ \\ / _` | '_ \\ / _ \\___ \\| '_ \\| | |_| __|\n" +
	"  ___) | | | | (_| | |_) |  __/___) | | | | |  _| |_ \n" +
	" |____/|_| |_|\\__,_| .__/ \\___|____/|_| |_|_|_|  \\__|\n" +
	"                   |_|                               "
	/*
	   const SHAPESHIFT = `
	      _____ __                    _____ __    _ ______
	     / ___// /_  ____ _____  ___ / ___// /_  (_) __/ /_
	     \__ \/ __ \/ __ '/ __ \/ _ \\__ \/ __ \/ / /_/ __/
	    ___/ / / / / /_/ / /_/ /  __/__/ / / / / / __/ /_
	   /____/_/ /_/\__,_/ .___/\___/____/_/ /_/_/_/  \__/
	                   /_/
	   `
	*/
	/*
	   const SHAPESHIFT = `
	   ███████╗██╗  ██╗ █████╗ ██████╗ ███████╗███████╗██╗  ██╗██╗███████╗████████╗
	   ██╔════╝██║  ██║██╔══██╗██╔══██╗██╔════╝██╔════╝██║  ██║██║██╔════╝╚══██╔══╝
	   ███████╗███████║███████║██████╔╝█████╗  ███████╗███████║██║█████╗     ██║
	   ╚════██║██╔══██║██╔══██║██╔═══╝ ██╔══╝  ╚════██║██╔══██║██║██╔══╝     ██║
	   ███████║██║  ██║██║  ██║██║     ███████╗███████║██║  ██║██║██║        ██║
	   ╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚══════╝╚══════╝╚═╝  ╚═╝╚═╝╚═╝        ╚═╝
	   `
	*/
