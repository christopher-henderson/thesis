from backtrack import backtrack
import os

db = '/Users/chris/Documents/ASU/thesis/backtrack/map_results'

AL = "Alabama"
AK = "Alaska"
AZ = "Arizona"
AR = "Arkansas"
CA = "California"
CO = "Colorado"
CT = "Connecticut"
DE = "Delaware"
FL = "Florida"
GA = "Georgia"
HI = "Hawaii"
ID = "Idaho"
IL = "Illinois"
IN = "Indiana"
IA = "Iowa"
KS = "Kansas"
KY = "Kentucky"
LA = "Louisiana"
ME = "Maine"
MD = "Maryland"
MA = "Massachusetts"
MI = "Michigan"
MN = "Minnesota"
MS = "Mississippi"
MO = "Missouri"
MT = "Montana"
NE = "Nebraska"
NV = "Nevada"
NH = "NewHampshire"
NJ = "NewJersey"
NM = "NewMexico"
NY = "NewYork"
NC = "NorthCarolina"
ND = "NorthDakota"
OH = "Ohio"
OK = "Oklahoma"
OR = "Oregon"
PA = "Pennsylvania"
RI = "RhodeIsland"
SC = "SouthCarolina"
SD = "SouthDakota"
TN = "Tennessee"
TX = "Texas"
UT = "Utah"
VT = "Vermont"
VA = "Virginia"
WA = "Washington"
WV = "WestVirginia"
WI = "Wisconsin"
WY = "Wyoming"

united_states_of_america = {
    AL: {GA, FL, TN, MS},
    AK: {MO, TN, MS},
    AZ: {CA, NV, UT, CO, NM},
    AR: {MO, OK, TX, LA, TN, MS},
    CA: {OR, NV, AZ},
    CO: {WY, NE, KS, OK, NM, AZ, UT},
    CT: {RI, MA, NY},
    DE: {PA, NJ, MD},
    FL: {AL, GA},
    GA: {SC, NC, TN, AL, FL},
    HI: {},
    ID: {WA, MT, OR, WY, UT, NV},
    IL: {WI, IA, MO, KY, IN, MI},
    IN: {MI, WI, IL, KY, OH},
    IA: {MN, SD, NE, MO, WI, IL},
    KS: {NE, CO, OK, MO},
    KY: {IN, IL, MO, TN, OH, WV, VA},
    LA: {AR, TX, MS},
    ME: {NH},
    MD: {VA, WV, PA},
    MA: {CT, RI, NH, VT},
    MI: {IL, WI, IN, OH},
    MN: {ND, SD, IA, WI},
    MS: {TN, AR, LA, AL},
    MO: {IA, NE, KS, OK, AR, IL, KY, TN},
    MT: {ID, WY, SD, ND},
    NE: {SD, WY, CO, KS, MO, IA},
    NV: {OR, ID, UT, AZ, CA},
    NH: {MA, VT, ME},
    NJ: {NY, PA},
    NM: {AZ, UT, CO, OK, TX},
    NY: {NJ, PA, CT, MA, VT},
    NC: {GA, TN, SC, VA},
    ND: {MT, SD, MN},
    OH: {MI, IN, KY, WV},
    OK: {KS, CO, NM, TX, AR, MO},
    OR: {WA, ID, NV, CA},
    PA: {DE, MD, WV, OH, NY, NJ, MD},
    RI: {CT, MA},
    SC: {GA, NC},
    SD: {ND, MT, WY, NE, MN, IA},
    TN: {KY, MO, AR, MS, MO, AL, GA, NC},
    TX: {OK, NM, AR, LA},
    UT: {ID, NV, WY, CO, AZ, NM},
    VT: {MA, NH, NY},
    VA: {WV, KY, NC},
    WA: {OR, ID},
    WV: {OH, VA, KY, MD},
    WI: {MN, IA, IL, MI, IN},
    WY: {MT, SD, NE, CO, UT, ID},
}

RED = 0
BLUE = 1
GREEN = 2
YELLOW = 3

num_solutions = 0

class State(object):

    CONSTRUCTED = {}

    def __init__(self, state):
        self.CONSTRUCTED[state] = self
        self.color = None
        self.state = state
        # If a neighboring state has already been constructed, then use that. Otherwise construct a new one.
        self.neighbors = [
            self.CONSTRUCTED[neighbor] if neighbor in self.CONSTRUCTED else State(neighbor)
            for neighbor in united_states_of_america[state]]

    @staticmethod
    def reject(solution, state):
        return any(neighbor.color == state.color for neighbor in state.neighbors)

    @classmethod
    def accept(cls, solution):
        return len(solution) == len(united_states_of_america)

    def first(self):
        # Find the next neighbor that hasn't been colored yet
        for neighbor in self.neighbors:
            if neighbor.color is None:
                neighbor.color = RED
                return neighbor
        # If all neighbors have been colored, then try the rest of the coutry.
        # This is because the non-continental US states (Alaska and Hawaii). They make
        # the tree actually a forest. This finds the next tree and continues on.
        for _, state in self.CONSTRUCTED.items():
            if state.color is None:
                state.color = RED
                return state
        return None

    def next(self):
        if self.color < YELLOW:
            self.color += 1
            return self
        self.color = None
        return None

    @staticmethod
    def output(solution):
        global num_solutions
        num_solutions += 1
        if num_solutions % 1000 is 0:
            print(num_solutions)

        # Get the first solution and dump it out to CSV, otherwise this will just go on for
        # quite some time. Supposedly there are over 19 trillion solutions 0_0 
        # color_map = {
        #     0: 'red',
        #     1: 'blue',
        #     2: 'green',
        #     3: 'yellow'
        # }
        # with open(os.path.join(db, 'us_map_coloring.csv'), 'w+') as s:
        #     s.write('state, color\n')
        #     for row in solution:
        #         r_str = str(row)[1:-1] + '\n'
        #         s.write(r_str)
        # exit(0)

    @classmethod
    def construct_map(cls):
        for state in (s for s in united_states_of_america.keys() if s not in cls.CONSTRUCTED):
            State(state)

    def __str__(self):
        return str(tuple((self.state, self.color)))

    def __unicode__(self):
        return str(self)

    def __repr__(self):
        return str(self)


State.construct_map()
State.CONSTRUCTED[CA].color = RED
backtrack(State.CONSTRUCTED[CA], State.first, State.next, State.reject, State.accept, output=State.output)