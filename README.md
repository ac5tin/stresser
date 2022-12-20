# Stresser
Minimalistic Stress testing tool



## Usage
**TLDR**
```sh
stresser -c 100 -t 1000 -f ./req.toml
```

Export requests to csv
```sh
stresser -c 100 -t 1000 -f ./req.toml
```

Export csv of errors only
```sh
stresser -c 100 -t 1000 -f ./req.toml -e ./resp.csv -err
```
