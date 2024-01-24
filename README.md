# Ecosystem of Programming Languages

## Go

### Goal
The goal of the project is to code the algorithm of Monte-Carlo.
This algorithm is used to estimate the volume of an object due to a stochastic process.
That's to say that the algorithm gets a big amount of random points in a range, and checks for each if the point is or not in the volume.
The volume is defined as inequalities, in n dimensions.
The software needs to parse the user input to gets information and then resolve

### Implementation
We implemented our code in GO Lang, using go routines (multiprocessing) to improve performances, and using a client/server structure to remotely resolve problems.
Also, we allow clients to manually prompt information or select a file.

### How to use
The format of the input must be of this type:
```
x y
~x->0:1
~y->0:1
#x>y
N=10000000
```

On the first row, the variables separated by a space.

For each variable, the client must define a range. Each range definition is written on a line, the first character is a ```~```. The variable and the range must be separated by ```->```, and the range of type ```float:float```. There is no order in the input range, even if the first float is greater than the second, the program reverse them.

Each inequality is preceded by ```#```.
The only operators allowed are ```+```, ```-```, ```/```, ```*``` and ```( )```, and one of ```>``` or ```<```.
The operational priorities apply.

The number of points is introduced due to ```N=``` prefix.