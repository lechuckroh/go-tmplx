# Usage Guide

go-tmplx의 사용법과 다양한 예제를 설명합니다.

## Basic Usage

### 기본 문법
```bash
tmplx <input_template> <output_file> [options]
```

### 필수 인수
- `input_template`: 템플릿 파일 경로
- `output_file`: 출력 파일 경로

### 옵션
- `-e, --env`: 환경 변수 파일 (.env) 경로
- `-p, --env-prefix`: 환경 변수 접두사 (예: APP_, DB_)
- `--version`: 버전 정보 출력
- `--help`: 도움말 출력

## Examples

### 1. 기본 템플릿 처리

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

**결과 (config.yaml):**
```yaml
server:
  host: localhost
  port: 8080
  debug: true
```

### 2. 시스템 환경 변수 사용

```bash
export HOST=production.example.com
export PORT=443
export DEBUG=false

tmplx config.tmpl config.yaml
```

### 3. 환경 변수 접두사 사용

환경 변수에 접두사를 사용하여 네임스페이스를 분리할 수 있습니다.

**환경 변수 파일 (.env):**
```env
APP_HOST=localhost
APP_PORT=8080
APP_DEBUG=true
DB_HOST=database
DB_PORT=5432
DB_NAME=myapp
```

**실행:**
```bash
# APP_ 접두사만 사용
tmplx config.tmpl config.yaml -e .env -p APP_

# DB_ 접두사만 사용
tmplx db-config.tmpl db-config.yaml -e .env -p DB_
```

**템플릿에서는 접두사 없이 사용:**
```go
server:
  host: {{.HOST}}  # APP_HOST 값
  port: {{.PORT}}  # APP_PORT 값
  debug: {{.DEBUG}}  # APP_DEBUG 값
```

### 4. 복잡한 템플릿 예제

**템플릿 파일 (docker-compose.tmpl):**
```yaml
version: '3.8'
services:
  app:
    image: {{.APP_IMAGE}}
    ports:
      - "{{.APP_PORT}}:8080"
    environment:
      - DB_HOST={{.DB_HOST}}
      - DB_PORT={{.DB_PORT}}
      - DB_NAME={{.DB_NAME}}
      - DB_USER={{.DB_USER}}
      - DB_PASSWORD={{.DB_PASSWORD}}
    volumes:
      - {{.APP_DATA}}:/app/data
    restart: unless-stopped

  database:
    image: {{.DB_IMAGE}}
    ports:
      - "{{.DB_PORT}}:5432"
    environment:
      - POSTGRES_DB={{.DB_NAME}}
      - POSTGRES_USER={{.DB_USER}}
      - POSTGRES_PASSWORD={{.DB_PASSWORD}}
    volumes:
      - {{.DB_DATA}}:/var/lib/postgresql/data
```

**환경 변수 파일 (.env):**
```env
APP_IMAGE=myapp:latest
APP_PORT=3000
APP_DATA=./data
DB_IMAGE=postgres:15
DB_HOST=database
DB_PORT=5432
DB_NAME=myapp
DB_USER=myuser
DB_PASSWORD=mypassword
DB_DATA=./postgres-data
```

**실행:**
```bash
tmplx docker-compose.tmpl docker-compose.yml -e .env
```

### 5. 조건문과 반복문 사용

**템플릿 파일 (nginx.tmpl):**
```nginx
server {
    listen {{.NGINX_PORT}};
    server_name {{.NGINX_HOST}};

    {{if .NGINX_SSL}}
    ssl_certificate {{.SSL_CERT}};
    ssl_certificate_key {{.SSL_KEY}};
    {{end}}

    location / {
        proxy_pass http://{{.APP_HOST}}:{{.APP_PORT}};
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    {{range .STATIC_PATHS}}
    location {{.Path}} {
        alias {{.Alias}};
        expires 1y;
    }
    {{end}}
}
```

**환경 변수 파일 (.env):**
```env
NGINX_PORT=80
NGINX_HOST=example.com
NGINX_SSL=true
SSL_CERT=/etc/ssl/certs/example.com.crt
SSL_KEY=/etc/ssl/private/example.com.key
APP_HOST=localhost
APP_PORT=8080
STATIC_PATHS='[{"Path":"/static","Alias":"/var/www/static"},{"Path":"/images","Alias":"/var/www/images"}]'
```

## Template Syntax

go-tmplx는 Go의 `text/template` 패키지를 사용합니다.

### 변수 출력
```go
{{.VARIABLE_NAME}}
```

### 조건문
```go
{{if .CONDITION}}
  조건이 참일 때 실행
{{else}}
  조건이 거짓일 때 실행
{{end}}
```

### 반복문
```go
{{range .ARRAY}}
  {{.}} - 현재 요소
{{end}}
```

### 함수 사용
```go
{{.VARIABLE | upper}}  // 대문자 변환
{{.VARIABLE | lower}}  // 소문자 변환
{{.VARIABLE | title}}  // 제목 케이스 변환
```

## Environment Variables

### 우선순위
1. 명령줄에서 지정한 환경 변수
2. .env 파일의 변수
3. 시스템 환경 변수

### 네이밍 규칙
- 환경 변수는 접두사를 사용할 수 있습니다 (예: `APP_`, `DB_`, `GW_`)
- `--env-prefix` 옵션으로 특정 접두사만 필터링할 수 있습니다
- 템플릿에서는 접두사 없이 사용합니다

예시:
```env
APP_HOST=localhost  # 환경 변수
DB_HOST=database    # 다른 접두사
```

템플릿에서:
```go
{{.HOST}}  # 접두사에 따라 다른 값 출력
```

**접두사 필터링:**
```bash
# APP_ 접두사만 사용
tmplx config.tmpl config.yaml -e .env -p APP_

# DB_ 접두사만 사용  
tmplx db-config.tmpl db-config.yaml -e .env -p DB_
```

## Best Practices

### 1. 템플릿 파일 관리
- 템플릿 파일은 `.tmpl` 확장자 사용
- 환경별로 다른 .env 파일 사용
- 템플릿 파일을 버전 관리에 포함

### 2. 환경 변수 관리
- 민감한 정보는 .env 파일에 저장
- .env 파일을 .gitignore에 추가
- 환경별로 다른 .env 파일 사용

### 3. CI/CD 파이프라인에서 사용
```yaml
# GitHub Actions 예제
- name: Generate config
  run: |
    tmplx config.tmpl config.yaml -e .env.production
```

## Troubleshooting

### 템플릿 오류
```bash
# 템플릿 문법 검증
tmplx --dry-run config.tmpl
```

### 환경 변수 문제
```bash
# 환경 변수 확인
env | grep APP_  # 특정 접두사 확인
env | grep DB_   # 다른 접두사 확인
```

### 권한 문제
```bash
# 출력 파일 권한 확인
ls -la output.yaml
``` 