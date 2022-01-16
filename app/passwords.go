package app

import (
	"math/rand"
	"time"
)

func Generate(length int, disable_special bool) string {
	small_set := "abcdefghijklmnopqrstuvwxyz"
	capital_set := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	number_set := "01234567890"
	special_set := "!@#$%^&*()_+"

	base_set := small_set + capital_set + number_set

	var final_set string

	if disable_special {
		final_set = base_set
	} else {
		final_set = base_set + special_set
	}

	rand.Seed(time.Now().Unix())

	var new_pass string
	for x := 0; x < length; x++ {
		new_pass = new_pass + string(final_set[rand.Intn(len(final_set))])
	}

	return new_pass
}
