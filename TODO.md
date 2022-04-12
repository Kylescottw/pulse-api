1.  get comment endpoint: if comment not found, return message to client.

2.  env variables:: check if all are defined at build time, throw error if one is missing:: https://github.com/spf13/viper::

3.  Create auth endpoints for user creation

4.  taskfile: remove staticly defined env variables, inject from .env file. right now the taskfile will not pass .env into docker-compose up command. experiment with writing a test that access .env varaible and prints to console. - IN PROGRESS
