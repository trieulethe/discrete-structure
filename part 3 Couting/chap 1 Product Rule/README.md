<h1> The Product Rule</h1>


- Question:  
   -  How many string of two charaters are there, if the first charater must be an uppercase letter and the second charater must be a digit ?

- Example :
   - string: A0, K7, Z9...
- Answer :
   - 26 * 10 = 260, 26 choices for first charater and 10 choices for second

> **Product Rule :** Assume a procedure consists of performing a sequence of *m* task in order. Furthermore, assume that for each  *i* = 1, 2, ..., *m*, there are N-i ways to do the i-th task, regardless of how the first *i* - 1 task were done. Then, there are N1N2...Nm ways to do the entire procedure.

- Source code:  Define

--------------------------------------------------------------

<h2>I.  Couting Bitstrings of Length n </h2>

- Question:
   - How many bitstrings of length n are there ?

- Step procedure:
   - The procedure is “write down a bitstring of length n”.
   - For i = 1, 2, . . . , n, the i-th task is “write down one bit”.

- Answer:
   - There are two ways to do the i-th task, regardless of how we did the first i − 1 tasks. Therefore, we can apply the Product Rule with N i = 2 for i = 1, 2, . . . , n, and conclude that there are N1 N2 ··· Nn = 2^n ways to do the entire procedure. As a result, the number of bitstrings of length n is equal to 2^n .

> **Theorem**: For any integer n ≥ 1, the number of bitstrings of length n
is equal to 2^n.
- Source code: Couting Bitstrings Length n

--------------------------------------------------------------

<h2>II. Couting Funtions </h2>

Section 1:

-  Question: 
   - Let m ≥ 1 and n ≥ 1 be integers, let A be a set of size m, and let B be a set of size n. How many functions *f* : A → B are there ?

-  Convention:
   - Write the set A as A = {a1, a2 , . . . , am}. Any function f : A → B is completely specified by the values f(a1), f(a2), . . . , f(am), where each such value can be any element of B. Again, we are going to apply the Product Rule.

-  Step Procedure:
   - The procedure is “specify the values f(a1), f(a2), . . . , f(am)”.
   - For i = 1, 2, . . . , m, the i-th task is “specify the value f(ai)”.
-  Note:
   - For each i, f(ai) can be any of the n elements of B. As a result, there are Ni = n ways to do the i-th task, regardless of how we did the first i − 1 tasks. By the Product Rule, there are N1 N2 · · · Nm = n m ways to do the entire procedure and, hence, this many functions f : A → B.
> **Theorem** Let m ≥ 1 and n ≥ 1 be integers, let A be a set of size m, and let B be a set of size n. The number of functions f : A → B is equal to n^m .
                     
Section 2:

- Question:
   - Recall that a function f : A → B is one-to-one if for any i and j with
   i # j, we have f(ai) # f(aj). How many one-to-one functions f : A → B
   are there?
- Note:
   - If m > n, then there is no such function. Assume m <= n. The procedure:
- Step procedure:
   - Since f(a1) can be any of the n elements of B, there are N1 = n ways to do the first task.
   - In the second task, we have to specify the value f(a2). Since the function f is one-to-one and since we have already specified f(a1), we can choose f(a2) to be any of the n - 1 elements in the set B \ {f(a1)}. As a result, there are N2 = n - 1 ways to do the second task. Note that this is true, regardless of how we did the first task.
   - In general, in the i-th task, we have to specidy the value f(ai). Since we have already specified f(a1), f(a2), ... , f(ai-1), we can choose f(ai) to be any of the n - i + 1 elements in the set. As a result, there are Ni = n − i + 1 ways to do the i-th task.

> **Theorem**  Let m ≥ 1 and n ≥ 1 be integers, let A be a set of size m, and let B be a set of size n.
   - 1. If m > n, then there is no one-to-one function f : A → B.
   - 2. If m ≤ n, then the number of one-to-one functions f : A → B is equal to:  n! / (n − m)!
---------------------------------------------------------------------------------
<h2>III. Placing Books on Shelves</h2>

- Question:
   - Let m ≥ 1 and n ≥ 1 be integers, and consider m books B1, B2 , . . . , Bm and n book shelves S1, S2, . . ., Sn . How many ways are there to place the books on the shelves?
- Note:
   - We specify for each book the shelf at which this book is placed, and
   - We specify for each shelf the left-to-right order of the books that are
   placed on that shelf.
   - Some book shelves may be empty. We assume that each shelf is large enough
   to fit all books.
- Step procedure:
   - The procedure is “place the m books on the n shelves”.
   - For i = 1, 2, . . . , m, the i-th task is “place book Bi on the shelves”.
   When placing book Bi , we can place it on the far left or far right of any shelf or between any two of the books B1 , . . . , Bi−1 that have already been placed.
- Discussion:
   - Just before we place book B1, all shelves are empty. Therefore, there
   are N1 = n ways to do the first task.
   - In the second task, we have to place book B2. Since B1 has already been placed, we have the following possibilities for placing B2 :
      - We place B2 on the far left of any of the n shelves.
      - We place B2 immediately to the right of B1.
      - As a result, there are N2 = n + 1 ways to do the second task. Note that this is true, regardless of how we did the first task.
   - In general, in the i-th task, we have to place book Bi . Since the books B1, B2 , . . . , Bi−1 have already been placed, we have the following possibilities for placing Bi :
      - We place Bi on the far left of any of the n shelves.
      - We place Bi immediately to the right of one of B1 , B2 , . . . , Bi−1.
      - As a result, there are Ni = n + i − 1 ways to do the i-th task. Note that this is true, regardless of how we did the first i − 1 tasks.




