package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/rromulos/aps-finder/helpers"
	"github.com/rromulos/aps-finder/helpers/logger"
)

func main() {
	logger.InitLogs()
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	var prefix = helpers.WhatProject()
	println(prefix)
	helpers.PerformAnalysis("app", ".php", prefix)

	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
	// for _, s := range helpers.PerformAnalysis("app", ".php", prefix) {
	// 	println(s)
	// }
}
