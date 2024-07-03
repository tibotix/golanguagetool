package main

import "github.com/fatih/color"

var backgroundColor = color.New(color.FgHiBlack)
var redColor = color.New(color.FgRed)
var greenColor = color.New(color.FgGreen)

var backgroundText = backgroundColor.SprintFunc()
var redText = redColor.SprintFunc()
var greenText = greenColor.SprintFunc()
var tick = greenText("\u2713")
var cross = redText("\u2717")
