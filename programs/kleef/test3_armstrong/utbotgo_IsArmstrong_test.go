package armstrong

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

type utbotgo_type_args_IsArmstrong struct {
	An int `json:"n"`
}

var utbotgo_args_IsArmstrong = []utbotgo_type_args_IsArmstrong{
	{ An: 0 },
	{ An: 1 },
	{ An: 10 },
	{ An: 796 },
	{ An: 9884 },
	{ An: 32604 },
	{ An: 824844 },
	{ An: 5331324 },
	{ An: 69750788 },
	{ An: 867630338 },
	{ An: 2937676874 },
	{ An: 13947401994 },
	{ An: 524345888162 },
	{ An: 4725852160492 },
	{ An: 87341710785132 },
	{ An: 253322479619842 },
	{ An: 1014791672863372 },
	{ An: 73139181584118402 },
	{ An: 540477418035203140 },
	{ An: 1081174791796680714 },
}

type utbotgo_type_results_IsArmstrong struct {
	R bool `json:"R"`
}

var utbotgo_answers_IsArmstrong = []utbotgo_type_results_IsArmstrong{
	
}

func utbotgo_writeResults_IsArmstrong(results []utbotgo_type_results_IsArmstrong) (err error) {
	var data []byte
	data, err = json.Marshal(results)
	if err != nil {
		return
	}
	err = os.WriteFile("utbotgo/functions/IsArmstrong/results.json", data, 0666)
	return
}

func Test_utbotgo_IsArmstrong(t *testing.T) {
	args := utbotgo_args_IsArmstrong
	answers := utbotgo_answers_IsArmstrong
	results := make([]utbotgo_type_results_IsArmstrong, len(args))
	for i := range args {
		t.Run("test_#" + strconv.Itoa(i+1), func(t_ *testing.T) {
			results[i].R = IsArmstrong(args[i].An)
			if i >= len(answers) {
				t_.Logf("\narguments: %v\nexpected is unknown\nactual: %v\n", args[i], results[i])
				t_.Fail()
			} else if answers[i].R != results[i].R {
				t_.Logf("\narguments: %v\nexpected: %v\nactual: %v\n", args[i], answers[i], results[i])
				t_.Fail()
			}
		})
	}
	err := utbotgo_writeResults_IsArmstrong(results)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
