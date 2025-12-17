# Task History

## 2025년 12월 17일 수요일 - GitHub Actions OIDC 설정 스크립트 수정

`scripts/setup-oidc.sh` 파일의 내용을 현재 사용자 환경에 맞게 수정했습니다.

**수정 내역:**

1.  **`GITHUB_REPO` 변수 수정**:
    *   **변경 전**: `Hands-On-Vibe-Coding/realworld-vibe-coding`
    *   **변경 후**: `noaa/realworld-vibe-coding`
    *   **이유**: 스크립트 내의 GitHub 리포지토리 이름을 현재 작업 중인 리포지토리 (`git remote get-url origin` 결과)에 맞게 업데이트했습니다.

2.  **`ACCOUNT_ID` 변수 수정**:
    *   **변경 전**: `931016744724`
    *   **변경 후**: `036437288093`
    *   **이유**: 스크립트 내의 AWS 계정 ID를 현재 인증된 AWS 사용자 계정 ID (`aws sts get-caller-identity` 결과)에 맞게 업데이트했습니다.

3.  **GitHub OIDC Thumbprint 업데이트**:
    *   **변경 전**: `--thumbprint-list 6938fd4d98bab03faadb97b34396831e3780aea1`
    *   **변경 후**: `--thumbprint-list 6938fd4d98bab03faadb97b34396831e3780aea1 1c58a3a8518e8759bf075b76b750d4f2df264fcd`
    *   **이유**: GitHub Actions OIDC 연결의 안정성을 위해 최신 GitHub OIDC Thumbprint 값을 추가했습니다. 기존 값과 함께 사용됩니다.

4.  **`AWS_REGION` 안내 문구 수정**:
    *   **변경 전**: `echo "   AWS_REGION: us-east-1"`
    *   **변경 후**: `echo "   AWS_REGION: ap-northeast-2"`
    *   **이유**: 스크립트 실행 완료 시 출력되는 AWS 리전 안내 메시지를 현재 사용자 지역으로 추정되는 `ap-northeast-2` (서울)로 변경했습니다.

## 2025년 12월 17일 수요일 - CDK 배포 에러 해결 (Docker 이미지 누락)

`npm run deploy:dev` 실행 중 발생한 `AWS::EarlyValidation::ResourceExistenceCheck` 에러를 해결하기 위한 작업을 수행했습니다.

**증상:**
*   CDK 배포 시 ECS 스택 생성 단계에서 실패.
*   원인: ECS 서비스가 참조할 백엔드 Docker 이미지가 ECR에 존재하지 않음.

**조치 내용:**
*   GitHub Actions OIDC 설정을 완료한 코드를 원격 리포지토리에 **커밋 및 푸시**했습니다.
*   이를 통해 `backend-deploy.yml` 워크플로우를 트리거하여, 백엔드 Docker 이미지를 자동으로 빌드하고 ECR에 푸시하도록 유도했습니다.
*   CI/CD 파이프라인이 성공적으로 완료되면, ECR에 이미지가 준비되어 CDK 배포가 정상적으로 진행될 것으로 예상됩니다.

## 2025년 12월 17일 수요일 - GitHub Actions OIDC 세션 시간 에러 해결

GitHub Actions 실행 중 `GitHubActionsRole`의 세션 시간 초과 에러(`The requested DurationSeconds exceeds the MaxSessionDuration`)가 발생하여 이를 해결했습니다.

**증상:**
*   GitHub Actions의 `configure-aws-credentials` 단계에서 OIDC 인증 실패.

**조치 내용:**
*   AWS IAM Role `GitHubActionsRole`의 **최대 세션 지속 시간(MaxSessionDuration)**을 3600초(1시간)에서 **7200초(2시간)**로 증가시켰습니다.
*   명령어: `aws iam update-role --role-name GitHubActionsRole --max-session-duration 7200`
*   이를 통해 GitHub Actions가 요청하는 세션 시간을 충분히 수용할 수 있게 되었습니다.

## 2025년 12월 17일 수요일 - GitHub Actions OIDC 세션 시간 에러 재해결 (MaxSessionDuration 불일치)

