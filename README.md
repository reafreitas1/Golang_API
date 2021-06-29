**It's working!**
![Captura de ecrã 2021-06-29 161652](https://user-images.githubusercontent.com/79333175/123824488-fe4d5480-d8f5-11eb-83e5-1234e91a1f64.jpg)

- Func http call method HandleFunc
- Server route: "/"
 - 2 args:
  1. w ResponseWriter: user response
  2. r Request: request that was made, can get customer information
- ListenAndServe:
 - 2 args:
  1. localhost (port)
  2. nil: DefaultServerMux

