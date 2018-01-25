#!/usr/bin/env swipl

:- discontiguous expression/3.
:- discontiguous eval_expr/3.
:- discontiguous eval_expr/2.
:- discontiguous eval/2.
:- discontiguous eval/3.
:- discontiguous eval/4.

:- initialization main.

main() :-
  run("python.pl").

% Reads in a Spark source file and "executes" the program.
%
%                        +-------+                           +-------+
%   +---[Source file]--->| run/1 |---[List of characters]--->| lexer |---+
%                        +-------+                           +-------+   |
%                                                                        |
%   +-------------------------[List of tokens]---------------------------+
%   |
%   |     +--------+                     +-------------+
%   +---->| parser |----[Parse tree]---->| interpreter |----[Output]----->
%         +--------+              
       %% +-------------+
%
% Executing the query '?- working_directory(CWD, CWD).' will output the
% current working directory in the Prolog interpreter.

run(Filename) :-
    % Read file into list of characters.
    parse(Filename, Tree),
    % print_term(Tree, []), nl,
    % write("Interpreting..."), nl,
    eval(Tree), !.

parse(Filename, Tree) :-
  % Read file into list of characters.
  read_file_to_codes(Filename, Codes, []),
  atom_codes(Atom, Codes),
  atom_chars(Atom, Characters),

  % Tokenize, parse, and evaluate.
  write("Lexing..."), nl,
  lexer(Tokens, Characters, []),
  write(Tokens), nl, !,
  parser(Tree, Tokens, _),
  write(Tree).

backtrack(root(R), sr(SOLUTION_REJECT), c(CANDIDATE), reject(RJ_SL), sa(SOLUTION_ACCEPT), accept(ACC_SL), p(PARENT), children(CH_SL)) -->
  ["bracktrack"], ["on"], expression(R), colon, newline,
    ["reject"], open_parenthesis, name(SOLUTION_REJECT), comma, name(CANDIDATE), close_parenthesis, colon, newline,
      statement_list(RJ_SL), newline,
    ["accept"], open_parenthesis, name(SOLUTION_ACCEPT), close_parenthesis, colon, newline,
      statement_list(ACC_SL), newline,
    ["children"], open_parenthesis, name(PARENT), close_parenthesis, colon, newline,
      statement_list(CH_SL).

%% % ----------------------------------------------------------------------
%% %   LEXER
%% % ----------------------------------------------------------------------

bckt --> ["backtrack"].
on --> ["on"].
reject --> ["reject"].
accept --> ["accept"].
children --> ["children"].

colon --> [":"].
newline --> ["\n"].
open_parenthesis --> ["("].
close_parenthesis --> [")"].
comma --> [","].

control_characters([
  bkct,
  on,
  reject,
  accept, children,

  colon,
  newline,
  open_parenthesis,
  close_parenthesis,
  comma
  ]).

% Whitespace characters.
whitespace --> [' '].
whitespace --> ['\t'].
%% whitespace --> ['\n'].

% The longest_match_alpha/3 predicate unifies with the next longest string of
% uppercase or lowercase letters in the list of remaining characters.
longest_match_alpha([H|T]) --> [H], longest_match_alpha_(T), {char_type(H, alpha)}.
longest_match_alpha_([H|T]) --> [H], longest_match_alpha_(T), {char_type(H, alpha)}.
longest_match_alpha_([]) --> [].

% The longest_match_no_nl/3 predicate unifies with the next longest string of
% non-newline characters in the list of remaining characters.
longest_match_no_nl([H|T]) --> [H], longest_match_no_nl_(T), {H \== '\n'}.
longest_match_no_nl_([H|T]) --> [H], longest_match_no_nl_(T), {H \== '\n'}.
longest_match_no_nl_([]) --> [].

% Whitespace characters.
lexer(Tokens) --> whitespace, lexer(R), {Tokens = R}.

% Everything else.
lexer(Tokens) --> [C], lexer(R), {Tokens = [C | R]}.
lexer(Tokens) --> [], {Tokens = []}.

