package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func Over_write(file_name string, bin []uint8) { //bin = nil (none) if file creation
	err := os.WriteFile(file_name, bin, 0660)

	if err != nil {
		fmt.Println("error writing to file")
	}
}

func Bite() []uint8 { //includes pretty print via MarshallIndent
	data := map[string]interface{}{
		"CHECK": map[string]interface{}{
			"desc":        "test run 1",
			"status":      "com",
			"taskID":      "sampleID",
			"timeCreated": "sample time",
			"timeUpdated": "sample time",
		},
	} //close
	json_data, _ := json.MarshalIndent(data, "", "    ")

	return json_data
}

func Two_inputs() (string, string) {
	var key string
	var pair string
	fmt.Print("Key: ")
	fmt.Scan(&key)
	fmt.Print("Pair: ")
	fmt.Scan(&pair)
	return key, pair
}

// ONE) ADD
func Bite_one(a string, b string, c string) []uint8 {
	task_name, description, status := a, b, c

	data := map[string]interface{}{
		task_name: map[string]interface{}{
			"desc":        description,
			"status":      status,
			"taskID":      "sampleID",
			"timeCreated": time.TimeOnly,
			"timeUpdated": time.TimeOnly,
		},
	} //close
	json_data, _ := json.MarshalIndent(data, "", "    ")

	return json_data
}
func In_bite_one() (string, string, string) {
	var taskname string
	var desc string
	var status int

	vstatus_map := map[int]string{
		1: "COM",
		2: "IPR",
		3: "INC",
	}

	fmt.Println("name of task?: ")
	fmt.Scan(&taskname)
	fmt.Println("write a descripion for the task: ")
	fmt.Scan(&desc)
	print(`What is the status of this task?:
                       1) Completed
                       2) In Progress
                       3) Incomplete`)
	fmt.Scan(&status)

	return taskname, desc, vstatus_map[status]

}

// IDKKK PROLLY ASS
func Byte_str(format_str string) []uint8 {
	type Message struct {
		key_pair string
	}

	m := Message{format_str}

	b, _ := json.Marshal(m)
	return b
}
