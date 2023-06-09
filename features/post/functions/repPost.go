package post

import (
	"fmt"
	postStruct "projek/features/post/structs"

	common "projek/common"
)

func RepPost(tPost *postStruct.TabPost) {
	//Pengguna menginputkan balasan pada postingan
	var IDPost string
	var idx int
	var inputUser string
	var response int

	fmt.Println("Pilih ID postingan yang ingin dibalas: ")
	fmt.Scan(&IDPost)
	idx = SearchPost(*tPost, IDPost)

	if idx != -1 {
		fmt.Println("Balasan anda: ")
		common.InputMultipleString(&inputUser)
		fmt.Println("[1] Kirim          [2]Batal")
		fmt.Print("Action : ")
		fmt.Scan(&response)
		for response == 1 {
			if tPost.ArrPost[idx].Nreply < len(tPost.ArrPost[idx].ArrReply) {
				tPost.ArrPost[idx].ArrReply[tPost.ArrPost[idx].Nreply] = inputUser
				tPost.ArrPost[idx].Nreply++
				fmt.Println("Balasan anda telah diposting!")
			}
			fmt.Println("Balasan anda: ")
			common.InputMultipleString(&inputUser)
			fmt.Println("[1] Balas          [2]Batal")
			fmt.Print("Action : ")
			fmt.Scan(&response)
		}
	} else {
		fmt.Print("postingan tidak ditemukan!")
	}
}
