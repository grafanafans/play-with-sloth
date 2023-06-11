# play-with-sloth

An example with sloth cli generate and dashboard.

## Usage

1. clone code and build app

```bash
git clone git@github.com:grafanafans/play-with-sloth.git
cd play-with-sloth
```

2. generate rules with `sloth generate` command

```
make generate
```

3. start app

```
make start
```

4. change `myservice` error rate

```
curl http://localhost:8080/errrate?value=0.005
```

5. visit sloth slo dashboard 

Go `http://localhost:3000` page and new dashboard with import by ids(`14348` and `14643`).

6. stop app

```
make down
```
