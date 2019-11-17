
class Node:
    def __init__(self, dataval=None):
        self.dataval = dataval
        self.nextval = None

class SLinkedList:
    def __init__(self):
        self.headval = None

    def atbeginning(self,num,added):
        AtBegining(self,added)
    

# Print the linked list
    def listprint(self):
        printval = self.headval
        while printval is not None:
            print (printval.dataval)
            printval = printval.nextval


      

list = SLinkedList()
list.headval = Node("Mon")
e2 = Node("Tue")
e3 = Node("Wed")

list.headval.nextval = e2
e2.nextval = e3

def AtBegining(self,newdata):
    NewNode = Node(newdata)
    NewNode.nextval = self.headval
    self.headval = NewNode

list.atbeginning(2,"Sun")

list.listprint()
