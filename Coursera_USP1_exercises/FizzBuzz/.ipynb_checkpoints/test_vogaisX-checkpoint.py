import pytest
from vogaisX import vogalx
from FizzBuzzFunc import fizzbuzz

def testa_vogal_minuscula():
    assert vogalx("a") == True
    assert vogalx("u") == True

def testa_vogal_maiuscula():
    assert vogalx("A") == True
    assert vogalx("U") == True

def testa_consoante_minuscula():
    assert vogalx("b") == False
    assert vogalx("z") == False

def testa_consoante_Maiuscula():
    assert vogalx("B") == False
    assert vogalx("B") == False

def testa_consoante_numeros():
    assert vogalx(0) == False
    assert vogalx(1) == False
    assert vogalx(-1) == False

def dnot35():
    assert fizzbuzz(15) == 15