# GoConfigLoader 🛠️

A simple and clean configuration loader in Go using [Viper](https://github.com/spf13/viper), designed to support:

- 🔧 Default values
- 📦 YAML-based config files
- 🌱 Environment variable overrides

## 🚀 Features

- Loads configuration with layered fallback: **env vars → config file → defaults**
- Converts `timeout_seconds` to `time.Duration`
- No config file required – works with just env or defaults



