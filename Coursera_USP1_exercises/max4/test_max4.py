import pytest
from max4 import maximo
def teste_x_maior():
    assert maximo(30,14,10) == 30
def teste_y_maior():
    assert maximo(2,3,1) == 3
def teste_z_maior():
    assert maximo(0,-1,1) == 1
