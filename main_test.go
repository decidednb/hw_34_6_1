package main

import (
	"reflect"
	"testing"
)

func Test_splitMathExpression(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Кейс 1 - выражение не соответсвует шаблону",
			args: args{
				s: "7564.....0123-1=?",
			},
			want:    [][]string{},
			wantErr: true,
		}, {
			name: "Кейс 2 - отрицательные числа",
			args: args{
				s: "-3.14/3.56=?",
			},
			want:    [][]string{{"-3.14/3.56=?", "-3.14", "/", "3.56", "=", "?"}},
			wantErr: false,
		}, {
			name: "Кейс 3 - числа с дробной частью",
			args: args{
				s: "3.14+3.56=?",
			},
			want:    [][]string{{"3.14+3.56=?", "3.14", "+", "3.56", "=", "?"}},
			wantErr: false,
		}, {
			name: "Кейс 4 - строка",
			args: args{
				s: "",
			},
			want:    [][]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitMathExpression(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("%v", got)
				t.Errorf("splitMathExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitMathExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		firstNum  float64
		secondNum float64
		operation string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Провал тестов - деление на 0",
			args: args{
				firstNum:  123.26,
				secondNum: 0.00,
				operation: "/"},
			want:    0.0,
			wantErr: true,
		}, {
			name: "Провал тестов - неизвестный арифметический оператор",
			args: args{
				firstNum:  26,
				secondNum: 0.00,
				operation: "?"},
			want:    0.0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.args.firstNum, tt.args.secondNum, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
