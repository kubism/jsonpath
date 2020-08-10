# jp (jsonPath)

A simple commandline tool to find the path to a key in nested JSON structures. Reads json from stdin.

## Usage

```
Usage:
  jp [flags]

Flags:
  -h, --help          help for jp
  -i, --ignore-case   If enabled, will ignores case when matching key
  -k, --key string    Key to search for
  -t, --show-type     If enabled, will show the type of the value for found path
  -v, --show-value    If enabled, will show the value for found path
```

## Example

```bash
$ kubectl get node node01 -ojson | jp -k osImage
.status.nodeInfo.osImage
$ kubectl get node node01 -ojson | jp -k osImage -v
.status.nodeInfo.osImage:Flatcar Container Linux
$ kubectl get node node01 -ojson | jp -k osImage -v -t
.status.nodeInfo.osImage:Flatcar Container Linux:string
$ kubectl get node node01 -ojson | jp -k osImage -t
.status.nodeInfo.osImage:string
kubectl get node node01 -ojson | ./jp.linux -k containerruntimeversion -v -i
.status.nodeInfo.containerRuntimeVersion:docker://47.1.1
```

## Releases

See the [releases](https://github.com/kubism/jsonpath/releases) to download the binary.

## Acknowledgements

* [cespare/jsonpath](https://github.com/cespare/jsonpath) - The source of this codebase
