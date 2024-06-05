package main
import "fmt"

type userAtt struct {
	id			int
	nama		string
	alamat		string
	npwp		int
	tier		string
	harta		int
	aset		int
}

const NMAX int = 100
var usersData = [NMAX]userAtt{
	{1, "A", "Jl.A", 123, "MNWI", 1000000, 1000000},
	{2, "B", "Jl.B", 456, "UHNWI", 2000000, 2000000},
	{3, "C", "Jl.C", 789, "MNWI", 3000000, 3000000},
	{4, "D", "Jl.D", 101, "HNWI", 4000000, 4000000},
	{5, "E", "Jl.E", 112, "UHNWI", 5000000, 5000000}}

func main(){
	var optionInput, deleteInput, tax int
	var searchInput, editInput string
	
	fmt.Println("+----------------------+")
	fmt.Println("|     LAPOR PAJAK      |")
	fmt.Println("+----------------------+")
	fmt.Println("| 1. Tampilkan Data    |")
	fmt.Println("| 2. Cari Data (Nama)  |")
	fmt.Println("| 3. Tambah Data       |")
	fmt.Println("| 4. Edit Data (Nama)  |")
	fmt.Println("| 5. Hapus Data (ID)   |")
	fmt.Println("| 6. Hitung Pajak      |")
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
		exitValidation()
		
	} else if optionInput == 2 {
		fmt.Println("| 2. Cari Data (Nama)  |")
		fmt.Print("| Cari Pengguna: ")
		fmt.Scanln(&searchInput)
		fmt.Println("+----------------------+")
		SearchData(usersData, searchInput)
		exitValidation()
	} else if optionInput == 3 {
		fmt.Println("| 3. Tambah Data       |")
		fmt.Println("+----------------------+")
		for i := 0; i < len(usersData); i++ {
			if usersData[i].id == 0 {
				usersData[i].id = i+1
				fmt.Print("Nama: ")
				fmt.Scanln(&usersData[i].nama)
				fmt.Print("Alamat: ")
				fmt.Scanln(&usersData[i].alamat)
				fmt.Print("NPWP: ")
				fmt.Scanln(&usersData[i].npwp)
				fmt.Print("Tier: ")
				fmt.Scanln(&usersData[i].tier)
				fmt.Print("Harta: ")
				fmt.Scanln(&usersData[i].harta)
				fmt.Print("Aset: ")
				fmt.Scanln(&usersData[i].aset)
				break
			}
		}
		exitValidation()
	} else if optionInput == 4 {
		fmt.Println("| 4. Edit Data (Nama)  |")
		fmt.Print("| Edit Pengguna: ")
		fmt.Scanln(&editInput)
		fmt.Println("+----------------------+")
		EditData(&usersData, editInput)
		exitValidation()
	} else if optionInput == 5 {
		fmt.Println("| 5. Hapus Data (ID)   |")
		fmt.Print("| Hapus Pengguna: ")
		fmt.Scanln(&deleteInput)
		fmt.Println("+----------------------+")
		if deleteInput == 0 {
			fmt.Println("| Delete Failed!       |")
		} else {
			DeleteData(&usersData, deleteInput)
		}
		exitValidation()
	} else if optionInput == 6 {
		fmt.Println("| 6. Hitung Pajak      |")
		fmt.Print("| ID Pengguna: ")
		fmt.Scanln(&tax)
		fmt.Println("+----------------------+")
		if usersData[tax].id == 0{
			fmt.Println("| ID Tidak Ditemukan   |")
		} else {
			annualTax(usersData, tax)
		}
		exitValidation()
	} else {
		fmt.Println("| Opsi Tidak Tersedia  |")
		exitValidation()
	}
}

func ShowData(arr [100]userAtt){
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

func Ascending(arr [100]userAtt){
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].harta > arr[j+1].harta {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	ShowData(arr)
}

func Descending(arr [100]userAtt){
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].harta < arr[j+1].harta {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	ShowData(arr)
}

func SearchData(arr [100]userAtt, search string){
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

func EditData(arr *[100]userAtt, edit string){
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
			fmt.Println("Edit Data")
			fmt.Print("Nama: ")
			fmt.Scanln(&arr[i].nama)
			fmt.Print("Alamat: ")
			fmt.Scanln(&arr[i].alamat)
			fmt.Print("NPWP: ")
			fmt.Scanln(&arr[i].npwp)
			fmt.Print("Tier: ")
			fmt.Scanln(&arr[i].tier)
			fmt.Print("Harta: ")
			fmt.Scanln(&arr[i].harta)
			fmt.Print("Aset: ")
			fmt.Scanln(&arr[i].aset)
			fmt.Println()
		}
	}
}

func DeleteData(arr *[100]userAtt, delete int){
	//sort ascending
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j].harta > arr[j+1].harta {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	var left, mid, right int
	left = 0
	right = len(arr)-1
	mid = right/2

	for left<=right && arr[mid].id!=delete {
		if delete<arr[mid].id {
			right = mid-1 
		} else {
			left = mid+1
		}
		mid = (left+right)/2
	}
	
	if arr[mid].id==delete {
		for i:=mid; i<len(arr); i++{
			if i+1 <= len(arr)-1 {
				arr[i] = arr[i+1]
			} else {
				arr[i].id = 0
				arr[i].nama = ""
				arr[i].alamat = ""
				arr[i].npwp = 0
				arr[i].harta = 0
				arr[i].aset = 0
			}
		}
		fmt.Println("| Delete Success!      |")
	} else {
		fmt.Println("| Delete Failed!       |")
	}
}

func exitValidation(){
	var exitOption string
	fmt.Println("+----------------------+")
	fmt.Print("| Keluar (y/n): ")
	fmt.Scanln(&exitOption)
	fmt.Println("+----------------------+")
	if exitOption == "y" {
		fmt.Println("| Keluar               |")
		fmt.Println("+----------------------+")
	} else if exitOption == "n"{
		fmt.Println("| Kembali ke Menu      |")
		fmt.Println("+----------------------+")
		main()
	}
}

/*
	MNWI - Mass Net Worth Individual 0M-1M, tax 0%
	UHNWI - Ultra High Net Worth Individual 1M-100M, tax 10%
	HNWI - High Net Worth Individual 100M-1B, tax 20%
*/
func annualTax(arr [100]userAtt, tax int){
	var intConvert float64
	for i := 0; i < len(arr); i++ {
		if tax == arr[i].id {
			if arr[i].tier == "MNWI" {
				fmt.Println("ID:", arr[i].id)
				fmt.Println("Nama:", arr[i].nama)
				fmt.Println("Tier:", arr[i].tier)
				fmt.Println("Harta:", arr[i].harta)
				fmt.Println("Pajak:", 0)
			} else if arr[i].tier == "UHNWI" {
				intConvert = float64(arr[i].harta)
				fmt.Println("ID:", arr[i].id)
				fmt.Println("Nama:", arr[i].nama)
				fmt.Println("Tier:", arr[i].tier)
				fmt.Println("Harta:", arr[i].harta)
				fmt.Println("Pajak:", intConvert*0.1)
			} else if arr[i].tier == "HNWI" {
				intConvert = float64(arr[i].harta)
				fmt.Println("ID:", arr[i].id)
				fmt.Println("Nama:", arr[i].nama)
				fmt.Println("Tier:", arr[i].tier)
				fmt.Println("Harta:", arr[i].harta)
				fmt.Println("Pajak:", intConvert*0.2)
			}  else {
				fmt.Println("| Tier Tidak Ditemukan |")
			}
		} 
	}
}
