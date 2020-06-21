package redis

import (
	"fmt"
	"unicode"
)

/* state
-1 error
0 start
1 action
2 args
3 end
*/

type FSM struct {
	state     int8
	separator rune
	inQuote   bool
	inArg     bool
	action    string
	lastArg   string
	args      []string
	reason    string
}

func Create(sep rune) FSM {
	return FSM{
		0,
		sep,
		false,
		false,
		"",
		"",
		[]string{},
		"",
	}
}

func (fsm *FSM) Parse(p []byte) {
	for _, ch := range p {
		if fsm.state == -1 {
			return
		}
		fsm.process(rune(ch))
	}
	fsm.terminate()
}

func (fsm FSM) Action() string {
	return fsm.action
}

func (fsm FSM) Args() []string {
	return fsm.args
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (fsm *FSM) process(ch rune) {
	switch fsm.state {
	case 0:
		fsm.case0(ch)
	case 1:
		fsm.case1(ch)
	case 2:
		fsm.case2(ch)
	default:
		fsm.caseDefault(ch)
	}
}

func (fsm *FSM) case0(ch rune) {
	if ch != fsm.separator {
		fsm.state = 1
		fsm.action += string(toUpperCase(ch))
	}
}
func (fsm *FSM) case1(ch rune) {
	if ch == fsm.separator {
		if len(fsm.action) == 0 {
			return
		}
		fsm.state = 2
		return
	}
	if isLetter(ch) {
		fsm.action += string(toUpperCase(ch))
		return
	}
	// else, there is an error
	fsm.state = -1
	if isNum(ch) {
		fsm.reason = "no number is allowed in action"
	}
	if isQuote(ch) {
		fsm.reason = "no quote is allowed in action"
	}
}
func (fsm *FSM) case2(ch rune) {
	if ch == fsm.separator {
		if fsm.inQuote {
			fsm.lastArg += string(ch)
			return
		}
		if fsm.inArg {
			fsm.args = append(fsm.args, fsm.lastArg)
			fsm.inArg = false
			fsm.lastArg = ""
			return
		}
		if !fsm.inArg || !fsm.inQuote {
			return
		}
	}
	if isQuote(ch) {
		if !fsm.inArg {
			fsm.inArg = true
			fsm.inQuote = true
			return
		}
		if fsm.inArg && fsm.inQuote {
			fsm.inArg = false
			fsm.inQuote = false
			return
		}
		if fsm.inArg && !fsm.inQuote {
			fsm.lastArg += string(ch)
			return
		}
	}
	fsm.inArg = true
	fsm.lastArg += string(ch)
	return
}
func (fsm *FSM) caseDefault(ch rune) {
	fmt.Print("Unexpected invoking by", ch)
}
func (fsm *FSM) terminate() {
	fsm.args = append(fsm.args, fsm.lastArg)
	fsm.inArg = false
	fsm.lastArg = ""
	fsm.state = 3
}
func isQuote(ch rune) bool {
	return ch == '\'' || ch == '"'
}

func isAlphaNum(ch rune) bool {
	return isLetter(ch) || isNum(ch)
}

func isNum(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func toLowerCase(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return unicode.ToLower(ch)
	}
	return ch
}
func toUpperCase(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return unicode.ToUpper(ch)
	}
	return ch
}
