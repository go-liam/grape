#!/bin/bash
set -e

shell_dir=$(dirname $0)
cd ${shell_dir}
cd ../



# account
mockgen -destination ./test/mock/account/mock_bind/mock.go -package mock_bind -source ./internal/pkg/data/account/bind/interface.go
mockgen -destination ./test/mock/account/mock_user/mock.go -package mock_user -source ./internal/pkg/data/account/user/interface.go

#config
mockgen -destination ./test/mock/config/mock_config/mock.go -package mock_config -source ./internal/pkg/data/config/interface.go

#www home_site
mockgen -destination ./test/mock/www/mock_menu/mock.go -package mock_menu -source ./internal/pkg/data/home_site/menu/interface.go
mockgen -destination ./test/mock/www/mock_notice/mock.go -package mock_notice -source ./internal/pkg/data/home_site/notice/interface.go
mockgen -destination ./test/mock/www/mock_page/mock.go -package mock_page -source ./internal/pkg/data/home_site/page/interface.go
mockgen -destination ./test/mock/www/mock_site/mock.go -package mock_site -source ./internal/pkg/data/home_site/site/interface.go


#money
mockgen -destination ./test/mock/money/mock_change/mock.go -package mock_change -source ./internal/pkg/data/money/change/interface.go
mockgen -destination ./test/mock/money/mock_pay/mock.go -package mock_pay -source ./internal/pkg/data/money/pay/interface.go
mockgen -destination ./test/mock/money/mock_recharge/mock.go -package mock_recharge -source ./internal/pkg/data/money/recharge/interface.go

# rbac
mockgen -destination ./test/mock/rbac/mock_power/mock.go -package mock_power -source ./internal/pkg/data/rbac/power/interface.go
mockgen -destination ./test/mock/rbac/mock_role/mock.go -package mock_role -source ./internal/pkg/data/rbac/role/interface.go
mockgen -destination ./test/mock/rbac/mock_user/mock.go -package mock_user -source ./internal/pkg/data/rbac/user/interface.go

# region
mockgen -destination ./test/mock/region/mock_region/mock.go -package mock_region -source ./internal/pkg/data/region/interface.go

# sylog
mockgen -destination ./test/mock/sylog/mock_sylog/mock.go -package mock_sylog -source ./internal/pkg/data/sylog/interface.go

