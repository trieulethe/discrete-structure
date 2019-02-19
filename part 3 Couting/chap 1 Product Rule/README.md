<h1> The Product Rule</h1>


- Question:  
   -  How many string of two charaters are there, if the first charater must be an uppercase letter and the second charater must be a digit ?

- Example :
   - string: A0, K7, Z9...
- Answer :
   - 26 * 10 = 260, 26 choices for fisrt charater and 10 choices for second

> **Product Rule :** Assume a procedure consists of performing a sequence of *m* task in order. Furthermore, assume that for each  *i* = 1, 2, ..., *m*, there are N-i ways to do the i-th task, regardless of how the first *i* - 1 task were done. Then, there are N1N2...Nm ways to do the entire procedure.

- Source code:  Define

--------------------------------------------------------------

<h3>I.  Couting Bitstrings of Length n :</h3>

- Question:
   - How many bitstrings of length n are there ?

- Step procedure:
   - The procedure is “write down a bitstring of length n”.
   - For i = 1, 2, . . . , n, the i-th task is “write down one bit”.

- Answer:
   - There are two ways to do the i-th task, regardless of how we did the first i − 1 tasks. Therefore, we can apply the Product Rule with N i = 2 for i = 1, 2, . . . , n, and conclude that there are N1 N2 ··· Nn = 2^n ways to do the entire procedure. As a result, the number of bitstrings of length n is equal to 2^n .

> Theorem: For any integer n ≥ 1, the number of bitstrings of length n
is equal to 2^n.
- Source code: Couting Bitstrings Length n

--------------------------------------------------------------









