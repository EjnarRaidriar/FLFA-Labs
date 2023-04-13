from Grammar import Grammar
from Automaton import Automaton

grammar = Grammar.Grammar(
    ['S', 'B', 'D'],
    ['a', 'b', 'c', 'd'],
    {
        'S': ['aS', 'bB'],
        'B': ['cB', 'd', 'aD'],
        'D': ['aB', 'b']
    },
    'S')

# Laboratory work 1
word = grammar.generateString()
print(word)
A = Automaton.FiniteAutomaton([], [], {}, "", [])
A = Automaton.FiniteAutomaton.RG_to_NFA(A, grammar)
if A.check_word(word):
    print("check")

print(A.transitions)

# Laboratory work 2
# print(grammar.define_grammar())
