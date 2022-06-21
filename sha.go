package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	check := "0000"
	var ty string
	fmt.Scan(&ty)

	for i := 0; ; i++ {

		str := ty + strconv.Itoa(i)
		h := sha256.New()
		h.Write([]byte(str))
		sha := hex.EncodeToString(h.Sum(nil))
		checkk := sha[:4]

		if check == checkk {
			fmt.Println(" the added digits is  ", i, " in ", str, " gets 0000 in sha")
			fmt.Println("sha=", sha)
			return
		}

	}

}
