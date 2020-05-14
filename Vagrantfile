# -*- mode: ruby -*-
# vi: set ft=ruby :

ip = '192.168.33.12'
cpus = 2
memory = 1024 * 6

def fail_with_message(msg)
  fail Vagrant::Errors::VagrantError.new, msg
end


Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.box_url = "https://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  config.vm.network :private_network, ip: ip, hostsupdater: 'skip'
  config.vm.hostname = 'mn.local'

  config.vm.provider 'virtualbox' do |vb|
    vb.name = config.vm.hostname
    vb.customize ['modifyvm', :id, '--cpus', cpus]
    vb.customize ['modifyvm', :id, '--memory', memory]

    # Fix for slow external network connections
    vb.customize ['modifyvm', :id, '--natdnshostresolver1', 'on']
    vb.customize ['modifyvm', :id, '--natdnsproxy1', 'on']
  end

  config.vm.provision "shell", path: "deploy/vagrant/provision.sh"
end
