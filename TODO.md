1.  get comment endpoint: if comment not found, return message to client.

2.  Env variables:: remove static definitions, move into .env file and inject into docker file and task file at build time. -> https://github.com/spf13/viper

3.  env variables:: check if all are defined at build time, throw error if one is missing

4.  Create auth endpoints for user creation
