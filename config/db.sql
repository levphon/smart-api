# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20070
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 8.3.0)
# 数据库: smart-api
# 生成时间: 2024-10-27 08:00:45 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 exec_machine
# ------------------------------------------------------------

LOCK TABLES `exec_machine` WRITE;
/*!40000 ALTER TABLE `exec_machine` DISABLE KEYS */;

INSERT INTO `exec_machine` (`id`, `ip`, `hostname`, `username`, `password`, `port`, `heartbeat`, `status`, `auth_type`, `private_key`, `creator`, `regenerator`, `description`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,'10.50.183.112','10.50.183.112','root','MYw22lpe07PpATygnFByQhVd1nuQRA==',52829,'2024-10-27 16:00:00',1,'2','-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABlwAAAAdzc2gtcn\nNhAAAAAwEAAQAAAYEA1oBi3zK++hYwnk33huB61Fev/Ky8F/BR9/FgR5/3C7Z+zKJTPFh1\nE2NyZ0dcFa+AgoIDLsgGgNQRHuTAQb+gOk/tuxiNLXOyT+z4zAHnP+HJ+WDt2gt5P3AO7U\nJIUCoZhSW43Se6axMtPqXj6Vbs9tWzO6pPKPuDGK2dCNEu2SEzjOIDakS0Acqs2akMXrtO\n9v7ww2EQumzKYuw4vkdnCCRUGY9FTzrdspTdZd6IYSGxGnSTSspPNb+7m/ramg9MyGgUmY\ne4PLULbTE351Anl78SjF/g+x67zUyTLdVRQvnbJPudi9Qb9+k1IzVd6e9I0/wAJL/BUc4W\nvcYDRHkBfGKnIZrzz+GOH0iBbgLVkwKwjd+ixKROV/7/LfzGngYQk9uUm0v+Dlgq3HeoO6\ny5WeMca3zRCYA1sB/4PLwaWH7nphkQMQ+fa5dAfR4LjBFTA8gIpovWma2PB6bEdMNY5nXK\nvjS6acO5zt7RSnlBw69EqlgC3HFJXy/pR+25ybV1AAAFkFtTO1VbUztVAAAAB3NzaC1yc2\nEAAAGBANaAYt8yvvoWMJ5N94bgetRXr/ysvBfwUffxYEef9wu2fsyiUzxYdRNjcmdHXBWv\ngIKCAy7IBoDUER7kwEG/oDpP7bsYjS1zsk/s+MwB5z/hyflg7doLeT9wDu1CSFAqGYUluN\n0numsTLT6l4+lW7PbVszuqTyj7gxitnQjRLtkhM4ziA2pEtAHKrNmpDF67Tvb+8MNhELps\nymLsOL5HZwgkVBmPRU863bKU3WXeiGEhsRp0k0rKTzW/u5v62poPTMhoFJmHuDy1C20xN+\ndQJ5e/Eoxf4Pseu81Mky3VUUL52yT7nYvUG/fpNSM1XenvSNP8ACS/wVHOFr3GA0R5AXxi\npyGa88/hjh9IgW4C1ZMCsI3fosSkTlf+/y38xp4GEJPblJtL/g5YKtx3qDusuVnjHGt80Q\nmANbAf+Dy8Glh+56YZEDEPn2uXQH0eC4wRUwPICKaL1pmtjwemxHTDWOZ1yr40umnDuc7e\n0Up5QcOvRKpYAtxxSV8v6Uftucm1dQAAAAMBAAEAAAGAGgjyNzoPEQaxdv1qnk3Pyscr3q\nTOna83G7uJ3petYhgP8uF+7dOkviozYBK6vA0VsYF7RmnT1D4pJ9FG/pP2LC24Yp2bwRkK\nWwYdupE+krPikmiv5ee/mzIMNcL2SPibKVyHQByK1WU5+CEldRRuZZVRkFvfCM/iPRQRe9\nj78THE8oQaOwNEv/TsHu0USck9T+Bos6Yr5BzBQdl/F6VN/aB/Lq0DkhbIgtzrtGoaroNq\n3hWpLQo6LAFuEYQUlV9mzuHWkrS1DalYYn9dVxxdXbVfdgDfmGVml7M4cpkC4KygK2h73E\nS87SBVGtM7U0Q4E3TU2qMrLFIxjPJKF/FJcZSAo51XjN86YpAc2wBO3nKreHqacu/ud8Fi\nLiL+UM4LEveXhzAUp9vAd4f3TO1oeox2f0Pjp80e2O7fd0wL13T3DgqgkF523xXstX52G4\nH7uAuOo/R8zVH/wEEAoh36wbS34+RsiBaRBfyW44aINqUstVrVisRoBiCA0RTffJkxAAAA\nwBoD3tfSYUc7zcQZ94cmFCjweg74v+vY4oUErvuQXOJfiuOLKa9M7Xx8I7RT7Ji96/Keg5\nE5gdP8KA6VNZT2+JoLdEB1dQssga1a4GWhKVcbBJIxZ3lTfabDTAyKOgp3kU2NUXwNYRDT\nZBKrK4nRI5AA9a1aTsSh9SwtOAVYXuCLidCdAgbSx7s83la0kMgZlCgDKqNAL400xeC46p\nYteBicoWLNhdKVkooWc7/QO6hEK1IThsRScL1CrhgmcxZKEAAAAMEA+9p/8X0YyOOXNZ3Z\ntEZC0hHxbLJIumx7lMbdnxicBWgC+yB0U6cnYTCaUu5T5LchCJ3QR5xjVufrWtV/x/Z1/Y\nsDxZgj4oVqq5RzVxP2TINW6BLngM8qN+2oHKFWr0aKYRdFesY5p1YWpwMWaIHqaK+h3eIN\nNtsB8YSSNniyAy7SBos516tgTkp6uxeHq/WUjJSB5aoyyuTi404SjOUdoHXBNattmmiFv5\nPo/qSZwjhnJAp9/Vf2rRAr1hSlarqpAAAAwQDaCHT8pEW/FNwgBpamMvW/EjHc51Pznjef\n6pUK3lLucpVmEM0vFKrD5dSgFrIEE0+ZZACKJ0iQBgTLZQwQUHzjfxyStOhpXklWj8mRdv\np6V9UAHnqPRgKVox2B/gLrRuPVo83rUgm6Attha2cjMyF1/X19lL/PVnIIeeIVqZmvk45L\nGZKaoC4B702vg8StTfVWOIp8fQLmEyAnWPCgawlRw4vOhjqgTS8eMJHREi67/Bd8PoZpuq\nGTSrCqUraTD+0AAAAXcm9vdEBubS11Y2xvdWQtY3B1LWxkYXABAgME\n-----END OPENSSH PRIVATE KEY-----','admin','admin','This is a test machine',1,0,'2024-08-30 19:29:33.388','2024-09-03 19:49:30.643',NULL),
    (2,'10.50.209.150','10.50.209.150','root','jM256ifdbVb8jAfTyXaWxk0lEEk0lg==',52829,'2024-10-27 16:00:00',1,'1','','admin','admin','',1,0,'2024-09-02 19:54:17.001','2024-09-10 20:57:44.276',NULL);

/*!40000 ALTER TABLE `exec_machine` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 exec_machine_history
# ------------------------------------------------------------

LOCK TABLES `exec_machine_history` WRITE;
/*!40000 ALTER TABLE `exec_machine_history` DISABLE KEYS */;

INSERT INTO `exec_machine_history` (`id`, `task_id`, `task_name`, `machine_id`, `hostname`, `ip`, `status`, `stdout`, `stderr`, `executed_at`, `creator`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `execution_time`, `taskUUID`)
VALUES
    (1,6,'LDAP 账号注册',1,'10.50.183.112','10.50.183.112',0,'{\"userFirstOu\":\"develop\",\"userName\":\"zhangsan1\",\"userOnePasswd\":\"1\",\"userSecondOu\":\"sre\",\"userTwoPasswd\":\"1\"}\nsunwenbo 测试\n{\"message\":{\"code\":50001,\"message\":\"Failed to create user:zhangsan1. reason: LDAP Result Code 68 \\\"Entry Already Exists\\\": \"}}请求已发送，工单数据：{\"userFirstOu\":\"develop\",\"userName\":\"zhangsan1\",\"userOnePasswd\":\"1\",\"userSecondOu\":\"sre\",\"userTwoPasswd\":\"1\"}\n','{\"userFirstOu\":\"develop\",\"userName\":\"zhangsan1\",\"userOnePasswd\":\"1\",\"userSecondOu\":\"sre\",\"userTwoPasswd\":\"1\"}\nsunwenbo 测试\n{\"message\":{\"code\":50001,\"message\":\"Failed to create user:zhangsan1. reason: LDAP Result Code 68 \\\"Entry Already Exists\\\": \"}}请求已发送，工单数据：{\"userFirstOu\":\"develop\",\"userName\":\"zhangsan1\",\"userOnePasswd\":\"1\",\"userSecondOu\":\"sre\",\"userTwoPasswd\":\"1\"}\n','2024-10-23 14:59:08','system',0,0,'2024-10-23 14:59:08.470','2024-10-23 14:59:08.470',NULL,0,'922e337e');

/*!40000 ALTER TABLE `exec_machine_history` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 flow_manage
# ------------------------------------------------------------

LOCK TABLES `flow_manage` WRITE;
/*!40000 ALTER TABLE `flow_manage` DISABLE KEYS */;

