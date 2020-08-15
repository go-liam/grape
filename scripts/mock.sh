#!/bin/bash
set -e

shell_dir=$(dirname $0)
cd ${shell_dir}
cd ../

mockgen -destination ./test/mock/www/mock_site/site.go -package mock_site -source ./internal/pkg/data/home_site/site/interface.go

