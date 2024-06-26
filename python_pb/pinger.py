import grpc
import pickle
import time
import numpy as np 
import J2kResultsHub_pb2
import J2kResultsHub_pb2_grpc
import ResultsHub as rh

def run():
    # Setup gRPC channel and stub
    channel = grpc.insecure_channel('localhost:50051')
    stub = J2kResultsHub_pb2_grpc.ResultsHubStub(channel)
    # Testing SayHello
    response = stub.SayHello(J2kResultsHub_pb2.HelloRequest(senderId='your_sender_id', message='your_message'))
    print("Pong server responded: " + response.message)

    # These variables can be any type!
    var1 = {'name': 'var1', 'data': [1, 2, 3]}
    list2 = [1, 2, 3, 4]
    var2 = np.array(list2)

    # Using the ResultsHubSubmission class
    submission = rh.ResultsHubSubmission(cell_number=1)
    submission.addVar('var1', var1)
    submission.addVar('var2', var2)
    submission.submit()
    print("Submission Success.")

    # Retrieve the variable
    fetched_var = rh.fetchVarResult('var2', varAncestorCell=1)
    print(f"fetched variable: '{fetched_var}'")

    # Requiring a file that supposed to be previously accessed by cell1
    rh.waitForCell(1)
    print(f"Wait for succesfully returned\n")

if __name__ == '__main__':
    run()
