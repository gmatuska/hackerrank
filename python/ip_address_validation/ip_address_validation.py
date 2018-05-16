#!/bin/python3

import re


class IpAddressValidation:

    def is_ip4(self, ip_address):
        pattern = '^((25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})$'
        return len(re.findall(pattern,ip_address)) > 0

    def is_ip6(self, ip_address):
        pattern = '^([0-9abcdef]{1,4}:){7}[0-9abcdef]{1,4}$'
        return len(re.findall(pattern,ip_address)) > 0


def main():
    results = []
    ipAddress = IpAddressValidation()
    n = int(input())
    for _ in range(n):
        ip = input().strip()
        if ipAddress.is_ip4(ip):
            results.append('IPv4')
        elif ipAddress.is_ip6(ip):
            results.append('IPv6')
        else:
            results.append('Neither')
    print(*results, sep='\n')

main()
