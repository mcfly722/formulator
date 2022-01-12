# formulator

 Current status - <b>in progress</b><br>

 # idea
Every mathematical function f(x) could be represented as [Abstract Syntax Tree (AST)](https://en.wikipedia.org/wiki/Abstract_syntax_tree) where nodes of this tree could be an<br>
 - <b>operators</b> (has two input arguments): <b>*, +, -, /, mod, pow(x,p)</b>
 <br>or<br>
 - <b>functions</b> (has one input argument): <b>sqrt(x), sin(x), x!, log(x), exp(x),round(x)</b>

Complex formulas which contains [&#8721;](https://en.wikipedia.org/wiki/Summation),[&#8719;](https://en.wikipedia.org/wiki/Multiplication),[&#8970;...&#8971;](https://en.wikipedia.org/wiki/Continued_fraction), or other sequence operator could be represented as recursive function which accepts previous calculated value (<b>pv</b>), index(i) and argument(x) as input parameters.

f.e.

![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/6a91595ef0946463456b2d0184bdcdb2ae9da7a2) ([Euler formula](https://en.wikipedia.org/wiki/Euler%27s_formula))

recursive function would be = <b>x^i/i!+pv0</b> (i=...,6,5,4,3,2,1,0)
(<b>pv0</b> means what first value would be equal to 0)

# purpose
Is to create distributed computing program module which will brute force over all possible recursive functions and well known constants and found recursive function which most closely to the original set of (x,y) pairs (unknown function).

Instrument could be used to find currently unknown solutions for integrals, irrational constants and sequences (f.e. prime numbers).

 # how it works
todo
