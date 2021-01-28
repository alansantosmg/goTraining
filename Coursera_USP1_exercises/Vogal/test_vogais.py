import pytest
from vogais import vogal

def testa_vogal_minuscula():
    assert vogal("a") == True
    assert vogal("u") == True

def testa_vogal_maiuscula():
    assert vogal("A") == True
    assert vogal("U") == True

def testa_consoante_minuscula():
    assert vogal("b") == False
    assert vogal("z") == False

def testa_consoante_Maiuscula():
    assert vogal("B") == False
    assert vogal("B") == False

def testa_consoante_numeros():
    assert vogal(0) == False
    assert vogal(1) == False
    assert vogal(-1) == False