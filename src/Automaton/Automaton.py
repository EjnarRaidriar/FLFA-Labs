import sys

sys.path.insert(1, "../FLFA-Labs/src/Grammar")
import Grammar


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

    def RG_to_NFA(self, grammar) -> 'FiniteAutomaton':
        productions = grammar.productions
        states = set(productions.keys()) | {"X"}
        alphabet = set()
        transitions = {}
        initial_state = grammar.initial_symbol()
        final_states = set("X") | \
                       ({initial_state} if "&" in productions[initial_state] else set())
        for non_terminal, prods in productions.items():
            for production in prods:
                if production == "&":
                    continue
                new_transition = "X" if len(production) == 1 else production[1]
                transitions.setdefault(
                    (non_terminal, production[0]), set()
                ).add(new_transition)
                alphabet.add(production[0])
        return FiniteAutomaton(states, alphabet, transitions, initial_state, final_states)

    def check_word(self, word) -> bool:
        current_state = {self.initial_state}
        for symbol in word:
            next_state = set()
            for state in current_state:
                next_state.update(
                    self.transitions.get((state, symbol), set())
                )
            current_state = next_state
        return bool(current_state.intersection(self.final_states))

    def determine_FA(fa) -> str:
        nfa = False
        e_nfa = False
        transition_count = []
        for k, states in fa.transitions.items():
            for state in states:
                if k[1] == '&':
                    e_nfa = True
                if k[0] in transition_count:
                    nfa = True
                else:
                    transition_count.append(k[0])
        if e_nfa:
            return "&-NFA"
        elif nfa:
            return "NFA"
        return "DFA"

    def closure(fa) -> set:
        closure = set()
        closure.add(fa.initial_state)
        for element in closure:
            for k, states in fa.transitions.items():
                for state in states:
                    if k[0] == element and k[1] == '&':
                        closure.add(state)
        return closure

