import socket
import struct


class NamePool(object):
    def __init__(self, prefix):
        self.prefix = prefix
        self.index = 1

    def alloc(self):
        current = self.index
        self.index += 1
        return self.prefix + str(current)


class MacGen(object):
    def __init__(self, mac_range):
        mac1, mac2 = mac_range.split('-')
        self.start = self._atoi(mac1)
        self.end = self._atoi(mac2)
        self.current = self.start

    def available(self):
        return self.end - self.current >= 0

    def alloc(self):
        if self.available():
            mac = self.current
            self.current += 1
            return self._itoa(mac)
        else:
            raise Exception("sorry, all mac addresses have been allocated.")

    @staticmethod
    def _atoi(mac):
        return int(mac.replace(':', ''), 16)

    @staticmethod
    def _itoa(mac):
        mac_hex = '%x' % mac
        x = [s.encode('hex') for s in mac_hex.decode('hex')]
        remain = 6 - len(x)
        for _ in range(remain):
            x.insert(0,'00')
        return ":".join(x)


class IpGen(object):
    def __init__(self, ip_range):
        ip1, ip2 = ip_range.split('-')
        self.start = self._atoi(ip1)
        self.end = self._atoi(ip2)
        self.current = self.start

    def available(self):
        return self.end - self.current >= 0

    def alloc(self):
        if self.available():
            ip = self.current
            self.current += 1
            return self._itoa(ip)
        else:
            raise Exception("sorry, all ips have been allocated.")

    def _atoi(self, ip):
        d = struct.unpack('i', socket.inet_aton(ip))[0]
        return socket.ntohl(d)

    def _itoa(self, ip):
        d = struct.pack('i', socket.htonl(ip))
        return socket.inet_ntoa(d)


if __name__ == "__main__":
    pool = IpGen("192.168.188.1-192.168.188.10")
    while pool.available():
        print pool.alloc()

    pool = MacGen("00:54:00:00:00:01-00:54:00:00:00:11")
    while pool.available():
        print pool.alloc()

    pool = NamePool("vm")
    for i in range(10):
        print pool.alloc()