%% % ----------------------------------------------------------------------
%% %   PARSER
%% % ----------------------------------------------------------------------

%% % Reserved words.
%% reserved_words([
%%     while,
%%     for,
%%     if,
%%     else,
%%     print,
%%     true,
%%     false,
%%     and,
%%     or,
%%     not,
%%     xor,
%%     function,
%%     return
%% ]).

%% % Integers.
%% % All atom tokens that are composed entirely of digit characters are
%% % parsed as integers.
%% integer(int(I)) --> [I], {atom_number(I, _)}.
%% integer(neg(I)) --> ['-'], expression(I).

%% % Identifiers.
%% % All atom tokens that are composed entirely of lowercase and uppercase
%% % letters that are not reserved words are parsed as identifiers.
%% identifier(id(I)) -->
%%     [I],
%%     {
%%         atom_chars(I, C),
%%         all_letters(C),
%%         reserved_words(RW),
%%         \+ member(I, RW)
%%     }.
%% all_letters([H|T]) :- char_type(H, alpha), all_letters(T).
%% all_letters([]).

%% % Booleans.
%% boolean(bool('true')) --> ['true'].
%% boolean(bool('false')) --> ['false'].

%% % Print to stdout.
%% print(print(E)) --> ['print'], expression(E).

%% statement(E) --> expression(E), [';'].
%% statement(B) --> branch(B).
%% statement(L) --> loop(L).
%% statement(F) --> function_declaration(F).
%% % statement(C) --> conditional_branch(C).

%% expression(I) --> identifier(I).
%% expression(B) --> boolean(B).
%% expression(I) --> integer(I).
%% expression(P) --> print(P).
%% expression(F) --> function_call(F).
%% expression(R) --> return(R).
%% expression(L) --> lambda(L).

%% lefthand_expression(I) --> identifier(I).
%% lefthand_expression(B) --> boolean(B).
%% lefthand_expression(I) --> integer(I).
%% lefthand_expression(F) --> function_call(F).

%% expression(ea(I, E)) --> identifier(I), ['='], expression(E).

%% expression(ep(LE, E)) --> lefthand_expression(LE), ['+'], expression(E).
%% expression(em(LE, E)) --> lefthand_expression(LE), ['*'], expression(E).
%% expression(es(LE, E)) --> lefthand_expression(LE), ['-'], expression(E).
%% expression(ed(LE, E)) --> lefthand_expression(LE), ['/'], expression(E).
%% expression(er(LE, E)) --> lefthand_expression(LE), ['%'], expression(E).

%% expression(elt(LE, E)) --> lefthand_expression(LE), ['<'], expression(E).
%% expression(egt(LE, E)) --> lefthand_expression(LE), ['>'], expression(E).
%% expression(ele(LE, E)) --> lefthand_expression(LE), ['<='], expression(E).
%% expression(ege(LE, E)) --> lefthand_expression(LE), ['>='], expression(E).
%% expression(eeq(LE, E)) --> lefthand_expression(LE), ['=='], expression(E).
%% expression(enq(LE, E)) --> lefthand_expression(LE), ['!='], expression(E).

%% expression(eeq(B, E)) --> lefthand_expression(B), ['=='], expression(E).
%% expression(enq(B, E)) --> lefthand_expression(B), ['!='], expression(E).
%% expression(ebc(B, E)) --> lefthand_expression(B), ['and'], expression(E).
%% expression(ebd(B, E)) --> lefthand_expression(B), ['or'], expression(E).
%% expression(ebx(B, E)) --> lefthand_expression(B), ['xor'], expression(E).
%% expression(ebn(E)) --> ['not'], expression(E).

%% % If
%% branch(if(Cond, SL)) -->
%%     ['if'], ['('], expression(Cond), [')'], ['{'],
%%       statement_list(SL),
%%     ['}'].

