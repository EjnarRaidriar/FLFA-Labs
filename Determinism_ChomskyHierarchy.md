
# Topic: Determinism in Finite Automata. Conversion from NDFA 2 DFA. Chomsky Hierarchy.
### Course: Formal Languages & Finite Automata
### Author: Cucoș Emanuil

-----

## Theory

### Finite Automaton:

A finite automaton is a mathematical model of a finite-state machine. 
The state machine has a set of input symbols and jumps through the set of states based on the transition functions.

### Non-deterministic Finite Automaton (NFA):

NFA is a finite-state machine which accepts multiple transitions of the same symbol from a state

### Deterministic Finite Automaton (DFA):

DFA is a finite-state machine that accepts or rejects a given string of symbols,
by running through a state sequence uniquely determined by the string.

### Chomsky Hierarchy:

Chomsky Hierarchy is a classification of grammars based on the strictness of their rules.
Type 0 - Recursively enumerable grammar
Type 1 - Context-sensitive grammar
Type 2 - Context-free grammar
Type 3 - Regular grammar

## Objectives:

1. Understand what an automaton is and what it can be used for.

2. Continuing the work in the same repository and the same project, the following need to be added:

    a. Provide a function in your grammar type/class that could classify the grammar based on Chomsky hierarchy.

    b. For this you can use the variant from the previous lab.

3. According to your variant number (by universal convention it is register ID), get the finite automaton definition and do the following tasks:

    a. Implement conversion of a finite automaton to a regular grammar.

    b. Determine whether your FA is deterministic or non-deterministic.

    c. Implement some functionality that would convert an NDFA to a DFA.
    
    d. Represent the finite automaton graphically (Optional, and can be considered as a __*bonus point*__).

## Implementation

For implementation of this lab the GO language was used, the main reason for swtitching was the abbsence of types in python language.

### Grammar

#### Defition of a grammar

Grammar class was rewritten in GO language to include typization.
Here is the class:  
```GO
type Grammar struct {
	nonTerminals  []rune
	terminals     []rune
	productions   map[string][]string
	initialSymbol rune
}
```

The function ```DefineGrammar() GrammarType``` will return an enum that defines the type of the grammar according to Chomsky Hierarchy.

This is the enum:
```GO
type GrammarType string

const (
	RightRegular GrammarType = "Right Regular"
	LeftRegular  GrammarType = "Left Regular"
	Type_2       GrammarType = "Type 2"
	Type_1       GrammarType = "Type 1"
	Type_0       GrammarType = "Type 0"
)
```

In order to define grammar, the function uses exclusion method.
It goes through some checks in productions and eliminates grammars with stricter rules.  
The functinion has the following structure:
```GO
func (g Grammar) DefineGrammar() GrammarType {
  // 1. Variable initialization
	for nonTerminal, productions := range g.productions {
    // 2. Check Non-Terminal count in the left side of production
		for _, prod := range productions {
      // 3. Check if grammar is Type 0
      // 4. Check if grammar is not Type 3
      // 5. Check if Type 3 is left or right liniar
		}
	}
  // 6. Return
}
```

Function steps:
1. Variable initailization:\
  These variables will be used at the return to find the grammar type. ```isType3``` and ```isType2``` are set to ```true``` to perform checks if grammar is not of type 3 or 2.
    ```GO
    isType3 := true
    isRightLiniar := false
    isType2 := true
    ```
2. Check Non-Terminal count in the left side of production:\
  If the left side of production has more than one Non-Terminal it can't be of type 2 or 3.
    ```GO
    if utf8.RuneCountInString(nonTerminal) > 1 {
      isType3 = false
      isType2 = false
    }
    ```
3. Check if grammar is Type 0:\
  Grammar is type 0 if left side of production is longer than the right side or it is not of type 3 or 2 and left side is not initials sybmol while the right side is empty string.
    ```GO
    prodLen := utf8.RuneCountInString(prod)
    if utf8.RuneCountInString(nonTerminal) > prodLen {
      return Type_0
    }
    if (isType2 || isType3) == false &&
      nonTerminal != string(g.initialSymbol) &&
      prod != "&" {
      return Type_0
    }
    ```

4. Check if grammar is not Type 3:\
  Grammar can't be of type 3 if right side of production is one uppercase (non-terminal) or if it is longer and has multiple non-terminals.
    ```GO
    if isType2 && functions.IsLower(prod) == false {
      if functions.IsUpper(prod) {
        isType3 = false
      } else if functions.HasMultipleUpper(prod) {
        isType3 = false
      }
    }
    ```
5. Check if Type 3 is left or right liniar:\
  If last letter is upper (non-terminal) then it is right liniar, but if is both right and left liniar it is not type 3.
    ```GO
    lastLetter := rune(prod[prodLen-1])
    if isType3 && prodLen > 1 {
      if unicode.IsUpper(lastLetter) {
        isRightLiniar = true
      }
      if isRightLiniar && unicode.IsLower(lastLetter) {
        isType3 = false
      }
    }
    ```
