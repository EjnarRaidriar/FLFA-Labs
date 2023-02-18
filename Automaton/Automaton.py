import sys
sys.path.insert(0, "C:/Users/emanu/PycharmProjects/FLFA-Labs")
from Grammar import Grammar

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
        productions = grammar.productions()
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

