# go-eg

## usage

```sh
helm install go-eg ./charts \
  -n go-eg \
  --set host='go.eg' \
  --set ui.image=core.harbor.domain/library/xmh-ui \
  --set ui.tag=v1 \
  --set auther.image=core.harbor.domain/library/xmh-auther \
  --set auther.tag=v1 \
  --set cacher.image=core.harbor.domain/library/xmh-cacher \
  --set cacher.tag=v1
```