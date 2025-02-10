from Automaton import FiniteAutomaton
from Grammar import Grammar


def lab1():
    g = Grammar(
        ['S', 'B', 'C', 'D'],
        ['a', 'b', 'c'],
        {
            'S': 'aB',
            'B': ['bS', 'aC', 'c'],
            'C': ['bD'],
            'D': ['c', 'aC']
        },
        'S'
    )
    print('Words generated by grammar:')
    for i in range(5):
        print(g.generateString(), end=", ")
    fa = FiniteAutomaton.RG_to_NFA(g)
    print()
    userinput = input("Choose a word to check if it can be obtained from FA: ")
    print(fa.check_word(userinput))


def main():
    lab1()


if __name__ == "__main__":
    main()

# RG_to_NFA = FiniteAutomaton.RG_to_NFA
# determine_FA = FiniteAutomaton.determine_FA
# FA = FiniteAutomaton
# NFA_to_DFA = FiniteAutomaton.NFA_to_DFA
#
# nfa = FA(
#     ["q0", "q1", "q2"],
#     ["a", "b", "c"],
#     {
#         ('q0', 'a'): {'q0'},
#         ('q0', '&'): {'q1'},
#         ('q0', 'b'): {'q1'},
#         ('q1', 'c'): {'q1', 'q2'},
#         ('q1', '&'): {'q2'},
#         ('q2', 'a'): {'q0'},
#         ('q1', 'a'): {'q1'}
#     },
#     "q0",
#     ["q2"])
#
#
# print(nfa.determine_FA())
# print(nfa.initial_state)
# print(nfa.states)
# print(nfa.final_states)
# for key, values in nfa.transitions.items():
#     print(key, ": ", values)
# print()
#
# print('Grammar from ', nfa.determine_FA())
# g = Grammar.FA_to_RG(nfa)
# print(g.initial_symbol)
# print(g.non_terminal)
# print(g.terminal)
# print(g.productions)
# print()
#
# nfa.NFA_to_DFA()
# print(nfa.determine_FA())
# print(nfa.initial_state)
# print(nfa.states)
# print(nfa.final_states)
# for key, values in nfa.transitions.items():
#     print(key, ": ", values)
# print()
#
# print('Grammar from ', nfa.determine_FA())
# g = Grammar.FA_to_RG(nfa)
# print(g.initial_symbol)
# print(g.non_terminal)
# print(g.terminal)
# print(g.productions)
