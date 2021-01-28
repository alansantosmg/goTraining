
import pytest
from mx2 import maximo

def test_maximo1_equals_negativos():
    assert maximo(-2,-1) == -1

def test_maximo2_equals_negativos2():
    assert maximo(-1,-2) == -1

def test_maximo3_equals_zeroNegativo():
    assert maximo(0,-1) == 0

def test_maximo4_equals_negativoZero():
    assert maximo(-1,0) == 0

def test_maximo5_equals_Zero_zero():
    assert maximo(0,0) == 0

def test_maximo6_equals_positivosIguais():
    assert maximo(1,1) == 1

def test_maximo7_equals_negativosPositivos():
    assert maximo(-1,1) == 1

def test_maximo8_equals_PositivosNegativos():
    assert maximo(1,-1) == 1

def test_maximo9_equals_PositivosMaior1():
    assert maximo(1,2) == 2

def test_maximo10_equals_PositivosMaior1():
    assert maximo(2,1) == 2
