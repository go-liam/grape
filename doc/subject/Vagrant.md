# Vagrant 虚拟机


## 快速搭建虚拟机
https://vagrantup.com/
### 新建 Vagrantfile 文件，然后运行：
$ vagrant up
$ vagrant reload
$ vagrant ssh
$ exit
### 删除：
$ vagrant destroy
同步数据：
vagrant rsync 同步文件或目录，
vagrant rsync-auto
挂起 vagrant suspend
恢复 vagrant resume
vagrant halt    # 关闭虚拟机
### 换root user
$ su
$ vagrant
==> /VirtualBox/vagrant/centos7_test/ => /vagrant
==>  VirtualBox/vagrant/centos7_test/share/ => /db
账户 密码 :默认的登录用户名是vagrant，密码一样
vagrant vagrant
root vagrant
### 网络设置：
https://www.ityoudao.cn/posts/vagrant-network/
virtualbox 网络设置
Bridged Adapter模式 
这样主机和子子可以互相访问