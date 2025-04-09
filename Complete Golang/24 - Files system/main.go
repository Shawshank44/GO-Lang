package main

import "os"

func main() {

	// read files
	// f, err := os.Open("C:/Users/Shashank.BR/OneDrive/Desktop/Go programing/Complete Golang/24 - Files system/example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())

	// buf := make([]byte, fileInfo.Size())

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("Data", d, string(buf[i]))
	// }

	// println("data", d, buf)

	// I, err := os.ReadFile("C:/Users/Shashank.BR/OneDrive/Desktop/Go programing/Complete Golang/24 - Files system/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(I))

	// read folders
	// dir, err := os.Open("./")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(10)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	// writing a file :
	// f, err := os.Create("Example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// // f.WriteString("Hi Go")
	// // f.WriteString("Nice Language")

	// bytesf := []byte("Hello Go lang")
	// f.Write(bytesf)

	// source, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer source.Close()

	// destfile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destfile.Close()

	// reader := bufio.NewReader(source)
	// writer := bufio.NewWriter(destfile)

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

	// fmt.Println("Written in a new file")

	// deleting file :
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
}
