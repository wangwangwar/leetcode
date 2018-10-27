package main

import (
	"fmt"
	"os"

	stack "github.com/golang-collections/collections/stack"
)

func pushInt(stack *stack.Stack, elem int) {
	peek := stack.Peek()
	if peek == nil {
		stack.Push(elem)
		return
	}

	peekInt, ok := peek.(int)
	if ok {
		stack.Pop()
		pushInt(stack, peekInt+elem)
		return
	}

	stack.Push(elem)
}

func pushRune(stack *stack.Stack, elem rune) {
	if elem == '(' {
		stack.Push(elem)
		return
	}

	// elem is ')'
	peek := stack.Peek()
	if peek == nil {
		os.Exit(-1)
	}
	peekRune, ok := peek.(rune)
	if ok && peekRune == '(' {
		stack.Pop()
		pushInt(stack, 1)
		return
	}

	// if peek is int, double it
	peekInt, ok := peek.(int)
	if ok {
		stack.Pop()
		stack.Pop()
		pushInt(stack, peekInt*2)
	}
}

func scoreOfParentheses(S string) int {
	runeStack := stack.New()
	for i := 0; i < len(S); i++ {
		pushRune(runeStack, rune(S[i]))
	}
	return runeStack.Pop().(int)
}

func main() {
	score1 := scoreOfParentheses("()")
	fmt.Println(score1)
	score2 := scoreOfParentheses("(())")
	fmt.Println(score2)
	score3 := scoreOfParentheses("()()")
	fmt.Println(score3)
	score4 := scoreOfParentheses("(())()")
	fmt.Println(score4)
	score5 := scoreOfParentheses("(())(())")
	fmt.Println(score5)
}
