Basepath=$(cd `dirname $0`; pwd)
#echo ${Basepath}
echo "---start---"
##############################

# mq
cd ${Basepath}
protoc --go_out=plugins=grpc:. *.proto
echo "---proto finish---"

# 
echo "---all finish---"
# end