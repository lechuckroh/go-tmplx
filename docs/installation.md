# Installation Guide

go-tmplx를 설치하는 다양한 방법들을 설명합니다.

## Prerequisites

- Linux, macOS, 또는 Windows
- Docker (Docker 이미지 사용 시)

## Installation Methods

### 1. Homebrew (macOS/Linux)

가장 쉬운 설치 방법입니다.

```bash
brew install lechuckroh/tap/tmplx
```

### 2. Manual Installation

#### Step 1: 릴리즈 페이지에서 다운로드

[GitHub 릴리즈 페이지](https://github.com/lechuckroh/go-tmplx/releases)에서 운영체제에 맞는 바이너리를 다운로드하세요.

#### Step 2: 바이너리 설치

**Linux/macOS:**
```bash
# 다운로드한 파일을 실행 가능하게 만들기
chmod +x tmplx

# 시스템 PATH에 추가 (예: /usr/local/bin)
sudo mv tmplx /usr/local/bin/
```

**Windows:**
```bash
# 다운로드한 zip 파일을 압축 해제
# tmplx.exe를 PATH에 있는 디렉토리로 이동
```

### 3. Docker

Docker를 사용하여 실행할 수 있습니다.

```bash
# 최신 버전 실행
docker run --rm -v $(pwd):/work ghcr.io/lechuckroh/go-tmplx:latest input.tmpl output.txt

# 환경 변수 파일과 접두사 사용
docker run --rm -v $(pwd):/work ghcr.io/lechuckroh/go-tmplx:latest input.tmpl output.txt -e .env -p APP_

# 특정 버전 실행
docker run --rm -v $(pwd):/work ghcr.io/lechuckroh/go-tmplx:v1.0.0 input.tmpl output.txt
```

### 4. Go Install (개발자용)

Go가 설치되어 있다면 직접 설치할 수 있습니다.

```bash
go install github.com/lechuckroh/go-tmplx/cmd/tmplx@latest
```

## Verification

설치가 완료되었는지 확인하세요:

```bash
tmplx --version
```

## Troubleshooting

### 권한 오류 (Linux/macOS)
```bash
chmod +x tmplx
```

### PATH 문제
바이너리가 PATH에 있는지 확인하세요:
```bash
which tmplx
```

### Docker 권한 문제
Docker 그룹에 사용자를 추가하세요:
```bash
sudo usermod -aG docker $USER
```

## Uninstallation

### Homebrew
```bash
brew uninstall tmplx
```

### Manual Installation
```bash
# 바이너리 파일 삭제
rm /usr/local/bin/tmplx
```

### Docker
```bash
# 이미지 삭제
docker rmi ghcr.io/lechuckroh/go-tmplx:latest
``` 