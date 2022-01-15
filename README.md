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
* a) we have a square and recursively build all paths from bottom left corner to upper right corner. This path should not cross over main diagonal and number of steps on each diagonal should not overcome 2, otherwise we would get bracket sequences with 3 or more childs for node which is not represented in math forms.<br><br>

* b) each path represent bracket sequence where every left arrow (&#10142;) equal to opened bracket and up arrow (&#129045;) is closed bracket.<br><br>
![alt tag](https://raw.githubusercontent.com/mcfly722/formulator/main/doc/squares.svg)

To simplify all the process there are function <b>GetNextBracketsSequence(bracketSequence, maxChilds)</b> that return next brackets sequence based on current brackets sequence (Function build it own square with path and recursively obtain next path that match to <b>maxChilds</b> condition. Also it extends number of brackets if your specify last one sequence for this square size).
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

## step 2 - Iterate over Previous Values (pv) forms (0,1,x)

 From new tree form we know how many constants positions we have. (In sequence it is <b>"()"</b> brackets pairs)<br>
 We also know that previous value (0,1,x) should appear at least one time, otherwise iterations has no reason to be.

 Examples:<br>

* <b>pv = 0<br></b>
 Computing exp(z) using [Euler formula](https://en.wikipedia.org/wiki/Euler%27s_formula):<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/6a91595ef0946463456b2d0184bdcdb2ae9da7a2)<br>
 last recursive sum should be equal to 0. If we take pv = 1, we will get wrong final answer &#8776; exp(z)+1
<br><br><br>
* <b>pv = 1<br></b>
 Computing Pi using [Wallis product](https://en.wikipedia.org/wiki/Wallis_product):<br>
![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/df59bf8aa67b6dff8be6cffb4f59777cea828454)<br>
last product could not be equal to 0, otherwise all final product will be equal to 0 too
<br><br><br>
* <b>pv = x<br></b>
 Computing square root using [Geron iteration formula](https://ru.wikipedia.org/wiki/%D0%98%D1%82%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D0%B0%D1%8F_%D1%84%D0%BE%D1%80%D0%BC%D1%83%D0%BB%D0%B0_%D0%93%D0%B5%D1%80%D0%BE%D0%BD%D0%B0):<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/9935d6f7061161b29325d712518fb58496f58bfb)<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/cd0d9bc3389f73d8501bfef1303b06246d81f771)<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/a8788bf85d532fa88d1fb25eff6ae382a601c308) could not be 0 or 1 and should be equal to initial function argument x
<br><br><br>

## step 3 Itarate through number of Previous Values (pv)
Number of pv's should not be equal to zero, so we iterate it from 1 to <b>maxPVs</b>

## step 4 Constants with Previous Value (pv) combination
To combine pv's with constants in all tree leafs here used lexicographic method (see [The Art of computer Programming](https://www.kcats.org/csci/464/doc/knuth/fascicles/fasc3a.pdf) page 21)

test sample 3 from 8:

```
\catalan> go test -v .
...
=== RUN   Test_CombinationNKNext
--- PASS: Test_CombinationNKNext (0.00s)
    Catalan_test.go:210:  1 ***..... = 224
    Catalan_test.go:210:  2 **.*.... = 208
    Catalan_test.go:210:  3 **..*... = 200
    Catalan_test.go:210:  4 **...*.. = 196
    Catalan_test.go:210:  5 **....*. = 194
    Catalan_test.go:210:  6 **.....* = 193
    Catalan_test.go:210:  7 *.**.... = 176
    Catalan_test.go:210:  8 *.*.*... = 168
    Catalan_test.go:210:  9 *.*..*.. = 164
    Catalan_test.go:210: 10 *.*...*. = 162
    Catalan_test.go:210: 11 *.*....* = 161
    Catalan_test.go:210: 12 *..**... = 152
    Catalan_test.go:210: 13 *..*.*.. = 148
    Catalan_test.go:210: 14 *..*..*. = 146
    Catalan_test.go:210: 15 *..*...* = 145
    Catalan_test.go:210: 16 *...**.. = 140
    Catalan_test.go:210: 17 *...*.*. = 138
    Catalan_test.go:210: 18 *...*..* = 137
    Catalan_test.go:210: 19 *....**. = 134
    Catalan_test.go:210: 20 *....*.* = 133
    Catalan_test.go:210: 21 *.....** = 131
    Catalan_test.go:210: 22 .***.... = 112
    Catalan_test.go:210: 23 .**.*... = 104
    Catalan_test.go:210: 24 .**..*.. = 100
    Catalan_test.go:210: 25 .**...*. = 98
    Catalan_test.go:210: 26 .**....* = 97
    Catalan_test.go:210: 27 .*.**... = 88
    Catalan_test.go:210: 28 .*.*.*.. = 84
    Catalan_test.go:210: 29 .*.*..*. = 82
    Catalan_test.go:210: 30 .*.*...* = 81
    Catalan_test.go:210: 31 .*..**.. = 76
    Catalan_test.go:210: 32 .*..*.*. = 74
    Catalan_test.go:210: 33 .*..*..* = 73
    Catalan_test.go:210: 34 .*...**. = 70
    Catalan_test.go:210: 35 .*...*.* = 69
    Catalan_test.go:210: 36 .*....** = 67
    Catalan_test.go:210: 37 ..***... = 56
    Catalan_test.go:210: 38 ..**.*.. = 52
    Catalan_test.go:210: 39 ..**..*. = 50
    Catalan_test.go:210: 40 ..**...* = 49
    Catalan_test.go:210: 41 ..*.**.. = 44
    Catalan_test.go:210: 42 ..*.*.*. = 42
    Catalan_test.go:210: 43 ..*.*..* = 41
    Catalan_test.go:210: 44 ..*..**. = 38
    Catalan_test.go:210: 45 ..*..*.* = 37
    Catalan_test.go:210: 46 ..*...** = 35
    Catalan_test.go:210: 47 ...***.. = 28
    Catalan_test.go:210: 48 ...**.*. = 26
    Catalan_test.go:210: 49 ...**..* = 25
    Catalan_test.go:210: 50 ...*.**. = 22
    Catalan_test.go:210: 51 ...*.*.* = 21
    Catalan_test.go:210: 52 ...*..** = 19
    Catalan_test.go:210: 53 ....***. = 14
    Catalan_test.go:210: 54 ....**.* = 13
    Catalan_test.go:210: 55 ....*.** = 11
    Catalan_test.go:210: 56 .....*** = 7
PASS
```
Main function that generates next combination based on previous one is <b>CombinationNKNext(input string) (string, bool, error)</b>

 ## step 5 - Iterate over all Functions combinations
 Number of functions in tree is known - it is all nodes with only one child.<br>
 Unfortunately, trees contains different number of functions, so we cannot use simple <b>for(){}</b> inside other <b>for(){}</b>. To iterate variative number of <b>for's</b>, we also use <b>for(){}</b> through recursive function.

 ## step 6 - Iterate over all Operator combinations
 Same method (from step 5) used here.

 <br><br><br><br><br><br><br><br><br>

  TODO:


 ## step 6 - Build calculation sequence
 ## step 7 - Calculate
 ## step 8 - Calculating [Variance](https://en.wikipedia.org/wiki/Variance) of points
