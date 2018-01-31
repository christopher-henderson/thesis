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
    parse(Filename, Tree),
    write("Evaulating..."),
    eval(Tree), !.

parse(Filename, Tree) :-
  % Read file into list of characters.
  read_file_to_codes(Filename, Codes, []),
  atom_codes(Atom, Codes),
  atom_chars(Atom, Characters),

  % Tokenize, parse, and evaluate.
  write("Lexing..."), nl,
  lexer(Tokens, Characters, []),
  %% write(Tokens), nl, !,
  !, parser(Tree, Tokens, _),
  write("LOL\n"), write(Tree).

%% backtrack(root(R), sr(SOLUTION_REJECT), c(CANDIDATE), reject(RJ_SL), sa(SOLUTION_ACCEPT), accept(ACC_SL), p(PARENT), children(CH_SL)) -->
%%   ["bracktrack"], ["on"], expression(R), colon, newline,
%%     ["reject"], open_parenthesis, name(SOLUTION_REJECT), comma, name(CANDIDATE), close_parenthesis, colon, newline,
%%       statement_list(RJ_SL), newline,
%%     ["accept"], open_parenthesis, name(SOLUTION_ACCEPT), close_parenthesis, colon, newline,
%%       statement_list(ACC_SL), newline,
%%     ["children"], open_parenthesis, name(PARENT), close_parenthesis, colon, newline,
%%       statement_list(CH_SL).

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

% Comment delimiters.
comment_single --> ['/', '/'].
comment_multi_start --> ['/', '*'].
comment_multi_end --> ['*', '/'].

% The longest_match_alpha/3 predicate unifies with the next longest string of
% uppercase or lowercase letters in the list of remaining characters.
longest_match_alpha([H|T]) --> [H], longest_match_alpha_(T), {char_type(H, alpha)}.
longest_match_alpha_([H|T]) --> [H], longest_match_alpha_(T), {char_type(H, alpha)}.
longest_match_alpha_([]) --> [].

% The longest_match_int/3 predicate unifies with the next longest string of
% digits in the list of remaining characters.
longest_match_int([H|T]) --> [H], longest_match_int_(T), {char_type(H, digit)}.
longest_match_int_([H|T]) --> [H], longest_match_int_(T), {char_type(H, digit)}.
longest_match_int_([]) --> [].

% The longest_match_no_nl/3 predicate unifies with the next longest string of
% non-newline characters in the list of remaining characters.
longest_match_no_nl([H|T]) --> [H], longest_match_no_nl_(T), {H \== '\n'}.
longest_match_no_nl_([H|T]) --> [H], longest_match_no_nl_(T), {H \== '\n'}.
longest_match_no_nl_([]) --> [].

% The comment/3 predicate is used to consume a variable number of characters.
% This is useful for consuming all characters between multi-line comment delimiters.
comment --> [].
comment --> [_], comment.

% Comments.
lexer(Tokens) --> comment_single, longest_match_no_nl(_), lexer(R), {Tokens = R}.
lexer(Tokens) --> comment_multi_start, comment, comment_multi_end, lexer(R), {Tokens = R}.

% Identifiers and integer literals.
lexer(Tokens) --> longest_match_alpha(I), lexer(R), {atom_chars(IA, I), append([IA], R, Tokens)}.
lexer(Tokens) --> longest_match_int(I), lexer(R), {atom_chars(IA, I), append([IA], R, Tokens)}.

% Relational operators.
lexer(Tokens) --> ['<', '='], lexer(R), {append(['<='], R, Tokens)}.
lexer(Tokens) --> ['>', '='], lexer(R), {append(['>='], R, Tokens)}.
lexer(Tokens) --> ['=', '='], lexer(R), {append(['=='], R, Tokens)}.
lexer(Tokens) --> ['!', '='], lexer(R), {append(['!='], R, Tokens)}.

% Whitespace characters.
lexer(Tokens) --> whitespace, lexer(R), {Tokens = R}.

% Everything else.
lexer(Tokens) --> [C], lexer(R), {Tokens = [C | R]}.
lexer(Tokens) --> [], {Tokens = []}.

%% % ----------------------------------------------------------------------
%% %   PARSER
%% % ----------------------------------------------------------------------


backtrack(back(A, B)) --> ["backtrack"], ["on"], tuple(A, B), [":"].
tuple(A, B) --> ["("], [A],[","],[B],[")"].


parser(program(SL)) --> backtrack(SL).
%% statement_list(b(R)) --> backtrack(R).
%% statement_list(R) --> anything, backtrack(R), anything.
%% statement_list(_) --> anything.





 %% EVAL

eval(program(SL)) :- eval(SL).
eval(b(_, SL, _)) :- write(SL).

%% anything(c(C)) --> control_characters(CR),
%%                     \+ member(C, CR)
anything --> [].
anything --> [_], anything.
%% parser(program(SL)) --> statement_list(SL).
%% statement_list(sl(S)) --> statement(S).
%% statement_list(sl(S, SL)) --> statement(S), statement_list(SL).
%% statement(b(R, SOLUTION_REJECT, CANDIDATE, RJ_SL, SOLUTION_ACCEPT, ACC_SL, PARENT, CH_SL)) --> backtrack(R, SOLUTION_REJECT, CANDIDATE, RJ_SL, SOLUTION_ACCEPT, ACC_SL, PARENT, CH_SL).
%% statement(stmt(E)) --> expression(E), newline.
%% expression(e(E)) --> anything(E).
