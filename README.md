# SocialRecon 🔍

[![Go Version](https://img.shields.io/github/go-mod/go-version/ismailtsdln/socialrecon)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**SocialRecon** is a high-performance, open-source social media reconnaissance and OSINT security scanner. It identifies social media presence, abandoned profiles, impersonation risks, and brand abuse from domains or usernames.

## 🚀 Key Features

- **Concurrent Scanning**: Fast discovery using Go worker pools.
- **Domain Discovery**: Automatically extracts social links from website HTML/JS.
- **Risk Scoring**: Intelligent severity assessment (INFO to CRITICAL).
- **Multi-Format Reporting**: Export results to CLI, JSON, or professional HTML dashboards.
- **Modular Plugin System**: Easily extensible for new social platforms.

## 📦 Installation

```bash
go install github.com/ismailtsdln/socialrecon@latest
```

## 🛠️ Usage

### Scan a Username

```bash
socialrecon scan johndoe
```

### Scan a Domain (Discovery Mode)

```bash
socialrecon scan example.com
```

### Export HTML Report

```bash
socialrecon scan example.com --html-report report.html
```

### JSON Output for Automation

```bash
socialrecon scan johndoe --json
```

## 🧠 Risk Scoring System

SocialRecon uses a weighted scoring engine to evaluate OSINT findings:

| Status | Risk Level | Description |
|--------|------------|-------------|
| **Available** | HIGH | Profile can be registered/hijacked. |
| **Suspended** | MEDIUM | Platform-level enforcement detected. |
| **Exists** | INFO | Passive profile discovery. |

## 🛡️ Ethics & Legal Disclaimer

SocialRecon is designed for passive reconnaissance and authorized security testing.

- No credential stuffing.
- No brute-force checking.
- Respects `robots.txt`.
The author assumes no liability for misuse.

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---
Built with ❤️ by [Ismail Tasdelen](https://github.com/ismailtsdln)
