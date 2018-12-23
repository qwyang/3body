class Queue(object):
    def __init__(self, num):
        self.queue_num = num
        self.columns = []
        self.found = False
        self.count = 0
        for i in range(num):
            self.columns.append(0)

    def showAnswer(self):
        for col in range(0, self.queue_num):
            for pos in range(0, self.queue_num):
                if self.columns[col] == pos:
                    print 'Q',
                else:
                    print '-',
            print
        print
        self.count += 1

    def checkValid(self, col, pos):
        for i in range(col):
            p = self.columns[i]
            if p == pos:
                return False
            if abs(i - col) == abs(p - pos):
                return False
        return True

    def backTrack(self, col):
        if col >= self.queue_num:
            self.showAnswer()
            self.found = True
        else:
            for p in range(0, self.queue_num):
                if self.checkValid(col, p):
                    self.columns[col] = p
                    self.backTrack(col + 1)
                    # if self.found == True:
                    #    break
                    # else:
                    #    self.backTrack(col - 1)


if __name__ == "__main__":
    q = Queue(8)
    q.backTrack(0)
    print "total solutions: %d" % q.count
