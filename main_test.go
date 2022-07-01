package main

import "testing"

var tests = []struct{
	name string
	dividen float64
	divisor float64
	expected float64
	isErr bool
}{
	{"valid-data",100.0,10.0,10.0,false},
	{"invalid-data",100.0,0.0,0.0,true},
	{"expected-data",50.0,10.0,5.0,false},
	{"expected-data",1000.0,10.0,100.0,false},


}

func TestDivision(t *testing.T) {
	for _ , tt := range tests{
		got, err := Divide(tt.dividen,tt.divisor)
		if tt.isErr{
			if err == nil{
				t.Error("expected an error but didnt get one ")
			}
		}else{
			if err != nil{
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected{
			t.Errorf("expected %f but got %f",tt.expected,got)
		}

	}
}

//func TestDivide(t *testing.T){
//	_, err := Divide(10.0, 1.0)
//	if err != nil{
//		t.Error("got an error we should not have")
//	}
//}
//func TestBadDivide(t *testing.T){
//	_, err := Divide(10.0, 0)
//	if err == nil{
//		t.Error("Did not get an error we should  have")
//	}
//}