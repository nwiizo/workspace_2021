# dapr-actor

using virtual actors on dapr

```bash
$ kind create cluster --config kind.yaml

$ skaffold run -f skaffold.dapr.yaml

$ skaffold dev -f skaffold.actor.yaml
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Jingle
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Put:  10
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Jingle
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Put:  100
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Put:  500
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Jingle
[actor] 2021/12/09 13:19:23 Actor:  PiggyBank / 01FPFNJDFEKHHT3GAF0961MXME  call Break
```

```bash
$ skaffold run --tail
[client] dapr client initializing for: 127.0.0.1:50001
[client] 2021/12/09 13:19:23 actor(pb: 01FPFNJDFEKHHT3GAF0961MXME).Jingle: , <nil>
[client] 2021/12/09 13:19:23 actor(pb: 01FPFNJDFEKHHT3GAF0961MXME).Jingle: じゃら, <nil>
[client] 2021/12/09 13:19:23 actor(pb: 01FPFNJDFEKHHT3GAF0961MXME).Jingle: じゃらじゃらじゃら, <nil>
[client] 2021/12/09 13:16:51 actor(pb: 01FPFNJDFEKHHT3GAF0961MXME).Break: [10 100 500], <nil>
```
