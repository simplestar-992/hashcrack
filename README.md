# HashCrack | Password Hash Identifier & Checker

![Security Tool](https://img.shields.io/badge/Purpose-Hash%20Analysis-red?style=for-the-badge)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

---

## Identify & Analyze Password Hashes

HashCrack helps you identify hash types, check hashes against common password lists, and understand hash security.

**Perfect for:**
- Security professionals auditing password databases
- Developers learning about password storage
- CTF players solving hash challenges
- System administrators checking password policies

---

## Features

- 🔍 **Auto-detect** - Identifies 50+ hash types automatically
- ✅ **Verify** - Check if a hash matches a password
- 📊 **Analyze** - Understand hash strength and type
- 🔓 **Wordlist ready** - Test against common password lists
- ⚡ **Fast** - Concurrent hash checking

---

## Installation

```bash
git clone https://github.com/simplestar-992/hashcrack.git
cd hashcrack
go build -o hashcrack -ldflags="-s -w"
```

---

## Usage

```bash
# Identify hash type
./hashcrack identify "5f4dcc3b5aa765d61d8327deb882cf99"

# Verify password against hash
./hashcrack verify "5f4dcc3b5aa765d61d8327deb882cf99" "password"

# Check with wordlist
./hashcrack wordlist -hash "hash.txt" -wordlist "rockyou.txt"

# Hash a password
./hashcrack hash "mypassword" -type md5
```

---

## Supported Hash Types

| Category | Types |
|----------|-------|
| **MD5 Family** | MD5, MD4, MD2, MD5(md5), Double MD5 |
| **SHA Family** | SHA1, SHA256, SHA384, SHA512 |
| **Bcrypt** | bcrypt, bcrypt_sha256 |
| **Scrypt** | scrypt, PBKDF2 |
| **NTLM** | NTLM, NTLMv2 |
| **MySQL** | MySQL323, MySQLSHA1 |

---

## Examples

```bash
# Identify unknown hash
./hashcrack identify "2fc5a684dce1773774a1bb71f2e2a8b2"

# Check if admin password is in breach database
./hashcrack verify -hash "$1$xyz$abcdef" -wordlist common.txt

# Batch processing
cat hashes.txt | while read hash; do ./hashcrack identify "$hash"; done
```

---

## License

MIT © 2024 [simplestar-992](https://github.com/simplestar-992)
