# rebalance-tool
Interactive CLI utility to realign weightings of portfolio assets

## Usage
Specify whatever asset classes you'd like with a configuration file, e.g.:
```json
// conf.json
{
    "smallCap": 30,
    "largeCap": 30,
    "intl": 30,
    "bond": 10
}
```
Then, pass the configuration file to the tool via the `--config` flag:
```console
./rebalance_tool --config conf.json
```

