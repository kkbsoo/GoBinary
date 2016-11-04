package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"xutil"
	"sort"
)

func main(){
	sFilePath, iTailFlag , err := xutil.GetFileArgs(os.Args)

	file, err := os.Open(sFilePath)
	defer file.Close()
	xutil.ErrorNilCheck(err)

	for {
		wtmp := new(xutil.Utmp)

		if iTailFlag < 1{
			err := binary.Read(file, binary.LittleEndian, wtmp)
			 if xutil.ErrorEOFCheck(err) < 0 {
				 break;
			 }

			mapWtmp := xutil.UtmpMap(wtmp)

			//map sort - 그냥 출력하면 순서없이 출력됨
			var keys []string
			for k := range  mapWtmp{
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, k := range  keys {
				fmt.Printf("%s = ", k)
				fmt.Print(mapWtmp[k]," | ")
			}

			fmt.Printf("\n")
		}else if iTailFlag == 1{

		}
	}
}