# 개발 가이드

이 문서는 go-tmplx 프로젝트의 개발자를 위한 가이드입니다.

## 사전 준비

- Go 1.24 이상
- [go-task](https://taskfile.dev/) (빌드 도구)
- [go-releaser](https://goreleaser.com/) (릴리즈 도구)

## 개발 환경 설정

### 1. 저장소 체크아웃
```bash
git clone https://github.com/lechuckroh/go-tmplx.git
cd go-tmplx
```

### 2. 의존성 설치
```bash
task install
```

### 3. go-releaser 설치
```bash
task install-goreleaser
```

## 개발 워크플로우

### 코드 포맷팅
```bash
task format
```

### 테스트 실행
```bash
task test
```

### 로컬 빌드
```bash
task build
```

### 로컬 실행
```bash
task run -- input.tmpl output.txt
```

## GoReleaser 사용

이 프로젝트는 [GoReleaser](https://goreleaser.com/)를 사용하여 여러 플랫폼에서 바이너리를 빌드하고 릴리즈합니다.

### 설정 검증
```bash
task goreleaser-check
```

### 스냅샷 릴리즈 생성 (테스트용)
```bash
task goreleaser-snapshot
```

### 현재 플랫폼용 빌드
```bash
task goreleaser-build-local
```

### 실제 릴리즈 생성
1. 태그 생성 및 푸시:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. GitHub Actions가 자동으로 실행되어:
   - 모든 지원 플랫폼에 대한 바이너리 빌드
   - GitHub 릴리즈 생성
   - Docker 이미지 푸시
   - Homebrew formula 업데이트 (설정된 경우)

## 지원 플랫폼

- **Linux**: amd64, arm64, armv6, armv7
- **macOS**: amd64, arm64
- **Windows**: amd64, 386

## 프로젝트 구조

```
go-tmplx/
├── cmd/tmplx/          # 메인 애플리케이션
├── internal/app/tmplx/ # 내부 패키지
├── docs/               # 문서
├── .github/            # GitHub Actions
├── Taskfile.yml        # go-task 설정
├── .goreleaser.yml     # go-releaser 설정
└── Dockerfile          # Docker 이미지
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `task test`
5. Format code: `task format`
6. Submit a pull request

## 릴리즈 프로세스

1. 코드 변경사항을 main 브랜치에 머지
2. 새로운 버전 태그 생성:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. GitHub Actions가 자동으로 릴리즈 생성
4. 릴리즈 노트 확인 및 필요시 수정
