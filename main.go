package main

import (
	"github.com/rromulos/aps-finder/helpers"
)

func main() {
	var prefix = helpers.WhatProject()
	println(prefix)
	helpers.PerformAnalysis("app", ".php", prefix)
	// for _, s := range helpers.PerformAnalysis("app", ".php", prefix) {
	// 	println(s)
	// }
}
