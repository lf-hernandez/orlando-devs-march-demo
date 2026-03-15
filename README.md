# Git Branching Strategies Demo - ODevs March 2026

[![ci](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/ci.yml/badge.svg)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/ci.yml)

[![cd (trunk-based)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-trunk-based.yml/badge.svg)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-trunk-based.yml)

[![cd (git-flow)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-gitflow.yml/badge.svg)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-gitflow.yml)

[![cd (github-flow)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-github-flow.yml/badge.svg)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-github-flow.yml)

[![cd (gitlab-flow)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-gitlab-flow.yml/badge.svg)](https://github.com/lf-hernandez/orlando-devs-march-demo/actions/workflows/cd-gitlab-flow.yml)

## CD Workflows

Each CD workflow triggers on successful completion of the CI workflow on specific branches and deploys to the corresponding Railway environment.

| Workflow | Trigger Branch(es) | Deploys To |
|---|---|---|
| Trunk-based | `main` | production |
| Git Flow | `develop` | develop |
| Git Flow | `hotfix/**` | develop |
| Git Flow | `release/**` | staging |
| Git Flow | `main` | production |
| GitHub Flow | `feature/**` | staging (preview) |
| GitHub Flow | `main` | production |
| GitLab Flow | `main` | develop |
| GitLab Flow | `staging` | staging |
| GitLab Flow | `production` | production |

## Getting Started

### To build and run locally:

```bash
  docker build -t orlando-devs-demo .
  docker run -p 3000:3000 orlando-devs-demo
```

### To inject version info at build time:

```bash
  docker build \
    --build-arg VERSION=1.0.0 \
    --build-arg COMMIT=$(git rev-parse --short HEAD) \
    --build-arg BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
    -t orlando-devs-demo .
```
