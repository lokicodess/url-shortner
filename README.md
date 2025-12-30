# 🔗 Production-Grade URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://golang.org)
[![React](https://img.shields.io/badge/React-18+-61DAFB?logo=react)](https://reactjs.org)
[![Kubernetes](https://img.shields.io/badge/K3s-1.27+-326CE5?logo=kubernetes)](https://k3s.io)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A production-ready URL shortening service built with **Go, React, and K3s**, deployed on AWS with full CI/CD pipeline.

## ✨ Features

- ✅ URL shortening custom slugs
- ✅ Automatic expiration (TTL-based)
- ✅ High availability with K3s cluster
- ✅ Jenkins CI/CD pipeline
- ✅ AWS RDS
- ✅ Monitoring and Observability using prometheus and Grafana
- ✅ GitOps workflow

## 🏗️ Architecture
<img width="1251" height="200" alt="image" src="https://github.com/user-attachments/assets/c5299ddd-61ab-4f39-b1e4-a3cff43024f8" />
For detailed architecture documentation, see [ARCHITECTURE.md](docs/ARCHITECTURE.md)


## 🚀 Quick Start

```bash
git clone https://github.com/lokicodess/url-shortener.git
cd url-shortener
docker-compose up -d
```
## 📚 Documentation

- [📖 **Architecture Details**](docs/ARCHITECTURE.md) - System design, infrastructure decisions, and cost optimization
- [🚀 **Deployment Guide**](docs/DEPLOYMENT.md) - Step-by-step setup and production deployment
- [🔧 **API Documentation**](docs/API.md) - Complete API reference with examples
- [⚙️ **Development Setup**](docs/DEVELOPMENT.md) - Local development environment and coding guidelines


## 🔗 Live Demo

| Environment | URL | Status |
|------------|-----|--------|
| **Frontend** | [https://clck.dev](https://clck.dev) | ✅ Live |
| **API** | [https://api.clck.dev](https://api.clck.dev/) | ✅ Live |
---

## 🛠️ Tech Stack

### Frontend
- **React** - UI library with modern hooks
- **Tailwind CSS** - Utility-first CSS framework

### Backend
- **Go (Golang) 1.24+** - High-performance server

### Database
- **AWS RDS** - Managed SQL service 
- **Database Migrations** - Version-controlled schema changes

### Cache
- **AWS ElastiCache** - Managed Redis service

### Infrastructure
- **AWS** - Cloud hosting (EC2, RDS, ElastiCache)
- **K3s** - Lightweight Kubernetes distribution
- **Jenkins/ArgoCD** - CI/CD automation server
- **Docker** - Containerization
- **Nginx** - Reverse proxy and load balancer

### Monitoring & Logging
- **Prometheus** - Metrics collection
- **Grafana** - Metrics visualization