6. Return:\
  First part checks if is type 3 and it's subtype, then checks if it is type 2. If none of those are true then the ramaining part is type 1. Type 0 was checked earlier.
    ```GO
    if isType3 && isRightLiniar {
      return RightRegular
    } else if isType3 {
      return LeftRegular
    }
    if isType2 {
      return Type_2
    }
    return Type_1
    ```
  
### Finite Automata

#### Conversion from Fintite Automata to Regular Grammar
Conversion from ```FiniteAutomaton``` to ```Grammar``` is done using three functions: ```FaToRg()```, ```addProduction()``` and ```statesToNonTerminals()```.

Functions ```statesToNonTerminals()``` return mapping of states to the non-terminals which later will be used to make the grammar.
The function takes as parameters FA states and start state.
Grammar's initial state is set to 'S', then the fuction loops through an alphabet and matches states with an uppercase letter.
```GO
func statesToNonTerminals(states []string, initialState string) map[string]rune {
	var upperCaseAlphabet []rune = []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I',
		'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R',
		/* 'S',*/ 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
  // set initial symbol to S
	result := map[string]rune{initialState: 'S'}
	counter := 0
  // match states to non-terminals in alphabetical order
	for _, state := range states {
		if state != initialState {
			result[state] = upperCaseAlphabet[counter]
			counter++
		}
	}
	return result
}
```

The function ```addProduction()``` doesn't return anything, instead it modifies pointer parameters.
First part checks if the non-terminal is already registered.
If true, production is appended, if false, new non-terminal is saved and a new list of productions is made.
```GO
func addProduction(productions *map[string][]string, nonTerminals *[]rune, nonTerminal rune, production string) {
	newProduction := *productions
	if _, ok := newProduction[string(nonTerminal)]; ok {
    // append production
		newProduction[string(nonTerminal)] = append(newProduction[string(nonTerminal)], production)
	} else {
    // create new production list and save new non-terminal
		newProduction[string(nonTerminal)] = []string{production}
		*nonTerminals = append(*nonTerminals, nonTerminal)
	}
	*productions = newProduction
}
```

The conversion function of Finite Automaton to Regular Grammar has the following structure:
```GO
func FaToRg(f *automaton.FiniteAutomaton) *grammar.Grammar {
  // 1. Variable initialization
  // 2. Check if final state has loop transitions

	for _, transition := range f.GetTransitions() {
    // 3. Add productions and record new states
	}
	return grammar.NewGrammar(nonTerminals, f.GetAlphabet(), productions, conversionMap[f.GetInitialState()])
}
```

1. Variable initialization:
    ```GO
    conversionMap := statesToNonTerminals(f.GetStates(), f.GetInitialState())
    var nonTerminals []rune
    productions := make(map[string][]string)
    finalStates := f.GetFinalStates()
    finalNotFinal := false
    ```


2. Check if final state has loop transitions:\
  This check is needed for form productions because.
  Presence of a final state that has transition into itself will change the way states that have transitions into final state will be converted into productions.
    ```GO
    for _, transition := range f.GetTransitions() {
      if transition.InitialState == transition.NextState &&
        slices.Contains(finalStates, transition.NextState) {
        finalNotFinal = true
      }
    }
    ```
3. Add productions and record new states:\
  There are three cases of making a new production depending on where and from where the transition is being made.\
  First, a state transitions into another non-final state. Then the production will be of the form A->aB.\
  Second, a state transitions into final state and final state doesn't have transitions into itself.
  Then the production will be of the form A->a.\
  Third, a state transitions into final state and the final state has a transiton into itself.
  Then ther will be two productions added A->aB and A->a.
    ```GO
		var production string
		nonTerminal := conversionMap[transition.InitialState]
		if slices.Contains(finalStates, transition.NextState) {
			if finalNotFinal {
				if transition.InitialState == transition.NextState {
          // third case
					production = strings.Join([]string{string(transition.Transition), string(conversionMap[transition.NextState])}, "")
					addProduction(&productions, &nonTerminals, nonTerminal, production)
					production = string(transition.Transition)
					addProduction(&productions, &nonTerminals, nonTerminal, production)
				} else {
          // second case
					production = string(transition.Transition)
					addProduction(&productions, &nonTerminals, nonTerminal, production)
				}
			}
		} else {
      // first case
			production = strings.Join([]string{string(transition.Transition), string(conversionMap[transition.NextState])}, "")
			addProduction(&productions, &nonTerminals, nonTerminal, production)
		}
    ```

#### Determining if Finite Automaton is DFA, NFA or &-NFA

Determination is done by the function ```DetermineFA()``` which returns the definition through a string.

```GO
func (f *FiniteAutomaton) DetermineFA() string {
  // 1. Variable initialization
	for _, transition := range f.transitions {
    // 2. Check if is &-NFA
    // 3. Check if is NFA and store new transition
	}
	if isNFA {
		return "NFA"
	}
	return "DFA"
}
```
1. Variable initialization:\
  <i>transitions</i> variable will store traversed transitions without the state it goes into and <i>isNFA</i> will be used to return definition.
    ```GO
    transitions := make(map[string]rune)
    isNFA := false
    ```
