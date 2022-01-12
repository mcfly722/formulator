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

recursive function would be:<br><br> ![alt tag](https://chart.googleapis.com/chart?cht=tx&chl=z^n/n!%2bpv0) (n=...,6,5,4,3,2,1,0)
(<b>pv0</b> means what first value would be equal to 0)

# purpose
Is to create distributed computing program module which will brute force over all possible recursive functions and well known constants and found recursive function which most closely to the original set of (x,y) pairs (unknown function).

Instrument could be used to find currently unknown solutions for integrals, irrational constants and sequences (f.e. prime numbers).

 # how it works
 ## step 1 - Iterate over trees
We need trees with only one or two childs for any node.
Nodes without childs represents Constants/Arguments or recursive values (pv\*) of previous iteration.
To iterate over all possible trees there used [Catalan method](https://en.wikipedia.org/wiki/Catalan_number).<br><br>
basic steps are:<br>
* a) we have a square and recursively build all paths from bottom left corner to upper right corner. This path should not cross over main diagonal and number of steps on each diagonal should not overcome 2, otherwise we would get bracket sequences with 3 or more childs for node which is not represented in math forms.<br>
![alt tag](https://raw.githubusercontent.com/mcfly722/formulator/main/doc/Catalan_number_4x4_grid_example.svg)

* b) each path represent bracket sequence where every left arrow (&#10142;) equal to opened bracket and up arrow (&#129045;) is closed bracket.<br><br>
For upper squares it would be:
```
(((())))
((()()))
((())())
(()(()))
...
()((()))
(()()()) - excluded (has 3 childs for one node)
((()))()
()(()())
(())(())
(()())()
...
(())()() - excluded (has 3 childs for one node)
()(())() - excluded (has 3 childs for one node)
()()(()) - excluded (has 3 childs for one node)
()()()() - excluded (has 4 childs for one node)
```

To simplify all the process there are function <b>GetNextBracketsSequence(bracketSequence, maxChilds)</b> that return next sequence based on current sequence (Function build it own square with path and recursively obtain next path that match to <b>maxChilds</b> condition.Also it extends number of brackets if your specify last one sequence for this square size).
```
ZeroOneTwoTree\go test . -v
...
=== RUN   Test_IterateOverZeroOneTwoTrees
--- PASS: Test_IterateOverZeroOneTwoTrees (0.00s)
    ZeroOneTwoTree_test.go:145:   1)   0 () -> ()()
    ZeroOneTwoTree_test.go:145:   2)   2 ()() -> (())
    ZeroOneTwoTree_test.go:145:   3)   1 (()) -> ()(())
    ZeroOneTwoTree_test.go:145:   4)   2 ()(()) -> (()())
    ZeroOneTwoTree_test.go:145:   5)   2 (()()) -> (())()
    ZeroOneTwoTree_test.go:145:   6)   2 (())() -> ((()))
    ZeroOneTwoTree_test.go:145:   7)   1 ((())) -> ()(()())
    ZeroOneTwoTree_test.go:145:   8)   2 ()(()()) -> ()((()))
    ZeroOneTwoTree_test.go:145:   9)   2 ()((())) -> (()())()
    ZeroOneTwoTree_test.go:145:  10)   2 (()())() -> (()(()))
    ZeroOneTwoTree_test.go:145:  11)   2 (()(())) -> (())(())
    ZeroOneTwoTree_test.go:145:  12)   2 (())(()) -> ((()()))
    ZeroOneTwoTree_test.go:145:  13)   2 ((()())) -> ((())())
    ZeroOneTwoTree_test.go:145:  14)   2 ((())()) -> ((()))()
    ZeroOneTwoTree_test.go:145:  15)   2 ((()))() -> (((())))
    ZeroOneTwoTree_test.go:145:  16)   1 (((()))) -> ()(()(()))
    ZeroOneTwoTree_test.go:145:  17)   2 ()(()(())) -> ()((()()))
```

* c) based on bracket sequence we build required form of tree
![alt tag](https://raw.githubusercontent.com/mcfly722/formulator/main/doc/exp.svg)

 ## step 2 - Iterate over Constants
 From new tree form we know how many constants positions we have. (In sequence it is <b>"()"</b> brackets pairs)<br>
 We also know that previous value (0,1,x) should appear at least one time, otherwise iterations has no reason to be.
 TODO

 ## step 3 - Iterate over Operators and Functions

 ## step 4 - Build calculation sequence
 ## step 5 - Calculate
 ## step 6 - Calculating [Variance](https://en.wikipedia.org/wiki/Variance) of points
