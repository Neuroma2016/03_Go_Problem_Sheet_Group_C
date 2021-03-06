//Andrew Usevas group-c
//Regular expressions

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//Eliza responses
func ElizaResponse(input string) string {
	//if word "father" used in users input + Eliza's response
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?"
	}

	//if string begins with variant of "i am",capture rest of the string and
	//concatenate that to your response
	re := regexp.MustCompile(`(?i)I am ([^.?!]*)([.?!]?)`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "I am $1 too$2")
	}

	//random Eliza responses
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	return answers[rand.Intn(len(answers))]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//1.
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("I'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()
	//3.
	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()
}
