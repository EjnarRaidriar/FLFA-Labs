from Grammar import Grammar
from Automaton import Automaton

G = Grammar.Grammar( ['S', 'B', 'D'],
            ['a', 'b', 'c', 'd'],
            {
                'S': ['aS', 'bB'],
                'B': ['cB', 'd', 'aD'],
                'D': ['aB', 'b']
            },
            'S')

word = G.generateString()
print(word)
A = Automaton.FiniteAutomaton([], [], {}, "", [])
A = Automaton.FiniteAutomaton.RG_to_NFA(A, G)
if A.check_word(word):
    print("check")

