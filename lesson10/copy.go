package lesson10

import (
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
	"time"
)

func Copy(from string, to string, limit int, offset int) error {
	bufSize := 5
	file, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fileTo, err := os.Create(to)
	defer fileTo.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return err
	}

	lOffset := offset

	var bLen int
	if limit <= int(fi.Size())-offset {
		bLen = int(fi.Size()) - offset
	} else {
		bLen = limit - offset
	}
	buf := make([]byte, bufSize)
	i := 1
	//var percent float32
	for lOffset < (int(fi.Size()) - offset) {
		count := bLen
		bar := pb.StartNew(count)
		i++
		read, err := file.ReadAt(buf, int64(lOffset))
		fileTo.WriteAt(buf, int64(lOffset))

		lOffset += read
		if err == io.EOF {
			bar.Finish()
			break
		}
		bar.Increment()
		time.Sleep(time.Millisecond)
	}

	/*	percent = float32(len(buf)*i)/float32(bLen) * 100
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("Прогресс=%.0f%%\n",percent)
	*/
	if err != nil {
		log.Printf("failed to read: %v", err)
		return err
	}

	return nil
}
