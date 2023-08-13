

# Client Test (evans)

```
brew install evans
evans proto/hello-world.proto

helloworld.Greeter@127.0.0.1:50051> call SayHello
name (TYPE_STRING) => hi
{
  "message": "hi"
}
```