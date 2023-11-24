package main

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

type utbotgo_type_args_goroutine struct {
	Aa int `json:"a"`
}

var utbotgo_args_goroutine = []utbotgo_type_args_goroutine{
	{ Aa: 0 },
}

type utbotgo_type_results_goroutine struct {
	R int `json:"R"`
}

var utbotgo_answers_goroutine = []utbotgo_type_results_goroutine{
	
}

func utbotgo_writeResults_goroutine(results []utbotgo_type_results_goroutine) (err error) {
	var data []byte
	data, err = json.Marshal(results)
	if err != nil {
		return
	}
	err = os.WriteFile("utbotgo/functions/goroutine/results.json", data, 0666)
	return
}

func Test_utbotgo_goroutine(t *testing.T) {
	args := utbotgo_args_goroutine
	answers := utbotgo_answers_goroutine
	results := make([]utbotgo_type_results_goroutine, len(args))
	for i := range args {
		t.Run("test_#" + strconv.Itoa(i+1), func(t_ *testing.T) {
			results[i].R = goroutine(args[i].Aa)
			if i >= len(answers) {
				t_.Logf("\narguments: %v\nexpected is unknown\nactual: %v\n", args[i], results[i])
				t_.Fail()
			} else if answers[i].R != results[i].R {
				t_.Logf("\narguments: %v\nexpected: %v\nactual: %v\n", args[i], answers[i], results[i])
				t_.Fail()
			}
		})
	}
	err := utbotgo_writeResults_goroutine(results)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
