package main

func Checkerror(err error) { // handling error in a separate function
	if err != nil {
		panic(err)
	}
}

func main() {
	// Creating a directory :
	// err := os.Mkdir("subdir", 0755)
	// Checkerror(err)

	// // Creating a file inside directory :
	// err = os.WriteFile("subdir/file.txt", []byte("Hello Subdir"), 0755)
	// Checkerror(err)

	// Creating a nested directory :
	// Checkerror(os.MkdirAll("Grandparent/Parent/Child", 0755))
	// Checkerror(os.MkdirAll("Grandparent/Parent/Child1", 0755)) // parent has multiple childrens
	// Checkerror(os.MkdirAll("Grandparent/Parent/Child2", 0755))
	// Checkerror(os.MkdirAll("Grandparent/Parent/Child3", 0755))

	// Creating files for all the childrens in directories
	// Checkerror(os.WriteFile("Grandparent/Parent/Child/file.txt", []byte("Hello child"), 0755))
	// Checkerror(os.WriteFile("Grandparent/Parent/Child1/file1.txt", []byte("Hello child1"), 0755))
	// Checkerror(os.WriteFile("Grandparent/Parent/Child2/file2.txt", []byte("Hello child2"), 0755))
	// Checkerror(os.WriteFile("Grandparent/Parent/Child3/file3.txt", []byte("Hello child3"), 0755))

	// Reading all the directories :
	// dir, err := os.ReadDir("Grandparent/Parent")
	// Checkerror(err)
	// for _, v := range dir {
	// 	// fmt.Println(v) // returns the directories
	// 	// fmt.Println(v.Name()) // returns the name of the directories
	// 	// fmt.Println(v.Type()) // returns d--------- (directory)
	// 	fmt.Println(v.IsDir()) // returns bool True if dir; false if not dir
	// }

	// Changing the directory :
	// Checkerror(os.Chdir("Grandparent/Parent")) // changes to the current working directory to the named directory.
	// dirs, err := os.ReadDir(".")               // just pass "."
	// Checkerror(err)
	// fmt.Println(dirs)

	// Get the working directory location :
	// before changing the working directory
	// loc, err := os.Getwd()
	// Checkerror(err)
	// fmt.Println(loc) // C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Ultimate Go\Intermediate\32_Directories\Grandparent\Parent

	// After changing the working directory
	// Checkerror(os.Chdir("../../"))
	// loc, err = os.Getwd()
	// fmt.Println(loc) // C:\Users\Shashank.BR\OneDrive\Desktop\Go programing\Ultimate Go\Intermediate\32_Directories

	// Deleting the directory:
	// Checkerror(os.Remove("Grandparent")) // only removes single directory or file
	// Checkerror(os.RemoveAll("Grandparent")) // removes all the directories or file

	// Using file path package for directories :
	// pathfile := "Grandparent/Parent/Child"
	// err := filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
	// 	Checkerror(err)
	// 	fmt.Println(path)
	// 	return nil
	// })
	// Checkerror(err)
}

/*
In Go, file and directory permissions are handled using Unix-style file mode bits, represented as os.FileMode. These permissions control how users can read, write, or execute files and directories.

| Octal | Meaning                                    	 | Owner | Group | Others |
| ----- | ---------------------------------------------- | ----- | ----- | ------ |
| 0644  | File readable by everyone, writable by owner   | rw-   | r--   | r--    |
| 0600  | File readable and writable only by owner       | rw-   | ---   | ---    |
| 0755  | Directory: owner full access, others read/exec | rwx   | r-x   | r-x    |
| 0700  | Directory/file access for owner only           | rwx   | ---   | ---    |

*/
