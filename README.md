# Interpreter
A golang interpreter implementation loosely following Thorsten Ball's Monkey

# Purpose
In an effort to gain a higher level of understanding of developing complex projects in golang, produce a complete interpreter.  This
will require a great deal of work with various facets of golang and help spur increased levels of skill in using the language.

# Potential TODO Items
1.  Currently, token types are defined as strings, making them easy to work with.  Refactoring this to use a wider variety of types
(int, byte, pointer, etc) could lead to better performance.
2.  Add more detailed information to lexer logging, such as line numbers, file names, detailed error reports, etc
3.  Add numbers and other characters to the alphabetic listing if what is allowed in identifiers and keywords
