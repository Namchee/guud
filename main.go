package main

import (
	"fmt"
	"time"

	"github.com/goodsign/monday"
)

func main() {
	// All date is presented in 27 May 2023

	fullDate := "27-05-2023"

	ddmmyyDate := "27/05/2023"
	muricanDate := "05/27/2023"

	shortened := "27 May 2023"
	local := "27 Mei 2023"

	// format test
	idLang := monday.Format(time.Now(), "Monday, 2 January 2006", monday.LocaleIdID)

	fmt.Println(idLang)

	// parse test

	// this should be equal
	baseID, _ := monday.Parse("2-1-2006", fullDate, monday.LocaleIdID)
	baseEN, _ := monday.Parse("2-1-2006", fullDate, monday.LocaleEnUS)

	fmt.Println(baseID == baseEN)

	// this should also be equal, provided that the format is correct

	drunkID, _ := monday.Parse("2-1-2006", ddmmyyDate, monday.LocaleIdID)
	drunkEN, _ := monday.Parse("1-2-2006", muricanDate, monday.LocaleEnUS)

	fmt.Println(drunkID == drunkEN)

	// this should fail, as americans are drunk on the based date format
	_, err := monday.Parse("2-1-2006", muricanDate, monday.LocaleEnUS)
	fmt.Println(err != nil)

	// this should also be same
	shortID, _ := monday.Parse("2 January 2006", local, monday.LocaleIdID)
	shortEN, _ := monday.Parse("2 January 2006", shortened, monday.LocaleEnUS)

	fmt.Println(shortEN == shortID)
}
