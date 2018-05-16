#!/bin/python3

from unittest import TestCase
from alien_username import AlienUsername


class TestAlienUsername(TestCase):

    def test_is_valid_alien_name(self):
        pass
        alien = AlienUsername()
        self.assertTrue(alien.is_valid_alien_name('_09090909abcD0') > 0)
