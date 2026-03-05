# Git Branching Strategies Demo - ODevs March 2026

## Getting Started

### To build and run locally:

```bash
  docker build -t orlando-devs-demo .
  docker run -p 5000:5000 orlando-devs-demo
```

### To inject version info at build time:

```bash
  docker build \
    --build-arg VERSION=1.0.0 \
    --build-arg COMMIT=$(git rev-parse --short HEAD) \
    --build-arg BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
    -t orlando-devs-demo .
```
