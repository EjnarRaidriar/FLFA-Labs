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

2. Provide the initial setup for the evolving project that you will work on during this semester. You can deal with each laboratory work as a separate task or project to demonstrate your understanding of the given themes, but you also can deal with labs as stages of making your own big solution, your own project. Do the following:

    a. Create GitHub repository to deal with storing and updating your project;

    b. Choose a programming language. Pick one that will be easiest for dealing with your tasks, you need to learn how to solve the problem itself, not everything around the problem (like setting up the project, launching it correctly and etc.);

    c. Store reports separately in a way to make verification of your work simpler

3. According to your variant number, get the grammar definition and do the following tasks:

    a. Implement a type/class for your grammar;

    b. Add one function that would generate 5 valid strings from the language expressed by your given grammar;

    c. Implement some functionality that would convert and object of type Grammar to one of type Finite Automaton;
    
    d. For the Finite Automaton, please add a method that checks if an input string can be obtained via the state transition from it;

## Implementation

For labs implementation was chosen the Python programming Language.
The main reasons for this choice where the ease of implementation in this language without sacrifices any sacrifices in understanding the algorithms.

The laboratory work is made of two classes: <i>Grammar</i>, and <i>FiniteAutomaton</i>.

### Grammar

The <i>Grammar</i> class requires in its constructor the following parameters:
Non-Terminal symbols, Terminal symbols, Production rules and Initial symbol.

Grammar class contains a function that generates a word out of the grammar:\
```def generateString(self) -> str:```

The function has the following structure:
```python
def generateString(self) -> str:
    # 1. variable initialization
    x = True
    while x:
        # 2. non-terminal states check
        # 3. production
    return word
```
Function steps:

1. Variables initialization:
    ```python
    word = self._initial_symbol
    x = True
    ```
    The <i>word</i> variable will be the final word, but now it starts with initial symbol.\
    The x variable is a counter for the while loop.
2. Non-terminal states check:
    ```python
    if not any(n in word for n in self._non_terminal):
        x = False
        break
    ```
    The if statement checks if there are any non-terminal states in the word
and exits the loop if there are no non-terminal states.
The negation is used to avoid nesting.
3. Production:
    ```python
    for n in self._non_terminal:
        if n in word:
            transition = self._productions.get(n)
            transition = random.choice(transition)
            word = word.replace(n, transition)
    ```
    The function loops through the non-terminal symbols of the grammar and then checks if that symbol is present in the word. 
    If it is true, a production random production rule from the non-terminal symbol is assigned to the transition variable. 
    Then inside the word the non-terminal symbol is replaced with the transition.
4. Function end:\
    The while loop will be repeated until there are no non-terminal symbols in the word. 
    When that happens, the word is returned from the function.

### FiniteAutomaton

The <i>FiniteAutomaton</i> class requires in its constructor:
States, Alphabet, Transitions, Initial state and Final state.

<i>FiniteAutomaton</i> class contains a static function to convert from a grammar into a finite automaton:
```python
@staticmethod
def RG_to_NFA(grammar: 'Grammar') -> 'FiniteAutomaton':
    # 1. variable initialization
    # 2. looping over items
        # 3. add transitions and update alphabet
    # 4. return
    return FiniteAutomaton()
```
1. Variable initialization:
    ```python
    productions = grammar.productions()
    states = set(productions.keys()) | {"X"}
    alphabet = set()
    transitions = {}
    initial_state = grammar.initial_symbol()
    final_states = set("X")
    ```
    The states, including final and initial ones are defined right away and won't be changed. 
    The main working type is set in order to avoid duplicates later on. 
    The final state is added by default as an addition to feature non-terminal state.

2. Looping:
    ```python
    for non_terminal, prods in productions.items():
        for production in prods:
    ```
    The first iteration goes through the non-terminals of grammar productions and the second one through each production rule of a non-terminal.

3. Adding transitions and updating the alphabet:
    ```python
    # inside second loop
    new_transition = "X" if len(production) == 1 else production[1]
    transitions.setdefault(
        (non_terminal, production[0]), set()
    ).add(new_transition)
    alphabet.add(production[0])
    ```
    <i>new_transition</i> is the state in which the transition is being made. 
    If the production is of length one then it has to be a terminal state (assuming the regular grammar is right liniar) otherwise it must be a non-terminal state. 
    Then the transitions dictionary is updated where the key is a tuple of a non-terminal state with transition through the terminal state and the value is the ending state. 
    At the end the alphabet set is updated.
4. Return:
    ```python
    return FiniteAutomaton(states, alphabet, transitions, initial_state, final_states)
    ```
    The function returns an object of <i>FiniteAutomaton</i> with the attributes according to the grammar.


## Conclusion

This laboratory was intended to understand the concepts of formal languages and finite automata. 
The implementation was done in Python language. 
The `Grammar` class was contains a function to generate valid strings from a given grammar, with the use of production rules. 
The `FiniteAutomaton` class was implemented with a method to convert a regular grammar into a non-deterministic finite automaton (NFA). 
