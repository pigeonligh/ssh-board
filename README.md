# ssh-board
ssh-board

## Usage

```
docker run -tid -p 13579:22 --name board -v $(pwd)/data/keys:/etc/ssh/keys -v $(pwd)/data/auth:/etc/sshboard -v $(pwd)/data/data:/home/noname/data pigeonligh/ssh-board:latest
```
