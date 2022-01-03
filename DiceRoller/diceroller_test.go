package DiceRoller

import (
	"testing"
)

//Test_roll - tests roll function
func Test_roll(t *testing.T) {
	type args struct {
		rollsCount int
		diceSize   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_roll",
			args: args{
				rollsCount: 5,
				diceSize:   6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roll(tt.args.rollsCount, tt.args.diceSize)
		})
	}
}

//TestIsValidDiceString - tests input is valid dice string or not
func TestIsValidDiceString(t *testing.T) {
	type args struct {
		diceInput string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsValidDiceString",
			args: args{diceInput: "3d6+5d10+10d20"},
			want: true,
		},
		{
			name: "TestIsValidDiceString",
			args: args{diceInput: "5d6+8d10"},
			want: true,
		},
		{
			name: "TestIsValidDiceString",
			args: args{diceInput: "4d6"},
			want: true,
		},
		{
			name: "TestIsValidDiceString",
			args: args{diceInput: "4d6+5"},
			want: true,
		},
		{
			name: "TestIsValidDiceString",
			args: args{diceInput: "5d6"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidDiceString(tt.args.diceInput); got != tt.want {
				t.Errorf("IsValidDiceString() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TestSumRollValue - tests sum of roll values
func TestSumRollValue(t *testing.T) {
	type args struct {
		rolls         []string
		ResultChannel chan int
	}
	var channel chan int
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSumRollValue",
			args: args{
				rolls:         []string{"5d10", "3d6", "12d20", "4d10"},
				ResultChannel: channel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sumRollValue(tt.args.rolls, tt.args.ResultChannel)
		})
	}
}

//Test_RollDices - tests the sum of roll dice values by taking dice roller input string
func TestRollDices(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestRollDices",
			args: args{
				inputString: "3d6+5d10",
			},
		},
		{
			name: "TestRollDices",
			args: args{
				inputString: "2d6+6d10+12d20",
			},
		},
		{
			name: "TestRollDices",
			args: args{
				inputString: "5d10+9",
			},
		},
		{
			name: "TestRollDices",
			args: args{
				inputString: "4d20",
			},
		},
		{
			name: "TestRollDices",
			args: args{
				inputString: "4d6+2d6",
			},
		},
		{
			name: "TestRollDices",
			args: args{
				inputString: "4d+23",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RollDices(tt.args.inputString)
		})
	}
}




