package main

import "fmt"

func main() {
	s := "egg"
	t := "ady"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		fmt.Println("sChar:", string(sChar), "tChar:", string(tChar))

		if val, ok := sMap[sChar]; ok {
			fmt.Println(string(val), string(tChar))
			if val != tChar {
				fmt.Println(val, tChar, 3)
				return false
			}
		} else {
			fmt.Println(i)
			sMap[sChar] = tChar
		}
		fmt.Println("SMAP:", sMap)

		if val, ok := tMap[tChar]; ok {

			if val != sChar {
				return false
			}
		} else {
			fmt.Println(i)
			tMap[tChar] = sChar
		}
		fmt.Println("TMAP:", tMap)

	}

	return true
}