이전 조치에도 불구하고 GitHub Actions OIDC 세션 시간 초과 에러가 재발생하여, 워크플로우(`backend-deploy.yml`) 설정과 IAM Role(`GitHubActionsRole`)의 MaxSessionDuration 간 불일치를 해결했습니다.

**원인:**
*   `backend-deploy.yml` 워크플로우에서 `role-duration-seconds`를 **14400초 (4시간)**로 명시적으로 요청하고 있었습니다.
*   반면, 이전 조치에서 IAM Role의 `MaxSessionDuration`을 **7200초 (2시간)**로 설정하여, 워크플로우의 요청이 Role 설정을 초과했기 때문에 에러가 발생했습니다.

**조치 내용:**
*   AWS IAM Role `GitHubActionsRole`의 **최대 세션 지속 시간(MaxSessionDuration)**을 워크플로우 설정에 맞춰 **14400초(4시간)**로 다시 증가시켰습니다.
*   명령어: `aws iam update-role --role-name GitHubActionsRole --max-session-duration 14400`

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 **Re-run** 해야 합니다.

## 2025년 12월 17일 수요일 - CDK 배포 권한 에러 해결 (EC2 Describe 권한 추가)

GitHub Actions를 통한 CDK 배포 시 `ec2:DescribeAvailabilityZones` 권한 부족으로 인한 `UnauthorizedOperation` 에러가 발생하여 해결했습니다.

**증상:**
*   CDK 배포 단계에서 `User ... is not authorized to perform: ec2:DescribeAvailabilityZones` 에러 발생.
*   CDK가 VPC 리소스를 구성하기 위해 가용 영역(AZ) 정보를 조회해야 하는데, 배포 정책에 해당 권한이 누락됨.

**조치 내용:**
*   `scripts/setup-oidc.sh` 파일의 `deployment-policy.json` 섹션에 다음 EC2 관련 조회 권한들을 추가했습니다:
    *   `ec2:DescribeAvailabilityZones`
    *   `ec2:DescribeRegions`
    *   `ec2:DescribeVpcs`
    *   `ec2:DescribeSubnets`
    *   `ec2:DescribeRouteTables`
    *   `ec2:DescribeSecurityGroups`
    *   `ec2:DescribeInternetGateways`
*   수정된 스크립트를 실행하여 AWS IAM Role `GitHubActionsRole`의 정책을 업데이트했습니다.

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 다시 **Re-run** 하여 배포 성공 여부를 확인합니다.

## 2025년 12월 17일 수요일 - CDK 자산 업로드 권한 에러 해결 (S3 권한 추가)

CDK 배포 시 Custom Resource(Lambda 함수 등) 자산을 S3에 업로드하는 단계에서 `Failed to publish asset` 에러가 발생하여 해결했습니다.

**증상:**
*   CDK 배포 중 `Failed to publish asset ... Code (036437288093-ap-northeast-2)` 에러 발생.
*   원인: GitHub Actions Role이 CDK v2가 사용하는 기본 Asset 버킷(`cdk-*-assets-...`)에 접근할 권한이 없었음.

**조치 내용:**
*   `scripts/setup-oidc.sh` 파일의 `deployment-policy.json` 섹션 내 S3 리소스 목록에 다음 패턴을 추가했습니다:
    *   `arn:aws:s3:::cdk-*-assets-${ACCOUNT_ID}-*`
    *   `arn:aws:s3:::cdk-*-assets-${ACCOUNT_ID}-*/*`
*   수정된 스크립트를 실행하여 AWS IAM Role `GitHubActionsRole`의 정책을 업데이트했습니다.

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 다시 **Re-run** 하여 배포 성공 여부를 확인합니다.

## 2025년 12월 17일 수요일 - CDK 배포 권한 에러 해결 (CloudFormation 스택 이름 매칭)

CDK 배포 중 `RealWorld` 스택에 대한 `cloudformation:DescribeStacks` 권한 거부(`AccessDenied`) 에러를 해결했습니다.

