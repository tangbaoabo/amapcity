package fetcher

import (
	"fmt"
	"testing"
)

func TestFetcher(t *testing.T) {
	contents, err := Fetcher("https://restapi.amap.com/v3/config/district?key=d920349477ba502b313c354d564335b4&keywords=640381&subdistrict=1&filter=640381")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))

}
