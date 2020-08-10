# jp (jsonPath)

A simple commandline tool to find the path to a key in nested JSON structures. Reads json from stdin.

## Usage

```
Usage:
  jp [flags]

Flags:
  -h, --help         help for jp
  -k, --key string   Key to search for
  -v, --show-value   If enabled, will show the value for found path
```

## Example

```bash
$ kubectl get node node01 -ojson | jp -k osImage
.status.nodeInfo.osImage
$ kubectl get node node01 -ojson | jp -k osImage -v
.status.nodeInfo.osImage [Flatcar Container Linux by Kinvolk 2512.2.0 (Oklo)]
```

## Acknowledgements

* [cespare/jsonpath](https://github.com/cespare/jsonpath) - The source of this codebase
