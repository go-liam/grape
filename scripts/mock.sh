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