%% % If-Else
%% branch(if(Cond, SL, SL2)) -->
%%     ['if'], ['('], expression(Cond), [')'], ['{'],
%%       statement_list(SL),
%%     ['}'], ['else'], ['{'],
%%       statement_list(SL2),
%%     ['}'].

%% % If-ElseIf
%% branch(if(Cond, SL, ElseIf)) -->
%%     ['if'], ['('], expression(Cond), [')'], ['{'],
%%       statement_list(SL),
%%     ['}'], branch(ElseIf).

%% % ElseIf
%% branch(elseif(Cond, SL)) -->
%%   ['else'], ['if'], ['('], expression(Cond), [')'], ['{'],
%%     statement_list(SL),
%%   ['}'].

%% % ElseIf-ElseIf
%% branch(elseif(Cond, SL, ElseIf)) -->
%%   ['else'], ['if'], ['('], expression(Cond), [')'], ['{'],
%%     statement_list(SL),
%%   ['}'], branch(ElseIf).

%% % ElseIf-Else
%% branch(elseif(Cond, SL, SL2)) -->
%%   ['else'], ['if'], ['('], expression(Cond), [')'], ['{'],
%%     statement_list(SL),
%%   ['}'], ['else'], ['{'],
%%     statement_list(SL2),
%%   ['}'].

%% loop(while(Cond, SL)) -->
%%   ['while'], ['('], expression(Cond), [')'], ['{'],
%%   statement_list(SL),
%%   ['}'].

%% loop(for(Initializer, Cond, Increment, SL)) -->
%%   ['for'], ['('], expression(Initializer), [';'], expression(Cond), [';'], expression(Increment), [')'], ['{'],
%%     statement_list(SL),
%%   ['}'].

%% function_declaration(function(Identifier, IdentifierList, SL)) -->
%%   ['function'], identifier(Identifier), ['('], identifier_list(IdentifierList), [')'], ['{'],
%%     statement_list(SL),
%%   ['}'].

%% function_declaration(function(Identifier, il(void), SL)) -->
%%   ['function'], identifier(Identifier), ['('], [')'], ['{'],
%%     statement_list(SL),
%%   ['}'].

%% lambda(lambda(IdentifierList, SL)) -->
%%   ['function'], ['('], identifier_list(IdentifierList), [')'], ['{'],
%%     statement_list(SL),
%%   ['}'].

%% lambda(lambda(il(void), SL)) -->
%%   ['function'], ['('], [')'], ['{'],
%%     statement_list(SL),
%%   ['}'].

%% function_call(function(Identifier, ValueList)) -->
%%   identifier(Identifier), ['('], value_list(ValueList), [')'].

%% function_call(function(Identifier, vl(void))) -->
%%   identifier(Identifier), ['('], [')'].

%% return(return(E)) -->
%%   ['return'], expression(E).

%% identifier_list(il(I)) --> identifier(I).
%% identifier_list(il(I, IL)) --> identifier(I), [','], identifier_list(IL).

%% value_list(vl(V)) --> expression(V).
%% value_list(vl(V, VL)) --> expression(V), [','], value_list(VL).

%% statement_list(sl(S)) --> statement(S).
%% statement_list(sl(S, SL)) --> statement(S), statement_list(SL).
%% parser(program(SL)) --> statement_list(SL).

anything(c(C)) --> control_characters(CR),
                    \+ member(C, CR)
%% anything --> [].
%% anything --> [_], anything.
parser(program(SL)) --> statement_list(SL).
statement_list(sl(S)) --> statement(S).
statement_list(sl(S, SL)) --> statement(S), statement_list(SL).
statement(b(R, SOLUTION_REJECT, CANDIDATE, RJ_SL, SOLUTION_ACCEPT, ACC_SL, PARENT, CH_SL)) --> backtrack(R, SOLUTION_REJECT, CANDIDATE, RJ_SL, SOLUTION_ACCEPT, ACC_SL, PARENT, CH_SL).
statement(stmt(E)) --> expression(E), newline.
expression(e(E)) --> anything(E).
