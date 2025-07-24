# go-tmplx

**go-tmplx**는 Go 템플릿 파일을 환경 변수와 함께 렌더링하는 경량 CLI 도구입니다.

개발자와 DevOps 엔지니어가 동적으로 설정 파일, 스크립트 또는 텍스트 파일을 생성해야 할 때 사용하도록 설계되었습니다.

## 주요 기능

- Go의 강력한 `text/template` 문법을 사용한 템플릿 렌더링
- 환경 변수 파일(.env)에서 변수 로드
- 시스템 환경 변수와의 우선순위 규칙
- CI/CD 파이프라인에서 자동화에 이상적인 간단하고 빠른 CLI 인터페이스

## 빠른 시작

### 설치

**수동 설치:**
[릴리즈 페이지](https://github.com/lechuckroh/go-tmplx/releases)에서 운영체제에 맞는 바이너리를 다운로드하세요.


### 기본 사용법

**템플릿 파일 (config.tmpl):**
```go
server:
  host: {{.HOST}}
  port: {{.PORT}}
  debug: {{.DEBUG}}
```

**환경 변수 파일 (.env):**
```env
HOST=localhost
PORT=8080
DEBUG=true
```

**실행:**
```bash
tmplx config.tmpl config.yaml -e .env
```

**환경 변수 접두사 사용:**
```bash
tmplx config.tmpl config.yaml -e .env -p APP_
```

**결과:**
```yaml
server:
  host: localhost
  port: 8080
  debug: true
```

## 문서

- **[설치 가이드](docs/installation.md)** - 다양한 설치 방법
- **[사용법 가이드](docs/usage.md)** - 상세한 사용법과 예제
- **[개발자 가이드](docs/development.md)** - 개발 환경 설정 및 빌드

## 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.
