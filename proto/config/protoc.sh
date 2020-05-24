Basepath=$(cd `dirname $0`; pwd)
#echo ${Basepath}
echo "---start---"
##############################
source ~/.bash_profile
# mq
cd ${Basepath}
protoc --go_out=plugins=grpc:. *.proto

echo "---all finish---"
# end