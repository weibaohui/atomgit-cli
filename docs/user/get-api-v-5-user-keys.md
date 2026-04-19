# 列出授权用户的所有公钥

## 基本信息

| 项目 | 值 |
|------|-----|
| **HTTP Method** | GET |
| **Endpoint** | `https://api.atomgit.com/api/v5/user/keys` |
| **文档链接** | https://docs.atomgit.com/docs/apis/get-api-v-5-user-keys |

## Path Parameters

| - | - | - | - |

## Query Parameters

| 参数名 | 类型 | 必填 | 描述 |
|--------|------|------|------|
| access_token | string | 是 | 用户授权码 |

## 请求示例

```bash
curl "https://api.atomgit.com/api/v5/user/keys?access_token=$ATOMGIT_TOKEN" \\n  -H "Authorization: token $ATOMGIT_TOKEN"
```


## 响应示例

```json
[{"id":10122,"title":"weibaohui@yeah.net","key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCr7igSjGaFfuo7TDaMR6STG7ecRpHy0Up690Gm8AiL+rmapMZTWy0YAfDrRS9h9k8QeDtFJDAQcbee28t3kLmcPBsoBHV38wcwWqDq0pF6wWInE0J474Cq5xL8pRhMwIF75EeshHzw4l4JbdSVK8IJ0ztWDdfQtsoLFinf5ckxmQt/o+OanE0sAjpvYDX5HKS8tsZhtl/h4IXeEbSZKR90b5gfH3Ijl02Ikrkvd4Hj0q0/S41WJkfjx9t0nS4y7hqzzYASjTqHgG/M0G8bxlUSkvV6DR9QbHLNiSei1EtbRjAYa2/v+nMTN1AhIgvNWJieWzOJ6IhZ8E5EJwxuPz4mEMjS1CjjfBZFa64BKyGpvplZH/8c3ZCOHlgBfHYeaYLC3VZUXZlXhpaA3wrL6ZXxUoCaG9Pbohi5l745qOP4At2nEnfFqdU0XbnV4RqC3CwjSxADb4DAdhuDIdYTCu/VokqyO54/6DuEHM/MFuczzaF/g6vshyVt4E0cOLwVm2VpBxS84ZuDM8Td8o2wa0U2lFKycGDvLd+j48952d+z3MIXR3zq0HDOUrNg2evrr/ZmlwiIy0lcP5Cgv1smLmm0UbaIL/+pmz+TlxxDBzYFWcPD+i8uyMD5aBP+bD+vzKU751HRbenzfWGRcT/xqrGUUotlOG2c5xi8fR3Mf6cPxw== weibaohui@yeah.net","created_at":"2024-09-11T14:28:03.509+00:00","url":"https://api.atomgit.com/v5/user/keys/10122"}]
```

## 相关 CLI 命令

参见 [API列表](../../apis_list.md)
