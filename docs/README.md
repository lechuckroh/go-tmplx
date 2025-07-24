# Documentation

go-tmplx 프로젝트의 문서입니다.

## 문서 목록

### 사용자 가이드
- **[Installation Guide](installation.md)** - 설치 방법
- **[Usage Guide](usage.md)** - 사용법 및 예제

### 개발자 가이드
- **[Development Guide](development.md)** - 개발 환경 설정 및 빌드

## 빠른 시작

### 설치
```bash
# Homebrew (macOS/Linux)
brew install lechuckroh/tap/tmplx

# 또는 수동 설치
# GitHub 릴리즈 페이지에서 다운로드
```

### 기본 사용법
```bash
# 환경 변수 파일과 함께 템플릿 처리
tmplx config.tmpl config.yaml -e .env

# 시스템 환경 변수 사용
export HOST=localhost
export PORT=8080
tmplx config.tmpl config.yaml
```

## 기여하기

문서 개선에 기여하고 싶으시다면:

1. 이슈를 생성하여 개선 사항을 제안
2. Pull Request를 통해 직접 수정
3. 오타나 오류 발견 시 이슈로 보고

## 문서 작성 가이드

- 모든 문서는 한국어로 작성
- 코드 예제는 실제 동작하는 것을 포함
- 명령어는 복사하여 바로 사용할 수 있도록 작성
- 스크린샷이나 다이어그램이 필요한 경우 추가 고려 