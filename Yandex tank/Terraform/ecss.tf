data "sbercloud_images_image" "ubuntu_image" {
  name = "Ubuntu 20.04 server 64bit"
  most_recent = true
}

resource "sbercloud_compute_instance" "ecs_for_tank" {
  name = "ecs-for-tank"
  image_id = data.sbercloud_images_image.ubuntu_image.id
  flavor_id = "s6.large.2"
  security_groups = ["sg-1"]
  availability_zone = "ru-moscow-1a"
  key_pair = "KeyPair-borodin"

  system_disk_type = "SAS"
  system_disk_size = 20

  network {
    uuid = "11257ffd-fdac-4ef8-914d-8a30dd59958a"
  }
}

resource "sbercloud_vpc_eip" "eip_for_tank" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "eip-for-tank"
    size        = 5
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "sbercloud_compute_eip_associate" "associated" {
  public_ip   = sbercloud_vpc_eip.eip_for_tank.address
  instance_id = sbercloud_compute_instance.ecs_for_tank.id
}