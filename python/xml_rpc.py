from SimpleXMLRPCServer import SimpleXMLRPCServer
import xmlrpclib
IP = "127.0.0.1"
PORT = 6118


def hello():
    """
    a simple function, returns nothing.
    """
    return "function: hello world"


class FunctionGroup1(object):

    @staticmethod
    def func1():
        return "func1"

    @staticmethod
    def func2():
        return "func2"

class Result(object):
    def __init__(self):
        self.success = True
        self.containers = ['a','b']

class Api(object):
    def __init__(self):
        self.fg1 = FunctionGroup1()

    @staticmethod
    def hello():
        return Result()

    @staticmethod
    def add(x, y):
        return x + y


def server():
    svr = SimpleXMLRPCServer((IP, PORT), allow_none=True)
    svr.register_introspection_functions()  # so that the client can do system.listMthods()
    svr.register_multicall_functions()  # so that the client can do multicall
    svr.register_function(hello, name="default.hello")  # specify the name so that the function can be grouped
    svr.register_instance(Api(), allow_dotted_names=True)  # only support one single instance
    svr.serve_forever()


def client():
    url = "http://" + IP + ":" + str(PORT)
    proxy = xmlrpclib.ServerProxy(url)
    result = proxy.hello()
    print type(result)
    print result
    print proxy.system.listMethods()
    print proxy.system.methodHelp('hello')
    print proxy.system.methodSignature('hello')
    print proxy.fg1.func1()
    multi = xmlrpclib.MultiCall(proxy)
    multi.add(5, 1)
    multi.add(24, 11)
    try:
        for response in multi():
            print response
    except xmlrpclib.Error, v:
        print "ERROR", v


if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument('--type', required=True, choices=("server", "client"), help="type")
    args = parser.parse_args()
    if args.type == "server":
        server()
    else:
        client()
