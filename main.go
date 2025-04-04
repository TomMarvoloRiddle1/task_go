package main

import (
	"fmt"
	"os"
	"tt/lib.go"
)

func main() {

	menu()

}

func menu() { //control flow
	var a int
	fmt.Println(`Select action by entering a number as a prompt.
        1) Add task
        2) View Tasks (completed, inprogress and not started)
        3) Update Tasks
        4) Delete Tasks
        5) Exit Program'`)
	fmt.Scan(&a)
	switch a {
	case 1:
		one()
	case 2:
	case 3:
	case 4:
	case 5:
		os.Exit(3)
	default:
		fmt.Println("invalid selection!")
		menu()
	}
}

// all below module2

// func two() {

// 	data := json_to_bite("sample.json")
// 	bite_to_interface(data)
// }

// func bite_to_interface(data []uint8) {

// 	type Set struct {
// 		//
// 		keypairs map[string]interface{} {
// 		"task_name": map[string]interface{} {
// 			"desc":        description,
// 			"status":      status,
// 			"taskID":      "sampleID",
// 			"timeCreated": time.TimeOnly,
// 			"timeUpdated": time.TimeOnly,
// 		},
// 	}	//close

// }
// 	var s []Set
// 	err := json.Unmarshal(data, &s)
// 	fmt.Println(err)

// }

// working mod2

// func json_to_bite(file_name string) []uint8 {
// 	data, _ := os.ReadFile(file_name)
// 	return data
// }

//all above module2

// data := map[string]interface{}{
// 		task_name: map[string]interface{}{
// 			"desc":        description,
// 			"status":      status,
// 			"taskID":      "sampleID",
// 			"timeCreated": time.TimeOnly,
// 			"timeUpdated": time.TimeOnly,
// 		},

// all below module1
func one() {
	a, b, c := lib.In_bite_one()
	data := lib.Bite_one(a, b, c)
	lib.Over_write("new.json", data)

	// exit flow
	end := end_selection() // 1 = one(), 2 = menu()
	end_cycle(end, one)
}
func end_selection() int {
	var in int
	fmt.Println(`Would you like to repeat this action?
              1) Yes
              2) Return to menu`)
	fmt.Scan(&in)
	return in
}
func end_cycle(num int, f func()) {
	switch num {
	case 1:
		f()
	case 2:
		main()
	default:
		fmt.Println("invalid selection")
		end := end_selection() // 1 = one(), 2 = menu()
		end_cycle(end, one)
	}
}

//all above module 1
