# Gitodo

This is my todoApp using phoenix + liveview (Realtime apps with steroids)

### Start app
```
$ docker-compose up -d
```

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.

### Execute elixir trough docker-compose
```
$ docker-compose run gitodo_web command

// ex:

$ docker-compose run gitodo_web mix ecto.migrate
```
