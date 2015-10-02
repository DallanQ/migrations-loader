package main

func main() {
    flag.Parse()
   	if *typ != "i" && *typ != "e" {
   		flag.Usage()
   		os.Exit(1)
   	}

   	fmt.Println("Hello world fs2")

}
