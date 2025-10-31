package lectormail

import (
	"fmt"
	"os"
	"strings"

	"github.com/richardlehane/mscfb"
)

func ListarStream(mailPath string) {
	file, err := os.Open(mailPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := mscfb.New(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Contenido del archivo .msg:")

	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {

		if !strings.HasSuffix(entry.Name, "1F") || entry.Size < 512 {
			continue
		}

		buf := make([]byte, entry.Size)
		doc.Read(buf)

		text := utf16ToString(buf)
		fmt.Printf("[  ] %s : %s\n", entry.Name, text)
		/*
			ft := binary.LittleEndian.Uint64(buf)
			t := filetimeToTime(ft)
			fmt.Printf("  : %s\n", t.Format("2006-01-02 15:04:05"))
		*/
	}
}

/*
// Convierte FILETIME (Windows) a time.Time
func filetimeToTime(ft uint64) time.Time {
	const ticksPerSecond = 10000000
	const epochDiff = 11644473600 // segundos entre 1601 y 1970
	secs := ft / ticksPerSecond
	nanos := (ft % ticksPerSecond) * 100
	return time.Unix(int64(secs-epochDiff), int64(nanos))
}
*/
