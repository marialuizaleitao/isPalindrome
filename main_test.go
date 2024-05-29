package main

import "testing"

func Test_invertText(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "PalindromeOddLength",
			args: args{s: "ana"},
			want: "ana",
		},
		{
			name: "NonPalindrome",
			args: args{s: "Maria"},
			want: "airaM",
		},
		{
			name: "EmptyString",
			args: args{s: ""},
			want: "",
		},
		{
			name: "SingleCharacter",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Spaces",
			args: args{s: "hello world"},
			want: "dlrow olleh",
		},
		{
			name: "SpecialCharacters",
			args: args{s: "hello!"},
			want: "!olleh",
		},
		{
			name: "NumericCharacters",
			args: args{s: "12345"},
			want: "54321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertText(tt.args.s); got != tt.want {
				t.Errorf("invertText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "EmptyString",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "SingleCharacter",
			args: args{
				s: "a",
			},
			want: true,
		},
		{
			name: "TwoCharactersPalindrome",
			args: args{
				s: "aa",
			},
			want: true,
		},
		{
			name: "TwoCharactersNonPalindrome",
			args: args{
				s: "ab",
			},
			want: false,
		},
		{
			name: "PalindromeOddLength",
			args: args{
				s: "ana",
			},
			want: true,
		},
		{
			name: "PalindromeEvenLength",
			args: args{
				s: "abba",
			},
			want: true,
		},
		{
			name: "NonPalindrome",
			args: args{
				s: "hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