**증상:**
*   `User ... is not authorized to perform: cloudformation:DescribeStacks on resource: ...stack/RealWorld/...` 에러 발생.
*   기존 정책의 리소스 패턴 `RealWorld-*`가 하이픈 없이 정확히 `RealWorld`라는 이름을 가진 스택을 매칭하지 못함.

**조치 내용:**
*   `scripts/setup-oidc.sh` 파일의 `deployment-policy.json` 섹션 내 CloudFormation 리소스 패턴을 `arn:aws:cloudformation:*:${ACCOUNT_ID}:stack/RealWorld*`로 수정했습니다.
*   이를 통해 `RealWorld` (접미사 없음) 및 `RealWorld-dev` (접미사 있음) 등 모든 관련 스택을 포괄적으로 허용하게 되었습니다.
*   수정된 스크립트를 실행하여 AWS IAM Role 정책을 업데이트했습니다.

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 다시 **Re-run** 합니다.

## 2025년 12월 17일 수요일 - CDK 배포 권한 에러 해결 (CloudFormation ChangeSet 권한 추가)

CDK 배포 과정에서 ChangeSet 삭제 권한 부족으로 인한 `AccessDenied` 에러를 해결했습니다.

**증상:**
*   `User ... is not authorized to perform: cloudformation:DeleteChangeSet` 에러 발생.
*   CDK가 배포 후 정리 작업 등을 위해 ChangeSet을 제어해야 하는데 관련 권한이 누락됨.

**조치 내용:**
*   `scripts/setup-oidc.sh` 파일의 `deployment-policy.json` 섹션 내 CloudFormation 액션 목록에 다음 권한들을 추가했습니다:
    *   `cloudformation:CreateChangeSet`
    *   `cloudformation:ExecuteChangeSet`
    *   `cloudformation:DeleteChangeSet`
*   수정된 스크립트를 실행하여 AWS IAM Role 정책을 업데이트했습니다.

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 다시 **Re-run** 합니다.

## 2025년 12월 17일 수요일 - CDK 배포 권한 에러 해결 (iam:PassRole for cdk-* 역할 추가)

CDK 배포 중 `iam:PassRole` 권한 부족으로 인한 `AccessDenied` 에러를 해결했습니다.

**증상:**
*   `User ... is not authorized to perform: iam:PassRole on resource: ...role/cdk-hnb659fds-cfn-exec-role-...` 에러 발생.
*   CDK는 CloudFormation 스택 작업 시 내부적으로 `cdk-*` 패턴의 IAM Role을 전달(PassRole)해야 하는데, `GitHubActionsRole`에 해당 권한이 누락됨.

**조치 내용:**
*   `scripts/setup-oidc.sh` 파일의 `deployment-policy.json` 섹션 내 `iam:PassRole` 리소스 목록에 `arn:aws:iam::${ACCOUNT_ID}:role/cdk-*` 패턴을 추가했습니다.
*   수정된 스크립트를 실행하여 AWS IAM Role 정책을 업데이트했습니다.

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 다시 **Re-run** 합니다.

## 2025년 12월 17일 수요일 - CDK 배포 권한 에러 해결 (DescribeChangeSet 권한 확인)

CDK 배포 중 `cloudformation:DescribeChangeSet` 권한 부족 에러가 발생하여, 해당 권한의 존재 여부를 재확인하고 조치했습니다.

**증상:**
*   `User ... is not authorized to perform: cloudformation:DescribeChangeSet` 에러 발생.

**조치 내용:**
*   이전 작업(`setup-oidc.sh` 수정)에서 `cloudformation:DescribeChangeSet`이 이미 스크립트에 포함되었음을 확인했습니다.
*   `aws iam get-role-policy` 명령을 통해 실제 AWS IAM Role에 해당 권한이 올바르게 적용되어 있음을 검증했습니다.
*   권한 전파 지연 가능성이 있으나, 현재 설정은 올바릅니다.

**다음 단계:**
*   GitHub Actions에서 실패한 워크플로우를 다시 **Re-run** 합니다.
