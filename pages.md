# GitHub Pages 설정 가이드

이 문서는 GitHub Pages의 "소스(Source)" 설정을 이해하고 올바르게 구성하는 방법을 설명합니다. 이 프로젝트의 프론트엔드는 GitHub Pages를 통해 배포됩니다.

## GitHub Pages란?

GitHub Pages는 GitHub 저장소에서 정적 웹사이트를 무료로 호스팅해주는 서비스입니다. 이를 통해 개인 프로젝트, 블로그, 문서 등을 웹에 쉽게 게시할 수 있습니다.

## "소스(Source)" 설정의 중요성

GitHub Pages의 핵심 설정 중 하나는 "소스(Source)" 설정입니다. 이 설정은 GitHub Pages가 웹사이트 콘텐츠를 빌드하고 게시하기 위해 저장소의 **어떤 브랜치와 폴더를 참조해야 하는지**를 지정합니다. 올바른 설정을 통해 GitHub Actions 워크플로우가 게시하는 결과물이 웹사이트로 서비스될 수 있습니다.

## 설정 옵션 및 선택

GitHub 저장소의 `Settings` 탭에서 `Pages` 섹션으로 이동하면 "Source"를 구성하는 옵션을 찾을 수 있습니다.

### 1. 브랜치에서 배포 (Deploy from a branch)

*   **설정 방식**: 특정 Git 브랜치(예: `main`, `gh-pages`)와 해당 브랜치 내의 특정 폴더(예: `/ (root)`, `/docs`)를 선택하여 콘텐츠를 가져옵니다.
*   **주요 특징**:
    *   **수동 또는 간접적**: 별도의 빌드 과정 없이 정적 파일만 있는 경우에 주로 사용됩니다. 웹사이트 콘텐츠가 이미 해당 브랜치/폴더에 준비되어 있다고 가정합니다.
    *   **제한적**: 복잡한 빌드 스텝이나 환경 변수 주입이 필요한 프로젝트에는 적합하지 않습니다.
*   **사용 시기**: 아주 간단한 HTML/CSS/JS 웹사이트, 또는 별도의 빌드 시스템이 이미 최종 정적 파일을 Git에 커밋하는 경우.

### 2. GitHub Actions로 배포 (GitHub Actions)

*   **설정 방식**: GitHub Actions 워크플로우를 사용하여 웹사이트 콘텐츠를 빌드하고, 빌드된 결과물(아티팩트)을 GitHub Pages에 게시하도록 설정합니다.
*   **주요 특징**:
    *   **강력한 자동화**: `frontend-deploy.yml`과 같은 GitHub Actions 워크플로우를 통해 복잡한 빌드(예: React, Vue, Angular 프로젝트 빌드), 테스트, 린팅, 타입 체크 등 모든 단계를 자동화할 수 있습니다.
    *   **동적 환경 변수**: 백엔드 API URL(`VITE_API_BASE_URL`)과 같이 빌드 시 필요한 환경 변수를 GitHub Actions의 `vars` 또는 `secrets`를 통해 안전하게 주입할 수 있습니다.
    *   **`actions/configure-pages` 및 `actions/deploy-pages`**: 이 프로젝트처럼 GitHub Pages 전용 액션을 사용하여 빌드된 아티팩트를 게시합니다.
*   **사용 시기**: 대부분의 최신 웹 프레임워크 기반 프로젝트, CI/CD 파이프라인을 통해 배포를 자동화하려는 프로젝트.

## 이 프로젝트의 권장 설정

이 `realworld-vibe-coding` 프로젝트의 프론트엔드는 `frontend-deploy.yml` GitHub Actions 워크플로우를 사용하므로, GitHub Pages 설정에서 **"GitHub Actions"를 "Source"로 선택**해야 합니다.

**설정 방법:**

1.  GitHub 저장소로 이동합니다.
2.  `Settings` 탭을 클릭합니다.
3.  좌측 사이드바에서 `Pages`를 클릭합니다.
4.  "Source" 드롭다운 메뉴에서 **"GitHub Actions"**를 선택합니다.

이 설정이 완료되면, `frontend-deploy.yml` 워크플로우가 성공적으로 실행될 때마다 빌드된 프론트엔드 애플리케이션이 자동으로 GitHub Pages에 배포됩니다.
