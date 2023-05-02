from Grammar import Grammar
from Automaton import Automaton

RG_to_NFA = Automaton.FiniteAutomaton.RG_to_NFA
determine_FA = Automaton.FiniteAutomaton.determine_FA
FA = Automaton.FiniteAutomaton
NFA_to_DFA = Automaton.FiniteAutomaton.NFA_to_DFA

nfa = FA(
    ["q0", "q1", "q2"],
    ["a", "b", "c"],
    {
        ('q0', 'a'): {'q0'},
        ('q0', '&'): {'q1'},
        ('q0', 'b'): {'q1'},
        ('q1', 'c'): {'q1', 'q2'},
        ('q1', '&'): {'q2'},
        ('q2', 'a'): {'q0'},
        ('q1', 'a'): {'q1'}
    },
    "q0",
    ["q2"])


print(nfa.determine_FA())
print(nfa.initial_state)
print(nfa.states)
print(nfa.final_states)
for key, values in nfa.transitions.items():
    print(key, ": ", values)
print()

print('Grammar from ', nfa.determine_FA())
g = Grammar.Grammar.FA_to_RG(nfa)
print(g.initial_symbol)
print(g.non_terminal)
print(g.terminal)
print(g.productions)
print()

nfa.NFA_to_DFA()
print(nfa.determine_FA())
print(nfa.initial_state)
print(nfa.states)
print(nfa.final_states)
for key, values in nfa.transitions.items():
    print(key, ": ", values)
print()

print('Grammar from ', nfa.determine_FA())
g = Grammar.Grammar.FA_to_RG(nfa)
print(g.initial_symbol)
print(g.non_terminal)
print(g.terminal)
print(g.productions)