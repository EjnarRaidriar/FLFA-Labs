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

## Implementation:

* The function _define_grammar({Grammar})_ defines the grammar according to Chomski hierarchy.
In order to define it, the function uses exclusion method.
It goes through some checks in productions and eliminates grammars with stricter rules.  
Possible outputs are: 'Type_0', 'Type_1', 'Type_2', 'Left_Regular' or 'Right_Regular'.


* The function _determine_FA({FA})_ returns a string that represents the FA:
'&-NFA', 'NFA', or 'DFA'.
The variable _transition_count_ is a list that saves the state that was already checked.
The function loops over the transition's directory keys and elements of each their values.
If it finds an epsilon transition, then it is '&-NFA'.
At the end of the iteration of an element in the key, it's key is added to _transition_count_.
If there is another transition from the same key, then it is 'NFA'.
If the loop is completed the function returns 'DFA'


* The function _NFA_to_DFA_() uses table method to convert to DFA.
This function uses another two functions, from the FiniteAutomaton class, and nested functions.  
First function is _determine_FA_() and used at the beginning to check if the automaton is already a DFA. It raises an error if it is one.  
The second function is ___closure({state})_.
This function returns epsilon closure of the parameter in form of a list.  
Function structure:
  1. The function checks if it is a DFA.
  2. Declaration of inner functions: _add_new_transition()_, _update_final_states()_, _find_new_state()_, _append_closure()_.
  3. Initialization of variables:
  _dfa_states_, _dfa_transitions_, _dfa_final_states_ and _closures_.
  4. Fill _closures_ list with epsilon closure of each state of the NFA.
  5. Epsilon closure of the initial state becomes initial state of the FA.
  6. Loop over _dfa_states_ and letters of the FA alphabet.
  7. Inside this loop is initialized _new_state_ as list. This will be the state formed in the transition table.
  8. 

