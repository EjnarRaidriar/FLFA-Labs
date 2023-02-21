# Intro to formal languages. Regular grammar. Finite Automata.
### Course: Formal Languages & Finite Automata
### Author: Cucoș Emanuil

-----

## Theory

### Formal language:
Any language has a grammar, which is a set of rules how to make a word from that language.  
A grammar contains:

* Vn - set of non-terminal symbols;
* Vt - set of terminal symbols;
* S - set of start symbols;
* P - a finite set of production of rules

### Finite automata:
Automata theory is the study of abstract computational devices
(abstract state machine).
A finite automaton is a mathematical model of a finite-state machine. 
The state machine has a set of input symbols and jumps through the set of states based on the transition functions.

## Objectives:

1. Understand what a language is and what it needs to have in order to be considered a formal one.

2. Provide the initial setup for the evolving project that you will work on during this semester. I said project because usually at lab works, I encourage/impose students to treat all the labs like stages of development of a whole project. Basically you need to do the following:

    a. Create a local && remote repository of a VCS hosting service;

    b. Choose a programming language, and my suggestion would be to choose one that supports all the main paradigms;

    c. Create a separate folder where you will be keeping the report;

3. According to your variant number (by universal convention it is register ID), get the grammar definition and do the following tasks:

    a. Implement a type/class for your grammar;

    b. Add one function that would generate 5 valid strings from the language expressed by your given grammar;

    c. Implement some functionality that would convert and object of type Grammar to one of type Finite Automaton;
    
    d. For the Finite Automaton, please add a method that checks if an input string can be obtained via the state transition from it;

## Implementation

* For the grammar I made a separate class having all the necessary sets and the productions as a dictionary

    ```
    class Grammar:   
        def __init__(self, Vn, Vt, P, S):
            self.Vn = Vn
            self.Vt = Vt
            self.P = P
            self.S = S
    ```
    Here is the call of Grammar constructor in <i>main.py</i> file
    ```
    G = Grammar.Grammar(['S', 'B', 'D'],
                        ['a', 'b', 'c', 'd'],
                        {
                            'S': ['aS', 'bB'],
                            'B': ['cB', 'd', 'aD'],
                            'D': ['aB', 'b']
                        },
                        'S')
    ```

* Grammar class contains a function that generates a word out of the grammar
```def generateString(self) -> str:```  
This function does the following:
  1. Checks if there are any non-terminal symbols in the word;
  2. Gets a random transition from the productions' dictionary;
  3. Replaces the non-terminal symbol with chosen transition;
  4. Repeats the process.

* The Automaton class is similar to the grammar class

    ```
    class FiniteAutomaton:
        def __init__(self,
                     states,
                     alphabet,
                     transitions,
                     initial_state,
                     final_sates) -> None:
            self.states = states
            self.alphabet = alphabet
            self.transitions = transitions
            self.initial_state = initial_state
            self.final_states = final_sates
    ```
* Automaton class contains two functions

    a. Function that converts Grammar to Finite Automata and returns the NFA:
        ```def RG_to_NFA(self, grammar) -> 'FiniteAutomaton':```

    b. Function that checks if a word belongs to the language
        ```def check_word(self, word) -> bool:```

## Conclusions
To implement grammar and finite automata into code is easy if you understand the theory.
In order to make it, I had to do some more studying before starting to code.
