# Go RabbitmMQ with Gin

### Run app:

You need to copy .env.example, change all secret keys. And run the app
```bash
cp .env.example .env
make start
```

To generate secret keys you can use for example this command:

```bash
openssl rand -base64 42
```

### Note

Later I will add more services for this toy architecure. To show ( and learn ) how to work with another types of RabbitmMQ connections