INSERT INTO `flow_manage` (`id`, `name`, `categoryId`, `notice`, `comments`, `ratings`, `description`, `creator`, `regenerator`, `structure`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,'k8s 接入',3,'[1, 2]',1,1,'新服务接入k8s集群','admin','admin','{\"edges\": [{\"id\": \"flow1725786420505\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"工单提交\", \"style\": {}, \"source\": \"start1725786402706\", \"target\": \"userTask1725786404846\", \"endPoint\": {\"x\": 268.3046875, \"y\": 188, \"index\": 3}, \"startPoint\": {\"x\": 96.3046875, \"y\": 173, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1725786447518\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1725786404846\", \"target\": \"userTask1725786406939\", \"endPoint\": {\"x\": 706.3046875, \"y\": 169.5, \"index\": 3}, \"startPoint\": {\"x\": 349.3046875, \"y\": 188, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1725786453569\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1725786404846\", \"target\": \"end1725786416842\", \"endPoint\": {\"x\": 198.3046875, \"y\": 426.5, \"index\": 2}, \"startPoint\": {\"x\": 308.8046875, \"y\": 210.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 2}, {\"id\": \"flow1725786478581\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1725786406939\", \"target\": \"receiveTask1725786410456\", \"endPoint\": {\"x\": 596.8046875, \"y\": 420.5, \"index\": 0}, \"startPoint\": {\"x\": 746.8046875, \"y\": 192, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1725786481665\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1725786406939\", \"target\": \"userTask1725786404846\", \"endPoint\": {\"x\": 308.8046875, \"y\": 165.5, \"index\": 0}, \"startPoint\": {\"x\": 746.8046875, \"y\": 147, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1725786512548\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理成功\", \"style\": {}, \"source\": \"receiveTask1725786410456\", \"target\": \"end1725786416842\", \"endPoint\": {\"x\": 213.8046875, \"y\": 411, \"index\": 0}, \"startPoint\": {\"x\": 556.3046875, \"y\": 443, \"index\": 3}, \"sourceAnchor\": 3, \"targetAnchor\": 0}, {\"id\": \"flow1725786519040\", \"seq\": \"7\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1725786410456\", \"target\": \"userTask1725786406939\", \"endPoint\": {\"x\": 787.3046875, \"y\": 169.5, \"index\": 1}, \"startPoint\": {\"x\": 637.3046875, \"y\": 443, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}], \"nodes\": [{\"x\": 80.8046875, \"y\": 173, \"id\": \"start1725786402706\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 308.8046875, \"y\": 188, \"id\": \"userTask1725786404846\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"直属领导\", \"shape\": \"user-task-node\", \"style\": {}, \"dueDate\": \"2024-09-30 17:10:19\", \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [14], \"isCounterSign\": false}, {\"x\": 746.8046875, \"y\": 169.5, \"id\": \"userTask1725786406939\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维负责人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [1], \"isCounterSign\": false}, {\"x\": 596.8046875, \"y\": 443, \"id\": \"receiveTask1725786410456\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.183.112\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 213.8046875, \"y\": 426.5, \"id\": \"end1725786416842\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',0,1,'2024-09-08 17:10:45.073','2024-10-09 19:11:00.882',NULL),
    (2,'LDAP账号注册',1,'[1, 2, 3]',1,1,'LDAP账号','admin','admin','{\"edges\": [{\"id\": \"flow1726408760137\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交\", \"style\": {}, \"source\": \"start1726408739841\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 382.8046875, \"y\": 75, \"index\": 3}, \"startPoint\": {\"x\": 210.3046875, \"y\": 68.5, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726408765640\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理成功\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 597.3046875, \"y\": 209, \"index\": 3}, \"sourceAnchor\": 3, \"targetAnchor\": 2}, {\"id\": \"flow1726408806929\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"receiveTask1726408752430\", \"endPoint\": {\"x\": 637.8046875, \"y\": 186.5, \"index\": 0}, \"startPoint\": {\"x\": 463.8046875, \"y\": 75, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 0}, {\"id\": \"flow1726408835676\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 423.3046875, \"y\": 97.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 2}, {\"id\": \"flow1726408890148\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 423.3046875, \"y\": 52.5, \"index\": 0}, \"startPoint\": {\"x\": 637.8046875, \"y\": 231.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 194.8046875, \"y\": 68.5, \"id\": \"start1726408739841\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 423.3046875, \"y\": 75, \"id\": \"userTask1726408742477\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维人员\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 637.8046875, \"y\": 209, \"id\": \"receiveTask1726408752430\", \"size\": [80, 44], \"task\": [\"LDAP 账号注册\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.183.112\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 382.8046875, \"y\": 223, \"id\": \"end1726408758173\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-15 22:02:06.143','2024-10-17 19:54:44.780',NULL),
    (3,'LDAP账号密码修改',1,'[1, 2, 3]',0,0,'修改LDAP账号','admin','','{\"edges\": [{\"id\": \"flow1726409531712\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交\", \"style\": {}, \"source\": \"start1726409502219\", \"target\": \"userTask1726409505077\", \"endPoint\": {\"x\": 403.8046875, \"y\": 101, \"index\": 3}, \"startPoint\": {\"x\": 232.3046875, \"y\": 101.5, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726409577610\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726409505077\", \"target\": \"userTask1726409518931\", \"endPoint\": {\"x\": 597.3046875, \"y\": 97, \"index\": 3}, \"startPoint\": {\"x\": 484.8046875, \"y\": 101, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726409579416\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726409518931\", \"target\": \"receiveTask1726409508537\", \"endPoint\": {\"x\": 637.8046875, \"y\": 232.5, \"index\": 0}, \"startPoint\": {\"x\": 637.8046875, \"y\": 119.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726409593415\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理成功\", \"style\": {}, \"source\": \"receiveTask1726409508537\", \"target\": \"end1726409530005\", \"endPoint\": {\"x\": 384.3046875, \"y\": 263, \"index\": 2}, \"startPoint\": {\"x\": 597.3046875, \"y\": 255, \"index\": 3}, \"sourceAnchor\": 3, \"targetAnchor\": 2}, {\"id\": \"flow1726409609060\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726409505077\", \"target\": \"end1726409530005\", \"endPoint\": {\"x\": 384.3046875, \"y\": 263, \"index\": 2}, \"startPoint\": {\"x\": 444.3046875, \"y\": 78.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726409657627\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726409518931\", \"target\": \"userTask1726409505077\", \"endPoint\": {\"x\": 444.3046875, \"y\": 78.5, \"index\": 0}, \"startPoint\": {\"x\": 637.8046875, \"y\": 74.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726409703536\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726409508537\", \"target\": \"userTask1726409518931\", \"endPoint\": {\"x\": 678.3046875, \"y\": 97, \"index\": 1}, \"startPoint\": {\"x\": 678.3046875, \"y\": 255, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}], \"nodes\": [{\"x\": 216.8046875, \"y\": 101.5, \"id\": \"start1726409502219\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 444.3046875, \"y\": 101, \"id\": \"userTask1726409505077\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门领导\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [14], \"isCounterSign\": false}, {\"x\": 637.8046875, \"y\": 255, \"id\": \"receiveTask1726409508537\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"运维值班人\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 637.8046875, \"y\": 97, \"id\": \"userTask1726409518931\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维领导\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 399.8046875, \"y\": 263, \"id\": \"end1726409530005\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,0,'2024-09-15 22:15:18.369','2024-09-15 22:15:18.369',NULL),
    (4,'LDAP权限申请',1,'[1, 2, 3]',1,1,'登陆LDAP关联的系统','admin','','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,0,'2024-09-19 16:01:04.378','2024-09-19 16:01:04.378',NULL),
    (5,'服务器权限',1,'[1, 2, 3]',1,1,'服务器登陆权限申请','admin','admin','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-19 16:01:04.378','2024-09-19 16:24:10.452',NULL),
    (6,'新增ingress',5,'[1, 2, 3]',1,1,'k8s集群服务暴露','admin','admin','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-19 16:01:04.378','2024-09-20 16:47:32.734',NULL),
    (7,'监控数据采集',3,'[1, 2, 3]',1,1,'监控数据采集','admin','admin','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-19 16:01:04.378','2024-09-20 16:50:55.516',NULL),
    (8,'DNS 域名解析',5,'[1, 2, 3]',1,1,'域名解析','admin','admin','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-19 16:01:04.378','2024-09-20 16:53:32.385',NULL),
    (9,'数据分析',4,'[1, 2, 3]',1,1,'数据分析','admin','admin','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-19 16:01:04.378','2024-09-20 16:55:19.126',NULL),
    (10,'数据分析-copy',4,'[1, 2, 3]',1,1,'数据分析','admin','admin','{\"edges\": [{\"id\": \"flow1726732590211\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交申请\", \"style\": {}, \"source\": \"start1726732577127\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 284.3046875, \"y\": 182, \"index\": 3}, \"startPoint\": {\"x\": 103.3046875, \"y\": 90, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732635153\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"userTask1726732582739\", \"reverse\": false, \"endPoint\": {\"x\": 719.3046875, \"y\": 66.5, \"index\": 3}, \"startPoint\": {\"x\": 365.3046875, \"y\": 182, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726732648043\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732580428\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 579.8046875, \"y\": 404, \"index\": 2}, \"startPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 2}, {\"id\": \"flow1726732676564\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"receiveTask1726732585122\", \"endPoint\": {\"x\": 161.8046875, \"y\": 248.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 89, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}, {\"id\": \"flow1726732695908\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726732582739\", \"target\": \"userTask1726732580428\", \"endPoint\": {\"x\": 324.8046875, \"y\": 159.5, \"index\": 0}, \"startPoint\": {\"x\": 759.8046875, \"y\": 44, \"index\": 0}, \"sourceAnchor\": 0, \"targetAnchor\": 0}, {\"id\": \"flow1726732751482\", \"seq\": \"6\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"userTask1726732582739\", \"endPoint\": {\"x\": 800.3046875, \"y\": 66.5, \"index\": 1}, \"startPoint\": {\"x\": 202.3046875, \"y\": 271, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 1}, {\"id\": \"flow1726732812515\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"登陆关联的系统\", \"style\": {}, \"source\": \"receiveTask1726732585122\", \"target\": \"end1726732588377\", \"endPoint\": {\"x\": 595.3046875, \"y\": 388.5, \"index\": 0}, \"startPoint\": {\"x\": 161.8046875, \"y\": 293.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 87.8046875, \"y\": 90, \"id\": \"start1726732577127\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 324.8046875, \"y\": 182, \"id\": \"userTask1726732580428\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"部门责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 759.8046875, \"y\": 66.5, \"id\": \"userTask1726732582739\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维责任人\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 161.8046875, \"y\": 271, \"id\": \"receiveTask1726732585122\", \"size\": [80, 44], \"task\": [\"测试任务\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.209.150\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 595.3046875, \"y\": 404, \"id\": \"end1726732588377\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}',1,1,'2024-09-19 16:01:04.378','2024-09-20 16:55:19.126','2024-10-09 18:49:50.789');

/*!40000 ALTER TABLE `flow_manage` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 flow_templates
# ------------------------------------------------------------

LOCK TABLES `flow_templates` WRITE;
/*!40000 ALTER TABLE `flow_templates` DISABLE KEYS */;

INSERT INTO `flow_templates` (`id`, `name`, `description`, `bindCount`, `form_data`, `categoryId`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `creator`, `regenerator`, `bindFlow`)
VALUES
    (1,'K8sServiceAdd','服务接入k8s集群',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid106178\", \"cols\": [{\"id\": \"grid-col-29207\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol29207\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input50983\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input50983\", \"size\": \"\", \"type\": \"text\", \"label\": \"input\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-98439\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol98439\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input37237\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input37237\", \"size\": \"\", \"type\": \"text\", \"label\": \"input\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid106178\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"input60459\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input60459\", \"size\": \"\", \"type\": \"text\", \"label\": \"input\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}, {\"id\": \"radio18185\", \"icon\": \"radio-field\", \"type\": \"radio\", \"options\": {\"name\": \"radio18185\", \"size\": \"\", \"label\": \"radio\", \"border\": false, \"hidden\": false, \"disabled\": false, \"onChange\": \"\", \"required\": false, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"buttonStyle\": false, \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"optionItems\": [{\"label\": \"radio 1\", \"value\": 1}, {\"label\": \"radio 2\", \"value\": 2}, {\"label\": \"radio 3\", \"value\": 3}], \"defaultValue\": null, \"displayStyle\": \"inline\", \"labelTooltip\": null, \"requiredHint\": \"\", \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',3,1,1,'2024-09-11 15:57:02.520','2024-10-17 11:53:37.555',NULL,'admin','admin',1),
    (2,'LdapAccountRegister','注册LDAP账号、可以登陆',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"card82642\", \"icon\": \"card\", \"type\": \"card\", \"options\": {\"name\": \"card82642\", \"label\": \"LDAP 用户注册\", \"folded\": false, \"hidden\": false, \"shadow\": \"never\", \"showFold\": false, \"cardWidth\": \"100%\", \"customClass\": []}, \"category\": \"container\", \"widgetList\": [{\"id\": \"grid6333\", \"cols\": [{\"id\": \"grid-col-16762\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol16762\", \"pull\": 0, \"push\": 0, \"span\": 24, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": []}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"userName\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letterAndNumber\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"请输入名字的拼音\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"名字必须为拼音字母或者拼音字母和数字结合\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid6333\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": \"\"}, \"category\": \"container\"}, {\"id\": \"grid68334\", \"cols\": [{\"id\": \"grid-col-42971\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol42971\", \"pull\": 0, \"push\": 0, \"span\": 24, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input35125\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"userFirstOu\", \"size\": \"\", \"type\": \"text\", \"label\": \"一级部门:\", \"hidden\": true, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"develop\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid68334\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": \"\"}, \"category\": \"container\"}, {\"id\": \"grid35254\", \"cols\": [{\"id\": \"grid-col-28449\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol28449\", \"pull\": 0, \"push\": 0, \"span\": 24, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": []}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"select82295\", \"icon\": \"select-field\", \"type\": \"select\", \"options\": {\"name\": \"userSecondOu\", \"size\": \"\", \"label\": \"部门:\", \"hidden\": false, \"onBlur\": \"\", \"remote\": false, \"onFocus\": \"\", \"disabled\": false, \"multiple\": false, \"onChange\": \"\", \"required\": true, \"clearable\": true, \"onCreated\": \"\", \"onMounted\": \"\", \"filterable\": false, \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"letterAndNumber\", \"allowCreate\": false, \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"optionItems\": [{\"label\": \"sre\", \"value\": \"sre\"}, {\"label\": \"qa\", \"value\": \"qa\"}, {\"label\": \"develop\", \"value\": \"develop\"}], \"placeholder\": \"请选择所属部门\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"multipleLimit\": 0, \"onRemoteQuery\": \"\", \"labelIconClass\": null, \"validationHint\": \"\", \"automaticDropdown\": false, \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid35254\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"grid58765\", \"cols\": [{\"id\": \"grid-col-29276\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol29276\", \"pull\": 0, \"push\": 0, \"span\": 24, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input11242\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"userOnePasswd\", \"size\": \"\", \"type\": \"password\", \"label\": \"用户密码\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": 15, \"minLength\": 8, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letterAndNumber\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"请输入密码\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"请输入密码\", \"showPassword\": true, \"showWordLimit\": true, \"labelIconClass\": null, \"validationHint\": \"可以是数字和字母组成\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid58765\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": \"\"}, \"category\": \"container\"}, {\"id\": \"grid106378\", \"cols\": [{\"id\": \"grid-col-32892\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32892\", \"pull\": 0, \"push\": 0, \"span\": 24, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input46762\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"userTwoPasswd\", \"size\": \"\", \"type\": \"password\", \"label\": \"确认密码:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": 15, \"minLength\": 8, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letterAndNumber\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"请确认密码\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"请输入密码\", \"showPassword\": true, \"showWordLimit\": true, \"labelIconClass\": null, \"validationHint\": \"可以是数字和字母组成\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid106378\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": \"\"}, \"category\": \"container\"}]}]}',1,1,1,'2024-09-15 22:05:40.443','2024-10-22 14:47:02.318',NULL,'admin','admin',2),
    (3,'LDAPPassWordUpdate','LDAP账号密码修改',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid106178\", \"cols\": [{\"id\": \"grid-col-29207\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol29207\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input50983\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input50983\", \"size\": \"\", \"type\": \"text\", \"label\": \"input\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-98439\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol98439\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input37237\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input37237\", \"size\": \"\", \"type\": \"text\", \"label\": \"input\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid106178\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"input60459\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input60459\", \"size\": \"\", \"type\": \"text\", \"label\": \"input\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}',1,1,1,'2024-09-15 22:35:55.720','2024-10-17 11:53:44.806',NULL,'admin','admin',3),
    (4,'LDAPPermission','申请LDAP关联的系统登陆或者操作权限',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid111008\", \"cols\": [{\"id\": \"grid-col-24534\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol24534\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67790\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letter\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-32750\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32750\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input67791\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67791\", \"size\": \"\", \"type\": \"text\", \"label\": \"部门：\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid111008\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"textarea91560\", \"icon\": \"textarea-field\", \"type\": \"textarea\", \"options\": {\"name\": \"textarea91560\", \"rows\": 3, \"size\": \"\", \"label\": \"textarea\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',1,1,1,'2024-09-19 16:02:44.287','2024-10-17 11:53:57.977',NULL,'admin','admin',4),
    (5,'ServerPermission','登陆服务器',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid111008\", \"cols\": [{\"id\": \"grid-col-24534\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol24534\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67790\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letter\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-32750\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32750\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input67791\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67791\", \"size\": \"\", \"type\": \"text\", \"label\": \"部门：\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid111008\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"textarea91560\", \"icon\": \"textarea-field\", \"type\": \"textarea\", \"options\": {\"name\": \"textarea91560\", \"rows\": 3, \"size\": \"\", \"label\": \"textarea\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',1,1,1,'2024-09-19 16:24:20.950','2024-10-17 11:54:04.772',NULL,'admin','admin',5),
    (6,'AddIngress','k8s集群服务暴露',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid111008\", \"cols\": [{\"id\": \"grid-col-24534\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol24534\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67790\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letter\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-32750\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32750\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input67791\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67791\", \"size\": \"\", \"type\": \"text\", \"label\": \"部门：\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid111008\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"textarea91560\", \"icon\": \"textarea-field\", \"type\": \"textarea\", \"options\": {\"name\": \"textarea91560\", \"rows\": 3, \"size\": \"\", \"label\": \"textarea\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',5,1,1,'2024-09-20 16:47:43.378','2024-10-17 11:54:11.843',NULL,'admin','admin',6),
    (7,'PrometheusMonitory','Prometheus监控',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid111008\", \"cols\": [{\"id\": \"grid-col-24534\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol24534\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67790\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letter\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-32750\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32750\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input67791\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67791\", \"size\": \"\", \"type\": \"text\", \"label\": \"部门：\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid111008\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"textarea91560\", \"icon\": \"textarea-field\", \"type\": \"textarea\", \"options\": {\"name\": \"textarea91560\", \"rows\": 3, \"size\": \"\", \"label\": \"textarea\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',3,1,1,'2024-09-20 16:51:41.956','2024-10-17 11:54:21.281',NULL,'admin','admin',7),
    (8,'DomainNameResolution','域名解析',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid111008\", \"cols\": [{\"id\": \"grid-col-24534\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol24534\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67790\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letter\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-32750\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32750\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input67791\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67791\", \"size\": \"\", \"type\": \"text\", \"label\": \"部门：\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid111008\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"textarea91560\", \"icon\": \"textarea-field\", \"type\": \"textarea\", \"options\": {\"name\": \"textarea91560\", \"rows\": 3, \"size\": \"\", \"label\": \"textarea\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',5,1,1,'2024-09-20 16:53:50.040','2024-10-17 11:54:26.997',NULL,'admin','admin',8),
    (9,'DataAnalysis','数据分析',1,'{\"formConfig\": {\"size\": \"\", \"cssCode\": \"\", \"refName\": \"vForm\", \"functions\": \"\", \"modelName\": \"formData\", \"rulesName\": \"rules\", \"labelAlign\": \"label-left-align\", \"labelWidth\": 80, \"layoutType\": \"PC\", \"customClass\": \"\", \"labelPosition\": \"left\", \"onFormCreated\": \"\", \"onFormMounted\": \"\", \"onFormDataChange\": \"\"}, \"widgetList\": [{\"id\": \"grid111008\", \"cols\": [{\"id\": \"grid-col-24534\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol24534\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input26051\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67790\", \"size\": \"\", \"type\": \"text\", \"label\": \"姓名:\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": true, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"label-center-align\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"letter\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"姓名拼音\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}, {\"id\": \"grid-col-32750\", \"icon\": \"grid-col\", \"type\": \"grid-col\", \"options\": {\"md\": 12, \"sm\": 12, \"xs\": 12, \"name\": \"gridCol32750\", \"pull\": 0, \"push\": 0, \"span\": 12, \"hidden\": false, \"offset\": 0, \"responsive\": false, \"customClass\": \"\"}, \"category\": \"container\", \"internal\": true, \"widgetList\": [{\"id\": \"input67791\", \"icon\": \"text-field\", \"type\": \"input\", \"options\": {\"name\": \"input67791\", \"size\": \"\", \"type\": \"text\", \"label\": \"部门：\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"clearable\": true, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"buttonIcon\": \"el-icon-search\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"prefixIcon\": \"\", \"suffixIcon\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": \"\", \"labelHidden\": false, \"placeholder\": \"\", \"appendButton\": false, \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showPassword\": false, \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\", \"onAppendButtonClick\": \"\", \"appendButtonDisabled\": false}, \"formItemFlag\": true}]}], \"icon\": \"grid\", \"type\": \"grid\", \"options\": {\"name\": \"grid111008\", \"gutter\": 12, \"hidden\": false, \"colHeight\": null, \"customClass\": []}, \"category\": \"container\"}, {\"id\": \"textarea91560\", \"icon\": \"textarea-field\", \"type\": \"textarea\", \"options\": {\"name\": \"textarea91560\", \"rows\": 3, \"size\": \"\", \"label\": \"textarea\", \"hidden\": false, \"onBlur\": \"\", \"onFocus\": \"\", \"onInput\": \"\", \"disabled\": false, \"onChange\": \"\", \"readonly\": false, \"required\": false, \"maxLength\": null, \"minLength\": null, \"onCreated\": \"\", \"onMounted\": \"\", \"labelAlign\": \"\", \"labelWidth\": null, \"onValidate\": \"\", \"validation\": \"\", \"columnWidth\": \"200px\", \"customClass\": [], \"labelHidden\": false, \"placeholder\": \"\", \"defaultValue\": \"\", \"labelTooltip\": null, \"requiredHint\": \"\", \"showWordLimit\": false, \"labelIconClass\": null, \"validationHint\": \"\", \"labelIconPosition\": \"rear\"}, \"formItemFlag\": true}]}',4,1,1,'2024-09-20 16:55:55.436','2024-10-17 11:54:33.620',NULL,'admin','admin',9);

/*!40000 ALTER TABLE `flow_templates` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 operation_history
# ------------------------------------------------------------



# 转储表 order_category
# ------------------------------------------------------------

LOCK TABLES `order_category` WRITE;
/*!40000 ALTER TABLE `order_category` DISABLE KEYS */;

INSERT INTO `order_category` (`id`, `name`, `chineseName`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `creator`, `regenerator`)
VALUES
    (1,'Permission','权限申请',1,0,'2024-07-16 16:28:49.140','2024-07-16 16:28:49.140',NULL,NULL,NULL),
    (2,'MessageQueue','消息队列',1,0,'2024-07-16 16:28:49.140','2024-07-16 16:28:49.140',NULL,NULL,NULL),
    (3,'Kubernetes','k8s相关',1,1,'2024-07-16 16:28:49.140','2024-08-15 20:23:12.961',NULL,'','admin'),
    (4,'BigData','大数据',1,0,'2024-07-16 16:28:49.140','2024-07-16 16:28:49.140',NULL,NULL,NULL),
    (5,'ServiceExport','服务暴露',1,0,'2024-07-16 16:28:49.140','2024-07-16 16:28:49.140',NULL,NULL,NULL);

/*!40000 ALTER TABLE `order_category` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 order_comment
# ------------------------------------------------------------

LOCK TABLES `order_comment` WRITE;
/*!40000 ALTER TABLE `order_comment` DISABLE KEYS */;

INSERT INTO `order_comment` (`id`, `order_id`, `parent_id`, `comment`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,3,NULL,'留言',1,0,'2024-10-22 14:55:25.689','2024-10-22 14:55:25.689',NULL),
    (2,3,NULL,'留言2',2,0,'2024-10-22 14:55:59.584','2024-10-22 14:55:59.584',NULL);

/*!40000 ALTER TABLE `order_comment` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 order_items
# ------------------------------------------------------------

LOCK TABLES `order_items` WRITE;
/*!40000 ALTER TABLE `order_items` DISABLE KEYS */;

INSERT INTO `order_items` (`id`, `title`, `bindTempLate`, `description`, `icon`, `categoryId`, `link`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `creator`, `regenerator`)
VALUES
    (1,'k8s服务接入','K8sServiceAdd','将服务部署在k8s集群','fa fa-cloud-upload',3,'/order/classify/K8sServiceAdd',1,1,'2024-09-08 17:17:43.606','2024-10-09 19:05:41.915',NULL,'admin','admin'),
    (2,'LDAP账号注册','LdapAccountRegister','注册LDAP账号','fa fa-user',1,'/order/classify/LdapAccountRegister',1,1,'2024-09-15 22:07:29.949','2024-10-09 15:48:57.993',NULL,'admin','admin'),
    (3,'LDAP密码修改','LDAPPassWordUpdate','修改LDAP账号密码','fa fa-user-circle',1,'/order/classify/LDAPPassWordUpdate',1,1,'2024-09-15 22:36:27.849','2024-10-09 15:49:38.867',NULL,'admin','admin'),
    (4,'LDAP权限申请','LDAPPermission','开通LDAP账号关联的系统','fa fa-user-plus',1,'/order/classify/LDAPPermission',1,1,'2024-09-19 16:03:43.988','2024-10-09 15:50:13.831',NULL,'admin','admin'),
    (5,'服务器权限','ServerPermission','服务器登陆权限','el-icon-s-platform',1,'/order/classify/ServerPermission',1,1,'2024-09-19 16:26:53.357','2024-10-22 14:23:52.951',NULL,'admin','admin'),
    (6,'新增ingress','AddIngress','k8s集群服务暴露','fa fa-share-alt-square',5,'/order/classify/AddIngress',1,0,'2024-09-20 16:49:40.373','2024-09-20 16:49:40.373',NULL,'admin',''),
    (7,'k8s监控数据采集','PrometheusMonitory','Prometheus监控','fa fa-bar-chart',3,'/order/classify/PrometheusMonitory',1,0,'2024-09-20 16:52:33.997','2024-09-20 16:52:33.997',NULL,'admin',''),
    (8,'DNS 域名解析','DomainNameResolution','域名解析','el-icon-set-up',5,'/order/classify/DomainNameResolution',1,0,'2024-09-20 16:54:23.744','2024-09-20 16:54:23.744',NULL,'admin',''),
    (9,'数据分析','DataAnalysis','数据分析','el-icon-data-line',4,'/order/classify/DataAnalysis',1,0,'2024-09-20 16:56:26.725','2024-09-20 16:56:26.725',NULL,'admin','');

/*!40000 ALTER TABLE `order_items` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 order_rating
# ------------------------------------------------------------

LOCK TABLES `order_rating` WRITE;
/*!40000 ALTER TABLE `order_rating` DISABLE KEYS */;

INSERT INTO `order_rating` (`id`, `order_id`, `rating`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `task_handler`)
VALUES
    (1,3,5,1,0,'2024-10-22 15:19:44.957','2024-10-22 15:19:44.957',NULL,2);

/*!40000 ALTER TABLE `order_rating` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 order_task
# ------------------------------------------------------------

LOCK TABLES `order_task` WRITE;
/*!40000 ALTER TABLE `order_task` DISABLE KEYS */;

INSERT INTO `order_task` (`id`, `name`, `taskType`, `interpreter`, `content`, `creator`, `description`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `regenerator`)
VALUES
    (1,'111','1','/usr/bin/python3','#!/usr/bin/env python3\n\n# 在这里编写 Python 脚本\nprint(\"task 孙文波测试\")\n# workOrderForm=$1 # 接受工单工单数据\n\n# 推荐使用 jq 命令来获取 Json 结构对应的键值数据。具体使用方法，还请自行百度。\n# 此外，还需注意 jq 命令，若没有还需在任务工作节点安装此命令。\n\n# ------------- 在下面编写您的业务逻辑代码 -------------','admin','111',1,1,'2024-08-19 19:20:39.940','2024-08-19 19:20:39.940',NULL,'admin'),
    (2,'LDAP 账号注册','0','/bin/bash','#!/bin/bash\n\n# 接收工单数据\nworkOrderForm=$1\necho $workOrderForm\n\necho \"sunwenbo 测试\"\n\n# 使用 curl 调用接口\ncurl --location \'http://ldap-user-manage.xxx/api/ldap/addUsers\' \\\n--header \'Content-Type: application/json\' \\\n--data \"$workOrderForm\"\n\n# 输出请求结果\necho \"请求已发送，工单数据：$workOrderForm\"\n','sunwenbo','',2,2,'2024-09-10 17:34:50.094','2024-09-10 17:34:50.094',NULL,'sunwenbo');

/*!40000 ALTER TABLE `order_task` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 order_works
# ------------------------------------------------------------

LOCK TABLES `order_works` WRITE;
/*!40000 ALTER TABLE `order_works` DISABLE KEYS */;

INSERT INTO `order_works` (`id`, `title`, `currentNode`, `currentHandler`, `process`, `priority`, `status`, `department`, `description`, `form_data`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `creator`, `regenerator`, `template`, `currentHandlerId`, `bindFlowData`)
VALUES
    (1,'LDAP账号注册-admin-20241017203343','工单结束','admin','LDAP账号注册','very-urgent','termination','数据中台','1111','{\"userName\": \"zhangsan\", \"userSecondOu\": \"develop1\", \"userOnePasswd\": \"123\", \"userTwoPasswd\": \"123\"}',1,0,'2024-10-17 20:34:00.904','2024-10-21 19:54:57.933',NULL,'admin','','LdapAccountRegister',1,'{\"id\": 2, \"name\": \"LDAP账号注册\", \"notice\": [1, 2, 3], \"creator\": \"admin\", \"ratings\": true, \"comments\": true, \"createBy\": 1, \"updateBy\": 1, \"createdAt\": \"2024-09-15 22:02:06\", \"structure\": {\"edges\": [{\"id\": \"flow1726408760137\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交\", \"style\": {}, \"source\": \"start1726408739841\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 382.8046875, \"y\": 75, \"index\": 3}, \"startPoint\": {\"x\": 210.3046875, \"y\": 68.5, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726408765640\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理成功\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 597.3046875, \"y\": 209, \"index\": 3}, \"sourceAnchor\": 3, \"targetAnchor\": 2}, {\"id\": \"flow1726408806929\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"receiveTask1726408752430\", \"endPoint\": {\"x\": 637.8046875, \"y\": 186.5, \"index\": 0}, \"startPoint\": {\"x\": 463.8046875, \"y\": 75, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 0}, {\"id\": \"flow1726408835676\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 423.3046875, \"y\": 97.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 2}, {\"id\": \"flow1726408890148\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 423.3046875, \"y\": 52.5, \"index\": 0}, \"startPoint\": {\"x\": 637.8046875, \"y\": 231.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 194.8046875, \"y\": 68.5, \"id\": \"start1726408739841\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 423.3046875, \"y\": 75, \"id\": \"userTask1726408742477\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维人员\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 637.8046875, \"y\": 209, \"id\": \"receiveTask1726408752430\", \"size\": [80, 44], \"task\": [\"LDAP 账号注册\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.183.112\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 382.8046875, \"y\": 223, \"id\": \"end1726408758173\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}, \"updatedAt\": \"2024-10-17 19:54:44\", \"categoryId\": 1, \"description\": \"LDAP账号\", \"regenerator\": \"admin\"}'),
    (2,'LDAP账号注册-孙文波-20241022115652','工单结束','sunwenbo','LDAP账号注册','very-urgent','termination','数据中台','','{\"userName\": \"zhangsan\", \"userFirstOu\": \"develop\", \"userSecondOu\": \"sre\", \"userOnePasswd\": \"1234567\", \"userTwoPasswd\": \"123456\"}',2,0,'2024-10-22 11:59:02.715','2024-10-22 12:07:45.435',NULL,'sunwenbo','','LdapAccountRegister',2,'{\"id\": 2, \"name\": \"LDAP账号注册\", \"notice\": [1, 2, 3], \"creator\": \"admin\", \"ratings\": true, \"comments\": true, \"createBy\": 1, \"updateBy\": 1, \"createdAt\": \"2024-09-15 22:02:06\", \"structure\": {\"edges\": [{\"id\": \"flow1726408760137\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交\", \"style\": {}, \"source\": \"start1726408739841\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 382.8046875, \"y\": 75, \"index\": 3}, \"startPoint\": {\"x\": 210.3046875, \"y\": 68.5, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726408765640\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理成功\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 597.3046875, \"y\": 209, \"index\": 3}, \"sourceAnchor\": 3, \"targetAnchor\": 2}, {\"id\": \"flow1726408806929\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"receiveTask1726408752430\", \"endPoint\": {\"x\": 637.8046875, \"y\": 186.5, \"index\": 0}, \"startPoint\": {\"x\": 463.8046875, \"y\": 75, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 0}, {\"id\": \"flow1726408835676\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 423.3046875, \"y\": 97.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 2}, {\"id\": \"flow1726408890148\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 423.3046875, \"y\": 52.5, \"index\": 0}, \"startPoint\": {\"x\": 637.8046875, \"y\": 231.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 194.8046875, \"y\": 68.5, \"id\": \"start1726408739841\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 423.3046875, \"y\": 75, \"id\": \"userTask1726408742477\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维人员\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 637.8046875, \"y\": 209, \"id\": \"receiveTask1726408752430\", \"size\": [80, 44], \"task\": [\"LDAP 账号注册\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.183.112\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 382.8046875, \"y\": 223, \"id\": \"end1726408758173\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}, \"updatedAt\": \"2024-10-17 19:54:44\", \"categoryId\": 1, \"description\": \"LDAP账号\", \"regenerator\": \"admin\"}'),
    (3,'LDAP账号注册-admin-20241022144951','工单结束','admin','LDAP账号注册','urgent','termination','解决方案团队','测试','{\"userName\": \"zhangsan1\", \"userFirstOu\": \"develop\", \"userSecondOu\": \"sre\", \"userOnePasswd\": \"1\", \"userTwoPasswd\": \"1\"}',1,0,'2024-10-22 14:54:46.414','2024-10-23 14:59:08.485',NULL,'admin','','LdapAccountRegister',1,'{\"id\": 2, \"name\": \"LDAP账号注册\", \"notice\": [1, 2, 3], \"creator\": \"admin\", \"ratings\": true, \"comments\": true, \"createBy\": 1, \"updateBy\": 1, \"createdAt\": \"2024-09-15 22:02:06\", \"structure\": {\"edges\": [{\"id\": \"flow1726408760137\", \"seq\": \"1\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"提交\", \"style\": {}, \"source\": \"start1726408739841\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 382.8046875, \"y\": 75, \"index\": 3}, \"startPoint\": {\"x\": 210.3046875, \"y\": 68.5, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 3}, {\"id\": \"flow1726408765640\", \"seq\": \"4\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理成功\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 597.3046875, \"y\": 209, \"index\": 3}, \"sourceAnchor\": 3, \"targetAnchor\": 2}, {\"id\": \"flow1726408806929\", \"seq\": \"2\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"同意\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"receiveTask1726408752430\", \"endPoint\": {\"x\": 637.8046875, \"y\": 186.5, \"index\": 0}, \"startPoint\": {\"x\": 463.8046875, \"y\": 75, \"index\": 1}, \"sourceAnchor\": 1, \"targetAnchor\": 0}, {\"id\": \"flow1726408835676\", \"seq\": \"3\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"拒绝\", \"style\": {}, \"source\": \"userTask1726408742477\", \"target\": \"end1726408758173\", \"endPoint\": {\"x\": 367.3046875, \"y\": 223, \"index\": 2}, \"startPoint\": {\"x\": 423.3046875, \"y\": 97.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 2}, {\"id\": \"flow1726408890148\", \"seq\": \"5\", \"type\": \"flow-polyline-round\", \"clazz\": \"flow\", \"label\": \"处理失败\", \"style\": {}, \"source\": \"receiveTask1726408752430\", \"target\": \"userTask1726408742477\", \"endPoint\": {\"x\": 423.3046875, \"y\": 52.5, \"index\": 0}, \"startPoint\": {\"x\": 637.8046875, \"y\": 231.5, \"index\": 2}, \"sourceAnchor\": 2, \"targetAnchor\": 0}], \"nodes\": [{\"x\": 194.8046875, \"y\": 68.5, \"id\": \"start1726408739841\", \"size\": [30, 30], \"type\": \"start-node\", \"clazz\": \"start\", \"depth\": 0, \"label\": \"发起申请\", \"shape\": \"start-node\", \"style\": {}}, {\"x\": 423.3046875, \"y\": 75, \"id\": \"userTask1726408742477\", \"size\": [80, 44], \"type\": \"user-task-node\", \"clazz\": \"userTask\", \"depth\": 0, \"label\": \"运维人员\", \"shape\": \"user-task-node\", \"style\": {}, \"assignType\": \"assignee\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 637.8046875, \"y\": 209, \"id\": \"receiveTask1726408752430\", \"size\": [80, 44], \"task\": [\"LDAP 账号注册\"], \"type\": \"receive-task-node\", \"clazz\": \"receiveTask\", \"depth\": 0, \"label\": \"处理节点\", \"shape\": \"receive-task-node\", \"style\": {}, \"machine\": [\"10.50.183.112\"], \"assignType\": \"person\", \"activeOrder\": false, \"assignValue\": [2], \"isCounterSign\": false}, {\"x\": 382.8046875, \"y\": 223, \"id\": \"end1726408758173\", \"size\": [30, 30], \"type\": \"end-node\", \"clazz\": \"end\", \"depth\": 0, \"label\": \"工单结束\", \"shape\": \"end-node\", \"style\": {}}], \"combos\": [], \"groups\": []}, \"updatedAt\": \"2024-10-17 19:54:44\", \"categoryId\": 1, \"description\": \"LDAP账号\", \"regenerator\": \"admin\"}');

/*!40000 ALTER TABLE `order_works` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_api
# ------------------------------------------------------------

LOCK TABLES `sys_api` WRITE;
/*!40000 ALTER TABLE `sys_api` DISABLE KEYS */;

INSERT INTO `sys_api` (`id`, `handle`, `title`, `path`, `type`, `action`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
    (5,'go-admin/app/admin/apis.SysLoginLog.Get-fm','登录日志通过id获取','/api/v1/sys-login-log/:id','BUS','GET','2021-05-13 19:59:00.728','2024-08-13 21:10:18.599',NULL,0,0),
    (6,'go-admin/app/admin/apis.SysOperaLog.GetPage-fm','操作日志列表','/api/v1/sys-opera-log','BUS','GET','2021-05-13 19:59:00.778','2024-08-13 21:10:18.599',NULL,0,0),
    (7,'go-admin/app/admin/apis.SysOperaLog.Get-fm','操作日志通过id获取','/api/v1/sys-opera-log/:id','BUS','GET','2021-05-13 19:59:00.821','2024-08-13 21:10:18.599',NULL,0,0),
    (8,'go-admin/common/actions.IndexAction.func1','分类列表','/api/v1/syscategory','BUS','GET','2021-05-13 19:59:00.870','2024-08-13 21:10:18.599',NULL,0,0),
    (9,'go-admin/common/actions.ViewAction.func1','分类通过id获取','/api/v1/syscategory/:id','BUS','GET','2021-05-13 19:59:00.945','2024-08-13 21:10:18.599',NULL,0,0),
    (10,'go-admin/common/actions.IndexAction.func1','内容列表','/api/v1/syscontent','BUS','GET','2021-05-13 19:59:01.005','2024-08-13 21:10:18.599',NULL,0,0),
    (11,'go-admin/common/actions.ViewAction.func1','内容通过id获取','/api/v1/syscontent/:id','BUS','GET','2021-05-13 19:59:01.056','2024-08-13 21:10:18.599',NULL,0,0),
    (15,'go-admin/common/actions.IndexAction.func1','job列表','/api/v1/sysjob','BUS','GET','2021-05-13 19:59:01.248','2024-08-13 21:10:18.599',NULL,0,0),
    (16,'go-admin/common/actions.ViewAction.func1','job通过id获取','/api/v1/sysjob/:id','BUS','GET','2021-05-13 19:59:01.298','2024-08-13 21:10:18.599',NULL,0,0),
    (21,'go-admin/app/admin/apis.SysDictType.GetPage-fm','字典类型列表','/api/v1/dict/type','BUS','GET','2021-05-13 19:59:01.525','2024-08-13 21:10:18.599',NULL,0,0),
    (22,'go-admin/app/admin/apis.SysDictType.GetAll-fm','字典类型查询【代码生成】','/api/v1/dict/type-option-select','SYS','GET','2021-05-13 19:59:01.582','2021-06-13 20:53:48.388',NULL,0,0),
    (23,'go-admin/app/admin/apis.SysDictType.Get-fm','字典类型通过id获取','/api/v1/dict/type/:id','BUS','GET','2021-05-13 19:59:01.632','2024-08-13 21:10:18.599',NULL,0,0),
    (24,'go-admin/app/admin/apis.SysDictData.GetPage-fm','字典数据列表','/api/v1/dict/data','BUS','GET','2021-05-13 19:59:01.684','2024-08-14 16:51:38.287',NULL,0,0),
    (25,'go-admin/app/admin/apis.SysDictData.Get-fm','字典数据通过code获取','/api/v1/dict/data/:dictCode','BUS','GET','2021-05-13 19:59:01.732','2024-08-13 21:10:18.599',NULL,0,0),
    (26,'go-admin/app/admin/apis.SysDictData.GetSysDictDataAll-fm','数据字典根据key获取','/api/v1/dict-data/option-select','SYS','GET','2021-05-13 19:59:01.832','2021-06-17 11:48:40.732',NULL,0,0),
    (27,'go-admin/app/admin/apis.SysDept.GetPage-fm','部门列表','/api/v1/dept','BUS','GET','2021-05-13 19:59:01.940','2024-09-29 17:39:23.384',NULL,0,0),
    (28,'go-admin/app/admin/apis.SysDept.Get-fm','部门通过id获取','/api/v1/dept/:id','BUS','GET','2021-05-13 19:59:02.009','2024-08-13 21:10:18.599',NULL,0,0),
    (29,'go-admin/app/admin/apis.SysDept.Get2Tree-fm','查询部门下拉树【角色权限-自定权限】','/api/v1/deptTree','SYS','GET','2021-05-13 19:59:02.050','2021-06-17 11:48:40.732',NULL,0,0),
    (30,'go-admin/app/admin/apis/tools.(*Gen).GetDBTableList-fm','数据库表列表','/api/v1/db/tables/page','SYS','GET','2021-05-13 19:59:02.098','2021-06-13 20:53:48.730',NULL,0,0),
    (31,'go-admin/app/admin/apis/tools.(*Gen).GetDBColumnList-fm','数据表列列表','/api/v1/db/columns/page','SYS','GET','2021-05-13 19:59:02.140','2021-06-13 20:53:48.771',NULL,0,0),
    (32,'go-admin/app/admin/apis/tools.Gen.GenCode-fm','数据库表生成到项目','/api/v1/gen/toproject/:tableId','SYS','GET','2021-05-13 19:59:02.183','2021-06-13 20:53:48.812',NULL,0,0),
    (33,'go-admin/app/admin/apis/tools.Gen.GenMenuAndApi-fm','数据库表生成到DB','/api/v1/gen/todb/:tableId','SYS','GET','2021-05-13 19:59:02.227','2021-06-13 20:53:48.853',NULL,0,0),
    (34,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesTree-fm','关系表数据【代码生成】','/api/v1/gen/tabletree','SYS','GET','2021-05-13 19:59:02.271','2021-06-13 20:53:48.893',NULL,0,0),
    (35,'go-admin/app/admin/apis/tools.Gen.Preview-fm','生成预览通过id获取','/api/v1/gen/preview/:tableId','SYS','GET','2021-05-13 19:59:02.315','2021-06-13 20:53:48.935',NULL,0,0),
    (36,'go-admin/app/admin/apis/tools.Gen.GenApiToFile-fm','生成api带文件','/api/v1/gen/apitofile/:tableId','SYS','GET','2021-05-13 19:59:02.357','2021-06-13 20:53:48.977',NULL,0,0),
    (37,'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm','验证码获取','/api/v1/getCaptcha','SYS','GET','2021-05-13 19:59:02.405','2021-06-13 20:53:49.020',NULL,0,0),
    (38,'go-admin/app/admin/apis.SysUser.GetInfo-fm','用户信息获取','/api/v1/getinfo','SYS','GET','2021-05-13 19:59:02.447','2021-06-13 20:53:49.065',NULL,0,0),
    (39,'go-admin/app/admin/apis.SysMenu.GetPage-fm','菜单列表','/api/v1/menu','BUS','GET','2021-05-13 19:59:02.497','2024-08-13 21:10:18.599',NULL,0,0),
    (40,'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm','查询菜单下拉树结构【废弃】','/api/v1/menuTreeselect','SYS','GET','2021-05-13 19:59:02.542','2021-06-03 22:37:21.857',NULL,0,0),
    (41,'go-admin/app/admin/apis.SysMenu.Get-fm','菜单通过id获取','/api/v1/menu/:id','BUS','GET','2021-05-13 19:59:02.584','2024-08-13 21:10:18.599',NULL,0,0),
    (42,'go-admin/app/admin/apis.SysMenu.GetMenuRole-fm','角色菜单【顶部左侧菜单】','/api/v1/menurole','SYS','GET','2021-05-13 19:59:02.630','2021-06-13 20:53:49.574',NULL,0,0),
    (43,'go-admin/app/admin/apis.SysMenu.GetMenuIDS-fm','获取角色对应的菜单id数组【废弃】','/api/v1/menuids','SYS','GET','2021-05-13 19:59:02.675','2021-06-03 22:39:52.500',NULL,0,0),
    (44,'go-admin/app/admin/apis.SysRole.GetPage-fm','角色列表','/api/v1/role','BUS','GET','2021-05-13 19:59:02.720','2024-08-14 20:21:50.984',NULL,0,0),
    (45,'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm','菜单权限列表【角色配菜单使用】','/api/v1/roleMenuTreeselect/:roleId','SYS','GET','2021-05-13 19:59:02.762','2021-06-17 11:48:40.732',NULL,0,0),
    (46,'go-admin/app/admin/apis.SysDept.GetDeptTreeRoleSelect-fm','角色部门结构树【自定义数据权限】','/api/v1/roleDeptTreeselect/:roleId','SYS','GET','2021-05-13 19:59:02.809','2021-06-17 11:48:40.732',NULL,0,0),
    (47,'go-admin/app/admin/apis.SysRole.Get-fm','角色通过id获取','/api/v1/role/:id','BUS','GET','2021-05-13 19:59:02.850','2024-08-13 21:10:18.599',NULL,0,0),
    (48,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).RefreshHandler-fm','刷新token','/api/v1/refresh_token','SYS','GET','2021-05-13 19:59:02.892','2021-06-13 20:53:49.278',NULL,0,0),
    (53,'go-admin/app/admin/apis.SysConfig.GetPage-fm','参数列表','/api/v1/config','BUS','GET','2021-05-13 19:59:03.116','2024-08-13 21:10:18.599',NULL,0,0),
    (54,'go-admin/app/admin/apis.SysConfig.Get-fm','参数通过id获取','/api/v1/config/:id','BUS','GET','2021-05-13 19:59:03.157','2024-08-13 21:10:18.599',NULL,0,0),
    (55,'go-admin/app/admin/apis.SysConfig.GetSysConfigByKEYForService-fm','参数通过键名搜索【基础默认配置】','/api/v1/configKey/:configKey','SYS','GET','2021-05-13 19:59:03.198','2021-06-13 20:53:49.745',NULL,0,0),
    (57,'go-admin/app/jobs/apis.SysJob.RemoveJobForService-fm','job移除','/api/v1/job/remove/:id','BUS','GET','2021-05-13 19:59:03.295','2024-08-13 21:10:18.599',NULL,0,0),
    (58,'go-admin/app/jobs/apis.SysJob.StartJobForService-fm','job启动','/api/v1/job/start/:id','BUS','GET','2021-05-13 19:59:03.339','2024-08-13 21:10:18.599',NULL,0,0),
    (59,'go-admin/app/admin/apis.SysPost.GetPage-fm','岗位列表','/api/v1/post','BUS','GET','2021-05-13 19:59:03.381','2024-08-14 17:47:19.864',NULL,0,0),
    (60,'go-admin/app/admin/apis.SysPost.Get-fm','岗位通过id获取','/api/v1/post/:id','BUS','GET','2021-05-13 19:59:03.433','2024-08-13 21:10:18.599',NULL,0,0),
    (62,'go-admin/app/admin/apis.SysConfig.GetSysConfigBySysApp-fm','系统前端参数','/api/v1/app-config','SYS','GET','2021-05-13 19:59:03.526','2021-06-13 20:53:49.994',NULL,0,0),
    (63,'go-admin/app/admin/apis.SysUser.GetProfile-fm','*用户信息获取','/api/v1/user/profile','SYS','GET','2021-05-13 19:59:03.567','2021-06-13 20:53:50.038',NULL,0,0),
    (66,'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).WsClient-fm','链接ws【定时任务log】','/ws/:id/:channel','BUS','GET','2021-05-13 19:59:03.705','2024-08-13 21:10:18.599',NULL,0,0),
    (67,'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).UnWsClient-fm','退出ws【定时任务log】','/wslogout/:id/:channel','BUS','GET','2021-05-13 19:59:03.756','2024-08-13 21:10:18.599',NULL,0,0),
    (68,'go-admin/common/middleware/handler.Ping','*用户基本信息','/info','SYS','GET','2021-05-13 19:59:03.800','2021-06-13 20:53:50.251',NULL,0,0),
    (72,'go-admin/common/actions.CreateAction.func1','分类创建','/api/v1/syscategory','BUS','POST','2021-05-13 19:59:03.982','2024-08-13 21:10:18.599',NULL,0,0),
    (73,'go-admin/common/actions.CreateAction.func1','内容创建','/api/v1/syscontent','BUS','POST','2021-05-13 19:59:04.027','2024-08-13 21:10:18.599',NULL,0,0),
    (76,'go-admin/common/actions.CreateAction.func1','job创建','/api/v1/sysjob','BUS','POST','2021-05-13 19:59:04.164','2024-08-13 21:10:18.599',NULL,0,0),
    (80,'go-admin/app/admin/apis.SysDictData.Insert-fm','字典数据创建','/api/v1/dict/data','BUS','POST','2021-05-13 19:59:04.347','2024-08-13 21:10:18.599',NULL,0,0),
    (81,'go-admin/app/admin/apis.SysDictType.Insert-fm','字典类型创建','/api/v1/dict/type','BUS','POST','2021-05-13 19:59:04.391','2024-08-13 21:44:39.789',NULL,0,0),
    (82,'go-admin/app/admin/apis.SysDept.Insert-fm','部门创建','/api/v1/dept','BUS','POST','2021-05-13 19:59:04.435','2024-08-13 21:10:18.599',NULL,0,0),
    (85,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).LoginHandler-fm','*登录','/api/v1/login','SYS','POST','2021-05-13 19:59:04.597','2021-06-13 20:53:50.784',NULL,0,0),
    (86,'go-admin/common/middleware/handler.LogOut','*退出','/api/v1/logout','SYS','POST','2021-05-13 19:59:04.642','2021-06-13 20:53:50.824',NULL,0,0),
    (87,'go-admin/app/admin/apis.SysConfig.Insert-fm','参数创建','/api/v1/config','BUS','POST','2021-05-13 19:59:04.685','2024-08-13 21:10:18.599',NULL,0,0),
    (88,'go-admin/app/admin/apis.SysMenu.Insert-fm','菜单创建','/api/v1/menu','BUS','POST','2021-05-13 19:59:04.777','2024-08-13 21:10:18.599',NULL,0,0),
    (89,'go-admin/app/admin/apis.SysPost.Insert-fm','岗位创建','/api/v1/post','BUS','POST','2021-05-13 19:59:04.886','2024-08-13 21:10:18.599',NULL,0,0),
    (90,'go-admin/app/admin/apis.SysRole.Insert-fm','角色创建','/api/v1/role','BUS','POST','2021-05-13 19:59:04.975','2024-08-13 21:10:18.599',NULL,0,0),
    (91,'go-admin/app/admin/apis.SysUser.InsetAvatar-fm','*用户头像编辑','/api/v1/user/avatar','SYS','POST','2021-05-13 19:59:05.058','2021-06-13 20:53:51.079',NULL,0,0),
    (92,'go-admin/app/admin/apis.SysApi.Update-fm','接口编辑','/api/v1/sys-api/:id','BUS','PUT','2021-05-13 19:59:05.122','2024-08-13 21:10:18.599',NULL,0,0),
    (95,'go-admin/common/actions.UpdateAction.func1','分类编辑','/api/v1/syscategory/:id','BUS','PUT','2021-05-13 19:59:05.255','2024-08-14 20:21:50.984',NULL,0,0),
    (96,'go-admin/common/actions.UpdateAction.func1','内容编辑','/api/v1/syscontent/:id','BUS','PUT','2021-05-13 19:59:05.299','2024-08-13 21:10:18.599',NULL,0,0),
    (97,'go-admin/common/actions.UpdateAction.func1','job编辑','/api/v1/sysjob','BUS','PUT','2021-05-13 19:59:05.343','2024-08-13 21:10:18.599',NULL,0,0),
    (101,'go-admin/app/admin/apis.SysDictData.Update-fm','字典数据编辑','/api/v1/dict/data/:dictCode','BUS','PUT','2021-05-13 19:59:05.519','2024-08-13 21:10:18.599',NULL,0,0),
    (102,'go-admin/app/admin/apis.SysDictType.Update-fm','字典类型编辑','/api/v1/dict/type/:id','BUS','PUT','2021-05-13 19:59:05.569','2024-08-13 21:10:18.599',NULL,0,0),
    (103,'go-admin/app/admin/apis.SysDept.Update-fm','部门编辑','/api/v1/dept/:id','BUS','PUT','2021-05-13 19:59:05.613','2024-08-13 21:10:18.599',NULL,0,0),
    (104,'go-admin/app/other/apis.SysFileDir.Update-fm','文件夹编辑','/api/v1/file-dir/:id','BUS','PUT','2021-05-13 19:59:05.662','2024-08-13 21:10:18.599',NULL,0,0),
    (105,'go-admin/app/other/apis.SysFileInfo.Update-fm','文件编辑','/api/v1/file-info/:id','BUS','PUT','2021-05-13 19:59:05.709','2024-08-13 21:10:18.599',NULL,0,0),
    (106,'go-admin/app/admin/apis.SysRole.Update-fm','角色编辑','/api/v1/role/:id','BUS','PUT','2021-05-13 19:59:05.752','2024-08-13 21:10:18.599',NULL,0,0),
    (107,'go-admin/app/admin/apis.SysRole.Update2DataScope-fm','角色数据权限修改','/api/v1/roledatascope','BUS','PUT','2021-05-13 19:59:05.803','2024-08-13 21:10:18.599',NULL,0,0),
    (108,'go-admin/app/admin/apis.SysConfig.Update-fm','参数编辑','/api/v1/config/:id','BUS','PUT','2021-05-13 19:59:05.848','2024-08-13 21:10:18.599',NULL,0,0),
    (109,'go-admin/app/admin/apis.SysMenu.Update-fm','编辑菜单','/api/v1/menu/:id','BUS','PUT','2021-05-13 19:59:05.891','2024-08-13 21:10:18.599',NULL,0,0),
    (110,'go-admin/app/admin/apis.SysPost.Update-fm','岗位编辑','/api/v1/post/:id','BUS','PUT','2021-05-13 19:59:05.934','2024-08-13 21:10:18.599',NULL,0,0),
    (111,'go-admin/app/admin/apis.SysUser.UpdatePwd-fm','*用户修改密码','/api/v1/user/pwd','SYS','PUT','2021-05-13 19:59:05.987','2021-06-13 20:53:51.724',NULL,0,0),
    (112,'go-admin/common/actions.DeleteAction.func1','分类删除','/api/v1/syscategory','BUS','DELETE','2021-05-13 19:59:06.030','2024-08-14 20:21:50.984',NULL,0,0),
    (113,'go-admin/common/actions.DeleteAction.func1','内容删除','/api/v1/syscontent','BUS','DELETE','2021-05-13 19:59:06.076','2024-08-13 21:10:18.599',NULL,0,0),
    (114,'go-admin/app/admin/apis.SysLoginLog.Delete-fm','登录日志删除','/api/v1/sys-login-log','BUS','DELETE','2021-05-13 19:59:06.118','2024-08-13 21:10:18.599',NULL,0,0),
    (115,'go-admin/app/admin/apis.SysOperaLog.Delete-fm','操作日志删除','/api/v1/sys-opera-log','BUS','DELETE','2021-05-13 19:59:06.162','2024-08-13 21:10:18.599',NULL,0,0),
    (116,'go-admin/common/actions.DeleteAction.func1','job删除','/api/v1/sysjob','BUS','DELETE','2021-05-13 19:59:06.206','2024-08-13 21:10:18.599',NULL,0,0),
    (117,'go-admin/app/other/apis.SysChinaAreaData.Delete-fm','行政区删除','/api/v1/sys-area-data','BUS','DELETE','2021-05-13 19:59:06.249','2024-08-13 21:10:18.599',NULL,0,0),
    (120,'go-admin/app/admin/apis.SysDictData.Delete-fm','字典数据删除','/api/v1/dict/data','BUS','DELETE','2021-05-13 19:59:06.387','2024-08-13 21:10:18.599',NULL,0,0),
    (121,'go-admin/app/admin/apis.SysDictType.Delete-fm','字典类型删除','/api/v1/dict/type','BUS','DELETE','2021-05-13 19:59:06.432','2024-08-13 21:10:18.599',NULL,0,0),
    (122,'go-admin/app/admin/apis.SysDept.Delete-fm','部门删除','/api/v1/dept/:id','BUS','DELETE','2021-05-13 19:59:06.475','2024-08-13 21:10:18.599',NULL,0,0),
    (123,'go-admin/app/other/apis.SysFileDir.Delete-fm','文件夹删除','/api/v1/file-dir/:id','BUS','DELETE','2021-05-13 19:59:06.520','2024-08-13 21:10:18.599',NULL,0,0),
    (124,'go-admin/app/other/apis.SysFileInfo.Delete-fm','文件删除','/api/v1/file-info/:id','BUS','DELETE','2021-05-13 19:59:06.566','2024-08-13 21:10:18.599',NULL,0,0),
    (125,'go-admin/app/admin/apis.SysConfig.Delete-fm','参数删除','/api/v1/config','BUS','DELETE','2021-05-13 19:59:06.612','2024-08-13 21:10:18.599',NULL,0,0),
    (126,'go-admin/app/admin/apis.SysMenu.Delete-fm','删除菜单','/api/v1/menu','BUS','DELETE','2021-05-13 19:59:06.654','2024-08-13 21:10:18.599',NULL,0,0),
    (127,'go-admin/app/admin/apis.SysPost.Delete-fm','岗位删除','/api/v1/post/:id','BUS','DELETE','2021-05-13 19:59:06.702','2024-08-13 21:10:18.599',NULL,0,0),
    (128,'go-admin/app/admin/apis.SysRole.Delete-fm','角色删除','/api/v1/role','BUS','DELETE','2021-05-13 19:59:06.746','2024-08-13 21:10:18.599',NULL,0,0),
    (131,'github.com/go-admin-team/go-admin-core/tools/transfer.Handler.func1','系统指标','/api/v1/metrics','SYS','GET','2021-05-17 22:13:55.933','2021-06-13 20:53:49.614',NULL,0,0),
    (132,'go-admin/app/other/router.registerMonitorRouter.func1','健康状态','/api/v1/health','SYS','GET','2021-05-17 22:13:56.285','2021-06-13 20:53:49.951',NULL,0,0),
    (133,'go-admin/app/admin/apis.HelloWorld','项目默认接口','/','SYS','GET','2021-05-24 10:30:44.553','2021-06-13 20:53:47.406',NULL,0,0),
    (134,'go-admin/app/other/apis.ServerMonitor.ServerInfo-fm','服务器基本状态','/api/v1/server-monitor','SYS','GET','2021-05-24 10:30:44.937','2021-06-13 20:53:48.255',NULL,0,0),
    (135,'go-admin/app/admin/apis.SysApi.GetPage-fm','接口列表','/api/v1/sys-api','BUS','GET','2021-05-24 11:37:53.303','2024-08-13 21:10:18.599',NULL,0,0),
    (136,'go-admin/app/admin/apis.SysApi.Get-fm','接口通过id获取','/api/v1/sys-api/:id','BUS','GET','2021-05-24 11:37:53.359','2024-08-13 21:10:18.599',NULL,0,0),
    (137,'go-admin/app/admin/apis.SysLoginLog.GetPage-fm','登录日志列表','/api/v1/sys-login-log','BUS','GET','2021-05-24 11:47:30.397','2024-08-13 21:10:18.599',NULL,0,0),
    (138,'go-admin/app/other/apis.File.UploadFile-fm','文件上传','/api/v1/public/uploadFile','SYS','POST','2021-05-25 19:16:18.493','2021-06-13 20:53:50.866',NULL,0,0),
    (139,'go-admin/app/admin/apis.SysConfig.Update2Set-fm','参数信息修改【参数配置】','/api/v1/set-config','BUS','PUT','2021-05-27 09:45:14.853','2024-08-13 21:10:18.599',NULL,0,0),
    (140,'go-admin/app/admin/apis.SysConfig.Get2Set-fm','参数获取全部【配置使用】','/api/v1/set-config','BUS','GET','2021-05-27 11:54:14.384','2024-08-13 21:10:18.599',NULL,0,0),
    (141,'go-admin/app/admin/apis.SysUser.GetPage-fm','管理员列表','/api/v1/sys-user','BUS','GET','2021-06-13 19:24:57.111','2024-09-29 17:39:23.384',NULL,0,0),
    (142,'go-admin/app/admin/apis.SysUser.Get-fm','管理员通过id获取','/api/v1/sys-user/:id','BUS','GET','2021-06-13 19:24:57.188','2024-08-13 21:10:18.599',NULL,0,0),
    (143,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesInfo-fm','','/api/v1/sys/tables/info','','GET','2021-06-13 19:24:57.437','2021-06-13 20:53:48.047',NULL,0,0),
    (144,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTables-fm','','/api/v1/sys/tables/info/:tableId','','GET','2021-06-13 19:24:57.510','2021-06-13 20:53:48.088',NULL,0,0),
    (145,'go-admin/app/admin/apis/tools.(*SysTable).GetSysTableList-fm','','/api/v1/sys/tables/page','','GET','2021-06-13 19:24:57.582','2021-06-13 20:53:48.128',NULL,0,0),
    (146,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/static/*filepath','','GET','2021-06-13 19:24:59.641','2021-06-13 20:53:50.081',NULL,0,0),
    (147,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','','/swagger/*any','','GET','2021-06-13 19:24:59.713','2021-06-13 20:53:50.123',NULL,0,0),
    (148,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/form-generator/*filepath','','GET','2021-06-13 19:24:59.914','2021-06-13 20:53:50.295',NULL,0,0),
    (149,'go-admin/app/admin/apis/tools.(*SysTable).InsertSysTable-fm','','/api/v1/sys/tables/info','','POST','2021-06-13 19:25:00.163','2021-06-13 20:53:50.539',NULL,0,0),
    (150,'go-admin/app/admin/apis.SysUser.Insert-fm','管理员创建','/api/v1/sys-user','BUS','POST','2021-06-13 19:25:00.233','2024-08-13 21:10:18.599',NULL,0,0),
    (151,'go-admin/app/admin/apis.SysUser.Update-fm','管理员编辑','/api/v1/sys-user','BUS','PUT','2021-06-13 19:25:00.986','2024-08-13 21:10:18.599',NULL,0,0),
    (152,'go-admin/app/admin/apis/tools.(*SysTable).UpdateSysTable-fm','','/api/v1/sys/tables/info','','PUT','2021-06-13 19:25:01.149','2021-06-13 20:53:51.375',NULL,0,0),
    (153,'go-admin/app/admin/apis.SysRole.Update2Status-fm','更新角色状态','/api/v1/role-status','BUS','PUT','2021-06-13 19:25:01.446','2024-08-13 21:10:18.599',NULL,0,0),
    (154,'go-admin/app/admin/apis.SysUser.ResetPwd-fm','重置用户密码','/api/v1/user/pwd/reset','BUS','PUT','2021-06-13 19:25:01.601','2024-08-13 21:10:18.599',NULL,0,0),
    (155,'go-admin/app/admin/apis.SysUser.UpdateStatus-fm','更新用户状态','/api/v1/user/status','BUS','PUT','2021-06-13 19:25:01.671','2024-08-13 21:10:18.599',NULL,0,0),
    (156,'go-admin/app/admin/apis.SysUser.Delete-fm','管理员删除','/api/v1/sys-user','BUS','DELETE','2021-06-13 19:25:02.043','2024-08-13 21:10:18.599',NULL,0,0),
    (157,'go-admin/app/admin/apis/tools.(*SysTable).DeleteSysTables-fm','','/api/v1/sys/tables/info/:tableId','','DELETE','2021-06-13 19:25:02.283','2021-06-13 20:53:52.367',NULL,0,0),
    (158,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/static/*filepath','','HEAD','2021-06-13 19:25:02.734','2021-06-13 20:53:52.791',NULL,0,0),
    (159,'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1','','/form-generator/*filepath','','HEAD','2021-06-13 19:25:02.808','2021-06-13 20:53:52.838',NULL,0,0),
    (160,'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm','验证码','/api/v1/captcha','SYS','GET','2024-07-08 16:17:04.505','2024-08-13 11:40:14.526',NULL,0,0),
    (161,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','','/swagger/admin/*any','','GET','2024-07-08 16:17:04.514','2024-07-08 16:17:04.514',NULL,0,0),
    (162,'go-admin/app/admin/apis.SysUser.UpdatePwd-fm','用户密码设置','/api/v1/user/pwd/set','BUS','PUT','2024-07-08 16:17:04.528','2024-08-13 21:10:18.599',NULL,0,0),
    (163,'go-admin/app/admin/apis.SysDept.Delete-fm','删除部门','/api/v1/dept','BUS','DELETE','2024-07-08 16:17:04.532','2024-08-13 21:10:18.599',NULL,0,0),
    (164,'go-admin/app/admin/apis.SysPost.Delete-fm','删除岗位','/api/v1/post','BUS','DELETE','2024-07-08 16:17:04.534','2024-08-13 21:10:18.599',NULL,0,0),
    (225,'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).LoginHandler-fm','ldap用户登陆','/api/v1/ldap/login','SYS','POST','2024-08-11 17:16:07.037','2024-08-13 12:21:43.105',NULL,0,0),
    (229,'go-admin/app/system/apis.SysNotify.NotifySend-fm','发送飞书通知消息','/api/v1/notify','SYS','POST','2024-08-16 14:14:01.595','2024-08-17 17:29:37.148',NULL,0,0),
    (230,'go-admin/app/system/apis.SysUser.GetSpecify-fm','根据用户名获取指定用户信息','/api/v1/sys-user/specify-user','SYS','GET','2024-08-17 17:41:34.158','2024-08-17 17:44:51.421',NULL,0,0),
    (243,'github.com/swaggo/gin-swagger.CustomWrapHandler.func1','swagger接口页面','/swagger/smart/*any','SYS','GET','2024-09-01 18:41:58.252','2024-09-08 16:10:24.164',NULL,0,0),
    (244,'go-admin/app/smart/apis.OrderWorks.GetPage-fm','获取工单列表','/api/v1/order-works','BUS','GET','2024-09-02 11:37:58.114','2024-09-29 17:39:23.384',NULL,0,0),
    (245,'go-admin/app/smart/apis.OrderWorks.MyBacklogGet-fm','工单列表我的待办','/api/v1/order-works/myBacklog','BUS','GET','2024-09-02 11:37:58.123','2024-09-29 17:39:23.384',NULL,0,0),
    (246,'go-admin/app/smart/apis.OrderWorks.CreatedByMe-fm','工单列表我的创建','/api/v1/order-works/createdByMe','BUS','GET','2024-09-02 11:37:58.127','2024-09-29 17:39:23.384',NULL,0,0),
    (247,'go-admin/app/smart/apis.OrderWorks.RelatedToMe-fm','工单列表与我相关','/api/v1/order-works/relatedToMe','BUS','GET','2024-09-02 11:37:58.133','2024-09-29 17:39:23.384',NULL,0,0),
    (248,'go-admin/app/smart/apis.OperationHistory.GetOperationHistory-fm','工单历史记录','/api/v1/order-works/history','BUS','GET','2024-09-02 11:37:58.137','2024-09-29 17:39:23.384',NULL,0,0),
    (249,'go-admin/app/smart/apis.WorksNotify.GetNotify-fm','系统通知','/api/v1/order-works/notify','BUS','GET','2024-09-02 11:37:58.139','2024-09-29 17:39:23.384',NULL,0,0),
    (250,'go-admin/app/smart/apis.OrderWorks.Get-fm','根据id查询工单','/api/v1/order-works/:id','BUS','GET','2024-09-02 11:37:58.140','2024-09-29 17:39:23.384',NULL,0,0),
    (251,'go-admin/app/smart/apis.OrderCategory.GetPage-fm','工单分类列表','/api/v1/order-category','BUS','GET','2024-09-02 11:37:58.142','2024-09-29 17:39:23.384',NULL,0,0),
    (252,'go-admin/app/smart/apis.OrderCategory.Get-fm','根据id查询工单分类','/api/v1/order-category/:id','BUS','GET','2024-09-02 11:37:58.144','2024-09-08 16:23:59.815',NULL,0,0),
    (253,'go-admin/app/smart/apis.OrderItems.GetPage-fm','获取工单项目列表','/api/v1/order-items','BUS','GET','2024-09-02 11:37:58.146','2024-09-29 17:39:23.384',NULL,0,0),
    (254,'go-admin/app/smart/apis.OrderItems.Get-fm','根据id获取工单项目','/api/v1/order-items/:id','BUS','GET','2024-09-02 11:37:58.147','2024-09-08 16:23:59.815',NULL,0,0),
    (255,'go-admin/app/smart/apis.OrderTask.GetPage-fm','工单任务列表','/api/v1/order-task','BUS','GET','2024-09-02 11:37:58.149','2024-09-29 17:39:23.384',NULL,0,0),
    (256,'go-admin/app/smart/apis.OrderTask.Get-fm','根据id查询工单任务','/api/v1/order-task/:id','BUS','GET','2024-09-02 11:37:58.150','2024-09-29 17:39:23.384',NULL,0,0),
    (257,'go-admin/app/smart/apis.FlowManage.GetPage-fm','工单流程管理列表','/api/v1/flow-manage','BUS','GET','2024-09-02 11:37:58.158','2024-09-08 16:20:18.821',NULL,0,0),
    (258,'go-admin/app/smart/apis.FlowManage.Get-fm','根据id查询流程','/api/v1/flow-manage/:id','BUS','GET','2024-09-02 11:37:58.160','2024-09-08 16:23:39.672',NULL,0,0),
    (259,'go-admin/app/smart/apis.FlowTemplates.GetPage-fm','工单模板列表','/api/v1/flow-templates','BUS','GET','2024-09-02 11:37:58.162','2024-09-29 17:39:23.384',NULL,0,0),
    (260,'go-admin/app/smart/apis.FlowTemplates.Get-fm','根据id获取工单模板数据	','/api/v1/flow-templates/:id','BUS','GET','2024-09-02 11:37:58.164','2024-09-08 16:23:24.404',NULL,0,0),
    (261,'go-admin/app/smart/apis.ExecMachine.GetPage-fm','执行节点列表','/api/v1/exec-machine','BUS','GET','2024-09-02 11:37:58.168','2024-09-08 15:53:16.464',NULL,0,0),
    (262,'go-admin/app/smart/apis.ExecMachine.Get-fm','根据id查询执行节点','/api/v1/exec-machine/:id','BUS','GET','2024-09-02 11:37:58.169','2024-09-08 15:53:30.896',NULL,0,0),
    (263,'go-admin/app/smart/apis.OrderWorks.Insert-fm','新建工单','/api/v1/order-works','BUS','POST','2024-09-02 11:37:58.178','2024-09-29 17:39:23.384',NULL,0,0),
    (264,'go-admin/app/smart/apis.WorksNotify.CreateNotify-fm','创建通知消息','/api/v1/order-works/notify','BUS','POST','2024-09-02 11:37:58.179','2024-09-29 17:39:23.384',NULL,0,0),
    (265,'go-admin/app/smart/apis.OrderCategory.Insert-fm','创建工单分类','/api/v1/order-category','BUS','POST','2024-09-02 11:37:58.181','2024-09-08 16:19:25.627',NULL,0,0),
    (266,'go-admin/app/smart/apis.OrderItems.Insert-fm','新建工单项目','/api/v1/order-items','BUS','POST','2024-09-02 11:37:58.183','2024-09-08 16:24:23.202',NULL,0,0),
    (267,'go-admin/app/smart/apis.OrderTask.Insert-fm','创建工单任务','/api/v1/order-task','BUS','POST','2024-09-02 11:37:58.185','2024-09-29 17:39:23.384',NULL,0,0),
    (268,'go-admin/app/smart/apis.FlowManage.Insert-fm','创建工单流程','/api/v1/flow-manage','BUS','POST','2024-09-02 11:37:58.187','2024-09-08 16:20:29.403',NULL,0,0),
    (269,'go-admin/app/smart/apis.FlowManage.Clone-fm','克隆工单流程','/api/v1/flow-manage/:id','BUS','POST','2024-09-02 11:37:58.189','2024-09-08 16:21:00.032',NULL,0,0),
    (270,'go-admin/app/smart/apis.FlowTemplates.Insert-fm','创建工单模板','/api/v1/flow-templates','BUS','POST','2024-09-02 11:37:58.191','2024-09-08 16:21:39.897',NULL,0,0),
    (271,'go-admin/app/smart/apis.ExecMachine.Insert-fm','创建执行节点','/api/v1/exec-machine','BUS','POST','2024-09-02 11:37:58.197','2024-09-08 15:54:21.248',NULL,0,0),
    (272,'go-admin/app/smart/apis.(*ExecMachine).TestConnection-fm','测试连接执行节点','/api/v1/exec-machine/testConnection','BUS','POST','2024-09-02 11:37:58.198','2024-09-08 15:53:54.510',NULL,0,0),
    (273,'go-admin/app/smart/apis.OrderWorks.Update-fm','更新工单','/api/v1/order-works','BUS','PUT','2024-09-02 11:37:58.207','2024-09-29 17:39:23.384',NULL,0,0),
    (274,'go-admin/app/smart/apis.OrderWorks.Handle-fm','处理工单','/api/v1/order-works/handle','BUS','PUT','2024-09-02 11:37:58.212','2024-09-29 17:39:23.384',NULL,0,0),
    (275,'go-admin/app/smart/apis.WorksNotify.UpdateNotify-fm','处理通知消息','/api/v1/order-works/notify','BUS','PUT','2024-09-02 11:37:58.215','2024-09-29 17:39:23.384',NULL,0,0),
    (276,'go-admin/app/smart/apis.OrderCategory.Update-fm','更新工单分类','/api/v1/order-category','BUS','PUT','2024-09-02 11:37:58.217','2024-09-08 16:19:57.195',NULL,0,0),
    (277,'go-admin/app/smart/apis.OrderItems.Update-fm','更新工单项目','/api/v1/order-items','BUS','PUT','2024-09-02 11:37:58.219','2024-09-08 16:07:16.325',NULL,0,0),
    (278,'go-admin/app/smart/apis.OrderTask.Update-fm','更新工单任务','/api/v1/order-task','BUS','PUT','2024-09-02 11:37:58.221','2024-09-29 17:39:23.384',NULL,0,0),
    (279,'go-admin/app/smart/apis.FlowManage.Update-fm','更新流程管理','/api/v1/flow-manage','BUS','PUT','2024-09-02 11:37:58.233','2024-09-08 16:20:43.772',NULL,0,0),
    (280,'go-admin/app/smart/apis.FlowTemplates.Update-fm','更新工单模板','/api/v1/flow-templates','BUS','PUT','2024-09-02 11:37:58.235','2024-09-08 16:21:46.311',NULL,0,0),
    (281,'go-admin/app/smart/apis.ExecMachine.Update-fm','更新执行节点','/api/v1/exec-machine','BUS','PUT','2024-09-02 11:37:58.236','2024-09-08 15:54:07.100',NULL,0,0),
    (282,'go-admin/app/smart/apis.OrderCategory.Delete-fm','删除工单分类','/api/v1/order-category','BUS','DELETE','2024-09-02 11:37:58.242','2024-09-08 16:19:46.850',NULL,0,0),
    (283,'go-admin/app/smart/apis.OrderItems.Delete-fm','删除工单项目','/api/v1/order-items','BUS','DELETE','2024-09-02 11:37:58.243','2024-09-08 16:07:24.286',NULL,0,0),
    (284,'go-admin/app/smart/apis.OrderTask.Delete-fm','删除工单任务','/api/v1/order-task','BUS','DELETE','2024-09-02 11:37:58.245','2024-09-29 17:39:23.384',NULL,0,0),
    (285,'go-admin/app/smart/apis.OrderWorks.Delete-fm','关闭工单','/api/v1/order-works','BUS','DELETE','2024-09-02 11:37:58.246','2024-09-08 15:59:32.750',NULL,0,0),
    (286,'go-admin/app/smart/apis.FlowManage.Delete-fm','删除工单流程','/api/v1/flow-manage','BUS','DELETE','2024-09-02 11:37:58.249','2024-09-08 16:20:53.283',NULL,0,0),
    (287,'go-admin/app/smart/apis.FlowTemplates.Delete-fm','删除工单模板','/api/v1/flow-templates','BUS','DELETE','2024-09-02 11:37:58.251','2024-09-08 16:21:53.786',NULL,0,0),
    (288,'go-admin/app/smart/apis.ExecMachine.Delete-fm','删除执行节点','/api/v1/exec-machine','BUS','DELETE','2024-09-02 11:37:58.253','2024-09-08 15:54:13.902',NULL,0,0),
    (289,'go-admin/app/system/apis.SysApi.DeleteSysApi-fm','更新接口','/api/v1/sys-api','SYS','PUT','2024-09-08 15:40:00.260','2024-10-27 15:51:49.136',NULL,0,0),
    (290,'go-admin/app/system/apis.SysApi.DeleteSysApi-fm','删除接口','/api/v1/sys-api','SYS','DELETE','2024-09-08 15:42:00.584','2024-10-27 15:51:52.811',NULL,0,0),
    (291,'go-admin/app/smart/apis.ExecutionHistory.GetHistoryTaskPage-fm','','/api/v1/exec-machine/history','','GET','2024-09-12 11:39:31.010','2024-09-12 11:39:31.010','2024-10-27 15:58:50.120',0,0),
    (292,'go-admin/app/smart/apis.WebSocketAPI.WebSocketHandler-fm','','/api/v1/order-works/ws/exec','','GET','2024-09-12 17:36:24.754','2024-09-12 17:36:24.754','2024-10-27 15:58:04.896',0,0),
    (293,'go-admin/app/smart/apis.WebSocketAPI.WebSocketHandler-fm','','/api/v1/order-works/ws/exec-task','','GET','2024-09-12 19:06:52.563','2024-09-12 19:06:52.563','2024-10-27 15:57:54.933',0,0),
    (294,'go-admin/app/system/apis.GoAdmin','','/ws/task','','GET','2024-09-13 14:26:32.848','2024-09-13 14:26:32.848','2024-10-27 15:57:32.531',0,0),
    (295,'go-admin/common/utils.WsHandler','','/api/v1/ws/:id/task','','GET','2024-09-13 14:56:55.550','2024-09-13 14:56:55.550','2024-10-27 15:58:21.809',0,0),
    (296,'go-admin/common/utils.WsHandler','','/api/v1/api/v1/ws/task/:id','','GET','2024-09-14 14:57:53.781','2024-09-14 14:57:53.781','2024-10-27 15:56:43.114',0,0),
    (297,'go-admin/common/utils.WsHandler','查询任务实时日志','/api/v1/ws/task/:id','BUS','GET','2024-09-14 15:23:20.312','2024-10-27 15:56:28.586',NULL,0,0),
    (298,'go-admin/app/smart/apis.ExecutionHistory.Delete-fm','','/api/v1/exec-machine/history','','DELETE','2024-09-19 16:21:35.460','2024-09-19 16:21:35.460','2024-10-27 15:53:45.309',0,0),
    (301,'go-admin/app/smart/apis.OrderRatingAPI.Insert-fm','新增工单评价','/api/v1/order-ratings','BUS','POST','2024-09-23 17:41:13.101','2024-09-29 17:39:23.384',NULL,0,0),
    (308,'go-admin/app/smart/apis.OrderRatingAPI.Update-fm','更新工单评价','/api/v1/order-ratings','BUS','PUT','2024-09-23 18:02:46.843','2024-09-29 17:39:23.384',NULL,0,0),
    (309,'go-admin/app/smart/apis.OrderRatingAPI.Delete-fm','删除工单评价','/api/v1/order-ratings','BUS','DELETE','2024-09-23 18:05:42.609','2024-09-29 17:39:23.384',NULL,0,0),
    (310,'go-admin/app/smart/apis.OrderCommentAPI.Insert-fm','新增工单留言','/api/v1/order-comments','BUS','POST','2024-09-23 18:18:09.642','2024-09-29 17:39:23.384',NULL,0,0),
    (311,'go-admin/app/smart/apis.OrderCommentAPI.Update-fm','更新工单留言','/api/v1/order-comments','BUS','PUT','2024-09-23 18:18:09.655','2024-09-29 17:39:23.384',NULL,0,0),
    (312,'go-admin/app/smart/apis.OrderCommentAPI.Delete-fm','删除工单留言','/api/v1/order-comments','BUS','DELETE','2024-09-23 18:18:09.662','2024-09-29 17:39:23.384',NULL,0,0),
    (314,'go-admin/app/smart/apis.OrderCommentAPI.Get-fm','根据工单ID获取工单留言','/api/v1/order-comments/:orderID','BUS','GET','2024-09-23 18:24:10.989','2024-09-29 17:39:23.384',NULL,0,0),
    (315,'go-admin/app/smart/apis.OrderRatingAPI.Get-fm','获取工单评价','/api/v1/order-ratings/:orderID','BUS','GET','2024-09-24 11:03:13.995','2024-09-29 17:39:23.384',NULL,0,0),
    (316,'go-admin/app/smart/apis.OrderStatistics.GetOrderStatistics-fm','获取统计工单数据','/api/v1/statistics/orders','BUS','GET','2024-09-25 17:01:25.826','2024-09-26 17:48:54.310',NULL,0,0),
    (317,'go-admin/app/smart/apis.OrderStatistics.GetOrderCountByPeriod-fm','获取统计工单折线图数据','/api/v1/statistics/orders/count','BUS','GET','2024-09-25 17:01:25.831','2024-09-26 17:49:15.676',NULL,0,0),
    (318,'go-admin/app/smart/apis.OrderStatistics.GetOrderRatings-fm','获取统计评分柱状图数据','/api/v1/statistics/ratings','BUS','GET','2024-09-25 18:30:44.327','2024-09-26 17:49:30.387',NULL,0,0),
    (322,'go-admin/app/smart/apis.UserFavorite.GetUserFavorites-fm','获取收藏的工单','/api/v1/favorite','BUS','GET','2024-09-29 15:24:18.864','2024-09-29 17:39:23.384',NULL,0,0),
    (323,'go-admin/app/smart/apis.UserFavorite.AddFavorite-fm','添加收藏的工单','/api/v1/favorite','BUS','POST','2024-09-29 15:24:18.880','2024-09-29 17:39:23.384',NULL,0,0),
    (324,'go-admin/app/smart/apis.UserFavorite.RemoveFavorite-fm','取消收藏的工单','/api/v1/favorite','BUS','DELETE','2024-09-29 15:24:18.902','2024-09-29 17:39:23.384',NULL,0,0),
    (325,'smart-api/app/smart/apis.OrderTask.GetHistoryTaskPage-fm','获取历史执行后的任务','/api/v1/order-task/history','BUS','GET','2024-10-23 11:14:10.937','2024-10-27 15:53:21.363',NULL,0,0),
    (326,'smart-api/app/smart/apis.OrderTask.DeleteHistoryTask-fm','删除历史执行后的任务','/api/v1/order-task/history','BUS','DELETE','2024-10-23 11:14:10.983','2024-10-27 15:53:10.164',NULL,0,0),
    (327,'smart-api/app/smart/apis.(*OrderTask).GetTaskLogsByID-fm','查询历史任务日志','/api/v1/order-task/history/:id/logs','BUS','GET','2024-10-23 13:21:10.047','2024-10-27 15:52:50.392',NULL,0,0);

/*!40000 ALTER TABLE `sys_api` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_casbin_rule
# ------------------------------------------------------------

LOCK TABLES `sys_casbin_rule` WRITE;
/*!40000 ALTER TABLE `sys_casbin_rule` DISABLE KEYS */;

INSERT INTO `sys_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`)
VALUES
    (762,'p','general','/api/v1/dept','GET','','','','',''),
    (773,'p','general','/api/v1/favorite','DELETE','','','','',''),
    (771,'p','general','/api/v1/favorite','GET','','','','',''),
    (772,'p','general','/api/v1/favorite','POST','','','','',''),
    (781,'p','general','/api/v1/flow-templates','GET','','','','',''),
    (769,'p','general','/api/v1/order-category','GET','','','','',''),
    (793,'p','general','/api/v1/order-comments','DELETE','','','','',''),
    (791,'p','general','/api/v1/order-comments','POST','','','','',''),
    (792,'p','general','/api/v1/order-comments','PUT','','','','',''),
    (794,'p','general','/api/v1/order-comments/:orderID','GET','','','','',''),
    (770,'p','general','/api/v1/order-items','GET','','','','',''),
    (789,'p','general','/api/v1/order-ratings','DELETE','','','','',''),
    (787,'p','general','/api/v1/order-ratings','POST','','','','',''),
    (788,'p','general','/api/v1/order-ratings','PUT','','','','',''),
    (790,'p','general','/api/v1/order-ratings/:orderID','GET','','','','',''),
    (786,'p','general','/api/v1/order-task','DELETE','','','','',''),
    (763,'p','general','/api/v1/order-task','GET','','','','',''),
    (783,'p','general','/api/v1/order-task','POST','','','','',''),
    (784,'p','general','/api/v1/order-task','PUT','','','','',''),
    (785,'p','general','/api/v1/order-task/:id','GET','','','','',''),
    (778,'p','general','/api/v1/order-works','GET','','','','',''),
    (764,'p','general','/api/v1/order-works','POST','','','','',''),
    (774,'p','general','/api/v1/order-works','PUT','','','','',''),
    (780,'p','general','/api/v1/order-works/:id','GET','','','','',''),
    (766,'p','general','/api/v1/order-works/createdByMe','GET','','','','',''),
    (768,'p','general','/api/v1/order-works/handle','PUT','','','','',''),
    (779,'p','general','/api/v1/order-works/history','GET','','','','',''),
    (765,'p','general','/api/v1/order-works/myBacklog','GET','','','','',''),
    (775,'p','general','/api/v1/order-works/notify','GET','','','','',''),
    (776,'p','general','/api/v1/order-works/notify','POST','','','','',''),
    (777,'p','general','/api/v1/order-works/notify','PUT','','','','',''),
    (767,'p','general','/api/v1/order-works/relatedToMe','GET','','','','',''),
    (782,'p','general','/api/v1/sys-user','GET','','','','','');

/*!40000 ALTER TABLE `sys_casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_columns
# ------------------------------------------------------------

LOCK TABLES `sys_columns` WRITE;
/*!40000 ALTER TABLE `sys_columns` DISABLE KEYS */;

INSERT INTO `sys_columns` (`column_id`, `table_id`, `column_name`, `column_comment`, `column_type`, `go_type`, `go_field`, `json_field`, `is_pk`, `is_increment`, `is_required`, `is_insert`, `is_edit`, `is_list`, `is_query`, `query_type`, `html_type`, `dict_type`, `sort`, `list`, `pk`, `required`, `super_column`, `usable_column`, `increment`, `insert`, `edit`, `query`, `remark`, `fk_table_name`, `fk_table_name_class`, `fk_table_name_package`, `fk_label_id`, `fk_label_name`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
    (1,1,'id','','bigint','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.115','2024-08-12 20:27:59.115',NULL,0,0),
    (2,1,'name','','varchar(100)','string','Name','name','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.122','2024-08-12 20:27:59.122',NULL,0,0),
    (3,1,'icon','','varchar(50)','string','Icon','icon','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.126','2024-08-12 20:27:59.126',NULL,0,0),
    (4,1,'categoryId','','bigint unsigned','string','CategoryId','categoryId','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.131','2024-08-12 20:27:59.131',NULL,0,0),
    (5,1,'task','','json','string','Task','task','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.133','2024-08-12 20:27:59.133',NULL,0,0),
    (6,1,'notice','','json','string','Notice','notice','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.134','2024-08-12 20:27:59.134',NULL,0,0),
    (7,1,'comments','','tinyint(1)','string','Comments','comments','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.136','2024-08-12 20:27:59.136',NULL,0,0),
    (8,1,'ratings','','tinyint(1)','string','Ratings','ratings','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.137','2024-08-12 20:27:59.137',NULL,0,0),
    (9,1,'description','','varchar(512)','string','Description','description','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.139','2024-08-12 20:27:59.139',NULL,0,0),
    (10,1,'creator','','varchar(20)','string','Creator','creator','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.140','2024-08-12 20:27:59.140',NULL,0,0),
    (11,1,'regenerator','','varchar(20)','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.142','2024-08-12 20:27:59.142',NULL,0,0),
    (12,1,'structure','','json','string','Structure','structure','0','','0','1','','','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.144','2024-08-12 20:27:59.144',NULL,0,0),
    (13,1,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.146','2024-08-12 20:27:59.146',NULL,0,0),
    (14,1,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.148','2024-08-12 20:27:59.148',NULL,0,0),
    (15,1,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',15,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.150','2024-08-12 20:27:59.150',NULL,0,0),
    (16,1,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',16,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.152','2024-08-12 20:27:59.152',NULL,0,0),
    (17,1,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',17,'',0,0,0,0,0,1,0,0,'','','','','','','2024-08-12 20:27:59.153','2024-08-12 20:27:59.153',NULL,0,0),
    (18,2,'id','','bigint','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.303','2024-09-01 23:28:33.303',NULL,0,0),
    (19,2,'ip','','varchar(15)','string','Ip','ip','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.306','2024-09-01 23:32:57.624',NULL,0,0),
    (20,2,'hostname','','varchar(255)','string','Hostname','hostname','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.308','2024-09-01 23:32:57.627',NULL,0,0),
    (21,2,'username','','varchar(45)','string','Username','username','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.316','2024-09-01 23:32:57.630',NULL,0,0),
    (22,2,'password','','varchar(100)','string','Password','password','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.319','2024-09-01 23:32:57.632',NULL,0,0),
    (23,2,'port','','bigint','string','Port','port','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.320','2024-09-01 23:32:57.633',NULL,0,0),
    (24,2,'heartbeat','','timestamp','time.Time','Heartbeat','heartbeat','0','','0','1','','','','EQ','datetime','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.321','2024-09-01 23:32:57.635',NULL,0,0),
    (25,2,'status','','bigint','string','Status','status','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.322','2024-09-01 23:32:57.636',NULL,0,0),
    (26,2,'auth_type','','bigint','string','AuthType','authType','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.323','2024-09-01 23:32:57.638',NULL,0,0),
    (27,2,'key_path','','varchar(255)','string','KeyPath','keyPath','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.326','2024-09-01 23:32:57.640',NULL,0,0),
    (28,2,'creator','','varchar(45)','string','Creator','creator','0','','0','1','','','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.328','2024-09-01 23:32:57.641',NULL,0,0),
    (29,2,'regenerator','','varchar(20)','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.329','2024-09-01 23:32:57.643',NULL,0,0),
    (30,2,'description','','longtext','string','Description','description','0','','0','1','','','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.330','2024-09-01 23:32:57.644',NULL,0,0),
    (31,2,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.332','2024-09-01 23:28:33.332',NULL,0,0),
    (32,2,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.333','2024-09-01 23:28:33.333',NULL,0,0),
    (33,2,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',16,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.334','2024-09-01 23:28:33.334',NULL,0,0),
    (34,2,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',17,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.335','2024-09-01 23:28:33.335',NULL,0,0),
    (35,2,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',18,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-01 23:28:33.336','2024-09-01 23:28:33.336',NULL,0,0),
    (36,3,'id','','bigint unsigned','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.440','2024-09-02 10:44:06.440',NULL,0,0),
    (37,3,'name','','varchar(100)','string','Name','name','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.448','2024-09-02 10:44:06.448',NULL,0,0),
    (38,3,'description','','varchar(200)','string','Description','description','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.449','2024-09-02 10:44:06.449',NULL,0,0),
    (39,3,'bindCount','','bigint','string','BindCount','bindCount','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.451','2024-09-02 10:44:06.451',NULL,0,0),
    (40,3,'form_data','','json','string','FormData','formData','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.452','2024-09-02 10:44:06.452',NULL,0,0),
    (41,3,'categoryId','','bigint unsigned','string','CategoryId','categoryId','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.454','2024-09-02 10:44:06.454',NULL,0,0),
    (42,3,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.455','2024-09-02 10:44:06.455',NULL,0,0),
    (43,3,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.457','2024-09-02 10:44:06.457',NULL,0,0),
    (44,3,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.458','2024-09-02 10:44:06.458',NULL,0,0),
    (45,3,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.460','2024-09-02 10:44:06.460',NULL,0,0),
    (46,3,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.461','2024-09-02 10:44:06.461',NULL,0,0),
    (47,3,'creator','','varchar(20)','string','Creator','creator','0','','0','1','','','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.463','2024-09-02 10:44:06.463',NULL,0,0),
    (48,3,'regenerator','','varchar(20)','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.464','2024-09-02 10:44:06.464',NULL,0,0),
    (49,3,'bindFlow','','bigint unsigned','string','BindFlow','bindFlow','0','','0','1','','','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:44:06.465','2024-09-02 10:44:06.465',NULL,0,0),
    (50,4,'id','','bigint unsigned','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.244','2024-09-02 10:51:16.244',NULL,0,0),
    (51,4,'title','','varchar(100)','string','Title','title','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.249','2024-09-02 10:51:16.249',NULL,0,0),
    (52,4,'node_name','','varchar(255)','string','NodeName','nodeName','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.251','2024-09-02 10:51:16.251',NULL,0,0),
    (53,4,'transfer','','varchar(255)','string','Transfer','transfer','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.252','2024-09-02 10:51:16.252',NULL,0,0),
    (54,4,'remark','','text','string','Remark','remark','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.253','2024-09-02 10:51:16.253',NULL,0,0),
    (55,4,'status','','varchar(255)','string','Status','status','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.256','2024-09-02 10:51:16.256',NULL,0,0),
    (56,4,'handlerId','','bigint','string','HandlerId','handlerId','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.257','2024-09-02 10:51:16.257',NULL,0,0),
    (57,4,'handlerName','','varchar(125)','string','HandlerName','handlerName','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.259','2024-09-02 10:51:16.259',NULL,0,0),
    (58,4,'handleTime','','datetime(3)','time.Time','HandleTime','handleTime','0','','0','1','','','','EQ','datetime','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.260','2024-09-02 10:51:16.260',NULL,0,0),
    (59,4,'handleDuration','','bigint','string','HandleDuration','handleDuration','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.262','2024-09-02 10:51:16.262',NULL,0,0),
    (60,4,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.263','2024-09-02 10:51:16.263',NULL,0,0),
    (61,4,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.264','2024-09-02 10:51:16.264',NULL,0,0),
    (62,4,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 10:51:16.265','2024-09-02 10:51:16.265',NULL,0,0),
    (63,5,'id','','bigint','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.952','2024-09-02 11:00:07.952',NULL,0,0),
    (64,5,'name','','varchar(50)','string','Name','name','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.954','2024-09-02 11:00:07.954',NULL,0,0),
    (65,5,'chineseName','','varchar(50)','string','ChineseName','chineseName','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.955','2024-09-02 11:00:07.955',NULL,0,0),
    (66,5,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.957','2024-09-02 11:00:07.957',NULL,0,0),
    (67,5,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.959','2024-09-02 11:00:07.959',NULL,0,0),
    (68,5,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.961','2024-09-02 11:00:07.961',NULL,0,0),
    (69,5,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.962','2024-09-02 11:00:07.962',NULL,0,0),
    (70,5,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.964','2024-09-02 11:00:07.964',NULL,0,0),
    (71,5,'creator','','varchar(20)','string','Creator','creator','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.965','2024-09-02 11:00:07.965',NULL,0,0),
    (72,5,'regenerator','','varchar(20)','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:00:07.968','2024-09-02 11:00:07.968',NULL,0,0),
    (73,6,'id','','bigint unsigned','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.766','2024-09-02 11:09:12.766',NULL,0,0),
    (74,6,'title','','varchar(50)','string','Title','title','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.770','2024-09-02 11:09:12.770',NULL,0,0),
    (75,6,'bindTempLate','','varchar(50)','string','BindTempLate','bindTempLate','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.774','2024-09-02 11:09:12.774',NULL,0,0),
    (76,6,'description','','varchar(256)','string','Description','description','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.776','2024-09-02 11:09:12.776',NULL,0,0),
    (77,6,'favorite','','tinyint(1)','string','Favorite','favorite','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.778','2024-09-02 11:09:12.778',NULL,0,0),
    (78,6,'icon','','varchar(50)','string','Icon','icon','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.779','2024-09-02 11:09:12.779',NULL,0,0),
    (79,6,'categoryId','','bigint unsigned','string','CategoryId','categoryId','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.781','2024-09-02 11:09:12.781',NULL,0,0),
    (80,6,'link','','varchar(256)','string','Link','link','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.782','2024-09-02 11:09:12.782',NULL,0,0),
    (81,6,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.784','2024-09-02 11:09:12.784',NULL,0,0),
    (82,6,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.785','2024-09-02 11:09:12.785',NULL,0,0),
    (83,6,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.786','2024-09-02 11:09:12.786',NULL,0,0),
    (84,6,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.788','2024-09-02 11:09:12.788',NULL,0,0),
    (85,6,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.789','2024-09-02 11:09:12.789',NULL,0,0),
    (86,6,'creator','','longtext','string','Creator','creator','0','','0','1','','','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.791','2024-09-02 11:09:12.791',NULL,0,0),
    (87,6,'regenerator','','longtext','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:09:12.792','2024-09-02 11:09:12.792',NULL,0,0),
    (88,7,'id','','bigint','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.146','2024-09-02 11:24:26.146',NULL,0,0),
    (89,7,'name','','varchar(256)','string','Name','name','0','','0','1','','0','0','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.148','2024-09-02 17:03:59.724',NULL,0,0),
    (90,7,'taskType','','varchar(45)','string','TaskType','taskType','0','','0','1','','0','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.149','2024-09-02 17:03:59.728',NULL,0,0),
    (91,7,'interpreter','','varchar(100)','string','Interpreter','interpreter','0','','0','1','','0','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.151','2024-09-02 17:03:59.729',NULL,0,0),
    (92,7,'content','','longtext','string','Content','content','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.152','2024-09-02 17:03:59.731',NULL,0,0),
    (93,7,'creator','','varchar(45)','string','Creator','creator','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.154','2024-09-02 17:03:59.733',NULL,0,0),
    (94,7,'description','','longtext','string','Description','description','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.155','2024-09-02 17:03:59.739',NULL,0,0),
    (95,7,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.160','2024-09-02 11:24:26.160',NULL,0,0),
    (96,7,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.162','2024-09-02 11:24:26.162',NULL,0,0),
    (97,7,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.163','2024-09-02 11:24:26.163',NULL,0,0),
    (98,7,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.164','2024-09-02 11:24:26.164',NULL,0,0),
    (99,7,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.166','2024-09-02 11:24:26.166',NULL,0,0),
    (100,7,'regenerator','','varchar(20)','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:24:26.167','2024-09-02 17:03:59.748',NULL,0,0),
    (101,8,'id','','bigint','int','Id','id','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.566','2024-09-02 11:26:24.566',NULL,0,0),
    (102,8,'title','','varchar(100)','string','Title','title','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.567','2024-09-02 11:26:24.567',NULL,0,0),
    (103,8,'currentNode','','varchar(50)','string','CurrentNode','currentNode','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.569','2024-09-02 11:26:24.569',NULL,0,0),
    (104,8,'currentHandler','','varchar(50)','string','CurrentHandler','currentHandler','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.571','2024-09-02 11:26:24.571',NULL,0,0),
    (105,8,'process','','varchar(50)','string','Process','process','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.572','2024-09-02 11:26:24.572',NULL,0,0),
    (106,8,'priority','','varchar(20)','string','Priority','priority','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.574','2024-09-02 11:26:24.574',NULL,0,0),
    (107,8,'status','','varchar(20)','string','Status','status','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.575','2024-09-02 11:26:24.575',NULL,0,0),
    (108,8,'department','','varchar(50)','string','Department','department','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.577','2024-09-02 11:26:24.577',NULL,0,0),
    (109,8,'description','','varchar(512)','string','Description','description','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.578','2024-09-02 11:26:24.578',NULL,0,0),
    (110,8,'form_data','','json','string','FormData','formData','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.579','2024-09-02 11:26:24.579',NULL,0,0),
    (111,8,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.580','2024-09-02 11:26:24.580',NULL,0,0),
    (112,8,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.581','2024-09-02 11:26:24.581',NULL,0,0),
    (113,8,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.583','2024-09-02 11:26:24.583',NULL,0,0),
    (114,8,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.584','2024-09-02 11:26:24.584',NULL,0,0),
    (115,8,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',15,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.585','2024-09-02 11:26:24.585',NULL,0,0),
    (116,8,'creator','','varchar(20)','string','Creator','creator','0','','0','1','','','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.585','2024-09-02 11:26:24.585',NULL,0,0),
    (117,8,'regenerator','','varchar(20)','string','Regenerator','regenerator','0','','0','1','','','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.587','2024-09-02 11:26:24.587',NULL,0,0),
    (118,8,'template','','varchar(50)','string','Template','template','0','','0','1','','','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.588','2024-09-02 11:26:24.588',NULL,0,0),
    (119,8,'currentHandlerId','','bigint','string','CurrentHandlerId','currentHandlerId','0','','0','1','','','','EQ','input','',19,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.589','2024-09-02 11:26:24.589',NULL,0,0),
    (120,8,'bindFlowData','','json','string','BindFlowData','bindFlowData','0','','0','1','','','','EQ','input','',20,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 11:26:24.589','2024-09-02 11:26:24.589',NULL,0,0),
    (121,9,'menu_id','','bigint','int','MenuId','menuId','1','','1','1','','','','EQ','input','',1,'',1,1,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.476','2024-09-02 17:04:30.476',NULL,0,0),
    (122,9,'menu_name','','varchar(128)','string','MenuName','menuName','0','','0','1','','','','EQ','input','',2,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.481','2024-09-02 17:04:30.481',NULL,0,0),
    (123,9,'title','','varchar(128)','string','Title','title','0','','0','1','','','','EQ','input','',3,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.483','2024-09-02 17:04:30.483',NULL,0,0),
    (124,9,'icon','','varchar(128)','string','Icon','icon','0','','0','1','','','','EQ','input','',4,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.488','2024-09-02 17:04:30.488',NULL,0,0),
    (125,9,'path','','varchar(128)','string','Path','path','0','','0','1','','','','EQ','input','',5,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.489','2024-09-02 17:04:30.489',NULL,0,0),
    (126,9,'paths','','varchar(128)','string','Paths','paths','0','','0','1','','','','EQ','input','',6,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.490','2024-09-02 17:04:30.490',NULL,0,0),
    (127,9,'menu_type','','varchar(1)','string','MenuType','menuType','0','','0','1','','','','EQ','input','',7,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.493','2024-09-02 17:04:30.493',NULL,0,0),
    (128,9,'action','','varchar(16)','string','Action','action','0','','0','1','','','','EQ','input','',8,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.497','2024-09-02 17:04:30.497',NULL,0,0),
    (129,9,'permission','','varchar(255)','string','Permission','permission','0','','0','1','','','','EQ','input','',9,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.499','2024-09-02 17:04:30.499',NULL,0,0),
    (130,9,'parent_id','','smallint','string','ParentId','parentId','0','','0','1','','','','EQ','input','',10,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.500','2024-09-02 17:04:30.500',NULL,0,0),
    (131,9,'no_cache','','tinyint(1)','string','NoCache','noCache','0','','0','1','','','','EQ','input','',11,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.502','2024-09-02 17:04:30.502',NULL,0,0),
    (132,9,'breadcrumb','','varchar(255)','string','Breadcrumb','breadcrumb','0','','0','1','','','','EQ','input','',12,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.504','2024-09-02 17:04:30.504',NULL,0,0),
    (133,9,'component','','varchar(255)','string','Component','component','0','','0','1','','','','EQ','input','',13,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.506','2024-09-02 17:04:30.506',NULL,0,0),
    (134,9,'sort','','tinyint','string','Sort','sort','0','','0','1','','','','EQ','input','',14,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.507','2024-09-02 17:04:30.507',NULL,0,0),
    (135,9,'visible','','varchar(1)','string','Visible','visible','0','','0','1','','','','EQ','input','',15,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.508','2024-09-02 17:04:30.508',NULL,0,0),
    (136,9,'is_frame','','varchar(1)','string','IsFrame','isFrame','0','','0','1','','','','EQ','input','',16,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.509','2024-09-02 17:04:30.509',NULL,0,0),
    (137,9,'create_by','创建者','bigint','string','CreateBy','createBy','0','','0','1','','','','EQ','input','',17,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.511','2024-09-02 17:04:30.511',NULL,0,0),
    (138,9,'update_by','更新者','bigint','string','UpdateBy','updateBy','0','','0','1','','','','EQ','input','',18,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.512','2024-09-02 17:04:30.512',NULL,0,0),
    (139,9,'created_at','创建时间','datetime(3)','time.Time','CreatedAt','createdAt','0','','0','1','','','','EQ','datetime','',19,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.512','2024-09-02 17:04:30.512',NULL,0,0),
    (140,9,'updated_at','最后更新时间','datetime(3)','time.Time','UpdatedAt','updatedAt','0','','0','1','','','','EQ','datetime','',20,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.513','2024-09-02 17:04:30.513',NULL,0,0),
    (141,9,'deleted_at','删除时间','datetime(3)','time.Time','DeletedAt','deletedAt','0','','0','1','','','','EQ','datetime','',21,'',0,0,0,0,0,1,0,0,'','','','','','','2024-09-02 17:04:30.515','2024-09-02 17:04:30.515',NULL,0,0);

/*!40000 ALTER TABLE `sys_columns` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_config
# ------------------------------------------------------------

LOCK TABLES `sys_config` WRITE;
/*!40000 ALTER TABLE `sys_config` DISABLE KEYS */;

INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,'皮肤样式','sys_index_skinName','skin-green','Y','0','主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow',1,1,'2021-05-13 19:56:37.913','2021-06-05 13:50:13.123',NULL),
    (2,'初始密码','sys_user_initPassword','123456','Y','0','用户管理-账号初始密码:123456',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
    (3,'侧栏主题','sys_index_sideTheme','theme-light','Y','0','主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
    (4,'系统名称','sys_app_name','smart 工单系统','Y','1','',1,0,'2021-03-17 08:52:06.067','2021-05-28 10:08:25.248',NULL),
    (5,'系统logo','sys_app_logo','https://github.com/sunwenbo/golang/raw/master/logo.png\n','Y','1','',1,0,'2021-03-17 08:53:19.462','2021-03-17 08:53:19.462',NULL);

/*!40000 ALTER TABLE `sys_config` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_dept
# ------------------------------------------------------------

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;

INSERT INTO `sys_dept` (`dept_id`, `parent_id`, `dept_path`, `dept_name`, `sort`, `leader`, `phone`, `email`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,0,'/0/1/','数据中台',0,'张三','13788888888','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.000','2024-10-14 14:48:18.422',NULL),
    (7,1,'/0/1/7/','研发部',1,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2024-09-23 20:43:52.181',NULL),
    (8,1,'/0/1/8/','运维部',0,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2024-09-23 20:43:52.181',NULL),
    (9,1,'/0/1/9/','客服部',0,'aituo','13782218188','atuo@aituo.com',2,1,1,'2021-05-13 19:56:37.913','2024-09-23 20:43:52.181',NULL),
    (10,1,'/0/1/10/','人力资源',3,'aituo','13782218188','atuo@aituo.com',1,1,1,'2021-05-13 19:56:37.913','2024-09-23 20:43:52.181',NULL),
    (11,0,'/0/11/','解决方案团队',1,'xxx','','',2,0,0,'2024-07-28 23:29:56.623','2024-10-14 14:47:21.786',NULL);

/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_dict_data
# ------------------------------------------------------------

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;

INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `status`, `default`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,0,'正常','2','sys_normal_disable','','','',2,'','系统正常',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:40.168',NULL),
    (2,0,'停用','1','sys_normal_disable','','','',2,'','系统停用',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (3,0,'男','0','sys_user_sex','','','',2,'','性别男',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (4,0,'女','1','sys_user_sex','','','',2,'','性别女',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (5,0,'未知','2','sys_user_sex','','','',2,'','性别未知',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (6,0,'显示','0','sys_show_hide','','','',2,'','显示菜单',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (7,0,'隐藏','1','sys_show_hide','','','',2,'','隐藏菜单',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (8,0,'是','Y','sys_yes_no','','','',2,'','系统默认是',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (9,0,'否','N','sys_yes_no','','','',2,'','系统默认否',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (10,0,'正常','2','sys_job_status','','','',2,'','正常状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (11,0,'停用','1','sys_job_status','','','',2,'','停用状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (12,0,'默认','DEFAULT','sys_job_group','','','',2,'','默认分组',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (13,0,'系统','SYSTEM','sys_job_group','','','',2,'','系统分组',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (14,0,'通知','1','sys_notice_type','','','',2,'','通知',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (15,0,'公告','2','sys_notice_type','','','',2,'','公告',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (16,0,'正常','2','sys_common_status','','','',2,'','正常状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (17,0,'关闭','1','sys_common_status','','','',2,'','关闭状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (18,0,'新增','1','sys_oper_type','','','',2,'','新增操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (19,0,'修改','2','sys_oper_type','','','',2,'','修改操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (20,0,'删除','3','sys_oper_type','','','',2,'','删除操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (21,0,'授权','4','sys_oper_type','','','',2,'','授权操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (22,0,'导出','5','sys_oper_type','','','',2,'','导出操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (23,0,'导入','6','sys_oper_type','','','',2,'','导入操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (24,0,'强退','7','sys_oper_type','','','',2,'','强退操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (25,0,'生成代码','8','sys_oper_type','','','',2,'','生成操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (26,0,'清空数据','9','sys_oper_type','','','',2,'','清空操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (27,0,'成功','0','sys_notice_status','','','',2,'','成功状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (28,0,'失败','1','sys_notice_status','','','',2,'','失败状态',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (29,0,'登录','10','sys_oper_type','','','',2,'','登录操作',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (30,0,'退出','11','sys_oper_type','','','',2,'','',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (31,0,'获取验证码','12','sys_oper_type','','','',2,'','获取验证码',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (32,0,'正常','1','sys_content_status','','','',1,'','',1,1,'2021-05-13 19:56:40.845','2021-05-13 19:56:40.845',NULL),
    (33,1,'禁用','2','sys_content_status','','','',1,'','',1,1,'2021-05-13 19:56:40.845','2021-05-13 19:56:40.845',NULL),
    (34,0,'进行中','under-way','order_works_status','',NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (35,0,'已结束','termination','order_works_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (36,0,'手动结束','manual-termination','order_works_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (37,0,'重开','reopen','order_works_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (38,0,'驳回','reject','order_works_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (39,0,'一般','normal','order_works_priority_type',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (40,0,'紧急','urgent','order_works_priority_type',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (41,0,'非常紧急','very-urgent','order_works_priority_type',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (42,0,'normal','success','order_works_priority_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (43,0,'urgent','warning','order_works_priority_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (44,0,'very-urgent','danger','order_works_priority_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (45,0,'同意','1','order_works_flow_result',NULL,NULL,NULL,2,NULL,'流程审批同意',NULL,NULL,NULL,NULL,NULL),
    (46,0,'拒绝','0','order_works_flow_result',NULL,NULL,NULL,2,NULL,'流程审批拒绝',NULL,NULL,NULL,NULL,NULL),
    (47,0,'其他','2','order_works_flow_result',NULL,NULL,NULL,2,NULL,'流程审批其他',NULL,NULL,NULL,NULL,NULL),
    (48,0,'Shell','0','order_works_task_type',NULL,NULL,NULL,2,NULL,'shell',NULL,NULL,NULL,NULL,NULL),
    (49,0,'Python','1','order_works_task_type',NULL,NULL,NULL,2,NULL,'python',NULL,NULL,NULL,NULL,NULL),
    (50,0,'Javascript','2','order_works_task_type',NULL,NULL,NULL,2,NULL,'javascript',NULL,NULL,NULL,NULL,NULL),
    (51,0,'在线','1','order_exec_machine_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (52,0,'离线','2','order_exec_machine_status',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (53,0,'密码','1','order_exec_machine_auth',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (54,0,'私钥','2','order_exec_machine_auth',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (55,0,'成功','0','task_history_type',NULL,NULL,NULL,2,NULL,'任务执行成功',NULL,NULL,NULL,NULL,NULL),
    (56,0,'失败','1','task_history_type',NULL,NULL,NULL,2,NULL,'任务执行失败',NULL,NULL,NULL,NULL,NULL),
    (57,0,'未知','2','task_history_type',NULL,NULL,NULL,2,NULL,'未知状态',NULL,NULL,NULL,NULL,NULL),
    (58,0,'邮件','1','order_message_notice',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (59,0,'飞书','2','order_message_notice',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
    (60,0,'电话','3','order_message_notice',NULL,NULL,NULL,2,NULL,NULL,NULL,NULL,NULL,NULL,NULL);

/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_dict_type
# ------------------------------------------------------------

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type` DISABLE KEYS */;

INSERT INTO `sys_dict_type` (`dict_id`, `dict_name`, `dict_type`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,'系统开关','sys_normal_disable',2,'系统开关列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (2,'用户性别','sys_user_sex',2,'用户性别列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (3,'菜单状态','sys_show_hide',2,'菜单状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (4,'系统是否','sys_yes_no',2,'系统是否列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (5,'任务状态','sys_job_status',2,'任务状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (6,'任务分组','sys_job_group',2,'任务分组列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (7,'通知类型','sys_notice_type',2,'通知类型列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (8,'系统状态','sys_common_status',2,'登录状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (9,'操作类型','sys_oper_type',2,'操作类型列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (10,'通知状态','sys_notice_status',2,'通知状态列表',1,1,'2021-05-13 19:56:37.914','2021-05-13 19:56:37.914',NULL),
    (11,'内容状态','sys_content_status',2,'',1,1,'2021-05-13 19:56:40.813','2021-05-13 19:56:40.813',NULL);

/*!40000 ALTER TABLE `sys_dict_type` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_job
# ------------------------------------------------------------

LOCK TABLES `sys_job` WRITE;
/*!40000 ALTER TABLE `sys_job` DISABLE KEYS */;

INSERT INTO `sys_job` (`job_id`, `job_name`, `job_group`, `job_type`, `cron_expression`, `invoke_target`, `args`, `misfire_policy`, `concurrent`, `status`, `entry_id`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
    (1,'接口测试','DEFAULT',1,'0/5 * * * * ','http://localhost:8000','',1,1,1,0,'2021-05-13 19:56:37.914','2024-09-03 16:38:39.679',NULL,1,1),
    (2,'函数测试','DEFAULT',2,'0/5 * * * * ','ExamplesOne','参数',1,1,1,0,'2021-05-13 19:56:37.914','2021-05-31 23:55:37.221',NULL,1,1),
    (3,'节点存活检测','DEFAULT',2,'0 */15 * * * *','CheckAllMachines',NULL,1,1,2,1,'2021-05-13 19:56:37.914','2024-10-27 15:50:49.417',NULL,1,1);

/*!40000 ALTER TABLE `sys_job` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_login_log
# ------------------------------------------------------------



# 转储表 sys_menu
# ------------------------------------------------------------

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`menu_id`, `menu_name`, `title`, `icon`, `path`, `paths`, `menu_type`, `action`, `permission`, `parent_id`, `no_cache`, `breadcrumb`, `component`, `sort`, `visible`, `is_frame`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (2,'Admin','系统管理','api-server','/admin','/0/2','M','无','',0,1,'','Layout',10,'0','1',0,1,'2021-05-20 21:58:45.679','2024-10-23 14:43:43.971',NULL),
    (3,'SysUserManage','用户管理','user','/admin/sys-user','/0/2/3','C','无','admin:sysUser:list',2,0,'','/admin/sys-user/index',10,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.975',NULL),
    (43,'','新增管理员','app-group-fill','','/0/2/3/43','F','POST','admin:sysUser:add',3,0,'','',10,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.983',NULL),
    (44,'','查询管理员','app-group-fill','','/0/2/3/44','F','GET','admin:sysUser:query',3,0,'','',40,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.984',NULL),
    (45,'','修改管理员','app-group-fill','','/0/2/3/45','F','PUT','admin:sysUser:edit',3,0,'','',30,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.987',NULL),
    (46,'','删除管理员','app-group-fill','','/0/2/3/46','F','DELETE','admin:sysUser:remove',3,0,'','',20,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.988',NULL),
    (51,'SysMenuManage','菜单管理','tree-table','/admin/sys-menu','/0/2/51','C','无','admin:sysMenu:list',2,1,'','/admin/sys-menu/index',30,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.975',NULL),
    (52,'SysRoleManage','角色管理','peoples','/admin/sys-role','/0/2/52','C','无','admin:sysRole:list',2,1,'','/admin/sys-role/index',20,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.976',NULL),
    (56,'SysDeptManage','部门管理','tree','/admin/sys-dept','/0/2/56','C','无','admin:sysDept:list',2,0,'','/admin/sys-dept/index',40,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.977',NULL),
    (57,'SysPostManage','岗位管理','pass','/admin/sys-post','/0/2/57','C','无','admin:sysPost:list',2,0,'','/admin/sys-post/index',50,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.978',NULL),
    (58,'Dict','字典管理','education','/admin/dict','/0/2/58','C','无','admin:sysDictType:list',2,0,'','/admin/dict/index',60,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.978',NULL),
    (59,'SysDictDataManage','字典数据','education','/admin/dict/data/:dictId','/0/2/59','C','无','admin:sysDictData:list',2,0,'','/admin/dict/data',100,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.979',NULL),
    (60,'Tools','开发工具','dev-tools','/dev-tools','/0/60','M','无','',0,0,'','Layout',40,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.972',NULL),
    (61,'Swagger','系统接口','guide','/dev-tools/swagger','/0/60/61','C','无','',60,0,'','/dev-tools/swagger/index',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.001',NULL),
    (62,'SysConfigManage','参数管理','swagger','/admin/sys-config','/0/2/62','C','无','admin:sysConfig:list',2,0,'','/admin/sys-config/index',70,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.980',NULL),
    (211,'Log','日志管理','log','/log','/0/2/211','M','','',2,0,'','/log/index',80,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.981',NULL),
    (212,'SysLoginLogManage','登录日志','logininfor','/admin/sys-login-log','/0/2/211/212','C','','admin:sysLoginLog:list',211,0,'','/admin/sys-login-log/index',1,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:44.005',NULL),
    (216,'OperLog','操作日志','skill','/admin/sys-oper-log','/0/2/211/216','C','','admin:sysOperLog:list',211,0,'','/admin/sys-oper-log/index',1,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:44.005',NULL),
    (220,'','新增菜单','app-group-fill','','/0/2/51/220','F','','admin:sysMenu:add',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.988',NULL),
    (221,'','修改菜单','app-group-fill','','/0/2/51/221','F','','admin:sysMenu:edit',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.989',NULL),
    (222,'','查询菜单','app-group-fill','','/0/2/51/222','F','','admin:sysMenu:query',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.990',NULL),
    (223,'','删除菜单','app-group-fill','','/0/2/51/223','F','','admin:sysMenu:remove',51,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.991',NULL),
    (224,'','新增角色','app-group-fill','','/0/2/52/224','F','','admin:sysRole:add',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.991',NULL),
    (225,'','查询角色','app-group-fill','','/0/2/52/225','F','','admin:sysRole:query',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.992',NULL),
    (226,'','修改角色','app-group-fill','','/0/2/52/226','F','','admin:sysRole:update',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.993',NULL),
    (227,'','删除角色','app-group-fill','','/0/2/52/227','F','','admin:sysRole:remove',52,0,'','',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.994',NULL),
    (228,'','查询部门','app-group-fill','','/0/2/56/228','F','','admin:sysDept:query',56,0,'','',40,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.994',NULL),
    (229,'','新增部门','app-group-fill','','/0/2/56/229','F','','admin:sysDept:add',56,0,'','',10,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.995',NULL),
    (230,'','修改部门','app-group-fill','','/0/2/56/230','F','','admin:sysDept:edit',56,0,'','',30,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.995',NULL),
    (231,'','删除部门','app-group-fill','','/0/2/56/231','F','','admin:sysDept:remove',56,0,'','',20,'0','1',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.995',NULL),
    (232,'','查询岗位','app-group-fill','','/0/2/57/232','F','','admin:sysPost:query',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.996',NULL),
    (233,'','新增岗位','app-group-fill','','/0/2/57/233','F','','admin:sysPost:add',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.996',NULL),
    (234,'','修改岗位','app-group-fill','','/0/2/57/234','F','','admin:sysPost:edit',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.997',NULL),
    (235,'','删除岗位','app-group-fill','','/0/2/57/235','F','','admin:sysPost:remove',57,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.997',NULL),
    (236,'','查询字典','app-group-fill','','/0/2/58/236','F','','admin:sysDictType:query',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.998',NULL),
    (237,'','新增类型','app-group-fill','','/0/2/58/237','F','','admin:sysDictType:add',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.998',NULL),
    (238,'','修改类型','app-group-fill','','/0/2/58/238','F','','admin:sysDictType:edit',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.998',NULL),
    (239,'','删除类型','app-group-fill','','/0/2/58/239','F','','system:sysdicttype:remove',58,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.999',NULL),
    (240,'','查询数据','app-group-fill','','/0/2/59/240','F','','admin:sysDictData:query',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:43.999',NULL),
    (241,'','新增数据','app-group-fill','','/0/2/59/241','F','','admin:sysDictData:add',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.000',NULL),
    (242,'','修改数据','app-group-fill','','/0/2/59/242','F','','admin:sysDictData:edit',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.000',NULL),
    (243,'','删除数据','app-group-fill','','/0/2/59/243','F','','admin:sysDictData:remove',59,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.000',NULL),
    (244,'','查询参数','app-group-fill','','/0/2/62/244','F','','admin:sysConfig:query',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.003',NULL),
    (245,'','新增参数','app-group-fill','','/0/2/62/245','F','','admin:sysConfig:add',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.003',NULL),
    (246,'','修改参数','app-group-fill','','/0/2/62/246','F','','admin:sysConfig:edit',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.004',NULL),
    (247,'','删除参数','app-group-fill','','/0/2/62/247','F','','admin:sysConfig:remove',62,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.004',NULL),
    (248,'','查询登录日志','app-group-fill','','/0/2/211/212/248','F','','admin:sysLoginLog:query',212,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.006',NULL),
    (249,'','删除登录日志','app-group-fill','','/0/2/211/212/249','F','','admin:sysLoginLog:remove',212,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.006',NULL),
    (250,'','查询操作日志','app-group-fill','','/0/2/211/216/250','F','','admin:sysOperLog:query',216,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.007',NULL),
    (251,'','删除操作日志','app-group-fill','','/0/2/211/216/251','F','','admin:sysOperLog:remove',216,0,'','',0,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.007',NULL),
    (261,'Gen','代码生成','code','/dev-tools/gen','/0/60/261','C','','',60,0,'','/dev-tools/gen/index',2,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.001',NULL),
    (262,'EditTable','代码生成修改','build','/dev-tools/editTable','/0/60/262','C','','',60,0,'','/dev-tools/gen/editTable',100,'1','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.002',NULL),
    (264,'Build','表单构建','build','/dev-tools/build','/0/60/264','C','','',60,0,'','/dev-tools/build/index',1,'0','1',1,1,'2020-04-11 15:52:48.000','2024-10-23 14:43:44.002',NULL),
    (269,'ServerMonitor','服务监控','druid','/sys-tools/monitor','/0/537/269','C','','sysTools:serverMonitor:list',537,0,'','/sys-tools/monitor',0,'0','1',1,1,'2020-04-14 00:28:19.000','2024-10-23 14:43:44.011',NULL),
    (459,'Schedule','定时任务','time-range','/schedule','/0/459','M','无','',0,0,'','Layout',20,'0','1',1,1,'2020-08-03 09:17:37.000','2024-10-23 14:43:43.973',NULL),
    (460,'ScheduleManage','定时任务','job','/schedule/manage','/0/459/460','C','无','job:sysJob:list',459,0,'','/schedule/index',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-10-23 14:43:44.008',NULL),
    (461,'sys_job','分页获取定时任务','app-group-fill','','/0/459/460/461','F','无','job:sysJob:query',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-10-23 14:43:44.008',NULL),
    (462,'sys_job','创建定时任务','app-group-fill','','/0/459/460/462','F','无','job:sysJob:add',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-10-23 14:43:44.009',NULL),
    (463,'sys_job','修改定时任务','app-group-fill','','/0/459/460/463','F','无','job:sysJob:edit',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-10-23 14:43:44.010',NULL),
    (464,'sys_job','删除定时任务','app-group-fill','','/0/459/460/464','F','无','job:sysJob:remove',460,0,'','',0,'0','1',1,1,'2020-08-03 09:17:37.000','2024-10-23 14:43:44.010',NULL),
    (471,'JobLog','日志','bug','/schedule/log','/0/459/471','C','','',459,0,'','/schedule/log',0,'1','1',1,1,'2020-08-05 21:24:46.000','2024-10-23 14:43:44.008',NULL),
    (528,'SysApiManage','接口管理','api-doc','/admin/sys-api','/0/2/528','C','无','admin:sysApi:list',2,0,'','/admin/sys-api/index',0,'0','0',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:43.981',NULL),
    (529,'','查询接口','app-group-fill','','/0/2/528/529','F','无','admin:sysApi:query',528,0,'','',40,'0','0',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:44.010',NULL),
    (531,'','修改接口','app-group-fill','','/0/2/528/531','F','无','admin:sysApi:edit',528,0,'','',30,'0','0',0,1,'2021-05-20 22:08:44.526','2024-10-23 14:43:44.011',NULL),
    (537,'SysTools','系统工具','system-tools','/sys-tools','/0/537','M','','',0,0,'','Layout',30,'0','1',1,1,'2021-05-21 11:13:32.166','2024-10-23 14:43:43.973',NULL),
    (540,'SysConfigSet','参数设置','system-tools','/admin/sys-config/set','/0/2/540','C','','admin:sysConfigSet:list',2,0,'','/admin/sys-config/set',0,'0','1',1,1,'2021-05-25 16:06:52.560','2024-10-23 14:43:43.982',NULL),
    (542,'','修改','upload','','/0/2/540/542','F','','admin:sysConfigSet:update',540,0,'','',0,'0','1',1,1,'2021-06-13 11:45:48.670','2024-10-23 14:43:44.011',NULL),
    (543,'OrderCenter','工单中心','education','/orderCenter','/0/543','M','','',0,0,'','Layout',1,'0','1',1,1,'2024-07-09 11:42:28.535','2024-10-23 14:43:43.973',NULL),
    (544,'OrderApply','工单申请','json-push','/orderCenter/apply','/0/543/544','C','','',543,0,'','/orderCenter/index',0,'0','0',1,1,'2024-07-09 11:50:54.940','2024-10-23 14:43:44.012',NULL),
    (545,'OrderList','工单列表','base-info','/orderCenter/list','/0/543/545','C','','',543,0,'','/orderCenter/list/workOrderList',1,'0','1',1,1,'2024-07-09 11:56:50.887','2024-10-23 14:43:44.012',NULL),
    (546,'WorkOrderDetails','工单详情','','/orderCenter/workflow/work-order/details/:id/:title','/0/543/546','C','','general:orderWorks:details',543,0,'','/orderCenter/list/workOrderDetails',3,'1','1',1,1,'2024-07-09 11:58:32.874','2024-10-23 14:43:44.012',NULL),
    (547,'FormRender','表单提交','clipboard','/order/classify/:bindTempLate','/0/543/547','C','','',543,0,'','/orderCenter/list/formRender',4,'1','1',1,1,'2024-07-09 12:00:13.000','2024-10-23 14:43:44.013',NULL),
    (548,'flowCenter','流程中心','app-group-fill','/flowCenter','/0/548','M','','',0,0,'','Layout',2,'0','1',1,1,'2024-07-09 12:03:07.870','2024-10-23 14:43:43.974',NULL),
    (549,'OrderCateGory','工单类型','log','/flowCenter/order-category','/0/548/549','C','','',548,0,'','/flowCenter/orderCateGory',0,'0','1',1,1,'2024-07-09 12:04:55.000','2024-10-23 14:43:44.019',NULL),
    (550,'TemplateManage','模板管理','dict','/flowCenter/template-manage','/0/548/550','C','','',548,0,'','/flowCenter/templateManage',2,'0','1',1,1,'2024-07-09 12:06:10.000','2024-10-23 14:43:44.020',NULL),
    (551,'OrderManage','工单管理','row','/flowCenter/order-manage','/0/548/551','C','','',548,0,'','/flowCenter/orderManage',3,'0','1',1,1,'2024-07-09 12:19:39.000','2024-10-23 14:43:44.020',NULL),
    (552,'FlowManage','流程管理','operators','/flowCenter/flow-manage','/0/548/552','C','','',548,0,'','/flowCenter/flowManage',1,'0','1',1,1,'2024-07-09 12:21:22.000','2024-10-23 14:43:44.020',NULL),
    (553,'TaskCenter','任务中心','example','/task','/0/553','M','','',0,0,'','Layout',3,'0','1',1,1,'2024-07-09 12:24:37.821','2024-10-23 14:43:43.974',NULL),
    (554,'TaskList','任务列表','list','/taskCenter/list','/0/553/554','C','','general:orderTask:list',553,0,'','/taskCenter/index',0,'0','1',1,1,'2024-07-09 12:26:49.198','2024-10-23 14:43:44.029',NULL),
    (555,'TaskHistory','历史任务','project-manage','/taskCenter/history','/0/553/555','C','','',553,0,'','/taskCenter/history',1,'0','1',1,1,'2024-07-09 12:28:39.180','2024-10-23 14:43:44.029',NULL),
    (556,'TaskExecute','执行节点','access','/taskCenter/execMachine','/0/553/556','C','','',553,0,'','/taskCenter/execMachine',2,'0','1',1,1,'2024-07-09 12:29:44.818','2024-10-23 14:43:44.029',NULL),
    (569,'','提交工单','app-group-fill','','/0/543/547/569','F','','general:orderApply:add',547,0,'','',1,'0','1',1,1,'2024-08-13 14:55:55.057','2024-10-23 14:43:44.019',NULL),
    (570,'','创建工单','app-group-fill','','/0/543/545/570','F','','general:orderWorks:create',545,0,'','',1,'0','1',1,0,'2024-08-13 15:55:45.791','2024-08-13 19:06:05.570','2024-08-13 19:42:41.357'),
    (571,'','我的待办','app-group-fill','','/0/543/545/571','F','','general:orderWorks:myBacklog',545,0,'','',10,'0','1',1,1,'2024-08-13 15:57:01.071','2024-10-23 14:43:44.014',NULL),
    (572,'','我创建的','app-group-fill','','/0/543/545/572','F','','general:orderWorks:createdByMe',545,0,'','',20,'0','1',1,1,'2024-08-13 15:57:46.963','2024-10-23 14:43:44.015',NULL),
    (573,'','我相关的','app-group-fill','','/0/543/545/573','F','','general:orderWorks:relatedToMe',545,0,'','',30,'0','1',1,1,'2024-08-13 15:58:41.259','2024-10-23 14:43:44.016',NULL),
    (574,'','工单处理','app-group-fill','','/0/543/546/574','F','','general:orderWorks:handler',546,0,'','',10,'0','1',1,1,'2024-08-13 16:01:47.417','2024-10-23 14:43:44.018',NULL),
    (575,'','工单消息通知','app-group-fill','','/0/543/545/575','F','','general:orderWorks:notifyGet',545,0,'','',40,'0','1',1,1,'2024-08-13 17:28:11.566','2024-08-14 19:46:17.017','2024-08-14 19:46:39.569'),
    (576,'','获取工单类型','app-group-fill','','/0/543/544/576','F','','general:orderCategory:list',544,0,'','',0,'0','1',1,1,'2024-08-13 19:37:16.887','2024-10-23 14:43:44.013',NULL),
    (577,'','获取工单项','app-group-fill','','/0/543/544/577','F','','general:orderItems:list',544,0,'','',0,'0','1',1,1,'2024-08-13 19:38:46.741','2024-10-23 14:43:44.014',NULL),
    (578,'','获取指定的工单数据','app-group-fill','','/0/543/545/578','F','','general:orderWorks:listid',545,0,'','',2,'0','1',1,1,'2024-08-13 20:02:37.576','2024-10-23 14:43:44.016',NULL),
    (579,'','工单更新','app-group-fill','','/0/543/545/579','F','','general:orderWorks:update',545,0,'','',3,'0','1',1,1,'2024-08-13 20:07:11.853','2024-10-23 14:43:44.016',NULL),
    (580,'','类型列表','app-group-fill','','/0/548/549/580','F','','general:orderCategory:list',549,0,'','',10,'0','1',1,1,'2024-08-13 20:12:45.390','2024-10-23 14:43:44.021',NULL),
    (581,'','新建工单分类','app-group-fill','','/0/548/549/581','F','','general:orderCategory:create',549,0,'','',20,'0','1',1,1,'2024-08-13 20:14:57.885','2024-10-23 14:43:44.021',NULL),
    (582,'','更新工单分类','app-group-fill','','/0/548/549/582','F','','general:orderCategory:update',549,0,'','',50,'0','1',1,1,'2024-08-13 20:15:18.734','2024-10-23 14:43:44.021',NULL),
    (583,'','删除工单分类','app-group-fill','','/0/548/549/583','F','','general:orderCategory:delete',549,0,'','',40,'0','1',1,1,'2024-08-13 20:15:44.973','2024-10-23 14:43:44.022',NULL),
    (584,'','根据id查询工单分类','app-group-fill','','/0/548/549/584','F','','general:orderCategory:listId',549,0,'','',30,'0','1',1,1,'2024-08-13 20:16:18.138','2024-10-23 14:43:44.022',NULL),
    (585,'','流程列表','app-group-fill','','/0/548/552/585','F','','general:orderFlowManage:list',552,0,'','',10,'0','1',1,1,'2024-08-13 20:34:36.752','2024-10-23 14:43:44.025',NULL),
    (586,'','新建流程','app-group-fill','','/0/548/552/586','F','','general:orderFlowManage:create',552,0,'','',20,'0','1',1,1,'2024-08-13 20:35:05.574','2024-10-23 14:43:44.025',NULL),
    (587,'','更新流程','app-group-fill','','/0/548/552/587','F','','general:orderFlowManage:update',552,0,'','',30,'0','1',1,1,'2024-08-13 20:35:32.835','2024-10-23 14:43:44.026',NULL),
    (588,'','删除流程','app-group-fill','','/0/548/552/588','F','','general:orderFlowManage:delete',552,0,'','',40,'0','1',1,1,'2024-08-13 20:36:00.732','2024-10-23 14:43:44.026',NULL),
    (589,'','克隆流程','app-group-fill','','/0/548/552/589','F','','general:orderFlowManage:clone',552,0,'','',50,'0','1',1,1,'2024-08-13 20:36:16.508','2024-10-23 14:43:44.027',NULL),
    (590,'','根据id查询流程','','','/0/548/552/590','F','','general:orderFlowManage:listId',552,0,'','',70,'0','1',1,1,'2024-08-13 20:36:48.647','2024-10-23 14:43:44.027',NULL),
    (591,'','模板列表','app-group-fill','','/0/548/550/591','F','','general:orderTemplate:list',550,0,'','',10,'0','1',1,1,'2024-08-13 20:37:38.658','2024-10-23 14:43:44.023',NULL),
    (592,'','创建模板','app-group-fill','','/0/548/550/592','F','','general:orderTemplate:create',550,0,'','',20,'0','1',1,1,'2024-08-13 20:38:16.755','2024-10-23 14:43:44.023',NULL),
    (593,'','更新模板','app-group-fill','','/0/548/550/593','F','','general:orderTemplate:update',550,0,'','',30,'0','1',1,1,'2024-08-13 20:38:42.278','2024-10-23 14:43:44.023',NULL),
    (594,'','删除模板','app-group-fill','','/0/548/550/594','F','','general:orderTemplate:delete',550,0,'','',50,'0','1',1,1,'2024-08-13 20:39:01.732','2024-10-23 14:43:44.023',NULL),
    (595,'','根据id查询模板','app-group-fill','','/0/548/550/595','F','','general:orderTemplate:listId',550,0,'','',50,'0','1',1,1,'2024-08-13 20:39:31.285','2024-10-23 14:43:44.024',NULL),
    (596,'','工单列表','app-group-fill','','/0/548/551/596','F','','general:orderItems:list',551,0,'','',10,'0','1',1,1,'2024-08-13 20:41:17.251','2024-10-23 14:43:44.024',NULL),
    (597,'','创建工单项','app-group-fill','','/0/548/551/597','F','','general:orderItems:create',551,0,'','',20,'0','1',1,1,'2024-08-13 20:42:31.063','2024-10-23 14:43:44.024',NULL),
    (598,'','更新工单项','app-group-fill','','/0/548/551/598','F','','general:orderItems:update',551,0,'','',30,'0','1',1,0,'2024-08-13 20:43:12.034','2024-08-13 21:46:42.897','2024-08-14 11:13:27.906'),
    (599,'','删除工单项','app-group-fill','','/0/548/551/599','F','','general:orderItems:delete',551,0,'','',40,'0','1',1,0,'2024-08-13 20:43:57.384','2024-08-13 21:46:42.898','2024-08-14 11:13:17.311'),
    (600,'','根据id查询工单项','app-group-fill','','/0/548/550/600','F','','general:orderItems:listId',550,0,'','',50,'0','1',1,1,'2024-08-13 20:44:48.263','2024-10-23 14:43:44.024',NULL),
    (601,'','获取模板','app-group-fill','','/0/548/551/601','F','','general:orderTemplate:list',551,0,'','',50,'0','1',1,1,'2024-08-13 21:04:51.547','2024-10-23 14:43:44.025',NULL),
    (602,'','工单消息通知','app-group-fill','','/0/543/602','F','','general:orderNotify:list',543,0,'','',5,'0','1',1,1,'2024-08-14 16:27:19.613','2024-10-23 14:43:44.013',NULL),
    (603,'','获取工单列表','app-group-fill','','/0/543/545/603','F','','general:orderWorks:list',545,0,'','',1,'0','1',1,1,'2024-08-14 16:53:29.195','2024-10-23 14:43:44.017',NULL),
    (604,'','获取工单历史记录','api-monitor','','/0/543/546/604','F','','general:orderWorks:historyGet',546,0,'','',20,'0','1',1,1,'2024-08-14 16:56:28.348','2024-10-23 14:43:44.018',NULL),
    (605,'','部门列表','app-group-fill','','/0/548/552/605','F','','general:orderDept:list',552,0,'','',80,'0','1',1,0,'2024-08-14 19:35:10.803','2024-10-23 14:43:44.027',NULL),
    (606,'','用户列表','app-group-fill','','/0/548/552/606','F','','general:orderUser:list',552,0,'','',90,'0','1',1,0,'2024-08-14 19:36:15.750','2024-10-23 14:43:44.028',NULL),
    (607,'','权限列表','app-group-fill','','/0/548/552/607','F','','general:orderRole:list',552,0,'','',60,'0','1',1,0,'2024-08-14 19:37:31.487','2024-10-23 14:43:44.028',NULL),
    (608,'','查看工单','app-group-fill','','/0/543/545/608','F','','general:orderWorks:view',545,0,'','',0,'0','1',1,1,'2024-08-14 19:56:37.397','2024-10-23 14:43:44.017',NULL),
    (609,'','获取用户列表','app-group-fill','','/0/543/545/609','F','','general:orderWorks:getUsers',545,0,'','',40,'0','1',1,1,'2024-08-17 17:08:28.627','2024-10-23 14:43:44.017',NULL),
    (610,'','创建任务','app-group-fill','','/0/553/554/610','F','','general:orderTask:create',554,0,'','',1,'0','1',1,1,'2024-08-17 20:51:15.230','2024-10-23 14:43:44.030',NULL),
    (611,'','更新任务','app-group-fill','','/0/553/554/611','F','','general:orderTask:update',554,0,'','',2,'0','1',1,1,'2024-08-17 20:52:31.690','2024-10-23 14:43:44.030',NULL),
    (612,'','根据id查询任务','app-group-fill','','/0/553/612','F','','general:orderTask:listId',553,0,'','',3,'0','1',1,0,'2024-08-17 20:53:02.124','2024-08-17 20:53:02.131','2024-08-17 20:53:44.397'),
    (613,'','根据id查询任务','app-group-fill','','/0/553/554/613','F','','general:orderTask:listId',554,0,'','',3,'0','1',1,1,'2024-08-17 20:53:41.157','2024-10-23 14:43:44.030',NULL),
    (614,'','删除任务','app-group-fill','','/0/553/554/614','F','','general:orderTask:delete',554,0,'','',4,'0','1',1,1,'2024-08-17 20:54:13.565','2024-10-23 14:43:44.031',NULL),
    (615,'','ExecMachine','pass','/exec-machine','/0/615','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-09-01 23:33:35.487','2024-09-01 23:33:35.492','2024-09-02 12:12:37.244'),
    (616,'ExecMachineManage','ExecMachine','pass','/admin/exec-machine','/0/615/616','C','无','admin:execMachine:list',615,0,'','/admin/exec-machine/index',0,'0','0',1,1,'2024-09-01 23:33:35.494','2024-09-01 23:33:35.496',NULL),
    (617,'','分页获取ExecMachine','','exec_machine','/0/615/616/617','F','无','admin:execMachine:query',616,0,'','',0,'0','0',1,1,'2024-09-01 23:33:35.498','2024-10-23 14:43:44.031',NULL),
    (618,'','创建ExecMachine','','exec_machine','/0/615/616/618','F','无','admin:execMachine:add',616,0,'','',0,'0','0',1,1,'2024-09-01 23:33:35.501','2024-10-23 14:43:44.032',NULL),
    (619,'','修改ExecMachine','','exec_machine','/0/615/616/619','F','无','admin:execMachine:edit',616,0,'','',0,'0','0',1,1,'2024-09-01 23:33:35.503','2024-10-23 14:43:44.032',NULL),
    (620,'','删除ExecMachine','','exec_machine','/0/615/616/620','F','无','admin:execMachine:remove',616,0,'','',0,'0','0',1,1,'2024-09-01 23:33:35.506','2024-10-23 14:43:44.033',NULL),
    (621,'','ExecMachine','pass','/exec-machine','/0/621','M','无','',0,0,'','Layout',0,'0','0',1,0,'2024-09-01 23:34:34.171','2024-09-01 23:34:34.172','2024-09-02 12:12:39.426'),
    (622,'ExecMachineManage','ExecMachine','pass','/admin/exec-machine','/0/621/622','C','无','admin:execMachine:list',621,0,'','/admin/exec-machine/index',0,'0','0',1,1,'2024-09-01 23:34:34.175','2024-09-01 23:34:34.177',NULL),
    (623,'','分页获取ExecMachine','','exec_machine','/0/621/622/623','F','无','admin:execMachine:query',622,0,'','',0,'0','0',1,1,'2024-09-01 23:34:34.180','2024-10-23 14:43:44.033',NULL),
    (624,'','创建ExecMachine','','exec_machine','/0/621/622/624','F','无','admin:execMachine:add',622,0,'','',0,'0','0',1,1,'2024-09-01 23:34:34.185','2024-10-23 14:43:44.033',NULL),
    (625,'','修改ExecMachine','','exec_machine','/0/621/622/625','F','无','admin:execMachine:edit',622,0,'','',0,'0','0',1,1,'2024-09-01 23:34:34.190','2024-10-23 14:43:44.034',NULL),
    (626,'','删除ExecMachine','','exec_machine','/0/621/622/626','F','无','admin:execMachine:remove',622,0,'','',0,'0','0',1,1,'2024-09-01 23:34:34.195','2024-10-23 14:43:44.034',NULL),
    (627,'','工单评价','app-group-fill','','/0/543/546/627','F','','general:orderRating:rating',546,0,'','',30,'0','1',1,1,'2024-09-23 20:41:04.810','2024-10-23 14:43:44.018',NULL),
    (628,'','工单留言','app-group-fill','','/0/543/546/628','F','','general:orderComment:comment',546,0,'','',40,'0','1',1,1,'2024-09-23 20:42:44.303','2024-10-23 14:43:44.019',NULL);

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_menu_api_rule
# ------------------------------------------------------------

LOCK TABLES `sys_menu_api_rule` WRITE;
/*!40000 ALTER TABLE `sys_menu_api_rule` DISABLE KEYS */;

INSERT INTO `sys_menu_api_rule` (`sys_menu_menu_id`, `sys_api_id`)
VALUES
    (528,5),
    (216,6),
    (250,6),
    (528,6),
    (528,7),
    (528,8),
    (528,9),
    (528,10),
    (528,11),
    (528,15),
    (528,16),
    (58,21),
    (236,21),
    (528,21),
    (238,23),
    (528,23),
    (59,24),
    (240,24),
    (528,24),
    (242,25),
    (528,25),
    (58,26),
    (236,26),
    (56,27),
    (228,27),
    (528,27),
    (544,27),
    (605,27),
    (230,28),
    (528,28),
    (226,29),
    (51,39),
    (222,39),
    (528,39),
    (221,41),
    (528,41),
    (52,44),
    (225,44),
    (528,44),
    (607,44),
    (226,45),
    (226,46),
    (226,47),
    (528,47),
    (62,53),
    (244,53),
    (528,53),
    (246,54),
    (528,54),
    (528,57),
    (528,58),
    (57,59),
    (232,59),
    (528,59),
    (234,60),
    (528,60),
    (528,66),
    (528,67),
    (528,72),
    (528,73),
    (528,76),
    (241,80),
    (528,80),
    (237,81),
    (528,81),
    (229,82),
    (528,82),
    (245,87),
    (528,87),
    (220,88),
    (528,88),
    (233,89),
    (528,89),
    (224,90),
    (528,90),
    (528,92),
    (531,92),
    (528,95),
    (528,96),
    (528,97),
    (242,101),
    (528,101),
    (238,102),
    (528,102),
    (230,103),
    (528,103),
    (528,104),
    (528,105),
    (226,106),
    (528,106),
    (226,107),
    (528,107),
    (246,108),
    (528,108),
    (221,109),
    (528,109),
    (234,110),
    (528,110),
    (528,112),
    (528,113),
    (249,114),
    (528,114),
    (251,115),
    (528,115),
    (528,116),
    (528,117),
    (243,120),
    (528,120),
    (239,121),
    (528,121),
    (231,122),
    (528,122),
    (528,123),
    (528,124),
    (247,125),
    (528,125),
    (223,126),
    (528,126),
    (235,127),
    (528,127),
    (227,128),
    (528,128),
    (51,135),
    (528,135),
    (529,135),
    (528,136),
    (531,136),
    (212,137),
    (248,137),
    (528,137),
    (528,139),
    (542,139),
    (528,140),
    (540,140),
    (3,141),
    (44,141),
    (528,141),
    (606,141),
    (609,141),
    (45,142),
    (528,142),
    (43,150),
    (528,150),
    (45,151),
    (528,151),
    (46,156),
    (528,156),
    (603,244),
    (571,245),
    (572,246),
    (573,247),
    (604,248),
    (602,249),
    (608,250),
    (576,251),
    (580,251),
    (584,252),
    (600,252),
    (577,253),
    (596,253),
    (600,254),
    (554,255),
    (613,256),
    (585,257),
    (590,258),
    (591,259),
    (601,259),
    (608,259),
    (595,260),
    (569,263),
    (602,264),
    (581,265),
    (597,266),
    (610,267),
    (586,268),
    (589,269),
    (592,270),
    (579,273),
    (574,274),
    (602,275),
    (582,276),
    (611,278),
    (587,279),
    (593,280),
    (583,282),
    (614,284),
    (588,286),
    (594,287),
    (627,301),
    (627,308),
    (627,309),
    (628,310),
    (628,311),
    (628,312),
    (628,314),
    (627,315),
    (577,322),
    (577,323),
    (577,324);

/*!40000 ALTER TABLE `sys_menu_api_rule` ENABLE KEYS */;
UNLOCK TABLES;

# 转储表 sys_post
# ------------------------------------------------------------

LOCK TABLES `sys_post` WRITE;
/*!40000 ALTER TABLE `sys_post` DISABLE KEYS */;

INSERT INTO `sys_post` (`post_id`, `post_name`, `post_code`, `sort`, `status`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,'首席执行官','CEO',0,2,'首席执行官',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
    (2,'首席技术执行官','CTO',2,2,'首席技术执行官',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
    (3,'首席运营官','COO',3,2,'测试工程师',1,1,'2021-05-13 19:56:37.913','2021-05-13 19:56:37.913',NULL),
    (4,'后端工程师','RD',10,2,'',1,1,'2024-08-15 19:33:03.640','2024-08-15 19:38:33.836',NULL),
    (5,'测试工程师','TEST',11,2,'测试工程师',1,1,'2024-08-15 19:33:33.276','2024-08-15 19:33:42.630',NULL),
    (6,'数据工程师','BIGDATA',12,2,'',1,0,'2024-08-15 19:34:50.562','2024-08-15 19:34:50.562',NULL),
    (7,'运维工程师','OPS',13,2,'',1,0,'2024-08-15 19:35:11.535','2024-08-15 19:35:11.535',NULL),
    (8,'运维开发工程师','DEVOPS',14,2,'',1,1,'2024-08-15 19:35:22.771','2024-08-15 19:36:44.342',NULL),
    (9,'产品经理','PM',15,2,'',1,0,'2024-08-15 19:35:59.535','2024-08-15 19:35:59.535',NULL),
    (10,'前端工程师','FE',16,2,'',1,0,'2024-08-15 19:39:01.303','2024-08-15 19:39:01.303',NULL);

/*!40000 ALTER TABLE `sys_post` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_role
# ------------------------------------------------------------

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`role_id`, `role_name`, `status`, `role_key`, `role_sort`, `flag`, `remark`, `admin`, `data_scope`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,'系统管理员','2','admin',1,'','',1,'1',1,1,'2021-05-13 19:56:37.913','2024-07-11 21:19:30.625',NULL),
    (2,'普通用户','2','general',2,'','',0,'2',0,0,'2024-08-08 21:13:00.494','2024-09-29 17:39:23.379',NULL);

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_role_dept
# ------------------------------------------------------------

LOCK TABLES `sys_role_dept` WRITE;
/*!40000 ALTER TABLE `sys_role_dept` DISABLE KEYS */;

INSERT INTO `sys_role_dept` (`role_id`, `dept_id`)
VALUES
    (2,1),
    (2,7),
    (2,8),
    (2,9),
    (2,10),
    (2,11);

/*!40000 ALTER TABLE `sys_role_dept` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_role_menu
# ------------------------------------------------------------

LOCK TABLES `sys_role_menu` WRITE;
/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;

INSERT INTO `sys_role_menu` (`role_id`, `menu_id`)
VALUES
    (2,269),
    (2,459),
    (2,460),
    (2,461),
    (2,462),
    (2,463),
    (2,464),
    (2,471),
    (2,537),
    (2,543),
    (2,544),
    (2,545),
    (2,546),
    (2,547),
    (2,553),
    (2,554),
    (2,555),
    (2,556),
    (2,569),
    (2,570),
    (2,571),
    (2,572),
    (2,573),
    (2,574),
    (2,575),
    (2,576),
    (2,577),
    (2,578),
    (2,579),
    (2,602),
    (2,603),
    (2,604),
    (2,608),
    (2,609),
    (2,610),
    (2,611),
    (2,613),
    (2,614),
    (2,627),
    (2,628);

/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_tables
# ------------------------------------------------------------

LOCK TABLES `sys_tables` WRITE;
/*!40000 ALTER TABLE `sys_tables` DISABLE KEYS */;

INSERT INTO `sys_tables` (`table_id`, `table_name`, `table_comment`, `class_name`, `tpl_category`, `package_name`, `module_name`, `module_front_name`, `business_name`, `function_name`, `function_author`, `pk_column`, `pk_go_field`, `pk_json_field`, `options`, `tree_code`, `tree_parent_code`, `tree_name`, `tree`, `crud`, `remark`, `is_data_scope`, `is_actions`, `is_auth`, `is_logical_delete`, `logical_delete`, `logical_delete_column`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`)
VALUES
    (1,'flow_manage','FlowManage','FlowManage','crud','admin','flow-manage','','flowManage','FlowManage','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-08-12 20:27:59.109','2024-08-12 20:27:59.109',NULL,0,0),
    (2,'exec_machine','ExecMachine','ExecMachine','crud','admin','exec-machine','','execMachine','ExecMachine','wenjianzhang','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-01 23:28:33.293','2024-09-01 23:32:57.619',NULL,0,0),
    (3,'flow_templates','FlowTemplates','FlowTemplates','crud','admin','flow-templates','','flowTemplates','FlowTemplates','sunwenbo','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 10:44:06.422','2024-09-02 10:44:06.422',NULL,0,0),
    (4,'operation_history','OperationHistory','OperationHistory','crud','admin','operation-history','','operationHistory','OperationHistory','sunwenbo','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 10:51:16.236','2024-09-02 10:51:16.236',NULL,0,0),
    (5,'order_category','OrderCategory','OrderCategory','crud','admin','order-category','','orderCategory','OrderCategory','sunwenbo','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 11:00:07.949','2024-09-02 11:00:07.949',NULL,0,0),
    (6,'order_items','OrderItems','OrderItems','crud','admin','order-items','','orderItems','OrderItems','sunwenbo','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 11:09:12.761','2024-09-02 11:09:12.761',NULL,0,0),
    (7,'order_task','OrderTask','OrderTask','crud','admin','order-task','','orderTask','工单任务','sunwenbo','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 11:24:26.143','2024-09-02 17:03:59.707',NULL,0,0),
    (8,'order_works','OrderWorks','OrderWorks','crud','admin','order-works','','orderWorks','OrderWorks','sunwenbo','id','Id','id','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 11:26:24.562','2024-09-02 11:26:24.562',NULL,0,0),
    (9,'sys_menu','SysMenu','SysMenu','crud','admin','sys-menu','','sysMenu','SysMenu','sunwenbo','menu_id','MenuId','menuId','','','','',0,1,'',1,2,1,'1',1,'is_del','2024-09-02 17:04:30.470','2024-09-02 17:04:30.470',NULL,0,0);

/*!40000 ALTER TABLE `sys_tables` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 sys_user
# ------------------------------------------------------------

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `dept_id`, `post_id`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `source`)
VALUES
    (1,'admin','$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.','admin','13311246030',1,'','','0','swb956721830@163.com',1,2,'','2',1,1,'2021-05-13 19:56:37.914','2024-08-16 17:29:33.360',NULL,'SYSTEM'),
    (2,'sunwenbo','$2a$10$ZBzBhm1Z6Rn0I.ZkECdoAeBjO/lzigy1XT3ObBt/DS6ZfRuLtfxtq','孙文波','13311246030',1,'','','0','123@qq.com',8,8,'','2',1,0,'2024-07-17 20:05:43.749','2024-08-15 19:36:25.064',NULL,'SYSTEM'),
    (3,'zhangsan','$2a$10$bttFFWqTZMyyoBzHLDE.A.hHJ5q5oqcukh16CUXyQvFTPg0bRhK8m','张三','13312312312',1,'','','0','123@qq.com',1,5,'','2',1,0,'2024-07-30 20:35:19.622','2024-08-15 19:37:01.935',NULL,'SYSTEM'),
    (4,'lisi','$2a$10$ZN4TSBEBvxXuVIHx3sAeceu/M1T7T4t1/1dgA9clOhgAsjyjIumHy','李四','13311246030',2,'','','1','123@qq.com',7,4,'测试用户','2',1,1,'2024-08-08 21:14:03.237','2024-08-15 19:37:10.409',NULL,'SYSTEM'),

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 tb_demo
# ------------------------------------------------------------



# 转储表 user_favorites
# ------------------------------------------------------------

LOCK TABLES `user_favorites` WRITE;
/*!40000 ALTER TABLE `user_favorites` DISABLE KEYS */;

INSERT INTO `user_favorites` (`id`, `user_id`, `order_item_id`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1,1,2,1,0,'2024-10-17 11:51:01.833','2024-10-17 11:51:01.833',NULL);

/*!40000 ALTER TABLE `user_favorites` ENABLE KEYS */;
UNLOCK TABLES;
