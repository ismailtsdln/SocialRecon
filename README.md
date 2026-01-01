# SocialRecon 🔍

[![Go Version](https://img.shields.io/github/go-mod/go-version/ismailtsdln/socialrecon)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**SocialRecon** is a high-performance, open-source social media reconnaissance and OSINT security scanner. It identifies social media presence, abandoned profiles, impersonation risks, and brand abuse from domains or usernames.

---

## 🚀 Key Features

- **Concurrent Scanning**: High-speed discovery using Go worker pools and goroutines.
- **Domain Discovery**: Automatically extracts social links from website HTML, Meta tags, and JS.
- **Risk Scoring Engine**: Intelligent severity assessment based on platform authority and profile status.
- **Executive Reporting**: Export results to CLI (color-coded), JSON, or professional HTML dashboards.
- **Modular Plugin System**: Easily extensible architecture for adding new platforms.

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
socialrecon scan example.com --verbose
```

### Export HTML Dashboard

```bash
socialrecon scan example.com --html-report report.html
```

### Options

| Flag | Description |
|------|-------------|
| `--json` | Output results in machine-readable JSON format |
| `--html-report [path]` | Generate a professional HTML report |
| `--verbose` | Enable detailed scan logging |

## 🧠 Risk Scoring System

SocialRecon evaluates OSINT findings using a weighted algorithm:

| Status | Risk Level | Description |
| :--- | :--- | :--- |
| **Available** | HIGH | Profile is available for registration (potential hijacking/squatting). |
| **Suspended** | MEDIUM | Profile exists but has been suspended by the platform. |
| **Exists** | INFO | Social media presence identified for the specified target. |

## 🛡️ Threat Model & Ethics

- **Passive Reconnaissance**: The tool only performs passive checks (HTTP GET) and does not interact with platform APIs in a way that requires credentials.
- **No Credential Stuffing**: Does not attempt to log in or use leaked credentials.
- **Legal Compliance**: Designed for authorized security audits, bug bounties, and brand protection.

## 🤝 Contribution

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the Repository
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---
Built with ❤️ by [Ismail Tasdelen](https://github.com/ismailtsdln)
