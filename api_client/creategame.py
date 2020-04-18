from client import Client


if __name__ == '__main__':
    client = Client('http://localhost:5000/')
    response = client.create_game(rows=5, columns=5)
    print(response)
   # response = client.reveal_cell(response['gameID'], rows=0, columns=0)
    print(response)

