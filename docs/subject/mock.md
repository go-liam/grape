#mock

```shell script
chmod 755 ./scripts/mock.sh

./scripts/mock.sh
```

example:

```shell admin

mockgen -destination ./test/mock/admin/mock_admin.go -package mock_admin -source ./internal/app/boxapi/services/data/admin/interface.go

```
