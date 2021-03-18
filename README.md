go-simple-blockchain
==========================

The simplest blockchain implementd with Golang.

How to use it
---------------

[LeafServer](https://github.com/name5566/leafserver) is a game server developped with Leaf. Let's start with it.

Download the source code of LeafServer：

```
git clone https://github.com/iamcmnut/go-simple-blockchain
```

Install dependencies:

```
go get
```

Run server：

```
go run main.go
```

You will get below screen output if everything is successful.

```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.2.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1234
```

Your server will be run on this url: http://localhost:1234

### API Document

* Get all blocks : `GET /`

* Create new block : `POST /`
```json
{
    "data": "A transfer 100BTC to B"
}
```
