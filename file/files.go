package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// buf := make([]byte, 10)
	// d, erro := f.Read(buf)
	// if erro != nil {
	// 	panic(erro)
	// }
	// for i := range d {
	// 	fmt.Println(string(buf[i]))
	// }
	// fmt.Println(buf, d)

	//for small files
	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	//read directory
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// fileInfo, err := dir.ReadDir(100)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//create file
	// fi, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer fi.Close()
	// bytes := []byte("Hello golang")
	// fi.Write(bytes)
	// fi.WriteString("hi go")
	// fi.WriteString("Hiiiiiii go")

	//read and write usinf sream fashion
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()
	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("Written to new file successfully!")

	//delecting file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")

}
