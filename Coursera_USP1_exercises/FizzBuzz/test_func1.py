import pytest
from func1 import fizzbuzz
def teste_1(): 
    assert fizzbuzz(3) == "Fizz"
def teste_3(): 
    assert fizzbuzz(5) == "Buzz"
def teste_2(): 
    assert fizzbuzz(30) == "FizzBuzz"
def teste_4(): 
    assert fizzbuzz(4) == 4