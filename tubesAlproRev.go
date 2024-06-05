package main
import "fmt"

type userAtt struct {
	id			int
	nama		string
	alamat		string
	npwp		string
	tier		string
	harta		int
	aset		int
}

func main(){
	var usersData = [10]userAtt{
		{1, "A", "Jl. A", "123", "MNWI", 1000000, 1000000},
		{2, "B", "Jl. B", "456", "UHNWI", 2000000, 2000000},
		{3, "C", "Jl. C", "789", "MNWI", 3000000, 3000000},
		{4, "D", "Jl. D", "101", "HNWI", 4000000, 4000000},
		{5, "E", "Jl. E", "112", "UHNWI", 5000000, 5000000}}
	var optionInput, deleteInput int
	var searchInput, editInput, exitOption string
	
	fmt.Println("+----------------------+")
	fmt.Println("|     LAPOR PAJAK      |")
	fmt.Println("+----------------------+")
	fmt.Println("| 1. Tampilkan Data    |")
	fmt.Println("| 2. Cari Data (Nama)  |")
	fmt.Println("| 3. Edit Data (Nama)  |")
	fmt.Println("| 4. Hapus Data (ID)   |")
	fmt.Println("| 5. Hitung Pajak      |")
	fmt.Println("+----------------------+")
	fmt.Print("| Pilih Opsi -> ")
	fmt.Scanln(&optionInput)
	fmt.Println("+----------------------+")

	if optionInput == 1 {
		fmt.Println("| 1. Tampilkan Data    |")
		fmt.Println("+----------------------+")
		ShowData(usersData)
		fmt.Println("+----------------------+")
		fmt.Println("| 2. Ascending         |")
		fmt.Println("+----------------------+")
		Ascending(usersData)
		fmt.Println("+----------------------+")
		fmt.Println("| 3. Descending        |")
		fmt.Println("+----------------------+")
		Descending(usersData)
		fmt.Print("| Keluar (y/n): ")
		fmt.Scanln(&exitOption)
		fmt.Println("+----------------------+")
		if exitOption == "y" {
			fmt.Println("| Keluar               |")
			fmt.Println("+----------------------+")
		} else {
			fmt.Println("| Kembali ke Menu      |")
			fmt.Println("+----------------------+")
			main()
		}
	} else if optionInput == 2 {
		fmt.Println("| 2. Cari Data (Nama)  |")
		fmt.Print("| Cari Pengguna: ")
		fmt.Scanln(&searchInput)
		fmt.Println("+----------------------+")
		SearchData(usersData, searchInput)
		// ShowData(arr)
		// fmt.Println("
		fmt.Println("+----------------------+")
	} else if optionInput == 3 {
		fmt.Println("| 3. Edit Data (Nama)  |")
		fmt.Print("| Edit Pengguna: ")
		fmt.Scanln(&editInput)
		fmt.Println("+----------------------+")
		EditData(usersData, editInput)
		// ShowData(arr)
		// fmt.Println("
		fmt.Println("+----------------------+")
	} else if optionInput == 4 {
		fmt.Println("| 4. Hapus Data (ID)   |")
		fmt.Print("| Hapus Pengguna: ")
		fmt.Scanln(&deleteInput)
		// ShowData(arr)
		// fmt.Println("
		fmt.Println("+----------------------+")
	} else if optionInput == 5 {
		fmt.Println("| 5. Hitung Pajak      |")
		// ShowData(arr)
		// fmt.Println("
		fmt.Println("+----------------------+")
	}
}

func ShowData(arr [10]userAtt){
	for i := 0; i < len(arr); i++ {
		if arr[i].id != 0 {
			fmt.Println("ID:", arr[i].id)
			fmt.Println("Nama:", arr[i].nama)
			fmt.Println("Alamat:", arr[i].alamat)
			fmt.Println("NPWP:", arr[i].npwp)
			fmt.Println("Tier:", arr[i].tier)
			fmt.Println("Harta:", arr[i].harta)
			fmt.Println("Aset:", arr[i].aset)
			fmt.Println()
		}
	}
}

func Ascending(arr [10]userAtt){
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].harta > arr[j+1].harta {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	ShowData(arr)
}

func Descending(arr [10]userAtt){
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].harta < arr[j+1].harta {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	ShowData(arr)
}

func SearchData(arr [10]userAtt, search string){
	for i := 0; i < len(arr); i++ {
		if search == arr[i].nama {
			fmt.Println("ID:", arr[i].id)
			fmt.Println("Nama:", arr[i].nama)
			fmt.Println("Alamat:", arr[i].alamat)
			fmt.Println("NPWP:", arr[i].npwp)
			fmt.Println("Tier:", arr[i].tier)
			fmt.Println("Harta:", arr[i].harta)
			fmt.Println("Aset:", arr[i].aset)
			fmt.Println()
		}
	}
}

func EditData(arr [10]userAtt, edit string){
	for i := 0; i < len(arr); i++ {
		if edit == arr[i].nama {
			fmt.Println("ID:", arr[i].id)
			fmt.Println("Nama:", arr[i].nama)
			fmt.Println("Alamat:", arr[i].alamat)
			fmt.Println("NPWP:", arr[i].npwp)
			fmt.Println("Tier:", arr[i].tier)
			fmt.Println("Harta:", arr[i].harta)
			fmt.Println("Aset:", arr[i].aset)
			fmt.Println()
		}
	}
}

// func ShowData(arr [5]userAtt){