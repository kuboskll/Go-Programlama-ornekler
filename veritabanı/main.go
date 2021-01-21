package main

func main() {
	vt, _ := sql.Open("sqlite3", "./veritabanı.db") //veri tabanı dosyamız
	//Veri Silme
	işlem, _ := vt.Prepare("delete from kisiler where id=?")
 	veri, _ := işlem.Exec(id) //Silinecek id
 	değişiklik, _ := veri.RowsAffected() //Silinen kişinin id'sini aldık
 	fmt.Println("Silinen kişinin id'si: ", değişiklik)
 	vt.Close() //veri tabanımızı kapatıyoruz
}
