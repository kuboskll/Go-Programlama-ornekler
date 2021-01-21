package main

import (
	"log"
	"os"
)

//Dosya Silme İşlemi
func main() {
	err := os.Remove("new.txt")
	if err != nil {
		log.Fatal(err)
	}
}

//***************************************************************************
//Dosya oluşturma
/*var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("deneme.txt")
	if err != nil {
		log.Fatal(err)
	}
}*/
//***************************************************************************
/*//Dosya bilgisi Alam (Get File info)
var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	//Dosya bilgisini dödürür. Dosya yoksa hata döner
	fileInfo, err := os.Stat("deneme.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dosya Adı: ", fileInfo.Name())
	fmt.Println("uzunluk: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Son güncelleme tarihi: ", fileInfo.ModTime())
	fmt.Println("Bu yol dizi mi dosya mı: ", fileInfo.IsDir())
	fmt.Println("System ınterface type: ", fileInfo.Sys())
	fmt.Println("Sytem info: \n\n", fileInfo.Sys())

}
//***************************************************************************
//Yeniden isimlendirme ve Taşıma işlemi
func main() {
	orjinalPath := "deneme.txt"
	newPath := "./yeni klasör/demo.txt"
	err := os.Rename(orjinalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
//***************************************************************************
//Dosyayı Açma ve Kapama
func main() {
	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Veri yazma kısmı
	file.Close()
}
//***************************************************************************
func main() {
	//Dosya Yazma
	err := ioutil.WriteFile("new.txt", []byte("Merhaba!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
*/
