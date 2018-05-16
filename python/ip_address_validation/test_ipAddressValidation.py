from unittest import TestCase
from ip_address_validation import IpAddressValidation


class TestIpAddressValidation(TestCase):

    def setUp(self):
        self.ipAddress = IpAddressValidation()

    def test_is_ip4(self):
        self.assertTrue(self.ipAddress.is_ip4('255.255.255.1'))

    def test_is_ip6(self):
        self.assertTrue(self.ipAddress.is_ip6('2001:0db8:0000:0000:0000:ff00:0042:8329'))
