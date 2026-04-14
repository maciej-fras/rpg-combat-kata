place to store what you want to present during retrospective session

What problems did you encounter?
What have you learned? Any new technique or pattern?
Share your design with others, and get feedback on different approaches.

My approach
------------------------------------------------------------------------------------
Using AI for learning new language (Go) but ... not solving tasks!

"Hello! 
I would like from you to be my teacher for a couple of next days.
I'm doing a coding kata and I need your help because I want to do this kata in GO which I don't know. What I would like to achieve is that you learn me a bit how to code in golang but at the same time you will not guide me through doing the exercises - that's my job. 
Do you need any more info from me to start?"

A couple of things that would help me calibrate:

1. What languages do you already know? (so I can draw analogies)
2. Do you have Go installed? If not, we should set that up first.
3. Are you planning to do TDD (test-driven development) for this kata? It's common practice and Go has great built-in testing support

Concept	C#/Java	Go
Variable declaration	var x int = 5;	x := 5 (short) or var x int = 5
No semicolons	—	Correct, no semicolons
No inheritance	class B : A	Not possible. Use composition
Interfaces	Explicit implements	Implicit — just implement the methods
Error handling	try/catch	Return error as second value
No generics*	Generics	Go has generics (since 1.18) but you likely won't need them here