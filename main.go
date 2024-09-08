package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
)

func main() {
	var jsondata = []byte(`{
  "apiVersion": "rbac.authorization.k8s.io/v1",
  "kind": "Role",
  "metadata": {
    "namespace": "default",
    "name": "configmap-updater"
  },
  "rules": [
    {
      "apiGroups": [
        ""
      ],
      "resources": [
        "configmaps"
      ],
      "resourceNames": [
        "my-configmap"
      ],
      "verbs": [
        "update",
        "get"
      ]
    },
    {
      "apiGroups": [
        ""
      ],
      "resources": [
        "secrets"
      ],
      "resourceNames": [
        "my-secret"
      ],
      "verbs": [
        "update",
        "get"
      ]
    }
  ]
}`) // your data
	var data interface{}
	err := json.Unmarshal(jsondata, &data)
	if err != nil {
		panic(err)
	}
	result, err := jmespath.Search("rules[?contains(resources, 'secrets') == `true`].resourceNames[]", data)
	if err != nil {
		panic(err)
	}
	jb, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jb))
}
