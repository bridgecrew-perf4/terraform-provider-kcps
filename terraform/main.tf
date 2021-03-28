terraform {
  required_version = ">= 0.12.7"
  required_providers {
    kcps = {
      source = "hashicorp.com/prd/kcps"
      version = "~> 1.0.0"
    }
  }
}

# 作成するVMの数（この値はポートフォワーディングのルールを作成する際などにも使う）
variable "_count" {
  default = 1
}

# SSH接続に使用するポート
variable "ports" {
  default = {
    public_port  = 1000
    private_port = 22
  }
}

# SSH接続をするクライアントのグローバルIP
variable "cidr" {
  default = "159.28.152.155/32"
}

# 利用可能なZoneのデータを取得
data "kcps_zone" "ex" {
}

# 「Ubuntu16.04」という名前のテンプレートのデータを取得する
data "kcps_template" "ex" {
  templatefilter = "community"
  name           = "CentOS7.8(64bit)100GB"
  zoneid         = data.kcps_zone.ex.id
}

# ServiceOfferingのデータを取得
data "kcps_service_offering" "ex" {
  name = "Small1(1vCPU,Mem2GB)"
}

# VMの作成
resource "kcps_value_vm" "ex" {
  count             = var._count
  name              = "example-${count.index}"
  serviceofferingid = data.kcps_service_offering.ex.id
  templateid        = data.kcps_template.ex.id
  zoneid            = data.kcps_zone.ex.id
}

# 既存のKCPSのNetwork「PublicFrontSegment」のデータを取得
data "kcps_network" "ex" {
  name   = "PublicFrontSegment"
  zoneid = data.kcps_zone.ex.id
}

# BackSegmentを付与する場合
/*
data "kcps_network" "bx" {
  name   = "BackSegment"
  zoneid = data.kcps_zone.ex.id
}

resource "kcps_nic" "ex" {
  count       = var._count
  networkid = data.kcps_network.bx.id
  virtualmachineid = element(kcps_value_vm.ex.*.id, count.index)
}
*/

# 外部ネットワークからVMに接続するためのPublic IPの取得
resource "kcps_publicip" "ex" {
  networkid = data.kcps_network.ex.id
}

# ポートフォワーディングのルールを作成
resource "kcps_nat_portforward" "ex" {
  count       = var._count
  ipaddressid = kcps_publicip.ex.id

  port {
    privateport = var.ports["private_port"]
    publicport  = var.ports["public_port"] + count.index
  }

  protocol         = "tcp"
  virtualmachineid = element(kcps_value_vm.ex.*.id, count.index)
}

# Firewallのルールの作成
resource "kcps_firewall" "ex" {
  ipaddressid = kcps_publicip.ex.id

  port {
    startport = var.ports["public_port"]
    endport   = var.ports["public_port"] + var._count - 1
  }

  cidrlist = [var.cidr]
  protocol = "tcp"
}


# 各VMのパスワードの出力
output "vm_pass" {
  value = join(",", kcps_value_vm.ex.*.password)
}

# 接続先のPublic IPの出力
output "publicip" {
  value = kcps_publicip.ex.ipaddress
}

# VMのSSH接続に使用するポート番号の出力
output "vm_port" {
  value = join(",", kcps_nat_portforward.ex.*.port.0.publicport)
}