2. Check if is &-NFA:\
  If any epsilon transtion is found it is automaticaly &-NFA.
    ```GO
		if transition.Transition == '&' {
			return "&-NFA"
		}
    ```
3. Check if is NFA and store new transition
  In case a transition already exists then it is NFA, but the loop will continue to check if there are any epsilon transitions.
  If the transition is new it is stored in the transitions map.
    ```GO
		if initialState, ok := transitions[transition.InitialState]; ok &&
			initialState == transition.Transition {
			isNFA = true
		} else {
			transitions[transition.InitialState] = transition.Transition
		}
    ```
  
#### Conversion from &-NFA and NFA into DFA

The conversion is supported by two private methods ```closure()``` and ```joinStates```.
Additionally there are two custom functions for type conversions and map search.

The function ```closure()``` returns an epsilon closure of a state.
It stores the states of the clousure in a set by traversing through transitions and saving a state if is reached by epsilon transition.

Here is the function:
  ```GO
  func (f *FiniteAutomaton) closure(state string) map[string]bool {
    closure := map[string]bool{state: true}
    for element := range closure {
      for _, transition := range f.transitions {
        if transition.InitialState == element &&
          transition.Transition == '&' {
          closure[transition.NextState] = true
        }
      }
    }
    return closure
  }
  ```

The function ```joinStates()``` transforms from a list of states into a single state.

```GO
func joinStates(states []string) string {
	if len(states) > 1 {
		return "{" + strings.Join(states, ", ") + "}"
	}
	return states[0]
}
```

Here is the structure of the function to convert into DFA:
```GO
func MakeDFA(f FiniteAutomaton) *FiniteAutomaton {
  // 1. Check if is already DFA
  // 2. Variable initialization
  // 3. Finding epsilon closures
	i := 0
	for i < len(dfaStatesMap) {
		state := dfaStatesMap[i]
		for _, symbol := range f.alphabet {
      // 4. Form new state
      // 5. Save new state
		}
		i++
	}
  // 6. Form states list
	return NewFiniteAutomaton(dfaStates, f.alphabet, dfaTransitions, f.initialState, dfaFinalStates)
}
```
A while loop is used to iterate through the states map because the map could be updated with new states that have to be iterated through.

1. Check if is already DFA:\
    There is no need to do go through the function if it is already a DFA.
    ```GO
    if f.DetermineFA() == "DFA" {
      return &f
    }
    ```

2. Variable initialization:\
    <i>dfaStatesMap</i> is a list of states which are stored as maps. The map keys will form a new complex state.
    ```GO
    dfaTransitions := make([]Transition, 0, 10)
    dfaFinalStates := make([]string, 0, 2)
    dfaStatesMap := make([]map[string]bool, 0)
    ```

3. Finding epsilon closures:\
    Epsilon closures are required to convert from &-NFA to DFA.
    A closure of each state consists from states that can be accessed via epsilon transitions in chain.
    ```GO
    closures := make(map[string]map[string]bool)
    for _, state := range f.states {
      closures[state] = f.closure(state)
      var newState map[string]bool = closures[state]
      dfaStatesMap = append(dfaStatesMap, newState)
    }
    ```

4. Form new state:\
    New state will be non-empty if a state has multiple transitions through the same symbol.
    ```GO
    newStateMap := make(map[string]bool)
    for element := range state {
      for _, transition := range f.transitions {
        if transition.Transition == symbol &&
          transition.InitialState == element {
          for closure, ok := range closures[transition.NextState] {
            newStateMap[closure] = ok
          }
        }
      }
    }
    ```

5. Save new state:\
    If there is a new state, then the map and transitions are updated and the while loop will iterate through it too.
    ```GO
    if len(newStateMap) > 0 {
      stateSlice := functions.KeyList(state)
      newStateSlice := functions.KeyList(newStateMap)
      dfaTransitions = append(dfaTransitions, Transition{
        joinStates(stateSlice),
        symbol,
        joinStates(newStateSlice),
      })
      if functions.ContainsMap(dfaStatesMap, newStateMap) == false {
        dfaStatesMap = append(dfaStatesMap, newStateMap)
      }
    }
    ```

6. Form states list:\
    This part converts from a list of states map into a list of states and list of final states.
    ```GO
    dfaStates := make([]string, 0, len(dfaStatesMap))
    for _, stateMap := range dfaStatesMap {
      state := functions.KeyList(stateMap)
      dfaStates = append(dfaStates, joinStates(state))
      for _, finalState := range f.finalStates {
        if slices.Contains(state, finalState) {
          dfaFinalStates = append(dfaFinalStates, joinStates(state))
        }
      }
    }
    ```
  

## Conclusion

This laboratory was intedet to understand the principles of determinism in finite automata, the conversion process from NDFA to DFA, and the classification of grammars according to the Chomsky hierarchy.
The implementation was carried out in the Go language, main reasons being type safety.
The `Grammar` updated with a function to classify grammars based on their production rules, while the `FiniteAutomaton` class received methods for converting finite automata into regular grammars and determining whether an automaton is deterministic or non-deterministic.
Additionally, a conversion function was implemented to transform an &-NDFA into a DFA.
