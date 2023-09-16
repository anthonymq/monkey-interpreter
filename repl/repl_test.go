package repl

import (
	"bytes"
	"testing"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name    string
		input   *bytes.Buffer
		wantOut string
	}{
		{
			"OneLine",
			bytes.NewBufferString("let test = 1;\n"),
			`🦍>> {Type:LET Literal:let}
{Type:IDENT Literal:test}
{Type:= Literal:=}
{Type:INT Literal:1}
{Type:; Literal:;}
🦍>> `,
		},
		{
			"TwoLines",
			bytes.NewBufferString("let one = 1;\n let two = 2;"),
			`🦍>> {Type:LET Literal:let}
{Type:IDENT Literal:one}
{Type:= Literal:=}
{Type:INT Literal:1}
{Type:; Literal:;}
🦍>> {Type:LET Literal:let}
{Type:IDENT Literal:two}
{Type:= Literal:=}
{Type:INT Literal:2}
{Type:; Literal:;}
🦍>> `,
		},
		{
			"Illegal char",
			bytes.NewBufferString("@@@\n"),
			`🦍>> {Type:ILLEGAL Literal:@}
{Type:ILLEGAL Literal:@}
{Type:ILLEGAL Literal:@}
🦍>> `,
		},
		{
			"Exit",
			bytes.NewBufferString("exit\n"),
			`🦍>> Harambe says goodbye`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			Start(tt.input, &output)
			if gotOut := output.String(); gotOut != tt.wantOut {
				t.Errorf("Start() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
