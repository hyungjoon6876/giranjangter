# 리니지 클래식 개인거래 중개 서비스 PRD

- 프로젝트: lincle
- 문서 버전: v2.12
- 마지막 업데이트: 2026-03-14 11:02 KST
- 상태: Draft

## 변경 이력
### v2.12 (2026-03-14 11:02 KST)
- `정책 집행 결과 통지(Actionable Policy Notice) / 사용자 액션 카드 / acknowledgment canonical 계약` 섹션을 추가
- policy 판단이 단순 로그에 머물지 않고 사용자 앱/알림/지원센터/운영큐에서 동일한 `policyNoticeState`, `requiredUserActionType`, `ackRequirementMode`, `recoveryChecklistItem` vocabulary로 실행되도록 구조화
- 제재/검토/복구/재동의/추가증빙 요청 결과가 어떤 통지 카드, CTA, 확인 여부, 만료, escalation 규칙을 가져야 하는지 명문화해 화면·API·DB·운영 runbook 파생성을 보강

### v2.11 (2026-03-14 10:58 KST)
- `운영 정책 문서 파생 팩(Operational Policy Document Pack) / policyDocId·policyRuleId·enforcementSurface canonical 계약` 섹션을 추가
- PRD에서 바로 운영정책서/관리자 가이드/사용자 고지문 템플릿으로 내려갈 수 있도록 `policyDomain`, `policyAudience`, `policyEnforcementMode`, `policyOutcomeClass`, `policyCommunicationTemplate` vocabulary를 구조화
- 운영 문서, 백오피스 액션 가이드, 정책 위반 고지, QA 정책 시나리오, 감사로그가 같은 기준으로 어떤 정책이 누구에게 어떻게 적용되고 어떤 문구/조치/예외를 가져야 하는지 해석하도록 admin policy derivation 기준을 보강

### v2.10 (2026-03-14 10:55 KST)
- `거래 결과 증빙 패키지(Trade Outcome Evidence Pack) / 완료·노쇼·분쟁 판정 입력 계약` 섹션을 추가
- 완료/노쇼/분쟁/이의제기 판단 시 어떤 증빙 묶음이 필수/선택이고, 어떤 상태에서 어떤 공개범위·보존정책·다운로드 제한을 가져야 하는지 `outcomeEvidencePack`, `evidenceArtifactType`, `decisionReadinessTier`, `evidenceDisclosureScope` vocabulary로 구조화
- 사용자 제출 UX, 운영 판정 큐, DB 첨부모델, API 응답, 감사로그, analytics가 동일한 기준으로 증빙 패키지를 해석하도록 operational evidence contract를 보강

### v2.09 (2026-03-14 10:46 KST)
- `화면 전환 / 딥링크 복귀 / 로컬 상태 복원(Session Continuation) canonical 계약` 섹션을 추가
- 푸시/알림/검색/내 거래/등록 중단 후 복귀가 제각각 구현되지 않도록 `resumeIntent`, `returnContext`, `draftRecoveryKey`, `routeReentryPolicy`, `stalenessCheckMode` vocabulary를 정리
- 화면 설계, 모바일 라우팅, 로컬 저장소, API freshness check, analytics가 같은 기준으로 `어디로 돌아가고 무엇을 복구하며 언제 새로고침하는지`를 해석하도록 mobile continuation contract를 보강

### v2.08 (2026-03-14 10:43 KST)
- `화면 명세 패키지(Screen Spec Package) / routeId·moduleId·state contract` 섹션을 추가
- PRD에서 바로 화면설계서와 프론트 구현 backlog를 파생할 수 있도록 `screenFamily`, `routeId`, `moduleId`, `screenStateVariant`, `primaryCommandSurface` vocabulary를 정리
- 홈/목록/상세/등록/채팅/내 거래/프로필/운영 화면이 어떤 모듈·상태·API·analytics 묶음으로 명세화되어야 하는지 canonical screen package 기준을 보강

### v2.07 (2026-03-14 10:40 KST)
- `매물 등록 폼 schema / 단계형 입력 / validation severity canonical 계약` 섹션을 추가
- listing write UX가 자유 입력 폼이 아니라 `stepId`, `fieldGroup`, `validationSeverity`, `draftPromotionRule`, `publishReadinessState` vocabulary를 공유하도록 구조화
- 등록 화면/임시저장/API write DTO/DB draft 모델/analytics가 같은 기준으로 필수값, 경고성 입력, 검토 필요 상태, 게시 가능 조건을 해석하도록 form contract를 보강

### v2.06 (2026-03-14 10:37 KST)
- `공개 신뢰 요약(Public Trust Summary) / 프로필·매물·채팅 surface 계약` 섹션을 추가
- 완료거래/후기/응답성/검증/제재/보호조치 신호를 사용자 화면에서 어떤 공개 배지·설명·경고로 번역해야 하는지 `publicTrustTier`, `trustEvidencePill`, `trustDisclosureMode`, `trustWarningBannerType` vocabulary로 구조화
- 프로필/매물상세/목록카드/채팅상대요약/API/DB/analytics가 같은 기준으로 신뢰 요약 객체를 해석하도록 공개 범위, fallback, edge case, projection 기준을 보강

### v2.05 (2026-03-14 10:26 KST)
- `보호조치 이의제기(Protection Appeal) / 추가 증빙 제출 / 해제 결과 반영 canonical 계약` 섹션을 추가
- fraud suspicion, visibility penalty, restriction, support case가 따로 놀지 않도록 `protectionAppealState`, `appealSubmissionWindowState`, `evidenceReviewDisposition`, `protectionLiftEffectScope` vocabulary를 구조화
- 사용자 앱/알림/지원센터/운영큐/API/DB/analytics가 같은 기준으로 보호조치에 대한 설명, 재심 접수, 증빙 보완, 해제/유지/부분완화 결과를 해석하도록 recovery contract를 보강

### v2.04 (2026-03-14 10:08 KST)
- `사기 의심 리스크 신호(Fraud Suspicion Signal) / 거래 중단 보호 / 수동 검토 승격 canonical 계약` 섹션을 추가
- 외부연락처 유도, 선입금 강요, 조건 급변, 다계정/반복 패턴, 증빙 불일치 같은 리스크를 `fraudSuspicionState`, `fraudSignalFamily`, `tradeProtectionMode`, `counterpartyDisclosureLevel` vocabulary로 구조화
- 채팅/예약/당일 실행/운영큐/API/DB/analytics가 같은 기준으로 거래중단 경고, 임시 보호조치, 사용자 안내, 운영 검토 승격을 해석하도록 anti-fraud read model 기준을 보강

### v2.03 (2026-03-14 09:59 KST)
- `노출 설명(Exposure Explanation) / 사용자 복구 액션 / 운영 해제 UX canonical 계약` 섹션을 추가
- 검색 후순위·추천 제외·검토 대기 같은 `가시성 패널티`가 사용자 화면에서는 어떤 힌트/배너/복구 CTA로 번역되고, 운영자 화면에서는 어떤 해제·재계산·재심 워크플로우로 다뤄져야 하는지 `exposureDecisionState`, `userRecoverabilityTier`, `exposureHintMode`, `recoveryActionType` vocabulary로 구조화
- 검색/홈/내 매물/알림/백오피스/API/analytics가 같은 기준으로 “왜 덜 노출되는지”, “사용자가 무엇을 하면 회복되는지”, “운영이 언제 명시 제재로 승격해야 하는지”를 해석하도록 transparency와 recovery contract를 보강

### v2.02 (2026-03-14 09:56 KST)
- `가시성 패널티(Visibility Penalty) / 검색 랭킹 제한 / 조용한 제한(shadow moderation) canonical 계약` 섹션을 추가
- 명시 제재 전 단계에서 어떤 매물/계정/후기/채팅이 검색 후순위, 추천 제외, 알림 억제, 등록 검토 강화 대상으로 들어가는지 `visibilityPenaltyState`, `penaltyReasonFamily`, `exposureSurface`, `userDisclosureLevel` vocabulary로 구조화
- 검색/추천/운영큐/백오피스/API/DB/analytics가 같은 기준으로 조용한 제한과 사용자 고지 수준을 해석하도록 패널티 수명주기, 허용 액션, 해제/재심, observability 기준을 보강

### v2.01 (2026-03-14 09:49 KST)
- `당일 실행 알림 cadence / suppression / escalation canonical 계약` 섹션을 추가
- 예약 확정 이후 day-of-trade reminder, reconfirm, arrival, delay, no-show/dispute 전환이 어떤 이벤트/시간창/억제 규칙/운영 에스컬레이션으로 알림화되어야 하는지 구조화
- 알림센터 IA, push/inbox fanout, scheduler/job, notification projection, 운영 runbook으로 직접 파생 가능한 execution notification vocabulary를 보강

### v2.00 (2026-03-14 09:44 KST)
- `운영 큐 taxonomy / 작업 소유권 / SLA 에스컬레이션 canonical 계약` 섹션을 추가
- Report, ModerationCase, SupportCase, No-show/Dispute, Restriction review가 어떤 운영 큐로 들어가고 누구의 작업으로 간주되며 어떤 상태/SLA/에스컬레이션 규칙을 공유해야 하는지 구조화
- 백오피스 IA, 운영 runbook, queue projection, admin API, staffing/모니터링 문서로 직접 파생 가능한 queue vocabulary와 action gating 기준을 보강

### v1.99 (2026-03-14 09:33 KST)
- `거래 실행 projection 응답 계약 / 내 거래·거래상세·당일 실행 surface canonical payload` 하위 섹션을 추가
- `executionReadinessSnapshot`, `mutualExecutionAck`, `rescheduleDecisionState`, `claimSuppressionReasonCode`가 `GET /me/trades`, `GET /me/trades/{tradeThreadId}`, 당일 실행 카드에서 어떤 필드/정렬/CTA gating으로 노출돼야 하는지 구체화
- 모바일 화면, projection 스키마, OpenAPI 응답 모델, QA fixture가 같은 payload shape를 바라보도록 response block과 fallback 규칙을 보강

### v1.98 (2026-03-14 09:30 KST)
- `재일정 승인 상태 / 당일 실행 재확인 / 노쇼·분쟁 action bridge 계약` 섹션을 추가
- accepted/rejected/expired reschedule이 당일 실행 카드, readiness snapshot, no-show claim 생성 가능 여부, dispute linkage를 어떤 공통 action code와 suppression reason으로 해석해야 하는지 구체화
- 화면 CTA 문구군, API command surface, DB/action history, 운영 adjudication 기준으로 직접 파생 가능한 execution-decision 기준을 보강

### v1.97 (2026-03-14 09:27 KST)
- `정책 위반군(Violation Family) / 기본 집행 레벨 / 증빙 기준 / 사용자 커뮤니케이션 계약` 섹션을 추가
- 허위매물, 외부연락처 유도, 노쇼, 괴롭힘, 사기 의심, 리뷰 악용, 다계정/스팸성 행위 등 주요 위반군별로 어떤 증빙이면 경고/제한/정지로 에스컬레이션되는지 기본 매트릭스를 정리
- 운영자 교육자료, 백오피스 액션 가이드, 제재 정책 문서, 사용자 통지 템플릿으로 직접 파생 가능한 기준을 보강

### v1.96 (2026-03-14 09:23 KST)
- `Starter DDL / 관계형 테이블 책임 / FK·유니크·삭제 정책` 섹션을 추가
- PRD에서 실제 초기 스키마 문서로 바로 파생할 수 있도록 aggregate별 write table, read model, soft delete, FK on delete, 유니크/인덱스/보관 전략을 한 번 더 정리
- Listing/TradeThread/Reservation/Completion/Moderation 계열이 어떤 테이블 경계와 불변식을 가져야 하는지 명문화해 DB 설계와 서버 구현의 공통 기준을 보강

### v1.95 (2026-03-14 09:11 KST)
- `거래 실행 상호확인(Mutual Execution Ack) / 명시 인정·암묵 인정·최종 출발 신호 canonical 계약` 섹션을 추가
- 예약은 성립했지만 상대가 실제로 같은 조건을 마지막으로 인지했는지 모호한 문제를 줄이기 위해 `mutualExecutionAckState`, `ackEvidenceType`, `ackTriggerMode`, `departureReadinessState` vocabulary를 정리
- 채팅/내 거래/당일 실행 카드/API/DB/운영이 같은 기준으로 `도착`, `확인했어요`, `변경 내용 봤어요`, `노쇼 claim 억제/허용`, `완료 가능 조건`을 판단하도록 명문화

### v1.94 (2026-03-14 08:57 KST)
- `거래 실행 준비도(Execution Readiness Snapshot) / 예약 성립과 실제 실행 가능성 구분 canonical 계약` 섹션을 추가
- 시간 슬롯, 장소, 식별 정보, 합의 조건, 정산 방식, 취소 의사, 용량 상태가 흩어진 채 해석되지 않도록 `executionReadinessState`, `blockingPrerequisiteType`, `readinessDriftReasonCode`, `finalGoSignalSource` vocabulary를 정리
- 채팅/내 거래/당일 실행 카드/API/DB/운영이 같은 준비도 snapshot을 기준으로 CTA 노출, reminder, no-show 억제, 완료 가능 조건을 판단하도록 명문화

### v1.93 (2026-03-14 08:50 KST)
- `거래 취소 의사(Cancel Intent) / 상호 취소·일방 취소·재개 가능성 canonical 계약` 섹션을 추가
- 예약 이후 완료 전 단계에서 발생하는 취소를 `cancelIntentState`, `cancelResolutionType`, `reopenDisposition`, `counterpartyImpactTier` vocabulary로 구조화
- 채팅/내 거래/알림/API/DB/운영이 같은 기준으로 취소 제안, 승인/거절, 자동 만료, available 복귀, no-show·분쟁과의 경계 해석을 공유하도록 명문화

### v1.92 (2026-03-14 08:47 KST)
- `동시 거래 수용량(Capacity) / overcommit 방지 / active commitment canonical 계약` 섹션을 추가
- 여러 문의/리드/예약/당일 실행이 겹칠 때 화면·알림·상태머신·운영이 같은 기준으로 `capacitySnapshot`, `capacityState`, `commitmentLoadTier`, `overcommitRiskDisposition`를 해석하도록 vocabulary를 정리
- 매물 상세/내 거래/예약/API/DB/analytics가 동일한 기준으로 활성 예약 상한, 시간 중복, 응답 여력, 과잉 약속 감지와 완화 규칙을 해석하도록 명문화

### v1.91 (2026-03-14 08:38 KST)
- `거래 가능 시간 Slot / 예약 제안 가능 범위 / availability-to-booking canonical 계약` 섹션을 추가
- 프로필 가용시간, 매물 가용시간 스냅샷, 예약 제안, 재일정, 노쇼 분쟁이 서로 다른 시간 해석을 쓰지 않도록 `bookingAvailabilitySnapshot`, `bookabilityState`, `slotDisclosureLevel`, `bookingConflictDisposition` vocabulary를 정리
- 홈/상세/채팅/예약/API/DB/analytics가 동일한 시간 슬롯 기준으로 예약 가능 여부, 최소 리드타임, 버퍼, blackout, 확인 필요 상태를 해석하도록 명문화

### v1.90 (2026-03-14 08:34 KST)
- `거래 단위(Unit of Trade) / 묶음·개당가·총액 canonical 계약` 섹션을 추가
- stackable/non-stackable, package listing, partial trade, offer/fixed pricing이 뒤섞여도 화면·검색·채팅·완료·정산·운영이 같은 `tradeUnitSnapshot` 해석을 쓰도록 vocabulary를 정리
- 매물 등록/상세/API/DB/analytics가 동일한 기준으로 `unitBasis`, `priceBasis`, `bundleRule`, `residualUnitPolicy`를 해석하도록 거래 단위/가격 단위/잔여 수량 규칙을 명문화

### v1.89 (2026-03-14 08:24 KST)
- `매물 품질 스냅샷(Listing Quality Snapshot) / 등록 완성도·검색 품질·운영 개입 canonical 계약` 섹션을 추가
- 단순 랭킹 점수가 아니라 등록 폼, 검색 노출, 운영 검토, 사용자 수정 유도를 함께 묶는 `listingQualitySnapshot` vocabulary와 상태 전이를 정리
- 홈/목록/상세/등록 UX/API/DB/analytics가 동일한 품질 해석을 쓰도록 completeness, clarity, freshness, policy, readiness 기준과 파생 규칙을 명문화

### v1.88 (2026-03-14 08:21 KST)
- `거래 작업 기한 / 응답 마감 / action deadline snapshot canonical 계약` 섹션을 추가
- 문의 응답, 예약 응답, 위치 재확인, 완료 확인, 노쇼 소명, 운영 추가자료 요청 등 시간제한 액션을 `actionDeadlineSnapshot` 하나로 묶어 화면/알림/API/운영이 같은 기한 해석을 사용하도록 정리
- 홈/내 거래/채팅/알림/운영큐 우선순위와 `deadlineState`, `urgencyTier`, `expiryConsequenceType`, `autoTransitionPolicy` vocabulary를 명문화

### v1.87 (2026-03-14 08:15 KST)
- `후기 자격 스냅샷(Review Eligibility Snapshot) / 작성 가능 창 / 봉인(freeze) canonical 계약` 섹션을 추가
- 완료/자동확정/노쇼/분쟁/제재/후기 숨김 재검토가 섞여도 화면·API·운영이 같은 기준으로 후기 작성 가능 여부를 해석하도록 `reviewEligibilityState`, `reviewWindowState`, `publicationFreezeReasonCode`, `reviewMutabilityState` vocabulary를 정리
- 프로필/내 거래/최종 결과/운영 큐/analytics가 동일한 review gate 객체를 바라보도록 생성 조건, 만료, 봉인 해제, 수정 가능 범위, DB/API 파생 기준을 명문화

### v1.86 (2026-03-14 08:06 KST)
- `제한 해제 / 재활성화 / probation(관찰기간) canonical 계약` 섹션을 추가
- 경고/일시제한/정지 이후 사용자가 어떤 조건으로 기능을 회복하는지, appeal 승인과 auto-expire 해제가 어떻게 다른지, trust/restriction/support/notification 화면이 동일한 해석을 쓰도록 `restrictionRecoveryState`, `liftDecisionType`, `probationTier`, `recoveryRequirementType` vocabulary를 정리
- 운영 해제, 자동 해제, 조건부 복귀, 재위반 시 재에스컬레이션, 사용자 노출 카피, API/DB/analytics 파생 기준을 명문화

### v1.85 (2026-03-14 08:03 KST)
- `콘텐츠 검수 케이스(Content Moderation Review Case) / 매물·채팅·후기·프로필 정책판정 canonical 계약` 섹션을 추가
- 자동 탐지, 사용자 수정 유도, 운영 검토, 임시숨김, 복구, strike 반영이 객체마다 다르게 해석되지 않도록 `contentModerationCase` vocabulary와 상태 전이를 정리
- 등록 UX, 운영 큐, API/DB, audit/analytics가 동일한 검수 사건 단위를 바라보도록 review trigger, evidence bundle, decision visibility, reopen/appeal 기준을 명문화

### v1.84 (2026-03-14 07:00 KST)
- `거래 신뢰 스냅샷(Trade Trust Snapshot) / 공개 신뢰·내부 리스크·후기 게이트 canonical 계약` 섹션을 추가
- 완료/노쇼/분쟁/재거래/제재/검증 상태가 프로필/상세/내 거래/운영도구에서 서로 다르게 해석되지 않도록 `tradeTrustSnapshot` vocabulary와 계산 규칙을 정리
- 공개 신뢰 신호, 내부 리스크 반영, 후기 가능 여부, safety 배너, API/DB/analytics 파생 기준을 하나의 canonical read model 기준으로 명문화

### v1.83 (2026-03-14 06:43 KST)
- `거래 사건 요약(Trade Case Summary) / 사용자·운영·지원용 canonical read model 계약` 섹션을 추가
- listing/chat/reservation/reschedule/no-show/dispute/final outcome/review eligibility를 하나의 `tradeCaseSummary` 단위로 수렴시키는 vocabulary와 계산 규칙을 정리
- 내 거래/상세/지원 케이스/운영 큐/API/DB/analytics가 같은 사건 요약 객체를 바라보도록 결정 근거, action-needed, evidence linkage, 보존/비공개 규칙을 명문화

### v1.82 (2026-03-14 06:36 KST)
- `재일정-노쇼-최종결과 연결 판정(Re-schedule to No-show Outcome Bridge) 계약` 섹션을 추가
- accepted/rejected/expired reschedule이 no-show claim, dispute, final outcome으로 어떻게 이어지는지 `executionDecisionState`, `rescheduleNoShowBridgeType`, `claimSuppressionReasonCode`, `outcomeCandidateSource` vocabulary로 구조화
- 당일 실행 카드, no-show 사건 생성, 결과 확정, 후기 게이트, trust/analytics가 동일한 판정 단위를 바라보도록 화면/API/DB/운영 기준을 명문화

### v1.81 (2026-03-14 06:27 KST)
- `최종 거래 결과(Final Trade Outcome) / 완료·취소·부분완료·분쟁 종결 canonical 계약` 섹션을 추가
- no-show, deal terms, exchange confirmation, completion dispute가 결국 어떤 `최종 결과`로 닫히는지 `finalOutcomeType`, `outcomeClosureMode`, `reviewEligibilityOutcome`, `trustImpactMode` vocabulary로 구조화
- 후기 공개, 신뢰 집계, 잔여 수량 처리, 운영 결과 통지, analytics가 동일한 종결 단위를 바라보도록 결과 판정 매트릭스와 API/DB/운영 기준을 명문화

### v1.80 (2026-03-14 06:22 KST)
- `거래 인도/수령 확인(Exchange Confirmation) / 실제 물건 전달·수령 확인 계약` 섹션을 추가
- 거래 직전 합의된 조건(`Deal Terms Snapshot`)이 실제 전달/수령 행위로 어떻게 닫히는지 `handoverState`, `receiptState`, `exchangeConfirmationState`, `exchangeMismatchType` vocabulary로 구조화
- 채팅/당일 실행/완료 확인/분쟁/API/DB/운영이 같은 기준으로 인도·수령 체크, 부분 인도, 현장 불일치, 완료 요청 가능 조건을 해석하도록 명문화

### v1.79 (2026-03-14 06:14 KST)
- `거래 조건 확정(Deal Terms Snapshot) / 최종 합의 조건 / 변경 이력 계약` 섹션을 추가
- 매물 원문, 가격 제안, 예약, 식별 정보가 실제 성사 직전 어떤 `최종 거래 조건`으로 수렴하는지 `dealTermsState`, `termsSourceType`, `termsChangeImpact`, `executionReadinessState` vocabulary로 구조화
- 채팅/예약/당일 실행/완료/분쟁/API/DB/운영이 동일한 조건 스냅샷을 기준으로 가격·수량·장소·식별자·지급방식 변경과 동의 여부를 해석하도록 명문화

### v1.78 (2026-03-14 06:11 KST)
- `거래 상대 식별 정보(Trade Counterparty Identity Handshake) / 캐릭터명·접선 식별자 공개/확인 계약` 섹션을 추가
- 실제 거래 직전 필요한 상대 식별 정보를 `counterpartyIdentifierType`, `identityDisclosureLevel`, `identifierVerificationState`, `tradeRecognitionMethod` vocabulary로 구조화
- 매물/채팅/예약/당일 실행 카드/노쇼 분쟁/API/DB/운영이 동일한 기준으로 식별 정보 공개 범위, 확인 상태, 변경 이력, 사칭·혼선 대응을 해석하도록 명문화

### v1.77 (2026-03-14 05:59 KST)
- `상대 사용자 관계(User Relationship Edge) / 차단·단골·위험 플래그 read model 계약` 섹션을 추가
- 거래 상대 간 관계를 단순 block 여부가 아니라 `relationshipState`, `safetyInteractionState`, `repeatTradeAffinityState`, `contactPolicyState` vocabulary로 구조화
- 매물 상세/채팅/내 거래/프로필/운영도구/API/DB가 동일한 관계 요약 객체를 공유하도록 CTA gating, 공개 범위, 운영 개입, analytics 기준을 명문화

### v1.76 (2026-03-14 05:54 KST)
- `사용자 문의/지원 케이스(Support Case) / 운영 커뮤니케이션 계약` 섹션을 추가
- 단순 FAQ 문의, 신고 후속 문의, 제재 이의제기, 분쟁 보완 요청이 DM/이메일/채팅/운영 메모로 흩어지지 않도록 `supportCaseType`, `caseOriginSurface`, `resolutionCommunicationMode`, `userExpectationState` vocabulary로 구조화
- Help/내 문의/신고 상세/분쟁 상세/운영 백오피스/API/DB가 같은 케이스 단위를 공유하도록 사용자 커뮤니케이션 수명주기, SLA, 상태전이, 메시지 템플릿, 분석 기준을 명문화

### v1.75 (2026-03-14 05:50 KST)
- `운영 공지 / 시스템 배너 / 정책 메시지 노출 계약` 섹션을 추가
- 긴급 점검, 정책 변경, 기능 제한, 거래 안전 경고를 단순 텍스트 공지로 흩어두지 않도록 `noticeAudienceScope`, `noticeSurface`, `noticeSeverity`, `actionRequirementLevel`, `ackRequirementType` vocabulary로 구조화
- 홈/목록/상세/채팅/내 거래/알림/API/DB/운영도구가 같은 기준으로 공지 우선순위, 표시 조건, 재노출, 확인 필요 여부를 해석하도록 명문화

### v1.74 (2026-03-14 05:43 KST)
- `재거래 시작(Start Repeat Trade) / 과거 맥락 승계 / 새 스레드 생성 계약` 섹션을 추가
- 완료/취소/보관된 거래에서 새 거래를 시작할 때 기존 thread를 되살리지 않고 `repeatTradeSourceType`, `contextCarryForwardPolicy`, `carryForwardDisclosureLevel`, `repeatTradeSpawnMode` vocabulary로 새 thread 생성 규칙을 구조화
- 홈/프로필/보관함/채팅/내 거래/API/DB/analytics가 같은 기준으로 재거래 CTA, 과거 후기·제재·차단 가드, 이전 약속 정보 승계 범위를 해석하도록 명문화

### v1.73 (2026-03-14 05:40 KST)
- `거래 스레드 종료 / 아카이브 / 재진입 계약` 섹션을 추가
- 활성 거래 워크스페이스와 기록 보관함의 경계를 `threadClosureReason`, `archiveVisibilityLevel`, `reopenEligibilityState`, `threadArchiveTrigger` vocabulary로 구조화
- 채팅목록/내 거래/알림/API/DB/운영이 동일한 기준으로 스레드 종료, 보관, 재진입, 재거래 CTA를 해석하도록 화면 동작·명령 surface·분석 파생 기준을 명문화

### v1.72 (2026-03-14 05:37 KST)
- `노쇼 사건(No-show Case) canonical 객체 / 화면 표시 / API·DB 파생 기준` 하위 섹션을 추가
- no-show claim, counter-claim, evidence, adjudication, appeal이 여러 객체로 흩어지지 않도록 `noShowCase`, `casePartyRole`, `caseOutcomeType`, `caseTimelineEventType`, `evidenceBundleScope` vocabulary를 정리
- 내 거래/채팅/운영큐/API/DB가 같은 사건 단위를 공유하도록 사건 생성 단위, 필수 필드, 읽기 projection, command surface, 상태 전이, analytics 파생 기준을 명문화

### v1.71 (2026-03-14 05:28 KST)
- `거래 당일 실행 카드 CTA family / 상태카피 / action gating 계약` 하위 섹션을 추가
- 당일 실행 카드가 `도착/지연/재일정/노쇼/완료`를 서로 다른 의미의 CTA로 분리하고, 상황별로 어떤 버튼과 경고문구를 보여야 하는지 `dayOfTradeCardMode`, `executionCTAType`, `ctaAvailabilityReasonCode`, `checkinConfidenceTier` vocabulary로 구조화
- 채팅/내 거래/알림/운영/analytics가 같은 action family를 기준으로 당일 실행 UX를 해석하도록 CTA 노출 우선순위, disabled 사유, copy tone, API/DB/분석 파생 기준을 명문화

### v1.70 (2026-03-14 05:25 KST)
- `노쇼 판정 이후 appeal / 재심 / 신뢰도 재계산 계약` 하위 섹션을 추가
- no-show 사건이 판정된 뒤 appeal 생성 가능 조건, 재심 입력 요건, 원판정 유지/수정/철회 시 후기 게이트·restriction candidate·trust aggregate를 어떻게 재계산하는지 `appealState`, `recomputationScope`, `resolutionRevisionType`, `trustReconciliationMode` vocabulary로 구조화
- 운영 큐, 사용자 결과 화면, API, analytics, 배치 재집계가 같은 사건 revision 체계를 공유하도록 판정 변경의 수명주기와 side-effect freeze/replay 원칙을 명문화

### v1.69 (2026-03-14 05:18 KST)
- `노쇼 case evidence sufficiency / adjudication SLA / 사용자 결과 노출 계약` 하위 섹션을 추가
- 노쇼 사건에서 운영자가 언제 추가 증빙을 요청하고 언제 `insufficient_evidence`로 닫는지, 어떤 증빙 조합이 최소 판정 입력으로 간주되는지 `evidenceSufficiencyTier`, `adjudicationSlaTier`, `userOutcomeVisibility` vocabulary로 구조화
- 판정 결과가 내 거래/후기/제한/이의제기/분석에 어떻게 같은 사건 단위로 노출되는지 명문화

### v1.68 (2026-03-14 05:16 KST)
- `노쇼 claim adjudication / case linkage / 후기·제한 반영 계약` 섹션을 추가
- 노쇼 claim이 생성된 뒤 Report, Dispute, Restriction, Review publication, analytics가 서로 다른 사건 키를 쓰지 않도록 `caseLinkPolicy`, `claimAdjudicationState`, `reviewPublicationGate`, `restrictionEscalationTier` vocabulary를 정리
- 운영 큐/백오피스/API/DB/화면이 같은 결정 단위를 공유하도록 claim 단위 병합, 단건 vs 양측 claim 해석, 판정 이후 후기 공개/제한/알림 반영 순서를 명문화

### v1.67 (2026-03-14 05:09 KST)
- `노쇼 claim / 증빙 / 분쟁·신고 연결 계약` 섹션을 추가
- 당일 실행 이후 한쪽이 `no_show`를 주장할 때 claim, counter-claim, evidence, moderation decision, 후기/신뢰도 반영이 서로 다른 기준으로 흩어지지 않도록 `noShowClaimState`, `arrivalEvidenceDisposition`, `noShowResolutionType`, `claimLinkagePolicy` vocabulary를 정리
- 채팅/내 거래/API/DB/운영 runbook이 같은 사건 단위를 공유하도록 no-show claim의 생성 조건, 중복 제한, accepted reschedule와의 관계, Report·Dispute·Restriction 연결 기준을 명문화

### v1.66 (2026-03-14 05:02 KST)
- `부분거래 / 잔여 수량 / 수량 할당 계약` 섹션을 추가
- 하나의 매물에서 일부 수량만 먼저 거래되는 상황을 `listingUnitState`, `allocationState`, `partialCompletionReasonCode`, `residualHandlingPolicy` vocabulary로 구조화
- 매물 상태, 채팅 우선상대, 완료/후기, 검색 노출, API/DB/운영이 부분거래를 서로 다르게 해석하지 않도록 잔여 수량 감소·재오픈·분쟁 연결 기준을 명문화

### v1.65 (2026-03-14 03:58 KST)
- `예약 재일정(Reschedule) / 시간·장소 변경 협상 계약` 섹션을 추가
- 예약 확정 이후 시간/장소 변경이 단순 텍스트 대화가 아니라 별도 상태 머신과 승인 흐름을 가지도록 `rescheduleState`, `rescheduleScope`, `rescheduleInitiatorType`, `rescheduleFailureReasonCode` vocabulary를 정리
- 당일 실행, 노쇼 판단, 알림, 내 거래 우선순위, Dispute/운영 검토까지 같은 기준으로 재일정 이벤트를 해석할 수 있도록 화면/API/DB/운영 파생 기준을 명문화

### v1.64 (2026-03-14 03:25 KST)
- `대금 전달 / 정산 방식 / 결제 증빙 / 위험 가드레일 계약` 섹션을 추가
- 에스크로 없이도 거래 직전/직후 어떤 대금 전달 방식이 허용되고 어떤 표현/행동이 고위험 신호인지 정리해 chat/reservation/completion/moderation 파생 기준을 보강
- `settlementMethod`, `paymentExecutionState`, `paymentEvidenceType`, `paymentRiskSignalCode` vocabulary를 추가해 화면/API/DB/운영 정책이 같은 언어를 쓰도록 명문화

### v1.63 (2026-03-14 03:15 KST)
- `약관/정책 동의 버전 / 재동의 / 기능 게이트 계약` 섹션을 추가
- 회원가입 이후 실제 거래 가능 상태를 정책 동의 이력과 연결해 `policyDocumentType`, `policyAcceptanceState`, `reConsentTriggerCode`, `policyGateScope` vocabulary를 정리
- 온보딩/거래 쓰기 API/운영 백오피스/감사로그/배치가 같은 기준으로 약관 재동의와 기능 차단을 해석할 수 있도록 동의 저장 단위, 재고지 UX, 시행 시점, 예외 규칙을 명문화

### v1.62 (2026-03-14 03:11 KST)
- `거래 스레드(Trade Thread) canonical read model / 작업 우선순위 계약` 섹션을 추가
- 홈/내 거래/채팅목록/알림이 서로 다른 기준으로 같은 거래를 표현하지 않도록 `tradeThreadState`, `threadUrgencyTier`, `nextBestAction`, `counterpartyCommitmentSignal` vocabulary를 정리
- screen/API/read model/analytics가 공통으로 사용할 요약 카드 필드, 정렬 우선순위, 상태 계산 원칙, edge case를 명문화

### v1.61 (2026-03-14 03:07 KST)
- `거래 당일 실행 / 도착 확인 / 노쇼 판단 계약` 섹션을 추가
- 예약 확정 이후 실제 만남 직전/직후에 필요한 `arrivalState`, `meetingExecutionState`, `noShowDecisionState`, `arrivalEvidenceType` vocabulary를 정의해 거래 성사 직전 UX와 분쟁/운영 기준을 보강
- 내 거래/채팅/알림/API/DB/운영 runbook이 같은 실행 단계 언어를 쓰도록 당일 체크인, grace period, 자동 리마인드, 노쇼 신고 창구, 상호 불일치 처리 규칙을 명문화

### v1.60 (2026-03-14 02:12 KST)
- `가정(Assumption) / 미확정 결정 / 결정 필요도(Decision Priority) 레지스터` 섹션을 추가
- 문서 전반에 흩어져 있던 `정책 결정 필요`, `가정`, `추후 확정` 성격의 항목을 제품/운영/데이터/API 관점의 결정 레지스터로 구조화
- 어떤 항목이 MVP blocker인지, 어떤 문서 산출물(화면/API/DB/운영정책)에 직접 영향을 주는지 추적할 수 있도록 `decisionArea`, `decisionPriority`, `targetArtifact`, `defaultAssumptionAction` 후보를 명문화

### v1.59 (2026-03-13 23:56 KST)
- `재거래(Repeat Trade) / 다시 거래하기 / 단골 상대 계약` 섹션을 추가
- 완료 거래 이후 동일 상대와의 반복 거래를 별도 `repeat-trade` 문맥으로 구조화해 프로필/채팅/홈/신뢰 정책 파생 기준을 보강
- `repeatTradeAffinityLevel`, `repeatTradeEligibilityState`, `retradeCTAType`, `counterpartyPreferenceSignal` 후보를 명문화

### v1.58 (2026-03-13 21:55 KST)
- `신규 사용자 첫 거래 활성화(Activation) / 첫 행동 미션 / 전환 보호 계약` 섹션을 추가
- 회원가입 직후 사용자를 `거래가능` 상태로만 보지 않고 `첫 매물 등록`, `첫 문의 응답`, `첫 예약 확정`, `첫 완료`까지의 activation 단계로 구조화해 홈/온보딩/알림/신뢰 정책 파생 기준을 보강
- `activationStage`, `activationBlockerCode`, `firstTradeMissionType`, `activationNudgePolicy` 후보를 명문화

### v1.57 (2026-03-13 20:52 KST)
- `신원확인(Verification) / 신뢰 레벨 / 제한 해제 계약` 섹션을 추가
- 게스트→회원→거래가능 사용자 사이의 신뢰 부트스트랩을 `verificationLevel`, `tradeEligibilityState`, `trustSignalLevel`, `restrictionLiftPolicy` vocabulary로 구조화해 프로필/API/운영정책/온보딩 문서 파생 기준을 보강
- 신고 다발 신규계정, 고위험 매물, 추가 인증 요구, 공개 프로필 노출 범위와의 연결 규칙을 명문화

### v1.56 (2026-03-13 19:42 KST)
- `증빙 다운로드 / 반출 워크플로우 상태 계약` 섹션을 추가
- preview/download/export 요청이 어떤 승인 단계와 토큰 수명주기를 거쳐 실행되는지, 어떤 경우 즉시 허용/승인대기/만료/회수되는지 정리해 Evidence API·백오피스·스토리지 문서 파생 기준을 보강
- `evidenceAccessRequestState`, `exportBundleState`, `downloadGrantScope`, `evidenceTokenRevocationReason` 후보를 명문화

### v1.55 (2026-03-13 19:35 KST)
- 모바일 중심 거래 흐름의 이탈/중단 복구 기준을 명문화하기 위해 `드래프트 / 자동저장 / 오프라인 복구 계약` 섹션을 추가
- 매물 작성 draft, 채팅 입력 draft, 예약 제안 draft를 어떤 단위로 저장하고 언제 복구/폐기/동기화할지 정리해 화면명세·로컬 저장소·API·DB 파생 기준을 보강
- `draftType`, `draftScope`, `draftRecoveryPolicy`, `conflictResolutionPolicy` 후보를 명문화

### v1.54 (2026-03-13 19:28 KST)
- `알림 전달 파이프라인 / suppression / retry / fallback 계약` 섹션을 추가
- in-app/push fanout, dedup, suppression, retry, provider callback, deep link fallback, 장애 시 degraded mode 원칙을 정리해 알림 인프라·운영 runbook·API/DB 파생 기준을 보강
- `notificationDispatchState`, `deliveryAttemptState`, `suppressionReasonCode`, `deliveryDegradedMode` 후보를 명문화

### v1.53 (2026-03-13 19:20 KST)
- 알림함을 단순 이벤트 리스트가 아니라 거래 재개용 작업 큐로 정의하는 `알림 인박스(Notification Inbox) 상태 / 액션 / 보존 계약` 섹션을 추가
- unread/read/acted/dismissed/expired 상태, notificationActionType, badge 계산, CTA 우선순위, 만료/자동정리 규칙을 정리해 화면명세·API·운영 문서 파생 기준을 보강
- 예약/완료/분쟁/운영 알림이 같은 vocabulary를 쓰도록 `notificationState`, `actionState`, `inboxPriorityTier`, `threadMergeKey` 후보를 명문화

### v1.52 (2026-03-13 19:13 KST)
- 분산돼 있던 정책/검증 규칙을 런타임 평가 관점으로 묶기 위해 `정책 평가 엔진 / 결정 결과 / 노출 계약` 섹션을 추가
- 인증/권한/상태/콘텐츠/안티어뷰즈/운영 제한을 어떤 순서로 평가하고, `allow/warn/block/review_required`를 어떻게 API·화면·감사로그에 공통 반영할지 구조화
- policy evaluation trace, decision source, userFacingPolicyHint, admin override 경계를 명문화해 기능명세·OpenAPI·정책엔진·백오피스 파생 기준을 보강

### v1.51 (2026-03-13 19:08 KST)
- 운영 정책/백오피스/API 인가 문서로 직접 파생할 수 있도록 `운영 액션 코드 / Admin RBAC 매트릭스 / 승인 체인 계약` 섹션을 추가
- 역할별 허용 액션, 2인 승인 필요 액션, 자기사건 처리 금지, break-glass, 감사 로그 필수 필드를 한 표로 정리해 admin 실행 규칙을 구체화
- actionCode, permission bundle, approvalPolicy, case assignment vocabulary를 통일해 추후 OpenAPI·백오피스·운영 runbook 작성 기준을 보강

### v1.50 (2026-03-13 19:02 KST)
- 신고/분쟁/운영 문서로 바로 파생할 수 있도록 `증빙(Evidence) 접근등급 / 다운로드 / 외부반출 통제 정책` 섹션을 추가
- 일반 채팅 첨부와 분쟁 증빙 첨부의 권한·워터마크·원본 다운로드·감사로그 기준을 분리하고, 사용자/운영자/API가 같은 vocabulary를 쓰도록 정리
- evidence access level, download disposition, share token, redaction, export approval 기준을 명문화해 백오피스/스토리지/보안 설정 문서의 공통 기준을 보강

### v1.49 (2026-03-13 18:55 KST)
- PRD에서 바로 OpenAPI/DB 설계 산출물을 뽑아낼 수 있도록 `MVP OpenAPI 산출물 팩 / Starter DDL 파생 기준` 섹션을 추가
- endpoint group별 필수 명세 항목, 공통 schema, typed error, enum freeze 범위와 starter migration에 반드시 포함할 테이블/제약/인덱스를 정리
- 다음 단계에서 실제 `openapi.yaml`/초기 migration을 작성할 때 무엇이 blocker 없이 채워져야 하는지 acceptance 기준을 보강

### v1.48 (2026-03-13 18:41 KST)
- 운영정책/DB/백오피스 문서로 직접 파생할 수 있도록 `데이터 보존(Retention) / 접근등급 / 증빙 라이프사이클 계약` 섹션을 추가
- 객체별 기본 보관기간, 공개 비노출 시점, 법적/운영 hold, 증빙 첨부 다운로드 통제, 배치 파기 원칙을 하나의 매트릭스로 정리
- 메시지/신고/분쟁/첨부/감사로그가 같은 retention vocabulary를 쓰도록 `retentionClass`, `legalHold`, `purgeState`, `accessLevel` 후보를 명문화

### v1.47 (2026-03-13 18:31 KST)
- DB enum, OpenAPI schema, 프론트 타입 정의가 같은 기준을 쓰도록 `핵심 enum / 코드셋 카탈로그 계약` 섹션을 추가
- 공개 상태값과 내부 상태값을 분리해 어떤 enum이 외부 API에 노출되고 어떤 값이 내부 운영/이벤트/DDL 전용인지 명문화
- 코드셋 버전관리, 미지원/unknown fallback, 다국어 라벨 분리 원칙을 정리해 다음 단계(OpenAPI/DDL/TypeScript 타입)로 직접 파생 가능한 기준을 보강

### v1.46 (2026-03-13 18:24 KST)
- 화면 버튼/CTA, API `availableActions`, 인가 정책, 감사 로그가 같은 언어를 쓰도록 `공개 액션 코드(Action Code) / 권한 번들 계약` 섹션을 추가
- 제품/디자인/모바일/서버/운영/QA가 같은 승인 체크포인트를 쓰도록 `오너십 / 승인 매트릭스 / 릴리즈 산출물` 섹션을 추가
- 남은 구현 산출물을 실제 문서·결정 단위로 재정리해 다음 단계(OpenAPI/DDL/RBAC 표/상태도)로 مباشرة 연결되도록 보강

### v1.45 (2026-03-13 18:12 KST)
- 거래 성사율과 판매자/구매자 응답 피로 관리를 직접 다루는 `문의 SLA / 거래 스레드 Aging / 자동 휴면 정리 정책` 섹션을 추가
- 문의 생성 이후 응답 지연, 휴면 스레드, 무의미한 장기 열린 대화, 재활성화 조건을 구조화해 내 거래/채팅/알림/운영 큐에서 동일 기준을 쓰도록 보강
- `TradeThreadAging`, `responseDueAt`, `staleReasonCode`, 관련 API/분석/운영 가드레일을 정리해 화면명세·DB·운영정책으로 직접 파생 가능한 기준을 추가

### v1.44 (2026-03-13 17:10 KST)
- 판매자가 여러 문의를 어떻게 정리하고 우선 거래 상대를 어떻게 선택/교체/종결하는지 다루는 `거래 후보 상대(Lead) / 우선상대 선정 / 대기열 계약` 섹션을 추가
- `reserved` 상태의 실제 운영 단위인 후보군, waiting 대화, 우선상대 승격, 응답기한, 일괄 안내, 차단/신고/랭킹 연동 기준을 구조화
- 내 거래/내 매물/채팅/DB/API/운영 큐로 직접 파생 가능한 `TradeLead`, `leadStatus`, 관련 API·분석 이벤트·가드레일을 보강

### v1.43 (2026-03-13 16:52 KST)
- 재방문 탐색과 거래 재개를 연결하는 `찜(Favorite) / 관심목록 / 상태복귀 알림 계약` 섹션을 추가
- 단순 북마크를 넘어 목록 조직 방식, 상태별 보존/정리, 알림 억제, `FavoriteListing` read model/API/분석 이벤트를 구체화
- 저장검색·가격변경·홈 진입 모듈과 연결되는 기준을 보강해 홈/찜목록/알림/DB/API 문서로 직접 파생 가능한 해상도를 높임

### v1.42 (2026-03-13 16:35 KST)
- 거래 성사 직결 정보인 `접속 가능 시간 / Availability Window / 미팅 가능 슬롯 정책` 섹션을 추가
- 자유서술 `availableTimeText`만으로는 부족했던 반복 가능 시간대, 예약 제안 가능 슬롯, 공개 범위, 서버 시간대 해석 규칙을 구조화
- `UserAvailabilityProfile`, `ListingAvailabilitySnapshot`, 관련 API/검색 랭킹/알림/운영 가드레일을 정리해 매물 상세·예약 UX·DB/API·분석 문서로 직접 파생 가능한 기준을 보강

### v1.41 (2026-03-13 16:23 KST)
- 거래 성사율과 사용자 신뢰에 직접 연결되는 `응답성 / 활동성 신호 계약` 섹션을 추가
- 최근 활동, 첫 응답 시간, 응답률, 휴면/자리비움 상태를 어떤 공개 레벨로 노출하고 검색/랭킹/알림 억제에 어떻게 활용할지 정의
- `UserActivitySnapshot`, `ResponseMetricDaily`, 관련 API/분석/운영 가드레일을 정리해 프로필·목록·홈·알림·운영 정책으로 직접 파생 가능한 기준을 보강

### v1.40 (2026-03-13 16:19 KST)
- 재방문 탐색과 즉시 거래 전환을 연결하기 위해 `저장 검색(Saved Search) / 재알림 구독 정책` 섹션을 추가
- 어떤 검색 조건을 저장할 수 있는지, 신규 매물/가격하락/상태복귀 알림을 어떤 억제 규칙으로 보낼지, 결과 기준선(baseline)과 만료/중지 정책을 구체화
- `SavedSearch`, `SavedSearchNotification`, 관련 API/분석 이벤트/운영 가드레일을 정리해 검색 구독 기능명세·DB·알림 정책으로 직접 파생 가능한 기준을 보강

### v1.39 (2026-03-13 14:12 KST)
- 검색/목록/딥링크/인덱스가 같은 해석을 쓰도록 `검색 쿼리 / 필터 패싯 / facet count 계약` 섹션을 추가
- 어떤 필터가 기본 포함값인지, 상태/서버/카테고리/가격/거래방식/속성 패싯을 어떻게 계산·노출할지와 `GET /listings` 응답 계약을 구체화
- facet count의 스냅샷 일관성, unavailable 상태 제외 규칙, zero-result 복구 UX, 인덱스/캐시/분석 포인트를 정리해 화면명세·API·검색 인프라 파생성을 높임

### v1.38 (2026-03-13 13:01 KST)
- read model, 알림, 감사로그, analytics가 같은 사건 경계를 공유할 수 있도록 `도메인 이벤트 / 아웃박스 / 발행 계약` 섹션을 추가
- 매물/채팅/예약/완료/분쟁/운영 제한에서 어떤 이벤트를 반드시 발행해야 하는지와 eventId, aggregateVersion, actor, ordering, idempotency 기준을 정리
- projection 재구성, 운영 타임라인, 외부 알림 fanout, dead-letter 대응까지 연결되는 비동기 계약을 보강해 DB/API/운영 문서 파생성을 높임

### v1.37 (2026-03-13 11:59 KST)
- 목록형/피드형 API가 실제 모바일 구현 단계에서 바로 쓰일 수 있도록 `목록 조회 / 커서 페이지네이션 / 스냅샷 일관성 계약` 섹션을 추가
- `GET /listings`, `GET /chats`, `GET /me/trades`, `GET /notifications`, 운영 큐 목록이 공통으로 따라야 하는 cursor, 정렬, 중복 방지, 새로고침, read model 반영 규칙을 정리
- 무한 스크롤, 딥링크 복귀, 실시간 갱신, soft delete/상태 변경 이후 목록 흔들림을 줄이기 위한 API/DB/read model/QA 기준을 보강

### v1.36 (2026-03-13 10:51 KST)
- 백엔드 구현·DB 설계·API 멱등성 규칙으로 직접 파생 가능하도록 `거래 집계 루트(Aggregate) / 동시성 / 커맨드 소유권 계약` 섹션을 추가
- Listing, ChatThread, Reservation, TradeCompletion/Dispute, Restriction이 각각 어떤 불변식을 소유하는지와 어느 경계에서 트랜잭션/락/보상 처리를 해야 하는지 정리
- 상태 충돌, 이벤트 발행 순서, read model 반영, 운영 강제 액션의 우선순위를 명문화해 기능명세와 서버 설계의 해석 차이를 줄임

### v1.35 (2026-03-13 10:23 KST)
- API/DB/딥링크/SEO/운영 로그가 같은 식별자 원칙을 쓰도록 `식별자(ID) / 공개 슬러그 / URL 수명주기 계약` 섹션을 추가
- 내부 불변 ID와 사용자 친화 슬러그를 분리하고, 매물/채팅/거래/분쟁/알림이 어떤 식별자를 외부에 노출해야 하는지와 재사용/변경 금지 원칙을 정리
- soft landing, canonical URL, deep link payload, 감사 로그, 비식별화/복구 시 식별자 유지 기준을 보강해 API 스펙·DB 스키마·라우팅 문서 파생성을 높임

### v1.34 (2026-03-13 10:14 KST)
- 운영 제재가 실제 제품 행동과 정확히 연결되도록 `계정/기능 제한(Restriction) 정책 계약` 섹션을 추가
- warning, listing_only, chat_only, trust_limited, temporary_suspend, permanent_ban 등 제재 scope와 수명주기, 사용자 노출, 해제/이의제기, API/DB/read model 기준을 정리
- 신고/안티어뷰즈/온보딩/권한 모델을 하나의 restriction 객체로 수렴시켜 운영정책·백오피스·화면 배너·쓰기 API 차단 규칙으로 직접 파생 가능한 기준을 보강

### v1.33 (2026-03-13 10:06 KST)
- 장기 방치 매물, stale 검색 품질 저하, 끌어올리기 남용 문제를 직접 다루는 `매물 수명주기 / 자동 만료 / 리프레시(끌어올리기) 정책` 섹션을 추가
- draft → active → stale → expired → closed 흐름과 상태·가시성·랭킹·알림·복구·운영 규칙을 분리해 화면명세/배치/job/API/DB 문서로 직접 파생 가능한 기준을 보강
- bump cooldown, daily quota, reserved/pending/completed 상태 제한, 재게시/복제 원칙, 분석 이벤트를 정의해 운영 악용 탐지와 검색 신뢰 정책 연결을 강화

### v1.32 (2026-03-13 09:53 KST)
- 매물 가격 변경이 검색 신뢰, 찜 알림, 협상 UX, 운영 악용 탐지와 직접 연결되도록 `가격 변경 / 가격 이력 / 가격 노출 정책` 섹션을 추가
- `sell`/`buy`, `fixed`/`negotiable`/`offer` 조합별 가격 수정 허용 범위, 상태별 변경 제한, 급격한 변동/도배성 수정 대응 기준을 구체화
- PriceSnapshot/PriceChangeEvent/read model/API/analytics 후보를 정리해 화면명세·DB·알림·운영정책으로 직접 파생 가능한 기준을 보강

### v1.31 (2026-03-13 09:50 KST)
- 가격 제안/역제안 기능을 화면·상태·운영까지 일관되게 내릴 수 있도록 `가격 제안(Offer) 화면/데이터/API 계약` 섹션을 추가
- 상세/채팅/내 거래에서 제안이 어떤 CTA와 배지로 보여야 하는지, 동시 제안·만료·중복·reserved/pending_trade 연동 규칙을 구체화
- Offer read model, API 후보, 분석 이벤트, 운영 모니터링 포인트를 정리해 이후 기능명세·DB·API 스펙 파생성을 높임

### v1.30 (2026-03-13 09:47 KST)
- 홈 화면을 단순 진입점이 아니라 거래 재개와 탐색 전환을 함께 담당하는 핵심 surface로 구체화하기 위해 `홈 화면(Home Surface) 계약` 섹션을 추가
- 개인화 모듈 우선순위, 액션 필요 카드, 서버 컨텍스트, 빈 상태/콜드스타트, read model/API projection, 분석 이벤트를 정리해 화면명세·API·랭킹·운영 문서 파생성을 높임
- 홈 진입이 `거래 성사율`에 직접 연결되도록 딥링크/배지/추천 슬롯/안전 배너의 노출 원칙을 보강

### v1.29 (2026-03-13 09:37 KST)
- 기능 출시 리스크를 줄이기 위해 `피처 플래그 / 단계적 출시 / 롤백 정책` 섹션을 추가
- 채팅, 예약, 완료확정, 신고/모더레이션, 검색 랭킹, 알림 고도화 기능을 어떤 단위로 점진 배포할지와 롤백 기준을 정의
- 운영자 커뮤니케이션, 관측성, QA, 데이터 보호 관점에서 출시 게이트와 연동되는 운영 규칙을 보강

### v1.28 (2026-03-13 09:34 KST)
- 멀티디바이스 환경에서의 푸시 억제 정밀도와 `ChatDeviceState` 저장 전략을 정리한 `디바이스 상태 / 멀티디바이스 읽음·푸시 억제 정책` 섹션을 추가
- 사용자 단위 읽음과 디바이스 단위 연결 상태를 어떻게 분리하고, 어떤 경우에 푸시를 억제/발송할지 우선순위 규칙을 보강
- MVP와 Post-MVP를 나누어 영속 저장소, TTL, observability, QA 시나리오까지 파생 가능한 기준을 추가

### v1.27 (2026-03-13 09:29 KST)
- projection을 실제 운영 가능한 수준으로 다루기 위해 재빌드/정합성 복구/운영 runbook 원칙을 추가
- 전체 리빌드보다 대상별 부분 재계산을 우선하는 기준, freeze/replay/cutover 절차, 사용자 영향 최소화 원칙을 정의
- projection dead-letter, 재처리, 감사로그, QA/출시 게이트에 연결되는 체크리스트를 보강

### v1.26 (2026-03-13 09:26 KST)
- 채팅/내거래/알림/운영 화면이 공유해야 할 read model 및 materialized projection 전략을 추가
- `ChatThreadSummary`, `TradeThreadProjection`, `NotificationFeedProjection`, `ModerationQueueProjection` 단위와 갱신 트리거를 정의해 화면/API/DB 파생성을 보강
- 분쟁 첨부와 일반 채팅 첨부의 저장소/권한 분리 원칙, projection 재빌드/정합성/관측성 기준을 추가

### v1.25 (2026-03-13 09:19 KST)
- 회원가입 이후 거래 시작까지의 병목을 줄이기 위해 `신규 사용자 온보딩 / 인증 / 신뢰 부트스트랩` 섹션을 추가
- 게스트→회원→프로필완료→거래가능 상태 전이, 초기 제한/해제 조건, 화면/API/운영 연계 기준을 정리해 인증·프로필·신뢰 정책 파생성을 보강
- 신규 계정의 스팸 억제와 정상 사용자 전환을 함께 고려한 단계별 가드레일/리마인드/분석 이벤트 후보를 추가

### v1.24 (2026-03-13 09:13 KST)
- 거래 허용 범위, 금지 품목/행위, 콘텐츠 검수 단계, 제재 연동 규칙을 정리한 `거래 가능 품목 및 콘텐츠 정책` 섹션을 추가
- 매물/채팅/이미지/OCR 탐지 정책을 연결해 운영정책·관리자 백오피스·신고 사유 코드 문서로 직접 파생 가능한 기준을 보강
- `사전 차단 / 사후 숨김 / 운영 검토`의 경계를 명확히 해 모바일 등록 UX와 안전 정책 충돌을 줄이는 기준을 추가

### v1.23 (2026-03-13 08:47 KST)
- 모바일 앱 기준 주요 화면의 route/deep link/진입 가드/뒤로가기 복귀 규칙을 정리한 `화면 라우팅 및 딥링크 계약` 섹션을 추가
- 홈/목록/상세/채팅/내거래/알림/프로필/운영 화면이 어떤 식별자와 컨텍스트로 연결되는지 명문화해 화면명세·API·푸시 딥링크 파생성을 높임
- soft landing, 만료 딥링크 fallback, 인증 요구/권한 부족 시 라우팅 원칙을 정의해 모바일 거래 흐름 복구성과 운영 일관성을 보강

### v1.22 (2026-03-13 08:44 KST)
- 액션별 rate limit / cooldown / 안티어뷰즈 정책 매트릭스를 추가해 채팅·매물·예약·신고·운영 API의 제어 기준을 구체화
- 위험도 기반 단계 제한과 사용자 노출 원칙을 정리해 전환 저해를 최소화하면서도 스팸·도배·괴롭힘 패턴을 억제하는 운영 기준을 보강
- API 응답/분석/운영대시보드가 공유할 `limitKey`, `restrictionScope`, `restrictionReasonCode` 후보를 정리해 백엔드/운영 문서 파생성을 높임

### v1.21 (2026-03-13 08:34 KST)
- 카테고리별 속성 템플릿(Item Attribute Template) 구조를 추가해 매물 등록 폼, 검색 필터, DB 스키마, API DTO가 동일한 속성 모델을 공유하도록 보강
- 자유입력 설명과 구조화 속성을 분리하는 원칙, 속성 타입/검증/표시 규칙, 카테고리별 필수 속성 정책을 정의해 구현 해상도를 높임
- 카탈로그/운영 도구/분석 이벤트 관점에서 어떤 속성이 승격·표준화 대상인지 정리해 이후 기능명세와 운영 문서 파생성을 강화

### v1.20 (2026-03-13 08:31 KST)
- 거래 성사율을 높이기 위한 가격 제안/역제안(offer / counter-offer) 도메인을 추가해 단순 채팅과 실행 액션의 경계를 명확화
- 제안 상태 머신, 만료/중복/동시성 규칙, 채팅/예약/매물 상태와의 연동 원칙을 정의해 API/DB/화면 명세로 직접 파생 가능한 기준을 보강
- 내 거래/채팅에서 필요한 CTA, 알림, 분석 이벤트를 정리해 모바일 중심 협상 UX와 운영 추적 기준을 추가

### v1.19 (2026-03-13 08:14 KST)
- 사용자 기준 읽음 커서와 기기/세션 기준 동기화 상태를 분리하는 채팅 participant/device state 모델을 추가
- SSE 중심 실시간 전송의 연결 수명주기, 재연결/backfill 규칙, 핵심 SLI/SLO 후보를 정의해 채팅/알림/운영 모니터링 기준을 보강
- trade/chat read model이 어떤 projection을 가져야 하는지와 QA/운영 체크포인트를 보강해 이후 API/DB/관측성 문서로 직접 파생 가능한 기준을 추가

### v1.18 (2026-03-13 08:12 KST)
- 홈/검색 결과에서 `sell`/`buy` 매물을 어떻게 혼합·분리 노출할지에 대한 정보구조 및 랭킹/필터 기본 정책을 추가
- 사용자의 현재 의도(사기/팔기/둘다)에 따라 목록 CTA, 카드 라벨, 추천 슬롯, 빈 상태 문구가 달라져야 하는 화면 계약을 보강
- 홈/검색 surface용 read model, API projection, 분석 이벤트를 정의해 이후 화면명세/API/랭킹 문서로 직접 파생 가능한 기준을 추가

### v1.17 (2026-03-13 07:57 KST)
- `sell`/`buy` 매물의 의미 차이, 역할 해석, 상태 전이/완료 처리 기준을 분리해 거래 플로우 및 데이터 계약의 모호함을 줄이는 섹션을 추가
- 매수글 기준 가격/수량/채팅 시작/예약/완료/후기 연결 규칙을 정리해 화면명세, API 명세, QA 시나리오 파생성을 보강
- buy/sell 혼합 검색·랭킹·알림·운영 정책 차이를 명문화해 이후 기능명세와 정책 문서의 기준 해상도를 높임

### v1.16 (2026-03-13 07:54 KST)
- 계정 신뢰 신호를 `인증/평판/행동위험` 3계층으로 분리하고, 공개 가능한 신뢰 신호와 내부 리스크 신호를 구분하는 정책을 추가
- 신규 가입/닉네임/프로필/차단/제한과 연결되는 계정 수명주기, 안티어뷰즈 제한, 장치/세션 추적 최소 기준을 정리해 운영정책/DB/API 파생성을 보강
- 거래 전환을 해치지 않으면서도 다계정/대량문의/스팸성 등록을 제어하기 위한 rate limit·cooldown·리스크 기반 단계 제한 초안을 추가

### v1.15 (2026-03-13 07:48 KST)
- 채팅 실시간성 MVP를 `풀 duplex 실시간 강제`가 아닌 `동기화 일관성 우선` 관점으로 정리하고, 권장 전송 방식/폴백/재연결 계약을 추가
- typing/presence를 MVP 포함 범위와 Post-MVP 확장 범위로 분리해 모바일 UX·푸시·읽음 모델 간 충돌을 줄임
- `ChatParticipantState`를 별도 읽기/알림/뮤트/읽음 커서 원천 객체로 명문화해 DB/API/동기화 명세 파생성을 보강

### v1.14 (2026-03-13 07:38 KST)
- 채팅 첨부/미디어를 MVP 범위에서 어떻게 저장·전달·검수할지에 대한 정책을 추가해 메시지/스토리지/API/운영 명세 파생성을 높임
- 이미지 첨부 허용 범위, 예약/분쟁 증빙 첨부 차이, 썸네일/원본/보관/마스킹 원칙을 정리해 모바일 업로드 UX와 안전 정책 기준을 보강
- 업로드 수명주기(임시 업로드 → 메시지 확정 → 미참조 정리)와 미디어 관련 분석/QA 포인트를 추가

### v1.13 (2026-03-13 07:35 KST)
- 채팅 동기화 기준을 `메시지 전송 상태 + 읽음 커서 + 타임라인 앵커` 관점으로 보강해 실시간/폴링/오프라인 복구 설계의 기준을 추가
- 채팅 목록/상세/재진입 시 unread 계산, lastReadMessage, delivery 상태, gap 복구 정책을 정리해 API/DB/QA 파생성을 높임
- 알림 중복 억제와 채팅 재전송/중복 전송 방지 기준을 연결해 모바일 네트워크 불안정 환경 대응 기준을 보강

### v1.12 (2026-03-13 07:32 KST)
- 채팅 메시지 도메인을 `일반 메시지`가 아니라 거래 실행 로그까지 포함하는 event timeline으로 재정의하고, 메시지 타입/읽음/전송 상태/시스템 이벤트 계약을 추가
- 첨부/민감정보/빠른문구/예약 연동 규칙을 보강해 채팅 화면명세, 메시지 DB 스키마, 실시간/폴링 API 설계로 직접 파생 가능한 기준을 마련
- 채팅 목록/상세 응답 projection과 분석 이벤트 후보를 추가해 `채팅`과 `내 거래`가 같은 이벤트 원천을 공유하도록 정리

### v1.11 (2026-03-13 07:26 KST)
- `내 거래`를 단순 채팅 목록이 아닌 거래 실행 워크스페이스로 정의하고, 리스트/상세/타임라인/액션 우선순위 계약을 추가해 화면명세와 API projection 파생성을 높임
- 거래 객체 집계 단위(`trade thread`)와 상태 배지, SLA 타이머, CTA 묶음을 명문화해 예약/완료/분쟁 흐름이 모바일 화면에서 어떻게 조직되어야 하는지 구체화
- `GET /me/trades` 및 거래 상세 응답 후보 구조를 추가해 홈/내거래/알림 딥링크/운영 추적이 공통 식별자를 사용할 수 있도록 보강

### v1.10 (2026-03-13 07:23 KST)
- 매물 도메인 필드를 `검색/거래/운영/표시` 관점으로 재정리한 계약 섹션을 추가해 화면명세, DB 스키마, API 응답 설계의 기준 해상도를 높임
- 가격/수량/아이템명/서버/이미지/노출용 요약 필드의 파생 규칙과 수정 제한 규칙을 보강해 목록/상세/랭킹/검증 로직 해석 차이를 줄임
- Listing API 읽기/쓰기 DTO 분리 원칙과 상태별 수정 가능 범위를 명시해 이후 기능명세/백엔드 설계 파생성을 강화

### v1.9 (2026-03-13 07:16 KST)
- 거래완료 공개 상태와 내부 완료 단계 enum을 분리하고, 별도 Dispute 객체 도입 초안을 추가해 완료/분쟁/운영/DB 설계 경계를 명확화
- 완료 확인 대기·자동확정·분쟁 전환의 API/데이터/화면 해석 기준을 보강해 기능명세와 상태머신 파생성을 높임
- 공개 프로필 응답 레벨(`summary`/`member`/`participant`)과 필드별 계약 초안을 추가해 비회원/회원/거래상대 노출 차이를 API 수준에서 구체화

### v1.8 (2026-03-13 07:09 KST)
- 인게임/오프라인 거래 약속을 공통 구조로 다루는 미팅 포인트 모델, 공개 범위, 검증 규칙을 추가해 예약/화면/DB/API 파생 기준을 보강
- 예약 협의 단계별 장소 정보 노출 수준과 안전 가드레일을 명문화해 개인정보/실행 UX/운영정책 간 충돌을 줄임
- 장소 후보/최근 사용/빠른 선택 중심의 예약 UX 요구사항과 관련 데이터/API 후보를 추가해 모바일 거래 완료 플로우의 구현 해상도를 높임

### v1.7 (2026-03-13 06:54 KST)
- 알림 이벤트별 수신자/채널/억제 조건/딥링크를 정리한 발송 매트릭스를 추가해 푸시/인앱/운영 명세 파생성을 높임
- 사용자 알림 설정 기본값과 이벤트 중요도 연결 규칙, 재알림 제한 원칙을 구체화해 모바일 UX와 운영 정책 해석 차이를 줄임
- 문서 말미의 중복/깨진 섹션 번호를 정리해 PRD 구조 일관성을 보정

### v1.6 (2026-03-13 06:51 KST)
- 공개 프로필 범위와 인덱싱 기본안을 추가해 SEO/프로필 화면/신뢰 정책 간 충돌을 줄임
- 후기 보복 리스크 완화 장치와 후기 분쟁 처리 원칙을 추가해 리뷰 정책을 운영 가능한 수준으로 구체화
- `reserved` 장기 유지/우선상대 반복 교체 등 악용 탐지 룰과 분쟁 중 추가 소명 UX를 추가해 운영/QA/백오피스 명세 파생성을 강화

### v1.5 (2026-03-13 06:47 KST)
- MVP 출시 게이트, 화면/운영/데이터 관점의 수용 기준을 추가해 기능명세/QA/운영 준비 문서로 직접 파생 가능한 기준을 보강
- API 표면을 MVP 필수/확장 후보/내부 전용으로 재분류하고 버전 전략, 멱등성/비호환 변경 원칙을 추가해 API 스펙 파생성을 강화
- 핵심 KPI에 대한 초기 목표값 후보와 출시 후 30일 관찰 항목을 추가해 런칭 판단 기준을 구체화

### v1.4 (2026-03-13 06:36 KST)
- API/DB/운영이 공통으로 사용하는 사유 코드 체계를 추가해 상태 변경, 제재, 경고, 자동화 실패 원인 표준화를 보강
- listing/chat/reservation/completion/review/report/user 객체별 reason code 후보를 정리해 운영정책/에러모델/감사로그 파생 입력을 강화
- 사용자 노출 문구와 내부 상세 사유를 분리하는 원칙, 분석 이벤트 연계 기준, 버전 관리 규칙을 추가

### v1.3 (2026-03-13 06:33 KST)
- 목록/상세/채팅/내거래/운영 화면에서 직접 파생 가능한 공통 화면 상태(loading/empty/error/partial/policy-blocked) 규격을 추가해 화면명세와 QA 기준을 보강
- 서버/클라이언트/운영이 동일하게 해석할 수 있도록 `availableActions`, `viewerContext`, 경고 배너/제한 사유 코드의 응답 표준 초안을 추가
- 핵심 객체별 삭제/탈퇴/비식별화/복구 원칙을 정리해 DB 스키마, 운영정책, 개인정보 문서 파생 입력을 보강
### v1.2 (2026-03-13 06:30 KST)
- 상태 변화와 운영 처리를 자동화하는 타이머/배치 정책을 추가해 예약 만료, 완료 자동확정, 알림 리마인드, 휴면 매물 처리 기준을 명문화
- 주요 사용자 화면별 목적/주요 API/핵심 상태/권한 가드레일을 연결한 화면-API 매핑 섹션을 추가해 기능명세와 API 명세 파생이 쉬워지도록 정리
- 자동화 처리 실패 시 재시도, 감사 로그, 사용자 노출 원칙을 추가해 운영/백엔드 설계 입력을 보강
### v1.1 (2026-03-13 06:21 KST)
- 객체별 권한/가시성/행위 가능 조건 매트릭스를 추가해 화면 요구사항, API 인가 규칙, 운영 정책이 공통 기준을 쓰도록 정리
- 비회원/회원/거래참여자/작성자/운영자 기준으로 매물·채팅·예약·후기·신고 데이터의 조회/수정 범위를 구체화
- 차단/제재/탈퇴/숨김 상태가 권한에 미치는 영향과 우선순위 규칙을 명문화
### v1.0 (2026-03-13 06:18 KST)
- 거래완료 요청/상호확정/자동확정/분쟁 전환 규칙을 명문화해 완료 상태머신과 운영 처리 기준을 보강
- `reserved` 상태의 신규 문의 허용 기본안을 추가해 검색/상세/채팅 CTA와 랭킹 정책의 기준을 구체화
- 완료 후 재오픈, 부분거래, 예약 상대 이탈 시 재게시 규칙을 추가해 예외 흐름을 정리
### v0.9 (2026-03-13 06:10 KST)
- 데이터 정합성 제약, 유니크/인덱스 후보, 상태 이력 저장 원칙을 추가해 DB 설계 입력을 보강
- 화면별 주요 액션 가드레일과 상태별 CTA 규칙을 추가해 화면 요구사항을 구현 단계 수준으로 구체화
- 운영/법적 고지 초안을 추가해 서비스 책임 범위와 사용자 안내 문구 방향을 명확화

### v0.8 (2026-03-13 06:07 KST)
- 후기 공개 시점, 수정/비공개 처리, 상호 작성 마감 정책을 구체화
- SEO/공개 범위, 인덱싱 원칙, 공개 URL 수명주기 정책을 추가
- 푸시 알림 opt-in, 채널 우선순위, 소음 제어 정책을 추가
- 검색 자동완성/추천어/오타 보정 정책과 운영 룰을 추가
- 운영 대시보드 지표와 이상징후 탐지 룰 초안을 추가

### v0.7 (2026-03-13 05:58 KST)
- 관리자 백오피스 권한 레벨과 화면별 처리 절차를 추가
- API 에러 모델, 권한/상태 충돌 규칙, 멱등성 고려사항을 추가
- 서버/아이템 카탈로그 표준화 전략과 자유입력 혼합 정책을 추가
- 차단/뮤트/사용자 안전 UX 정책 초안을 추가

### v0.6 (2026-03-13 04:56 KST)
- 비회원 공개 범위와 공개 정보 레벨 정책 초안을 추가
- 채팅 내 개인정보/외부연락처 교환 정책과 자동 제어 원칙을 추가
- 매물/예약 입력 validation 및 제한 규칙을 추가해 API/화면 검증 기준 보강
- 비기능 요구사항을 SLI/SLO, 보안, 보관, 확장성 관점까지 확장

### v0.5 (2026-03-13 03:50 KST)
- 핵심 엔티티별 필수/선택 컬럼 표를 추가해 DB 설계 입력 수준으로 구체화
- 검색/정렬/노출 랭킹 규칙 초안을 추가해 목록 UX와 추천 로직 기준 정리
- 운영 SLA, 우선순위, 이의제기 처리 정책 초안 추가
- 채팅/예약/완료 관련 API 요청/응답 예시를 추가해 인터페이스 구체화

### v0.4 (2026-03-13 02:48 KST)
- 후기/신뢰도 정책과 노출 원칙 초안 추가
- 신고/분쟁/모더레이션 운영 정책 상세화
- 매물/채팅/예약 상태별 사용자 액션 매트릭스 추가
- 분석 이벤트, 핵심 KPI, 단계별 출시 범위 초안 추가
- 오픈 질문 섹션 번호 정리 및 운영 관점 가정 보강

### v0.3 (2026-03-13 01:46 KST)
- 채팅방/예약 워크플로우와 세부 상태 정의 추가
- 알림 이벤트 및 발송 규칙 초안 추가
- 화면별 요구사항 및 정보구조 구체화
- 핵심 데이터 엔티티 후보와 관계 초안 추가
- API 후보 목록과 관리자 운영 API 범위 초안 추가

### v0.2 (2026-03-13 01:03 KST)
- 제품 비전과 서비스 원칙을 명확화
- 도메인 용어집 추가
- 사용자 역할과 권한 초안 추가
- 매물 데이터 스키마 초안 구체화
- 거래 상태 정의 및 상태 전이 규칙 추가
- 핵심 비즈니스 규칙/예외 규칙 추가
- 오픈 질문 및 가정 섹션 추가

### v0.1 (2026-03-13 00:49 KST)
- 초기 PRD 초안 작성

## 0. 문서 목적
이 문서는 리니지 클래식 유저를 위한 개인 간 거래 중개 서비스의 MVP 및 확장 방향을 정의한다. 목표는 이 문서를 점진적으로 발전시켜, 이후 기능명세서/화면설계/DB설계/API스펙/운영정책 문서를 작성할 수 있는 수준까지 구체화하는 것이다.

## 1. 제품 개요
### 1.1 한 줄 정의
리니지 클래식 유저가 아이템/재화를 사고팔기 위해 매물 등록, 관심 표시, 1:1 채팅, 거래 예약, 거래 상태 관리까지 수행할 수 있는 게임 특화 거래 중개 플랫폼.

### 1.2 해결하려는 문제
- 기존 거래 커뮤니티/사이트는 검색과 분류가 불편하다.
- 거래 상태(가능/예약중/완료)가 명확하지 않다.
- 거래 약속과 장소 조율이 서비스 안에서 자연스럽게 이어지지 않는다.
- 모바일 UX가 떨어진다.
- 신뢰도/후기/신고 체계가 약하다.

### 1.3 핵심 가치
1. 쉽게 올린다.
2. 쉽게 연결된다.
3. 쉽게 완료된다.

### 1.4 제품 비전
린클은 단순 게시판이 아니라, “매물 발견 → 대화 → 예약 → 실제 거래 완료”까지 이어지는 거래 실행 도구를 지향한다. 핵심은 매물 노출량보다 거래 성사율과 거래 완료 경험을 높이는 것이다.

### 1.5 서비스 원칙
1. 모바일에서 한 손으로 거래를 이어갈 수 있어야 한다.
2. 매물 정보보다 거래 가능성 신호(응답성, 상태, 최근 활동)를 더 잘 보여줘야 한다.
3. 거래 상태는 게시글 단위가 아니라 실제 협의 흐름과 연결되어야 한다.
4. 분쟁은 완전히 없앨 수 없으므로, 예방 정보와 사후 운영 도구를 함께 제공해야 한다.

## 2. 제품 목표
### 2.1 상위 목표
- 린클 개인거래를 더 빠르게 연결한다.
- 등록부터 거래완료까지 한 서비스 안에서 처리한다.
- 서버/아이템/장소/접속 시간 중심의 게임 특화 UX를 제공한다.
- 반복 거래 유저를 위한 상태/채팅/신뢰 흐름을 만든다.

### 2.2 MVP 성공 기준(초안)
- 사용자가 3분 이내 매물 등록 가능
- 관심 사용자와 판매자 간 채팅 진입률 확보
- 등록된 매물 중 일정 비율 이상이 상태 업데이트를 거쳐 완료/취소로 종결
- 거래완료 후 후기 작성률 확보

## 3. 핵심 사용자
### 3.1 판매자
- 아이템/재화를 빠르게 판매하고 싶은 유저
- 시세보다 “빠른 처분”이 중요할 수 있음

### 3.2 구매자
- 원하는 아이템을 검색하고 바로 거래 상대를 찾고 싶은 유저
- 가격, 강화, 서버, 접속 시간 조건을 빠르게 비교하고 싶음

### 3.3 고빈도 거래자
- 여러 매물을 동시에 운영하며 빠른 응답/예약 관리가 필요한 유저
- 여러 채팅과 예약을 병렬로 관리해야 함

### 3.4 운영자(관리자/모더레이터)
- 허위매물, 사기 의심, 욕설, 노쇼, 계정 악용을 감시/조치하는 역할
- 신고 처리, 제재, 매물 숨김, 계정 제한 수행

## 4. 도메인 용어집
- **매물(Listing)**: 판매/구매 의사를 표현하는 기본 게시 단위
- **판매 매물**: 판매자가 특정 아이템/재화를 팔기 위해 등록한 매물
- **구매 매물**: 구매자가 특정 아이템/재화를 사고 싶어 등록한 매물
- **거래 상대(Interested User)**: 매물에 관심을 보이고 채팅을 시작한 사용자
- **채팅방(Trade Chat)**: 특정 매물을 기반으로 생성된 1:1 대화 공간
- **예약(Reservation)**: 특정 상대와 거래 시간/장소/서버/캐릭터명 등을 합의한 상태 또는 객체
- **거래 상태(Trade Status)**: 매물이 현재 거래 가능한지, 예약되었는지, 완료되었는지 표현하는 상태값
- **거래 완료(Completed)**: 당사자가 실제 거래가 끝났다고 표시한 종결 상태
- **노쇼(No-show)**: 약속된 시간/장소에 나타나지 않거나 응답 없이 거래를 무산시키는 행위
- **신뢰도(Trust)**: 후기, 완료 이력, 신고/제재 기록, 응답성 등의 조합으로 표현되는 사용자 신뢰 지표
- **거래 장소(Meeting Point)**: 실제 오프라인(예: PC방 인근) 또는 인게임 내 접선 장소

## 5. 사용자 역할 및 권한(초안)
### 5.1 비회원
- 매물 목록/상세 일부 조회 가능 여부는 정책 결정 필요
- 채팅 시작, 찜, 등록, 신고는 불가

### 5.2 회원
- 매물 등록/수정/종료
- 찜 등록/해제
- 채팅 시작/응답
- 예약 제안/수정/확정
- 거래 상태 변경(권한 범위 내)
- 후기 작성
- 신고 접수

### 5.3 판매자/구매자 특수 권한
- 본인 매물의 공개 상태 및 거래 상태 변경
- 본인 채팅방 내 예약 확정/취소
- 거래완료/거래취소 제안

### 5.4 운영자
- 신고 검토 및 상태 변경
- 매물 숨김/차단
- 채팅 로그 및 신고 증빙 열람(정책 범위 내)
- 계정 제재, 경고, 이용 제한

## 6. 핵심 기능
### 6.1 매물 등록
- 판매 / 구매 유형 선택
- 서버, 카테고리, 아이템명, 옵션/강화/수량, 가격, 설명, 이미지
- 거래 가능 시간대 및 희망 장소 등록

### 6.2 매물 탐색
- 서버 / 아이템명 / 카테고리 / 가격 / 상태 필터
- 최신순 / 가격순 / 관심순 정렬

### 6.3 관심 표시
- 찜 저장
- 상태 변경 및 가격 변경 알림 연동(확장)

### 6.4 1:1 채팅
- 매물 상세에서 바로 채팅 시작
- 빠른 제안 문구 제공
- 거래 약속과 상태 변경 액션 연계

### 6.5 거래 예약
- 날짜/시간/서버/캐릭터명/장소/메모 입력
- 채팅방 내 예약 카드 형태 제공

### 6.6 거래 상태 관리
- 거래가능
- 예약중
- 거래대기
- 거래완료
- 거래취소

### 6.7 후기 / 신뢰도
- 거래 완료 후 상호 후기
- 거래 횟수 / 긍정 후기 / 신고 이력 기반 신뢰도 표시

### 6.8 신고 기능
- 허위매물 / 사기 의심 / 욕설 / 노쇼 등 신고

## 7. MVP 범위
### 7.1 포함
- 로그인
- 매물 등록/수정
- 리스트/상세
- 검색/필터
- 찜
- 1:1 채팅
- 거래 예약
- 상태 변경
- 후기(기본)
- 신고(기본)

### 7.2 제외
- 결제/에스크로
- 실시간 시세 엔진
- PC방 지도 자동 추천
- 고급 평판 알고리즘
- 관리자 고급 통계

## 8. 사용자 플로우
### 8.1 판매자 플로우
1. 로그인
2. 판매글 등록
3. 문의 수신
4. 채팅 협의
5. 예약 설정
6. 거래대기/예약중 전환
7. 거래완료
8. 후기 수신

### 8.2 구매자 플로우
1. 검색/필터
2. 상세 확인
3. 찜 또는 채팅 시작
4. 시간/장소 협의
5. 거래 진행
6. 거래 완료
7. 후기 작성

### 8.3 운영자 플로우(초안)
1. 신고 접수 확인
2. 신고 유형 분류
3. 증빙 검토(매물/채팅/이력)
4. 임시 숨김 또는 경고 여부 판단
5. 조치 확정 및 기록 저장
6. 반복 위반 사용자 누적 관리

## 9. 정보 구조(초안)
- 홈
- 거래소
- 찜 목록
- 채팅
- 내 매물
- 내 거래
- 마이페이지
- 신고/운영센터(운영자)

## 10. 매물 스키마 초안
아래 항목은 이후 DB 스키마와 API 명세의 기초가 되는 후보 필드다.

### 10.1 공통 필드
- listingId
- listingType: `sell` | `buy`
- authorUserId
- serverId
- categoryId
- title
- description
- priceType: `fixed` | `negotiable` | `offer`
- priceAmount
- quantity
- status
- visibility: `public` | `hidden` | `blocked`
- createdAt
- updatedAt
- bumpedAt
- expiresAt (정책 결정 필요)

### 10.2 게임 특화 필드
- itemName
- itemGrade / rarity (가정)
- enhancementLevel
- optionsText
- packageYn (묶음 판매 여부)
- tradeMethod: `in_game` | `offline_pc_bang` | `either`
- preferredMeetingAreaText
- preferredMeetingServerText
- availableTimeText
- sellerCharacterName / buyerCharacterName 공개 여부

### 10.3 메타 필드
- viewCount
- favoriteCount
- chatCount
- reservationCount
- lastActivityAt
- completionAt
- cancellationReasonCode

### 10.4 이미지 정책(초안)
- 매물당 최대 N장
- 첫 번째 이미지를 대표 이미지로 사용
- 선정성/불법/무관 이미지 업로드 금지
- 과도한 개인정보 노출 이미지 금지

## 11. 거래 상태 정의 및 전이 규칙
매물 상태는 단순 게시 상태가 아니라 거래 진행 가능성을 표현해야 한다.

### 11.1 매물 상태 정의
- **available**: 거래 가능. 새로운 채팅과 제안을 받을 수 있음.
- **reserved**: 특정 상대와 우선 거래 협의가 확정된 상태. 신규 문의는 받을 수 있으나 제한 정책 결정 필요.
- **pending_trade**: 거래 약속이 잡혔고 실제 실행 대기 상태.
- **completed**: 거래가 끝난 종결 상태. 더 이상 신규 문의 불가.
- **cancelled**: 작성자가 거래를 중단하거나 거래가 무산된 종결 상태.

### 11.2 권장 상태 전이
- `available -> reserved`
- `reserved -> pending_trade`
- `pending_trade -> completed`
- `reserved -> available` (예약 해제)
- `pending_trade -> available` (거래 불발 후 재오픈)
- `available -> cancelled`
- `reserved -> cancelled`
- `pending_trade -> cancelled`

### 11.3 금지 또는 제한 전이
- `completed -> available` 금지
- `completed -> reserved` 금지
- `cancelled -> completed` 금지
- 종결 상태 재오픈이 필요하면 새 매물 생성 또는 운영 정책 기반 복구 처리

### 11.4 상태 변경 권한
- 기본적으로 매물 작성자가 상태를 변경
- 거래완료는 상대방 확인 흐름을 둘지 여부는 정책 결정 필요
- 운영자는 정책 위반 시 강제 숨김/강제 종료 가능

### 11.5 상태와 채팅의 관계
- 하나의 매물에 여러 채팅방이 생길 수 있음
- `reserved` 또는 `pending_trade`로 진입하면 어떤 채팅 상대와 진행 중인지 연결돼야 함
- 특정 채팅방의 예약 확정이 매물 전체 상태에 영향을 주는 구조 필요

## 12. 핵심 비즈니스 규칙(초안)
### 12.1 매물 생성/수정
- 작성자는 본인 매물만 수정 가능
- 거래 진행 중(`pending_trade`)에는 일부 필드 수정 제한 필요
- 완료/취소 상태의 매물은 수정 가능 범위를 제한하거나 재등록 유도

### 12.2 중복/도배 방지
- 동일 유저가 짧은 시간 내 유사 매물을 과도하게 올리는 경우 제한 가능
- 동일 아이템 다중 등록은 허용하되 스팸성 패턴은 탐지 대상

### 12.3 가격 정책
- 가격 미입력 허용 여부 결정 필요
- `offer` 타입일 경우 채팅 유도형 매물로 분류 가능
- 비정상적 가격(0원, 과도한 숫자)은 validation 필요

### 12.4 예약 정책
- 한 매물에 동시에 복수 예약은 기본적으로 불가
- 예약에는 만료 시각 또는 예약 해제 기준이 필요
- 예약 상대, 시간, 장소가 기록 가능한 구조여야 함

### 12.5 거래 완료 정책
- 거래완료 처리 시 매물은 종결
- 거래완료 후 후기 작성 가능
- 후기 작성 가능 기간 제한 여부는 정책 결정 필요

### 12.6 취소/불발 정책
- 예약 후 무응답, 노쇼, 조건 변경 등의 사유를 사유 코드로 남길 수 있어야 함
- 취소 사유는 사용자 개인 메모와 운영용 코드 분리 가능성이 있음

## 13. 예외/엣지 케이스(초안)
- 판매자가 여러 명과 동시에 채팅하다가 한 명과만 예약 확정하는 경우
- 예약 후 거래 시간이 지나도 완료/취소 처리가 안 되는 경우
- 거래완료 후 상대가 실제로는 미완료라고 주장하는 경우
- 판매자가 아이템 일부만 판매해 수량이 줄어드는 경우
- 구매자가 제안 후 장시간 무응답 상태가 되는 경우
- 동일 사용자 간 반복적인 허위 문의/괴롭힘 발생
- 서버/아이템 정보가 불명확한 자유서술 매물 등록

## 14. 비기능 요구사항
- 모바일 우선 설계
- 리스트/검색 응답 속도 최적화
- 채팅은 준실시간 또는 실시간 제공
- 신고/분쟁 대응 로그 보관
- 개인정보 최소 수집
- 운영자 감사 로그 보관 필요
- 상태 변경, 예약 변경, 신고 처리 이력은 추적 가능해야 함

## 15. 초기 차별화 포인트
- 게시판형이 아닌 거래 플로우형 UX
- 게임 특화 필드(서버/강화/거래장소/접속시간)
- 채팅 + 예약 + 상태관리 통합
- 모바일 UX 개선
- 후기/신뢰 구조화
- 거래 성사 중심의 상태/알림 설계

## 16. 향후 문서 발전 목표
이 문서는 아래 문서를 작성할 수 있는 수준까지 발전시킨다.
- 기능명세서
- 화면별 요구사항 문서
- DB 스키마 문서
- API 스펙
- 운영정책 문서

## 17. 채팅/예약 워크플로우 상세
### 17.1 채팅방 생성 규칙
- 채팅방은 `listingId + buyerUserId + sellerUserId` 조합 기준으로 1개를 기본 생성 단위로 본다.
- 동일 매물에 대해 동일 상대와 중복 채팅방 생성은 금지한다.
- 매물 작성자가 구매 매물을 올린 경우에도 상대 역할만 바뀔 뿐 구조는 동일하다.
- 차단된 사용자, 제재 중 사용자, 종료된 매물에는 신규 채팅 생성이 제한될 수 있다.

### 17.2 채팅방 상태 정의
- `open`: 일반 대화 가능 상태
- `reservation_proposed`: 한쪽이 예약안을 제안한 상태
- `reservation_confirmed`: 양측 또는 작성자 정책 기준으로 예약 확정된 상태
- `trade_due`: 예약 시각이 임박했거나 도래한 상태
- `deal_completed`: 해당 채팅 기준 거래 완료 처리된 상태
- `deal_cancelled`: 해당 채팅 기준 협의 종료 상태
- `report_locked`: 신고/분쟁 처리로 운영 잠금된 상태

### 17.3 예약 객체 필수 속성(초안)
- reservationId
- listingId
- chatRoomId
- proposerUserId
- counterpartUserId
- scheduledAt
- timeWindowText
- meetingType: `in_game` | `offline_pc_bang` | `either`
- serverId 또는 serverText
- meetingPointText
- characterNameA / characterNameB (선택)
- noteToCounterparty
- reservationStatus: `proposed` | `confirmed` | `expired` | `cancelled` | `fulfilled` | `no_show_reported`
- expiresAt
- confirmedAt
- fulfilledAt
- cancelledAt
- cancellationReasonCode

### 17.4 예약/상태 연동 규칙
- 예약 제안만으로 매물 상태를 즉시 `reserved`로 바꾸지 않는다.
- 예약 확정 시 매물 상태를 `reserved` 또는 `pending_trade`로 변경한다. 세부 기준은 예약 시각과의 거리 기준으로 나눈다.
- 예약 시각이 임박한 경우(예: 2시간 이내, 가정) `pending_trade` 진입이 가능하다.
- 예약이 취소/만료되면 매물은 자동으로 `available` 복귀 후보가 된다. 단, 작성자가 이미 종결 처리한 경우 제외한다.
- 하나의 활성 예약이 존재하면 동일 매물에 다른 예약 확정은 불가하다.

### 17.5 거래 완료 확정 정책(초안)
- 기본안 A: 매물 작성자 1차 완료 처리 + 상대방 확인 요청
- 상대방이 일정 시간 내 이의 제기하지 않으면 자동 확정(가정)
- 상대방이 이의를 제기하면 상태를 `disputed_completion` 성격의 운영 검토 큐로 보낸다.
- MVP에서는 복잡도를 줄이기 위해 “작성자 완료 처리 후 상대방 후기/신고 가능” 구조로 시작하는 방안도 가능하며, 이는 정책 결정 필요

### 17.6 무응답/노쇼 처리
- 예약 시각 경과 후 양측은 노쇼 또는 불발 사유를 남길 수 있다.
- 노쇼 리포트는 즉시 제재가 아니라 누적 신뢰도/운영 검토 신호로 반영한다.
- 동일 사용자에게 짧은 기간 반복 노쇼 신고가 누적되면 경고/채팅 제한/매물 노출 제한 후보가 된다.

## 18. 알림 정책(초안)
### 18.1 알림 목표
- 사용자가 앱을 열지 않아도 거래 성사에 필요한 행동을 놓치지 않게 한다.
- 단순 홍보성 푸시보다 거래 진행 이벤트를 우선한다.

### 18.2 핵심 알림 이벤트
- 새 채팅 시작
- 새 메시지 수신
- 예약 제안 도착
- 예약 확정/수정/취소
- 예약 시간 임박
- 거래 상태 변경(`available -> reserved`, `reserved -> pending_trade`, `pending_trade -> completed` 등)
- 후기 작성 요청
- 신고 처리 결과 / 운영 경고
- 찜한 매물 가격 변경 또는 상태 복귀(확장)

### 18.3 발송 규칙
- 새 메시지 푸시는 사용자가 현재 해당 채팅방을 보고 있지 않을 때만 발송
- 예약 임박 알림은 최소 1회, 필요 시 24시간 전/1시간 전 2단계 발송
- 심야 시간(예: 23:00~08:00, 가정)에는 마케팅성 알림 금지, 거래 필수 알림만 허용
- 동일 이벤트의 중복 알림은 디바운스 처리 필요
- 앱 푸시 미수신 사용자는 앱 내 알림함에서 동일 이벤트를 확인 가능해야 함

### 18.4 알림 우선순위
- P1: 거래 진행 필수 이벤트(예약 확정, 예약 임박, 완료 요청)
- P2: 대화 이벤트(새 메시지, 새 문의)
- P3: 관심/리마인드 이벤트(가격 변경, 상태 재오픈, 후기 요청)
- P4: 운영 이벤트(경고, 신고 결과)

## 19. 화면별 요구사항 / IA 상세
### 19.1 홈
목적: 빠른 탐색 진입과 개인화 진입점 제공
- 주요 영역: 검색바, 서버 빠른 선택, 인기/최근 매물, 내 거래 상태 요약, 미확인 채팅 배지
- 핵심 액션: 검색, 매물 등록, 찜/채팅/내거래 진입
- 비회원 노출 범위는 정책 결정 필요

### 19.2 거래소 목록 화면
목적: 서버/아이템/상태 기반으로 빠르게 후보를 좁힌다.
- 필수 요소: 서버 필터, 카테고리, 검색어, 가격/상태/거래방식 필터, 정렬
- 카드 정보: 대표 이미지, 아이템명, 가격, 서버, 상태, 최근 활동, 작성자 신뢰 신호
- 행동: 찜, 상세 진입, 빠른 채팅
- 빈 상태: 검색 결과 없음, 서버 미선택, 필터 과다 적용

### 19.3 매물 상세 화면
목적: 거래 여부 판단과 즉시 행동 유도
- 필수 정보: 아이템/수량/옵션/가격/설명/이미지/서버/가능 시간/희망 장소/상태
- 판매자 정보: 닉네임, 최근 활동, 완료 거래 수, 후기 요약, 응답성 배지(가정)
- 핵심 CTA: 채팅 시작, 찜, 신고
- 상태별 UI: 예약중/거래대기/완료/취소에 따라 CTA 비활성 또는 대체 문구 제공

### 19.4 매물 등록/수정 화면
목적: 3분 내 등록 가능
- 단계형 또는 단일 폼 UX 후보
- 필수값 검증: 거래 유형, 서버, 아이템명, 가격 정책, 설명 최소 길이(가정)
- 선택값: 이미지, 접속 가능 시간, 희망 장소, 캐릭터명 공개 범위
- 저장 방식: 임시저장 여부는 확장 검토

### 19.5 채팅 화면
목적: 협의와 예약 확정을 한 화면 흐름 안에서 처리
- 메시지 타임라인 + 예약 카드 영역 결합
- 빠른 액션: 예약 제안, 상태 변경 요청, 완료 처리, 신고, 차단
- 시스템 메시지: 예약 생성/수정/취소, 상태 변경, 완료 요청 로그
- 입력 UX: 자주 쓰는 제안 문구, 장소/시간 템플릿

### 19.6 내 매물 화면
목적: 작성자가 매물 운영 상태를 빠르게 관리
- 탭: 진행중 / 예약중 / 완료 / 취소
- 카드 액션: 상태 변경, 수정, 끌어올리기(정책 결정), 채팅 보기
- 지표: 조회수, 찜수, 채팅수, 최근 활동

### 19.7 내 거래 화면
목적: 사용자가 본인이 참여 중인 거래를 매물 중심이 아닌 진행 중심으로 본다.
- 탭: 문의중 / 예약됨 / 거래대기 / 완료대기 / 종료
- 각 항목은 상대 사용자, 매물명, 약속 시간, 현재 액션 필요 여부를 보여줘야 함
- 채팅 미읽음과 예약 임박을 우선 노출

### 19.8 알림함
목적: 푸시를 놓친 경우에도 거래 이벤트 추적 가능
- 읽지 않음/전체 구분
- 이벤트 유형별 아이콘/분류
- 클릭 시 상세 화면 또는 채팅방으로 딥링크

### 19.9 프로필/신뢰 화면
목적: 사용자 신뢰 판단에 필요한 요약 정보 제공
- 완료 거래 수, 최근 후기, 노쇼/신고 경고 여부(노출 범위 정책 필요), 응답성, 가입 시점
- 후기 작성 이력, 받은 후기, 활동 서버 요약

### 19.10 관리자 운영 화면(백오피스)
- 신고 큐 목록
- 신고 상세(매물/채팅/사용자 이력/증빙)
- 조치 실행: 경고, 숨김, 차단, 계정 제한
- 감사 로그 및 조치 메모 저장

## 20. 데이터 모델 후보
### 20.1 핵심 엔티티
- User
- UserProfile
- Listing
- ListingImage
- Favorite
- ChatRoom
- ChatMessage
- Reservation
- TradeCompletion
- Review
- ReviewRatingDimension(확장)
- Report
- ModerationAction
- Notification
- BlockRelation
- AuditLog

### 20.2 관계 초안
- User 1:N Listing
- Listing 1:N ListingImage
- User N:N Listing (Favorite)
- Listing 1:N ChatRoom
- ChatRoom 1:N ChatMessage
- ChatRoom 1:N Reservation (활성 예약은 최대 1개)
- Listing 1:N Report
- User 1:N Report (reporter)
- User 1:N ModerationAction (target)
- Listing 1:0..1 TradeCompletion (MVP 기준 최종 종결 기록)
- TradeCompletion 1:0..2 Review

### 20.3 설계상 중요 포인트
- 매물 상태와 채팅방 상태를 분리 저장해야 한다.
- 거래 완료는 단순 매물 상태값 외에 별도 이벤트/엔티티로 보관하는 것이 운영 추적에 유리하다.
- 신고와 제재는 사용자/매물/채팅/메시지 단위 연결이 가능해야 한다.
- 알림은 발송 결과와 읽음 여부를 추적해야 한다.

## 21. API 후보 목록(초안)
### 21.1 인증/계정
- `POST /auth/login`
- `POST /auth/logout`
- `GET /me`
- `PATCH /me/profile`

### 21.2 매물
- `GET /listings`
- `POST /listings`
- `GET /listings/{listingId}`
- `PATCH /listings/{listingId}`
- `POST /listings/{listingId}/status`
- `POST /listings/{listingId}/favorite`
- `DELETE /listings/{listingId}/favorite`
- `GET /me/listings`

### 21.3 채팅
- `POST /listings/{listingId}/chats`
- `GET /chats`
- `GET /chats/{chatRoomId}`
- `GET /chats/{chatRoomId}/messages`
- `POST /chats/{chatRoomId}/messages`
- `POST /chats/{chatRoomId}/block`

### 21.4 예약/거래
- `POST /chats/{chatRoomId}/reservations`
- `PATCH /reservations/{reservationId}`
- `POST /reservations/{reservationId}/confirm`
- `POST /reservations/{reservationId}/cancel`
- `POST /reservations/{reservationId}/expire` (system/internal 가능)
- `POST /listings/{listingId}/complete`
- `POST /listings/{listingId}/cancel`
- `POST /listings/{listingId}/reopen` (정책상 미지원 가능)

### 21.5 후기/신고
- `POST /trade-completions/{completionId}/reviews`
- `GET /users/{userId}/reviews`
- `POST /reports`
- `GET /me/reports`

### 21.6 알림
- `GET /notifications`
- `POST /notifications/read`
- `POST /push-tokens`

### 21.7 관리자
- `GET /admin/reports`
- `GET /admin/reports/{reportId}`
- `POST /admin/reports/{reportId}/actions`
- `GET /admin/users/{userId}/history`
- `POST /admin/listings/{listingId}/hide`
- `POST /admin/users/{userId}/restrict`

## 22. 후기/신뢰도 정책(초안)
### 22.1 설계 원칙
- 신뢰도는 “절대적인 안전 보장”이 아니라 거래 판단 보조 신호로 제공한다.
- 조작이 쉬운 단일 점수보다, 근거가 보이는 다중 신호 구조를 우선한다.
- 신고 이력은 민감 정보이므로 공개 범위를 최소화하고 운영 판단용 데이터와 사용자 노출 데이터를 분리한다.

### 22.2 사용자 공개 신뢰 신호
- 완료 거래 수
- 최근 30/90일 완료 거래 수
- 받은 후기 수
- 긍정 후기 비율 또는 추천 여부
- 평균 응답 시간 구간(예: 빠름/보통/느림)
- 최근 접속 또는 최근 활동 시점
- 계정 생성 시점
- 본인 인증/추가 인증 여부(정책 결정 필요)

### 22.3 비공개 운영 신뢰 신호
- 누적 신고 건수 및 신고 유형 분포
- 노쇼 신고 누적 수
- 제재 이력(경고/일시정지/영구제한)
- 동일 기기/계정군 이상 패턴(확장)
- 비정상 채팅 생성/취소 반복 패턴

### 22.4 후기 구조
- 후기 작성 가능 조건: 거래 완료된 당사자만 작성 가능
- 후기 대상: 상대 사용자
- 후기 형식(MVP): 추천/비추천 2값 + 선택 코멘트
- 확장 형식: 응답속도, 약속준수, 설명정확성 등 다차원 평가
- 후기 수정 가능 여부: 작성 후 짧은 유예시간 내 1회 수정 허용 여부 결정 필요
- 후기 공개 시점: 양측 작성 완료 시 공개 또는 작성 마감 기한 후 공개

### 22.5 후기 작성 정책
- 거래 완료 후 일정 기간 내 작성 가능(예: 7일, 가정)
- 본인 자신에 대한 후기 작성 불가
- 동일 거래 완료 건당 최대 1회 작성
- 욕설/개인정보/협박 포함 후기는 운영 숨김 대상
- 거래와 무관한 정치/광고/외부 홍보성 후기는 금지

### 22.6 신뢰 배지 표현 방식(초안)
- 숫자 점수 대신 단계형 배지 우선 검토: `신규` / `거래중` / `거래경험 많음` / `신뢰 우수`
- 배지 산정에는 완료 거래 수, 최근 활동성, 후기 품질을 반영
- 제재/노쇼 누적은 공개 배지 하향 요인일 수 있으나, 직접적인 낙인 표현은 지양

## 23. 신고/분쟁/모더레이션 정책(초안)
### 23.1 신고 대상 단위
- 사용자
- 매물
- 채팅방
- 개별 메시지
- 후기

### 23.2 신고 유형
- 허위매물
- 사기 의심
- 거래 후 잠수/노쇼
- 욕설/괴롭힘/성희롱
- 스팸/광고
- 금지 품목 또는 운영정책 위반
- 개인정보 노출
- 기타

### 23.3 신고 접수 최소 요건
- 신고 유형 선택
- 설명 텍스트
- 대상 객체 식별값
- 필요 시 증빙 첨부(확장)
- 허위 신고 방지를 위한 신고자 식별 및 이력 저장

### 23.4 운영 처리 단계
1. 접수
2. 자동 분류/우선순위 부여
3. 증빙 확인
4. 임시 조치 여부 판단
5. 최종 조치 확정
6. 당사자 통지
7. 감사 로그 기록

### 23.5 임시 조치 정책
- 사기/불법/심각한 괴롭힘 의심 시 매물 임시 숨김 가능
- 반복 신고 누적 사용자는 신규 채팅/등록 일시 제한 가능
- 임시 조치는 확정 제재와 구분되어 저장되어야 함

### 23.6 제재 레벨(초안)
- L1 경고: 안내 및 재발 방지 고지
- L2 기능 제한: 채팅/매물등록/후기작성 일부 제한
- L3 기간 정지: 일정 기간 서비스 이용 정지
- L4 영구 제한: 반복 악성 행위 또는 중대한 정책 위반

### 23.7 분쟁 처리 원칙
- 운영자는 거래 성사 자체를 보증하지 않는다.
- 다만 플랫폼 내 기록(매물, 메시지, 예약, 완료 로그)을 기준으로 운영 판단을 내릴 수 있다.
- 금전/현물 반환 강제는 불가하나, 반복 악성 사용자 식별과 제재는 가능해야 한다.

### 23.8 개인정보/민감정보 정책
- 전화번호, 계좌번호, 실명 등 직접 식별 정보는 기본적으로 공개 프로필에 노출하지 않는다.
- 채팅 내 개인정보 교환 허용 범위는 별도 정책 필요
- 운영자 열람은 신고 처리 등 정당한 사유와 감사 로그 하에서만 허용

## 24. 상태별 액션 매트릭스(초안)
### 24.1 매물 상태별 기본 액션
| 매물 상태 | 작성자 가능 액션 | 상대 사용자 가능 액션 | 시스템 동작 |
|---|---|---|---|
| available | 수정, 예약확정, 취소, 숨김 | 채팅 시작, 찜, 신고 | 검색/노출 가능 |
| reserved | 예약해제, 거래대기 전환, 취소 | 기존 채팅 지속, 신고 | 신규 채팅 허용 여부 정책 결정 필요 |
| pending_trade | 완료 처리, 거래불발 처리 | 기존 채팅 지속, 노쇼/불발 표시 | 예약 임박/경과 알림 |
| completed | 후기 작성 유도, 기록 열람 | 후기 작성, 신고 | 신규 문의 차단 |
| cancelled | 재등록, 기록 열람 | 기록 열람(제한적) | 노출 중단 |

### 24.2 채팅방 상태별 기본 액션
| 채팅 상태 | 발신/수신 가능 | 예약 액션 | 기타 |
|---|---|---|---|
| open | 가능 | 예약 제안 가능 | 신고/차단 가능 |
| reservation_proposed | 가능 | 수정/확정/거절 가능 | 시스템 카드 노출 |
| reservation_confirmed | 가능 | 취소/시간조정 가능 | 매물 상태 연동 |
| trade_due | 가능 | 완료/불발 처리 가능 | 임박 알림 우선 |
| deal_completed | 제한적 또는 읽기전용(정책 결정) | 불가 | 후기 진입 |
| deal_cancelled | 제한적 또는 읽기전용(정책 결정) | 불가 | 재문의 유도 여부 정책 결정 |
| report_locked | 제한 | 불가 | 운영 검토 중 배너 |

### 24.3 예약 상태별 기본 액션
| 예약 상태 | 제안자 | 상대방 | 시스템 |
|---|---|---|---|
| proposed | 수정, 취소 | 수락, 거절, 대안 제시 | 만료 타이머 작동 |
| confirmed | 취소 요청, 시간조정 요청 | 취소 요청, 시간조정 요청 | 매물 reserved/pending_trade 연동 |
| expired | 재제안 | 재제안 | 자동 종료 로그 |
| cancelled | 재제안 | 재제안 | 사유 코드 저장 |
| fulfilled | 후기 작성 | 후기 작성 | 완료 로그 생성 |
| no_show_reported | 소명/신고 | 소명/신고 | 운영 검토 큐 적재 가능 |

## 25. 분석 이벤트 및 KPI(초안)
### 25.1 핵심 퍼널
1. 매물 조회
2. 상세 진입
3. 채팅 시작
4. 예약 제안
5. 예약 확정
6. 거래 완료
7. 후기 작성

### 25.2 필수 이벤트 후보
- `listing_create`
- `listing_publish`
- `listing_view`
- `listing_favorite`
- `chat_room_create`
- `chat_message_send`
- `reservation_propose`
- `reservation_confirm`
- `reservation_cancel`
- `listing_status_change`
- `trade_complete_request`
- `trade_complete_confirm`
- `review_submit`
- `report_submit`
- `notification_open`

### 25.3 이벤트 공통 속성
- userId
- listingId
- chatRoomId(optional)
- reservationId(optional)
- listingType
- serverId
- itemCategoryId
- sourceScreen
- timestamp

### 25.4 제품 KPI(초안)
- 매물 등록 완료율
- 매물 상세 대비 채팅 시작 전환율
- 채팅 시작 대비 예약 확정 전환율
- 예약 확정 대비 거래 완료 전환율
- 거래 완료 대비 후기 작성률
- 평균 첫 응답 시간
- 매물 등록 후 첫 문의까지 걸린 시간
- 신고율 / 완료 거래 100건당 신고 건수
- 노쇼율 / 예약 100건당 노쇼 신고 건수

### 25.5 운영 KPI(초안)
- 신고 최초 응답 시간
- 신고 처리 완료 시간
- 허위/악성 신고 비율
- 제재 후 재위반율
- 숨김/차단 조치의 이의 제기율

## 26. 출시 범위 / 릴리즈 단계(초안)
### 26.1 MVP 필수
- 회원가입/로그인
- 매물 등록/수정/목록/상세
- 검색/필터
- 찜
- 1:1 채팅
- 예약 제안/확정/취소
- 매물 상태 변경
- 거래 완료 및 기본 후기
- 신고 접수 및 관리자 기본 처리
- 앱 내 알림함

### 26.2 Post-MVP 우선 후보
- 푸시 알림 고도화
- 신뢰 배지 세분화
- 차단/뮤트 고도화
- 매물 끌어올리기/만료 정책
- 추천 검색어/시세 보조 기능
- 예약 자동 리마인드 고도화

### 26.3 장기 확장 후보
- 서버/아이템 표준 카탈로그 고도화
- 사기 패턴 탐지 자동화
- 반복 거래자용 일괄 관리 도구
- 지도 기반 오프라인 만남 포인트 추천
- 외부 커뮤니티 연동 또는 콘텐츠 피드


## 27. 엔티티별 필수/선택 컬럼 표
### 27.1 User / UserProfile
| 엔티티 | 컬럼 | 필수 여부 | 타입/형태 | 설명 |
|---|---|---|---|---|
| User | userId | 필수 | UUID/ULID | 내부 식별자 |
| User | loginProvider | 필수 | enum | 소셜/휴대폰 등 로그인 수단 |
| User | loginProviderUserKey | 필수 | string | 외부 식별자 매핑 키 |
| User | accountStatus | 필수 | enum | active / restricted / suspended / withdrawn |
| User | role | 필수 | enum | user / moderator / admin |
| User | lastLoginAt | 선택 | datetime | 최근 로그인 시각 |
| User | createdAt | 필수 | datetime | 가입 시각 |
| User | withdrawnAt | 선택 | datetime | 탈퇴 시각 |
| UserProfile | userId | 필수 | FK | User 참조 |
| UserProfile | nickname | 필수 | string | 서비스 표시 닉네임 |
| UserProfile | avatarUrl | 선택 | string | 프로필 이미지 |
| UserProfile | introduction | 선택 | string | 자기소개 |
| UserProfile | primaryServerId | 선택 | FK/string | 주 활동 서버 |
| UserProfile | responseBadge | 선택 | enum | fast / normal / slow |
| UserProfile | trustBadge | 선택 | enum | 신규/거래중/거래경험많음/신뢰우수 |
| UserProfile | completedTradeCount | 필수 | integer | 완료 거래 수 캐시 |
| UserProfile | positiveReviewCount | 필수 | integer | 긍정 후기 수 캐시 |
| UserProfile | lastActiveAt | 선택 | datetime | 최근 활동 시각 |

### 27.2 Listing
| 컬럼 | 필수 여부 | 타입/형태 | 설명 |
|---|---|---|---|
| listingId | 필수 | UUID/ULID | 매물 식별자 |
| listingType | 필수 | enum | sell / buy |
| authorUserId | 필수 | FK | 작성자 |
| serverId | 필수 | FK/string | 거래 서버 |
| categoryId | 필수 | FK/string | 아이템 카테고리 |
| itemName | 필수 | string | 대표 아이템명 |
| title | 필수 | string | 리스트 표시 제목 |
| description | 필수 | text | 상세 설명 |
| priceType | 필수 | enum | fixed / negotiable / offer |
| priceAmount | 선택 | decimal/integer | 가격, offer면 null 허용 가능 |
| currencyType | 선택 | enum | adena / KRW / mixed 등, 정책 결정 필요 |
| quantity | 필수 | decimal/integer | 수량 |
| enhancementLevel | 선택 | integer | 강화 수치 |
| optionsText | 선택 | text | 자유 옵션 텍스트 |
| tradeMethod | 필수 | enum | in_game / offline_pc_bang / either |
| preferredMeetingAreaText | 선택 | string | 오프라인 또는 인게임 접선 지역 |
| availableTimeText | 선택 | string | 가능 시간 자유기입 |
| status | 필수 | enum | available / reserved / pending_trade / completed / cancelled |
| visibility | 필수 | enum | public / hidden / blocked |
| reservedChatRoomId | 선택 | FK | 현재 우선 진행 중인 채팅방 |
| lastActivityAt | 필수 | datetime | 정렬/노출용 활동 시각 |
| viewCount | 필수 | integer | 조회수 캐시 |
| favoriteCount | 필수 | integer | 찜 수 캐시 |
| chatCount | 필수 | integer | 채팅방 수 캐시 |
| createdAt | 필수 | datetime | 생성 시각 |
| updatedAt | 필수 | datetime | 수정 시각 |
| expiresAt | 선택 | datetime | 자동 만료 정책용 |
| completionAt | 선택 | datetime | 완료 시각 |

### 27.3 ChatRoom / ChatMessage
| 엔티티 | 컬럼 | 필수 여부 | 타입/형태 | 설명 |
|---|---|---|---|---|
| ChatRoom | chatRoomId | 필수 | UUID/ULID | 채팅방 식별자 |
| ChatRoom | listingId | 필수 | FK | 매물 참조 |
| ChatRoom | sellerUserId | 필수 | FK | 판매 역할 사용자 |
| ChatRoom | buyerUserId | 필수 | FK | 구매 역할 사용자 |
| ChatRoom | chatStatus | 필수 | enum | open / reservation_proposed / reservation_confirmed / trade_due / deal_completed / deal_cancelled / report_locked |
| ChatRoom | lastMessageId | 선택 | FK | 최근 메시지 |
| ChatRoom | lastMessageAt | 선택 | datetime | 최근 메시지 시각 |
| ChatRoom | unreadCountForSeller | 필수 | integer | 판매자 미읽음 캐시 |
| ChatRoom | unreadCountForBuyer | 필수 | integer | 구매자 미읽음 캐시 |
| ChatRoom | closedAt | 선택 | datetime | 종료 시각 |
| ChatMessage | messageId | 필수 | UUID/ULID | 메시지 식별자 |
| ChatMessage | chatRoomId | 필수 | FK | 채팅방 참조 |
| ChatMessage | senderUserId | 필수 | FK | 발신자 |
| ChatMessage | messageType | 필수 | enum | text / system / reservation_card / image(확장) |
| ChatMessage | bodyText | 선택 | text | 본문 |
| ChatMessage | metadataJson | 선택 | json | 예약/상태 변경 부가정보 |
| ChatMessage | sentAt | 필수 | datetime | 발송 시각 |
| ChatMessage | deletedAt | 선택 | datetime | 소프트 삭제 시각 |

### 27.4 Reservation / TradeCompletion / Review
| 엔티티 | 컬럼 | 필수 여부 | 타입/형태 | 설명 |
|---|---|---|---|---|
| Reservation | reservationId | 필수 | UUID/ULID | 예약 식별자 |
| Reservation | listingId | 필수 | FK | 매물 참조 |
| Reservation | chatRoomId | 필수 | FK | 채팅방 참조 |
| Reservation | proposerUserId | 필수 | FK | 제안자 |
| Reservation | counterpartUserId | 필수 | FK | 상대방 |
| Reservation | reservationStatus | 필수 | enum | proposed / confirmed / expired / cancelled / fulfilled / no_show_reported |
| Reservation | scheduledAt | 필수 | datetime | 약속 기준 시각 |
| Reservation | timeWindowText | 선택 | string | 시간 범위 보조 텍스트 |
| Reservation | meetingType | 필수 | enum | in_game / offline_pc_bang / either |
| Reservation | meetingPointText | 필수 | string | 접선 장소 |
| Reservation | serverId | 선택 | FK/string | 인게임 서버 |
| Reservation | characterNameA | 선택 | string | 캐릭터명 A |
| Reservation | characterNameB | 선택 | string | 캐릭터명 B |
| Reservation | noteToCounterparty | 선택 | text | 전달 메모 |
| Reservation | expiresAt | 선택 | datetime | 제안 만료 시각 |
| Reservation | confirmedAt | 선택 | datetime | 확정 시각 |
| Reservation | cancelledAt | 선택 | datetime | 취소 시각 |
| Reservation | fulfilledAt | 선택 | datetime | 실제 이행 시각 |
| Reservation | cancellationReasonCode | 선택 | enum/string | 취소 사유 |
| TradeCompletion | completionId | 필수 | UUID/ULID | 완료 기록 식별자 |
| TradeCompletion | listingId | 필수 | FK | 매물 참조 |
| TradeCompletion | chatRoomId | 필수 | FK | 실제 완료 채팅방 |
| TradeCompletion | completedByUserId | 필수 | FK | 완료 처리자 |
| TradeCompletion | completionStatus | 필수 | enum | requested / confirmed / disputed / closed |
| TradeCompletion | completedAt | 필수 | datetime | 완료 요청 시각 |
| TradeCompletion | confirmedAt | 선택 | datetime | 최종 확정 시각 |
| Review | reviewId | 필수 | UUID/ULID | 후기 식별자 |
| Review | completionId | 필수 | FK | 완료 기록 참조 |
| Review | reviewerUserId | 필수 | FK | 작성자 |
| Review | revieweeUserId | 필수 | FK | 대상자 |
| Review | recommendation | 필수 | boolean/enum | recommend / not_recommend |
| Review | commentText | 선택 | text | 후기 본문 |
| Review | visibilityStatus | 필수 | enum | visible / hidden / pending_moderation |
| Review | createdAt | 필수 | datetime | 생성 시각 |

### 27.5 Report / ModerationAction / Notification
| 엔티티 | 컬럼 | 필수 여부 | 타입/형태 | 설명 |
|---|---|---|---|---|
| Report | reportId | 필수 | UUID/ULID | 신고 식별자 |
| Report | reporterUserId | 필수 | FK | 신고자 |
| Report | targetType | 필수 | enum | user / listing / chat_room / message / review |
| Report | targetId | 필수 | string | 대상 식별자 |
| Report | reportReasonCode | 필수 | enum | 신고 유형 |
| Report | descriptionText | 필수 | text | 상세 설명 |
| Report | reportStatus | 필수 | enum | submitted / triaged / investigating / resolved / rejected |
| Report | priority | 필수 | enum | P1 / P2 / P3 / P4 |
| Report | createdAt | 필수 | datetime | 접수 시각 |
| Report | resolvedAt | 선택 | datetime | 해결 시각 |
| ModerationAction | actionId | 필수 | UUID/ULID | 조치 식별자 |
| ModerationAction | reportId | 선택 | FK | 연결 신고 |
| ModerationAction | targetUserId | 선택 | FK | 대상 사용자 |
| ModerationAction | targetListingId | 선택 | FK | 대상 매물 |
| ModerationAction | actionType | 필수 | enum | warn / hide_listing / restrict_chat / suspend / ban |
| ModerationAction | actionReason | 필수 | text | 조치 사유 |
| ModerationAction | startsAt | 필수 | datetime | 시작 시각 |
| ModerationAction | endsAt | 선택 | datetime | 종료 시각 |
| ModerationAction | createdByAdminId | 필수 | FK | 처리 운영자 |
| Notification | notificationId | 필수 | UUID/ULID | 알림 식별자 |
| Notification | userId | 필수 | FK | 수신자 |
| Notification | notificationType | 필수 | enum | chat / reservation / status / review / report / system |
| Notification | title | 필수 | string | 제목 |
| Notification | body | 필수 | string | 본문 |
| Notification | deepLink | 선택 | string | 앱 이동 경로 |
| Notification | readAt | 선택 | datetime | 읽음 시각 |
| Notification | deliveredAt | 선택 | datetime | 발송 시각 |

## 28. 검색/정렬/노출 랭킹 규칙(초안)
### 28.1 검색 목표
- 사용자가 원하는 서버/아이템 매물을 빠르게 좁히게 한다.
- 이미 거래 불가능한 매물보다 실제로 대화/완료 가능성이 높은 매물을 우선 노출한다.
- 과도한 끌어올리기나 스팸 등록이 상단을 독점하지 못하게 한다.

### 28.2 기본 필터
- 서버: 단일 선택 기본, 다중 선택은 확장 후보
- 거래 유형: 판매 / 구매
- 카테고리
- 아이템명 키워드
- 가격 범위
- 거래 방식(in-game / offline / either)
- 상태(기본값은 `available`, 선택 시 `reserved` 포함 가능)
- 최근 활동 기간(예: 최근 24시간/3일/7일)

### 28.3 기본 정렬 옵션
- 추천순(기본)
- 최신순
- 가격 낮은순 / 높은순
- 문의 많은순(확장)
- 곧 거래 가능순(예약 임박/활동성 기반, 확장)

### 28.4 추천순 랭킹 신호(초안)
추천순은 아래 신호의 조합으로 계산한다.
1. 검색 적합도: 서버 일치, 아이템명 exact/partial match, 카테고리 일치
2. 거래 가능성: `available` 우선, 최근 응답/활동 있는 매물 가산
3. 신뢰 신호: 작성자 배지, 완료 거래 수, 최근 신고/제재 여부
4. 콘텐츠 품질: 제목/설명 충실도, 이미지 존재, 옵션 정보 존재
5. 신선도: 최근 등록/수정/상태 변경 시점
6. 패널티: 반복 유사 매물, 신고 검토 중, 비정상 클릭/찜 패턴

### 28.5 상태별 노출 원칙
- `available`: 목록 기본 노출 대상
- `reserved`: 기본 검색 결과에서는 후순위, 필터 선택 시 노출
- `pending_trade`: 기본 검색 결과 제외, 직접 링크/내 거래에서는 노출
- `completed`, `cancelled`: 공개 목록 제외, 작성자/참여자 기록 화면에서만 노출
- `hidden`, `blocked`: 운영/작성자 전용 뷰에서만 노출

### 28.6 스팸/도배 완화 규칙
- 동일 작성자의 유사 제목/아이템/가격 조합이 짧은 시간 반복되면 랭킹 감점
- 동일 작성자의 상위 노출 개수 제한을 둘 수 있다(예: 같은 서버/카테고리 상위 N개, 가정)
- 신고 검토 중인 매물은 검색 노출 가중치를 낮추거나 임시 제외할 수 있다.

## 29. 운영 SLA / 우선순위 / 이의제기 정책(초안)
### 29.1 운영 우선순위 정의
- **P1 긴급**: 사기 의심, 불법 거래, 심각한 괴롭힘/성희롱, 개인정보 대량 노출
- **P2 높음**: 반복 노쇼, 욕설/협박, 대량 스팸, 허위매물 반복 등록
- **P3 보통**: 일반 허위 정보, 후기 분쟁, 경미한 운영정책 위반
- **P4 낮음**: 단순 문의, 기능 오남용 의심 but 증빙 부족, 정책 해석 요청

### 29.2 목표 SLA
| 우선순위 | 1차 확인 목표 | 임시 조치 목표 | 최종 처리 목표 |
|---|---|---|---|
| P1 | 1시간 이내 | 2시간 이내 | 24시간 이내 |
| P2 | 6시간 이내 | 24시간 이내 | 72시간 이내 |
| P3 | 24시간 이내 | 필요 시만 | 5영업일 이내 |
| P4 | 48시간 이내 | 없음 | 7영업일 이내 |

### 29.3 야간/비영업 시간 정책
- 24/7 완전 대응을 전제로 하지 않는다.
- 다만 P1은 비영업 시간에도 최소 임시 숨김/잠금 조치가 가능해야 한다.
- P2 이하의 상세 검토는 운영 시간 내 순차 처리할 수 있다.

### 29.4 증빙 우선순위
- 플랫폼 내부 로그(채팅/예약/상태 변경/신고 이력)를 1차 기준으로 본다.
- 외부 메신저 캡처, 계좌 이체 내역 등은 보조 자료로만 활용한다.
- 증빙 불충분 시 조치 보류 또는 경고 수준에 그칠 수 있다.

### 29.5 이의제기 정책
- 제재/숨김 대상자는 결과 통지 후 일정 기간 내 이의제기 가능(예: 7일, 가정)
- 이의제기는 1회 기본 접수 원칙으로 하고, 신규 증빙이 있을 때만 재심 허용
- 동일 사안 반복 이의제기는 제한 가능
- 이의제기 처리 결과 역시 감사 로그에 남겨야 한다.

### 29.6 운영자 기록 원칙
- 모든 숨김/제재/복구 조치는 사유 코드와 자유 메모를 함께 남긴다.
- 운영자 식별자, 시각, 대상, 근거 링크를 저장한다.
- 고위험 조치(L3 이상)는 가능하면 2인 검토 또는 사후 승인 로그를 남긴다.

## 30. API 요청/응답 예시(초안)
### 30.1 예약 생성
`POST /chats/{chatRoomId}/reservations`

Request 예시:
```json
{
  "scheduledAt": "2026-03-14T21:00:00+09:00",
  "timeWindowText": "21:00~21:30",
  "meetingType": "in_game",
  "serverId": "ken-01",
  "meetingPointText": "기란 마을 창고 앞",
  "characterName": "린클상인",
  "noteToCounterparty": "도착 5분 전에 채팅 드릴게요"
}
```

Response 예시:
```json
{
  "reservationId": "res_123",
  "chatRoomId": "chat_123",
  "reservationStatus": "proposed",
  "scheduledAt": "2026-03-14T21:00:00+09:00",
  "expiresAt": "2026-03-14T18:00:00+09:00",
  "listingStatusImpact": "none"
}
```

### 30.2 예약 확정
`POST /reservations/{reservationId}/confirm`

Response 예시:
```json
{
  "reservationId": "res_123",
  "reservationStatus": "confirmed",
  "confirmedAt": "2026-03-13T10:12:00+09:00",
  "chatStatus": "reservation_confirmed",
  "listingStatus": "reserved"
}
```

### 30.3 거래 완료 요청
`POST /listings/{listingId}/complete`

Request 예시:
```json
{
  "chatRoomId": "chat_123",
  "note": "거래 완료했습니다"
}
```

Response 예시:
```json
{
  "completionId": "comp_123",
  "listingId": "listing_123",
  "completionStatus": "requested",
  "listingStatus": "completed",
  "counterpartyReviewEligible": true
}
```

### 30.4 신고 접수
`POST /reports`

Request 예시:
```json
{
  "targetType": "chat_room",
  "targetId": "chat_123",
  "reportReasonCode": "abuse",
  "descriptionText": "욕설과 반복적인 비방 메시지를 보냈습니다."
}
```

Response 예시:
```json
{
  "reportId": "rep_123",
  "reportStatus": "submitted",
  "priority": "P2",
  "nextAction": "moderation_review"
}
```

## 31. 비회원 공개 정책(초안)
### 31.1 정책 목표
- 검색 유입을 확보하되, 거래 당사자 보호와 스팸 유입 방지를 우선한다.
- 비회원에게는 “탐색 가능한 정보”만 공개하고, “접촉/신뢰 판단/직접 식별” 정보는 회원 이후 단계로 제한한다.

### 31.2 공개 레벨 정의
| 정보 항목 | 비회원 | 로그인 회원 | 거래 상대 확정/예약 단계 |
|---|---|---|---|
| 매물 제목/아이템명/카테고리 | 공개 | 공개 | 공개 |
| 서버/거래 방식/상태 | 공개 | 공개 | 공개 |
| 가격/수량/설명 요약 | 공개 | 공개 | 공개 |
| 상세 설명 전문 | 부분 공개 또는 전체 공개(정책 결정) | 공개 | 공개 |
| 이미지 | 대표 1장 또는 워터마크 버전 공개 가능 | 공개 | 공개 |
| 작성자 닉네임 | 마스킹 또는 비공개 후보 | 공개 | 공개 |
| 완료 거래 수/후기 요약 | 요약값만 제한 공개 가능 | 공개 | 공개 |
| 최근 활동 시각 | 상대 시간 단위로만 제한 공개 | 공개 | 공개 |
| 채팅 시작 / 예약 / 찜 | 불가 | 가능 | 가능 |
| 캐릭터명/직접 연락처 | 비공개 | 비공개 | 정책 허용 범위 내 제한 공개 |

### 31.3 비회원 허용 액션
- 매물 목록 조회
- 매물 상세 제한 조회
- 검색/필터 사용
- 회원가입/로그인 유도 진입
- 신고, 채팅, 찜, 후기 작성, 예약 제안은 불가

### 31.4 SEO/공개 노출 원칙
- 공개 매물은 검색엔진 색인 허용 여부를 별도 정책으로 관리한다.
- 색인 허용 시에도 닉네임, 캐릭터명, 연락처, 채팅 내용은 색인 대상에서 제외해야 한다.
- 완료/취소 매물은 공개 인덱스에서 제거하거나 `noindex` 처리하는 방안을 우선 검토한다.

### 31.5 비회원 전환 UX 원칙
- 채팅 시작, 찜, 상세 정보 확장 시 로그인 바텀시트/모달로 전환시킨다.
- 로그인 이후 원래 보던 매물/필터/스크롤 위치를 최대한 복원해야 한다.
- 비회원 유입이 많은 화면에서는 “거래 안전/후기/예약 기능은 로그인 후 제공” 메시지를 명확히 노출한다.

## 32. 개인정보/외부 연락처 교환 정책(초안)
### 32.1 정책 목표
- 플랫폼 밖으로 대화가 바로 이탈하는 것을 줄여, 신고/분쟁 처리 가능한 기록을 최대한 플랫폼 내에 남긴다.
- 동시에 실제 거래 성사를 위해 필요한 최소한의 식별 정보 교환은 과도하게 막지 않는다.

### 32.2 기본 원칙
- 전화번호, 카카오톡 ID, 오픈채팅 링크, 계좌번호, 실명, SNS 아이디 등 직접 식별/외부 이동 정보는 기본적으로 공개 프로필 및 매물 본문에 금지한다.
- 채팅에서도 초기 협의 단계에서는 외부 연락처 교환을 제한하거나 경고한다.
- 예약 확정 또는 거래 직전 단계에서만 제한적으로 허용하는 정책을 기본안으로 둔다.

### 32.3 단계별 허용 기준(기본안)
| 단계 | 전화번호/메신저 ID | 캐릭터명 | 계좌번호 | 비고 |
|---|---|---|---|---|
| 매물 공개 상태 | 금지 | 선택 공개 가능 | 금지 | 본문/이미지 OCR 포함 차단 대상 |
| 일반 채팅(open) | 제한 또는 경고 후 차단 | 공유 가능 | 금지 | 플랫폼 내 협의 우선 |
| 예약 확정 후 | 제한적 허용 가능 | 공유 권장 | 금지 또는 강한 경고 | 실거래 접선 목적 |
| 분쟁/신고 처리 | 사용자 간 노출 아님 | 필요 시 운영 열람 | 필요 시 운영 열람 | 감사 로그 전제 |

### 32.4 자동 감지/제어 정책
- 메시지/설명/이미지 OCR 결과에서 전화번호, 메신저 URL, 오픈채팅 초대 링크, 계좌번호 패턴을 감지할 수 있어야 한다.
- 감지 시 처리 옵션:
  1. 전송 전 경고 후 수정 유도
  2. 민감정보 일부 마스킹 후 전송
  3. 정책 위반 수준이 높으면 전송 차단 및 운영 로그 적재
- 반복 위반자는 채팅 기능 제한 후보가 된다.

### 32.5 운영 예외
- 운영자가 신고 처리, 법적 요청 대응, 계정 본인 확인 등 정당한 사유로만 제한 열람 가능해야 한다.
- 운영자 열람 시 대상, 열람 사유, 시각을 감사 로그에 남긴다.

## 33. 입력 validation / 제한 규칙(초안)
### 33.1 매물 등록 validation
- 제목: 최소 4자~최대 80자(가정)
- 설명: 최소 10자~최대 2,000자(가정)
- 아이템명: 최소 1자~최대 50자
- 가격:
  - `fixed`, `negotiable`일 때 필수
  - `offer`일 때 null 허용 가능
  - 0 이하 금지
  - 비정상 고가/저가 임계치는 경고 또는 운영 탐지 대상
- 수량: 0 초과 필수, 소수 허용 여부는 품목 정책에 따름
- 이미지 수: 0~10장(가정), 파일당 최대 용량 제한 필요
- 동일 사용자의 활성 매물 수 상한은 운영정책으로 별도 관리 가능

### 33.2 예약 validation
- `scheduledAt`은 현재 시각 이후만 허용
- 예약 제안 최소 lead time 필요(예: 현재로부터 10분 이후, 가정)
- 예약 가능 최대 범위 필요(예: 14일 이내, 가정)
- `meetingPointText`는 최소 2자 이상
- `meetingType=in_game`일 때 서버 정보 또는 인게임 위치 정보 중 최소 1개 필요
- 이미 활성 예약이 있는 매물에는 신규 예약 확정 불가

### 33.3 메시지/채팅 제한
- 단일 메시지 길이 제한 필요(예: 1,000자, 가정)
- 짧은 시간 반복 전송, 동일 문구 반복, 다중 채팅방 복붙 전송은 스팸 탐지 대상
- 차단 관계가 성립하면 신규 메시지 전송 금지
- 신고 잠금(`report_locked`) 상태에서는 읽기 전용 또는 운영 안내 메시지만 허용

### 33.4 상태 변경 validation
- `completed`, `cancelled` 상태로의 변경 시 사유 또는 관련 채팅방 지정이 필요할 수 있다.
- `reserved` 전환 시 활성 채팅방/예약 연결값이 필수다.
- `pending_trade` 전환 시 확정 예약 존재가 필수다.
- 시스템 자동 상태 변경과 사용자 수동 상태 변경은 actorType을 분리 저장해야 한다.

## 34. 비기능 요구사항 상세(SLI/SLO/보안/보관)
### 34.1 성능 목표(초안)
- 목록 API p95 응답시간: 800ms 이하
- 상세 API p95 응답시간: 700ms 이하
- 채팅 메시지 송신 후 상대 표시까지 p95 3초 이하(실시간 또는 준실시간 기준)
- 알림 이벤트 생성 후 앱 내 알림함 반영까지 p95 10초 이하

### 34.2 가용성 목표(초안)
- 사용자 핵심 거래 기능(목록/상세/채팅/예약)의 월간 가용성 목표: 99.5% 이상
- 관리자 신고 조회/처리 기능의 월간 가용성 목표: 99.0% 이상
- 단, 사전 공지된 점검 시간은 제외한다.

### 34.3 데이터 일관성 요구
- 매물 상태, 활성 예약, 완료 기록 간 모순이 발생하지 않도록 트랜잭션 경계 또는 보상 로직이 필요하다.
- 동일 매물에 대한 동시 예약 확정 요청은 1건만 성공해야 한다.
- 미읽음 수, 찜 수, 조회수 캐시는 지연 허용 가능하나 최종 일관성을 보장해야 한다.

### 34.4 보안 요구사항
- 모든 인증 세션과 주요 쓰기 API는 사용자 식별 가능한 감사 로그를 남긴다.
- 관리자 API는 일반 사용자 API와 권한/네트워크/감사 정책을 분리한다.
- 비밀번호를 직접 운영하지 않는 경우에도 토큰 탈취 방지, 세션 만료, 기기 관리 정책이 필요하다.
- 신고 증빙, 채팅 로그, 개인정보성 필드는 저장 시 암호화 또는 접근 통제가 필요하다.

### 34.5 보관/삭제 정책(초안)
- 채팅/예약/상태 변경/신고 로그는 분쟁 대응을 위해 최소 보관 기간이 필요하다(예: 1년, 가정).
- 탈퇴 사용자의 공개 프로필 정보는 비식별화하되, 운영상 필요한 거래/신고 로그는 정책에 따라 별도 보관 가능하다.
- 완료/취소 매물은 사용자 기록 화면에서 조회 가능하되, 공개 목록 노출은 중단한다.

### 34.6 확장성 고려사항
- 서버/아이템 카탈로그가 확장되어도 검색 성능이 유지되도록 색인 전략이 필요하다.
- 채팅 메시지와 알림은 높은 쓰기 빈도를 고려해 별도 스토리지/큐 설계를 검토한다.
- 매물 목록 조회와 채팅/알림 처리 부하는 분리 확장 가능해야 한다.

## 35. 다음 구체화 대상
1. 후기 공개 타이밍/수정 가능 정책 확정안
2. 색인/SEO 공개 여부 정책 확정
3. 푸시 알림 채널별 우선순위 및 opt-in 정책
4. 검색 추천/자동완성/오타 보정 정책
5. 운영 지표 대시보드와 이상징후 탐지 룰
6. 정산/에스크로 미도입 상태에서의 법적 고지 문구 확정

## 36. 관리자 백오피스 권한 및 작업 절차(초안)
### 36.1 운영 권한 레벨
| 역할 | 주요 권한 | 제한 사항 |
|---|---|---|
| CS Operator | 신고 열람, 기본 메모 작성, P3/P4 1차 분류 | 숨김/제재 확정 불가 |
| Moderator | 매물 임시 숨김, 채팅 잠금, 경고, P1/P2 1차 조치 | 영구 정지, 민감정보 전체 열람은 제한 |
| Senior Moderator | 기간 정지, 이의제기 재검토, 고위험 복구 승인 | 시스템 설정 변경 불가 |
| Admin | 모든 운영 조치, 정책 설정, 감사 로그 열람 | 최소 권한 원칙 적용 대상 |

### 36.2 신고 큐 화면 요구사항
- 기본 컬럼: 접수시각, 우선순위, 신고 유형, 대상 유형, 대상 식별 요약, 신고자, 현재 상태, 담당자
- 필터: 우선순위, 상태, 신고 유형, 반복 신고 대상 여부, 접수 기간
- 정렬 기본값: 우선순위 높은순 + 접수 오래된순
- 대량 작업은 P3/P4 범위에서만 허용하고, P1/P2는 개별 검토를 기본으로 한다.

### 36.3 신고 상세 화면 처리 절차
1. 대상 요약 확인: 사용자/매물/채팅/메시지 기본 정보 확인
2. 내부 로그 확인: 관련 채팅, 예약, 상태 변경, 과거 신고/제재 이력 확인
3. 위험도 판정: 즉시 차단 필요한지, 임시 숨김이 필요한지 판단
4. 조치 실행: 경고/임시숨김/채팅잠금/기간정지/영구제한 중 선택
5. 사유 기록: 코드 + 자유 메모 + 근거 링크 저장
6. 통지: 대상자 및 신고자에게 결과/상태를 정책 범위 내 통지
7. 사후 추적: 재발 여부, 이의제기 가능 기간, 후속 검토 일자 등록

### 36.4 운영 화면별 최소 기능
- **사용자 상세**: 기본 프로필, 완료 거래 수, 최근 활동, 최근 신고/제재 이력, 연결된 활성 매물/채팅
- **매물 상세**: 현재 상태, 변경 이력, 연결된 채팅방 수, 예약/완료 이력, 숨김/복구 액션
- **채팅 상세**: 메시지 타임라인, 시스템 이벤트, 예약 카드, 민감정보 탐지 로그, 잠금 액션
- **이의제기 상세**: 원조치, 제출 사유, 신규 증빙, 재판정 결과, 승인자 기록
- **감사 로그**: 운영자별 액션, 열람 사유, 대상, 시각, 결과, diff 요약

### 36.5 민감정보 열람 통제
- 전화번호/계좌번호/직접 식별 정보가 감지된 경우 기본적으로 마스킹 상태로 노출한다.
- 민감정보 원문 열람은 `Senior Moderator` 이상 + 열람 사유 입력 + 감사 로그 저장이 필요하다.
- 동일 사건 처리에 불필요한 민감정보는 기본 화면에 표시하지 않는다.

### 36.6 복구/오조치 대응
- 임시 숨김/제재는 복구 API 또는 복구 액션을 별도로 가져야 하며, 덮어쓰기식 수정이 아닌 원조치 + 복구조치의 이력형 저장을 원칙으로 한다.
- 오조치 복구 시 사용자에게 사유를 통지하고, 검색 랭킹/신뢰도에 미친 불이익 복원 여부를 검토해야 한다.

## 37. API 에러 모델 / 권한 / 상태 충돌 규칙(초안)
### 37.1 공통 응답 원칙
- 성공/실패 모두 기계 처리 가능한 코드와 사용자 표시 가능한 메시지를 함께 제공한다.
- 에러 응답은 최소한 `code`, `message`, `requestId`, `details(optional)`를 포함한다.
- 검증 오류는 필드 단위 에러 목록을 포함할 수 있어야 한다.

에러 응답 예시:
```json
{
  "code": "LISTING_STATUS_CONFLICT",
  "message": "현재 매물 상태에서는 예약을 확정할 수 없습니다.",
  "requestId": "req_123",
  "details": {
    "listingId": "listing_123",
    "currentStatus": "completed"
  }
}
```

### 37.2 대표 에러 코드 후보
| 코드 | HTTP 후보 | 설명 |
|---|---|---|
| UNAUTHORIZED | 401 | 로그인 필요 |
| FORBIDDEN | 403 | 대상 리소스 접근 권한 없음 |
| VALIDATION_ERROR | 422 | 입력값 검증 실패 |
| LISTING_NOT_FOUND | 404 | 매물 없음 또는 비공개 |
| CHAT_NOT_FOUND | 404 | 채팅방 없음 또는 접근 불가 |
| RESERVATION_NOT_FOUND | 404 | 예약 없음 |
| LISTING_STATUS_CONFLICT | 409 | 현재 매물 상태와 요청 액션 충돌 |
| RESERVATION_CONFLICT | 409 | 활성 예약 중복/동시 확정 충돌 |
| RATE_LIMITED | 429 | 전송/등록/신고 빈도 제한 |
| POLICY_BLOCKED | 423 또는 403 | 정책 위반으로 기능 차단 |
| IDEMPOTENCY_REPLAY | 200/409 | 동일 멱등 키 재시도 처리 |

### 37.3 상태 충돌 규칙
- `completed`, `cancelled` 매물에는 새 채팅 생성 불가
- `reserved`, `pending_trade` 상태의 매물에 다른 예약 확정 요청이 들어오면 409 반환
- 이미 `fulfilled` 또는 `cancelled`된 예약에 추가 확정/취소 요청 시 409 반환
- `report_locked` 채팅방에서는 메시지 발송 시 `POLICY_BLOCKED` 또는 읽기전용 응답 반환

### 37.4 권한 규칙
- 매물 수정/상태 변경은 작성자 또는 운영 권한만 가능
- 채팅 메시지는 채팅방 참여자만 가능
- 예약 확정/취소는 채팅방 참여자만 가능하되, 상대방이 제안한 예약을 제3자가 조작할 수 없어야 한다.
- 관리자 API는 일반 인증 외에 역할 검증과 감사 로그 기록이 필수다.

### 37.5 멱등성/재시도 고려
- 거래 완료, 예약 확정, 신고 접수, 제재 실행처럼 중복 실행 위험이 큰 API는 `Idempotency-Key` 또는 서버측 중복 방지 키를 지원해야 한다.
- 모바일 네트워크 불안정 환경에서 동일 요청이 재전송될 수 있으므로, 서버는 안전한 중복 처리 결과를 반환해야 한다.
- 푸시 토큰 등록, 알림 읽음 처리 등 반복 호출이 예상되는 API는 upsert 성격으로 설계 가능하다.

### 37.6 동시성 제어 포인트
- 예약 확정은 `listingId` 또는 `active_reservation` 기준 잠금/원자 갱신이 필요하다.
- 거래 완료는 연결된 채팅방과 매물 상태를 함께 검증해야 하며, 이미 다른 채팅방으로 완료 처리된 경우 거절해야 한다.
- 끌어올리기/노출 순위 변경 기능이 도입되면 쿨다운과 중복 요청 방지 규칙이 필요하다.

## 38. 서버/아이템 카탈로그 표준화 전략(초안)
### 38.1 목표
- 검색 품질과 입력 편의성을 동시에 확보하기 위해 표준 카탈로그와 자유 입력을 혼합한다.
- 초기에는 과도한 정규화보다 빠른 등록/검색을 우선하고, 반복 입력 데이터를 기반으로 표준화를 강화한다.

### 38.2 서버 카탈로그 원칙
- 서버는 가능한 한 표준 선택형 목록을 기본으로 제공한다.
- 운영자는 서버 추가/비활성화/표기명 변경을 관리할 수 있어야 한다.
- 종료되었거나 병합된 서버는 신규 등록에는 비활성화하되, 과거 매물 조회를 위해 레거시 값 보존이 필요하다.

### 38.3 아이템 카탈로그 원칙
- MVP는 `카테고리 선택 + 대표 아이템명 자유입력` 혼합 구조를 기본안으로 한다.
- 인기 아이템은 자동완성 후보를 제공하되, 후보에 없더라도 자유 입력 가능해야 한다.
- 자유 입력값은 정규화 후보 테이블에 적재하여 추후 alias/synonym 사전으로 승격할 수 있어야 한다.

### 38.4 정규화 계층(초안)
1. 표시 입력값(raw item name)
2. 정규화 후보(normalized candidate)
3. 표준 아이템(master item)
4. 동의어/오타 사전(alias dictionary)

### 38.5 검색 처리 원칙
- 검색 시 표준 아이템 일치, 동의어 일치, raw 문자열 부분 일치를 함께 지원한다.
- 오탈자 허용 검색은 확장 범위로 두되, 초반에는 동의어/별칭 사전 기반 보정을 우선한다.
- 제목, 아이템명, 옵션 텍스트는 색인 전략을 분리해 relevance를 조정한다.

### 38.6 데이터 설계 시사점
- Listing에는 `itemNameRaw`, `normalizedItemId(optional)`, `categoryId`, `serverId`를 함께 보관하는 구조가 유리하다.
- Catalog 테이블은 사용자 입력값을 강제로 덮어쓰지 않고, 검색/필터/추천에 활용하는 보조 레이어로 동작해야 한다.
- 관리자 화면에서 빈도 높은 자유 입력어를 검토해 표준 항목으로 승격할 수 있어야 한다.

## 39. 차단/뮤트 정책과 사용자 안전 UX(초안)
### 39.1 목표
- 악성 사용자와의 반복 접촉을 줄이고, 신고 이전 단계에서도 사용자가 스스로 대화 노출을 제어할 수 있게 한다.

### 39.2 차단(Block) 정책
- 사용자는 특정 사용자를 차단할 수 있어야 한다.
- 차단 시 효과:
  - 신규 채팅 생성 불가
  - 기존 채팅방은 읽기 전용 또는 숨김 후보
  - 차단 관계 상대의 푸시/알림 수신 중단
  - 상대 매물의 노출을 줄이거나 숨김 처리(정책 선택)
- 상호 차단 여부와 차단 시각을 저장해야 한다.

### 39.3 뮤트(Mute) 정책
- 뮤트는 안전 목적보다 알림 피로도 완화 목적이다.
- 뮤트 시 채팅/알림은 유지하되, 푸시만 끄거나 목록 우선순위를 낮출 수 있다.
- 뮤트는 신고/차단보다 가벼운 조치로, 기록 공개/운영 판단 지표로 직접 사용하지 않는다.

### 39.4 UX 원칙
- 차단 액션은 채팅방, 프로필, 매물 상세에서 접근 가능해야 한다.
- 차단 직전에는 효과를 명확히 안내한다(예: “이 사용자는 더 이상 새 채팅을 보낼 수 없습니다”).
- 차단 후 신고를 이어서 할 수 있는 CTA를 제공해 악성 행위 기록을 놓치지 않게 한다.
- 차단 해제는 마이페이지 안전 설정 또는 사용자 상세에서 가능해야 한다.

### 39.5 운영 연계
- 다수 사용자에게 반복 차단된 계정은 운영 모니터링 신호로 사용할 수 있다. 단, 자동 제재의 단독 근거로는 쓰지 않는다.
- 차단 직전/직후 메시지, 신고 여부, 노쇼 패턴을 함께 봐야 한다.

### 39.6 안전 메시지/가드레일
- 예약 직전 화면에 “계좌이체/외부 메신저 유도 주의”, “가능하면 플랫폼 내 기록 유지” 등의 안전 문구를 노출한다.
- 처음 거래하는 사용자 조합에는 후기/완료 수/가입 시점 등 판단 신호를 더 강조한다.
- 신고가 잦은 사용자와 대화 중일 때는 운영 검토 중 안내 배너 또는 주의 배지를 노출할 수 있다(정책 확정 필요).

## 40. 후기 공개/수정/비노출 정책(초안)
### 40.1 목표
- 후기가 거래 신뢰도를 높이되, 보복성 후기와 선제적 감정 싸움을 줄여야 한다.
- 사용자가 후기를 작성하도록 유도하면서도, 공개 규칙은 단순하고 예측 가능해야 한다.

### 40.2 기본 공개 정책
- 기본안: **상호 작성 완료 시 즉시 공개, 미작성 시 마감 시점 도래 후 작성된 후기만 공개**
- 후기 작성 가능 기간은 거래 완료 확정 후 7일(가정)
- 양측 모두 후기 작성 시, 상대 후기 내용을 작성 전에는 볼 수 없어야 한다.
- 한쪽만 작성하고 상대방이 기간 내 미작성한 경우, 마감 시점 이후 단독 후기 공개 허용
- 거래 완료가 `disputed` 상태로 전환되면 후기 공개를 보류하고 운영 검토 결과에 따름

### 40.3 수정/삭제 정책
- 작성 후 15분 이내 1회 수정 허용(가정)
- 추천/비추천 방향 전환까지 허용할지 여부는 정책 결정 필요하나, MVP에서는 전체 수정 1회 허용이 단순하다.
- 사용자 자의 삭제는 불가 또는 작성 직후 짧은 유예 시간 내만 허용하는 안을 검토한다.
- 운영 숨김은 가능하되, 공개 화면에서는 단순 삭제가 아니라 “운영정책에 따라 비노출 처리됨” 상태를 선택적으로 표시할 수 있다.

### 40.4 비노출/운영 개입 기준
- 욕설, 개인정보, 외부 홍보, 거래 무관 비방, 법적 위험 표현은 비노출 또는 수정 요청 대상
- 신고 접수된 후기는 `pending_moderation` 상태로 전환 가능
- 후기 숨김 여부와 별개로 내부 운영 기록과 산정용 원본 보관 정책은 분리해야 한다.

### 40.5 데이터/화면 설계 시사점
- Review 엔티티에 `publishedAt`, `editedAt`, `editCount`, `moderationReasonCode` 필드 추가 후보
- 프로필 화면에는 최근 후기 3~5개, 전체 후기 탭, 숨김 후기 제외 집계를 제공
- 후기 CTA는 완료 직후/24시간 후/마감 전 리마인드 1회 등 과도하지 않게 설계

## 41. SEO / 공개 URL / 인덱싱 정책(초안)
### 41.1 목표
- 검색 유입을 확보하되, 실제 거래가 종료된 매물과 민감 정보가 외부 검색에 오래 남지 않게 한다.

### 41.2 인덱싱 기본안
- `available` 상태 공개 매물만 검색엔진 인덱싱 허용 후보
- `reserved`, `pending_trade`, `completed`, `cancelled` 매물은 기본적으로 `noindex` 또는 색인 제외
- 프로필 페이지는 기본 `noindex`를 우선 검토하고, 공개 신뢰 배지/후기 중심 공개 프로필은 Post-MVP 검토 범위로 둔다.

### 41.3 공개 페이지 노출 원칙
- 공개 매물 페이지에는 제목, 아이템명, 서버, 가격, 거래 방식, 설명 요약, 대표 이미지 정도만 노출
- 닉네임은 마스킹 또는 제한 공개 우선
- 캐릭터명, 채팅 이력, 예약 정보, 최근 활동의 정확한 타임스탬프는 비공개
- 구조화 데이터(schema.org)는 상품/게시물 유사 형태로 검토 가능하나, 거래 당사자 식별 정보는 제외

### 41.4 URL 수명주기 정책
- 매물이 `completed` 또는 `cancelled` 되면 공개 URL은 유지하되 내용은 축소/비공개 전환하는 소프트 랜딩 방식을 우선 검토
- 소프트 랜딩 페이지에는 “거래 종료됨”, “유사 매물 보기” CTA를 제공할 수 있다.
- 완전 삭제 요청 또는 법적 이슈가 있는 경우 404/410 처리 가능해야 한다.

### 41.5 검색 유입 보호 장치
- 공개 페이지에서 채팅/찜 CTA는 로그인 전환으로 연결
- 크롤러 대상 페이지에는 연락처/민감 문자열 노출 차단 및 이미지 OCR 리스크 점검 필요
- 완료 매물의 캐시 잔존을 줄이기 위한 sitemap/robots/meta 정책 정리가 필요하다.

## 42. 푸시 opt-in / 채널 우선순위 / 소음 제어 정책(초안)
### 42.1 알림 채널 정의
- 인앱 알림함: 기본 필수 채널
- 모바일 푸시: 거래 진행 이벤트 중심 핵심 채널
- 이메일/SMS: MVP 제외 또는 비상/복구 수단으로만 검토

### 42.2 opt-in 원칙
- 앱 설치 직후 즉시 시스템 푸시 권한을 강요하지 않고, 첫 거래 행동 맥락에서 프리프롬프트를 제공한다.
- 추천 노출 시점:
  1. 첫 채팅 시작 직후
  2. 첫 예약 제안 직전 또는 직후
  3. 미응답으로 거래 기회를 놓칠 수 있는 시점
- 푸시 거부 사용자는 인앱 알림함과 배지로 동일 핵심 이벤트를 확인 가능해야 한다.

### 42.3 이벤트별 채널 우선순위
| 이벤트 | 인앱 | 푸시 | 비고 |
|---|---|---|---|
| 새 채팅 시작 | 필수 | 기본 ON 권장 | 거래 연결 시작점 |
| 새 메시지 | 필수 | 사용자 설정 따름 | 현재 대화방 열람 중이면 푸시 생략 |
| 예약 제안/확정/변경/취소 | 필수 | 강한 권장 | 시간 민감도 높음 |
| 거래 완료 요청 | 필수 | 강한 권장 | 후기/분쟁 흐름 시작 |
| 후기 작성 요청 | 필수 | 선택적 | 과도한 재알림 금지 |
| 운영 경고/제재 | 필수 | 중요도 따라 발송 | 계정 상태 변화 |
| 마케팅/추천 | 선택 | 기본 OFF | MVP 범위 제외 우선 |

### 42.4 사용자 설정 모델
- 최소 설정 단위:
  - 거래 필수 알림
  - 새 메시지 알림
  - 후기/리마인드 알림
  - 마케팅/이벤트 알림
- 기본값은 거래 필수 알림 ON, 마케팅 알림 OFF를 우선안으로 한다.

### 42.5 소음 제어 규칙
- 동일 채팅방 메시지가 짧은 시간 연속 도착하면 묶음 푸시 처리
- 예약 관련 이벤트는 최신 상태 기준으로 갱신형 알림을 우선 검토
- 심야 시간에는 거래 필수 알림만 허용하고, 후기 요청/홍보성 알림은 지연 발송
- 완료된 거래에 대한 재알림은 제한 횟수(예: 후기 리마인드 최대 2회, 가정)를 둔다.

## 43. 알림 이벤트 발송 매트릭스(초안)
### 43.1 목적
- 거래 성사에 직접 영향을 주는 이벤트만 적시에 보내고, 같은 사건에 대해 채널별 중복 체감을 줄인다.
- 모바일 푸시, 인앱 알림함, 배지, 운영 통지가 어떤 기준으로 발송되는지 명시해 화면/백엔드/운영 해석 차이를 줄인다.

### 43.2 이벤트 분류
- **T1 즉시행동 필요**: 예약 응답, 완료 확인, 분쟁 소명 요청
- **T2 대화/연결 유지**: 새 문의, 새 메시지, 예약 변경
- **T3 사후 정리/리텐션**: 후기 작성 요청, 상태 점검 리마인드
- **T4 운영/정책**: 경고, 제한, 신고 처리 결과

### 43.3 발송 매트릭스
| 이벤트 | 주 수신자 | 기본 채널 | 푸시 기본값 | 억제 조건 | 딥링크 대상 |
|---|---|---|---|---|---|
| 새 채팅 시작 | 매물 작성자 | 인앱 + 푸시 | ON | 작성자가 해당 채팅방을 현재 열람 중 | 해당 채팅방 |
| 새 메시지 | 상대 채팅 참여자 | 인앱 + 푸시 | ON(메시지 알림 설정 시) | 수신자가 같은 채팅방 활성 열람 중, 차단/잠금 상태 | 해당 채팅방 |
| 예약 제안 도착 | 상대 예약 참여자 | 인앱 + 푸시 | ON | 예약이 이미 만료/취소됨 | 예약 카드가 열린 채팅방 |
| 예약 확정 | 양측 참여자 | 인앱 + 푸시 | ON | 직전 동일 상태 재전송, 사용자가 이미 확정 화면 열람 중 | 채팅방 또는 내 거래 상세 |
| 예약 취소/만료 | 양측 참여자 | 인앱 + 푸시 | ON | 사용자가 동일 취소 사유를 이미 확인함 | 채팅방 |
| 예약 임박(24h/1h) | 양측 참여자 | 인앱 + 푸시 | ON | 예약이 취소/완료됨, 최근 10분 내 동일 단계 알림 발송 | 내 거래 상세 |
| 거래 완료 요청 | 상대방 | 인앱 + 푸시 | ON | 이미 확인/이의제기 완료 | 완료 확인 화면 |
| 거래 완료 자동확정 | 양측 참여자 | 인앱 + 푸시 | ON | 이미 앱 내에서 동일 이벤트 확인됨 | 완료 기록/후기 진입 |
| 분쟁 접수/추가 소명 요청 | 분쟁 당사자 | 인앱 + 푸시 | ON | 사용자가 사건 화면을 이미 열람 중 | 분쟁 상세 |
| 후기 작성 요청 | 후기 미작성 당사자 | 인앱 + 선택적 푸시 | 설정 따름 | 이미 후기 작성, 분쟁 상태, 최대 리마인드 횟수 도달 | 후기 작성 화면 |
| reserved 장기 유지 점검 | 매물 작성자 | 인앱 우선, 푸시 선택 | 설정 따름 | 최근 24시간 내 동일 점검 알림 발송 | 내 매물 상세 |
| 신고 처리 결과 | 신고자 | 인앱 우선, 중요도 따라 푸시 | 중요도 따라 ON | 동일 reportStatus 재발송 | 내 신고 상세 |
| 경고/기능 제한/정지 | 제재 대상자 | 인앱 + 푸시 | ON | 로그인 중 실시간 배너로 이미 확인 | 계정 상태/정책 안내 |

### 43.4 알림 본문/카피 원칙
- 푸시 본문에는 아이템명, 상태 변화, 다음 액션 정도만 포함하고 민감 정보(전화번호, 정확한 접선 위치, 상대 실명 추정 정보)는 포함하지 않는다.
- 잠금화면 노출을 고려해 메시지 원문 전체를 그대로 푸시에 복사하지 않는다.
- 예약/완료/분쟁 알림은 반드시 다음 행동 CTA가 분명해야 한다.
- 운영/제재 알림은 감정적 표현보다 `무엇이 제한되었는지`, `어디서 확인하는지`, `이의제기 가능 여부`를 우선 전달한다.

### 43.5 배지/인앱/푸시 우선순위 규칙
- T1 이벤트는 푸시 비허용 상태여도 인앱 알림함, 내 거래 배지, 해당 객체 상태 배너가 모두 남아야 한다.
- T2 이벤트는 푸시 실패 시 인앱 알림함 저장만으로도 기능적으로 복구 가능해야 한다.
- T3 이벤트는 푸시보다 인앱 우선이며, 사용자가 반복 무시하면 재알림 빈도를 낮춘다.
- T4 이벤트는 사용자의 권리/이용 상태에 영향을 주므로 인앱 기록 보존을 필수로 한다.

### 43.6 사용자 설정 연결 규칙
| 설정 항목 | 포함 이벤트 | 기본값 | 비고 |
|---|---|---|---|
| 거래 필수 알림 | 예약 제안/확정/임박, 완료 요청, 분쟁 요청, 제재 알림 | ON | 사용자가 완전 OFF 하더라도 인앱 기록은 유지 |
| 새 메시지 알림 | 새 채팅 시작, 새 메시지 | ON | 채팅방 단위 뮤트가 전역 설정보다 우선 |
| 후기/리마인드 알림 | 후기 요청, 장기 reserved 점검, 휴면 매물 점검 | ON 또는 약한 ON 후보 | MVP에서는 단순 토글 우선 |
| 마케팅/추천 알림 | 추천 매물, 인기 검색어, 이벤트 | OFF | MVP 범위 제외 우선 |

### 43.7 재알림/묶음 처리 원칙
- 같은 채팅방에서 짧은 시간에 여러 메시지가 오면 메시지 수를 묶어 1개의 푸시로 요약한다.
- 예약 상태는 최신 상태만 남기는 갱신형 알림을 우선 검토한다.
- 후기 요청, 상태 점검, 휴면 리마인드는 누적형이 아니라 가장 최신 건 1개만 유지하는 구조를 우선한다.
- 동일 객체에 대해 사용자가 이미 화면 진입으로 확인한 이벤트는 read 동기화 후 추가 푸시를 억제한다.

### 43.8 데이터/API 시사점
- Notification 엔티티/이벤트 로그에 아래 필드 후보를 추가 검토한다.
  - `eventType`
  - `eventPriorityTier`
  - `deliveryChannels`
  - `pushSuppressedReasonCode`
  - `groupingKey`
  - `deepLinkPayloadJson`
- 푸시 발송 결과와 사용자 행동을 연결해 `sent -> delivered -> opened -> acted` 퍼널을 추적해야 한다.
- 알림 읽음은 단순 목록 읽음 외에 `actionCompletedAt`과 분리해 저장하는 것이 운영 분석에 유리하다.

## 45. 검색 자동완성 / 추천어 / 오타 보정 정책(초안)
### 43.1 목표
- 사용자가 정확한 아이템명/서버명을 몰라도 원하는 매물에 빠르게 도달하게 한다.
- 초기에는 과도한 검색 복잡도보다 체감 편의 개선에 집중한다.

### 43.2 자동완성 소스 우선순위
1. 표준 서버 카탈로그
2. 표준/승격된 아이템 카탈로그
3. 최근 인기 검색어
4. 최근 등록/거래가 활발한 raw item name
5. 사용자 개인 최근 검색어

### 43.3 추천어 노출 원칙
- 검색어 입력 전: 서버별 인기 검색어, 최근 많이 거래된 카테고리 노출
- 1~2자 입력: prefix 기반 자동완성 + 카테고리 제안
- 3자 이상 입력: 아이템명 후보 + 동의어/별칭 매핑 결과 + raw 부분 일치
- 추천어는 단순 트래픽이 아니라 실제 상세 진입/채팅 시작/예약 전환 기여도를 반영해 정렬한다.

### 43.4 오타/별칭 보정 정책
- 서버명/아이템명에 대해 alias dictionary 기반 정규화 제안 제공
- 자동 치환보다는 “이 검색어를 찾으셨나요?” 보정 UX를 우선 적용
- 잘못된 보정으로 탐색 실패가 생기지 않도록, 원문 검색 결과도 함께 유지해야 한다.
- 운영자는 자주 발생하는 오타/별칭을 사전에 추가할 수 있어야 한다.

### 43.5 금칙어/운영 룰
- 욕설, 사기 유도 문구, 외부 연락처 유도 키워드는 추천어에서 제외
- 특정 사용자 닉네임/개인정보성 문자열은 자동완성 후보로 승격되지 않도록 차단
- 신고/차단 이슈가 많은 키워드는 운영 검토 후 노출 제한 가능

### 43.6 분석 포인트
- `search_autocomplete_impression`
- `search_autocomplete_click`
- `search_zero_result`
- `search_correction_shown`
- `search_correction_accepted`
- 추천어별 상세 진입률, 채팅 시작률, 예약 전환율 추적 필요

## 45. 운영 대시보드 / 이상징후 탐지 룰(초안)
### 44.1 운영 대시보드 목표
- 단순 신고 처리 현황을 넘어, 거래 성사율 저하와 악성 행위 패턴을 조기 감지해야 한다.

### 44.2 핵심 대시보드 묶음
- **거래 퍼널 대시보드**: 조회→채팅→예약→완료 전환율, 서버별/카테고리별 비교
- **신뢰/안전 대시보드**: 신고 건수, 노쇼율, 차단 비율, 제재 집행 현황
- **운영 처리 대시보드**: SLA 준수율, 큐 적체, 우선순위별 미처리 건수
- **검색 품질 대시보드**: zero-result 비율, 자동완성 클릭률, 보정 수락률
- **알림 성과 대시보드**: 푸시 수신/오픈/행동 전환율, opt-in 전환율

### 44.3 이상징후 탐지 룰 후보
- 동일 사용자가 짧은 시간에 유사 매물 다건 등록
- 동일 사용자가 여러 상대에게 예약을 반복 확정 후 취소
- 특정 계정이 다수 사용자에게 단기간 차단/신고됨
- 새 계정이 고가 매물 다건 등록 후 외부 연락처 교환 시도
- 특정 서버/카테고리에서 zero-result 급증 또는 채팅 전환율 급락
- 푸시 발송 성공은 높지만 예약/완료 전환이 급락하는 경우(알림 피로/품질 이슈 신호)

### 44.4 룰 처리 단계
1. 탐지
2. 점수화 또는 심각도 분류
3. 운영 큐 적재 또는 자동 임시 제한
4. 운영자 검토
5. 해제/확정 제재/룰 튜닝

### 44.5 자동 조치 원칙
- 자동 영구 제재는 금지하고, 자동 조치는 임시 숨김/추가 인증 요구/등록 속도 제한 수준으로 제한한다.
- 사용자 영향이 큰 자동 조치는 이력과 사유를 사용자에게 설명 가능해야 한다.
- 탐지 룰은 오탐률을 측정하고, 운영자가 해제한 사례를 기반으로 조정해야 한다.

## 46. 데이터 정합성 제약 / 인덱스 / 이력 저장 원칙(초안)
### 45.1 정합성 제약
- 한 `listingId`에는 동시에 1개의 활성 예약(`proposed` 또는 `confirmed` 중 정책상 활성로 간주하는 상태)만 허용하는 것을 기본안으로 한다.
- `Listing.reservedChatRoomId`가 존재하면, 해당 채팅방은 반드시 동일 `listingId`를 참조해야 한다.
- `Listing.status=reserved`이면 `reservedChatRoomId`가 필수다.
- `Listing.status=pending_trade`이면 `reservedChatRoomId`와 `confirmed` 상태의 Reservation이 모두 존재해야 한다.
- `TradeCompletion`은 매물 종결 기준 채팅방 1건만 연결할 수 있도록, `listingId + completionStatus in (requested, confirmed, disputed)` 수준의 중복 방지 제약이 필요하다.
- `Review`는 `completionId + reviewerUserId` 기준 유니크해야 한다.
- `ChatRoom`은 `listingId + sellerUserId + buyerUserId` 기준 유니크해야 한다.
- 차단 관계는 `blockerUserId + blockedUserId` 기준 유니크해야 한다.
- 사용자 탈퇴 후에도 거래/신고/감사 로그 참조 무결성이 깨지지 않도록 soft delete 또는 비식별화 전략을 우선 검토한다.

### 45.2 인덱스 후보
- Listing: `(status, visibility, serverId, categoryId, lastActivityAt desc)`
- Listing 검색: `itemNameRaw`, `title`, `description` 전문검색 또는 prefix 인덱스 전략
- ChatRoom: `(sellerUserId, lastMessageAt desc)`, `(buyerUserId, lastMessageAt desc)`
- Reservation: `(listingId, reservationStatus)`, `(chatRoomId, scheduledAt)`
- Report: `(reportStatus, priority, createdAt)`, `(targetType, targetId)`
- Notification: `(userId, readAt, createdAt desc)`
- ModerationAction/AuditLog: `(targetUserId, createdAt desc)`, `(createdByAdminId, createdAt desc)`
- 검색 자동완성용 카탈로그: `(normalizedTerm)`, `(aliasTerm)` 인덱스 필요

### 45.3 상태 이력 저장 원칙
- `Listing`, `Reservation`, `ChatRoom`, `TradeCompletion`, `Report`는 현재 상태 컬럼과 별도로 상태 변경 이력 테이블 또는 이벤트 로그를 가져야 한다.
- 상태 이력 레코드 최소 필드: `entityType`, `entityId`, `fromStatus`, `toStatus`, `changedByUserId or system`, `reasonCode`, `memo`, `createdAt`
- 사용자 액션과 시스템 자동 액션을 구분해야 하며, 자동 만료/자동 복귀/운영 강제 변경도 이력에 남겨야 한다.
- 운영자 화면에서는 현재 상태뿐 아니라 “왜 이 상태가 되었는지”를 시간순으로 재구성 가능해야 한다.

### 45.4 카운터 캐시/집계 원칙
- `viewCount`, `favoriteCount`, `chatCount`, `completedTradeCount` 등은 캐시 값으로 취급하고, 원본 이벤트/테이블로부터 재계산 가능해야 한다.
- 사용자 신뢰 지표는 실시간 집계 대신 배치 또는 비동기 갱신을 허용한다. 단, 후기 작성/제재 직후 프로필 핵심 지표는 합리적인 시간 내 반영되어야 한다.
- 랭킹/검색 노출에 쓰이는 캐시는 지연 갱신 가능하나, 거래 상태와 충돌하는 stale 노출은 최소화해야 한다.

## 47. 화면별 주요 액션 가드레일 / CTA 규칙(초안)
### 46.1 목록 화면 카드 액션 규칙
- `available`: `채팅하기`, `찜`, `상세보기` 노출
- `reserved`: 기본적으로 `예약중` 배지 우선 노출, `채팅하기`는 정책에 따라 유지 또는 비활성
- `pending_trade`: 기본 목록에서는 숨김이 원칙이며, 노출 시 `거래대기`와 직접 문의 불가 표시 필요
- `completed`, `cancelled`: 공개 목록 카드 노출 제외
- 비회원은 `채팅하기`, `찜` 클릭 시 로그인 유도 바텀시트 노출

### 46.2 상세 화면 CTA 규칙
| 상태 | 주 CTA | 보조 CTA | 금지/제한 |
|---|---|---|---|
| available | 채팅 시작 | 찜, 신고 | 없음 |
| reserved | 진행 중 상대가 아닌 경우 대기/알림 유도 또는 채팅 제한 | 찜, 신고 | 즉시 예약 확정 불가 |
| pending_trade | 직접 거래 문의 CTA 비노출 | 신고, 유사 매물 보기 | 신규 채팅 생성 제한 |
| completed | 유사 매물 보기 | 신고, 기록 보기 | 채팅 시작 불가 |
| cancelled | 유사 매물 보기 | 신고 | 채팅 시작 불가 |

### 46.3 채팅 화면 액션 규칙
- 참여자가 아닌 사용자는 접근 불가
- `open`: 메시지 발송, 예약 제안, 차단, 신고 허용
- `reservation_proposed`: 수락/거절/대안 제시 CTA 강조
- `reservation_confirmed`: 시간 변경 요청, 취소 요청, 거래 준비 체크리스트 노출
- `trade_due`: `거래 완료`, `노쇼/불발`, `상대에게 도착 메시지 보내기` CTA 우선 노출
- `deal_completed`: 메시지 입력창 제거 또는 읽기 전용 전환, 후기 CTA 노출
- `report_locked`: 일반 입력창 제거, 운영 검토 중 안내와 신고번호 노출

### 46.4 내 매물 화면 액션 규칙
- 작성자는 진행 중 매물 카드에서 `수정`, `상태 변경`, `채팅 보기`를 바로 실행 가능해야 한다.
- `reserved`, `pending_trade` 상태 매물은 제목/가격/핵심 거래 조건 수정 시 강한 경고 또는 제한이 필요하다.
- 완료/취소 매물은 `복제 후 재등록`을 기본 CTA로 두고, 직접 재오픈은 정책적으로 제한한다.

### 46.5 내 거래 화면 우선순위 규칙
- 사용자의 액션이 필요한 거래(예약 응답 대기, 완료 확인 대기, 노쇼 판단 필요)를 최상단 정렬
- 예약 시각 임박 거래는 `scheduledAt` 기준 정렬 우선순위를 높인다.
- 단순 미읽음보다 `거래 기한 임박 + 액션 필요` 상태를 우선 배지로 노출한다.

### 46.6 후기 작성 화면 규칙
- 추천/비추천 선택 후 선택 코멘트 입력 구조를 기본안으로 한다.
- 상대 후기 내용은 공개 전까지 블라인드 처리한다.
- 거래와 무관한 표현, 개인정보 포함 시 전송 전 경고를 줄 수 있어야 한다.
- 작성 마감 시점이 명확히 보여야 하며, 마감 이후 CTA는 비활성화한다.

## 48. 거래 가능 품목 및 콘텐츠 정책(초안)
### 48.1 목표
- 플랫폼이 어떤 거래를 허용하고 어떤 거래를 제한하는지 명확히 해, 매물 등록 UX·검색 노출·신고 처리·운영 제재가 동일 기준을 사용하도록 한다.
- 단순 커뮤니티 게시판이 아니라 `실제 거래 완료를 돕는 중개 서비스`라는 제품 방향에 맞게, 거래 가능성이 낮거나 분쟁/불법 리스크가 큰 콘텐츠를 조기에 제어한다.
- 사전 차단, 사후 숨김, 운영 검토의 경계를 정의해 과차단과 무대응을 동시에 줄인다.

### 48.2 정책 분류 원칙
거래 관련 콘텐츠는 아래 4단계로 분류한다.

| 분류 | 의미 | 등록 시 처리 | 노출 정책 | 운영 처리 |
|---|---|---|---|---|
| `allowed` | 일반 허용 품목/표현 | 즉시 등록 가능 | 정상 노출 | 일반 모니터링 |
| `restricted` | 조건부 허용, 추가 검토/경고 필요 | 경고 또는 일부 필드 제한 | 제한 노출 가능 | 반복 시 검토 |
| `review_required` | 자동 확신 불가, 운영 판단 필요 | 임시 보류 또는 제한 공개 | 검토 전 노출 제한 후보 | 큐 적재 |
| `prohibited` | 명백한 금지 품목/행위 | 등록 차단 또는 즉시 숨김 | 비노출 | 제재 가능 |

원칙:
- 사용자에게는 가능한 한 단순한 문구(`등록 불가`, `검토 필요`, `수정 후 등록`)로 안내한다.
- 내부적으로는 `contentPolicyCode`, `riskTier`, `moderationDecision` 같은 구조화 값을 남겨 API/운영/분석이 같은 기준을 쓰게 한다.

### 48.3 기본 허용 거래 범위
MVP 기본안에서 허용 대상으로 보는 거래는 아래와 같다.
- 리니지 클래식 관련 아이템
- 리니지 클래식 관련 재화/소모품/장비/재료
- 게임 내 거래를 전제로 한 판매/구매 의사 표현
- 실제 거래 실행에 필요한 시간/장소/서버/캐릭터명 협의

허용 원칙:
- 서비스는 `리니지 클래식 개인 간 거래 연결`에 집중하며, 게임과 무관한 일반 중고거래/커뮤니티 글은 범위 밖으로 본다.
- 품목 자체가 허용 범위여도 허위 정보, 사기 유도, 외부 결제 강요, 계정 거래 등 위험 행위와 결합되면 제한 또는 금지 대상으로 전환될 수 있다.

### 48.4 금지 품목/행위 카테고리
아래는 MVP 기준 금지 또는 강한 제한 대상 후보다.

| 카테고리 | 기본 정책 | 비고 |
|---|---|---|
| 계정 판매/구매/대여 | `prohibited` | 계정 양도 리스크 높음 |
| 개인정보/연락처/계좌 직접 매매성 게시 | `prohibited` | 거래 목적 무관 민감정보 노출 금지 |
| 사기 유도 문구(선입금 강요, 외부 링크 결제 유도 등) | `prohibited` | 자동 탐지 + 운영 검토 |
| 현금성/불법성 직접 강조 표현 | `review_required` 또는 `prohibited` | 실제 정책 문구는 별도 확정 필요 |
| 게임과 무관한 일반 물품 거래 | `prohibited` 또는 `off-topic` 제한 | 플랫폼 범위 밖 |
| 도배/광고/홍보성 매물 | `prohibited` | 반복 시 제재 |
| 욕설/혐오/성적 표현 포함 매물/후기/채팅 | `prohibited` | 안전 정책 연계 |
| 불법 촬영물/저작권 침해/위법 콘텐츠 이미지 | `prohibited` | 즉시 숨김/차단 후보 |

주의:
- 실제 법률/게임 운영정책 해석이 필요한 세부 품목은 별도 정책 문서에서 확정해야 하며, 본 PRD에서는 제품/운영 관점의 기본 분류 기준만 정의한다.
- 외부 법적 사실을 단정하지 않고, 정책상 위험도가 높아 플랫폼이 허용하지 않는다는 관점으로 기술한다.

### 48.5 조건부 허용(`restricted`) 사례
다음은 완전 금지보다 `조건부 허용 + 경고/검토`가 적절할 수 있는 사례다.
- 가격/수량/품목 설명이 지나치게 불명확한 매물
- 대표 이미지가 없거나 실제 거래 품목 확인이 어려운 매물
- 과도한 시세 언급, 과장 표현, 이모지/반복 문자로 가독성이 낮은 제목
- 오프라인 직거래를 암시하지만 장소 설명이 지나치게 모호하거나 안전성이 낮은 경우
- 외부 연락처를 직접 적지 않았지만 외부 이동을 강하게 유도하는 문구
- 운영상 민감 품목으로 분류된 고가/고위험 매물

처리 원칙:
- 가능한 한 즉시 차단보다 수정 유도/경고/노출 감점부터 적용한다.
- 반복적이거나 다수 신고가 결합되면 `review_required` 또는 `prohibited`로 상향 가능하다.

### 48.6 등록 시점 검수 단계
매물 생성/수정 시 아래 순서의 검수를 기본안으로 둔다.

1. **구조 검증**
   - 필수값, 길이, enum, 수량/가격 범위 확인
2. **정책 키워드 검수**
   - 금지 품목, 욕설, 외부 연락처 유도, 사기 고위험 문구 패턴 탐지
3. **민감정보 검수**
   - 전화번호, 계좌번호, 오픈채팅 링크, 메신저 ID, 실명 추정 정보 탐지
4. **이미지/OCR 검수**
   - 이미지 내 연락처/계좌/무관 콘텐츠/부적절 표현 탐지 후보
5. **위험도 산정**
   - 신규 계정, 반복 위반, 유사 도배, 고가 매물 여부 등 결합
6. **결정**
   - 등록 허용 / 수정 후 재시도 / 검토 대기 / 즉시 차단

### 48.7 채팅/후기/이미지까지 포함한 콘텐츠 정책 범위
콘텐츠 정책은 매물 본문에만 적용되지 않는다.

| 객체 | 검수 대상 | 대표 리스크 |
|---|---|---|
| Listing | 제목, 설명, 이미지, 구조화 속성 | 허위매물, 금지 품목, 광고 |
| ChatMessage | 본문, 첨부, 링크, OCR | 외부 유도, 욕설, 사기 정황 |
| Review | 코멘트 | 보복성 비방, 개인정보, 무관 표현 |
| Profile | 닉네임, 소개, 아바타 | 광고, 욕설, 연락처 노출 |
| Reservation note | 메모, 장소 설명 | 과도한 개인정보, 위험한 만남 유도 |

원칙:
- 동일 정책 엔진/사유 코드 체계를 재사용하되, 객체별 허용 수준은 다를 수 있다.
- 예를 들어 캐릭터명은 예약에서는 허용 가능하지만 공개 프로필/공개 매물 본문에서는 제한될 수 있다.

### 48.8 사전 차단 vs 사후 숨김 기준
| 처리 방식 | 사용 시점 | 예시 |
|---|---|---|
| 사전 차단(pre-block) | 명백한 금지 패턴 | 전화번호/오픈채팅 링크 본문 직입력, 계정 거래 문구 |
| 수정 유도(soft block) | 오탐 가능성 높음, 개선 가능한 입력 | 제목 도배, 설명 부족, 모호한 가격 표현 |
| 제한 공개(limited visibility) | 위험 신호 있으나 확정 어려움 | 신규 계정 고가 매물, 민감 키워드 포함 |
| 사후 숨김(post-hide) | 신고/운영 판단 후 | 허위매물, 욕설, 사기 의심 |
| 운영 검토(review queue) | 자동 판단 애매함 | 이미지 OCR 탐지, 문맥 의존 표현 |

설계 원칙:
- 등록 UX에서 너무 많은 오탐으로 전환율이 깨지지 않도록, 명백한 금지 외에는 수정 유도/검토 대기를 우선 고려한다.
- 단, 사기/불법/개인정보 노출은 보수적으로 차단한다.

### 48.9 사용자 노출 문구 원칙
- 사용자에게 내부 위험 점수나 탐지 룰 이름을 직접 공개하지 않는다.
- 대신 아래 수준의 설명을 제공한다.
  - `연락처/외부 링크는 본문에 직접 적을 수 없어요`
  - `거래와 무관한 홍보성 내용은 등록할 수 없어요`
  - `일부 내용은 운영 검토 후 노출될 수 있어요`
  - `과도한 개인정보가 포함되어 수정이 필요해요`
- 제재/반복 위반 시에는 사유 코드와 재시도 가능 여부를 분리 안내한다.

### 48.10 운영/제재 연동 규칙
- 동일 `contentPolicyCode`가 짧은 기간 반복되면 경고 없이 더 강한 제한으로 상향 가능하다.
- 콘텐츠 정책 위반은 객체 단위 처리와 계정 단위 처리로 나뉜다.
  - 객체 단위: 숨김, 수정 요청, 등록 차단
  - 계정 단위: 등록 제한, 채팅 제한, 기간 정지
- 운영자는 `무엇이 탐지되었는지`, `자동 결정인지 수동 결정인지`, `반복 위반인지`를 함께 볼 수 있어야 한다.

후보 필드:
- `contentPolicyCode`
- `contentRiskTier`
- `moderationDecision`
- `decisionSource` (`auto` / `manual` / `appeal`)
- `policyStrikeCount`

### 48.11 백오피스/분석 시사점
운영 화면에서 아래를 확인 가능해야 한다.
- 객체 원문/마스킹 버전
- 탐지된 정책 코드 목록
- 자동 탐지 confidence 또는 rule hit 목록
- 과거 동일 코드 위반 이력
- 사용자가 수정 후 재시도한 이력
- 숨김/복구/이의제기 결과

분석 이벤트 후보:
- `listing_policy_blocked`
- `listing_policy_warned`
- `listing_submitted_for_review`
- `chat_message_policy_blocked`
- `content_policy_appeal_submitted`

핵심 분석 포인트:
- 정책 차단이 등록 완료율에 미치는 영향
- 오탐률(운영 복구율)
- 금지 품목/연락처 유도 패턴의 신규/재범 비율
- 객체별(policy on listing/chat/review/profile) 위반 분포

### 48.12 오픈 질문
- 계정 거래, 현금 직거래 표현, 재화 표현 등 민감 카테고리를 어디까지 금지/검토 대상으로 둘 것인가?
- 이미지 OCR/정책 검수를 MVP에 어느 수준까지 포함할 것인가?
- 운영 검토 대기 매물을 사용자에게 `임시 비공개`로 둘지 `일부 제한 공개`로 둘지?
- 정책 위반 재범 기준을 계정 단위/기기 단위 어디까지 연결할 것인가?

## 49. 운영/법적 고지 초안(서비스 책임 범위)
### 47.1 서비스 책임 범위 원칙
- 린클은 개인 간 거래를 연결하고 거래 기록/신고/운영 도구를 제공하는 중개형 플랫폼이다.
- 린클은 거래 대금 보관, 결제 보증, 물품 진위 보증, 거래 이행 강제의 주체가 아니다.
- 다만 플랫폼 내 기록에 근거한 운영 제재, 위험 사용자 제한, 안전 가이드 제공 책임은 가진다.

### 47.2 사용자 고지 필요 문구 방향
- 거래 전: “외부 메신저/계좌 유도에 주의하세요. 가능하면 플랫폼 내 대화를 유지하세요.”
- 예약 확정 시: “약속 시간/장소/캐릭터명을 다시 확인하세요. 거래 완료 전 선입금/선전달은 주의가 필요합니다.”
- 완료 처리 시: “완료 처리 후에는 거래가 종결되며, 분쟁 시 플랫폼 기록을 기준으로 검토될 수 있습니다.”
- 신고 화면: “운영 검토는 거래 결과를 보증하지 않으며, 제재 여부는 제출된 기록과 증빙에 따라 결정됩니다.”

### 47.3 약관/정책 문서로 분리될 항목
- 금지 거래 품목/행위 정책
- 외부 연락처/개인정보 공유 정책
- 신고 및 제재 기준표
- 로그 보관 및 개인정보 처리방침
- 거래 분쟁 비보증/면책 조항
- 탈퇴/콘텐츠 비노출/보관 정책

### 47.4 고위험 시나리오 대응 원칙
- 고가 매물, 신규 계정, 외부 연락처 유도, 반복 노쇼/신고 패턴 조합은 고위험 거래 신호로 간주한다.
- 고위험 시나리오에서는 추가 경고 배너, 채팅 제약, 운영 검토 대기, 등록 속도 제한을 적용할 수 있어야 한다.
- 사용자에게는 “왜 제한되었는지”를 설명 가능한 수준의 사유 코드를 제공해야 한다.

## 49. 거래완료 확정 / 분쟁 / 재오픈 정책(초안)
### 48.1 목표
- 거래완료를 너무 쉽게 종결시켜 오완료를 만들지 않으면서도, 실제로 끝난 거래가 장기간 미종결 상태로 남지 않게 한다.
- 완료 직후 후기/신뢰도 반영과 분쟁 접수를 자연스럽게 연결한다.
- 거래 불발 후 재오픈과 부분거래를 별도 규칙으로 정의해 상태 혼선을 줄인다.

### 48.2 기본 완료 확정안
- 기본안은 **한쪽의 완료 요청 + 상대방 확인 기회 + 무응답 시 자동확정** 구조를 채택한다.
- 완료 요청은 해당 매물의 `reservedChatRoomId` 또는 실제 예약/진행 중인 채팅방 참여자만 할 수 있다.
- 완료 요청 시 즉시 `TradeCompletion.completionStatus=requested`를 생성하고, 매물 상태는 `completed_pending_confirmation` 성격의 내부 상태로 관리한다.
- 사용자 노출 단순화를 위해 UI에는 `거래완료 확인 대기` 배지/문구를 사용하고, 외부 공개 목록에서는 이미 종결 후보로 취급한다.
- 상대방이 확인하면 `confirmed`로 전환되고, 매물 최종 상태는 `completed`로 확정된다.
- 상대방이 일정 시간(예: 24시간, 가정) 내 응답하지 않으면 자동확정한다.

### 48.3 완료 요청 후 가능한 상대 액션
| 상대 액션 | 결과 | 비고 |
|---|---|---|
| 완료 확인 | `confirmed` | 후기 작성 가능 즉시 열림 |
| 미응답 | 시간 경과 후 `confirmed` 자동 전환 | 자동확정 시점 기록 필요 |
| 이의 제기 | `disputed` | 운영 검토 큐 적재 |
| 거래 불발 주장 | `disputed` 또는 `cancelled_after_request` 성격 이력 기록 | 실제 상태는 운영 판단 전 잠정 보류 |

### 48.4 분쟁 처리 규칙
- 완료 요청에 대해 이의 제기되면 공개 신뢰 지표 반영과 후기 공개를 일시 보류한다.
- 분쟁 상태에서는 해당 채팅방을 기본적으로 읽기 전용으로 두고, 운영 요청 시에만 추가 소명 입력을 허용하는 안을 기본으로 한다.
- 운영 판단 결과:
  1. 완료 인정 → `confirmed`
  2. 거래 미완료/불발 인정 → `cancelled` 또는 `available` 복귀
  3. 악의적 허위 완료 요청 → 경고/기능 제한 후보
- 분쟁 처리 중에도 원본 예약/채팅/상태 이력은 변경 불가 이력으로 보존해야 한다.

### 48.5 완료 후 후기/신뢰 반영 규칙
- 후기 작성 가능 시점은 `confirmed` 직후부터 시작한다.
- 자동확정 완료도 정상 완료 건으로 집계하되, 내부적으로는 `autoConfirmed=true` 플래그를 저장해 분쟁률 분석에 활용한다.
- `disputed`로 종료된 거래는 후기 작성 허용 여부를 운영정책으로 분리할 수 있으나, MVP에서는 운영 확정 후에만 후기 작성 가능하도록 제한하는 것이 단순하다.

### 48.6 재오픈/부분거래 규칙
- `completed` 확정 매물은 직접 재오픈하지 않고 `복제 후 새 매물 생성`을 기본 원칙으로 한다.
- 단, 수량형 매물에서 일부만 거래된 경우에는 `partial_completed` 내부 이벤트를 기록하고, 남은 수량으로 동일 매물을 계속 `available` 상태 유지하는 정책을 검토할 수 있다.
- MVP에서는 복잡도 축소를 위해 **부분거래를 지원하더라도 작성자가 잔여 수량을 수동 수정하고 재오픈 확인을 거치는 방식**을 우선안으로 한다.
- 예약 상대가 이탈하여 거래가 불발된 경우, 작성자는 사유 코드를 남기고 `available`로 복귀할 수 있어야 한다.

### 48.7 데이터/API 시사점
- `Listing.status` 공개 enum과 별도로 내부 완료 확인 단계를 표현할 보조 필드(`completionStage` 등) 검토 필요
- `POST /listings/{listingId}/complete` 이후 상대방 액션 API 후보:
  - `POST /trade-completions/{completionId}/confirm`
  - `POST /trade-completions/{completionId}/dispute`
- 완료 자동확정 배치/워커가 필요하며, 자동 처리도 감사 로그를 남겨야 한다.

## 50. `reserved` 상태 신규 문의 허용 정책(초안)
### 49.1 목표
- 판매자/구매자가 실제 우선 협의 상대를 보호하면서도, 예약 불발 시 거래 기회를 완전히 잃지 않게 한다.
- 사용자에게 현재 문의 가능 여부를 예측 가능하게 보여준다.

### 49.2 기본 정책안
- `reserved` 상태에서도 **신규 문의는 제한적으로 허용**한다.
- 다만 신규 문의 사용자는 즉시 예약 확정 상대가 될 수 없고, 대기열 성격의 `open` 채팅만 생성 가능하다.
- 이미 `reserved` 상태인 매물 상세에는 `현재 다른 상대와 우선 거래 중` 문구를 명확히 노출한다.
- `pending_trade` 상태부터는 신규 채팅 생성을 차단한다.

### 49.3 UX/CTA 규칙
| 상태 | 비참여 사용자 CTA | 설명 |
|---|---|---|
| available | 채팅 시작 | 일반 문의 가능 |
| reserved | 문의 남기기 또는 대기 등록 | 즉시 거래 확정 불가 문구 동반 |
| pending_trade | 문의 불가 / 유사 매물 보기 | 실제 거래 직전 단계 |
| completed/cancelled | 유사 매물 보기 | 종결 |

### 49.4 상세 동작 규칙
- `reserved` 상태 신규 채팅 생성 시 시스템은 구매자에게 “거래가 불발되면 답변을 받을 수 있습니다” 수준의 기대값을 안내한다.
- 작성자는 `reserved` 상태 채팅 목록에서 **현재 우선 상대**와 **대기 문의**를 구분해 볼 수 있어야 한다.
- 우선 상대와 예약이 취소/만료되면 대기 문의 채팅은 다시 일반 `open` 문의로 전환되고, 작성자는 다음 상대를 선택해 `reserved`로 재진입할 수 있다.
- 대기 문의가 많을 경우 스팸성 자동 메시지를 줄이기 위해 quick-reply 또는 일괄 안내 문구를 지원할 수 있다.

### 49.5 랭킹/노출 규칙 연계
- `reserved` 매물은 기본 검색 결과에서 후순위 유지
- 사용자가 `reserved 포함` 필터를 켠 경우 노출 가능
- 찜 사용자와 기존 문의 참여자에게는 상태 복귀(`reserved -> available`)를 우선 알릴 수 있다.

### 49.6 운영/안전 고려사항
- `reserved` 상태를 악용해 허위 희소성(이미 거래 중인 척) 연출하는 패턴은 탐지 대상이다.
- 장시간 `reserved` 유지 후 반복 복귀하는 매물은 랭킹 감점 또는 운영 검토 후보가 된다.
- 우선 거래 상대를 자주 바꾸는 사용자는 신뢰 지표 또는 운영 이상징후 룰에 반영 가능하다.

## 51. 객체별 권한 / 가시성 / 행위 가능 조건 매트릭스(초안)
### 50.1 목적
- 화면 요구사항, API 인가 정책, 운영 정책이 서로 다른 기준을 쓰지 않도록 공통 권한 모델을 정의한다.
- “누가 무엇을 볼 수 있고, 무엇을 바꿀 수 있는가”를 객체 단위로 명확히 해 구현 시 해석 차이를 줄인다.

### 50.2 역할 계층 정의
- **비회원(Guest)**: 로그인 이전 탐색 사용자
- **회원(Member)**: 로그인 완료 일반 사용자
- **거래참여자(Participant)**: 특정 매물의 채팅/예약/완료 흐름에 실제 참여한 사용자
- **작성자(Owner)**: 해당 매물 작성자
- **운영자(Staff)**: CS Operator / Moderator / Senior Moderator / Admin
- 하나의 사용자는 객체마다 다른 역할을 가질 수 있으며, 인가 판단은 “전역 역할 + 객체 관계 + 현재 상태”의 조합으로 결정한다.

### 50.3 권한 판단 우선순위
1. 법적/운영 강제 제한(`blocked`, `report_locked`, 계정 정지)
2. 객체 소유/참여 여부(작성자, 거래참여자, 신고 당사자)
3. 공개 범위(`public`, `hidden`, `blocked`, `noindex`와는 별개)
4. 기능별 추가 상태 조건(`reserved`, `pending_trade`, `completed` 등)
5. 사용자 개인 설정(알림, 뮤트 등)

### 50.4 매물(Listing) 권한 매트릭스
| 대상/액션 | 비회원 | 일반 회원 | 거래참여자 | 작성자 | 운영자 |
|---|---|---|---|---|---|
| 공개 목록 조회 | 가능(정책 범위 내) | 가능 | 가능 | 가능 | 가능 |
| 비공개/숨김 매물 조회 | 불가 | 불가 | 참여 기록 화면에서 제한 가능 | 본인 것만 가능 | 권한 범위 내 가능 |
| 매물 상세 조회 | 부분 가능 | 가능 | 가능 | 가능 | 가능 |
| 채팅 시작 | 불가 | 상태 조건 충족 시 가능 | 이미 참여 중이면 재진입 | 가능하지 않음(본인 매물) | 테스트/운영 목적 제외 불가 |
| 찜 | 불가 | 가능 | 가능 | 가능(정책 선택) | 불가 |
| 수정 | 불가 | 불가 | 불가 | 상태 조건 충족 시 가능 | 운영 수정은 제한적으로 가능 |
| 상태 변경 | 불가 | 불가 | 불가 | 가능 | 강제 변경 가능 |
| 숨김/차단 | 불가 | 불가 | 불가 | 자가 숨김 가능 | 정책 조치 가능 |
| 삭제/철회 | 불가 | 불가 | 불가 | 소프트 삭제/취소 가능 | 복구/강제종료 가능 |

세부 규칙:
- 작성자는 본인 매물에 채팅 시작을 할 수 없다.
- `hidden`, `blocked` 매물은 공개 검색 결과에서 제외되며, 직접 URL 접근도 제한한다.
- 거래참여자는 종결 후에도 자신이 참여한 거래 기록으로서 제한 조회 권한을 가질 수 있다.

### 50.5 채팅(ChatRoom/Message) 권한 매트릭스
| 대상/액션 | 비회원 | 일반 회원 | 채팅 참여자 | 매물 작성자(=참여자일 수도 있음) | 운영자 |
|---|---|---|---|---|---|
| 채팅방 목록 조회 | 불가 | 본인 참여방만 | 가능 | 가능 | 권한 범위 내 가능 |
| 메시지 조회 | 불가 | 본인 참여방만 | 가능 | 가능 | 신고/운영 목적 범위 내 가능 |
| 메시지 발송 | 불가 | 참여 중이고 상태 허용 시 가능 | 가능 | 가능 | 일반적으로 불가 |
| 예약 제안 | 불가 | 참여 중일 때만 가능 | 가능 | 가능 | 불가 |
| 차단/신고 | 불가 | 참여 중일 때만 가능 | 가능 | 가능 | 운영 액션 별도 |
| 채팅 잠금 해제 | 불가 | 불가 | 불가 | 불가 | 권한 범위 내 가능 |

세부 규칙:
- 채팅 메시지 원문은 참여자와 운영자 외 공개 금지다.
- 운영자도 기본은 마스킹/필요 최소 열람 원칙을 적용한다.
- `report_locked` 상태에서는 참여자도 신규 메시지 발송 불가다.

### 50.6 예약/완료/후기 권한 규칙
#### 예약(Reservation)
- 생성/수정/확정/취소는 해당 채팅방 참여자만 가능
- 운영자는 강제 만료/강제 취소 가능하나, 사용자 예약 내용을 임의 수정하지 않는다.
- 예약 상세는 참여자와 운영자만 조회 가능
- 예약 메타데이터 중 캐릭터명/장소/시간은 공개 매물 영역에 노출하지 않는다.

#### 거래 완료(TradeCompletion)
- 완료 요청은 실제 진행 채팅방 참여자만 가능
- 완료 확정/이의제기는 상대방 또는 운영자만 가능
- 완료 기록은 작성자/상대방/운영자만 전체 조회 가능
- 완료 결과 중 후기 가능 여부, 분쟁 여부는 프로필 집계에 반영되되 민감 사유는 비공개다.

#### 후기(Review)
- 후기 작성은 해당 completion의 당사자만 가능
- 공개 후기는 전체 회원/비회원에게 일부 또는 전체 공개 가능하나, 작성자 식별/운영 숨김 여부 정책을 따른다.
- 후기 신고는 누구나 가능하게 열 수 있으나, 무분별한 신고 방지를 위해 로그인은 필수로 본다.

### 50.7 신고/운영 데이터 권한 규칙
| 대상/액션 | 신고자 | 피신고자 | 일반 회원 | 운영자 |
|---|---|---|---|---|
| 내 신고 접수 여부 확인 | 가능 | 불가 | 불가 | 가능 |
| 신고 상세 원문 조회 | 본인 작성분 일부만 | 원칙적 불가 | 불가 | 가능 |
| 처리 상태 조회 | 접수/처리중/완료 수준만 | 제재 결과 본인 건만 | 불가 | 가능 |
| 증빙 열람 | 본인 제출분만 | 원칙적 불가 | 불가 | 가능 |
| 제재 이력 전체 조회 | 불가 | 본인 계정 관련 일부만 | 불가 | 가능 |

세부 규칙:
- 피신고자에게는 신고자 신원과 내부 메모를 공개하지 않는다.
- 신고자는 “조치 완료/미조치” 등 결과 수준만 확인 가능하며, 구체적 제재 수단 전체를 공개하지 않을 수 있다.
- 운영 내부 메모, 위험 점수, 탐지 룰 적중 여부는 외부 비공개다.

### 50.8 차단/제재/탈퇴 상태의 우선 규칙
- **차단(Block)**: 사용자 간 상호작용 제한. 기존 기록 조회는 제한적으로 허용될 수 있으나 신규 접촉은 차단한다.
- **기능 제한(Restricted)**: 매물 등록/채팅/후기/신고 중 일부만 제한될 수 있으며, 제한 사유 코드가 필요하다.
- **정지(Suspended)**: 로그인 유지 여부와 무관하게 쓰기 액션 전면 제한, 읽기 범위는 최소한으로 축소
- **탈퇴(Withdrawn)**: 본인 공개 프로필 접근은 종료되나, 거래/운영 로그는 비식별 상태로 남길 수 있다.
- 우선순위는 `정지 > 기능 제한 > 차단 > 일반 상태`로 적용한다.

### 50.9 화면/백엔드 설계 시사점
- 화면은 버튼 노출 여부만으로 권한을 판단하지 않고, 서버도 동일 규칙으로 재검증해야 한다.
- 각 상세 API 응답에는 현재 사용자의 가능 액션 집합(`availableActions`)을 포함하는 구조를 검토한다.
- 예시:
```json
{
  "listingId": "listing_123",
  "status": "reserved",
  "viewerRole": "member",
  "availableActions": ["favorite", "report", "create_waiting_chat"]
}
```
- 이 방식은 클라이언트 화면 제어, QA 시나리오, 권한 테스트케이스 작성에 직접 활용 가능하다.

### 50.10 QA/운영 테스트케이스 파생 포인트
- 비회원이 직접 URL로 `completed` 매물 상세 접근 시 어느 수준까지 보이는가
- 차단된 상대가 기존 채팅방에서 과거 메시지를 읽을 수 있는가
- 정지 계정이 알림함/내 거래 기록을 어디까지 조회 가능한가
- 신고된 후기의 공개/비공개 전환 시 프로필 집계가 즉시 일관되게 반영되는가
- 동일 사용자가 작성자이면서 참여자인 케이스(구매 매물 구조 포함)에서 액션 중복 노출이 없는가

## 52. 자동화 타이머 / 배치 / 시스템 액션 정책(초안)
### 51.1 목표
- 사용자가 매번 수동으로 상태를 정리하지 않아도 거래 흐름이 자연스럽게 종결되도록 한다.
- 자동 처리로 인한 혼란을 줄이기 위해, 모든 시스템 액션은 예측 가능하고 되돌릴 수 있는 범위에서 동작해야 한다.
- 예약/완료/알림/노출의 시간 기준을 명시해 운영 정책과 백엔드 스케줄러 설계 기준을 통일한다.

### 51.2 자동화 대상 이벤트
| 대상 | 트리거 | 기본 시스템 액션 | 사용자 노출 |
|---|---|---|---|
| 예약 제안 | `expiresAt` 도달 | `proposed -> expired` | 채팅방 시스템 메시지 + 알림함 |
| 확정 예약 | 예약 시각 경과 후 미정리 | `trade_due` 유지 + 불발/완료 리마인드 | 거래 후속처리 배너 |
| 완료 요청 | 상대방 무응답 타임아웃 | `requested -> confirmed(auto)` | 완료 확정 알림 |
| reserved 장기 유지 | 기준 시간 초과 | 작성자에게 상태 점검 리마인드 | 매물 관리 알림 |
| 오래된 available 매물 | 활동 없음 기준 도달 | 만료 예정 안내 또는 자동 비노출 후보 | 내 매물 배너/알림 |
| 후기 미작성 | 작성 마감 임박 | 1회 또는 제한적 리마인드 | 알림함/푸시 |

### 51.3 예약 관련 타이머 기본안
- 예약 제안(`proposed`)은 제안 시점에 `expiresAt`을 가져야 한다.
- 기본안: 예약 제안 만료는 **제안 후 12시간 또는 예약 시각 2시간 전 중 더 이른 시점**을 사용한다(가정).
- 만료 시:
  - 예약 상태는 `expired`
  - 채팅 상태는 `open` 또는 직전 협의 상태로 복귀
  - 매물 상태는 해당 예약이 유일한 활성 흐름이었다면 `available` 복귀 후보
- 사용자가 만료 직전 응답 중인 상황을 고려해, 만료 직전 1회 연장 제안 UX를 둘 수 있다(Post-MVP)

### 51.4 거래 시각 경과 후 후속 처리
- `confirmed` 예약의 `scheduledAt`이 지나면 채팅 상태를 `trade_due`로 간주한다.
- 예약 시각 경과 후 시스템은 아래 순서로 후속 액션을 유도한다.
  1. 예약 시각 직후: “거래가 끝났나요?” 인앱 알림
  2. +30분(가정): 완료/노쇼/불발 선택 리마인드
  3. +24시간(가정): 미정리 거래를 운영 분석용 `stale_trade_due` 후보로 집계
- 시스템은 거래 결과를 자동으로 `cancelled` 처리하지 않는다. 다만 장기 미정리 건은 작성자/참여자에게 정리 행동을 요구할 수 있어야 한다.

### 51.5 완료 자동확정 타이머
- 완료 요청 후 상대방 응답 기한 기본안은 **24시간**이다(가정).
- 자동확정 전:
  - 상대방에게 최소 1회 리마인드 발송
  - 이의 제기 CTA를 명확히 제공
- 자동확정 실행 시:
  - `TradeCompletion.confirmationMethod=auto` 성격 필드 저장 후보
  - `Listing.status=completed` 최종 반영
  - 후기 작성 가능 상태 오픈
  - 감사 로그 및 사용자 알림 생성
- 자동확정 후에도 신고는 가능하되, 상태 자체는 운영 개입 없이 사용자가 임의 롤백할 수 없게 한다.

### 51.6 휴면/장기 방치 매물 정책
- `available` 상태에서 장기간 활동 없는 매물은 거래 품질을 저하시키므로 만료 정책이 필요하다.
- 기본안:
  - 최근 활동 7일 경과 시 “여전히 거래 가능 여부 확인” 리마인드
  - 최근 활동 14일 경과 시 `stale_available` 플래그 부여 및 랭킹 감점
  - 최근 활동 30일 경과 시 자동 비노출 또는 `expired` 성격 내부 상태 후보
- 자동 비노출 시에도 작성자는 내 매물에서 재게시/복제 후 재등록 가능해야 한다.
- 자동 만료된 매물은 공개 SEO 인덱스에서 우선 제거 대상이다.

### 51.7 후기/알림 리마인드 제한
- 후기 요청 리마인드는 과도하지 않게 제한한다.
- 기본안:
  - 완료 확정 직후 1회
  - 48시간 후 미작성 시 1회
  - 마감 24시간 전 1회
- 최대 3회 초과 금지, 사용자가 명시적으로 알림을 끈 경우 인앱 알림만 유지한다.

### 51.8 자동화 처리 실패/재시도 원칙
- 자동화 액션은 워커/배치 실패 시 재시도 가능해야 하나, 중복 실행으로 상태가 꼬이지 않도록 멱등하게 설계해야 한다.
- 각 시스템 액션은 최소한 아래를 기록해야 한다.
  - 대상 객체
  - 예정 시각
  - 실제 실행 시각
  - 실행 결과(success/failure/skipped)
  - 재시도 횟수
  - 실패 사유 코드
- 반복 실패 건은 운영 대시보드 또는 기술 알림 채널로 집계되어야 한다.

### 51.9 시스템 액션 감사/가시성 원칙
- 사용자가 체감하는 상태 변경은 가능하면 채팅 시스템 메시지, 알림함, 내 거래 화면 중 최소 1곳에서 근거를 볼 수 있어야 한다.
- 운영자 화면에서는 “사용자 액션”과 “시스템 자동 처리”를 명확히 구분해 보여야 한다.
- 자동화는 사용자 의도를 대체하는 것이 아니라 미정리 흐름을 종결/보조하는 역할로 제한한다.

## 53. 화면-API 매핑 / 기능명세 파생 프레임(초안)
### 52.1 목적
- PRD에서 기능명세, 화면설계, API 명세로 내려갈 때 누락되기 쉬운 연결점을 미리 정의한다.
- 각 화면이 어떤 읽기/쓰기 API와 어떤 상태 조건에 의존하는지 구조화해 개발 범위와 QA 범위를 명확히 한다.

### 52.2 거래소 목록 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 매물을 탐색하고 상세/채팅으로 진입 |
| 핵심 읽기 API | `GET /listings` |
| 핵심 쓰기 API | `POST /listings/{listingId}/favorite`, `DELETE /listings/{listingId}/favorite` |
| 주요 상태 의존성 | `available`, `reserved` 중심 노출 |
| 권한 가드레일 | 비회원은 읽기만 가능, 찜/채팅 CTA는 로그인 유도 |
| 파생 명세 포인트 | 필터 파라미터, 정렬 옵션, 카드별 `availableActions`, 배지 노출 규칙 |

### 52.3 매물 상세 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 거래 판단과 즉시 행동 유도 |
| 핵심 읽기 API | `GET /listings/{listingId}` |
| 핵심 쓰기 API | `POST /listings/{listingId}/chats`, `POST /listings/{listingId}/favorite`, `POST /reports` |
| 주요 상태 의존성 | 상태별 CTA(`available`, `reserved`, `pending_trade`, `completed`, `cancelled`) |
| 권한 가드레일 | 작성자는 자기 매물에 채팅 시작 불가 |
| 파생 명세 포인트 | 공개 정보 레벨, 판매자 신뢰 배지, reserved 안내 문구, 유사 매물 추천 슬롯 |

### 52.4 매물 등록/수정 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 빠른 등록과 안전한 수정 |
| 핵심 읽기 API | 카탈로그/서버 후보 조회(후보 API 필요) |
| 핵심 쓰기 API | `POST /listings`, `PATCH /listings/{listingId}` |
| 주요 상태 의존성 | 진행 상태에 따라 수정 가능 필드 제한 |
| 권한 가드레일 | 작성자만 접근 가능, `pending_trade` 이상에서는 핵심 필드 수정 제한 |
| 파생 명세 포인트 | validation 메시지, 임시저장 여부, 이미지 업로드 정책, 금칙어/민감정보 사전검사 |

### 52.5 채팅 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 협의, 예약, 완료, 신고를 하나의 흐름에서 처리 |
| 핵심 읽기 API | `GET /chats/{chatRoomId}`, `GET /chats/{chatRoomId}/messages` |
| 핵심 쓰기 API | `POST /chats/{chatRoomId}/messages`, `POST /chats/{chatRoomId}/reservations`, `POST /reports`, `POST /chats/{chatRoomId}/block` |
| 주요 상태 의존성 | `open`, `reservation_proposed`, `reservation_confirmed`, `trade_due`, `deal_completed`, `report_locked` |
| 권한 가드레일 | 참여자만 접근 가능, `report_locked` 시 읽기 전용 |
| 파생 명세 포인트 | 메시지 타입별 렌더링, 예약 카드 상호작용, 시스템 메시지 포맷, 입력창 disabled 사유 |

### 52.6 내 매물 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 작성자 기준 매물 운영 현황 관리 |
| 핵심 읽기 API | `GET /me/listings` |
| 핵심 쓰기 API | `PATCH /listings/{listingId}`, `POST /listings/{listingId}/status` |
| 주요 상태 의존성 | 진행중/예약중/완료/취소 탭 구분 |
| 권한 가드레일 | 본인 매물만 조회/조작 가능 |
| 파생 명세 포인트 | 탭 기준 상태 매핑, 대량 액션 허용 여부, 지표 캐시 신선도, 재등록 CTA |

### 52.7 내 거래 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 참여자 관점에서 액션이 필요한 거래를 우선 관리 |
| 핵심 읽기 API | 별도 `GET /me/trades` 또는 `GET /chats`, `GET /reservations` 조합 후보 |
| 핵심 쓰기 API | `POST /reservations/{reservationId}/confirm`, `POST /reservations/{reservationId}/cancel`, `POST /listings/{listingId}/complete` |
| 주요 상태 의존성 | 예약 응답 대기, 거래대기, 완료 확인 대기 |
| 권한 가드레일 | 참여 거래만 노출 |
| 파생 명세 포인트 | 우선순위 정렬 규칙, 액션 필요 배지, 시한 임박 표기, 완료 확인 플로우 |

### 52.8 알림함 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 푸시를 놓쳐도 거래 진행 이벤트를 복구 |
| 핵심 읽기 API | `GET /notifications` |
| 핵심 쓰기 API | `POST /notifications/read` |
| 주요 상태 의존성 | unread/read, notificationType |
| 권한 가드레일 | 본인 알림만 조회 가능 |
| 파생 명세 포인트 | 딥링크 매핑표, 묶음 알림 표시 규칙, 만료된 딥링크 fallback 처리 |

### 52.9 프로필/신뢰 화면
| 항목 | 내용 |
|---|---|
| 화면 목적 | 거래 판단을 위한 신뢰 요약 제공 |
| 핵심 읽기 API | `GET /users/{userId}/reviews`, 사용자 프로필 조회 API 후보 |
| 핵심 쓰기 API | 차단/신고/프로필 수정 관련 API 후보 |
| 주요 상태 의존성 | 공개/비공개 후기, 제재 노출 정책 |
| 권한 가드레일 | 민감 운영 데이터 비노출 |
| 파생 명세 포인트 | 배지 산정값, 후기 페이징, 마스킹 규칙, 비회원 공개 레벨 |

### 52.10 관리자 백오피스
| 항목 | 내용 |
|---|---|
| 화면 목적 | 신고 처리, 제재, 복구, 감사 추적 |
| 핵심 읽기 API | `GET /admin/reports`, `GET /admin/reports/{reportId}`, `GET /admin/users/{userId}/history` |
| 핵심 쓰기 API | `POST /admin/reports/{reportId}/actions`, `POST /admin/listings/{listingId}/hide`, `POST /admin/users/{userId}/restrict` |
| 주요 상태 의존성 | reportStatus, priority, moderation level |
| 권한 가드레일 | 운영 역할 레벨별 접근 제어 필수 |
| 파생 명세 포인트 | 처리 단계별 폼, 사유 코드 체계, 대량 작업 제한, 감사 로그 diff |

### 52.11 명세 산출물 연결 원칙
- 각 화면 명세서는 최소한 다음을 포함해야 한다.
  1. 진입 조건/권한
  2. API 요청/응답 계약
  3. 상태별 UI 차이
  4. 예외/오류 문구
  5. 분석 이벤트
  6. empty/loading/error 상태
- API 명세서는 화면 기준 유스케이스와 1:N 관계일 수 있으나, 어떤 화면에서 어떤 응답 필드가 필요한지 역추적 가능해야 한다.
- QA 테스트케이스는 본 섹션의 화면별 상태/권한 조합을 기준으로 파생한다.

## 54. 공통 화면 상태 / 응답 컨텍스트 표준(초안)
### 53.1 목적
- 화면 설계, API 응답, QA 케이스가 서로 다른 상태명을 쓰지 않도록 공통 UI 상태 모델을 정의한다.
- 단순 성공/실패 외에 `권한 제한`, `정책 제한`, `부분 데이터`, `복구 불가` 상태를 명시해 모바일 UX와 운영 메시지를 일관되게 만든다.

### 53.2 공통 화면 상태 정의
| 상태 | 의미 | 대표 화면 예시 | 사용자 기대 액션 |
|---|---|---|---|
| `loading` | 최초/재조회 중 | 목록, 상세, 채팅 진입 직후 | 대기, skeleton 노출 |
| `ready` | 정상 표시 가능 | 대부분의 정상 화면 | 주요 CTA 사용 |
| `empty` | 데이터는 정상 조회됐으나 항목 없음 | 검색결과 없음, 알림 없음 | 필터 변경, 새 등록, 복귀 |
| `partial` | 핵심 정보는 보이지만 일부 영역 비노출/미지원 | 삭제된 사용자 프로필 요약, 종료 매물 소프트 랜딩 | 기록 조회 또는 대체 액션 |
| `error_retryable` | 네트워크/일시 오류 | 목록 재시도, 메시지 재전송 | 다시 시도 |
| `error_terminal` | 대상 없음/영구 접근 불가 | 삭제된 매물, 권한 없는 운영 URL | 뒤로 가기, 홈 이동 |
| `policy_blocked` | 기능은 존재하지만 현재 정책상 차단 | 정지 계정의 채팅 발송, `report_locked` 입력창 | 제한 사유 확인 |
| `auth_required` | 로그인 필요 | 비회원 찜/채팅 진입 | 로그인 |

### 53.3 화면별 최소 상태 조합
- **거래소 목록**: `loading` / `ready` / `empty` / `error_retryable`
- **매물 상세**: `loading` / `ready` / `partial` / `error_terminal` / `auth_required(행동 시)`
- **채팅 화면**: `loading` / `ready` / `empty(메시지 없음)` / `policy_blocked` / `error_terminal`
- **내 매물/내 거래**: `loading` / `ready` / `empty` / `error_retryable`
- **운영 백오피스**: `loading` / `ready` / `empty` / `policy_blocked` / `error_terminal`

### 53.4 공통 응답 컨텍스트 필드 후보
화면 주요 읽기 API는 도메인 객체 외에 아래 컨텍스트를 함께 반환하는 패턴을 검토한다.

```json
{
  "viewerContext": {
    "viewerRole": "member",
    "relationshipToResource": "listing_participant",
    "accountStatus": "active"
  },
  "availableActions": ["send_message", "report", "mute_chat"],
  "policyHints": [
    {
      "code": "RESERVED_WAITING_ONLY",
      "severity": "info",
      "message": "현재 다른 상대와 우선 거래 중입니다. 문의는 남길 수 있지만 즉시 예약 확정되지는 않습니다."
    }
  ]
}
```

### 53.5 `availableActions` 표준화 원칙
- 버튼 노출 여부를 프론트가 임의 추론하지 않고 서버가 허용 액션 집합을 내려주는 구조를 우선 검토한다.
- 액션명은 화면 문구가 아니라 **도메인 동사 기반**으로 관리한다.
- 예시 후보:
  - Listing: `create_chat`, `favorite`, `edit_listing`, `change_listing_status`, `duplicate_listing`, `report_listing`
  - Chat: `send_message`, `propose_reservation`, `confirm_reservation`, `cancel_reservation`, `mark_completed`, `report_chat`, `block_user`
  - Review: `create_review`, `edit_review`, `report_review`
  - Admin: `triage_report`, `hide_listing`, `restrict_user`, `unlock_chat`

### 53.6 정책 힌트/제한 사유 코드 원칙
- `policyHints.code`는 화면 배너, 토스트, 고객센터 응답, 운영 로그에서 재사용 가능해야 한다.
- 사용자 노출 가능한 사유와 내부 운영 사유를 분리한다.
- 후보 코드 예시:
  - `AUTH_REQUIRED`
  - `OWNER_CANNOT_CHAT_OWN_LISTING`
  - `LISTING_ALREADY_COMPLETED`
  - `RESERVED_WAITING_ONLY`
  - `CHAT_LOCKED_BY_REPORT`
  - `ACCOUNT_RESTRICTED_TEMPORARILY`
  - `SENSITIVE_INFO_BLOCKED`
- 하나의 화면에 여러 제한이 동시에 존재할 수 있으므로 배열 구조를 유지한다.

### 53.7 QA/명세 파생 포인트
- 각 화면 명세서는 `화면 상태 x 사용자 역할 x 도메인 상태` 조합표를 가져야 한다.
- 각 읽기 API 명세서는 객체 본문 외에 `viewerContext`, `availableActions`, `policyHints` 포함 여부를 명시한다.
- QA는 “버튼 숨김”만 보지 않고, 숨김/비활성/경고 배너/에러 코드가 일관되게 매핑되는지 확인해야 한다.

## 55. 삭제 / 탈퇴 / 비식별화 / 복구 정책(초안)
### 54.1 목적
- 사용자 삭제 요청, 운영 숨김, 계정 탈퇴, 법적 보관 의무, 거래 분쟁 대응이 서로 충돌하지 않도록 객체별 수명주기 원칙을 정의한다.
- DB 스키마 단계에서 hard delete를 최소화하고, 공개 비노출과 내부 보관을 분리한다.

### 54.2 공통 원칙
1. 거래/신고/감사 추적이 필요한 핵심 객체는 hard delete보다 soft delete 또는 비식별화를 기본으로 한다.
2. 사용자에게 보이는 “삭제됨”과 내부 저장소에서의 “보관 중”은 다른 개념으로 관리한다.
3. 운영 조치 복구는 원본 레코드 덮어쓰기가 아니라 **복구 이벤트 추가** 방식으로 저장한다.
4. 법적 요청/정책 변경 시 객체별 보관 기간과 파기 가능 범위를 별도 문서로 분리 가능해야 한다.

### 54.3 객체별 기본 정책 매트릭스
| 객체 | 사용자 자가 삭제 | 운영 숨김/비공개 | 내부 보관 원칙 | 공개 화면 표시 |
|---|---|---|---|---|
| Listing | hard delete 불가, 취소/숨김/복제 후 재등록 유도 | 가능 | 거래/신고 연결 위해 보관 | 상태에 따라 비노출 또는 종료 랜딩 |
| ChatMessage | 자가 삭제는 제한적(soft delete 표시) | 신고/정책 위반 시 마스킹 가능 | 분쟁 대응 위해 원본/이력 보관 | `삭제된 메시지` 또는 마스킹 |
| ChatRoom | 참여자 임의 삭제 불가, 개인 보관함에서 숨김만 허용 | 잠금/비공개 가능 | 거래 기록 보관 | 참여자 화면에서는 종료/잠금 상태로 표시 |
| Reservation | 삭제 불가 | 운영 강제 취소 가능 | 상태 이력 포함 보관 | 취소/만료 이력 표시 |
| TradeCompletion | 삭제 불가 | 분쟁 상태 전환/비공개 가능 | 신뢰도/감사 목적 보관 | 완료/분쟁 상태로 표시 |
| Review | 자가 삭제 제한, 짧은 수정 유예만 | 숨김/비공개 가능 | 집계 근거와 원본 보관 분리 | 숨김 또는 비노출 처리 |
| Report | 신고자 취소 요청은 가능하나 삭제 아님 | 운영 처리 상태 변경 | 감사 목적 장기 보관 | 신고자에게 상태만 표시 |
| UserProfile | 탈퇴 시 비식별화 | 운영 제한 가능 | 핵심 거래 참조 위해 연결키 유지 | `탈퇴한 사용자` 등으로 치환 |

### 54.4 계정 탈퇴 정책 기본안
- 탈퇴 즉시:
  - 로그인/쓰기 기능 차단
  - 공개 프로필 비식별화
  - 활성 매물은 자동 `cancelled` 또는 `hidden` 처리 후보
  - 진행 중 거래가 있으면 탈퇴 전 경고 또는 탈퇴 유예 정책 필요
- 탈퇴 후 유지 데이터:
  - 거래, 신고, 감사 로그, 후기 집계 근거, 운영 제재 이력
  - 단, 공개 화면에서는 식별 가능한 닉네임/아바타/소개는 제거
- 탈퇴 계정 참조 표기 예시:
  - `탈퇴한 사용자`
  - `알 수 없는 사용자`는 지양하고, 거래 이력의 맥락이 보이도록 최소 의미를 유지

### 54.5 매물/채팅 복구 정책
- 운영 숨김된 매물은 `복구 가능` 상태를 가져야 하며, 복구 시 기존 ID와 이력을 유지한다.
- 사용자가 취소한 매물은 직접 되살리기보다 `복제 후 재등록`을 기본으로 한다.
- 신고로 잠긴 채팅은 운영 해제 시 `report_locked -> open` 또는 적절한 후속 상태로 이력형 복구한다.
- 복구된 객체는 랭킹/신뢰 집계에서 불이익 복원 여부를 별도 계산해야 한다.

### 54.6 데이터 모델 시사점
- 주요 테이블에 아래 필드 후보를 검토한다.
  - `deletedAt`
  - `hiddenAt`
  - `anonymizedAt`
  - `restoredAt`
  - `visibilityReasonCode`
  - `retentionUntil`
- 프로필 비식별화는 FK 파손 없이 표시용 컬럼만 치환하는 전략이 유리하다.
- 메시지/후기 마스킹은 원문 컬럼 삭제 대신 별도 `maskedBodyText`, `moderationSnapshot` 저장을 검토한다.

### 54.7 운영/정책 문서 파생 포인트
- 개인정보 처리방침: 보관 기간, 탈퇴 후 보관 항목, 비식별화 범위
- 운영정책: 숨김/복구 기준, 오조치 복구 원칙, 신고 원본 보존 기준
- DB 문서: soft delete 인덱스, 보관 배치, 비식별화 대상 필드
- 화면명세: 삭제된 메시지/탈퇴 사용자/종료 매물의 표시 문구

## 56. MVP 출시 게이트 / 수용 기준(초안)
### 55.1 목적
- MVP 범위를 단순 기능 목록이 아니라 **출시 가능 여부를 판단하는 기준 세트**로 정의한다.
- 제품, 디자인, 서버, QA, 운영이 서로 다른 기준으로 “준비 완료”를 선언하지 않도록 공통 게이트를 둔다.

### 55.2 출시 게이트 묶음
| 게이트 | 설명 | 통과 기준 초안 |
|---|---|---|
| G1 핵심 사용자 흐름 | 등록→문의→예약→완료→후기 흐름 동작 | 치명 결함 없이 E2E 1개 이상 시나리오 통과 |
| G2 상태 정합성 | 매물/채팅/예약/완료 상태 모순 없음 | 정의된 상태 전이 테스트 100% 통과 |
| G3 운영 대응 가능성 | 신고/숨김/제재/복구 가능 | 운영자 기본 시나리오 전부 재현 가능 |
| G4 모바일 사용성 | 한 손 사용/핵심 CTA 접근성 | 주요 화면 모바일 기준 사용성 리뷰 통과 |
| G5 성능/안정성 | 핵심 API와 채팅이 허용 성능 내 동작 | 34장 성능 목표의 초기 기준 충족 |
| G6 안전/정책 가드레일 | 민감정보/외부연락처/신고 흐름 최소 동작 | 차단/경고/신고 정책의 기본 동작 확인 |
| G7 분석/관측성 | 핵심 퍼널 추적 가능 | 25장의 필수 이벤트 누락 없이 적재 |

### 55.3 기능 수용 기준: 판매자 핵심 시나리오
**시나리오 A: 판매 매물 등록 후 거래 완료**
1. 판매자가 모바일에서 3분 이내 매물 등록 가능
2. 등록 직후 공개 목록/상세에서 본인 매물 확인 가능
3. 구매자 문의 발생 시 판매자는 채팅 알림과 내 거래 항목을 확인 가능
4. 채팅에서 예약 제안/확정 시 매물 상태와 채팅 상태가 정의대로 연동
5. 거래 완료 요청 후 상대 확인 또는 자동확정이 동작
6. 완료 후 후기 CTA와 기록 화면이 열린다.

수용 기준:
- 치명 오류 없이 완료 가능
- 상태 전환 이력과 감사 로그 남음
- 잘못된 중복 완료 요청은 멱등 또는 409 처리

### 55.4 기능 수용 기준: 구매자 핵심 시나리오
**시나리오 B: 검색 후 문의, 예약, 후기 작성**
1. 구매자는 서버/아이템 기준으로 결과를 찾을 수 있어야 한다.
2. `available` 매물에는 즉시 채팅 시작 가능해야 한다.
3. `reserved` 매물에는 대기 문의/제한 메시지가 정책대로 노출되어야 한다.
4. 예약 확정 후 시간/장소/서버 정보가 채팅과 내 거래에 일관되게 표시되어야 한다.
5. 완료 확정 후 후기 작성 기간과 공개 규칙이 안내되어야 한다.

수용 기준:
- 검색→상세→채팅 진입이 3탭 이내 기본 달성
- 예약/완료 관련 주요 알림 누락 없음
- 후기 작성 불가 상태에서는 명확한 사유 노출

### 55.5 운영 수용 기준
- 신고 접수 후 운영 큐에 우선순위와 대상 타입이 정확히 표시된다.
- 운영자는 매물 숨김, 채팅 잠금, 경고, 기간 제한, 복구 액션을 실행할 수 있어야 한다.
- 모든 운영 액션은 사유 코드, 대상, 처리자, 시각, 결과를 감사 로그에 남겨야 한다.
- P1/P2 사건에 대해 최소 임시조치가 가능해야 한다.
- 오조치 복구 후 사용자 노출 상태와 내부 이력이 모두 일관되어야 한다.

### 55.6 데이터/정합성 수용 기준
- 동일 매물에 활성 예약 중복 확정 불가
- 동일 completion에 동일 작성자의 후기 중복 생성 불가
- 종결 매물에 신규 채팅 생성 불가
- `reservedChatRoomId`, Reservation, Listing.status 간 정합성 검증 테스트 필요
- soft delete/비식별화 후에도 거래 기록 조회와 감사 추적이 깨지지 않아야 한다.

### 55.7 화면 품질 수용 기준
#### 거래소 목록
- 필수 필터: 서버, 거래유형, 카테고리, 상태
- 카드별 상태 배지/가격/서버/신뢰 신호 노출
- empty/loading/error 상태 분리 표현

#### 매물 상세
- 상태별 CTA가 46장 규칙과 일치
- 비회원 액션은 로그인 유도, 작성자 본인 액션은 관리 액션으로 대체
- 정책 제한/운영 제한 시 배너 노출

#### 채팅
- 시스템 메시지와 일반 메시지가 구분됨
- 예약 카드 상호작용 가능
- 잠금/완료 상태에서 입력창 정책이 명확히 바뀜

#### 내 거래
- 액션 필요 거래가 상단 우선 정렬
- 예약 임박, 완료 확인 대기, 분쟁 상태를 시각적으로 구분

### 55.8 KPI 초기 목표값 후보(출시 후 30일 관찰)
- 매물 등록 완료율: 70%+
- 매물 상세 대비 채팅 시작 전환율: 10%+
- 채팅 시작 대비 예약 확정 전환율: 20%+
- 예약 확정 대비 완료 확정 전환율: 50%+
- 완료 확정 대비 후기 작성률: 30%+
- 신고율: 완료 거래 100건당 5건 이하를 초기 목표로 관찰
- 노쇼율: 예약 100건당 15건 이하를 초기 목표로 관찰

※ 위 수치는 외부 벤치마크가 아닌 초기 운영 목표값 가정이며, 실제 런칭 후 보정 필요.

### 55.9 출시 차단 조건(Blockers)
- 거래 상태가 꼬여 같은 매물에 복수 완료/복수 활성 예약이 가능한 경우
- 차단/정지 사용자도 신규 채팅 또는 신규 매물 등록이 가능한 경우
- 신고 접수는 되지만 운영자가 대상 근거를 추적할 수 없는 경우
- 완료/취소 후 공개 목록에 계속 노출되는 치명적 정책 오류
- 민감정보 차단/경고가 완전히 비활성화된 상태

## 57. API 표면 우선순위 / 버전 전략(초안)
### 57.1 목적
- API 목록을 단순 나열이 아니라 **MVP에 꼭 필요한 표면**과 이후 확장 표면으로 나눠 구현 우선순위를 명확히 한다.
- 버전 정책을 미리 정의해 모바일 앱 배포 이후의 비호환 변경 리스크를 줄인다.

### 57.2 API 분류 원칙
- **Public App API**: 모바일/웹 클라이언트가 직접 호출하는 사용자용 API
- **Admin API**: 운영 백오피스 전용 API
- **Internal System API/Job**: 워커/배치/자동화 전용, 외부 비공개
- MVP 문서와 개발 일정은 Public App API를 최우선 기준으로 잡는다.

### 57.3 MVP 필수 Public App API
#### 인증/내 계정
- `POST /auth/login`
- `POST /auth/logout`
- `GET /me`
- `PATCH /me/profile`

#### 매물 탐색/관리
- `GET /listings`
- `POST /listings`
- `GET /listings/{listingId}`
- `PATCH /listings/{listingId}`
- `POST /listings/{listingId}/status`
- `POST /listings/{listingId}/favorite`
- `DELETE /listings/{listingId}/favorite`
- `GET /me/listings`

#### 채팅/예약/거래
- `POST /listings/{listingId}/chats`
- `GET /chats`
- `GET /chats/{chatRoomId}`
- `GET /chats/{chatRoomId}/messages`
- `POST /chats/{chatRoomId}/messages`
- `POST /chats/{chatRoomId}/reservations`
- `POST /reservations/{reservationId}/confirm`
- `POST /reservations/{reservationId}/cancel`
- `POST /listings/{listingId}/complete`
- `POST /trade-completions/{completionId}/confirm`
- `POST /trade-completions/{completionId}/dispute`

#### 후기/신고/알림
- `POST /trade-completions/{completionId}/reviews`
- `GET /users/{userId}/reviews`
- `POST /reports`
- `GET /me/reports`
- `GET /notifications`
- `POST /notifications/read`
- `POST /push-tokens`

### 56.4 Post-MVP 또는 확장 후보 API
- `GET /catalog/servers`
- `GET /catalog/items/suggest`
- `GET /me/trades`
- `POST /chats/{chatRoomId}/mute`
- `DELETE /chats/{chatRoomId}/mute`
- `POST /users/{userId}/block`
- `DELETE /users/{userId}/block`
- `POST /listings/{listingId}/duplicate`
- `POST /listings/{listingId}/bump`
- `GET /search/suggestions`
- `GET /users/{userId}/profile`

### 56.5 Admin API 최소 범위
- `GET /admin/reports`
- `GET /admin/reports/{reportId}`
- `POST /admin/reports/{reportId}/actions`
- `GET /admin/users/{userId}/history`
- `POST /admin/listings/{listingId}/hide`
- `POST /admin/listings/{listingId}/restore`
- `POST /admin/chats/{chatRoomId}/lock`
- `POST /admin/chats/{chatRoomId}/unlock`
- `POST /admin/users/{userId}/restrict`
- `POST /admin/users/{userId}/restore`

### 56.6 Internal System API / Job 후보
- 예약 만료 처리 job
- 완료 자동확정 job
- 후기 리마인드 job
- 휴면 매물 비노출 job
- 이상징후 탐지 score job
- 알림 fanout/dedup job

### 56.7 버전 전략 기본안
- 외부 클라이언트용 Public API는 `/v1` prefix 또는 헤더 버전 전략 중 하나를 명확히 선택해야 한다.
- 초기 운영 단순성을 위해 **경로 버전(`/v1`) 우선안**을 권장한다.
- 비호환 변경은 `/v2`처럼 명시적 버전 업을 원칙으로 하고, 같은 버전 내에서는 additive change만 허용한다.
- Admin/Internal API는 초기에는 별도 버전 없이 운영 가능하나, 외부 자동화 연동이 생기면 동일 원칙 적용 필요.

### 56.8 비호환 변경 금지 원칙
같은 major 버전에서 아래 변경은 금지 또는 deprecated 후 제거를 원칙으로 한다.
- 필수 응답 필드 제거/이름 변경
- enum 기존 값 삭제 또는 의미 변경
- 동일 상태에서 허용되던 액션의 무고지 제거
- 에러 코드 의미 변경
- 날짜/시간 포맷 변경

허용 가능한 변경 예시:
- optional 필드 추가
- 신규 enum 값 추가(단, 클라이언트 fallback 규칙 필요)
- 신규 endpoint 추가

## 58. 공개 신뢰 요약(Public Trust Summary) / 프로필·매물·채팅 surface 계약
### 58.1 목적
- 내부 trust/risk/restriction/protection signal이 많아져도, 사용자 화면에서는 예측 가능하고 과도하게 낙인찍지 않는 형태의 `공개 신뢰 요약`으로 번역되어야 한다.
- 프로필, 매물 카드, 매물 상세, 채팅 상대 요약, 내 거래가 서로 다른 기준으로 상대 신뢰를 표현하지 않도록 공통 vocabulary와 projection 기준을 정의한다.
- 공개 신뢰는 거래 판단 보조 장치이며, 안전 보증/사기 방지 확정 신호처럼 오해되지 않도록 설명 레이어를 함께 설계한다.

### 58.2 설계 원칙
1. **근거 기반 요약**: 단일 점수보다 완료 이력, 후기, 응답성, 검증 상태, 최근 제재 여부를 구조화된 evidence pill로 보여준다.
2. **낙인 최소화**: 내부 risk score, fraud suspicion raw signal, visibility penalty 원인은 그대로 노출하지 않고 사용자에게 필요한 수준의 warning/banner만 공개한다.
3. **surface 일관성**: 같은 사용자는 어느 화면에서 보더라도 같은 `publicTrustTier`와 핵심 evidence 집합을 공유해야 한다.
4. **맥락 차등 공개**: 목록 카드에서는 요약만, 상세/채팅/거래상세에서는 더 풍부한 근거와 주의 문구를 제공한다.
5. **시간 민감성 분리**: 완료 거래 수처럼 비교적 안정적인 신호와, 최근 활동/응답성/보호조치처럼 빠르게 바뀌는 신호를 분리해 projection freshness 정책을 다르게 둔다.

### 58.3 핵심 vocabulary
#### 58.3.1 `publicTrustTier`
- `new`: 공개 신뢰 근거가 아직 적은 신규/저활동 사용자
- `building`: 일부 완료 이력과 기본 신뢰 신호가 쌓인 사용자
- `established`: 반복 거래와 안정적 후기/응답성을 가진 사용자
- `high_confidence`: 장기간 안정적 거래 이력과 양호한 신호를 가진 사용자
- `limited`: 공개 신뢰 노출을 축소하거나 주의가 필요한 사용자

원칙:
- `limited`는 제재 낙인 배지가 아니라, 신뢰 요약을 정상 노출할 수 없는 상황을 포함하는 보호적 상태다.
- 공개 tier는 순위 경쟁 점수처럼 다루지 않고, badge/summary copy 용도로만 사용한다.

#### 58.3.2 `trustEvidencePill`
- `completed_trade_count`
- `positive_review_ratio`
- `recent_activity`
- `response_reliability`
- `repeat_trade_signal`
- `verification_status`
- `recent_completion_velocity`
- `safe_trade_guidance`

각 pill은 label + shortValue + optional tooltip/copy를 가진다.

#### 58.3.3 `trustDisclosureMode`
- `minimal`: 목록 카드/좁은 UI에서 핵심 1~2개 신호만 노출
- `standard`: 매물 상세/채팅 상단 등 일반 상세 surface
- `expanded`: 프로필/거래상세에서 근거와 설명을 더 풍부하게 노출
- `suppressed`: 공개 노출 최소화, 내부적으로만 full signal 유지

#### 58.3.4 `trustWarningBannerType`
- `none`
- `new_user_caution`
- `limited_history_notice`
- `recent_restriction_notice`
- `ongoing_protection_review`
- `safety_best_practice`

원칙:
- 배너는 내부 위반명 그대로 노출하지 않는다.
- `recent_restriction_notice`도 구체적 사유보다 “일부 기능/노출이 제한될 수 있음” 수준 설명을 우선한다.

### 58.4 공개 신뢰 요약 canonical 객체
```json
{
  "userId": "user_123",
  "publicTrustSummary": {
    "publicTrustTier": "building",
    "trustDisclosureMode": "standard",
    "headlineLabel": "거래 이력이 쌓이는 중",
    "evidencePills": [
      {
        "type": "completed_trade_count",
        "label": "완료 거래",
        "shortValue": "12회",
        "priority": 1
      },
      {
        "type": "response_reliability",
        "label": "응답성",
        "shortValue": "빠른 편",
        "priority": 2
      },
      {
        "type": "verification_status",
        "label": "확인 상태",
        "shortValue": "기본 확인",
        "priority": 3
      }
    ],
    "warningBanner": {
      "type": "none",
      "message": null
    },
    "asOf": "2026-03-14T10:37:00+09:00"
  }
}
```

### 58.5 공개 산정 입력 신호
공개 신뢰 요약은 아래 원천을 사용하되, 내부 원천과 공개 표현을 분리한다.

| 입력 신호 | 원천 | 공개 반영 방식 | 비고 |
|---|---|---|---|
| 완료 거래 수 | completion aggregate | pill + tier 산정 | 절대 수치 또는 구간값 |
| 후기 추천 비율 | review aggregate | pill 또는 profile 상세 | 표본이 적으면 과도한 강조 금지 |
| 최근 활동성 | activity snapshot | `recent_activity` pill | 정확한 시각 대신 상대적 표현 우선 |
| 응답성 | response metric | `response_reliability` pill | `빠름/보통/느림` 같은 구간화 |
| 검증 상태 | verification snapshot | `verification_status` pill | 세부 인증 수단은 필요 최소 공개 |
| 반복 거래 신호 | repeat trade aggregate | `repeat_trade_signal` pill | 단골/재거래 경험 신호 |
| 제재/보호조치 | restriction/protection snapshot | tier 조정 또는 warning banner | 내부 사유 직접 노출 금지 |
| fraud suspicion/internal risk | anti-abuse internal model | 직접 미노출 | 배너/행동 제한으로만 간접 반영 |

### 58.6 tier 산정 기본 가이드
구체 점수식은 별도 스펙에서 확정하되, PRD 수준에서는 아래 가이드를 고정한다.

| tier | 기본 조건 가이드 | 공개 카피 방향 |
|---|---|---|
| `new` | 완료 거래 거의 없음, 활동 이력 부족 | `첫 거래 단계` |
| `building` | 일부 완료 거래/후기/응답성 신호 존재 | `거래 이력이 쌓이는 중` |
| `established` | 반복 완료 거래 + 안정적 응답/후기 | `안정적으로 거래 중` |
| `high_confidence` | 충분한 거래 이력 + 장기간 안정 신호 + 보호조치 이슈 적음 | `꾸준히 거래한 사용자` |
| `limited` | 공개 신뢰 표시를 정상적으로 노출하기 어려움 또는 보호 문구 필요 | `일부 정보 확인 필요` |

금지 원칙:
- `사기 위험`, `문제 계정`, `주의 인물` 같은 단정/낙인 카피는 사용하지 않는다.
- 공개 tier는 내부 제재 레벨과 1:1 매핑하지 않는다.

### 58.7 surface별 노출 계약
#### 58.7.1 목록 카드(`minimal`)
- 노출 요소:
  - tier badge 1개
  - evidence pill 최대 2개
  - warning banner는 미노출 또는 매우 짧은 icon hint만 허용
- 우선 노출 순서:
  1. 완료 거래 수 또는 구간
  2. 응답성 또는 최근 활동
  3. 검증 상태
- 공간 부족 시 숫자/구간형 신호 우선

#### 58.7.2 매물 상세(`standard`)
- 노출 요소:
  - tier badge
  - evidence pill 3~4개
  - 필요 시 safety/warning banner 1개
  - `프로필 보기` CTA
- warning banner 예시:
  - `첫 거래 상대라면 플랫폼 안에서 대화를 이어가세요`
  - `최근 활동 이력이 적어 응답이 늦을 수 있어요`
  - `일부 신뢰 정보가 제한되어 있어 거래 전 조건을 다시 확인하세요`

#### 58.7.3 채팅 상대 요약(`standard`)
- 노출 요소:
  - tier badge
  - `recent_activity` 또는 `response_reliability`
  - 거래 당일에는 `safety_best_practice` 계열 배너 우선
- 채팅 화면에서는 후기 비율보다 **응답/약속/실행 관련 신호**를 우선한다.
- active protection/review가 있으면 상세 reason 대신 caution banner만 노출한다.

#### 58.7.4 프로필(`expanded`)
- 노출 요소:
  - tier badge + headline copy
  - evidence pill 4~6개
  - 최근 후기 요약
  - 완료 거래/재거래/응답성/검증 신호의 상세 설명
  - warning banner/history notice
- profile에서는 `limited` 상태라도 최소한의 설명 카피가 있어야 하며, blank/silent suppression은 지양한다.

#### 58.7.5 내 거래 / 거래상세(`expanded`)
- 상대 신뢰 자체보다 `이번 거래 실행 맥락`에 유리한 신호를 우선한다.
- 우선 노출 신호:
  - 최근 응답성
  - 최근 거래 실행 이력 요약(가능한 경우)
  - 보호 문구 / 안전 가이드
  - repeat trade 여부

### 58.8 warning banner 노출 규칙
| banner type | 노출 조건 가이드 | 사용자 문구 방향 |
|---|---|---|
| `new_user_caution` | 신규/이력 부족 + 첫 거래 가능성 높음 | `처음 거래하는 상대예요. 조건을 한 번 더 확인하세요.` |
| `limited_history_notice` | 활동/완료/후기 근거 부족 | `참고할 거래 이력이 아직 많지 않아요.` |
| `recent_restriction_notice` | 최근 제한 해제/관찰기간 등 공개 보호 필요 | `일부 정보는 제한적으로 표시돼요. 거래 전 기록을 서비스 안에서 남겨두세요.` |
| `ongoing_protection_review` | protection review/provisional state 영향 | `안전 검토 중인 정보가 있어 거래 전 조건을 다시 확인하세요.` |
| `safety_best_practice` | 고가/첫거래/오프플랫폼 유도 위험군 등 | `외부 연락처나 선입금 유도는 주의하세요.` |

### 58.9 edge case / fallback 규칙
- 후기 수가 너무 적으면 비율보다 `후기 있음/없음` 수준으로만 축약한다.
- 완료 거래 수가 0이어도 verification/status/activity가 있으면 `new` tier와 함께 일부 evidence는 노출 가능하다.
- 탈퇴/비식별화 사용자와의 과거 거래 기록에서는 public tier를 실시간 재계산하지 않고 archived snapshot을 우선 사용한다.
- 제한/보호 상태로 인해 disclosure가 `suppressed`인 경우:
  - tier는 `limited`
  - evidence pill 일부 또는 전체 감춤
  - warning banner 또는 explanation copy 필수
- 운영 오류/집계 지연 시에는 stale 수치보다 conservative fallback(`정보 준비 중`)을 우선한다.

### 58.10 API / projection 파생 기준
후속 API/DB 문서는 아래 projection을 공통 재사용해야 한다.
- `UserPublicTrustProjection`
- 주요 필드 후보:
  - `userId`
  - `publicTrustTier`
  - `trustDisclosureMode`
  - `headlineLabel`
  - `headlineReasonCode`
  - `evidencePillsJson`
  - `warningBannerType`
  - `warningBannerMessage`
  - `asOf`
  - `snapshotVersion`

응답 surface 후보:
- `GET /listings`
- `GET /listings/{listingId}`
- `GET /chats/{chatRoomId}`
- `GET /me/trades`
- `GET /users/{userId}/profile`

원칙:
- listing author, chat counterparty, trade counterparty 모두 동일 projection source를 사용한다.
- 화면별 차이는 source 데이터가 아니라 `trustDisclosureMode`와 field trimming으로 해결한다.

### 58.11 analytics / KPI 파생 기준
이 섹션은 다음 분석 이벤트의 기준이 된다.
- `trust_summary_impression`
- `trust_summary_expand`
- `trust_warning_banner_impression`
- `trust_warning_banner_click`
- `profile_trust_detail_view`
- `trust_signal_to_chat_start_conversion`

핵심 관찰 포인트:
- 공개 신뢰 요약이 채팅 시작/예약 확정/완료 전환에 미치는 영향
- `limited`/warning banner 노출이 거래 이탈을 과도하게 만드는지 여부
- 신규 사용자에 대한 `new_user_caution`이 안전성 향상과 전환 저하 사이에서 어떤 균형을 보이는지

### 58.12 오픈 질문
- 공개 tier를 완전 정적 4~5단계 배지로 둘지, 일부 surface에서는 badge 없이 headline copy만 쓸지 확정 필요
- 후기 비율을 숫자로 공개할지, `긍정 후기 다수` 같은 구간형 라벨로 축약할지 결정 필요
- `limited` 상태에서 어떤 수준까지 이유 설명을 보여줄지, 이의제기/보호조치 UX와 연결 방식 확정 필요
- 고가 거래, 첫 거래, 오프라인 거래에서 trust warning banner를 얼마나 적극적으로 노출할지 실험 설계 필요
- 내부 reason code 세분화 후 publicReasonCode 유지

### 56.9 클라이언트 호환성 원칙
- 앱은 알 수 없는 enum/배지/알림 타입에 대해 안전한 fallback UI를 가져야 한다.
- `availableActions` 기반 UI를 사용하면 상태 분기 변경 시 앱 hotfix 필요성을 줄일 수 있다.
- 강제 업데이트가 필요한 비호환 정책 변경은 최소화하고, 가능하면 서버 주도 힌트/배너로 완충한다.

### 56.10 후속 API 명세 문서 파생 포인트
- endpoint별 auth scope
- request/response schema
- idempotency 적용 여부
- rate limit
- 상태 충돌 예시
- publicReasonCode / availableActions 포함 규칙
- 분석 이벤트 트리거 위치


## 58. 공개 프로필 범위 / 인덱싱 정책(초안)
### 57.1 목표
- 프로필은 거래 상대를 판단하는 최소 신뢰 요약을 제공하되, 검색엔진에 개인 활동 이력이 과도하게 축적되지 않게 해야 한다.
- 매물 SEO와 달리 프로필은 재식별/낙인 리스크가 커서 보수적으로 공개한다.

### 57.2 기본 정책안
- **MVP 기본안: 공개 프로필 URL은 로그인 사용자에게만 노출, 검색엔진 인덱싱은 기본 `noindex`**
- 비회원에게는 상세 프로필 페이지 대신 매물 카드/상세 내 요약 정보만 제한 제공
- 공개 프로필 전면 SEO 인덱싱은 Post-MVP 검토 항목으로 둔다.

### 57.3 프로필 공개 정보 레벨
| 정보 항목 | 비회원 | 로그인 회원 | 거래 상대/후기 작성 가능 관계 | 검색엔진 |
|---|---|---|---|---|
| 닉네임 | 마스킹 또는 축약 | 공개 | 공개 | 비공개 또는 noindex |
| 아바타 | 선택적 축소 노출 | 공개 | 공개 | 비공개 권장 |
| 완료 거래 수 | 구간형 요약 가능 | 공개 | 공개 | 인덱싱 제외 |
| 후기 요약 | 최근 1~2개 요약만 가능 | 공개 | 공개 | 인덱싱 제외 |
| 최근 활동 | 상대시간/구간형 | 상대시간 | 상대시간 또는 더 상세 | 인덱싱 제외 |
| 주 활동 서버 | 선택 공개 | 공개 | 공개 | 인덱싱 제외 |
| 신고/제재 세부 | 비공개 | 비공개 | 비공개 | 비공개 |
| 차단/뮤트 여부 | 본인에게만 | 본인에게만 | 본인에게만 | 비공개 |

### 57.4 프로필 URL/수명주기 원칙
- 사용자별 영구 숫자 증가 ID 노출보다 난수형/비추정 가능 식별자를 우선 검토한다.
- 탈퇴/영구정지 사용자의 공개 프로필 URL은 소프트 랜딩 또는 비공개 전환이 가능해야 한다.
- 동일 사용자의 과거 닉네임 이력이 공개 URL에 남지 않도록 canonical/slug 정책이 필요하다.

### 57.5 화면/데이터 시사점
- 프로필 API는 `publicProfile`, `memberProfile`, `participantProfile` 수준의 응답 레벨을 지원할 수 있다.
- 후기/완료 수는 원시 숫자 대신 구간형 표현(예: 1-4건, 5-19건)도 검토 가능하다.
- 운영/법적 요청이 있는 경우 공개 프로필 전체 비활성화가 가능해야 한다.

## 59. 후기 보복 리스크 완화 정책(초안)
### 58.1 목표
- 거래 품질을 보여주는 후기를 유지하면서도, 맞보복/협박/선제 압박 때문에 리뷰 품질이 무너지는 것을 줄인다.

### 58.2 완화 원칙
1. 상대 후기 내용은 공개 시점 전까지 블라인드 처리한다.
2. 거래 완료 직후 채팅 화면에서 “후기를 쓰지 않으면 불이익”처럼 느껴지는 압박 문구를 금지한다.
3. 후기 요청 메시지는 중립적 문구만 사용하고, 상대의 추천/비추천 여부를 암시하지 않는다.
4. 분쟁 또는 신고 접수 상태에서는 후기 공개/집계를 보수적으로 처리한다.

### 58.3 보복성 후기 징후 후보
- 상대가 비추천을 남긴 직후 유사한 표현으로 즉시 맞비추천 작성
- 거래 사실과 무관한 인신공격/협박/외부 유도 문구 포함
- 채팅에서 후기 강요/후기 거래("좋게 써주면 나도 좋게") 정황 존재
- 동일 사용자에게 반복적으로 패턴화된 단문 비추천 작성

### 58.4 운영 처리 원칙
- 보복성 의심 후기는 자동 삭제 대신 `pending_moderation` 또는 `limited_visibility` 후보로 전환 가능해야 한다.
- 숨김 여부와 별개로 내부 집계 반영은 즉시가 아니라 운영 판정 후 반영하는 방안을 우선 검토한다.
- 운영자는 후기 원문, 직전 채팅 맥락, 완료/분쟁 상태, 상대방 신고 여부를 함께 봐야 한다.

### 58.5 사용자 UX 원칙
- 후기 작성 화면에 아래 가이드 문구를 노출한다.
  - 거래 사실 중심으로 작성해 주세요.
  - 욕설, 개인정보, 보복성 표현은 비노출될 수 있습니다.
  - 상대 후기 내용은 공개 전까지 보이지 않습니다.
- 후기 신고 CTA는 리뷰 카드와 상세 화면 모두에서 접근 가능해야 한다.
- 후기 비노출 시 작성자에게는 축약된 정책 사유와 1회 이의제기 경로를 제공한다.

## 60. `reserved` 장기 유지 악용 탐지 / 운영 기준(초안)
### 59.1 목표
- 실제 우선 거래 보호를 위해 필요한 `reserved` 상태가 허위 희소성, 관심 끌기, 경쟁자 차단 수단으로 악용되지 않게 한다.

### 59.2 이상징후 룰 후보
- 동일 매물이 짧은 기간 내 `available -> reserved -> available`를 반복
- `reserved` 평균 유지 시간이 지나치게 길고 실제 `pending_trade` 또는 `completed`로 거의 이어지지 않음
- 동일 작성자가 우선 거래 상대를 반복 교체

- `reserved` 상태에서 대기 문의만 과도하게 모으고 실제 응답/후속 조치가 낮음
- 신규/저신뢰 계정이 고가 매물을 장시간 `reserved`로 유지

### 59.3 운영 임계치 기본안(가정)
- 최근 14일 내 동일 매물 또는 동일 작성자 기준 `reserved -> available` 복귀가 3회 이상이면 모니터링 후보
- `reserved` 연속 유지 48시간 초과 시 작성자 상태 점검 리마인드
- 리마인드 후 24시간 추가 무응답이면 랭킹 감점 또는 운영 검토 큐 적재 후보
- 자동 영구 제재는 금지하고, 우선은 노출 감점/상태 점검/등록 제한부터 적용

### 59.4 사용자/운영 노출 원칙
- 사용자에게는 "장시간 예약중인 매물" 정도의 중립 표시만 제공하고 내부 탐지 점수는 비공개로 유지한다.
- 운영 백오피스에서는 `reserved` 유지 시간, 우선 상대 변경 횟수, 대기 문의 수, 실제 완료 전환율을 함께 보여야 한다.
- 반복 악용이 확인되면 `listing_hidden_policy_violation` 또는 별도 `listing_reserved_abuse_detected` 내부 코드 후보를 검토한다.

## 61. 분쟁 중 추가 소명 UX / 운영 상호작용 정책(초안)
### 60.1 목표
- 완료 분쟁/노쇼 분쟁이 발생했을 때, 사용자가 채팅 본문을 훼손하지 않고 필요한 소명만 추가 제출할 수 있어야 한다.
- 운영자는 사건 타임라인과 추가 소명을 한 곳에서 재구성 가능해야 한다.

### 60.2 기본 UX 원칙
- 분쟁 상태 진입 시 일반 채팅 입력창은 기본적으로 비활성화하고, 별도의 `추가 소명 제출` 액션을 노출한다.
- 소명은 일반 대화와 분리된 사건 스레드 또는 폼으로 저장한다.
- 소명 제출은 양측 각각 제한 횟수(예: 3회, 가정)와 길이 제한을 두어 감정적 반복 공방을 줄인다.

### 60.3 소명 데이터 항목 후보
- `disputeId`
- `submittedByUserId`
- `disputeType`: `completion_dispute` | `no_show_dispute` | `abuse_dispute`
- `statementText`
- `attachments(optional)`
- `submittedAt`
- `visibleToCounterparty`: 기본 false 또는 요약만 공개 정책 검토
- `visibleToModerator`: true

### 60.4 운영 처리 원칙
- 운영자는 채팅 원본, 예약 이력, 완료 요청 시점, 소명 제출 순서를 타임라인으로 같이 봐야 한다.
- 소명 제출 후 상대방에게 자동 원문 공유를 할지 여부는 보수적으로 접근하고, MVP에서는 운영자 전용 제출을 기본안으로 둔다.
- 운영자가 추가 자료를 요청할 때는 기한과 요청 사유를 명시해야 한다.
- 제출 기한 경과 또는 반복 무관 소명은 더 이상 접수하지 않을 수 있다.

### 60.5 화면/API 파생 포인트
- 사용자 화면: `분쟁 진행 상태`, `내 소명 제출`, `운영 추가 요청`, `다음 가능 액션`
- 운영 화면: `사건 타임라인`, `양측 소명 목록`, `추가 자료 요청`, `판정 기록`
- API 후보:
  - `POST /trade-completions/{completionId}/dispute-statements`
  - `GET /trade-completions/{completionId}/dispute`
  - `POST /admin/disputes/{disputeId}/request-more-info`

## 62. 거래 장소/미팅 포인트 모델 및 안전 정책(초안)
### 62.1 목표
- 리니지 클래식 거래의 실제 성사 지점인 `언제`, `어디서`, `어떤 식으로 만날지`를 게시판 텍스트가 아니라 구조화된 예약 데이터로 다룬다.
- 인게임 접선과 오프라인 접선(예: PC방 인근)을 하나의 예약 UX 안에서 지원하되, 개인정보/정확한 위치 노출은 단계적으로 제한한다.
- 화면/DB/API/운영정책이 동일한 미팅 포인트 개념을 사용하도록 공통 모델을 정의한다.

### 62.2 미팅 포인트 유형 정의
| meetingType | 의미 | 대표 예시 | MVP 지원 수준 |
|---|---|---|---|
| `in_game` | 게임 내 서버/마을/좌표성 텍스트 기반 접선 | 기란 마을 창고 앞 | 필수 |
| `offline_pc_bang` | 오프라인 PC방/PC방 인근 접선 | 강남역 OO PC방 앞 | 필수 |
| `offline_public_place` | PC방 외 공개장소 접선 | 지하철역 출구 앞 | 확장 후보 |
| `either` | 상대와 협의 후 인게임/오프라인 중 확정 | 추후 결정 | 필수 |

기본 원칙:
- MVP에서는 `in_game`, `offline_pc_bang`, `either`를 우선 지원하고, `offline_public_place`는 내부 모델만 열어두고 노출 여부를 추후 결정한다.
- `tradeMethod`는 매물 수준의 선호이고, `meetingType`은 예약 수준의 실제 약속 방식이다.

### 62.3 미팅 포인트 구조화 필드 후보
`Reservation` 또는 별도 `MeetingPointSnapshot` 성격으로 아래 구조를 검토한다.

| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `meetingType` | 필수 | `in_game` / `offline_pc_bang` / `either` |
| `meetingRegionText` | 선택 | 대지역/동네/마을명 수준 요약 |
| `meetingPointText` | 필수 | 사용자가 보는 접선 설명 |
| `serverId` | 조건부 필수 | 인게임 접선 시 필요 |
| `zoneText` | 선택 | 인게임 세부 지역/사냥터/마을 |
| `pcBangName` | 선택 | 오프라인 접선 시 PC방 명칭 |
| `pcBangAddressText` | 선택 | 운영/사용자 확인용 주소 텍스트 |
| `placeVisibilityLevel` | 필수 | `coarse` / `confirmed_only` / `exact_confirmed` |
| `arrivalGuideText` | 선택 | "도착하면 1층 입구에서 채팅" 등 |
| `counterpartyCharacterName` | 선택 | 상대 식별용 캐릭터명 |
| `locationConfirmedAt` | 선택 | 정확 장소 확정 시각 |
| `locationLastEditedAt` | 선택 | 장소 마지막 수정 시각 |

설계 원칙:
- 예약 시점의 장소는 이후 매물 본문이 바뀌어도 보존되도록 snapshot 성격으로 저장한다.
- 오프라인 주소 원문은 검색 인덱스/SEO/공개 프로필/푸시 본문에 재사용하지 않는다.

### 62.4 단계별 장소 정보 공개 정책
| 단계 | 상대방에게 보이는 수준 | 예시 |
|---|---|---|
| 매물 공개 | 대략적 선호만 | `기란 근처`, `강남/서초권 PC방 선호` |
| 일반 채팅(open) | 협의용 요약 수준 | `기란 마을 가능`, `강남역 부근 가능` |
| 예약 제안(proposed) | 구체 후보 + 시간 포함 | `오늘 21시 강남역 XX PC방` |
| 예약 확정(confirmed) | 실제 실행 정보 | `21시 강남역 XX PC방 2층 카운터 앞` |
| 거래 종료 후 | 기록 조회용 제한 표시 | 필요 최소 정보만 유지 |

세부 원칙:
- exact 주소, 층수, 좌석번호, 실명 추정 정보 등은 예약 확정 전 과도하게 노출하지 않는다.
- 푸시/잠금화면에는 `강남역 근처`, `기란 마을` 등 coarse 수준만 사용한다.
- 거래 종료 후 기록 화면에서는 향후 악용을 줄이기 위해 exact 장소보다 요약 정보 위주로 보관/표시하는 안을 우선 검토한다.

### 62.5 예약 생성/수정 규칙
- `meetingType=in_game`이면 `serverId`와 `meetingPointText`가 필수다.
- `meetingType=offline_pc_bang`이면 `meetingPointText`는 필수이며, `pcBangName` 또는 `meetingRegionText` 중 최소 1개가 필요하다.
- `meetingType=either`는 제안 단계에서만 허용하고, 예약 확정 시에는 실제 방식(`in_game` 또는 `offline_pc_bang`)으로 구체화해야 한다.
- 예약 확정 후 장소 수정은 가능하되, 상대방에게 변경 diff를 명확히 보여주고 재확인 상태를 요구할 수 있어야 한다.
- 예약 시각 1시간 이내의 장소 대폭 변경은 운영 분석/노쇼 분쟁 참고 신호로 남기는 것이 바람직하다.

### 62.6 장소 변경/재확인 상태 모델
장소 변경은 단순 텍스트 수정이 아니라 거래 실패 리스크가 큰 이벤트이므로 별도 상태를 검토한다.

| 상태/필드 | 의미 |
|---|---|
| `locationChangePending=true` | 한쪽이 장소/시간 변경안을 제시함 |
| `lastConfirmedMeetingSnapshotId` | 마지막 상호 인지된 장소 스냅샷 |
| `requiresCounterpartyAck=true` | 상대 재확인 필요 |
| `locationMismatchRisk=true` | 거래 직전 큰 변경 발생 |

운영/UX 원칙:
- 장소 변경 후 상대방이 읽지 않은 채 예약 시각이 다가오면 리마인드 알림을 보낸다.
- 직전 변경이 누적되면 노쇼 분쟁 판단 시 참고 신호로 활용할 수 있어야 한다.

### 62.7 모바일 예약 UX 요구사항
- 예약 제안 시 자유 입력만 두지 않고 아래 빠른 입력을 제공한다.
  1. 최근 사용 장소
  2. 자주 쓰는 인게임 장소 템플릿(예: `기란 마을 창고 앞`)
  3. 자주 쓰는 시간 슬롯(오늘 저녁, 30분 후, 내일 오후 등)
  4. 거래 방식 quick chip(`인게임`, `PC방`, `협의 후 결정`)
- 상세 주소를 길게 입력하지 않아도 되는 구조를 우선 검토한다.
- 예약 카드에는 항상 `시간`, `방식`, `요약 위치`, `상대 확인 상태`가 한눈에 보여야 한다.
- 거래 직전에는 `도착했어요`, `5분 늦어요`, `장소 다시 보기` 같은 실행형 CTA를 제공하는 것이 바람직하다.

### 62.8 화면 파생 요구사항
#### 매물 등록/상세
- 매물에는 정확 장소가 아니라 선호권역/선호방식만 노출한다.
- `offline_pc_bang` 선호 시 `강남역/선릉역 인근 가능` 같은 coarse 텍스트 예시를 제공한다.

#### 채팅/예약 카드
- 예약 카드 필수 노출 요소:
  - 거래 방식
  - 예약 시각
  - 서버 또는 지역
  - 장소 확정 여부
  - 마지막 수정 시각
  - 변경/재확인 필요 배지
- 예약 변경 시 이전 정보와 현재 정보를 비교 가능하게 보여야 한다.

#### 내 거래 화면
- 임박 거래는 `오늘 21:00 · 기란 마을`처럼 한 줄 요약 표시
- exact 장소는 거래 상세 진입 후에만 노출하는 계층형 UX를 우선 검토한다.
- 오프라인 거래는 안전 문구(공공장소/플랫폼 내 기록 유지)를 함께 노출한다.

### 62.9 데이터 모델 시사점
- `Reservation`에 모든 장소 필드를 직접 넣을 수도 있으나, 변경 이력/스냅샷 보존을 위해 `ReservationMeetingSnapshot` 분리도 검토 가치가 있다.
- 후보 엔티티:
  - `ReservationMeetingSnapshot`
  - `SavedMeetingTemplate` (사용자 최근/즐겨찾기 장소 템플릿)
- `SavedMeetingTemplate`는 공개 데이터가 아니라 사용자 개인 설정으로 취급해야 한다.
- 인게임 장소 자동완성 도입 시 `GameMeetingSpotCatalog` 같은 보조 테이블을 둘 수 있다.

### 62.10 API 후보 확장
- `GET /me/meeting-templates`
- `POST /me/meeting-templates`
- `GET /catalog/meeting-spots?serverId=...`
- `POST /reservations/{reservationId}/ack-location-change`

응답 필드 후보:
```json
{
  "reservationId": "res_123",
  "meeting": {
    "meetingType": "offline_pc_bang",
    "summaryLabel": "강남역 인근 PC방",
    "detailLabel": "XX PC방 2층 카운터 앞",
    "placeVisibilityLevel": "exact_confirmed",
    "requiresCounterpartyAck": true,
    "lastUpdatedAt": "2026-03-13T19:20:00+09:00"
  }
}
```

### 62.11 운영/분쟁 시사점
- 노쇼/장소 혼선 분쟁에서는 아래를 함께 재구성할 수 있어야 한다.
  1. 마지막 확정 장소
  2. 직전 장소 변경 횟수
  3. 상대방 읽음/재확인 여부
  4. 예약 시각 직전 알림 발송 여부
- 오프라인 거래의 exact 장소는 운영자도 최소 권한/감사 로그 하에 열람해야 한다.
- 반복적으로 직전 장소를 바꾸거나 모호한 장소만 제시하는 사용자는 안전 리스크 신호로 사용할 수 있다.

### 62.12 분석 이벤트 후보
- `reservation_location_suggested`
- `reservation_location_confirmed`
- `reservation_location_changed`
- `reservation_location_acknowledged`
- `trade_arrival_message_sent`
- `trade_location_template_used`

핵심 분석 포인트:
- 인게임/오프라인 방식별 예약 확정률
- 장소 변경 발생률과 노쇼율 상관관계
- 최근 사용 장소 템플릿 사용 시 예약 생성 시간 단축 여부

## 63. 완료 내부 상태 enum / Dispute 객체 도입안(초안)
### 63.1 목표
- 사용자에게 보이는 단순한 거래 상태(`available`, `reserved`, `pending_trade`, `completed`, `cancelled`)를 유지하면서도, 서버/운영/백오피스는 완료 확인 대기와 분쟁 단계를 더 정밀하게 추적할 수 있어야 한다.
- 완료 처리, 자동확정, 이의 제기, 노쇼/완료 분쟁을 하나의 일관된 내부 상태 모델로 다뤄 DB/API/운영정책 충돌을 줄인다.

### 63.2 공개 상태와 내부 완료 단계 분리 원칙
- `Listing.status`는 사용자 탐색/검색/목록 노출을 위한 **공개 상태**로 유지한다.
- 완료 처리 이후의 세부 단계는 `TradeCompletion.completionStage` 또는 동등한 내부 필드로 관리한다.
- 검색/목록/SEO 관점에서는 완료 요청 시점부터 사실상 종결 후보로 취급하되, 프로필 집계/후기 개방/분쟁 운영은 내부 단계에 따라 달라질 수 있다.

### 63.3 내부 완료 단계 enum 후보
| 내부 필드 | 의미 | 사용자 노출 문구 예시 | 주요 허용 액션 |
|---|---|---|---|
| `none` | 완료 흐름 미시작 | 노출 없음 | 완료 요청 가능 |
| `requested` | 한쪽이 완료 요청함 | 거래완료 확인 대기 | 상대 확인, 이의제기 |
| `confirmed_by_counterparty` | 상대가 명시 확인 | 거래완료 | 후기 작성 |
| `auto_confirmed` | 상대 무응답으로 자동확정 | 거래완료 | 후기 작성, 신고 |
| `disputed` | 완료 주장에 이의 제기됨 | 분쟁 진행 중 | 소명 제출 |
| `resolved_completed` | 운영이 완료로 판정 | 거래완료 | 후기 작성 |
| `resolved_not_completed` | 운영이 미완료/불발로 판정 | 거래불발 / 취소 | 재오픈 후보 |
| `closed_without_review` | 신고 철회/기각 등으로 종료 | 처리 완료 | 기록 열람 |

원칙:
- `Listing.status=completed`는 외부 공개용 최종 상태로 유지하되, 내부적으로는 `completionStage=requested` 같은 대기 상태가 가능하다.
- UI에는 내부 enum 전체를 그대로 노출하지 않고, `거래완료 확인 대기`, `분쟁 진행 중`, `거래완료` 정도의 사용자 문구로 매핑한다.

### 63.4 상태 조합 규칙
| Listing.status | completionStage | 허용 여부 | 설명 |
|---|---|---|---|
| `pending_trade` | `none` | 허용 | 거래 직전/직후 기본 상태 |
| `completed` | `requested` | 허용 | 외부 공개상 종결, 내부 확인 대기 |
| `completed` | `confirmed_by_counterparty` | 허용 | 정상 완료 |
| `completed` | `auto_confirmed` | 허용 | 무응답 자동확정 |
| `completed` | `disputed` | 허용 | 공개 목록에서는 종결 유지, 운영은 분쟁 처리 |
| `available` | `resolved_not_completed` | 허용 | 운영 판단으로 거래 미완료 후 재오픈 |
| `cancelled` | `resolved_not_completed` | 허용 | 불발/취소 확정 |
| `completed` | `none` | 금지 | 완료 상태에는 completion 객체 필요 |

### 63.5 Dispute 객체 도입 필요성
다음 이유로 `Report`만으로는 완료/노쇼 분쟁을 충분히 구조화하기 어렵다.
1. 분쟁은 단순 신고보다 상태머신과 긴밀히 연결된다.
2. 분쟁에는 양측 소명, 운영 요청, 판정, 판정 시각, 판정 결과가 필요하다.
3. 완료 분쟁과 일반 욕설 신고는 같은 큐에 있더라도 데이터 구조가 다르다.

따라서 MVP 또는 직후 버전에서 별도 `Dispute` 객체 도입을 우선 검토한다.

### 63.6 Dispute 엔티티 초안
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `disputeId` | 필수 | 분쟁 식별자 |
| `disputeType` | 필수 | `completion_dispute` / `no_show_dispute` / `meeting_mismatch_dispute` |
| `listingId` | 필수 | 관련 매물 |
| `chatRoomId` | 필수 | 관련 채팅방 |
| `reservationId` | 선택 | 관련 예약 |
| `completionId` | 선택 | 관련 완료 기록 |
| `openedByUserId` | 필수 | 분쟁 시작자 |
| `counterpartyUserId` | 필수 | 상대방 |
| `disputeStatus` | 필수 | `open` / `waiting_statement` / `under_review` / `resolved` / `closed` |
| `resolutionType` | 선택 | `completed_upheld` / `not_completed_upheld` / `no_fault` / `policy_violation_detected` |
| `resolutionReasonCode` | 선택 | 판정 사유 코드 |
| `openedAt` | 필수 | 시작 시각 |
| `resolvedAt` | 선택 | 해결 시각 |
| `resolvedByAdminId` | 선택 | 처리 운영자 |

### 63.7 DisputeStatement 하위 객체 초안
- `disputeStatementId`
- `disputeId`
- `submittedByUserId`
- `statementType`: `initial_claim` / `counter_statement` / `additional_evidence` / `admin_request_response`
- `bodyText`
- `attachmentsJson(optional)`
- `submittedAt`
- `visibilityScope`: `staff_only` / `shared_summary` 후보

원칙:
- 일반 채팅 메시지와 소명 텍스트는 분리 저장한다.
- 운영자 요청에 의한 추가자료 제출과 자발적 소명을 구분할 수 있어야 한다.
- 분쟁 소명은 수정 대신 append-only 이력 저장을 우선안으로 한다.

### 63.8 Report와 Dispute의 관계
- 모든 Dispute가 Report를 필수로 가지는 것은 아니다.
  - 예: 완료 확인 화면에서 상대가 `이의 제기`를 누르면 곧바로 Dispute 생성 가능
- 다만 고위험 분쟁은 운영 큐 통합을 위해 `linkedReportId(optional)`를 가질 수 있다.
- `Report`는 안전/정책 위반 접수 중심, `Dispute`는 거래 결과/사실관계 판정 중심으로 역할을 나눈다.

## 64. 매물 품질 스냅샷(Listing Quality Snapshot) / 등록 완성도·검색 품질·운영 개입 canonical 계약
### 64.1 목적
- 같은 매물을 등록 화면에서는 `작성 미완성`, 검색에서는 `품질 낮음`, 운영에서는 `검토 필요`로 제각각 해석하지 않도록 하나의 canonical 품질 읽기 모델을 정의한다.
- 품질 평가는 단순 랭킹 점수가 아니라 **등록 완료 가능성**, **구매자 이해 가능성**, **실제 거래 연결 가능성**, **정책 안전성**을 함께 반영해야 한다.
- 홈/목록/상세/등록 UX/API/운영 큐/analytics가 동일한 vocabulary를 사용해 `왜 이 매물이 덜 노출되는지`, `무엇을 보완하면 되는지`, `언제 운영 개입이 필요한지`를 일관되게 설명할 수 있어야 한다.

### 64.2 핵심 개념
| 용어 | 의미 |
|---|---|
| `listingQualitySnapshot` | 특정 시점의 매물 품질 상태를 요약한 canonical read model |
| `qualityTier` | 검색/홈/운영에서 공통으로 쓰는 품질 단계 (`excellent` / `good` / `needs_improvement` / `restricted`) |
| `qualityDimensionType` | 품질 평가 축 (`completeness` / `clarity` / `freshness` / `trust_readiness` / `policy_safety`) |
| `qualityFindingCode` | 구체 보완/경고 사유 코드 |
| `improvementActionType` | 사용자가 취할 수 있는 보완 액션 (`add_image`, `clarify_title`, `update_price`, `confirm_availability`, `fix_policy_issue`) |
| `rankingQualityImpact` | 검색/추천 랭킹에 미치는 영향 수준 |
| `qualityReviewState` | 운영/자동화의 품질 검토 상태 (`none` / `soft_warned` / `limited_visibility` / `review_required`) |

원칙:
- 품질은 `정책 위반 여부`와 같지 않다. 정책 위반이 아니어도 거래 성사 가능성이 낮으면 품질 개선 대상으로 분류할 수 있다.
- 반대로 품질이 좋아 보여도 정책 위반이면 `restricted` tier가 우선한다.
- 최종 사용자 노출은 점수보다 `보완 필요`, `거래 가능 정보 충분`, `검토 필요`처럼 행동 가능한 문구로 표현한다.

### 64.3 품질 평가 축 정의
#### 64.3.1 completeness
매물 작성에 필요한 기본 정보가 충분한가를 본다.
- 제목/아이템명/서버/거래유형/가격정책/수량/설명 최소 길이 충족 여부
- 이미지 존재 여부
- 거래 가능 시간/희망 장소/구조화 속성 입력 여부
- 카테고리별 필수 속성 누락 여부

## 65. 거래 단위(Unit of Trade) / 묶음·개당가·총액 canonical 계약
### 65.1 목적
- 같은 매물을 어떤 화면에서는 `수량 1`, 어떤 화면에서는 `묶음 1세트`, 어떤 채팅에서는 `개당 50만`, 어떤 완료 기록에서는 `총 200만`으로 다르게 해석하는 문제를 방지한다.
- `부분거래`, `가격 제안`, `잔여 수량`, `총액 계산`, `후기/분쟁`이 모두 동일한 거래 단위 기준을 바라보도록 canonical 계약을 정의한다.
- 매물 등록/검색/상세/채팅/완료/API/DB/운영정책이 `무엇을 몇 단위로 거래했는가`를 일관되게 해석하도록 한다.

### 65.2 핵심 개념
| 용어 | 의미 |
|---|---|
| `tradeUnitSnapshot` | 특정 매물 또는 실제 성사 건이 어떤 단위로 거래되는지 요약한 canonical read model |
| `unitBasis` | 거래 단위 기준 (`single_item` / `stack_count` / `bundle` / `lot`) |
| `priceBasis` | 가격 기준 (`total_price` / `per_unit_price` / `per_bundle_price` / `offer_based`) |
| `bundleRule` | 묶음 분할 가능 여부 (`fixed_bundle_only` / `divisible_bundle` / `seller_decides`) |
| `quantityGranularity` | 수량 증감 최소 단위 (`1`, `10`, `0.1` 등) |
| `residualUnitPolicy` | 부분거래 후 잔여 수량 처리 방식 (`auto_reduce`, `seller_reconfirm`, `close_on_any_completion`) |
| `executionQuantitySnapshot` | 실제 거래 직전 확정된 수량/단가/총액 스냅샷 |
| `priceComputationMode` | 총액 계산 방식 (`unit_times_quantity`, `bundle_count_times_bundle_price`, `manual_total`) |

원칙:
- `quantity` 하나만으로는 거래 의미를 충분히 표현하지 못한다. 반드시 `unitBasis`, `priceBasis`, `quantityGranularity`와 함께 해석해야 한다.
- 같은 품목이라도 매물마다 거래 단위가 다를 수 있다. 예: `화살 1,000개 묶음`, `젤 1장 단위`, `재료 100개 단위 일괄 판매`.
- 검색/상세/채팅/완료는 동일한 `tradeUnitSnapshot`을 참조하되, 화면마다 보여주는 요약 수준만 달라진다.

### 65.3 단위 기준 유형 정의
| `unitBasis` | 의미 | 예시 | MVP 기본 지원 |
|---|---|---|---|
| `single_item` | 비중첩 또는 개별 아이템 1개 단위 | 무기 1개, 방어구 1개 | 필수 |
| `stack_count` | 수량형 재화/소모품을 개수로 거래 | 아데나 1,000,000 / 재료 300개 | 필수 |
| `bundle` | 여러 개를 하나의 묶음/세트로 판매 | 주문서 10장 세트 1묶음 | 필수 |
| `lot` | 개별 단위보다 큰 로트/일괄 물량 거래 | 잡템 일괄, 창고 정리 일괄 | 확장 후보지만 모델은 열어둠 |

기본 원칙:
- `single_item`은 기본적으로 `quantityGranularity=1`이다.
- `stack_count`는 `quantity`와 `priceAmount`를 직접 연결해 총액을 계산할 수 있어야 한다.
- `bundle`은 `bundleSize` 또는 `bundleDescription`이 필요하며, 사용자가 `1묶음 = 무엇인지` 이해할 수 있어야 한다.
- `lot`은 세부 구성보다 일괄 거래 자체가 중요하므로, 품질 정책과 운영 검토 기준이 더 엄격해야 한다.

### 65.4 가격 기준 정의
| `priceBasis` | 의미 | 표시 예시 | 계산 규칙 |
|---|---|---|---|
| `total_price` | 전체 매물 총액 기준 | `총 500만 아데나` | `priceAmount`가 최종 총액 |
| `per_unit_price` | 1개/1카운트당 가격 | `개당 50만` | `quantity × unitPrice` |
| `per_bundle_price` | 1묶음당 가격 | `1세트 200만` | `bundleCount × bundlePrice` |
| `offer_based` | 제안 기반, 고정액 없음 | `가격 제안 받음` | 총액은 offer/deal terms에서 확정 |

원칙:
- 검색 목록에는 사용자가 오해하지 않도록 `개당/묶음당/총액` 라벨을 함께 표시해야 한다.
- `offer_based`라도 수량 기준은 명확해야 한다. 즉 `얼마에 팔지`는 미정이어도 `무엇을 얼마나`는 구조화되어야 한다.
- `priceBasis`가 `per_unit_price`인 매물은 부분거래와 자연스럽게 연결되지만, `total_price`인 매물은 부분거래 허용 여부를 별도 명시해야 한다.

### 65.5 매물 등록 규칙
등록 UX는 최소 아래 필드를 구조화해야 한다.

| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `unitBasis` | 필수 | 거래 단위 유형 |
| `quantity` | 필수 | 판매/구매 의도 수량 |
| `quantityLabel` | 선택 | `개`, `장`, `세트`, `묶음`, `만 아데나` 등 표시용 단위명 |
| `quantityGranularity` | 필수 | 수량 증감 최소 단위 |
| `bundleSize` | 조건부 | `bundle`일 때 1묶음 구성 수량 |
| `bundleDescription` | 조건부 | 묶음 구성 설명 |
| `priceBasis` | 필수 | 총액/개당/묶음당/제안 |
| `priceAmount` | 조건부 | `offer_based` 제외 기본 필수 |
| `minTradeQuantity` | 선택 | 허용 최소 거래 수량 |
| `maxTradeQuantity` | 선택 | 허용 최대 거래 수량 |
| `bundleRule` | 필수 | 묶음 분할 가능 여부 |
| `residualUnitPolicy` | 필수 | 부분거래 후 잔여 수량 처리 정책 |

기본 입력 규칙:
- `single_item`이면 `quantity=1`을 기본값으로 제안한다.
- `bundle`이면 `bundleSize >= 2` 또는 사람이 이해 가능한 `bundleDescription`이 필요하다.
- `priceBasis=per_unit_price`이면 `quantityGranularity`와 단가 해석이 충돌하지 않아야 한다.
- `offer_based`에서는 `희망 가격대`를 선택 입력으로 둘 수 있으나 canonical price는 null 허용으로 본다.

### 65.6 검색/목록 표시 규칙
목록 카드와 검색 정렬은 아래 기준을 따라야 한다.
- `총액만 있는 매물`과 `개당가 매물`을 같은 숫자값으로 단순 비교 정렬하지 않는다.
- 검색 결과의 가격 라벨은 최소한 아래 형식을 지원한다.
  - `총 500만`
  - `개당 50만 · 수량 10`
  - `1묶음 200만 · 3묶음`
  - `가격 제안 받음`
- 가격 필터는 `priceBasis`가 다른 매물끼리 혼동을 줄이기 위해 `총액 기준`과 `단위가 기준`을 내부적으로 구분 계산해야 한다.
- Post-MVP에서는 `정규화 표시 가격(normalizedComparablePrice)`를 별도 read model로 둘 수 있으나, MVP에서는 비교 가능한 경우에만 정렬 정확도를 보장하는 보수적 정책을 택한다.

### 65.7 부분거래와 잔여 수량 규칙
`부분거래 / 잔여 수량 / 수량 할당 계약`과 연결해 아래를 명시한다.
- `partial trade`는 `unitBasis`, `priceBasis`, `bundleRule`이 허용할 때만 가능하다.
- `bundleRule=fixed_bundle_only`이면 묶음을 쪼개는 부분거래는 금지다.
- `priceBasis=total_price`이고 `bundleRule`이 분할 허용이 아니면 부분거래 완료 후 자동 잔여 계산을 하지 않는다.
- `residualUnitPolicy=auto_reduce`이면 실제 완료 수량만큼 `remainingQuantity`를 자동 차감한다.
- `residualUnitPolicy=seller_reconfirm`이면 완료 후 작성자가 잔여 수량/가격을 재확인해야 다시 `available`로 노출될 수 있다.
- `residualUnitPolicy=close_on_any_completion`이면 일부만 거래되더라도 원매물은 종결하고, 필요 시 새 매물로 재등록한다.

### 65.8 채팅/예약/딜 확정 규칙
실제 협상 단계에서는 `executionQuantitySnapshot`이 필요하다.
- 매물 원본 수량이 `100개`라도 채팅에서 `20개만 먼저` 합의할 수 있다.
- 이때 deal terms는 최소 아래를 고정해야 한다.
  - `agreedUnitBasis`
  - `agreedQuantity`
  - `agreedPriceBasis`
  - `agreedUnitPrice(optional)`
  - `agreedTotalPrice`
  - `residualHandlingDecision`
- `offer_based` 매물은 예약 확정 전이라도 거래 실행 직전에는 반드시 `agreedTotalPrice`가 존재해야 한다.
- 예약 카드/완료 카드/분쟁 타임라인은 `매물 원본 수량`이 아니라 `이번 거래에서 실제 합의된 수량`을 우선 표시해야 한다.

### 65.9 완료/분쟁/후기 해석 규칙
- 후기와 완료 이력은 `매물 전체`보다 `실행된 거래 단위`를 기준으로 남겨야 한다.
- 한 매물에서 여러 번 부분거래가 일어날 수 있으므로, 후기 eligibility는 `executionQuantitySnapshot` 단위로 판정하는 것이 바람직하다.
- 분쟁 시 운영자는 아래를 함께 재구성해야 한다.
  1. 원매물 수량/가격 기준
  2. 실제 합의 수량/총액
  3. 부분거래 여부
  4. 잔여 수량 처리 정책
  5. 상대가 주장하는 불일치 유형 (`quantity_shortfall`, `bundle_mismatch`, `price_mismatch` 등)
- `bundle` 매물에서는 `묶음 구성과 실제 인도 수량이 다름`을 별도 mismatch 사유로 다뤄야 한다.

### 65.10 API/DB 파생 기준
#### API 후보 필드
```json
{
  "tradeUnit": {
    "unitBasis": "bundle",
    "quantity": 3,
    "quantityLabel": "세트",
    "bundleSize": 10,
    "bundleDescription": "주문서 10장 1세트",
    "quantityGranularity": 1,
    "priceBasis": "per_bundle_price",
    "priceAmount": 2000000,
    "bundleRule": "fixed_bundle_only",
    "residualUnitPolicy": "seller_reconfirm"
  }
}
```

#### DB 컬럼 후보
- `unit_basis`
- `quantity_label`
- `quantity_granularity`
- `bundle_size`
- `bundle_description`
- `price_basis`
- `min_trade_quantity`
- `max_trade_quantity`
- `bundle_rule`
- `residual_unit_policy`
- `normalized_total_price(optional)`
- `normalized_unit_price(optional)`

원칙:
- 숫자 컬럼만으로 의미를 잃지 않도록 표시용 단위 라벨과 계산용 basis를 분리 저장한다.
- 가격 정렬/필터를 위해 `normalized_total_price`, `normalized_unit_price` 같은 파생 컬럼을 둘 수 있으나, 원본 입력 의미를 훼손해서는 안 된다.
- `dealTerms`, `exchangeConfirmation`, `finalOutcome`는 원매물 필드 복사본이 아니라 실행 시점 스냅샷을 가져야 한다.

### 65.11 화면 요구사항 파생
#### 등록 화면
- `개당 / 총액 / 묶음당 / 제안받기`를 명시적으로 고르게 해야 한다.
- `묶음 판매`를 켠 경우 구성 설명 예시를 제공한다.
- 부분거래 허용 여부와 잔여 처리 정책을 선택 또는 기본값으로 보여줘야 한다.

#### 목록/상세 화면
- 가격 옆에 반드시 basis 라벨을 붙인다.
- `묶음`, `일괄`, `부분거래 가능` 같은 배지를 지원한다.
- 잔여 수량이 줄어든 매물은 `남은 수량` 기준으로 갱신하되, 필요하면 `원래 100개 중 40개 남음` 같은 history hint를 추가할 수 있다.

#### 채팅/완료 화면
- 합의 수량과 총액을 구조화된 카드로 보여줘야 한다.
- 완료 확인 시 `실제 받은 수량/묶음`을 체크하는 UX가 있어야 한다.
- 분쟁 진입 시 `수량이 달랐어요`, `묶음 구성이 달랐어요`, `가격이 달랐어요` 같은 구체 사유를 제공한다.

### 65.12 분석 이벤트 후보
- `listing_unit_basis_selected`
- `listing_price_basis_selected`
- `listing_partial_trade_enabled`
- `deal_quantity_confirmed`
- `deal_price_basis_changed`
- `partial_completion_recorded`
- `trade_unit_mismatch_reported`

핵심 분석 포인트:
- `per_unit_price` vs `total_price` vs `offer_based`에 따른 채팅 전환율 차이
- `bundle` 매물의 예약 확정률과 분쟁률
- 부분거래 허용 매물의 완료율/재노출율
- `residualUnitPolicy`별 재거래 효율과 사용자 혼선 비율

### 65.13 오픈 질문 / 가정
- `아데나`처럼 사실상 통화에 가까운 품목을 일반 `stack_count`로 볼지 별도 `currency_like` 단위로 분리할지 결정 필요
- 카테고리별 기본 `quantityLabel`, `quantityGranularity`, `bundleRule` 템플릿을 운영자가 얼마나 세밀하게 관리할지 결정 필요
- 목록 가격 정렬에서 basis가 다른 매물을 완전 분리할지, 비교 가능 가격으로 약식 정규화할지 결정 필요
- MVP에서 `lot` 타입을 실제 노출할지 내부 모델만 선반영할지 결정 필요

#### 64.3.2 clarity
구매자가 매물을 읽고 바로 이해할 수 있는가를 본다.
- 제목이 지나치게 짧거나 반복문자/이모지/광고성 표현 위주인지
- 설명이 품목, 수량, 옵션, 거래 조건을 구분해서 전달하는지
- 가격/수량/옵션 표현이 상충하지 않는지
- 자유입력 텍스트가 지나치게 모호한지 (`문의`, `가격제시`, `아무거나`, `급처`) 여부

#### 64.3.3 freshness
지금도 실제 거래 가능한 매물인지 본다.
- 최근 수정/상태 점검/응답 활동 시각
- `reserved` 장기 방치 또는 `available` 장기 미활동 여부
- 접속 가능 시간 정보가 너무 오래되었는지
- price/quantity/status가 최근 거래 상황과 맞는지 확인 시점

#### 64.3.4 trust_readiness
거래 상대가 안심하고 문의할 만한 최소 신호가 있는지 본다.
- 작성자 verification/trust/restriction 상태
- 첫 거래 사용자 여부 자체보다, 매물 정보가 실제 접선·완료까지 이어질 준비가 되었는지
- 응답성/활동성 신호 존재 여부
- 과거 허위성 패턴과 유사한 작성 패턴 여부

#### 64.3.5 policy_safety
정책 위반 또는 검토 필요 위험을 포함하는지 본다.
- 금지 키워드, 외부 연락처, 민감정보, 사기 유도 표현
- 이미지/OCR 위험 탐지
- 고위험 카테고리/신규 계정/반복 위반 조합
- 운영 검토 대기 또는 제한 조치 존재 여부

### 64.4 qualityTier 결정 규칙
| qualityTier | 의미 | 기본 해석 |
|---|---|---|
| `excellent` | 거래 판단/문의에 필요한 정보가 충분하고 위험 신호가 낮음 | 검색/추천 가산 가능 |
| `good` | 기본 거래에는 무리 없으나 일부 보완 여지 존재 | 정상 노출 |
| `needs_improvement` | 등록은 가능하지만 정보 부족/모호함/오래됨으로 전환 저하 우려 | 경고/보완 유도/일부 랭킹 감점 |
| `restricted` | 정책/위험 신호로 제한 노출 또는 검토가 필요 | 검색 제한 또는 운영 큐 |

결정 우선순위:
1. `policy_safety`가 block/review 수준이면 `restricted`
2. `completeness` 또는 `clarity`가 최소 기준 미달이면 `needs_improvement`
3. `freshness`가 기준 초과 stale이면 `needs_improvement` 또는 `good` 하향
4. 모든 축이 기준 충족이고 핵심 보완점이 없으면 `good` 이상
5. 이미지/속성/응답성/최근 활동까지 우수하면 `excellent`

### 64.5 qualityFindingCode 후보
| 코드 | 의미 | 기본 사용자 액션 |
|---|---|---|
| `TITLE_TOO_GENERIC` | 제목이 지나치게 일반적 | 제목 구체화 |
| `DESCRIPTION_TOO_SHORT` | 설명 정보 부족 | 설명 보강 |
| `MISSING_PRICE_CONTEXT` | 가격 타입 대비 가격 맥락 부족 | 가격/협상 조건 보강 |
| `MISSING_IMAGE_FOR_HIGH_VALUE` | 고가/고위험 매물에 이미지 없음 | 이미지 추가 |
| `MISSING_REQUIRED_ATTRIBUTES` | 카테고리 필수 속성 누락 | 속성 입력 |
| `STALE_AVAILABILITY_INFO` | 가능 시간/최근 활동 정보 오래됨 | 상태 점검 |
| `RESERVED_TOO_LONG` | 예약중 장기 방치 | 예약 정리 |
| `POLICY_RISK_CONTACT_INFO` | 연락처/외부 유도 위험 | 정책 위반 수정 |
| `POLICY_RISK_PROHIBITED_ITEM` | 금지 품목/행위 의심 | 등록 불가/운영 검토 |
| `TRUST_SIGNAL_WEAK` | 신규/저신뢰 계정 + 정보 부족 결합 | 정보 보강/인증 유도 |
| `PRICE_OUTLIER_WARNING` | 비정상 가격 패턴 | 가격 재확인 |
| `MEETING_INFO_NOT_READY` | 거래 방식 대비 가능 장소/시간 정보 미흡 | 시간/장소 보강 |

원칙:
- 동일 매물에 여러 finding이 공존할 수 있다.
- 사용자에게는 상위 1~3개만 노출해 과도한 경고를 피한다.
- 내부적으로는 full finding set을 보관해 운영/analytics가 재사용한다.

### 64.6 등록 화면 UX 파생 규칙
- 작성 중에도 `listingQualitySnapshot` preview를 계산해 `게시 가능`, `보완하면 더 잘 팔려요`, `정책 수정 필요` 수준으로 즉시 피드백을 제공할 수 있어야 한다.
- 저장 차단이 아닌 soft warn 대상은 게시를 허용하되, publish 직전에 보완 CTA를 제시한다.
- 등록 완료 직후 아래와 같은 요약을 제공할 수 있다.
  - `거래 정보 충분 · 정상 노출 예정`
  - `이미지를 추가하면 신뢰도가 더 올라가요`
  - `거래 가능 시간을 적으면 문의 전환율이 높아질 수 있어요`
- 정책/위험 이슈는 품질 개선 메시지와 섞지 않고 별도 섹션으로 분리한다.

### 64.7 검색/홈/상세 노출 규칙
#### 검색/목록
- `excellent`, `good`는 정상 노출 대상이다.
- `needs_improvement`는 노출 가능하나 `rankingQualityImpact=downrank_minor` 또는 `downrank_moderate`가 적용될 수 있다.
- `restricted`는 기본 검색 결과 제외 또는 제한 노출 대상으로 본다.
- 카드에 내부 점수를 직접 노출하지 않고, 필요 시 `정보보완 필요`, `오래된 매물`, `검토중` 정도의 사용자 문구만 사용한다.

#### 홈 추천/재방문 모듈
- 추천 슬롯은 품질과 freshness를 함께 본다.
- low quality 매물은 단순 클릭 수가 높아도 홈 추천 슬롯에 과도하게 노출하지 않는다.
- 작성자 본인 홈에는 `품질 개선 필요 매물` 모듈을 제공해 재활성화/수정 행동을 유도할 수 있다.

#### 매물 상세
- 상세에서는 구매자에게 신뢰 판단을 돕는 최소 힌트만 제공한다.
- 예: `최근 정보 업데이트됨`, `거래 가능 시간 확인됨`, `오래된 매물일 수 있음`
- 품질 경고가 있어도 정책상 허용 매물이면 채팅 CTA를 무조건 막지 않고, 필요한 경우 기대값만 조정한다.

### 64.8 운영/자동화 개입 규칙
| qualityReviewState | 의미 | 기본 동작 |
|---|---|---|
| `none` | 운영 개입 없음 | 정상 |
| `soft_warned` | 사용자 수정 유도 필요 | 등록/내 매물에 경고 노출 |
| `limited_visibility` | 제한 노출 | 검색/추천 감점, 본인에게 사유 표시 |
| `review_required` | 운영 검토 필요 | 공개 노출 제한 또는 보류 |

운영 원칙:
- 품질 저하만으로 즉시 제재하지 않는다. 다만 반복적 저품질/허위성 패턴은 운영 이상징후와 연결할 수 있다.
- `qualityFindingCode`가 반복되어도 먼저 수정 유도→제한 노출→검토 순으로 escalte하는 것이 원칙이다.
- 정책 위반형 finding은 content moderation case와 연결되며, 품질 시스템이 제재의 단독 원천이 되어서는 안 된다.

### 64.9 canonical read model 초안
```json
{
  "listingQualitySnapshot": {
    "listingId": "listing_123",
    "qualityTier": "needs_improvement",
    "qualityReviewState": "soft_warned",
    "rankingQualityImpact": "downrank_minor",
    "dimensionScores": {
      "completeness": "low",
      "clarity": "medium",
      "freshness": "high",
      "trustReadiness": "medium",
      "policySafety": "pass"
    },
    "topFindings": [
      "MISSING_REQUIRED_ATTRIBUTES",
      "MEETING_INFO_NOT_READY"
    ],
    "improvementActions": [
      "fill_required_attributes",
      "add_availability_info"
    ],
    "lastEvaluatedAt": "2026-03-14T08:24:00+09:00"
  }
}
```

### 64.10 생성/갱신 트리거
`listingQualitySnapshot`은 아래 이벤트 시 재계산 대상이 된다.
- 매물 생성/수정/재게시/상태 점검
- 가격/수량/속성/이미지 변경
- availability/meeting preference 변경
- trust/restriction/verification 상태 변화
- content moderation decision 반영
- stale batch, reserved abuse batch, refresh batch 실행

원칙:
- 랭킹용 계산과 사용자 안내용 계산이 완전히 다른 기준을 쓰지 않도록 같은 snapshot을 공유하되, surface별 노출 수준만 다르게 한다.
- 전체 재계산보다 대상 매물 부분 재계산을 우선한다.

### 64.11 DB/API 파생 기준
#### DB 후보
- `ListingQualitySnapshot` projection 또는 materialized table
- 주요 필드 후보:
  - `listingId`
  - `qualityTier`
  - `qualityReviewState`
  - `rankingQualityImpact`
  - `topFindingCodesJson`
  - `improvementActionsJson`
  - `lastEvaluatedAt`
  - `staleSince`
  - `qualityVersion`

#### Public API 응답 후보
- `GET /listings`
- `GET /listings/{listingId}`
- `GET /me/listings`
응답에 surface별로 아래 필드를 선택적으로 포함할 수 있다.
- `qualityTier` (본인 또는 제한적 공개)
- `qualityHints`
- `staleBadge`
- `improvementActions` (본인 뷰 중심)

#### Admin API 후보
- `GET /admin/listings/{listingId}/quality`
- `POST /admin/listings/{listingId}/quality/recompute`
- `POST /admin/listings/{listingId}/quality/override-visibility`

### 64.12 analytics / KPI 파생 포인트
이 섹션은 단순 품질 점수 추적보다 `품질 개선이 실제 거래 성사에 기여하는가`를 보는 데 목적이 있다.

이벤트 후보:
- `listing_quality_evaluated`
- `listing_quality_warning_shown`
- `listing_quality_improvement_clicked`
- `listing_quality_tier_changed`
- `listing_quality_recovered`

핵심 분석 포인트:
- `qualityTier`별 상세 진입 대비 채팅 시작 전환율
- 이미지/가용시간/속성 보강 후 문의율 상승 여부
- `needs_improvement -> good` 회복까지 걸리는 평균 시간
- soft warning 이후 실제 수정률
- stale quality가 reserved abuse/no-show와 상관관계가 있는지

### 64.13 가정 및 오픈 질문
- MVP에서 `qualityTier`를 사용자 공개 화면에 어디까지 노출할 것인가? 기본안은 본인 매물 중심 노출, 공개 화면에는 최소 힌트만 노출이다.
- quality evaluation을 동기 계산으로 둘지, publish 성공 후 비동기 projection으로 둘지 결정 필요.
- 카테고리별 required attribute 수준과 `고가 매물` 판정 기준은 별도 운영 표로 확정해야 한다.
- low quality와 low trust를 하나의 경고로 합칠지, 사용자에게는 분리된 이유로 보여줄지 디자인 결정이 필요하다.

### 63.9 API 후보 정리
#### 사용자용
- `POST /trade-completions/{completionId}/dispute`
- `GET /trade-completions/{completionId}/dispute`
- `POST /disputes/{disputeId}/statements`
- `GET /disputes/{disputeId}`

#### 관리자용
- `GET /admin/disputes`
- `GET /admin/disputes/{disputeId}`
- `POST /admin/disputes/{disputeId}/request-more-info`
- `POST /admin/disputes/{disputeId}/resolve`

응답 필드 후보:
```json
{
  "completionId": "comp_123",
  "listingStatus": "completed",
  "completionStage": "disputed",
  "dispute": {
    "disputeId": "disp_123",
    "disputeStatus": "under_review",
    "nextActionForViewer": "submit_statement"
  }
}
```

### 63.10 화면/운영 시사점
- 내 거래 화면은 `completed`만으로 끝내지 않고 `확인 대기`, `분쟁 진행 중`, `운영 검토 중`을 구분 표시해야 한다.
- 후기 작성 CTA는 `completionStage in (confirmed_by_counterparty, auto_confirmed, resolved_completed)`일 때만 활성화하는 것이 안전하다.

## 64. 예약 재일정(Reschedule) / 시간·장소 변경 협상 계약
### 64.1 목표
- 예약 확정 이후 발생하는 시간 변경, 장소 변경, 방식 변경(`in_game` ↔ `offline_pc_bang`)을 일반 채팅 잡담이 아니라 **명시적 협상 객체**로 다룬다.
- 당일 거래 흐름에서 `지각`, `장소 혼선`, `일정 변경`, `노쇼`가 서로 섞여 판정되지 않도록 경계를 분리한다.
- 화면/UI copy, API, DB, 알림, 운영 runbook이 모두 동일한 `재일정` vocabulary를 사용하게 한다.

### 64.2 적용 범위
재일정은 아래 상황에 적용한다.
- 확정된 예약(`reservationStatus=confirmed`)의 시간 변경
- 확정된 예약의 장소/서버/거래방식 변경
- 당일 `arrivalState` 이전 또는 이후의 지연/변경 요청
- 노쇼 신고 직전 제기된 "늦어요 / 장소 바꿔요 / 30분 뒤 가능" 유형의 협상

적용 제외:
- 아직 `proposed` 단계인 예약안 수정은 일반 예약 수정 흐름으로 처리한다.
- 거래 완료 후 회고성 메시지는 재일정 객체를 생성하지 않는다.
- 단순 사과/도착예고("5분 후 도착")는 임계치 이하일 경우 arrival signal로만 처리할 수 있다.

### 64.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `rescheduleState` | `none` / `requested` / `counter_proposed` / `accepted` / `rejected` / `expired` / `superseded` / `cancelled_after_request` | 재일정 협상 자체의 상태 |
| `rescheduleScope` | `time_only` / `location_only` / `time_and_location` / `method_change` | 어떤 축이 바뀌는지 |
| `rescheduleInitiatorType` | `seller` / `buyer` / `system_recovery` / `moderator` | 변경 시작 주체 |
| `rescheduleUrgencyTier` | `same_day` / `within_24h` / `advance_notice` | 거래 시점과의 거리 |
| `rescheduleFailureReasonCode` | `no_response` / `too_many_changes` / `past_cutoff` / `counterparty_rejected` / `meeting_window_expired` | 실패/종결 사유 |
| `rescheduleImpact` | `reservation_reconfirmed_required` / `arrival_timer_reset` / `no_show_grace_extended` / `listing_state_unchanged` | 승인 시 원 예약에 미치는 영향 |

원칙:
- 재일정은 "예약의 부분 수정"이지만, 감사/분쟁 추적을 위해 append-only 이력 객체로 남긴다.
- 최종 예약 스냅샷과 재일정 제안 이력을 분리해 저장해야 한다.

### 64.4 재일정 상태 머신
#### 기본 흐름
- `confirmed reservation` + 변경 요청 생성 → `rescheduleState=requested`
- 상대가 대안 시간/장소를 다시 제안하면 → `counter_proposed`
- 한쪽이 최종 수락하면 → `accepted`
- 명시 거절 시 → `rejected`
- 응답 기한 초과 시 → `expired`
- 새 재일정 제안이 이전 제안을 대체하면 → `superseded`

#### 해석 규칙
- `accepted`가 되기 전까지 원래 예약은 기본적으로 여전히 유효하다. 단, UI에서는 `재확인 필요` 배지를 띄운다.
- `same_day` 재일정이 `accepted`되면 `scheduledAt`, meeting snapshot, arrival grace 기준 시간을 새 값으로 재설정한다.
- `rejected` 또는 `expired`가 되어도 원예약이 아직 지나지 않았다면 원예약은 유지될 수 있다.
- 원예약 시각이 지났고 재일정도 성립하지 않으면 `meetingExecutionState`는 `at_risk` 또는 `missed` 후보로 전환된다.

### 64.5 재일정 생성 제한과 cutoff 규칙
- 재일정은 `reservationStatus=confirmed`인 예약에 대해서만 생성 가능하다.
- 같은 예약에 동시에 활성 재일정 제안은 1건만 허용한다.
- 당일 거래에서 예약 시각 직전 N분(가정: 10분) 이내의 시간 변경은 허용하되 `same_day` 고위험으로 태깅한다.
- 예약 시각이 이미 grace period를 초과한 뒤 생성된 재일정 요청은 기본적으로 `past_cutoff` 처리 후보다. 단, 상대방 명시 수락 시 운영상 유효한 재협상으로 인정 가능하다.
- 동일 예약에서 짧은 시간 내 재일정 요청이 과도하면(`too_many_changes`) 경고/운영 탐지 신호로 누적한다.

### 64.6 arrival / no-show와의 관계
- `arrivalState=arrived` 이후의 재일정은 기본적으로 `location_only` 또는 짧은 `time_only` 변경만 허용하는 보수안을 둔다.
- 한쪽이 `no_show_claimed`를 제출한 뒤에는 일반 재일정 생성이 불가하고, 운영자 또는 상대방 명시 동의가 있을 때만 `system_recovery` 성격으로 재협상 열 수 있다.
- `accepted`된 재일정은 노쇼 판단 grace period를 새 시각 기준으로 재설정한다.
- `requested` 상태의 재일정만 존재하고 상대 확인이 없는 경우, 노쇼 판단은 원시각 기준을 유지한다. 즉, 일방적 "늦어요" 메시지만으로 자동 연장되지 않는다.

### 64.7 화면 요구사항
#### 채팅 화면
- 확정 예약 카드에 `시간 변경`, `장소 변경`, `다시 제안` CTA를 노출할 수 있어야 한다.
- 재일정 제안은 일반 메시지가 아니라 시스템 카드/inline card로 표현한다.
- 카드 필수 필드:
  - 현재 예약 요약
  - 제안된 변경값
  - 변경 범위(`time_only`, `location_only` 등)
  - 응답 마감 시각
  - 원예약 유지 여부
  - 상대 응답 CTA(`수락`, `거절`, `대안 제시`)
- 재일정이 열린 동안 채팅 헤더/내 거래 요약에 `재확인 필요` 배지 노출이 필요하다.

#### 내 거래 화면
- 재일정 요청이 있는 거래는 일반 미읽음보다 높은 우선순위를 가져야 한다.
- `same_day` 재일정은 `action needed` 최상위 그룹으로 노출한다.
- 거래 카드에는 `원래 21:00 → 제안 21:30` 같은 diff 요약이 보여야 한다.

#### 당일 실행 카드
- `도착했어요`, `늦어요`, `장소 다시 보기` CTA와 재일정 CTA를 분리한다.
- `늦어요` CTA가 단순 메시지인지, 실제 `time_only` 재일정인지 사용자가 구분 가능해야 한다.

### 64.8 API 후보
#### 사용자용
- `POST /reservations/{reservationId}/reschedules`
- `POST /reschedules/{rescheduleId}/accept`
- `POST /reschedules/{rescheduleId}/reject`
- `POST /reschedules/{rescheduleId}/counter`
- `GET /reservations/{reservationId}/reschedules`

Request 예시:
```json
{
  "scope": "time_and_location",
  "proposedScheduledAt": "2026-03-14T21:30:00+09:00",
  "meeting": {
    "meetingType": "in_game",
    "serverId": "ken-01",
    "meetingPointText": "기란 마을 창고 앞"
  },
  "reasonCode": "running_late",
  "message": "30분 뒤 가능해요"
}
```

Response 예시:
```json
{
  "rescheduleId": "rsc_123",
  "reservationId": "res_123",
  "rescheduleState": "requested",
  "rescheduleScope": "time_and_location",
  "responseDueAt": "2026-03-14T21:05:00+09:00",
  "preservesOriginalReservation": true,
  "availableActions": ["accept_reschedule", "reject_reschedule", "counter_reschedule"]
}
```

### 64.9 데이터 모델 후보
#### RescheduleRequest
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `rescheduleId` | 필수 | 재일정 식별자 |
| `reservationId` | 필수 | 원 예약 참조 |
| `listingId` | 필수 | 조회/운영 편의용 |
| `chatRoomId` | 필수 | 채팅 연동 |
| `requestedByUserId` | 필수 | 제안자 |
| `counterpartyUserId` | 필수 | 상대 |
| `rescheduleState` | 필수 | 상태 |
| `rescheduleScope` | 필수 | 변경 범위 |
| `urgencyTier` | 필수 | same_day 등 |
| `proposedScheduledAt` | 선택 | 제안 시간 |
| `meetingSnapshotJson` 또는 FK | 선택 | 제안 장소/방식 스냅샷 |
| `reasonCode` | 선택 | 지각/장소이슈 등 |
| `messageText` | 선택 | 사용자 설명 |
| `responseDueAt` | 필수 | 응답 마감 |
| `acceptedAt` | 선택 | 수락 시각 |
| `rejectedAt` | 선택 | 거절 시각 |
| `supersededByRescheduleId` | 선택 | 대체 제안 연결 |
| `createdAt` | 필수 | 생성 시각 |

원칙:
- 수락 시 원 예약 row를 덮어쓰더라도, 변경 전/후 snapshot은 모두 복구 가능해야 한다.
- `Reservation.lastAcceptedRescheduleId` 같은 역참조 캐시를 둘 수 있다.

### 64.10 알림 규칙
- `requested`: 상대방에게 T1 또는 높은 T2 수준 알림 발송
- `same_day` 재일정: 푸시 기본 ON, 미열람 시 재리마인드 1회 허용
- `accepted`: 양측에 새 시간/장소 요약 알림 발송
- `rejected`/`expired`: 제안자에게 결과 알림, 원예약 유지 여부를 함께 고지
- 잠금화면/푸시 본문에는 exact 장소 대신 요약 라벨만 사용한다.

### 64.11 운영/안티어뷰즈 규칙
- 재일정 남용 지표 후보:
  - 예약당 재일정 횟수
  - same-day 재일정 비율
  - 수락 없는 일방 재일정 반복률
  - 재일정 후 no-show 전환률
- `too_many_changes` 패턴은 노쇼/괴롭힘/시간끌기 탐지의 보조 신호로 사용한다.
- 운영자는 아래를 한 화면에서 볼 수 있어야 한다.
  1. 원예약
  2. 각 재일정 제안 diff
  3. 상대 읽음/응답 여부
  4. accepted 시각
  5. 이후 no-show claim 또는 dispute 연결

### 64.12 Dispute 연결 규칙
- 분쟁 생성 시 `activeRescheduleId` 또는 `lastAcceptedRescheduleId`를 함께 연결할 수 있어야 한다.
- 노쇼 분쟁에서 핵심 판단 질문은 아래와 같다.
  1. 마지막으로 상호 수락된 약속 시각/장소는 무엇이었는가?
  2. 일방 요청만 있었는가, 상호 합의가 있었는가?
  3. 재일정 후 grace period가 재설정되었는가?
- 따라서 no-show 판정은 항상 `accepted` 재일정 기준으로만 원예약을 덮어쓴다.

### 64.13 분석 이벤트 후보
- `reschedule_requested`
- `reschedule_accepted`
- `reschedule_rejected`
- `reschedule_counter_proposed`
- `reschedule_expired`
- `reschedule_same_day_high_risk_flagged`

핵심 KPI 연결:
- 예약 대비 재일정 발생률
- 재일정 수락률
- 재일정 이후 거래 완료율
- 재일정 이후 no-show/분쟁 전환율
- same-day 재일정이 first-trade cohort에 미치는 영향

### 64.14 오픈 질문 / 기본 가정
- 가정: `requested` 재일정 응답 기한은 원예약 시각 또는 제안 후 30분 중 더 이른 값으로 둔다. 실제 값은 운영 테스트로 보정 필요.
- 가정: 당일 재일정은 최대 2회까지 허용하고, 초과 시 운영 신호를 남긴다.
- 결정 필요: `늦어요` 같은 short-delay를 별도 quick action으로 둘지, 항상 `time_only` 재일정으로 강제할지.
- 결정 필요: `accepted` 이후 원예약 로그를 사용자 화면에 얼마나 자세히 보여줄지(전체 diff vs 최신값 중심).
- 결정 필요: same-day 재일정에서 장소만 바뀌는 경우 no-show grace를 그대로 둘지, 일부 재설정할지.
- 운영 대시보드에서 `completion dispute rate`, `auto confirmed rate`, `dispute overturn rate`를 분리해 봐야 한다.
- DB 설계 시 `TradeCompletion`만 확장할지, `Dispute`를 별도 테이블로 분리할지 빠르게 결정해야 API/화면 명세가 안정된다.

## 64. 공개 프로필 레벨별 API 계약(초안)
### 64.1 목표
- 같은 프로필이라도 비회원, 로그인 회원, 거래 상대가 필요한 정보 수준이 다르므로, API 응답 계약을 레벨별로 명시한다.
- 화면에서 임의 마스킹 규칙을 중복 구현하지 않고 서버가 응답 레벨에 맞는 필드만 반환하도록 설계한다.

### 64.2 응답 레벨 정의
| 레벨 | 대상 | 사용 화면 | 목적 |
|---|---|---|---|
| `summary` | 비회원 포함 공개 뷰어 | 목록 카드, 매물 상세 상단 | 최소 신뢰 요약 |
| `member` | 로그인 회원 | 프로필 화면, 채팅 진입 전 | 일반 거래 판단 |
| `participant` | 거래 상대/후기 작성 가능 관계 | 채팅, 예약, 완료 후 프로필 참조 | 실제 거래 실행 판단 |
| `staff` | 운영자 | 백오피스 | 운영/안전 판단 |

### 64.3 레벨별 필드 계약 원칙
- 상위 레벨은 하위 레벨 필드를 포함하되, 민감정보는 별도 권한 플래그 없이는 내려주지 않는다.
- 클라이언트는 누락 필드를 `null`로 가정하지 말고, 레벨에 따라 미제공될 수 있음을 전제로 구현한다.
- `profileVisibilityLevel`과 `viewerRelationship`를 함께 반환해 디버깅/QA를 단순화한다.

### 64.4 `summary` 레벨 응답 후보
사용처: 목록 카드, 상세 상단, 비회원 공개 영역

```json
{
  "userId": "user_123",
  "profileVisibilityLevel": "summary",
  "viewerRelationship": "public_viewer",
  "nicknameMasked": "린***",
  "avatarUrl": null,
  "trust": {
    "badge": "거래경험많음",
    "completedTradeCountBand": "20_plus",
    "positiveReviewRatioBand": "high"
  },
  "activity": {
    "lastActiveRelative": "3일 이내",
    "primaryServerLabel": "데포로쥬"
  }
}
```

원칙:
- 정확한 최근 접속시각, 후기 원문, 신고/제재 세부는 미포함
- 닉네임은 마스킹 기본안 유지
- 완료 거래 수는 exact count보다 구간형 표현을 우선 검토

### 64.5 `member` 레벨 응답 후보
사용처: 로그인 후 프로필 화면, 거래 시작 전 판단

```json
{
  "userId": "user_123",
  "profileVisibilityLevel": "member",
  "viewerRelationship": "logged_in_member",
  "nickname": "린클상인",
  "avatarUrl": "https://...",
  "introduction": "저녁 시간대 거래 가능합니다.",
  "trust": {
    "badge": "거래경험많음",
    "completedTradeCount": 27,
    "positiveReviewCount": 21,
    "reviewSummary": {
      "recommend": 21,
      "notRecommend": 2
    }
  },
  "activity": {
    "lastActiveRelative": "24시간 이내",
    "primaryServerLabel": "데포로쥬"
  },
  "safety": {
    "isBlockedByViewer": false,
    "hasMutedByViewer": false
  }
}
```

원칙:
- 후기 원문은 일부 preview만 내려주거나 별도 `/users/{userId}/reviews`에서 조회
- 내부 제재 이력, 신고 누적 수, 탐지 점수는 포함하지 않음
- 차단/뮤트는 본인 기준 boolean만 응답

### 64.6 `participant` 레벨 응답 후보
사용처: 채팅/예약/거래 실행 중 상대 프로필 패널

```json
{
  "userId": "user_123",
  "profileVisibilityLevel": "participant",
  "viewerRelationship": "trade_counterparty",
  "nickname": "린클상인",
  "avatarUrl": "https://...",
  "trust": {
    "badge": "거래경험많음",
    "completedTradeCount": 27,
    "recentReviewHighlights": [
      "약속시간 잘 지켜요",
      "응답이 빨랐어요"
    ]
  },
  "activity": {
    "lastActiveRelative": "방금 전"
  },
  "tradeContext": {
    "sharedListingCount": 1,
    "currentTradeRole": "seller",
    "canShareCharacterName": true
  }
}
```

원칙:
- `participant`는 실제 거래 실행에 필요한 관계형 맥락만 추가 제공한다.
- 전화번호/실명/정확 위치 같은 직접 식별 정보는 여전히 별도 정책 없이 프로필 API로 제공하지 않는다.
- 캐릭터명 공유 여부는 프로필 고정값이 아니라 현재 거래 컨텍스트 기준 힌트로 다루는 것이 안전하다.

### 64.7 `staff` 레벨 응답 후보
- 백오피스는 일반 프로필 API를 재사용하기보다 운영 전용 DTO/엔드포인트를 쓰는 것이 바람직하다.
- `staff` 레벨에서만 가능한 정보 예시:
  - 최근 신고 건수 구간
  - 최근 제재 이력 요약
  - reserved 악용/노쇼/보복성 후기 탐지 플래그
  - 열람 시 감사 로그 필요 여부
- 단, 민감정보 원문은 `staff` 내부에서도 추가 열람 권한을 다시 검증해야 한다.

### 64.8 엔드포인트 전략 후보
옵션 A:
- `GET /users/{userId}/profile`
- 서버가 `viewer` 관계를 계산해 자동으로 레벨 선택

옵션 B:
- `GET /users/{userId}/profile-summary`
- `GET /users/{userId}/profile`
- `GET /chats/{chatRoomId}/counterparty-profile`

권장 기본안:
- 외부 앱 API는 옵션 A처럼 단일 엔드포인트를 우선하고,
- 응답에 `profileVisibilityLevel`을 포함해 실제 적용 레벨을 명시한다.
- 거래 맥락이 필요한 경우 `chatRoomId` 또는 `listingId` 컨텍스트를 쿼리로 받을 수 있다.

### 64.9 API/캐시/보안 시사점
- 프로필 응답은 viewer 관계에 따라 달라지므로 CDN 공개 캐시보다 사용자별 private cache를 우선 고려해야 한다.
- 목록/상세에서 필요한 `summary` 정보는 Listing 응답 내부에 임베드하는 것이 호출 수 절감에 유리하다.
- 차단/정지/탈퇴 직후 프로필 노출 레벨이 즉시 반영되도록 캐시 무효화 전략이 필요하다.

### 64.10 분석/QA 파생 포인트
- 분석 이벤트 후보:
  - `profile_view_summary`
  - `profile_view_member`
  - `profile_view_participant`
  - `profile_block_click`
- QA 체크포인트:
  - 비회원과 로그인 회원이 같은 사용자 프로필에서 서로 다른 필드 집합을 받는가
  - 거래 상대가 아닌 회원에게 `participant` 전용 필드가 노출되지 않는가
  - 차단/탈퇴/정지 후 프로필 응답 레벨과 CTA가 즉시 축소되는가

## 65. 매물(Listing) 계약 상세화 초안
### 65.1 목표
- 매물을 단순 게시글이 아니라 **검색 가능한 거래 단위 + 채팅/예약의 기준 객체 + 운영 판단 단위**로 정의한다.
- 이후 DB 스키마, API DTO, 화면 컴포넌트, validation, 랭킹 로직이 모두 같은 필드 의미를 쓰도록 한다.
- 매물 작성/수정 시 자유도를 유지하되, 목록 노출/검색/운영 판단에는 구조화 필드를 우선 사용한다.

### 65.2 매물 필드 묶음 정의
| 필드 묶음 | 목적 | 대표 필드 |
|---|---|---|
| 식별/소유 | 객체 식별과 권한 판단 | `listingId`, `authorUserId`, `listingType` |
| 검색/정규화 | 검색, 자동완성, 필터 | `serverId`, `categoryId`, `itemNameRaw`, `normalizedItemId` |
| 거래 조건 | 실제 거래 판단 | `priceType`, `priceAmount`, `quantity`, `tradeMethod` |
| 표시/카드 | 목록/상세 렌더링 | `title`, `summaryText`, `coverImageUrl`, `statusBadge` |
| 상태/흐름 | 거래 진행 연결 | `status`, `reservedChatRoomId`, `completionAt` |
| 운영/품질 | 노출 제어와 감사 | `visibility`, `lastActivityAt`, `expiresAt`, `reasonCode` |

원칙:
- `title`, `description`, `itemNameRaw`는 사용자 입력값을 최대한 보존한다.
- 검색/필터/랭킹은 가능하면 `serverId`, `categoryId`, `normalizedItemId`, `status`, `visibility` 같은 구조화 필드를 우선 사용한다.
- 목록 카드에서 필요한 축약값은 런타임 계산보다 저장/캐시 가능한 파생 필드로 관리할 수 있다.

### 65.3 사용자 입력 필드 vs 시스템 파생 필드
| 구분 | 필드 예시 | 생성 주체 | 수정 가능성 |
|---|---|---|---|
| 사용자 입력 원본 | `itemNameRaw`, `title`, `description`, `priceAmount`, `quantity`, `availableTimeText` | 작성자 | 상태 정책 범위 내 수정 가능 |
| 시스템 정규화 | `normalizedItemId`, `searchKeywords`, `summaryText` | 서버/배치 | 직접 수정 불가 |
| 집계/캐시 | `viewCount`, `favoriteCount`, `chatCount` | 시스템 | 직접 수정 불가 |
| 운영 제어 | `visibility`, `reasonCode`, `hiddenAt` | 운영/시스템 | 일반 사용자 수정 불가 |
| 흐름 연결 | `reservedChatRoomId`, `completionAt` | 사용자 액션 + 시스템 검증 | 임의 수정 불가, 상태 전이로만 변경 |

설계 원칙:
- 작성자가 수정 가능한 값과 시스템이 소유하는 값을 명확히 분리해야 API 권한 모델이 단순해진다.
- 클라이언트는 파생 필드를 신뢰하고, 직접 재계산하지 않는 방향이 바람직하다.

### 65.4 핵심 필드 상세 계약
#### 65.4.1 `itemNameRaw` / `normalizedItemId`
- `itemNameRaw`: 사용자가 입력한 대표 아이템명 원문
- `normalizedItemId`: 카탈로그/정규화가 가능한 경우 매핑되는 내부 식별자
- 원칙:
  - `itemNameRaw`는 항상 보존
  - `normalizedItemId`는 null 허용
  - 검색은 `normalizedItemId` exact match > alias match > raw partial match 순으로 가중치 부여 가능

#### 65.4.2 `title`
- 목록 카드와 상세 상단에서 가장 먼저 보이는 사용자 노출 제목
- 기본안: 사용자가 직접 입력하되, 입력 편의를 위해 `아이템명 + 강화/옵션 + 가격` 기반 추천 제목을 생성할 수 있다.
- 제한:
  - 과도한 이모지/반복 문자/외부 연락처 유도 문구는 validation 또는 후처리 대상
  - 상태 진행 중에는 제목을 바꿀 수 있어도, 거래 상대가 혼동할 정도의 핵심 의미 변경은 제한해야 한다.

#### 65.4.3 `summaryText`
- 목록 카드/공유 미리보기/SEO 설명에 쓰이는 짧은 요약 필드
- 작성자 자유 입력이 아니라 서버 파생 필드로 두는 것을 권장한다.
- 예시 조합:
  - `+7 장검 · 1개 · 데포로쥬 · 50만 아데나`
- 장점:
  - 목록/알림/공유 UI 문구 일관성 확보
  - 여러 화면에서 중복 문자열 조합 로직 제거

#### 65.4.4 `priceType` / `priceAmount` / `currencyType`
- `priceType=fixed`: `priceAmount` 필수
- `priceType=negotiable`: `priceAmount` 필수, 단 협의 가능 배지 노출
- `priceType=offer`: `priceAmount` null 허용, 목록에서는 `제안받기` 성격 문구 사용
- `currencyType`:
  - MVP 우선안은 단일 통화 체계로 단순화하되, 장기적으로는 enum 분리 가능성을 열어둔다.
  - 외부 규제/정책 민감성이 있으므로 실제 노출/입력 범위는 별도 정책 문서에서 확정 필요

#### 65.4.5 `quantity`
- 수량형 거래를 지원하되, MVP에서는 **단일 매물 = 하나의 대표 거래 단위**로 해석한다.
- `quantity > 1`인 경우에도 활성 예약은 1개만 유지하는 기본안을 유지한다.
- 일부 거래가 발생하면:
  1. 잔여 수량 수동 수정 + 상태 재검토
  2. 또는 복제 후 재등록
- 자동 부분분할은 Post-MVP로 미루는 것이 안전하다.

#### 65.4.6 `tradeMethod`
- 매물 수준 선호 방식: `in_game` / `offline_pc_bang` / `either`
- 예약 수준의 실제 방식(`meetingType`)과 구분한다.
- 규칙:
  - `tradeMethod=in_game`이면 예약 단계에서 오프라인 제안 차단 또는 강한 경고
  - `tradeMethod=either`이면 예약 확정 시 실제 방식 구체화 필수

#### 65.4.7 `status` / `visibility`
- `status`: 거래 흐름 상태
- `visibility`: 공개 노출 제어 상태
- 두 필드는 독립 축으로 관리한다.
- 예시:
  - `status=available`, `visibility=public` → 일반 노출
  - `status=available`, `visibility=hidden` → 작성자/운영자만 보임
  - `status=completed`, `visibility=public` → 공개 목록 제외, 소프트 랜딩 가능
- 같은 `completed`라도 공개 노출 여부는 `visibility` 정책에 따라 달라질 수 있다.

### 65.5 상태별 수정 가능 필드 규칙
| 상태 | 수정 가능 필드 | 제한 필드 | 비고 |
|---|---|---|---|
| `available` | 대부분의 사용자 입력 필드 | 소유/시스템 필드 | 일반 수정 가능 |
| `reserved` | 설명 보조문구, 가능시간, 이미지 일부 | 가격, 서버, 아이템명, 거래방식 핵심 변경 제한 권장 | 대기 문의 혼선 방지 |
| `pending_trade` | 매우 제한적 | 가격, 수량, 서버, 아이템명, 제목 핵심 변경 금지 | 실행 직전 혼란 방지 |
| `completed` | 수정 불가 또는 메모성 필드만 | 대부분 금지 | 기록 보존 우선 |
| `cancelled` | 직접 수정 대신 복제 후 재등록 권장 | 대부분 금지 | 과거 기록 보존 |

보충 원칙:
- 핵심 거래 조건 변경이 필요한 경우에는 수정이 아니라 `취소 후 재등록` 또는 `복제 후 재등록`을 우선 유도한다.
- 수정 제한은 UX에서 버튼 숨김만 하지 말고 서버도 동일하게 강제해야 한다.

### 65.6 목록 카드용 파생 필드 후보
목록 성능과 일관성을 위해 아래 필드는 API 응답에서 직접 제공하는 것이 유리하다.

| 필드 | 설명 |
|---|---|
| `displayPriceLabel` | `50만 아데나`, `가격 협의`, `제안받기` 등 |
| `displayStatusLabel` | `거래 가능`, `예약중`, `거래대기` 등 |
| `displayQuantityLabel` | `1개`, `세트`, `수량 10` 등 |
| `activityLabel` | `방금 전`, `3시간 전`, `3일 이내` |
| `trustSummaryLabel` | `거래 20+ · 후기 좋음` 등 |
| `thumbnailUrl` | 대표 이미지 축약 URL |
| `listingCardFlags` | `new`, `price_dropped`, `reserved_waiting`, `high_response` 등 후보 |

원칙:
- 화면은 원시 데이터로 라벨을 매번 조합하기보다 서버 제공 라벨을 우선 사용한다.
- 단, 접근성/다국어를 위해 원시값과 표시값을 함께 제공하는 것이 안전하다.

### 65.7 읽기 DTO와 쓰기 DTO 분리 원칙
#### 쓰기(Create/Update) DTO
- 사용자가 직접 입력하거나 선택한 값만 받는다.
- 예: `title`, `description`, `itemNameRaw`, `priceType`, `priceAmount`, `quantity`, `tradeMethod`, `serverId`

#### 읽기(Read) DTO
- 화면 렌더링과 정책 힌트까지 포함한다.
- 예: `summaryText`, `displayPriceLabel`, `availableActions`, `viewerContext`, `reasonCode`, `authorProfileSummary`

원칙:
- 쓰기 DTO에 파생 필드(`summaryText`, `viewCount`, `chatCount`)를 허용하지 않는다.
- 읽기 DTO는 화면 수를 줄이기 위해 카드/상세/내매물 관점으로 projection을 다르게 둘 수 있다.

### 65.8 Listing API projection 기본안
| projection | 사용 화면 | 필수 포함 필드 |
|---|---|---|
| `card` | 목록, 홈 추천, 찜 목록 | 제목, 가격 라벨, 서버, 상태, 썸네일, 작성자 요약 |
| `detail` | 매물 상세 | 카드 정보 + 설명, 이미지 전체, 거래 방식, 가능 시간, policyHints |
| `owner_manage` | 내 매물 | detail + 집계값, 최근 채팅/예약 요약, 수정 가능 여부 |
| `admin_review` | 백오피스 | detail + 신고/숨김/이력/사유 코드 |

권장 원칙:
- 같은 엔드포인트에서 viewer/role 기반으로 projection이 달라질 수 있다.
- 단, 문서화와 QA 편의를 위해 응답 레벨 이름은 명시적으로 관리하는 것이 좋다.

### 65.9 운영/랭킹 관점의 매물 품질 신호 후보
- 제목/설명 최소 충실도 충족 여부
- 이미지 존재 여부
- `normalizedItemId` 매핑 성공 여부
- 최근 응답/상태 갱신 여부
- 장기 `reserved`/반복 취소 패턴 여부
- 가격 이상치 여부
- 신고 검토/운영 제한 여부

활용 원칙:
- 이 신호는 검색 랭킹, 노출 감점, 운영 검토 우선순위에 쓰일 수 있다.
- 사용자에게는 점수 자체보다 행동 결과(노출 감점, 경고 배너, 수정 유도)로만 제한적으로 드러낸다.

### 65.10 후속 파생 문서 포인트
- 기능명세서:
  - 상태별 수정 가능 필드 표
  - 목록 카드/상세/내매물 projection 정의
- DB 스키마:
  - raw vs normalized 컬럼 분리
  - summary/label 캐시 필드 여부
- API 명세:
  - Create/Update payload와 Read DTO 분리
  - projection/query param 전략
- 화면명세:
  - 카드 라벨 우선순위, 상태 배지, 수정 제한 경고 UX
- 운영정책:
  - 가격 이상치, 허위 `reserved`, 제목 스팸, 구조화 필드 누락 대응 기준

## 66. 내 거래(Trade Workspace) 화면/데이터 계약 초안
### 66.1 목표
- `내 거래`를 단순 채팅 목록이 아니라 **거래 실행과 종결을 관리하는 워크스페이스**로 정의한다.
- 사용자가 지금 당장 해야 할 행동(예약 응답, 도착 확인, 완료 확인, 분쟁 소명)을 최우선으로 보게 한다.
- 채팅방/예약/완료/분쟁 객체가 분산되어 보여 사용자에게 상태 혼선을 주지 않도록, 화면 단위 집계 객체를 별도로 정의한다.

### 66.2 핵심 개념: Trade Thread
- `Trade Thread`는 사용자가 체감하는 하나의 거래 실행 단위다.
- 기본 기준:
  - 1개 `listingId`
  - 1개 상대 사용자
  - 1개 대표 `chatRoomId`
  - 현재 활성 `reservationId(optional)`
  - 현재 활성 `completionId(optional)`
  - 현재 활성 `disputeId(optional)`
- 구현 관점에서는 별도 테이블일 수도, 읽기 모델(read model)일 수도 있다.
- 핵심은 `내 거래` 화면이 채팅방 raw 목록을 그대로 쓰지 않고, 거래 실행 맥락으로 정리된 projection을 사용해야 한다는 점이다.

### 66.3 Trade Thread 상태 집계 규칙
| tradeThreadStatus | 의미 | 대표 근거 객체 | 사용자 우선 액션 |
|---|---|---|---|
| `inquiry_open` | 일반 문의 진행 중 | ChatRoom=open | 메시지 응답 |
| `reservation_waiting_response` | 예약안 응답 필요 | Reservation=proposed | 수락/거절/대안 제시 |
| `reservation_confirmed` | 예약 확정, 거래 전 | Reservation=confirmed | 장소/시간 확인 |
| `trade_due_soon` | 거래 임박 | Reservation + time window | 도착/지연 공유 |
| `completion_waiting_me` | 상대 완료 요청에 내 응답 필요 | TradeCompletion=requested | 완료 확인/이의 제기 |
| `completion_waiting_counterparty` | 내가 완료 요청했고 상대 응답 대기 | TradeCompletion=requested | 대기, 필요 시 신고 |
| `dispute_open` | 분쟁/운영 검토 중 | Dispute=open/under_review | 소명 제출/상태 확인 |
| `closed_completed` | 최종 완료 | completion confirmed | 후기 작성/기록 보기 |
| `closed_cancelled` | 거래 취소/불발 종료 | cancelled or resolved_not_completed | 유사 매물 보기/재문의 |

원칙:
- `tradeThreadStatus`는 `Listing.status`, `ChatRoom.chatStatus`, `Reservation.reservationStatus`, `TradeCompletion.completionStage`를 바탕으로 계산되는 **화면용 상태**다.
- 모바일 클라이언트는 원시 상태 4종을 직접 조합하기보다 서버가 내려주는 `tradeThreadStatus`를 우선 사용한다.

### 66.4 내 거래 목록 화면 요구사항
목적: 사용자가 여러 거래를 병렬 관리할 때, 행동 우선순위가 높은 거래를 놓치지 않게 한다.

필수 그룹/탭 후보:
1. `액션 필요`
2. `예정됨`
3. `진행 중`
4. `완료`
5. `종료/분쟁`

정렬 원칙:
1. 액션 필요 여부
2. 마감/예약 시각 임박도
3. 미읽음 존재 여부
4. 최신 활동 시각

카드 필수 노출 요소:
- 상대 사용자 요약
- 매물 제목 또는 요약명
- 현재 tradeThreadStatus 배지
- 다음 액션 문구 1개
- 예약 시각/장소 요약(optional)
- 미읽음 수
- 최근 메시지 또는 시스템 이벤트 1줄

예시 카드 문구:
- `오늘 21:00 · 기란 마을 · 예약 응답 필요`
- `상대가 거래완료를 요청했어요`
- `분쟁 소명 제출 마감: 내일 18:00`

### 66.5 거래 상세(Trade Detail) 화면 요구사항
목적: 특정 거래 1건의 실행 상태를 종단적으로 확인하고, 다음 행동을 한 곳에서 수행한다.

필수 섹션:
1. 헤더 요약
   - 상대 프로필 요약
   - 매물 요약
   - 현재 상태 배지
   - 타이머/SLA 문구
2. 다음 액션 영역
   - 사용자가 지금 할 수 있는 행동 1~3개만 강조
3. 예약/장소 카드
   - 시간, 방식, 장소 요약, 변경 이력, 재확인 필요 여부
4. 거래 타임라인
   - 문의 시작
   - 예약 제안/확정/변경
   - 상태 변경
   - 완료 요청/확정
   - 분쟁/운영 처리
5. 채팅 진입 영역
   - 전체 메시지 보기 CTA 또는 인라인 최근 메시지
6. 안전/정책 가드레일
   - 외부 연락처, 고위험 사용자, 분쟁 안내 배너

화면 원칙:
- 채팅 화면은 대화 중심, 거래 상세는 **결정과 실행 중심**으로 역할을 분리한다.
- 완료 대기, 분쟁, 임박 예약은 채팅방 안에 묻히지 않도록 거래 상세 상단에서 재강조해야 한다.

### 66.6 상태별 헤더/CTA 계약
| tradeThreadStatus | 헤더 배지 | 주 CTA | 보조 CTA | 금지/숨김 |
|---|---|---|---|---|
| `inquiry_open` | 문의중 | 채팅 보기 | 예약 제안 | 완료 처리 |
| `reservation_waiting_response` | 예약 응답 필요 | 수락 | 거절, 대안 제시 | 완료 처리 |
| `reservation_confirmed` | 예약 확정 | 채팅 보기 | 일정/장소 재확인 | 후기 작성 |
| `trade_due_soon` | 거래 임박 | 도착 메시지 보내기 | 늦음 알리기, 채팅 보기 | 새 예약 생성 |
| `completion_waiting_me` | 완료 확인 대기 | 완료 확인 | 이의 제기 | 새 예약 생성 |
| `completion_waiting_counterparty` | 상대 확인 대기 | 상태 보기 | 신고 | 후기 작성 |
| `dispute_open` | 분쟁 진행 중 | 소명 제출 | 진행 상태 보기 | 일반 완료 확정 |
| `closed_completed` | 거래완료 | 후기 작성/보기 | 기록 보기 | 예약/메시지 주요 입력 |
| `closed_cancelled` | 거래종료 | 유사 매물 보기 | 기록 보기 | 완료 확인 |

원칙:
- 화면은 항상 `주 CTA` 1개를 가장 강하게 보여주고, 나머지는 보조로 내린다.
- 상태가 복합적이어도 사용자가 지금 해야 할 행동이 하나로 정리되어야 한다.

### 66.7 타이머/SLA 노출 규칙
| 상황 | 사용자 노출 문구 예시 | 데이터 기준 |
|---|---|---|
| 예약 제안 응답 대기 | `오늘 18:00까지 응답 필요` | `Reservation.expiresAt` |
| 예약 임박 | `1시간 후 거래 예정` | `scheduledAt - now` |
| 완료 확인 대기 | `내일 10:00까지 확인하지 않으면 자동 완료` | completion timeout |
| 분쟁 소명 대기 | `운영 요청 자료 제출 마감: 3월 15일 18:00` | dispute request due |

원칙:
- 리스트에는 상대시간 중심(`1시간 후`, `내일까지`)을, 상세에는 절대시각을 함께 보여주는 것이 바람직하다.
- 자동화 배치 기준 시각과 화면 문구가 어긋나지 않도록 서버가 계산한 `deadlineLabel`, `deadlineAt`을 함께 주는 구조를 권장한다.

### 66.8 `GET /me/trades` projection 초안
목적: 홈 위젯, 내 거래 목록, 알림 딥링크, 배지 집계를 공통 projection으로 해결한다.

응답 item 후보:
```json
{
  "tradeThreadId": "tt_123",
  "listingId": "listing_123",
  "chatRoomId": "chat_123",
  "reservationId": "res_123",
  "completionId": null,
  "disputeId": null,
  "tradeThreadStatus": "reservation_waiting_response",
  "sortPriority": 920,
  "unreadCount": 2,
  "listing": {
    "title": "+7 장검 팝니다",
    "summaryText": "+7 장검 · 데포로쥬 · 50만 아데나",
    "thumbnailUrl": "https://..."
  },
  "counterparty": {
    "userId": "user_456",
    "nickname": "린클상인",
    "trustBadge": "거래경험많음"
  },
  "nextAction": {
    "code": "confirm_or_decline_reservation",
    "label": "예약 응답하기",
    "deadlineAt": "2026-03-13T18:00:00+09:00",
    "deadlineLabel": "오늘 18:00까지"
  },
  "meetingSummary": {
    "scheduledAt": "2026-03-13T21:00:00+09:00",
    "summaryLabel": "오늘 21:00 · 기란 마을"
  },
  "lastEvent": {
    "type": "reservation_proposed",
    "label": "상대가 예약을 제안했어요",
    "createdAt": "2026-03-13T07:10:00+09:00"
  },
  "availableActions": ["confirm_reservation", "decline_reservation", "open_chat"]
}
```

설계 원칙:
- `tradeThreadId`는 프론트가 목록/상세/푸시 딥링크에서 공통으로 쓰는 식별자다.
- `sortPriority`는 서버 계산값을 둘 수 있으나, 설명 가능한 규칙을 유지해야 한다.
- `nextAction`은 반드시 0개 또는 1개 대표 행동을 제공해 모바일 UI 결정을 단순화한다.

### 66.9 거래 상세 응답 projection 초안
후보 엔드포인트:
- `GET /me/trades/{tradeThreadId}`
- 또는 `GET /trade-threads/{tradeThreadId}`

필수 응답 묶음:
- `summary`: 현재 상태, 상대, 매물, 대표 타이머
- `timeline`: 시스템 이벤트/상태 전이 목록
- `meeting`: 현재 확정 장소/변경 대기 상태
- `completion`: 완료 확인/자동확정/분쟁 상태
- `safety`: 경고 배너, 정책 힌트
- `availableActions`: 현재 가능한 액션 목록
- `deepLinks`: `chat`, `listingDetail`, `counterpartyProfile`

### 66.10 알림/홈/배지 연계 규칙
- 홈은 `GET /me/trades` 중 상위 N개 `action required` 거래만 요약 노출할 수 있어야 한다.
- 푸시 딥링크는 가능하면 `tradeThreadId`를 기준으로 이동하고, 화면 내부에서 채팅/예약/분쟁 탭으로 분기한다.
- 전역 배지는 `미읽음 메시지 수`와 별도로 `액션 필요 거래 수`를 구분 집계하는 것이 좋다.
- 알림함에서 `예약 응답 필요`, `완료 확인 필요`, `분쟁 소명 필요`를 클릭하면 동일 거래 상세로 수렴시키는 구조가 바람직하다.

### 66.11 분석 이벤트 후보
- `trade_thread_list_view`
- `trade_thread_open`
- `trade_thread_primary_action_click`
- `trade_thread_chat_open`
- `trade_thread_deadline_missed`
- `trade_thread_resolved_completed`
- `trade_thread_resolved_cancelled`

핵심 분석 포인트:
- 내 거래 화면 진입 후 실제 필요한 액션 수행까지 걸리는 시간
- `completion_waiting_me` 상태에서 자동확정 전 직접 확인률
- `reservation_waiting_response` 상태에서 응답 전환율
- 분쟁 상태에서 소명 제출률과 처리 기간

### 66.12 후속 파생 문서 포인트
- 화면명세:
  - 내 거래 목록 카드 규격
  - 거래 상세 레이아웃, 타임라인, CTA 우선순위
- API 명세:
  - `GET /me/trades`, `GET /me/trades/{tradeThreadId}` 응답 계약
  - `tradeThreadStatus`, `nextAction`, `deadlineLabel` enum/필드 정의
- DB/백엔드 설계:
  - read model materialization 여부
  - `tradeThreadId` 생성 전략
  - 정렬 우선순위 계산 배치/실시간화 전략
- QA:
  - 예약 응답 대기/완료 확인 대기/분쟁 진행 중 상태가 목록과 상세에서 동일하게 표현되는지 검증

## 63. 가정 및 오픈 질문
### 63.1 현재 가정
- 로그인 사용자는 최소한 휴대폰/소셜 기반 식별이 가능하다고 가정
- 거래는 외부 결제 없이 당사자 간 직접 이행한다고 가정
- 서버/아이템 체계는 서비스 내부 표준 카탈로그 또는 자유 입력 혼합 구조일 수 있음
- 오프라인 만남(예: PC방)과 인게임 만남을 모두 지원 대상으로 본다
- MVP에서는 숫자 기반 평점보다 후기 + 배지 중심 신뢰 표현이 더 적합하다고 가정
- 운영자는 거래를 중개하지만 법적 에스크로/결제 보증 주체는 아니라고 가정
- 공개 인덱싱은 `available` 매물 중심 제한 공개가 기본안이라고 가정
- 후기 공개는 상호 작성 완료 또는 작성 마감 도래 후 공개가 기본안이라고 가정
- 매물/예약/신고/완료 객체는 현재 상태 + 상태 이력 저장 구조를 기본안으로 가정
- 권한 판단은 전역 역할보다 객체 관계와 현재 상태를 우선 반영하는 정책을 기본안으로 가정
- 핵심 읽기 API는 `viewerContext`, `availableActions`, `policyHints` 같은 UI 제어 컨텍스트를 포함하는 방향이 적합하다고 가정
- 거래/신고 관련 핵심 데이터는 hard delete보다 soft delete/비식별화를 기본 전략으로 가정
- 공개 프로필은 검색엔진 인덱싱보다 회원 내 신뢰 확인 수단으로 먼저 설계하는 것이 적합하다고 가정
- 분쟁 중 추가 소명은 일반 채팅이 아닌 별도 사건 단위 입력으로 분리하는 것이 적합하다고 가정

### 63.2 오픈 질문
- 비회원에게 어느 수준까지 매물을 공개할 것인가?
- 공개 매물의 검색엔진 색인을 어디까지 허용할 것인가?
- 거래완료는 단독 확정인지, 상호 확인인지?
- 예약 자동 만료 시간은 필요한가?
- 매물 만료일/끌어올리기 정책은 필요한가?
- 전화번호 등 직접 연락처 교환을 어느 단계부터 허용할 것인가?
- 계좌번호/오픈채팅 링크 등 민감 패턴을 경고만 할지 차단할지?
- 실시간 채팅을 MVP에 포함할지, 폴링 기반으로 시작할지?
- 신뢰도 점수를 숫자로 노출할지, 뱃지/단계형으로 노출할지?
- 후기 수정 허용 범위를 어디까지 둘 것인가?
- 후기 단독 공개 시 보복성 리스크를 어떻게 완화할 것인가?
- 신고 처리 SLA와 야간 운영 범위를 어디까지 둘 것인가?
- `reserved` 상태에서도 신규 문의를 받을 것인가, 차단할 것인가?
- 공개 SEO 유입과 로그인 전환 사이의 최적 균형은 무엇인가?
- 완료/취소 매물 공개 URL을 소프트 랜딩으로 유지할지, 즉시 404/410으로 전환할지?
- 차단 시 상대 매물을 전체 숨김할지, 단지 채팅만 막을지?
- 정지/탈퇴 계정의 본인 기록 조회 범위를 어디까지 허용할 것인가?
- `availableActions` 기반 응답 모델을 전 API 공통 패턴으로 채택할 것인가?
- 탈퇴 전 진행 중 거래가 있는 사용자의 탈퇴를 즉시 허용할지, 종결/경고/유예 흐름을 둘지?
- 메시지/후기 원문에 대한 마스킹과 원본보관의 저장 분리 수준을 어디까지 설계할 것인가?
- 공개 프로필을 로그인 회원에게만 보여줄지, 일부 비회원 공개를 허용할지?
- 보복성 후기 판정 시 자동 점수 반영을 보류할지, 공개만 가릴지?
- 분쟁 중 상대방에게 소명 원문을 공유할지, 운영자 전용으로만 둘지?

## 67. 채팅 동기화 / 읽음 / 전송 보장 계약(초안)
### 67.1 목표
- 채팅을 단순 메시지 목록이 아니라 **거래 이벤트 타임라인**으로 안정적으로 동기화한다.
- 모바일 네트워크 불안정, 앱 백그라운드 전환, 푸시 누락, 중복 전송 상황에서도 사용자가 거래 흐름을 잃지 않게 한다.
- 화면, API, DB, 알림 시스템이 동일한 읽음/미읽음/전송 상태 의미를 공유하도록 기준을 정의한다.

### 67.2 핵심 개념
| 개념 | 의미 | 설계 포인트 |
|---|---|---|
| `timeline event` | 일반 텍스트, 시스템 이벤트, 예약 카드, 완료/분쟁 로그를 모두 포함하는 채팅 타임라인 원소 | 단일 정렬축 필요 |
| `clientMessageId` | 클라이언트가 생성하는 전송 dedup 키 | 재시도/중복 방지 |
| `eventCursor` | 타임라인 재동기화 기준점 | 실시간 실패 시 backfill |
| `lastReadEventId` | 사용자가 마지막으로 읽은 이벤트 식별자 | unread 계산 기준 |
| `deliveryStatus` | 서버 수신/상대 전달/읽음 상태 | 메시지 버블/재시도 UI |
| `gapDetected` | 클라이언트가 이벤트 누락 가능성을 감지한 상태 | 목록/상세 재조회 트리거 |

원칙:
- unread는 단순 숫자 카운터보다 `lastReadEventId` 기준이 우선이다.
- 채팅 목록의 `unreadCount`는 캐시값일 수 있으나, 서버는 언제든 `lastReadEventId` 기준으로 재계산 가능해야 한다.
- 일반 메시지와 시스템 이벤트 모두 타임라인 정렬에는 포함하되, unread 집계 포함 여부는 이벤트 타입별 정책이 필요하다.

### 67.3 이벤트 타입별 unread 반영 원칙
| eventType | unread 포함 기본안 | 비고 |
|---|---|---|
| `text` | 포함 | 일반 메시지 |
| `image` | 포함 | 첨부 메시지 |
| `system_status_change` | 포함 | 중요한 상태 변경은 읽음 필요 |
| `reservation_card_created` | 포함 | 행동 유도가 필요함 |

## 68. 문의 SLA / 거래 스레드 Aging / 자동 휴면 정리 정책(초안)
### 68.1 목표
- 문의가 많이 쌓여도 판매자/구매자가 지금 응답해야 할 거래를 빠르게 식별할 수 있게 한다.
- 장기간 열린 채팅이 실제 진행 중 거래처럼 남아 검색/알림/내 거래 화면을 오염시키지 않게 한다.
- 응답성 신호, 리드 대기열, 내 거래 워크스페이스, 운영 이상징후 탐지가 같은 `stale / dormant / closed` 해석을 쓰도록 맞춘다.

### 68.2 설계 원칙
1. **채팅방을 자동 삭제하지 않는다.** 대신 활성도와 기대 행동을 기준으로 `aging state`를 부여한다.
2. **사용자 간 침묵 자체를 제재하지 않는다.** 다만 반복적 무응답과 방치 패턴은 랭킹/알림/운영 신호로 활용할 수 있다.
3. **거래 실행 흐름이 있는 스레드(예약/완료/분쟁)는 일반 문의보다 우선 보호**한다.
4. **휴면 전환은 가시성/우선순위 정책**이며, 이력 보존과 재활성화 가능성을 함께 가져야 한다.

### 68.3 Trade Thread Aging 상태 정의
| agingStatus | 의미 | 대표 조건 | 사용자 영향 |
|---|---|---|---|
| `fresh` | 최근 생성/응답되어 활성 흐름 | 최근 메시지 또는 상태 변경이 SLA 내 | 일반 우선순위 |
| `awaiting_my_response` | 내가 응답해야 하는 상태 | 상대 메시지/예약 제안 이후 내 응답 없음 | 액션 필요 최상단 |
| `awaiting_counterparty` | 상대 응답 대기 | 내가 마지막 행동 후 상대 무응답 | 대기 상태 표시 |
| `stale_conversation` | 일반 문의가 장시간 정리되지 않음 | 예약/완료 없이 일정 시간 경과 | 정렬 후순위, 리마인드 후보 |
| `dormant` | 사실상 종료됐으나 명시 종결 안 됨 | 장기 무활동 + 활성 예약/분쟁 없음 | 기본 접힘/보관함 이동 후보 |
| `reactivated` | 휴면 후 다시 메시지/상태 변경 발생 | dormant 이후 신규 이벤트 | 상단 재노출 |
| `locked_operational` | 정책/분쟁/제재로 잠김 | report_locked, dispute under review | 읽기 위주, 액션 제한 |

원칙:
- `agingStatus`는 `tradeThreadStatus`와 별도 축이다. 예를 들어 `reservation_confirmed` 거래는 aging 기준으로 stale 처리하지 않는다.
- 예약 확정, 완료 확인 대기, 분쟁 진행 중인 스레드는 자동 `dormant` 전환 대상에서 제외한다.

### 68.4 기본 응답 SLA 계약
| 상황 | 기준 시각 | 기본 SLA 초안 | 초과 시 처리 |
|---|---|---|---|
| 일반 첫 문의에 대한 작성자 첫 응답 | 첫 상대 메시지 시각 | 12시간 | 응답성 지표 반영, 리마인드 후보 |
| 일반 대화 중 후속 응답 | 마지막 상대 메시지 시각 | 24시간 | `awaiting_my_response` 유지 |
| 예약 제안 응답 | reservation proposedAt | 제안 만료 전 또는 12시간 이내 중 더 이른 시점 | 강한 액션 필요 상태 |
| 장소/시간 변경 재확인 | locationChangePending 시각 | 6시간 또는 예약 시각 전 | 임박 알림 우선 |
| 완료 확인 응답 | completion requestedAt | 24시간 | 자동확정 또는 분쟁 진입 |
| 운영 추가 소명 응답 | admin requestAt | 운영이 지정한 기한 | 분쟁 처리 기준 반영 |

주의:
- 위 SLA는 서비스 내부 행동 유도 기준이며, 사용자에게 계약상 의무처럼 표현하지 않는다.
- 실제 숫자는 운영 데이터 기반으로 조정 가능하며, 본 문서에서는 초기 가정값으로 둔다.

### 68.5 일반 문의 Aging 규칙
- 새 채팅방 생성 후 24시간 동안 상호 메시지가 1회도 오가지 않으면 `stale_conversation` 후보가 된다.
- 최근 72시간 동안 메시지/예약/상태 변경이 없고 활성 우선상대도 아니면 `stale_conversation`으로 본다.
- 최근 14일 동안 아무 활동이 없고 예약/완료/분쟁 연결이 없으면 `dormant`로 전환할 수 있다.
- `dormant` 스레드는 삭제하지 않고:
  - 내 거래 기본 목록에서는 접거나 보관함 탭으로 이동
  - 채팅 목록 검색/기록에서는 조회 가능
  - 신규 메시지 수신 시 즉시 `reactivated` 처리

### 68.6 판매자 우선상대/대기열과 Aging 연계
- `reservedLead` 또는 현재 우선상대 채팅은 일반 무응답 기준만으로 stale 처리하지 않는다.
- 다만 아래 경우는 예외적으로 경고/점검 대상이다.
  1. 우선상대 지정 후 일정 시간 내 아무 메시지/예약 진전 없음
  2. 우선상대가 반복 교체되지만 완료 전환이 없음
  3. 대기열 문의가 많음에도 일괄 안내/종결 없이 장기 방치
- 작성자는 대기 문의에 대해 `지금은 다른 상대와 진행 중`, `거래 불발 시 다시 연락`, `문의 종료` 같은 빠른 정리 액션을 사용할 수 있어야 한다.

### 68.7 자동 휴면 정리 액션
| 트리거 | 시스템 액션 | 사용자 노출 | 운영 기록 |
|---|---|---|---|
| 일반 문의 72시간 무활동 | `stale_conversation` 표시 | 내 거래/채팅에 `응답 없음` 또는 `정리 필요` 배지 | stale reason 저장 |
| 14일 무활동 | `dormant` 전환 | 기본 목록 후순위/보관함 이동 | dormant 전환 로그 |
| dormant 이후 새 메시지 | `reactivated` 전환 | 상단 재노출 + 알림 | reactivated 로그 |
| 대기열 문의 일괄 종료 | thread close reason 기록 | 상대방에게 안내 시스템 메시지 | 종료 사유 저장 |
| 우선상대 장기 무진전 | 작성자 점검 리마인드 | `우선 거래 상태를 정리해 주세요` | abuse/stale 분석 신호 |

원칙:
- 시스템은 일반 문의를 자동 `cancelled` 처리하지 않는다.
- 명시 종결이 필요한 경우 `chatClosedReasonCode` 또는 `leadClosedReasonCode`를 남기는 구조를 우선한다.

### 68.8 사용자 화면 요구사항
#### 내 거래
- `액션 필요`에는 `awaiting_my_response`, `completion_waiting_me`, `reservation_waiting_response`를 우선 노출한다.
- `stale_conversation`은 별도 `정리 필요` 섹션 또는 후순위 카드 배지로 표시한다.
- `dormant`는 기본 탭에서 숨기고 `보관된 문의`에서 다시 볼 수 있어야 한다.

#### 채팅 목록
- unread와 stale를 혼동하지 않도록 배지를 분리한다.
- 예시 배지:
  - `답변 필요`
  - `대기 중`
  - `휴면`
  - `운영 잠금`
- dormant 스레드도 검색 결과에는 나타날 수 있어야 하며, 복구 없이 바로 재대화 가능해야 한다.

#### 매물 관리 화면
- 매물별 문의 수를 단순 총합이 아니라 아래로 분리 표시하는 것이 좋다.
  - `응답 필요 문의`
  - `대기 문의`
  - `휴면 문의`
  - `종결 문의`
- 고빈도 거래자는 여러 문의를 일괄 정리할 수 있는 quick action을 필요로 할 수 있다.

### 68.9 데이터 모델 후보
#### TradeThreadAging
| 필드 | 설명 |
|---|---|
| `tradeThreadId` | 거래 스레드 식별자 |
| `agingStatus` | 현재 aging 상태 |
| `responseDueAt` | 현재 뷰어 기준 응답 기대 시각 |
| `staleAt` | stale 판정 시각 |
| `dormantAt` | dormant 전환 시각 |
| `reactivatedAt` | 휴면 후 재활성화 시각 |
| `staleReasonCode` | `no_reply`, `seller_not_responding`, `buyer_not_responding`, `waiting_queue_idle` 등 |
| `lastMeaningfulEventAt` | 일반 메시지/예약/완료/분쟁 등 의미 있는 마지막 이벤트 시각 |
| `lastNudgeAt` | 마지막 리마인드 발송 시각 |
| `nudgeCount` | 최근 기간 리마인드 횟수 |

설계 원칙:
- `lastMeaningfulEventAt`는 단순 읽음 동기화나 시스템 heartbeat성 이벤트를 제외할 수 있어야 한다.
- stale/dormant 판단은 읽기 모델에서 계산 가능하지만, 모바일 정렬/운영 분석 효율을 위해 materialized field를 둘 수 있다.

### 68.10 API 후보
- `GET /me/trades?agingStatus=awaiting_my_response,dormant`
- `POST /trade-threads/{tradeThreadId}/archive`
- `POST /trade-threads/{tradeThreadId}/unarchive`
- `POST /trade-threads/{tradeThreadId}/close`
- `POST /trade-threads/{tradeThreadId}/nudge`

응답 필드 후보:
```json
{
  "tradeThreadId": "tt_123",
  "tradeThreadStatus": "inquiry_open",
  "aging": {
    "agingStatus": "awaiting_my_response",
    "responseDueAt": "2026-03-14T10:00:00+09:00",
    "staleReasonCode": null,
    "nudgeAllowed": true
  },
  "availableActions": ["open_chat", "send_quick_reply", "close_thread"]
}
```

### 68.11 알림/리마인드 정책
- 일반 문의 무응답 리마인드는 과도하면 스팸처럼 느껴지므로 보수적으로 발송한다.
- 기본안:
  1. 첫 문의 후 응답 없음 6~12시간 구간 1회
  2. stale 전환 직전 1회
  3. dormant 전환 시 푸시 대신 인앱 요약 우선
- 상대방에게는 `읽었는지`보다 `현재 응답 대기 중인지` 중심으로 보여주고, 수치 경쟁처럼 느껴지는 표현은 피한다.
- 같은 스레드에 대해 24시간 내 1회 초과 nudging은 지양한다.

### 68.12 운영/안티어뷰즈 연계
- 다수 스레드가 반복적으로 `awaiting_counterparty -> dormant`로 끝나는 사용자는 단순 비매너 신호로 참고할 수 있다.
- 다만 초보 사용자, 푸시 미허용 사용자, 저활동 시간대를 고려해 자동 제재 단독 근거로는 사용하지 않는다.
- 다음 패턴은 운영 모니터링 후보가 될 수 있다.
  - 대량 문의 생성 후 무응답
  - 반복 문의 후 예약 직전 이탈
  - 대기열만 모으고 일괄 종결 없는 판매자 패턴
  - 완료/예약 없이 장기 `open` 스레드가 비정상적으로 많은 계정

### 68.13 분석 이벤트 후보
- `trade_thread_aging_changed`
- `trade_thread_marked_stale`
- `trade_thread_marked_dormant`
- `trade_thread_reactivated`
- `trade_thread_nudge_sent`
- `trade_thread_closed_by_owner`
- `trade_thread_archived`

핵심 KPI 후보:
- 첫 문의 후 작성자 첫 응답까지 걸린 시간
- stale 스레드 비율 / 전체 문의 대비
- dormant 전환 전 재활성화율
- stale 스레드 정리 이후 예약 전환율
- 고빈도 거래자의 문의 정리 완료율

### 68.14 후속 파생 문서 포인트
- 화면명세: `액션 필요`, `정리 필요`, `보관된 문의` 구획 정의
- DB 스키마: `TradeThreadAging`, close/archive reason code, nudge 이력
- API 명세: archive/close/nudge 권한, 응답 필드, 429 제한
- 운영정책: 반복 무응답/장기 방치 패턴의 모니터링 기준
- QA: stale/dormant 전환, 재활성화, 예약/완료 예외 처리 시나리오
| `reservation_card_updated` | 포함 | 시간/장소 변경 인지 중요 |
| `reservation_expired` | 포함 | 후속 행동 필요 |
| `completion_requested` | 포함 | 매우 중요 |
| `completion_auto_confirmed` | 포함 | 결과 인지 필요 |
| `review_prompt` | 선택 | 앱 내 배지 포함 여부는 정책 결정 가능 |
| `typing` | 미포함 | 저장 안 함 |

원칙:
- 거래 실행과 직접 관련된 시스템 이벤트는 unread에 포함한다.
- purely ephemeral 상태(typing, connection state)는 DB 저장/미읽음 집계 대상에서 제외한다.

### 67.4 전송 상태 모델
메시지/이벤트 전송 상태는 클라이언트 UX와 운영 디버깅을 위해 최소 아래 단계를 가진다.

| 상태 | 의미 | 사용자 노출 |
|---|---|---|
| `pending_local` | 클라이언트 작성 후 서버 미수신 | 회색 시계/재시도 가능 |
| `accepted` | 서버 저장 완료 | 기본 전송 완료 |
| `delivered` | 상대 디바이스 또는 세션에 전달됨(가정) | 선택적 표시 |
| `read` | 상대 `lastReadEventId`가 해당 이벤트 이상 | 읽음 표시 |
| `failed` | 유효하지 않거나 재시도 가능 실패 | 오류 문구/재전송 CTA |
| `blocked` | 정책/권한 차단 | 수정 필요 안내 |

원칙:
- MVP에서는 `accepted`와 `read`만 노출해도 되지만, 내부적으로는 `failed`/`blocked` 구분이 필요하다.
- `delivered`는 인프라 복잡도를 높이므로 내부 상태만 두고 UI 노출은 Post-MVP로 미룰 수 있다.

### 67.5 메시지 dedup / 재전송 정책
- 클라이언트는 쓰기 메시지마다 `clientMessageId`를 생성해 서버에 전송한다.
- 서버는 `chatRoomId + senderUserId + clientMessageId` 기준 중복 방지를 지원해야 한다.
- 동일 `clientMessageId` 재전송 시:
  - 이미 성공 저장된 경우 기존 message/event를 그대로 반환
  - 영구 실패였던 경우 실패 코드 유지 또는 새 전송 요구
- 이미지/첨부 업로드도 가능하면 동일 dedup 흐름을 재사용한다.
- 네트워크 타임아웃으로 응답을 못 받았더라도 서버 저장 가능성이 있으므로, 클라이언트는 즉시 새 ID로 재전송하지 않고 기존 `clientMessageId` 결과 조회 또는 재시도부터 해야 한다.

### 67.6 읽음 처리 기준
- 채팅 상세 진입만으로 즉시 전부 읽음 처리하지 않고, **타임라인이 실제 렌더링된 마지막 이벤트** 기준으로 읽음 커서를 갱신하는 안을 기본으로 한다.
- 기본 읽음 API 후보:
  - `POST /chats/{chatRoomId}/read-cursor`
- 요청 예시:
```json
{
  "lastReadEventId": "evt_123",
  "readAt": "2026-03-13T07:35:00+09:00"
}
```
- 서버는 기존 커서보다 과거 이벤트로 후퇴시키는 요청을 거절하거나 무시해야 한다.
- 예약/완료 확인처럼 매우 중요한 이벤트는 단순 읽음과 별도로 명시적 액션(`confirm_reservation`, `confirm_completion`)이 필요하다.

### 67.7 채팅 목록 projection 계약
채팅 목록은 단순 last message 목록이 아니라, 재진입 우선순위 판단용 요약 정보를 포함해야 한다.

응답 item 후보:
```json
{
  "chatRoomId": "chat_123",
  "listingId": "listing_123",
  "lastEvent": {
    "eventId": "evt_999",
    "eventType": "completion_requested",
    "previewLabel": "상대가 거래완료를 요청했어요",
    "createdAt": "2026-03-13T07:31:00+09:00"
  },
  "lastReadEventId": "evt_870",
  "unreadCount": 4,
  "hasActionRequired": true,
  "gapDetected": false,
  "availableActions": ["open_chat", "open_trade_thread"]
}
```

필수 요건:
- `previewLabel`은 시스템 이벤트도 사람이 이해 가능한 1줄 요약으로 제공
- `hasActionRequired`는 단순 unread와 별도로 예약 응답/완료 확인/분쟁 요청 같은 행동 필요 상태를 반영
- `gapDetected=true`면 클라이언트는 상세 진입 시 강제 backfill을 수행해야 함

### 67.8 상세 동기화 / backfill 정책
#### 실시간 연결 가능 시
- 새 이벤트는 push/socket/SSE 등으로 즉시 수신 가능해야 한다.
- 실시간 연결이 끊겼다가 복구되면 `lastKnownEventCursor` 이후 diff를 조회한다.

#### 실시간 연결 불가 또는 MVP 단순안
- 일정 주기 polling 또는 화면 진입/복귀 시 pull 기반 동기화를 수행한다.
- 기본 API 후보:
  - `GET /chats/{chatRoomId}/events?afterEventId=...`
  - `GET /chats/{chatRoomId}/events?beforeEventId=...` (과거 pagination)

원칙:
- 클라이언트는 최신 메시지만 append하지 말고, 누락 여부를 감지하면 anchor 기준 재조회할 수 있어야 한다.
- 이벤트 정렬은 `createdAt`만이 아니라 서버 일관 정렬 키(`eventSeq`, `eventId`)를 갖는 것이 안전하다.

### 67.9 오프라인/앱 재진입 UX 원칙
- 오프라인 상태에서 작성한 메시지는 로컬에 `pending_local`로 남기고, 복구 시 자동 재시도한다.
- 예약 확정/완료 확인 같이 멱등성이 중요한 액션은 오프라인 큐에 넣더라도 중복 실행 방지를 최우선으로 해야 한다.
- 앱 재진입 시 아래를 순서대로 복구한다.
  1. 채팅 목록 unread/lastEvent 갱신
  2. 액션 필요 거래 우선 갱신
  3. 현재 열려 있는 채팅방 diff backfill
- 사용자가 오래된 기기에서 재접속해 대량 backlog를 받더라도, `다음 행동` 이벤트가 가장 먼저 눈에 띄어야 한다.

### 67.10 알림과 읽음/동기화의 연결 원칙
- 푸시 오픈만으로 읽음 처리하지 않는다. 실제 채팅/거래 상세에서 이벤트를 확인한 뒤 읽음 커서를 올린다.
- 이미 `lastReadEventId` 뒤로 소비된 이벤트에 대해 푸시 중복 발송이 발생하지 않도록 notification dedup key가 필요하다.
- notification dedup 후보 키:
  - `userId + chatRoomId + eventId + channel`
- 푸시가 누락되더라도 목록 unread와 거래 워크스페이스 상태는 서버 원본 기준으로 복구 가능해야 한다.

### 67.11 데이터 모델 시사점
`ChatMessage` 또는 통합 이벤트 테이블에 아래 필드 후보를 추가 검토한다.
- `eventId`
- `eventSeq`
- `clientMessageId(optional)`
- `deliveryStatus`
- `acceptedAt`
- `readByCounterpartyAt(optional)`
- `isUnreadCounted`
- `dedupScopeKey`

`ChatParticipantState` 또는 동등 객체 후보:
- `chatRoomId`
- `userId`
- `lastReadEventId`
- `lastReadAt`
- `lastDeliveredEventId(optional)`
- `mutedUntil(optional)`
- `hasGapDetected`

원칙:
- unreadCount를 ChatRoom 본문에만 두지 말고 participant별 상태 객체를 분리하는 것이 자연스럽다.
- 다기기 사용을 고려하면 `lastReadEventId`는 사용자 기준 공통 커서 우선, 기기별 커서는 확장 영역으로 둔다.

### 67.12 QA / 운영 체크포인트
- 같은 메시지가 네트워크 재시도로 두 번 생성되지 않는가
- 예약 변경 시스템 이벤트가 목록 preview와 상세 타임라인에 동일하게 반영되는가
- 오래 오프라인이었다가 재접속해도 unread, action required, lastEvent가 일관되게 복구되는가
- `report_locked` 상태에서 일반 메시지는 차단되고 시스템/운영 이벤트만 표시되는가
- 완료 확인 푸시를 눌렀지만 읽음 처리 없이 앱이 꺼졌을 때, 재진입 후 여전히 action required가 유지되는가

## 64. 사유 코드 / 정책 코드 표준(초안)
### 64.1 목적
- 상태 변경, 운영 제재, 자동화 실패, UI 경고 배너, 감사 로그가 각각 다른 자유 텍스트를 쓰지 않도록 공통 코드 체계를 정의한다.
- DB enum/string, API 응답, 운영 백오피스 액션 로그, 분석 이벤트 속성값을 하나의 사전으로 정렬한다.
- 사용자에게 노출하는 설명과 내부 운영 상세 사유를 분리해 안전성과 일관성을 확보한다.

### 56.2 설계 원칙
1. 모든 코드에는 **안정적인 machine code**와 **변경 가능한 human label**을 분리한다.
2. 코드값은 영문 `snake_case`를 기본으로 하고, prefix로 도메인을 구분한다.
3. 사용자 노출 문구는 코드에 직접 종속되지 않고 locale 템플릿에서 관리한다.
4. 내부 코드 추가는 가능하되, 이미 배포된 코드는 의미 변경보다 deprecated 처리 후 신규 코드 추가를 우선한다.
5. 하나의 이벤트에는 `primaryReasonCode` 1개와 `secondaryReasonCodes` 0..n개를 허용할 수 있다.
6. 운영 제재/신고 관련 코드는 외부 노출 시 축약된 공개 코드로 매핑 가능해야 한다.

### 56.3 공통 데이터 필드 후보
- `reasonCode`: 단일 대표 사유
- `reasonCodes`: 다중 사유 배열
- `reasonNote`: 운영자/시스템 내부 메모
- `publicReasonCode`: 사용자 노출용 축약 사유
- `reasonSource`: `user` / `counterparty` / `system` / `moderator` / `automation`
- `reasonVersion`: 코드 사전 버전
- `reasonOccurredAt`: 사유가 확정된 시각
- `reasonActorUserId`: 직접 조치한 사용자/운영자 ID

### 56.4 코드 네이밍 규칙
- 상태 전환 사유: `listing_*`, `reservation_*`, `completion_*`
- 정책/제재 사유: `policy_*`, `moderation_*`, `safety_*`
- 입력/검증 오류: `validation_*`
- 자동화/배치 오류: `automation_*`, `delivery_*`
- 사용자 경고/가이드: `hint_*`, `warning_*`

예시:
- `listing_closed_by_seller`
- `reservation_expired_no_confirmation`
- `moderation_hidden_suspected_fraud`
- `policy_blocked_contact_exchange`
- `automation_retry_exhausted`

### 56.5 객체별 대표 사유 코드 후보
#### 56.5.1 Listing
| 코드 | 의미 | 사용자 공개 여부 | 발생 예시 |
|---|---|---|---|
| `listing_created` | 신규 등록 | 내부/일부 공개 | 매물 최초 생성 |
| `listing_updated_core_fields` | 핵심 정보 수정 | 내부 | 가격/수량/서버 수정 |
| `listing_price_changed` | 가격 변경 | 공개 가능 | 가격 인하/인상 |
| `listing_closed_by_author` | 작성자 종료 | 공개 가능 | 판매 완료 외 단순 종료 |
| `listing_auto_expired` | 자동 만료 | 공개 가능 | 장기 비활성 매물 만료 |
| `listing_hidden_policy_violation` | 정책 위반 숨김 | 축약 공개 | 금칙어/허위매물 |
| `listing_hidden_suspected_fraud` | 사기 의심 숨김 | 비공개/축약 공개 | 반복 신고 임계치 초과 |
| `listing_restored_after_review` | 운영 복구 | 공개 가능 | 오탐 해제 |
| `listing_reopened_after_reservation_cancel` | 예약 취소 후 재오픈 | 공개 가능 | 예약 상대 이탈 |

#### 56.5.2 Chat / Message
| 코드 | 의미 | 사용자 공개 여부 | 발생 예시 |
|---|---|---|---|
| `chat_opened_from_listing` | 매물 기반 채팅 시작 | 내부 | 상세 CTA 진입 |
| `chat_blocked_by_user_relation` | 차단 관계로 채팅 차단 | 공개 가능 | 상호 차단 상태 |
| `message_masked_abusive_language` | 욕설/혐오 표현 마스킹 | 축약 공개 | 자동/운영 마스킹 |
| `message_blocked_contact_exchange` | 외부 연락처 교환 차단 | 공개 가능 | 전화번호/오픈채팅 링크 감지 |
| `chat_locked_by_report` | 신고로 채팅 잠금 | 축약 공개 | 분쟁 조사 중 |
| `chat_closed_listing_unavailable` | 매물 종결로 신규 입력 제한 | 공개 가능 | 완료/취소 후 읽기 전용 |

#### 56.5.3 Reservation
| 코드 | 의미 | 사용자 공개 여부 | 발생 예시 |
|---|---|---|---|
| `reservation_proposed` | 예약 제안 생성 | 공개 | 시간/장소 제안 |
| `reservation_counter_proposed` | 상대 수정 제안 | 공개 | 시간 변경 |
| `reservation_confirmed_both` | 상호 확정 | 공개 | 양측 승인 완료 |
| `reservation_cancelled_by_buyer` | 구매자 취소 | 공개 | 개인 사정 취소 |
| `reservation_cancelled_by_seller` | 판매자 취소 | 공개 | 물품 소진/변심 |
| `reservation_expired_no_confirmation` | 미확정 만료 | 공개 | 응답 없음 |
| `reservation_expired_no_show` | 노쇼 추정 만료 | 축약 공개 | 약속 시각 이후 상호 미확인 |
| `reservation_forced_cancelled_moderation` | 운영 강제 취소 | 축약 공개 | 정책 위반/계정 제재 |

#### 56.5.4 Trade Completion / Review
| 코드 | 의미 | 사용자 공개 여부 | 발생 예시 |
|---|---|---|---|
| `completion_requested_by_seller` | 판매자 완료 요청 | 공개 | 거래 후 완료 요청 |
| `completion_confirmed_both` | 상호 완료 확정 | 공개 | 양측 완료 확인 |
| `completion_auto_confirmed_timeout` | 자동 완료 확정 | 공개 | 기한 내 이의 없음 |
| `completion_disputed_counterparty` | 상대 이의 제기 | 공개 | 완료 거부/분쟁 전환 |
| `review_hidden_retaliation_risk` | 보복성 후기 숨김 | 비공개/축약 공개 | 운영 판단 |
| `review_removed_policy_violation` | 후기 정책 위반 비노출 | 축약 공개 | 욕설/개인정보 포함 |

#### 56.5.5 Report / Moderation / User
| 코드 | 의미 | 사용자 공개 여부 | 발생 예시 |
|---|---|---|---|
| `report_reason_fake_listing` | 허위매물 신고 | 내부 | 실물 없는 매물 |
| `report_reason_no_show` | 노쇼 신고 | 일부 공개 | 약속 불이행 |
| `report_reason_abusive_language` | 욕설/괴롭힘 신고 | 내부 | 채팅 폭언 |
| `moderation_warned_first_offense` | 1차 경고 | 공개 가능 | 경미 위반 |
| `moderation_restricted_chat` | 채팅 제한 | 공개 가능 | 반복 위반 |
| `moderation_suspended_account` | 계정 정지 | 공개 가능 | 중대 위반 |
| `user_deactivated_self_requested` | 자진 탈퇴 | 공개 가능 | 회원 탈퇴 |
| `user_deactivated_policy_enforced` | 정책상 비활성화 | 축약 공개 | 중대 사기/악성 행위 |

### 56.6 공개 사유와 내부 사유 분리 원칙
- 내부 코드가 `listing_hidden_suspected_fraud_ring_pattern`처럼 세분화되더라도 공개 사유는 `policy_review_required` 또는 `운영 검토 중` 수준으로 축약 가능해야 한다.
- 신고자/피신고자/제3자 각각에 노출 가능한 사유 수준이 달라야 한다.
- 안전/악용 방지를 위해 탐지 로직, 임계치, 세부 신호는 사용자 응답에 노출하지 않는다.

### 56.7 API/화면 반영 원칙
- 상세/목록/채팅 응답은 필요 시 아래 구조를 포함한다.
```json
{
  "viewerContext": {
    "warningBanners": [
      {
        "code": "policy_review_required",
        "severity": "warning",
        "message": "현재 운영 검토 중인 항목이 있어 일부 기능이 제한됩니다."
      }
    ]
  },
  "availableActions": ["view", "report"],
  "stateReason": {
    "reasonCode": "listing_hidden_policy_violation",
    "publicReasonCode": "policy_review_required"
  }
}
```
- 화면은 `reasonCode`의 전체 분기보다 `publicReasonCode`, `availableActions`, `warningBanners`를 우선 사용한다.
- 운영 화면은 축약 코드뿐 아니라 원본 코드, 발생 시각, 조치자, 연관 신고 ID를 함께 보여야 한다.

### 56.8 분석/감사로그 연계
- 모든 상태 전환 이벤트는 `entityType`, `entityId`, `fromState`, `toState`, `reasonCode`, `actorType`, `actorId`, `sourceSurface`를 남겨야 한다.
- KPI 산출 시 `cancelled` 건수만 보지 말고 `reservation_cancelled_by_buyer`, `reservation_cancelled_by_seller`, `reservation_expired_no_show`를 분리 집계해야 운영 개선 포인트가 보인다.
- 정책 탐지 모델/룰 변경 전후 비교를 위해 `reasonVersion`을 로그에 남긴다.

### 56.9 버전 관리 원칙
- PRD 기준 코드 사전 버전은 문서 버전과 별도로 `reason taxonomy v0.x`처럼 관리 가능하다.
- API/DB 반영 후 제거가 어려운 코드는 rename보다 deprecated + alias 매핑을 권장한다.
- 운영툴/분석/알림 템플릿이 참조하는 코드는 분기마다 동기화 점검 체크리스트가 필요하다.

### 56.10 후속 파생 문서 포인트
- DB 스키마: enum vs lookup table 결정, 다국어 라벨 테이블 여부
- API 명세: `publicReasonCode`, `reasonSource`, `warningBanners` 필드 공통화
- 운영정책: 코드별 기본 제재 수위, 공개 범위, 이의제기 가능 여부
- QA 체크리스트: 코드별 화면 문구/CTA 차단/복구 시나리오 검증

## 67. 채팅 메시지 모델 / 전달 상태 / 시스템 이벤트 계약(초안)
### 67.1 목표
- 채팅을 단순 text stream이 아니라 **거래 실행 timeline**으로 정의한다.
- 일반 메시지, 예약 카드, 상태 변경, 완료 요청, 분쟁 안내를 모두 같은 타임라인 위에서 다루되, 렌더링/권한/분석은 타입별로 분리 가능해야 한다.
- 모바일 앱이 실시간 소켓이든 폴링이든 같은 메시지/읽음 계약을 사용하도록 데이터 모델을 정리한다.

### 67.2 메시지 타입 체계
| messageType | 의미 | 발신 주체 | 대표 UI |
|---|---|---|---|
| `text` | 일반 텍스트 메시지 | 사용자 | 말풍선 |
| `system_event` | 상태 변경/정책/자동화 로그 | 시스템/운영 | 회색 시스템 라인 |
| `reservation_card` | 예약 제안/확정/변경 카드 | 사용자+시스템 조합 | 카드 블록 |
| `completion_card` | 완료 요청/확정/자동확정/이의 제기 | 사용자+시스템 조합 | 거래완료 카드 |
| `safety_notice` | 외부연락처 경고, 고위험 안내 | 시스템 | 상단/인라인 배너 |
| `image` | 이미지 첨부 | 사용자 | 이미지 말풍선 |
| `quick_template` | 빠른문구 기반 전송 | 사용자 | text와 동일, metadata에 출처 표시 |
| `admin_notice` | 운영 안내/잠금/자료 요청 | 운영/시스템 | 강조 시스템 카드 |

원칙:
- `reservation_card`, `completion_card`는 단순 문자열이 아니라 관련 객체 ID를 metadata로 연결해야 한다.
- 사용자에게는 하나의 타임라인처럼 보여도, 백엔드는 `messageType`별 validation과 권한을 다르게 가져야 한다.
- `system_event`와 `admin_notice`는 사용자가 임의 생성할 수 없다.

### 67.3 메시지 레코드 필수 필드 보강
`ChatMessage` 또는 동등 구조는 아래 필드를 추가 검토한다.

| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `messageId` | 필수 | 메시지 식별자 |
| `chatRoomId` | 필수 | 채팅방 식별자 |
| `messageType` | 필수 | 렌더링/권한 기준 |
| `senderUserId` | 선택 | 시스템 메시지는 null 가능 |
| `senderType` | 필수 | `user` / `system` / `moderator` / `automation` |
| `clientMessageId` | 선택 | 모바일 낙관적 전송 dedupe 용도 |
| `bodyText` | 선택 | 일반 본문 |
| `attachmentsJson` | 선택 | 이미지/증빙/미디어 메타데이터 |
| `metadataJson` | 선택 | 예약/완료/정책 객체 연결 정보 |
| `sentAt` | 필수 | 서버 수신/정렬 시각 |
| `editedAt` | 선택 | 수정 시각 |
| `deletedAt` | 선택 | 소프트 삭제 |
| `maskedAt` | 선택 | 민감정보 마스킹 시각 |
| `deliveryStatus` | 선택 | 클라이언트 projection 용 `sending` / `sent` / `failed` 등 |

원칙:
- 서버 영속 저장소의 기준 상태는 `sentAt`이며, `deliveryStatus`는 주로 클라이언트 projection/동기화 보조로 본다.
- `clientMessageId`를 통해 네트워크 재시도 시 중복 메시지 생성을 줄여야 한다.
- 본문 수정은 MVP에서는 비활성 또는 매우 짧은 유예 시간만 허용하는 것이 안전하다.

### 67.4 타임라인 렌더링 원칙
- 일반 텍스트와 시스템 이벤트는 시각적으로 분리해야 한다.
- 예약/완료/분쟁 관련 이벤트는 단순 시스템 한 줄이 아니라 **행동 가능한 카드**로 렌더링해야 한다.
- 같은 상태 변화가 채팅과 내 거래에 모두 보일 때, 원천 이벤트는 하나여야 하며 문구만 surface별로 달라질 수 있다.
- `deleted` 또는 `masked` 메시지는 원문 제거 대신 `삭제된 메시지`, `정책상 일부 가림` 같은 상태 렌더링을 우선한다.

### 67.5 읽음(Read) 모델 기본안
#### 67.5.1 목적
- 미읽음 배지, 푸시 억제, 내 거래 우선순위가 같은 읽음 원천 데이터를 공유해야 한다.

#### 67.5.2 필드/모델 후보
| 필드 | 설명 |
|---|---|
| `lastReadMessageId` | 사용자가 마지막으로 읽은 메시지 |
| `lastReadAt` | 읽음 처리 시각 |
| `unreadCount` | projection/counter cache |
| `lastSeenAt` | 채팅방 실제 열람 시각 |

권장 원칙:
- 사용자별 읽음 포인터는 `ChatRoomParticipantState` 같은 별도 구조로 관리하는 것이 안전하다.
- 단순 `unreadCountForSeller`, `unreadCountForBuyer` 캐시만 두면 multi-device 동기화/부분 복구가 어렵다.
- 푸시 억제는 `현재 화면 열람 중인지` + `lastSeenAt` + 최신 메시지 시각을 함께 고려해야 한다.

### 67.6 전송/실패/재시도 정책
| 상태 | 의미 | 사용자 노출 |
|---|---|---|
| `sending` | 클라이언트 로컬 전송 중 | 회색/스피너 |
| `sent` | 서버 저장 완료 | 일반 표시 |
| `failed_retryable` | 네트워크/일시 오류 | 재시도 CTA |
| `blocked_policy` | 민감정보/제재 등 정책 차단 | 경고/수정 유도 |
| `blocked_room_state` | 채팅 잠금/종결 상태로 차단 | 배너 + 입력창 제한 |

원칙:
- 채팅 입력 UX는 실패 메시지를 조용히 삼키지 말고, 사용자가 수정/재시도/삭제를 선택할 수 있어야 한다.
- `blocked_policy`는 가능한 한 전송 전(pre-flight) 경고를 우선하고, 서버 최종 검증으로 재확인한다.
- 낙관적 UI를 쓰더라도 서버가 거절한 메시지는 타임라인에 영구 성공처럼 남아선 안 된다.

### 67.7 첨부/이미지 정책(채팅)
- MVP 기본안은 `text + image`까지를 우선 지원한다.
- 이미지 첨부 규칙 후보:
  - 1회 최대 5장
  - 파일당 최대 용량 제한 필요
  - 지원 포맷: jpg/png/webp 우선
  - 원본과 썸네일을 분리 저장
- 금지/제한:
  - 계좌번호/연락처/실명 등이 이미지에 포함될 수 있으므로 OCR 기반 사전검사 또는 사후 마스킹 검토 필요
  - 혐오/음란/무관 이미지 업로드 금지
- 채팅 내 이미지 업로드는 `pending_trade` 이후 허용 범위를 완화할지 여부가 오픈 질문이다.

### 67.8 민감정보/외부연락처 메시지 처리 상세
| 상황 | 기본 처리 | 사용자 경험 |
|---|---|---|
| 전화번호/메신저 ID 감지 | 경고 후 수정 유도 또는 차단 | “플랫폼 밖 연락처 공유는 제한될 수 있어요” |
| 계좌번호 감지 | 강한 경고 또는 차단 우선 | “계좌정보 공유는 지원하지 않아요” |
| 정확 주소/좌석번호 등 오프라인 민감 위치 | 예약 확정 전 차단/경고 | “정확한 장소는 예약 확정 후 공유해 주세요” |
| 욕설/혐오 표현 | 경고, 마스킹, 운영 로그 | “정책 위반 표현이 포함되어 있어요” |

원칙:
- 같은 정책이라도 `경고만`, `부분 마스킹`, `전송 차단` 단계를 구분해야 한다.
- 어떤 메시지가 왜 막혔는지 사용자가 이해 가능해야 하나, 탐지 규칙 상세는 과도하게 노출하지 않는다.

### 67.9 빠른문구(Quick Reply / Template) 정책
목적: 모바일 거래 속도를 높이되 스팸성 복붙을 줄인다.

후보 템플릿:
- `지금 거래 가능하신가요?`
- `오늘 저녁 거래 가능해요.`
- `기란 마을 창고 앞 가능해요.`
- `도착하면 채팅드릴게요.`
- `거래 완료했습니다. 확인 부탁드려요.`

원칙:
- 빠른문구는 사용자 수정 가능 텍스트로 삽입하되, 전송 후에는 일반 메시지로 저장한다.
- 지나친 자동 전송/대량 복붙 방지를 위해 짧은 시간 동일 템플릿 반복 사용은 rate limit 대상이 될 수 있다.
- 예약/완료 카드에 종속된 빠른문구는 context-aware 추천으로 노출할 수 있다.

### 67.10 채팅 목록 projection 기본안
후보 엔드포인트: `GET /chats`

응답 item 예시:
```json
{
  "chatRoomId": "chat_123",
  "listingId": "listing_123",
  "tradeThreadId": "tt_123",
  "chatStatus": "reservation_confirmed",
  "counterparty": {
    "userId": "user_456",
    "nickname": "린클상인",
    "trustBadge": "거래경험많음"
  },
  "listing": {
    "title": "+7 장검 팝니다",
    "thumbnailUrl": "https://..."
  },
  "lastMessage": {
    "messageType": "reservation_card",
    "previewText": "오늘 21:00 기란 마을로 예약이 확정됐어요",
    "sentAt": "2026-03-13T07:20:00+09:00"
  },
  "unreadCount": 2,
  "availableActions": ["open_chat", "open_trade_thread", "mute_chat"]
}
```

원칙:
- 채팅 목록은 raw 마지막 메시지뿐 아니라 `previewText`, `messageType`, `tradeThreadId`를 함께 제공해 내 거래/채팅 탭 이동을 단순화한다.
- 목록 정렬은 `lastMessageAt`와 별도로 `액션 필요 거래 여부`를 섞을지 제품 정책으로 확정 필요하다.

### 67.11 채팅 상세 응답/동기화 계약 기본안
후보 엔드포인트:
- `GET /chats/{chatRoomId}`
- `GET /chats/{chatRoomId}/messages?cursor=...`
- `POST /chats/{chatRoomId}/messages`
- `POST /chats/{chatRoomId}/read`

응답 설계 원칙:
- 메시지는 cursor pagination을 기본안으로 한다.
- 상세 응답에는 현재 활성 reservation/completion/dispute 요약을 함께 포함해 상단 카드 렌더링에 활용할 수 있어야 한다.
- 실시간 소켓 도입 전에도 폴링으로 동기화 가능한 `sinceMessageId` 또는 `updatedAfter` 기반 계약을 검토한다.
- 읽음 API는 단일 messageId 기준 ack를 우선안으로 한다.

### 67.12 시스템 이벤트 생성 규칙
아래 이벤트는 일반 메시지와 별개로 반드시 타임라인에 남아야 한다.
- 채팅방 생성
- 예약 제안/수정/확정/취소/만료
- 장소 변경 및 상대 재확인 필요
- 매물 상태 변경
- 완료 요청/확정/자동확정/이의 제기
- 신고 잠금/운영 자료 요청/잠금 해제

원칙:
- 이벤트 생성 실패로 인해 상태만 바뀌고 타임라인이 비어 보이는 상황을 피해야 한다.
- 시스템 이벤트는 사용자가 삭제/수정할 수 없다.
- 알림함/내 거래 카드 문구는 가능하면 이 시스템 이벤트 payload를 재활용한다.

### 67.13 분석 이벤트 후보
- `chat_list_view`
- `chat_room_view`
- `chat_message_send_attempt`
- `chat_message_send_success`
- `chat_message_send_failed`
- `chat_message_policy_blocked`
- `chat_read_ack`
- `chat_quick_template_used`
- `chat_attachment_uploaded`
- `chat_system_event_rendered`

핵심 분석 포인트:
- 첫 문의 후 첫 응답까지 걸리는 시간
- 예약 카드 노출 후 예약 응답 전환율
- 정책 차단 메시지 유형별 발생률
- 빠른문구 사용이 거래 완료율/응답속도에 미치는 영향

### 67.14 후속 파생 문서 포인트
- 화면명세:
  - 말풍선/시스템 이벤트/카드 렌더링 규격
  - 입력창 disabled 사유와 재시도 UX
- DB 스키마:
  - `ChatRoomParticipantState`, `clientMessageId`, `senderType`, `maskedAt` 반영 여부
- API 명세:
  - read ack, cursor sync, optimistic send/dedupe 계약
- 운영정책:
  - 메시지 마스킹 기준, 이미지 OCR 처리, admin notice 생성 기준
- QA:
  - 다중 기기 읽음 동기화, 정책 차단 메시지 재시도, 시스템 이벤트 누락 여부 검증


## 68. 채팅 첨부/미디어 저장·노출 정책(초안)
### 68.1 목표
- 채팅 첨부를 단순 파일 업로드가 아니라 **거래 실행 증빙 + 대화 보조 수단**으로 정의한다.
- MVP에서는 이미지 중심으로 시작하되, 이후 분쟁 증빙/운영 검토까지 확장 가능한 저장 구조를 마련한다.
- 모바일 업로드 UX, 저장 비용, 안전 검수, 신고/분쟁 대응이 서로 충돌하지 않도록 기준을 명문화한다.

### 68.2 MVP 첨부 범위
| 첨부 유형 | MVP 지원 | 주요 용도 | 비고 |
|---|---|---|---|
| `image` | 지원 | 아이템 스크린샷, 거래 증빙, 장소 안내 캡처 | 핵심 |
| `multi_image` | 제한 지원 | 한 번에 1~5장 | 카드형 묶음 노출 후보 |
| `file` | 미지원 또는 운영 전용 | 로그/문서/압축파일 | 악용 리스크 높음 |
| `video` | 미지원 | 사기 증빙/화면 녹화 | 저장비용 커서 Post-MVP |
| `voice` | 미지원 | 빠른 음성 협의 | 초기 복잡도 높음 |

원칙:
- MVP는 일반 채팅 첨부를 **이미지**로 제한한다.
- 분쟁/운영 제출은 장기적으로 file/video를 검토할 수 있으나, 일반 채팅 surface에는 열지 않는 기본안이 안전하다.
- 첨부 지원 여부는 `messageType=image` 또는 `text + attachments` 방식 중 하나로 일관되게 정해야 하며, MVP에서는 `image` 전용 타입을 우선 검토한다.

### 68.3 첨부 사용 시나리오 구분
| 시나리오 | 허용 범위 | 저장/노출 특성 |
|---|---|---|
| 일반 거래 대화 | 이미지 1~5장 | 참여자 간 노출, 타임라인 표시 |
| 예약/장소 안내 | 이미지 1장 권장 | 지도/위치 스크린샷 수준, exact 개인정보 주의 |
| 거래 완료 증빙 | 이미지 허용 | completion/dispute와 연결 가능 |
| 분쟁 소명 | 별도 제출 첨부 후보 | staff 전용 또는 제한 공유 |
| 운영 증빙 열람 | 원본 접근 가능 | 감사 로그/권한 통제 필수 |

원칙:
- 일반 채팅 첨부와 분쟁 증빙 첨부는 같은 저장소를 쓰더라도 **권한/노출 정책은 분리**해야 한다.
- 사용자가 채팅에 올린 이미지를 운영 분쟁 근거로 참조할 수는 있으나, 분쟁 폼 제출 첨부는 별도 객체 연결이 바람직하다.

### 68.4 첨부 메시지 데이터 모델 후보
`ChatMessage.attachmentsJson` 또는 별도 `ChatAttachment` 엔티티 후보:

| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `attachmentId` | 필수 | 첨부 식별자 |
| `messageId` | 필수 | 소속 메시지 |
| `attachmentType` | 필수 | `image` |
| `storageKey` | 필수 | 원본 저장 경로 |
| `thumbnailStorageKey` | 선택 | 목록/채팅용 썸네일 |
| `mimeType` | 필수 | `image/jpeg`, `image/png`, `image/webp` 등 |
| `fileSizeBytes` | 필수 | 용량 |
| `width` / `height` | 선택 | 렌더링 최적화 |
| `blurhash` 또는 placeholder | 선택 | 로딩 UX 개선 |
| `moderationStatus` | 필수 | `pending_scan` / `clean` / `masked` / `blocked` |
| `visibilityScope` | 필수 | `participants`, `staff_only`, `masked_preview` |
| `createdAt` | 필수 | 생성 시각 |
| `deletedAt` | 선택 | 소프트 삭제 |

원칙:
- 첨부 메타데이터는 메시지 본문과 분리 저장하는 것이 운영/스토리지/변환 관리에 유리하다.
- 썸네일, 원본, 검수 상태는 별도 필드로 관리해 모바일 렌더링과 정책 차단을 분리한다.

### 68.5 업로드 수명주기 정책
1. **임시 업로드 생성**
   - 사용자가 이미지를 선택하면 먼저 임시 업로드 토큰 또는 presigned 업로드 흐름을 사용한다.
2. **보안/형식 검증**
   - MIME, 확장자, 용량, 이미지 decode 가능 여부를 검사한다.
3. **메시지 확정 연결**
   - `POST /chats/{chatRoomId}/messages` 호출 시 attachment reference를 함께 전송해 실제 메시지와 연결한다.
4. **검수/후처리**
   - 썸네일 생성, OCR/민감정보 탐지, 정책 스캔을 비동기로 수행할 수 있다.
5. **미참조 정리**
   - 임시 업로드 후 메시지 전송에 사용되지 않은 파일은 TTL 기준 자동 삭제한다.

기본 원칙:
- 메시지 확정 전 업로드 파일은 `orphan` 상태로 간주하고 짧은 TTL(예: 24시간, 가정) 후 삭제한다.
- 메시지 삭제와 파일 물리 삭제는 분리할 수 있으며, 분쟁/감사 목적 보관 정책을 우선한다.

### 68.6 첨부 validation / 제한 규칙
- 허용 포맷: `jpg`, `jpeg`, `png`, `webp` 우선
- GIF/애니메이션은 MVP 비권장 또는 정지 이미지로만 처리
- 파일당 최대 용량: 예: 10MB 이하(가정)
- 메시지 1건당 첨부 수: 최대 5개(가정)
- 짧은 시간 대량 이미지 업로드는 rate limit 또는 스팸 탐지 대상
- 저해상도/깨진 이미지, 비이미지 파일 위장 업로드는 차단
- 이미지 내 전화번호, 계좌번호, 오픈채팅 QR/링크 등은 OCR 기반 경고/차단 후보

### 68.7 민감정보/정책 검수 원칙
- 채팅 첨부도 본문과 동일하게 외부 연락처, 계좌번호, 실명 노출 정책을 적용한다.
- 검수 단계는 최소 아래를 고려한다.
  1. 악성 파일/포맷 검증
  2. OCR 기반 민감정보 감지
  3. 성인물/무관 이미지/광고성 이미지 탐지
- 위반 감지 시 처리 옵션:
  - 전송 전 차단
  - 전송은 허용하되 상대방에는 마스킹/경고 표시
  - 운영 로그 적재 후 재검토

원칙:
- 자동 차단은 명백한 정책 위반 패턴에 우선 적용하고, 애매한 경우 `masked_preview` 또는 운영 검토 대기로 완충한다.
- 사용자에게는 `이미지에 연락처 정보가 포함되어 전송이 제한되었습니다` 같은 공개 사유만 노출한다.

### 68.8 화면 노출 규칙
#### 채팅 타임라인
- 첨부 이미지는 본문 없이도 단독 전송 가능
- 다중 이미지는 묶음 grid 또는 carousel 형태를 검토한다.
- 정책 마스킹 시 원본 대신 placeholder와 사유 배지를 노출한다.
- `report_locked` 상태에서는 새 첨부 업로드를 금지한다.

#### 채팅 목록/푸시/알림함
- 목록 preview는 `사진을 보냈어요` 같은 텍스트 요약만 사용한다.
- 푸시에는 이미지 썸네일을 기본 포함하지 않고, 잠금화면 민감도 리스크를 낮춘다.
- 알림함 deep link는 해당 채팅 메시지 앵커로 이동할 수 있으면 좋다.

#### 분쟁/운영 화면
- 운영자는 첨부 이미지의 검수 상태, 원본/썸네일, OCR 탐지 결과, 신고 연결 여부를 함께 볼 수 있어야 한다.
- 민감정보 원본 열람은 기존 백오피스 권한 정책을 따른다.

### 68.9 삭제/보관/복구 정책
- 사용자가 메시지를 soft delete하더라도 첨부 원본은 분쟁 대응 보관 기간 동안 유지 가능해야 한다.
- 공개/참여자 화면에서는 삭제된 메시지의 첨부를 즉시 비노출 처리한다.
- 운영 숨김된 첨부는 일반 참여자에게는 `정책상 비노출`로 보이고, 원본은 staff 권한 하에만 유지한다.
- 탈퇴 사용자 첨부도 거래/신고 맥락 보존을 위해 직접 hard delete보다 참조 차단 + 보관 만료 정책을 우선한다.

### 68.10 API 후보
#### 업로드 준비
- `POST /uploads/chat-images`
- 응답: 업로드 토큰, presigned URL, 임시 attachmentId

#### 메시지 전송
- `POST /chats/{chatRoomId}/messages`
```json
{
  "messageType": "image",
  "clientMessageId": "cmsg_123",
  "attachments": [
    {"attachmentId": "att_123"}
  ]
}
```

#### 응답 projection 후보
```json
{
  "messageId": "msg_123",
  "messageType": "image",
  "deliveryStatus": "accepted",
  "attachments": [
    {
      "attachmentId": "att_123",
      "attachmentType": "image",
      "thumbnailUrl": "https://...",
      "originalUrl": "https://...",
      "moderationStatus": "clean"
    }
  ]
}
```

### 68.11 스토리지/보안 시사점
- 공개 CDN 캐시를 쓰더라도 원본 URL은 짧은 만료 시간 또는 서명 URL 정책을 검토해야 한다.
- 썸네일은 참여자 경험을 위해 캐시 가능하지만, 권한 없는 직접 접근 방지가 필요하다.
- 원본 저장소와 썸네일/변환 저장소를 분리하면 정책 마스킹/재생성에 유리하다.
- 이미지 메타데이터(EXIF 위치 정보 등)는 업로드 시 제거하는 기본안을 권장한다.

### 68.12 분석 / QA 포인트
분석 이벤트 후보:
- `chat_attachment_upload_started`
- `chat_attachment_upload_completed`
- `chat_attachment_send_success`
- `chat_attachment_send_blocked`
- `chat_attachment_open`
- `chat_attachment_reported`

QA 체크포인트:
- 업로드 성공 후 메시지 전송 실패 시 orphan attachment 정리가 되는가
- 같은 첨부를 네트워크 재시도로 중복 메시지 생성 없이 처리하는가
- OCR/정책 차단된 이미지가 상대방/푸시에 노출되지 않는가
- 삭제된 메시지의 첨부가 일반 사용자 화면에서 즉시 사라지는가
- 운영자가 정책 숨김 후에도 감사 추적과 복구가 가능한가

## 69. 채팅 실시간 전송 방식 / Presence / Participant State 정책(초안)
### 69.1 목표
- 모바일 거래 UX에 필요한 체감 실시간성을 확보하되, 초기 MVP에서 과도한 소켓 복잡도를 강제하지 않는다.
- 읽음/미읽음/푸시/액션 필요 상태가 전송 방식에 따라 다르게 해석되지 않도록, **전송 수단보다 동기화 계약**을 우선한다.
- 이후 WebSocket, SSE, 폴링을 혼합하더라도 같은 `event stream + participant state` 모델을 재사용할 수 있게 한다.

### 69.2 전송 방식 결정 원칙
우선순위는 아래와 같다.
1. **정합성**: 이벤트 누락/중복 없이 복구 가능해야 한다.
2. **모바일 배터리/연결 안정성**: 백그라운드 전환, 네트워크 변경에 강해야 한다.
3. **구현 복잡도 관리**: MVP에서 운영/디버깅 가능한 수준이어야 한다.
4. **확장성**: 추후 typing/presence/다기기 동기화로 확장 가능해야 한다.

권장 기본안:
- **MVP는 `pull + server push 보조` 혼합 전략**을 채택한다.
- 기본 원천 동기화는 `GET /chats/{chatRoomId}/events?afterEventId=...` 같은 pull API로 보장한다.
- 가능하면 foreground 상태에서만 SSE 또는 WebSocket을 붙여 신규 이벤트 체감 지연을 줄인다.
- push/socket 실패 시에도 클라이언트는 pull backfill만으로 완전 복구 가능해야 한다.

### 69.3 권장 MVP 아키텍처
| 상황 | 권장 방식 | 이유 |
|---|---|---|
| 앱 foreground에서 채팅 상세 열람 중 | SSE 우선 또는 WebSocket 후보 | 서버→클라이언트 단방향 신규 이벤트 전달이면 충분한 경우가 많음 |
| 앱 foreground에서 채팅 목록/내 거래 화면 | 짧은 간격 폴링 또는 이벤트 요약 feed | 모든 방에 상시 socket 유지 비용 절감 |
| 앱 background/종료 상태 | 푸시 알림 + 화면 재진입 시 pull sync | 모바일 OS 제약 대응 |
| 실시간 연결 실패/재연결 직후 | `afterEventId` backfill pull | 누락 복구 핵심 |
| 운영/관리자 화면 | 폴링 우선 | 사용자용 실시간보다 안정성 우선 |

원칙:
- WebSocket이 반드시 필요한 것은 아니다. 초기 거래 채팅은 **신규 이벤트 수신 + 읽음 반영 + backfill**만 안정적이면 된다.
- 양방향 소켓을 도입하더라도 쓰기 성공의 기준은 소켓 ack가 아니라 서버 저장 성공 응답이다.
- SSE는 구현 단순성이 장점이 있으나, 양방향 interactive signal이 늘어나면 WebSocket 전환 가능성을 열어둔다.

### 69.4 실시간 연결 상태 모델
클라이언트는 최소 아래 연결 상태를 가져야 한다.

| connectionState | 의미 | 사용자 노출 |
|---|---|---|
| `live` | 신규 이벤트를 실시간 수신 중 | 별도 노출 불필요 |
| `reconnecting` | 연결 복구 중 | 선택적 약한 배너 |
| `degraded_polling` | 실시간 연결 없이 폴링만 동작 | 노출 생략 가능 |
| `offline` | 네트워크 없음 | 오프라인 배너 + 재시도 |
| `stale_sync` | 연결은 있으나 gap 가능성 존재 | 재동기화 수행 |

원칙:
- 연결 상태는 UX 보조 정보일 뿐, 거래 상태와 섞어 쓰지 않는다.
- 실시간 연결이 끊겨도 메시지 전송/읽음/예약 응답 기능 전체를 막지 않고, 가능한 범위에서 pull 기반으로 지속해야 한다.

### 69.5 SSE vs WebSocket 선택 기준(가정)
#### SSE 우선이 적합한 이유
- 채팅/내거래는 서버 이벤트 fanout 비중이 높고, 클라이언트→서버 쓰기는 기존 HTTP API로 충분하다.
- 인증/재연결/프록시/인프라 디버깅이 WebSocket보다 단순할 수 있다.
- MVP 범위의 typing/presence를 제외하면 대부분 단방향 이벤트 푸시로 해결 가능하다.

#### WebSocket이 필요해지는 조건
- typing/presence를 적극적으로 실시간 표시해야 함
- 다수 채팅방의 동시 상태 스트림을 하나의 연결로 통합해야 함
- 향후 통화/고빈도 상호작용 같은 richer realtime 기능이 필요함

권장 결론:
- **MVP 권장안은 `HTTP write + SSE read + pull backfill`**
- 단, 인프라/플랫폼 제약이 있으면 **`HTTP write + short polling read`**로도 시작 가능하며, 이 경우 동일 이벤트 커서 계약을 유지해야 한다.

### 69.6 폴링 정책 기본안
| 화면 | 폴링 기본안 | 조건 |
|---|---|---|
| 채팅 상세 | 3~5초 간격(가정) | 실시간 연결 미사용 시 |
| 채팅 목록 | 10~20초 간격 또는 앱 복귀 시 | foreground 한정 |
| 내 거래 | 15~30초 간격 또는 중요 이벤트 후 즉시 refresh | 액션 필요 상태 반영 |
| 알림함 | 화면 진입/당겨서 새로고침 중심 | 실시간 강제 불필요 |

최적화 원칙:
- 사용자가 해당 화면을 보고 있을 때만 폴링한다.
- 입력 중/예약 응답 대기/완료 대기 상황에서는 일시적으로 poll interval을 줄일 수 있다.
- background에서는 폴링을 유지하지 않고 푸시에 의존한다.

### 69.7 Typing / Presence MVP 정책
#### typing indicator
- **MVP 기본안: 미포함 또는 매우 제한적 포함**
- 이유:
  - 거래 채팅은 일반 메신저보다 예약/완료 카드 같은 구조화 액션 비중이 큼
  - typing 상태는 저장되지 않는 ephemeral signal이라 구현/복구 대비 제품 가치가 낮을 수 있음
  - 푸시/읽음/실시간 연결 안정화가 typing보다 우선순위가 높음

Post-MVP 포함 시 원칙:
- DB 저장 금지
- SSE/WebSocket ephemeral event로만 전달
- 일정 시간(예: 5초) 후 자동 만료
- 차단/잠금 채팅방에는 비활성

#### online/presence indicator
- **MVP 기본안: 정확한 `온라인` 상태 미노출**
- 대신 아래의 구간형 상태를 사용한다.
  - `방금 활동함`
  - `오늘 활동`
  - `최근 활동`
- 이유:
  - 정확한 접속 여부는 사생활/압박/스토킹 리스크가 있음
  - 거래 신뢰에는 절대 실시간 presence보다 응답성/최근 활동성이 더 중요함

### 69.8 Presence/Typing 확장 가드레일
- 차단 관계가 있으면 상대 typing/presence를 보여주지 않는다.
- 신고 잠금/분쟁 상태 채팅에서는 typing 비활성화가 기본안이다.
- 마지막 활동 시각은 exact timestamp보다 상대시간/구간형을 우선한다.
- 고위험 사용자 조합에서는 presence보다 안전 배너를 우선 노출한다.

### 69.9 ChatParticipantState 객체 계약(권장)
목적: 채팅방과 사용자 간 상태를 `ChatRoom` 본문에서 분리해 읽음/뮤트/액션 우선순위/동기화 상태의 원천 객체로 사용한다.

후보 필드:
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `chatRoomId` | 필수 | 채팅방 식별자 |
| `userId` | 필수 | 참여 사용자 |
| `participantRole` | 필수 | `seller` / `buyer` |
| `lastReadEventId` | 선택 | 마지막 읽은 이벤트 |
| `lastReadAt` | 선택 | 읽음 처리 시각 |
| `lastSeenAt` | 선택 | 화면 실제 열람 시각 |
| `muteLevel` | 필수 | `all`, `messages_only`, `none` 등 후보 |
| `mutedUntil` | 선택 | 임시 뮤트 만료 |
| `pushOptInEffective` | 필수 | 전역 + 채팅 설정 합성 결과 |
| `hasActionRequired` | 필수 | 이 참여자 관점 액션 필요 여부 캐시 |
| `lastNotifiedEventId` | 선택 | 마지막으로 알림 발송한 이벤트 |
| `gapDetected` | 필수 | 재동기화 필요 플래그 |
| `archivedAt` | 선택 | 개인 보관함 숨김/아카이브 시각 |
| `joinedAt` | 필수 | 채팅방 참여 시작 |

원칙:
- 미읽음 수는 이 객체에서 계산하거나 projection으로 캐시한다.
- `ChatRoom` 본문에는 공통 상태만 두고, 사용자별 상태는 `ChatParticipantState`에 둔다.
- 추후 다기기 세분화가 필요하면 `ChatDeviceState` 같은 하위 모델을 추가하되, MVP는 사용자 단위 상태를 기본으로 한다.

### 69.10 `hasActionRequired` 계산 원칙
아래 경우 true 후보:
- 상대가 보낸 미읽음 일반 메시지 존재
- 예약 제안 응답 대기
- 장소 변경 재확인 필요
- 완료 확인/이의 제기 대기
- 분쟁 추가 자료 제출 요청 대기
- 운영 경고/정책 확인 필요

아래 경우 false 후보:
- 단순 시스템 정보 로그만 존재
- 이미 확인 완료한 완료/예약 이벤트
- 뮤트되었지만 액션 필요는 없는 일반 대화

원칙:
- `hasActionRequired`는 unread와 동일하지 않다.
- 홈/내거래/푸시 우선순위는 unread보다 `hasActionRequired`를 우선 사용해야 한다.

### 69.11 API 후보
#### 연결/이벤트 조회
- `GET /chats/{chatRoomId}/events?afterEventId=...`
- `GET /me/chat-feed?afterCursor=...` (목록/내거래 요약용 확장 후보)
- `GET /chats/{chatRoomId}/stream` (SSE 후보)

#### participant state 업데이트
- `POST /chats/{chatRoomId}/read-cursor`
- `POST /chats/{chatRoomId}/mute`
- `DELETE /chats/{chatRoomId}/mute`
- `POST /chats/{chatRoomId}/archive`
- `DELETE /chats/{chatRoomId}/archive`

#### 응답 projection 예시
```json
{
  "chatRoomId": "chat_123",
  "participantState": {
    "lastReadEventId": "evt_120",
    "hasActionRequired": true,
    "muteLevel": "none",
    "gapDetected": false,
    "connectionMode": "sse"
  }
}
```

### 69.12 알림/실시간 연계 원칙
- 실시간으로 이미 이벤트를 수신하고 현재 채팅을 열람 중인 경우, 푸시는 억제한다.
- 채팅 목록만 열람 중이면 메시지 원문 푸시는 억제할 수 있으나, 예약/완료 요청 같은 T1 이벤트는 인앱 배지와 상단 상태 카드 갱신을 우선 보장한다.
- `lastNotifiedEventId`를 participant state 또는 notification log와 연결해 중복 푸시를 막는다.
- 푸시를 받은 뒤 앱이 열리면, 해당 eventId 이후 backfill을 반드시 수행해 타임라인/내거래/알림함이 일치해야 한다.

### 69.13 운영/관측성 포인트
추적 지표 후보:
- 실시간 연결 성공률 / SSE 연결 유지 시간
- 폴링 fallback 비율
- gapDetected 발생률
- read ack 지연 시간
- push 억제율 vs 실제 열람율
- action required 상태에서의 응답 완료 시간

운영 원칙:
- 실시간 연결 장애가 나도 거래 기능 전체 장애로 간주하지 않고, `degraded realtime / healthy sync` 상태를 구분 관측한다.
- 반대로 push는 정상인데 backfill이 실패하면 거래 UX는 깨질 수 있으므로, pull sync 성공률을 더 중요한 SLI로 본다.

### 69.14 후속 파생 문서 포인트
- 기술 설계:
  - SSE/WebSocket/polling 선택 근거와 재연결 시퀀스
  - event cursor / seq 키 설계
  - participant state 저장/캐시 전략
- API 명세:
  - stream endpoint, read cursor, mute/archive 계약
  - `hasActionRequired`, `gapDetected`, `connectionMode` 필드 정의
- 화면명세:
  - 연결 상태 배너, 읽음 갱신 타이밍, presence 미노출 원칙
- QA:
  - foreground 실시간 수신, background 푸시 후 복귀 sync, 네트워크 전환 중 중복/누락 검증

## 70. 계정 신뢰 신호 / 인증 레벨 / 안티어뷰즈 정책(초안)
### 70.1 목표
- 린클의 신뢰도는 단순 후기 점수 하나가 아니라 **계정 확인 수준 + 거래 평판 + 최근 위험 신호**의 조합으로 해석되어야 한다.
- 사용자에게는 거래 판단에 필요한 최소 신뢰 신호만 공개하고, 악용 방지에 필요한 리스크 신호는 운영 전용으로 분리한다.
- 가입/프로필/기기/행동 제한 정책을 명확히 해 DB 스키마, API 권한, 운영 제재, 분석 이벤트가 같은 기준을 쓰도록 한다.

### 70.2 신뢰 신호 3계층 모델
| 계층 | 의미 | 대표 데이터 | 공개 범위 |
|---|---|---|---|
| `identity_assurance` | 이 계정이 최소한 일관된 사용자인지 확인하는 계층 | 로그인 수단, 본인확인 여부, 계정 연령 | 제한 공개 또는 배지형 |
| `trade_reputation` | 실제 거래 경험과 상대 평가 | 완료 거래 수, 후기, 응답성, 노쇼율 | 회원/거래 상대에게 공개 가능 |
| `behavior_risk` | 스팸/사기/다계정/정책위반 위험 | 반복 신고, 차단 누적, 대량문의 패턴, 기기/세션 이상 | 운영 전용 |

원칙:
- 공개 프로필과 매물 카드에는 `identity_assurance`와 `trade_reputation`의 일부만 사용한다.
- `behavior_risk`는 사용자 낙인 방지를 위해 직접 점수로 공개하지 않는다.
- 운영자는 세 계층을 함께 보되, 자동 제재의 단독 근거는 `behavior_risk`만으로 삼지 않는 것을 원칙으로 한다.

### 70.3 계정 확인(Assurance) 레벨 초안
| assuranceLevel | 의미 | 획득 조건 후보 | 사용자 노출 예시 |
|---|---|---|---|
| `A0_guest` | 비회원 | 로그인 없음 | 비노출 |
| `A1_basic` | 기본 가입 계정 | 소셜/이메일/휴대폰 중 1개 | `가입 완료` |
| `A2_verified` | 기본 확인 계정 | 휴대폰 확인 또는 동등 수단 | `기본 확인됨` |
| `A3_consistent_trader` | 일정 거래 이력까지 쌓인 계정 | 완료 거래/활동 기간/제재 없음 조합 | `거래 이력 있음` |
| `A4_high_confidence` | 장기적으로 안정적 사용 패턴 확인 | 장기 활동 + 후기/분쟁 안정성 + 추가 확인 수단(정책 결정) | `안정적 거래자` |

원칙:
- MVP에서는 `A1~A3` 정도만 실제로 운영해도 충분하다.
- `A4`는 과한 인증을 강제하기보다 장기 활동 기반 신뢰 배지와 합쳐 해석하는 안을 우선 검토한다.
- assurance level은 후기를 대체하지 않으며, 신규 계정을 일괄 배척하는 낙인 UX로 사용하지 않는다.

### 70.4 가입/온보딩 기본 정책
- 가입 직후 즉시 가능한 기본 행동:
  - 목록/상세 조회
  - 프로필 기본 설정
  - 매물 초안 작성 또는 제한된 수의 매물 등록
  - 제한된 수의 채팅 시작
- 가입 직후 제한이 필요한 행동 후보:
  - 대량 매물 등록
  - 짧은 시간 내 다수 채팅 생성
  - 이미지 다량 업로드
  - 외부 연락처 포함 메시지 전송
- 닉네임 정책 기본안:
  - 최소/최대 길이 제한
  - 욕설/광고/연락처 포함 금지
  - 빈번한 닉네임 변경 쿨다운 필요(예: 7일, 가정)
- 프로필 사진/소개는 선택이지만, 신뢰 형성에 도움이 되는 선택 정보로 안내할 수 있다.

### 70.5 공개 신뢰 신호 계약
매물 카드/상세/프로필에서 공개 가능한 신호는 다음 범위를 기본안으로 둔다.

| 신호 | 목록 카드 | 상세 | 거래 상대 | 비고 |
|---|---|---|---|---|
| 가입 경과 기간 구간 | 선택 | 가능 | 가능 | exact 날짜 대신 `3개월+` 구간형 우선 |
| 완료 거래 수 구간/수치 | 요약 | 가능 | 가능 | summary/member/participant 레벨에 따라 차등 |
| 후기 요약 | 요약 | 가능 | 가능 | 보복/분쟁 중 제외 규칙 필요 |
| 응답성 배지 | 가능 | 가능 | 가능 | 최근 N일 기준 재계산 후보 |
| 기본 확인 여부 | 가능 | 가능 | 가능 | `휴대폰 확인됨` 같은 단순 표현 |
| 최근 활동 상대시간 | 가능 | 가능 | 가능 | exact timestamp 지양 |
| 제재/신고 세부 | 불가 | 불가 | 불가 | 운영 전용 |

원칙:
- 신뢰 신호는 “거래 가능성”을 높이는 정보로만 쓰고, “이 사람은 안전하다”는 보증성 카피를 사용하지 않는다.
- 운영 검토/분쟁 중인 거래는 관련 후기/배지 반영을 보수적으로 유예할 수 있어야 한다.

### 70.6 내부 리스크 신호 후보
운영/탐지 시스템이 참고할 수 있는 내부 신호 예시:
- 짧은 시간 내 다수 계정 생성 시도
- 동일 기기/네트워크 추정 집합에서 반복 신고/차단/제재 발생
- 매물 등록 직후 외부 연락처 교환 시도 빈도 높음
- 다수 상대에게 동일 복붙 문의 전송
- 예약 확정 대비 실제 완료율이 비정상적으로 낮음
- 장기 `reserved` 유지 후 반복 취소
- 신규 계정의 고가/고위험 키워드 매물 집중 등록

원칙:
- 내부 리스크 신호는 누적 점수, rule hit count, 최근성(window) 개념으로 관리하는 것이 유리하다.
- 룰 적중만으로 곧바로 영구 제재하지 않고, 먼저 rate limit/추가 확인/운영 검토 큐 적재를 우선한다.

### 70.7 행위 제한 / rate limit / cooldown 정책 초안
| 행동 | 기본 계정(A1) | 거래 이력 계정(A2~A3) | 비고 |
|---|---|---|---|
| 매물 등록 | 일/시간당 낮은 상한 | 더 높은 상한 | 카테고리/이미지 수에 따라 가중 가능 |
| 채팅 시작 | 짧은 시간당 제한 | 완화 | reserved 매물 대량문의 방지 |
| 메시지 전송 | burst 제한 필요 | 완화 | 동일문구 반복 탐지 병행 |
| 이미지 업로드 | 장당/일일 제한 | 완화 | 미디어 검수 비용 고려 |
| 닉네임 변경 | 긴 쿨다운 | 동일 | 사칭/회피 방지 |
| 후기 작성 | 거래 완료 기반만 허용 | 동일 | 별도 완화 불필요 |
| 신고 접수 | 반복 허위 신고 제한 | 동일 | 허위 신고 누적 시 추가 제한 |

설계 원칙:
- 모든 limit은 절대 숫자만이 아니라 risk-adjusted policy로 확장 가능해야 한다.
- 사용자에게는 상세 임계치를 과도하게 공개하지 않고, `잠시 후 다시 시도해 주세요` + 필요 시 고객센터 경로를 제공한다.
- 서버는 429/423 계열 코드와 함께 `retryAfterSeconds(optional)`를 반환할 수 있어야 한다.

### 70.8 세션/기기 추적 최소 기준
- 계정 보안과 안티어뷰즈를 위해 아래 최소 메타데이터 보관을 검토한다.
  - 최근 로그인 시각
  - 최근 활성 기기/세션 수
  - 로그인 수단별 연결 상태
  - 기기/브라우저 fingerprint의 직접 식별이 아닌 안정적 해시 또는 risk token
  - push token / device platform / app version
- 원칙:
  - 디바이스 추적은 개인정보 최소수집 원칙 하에 위험 탐지 목적 범위로 제한한다.
  - exact IP 장기 보관보다 risk aggregation 및 보안 감사 목적의 제한 보관을 우선 검토한다.
  - 사용자가 본인 세션을 관리/로그아웃할 수 있는 기능은 Post-MVP라도 고려 가치가 있다.

### 70.9 운영 액션과 계정 상태 연결 규칙
| accountStatus | 의미 | 기본 허용 범위 | 운영 액션 예시 |
|---|---|---|---|
| `active` | 정상 이용 | 전체 기능 | 없음 |
| `limited_listing` | 매물 등록만 제한 | 조회/채팅 일부 가능 | 스팸 등록 억제 |
| `limited_chat` | 채팅 시작/전송 일부 제한 | 조회/기존 기록 확인 가능 | 괴롭힘/복붙 문의 대응 |
| `review_required` | 운영 검토 필요 | 읽기 중심, 일부 쓰기 차단 | 고위험 패턴 임시 홀드 |
| `suspended` | 기간 정지 | 쓰기 전면 제한 | 중대 위반 |
| `withdrawn` | 자진 탈퇴/종료 | 본인 공개 프로필 종료 | 탈퇴 처리 |

원칙:
- 계정 상태는 `User.accountStatus`와 별도 `UserRestriction` 이벤트 이력으로 함께 관리하는 것이 바람직하다.
- 기능 제한은 가능한 한 세분화하고, 영구 정지 이전에 설명 가능한 중간 단계가 있어야 한다.
- 거래 진행 중인 계정에 제한이 걸릴 때는 상대방 보호를 위한 자동 취소/운영 안내 규칙이 함께 필요하다.

### 70.10 거래 중 제한 계정 처리 원칙
- 예약 확정 또는 완료 확인 대기 중인 사용자가 정지/제한되면, 상대방에게 아래 수준의 안내가 필요하다.
  - `상대방 계정 상태 변경으로 현재 거래 진행이 제한되었습니다.`
- 세부 위반 사유는 직접 노출하지 않는다.
- 시스템 기본안:
  1. 새 채팅/새 예약 생성 차단
  2. 기존 진행 거래는 읽기 가능 + 운영 안내 배너
  3. 필요 시 예약 강제 취소/완료 보류/분쟁 전환
- 운영자 화면에서는 제한 시점 당시의 활성 매물, 예약, 완료 대기 건수를 함께 보여야 한다.

### 70.11 데이터 모델 후보
| 엔티티/필드 | 설명 |
|---|---|
| `User.assuranceLevel` | 계정 확인 수준 |
| `User.accountStatus` | 현재 계정 상태 |
| `UserRiskProfile` | 내부 리스크 집계 스냅샷 |
| `UserRestriction` | 기능 제한/정지 이력 |
| `UserIdentityVerification` | 휴대폰/인증 수단 확인 기록 |
| `UserSession` | 활성 세션/기기 정보(확장) |
| `UserDeviceRiskToken` | 기기 단위 리스크 식별 보조값 |

`UserRiskProfile` 필드 후보:
- `userId`
- `riskTier`: `low` / `medium` / `high` / `critical`
- `riskScore(optional)`
- `ruleHitSummaryJson`
- `lastEvaluatedAt`
- `requiresManualReview`
- `lastRestrictionAt`

### 70.12 API/응답 시사점
사용자용 응답에 직접 내부 리스크를 노출하지 않되, 화면 제어를 위한 최소 힌트는 제공할 수 있다.

예시:
```json
{
  "me": {
    "userId": "user_123",
    "assuranceLevel": "A2_verified",
    "accountStatus": "limited_chat"
  },
  "availableActions": ["view_listings", "view_chats"],
  "policyHints": [
    {
      "code": "ACCOUNT_CHAT_LIMITED",
      "severity": "warning",
      "message": "현재 채팅 기능이 일시 제한되어 있습니다. 자세한 내용은 운영 안내를 확인해 주세요."
    }
  ]
}
```

관리자용 API 후보:
- `GET /admin/users/{userId}/risk-profile`
- `GET /admin/users/{userId}/restrictions`
- `POST /admin/users/{userId}/assurance-review`
- `POST /admin/users/{userId}/restrictions`
- `POST /admin/users/{userId}/restrictions/{restrictionId}/lift`

### 70.13 분석/KPI 후보
- 가입 후 첫 매물 등록까지 시간
- 가입 후 첫 채팅 시작까지 시간
- assurance level별 매물 등록/채팅/완료 전환율
- 신규 계정의 신고율/차단율/예약 취소율
- 제한 조치 후 재위반율
- rate limit 노출 대비 정상 전환 회복률
- high risk 판정 계정의 실제 운영 확정율(오탐 관리)

### 70.14 오픈 질문
- MVP에서 `휴대폰 확인`을 필수로 둘지, 위험 행동에서만 단계적으로 요구할지?
- assurance level을 숫자/등급으로 직접 노출할지, 더 부드러운 배지형 카피로만 보여줄지?
- 동일 기기/네트워크 기반 다계정 탐지를 어디까지 허용할지?
- 신규 계정의 채팅/등록 제한을 얼마나 강하게 두어야 거래 전환을 해치지 않을지?
- 제한 계정이 기존 진행 거래의 채팅 읽기/소명 제출까지 막히면 오히려 분쟁 처리가 어려워지지 않는지?

## 71. `sell` / `buy` 매물 대칭성 및 차이 규격(초안)
### 71.1 목표
- 린클은 판매글만 있는 게시판이 아니라 **판매 매물(`sell`)과 구매 매물(`buy`)을 모두 거래 실행 단위로 다뤄야 한다.**
- 다만 두 매물은 작성자의 의도, 가격 해석, 채팅 시작 맥락, 예약 후 역할, 완료 확인 UX가 완전히 같지 않으므로 이를 명시적으로 분리한다.
- 이후 화면명세, DB, API, 분석, 운영 정책이 `listingType`을 단순 라벨이 아니라 행동 규칙으로 해석할 수 있게 한다.

### 71.2 기본 해석 원칙
| 항목 | `sell` 매물 | `buy` 매물 |
|---|---|---|
| 작성자 의도 | 내가 가진 아이템/재화를 판매 | 내가 원하는 아이템/재화를 구매 |
| 거래 상대 역할 | 문의자는 구매 희망자 | 문의자는 판매 희망자 |
| 기본 CTA 카피 | `채팅하기`, `구매 문의` | `판매 제안하기`, `팔기 문의` |
| 가격 의미 | 판매 희망가/협의 기준 | 구매 희망가/제안 수용 기준 |
| 수량 의미 | 내가 판매 가능한 수량 | 내가 구매하고 싶은 수량 |
| 완료 후 후기 대상 | 거래 상대 구매자 | 거래 상대 판매자 |

원칙:
- `Listing.authorUserId`는 항상 작성자이지만, `sellerUserId`, `buyerUserId`는 채팅/거래 컨텍스트에서 파생되어야 한다.
- 즉, `listingType=buy`에서는 **매물 작성자가 buyer role**이고, 문의자가 seller role이 된다.
- `ChatRoom`과 `Trade Thread`는 매물 타입에 따라 역할을 뒤집어 계산하되, 완료/후기/신고 정책은 동일한 도메인 언어를 유지해야 한다.

### 71.3 역할 매핑 규칙
#### 71.3.1 채팅방 역할 결정
- `sell` 매물:
  - 작성자 = `sellerUserId`
  - 문의자 = `buyerUserId`
- `buy` 매물:
  - 작성자 = `buyerUserId`
  - 문의자 = `sellerUserId`

설계 원칙:
- `ChatRoom`은 화면/운영/후기 계산 편의를 위해 `sellerUserId`, `buyerUserId`를 명시 저장하는 것이 바람직하다.
- 단순히 `authorUserId`, `counterpartyUserId`만 두면 후기/완료/예약 카피와 운영 검색에서 의미가 모호해질 수 있다.
- 동일 사용자 조합이라도 `sell`과 `buy`는 다른 채팅방/거래 문맥으로 취급한다.

#### 71.3.2 예약/완료/후기에서의 역할 유지
- 예약 객체의 `proposerUserId`, `counterpartUserId`는 실제 제안 행위자 기준이다.
- 그러나 `seller`/`buyer` 역할은 거래 맥락 기준으로 고정 유지해야 한다.
- 후기 작성 시에도 “누가 매수자였고 누가 매도자였는가”는 거래 품질 분석 차원에서 별도 저장 가치가 있다.

후보 필드:
- `TradeCompletion.sellerUserId`
- `TradeCompletion.buyerUserId`
- `Review.tradeRoleOfReviewer`
- `Review.tradeRoleOfReviewee`

### 71.4 가격 해석 규칙
| 필드 | `sell` 매물 해석 | `buy` 매물 해석 |
|---|---|---|
| `priceType=fixed` | 이 가격에 판매 의사 | 이 가격에 구매 의사 |
| `priceType=negotiable` | 협의 가능하지만 기준 판매가 존재 | 협의 가능하지만 기준 구매가 존재 |
| `priceType=offer` | 구매자가 가격 제안 | 판매자가 가격 제안 |
| `priceAmount` | 판매 기준액 | 구매 기준액 |

세부 원칙:
- `buy` 매물에서 `offer`는 “가격 미정, 판매자가 조건 제시” 의미로 해석한다.
- `buy` 매물의 가격은 시세 보조 신호일 뿐, 판매자의 실제 제안가가 더 높거나 낮을 수 있다.
- 목록 카드/상세에서는 `buy` 매물일 때 가격 라벨을 단순히 `50만 아데나`로만 보여주지 말고, `50만 아데나에 구매 희망`처럼 방향성이 드러나는 카피를 검토한다.

후보 표시 라벨:
- `sell + fixed` → `50만 아데나`
- `sell + negotiable` → `50만 아데나 · 협의 가능`
- `buy + fixed` → `50만 아데나에 구해요`
- `buy + offer` → `가격 제안받아요`

### 71.5 수량/부분거래 해석 규칙
- `sell.quantity`: 판매자가 현재 양도 가능한 총 수량
- `buy.quantity`: 구매자가 현재 원하는 총 수량
- `buy` 매물도 `quantity > 1`을 가질 수 있으나, MVP에서는 **한 거래 스레드가 해당 수요의 일부 또는 전체를 충족할 수 있는지**를 자동 분할 계산하지 않는다.
- 따라서 `buy` 매물에서도 부분 체결은 아래 두 방식 중 하나로 단순화한다.
  1. 작성자가 남은 필요 수량을 수동 수정
  2. 해당 거래 완료 후 새 매수글 재등록

원칙:
- `buy` 매물에서 일부 수량만 충족된 경우에도 자동으로 동일 매물에 복수 완료를 누적시키는 구조는 MVP에서 지양한다.
- 부분 체결을 허용하더라도 상태는 단순히 `completed`로 끝내기보다 `needs_quantity_update` 성격의 내부 플래그를 둘 수 있다.

### 71.6 목록/상세/검색 UX 차이
#### 71.6.1 목록 카드
- `sell` 매물은 공급 카드, `buy` 매물은 수요 카드로 시각적 차이를 두는 것이 바람직하다.
- 최소 차이 요소:
  - 타입 배지: `판매`, `구매`
  - CTA 카피 차이
  - 가격 라벨 방향성 차이
  - 아이콘/색상 차이(과도한 분리보다 인지성 우선)

#### 71.6.2 상세 화면
- `buy` 매물 상세에는 아래 정보가 더 중요하다.
  - 원하는 상태/옵션/강화 조건
  - 희망 수량
  - 구매 가능 시간/장소
  - 판매자가 제안할 수 있는 범위 안내
- `sell` 상세는 보유품 설명 정확성과 상태가 더 중요하고,
- `buy` 상세는 원하는 스펙과 거래 가능성(예: 언제 받을 수 있는지)이 더 중요하다.

#### 71.6.3 검색/정렬
- 기본 검색 결과는 `sell` 중심이더라도, 사용자가 `buy 포함` 또는 `구매글만` 필터를 쉽게 켤 수 있어야 한다.
- `buy` 매물은 희소 공급자 유입 관점에서 별도 탭/필터 가치를 가진다.
- 추천순 계산 시 `buy`와 `sell`의 성공 신호를 동일 가중치로 쓰지 않을 수 있다.
  - `sell`: 조회→채팅 전환, 응답성, 가격 경쟁력
  - `buy`: 채팅→거래완료 전환, 조건 명확성, 최근 활동성

### 71.7 채팅 시작/빠른문구 규칙
| 상황 | 권장 CTA | 첫 빠른문구 예시 |
|---|---|---|
| `sell` 상세에서 문의 | `구매 문의하기` | `아직 거래 가능할까요?` |
| `buy` 상세에서 문의 | `판매 제안하기` | `해당 아이템 판매 가능합니다.` |
| `buy + offer` | `조건 제안하기` | `판매 가능하며 가격은 이렇게 생각합니다.` |

원칙:
- 빠른문구 템플릿은 `listingType`에 따라 달라져야 한다.
- `buy` 매물 문의에서는 판매자가 자신이 가진 아이템 상태/수량/희망 가격을 빠르게 제시하도록 유도하는 것이 전환에 유리하다.
- 시스템 메시지도 `판매자가 예약을 제안했습니다` / `구매자가 거래완료를 요청했습니다`처럼 역할 기반 카피가 가능해야 한다.

### 71.8 상태 전이와 완료 처리 차이
핵심 상태 enum은 `sell`/`buy` 모두 동일하게 유지하되, 상태 의미는 역할 기준으로 읽어야 한다.

| 상태 | `sell` 매물 의미 | `buy` 매물 의미 |
|---|---|---|
| `available` | 판매 가능 | 구매 의사 열려 있음 |
| `reserved` | 특정 구매자와 우선 협의 | 특정 판매자와 우선 협의 |
| `pending_trade` | 실제 판매 실행 직전 | 실제 구매 실행 직전 |
| `completed` | 판매 완료 | 구매 완료 |
| `cancelled` | 판매 중단/불발 | 구매 중단/불발 |

완료 처리 원칙:
- 완료 요청/확인/이의제기 플로우 자체는 `sell`/`buy` 동일하게 유지한다.
- 다만 운영/분석 관점에서는 아래를 분리해 저장하는 것이 유리하다.
  - `whoRequestedCompletion`: seller or buyer
  - `listingType`
  - `tradeRoleOfRequester`
- 이를 통해 `구매글에서 판매자 문의 후 완료까지의 전환율` 같은 지표를 분리 볼 수 있다.

### 71.9 후기/신뢰/운영 정책 시사점
- 후기는 기본 구조를 공유하되, 후속 분석에서는 `sell-origin trade`와 `buy-origin trade`를 분리할 필요가 있다.
- 예시:
  - 구매글 작성자가 실제로는 지나치게 까다로운 수요자인지
  - 판매 제안자가 buy 글에 반복 허위 접근하는지
  - 특정 사용자가 sell에서는 평판이 좋지만 buy에서는 노쇼가 많은지
- 운영 백오피스는 사용자 히스토리에서 아래를 구분 조회할 수 있어야 한다.
  - 판매글 작성자로서의 완료/취소/신고 이력
  - 구매글 작성자로서의 완료/취소/신고 이력
  - 상대 역할(매수자/매도자)별 후기 패턴

### 71.10 데이터/API 시사점
#### 71.10.1 데이터 모델
후보 필드/원칙:
- `Listing.listingType`는 필수 enum 유지
- `ChatRoom.sellerUserId`, `ChatRoom.buyerUserId` 명시 저장 권장
- `TradeThread.tradeRoleForViewer` projection 필요
- `TradeCompletion.requestedByTradeRole` 저장 후보
- `Review.tradeContextType`: `sell_listing_trade` / `buy_listing_trade` 후보

#### 71.10.2 API 응답
목록/상세/채팅/내거래 응답에는 최소 아래가 필요하다.
```json
{
  "listingId": "listing_123",
  "listingType": "buy",
  "displayTypeLabel": "구매",
  "displayPriceLabel": "50만 아데나에 구해요",
  "viewerTradeRoleIfStartChat": "seller",
  "availableActions": ["create_chat", "report"]
}
```

#### 71.10.3 QA 체크포인트
- `buy` 매물에서 작성자가 자동으로 buyer role로 계산되는가
- 후기/완료 집계가 `listingType`에 상관없이 동일 규칙으로 동작하되 역할 분석은 분리 가능한가
- `buy` 매물 카드/상세/채팅 CTA 문구가 `sell` 카피를 잘못 재사용하지 않는가
- `buy + offer` 케이스에서 가격 null 허용과 라벨 노출이 일관적인가

### 71.11 분석 이벤트 확장 후보
- `buy_listing_create`
- `buy_listing_chat_start`
- `buy_listing_offer_proposed`
- `buy_listing_completed`
- `sell_listing_completed`
- `trade_role_response_time_calculated`

핵심 분석 포인트:
- `buy` 매물의 채팅 시작률과 완료 전환율
- `buy` 매물에서 가격 미정(`offer`)이 전환에 미치는 영향
- 매수글 작성자 vs 매도글 작성자로서의 사용자 행동 품질 차이

### 71.12 오픈 질문
- MVP 홈/기본 검색에서 `buy` 매물을 `sell`과 동일 비중으로 섞어 보여줄지, 별도 탭으로 분리할지?
- `buy` 매물에서 판매자가 여러 조건 제안을 보낼 수 있게 할지, 단순 채팅 협의로만 둘지?
- 부분 체결을 `buy` 매물까지 정식 지원할지, MVP에서는 수동 수량 수정으로 제한할지?
- 후기/신뢰도 집계에서 `buy` 역할과 `sell` 역할을 분리 공개할지, 내부 분석 전용으로 둘지?

## 72. 채팅 Participant/Device State 및 실시간 연결 SLI 정책(초안)
### 72.1 목표
- 채팅 읽음 상태를 단순 `unreadCount` 캐시가 아니라 **사용자 기준 사실 상태**와 **기기/세션 기준 동기화 상태**로 분리해 다기기 환경에서도 일관성을 유지한다.
- 실시간 전송은 화려한 presence보다 **읽음/이벤트 유실 없음/복구 가능성**을 우선 목표로 정의한다.
- 채팅/내 거래/알림/운영 모니터링이 모두 같은 연결 상태 모델과 SLI를 사용하도록 기준을 만든다.

### 72.2 상태 계층 분리 원칙
채팅 상태는 아래 3계층으로 분리한다.

| 계층 | 주체 | 역할 | 대표 필드 |
|---|---|---|---|
| 사용자 참여 상태 | 사용자 1명 x 채팅방 1개 | 읽음/뮤트/알림/차단 후속 상태의 원천 | `ChatParticipantState` |
| 기기/세션 동기화 상태 | 사용자 기기/앱 세션 x 채팅방 1개 | 실시간 연결, 마지막 수신, gap 감지, backfill 기준 | `ChatDeviceState` |
| 이벤트 원천 로그 | 채팅방 x 이벤트 | 진실 원장(source of truth) | `ChatEvent` 또는 동등 모델 |

핵심 원칙:
- **사용자 기준 읽음 커서**는 여러 기기 중 가장 앞선 합의 커서를 대표값으로 관리한다.
- **기기/세션 상태**는 읽음 사실 그 자체가 아니라 `어디까지 동기화했는지`, `연결이 살아있는지`, `gap이 있는지`를 관리한다.
- unread 계산, 푸시 억제, 내 거래 액션 필요 판정은 사용자 기준 상태를 우선 사용한다.

### 72.3 `ChatParticipantState` 계약 초안
`ChatParticipantState`는 사용자 기준 채팅 참여 상태의 원천 객체다.

| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `chatRoomId` | 필수 | 채팅방 식별자 |
| `userId` | 필수 | 참여 사용자 |
| `participantStatus` | 필수 | `active` / `muted` / `blocked_view_only` / `left_hidden` 후보 |
| `lastReadEventId` | 필수 | 사용자 기준 마지막 읽은 이벤트 |
| `lastReadAt` | 선택 | 사용자 기준 읽음 시각 |
| `lastActionedEventId` | 선택 | 완료 확인/예약 응답 등 명시 행동이 마지막으로 반영된 이벤트 |
| `notificationPreference` | 선택 | `all` / `important_only` / `mute_push` |
| `muteUntil` | 선택 | 일시 뮤트 종료 시각 |
| `hasActionRequired` | 필수 | 예약 응답/완료 확인/분쟁 소명 등 사용자 액션 필요 여부 캐시 |
| `lastVisibleEventId` | 선택 | 실제 렌더링 확인 후 반영된 마지막 이벤트 |
| `updatedAt` | 필수 | 상태 갱신 시각 |

원칙:
- `lastReadEventId`는 후퇴 불가다.
- `lastVisibleEventId`는 "상세 화면에 실제 그려진 이벤트" 기준이라 읽음 품질 분석에 유용하지만, unread의 법적/기능적 원천은 `lastReadEventId`를 우선 사용한다.
- `hasActionRequired`는 계산 가능한 캐시이지만, 목록/알림/내 거래 우선순위 최적화를 위해 저장해둘 가치가 있다.

### 72.4 `ChatDeviceState` 계약 초안
`ChatDeviceState`는 기기/세션별 연결/복구 상태를 관리한다.

| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `deviceStateId` | 필수 | 기기 상태 식별자 |
| `chatRoomId` | 필수 | 채팅방 식별자 |
| `userId` | 필수 | 사용자 |
| `deviceId` | 필수 | 설치/디바이스 식별자 |
| `sessionId` | 선택 | 앱 세션 또는 웹 탭 식별자 |
| `transportType` | 필수 | `sse` / `polling` / `background_sync` |
| `connectionStatus` | 필수 | `connected` / `reconnecting` / `disconnected` / `backgrounded` |
| `lastAckedEventId` | 선택 | 이 기기가 서버 이벤트를 마지막으로 수신 확인한 이벤트 |
| `lastSyncedEventId` | 선택 | backfill 포함 반영 완료된 마지막 이벤트 |
| `gapDetected` | 필수 | 누락 가능성 감지 여부 |
| `gapSinceEventId` | 선택 | 누락 의심 시작점 |
| `lastHeartbeatAt` | 선택 | 연결 생존 확인 시각 |
| `pushEligible` | 필수 | 현재 세션이 포그라운드가 아니어서 푸시 대상이 되는지 여부 |
| `appLifecycleState` | 필수 | `foreground` / `background` / `terminated_unknown` |
| `updatedAt` | 필수 | 상태 갱신 시각 |

원칙:
- `ChatDeviceState`는 사용자 경험 복구용이며, 오래된 레코드는 TTL 기반 정리가 가능해야 한다.
- 같은 사용자가 여러 기기로 접속해도 unread 원천은 하나지만, 푸시 억제는 `deviceId/sessionId` 단위로 더 세밀하게 판단할 수 있다.
- `gapDetected=true`면 해당 기기는 읽음 커서를 올리기 전에 먼저 backfill을 완료해야 한다.

### 72.5 읽음 커서와 device ack의 관계
| 개념 | 의미 | 사용처 |
|---|---|---|
| `lastAckedEventId` | 기기가 서버 이벤트를 수신했다는 사실 | 재연결/차등 전송 |
| `lastSyncedEventId` | 기기 로컬 스토어까지 반영된 마지막 이벤트 | backfill 종료 판단 |
| `lastReadEventId` | 사용자가 읽었다고 간주되는 마지막 이벤트 | unread/푸시 억제/상대 읽음 |
| `lastActionedEventId` | 사용자가 행동으로 처리한 마지막 중요 이벤트 | 액션 필요 해제 |

설계 원칙:
- **수신됨 = 읽음 아님**. SSE로 받았더라도 사용자가 보지 않았으면 `lastReadEventId`를 올리지 않는다.
- **읽음 = 행동 완료 아님**. 완료 요청을 읽었더라도 확인/이의 제기를 하기 전까지 `hasActionRequired`는 유지될 수 있다.
- 다기기 중 어느 한 기기에서 읽음이 발생하면 사용자 기준 `lastReadEventId`를 갱신하고, 다른 기기는 다음 sync에서 그 사실을 받아 UI를 정리한다.

### 72.6 실시간 전송 기본안: SSE 우선 + polling 폴백
현재 PRD 기준 MVP 권장안은 **SSE(server-sent events) 우선, 실패 시 polling/backfill 폴백**이다.

선택 이유:
1. 모바일 거래 서비스의 핵심은 양방향 presence보다 `새 이벤트 전달 + 읽음 복구`다.
2. WebSocket 대비 운영/인프라 복잡도를 낮추면서도 서버→클라이언트 실시간성은 충분히 확보 가능하다.
3. 사용자→서버 쓰기 액션은 일반 HTTPS POST로도 충분히 모델링 가능하다.

기본 흐름:
1. 앱 포그라운드 진입 시 채팅 목록/열린 채팅방에 SSE 연결 시도
2. SSE 연결 성공 시 `afterEventId` 기준 diff 수신
3. 연결 끊김 시 지수 백오프로 재연결
4. 재연결 실패 또는 앱 백그라운드 상태면 polling/backfill로 전환
5. 기기별 `gapDetected`가 해소되면 정상 상태 복귀

### 72.7 연결 수명주기 정책
| 상태 | 진입 조건 | 시스템 동작 | 사용자 체감 |
|---|---|---|---|
| `connected` | SSE 정상 연결 | 새 이벤트 push, heartbeat 유지 | 실시간 수신 |
| `reconnecting` | 일시 끊김/네트워크 변경 | 백오프 재연결, `afterEventId` 준비 | 잠시 지연 가능 |
| `degraded_polling` | 재연결 실패 또는 정책상 SSE 미사용 | 주기 polling, 필요 시 백필 | 실시간성 약화 |
| `backgrounded` | 앱 백그라운드 | 연결 축소/해제, 푸시 전환 | 푸시 의존 |
| `recovering_gap` | 누락 의심 | 강제 backfill, 미읽음 재계산 | 로딩/복구 배너 가능 |

원칙:
- 앱이 백그라운드로 가면 무조건 연결을 유지하려 하기보다, 푸시 + 복귀 시 fast backfill을 우선한다.
- `recovering_gap` 동안에는 UI에 조용한 복구 상태를 둘 수 있으나, 액션 필요 이벤트는 숨기면 안 된다.

### 72.8 실시간 연결 장애/복구 규칙
- 서버는 클라이언트가 보낸 `afterEventId`가 너무 오래되었거나 정리된 경우 `gap_required` 성격 응답을 줄 수 있어야 한다.
- 이 경우 클라이언트는 단순 재연결 반복이 아니라 REST backfill로 전환해야 한다.
- 네트워크 전환(와이파이↔셀룰러), 앱 재시작, 장시간 sleep 이후에는 `lastAckedEventId`와 `lastSyncedEventId` 차이를 우선 점검한다.
- 동일 이벤트가 SSE와 polling 양쪽으로 들어오더라도 `eventId` 기준 dedup돼야 한다.

### 72.9 Read model / projection 보강 포인트
#### 채팅 목록 projection
목록은 단순 `last message`가 아니라 아래를 포함해야 한다.
- 사용자 기준 `lastReadEventId`
- 이 기기 기준 `gapDetected` 여부
- `hasActionRequired`
- `nextActionCode`
- `connectionHealthHint` (`live`, `delayed`, `syncing` 정도의 단순 라벨)

응답 예시 후보:
```json
{
  "chatRoomId": "chat_123",
  "lastEventId": "evt_999",
  "lastReadEventId": "evt_870",
  "unreadCount": 4,
  "hasActionRequired": true,
  "nextActionCode": "confirm_completion",
  "sync": {
    "connectionHealthHint": "syncing",
    "gapDetected": true
  }
}
```

#### 내 거래 projection
- `tradeThreadStatus` 외에 `syncState`를 추가할 수 있어야 한다.
- 예: `ready`, `delayed`, `recovering`
- 거래 실행에 필요한 핵심 이벤트가 아직 복구 중이면, 목록 상단 CTA보다 `복구 중` 배너가 먼저 보여야 할 수도 있다.

### 72.10 푸시 억제와 멀티디바이스 정책
- 같은 사용자가 A기기에서 채팅을 foreground로 보고 있으면, 동일 이벤트에 대해 A기기 푸시는 억제한다.
- B기기가 background/terminated 상태라면 사용자 설정에 따라 B기기 푸시는 유지 가능하다.
- 단, 사용자 경험 단순화를 위해 MVP에서는 **사용자 단위 푸시 억제**를 기본안으로 하고, 멀티디바이스 세밀 제어는 Post-MVP로 둘 수 있다.
- 액션 필요 이벤트(`completion_requested`, `reservation_proposed`, `dispute_more_info_requested`)는 보수적으로 푸시 누락보다 과소 억제를 피한다.

### 72.11 실시간/동기화 핵심 SLI·SLO 후보
#### 사용자 체감 SLI
| 지표 | 정의 | 초기 목표 후보 |
|---|---|---|
| `chat_event_delivery_latency_p95` | 서버 이벤트 생성→포그라운드 클라이언트 반영까지 | 3초 이하 |
| `action_required_event_delivery_latency_p95` | 예약 응답/완료 확인 등 중요 이벤트 생성→클라이언트 반영 | 2초 이하 |
| `chat_backfill_recovery_time_p95` | gap 감지→목록/상세 일관 상태 복구까지 | 5초 이하 |
| `read_cursor_propagation_p95` | 한 기기 읽음→상대 읽음/다른 기기 동기화 반영 | 5초 이하 |
| `duplicate_event_render_rate` | 동일 이벤트 중복 렌더링 비율 | 0.1% 이하 |

#### 시스템/운영 SLI
| 지표 | 정의 | 초기 목표 후보 |
|---|---|---|
| `sse_connect_success_rate` | 포그라운드 세션의 SSE 연결 성공 비율 | 98%+ |
| `sse_unexpected_disconnect_rate` | 비정상 연결 종료 비율 | 2% 이하 |
| `gap_detected_session_rate` | 세션당 gap 복구 필요 비율 | 3% 이하 |
| `backfill_api_success_rate` | backfill 요청 성공률 | 99%+ |
| `push_suppression_misfire_rate` | 읽지 않았는데 푸시가 억제된 추정 비율 | 낮을수록 좋음, 측정 체계 필요 |

원칙:
- 실시간 품질은 단순 연결 성공률보다 **중요 이벤트가 제때 도착했는지**로 평가한다.
- `action_required_event_delivery_latency`는 일반 메시지 지연보다 더 강하게 관리해야 한다.

### 72.12 관측성/로그 설계 시사점
- 실시간 계층은 최소 아래 이벤트를 남길 수 있어야 한다.
  - `chat_stream_connected`
  - `chat_stream_disconnected`
  - `chat_stream_reconnect_attempted`
  - `chat_gap_detected`
  - `chat_backfill_started`
  - `chat_backfill_completed`
  - `chat_read_cursor_updated`
- 로그 공통 속성 후보:
  - `userId`, `deviceId`, `sessionId`, `chatRoomId`, `afterEventId`, `lastAckedEventId`, `networkType`, `appVersion`
- 운영 대시보드에서는 서버 오류율보다 먼저 `gap_detected_session_rate`, `backfill_recovery_time`, `action_required_event_latency`를 봐야 한다.

### 72.13 QA 체크포인트 확장
- 두 기기 로그인 상태에서 A기기 읽음 후 B기기 unread가 합리적 시간 내 정리되는가
- SSE 끊김 후 polling 폴백으로도 완료 요청/예약 제안이 누락되지 않는가
- gap recovery 중 같은 이벤트가 목록/상세에 중복 표시되지 않는가
- 백그라운드 전환 직전 받은 이벤트와 복귀 후 backfill 결과가 충돌하지 않는가
- `hasActionRequired=false`가 되었는데 실제로 완료 확인 CTA가 남아 있거나, 반대로 액션 필요인데 목록 배지가 사라지는 불일치가 없는가

### 72.14 후속 오픈 질문
- MVP에서 `ChatDeviceState`를 영속 테이블로 둘지, Redis/TTL 상태 저장소로 둘지?
- 사용자 단위 푸시 억제와 기기 단위 푸시 억제 중 어느 수준까지 MVP에 포함할지?
- read cursor를 서버 authoritative로만 둘지, 클라이언트 optimistic update 허용 범위를 어디까지 둘지?
- SSE keepalive 주기와 모바일 배터리/데이터 비용 균형을 어떻게 잡을지?
- `tradeThread syncState`를 사용자에게 노출할지, 내부 디버그/운영용으로만 둘지?


## 73. 가격 제안 / 역제안(Offer Negotiation) 워크플로우(초안)
### 73.1 목표
- 거래 성사율을 높이기 위해 가격 협의를 자유 텍스트만이 아니라 구조화된 액션으로 지원한다.
- 판매자/구매자가 “얼마에 거래할지”를 빠르게 합의하고, 합의 결과가 예약/완료 플로우로 자연스럽게 이어지게 한다.
- 분쟁/운영 관점에서 누가 어떤 금액을 언제 제안·수락·거절했는지 감사 가능한 이력을 남긴다.

### 73.2 기본 원칙
- 가격 제안은 채팅의 하위 액션이지만, 별도 객체(`Offer`)로 저장한다.
- 텍스트 협상은 허용하되, 금액 합의의 시스템 기준은 구조화된 offer 액션이 우선이다.
- 한 채팅방에서 동시에 `active` 상태인 offer는 1개만 허용한다.
- offer는 예약과 동일하지 않다. 가격 합의 후에도 장소/시간 확정 전까지는 예약이 아니다.
- `sell`/`buy` 매물 모두 같은 구조를 쓰되, 기준가(anchor price)의 해석만 다르게 한다.

### 73.3 핵심 개념 정의
- `anchorPrice`: 매물에 게시된 기준 가격. `sell`에서는 판매 희망가, `buy`에서는 구매 희망가.
- `offer`: 특정 상대가 제안한 거래 금액 및 조건 묶음.
- `counterOffer`: 기존 offer를 거절하면서 새 금액으로 다시 제안하는 액션. 데이터 모델상 별도 타입이 아니라 `parentOfferId`를 가진 새 offer로 본다.
- `agreedPrice`: 양측이 명시적으로 수락한 최종 합의 금액. 예약/완료의 기준 금액이 된다.
- `offerExpiryAt`: 상대 응답 유효시간. 만료 후에는 자동 `expired` 처리된다.

### 73.4 Offer 상태 모델
- `active`: 상대 응답 대기 중인 제안
- `accepted`: 상대가 수락하여 합의 금액으로 확정됨
- `rejected`: 상대가 거절함
- `withdrawn`: 제안자가 응답 전에 철회함
- `expired`: 유효시간 경과로 자동 만료됨
- `superseded`: 같은 협상 흐름에서 더 최신 counter-offer가 생성되어 효력이 대체됨

### 73.5 상태 전이 규칙
- `active -> accepted | rejected | withdrawn | expired | superseded`
- `accepted/rejected/withdrawn/expired/superseded` 이후 재활성화는 불가하며 새 offer를 생성해야 한다.
- 새 counter-offer 생성 시 직전 `active` offer는 자동 `superseded` 처리한다.
- `accepted`된 offer가 존재하면 해당 chat room의 `agreedPrice`를 갱신하고, 이후 새 offer 생성은 허용하되 사용자에게 “기존 합의 금액 변경” 경고를 노출해야 한다.

### 73.6 제안 생성 규칙
- 제안자는 채팅 참여자여야 하며 차단/제재/채팅종료 상태에서는 offer를 생성할 수 없다.
- `listing.status`가 `available` 또는 `reserved`일 때만 offer 생성 가능하다. `trade_pending`, `completed`, `cancelled`에서는 불가하다.
- 제안 금액은 정수 단위이며 `currencyType`과 일치해야 한다.
- 금액 미입력 offer는 허용하지 않는다. MVP에서는 복합 조건(여러 아이템 묶음, 부분수량별 단가 차등)은 제외한다.
- 동일 사용자는 자신의 직전 `active` offer가 있는 상태에서 또 다른 offer를 만들 수 없고, 먼저 철회하거나 counter-offer 체인으로 대체해야 한다.

### 73.7 금액 validation / 정책 가정
- 최소 금액: 1 이상.
- 최대 금액: 시스템 정수 범위 및 게임 경제 기준을 넘지 않도록 별도 상한을 둔다(구현 시 정책 상수화).
- anchorPrice 대비 과도하게 낮거나 높은 금액은 차단보다 경고 우선으로 처리한다.
- 운영 리스크가 높은 계정은 일정 횟수 이상 저품질 제안(예: 다수 채팅방에 극단적 저가 제안 반복) 시 offer 생성 rate limit을 강화할 수 있다.

### 73.8 `sell` / `buy` 매물에서의 해석 차이
- `sell` 매물: 구매자가 더 낮은 금액을 제안할 수 있고, 판매자는 수락/거절/역제안 가능하다. 더 높은 금액 제안도 허용하지만 우선순위 자동 변경은 하지 않는다.
- `buy` 매물: 판매자가 더 높은 금액을 제안할 수 있고, 구매자는 수락/거절/역제안 가능하다. 더 낮은 금액 제안도 허용하되 UX 문구는 “이 가격에도 판매 가능”처럼 바뀐다.
- 어떤 경우에도 플랫폼은 가격 우선 자동 낙찰을 보장하지 않는다. 최종 선택권은 매물 작성자에게 있다.

### 73.9 채팅 UX 요구사항
- 채팅 입력창 상단 또는 빠른 액션에 `가격 제안`, `역제안`, `제안 수락`, `거절`, `철회` CTA를 둔다.
- 구조화된 offer는 일반 텍스트와 구분되는 카드형 시스템 메시지로 렌더링한다.
- offer 카드에는 최소 아래 정보를 노출한다.
  - 제안 금액
  - 기준가 대비 차이(금액/비율)
  - 제안 시각, 만료 시각
  - 현재 상태(active/accepted 등)
  - 가능한 CTA
- 만료 임박 시(예: 10분 전) 채팅방 헤더 및 내 거래 리스트에서 강조할 수 있어야 한다.
- 상대가 accepted한 offer가 있으면 예약 생성 CTA를 우선 노출한다.

### 73.10 예약/거래 상태와의 연동
- `accepted` offer가 생겨도 자동 예약 생성은 하지 않는다. 다만 예약 생성 폼의 기본 금액은 `agreedPrice`로 채운다.
- 예약이 이미 존재하는 경우에도 가격 재협상은 허용할 수 있으나, 예약 확정 후 가격 변경은 “예약 조건 변경”으로 명시하고 양측 재확인을 요구한다.
- `trade_pending` 진입 후에는 새 offer 생성을 막는다. 이 단계부터는 완료/취소/분쟁만 허용한다.
- 거래 완료 시 금액 기준은 마지막 `accepted` offer 또는 명시된 예약 금액 중 최신 확정값을 따른다.

### 73.11 알림 규칙
- 새 offer 수신 시 즉시 푸시/인앱 알림 대상이다.
- `counter-offer`는 일반 메시지보다 높은 우선순위를 가질 수 있다.
- offer 만료 임박 알림은 기본 off 또는 저소음 정책으로 시작한다.
- accepted/rejected/withdrawn/expired 전이는 모두 인앱 타임라인에 남기고, 푸시는 사용자 설정 및 중요도에 따라 제한한다.

### 73.12 데이터 모델 후보
- `Offer`
  - `id`, `chatRoomId`, `listingId`, `tradeThreadId?`
  - `parentOfferId?`, `sequenceNo`
  - `offeredByUserId`, `offeredToUserId`
  - `currencyType`, `amount`, `quantityScope`
  - `status`, `statusReasonCode?`
  - `expiresAt`, `respondedAt?`, `withdrawnAt?`
  - `acceptedReservationId?`
  - `createdAt`, `updatedAt`
- `ChatMessage`와는 `messageType=offer` + `offerId` 방식으로 연결한다.
- 채팅방/내거래 read model에는 `activeOfferSummary`, `lastOfferStatus`, `agreedPrice` 캐시 필드를 둘 수 있다.

### 73.13 API 후보
- `POST /chat-rooms/{chatRoomId}/offers`
- `POST /offers/{offerId}/accept`
- `POST /offers/{offerId}/reject`
- `POST /offers/{offerId}/withdraw`
- `GET /chat-rooms/{chatRoomId}/offers`
- `GET /me/trades?hasActiveOffer=true`
- 응답에는 항상 `availableActions`, `expiresAt`, `agreedPriceImpact`, `supersedesOfferId?`를 포함하는 방향을 권장한다.

### 73.14 분석 이벤트 후보
- `offer_created`
- `offer_viewed`
- `offer_accepted`
- `offer_rejected`
- `offer_withdrawn`
- `offer_expired`
- `offer_countered`
- `offer_to_reservation_started`
- 핵심 속성 후보: `listingType`, `anchorPrice`, `offerAmount`, `deltaFromAnchor`, `offerSequenceNo`, `timeToResponseSec`, `wasReserved`, `userRole`

### 73.15 운영/안전 체크포인트
- 극단적 저가/고가 반복 제안, 다중 채팅방 동시 제안, 수락 직전 반복 철회는 abuse 신호로 활용할 수 있어야 한다.
- 분쟁 시 운영자는 offer 체인과 최종 합의 금액을 한 화면에서 추적할 수 있어야 한다.
- 채팅 텍스트 합의와 구조화된 offer가 충돌할 경우, 운영 기본 원칙은 구조화된 시스템 기록 우선 + 텍스트는 보조 증빙으로 본다.

### 73.16 오픈 질문
- MVP에서 offer 만료 기본값을 30분/1시간/3시간 중 무엇으로 둘지?
- 부분 수량 거래가 본격화되면 `quantityScope`를 단순 `all` 외에 어떻게 확장할지?
- 수락된 offer 이후 추가 offer 생성 가능 범위를 어디까지 허용할지(혼선 방지 vs 유연성)?
- “즉시 거래 가능 가격(Buy Now 유사)”을 추후 별도 기능으로 둘지?


## 74. 카테고리별 속성 템플릿 / 아이템 속성 모델(초안)
### 74.1 목표
- `아이템명 + 자유설명`만으로는 검색 품질과 비교 가능성이 낮아지므로, 카테고리별로 반복되는 거래 속성을 구조화한다.
- 매물 등록 폼, 목록 카드, 상세 화면, 검색 필터, DB 스키마, API DTO가 같은 속성 정의를 재사용하도록 한다.
- 초기에는 과도한 정규화보다 **카테고리별 핵심 속성 몇 개를 정확히 구조화**하는 방향을 우선한다.

### 74.2 기본 개념
| 개념 | 의미 | 예시 |
|---|---|---|
| `ItemCategoryTemplate` | 카테고리별 등록/표시/검색 속성 정의 | 무기, 방어구, 소모품, 재화 |
| `ItemAttributeDefinition` | 개별 속성 정의 | 강화수치, 옵션 텍스트, 거래 단위 |
| `ListingAttributeValue` | 특정 매물에 저장된 속성 값 | `enhancementLevel=7` |
| `displayAttributeSet` | 카드/상세에 노출할 속성 묶음 | `+7`, `1개`, `옵션 있음` |

원칙:
- `categoryId`가 같아도 모든 매물이 동일한 속성 세트를 강제할 필요는 없지만, 검색/비교에 중요한 속성은 템플릿에서 표준화한다.
- 자유설명(`description`)은 보조 정보로 유지하고, 구조화 가능한 값은 가능한 한 `ListingAttributeValue`로 분리한다.

### 74.3 속성 타입 정의
| attributeType | 설명 | 예시 | 검색/필터 적합성 |
|---|---|---|---|
| `enum_single` | 단일 선택 | 거래 단위: 개별/세트 | 높음 |
| `enum_multi` | 다중 선택 | 클래스 제한, 옵션 태그 | 중간 |
| `number_int` | 정수 | 강화 수치 | 높음 |
| `number_decimal` | 소수 허용 숫자 | 수량, 묶음 수량 | 높음 |
| `text_short` | 짧은 텍스트 | 옵션 요약 | 낮음 |
| `boolean` | 예/아니오 | 안전거래 가능 여부(향후) | 중간 |
| `range_hint` | 구간성 값 | 사용 가능 레벨대 | 중간 |

원칙:
- MVP에서는 `enum_single`, `number_int`, `number_decimal`, `text_short` 위주로 시작하는 것이 안전하다.
- 자유 텍스트 속성은 검색 필터보다 상세 표시/운영 검토에 주로 활용한다.

### 74.4 카테고리 템플릿 필드 후보
| 필드 | 설명 |
|---|---|
| `templateId` | 템플릿 식별자 |
| `categoryId` | 적용 카테고리 |
| `templateVersion` | 속성 세트 버전 |
| `attributeKey` | 속성 키 (`enhancementLevel`, `unitType`) |
| `attributeLabel` | 사용자 노출 라벨 |
| `attributeType` | 값 타입 |
| `isRequiredForCreate` | 등록 시 필수 여부 |
| `isEditableAfterPublish` | 발행 후 수정 가능 여부 |
| `isFilterable` | 검색 필터 노출 여부 |
| `isSortable` | 정렬/랭킹 활용 가능 여부 |
| `isShownOnCard` | 목록 카드 노출 여부 |
| `isShownOnDetail` | 상세 노출 여부 |
| `validationRuleJson` | min/max/선택지/금칙 규칙 |
| `displayOrder` | 폼/상세 표시 순서 |

원칙:
- 템플릿 변경이 기존 매물 표시를 깨지 않도록 `templateVersion` 기반 snapshot 또는 하위호환 전략이 필요하다.
- 새 템플릿이 생겨도 기존 매물은 저장 당시 버전을 기준으로 해석 가능해야 한다.

### 74.5 카테고리별 MVP 기본 템플릿 예시
#### 74.5.1 무기/방어구형 장비
필수 후보:
- `enhancementLevel` (`number_int`)
- `quantity` (`number_decimal` 또는 공통 필드 재사용)
- `optionSummary` (`text_short`)
- `isPackage` (`boolean`)

카드 우선 노출 후보:
- 강화 수치
- 대표 옵션 존재 여부
- 수량

#### 74.5.2 소모품/재료형
필수 후보:
- `quantity`
- `unitType` (`enum_single`: 개별 / 묶음 / 스택)
- `packageCount` (`number_int`, 선택)

카드 우선 노출 후보:
- 총 수량
- 거래 단위

#### 74.5.3 재화/아데나형
필수 후보:
- `quantity`
- `unitScale` (`enum_single`: 만 / 백만 / 직접입력 보조 등, 정책 결정 필요)

원칙:
- 재화형은 별도 카테고리와 가격 표현이 충돌할 수 있으므로, `quantity`와 `priceAmount`의 의미가 혼동되지 않게 명세에서 분리해야 한다.

#### 74.5.4 구매글 전용 차이
- `buy` 매물도 동일 템플릿을 재사용할 수 있으나, 일부 속성은 `희망 조건` 의미로 해석된다.
- 예: `enhancementLevel=7`은 판매글에서는 실제 보유 속성, 구매글에서는 희망 최소/정확 조건일 수 있다.
- 따라서 속성별로 `matchMode` 또는 `valueInterpretation` 후보를 둔다.

### 74.6 등록 UX 원칙
- 등록 화면은 카테고리 선택 이후 해당 템플릿의 핵심 속성만 단계적으로 노출한다.
- 모든 속성을 처음부터 펼치기보다:
  1. 필수 속성
  2. 거래 조건(가격/수량/방식)
  3. 선택 속성
  4. 자유설명
  순서가 바람직하다.
- 속성 입력이 목록 카드 문구에 어떤 식으로 반영되는지 미리보기(preview)를 제공하면 입력 품질을 높일 수 있다.

### 74.7 검색/필터 연계 원칙
| 속성 | 필터 방식 후보 | 비고 |
|---|---|---|
| `enhancementLevel` | 최소값/정확값 | 장비 카테고리 핵심 |
| `unitType` | 단일 선택 | 소모품/재료형 유용 |
| `isPackage` | 토글 | 세트 거래 빠른 필터 |
| `quantity` | 최소 수량 | 대량 거래 탐색 |
| `optionSummary` | 키워드 보조 검색 | 구조화보다는 텍스트 보조 |

원칙:
- 모든 등록 속성을 필터로 노출하지 않는다. 검색 전환과 비교 가능성에 기여하는 속성만 `isFilterable=true`로 승격한다.
- 같은 속성도 카테고리에 따라 필터 가치가 다르므로 템플릿 수준에서 제어한다.

### 74.8 상세/카드 표시 원칙
- 카드에는 2~3개의 핵심 속성만 압축 노출한다.
- 상세 화면에서는 속성 라벨과 값을 구조화 표 형태 또는 칩 목록으로 노출한다.
- 속성이 비어 있을 때는 빈 라벨을 보여주지 않고 숨기되, 필수 속성 누락 매물은 등록 단계에서 차단하는 것이 우선이다.

예시 카드 라벨:
- `+7 · 옵션 있음 · 1개`
- `수량 500개 · 묶음 거래`

### 74.9 DB/API 설계 시사점
옵션 A: Listing 테이블에 컬럼 추가
- 장점: 단순, 조회 빠름
- 단점: 카테고리 확장 시 스키마 비대화

옵션 B: 속성 정의 + 값 테이블(EAV-lite)
- `ItemAttributeDefinition`
- `ListingAttributeValue`
- 장점: 카테고리 확장 유연
- 단점: 조회/필터 복잡도 증가

권장 기본안:
- **공통 핵심 필드(가격, 수량, 강화수치처럼 보편적이고 자주 쓰는 값)는 Listing 컬럼으로 승격**
- 그 외 카테고리별 확장 속성은 EAV-lite 또는 JSON snapshot으로 보조 저장
- 읽기 모델에서는 카드/상세용으로 평탄화된 projection을 제공

### 74.10 validation / 수정 제한 규칙
- `isRequiredForCreate=true`인 속성 누락 시 등록 차단
- `isEditableAfterPublish=false`인 속성은 `reserved` 이상에서 수정 금지 권장
- 숫자 속성은 min/max 및 비정상값 경고 규칙 필요
- 선택형 속성은 비활성 옵션 참조 시 기존 매물 표시 호환은 유지하되 신규 입력은 막아야 한다.

### 74.11 운영/카탈로그 도구 요구사항
운영자는 아래를 관리할 수 있어야 한다.
- 카테고리별 템플릿 생성/수정/비활성화
- 속성 선택지 추가/비활성화
- 자주 쓰이는 자유입력 속성값 승격 검토
- 특정 카테고리의 등록 품질 점검(필수 속성 누락률, 자유설명 의존률)

### 74.12 분석 이벤트 후보
- `listing_attribute_template_loaded`
- `listing_attribute_filled`
- `listing_attribute_skipped`
- `listing_attribute_validation_failed`
- `listing_attribute_filter_used`
- `listing_attribute_normalized`

핵심 분석 포인트:
- 카테고리별 등록 완료율
- 속성 입력률과 채팅 시작률/예약 전환율 상관관계
- 어떤 속성이 실제 검색/필터에서 쓰이는지
- 자유입력 설명만 있고 구조화 속성이 비어 있는 매물 비율

### 74.13 오픈 질문
- 초기 MVP에서 몇 개 카테고리까지 별도 템플릿을 둘 것인가?
- `enhancementLevel`, `quantity`처럼 자주 쓰는 속성을 Listing 공통 컬럼으로 승격할 범위는 어디까지인가?
- 카테고리 템플릿 변경 시 기존 매물 재색인/재계산을 어디까지 자동화할 것인가?
- 구매글의 희망 조건 표현(`최소 +7 이상`)을 별도 속성 모델로 분리할 것인가?


## 75. 액션별 Rate Limit / Cooldown / 안티어뷰즈 정책(초안)
### 75.1 목표
- 거래 성사에 필요한 정상 사용은 최대한 방해하지 않으면서, 도배 등록·대량 문의·괴롭힘·허위 신고·자동화 남용을 제어한다.
- limit 규칙을 단순 인프라 throttling이 아니라 **제품 정책**으로 정의해 API, 운영, UX, 분석이 같은 기준을 사용하도록 한다.
- 제재는 가능하면 즉시 영구 차단보다 `경고 -> 일시 제한 -> 가중 제한 -> 운영 검토` 순으로 단계화한다.

### 75.2 설계 원칙
1. 동일한 HTTP 429라도 원인은 다를 수 있으므로 `restrictionReasonCode`를 함께 반환한다.
2. limit는 `IP`, `device`, `account`, `listing`, `chatRoom`, `targetUser` 등 여러 키를 조합해 본다.
3. 신규 계정, 저신뢰 계정, 최근 제재 계정은 더 보수적인 한도를 적용할 수 있다.
4. 거래상 긴급성이 큰 액션(예약 응답, 완료 확인)은 일반 메시지보다 완화된 정책을 쓴다.
5. 단순 요청 횟수 외에 **중복도/유사도/상대 다양성/실패 패턴**도 함께 본다.
6. 사용자는 “왜 막혔는지”를 이해할 수 있어야 하지만, 악용 방지를 위해 정확한 임계치 전체를 모두 노출하지는 않는다.

### 75.3 제한 레벨 정의
| 레벨 | 의미 | 사용자 경험 | 운영 의미 |
|---|---|---|---|
| `L0_monitor` | 기록만 하고 차단 안 함 | 영향 없음 | 이상 패턴 탐지용 |
| `L1_soft_warn` | 경고/캡차/확인 요청 | 진행 가능하나 주의 표시 | 초기 억제 |
| `L2_cooldown` | 짧은 재시도 대기 | 몇 초~몇 분 후 재시도 | 순간 폭주 차단 |
| `L3_action_restrict` | 특정 기능 일시 제한 | 예: 새 채팅 1시간 제한 | 반복 남용 대응 |
| `L4_account_restrict` | 계정 수준 제한 | 등록/채팅 등 다중 기능 제한 | 운영 검토 필요 |
| `L5_staff_review` | 수동 검토 큐 진입 | 일부 기능 잠금 가능 | 고위험 사건 |

### 75.4 제한 키(limitKey) 후보
- `account_id`: 기본 사용자 단위
- `device_fingerprint`: 다계정 우회/신규 계정 폭주 탐지용
- `ip_prefix`: 공용망 과잉 차단 방지를 위해 prefix 또는 reputation 기반 완화 필요
- `listing_id`: 동일 매물 반복 변경/문의 남용 제어
- `chat_room_id`: 동일 대화방 스팸 연타 제어
- `target_user_id`: 특정 사용자 괴롭힘/다중 문의 집중 탐지
- `content_hash`: 동일 문구 복붙/동일 이미지 반복 업로드 탐지
- `risk_tier`: trust/reputation 기반 동적 limit 계층

### 75.5 액션별 기본 정책 매트릭스
| 액션 | 기본 정책 방향 | 주요 limitKey | 기본 대응 |
|---|---|---|---|
| 매물 생성 | 짧은 시간 다건 등록 제한 | `account_id`, `device_fingerprint`, `content_hash` | 쿨다운, 운영 검토 |
| 매물 핵심 수정(가격/제목/아이템) | 잦은 변경 제한 | `listing_id`, `account_id` | 짧은 쿨다운 |
| 매물 상태 변경 | 반복 토글 제한 | `listing_id`, `account_id` | 상태 토글 감점/경고 |
| 채팅방 생성 | 동일/다수 상대 대량 문의 제한 | `account_id`, `target_user_id`, `device_fingerprint` | 쿨다운, 신규 채팅 제한 |
| 메시지 전송 | 초당 연타/동일문구 복붙 제한 | `chat_room_id`, `account_id`, `content_hash` | 전송 지연/차단 |
| 예약 제안/수정 | 반복 예약 괴롭힘 제한 | `chat_room_id`, `listing_id` | 예약 기능 제한 |
| 완료 요청/분쟁 제기 | 중복 실행 방지 우선 | `completion_id`, `account_id` | 멱등 처리/409 |
| 신고 접수 | 허위/연속 신고 남용 제한 | `account_id`, `target_user_id`, `content_hash` | 신고 쿨다운, 검토 가중치 하향 |
| 이미지 업로드 | 대용량/반복 업로드 제한 | `account_id`, `device_fingerprint`, `content_hash` | 업로드 쿨다운 |

### 75.6 매물 등록/수정 제한 규칙
기본안(가정):
- 신규 계정은 첫 24시간 동안 활성 공개 매물 수 상한을 더 낮게 둔다.
- 동일 카테고리/서버/유사 제목 조합의 매물 연속 등록은 `content_hash` 기반으로 감지한다.
- 가격/제목/핵심 속성 변경은 짧은 시간 반복 시 검색 상단 점유를 노린 랭킹 악용으로 간주할 수 있다.

후보 규칙:
- 동일 사용자의 활성 공개 매물 수 상한
- 동일/유사 매물 연속 등록 쿨다운
- `available <-> reserved` 잦은 토글 시 랭킹 감점 + 운영 모니터링
- `bumpedAt` 기능 도입 시 별도 긴 쿨다운 필수

사용자 노출 원칙:
- 단순 상한 도달 시: `현재 활성 매물이 많아 새 등록은 잠시 후 가능합니다.`
- 의심 패턴 시: 구체 임계치 대신 `유사한 매물이 반복 등록되어 잠시 제한됩니다.` 수준으로 안내

### 75.7 채팅 시작 제한 규칙
목표는 판매자/구매자가 대화폭탄을 받지 않게 하면서도, 실제 탐색 중인 사용자의 정상 문의는 허용하는 것이다.

후보 규칙:
- 같은 사용자가 짧은 시간 내 많은 서로 다른 매물/상대에게 신규 채팅 생성 시 가중 제한
- 동일 상대에게 여러 매물 기반 문의를 연속 생성할 경우 `target_user_id` 기준 점진적 쿨다운
- 차단/신고 누적이 많은 계정은 새 채팅 생성 한도를 더 낮게 적용
- `reserved` 매물에 대한 대기 문의는 일반 `available` 매물 문의보다 더 보수적으로 제한 가능

예외/완화:
- 이미 거래 이력이 있거나 상대가 최근 응답한 관계는 limit를 완화할 수 있다.
- 동일 매물 재진입은 새 채팅 생성으로 보지 않고 기존 채팅방 재사용을 우선한다.

### 75.8 메시지 전송/복붙 방지 규칙
| 패턴 | 탐지 신호 | 기본 대응 |
|---|---|---|
| 초당 연속 전송 | 매우 짧은 간격 다건 전송 | 짧은 전송 지연 또는 쿨다운 |
| 동일문구 반복 | `content_hash` 반복 | 경고 후 차단 |
| 다중 채팅방 동일 홍보문 | 여러 `chat_room_id`에 동일 본문 전송 | 새 채팅/메시지 제한 |
| 외부 연락처 반복 전송 | 민감 패턴 다회 적중 | 메시지 차단 + 위험도 상승 |
| 욕설/괴롭힘 폭주 | 신고/정책 필터 + 반복 전송 | 채팅 잠금/운영 검토 |

원칙:
- 메시지 길이 제한과 속도 제한은 별도로 관리한다.
- 거래 실행형 액션 메시지(`도착했어요`, `예약 확인`)는 일반 자유텍스트보다 우선 전달되도록 완화 가능하다.
- 시스템 quick reply는 정상 사용으로 오탐하지 않도록 별도 분류가 필요하다.

### 75.9 예약/완료/분쟁 액션 제한 규칙
- 예약 제안은 같은 채팅방에서 연속 거절/취소가 누적되면 일정 시간 재제안을 제한할 수 있다.
- 시간/장소를 직전까지 반복 변경하는 패턴은 `locationMismatchRisk`와 함께 rate limit 점수에 반영한다.
- 완료 요청은 멱등 처리가 우선이며, 단순 재시도는 제재보다 같은 결과 반환을 기본으로 한다.
- 근거 없이 반복 분쟁 제기/취소를 반복하는 계정은 `dispute_open` 액션을 가중 제한하고 운영 검토 큐에 올린다.

### 75.10 신고 남용 방지 규칙
신고는 안전 기능이므로 지나친 차단은 위험하지만, 허위/보복 신고는 운영 비용을 크게 올린다.

후보 규칙:
- 동일 대상에 대한 짧은 시간 반복 신고 제한
- 거의 동일한 설명문으로 다수 대상 신고 시 가중 검토
- 과거 허위 신고 비율이 높은 계정은 자동 우선순위 상향 대신 `review_weight`를 낮춤
- P1 유형(사기/성희롱/개인정보 노출)은 일반 신고보다 완화된 limit 적용

원칙:
- 신고 버튼 자체를 과도하게 숨기지 않는다.
- 제한이 걸리더라도 고위험 신고 유형은 대체 접수 경로를 제공하는 것이 바람직하다.

### 75.11 위험도 기반 동적 제한
`riskTier` 후보:
- `trusted_high`
- `normal`
- `new_account`
- `watchlist`
- `restricted`

동적 반영 신호 후보:
- 가입 후 경과 시간
- 완료 거래 수/후기 품질
- 최근 신고/차단/노쇼 누적
- 동일 기기 다계정 패턴
- 외부 연락처/민감정보 차단 적중 횟수

원칙:
- 고신뢰 사용자는 약간 완화된 한도를 받을 수 있다.
- 저신뢰/감시 계정은 채팅 시작, 매물 생성, 이미지 업로드에서 더 보수적으로 제한한다.
- 위험도는 사용자에게 점수 그대로 노출하지 않고, 기능 제한 결과로만 간접 반영한다.

### 75.12 API 응답/상태 코드 원칙
제한 발생 시 응답 후보:
```json
{
  "code": "RATE_LIMITED",
  "message": "새 문의를 너무 빠르게 보내고 있어요. 잠시 후 다시 시도해 주세요.",
  "requestId": "req_123",
  "details": {
    "restrictionScope": "create_chat",
    "restrictionReasonCode": "rate_limit_new_chat_burst",
    "retryAfterSeconds": 120
  }
}
```

원칙:
- 인프라성 limit와 제품 정책성 limit를 `restrictionScope`/`restrictionReasonCode`로 구분한다.
- `retryAfterSeconds`가 있는 경우 UI는 타이머/비활성 CTA를 일관되게 제공해야 한다.
- 장기 제한은 429보다 `POLICY_BLOCKED` 또는 별도 account restriction 응답이 더 적합할 수 있다.

### 75.13 운영 도구/대시보드 요구사항
운영자는 아래를 볼 수 있어야 한다.
- 액션별 limit 적중 횟수
- 사용자/기기/IP 기준 상위 적중 주체
- 허위 양성 해제 비율
- rate limit 이후 실제 재위반/신고 전환율
- 정상 사용자가 과도하게 제한된 구간(전환 저하 징후)

운영 액션 후보:
- 사용자별 일시 완화/해제
- 특정 룰 일시 비활성화
- 공격 상황 시 전역 임계치 상향
- 특정 카테고리/서버만 별도 제한 정책 적용

### 75.14 분석 이벤트 후보
- `rate_limit_hit`
- `rate_limit_warning_shown`
- `rate_limit_retry_success`
- `rate_limit_escalated_to_restriction`
- `restriction_lifted`

공통 속성 후보:
- `restrictionScope`
- `restrictionReasonCode`
- `riskTier`
- `limitKeyType`
- `targetObjectType`
- `targetObjectId(optional)`

핵심 분석 포인트:
- 어떤 limit가 스팸 억제에는 효과적이고 전환 손실은 적은지
- 신규 계정과 고신뢰 계정의 제한 체감 차이
- 채팅 시작/메시지 전송 제한이 실제 신고율 감소에 기여하는지

### 75.15 후속 파생 문서 포인트
- API 명세: 429/423 응답 구조, `retryAfterSeconds`, `restrictionReasonCode` enum
- 운영정책: 제한 단계별 조치 기준, 수동 해제 절차, 오탐 대응
- DB/관측성: rate limit event log, device/account/IP 연계 전략
- QA: 정상 사용자와 남용 사용자 시나리오 분리 테스트, 경계값 테스트, 다계정/재시도 케이스 검증

## 76. 화면 라우팅 / 딥링크 / 복귀 계약(초안)
### 76.1 목표
- 모바일 중심 제품에서 사용자가 푸시, 검색, 찜, 채팅, 내 거래, 운영 알림 어디서 진입하든 **같은 거래 컨텍스트로 복귀**할 수 있어야 한다.
- 화면명세, 딥링크, 내비게이션 스택, API projection이 서로 다른 식별자를 쓰지 않도록 공통 라우팅 계약을 정의한다.
- 종료/권한제한/삭제된 객체로 들어온 경우에도 사용자가 길을 잃지 않도록 fallback 화면 규칙을 명시한다.

### 76.2 라우팅 설계 원칙
1. 도메인 객체 식별자(`listingId`, `chatRoomId`, `tradeThreadId`, `notificationId`)는 URL/deep link에서도 그대로 사용한다.
2. 가능하면 **행동 중심 딥링크**보다 **객체 중심 딥링크 + 초기 탭/액션 힌트** 구조를 우선한다.
3. 푸시/알림/운영 링크는 앱이 cold start 상태여도 최종 목적 화면으로 복구 가능해야 한다.
4. 접근 불가/만료/종결 상태에서는 404로 끝내지 않고, 정책에 맞는 soft landing 또는 대체 CTA를 제공한다.
5. 로그인 필요 액션은 원래 목적지와 필요한 컨텍스트를 보존한 뒤 인증 이후 복귀시킨다.

### 76.3 앱 주요 route 후보
| route | 목적 | 필수 path/query 식별자 | 대표 진입 소스 |
|---|---|---|---|
| `/home` | 개인화 홈 | 없음 | 앱 첫 진입, 로고 탭 |
| `/market` | 거래소 목록 | `serverId?`, `listingType?`, `query?` | 홈, 검색, 딥링크 |
| `/listings/{listingId}` | 매물 상세 | `listingId` | 목록 카드, 찜, 공유 링크 |
| `/listings/{listingId}/edit` | 매물 수정 | `listingId` | 내 매물 |
| `/chats` | 채팅 목록 | `filter?` | 하단 탭, 알림 |
| `/chats/{chatRoomId}` | 채팅 상세 | `chatRoomId`, `focusEventId?` | 매물 상세, 푸시, 내 거래 |
| `/trades` | 내 거래 목록 | `tab?` | 하단 탭, 홈 요약 |
| `/trades/{tradeThreadId}` | 거래 상세 워크스페이스 | `tradeThreadId`, `panel?` | 내 거래, 알림, 푸시 |
| `/notifications` | 알림함 | `filter?` | 하단/마이페이지 |
| `/users/{userId}` | 프로필 | `userId`, `contextChatRoomId?` | 매물 상세, 채팅, 후기 |
| `/me/listings` | 내 매물 | `statusTab?` | 하단/마이페이지 |
| `/reports/{reportId}` | 내 신고 상세 | `reportId` | 알림, 마이페이지 |
| `/admin/reports` | 운영 신고 큐 | `status?`, `priority?` | 운영 홈 |
| `/admin/reports/{reportId}` | 운영 신고 상세 | `reportId` | 운영 신고 큐 |
| `/admin/disputes/{disputeId}` | 운영 분쟁 상세 | `disputeId` | 운영 알림/큐 |

원칙:
- `chatRoomId`와 `tradeThreadId`는 역할이 다르므로 서로 대체하지 않는다.
- 채팅은 대화 중심 route, 거래 상세는 실행/상태 중심 route로 분리 유지한다.

### 76.4 route별 진입 가드 규칙
| route | 비회원 | 일반 회원 | 참여자/작성자 | 운영자 |
|---|---|---|---|---|
| `/home`, `/market`, `/listings/{listingId}` | 제한적 허용 | 허용 | 허용 | 허용 |
| `/chats`, `/chats/{chatRoomId}` | 로그인 유도 | 본인 참여방만 | 허용 | 운영 범위 내 허용 |
| `/trades`, `/trades/{tradeThreadId}` | 로그인 유도 | 본인 거래만 | 허용 | 운영 전용 뷰는 별도 |
| `/users/{userId}` | 정책 범위 내 summary만 | member/participant 레벨 | 허용 | staff 전용 별도 |
| `/reports/{reportId}` | 불가 | 본인 신고만 | 본인 건만 | 운영용 route 별도 |
| `/admin/*` | 불가 | 불가 | 불가 | 역할별 허용 |

가드 원칙:
- 권한 부족은 단순 `FORBIDDEN` toast로 끝내지 않고, 가능한 경우 대체 route로 이동시킨다.
- 예: 종결/비공개 채팅 딥링크 → 거래 기록 상세 또는 채팅 읽기전용 화면

### 76.5 목록 → 상세 → 행동 전환 계약
#### 거래소 목록 → 매물 상세
- 목록 카드는 기본적으로 `/listings/{listingId}`로 이동한다.
- 필터/정렬/스크롤 위치는 목록 상태로 보존해야 한다.
- 상세에서 뒤로 가기 시 마지막 목록 컨텍스트(`serverId`, `query`, `sort`, `scrollOffset`)를 복원한다.

#### 매물 상세 → 채팅 시작
- `create_chat` 성공 시 바로 `/chats/{chatRoomId}`로 진입할 수 있어야 한다.
- 단, 거래 실행 우선 제품 방향상 `tradeThread`가 즉시 생성 가능한 구조라면 내부적으로 `tradeThreadId`도 함께 확보해 후속 알림/내 거래와 연결한다.
- 첫 진입 overlay/coachmark는 1회성으로 제한하고, 거래 흐름을 가리지 않아야 한다.

#### 매물 상세 → 내 거래/거래 상세
- 이미 참여 중인 매물이라면 상세 CTA는 `채팅 시작`보다 `거래 보기` 또는 `대화 계속`으로 대체 가능해야 한다.
- 이 경우 우선 이동 대상은 `/trades/{tradeThreadId}` 또는 `/chats/{chatRoomId}` 중 현재 action required가 더 큰 화면으로 결정한다.

### 76.6 푸시/알림 딥링크 원칙
| 이벤트 | 기본 목적 route | 보조 파라미터 | 비고 |
|---|---|---|---|
| 새 메시지 | `/chats/{chatRoomId}` | `focusEventId`, `fromNotificationId` | 대화 맥락 우선 |
| 예약 응답 필요 | `/trades/{tradeThreadId}` | `panel=reservation` | 액션 우선 |
| 예약 확정/임박 | `/trades/{tradeThreadId}` | `panel=meeting` | 장소/시간 요약 우선 |
| 거래 완료 요청 | `/trades/{tradeThreadId}` | `panel=completion` | 확인/이의제기 CTA 우선 |
| 분쟁 소명 요청 | `/trades/{tradeThreadId}` | `panel=dispute` | 일반 채팅보다 분쟁 패널 우선 |
| 후기 작성 요청 | `/trades/{tradeThreadId}` | `panel=review` | 종료 거래 기록과 연결 |
| 신고 처리 결과 | `/reports/{reportId}` | `fromNotificationId` | 본인 신고 맥락 |
| 운영 경고/제한 | `/me/account-status` 또는 정책 화면 | `reasonCode` | 계정 상태 설명 우선 |

원칙:
- 단순 메시지 이벤트는 채팅으로, **행동이 필요한 거래 이벤트는 거래 상세로** 보낸다.
- 푸시 payload는 직접 문구보다 `targetRoute`, `targetId`, `panel`, `notificationId` 같은 구조화 필드를 우선 포함한다.

### 76.7 딥링크 payload 후보
```json
{
  "targetRoute": "/trades/{tradeThreadId}",
  "targetId": "tt_123",
  "panel": "completion",
  "notificationId": "noti_456",
  "fallbackRoute": "/trades",
  "authRequired": true,
  "context": {
    "chatRoomId": "chat_123",
    "listingId": "listing_123"
  }
}
```

설계 원칙:
- 앱이 특정 route를 아직 지원하지 않는 구버전일 수 있으므로 `fallbackRoute`를 함께 두는 것이 안전하다.
- 민감한 exact 장소, 연락처, 신고 상세 텍스트는 payload에 직접 넣지 않는다.

### 76.8 인증/권한 부족 시 복귀 규칙
- 비회원이 찜/채팅/내 거래/알림 딥링크로 진입하면 로그인 화면으로 리다이렉트한다.
- 로그인 성공 후 아래를 복원해야 한다.
  1. 원래 target route
  2. 필요한 query/panel
  3. 목록 스크롤/검색 조건(해당되는 경우)
- 권한이 없는 객체로 진입한 경우:
  - 본인 거래가 아님 → `/trades` 또는 `/home`으로 fallback
  - 운영 route 무권한 → 일반 사용자 홈으로 fallback + 오류 로그 저장
  - 숨김/차단된 매물 → 정책 안내가 포함된 soft landing 또는 유사 매물 목록으로 fallback

### 76.9 종료/삭제/만료 객체 fallback 규칙
| 원래 진입 대상 | 접근 불가 사유 | fallback 화면 | 사용자 문구 방향 |
|---|---|---|---|
| 매물 상세 | `completed`/`cancelled` 공개 종료 | 종료 랜딩 + 유사 매물 | `거래가 종료되었어요` |
| 매물 상세 | 운영 숨김 | 정책 안내 랜딩 | `현재 볼 수 없는 매물입니다` |
| 채팅 상세 | 참여 권한 상실/숨김 | 거래 상세 또는 채팅 목록 | `더 이상 대화할 수 없어요` |
| 거래 상세 | tradeThread 종료 후 압축 보관 | 종료 기록 화면 | `거래 기록으로 이동했어요` |
| 프로필 | 탈퇴/비공개 | 축약 프로필 또는 이전 화면 | `탈퇴한 사용자입니다` |
| 신고 상세 | 삭제/비공개 처리 | 내 신고 목록 | `확인 가능한 기록이 없습니다` |

원칙:
- fallback은 항상 사용자가 다음 행동을 할 수 있는 화면이어야 한다.
- 404는 개발/디버깅에는 필요하지만, 제품 UX에서는 가능한 마지막 수단으로 제한한다.

### 76.10 뒤로가기 / 스택 복귀 원칙
- 하단 탭 루트(`/home`, `/market`, `/chats`, `/trades`, `/notifications`, `/me`)는 각자 독립 navigation stack을 가지는 구조를 우선 검토한다.
- 푸시로 진입한 상세 화면은 뒤로가기 시 아래 우선순위를 따른다.
  1. 같은 탭 내 이전 화면이 있으면 해당 화면
  2. 없으면 목적 기능의 루트 탭(예: 거래 이벤트 → `/trades`, 메시지 → `/chats`)
  3. 그래도 없으면 `/home`
- 외부 공유 링크로 직접 매물 상세에 진입한 경우 뒤로가기는 `/market`으로 연결하는 것이 자연스럽다.
- 로그인 중간 개입이 있더라도 인증 이전의 navigation intent를 최대한 보존해야 한다.

### 76.11 탭/배지/글로벌 진입 규칙
- 하단 탭 배지는 최소 아래 두 종류를 분리한다.
  - `채팅 미읽음 수`
  - `액션 필요 거래 수`
- 홈의 거래 요약 위젯은 `/trades?tab=action_required` 또는 특정 `tradeThreadId`로 딥링크 가능해야 한다.
- 알림함의 각 항목은 자체 상세 화면을 갖기보다, 가능한 한 도메인 원본 객체로 연결해야 한다.
- 단, 운영/신고 결과처럼 원본 객체보다 사건 상태가 중요한 경우에는 `/reports/{reportId}` 같은 사건 route를 유지한다.

### 76.12 웹 공유/SEO용 공개 URL 원칙
- 공개 공유 링크는 우선 매물 상세에 한정한다.
- 공개 URL 후보:
  - `/listings/{listingId}`
  - 선택적으로 slug 포함: `/listings/{listingId}-{slug}`
- slug는 표시/SEO용 보조값이며, 라우팅의 진짜 키는 `listingId`다.
- 완료/취소/숨김 전환 시에도 기존 공개 URL은 가능하면 소프트 랜딩으로 응답해 깨진 링크 경험을 줄인다.
- 프로필/채팅/거래 상세는 공개 공유 URL 대상에서 제외한다.

### 76.13 데이터/API 시사점
- Notification/Push payload는 아래 필드를 공통으로 가질 수 있다.
  - `targetRoute`
  - `targetId`
  - `panel`
  - `fallbackRoute`
  - `authRequired`
  - `contextJson`
- 읽기 API는 현재 객체 외에도 `canonicalRoute`, `fallbackRoute`, `shareUrl(optional)` 필드를 포함할 수 있다.
- `GET /me/trades`, `GET /notifications`, `GET /chats`는 각 item마다 바로 이동 가능한 `deepLink`를 주는 구조를 검토할 수 있다.

예시:
```json
{
  "tradeThreadId": "tt_123",
  "deepLink": {
    "route": "/trades/tt_123",
    "panel": "reservation",
    "fallbackRoute": "/trades"
  }
}
```

### 76.14 분석 / QA 체크포인트
분석 이벤트 후보:
- `deep_link_open`
- `deep_link_fallback_used`
- `login_redirect_return_success`
- `screen_restore_after_auth`
- `soft_landing_view`

QA 체크포인트:
- 비회원이 매물 상세에서 로그인 후 원래 보던 매물과 스크롤 위치로 정확히 복귀하는가
- 예약 응답 푸시 탭 시 채팅이 아니라 거래 상세 `reservation` 패널이 열리는가
- 종료된 매물 공유 링크가 깨지지 않고 soft landing으로 연결되는가
- 권한 없는 운영 route 접근 시 민감 정보 노출 없이 안전하게 fallback 되는가
- 푸시 cold start / warm start / foreground 상태에서 동일 target route 해석이 일관적인가

### 76.15 후속 파생 문서 포인트
- 화면명세: route별 진입 조건, back behavior, empty/error/fallback 상태
- API 명세: `deepLink`, `canonicalRoute`, `fallbackRoute`, `panel` enum 계약
- 푸시/알림 문서: notification payload schema, auth redirect flow, route versioning
- 모바일 아키텍처: 탭별 navigation stack, cold start routing, login redirect persistence
- QA: 딥링크 매트릭스, 권한별 진입/복귀 케이스, 종료 객체 soft landing 시나리오
## 77. 신규 사용자 온보딩 / 인증 / 신뢰 부트스트랩(초안)
### 77.1 목표
- 첫 방문 사용자가 매물 탐색에서 이탈하지 않고, 최소 단계로 거래 가능한 상태까지 진입하게 한다.
- 신규 계정의 스팸/사기 리스크를 낮추되, 정상 사용자에게 과도한 장벽을 주지 않는다.
- 가입, 프로필 설정, 첫 매물 등록, 첫 채팅 시작, 첫 예약까지의 상태 전이를 화면/API/운영이 같은 기준으로 해석하게 한다.

### 77.2 온보딩 상태 모델
| 상태 | 의미 | 대표 진입 시점 | 제한/허용 |
|---|---|---|---|
| `guest` | 비로그인 탐색 사용자 | 첫 방문 | 목록/상세 제한 조회만 가능 |
| `registered` | 로그인 완료, 최소 계정 생성 | 로그인 직후 | 찜/채팅/등록 일부 제한 가능 |
| `profile_ready` | 닉네임/주 서버 등 최소 프로필 완료 | 첫 프로필 설정 후 | 채팅 시작, 매물 등록 가능 |
| `trust_bootstrap` | 첫 거래 전 기본 신뢰 부트스트랩 단계 | 초기 활동 중 | 속도 제한, 일부 고위험 행위 제한 |
| `trade_enabled` | 일반 거래 가능 상태 | 조건 충족 후 | 기본 기능 전부 사용 |
| `restricted_onboarding` | 온보딩 중 정책/위험 제한 | 스팸/반복 위반 탐지 시 | 등록/채팅 일부 또는 전체 제한 |

원칙:
- `registered`와 `trade_enabled` 사이에 최소한 `profile_ready` 단계를 둬, 익명성 높은 저품질 계정 대량 유입을 줄인다.
- `trust_bootstrap`은 별도 인증 강제가 아니라 초기 행동 가드레일을 의미한다.
- 사용자는 현재 상태명을 직접 보기보다, `거래 시작을 위해 닉네임을 설정해 주세요`, `첫 거래를 위해 프로필을 조금만 더 채워 주세요` 같은 행동 문구로 인지한다.

### 77.3 최소 가입/프로필 완료 요건
MVP 기본안의 `profile_ready` 진입 요건:
- 로그인 완료
- 서비스 닉네임 설정
- 주 활동 서버 1개 선택
- 약관/커뮤니티/거래안전 정책 동의
- 차단된 닉네임/광고성 소개문구 검수 통과

선택 입력(초기 스킵 가능):
- 아바타
- 자기소개
- 선호 거래 방식(`in_game`/`offline_pc_bang`/`either`)
- 자주 거래 가능한 시간대

원칙:
- 첫 거래 시작까지 필요한 필드는 최소화한다.
- 단, 운영/신뢰 관점에서 최소 식별 단서(닉네임, 서버)는 반드시 구조화한다.
- 선택 필드는 첫 채팅/첫 예약 직전 맥락형 리마인드로 보완한다.

### 77.4 신규 사용자 초기 제한 정책
| 기능 | `registered` | `profile_ready` | `trust_bootstrap` | `trade_enabled` |
|---|---|---|---|---|
| 목록/상세 조회 | 가능 | 가능 | 가능 | 가능 |
| 찜 | 가능 또는 프로필 완료 후 허용 | 가능 | 가능 | 가능 |
| 채팅 시작 | 제한 가능 | 가능 | 가능(속도 제한 강화) | 가능 |
| 매물 등록 | 제한 가능 | 가능 | 가능(활성 매물 수 제한) | 가능 |
| 예약 확정 | 제한 없음 또는 조건부 허용 | 가능 | 가능 | 가능 |
| 이미지 다중 업로드 | 제한 가능 | 제한적 허용 | 제한적 허용 | 가능 |
| 외부 연락처 허용 단계 진입 | 불가 | 불가 | 매우 제한적 | 정책 범위 내 가능 |

권장 기본안:
- `registered` 상태에서는 탐색/찜 중심으로 두고, 채팅/매물 등록은 `profile_ready` 이후 허용한다.
- `trust_bootstrap` 상태에서는 기능 자체를 막기보다 `활성 매물 수`, `채팅 생성 속도`, `동시 문의 수`를 보수적으로 제한한다.
- 정상 완료 거래/정상 활동이 누적되면 자동으로 `trade_enabled`로 승격한다.

### 77.5 신뢰 부트스트랩 해제 조건 후보
`trade_enabled` 전환 후보 신호:
- 프로필 최소 요건 충족
- 첫 신고/차단/정책 위반 없이 일정 기간 정상 활동
- 첫 채팅 또는 첫 매물 등록 이후 이상 패턴 미탐지
- 첫 완료 거래 1건 또는 운영상 충분한 정상행동 신호

주의:
- 첫 완료 거래를 절대 조건으로 두면 신규 판매자/구매자가 막힐 수 있으므로, **행동 기반 완화 + 완료 기반 가산** 구조를 권장한다.
- 내부적으로는 risk score 하향 또는 restriction tier 해제 형태가 구현상 유리하다.

### 77.6 화면 요구사항
#### 77.6.1 게스트 → 로그인 전환
- 채팅 시작, 찜, 매물 등록 CTA 클릭 시 전체 페이지 이탈보다 바텀시트/풀스크린 온보딩으로 전환한다.
- 로그인 후 원래 보던 매물, 필터, 스크롤 위치를 복원해야 한다.
- 로그인 프롬프트에는 `거래 상태 알림`, `예약 관리`, `후기/신뢰 확인` 같은 거래 맥락 가치를 강조한다.

#### 77.6.2 첫 프로필 설정 화면
목적: 1분 이내 `profile_ready` 진입
- 필수 입력만 먼저 묻고, 나머지는 스킵 가능하게 한다.
- 닉네임 정책/중복/금칙어 검사는 실시간에 가깝게 피드백한다.
- 서버 선택은 검색 가능한 단일 선택 기본안을 우선 적용한다.
- 완료 후 바로 이전 의도 화면(매물 상세, 채팅 시작, 등록 폼)로 복귀시킨다.

#### 77.6.3 첫 매물 등록/첫 채팅 시작 맥락형 가이드
- 첫 매물 등록 전: 금지 품목/연락처/과장 제목 금지 핵심만 짧게 보여준다.
- 첫 채팅 시작 전: `가능하면 플랫폼 안에서 조율하세요` 안전 가이드를 노출한다.
- 첫 예약 직전: 장소/시간/캐릭터명 공유 범위를 다시 확인시키는 체크리스트를 제공한다.

### 77.7 운영/리스크 정책 연계
- 신규 계정은 아래 패턴에 더 민감하게 탐지한다.
  - 가입 직후 다건 매물 등록
  - 가입 직후 다수 사용자 대상 복붙 문의
  - 외부 연락처/계좌/오픈채팅 유도
  - 고가 매물 다건 등록 후 빠른 예약 유도
- 자동 조치 기본안:
  1. 속도 제한 강화
  2. 추가 프로필 입력 유도 또는 경고
  3. 등록/채팅 임시 제한
  4. 운영 검토 큐 적재
- 신규 계정이라는 이유만으로 공개 낙인 배지를 노출하지 않는다.
- 사용자에게는 `새 계정이라서 제한`보다 `안전한 거래를 위해 잠시 일부 기능이 제한됩니다` 수준의 설명을 우선 사용한다.

### 77.8 데이터/API 시사점
후보 필드:
- `User.onboardingStage`
- `UserProfileCompletedAt`
- `User.tradeEnabledAt`
- `UserRestriction.bootstrapTier`
- `UserRestriction.bootstrapReasonCode`
- `User.firstListingAt`
- `User.firstChatAt`
- `User.firstCompletedTradeAt`

API 후보:
- `GET /me/onboarding-status`
- `POST /me/profile/bootstrap-complete`
- `GET /catalog/servers` (온보딩/프로필 입력 공용)

응답 예시:
```json
{
  "onboardingStage": "profile_ready",
  "nextRequiredAction": {
    "code": "none",
    "label": null
  },
  "restrictions": [
    {
      "code": "NEW_ACCOUNT_CHAT_RATE_LIMIT",
      "message": "새 계정은 잠시 채팅 시작 속도가 제한될 수 있어요."
    }
  ],
  "availableActions": ["create_listing", "create_chat", "favorite"]
}
```

원칙:
- 온보딩 상태는 화면별로 흩어지지 않고, 공통 `me` 계열 응답에서 재사용 가능해야 한다.
- `availableActions`와 동일한 설계 철학을 유지해, 클라이언트가 단계 조건을 하드코딩하지 않도록 한다.

### 77.9 분석 이벤트 / KPI 후보
이벤트:
- `signup_completed`
- `onboarding_profile_started`
- `onboarding_profile_completed`
- `first_chat_attempted`
- `first_chat_started`
- `first_listing_attempted`
- `first_listing_published`
- `onboarding_restriction_shown`
- `onboarding_restriction_released`

관찰 KPI:
- 로그인 완료 → `profile_ready` 전환율
- `profile_ready` → 첫 채팅 시작 전환율
- `profile_ready` → 첫 매물 등록 전환율
- 첫 기능 시도 시 정책/제한으로 막힌 비율
- 신규 계정 24시간 내 신고율 / 차단율 / 정상 완료 전환율

### 77.10 오픈 질문
- `registered` 단계에서 찜은 바로 허용하고, 채팅/등록만 제한하는 것이 최적인가?
- 휴대폰 인증/추가 인증을 MVP 필수로 둘지, 운영 리스크 상승 시점에만 도입할지?
- `trust_bootstrap` 해제 기준을 시간 기반, 행동 기반, 완료 거래 기반 중 어떤 조합으로 가져갈지?
- 신규 사용자에게 첫 예약 직전 체크리스트를 강제할지, 1회성 안내로 둘지?


## 78. Read Model / Materialized Projection 전략(초안)
### 78.1 목표
- 쓰기 모델(`Listing`, `ChatMessage`, `Reservation`, `TradeCompletion`, `Dispute`, `Notification`)과 화면용 읽기 모델을 분리해 모바일 화면 응답 속도와 상태 일관성을 함께 확보한다.
- `채팅`, `내 거래`, `알림함`, `운영 큐`가 각각 원천 이벤트를 매번 조인하지 않고, 목적에 맞는 projection을 공유하도록 설계한다.
- 실시간(SSE), polling, backfill, 배치 재계산이 모두 같은 projection 계약을 재사용하도록 해 구현/운영 복잡도를 낮춘다.

### 78.2 분리 원칙
1. **원천 진실(source of truth)**
   - 상태 전이와 감사 추적의 원천은 쓰기 모델 테이블/이벤트 로그다.
   - read model은 재생성 가능한 파생물이며, 단독 진실원으로 취급하지 않는다.
2. **화면 목적별 projection 분리**
   - 채팅 상세, 채팅 목록, 내 거래, 알림함, 운영 큐는 서로 다른 projection을 가진다.
   - 하나의 거대한 범용 DTO보다, 화면 목적별 얇은 projection을 여러 개 두는 쪽을 우선한다.
3. **강한 정합성이 필요한 필드와 지연 허용 필드 분리**
   - `availableActions`, `tradeThreadStatus`, `currentReservationSummary`, `completionStage` 같은 액션 판단 필드는 가능한 한 최신이어야 한다.
   - `viewCount`, `favoriteCount`, `responseBadge`, 일부 랭킹 점수는 지연 허용 가능하다.
4. **projection은 append/update 가능해야 하되 원본 훼손 없이 재빌드 가능해야 한다.**

### 78.3 핵심 projection 후보
| projection | 주요 소비 화면/API | 집계 단위 | 목적 |
|---|---|---|---|
| `ListingCardProjection` | 홈, 목록, 찜 | listingId | 빠른 검색/카드 렌더링 |
| `ChatThreadSummary` | 채팅 목록, 배지, 푸시 억제 | chatRoomId x viewerUserId | 채팅 목록 요약/미읽음/최근 메시지 |
| `TradeThreadProjection` | 내 거래 목록/상세 | listingId x counterpartyUserId x viewerUserId | 거래 실행 상태/다음 액션 |
| `NotificationFeedProjection` | 알림함 | notificationId x userId | 딥링크/묶음/읽음 상태 |
| `ModerationQueueProjection` | 백오피스 큐 | reportId 또는 disputeId | 운영 우선순위/대상 요약 |
| `ProfileTrustProjection` | 프로필/매물 작성자 요약 | userId x visibilityLevel | 공개 신뢰 요약 |

원칙:
- `TradeThreadProjection`은 `ChatThreadSummary`와 일부 필드를 공유할 수 있으나, `다음 액션`과 `예약/완료/분쟁` 문맥이 추가된 별도 projection으로 본다.
- 운영 큐는 `Report`와 `Dispute`를 내부적으로 분리 저장하더라도, 대기열 UI에서는 하나의 통합 projection으로 볼 수 있어야 한다.

### 78.4 `ChatThreadSummary` 계약
채팅 목록은 raw `ChatRoom`만으로 충분하지 않다. 사용자별 unread/뮤트/푸시 억제/현재 액션 필요 여부를 반영한 projection이 필요하다.

후보 필드:
- `chatRoomId`
- `viewerUserId`
- `listingId`
- `counterpartyUserId`
- `counterpartyProfileSummary`
- `listingCardSummary`
- `chatStatus`
- `lastMessagePreview`
- `lastMessageType`
- `lastMessageAt`
- `lastEventSequence`
- `unreadCount`
- `lastReadEventSequence`
- `hasMentionLikeActionNeeded` (후보)
- `activeReservationSummary(optional)`
- `completionSummary(optional)`
- `disputeSummary(optional)`
- `isMuted`
- `isBlockedByViewer`
- `pushSuppressedUntil(optional)`
- `sortKey`

원칙:
- unread는 메시지 개수보다 **이벤트 시퀀스 기반**으로 계산하는 것을 우선한다.
- 마지막 프리뷰는 정책상 노출 가능한 텍스트만 저장하며, 민감정보 마스킹 후 projection에 반영할 수 있어야 한다.
- 채팅 목록은 `activeReservationSummary` 정도만 포함하고, 상세 예약 diff는 채팅 상세/내 거래 상세에서 조회한다.

### 78.5 `TradeThreadProjection` 계약
`내 거래`는 채팅 목록보다 더 실행 중심이다. 아래 projection을 별도로 둔다.

후보 필드:
- `tradeThreadId` (파생 식별자)
- `viewerUserId`
- `listingId`
- `chatRoomId`
- `counterpartyUserId`
- `tradeRole` (`seller` / `buyer` from viewer perspective)
- `tradeThreadStatus`
- `nextPrimaryAction`
- `nextActionDeadlineAt(optional)`
- `listingSummary`
- `counterpartyProfileSummary`
- `currentReservationSummary(optional)`
- `meetingSummary(optional)`
- `completionStage(optional)`
- `disputeStatus(optional)`
- `lastActivityAt`
- `unreadCount`
- `priorityBucket` (`action_required` / `scheduled` / `waiting` / `closed`)
- `sortScore`
- `policyHints`

집계 규칙:
- `nextPrimaryAction`은 서버가 계산해 내려주는 것을 기본안으로 한다.
- 예시 값:
  - `reply_message`
  - `confirm_reservation`
  - `reconfirm_location`
  - `mark_completed`
  - `confirm_completion`
  - `submit_dispute_statement`
  - `write_review`
- `priorityBucket`은 화면 탭/섹션 구분에 직접 사용 가능해야 한다.
- `sortScore`는 앱이 복잡한 상태 조합을 다시 계산하지 않도록 서버 측에서 만들어 주는 것이 유리하다.

### 78.6 `NotificationFeedProjection` 계약
알림함은 단순 로그가 아니라, 중복 억제와 딥링크 복구를 지원하는 projection이어야 한다.

후보 필드:
- `notificationId`
- `userId`
- `eventType`
- `eventPriorityTier`
- `groupingKey`
- `title`
- `body`
- `displayBodyMasked`
- `deepLinkRoute`
- `deepLinkParamsJson`
- `targetResourceStatus`
- `isRead`
- `readAt`
- `isActionCompleted`
- `actionCompletedAt`
- `pushDeliverySummary`
- `createdAt`

원칙:
- 알림 목록은 원천 `Notification`과 거의 유사해 보여도, `targetResourceStatus`와 `isActionCompleted`가 붙어야 사용자에게 stale 알림을 다르게 보여줄 수 있다.
- 예: 이미 취소된 예약 알림은 `종료됨` 배지와 함께 표시하고, 클릭 시 soft landing을 제공한다.

### 78.7 `ModerationQueueProjection` 계약
운영 큐는 `Report`와 `Dispute`의 원천 구조 차이를 숨기고, 대기열/우선순위 관점에서 통합 정렬 가능한 projection이 필요하다.

후보 필드:
- `queueItemId`
- `queueItemType` (`report` / `dispute` / `policy_review`)
- `sourceId`
- `priority`
- `slaBucket`
- `status`
- `targetType`
- `targetId`
- `targetSummary`
- `reporterOrOpenerSummary`
- `riskSignals`
- `linkedListingId(optional)`
- `linkedChatRoomId(optional)`
- `assignedAdminId(optional)`
- `firstResponseDueAt`
- `finalResolutionDueAt`
- `lastUpdatedAt`

원칙:
- 운영 화면 정렬/필터는 projection 기준으로 하고, 상세 진입 시 원천 객체를 조회한다.
- `riskSignals`는 내부 전용 필드이며, 외부 API/사용자 응답으로 재사용하지 않는다.

### 78.8 projection 갱신 트리거
아래 이벤트는 projection 갱신 트리거가 되어야 한다.

| 원천 이벤트 | 갱신 대상 projection | 비고 |
|---|---|---|
| 새 메시지 생성 | `ChatThreadSummary`, `TradeThreadProjection`, `NotificationFeedProjection` | unread, last preview, 액션 필요 업데이트 |
| 읽음 커서 변경 | `ChatThreadSummary`, `TradeThreadProjection` | unread/푸시 억제 반영 |
| 예약 생성/확정/변경/취소/만료 | `ChatThreadSummary`, `TradeThreadProjection`, `NotificationFeedProjection` | meeting/타이머/CTA 변경 |
| 매물 상태 변경 | `ListingCardProjection`, `TradeThreadProjection` | 상태 배지/신규 문의 가능 여부 반영 |
| 완료 요청/확정/분쟁 전환 | `TradeThreadProjection`, `NotificationFeedProjection`, `ModerationQueueProjection` | 후기 개방/소명 요청 반영 |
| 신고/운영 조치 | `ModerationQueueProjection`, 일부 `TradeThreadProjection` | 잠금/제한 배너 반영 |
| 프로필 신뢰 집계 갱신 | `ProfileTrustProjection`, `ListingCardProjection`, `TradeThreadProjection` | 작성자/상대 요약 반영 |

원칙:
- 사용자 체감이 큰 projection(`ChatThreadSummary`, `TradeThreadProjection`)은 이벤트 기반 즉시 갱신을 우선한다.
- 무거운 랭킹/신뢰 집계성 projection은 배치/비동기 갱신을 허용한다.

### 78.9 projection materialization 전략 후보
옵션 A: **DB 테이블/뷰 기반 materialization**
- 장점: 단순 운영, SQL 친화적, 백오피스/배치 연동 쉬움
- 단점: write fanout 증가, 복잡한 viewer별 projection 비용 증가

옵션 B: **이벤트 소비자 기반 별도 projection 저장소**
- 장점: 채팅/알림처럼 고빈도 읽기 성능에 유리
- 단점: 운영 복잡도 증가, 재빌드/디버깅 비용 증가

권장 기본안:
- MVP는 **주요 projection을 애플리케이션 DB 내 materialized table 또는 캐시 테이블로 시작**하고,
- 채팅/알림 부하가 커지면 `ChatThreadSummary`, `NotificationFeedProjection`부터 별도 소비자/저장소로 분리하는 단계적 전략을 권장한다.

### 78.10 재빌드 / 백필 / 정합성 검증 원칙
- 모든 projection은 원천 이벤트 또는 현재 상태 테이블로부터 **전체/부분 재빌드 가능**해야 한다.
- 재빌드 단위 후보:
  - `listingId`
  - `chatRoomId`
  - `userId`
  - `reportId` / `disputeId`
- 운영/개발 도구에서 특정 객체 projection 재계산을 트리거할 수 있으면 디버깅에 유리하다.
- 정합성 체크 예시:
  - `TradeThreadProjection.unreadCount`와 `ChatParticipantState.lastReadEventSequence`의 관계
  - `TradeThreadProjection.tradeThreadStatus`와 `Listing.status + completionStage` 조합
  - `NotificationFeedProjection.targetResourceStatus`와 실제 대상 상태
- 정합성 실패는 조용히 무시하지 말고, SLI/오류 이벤트로 관측해야 한다.

### 78.11 분쟁 첨부와 일반 채팅 첨부의 분리 원칙
분쟁 소명 첨부는 일반 채팅 첨부와 보안/보관/권한 요구가 다르므로, 같은 저장 파이프라인을 쓰더라도 논리적 분리가 필요하다.

기본 원칙:
- 일반 채팅 첨부는 거래 당사자 열람을 전제로 한다.
- 분쟁 첨부는 **운영자 우선 열람, 상대방 직접 열람은 제한 또는 비공개 기본안**을 둔다.
- 따라서 아래 필드/구조 분리를 권장한다.

후보 필드:
- `attachmentPurpose` (`chat_message` / `dispute_evidence` / `report_evidence`)
- `visibilityScope` (`participants` / `staff_only` / `staff_and_reporter`)
- `retentionTier`
- `moderationHoldYn`
- `accessAuditRequiredYn`

권장 운영안:
- 같은 object storage bucket을 쓰더라도 prefix/ACL/서명 URL 정책을 purpose별로 분리한다.
- 분쟁 첨부는 일반 채팅 썸네일/프리뷰 캐시와 분리해, 앱의 일반 미디어 프리패치 대상에서 제외하는 것이 안전하다.
- `TradeThreadProjection`에는 분쟁 첨부 원문/썸네일을 직접 싣지 않고, `evidenceCount`와 `canSubmitMoreEvidence` 정도만 포함하는 것이 바람직하다.

### 78.12 API/화면 파생 포인트
- `GET /chats`는 `ChatThreadSummary` projection 기반으로 응답하는 것이 적절하다.
- `GET /me/trades`는 `TradeThreadProjection` 전용 endpoint로 두는 것을 권장한다.
- `GET /notifications`는 `NotificationFeedProjection` 기준으로, stale 상태 라벨과 딥링크 fallback 정보를 포함해야 한다.
- 운영 큐는 `GET /admin/queue` 또는 `GET /admin/reports` 내부 projection 통합층을 둘 수 있다.

화면 설계 시 유의점:
- 목록형 화면은 projection을 그대로 소비하고, 상세 진입 시에만 원천 객체/상세 projection을 추가 조회한다.
- 앱 클라이언트가 `Listing.status + Reservation + Completion + Dispute`를 직접 합성하는 구조는 피한다.
- `nextPrimaryAction`, `priorityBucket`, `availableActions`는 서버 계산 우선 원칙을 유지한다.

### 78.13 관측성 / SLI 후보
projection 운영 품질을 보기 위한 지표 후보:
- `projection_update_lag_ms`
- `projection_rebuild_count`
- `projection_rebuild_failure_count`
- `projection_staleness_detected_count`
- `chat_thread_summary_mismatch_count`
- `trade_thread_status_mismatch_count`
- `notification_target_stale_click_rate`
- `moderation_queue_projection_delay_ms`

초기 SLO 후보:
- 새 메시지 발생 후 `ChatThreadSummary` 반영 p95 2초 이하
- 예약 상태 변경 후 `TradeThreadProjection` 반영 p95 3초 이하
- 운영 조치 후 `ModerationQueueProjection` 반영 p95 10초 이하
- projection 재빌드 실패율 일일 0.5% 이하

### 78.14 오픈 질문
- `TradeThreadProjection`을 DB view/materialized table로 둘지, 애플리케이션 캐시 projection으로 둘지 초기 부하 예측에 따라 어디까지 분리할 것인가?
- `GET /me/trades`를 MVP에 포함해 클라이언트 복잡도를 낮출지, 초기에는 `GET /chats` + `GET /reservations` 조합으로 버틸지?
- 분쟁 첨부의 상대방 공개 범위를 완전 `staff_only`로 둘지, 운영 승인 후 일부 공유할지?
- projection 재빌드 도구를 관리자 백오피스에 둘지, 내부 운영 CLI/Job로만 둘지?


## 79. Projection 재빌드 / 정합성 복구 / 운영 Runbook(초안)
### 79.1 목표
- `ChatThreadSummary`, `TradeThreadProjection`, `NotificationFeedProjection`, `ModerationQueueProjection`을 단순 캐시가 아니라 **운영 가능한 파생 모델**로 다룬다.
- 이벤트 누락, 소비자 장애, 스키마 변경, 버그 수정 시 projection을 안전하게 재생성/복구할 수 있어야 한다.
- 재빌드 과정이 사용자 거래 흐름을 망가뜨리지 않도록, 부분 재계산·fallback·cutover 원칙을 미리 정의한다.

### 79.2 기본 원칙
1. projection은 재생성 가능한 파생물이며, 원천 이벤트/원천 상태를 대체하지 않는다.
2. 운영 기본값은 **전체 리빌드보다 대상 범위가 좁은 부분 재계산**이다.
3. 재빌드 중에도 사용자 읽기 화면은 가능한 한 degraded read 또는 마지막 정상 projection으로 서비스를 지속해야 한다.
4. 재빌드 성공 여부는 단순 job 성공이 아니라 **정합성 검증 통과**까지 포함해 판단한다.
5. 운영자 수동 수정은 원칙적으로 projection 직접 수정이 아니라 원천 객체 보정 또는 재계산 트리거 방식으로 수행한다.

### 79.3 재빌드 트리거 유형
| 유형 | 예시 | 권장 대응 |
|---|---|---|
| 이벤트 소비 누락 | consumer 다운, DLQ 적재, offset 유실 | 대상 범위 replay + 정합성 체크 |
| projection 버그 수정 | unread 계산식 오류, 상태 집계 로직 버그 | 코드 배포 후 affected scope rebuild |
| 스키마 변경 | 새 필드 추가, enum 확장, summary label 변경 | backfill job + 버전 전환 |
| 운영 보정 | 특정 거래/신고건 상태 꼬임 | single-object recompute |
| 대규모 장애 복구 | 저장소 손상, 캐시 전면 유실 | full rebuild + 단계별 cutover |

### 79.4 재빌드 범위 단위
재빌드는 아래 단위 중 하나로 실행 가능해야 한다.

| 범위 단위 | 대표 키 | 사용 시점 |
|---|---|---|
| 단일 채팅방 | `chatRoomId` | unread/preview 불일치, 단일 문의 복구 |
| 단일 거래 스레드 | `tradeThreadId` 또는 `listingId + counterpartUserId` | 내 거래 상태/CTA 복구 |
| 단일 사용자 | `userId` | 알림함/채팅목록/배지 재계산 |
| 단일 신고/분쟁 | `reportId` / `disputeId` | 운영 큐 정합성 복구 |
| 기간 범위 | `fromEventAt ~ toEventAt` | 배포 버그/consumer 누락 구간 복구 |
| 전체 projection 종류 | projection type | 저장소 손상/스키마 마이그레이션 |

원칙:
- 운영 도구는 최소 `chatRoomId`, `tradeThreadId`, `userId`, `reportId/disputeId` 단위 재계산을 지원하는 것이 바람직하다.
- 전체 재빌드는 비용과 리스크가 크므로 명시적 승인과 사전 공지 기준이 필요하다.

### 79.5 재빌드 입력 원천 우선순위
projection 재생성 시 어떤 데이터를 진실원으로 삼는지 고정해야 한다.

| projection | 1차 원천 | 2차 보조 원천 |
|---|---|---|
| `ChatThreadSummary` | ChatMessage/Event timeline + ChatParticipantState | ChatRoom 캐시 필드 |
| `TradeThreadProjection` | Listing + Reservation + TradeCompletion + Dispute + Chat state | 기존 projection snapshot |
| `NotificationFeedProjection` | Notification event log + read state | push delivery log |
| `ModerationQueueProjection` | Report + Dispute + ModerationAction | triage assignment cache |

원칙:
- 기존 projection을 다시 읽어 projection을 재구축하는 방식은 허용하되, 최종 진실원은 원천 도메인 객체/이벤트여야 한다.
- push 발송 성공 여부 같은 파생 신호는 `NotificationFeedProjection`을 보완할 수 있지만, 원천 알림 이벤트 자체를 대체하지 않는다.

### 79.6 재빌드 실행 모드
| 모드 | 설명 | 사용 시점 |
|---|---|---|
| `sync_inline` | 요청 직후 동기 재계산 | 단일 객체 디버깅, 운영 수동 복구 |
| `async_targeted` | 배경 job으로 특정 범위 재계산 | 사용자/거래/신고 단위 복구 |
| `async_bulk` | 기간/배치 단위 대량 재계산 | 배포 버그, 필드 backfill |
| `full_replay` | 이벤트 전체 재생성 | 저장소 재구축, 대형 장애 |

권장안:
- 운영 UI/CLI에서는 `async_targeted`를 기본으로 제공하고, `sync_inline`은 소규모/저비용 projection에만 제한적으로 허용한다.
- `full_replay`는 내부 운영 절차와 승인 체계를 요구한다.

### 79.7 Freeze / Replay / Cutover 절차
대규모 재빌드나 저장소 전환 시 아래 3단계 절차를 권장한다.

1. **Freeze scope 결정**
   - 어떤 projection 종류와 어떤 키 범위를 재빌드할지 확정
   - 실시간 소비를 완전히 멈출지, shadow write를 병행할지 결정
2. **Replay / Recompute 수행**
   - 원천 이벤트/상태에서 새 projection 생성
   - 진행률, 실패 건수, 대상 키 수 집계
3. **Cutover**
   - 샘플 검증 + 집계 검증 통과 후 새 projection 읽기로 전환
   - 전환 직후 mismatch 모니터링 강화

권장 원칙:
- 가능하면 write freeze보다 **shadow rebuild + atomic alias cutover**를 우선 검토한다.
- 채팅/내 거래처럼 사용자 체감이 큰 surface는 cutover 직후 일정 시간 fallback read 경로를 유지하는 것이 안전하다.

### 79.8 정합성 검증 체크리스트
재빌드 성공 판정 전 아래 수준의 검증이 필요하다.

#### 객체 수준 검증
- `ChatThreadSummary.unreadCount >= 0`
- `TradeThreadProjection.nextPrimaryAction`이 현재 상태에서 허용 가능한 액션인지
- `NotificationFeedProjection.readAt`이 사용자 읽음 이력과 모순되지 않는지
- `ModerationQueueProjection.priorityBucket`이 Report/Dispute 원천 우선순위와 일치하는지

#### 관계 수준 검증
- `TradeThreadProjection.tradeThreadStatus`와 `Listing.status + completionStage + reservationStatus` 조합 일치
- `ChatThreadSummary.lastEventId`가 실제 chat timeline 최신 이벤트와 일치 또는 허용 오차 범위 내인지
- `NotificationFeedProjection.targetDeepLink`가 현재 존재하는 대상/soft landing으로 연결 가능한지

#### 집계 수준 검증
- 사용자별 unread 총합
- 액션 필요 거래 수
- 운영 큐 미처리 건수
- 알림 unread 건수

### 79.9 Fallback / 사용자 영향 최소화 원칙
재빌드 중 projection이 일시 부정확하거나 비어 있을 수 있으므로 사용자 surface별 fallback을 정의한다.

| surface | fallback 원칙 | 사용자 노출 |
|---|---|---|
| 채팅 목록 | 최근 정상 projection 유지 또는 최소 last message만 조회 | 필요 시 짧은 복구 배너 |
| 내 거래 | 액션 필요 거래 우선 direct recompute 허용 | CTA 지연 시 재시도 안내 |
| 알림함 | 목록은 마지막 snapshot 유지, 상세 클릭 시 대상 직접 조회 | 일부 항목 최신화 중 문구 |
| 운영 큐 | 큐 읽기 지연 허용, 원천 신고/분쟁 상세 direct open 가능 | 운영 전용 stale 경고 |

원칙:
- 거래 실행에 직접 영향을 주는 `nextPrimaryAction`, 예약 시각, 완료 확인 기한은 fallback 중에도 가능한 한 direct source 검증을 우선한다.
- stale projection 때문에 잘못된 CTA를 노출하는 것보다, CTA를 잠시 숨기고 재검증하는 편이 안전하다.

### 79.10 Dead-letter / 재처리 정책
projection 갱신 소비 중 실패한 이벤트는 별도 dead-letter 흐름으로 관리해야 한다.

후보 필드:
- `projectionType`
- `sourceEventId`
- `targetKey`
- `failureCode`
- `retryCount`
- `firstFailedAt`
- `lastFailedAt`
- `lastErrorSnapshot`

운영 원칙:
- 일시 오류는 자동 재시도하되, 구조적 오류(스키마 불일치, enum 미지원)는 무한 재시도 금지
- 일정 횟수 이상 실패 시 DLQ 적재 + 운영 알림
- DLQ 수동 재처리 시 재빌드 실행자, 실행 사유, 범위를 감사 로그에 남긴다.

### 79.11 감사로그 / 운영 권한 원칙
- projection 재빌드/재처리/강제 cutover는 운영 액션으로 간주하고 감사 로그 대상에 포함한다.
- 감사 로그 최소 항목:
  - 실행자
  - projection 종류
  - 범위 키
  - 실행 모드
  - 실행 사유
  - 결과(success/partial/fail)
  - mismatch 검증 결과
- 관리자 백오피스에서 직접 실행 가능한 범위와 내부 운영 CLI 전용 범위를 분리하는 것이 바람직하다.
  - 백오피스: 단일 객체/소규모 범위
  - 내부 CLI/Job: 기간/전체 리플레이

### 79.12 QA / 출시 게이트 반영 포인트
출시 전 아래 테스트를 포함해야 한다.

1. 특정 채팅방 projection 삭제 후 targeted rebuild로 복구 가능한가
2. 완료/분쟁 상태가 꼬인 trade thread를 recompute 했을 때 올바른 `nextPrimaryAction`이 복구되는가
3. 알림 projection stale 상태에서 딥링크 fallback이 정상 동작하는가
4. DLQ 이벤트 재처리 후 unread/배지 집계가 정상화되는가
5. 재빌드 중 사용자 화면이 치명 오류 없이 degraded read로 유지되는가

출시 게이트 연결:
- MVP 운영 준비에는 `targeted rebuild capability`를 최소 요건으로 포함하는 것이 바람직하다.
- full replay는 MVP 필수는 아니지만, 절차 문서와 테스트 환경 검증은 필요하다.

### 79.13 오픈 질문
- projection 대상별로 저장소를 하나의 DB/materialized table로 시작할지, 일부는 Redis/검색엔진/별도 read DB로 분리할지?
- `TradeThreadProjection` targeted rebuild를 동기 API로 허용할지, 항상 비동기 job으로 제한할지?
- DLQ 재처리 권한을 Senior Moderator 이상 백오피스에 줄지, 엔지니어링 운영 전용으로 둘지?
- shadow rebuild 시점에 실시간 consumer dual-write를 어디까지 허용할지?

## 80. 디바이스 상태 / 멀티디바이스 읽음·푸시 억제 정책(초안)
### 80.1 목표
- 사용자 단위 읽음 상태(`ChatParticipantState`)와 디바이스/세션 단위 연결 상태(`ChatDeviceState`)를 분리해, 멀티디바이스 환경에서도 unread/푸시/실시간 복구가 일관되게 동작하도록 한다.
- 같은 사용자가 여러 기기에서 앱을 열고 있을 때 과도한 중복 푸시를 줄이되, 실제로 메시지를 놓치는 상황은 최소화해야 한다.
- MVP와 Post-MVP에서 필요한 저장소 복잡도를 구분해, 초기 구현은 단순하게 시작하면서도 정밀한 억제 정책으로 확장 가능하게 한다.

### 80.2 핵심 원칙
1. **읽음의 진실원은 사용자 단위**다. 최종 unread 계산과 상대방 read 표시는 `ChatParticipantState.lastReadEventSequence`를 기준으로 한다.
2. **푸시 억제 판단은 디바이스 단위 보조 신호**다. 어떤 기기가 현재 화면을 보고 있는지, 연결이 살아 있는지, 최근 ack가 있었는지는 `ChatDeviceState`가 담당한다.
3. 푸시 억제는 `완전 정확성`보다 `놓치지 않음`을 우선한다. 애매하면 푸시를 보내는 편이 안전하다.
4. 디바이스 상태는 고휘발성 데이터이므로, 거래/분쟁 원천 기록처럼 장기 보관하지 않는다.
5. 사용자 UI와 운영 지표는 `읽음`과 `푸시 억제`를 혼동하지 않아야 한다. 푸시를 안 보냈다고 읽은 것은 아니다.

### 80.3 상태 계층 분리
| 계층 | 객체 | 진실원 역할 | 대표 책임 |
|---|---|---|---|
| 사용자 계층 | `ChatParticipantState` | 영속 진실원 | 마지막 읽음 이벤트, unread 계산, mute, 알림 선호 |
| 디바이스 계층 | `ChatDeviceState` | 휘발성 보조 상태 | 현재 열람 여부, 최근 수신/ack, 연결 상태, push 억제 힌트 |
| 세션/연결 계층 | SSE/WebSocket connection state | 일시적 런타임 상태 | 연결 생존 여부, reconnect/backfill 트리거 |

원칙:
- `ChatParticipantState`는 DB 영속 저장 기본안 유지
- `ChatDeviceState`는 MVP에서 TTL 저장소 기반을 우선 검토
- 런타임 connection state는 프로세스 메모리/edge connection manager에 존재할 수 있으나, 푸시 억제에 필요한 최소 snapshot은 `ChatDeviceState`에 반영 가능해야 한다.

### 80.4 `ChatDeviceState` 필드 계약(정교화)
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `deviceStateId` | 필수 | 내부 식별자 |
| `userId` | 필수 | 사용자 |
| `deviceId` | 필수 | 앱 설치/기기 식별자 |
| `sessionId` | 선택 | 앱 실행 세션 식별자 |
| `platform` | 선택 | iOS / Android / Web |
| `appState` | 필수 | `foreground` / `background` / `terminated_assumed` |
| `activeChatRoomId` | 선택 | 현재 전면 열람 중인 채팅방 |
| `lastSeenAt` | 필수 | 최근 heartbeat/ack 시각 |
| `lastEventAckSequence` | 선택 | 이 기기가 수신 확인한 마지막 이벤트 sequence |
| `pushTokenId` | 선택 | 현재 푸시 토큰 참조 |
| `pushEligible` | 필수 | 이 기기에 푸시 발송 대상이 될 수 있는지 |
| `notificationPreferenceSnapshot` | 선택 | 기기 로컬 푸시 허용/권한 상태 snapshot |
| `ttlExpiresAt` | 필수 | 상태 만료 시각 |

원칙:
- `lastEventAckSequence`는 사용자 read cursor가 아니라 **이 디바이스가 어디까지 받아봤는지**를 보는 보조 필드다.
- `activeChatRoomId`는 push 억제 정밀도를 높이기 위한 핵심 필드지만, stale 가능성을 전제로 TTL을 짧게 둬야 한다.
- `terminated_assumed`는 OS 강제 종료를 서버가 정확히 알 수 없으므로 TTL 경과/heartbeat 부재로 추정한 상태다.

### 80.5 MVP 저장 전략 권장안
#### 기본안
- `ChatParticipantState`: 관계형 DB 영속 저장
- `ChatDeviceState`: Redis 같은 TTL 저장소 또는 DB TTL 테이블 성격 저장소
- SSE 연결 상태: 앱 서버/connection manager 메모리 + 필요 시 `ChatDeviceState`로 요약 반영

#### 이유
- unread/read는 정합성이 중요하므로 DB 영속이 적합하다.
- device state는 빈번히 갱신되고 빠르게 만료되어야 하므로 TTL 저장소가 운영상 유리하다.
- MVP에서 `ChatDeviceState`를 완전 이력형 DB 테이블로 두면 write volume과 정리 작업 복잡도가 커질 수 있다.

#### Post-MVP 확장안
- 다기기 푸시 억제 정확도가 중요해지면 `device_presence_history` 또는 최소 이벤트 로그를 추가해 디버깅 가능성을 높일 수 있다.
- 단, 영속 이력도 개인정보/행동 추적 리스크를 고려해 짧은 보관기간을 기본으로 한다.

### 80.6 푸시 억제 결정 규칙
메시지/예약/완료 이벤트 발생 시 수신자 사용자에 대해 아래 순서로 판단한다.

1. 사용자 전역/채팅방 mute 여부 확인
2. 계정 제한/차단/채팅 잠금 상태 확인
3. `ChatParticipantState` 기준 unread 생성 여부 확인
4. 활성 `ChatDeviceState` 조회
5. 같은 `chatRoomId`를 foreground로 열고 있는 디바이스가 있으면 푸시 억제 후보
6. 단, 거래 필수 이벤트(T1)는 억제보다 인지를 우선하며, 조건부로 약한 푸시 또는 인앱만 유지
7. 어떤 디바이스도 신뢰 가능한 foreground 상태가 아니면 정상 푸시 발송

### 80.7 이벤트 우선순위별 푸시 억제 기본안
| 이벤트 tier | 예시 | 같은 채팅방 active foreground 존재 시 | 다른 화면 foreground만 존재 시 | 모든 기기 background/불명확 시 |
|---|---|---|---|---|
| `T1` | 완료 요청, 분쟁 소명 요청, 예약 임박 | 기본 억제하지 않음 또는 약한 local/in-app 우선 | 푸시 발송 | 푸시 발송 |
| `T2` | 새 메시지, 예약 제안/변경 | 푸시 억제 | 푸시 발송 | 푸시 발송 |
| `T3` | 후기 요청, 상태 점검 | 푸시 억제 가능 | 사용자 설정 따름 | 사용자 설정 따름 |
| `T4` | 운영 경고/제한 | 푸시 발송 우선 | 푸시 발송 | 푸시 발송 |

원칙:
- `T1`, `T4`는 사용자가 이미 앱을 열고 있더라도 놓치면 안 되므로 억제를 보수적으로 적용한다.
- 일반 메시지(`T2`)는 동일 채팅방을 실시간 열람 중이면 푸시 억제가 기본안이다.
- `activeChatRoomId`가 같더라도 `lastSeenAt`가 너무 오래되었으면 억제하지 않는다.

### 80.8 staleness / TTL 규칙
| 필드/상태 | 권장 TTL/판단 기준 | 설명 |
|---|---|---|
| `activeChatRoomId` foreground snapshot | 30~90초(가정) | 짧게 유지, heartbeat 부재 시 무효 |
| 일반 `lastSeenAt` device presence | 5~15분(가정) | background 상태 추정 유지 |
| `terminated_assumed` 전환 | foreground snapshot TTL 경과 시 | 강제 종료 정확성보다 안전한 fallback 우선 |
| 푸시 억제 허용 최대 stale | 60초 이하 권장 | 오래된 foreground 정보로 푸시를 막지 않음 |

원칙:
- stale device state 때문에 중요한 푸시를 놓치는 것이 가장 위험하다.
- 따라서 `activeChatRoomId` 기반 억제는 매우 최근 heartbeat/ack가 있는 경우에만 적용한다.
- TTL 경과 시 상태를 삭제하거나 `terminated_assumed`로 간주해 보수적으로 푸시를 다시 허용한다.

### 80.9 멀티디바이스 읽음/ack 규칙
- 한 디바이스에서 채팅방을 열고 읽음 커서를 올리면, 사용자 단위 `ChatParticipantState.lastReadEventSequence`가 갱신된다.
- 다른 디바이스는 다음 동기화 시 해당 읽음 커서를 받아 unread를 줄인다.
- 디바이스 단위 `lastEventAckSequence`는 다음 목적에만 사용한다.
  1. push 억제 보조 판단
  2. reconnect/backfill 시작점 추정
  3. 디버깅/관측성
- 상대방에게 노출되는 read 표시는 사용자 단위 read만 반영하고, 디바이스 ack는 외부 노출하지 않는다.

### 80.10 API / 이벤트 후보
#### 클라이언트 -> 서버
- `POST /chats/{chatRoomId}/device-presence`
- `POST /chats/{chatRoomId}/ack`
- `POST /devices/push-tokens` (기존 `POST /push-tokens` 확장 가능)
- `DELETE /devices/{deviceId}/push-token` 또는 logout/unregister 흐름

`device-presence` 요청 예시:
```json
{
  "deviceId": "ios_abc123",
  "sessionId": "sess_789",
  "appState": "foreground",
  "activeChatRoomId": "chat_123",
  "lastSeenAt": "2026-03-13T09:34:00+09:00"
}
```

원칙:
- presence heartbeat는 메시지마다 보내기보다 화면 진입/이탈, 앱 전경 전환, 일정 주기 heartbeat 조합이 바람직하다.
- `ack`는 read보다 더 자주 올 수 있으므로 고빈도 쓰기 부담을 고려해 coalescing이 필요하다.

### 80.11 관측성 / 운영 지표 후보
- `push_suppressed_count`
- `push_suppressed_but_not_opened_count`
- `push_sent_while_active_chat_open_count`
- `device_presence_stale_rate`
- `active_chat_presence_count`
- `ack_to_read_promotion_latency_ms`
- `multi_device_read_sync_lag_ms`

해석 원칙:
- `push_suppressed_count`만 높다고 좋은 것이 아니다. suppression 후 실제 열람이 이어졌는지 함께 봐야 한다.
- `push_suppressed_but_not_opened_count`가 높으면 억제 기준이 너무 공격적일 수 있다.
- `push_sent_while_active_chat_open_count`가 높으면 presence TTL 또는 foreground 판정이 너무 보수적일 수 있다.

### 80.12 QA 시나리오 파생 포인트
1. A기기에서 채팅방 열람 중, B기기 background 상태일 때 일반 메시지 푸시가 억제되는가
2. A기기 foreground snapshot TTL 만료 후 새 메시지 수신 시 푸시가 다시 발송되는가
3. 완료 요청(T1)은 active chat open 상태여도 사용자가 인지 가능한 방식으로 전달되는가
4. 한 기기에서 읽음 처리 후 다른 기기 unread가 합리적 시간 내 동기화되는가
5. logout/앱 삭제 후 stale push token/device state가 남아 중복 푸시를 만들지 않는가

### 80.13 오픈 질문
- MVP에서 `ChatDeviceState`를 Redis/TTL 저장소로만 둘지, 디버깅용 최소 영속 로그를 함께 남길지?
- `T1` 이벤트를 foreground active chat 상태에서도 푸시로 보낼지, 인앱 배너/local notification으로 대체할지?
- iOS/Android 앱이 foreground 상태 heartbeat를 어느 주기로 보낼지, 배터리 비용과 억제 정확도 사이 균형을 어디에 둘지?
- 하나의 사용자가 같은 채팅방을 두 기기에서 동시에 열람할 때 `activeChatRoomId` 충돌/우선순위를 어떻게 해석할지?

## 72. 피처 플래그 / 단계적 출시 / 롤백 정책(초안)
### 72.1 목표
- 거래 도메인 특성상 기능 하나의 오동작이 상태 정합성, 사용자 신뢰, 운영 부담으로 바로 이어질 수 있으므로, 주요 기능은 전면 공개보다 점진 배포를 기본으로 한다.
- 화면 노출 여부, 쓰기 허용 여부, 자동화 워커 동작 여부를 분리 제어해 “보이기만 하는 기능”, “실제 쓰기 가능한 기능”, “백그라운드 자동화 기능”을 별도로 롤아웃할 수 있어야 한다.
- 기능 플래그는 단순 개발 편의가 아니라 운영 정책과 출시 판단의 일부로 다뤄야 한다.

### 72.2 플래그 분류 원칙
| 플래그 유형 | 목적 | 예시 | 운영 주의사항 |
|---|---|---|---|
| `ui_visibility_flag` | 특정 화면/CTA 노출 제어 | 예약 카드 노출, 후기 CTA 노출 | UI만 열고 서버 쓰기 API는 닫혀 있을 수 있음 |
| `write_enable_flag` | 실제 쓰기 액션 허용 여부 제어 | 예약 생성, 완료 요청, 분쟁 소명 제출 | 서버가 최종 판단해야 함 |
| `automation_flag` | 배치/워커/자동 상태 전환 제어 | 예약 자동 만료, 완료 자동확정 | 재실행/중복처리 안전성 필요 |
| `policy_flag` | 정책 수준/임계치 제어 | reserved 신규문의 허용, 연락처 차단 강도 | 실험과 정책 혼용 시 로그 분리 필요 |
| `ranking_flag` | 검색/추천 가중치 제어 | reserved 감점 강화, 신규계정 감점 | 랭킹 변화는 KPI 영향 모니터링 필수 |
| `ops_tool_flag` | 운영 백오피스 기능 공개 제어 | 분쟁 해결 버튼, 대량 숨김 도구 | 잘못 열면 치명적이므로 권한과 별도 관리 |

원칙:
- 클라이언트는 플래그만 믿고 쓰기 동작을 낙관하면 안 되며, 서버 응답의 `availableActions`, `policyHints`, `restrictionReasonCode`를 최종 기준으로 사용한다.
- 고위험 기능은 `UI 노출 ON`보다 `쓰기 허용 ON`을 더 보수적으로 열어야 한다.

### 72.3 고위험 기능별 롤아웃 단위
| 기능 묶음 | 위험도 | 권장 롤아웃 순서 | 비고 |
|---|---|---|---|
| 채팅 기본 메시지 | 중간 | 내부 QA → 제한 베타 → 전체 | 미읽음/푸시 동기화 확인 필요 |
| 예약 생성/확정 | 높음 | 내부 QA → 직원/테스터 → 서버/카테고리 제한 공개 → 전체 | 상태 정합성 핵심 |
| 완료 요청/자동확정 | 매우 높음 | 수동 완료만 공개 → 자동확정은 shadow 모드 → 제한 공개 → 전체 | 분쟁/후기와 직접 연결 |
| 분쟁 소명/운영 큐 | 높음 | 운영 전용 준비 → 제한 사용자군 공개 | 운영자 교육 필요 |
| 외부 연락처 탐지/차단 | 높음 | 로그 수집 only → 경고 모드 → 차단 모드 | 오탐률 관리 필수 |
| 검색 랭킹 변경 | 중간 | shadow score → 소량 실험 → 전체 | 주요 KPI와 함께 관찰 |
| 푸시/리마인드 고도화 | 중간 | 인앱 only → 푸시 일부 이벤트 → 전체 | 소음/중복 억제 검증 필요 |

### 72.4 Shadow / Read-only / Limited Write 모드
복잡한 기능은 아래 3단계 공개 모델을 우선 검토한다.

1. **Shadow mode**
   - 사용자 경험에는 직접 영향 주지 않고, 서버 내부에서만 계산/로그를 남긴다.
   - 예: 자동확정 예정 시각 계산, reserved 악용 점수 계산, 연락처 탐지 confidence 산출
2. **Read-only exposure**
   - 사용자/운영자가 결과를 볼 수는 있지만 실제 상태 변경은 하지 않는다.
   - 예: 내 거래에서 `자동확정 예정` 배지 표시만 하고 자동확정 워커는 OFF
3. **Limited write exposure**
   - 일부 사용자군/서버/카테고리/플랫폼에서만 실제 쓰기 허용
   - 예: 특정 서버의 sell 매물에만 예약 기능 ON

원칙:
- `shadow -> read-only -> limited write -> full rollout` 순서를 기본 템플릿으로 삼는다.
- 거래 상태를 바꾸는 기능은 가능하면 shadow 로그 없이 바로 전체 공개하지 않는다.

### 72.5 플래그 타게팅 기준 후보
- 사용자군: 운영자, 내부 QA, 화이트리스트 베타 사용자, 신규 사용자 제외군
- 거래 유형: `sell`만, `buy`만, 특정 카테고리만
- 서버/지역: 특정 게임 서버 단위
- 플랫폼: iOS / Android / Web
- 계정 신뢰 수준: 프로필 완료 사용자만, 제재 이력 없는 사용자만
- 앱 버전: 최소 지원 버전 이상

원칙:
- 거래 상대가 서로 다른 플래그 상태일 때도 핵심 흐름이 깨지지 않아야 한다.
- 예: 한 사용자는 새 예약 카드 UI를 보고, 다른 사용자는 구버전 텍스트 fallback만 보더라도 같은 예약 객체를 읽을 수 있어야 한다.

### 72.6 롤백 기준과 안전 스위치
다음 경우는 부분 롤백 또는 즉시 kill switch 후보로 본다.
- 동일 매물에 복수 활성 예약/복수 완료 요청이 발생
- `GET /me/trades`와 채팅 상세의 상태 표시가 일관되지 않음
- 푸시 중복/누락으로 사용자가 예약 응답 기한을 반복적으로 놓침
- 외부 연락처 차단 오탐으로 정상 거래 등록/채팅 전환이 급감
- 분쟁/운영 액션이 감사 로그 없이 실행되거나 재현 불가
- 랭킹 변경 후 채팅 시작률/예약 전환율이 비정상 급락

Kill switch 원칙:
- 최소 아래 고위험 기능은 즉시 비활성화 가능한 전역 플래그를 가져야 한다.
  - 예약 생성/확정 쓰기
  - 완료 자동확정 워커
  - 분쟁 자동 큐 적재
  - 외부 연락처 전송 차단
  - 실험적 랭킹 모델
- kill switch 후에도 기존 데이터 조회와 기록 열람은 가능해야 한다.
- kill switch는 “새 액션 차단”이 목적이지, 기존 거래 기록을 숨기기 위한 수단이 아니다.

### 72.7 배포 체크리스트와 관측성 연결
주요 플래그를 여는 배포에는 아래 체크리스트를 함께 갖는다.

#### 사전 체크
- 상태 전이 테스트/멱등성 테스트 통과
- 관련 read model/projection 지연 시간 허용 범위 확인
- 운영자 매뉴얼/FAQ/사용자 안내 문구 준비
- 알림/딥링크/분석 이벤트가 함께 동작하는지 확인

#### 배포 직후 30~60분 모니터링
- API 에러율, 409 충돌률, 429 증가 여부
- 예약/완료/분쟁 이벤트 적재량 급증 여부
- projection rebuild backlog, dead-letter, worker retry 증가 여부
- 푸시 sent/delivered/opened/acted 퍼널 이상 여부

#### 배포 후 24시간 관찰
- 채팅 시작률, 예약 확정률, 완료 전환율 변화
- 신고율/차단율/분쟁율 변화
- 운영 처리 시간과 문의량 변화
- 오탐/오조치 복구 건수

### 72.8 운영자 커뮤니케이션 정책
- 정책/기능 플래그 변경으로 사용자 행동이 달라지는 경우, 운영자는 최소한 아래를 인지해야 한다.
  - 무엇이 바뀌었는가
  - 누구에게 열렸는가
  - 문제가 생기면 어디서 끌 수 있는가
  - 사용자 문의에 어떤 문구로 답해야 하는가
- 사용자 공지가 필요한 경우 예시:
  - 예약 기능 베타 오픈
  - 자동확정 정책 시작/변경
  - 외부 연락처 차단 강화
  - 후기 공개 규칙 변경

원칙:
- 운영팀이 모르는 기능은 출시하지 않는다는 기준을 둔다.
- 기능 플래그 상태는 백오피스 또는 내부 운영 대시보드에서 최소 읽기 가능해야 한다.

### 72.9 데이터/분석 설계 시사점
플래그 영향 분석을 위해 주요 이벤트에 아래 컨텍스트를 같이 남기는 것을 권장한다.
- `featureFlagSnapshot`
- `experimentBucket`
- `appVersion`
- `projectionVersion`
- `policyMode`

예시:
```json
{
  "eventName": "reservation_confirm",
  "listingId": "listing_123",
  "featureFlagSnapshot": {
    "reservation_v2_ui": true,
    "reservation_write_enabled": true,
    "auto_complete_shadow": false
  },
  "experimentBucket": "beta_a",
  "appVersion": "1.2.0"
}
```

원칙:
- 실험 효과와 제품 전반 추세를 구분하기 위해 플래그/실험 컨텍스트 없는 KPI 집계만으로 판단하지 않는다.
- 정책 플래그와 UI A/B 실험은 분석 축을 분리하는 것이 바람직하다.

### 72.10 QA 파생 포인트
- 동일 기능이 ON/OFF일 때 화면/서버 응답/알림이 어떻게 달라지는지 케이스화 필요
- 구버전 앱과 신버전 앱이 같은 거래 객체를 볼 때 fallback이 안전한지 검증
- kill switch 작동 후 신규 쓰기만 막히고 기존 기록 조회는 유지되는지 확인
- shadow mode 로그가 실제 쓰기 없이도 운영 검토에 쓸 수 있는 수준인지 확인
- 플래그 조합이 많아질 경우 우선순위 규칙(`ops emergency off > experiment on`)을 명확히 테스트해야 한다.

### 72.11 오픈 질문
- 플래그 관리는 자체 어드민 UI로 할지, 배포 환경변수/원격 설정 시스템으로 할지?
- 운영자가 직접 바꿀 수 있는 플래그 범위를 어디까지 허용할지?
- 예약/완료처럼 상태 정합성에 영향을 주는 플래그는 DB migration/version lock과 어떻게 연동할지?
- A/B 실험을 MVP부터 할지, 안정화 이후에만 허용할지?
- 사용자별 실험 버킷을 거래 상대 간 일관되게 맞출 필요가 있는지, 아니면 read fallback만으로 충분한지?

## 73. 홈 화면(Home Surface) 계약(초안)
### 73.1 목표
- 홈 화면을 단순 배너/추천 피드가 아니라 **거래 재개 + 탐색 시작 + 신뢰/안전 리마인드**를 동시에 수행하는 첫 화면으로 정의한다.
- 사용자가 앱을 열었을 때 `지금 해야 할 일`과 `지금 볼 만한 매물`을 같은 surface 안에서 우선순위 있게 보여줘 거래 성사율을 높인다.
- 홈 화면이 목록/내 거래/알림으로 흩어진 핵심 상태를 요약하되, 각 기능의 완전한 대체가 아니라 `빠른 재진입 허브` 역할을 하도록 한다.

### 73.2 설계 원칙
1. **액션 필요 우선**: 추천 매물보다 예약 응답/완료 확인/분쟁 소명 같은 행동 필요 카드를 먼저 보여준다.
2. **서버 컨텍스트 명확화**: 사용자가 어느 서버를 기준으로 탐색 중인지 항상 드러나야 한다.
3. **모바일 한 손 사용 최적화**: 상단 1~2스크린 안에 검색, 액션 필요, 추천 슬롯, 등록 CTA가 들어와야 한다.
4. **콜드스타트와 리텐션 동시 대응**: 신규 사용자는 탐색 시작을 돕고, 기존 거래 사용자는 미완료 흐름 복귀를 돕는다.
5. **설명 가능한 개인화**: 추천은 블랙박스 피드보다 `최근 본 서버`, `관심 카테고리`, `거래 가능 상태` 같은 설명 가능한 신호를 우선 사용한다.

### 73.3 홈 화면 핵심 모듈 구성
| 모듈 | 목적 | 우선순위 | 노출 조건 |
|---|---|---|---|
| 상단 검색/서버 선택 바 | 탐색 시작점 제공 | 항상 최상단 | 항상 |
| 액션 필요 트레이 | 예약 응답/완료 확인/분쟁 소명 복귀 | 최우선 | action required 거래 존재 시 |
| 빠른 등록 CTA | 판매/구매 매물 등록 진입 | 상위 | 로그인 회원 기본 |
| 추천 매물 섹션 | 탐색 전환 및 재방문 유도 | 중간 | 공개 매물 존재 시 |
| 찜/최근 본 재진입 | 이전 탐색 복구 | 중간 | 관련 데이터 존재 시 |
| 내 거래 요약 | 진행 중 거래 수/미읽음/임박 예약 요약 | 상위 | 로그인 회원 |
| 안전/정책 배너 | 사기 예방, 기능 제한, 운영 공지 | 조건부 | 고위험/제재/중요 공지 시 |
| 빈 상태 온보딩 모듈 | 서버 선택, 첫 검색, 첫 등록 유도 | 조건부 | 콜드스타트/데이터 없음 |

원칙:
- 홈은 무한 스크롤 피드보다 **명확한 모듈형 레이아웃**을 우선한다.
- 액션 필요가 없는 경우에만 추천/탐색 모듈 비중을 더 크게 가져간다.

### 73.4 사용자 상태별 홈 우선순위 시나리오
| 사용자 상태 | 홈 1순위 | 홈 2순위 | 홈 3순위 |
|---|---|---|---|
| 비회원 | 서버 선택 + 검색 | 추천 매물 | 로그인 유도 CTA |
| 신규 회원(매물/채팅 없음) | 서버 선택 + 첫 탐색 | 첫 매물 등록 안내 | 인기 카테고리/추천 매물 |
| 기존 탐색 사용자 | 최근 본/찜한 매물 복귀 | 추천 매물 | 등록 CTA |
| 진행 중 거래 사용자 | 액션 필요 트레이 | 내 거래 요약 | 추천 매물 |
| 고빈도 거래자 | 액션 필요 트레이 | 내 매물/내 거래 요약 | 서버별 추천/재등록 CTA |
| 제재/제한 사용자 | 제한 사유 배너 | 읽기 가능한 내 거래/알림 | 정책 안내 |

원칙:
- 홈 personalization은 `광고성 추천`보다 `거래 맥락 복구`를 우선한다.
- 동일 사용자라도 시간대/상태 변화에 따라 홈 첫 모듈이 달라질 수 있다.

### 73.5 액션 필요 트레이(Action Required Rail) 계약
목적: 앱 진입 직후 가장 중요한 거래 행동을 놓치지 않게 한다.

카드 후보:
- 예약 응답 필요
- 거래 1시간 이내 임박
- 상대 완료 요청 확인 필요
- 분쟁 소명 제출 필요
- reserved 장기 유지 점검 필요

카드 필수 필드:
- `tradeThreadId`
- `priorityLabel`
- `actionLabel`
- `deadlineLabel`
- `listingSummary`
- `counterpartySummary`
- `deepLink`

노출 규칙:
- 최대 3건 우선 노출, 초과분은 `모두 보기`로 내 거래 화면 이동
- 우선순위는 `분쟁 소명 > 완료 확인 > 예약 응답 > 거래 임박 > 상태 점검` 순서를 기본안으로 둔다.
- 홈과 내 거래의 우선순위 기준은 동일하되, 홈은 더 강하게 축약된 카드 표현을 사용한다.

### 73.6 추천 매물 섹션 구성 원칙
추천 매물은 목적별 슬롯으로 분리하는 것이 바람직하다.

| 섹션 | 의미 | 추천 신호 예시 |
|---|---|---|
| `selected_server_fresh` | 현재 선택 서버의 최신/거래 가능 매물 | serverId, available, freshness |
| `because_you_viewed` | 최근 본 아이템/카테고리 기반 | recent views, favorites |
| `matching_buy_or_sell_intent` | 사용자의 최근 사기/팔기 의도 기반 | recent listingType actions |
| `reopened_or_back_available` | reserved/대기에서 다시 가능해진 매물 | status recovery, availability recovery |
| `high_response_sellers` | 응답성/신뢰 우수 작성자 중심 | response badge, trust badge |

원칙:
- 추천 섹션 이름은 사용자에게 의미가 전달되도록 구성한다. 예: `데포로쥬에서 지금 거래 가능한 매물`, `최근 본 아이템과 비슷한 매물`
- `completed`, `cancelled`, `pending_trade` 매물은 홈 추천 기본 노출 대상에서 제외한다.
- `reserved`는 기본 추천에서는 후순위이며, 상태를 명확히 표시할 수 있을 때만 일부 슬롯에 제한 노출한다.

### 73.7 서버 컨텍스트 / 필터 유지 정책
- 홈 상단에는 항상 현재 서버 컨텍스트가 보여야 한다.
- 서버 선택 기본 우선순위:
  1. 사용자의 수동 선택값
  2. 프로필 primary server
  3. 최근 검색/조회 서버
  4. 인기 서버 기본값(가정)
- 홈에서 서버를 바꾸면 추천 섹션, 빠른 필터, 검색 진입 기본값이 함께 갱신되어야 한다.
- 비회원도 서버 컨텍스트를 유지해야 하며, 로그인 후에도 가능한 한 복원해야 한다.

후보 UI:
- 서버 pill selector
- 최근 사용 서버 드롭다운
- `전체 서버`는 MVP에서 비권장, Post-MVP 후보

### 73.8 빠른 탐색 모듈 계약
목적: 검색어를 아직 정하지 못한 사용자에게 탐색 시작점을 제공한다.

후보 모듈:
- 인기 카테고리 chip
- 최근 많이 거래된 아이템 chip
- 최근 검색어/최근 본 매물
- `판매글 보기` / `구매글 보기` intent toggle

원칙:
- 홈 quick filter는 목록 화면의 정교한 필터를 완전히 대체하지 않는다.
- 홈에서 선택한 quick filter는 목록 진입 시 query context로 전달되어야 한다.
- 필터 텍스트는 내부 enum이 아니라 사용자가 이해 가능한 언어여야 한다.

### 73.9 콜드스타트 / 빈 상태 정책
#### 비회원/신규 회원
- 홈에 아래 순서로 노출:
  1. 서버 선택 유도
  2. 인기 카테고리
  3. 샘플 추천 매물
  4. 로그인/첫 매물 등록 CTA

#### 로그인 회원이지만 활동 데이터 없음
- `무엇을 찾고 계신가요?` 검색 시작 CTA
- `팔고 있다면 3분 안에 등록` 같은 등록 유도 CTA
- 추천은 전체 인기보다 현재 서버 인기/최신 매물 중심으로 단순화

#### 데이터 오류/부분 실패
- 액션 필요 데이터 실패와 추천 데이터 실패를 분리한다.
- 거래 재개 모듈 실패 시 추천만 보여주는 것으로 대체하지 말고, 명시적 재시도 영역을 둔다.

### 73.10 안전/운영 배너 노출 규칙
홈 배너는 마케팅 배너보다 안전/운영 우선이어야 한다.

배너 유형 후보:
- 외부 연락처/선입금 주의 안전 배너
- 계정 제한/경고 배너
- 분쟁 추가 자료 제출 요청 배너
- 예약 임박 주의 배너
- 운영 공지/점검 배너

우선순위 원칙:
1. 계정 상태/제재 관련
2. 분쟁/거래 실행 관련
3. 안전 리마인드
4. 일반 공지
5. 프로모션성 배너(Post-MVP)

### 73.11 홈 화면 API / read model 시사점
후보 엔드포인트:
- `GET /home`
- 또는 `GET /me/home` + 비회원용 `GET /home/public`

권장 응답 묶음:
```json
{
  "viewer": {
    "isLoggedIn": true,
    "homeMode": "member_active_trade",
    "selectedServerId": "dep-01"
  },
  "actionRequired": [
    {
      "tradeThreadId": "tt_123",
      "actionCode": "confirm_completion",
      "actionLabel": "거래완료 확인하기",
      "deadlineLabel": "오늘 18:00까지"
    }
  ],
  "quickActions": [
    {"code": "create_sell_listing", "label": "판매글 등록"},
    {"code": "create_buy_listing", "label": "구매글 등록"}
  ],
  "sections": [
    {
      "sectionType": "selected_server_fresh",
      "title": "데포로쥬에서 지금 거래 가능한 매물",
      "items": []
    }
  ],
  "policyHints": []
}
```

read model 후보:
- `HomeSurfaceProjection`
- `HomeActionRequiredProjection`
- `HomeRecommendationSectionProjection`

원칙:
- 홈은 여러 서비스의 fan-in 결과이므로 API 응답에서 섹션 단위 partial failure 식별이 가능해야 한다.
- 추천 섹션은 서버가 title/sectionType/items를 함께 내려주어 앱이 모듈 순서를 단순 렌더링할 수 있게 하는 편이 좋다.

### 73.12 홈 딥링크 / 복귀 정책
- 푸시, 알림함, 배지 클릭은 가능한 한 홈이 아니라 **해당 tradeThread/listing/chat**으로 직접 이동해야 한다.
- 단, 앱 cold start 또는 로그인 복구 중에는 홈을 중간 landing으로 사용할 수 있으며, 이때 `pendingDestination`을 유지해야 한다.
- 홈에서 목록/상세/내 거래로 이동 후 뒤로가기 시, 서버 선택과 스크롤 위치는 가능한 한 복원해야 한다.

### 73.13 분석 이벤트 후보
- `home_view`
- `home_action_required_impression`
- `home_action_required_click`
- `home_quick_action_click`
- `home_section_impression`
- `home_section_item_click`
- `home_server_change`
- `home_empty_state_cta_click`

핵심 KPI 연결 포인트:
- 홈 진입 후 5분 내 `채팅`, `예약 응답`, `완료 확인` 행동 전환율
- 액션 필요 카드 노출 대비 실제 행동 수행률
- 추천 섹션별 상세 진입률, 채팅 시작률
- 신규 사용자 홈 진입 후 첫 검색 또는 첫 등록 전환율

### 73.14 오픈 질문
- 홈 추천을 완전 개인화할지, 서버 중심 반개인화로 시작할지?
- 홈에서 `reserved` 매물을 얼마나 적극적으로 보여줄지?
- 비회원 홈과 로그인 홈을 같은 엔드포인트에서 처리할지 분리할지?
- 홈에 찜/최근 본 섹션을 MVP부터 넣을지, 이후 단계로 미룰지?
- 안전/운영 배너와 추천 섹션이 경쟁할 때 최대 몇 개까지 노출할지?


## 68. 가격 제안(Offer) 화면/데이터/API 계약(초안)
### 68.1 목표
- 가격 제안(`offer`)과 역제안(`counter-offer`)을 단순 채팅 메시지가 아니라 **거래 실행 액션**으로 정의한다.
- 사용자가 숫자 협상을 하다가도 현재 유효한 제안, 응답 마감, 마지막 제안자를 한눈에 이해할 수 있어야 한다.
- 매물 상세/채팅/내 거래/알림/운영 화면이 동일한 Offer 상태 모델을 사용해 혼선을 줄인다.

### 68.2 Offer 도메인 역할 정의
| 개념 | 의미 | 생성 주체 |
|---|---|---|
| `initial_offer` | 최초 가격 제안 | 매물 비작성자 또는 작성자 모두 가능 |
| `counter_offer` | 기존 제안에 대한 가격 수정 응답 | 현재 응답권자 |
| `offer_accept` | 제안 수락 | 제안을 받은 상대 |
| `offer_decline` | 제안 거절 | 제안을 받은 상대 |
| `offer_expire` | 응답 기한 만료 | 시스템 |
| `offer_withdraw` | 제안자가 수락 전 자발 철회 | 마지막 제안자 |

원칙:
- 한 시점의 유효 제안은 **채팅방 기준 최대 1개**를 기본안으로 한다.
- Offer는 채팅 메시지로도 보이지만, 원본 상태는 별도 Offer 객체가 소유한다.
- 같은 채팅방의 여러 메시지가 하나의 Offer를 설명할 수 있어도, 상태 변경의 원천은 Offer 이벤트다.

### 68.3 Offer 상태 모델
| offerStatus | 의미 | 다음 가능 액션 |
|---|---|---|
| `proposed` | 최초 제안 발신됨, 상대 응답 대기 | 수락, 거절, 역제안, 철회 |
| `countered` | 상대가 역제안함, 원제안자 응답 대기 | 수락, 거절, 재역제안 |
| `accepted` | 한쪽 제안이 수락됨 | 예약 제안, 거래 진행 |
| `declined` | 명시 거절됨 | 새 제안 가능 |
| `withdrawn` | 제안자가 수락 전 철회 | 새 제안 가능 |
| `expired` | 응답 기한 만료 | 새 제안 가능 |
| `superseded` | 더 최신 제안이 생겨 비활성화됨 | 없음 |

보충 원칙:
- `countered`는 새로운 제안이 유효 상태가 되었음을 뜻하며, 직전 Offer는 `superseded`로 본다.
- `accepted`는 곧바로 예약 확정과 동일하지 않다. 다만 예약 제안 CTA를 강하게 올리는 신호다.
- `accepted` 이후 가격 변경이 필요하면 새 Offer를 생성할 수 있으나, `pending_trade` 이후에는 금지하는 기본안을 권장한다.

### 68.4 Listing/Reservation 상태와의 연동 규칙
| 상황 | Offer 허용 여부 | 규칙 |
|---|---|---|
| `Listing.status=available` | 허용 | 일반 협상 가능 |
| `Listing.status=reserved` + 우선상대 채팅 | 허용 | 우선상대와의 가격 재협상 허용 |
| `Listing.status=reserved` + 대기문의 채팅 | 제한적 허용 또는 차단 | MVP 기본안은 `문의는 가능, 가격제안은 차단` 권장 |
| `Listing.status=pending_trade` | 기본 금지 | 거래 직전 혼선 방지 |
| `Listing.status=completed/cancelled` | 금지 | 종결 상태 |
| 활성 Reservation 존재 + `scheduledAt` 임박 | 제한 | 예: 1시간 이내 새 Offer 금지 가정 |

권장 기본안:
- `accepted`된 Offer가 있어도 Listing 상태를 자동 `reserved`로 바꾸지는 않는다.
- 다만 `accepted` 직후 채팅/내 거래/상세에서 `예약으로 이어가기` CTA를 최상단 노출한다.
- 실제 우선 거래 보호는 Reservation/Listing 상태가 담당하고, Offer는 **협상 기록과 실행 유도 신호**로 둔다.

### 68.5 채팅 화면 계약
#### 68.5.1 메시지 타입
- `offer_card`
- `offer_counter_card`
- `offer_accepted_system`
- `offer_declined_system`
- `offer_expired_system`
- `offer_withdrawn_system`

#### 68.5.2 카드 필수 노출 요소
- 제안 금액
- 제안 방향(`내가 제안`, `상대가 제안`)
- 제안 시각
- 응답 마감 시각 또는 상대시간
- 현재 상태 배지
- 대표 CTA 1개
- 보조 CTA 1~2개

#### 68.5.3 상태별 CTA 규칙
| 상태 | 제안 수신자 CTA | 제안 발신자 CTA |
|---|---|---|
| `proposed` | 수락, 거절, 역제안 | 철회, 상태 보기 |
| `countered` | 수락, 거절, 재역제안 | 상태 보기 |
| `accepted` | 예약 제안/채팅 보기 | 예약 제안/채팅 보기 |
| `declined` | 새 제안 | 새 제안 |
| `expired` | 새 제안 | 새 제안 |
| `withdrawn` | 새 제안 | 새 제안 |

화면 원칙:
- 채팅 입력창의 자유 텍스트와 Offer 액션은 시각적으로 구분해야 한다.
- 숫자 협상이 길어질수록 과거 제안 전체보다 **현재 유효 제안 1개**를 상단 sticky 또는 인라인 강조하는 것이 바람직하다.
- `accepted`된 Offer가 있으면 입력창 근처에 `이 가격으로 예약 제안하기` CTA를 노출한다.

### 68.6 매물 상세 / 내 거래 / 홈 노출 규칙
#### 매물 상세
- 현재 사용자가 참여 중인 채팅방에서 유효 Offer가 있으면 상세 상단 CTA 영역에 요약 배너를 노출할 수 있다.
- 예: `상대가 48만 아데나를 제안했어요 · 오늘 18:00까지 응답`
- 작성자 본인 상세에서는 `최근 협상중` 상태를 볼 수 있으나, 제3자에게는 개별 Offer 금액을 노출하지 않는다.

#### 내 거래
- `tradeThreadStatus`와 별개로 `activeOfferSummary`를 보조 정보로 둘 수 있다.
- 액션 필요 정렬에서 `상대 가격 제안 응답 대기`는 일반 미읽음보다 우선한다.
- 예시 문구:
  - `48만 아데나 제안에 응답 필요`
  - `내 역제안 응답 대기`

#### 홈
- 홈의 액션 필요 카드에 Offer 응답 대기를 포함할 수 있다.
- 단, 홈에서는 상세 금액 전체보다 `가격 제안 응답 필요` 수준의 요약 카드가 우선이며, 민감한 거래 상세는 내 거래 상세 진입 후 확인하게 한다.

### 68.7 Offer 데이터 모델 후보
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `offerId` | 필수 | Offer 식별자 |
| `listingId` | 필수 | 관련 매물 |
| `chatRoomId` | 필수 | 관련 채팅방 |
| `offerRound` | 필수 | 같은 채팅방 협상 라운드 번호 |
| `parentOfferId` | 선택 | 역제안 대상 Offer |
| `proposedByUserId` | 필수 | 현재 제안 발신자 |
| `targetUserId` | 필수 | 현재 응답 대상 |
| `amount` | 필수 | 제안 금액 |
| `currencyType` | 필수 | 통화 타입 |
| `offerStatus` | 필수 | `proposed` / `countered` / ... |
| `expiresAt` | 선택 | 응답 마감 시각 |
| `acceptedAt` | 선택 | 수락 시각 |
| `declinedAt` | 선택 | 거절 시각 |
| `withdrawnAt` | 선택 | 철회 시각 |
| `supersededByOfferId` | 선택 | 대체한 최신 Offer |
| `createdAt` | 필수 | 생성 시각 |

제약 후보:
- `chatRoomId` 기준 `offerStatus in (proposed, countered)` 활성 Offer는 최대 1개
- `parentOfferId`가 있으면 동일 `chatRoomId`와 양 당사자 조합을 공유해야 함
- `accepted` 이후 활성 Reservation/Completion 정책과 충돌하는 새 Offer 생성은 차단 가능

### 68.8 API 후보
#### 사용자용
- `POST /chats/{chatRoomId}/offers`
- `POST /offers/{offerId}/accept`
- `POST /offers/{offerId}/decline`
- `POST /offers/{offerId}/counter`
- `POST /offers/{offerId}/withdraw`
- `GET /chats/{chatRoomId}/offers`
- `GET /offers/{offerId}`

#### 응답 예시
```json
{
  "offerId": "offer_123",
  "chatRoomId": "chat_123",
  "offerStatus": "proposed",
  "amount": 480000,
  "currencyType": "adena",
  "expiresAt": "2026-03-13T18:00:00+09:00",
  "directionForViewer": "incoming",
  "availableActions": ["accept_offer", "decline_offer", "counter_offer"],
  "nextActionLabel": "오늘 18:00까지 응답"
}
```

### 68.9 만료 / 중복 / 동시성 규칙
- 기본안: Offer 응답 기한은 생성 후 12시간 또는 예약 시각 이전 등 정책식으로 계산 가능
- 상대가 거의 동시에 `accept`와 `counter`를 보내면, 먼저 확정된 액션만 성공하고 나머지는 409 충돌 처리
- `counter` 성공 시 직전 활성 Offer는 `superseded`
- 네트워크 재전송 대비 `POST /offers`, `/accept`, `/decline`, `/counter`, `/withdraw`는 멱등 키 지원 권장
- `accepted` 직후 한쪽이 새 Offer를 보내려면 예약/거래 상태와 충돌 검증을 먼저 수행해야 한다.

### 68.10 운영 / 안전 정책 연계
- 비정상 협상 패턴은 운영 신호가 된다.
  - 짧은 시간 반복 고저가 흔들기
  - 다수 채팅방에 동일 금액 스팸 제안
  - `accepted` 후 반복 철회
  - 예약 직전 과도한 가격 재협상
- 운영 화면에는 아래 요약이 보이면 좋다.
  - 최근 7일 Offer 생성/거절/수락 수
  - accepted 후 취소율
  - counter 횟수 분포
  - 다중 상대 병렬 제안 패턴
- 가격 자체의 적정성 판단보다는 **괴롭힘/낚시/시간 끌기 패턴** 탐지가 더 중요하다.

### 68.11 분석 이벤트 후보
- `offer_create`
- `offer_accept`
- `offer_decline`
- `offer_counter`
- `offer_withdraw`
- `offer_expire`
- `offer_to_reservation_convert`

핵심 분석 포인트:
- 상세 진입 후 Offer 생성률
- Offer 수락률 / 역제안률 / 만료율
- Offer 수락 후 예약 전환율
- `reserved` 상태 채팅에서 Offer 허용 여부가 성사율에 미치는 영향
- 과도한 역제안 횟수가 실제 완료율을 낮추는지 여부

### 68.12 후속 파생 문서 포인트
- 화면명세:
  - Offer 카드 컴포넌트, sticky 협상 바, 입력 모달, 만료 카운트다운
- DB 문서:
  - 활성 Offer 유니크 제약, superseded 관계, 라운드 번호 전략
- API 명세:
  - 액션별 409/422 규칙, idempotency, `directionForViewer`, `availableActions`
- 운영정책:
  - 협상 스팸, accepted 후 반복 철회, 예약 직전 가격변경 악용 대응 기준


## 68. 가격 변경 / 가격 이력 / 가격 노출 정책(초안)
### 68.1 목표
- 가격은 단순 숫자 필드가 아니라 **탐색 전환, 협상 시작, 찜 사용자 재유입, 운영 이상징후 탐지**에 모두 영향을 주는 핵심 도메인 신호로 다룬다.
- 사용자가 가격을 자주 바꾸더라도 거래 상대와 찜 사용자가 혼란스럽지 않게, `현재 가격`, `협상 가능 여부`, `최근 변경 이력`의 의미를 분리한다.
- 화면, DB, API, 알림, 랭킹이 동일한 가격 변경 이벤트를 기준으로 동작하도록 공통 계약을 정의한다.

### 68.2 기본 원칙
1. 목록/상세/채팅/내 거래에서 보이는 **현재 가격 표기**는 항상 동일한 원천 필드에서 계산되어야 한다.
2. 가격 변경은 단순 `Listing.priceAmount` overwrite가 아니라, 필요 시 이력 이벤트를 남기는 **의미 있는 상태 변화**로 취급한다.
3. `offer` 기반 매물과 `fixed/negotiable` 매물은 가격 변경 UX와 알림 규칙이 달라야 한다.
4. 거래 진행 중(`reserved`, `pending_trade`) 가격 변경은 허용 범위를 보수적으로 제한해 분쟁을 줄인다.
5. 가격 변경은 검색 랭킹/찜 알림에 활용할 수 있지만, 상단 노출을 위한 잦은 가격 장난은 감점/제한 대상이다.

### 68.3 가격 표현 모델
| 필드/개념 | 의미 | 사용자 노출 예시 |
|---|---|---|
| `priceType` | 가격 제시 방식 | `즉시 거래가`, `가격 협의`, `제안받기` |
| `currentPriceAmount` | 현재 대표 가격 숫자 | `500,000` |
| `currentPriceLabel` | 화면 표시용 라벨 | `50만 아데나`, `가격 협의`, `제안받기` |
| `priceVisibility` | 현재 가격 공개 범위 | `public`, `member_only`, `chat_only` 후보 |
| `lastPriceChangedAt` | 마지막 가격 변경 시각 | `3시간 전 가격 수정` |
| `priceChangeDirection` | 직전 변경 방향 | `up`, `down`, `none` |
| `priceChangeDeltaAmount` | 직전 변경 폭 | `-50,000` |
| `negotiationHint` | 협상 가능성 요약 | `가격 제안 가능`, `고정가 우선` |

원칙:
- `currentPriceLabel`은 목록/상세/푸시/홈 카드에서 재사용 가능한 파생 필드로 본다.
- `priceVisibility`는 공개 가격 정책 확장 여지를 위한 필드이며, MVP는 `public` 고정안으로 시작할 수 있다.
- `offer` 매물은 `currentPriceAmount`가 null일 수 있으나, 이후 제안/역제안과 별개로 `희망 범위`를 도입할지 여부는 오픈 질문으로 남긴다.

### 68.4 매물 유형별 가격 계약
| listingType | priceType | 기본 의미 | 목록 표시 | 상세 CTA/보조문구 |
|---|---|---|---|---|
| `sell` | `fixed` | 판매자가 즉시 거래 희망 가격 제시 | 가격 숫자 강조 | `채팅 시작`, 필요 시 `가격 제안` |
| `sell` | `negotiable` | 기준 가격은 있으나 협상 허용 | 가격 + `협의 가능` 배지 | `채팅`, `제안하기` |
| `sell` | `offer` | 판매자가 공개 가격 없이 제안 수집 | `제안받기` | `가격 제안 보내기` |
| `buy` | `fixed` | 구매자가 지불 의향 가격 제시 | `구매 희망가` 라벨 | `판매 제안`, `채팅` |
| `buy` | `negotiable` | 희망 가격은 있으나 협의 가능 | 희망가 + `협의 가능` | `가격 제안`, `채팅` |
| `buy` | `offer` | 구매자가 공개 가격 없이 판매자 제안 대기 | `가격 제안 기다림` | `판매 조건 제안` |

원칙:
- `buy` 매물의 가격은 `판매가`가 아니라 `구매 희망가`라는 의미를 UI에서 분명히 구분해야 한다.
- 동일 `priceAmount`라도 `listingType`에 따라 사용자 해석이 달라지므로, 화면/알림/분석 이벤트에 `listingType`이 항상 함께 있어야 한다.

### 68.5 상태별 가격 수정 허용 규칙
| Listing.status | 가격 변경 허용 여부 | 허용 범위 기본안 | 사용자 안내 원칙 |
|---|---|---|---|
| `available` | 허용 | 자유 수정 가능, 단 rate limit 적용 | 즉시 반영 |
| `reserved` | 제한 허용 | 우선 거래 상대가 없는 대기 문의 구간에서만 제한적 허용 후보 | 현재 대화 상대에게 변경 사실 명시 |
| `pending_trade` | 원칙적 금지 | 운영 예외 또는 예약 취소 후만 가능 | 거래 직전 가격 변경 불가 |
| `completed` | 불가 | 없음 | 기록 보존 |
| `cancelled` | 직접 수정 비권장 | 복제 후 재등록 권장 | 과거 기록 유지 |

세부 원칙:
- `reserved` 상태에서 가격을 바꾸면 현재 우선 상대에게는 **자동 시스템 메시지 + 변경 전/후 가격 비교**가 필요하다.
- `pending_trade` 상태의 가격 변경은 거래 조건 뒤집기로 해석될 수 있으므로 MVP에서는 차단이 안전하다.
- `offer` 매물에서 공개 가격을 새로 입력해 `fixed/negotiable`로 전환하는 것은 가능하나, 활성 제안과의 우선순위 규칙이 필요하다.

### 68.6 가격 변경 이벤트 모델
가격 변경은 이벤트 이력으로 남겨야 하는 대표 행위다.

후보 엔티티:
- `ListingPriceSnapshot`
- `ListingPriceChangeEvent`

#### `ListingPriceSnapshot` 후보 필드
| 필드 | 설명 |
|---|---|
| `listingId` | 매물 식별자 |
| `priceType` | 당시 가격 방식 |
| `priceAmount` | 당시 가격 값 |
| `priceCurrency` | 통화 단위 |
| `capturedAt` | 스냅샷 시각 |
| `capturedReason` | `create`, `user_update`, `offer_accept_ref`, `admin_adjustment` 등 |

#### `ListingPriceChangeEvent` 후보 필드
| 필드 | 설명 |
|---|---|
| `priceChangeEventId` | 이벤트 식별자 |
| `listingId` | 매물 식별자 |
| `beforePriceType` / `afterPriceType` | 변경 전/후 가격 방식 |
| `beforePriceAmount` / `afterPriceAmount` | 변경 전/후 가격 값 |
| `changeDirection` | `up`, `down`, `lateral`, `type_only` |
| `deltaAmount` | 변경 폭 |
| `changedByUserId` | 변경 주체 |
| `changedByActorType` | `owner`, `system`, `admin` |
| `changeReasonCode` | `seller_adjustment`, `buyer_adjustment`, `stock_remainder`, `market_repricing`, `policy_correction` 등 후보 |
| `createdAt` | 변경 시각 |

원칙:
- 단순 overwrite만 남기면 찜 알림, 급락 배지, 운영 분석, 분쟁 추적이 어려우므로 최소 이벤트 로그는 필요하다.
- 가격 변경은 감사/알림/랭킹 모두에서 재활용되므로 별도 이력 객체를 우선 검토한다.

### 68.7 화면 노출 계약
#### 목록 카드
- 현재 가격 라벨 우선 노출
- 최근 가격 인하가 있으면 `가격 인하` 또는 하향 배지 노출 가능
- 가격 상승은 기본 목록에서 굳이 강조하지 않되, 상세에서는 이력 확인 가능하게 한다.
- `offer` 매물은 숫자 대신 `제안받기` / `희망 조건 협의` 식 문구를 사용한다.

#### 매물 상세
- 현재 가격, 가격 방식, 마지막 수정 시점, 협상 가능 여부를 명확히 보여야 한다.
- 찜 사용자가 가격 변경으로 재방문했을 때 `이전 가격 대비 변화`를 확인할 수 있는 보조 라벨을 검토한다.
- `buy` 매물은 `구매 희망가`, `sell` 매물은 `판매가`로 레이블을 분리한다.

#### 채팅/내 거래
- 거래 진행 중 가격이 바뀌면 일반 텍스트가 아니라 시스템 이벤트 카드로 남겨야 한다.
- 예시 문구:
  - `판매자가 가격을 55만 → 50만 아데나로 변경했어요`
  - `구매 희망가가 45만 → 48만 아데나로 조정되었어요`
- 거래 상세에는 **현재 거래가 어떤 가격 기준에서 진행 중인지**를 요약 패널에 항상 표시해야 한다.

### 68.8 찜/알림 연계 규칙
| 이벤트 | 수신자 | 발송 조건 | 억제 조건 |
|---|---|---|---|
| 가격 인하 | 찜 사용자 | `available` 상태에서 의미 있는 가격 하향 | 너무 잦은 변경, 사용자 알림 OFF |
| 가격 상승 | 기본 비발송 | 원칙적으로 푸시 없음 | 인앱 기록만 후보 |
| `offer -> fixed/negotiable` 전환 | 찜 사용자, 기존 문의자 | 공개 가격 생김 | 이미 거래 종결 상태 |
| 거래 진행 중 가격 변경 | 현재 우선 상대/참여자 | `reserved` 상태 변경 발생 | 동일 세션에서 이미 확인 |

기본안:
- 가격 인하 알림은 **재유입 가치가 높으므로 허용**, 가격 상승 알림은 기본 비허용으로 시작한다.
- 가격 인하의 `의미 있는 변화` 임계값은 절대값 또는 비율 기준(예: 5% 이상, 가정)을 둘 수 있다.
- 동일 매물에서 짧은 시간 내 여러 번 가격을 바꿔도 푸시는 묶음/최신값 기준으로 1회만 보내는 것이 바람직하다.

### 68.9 랭킹/악용 방지 규칙
- 가격 인하만으로 최신 매물처럼 과도하게 상단 재노출되면 안 된다.
- 가격 변경은 추천 랭킹의 보조 신호일 수 있으나, `새 활동` 가중치는 제한적으로 부여한다.
- 아래 패턴은 운영/랭킹 감점 후보다.
  1. 짧은 시간에 가격을 반복 상하 조정
  2. 의미 없는 소폭 변경으로 찜 알림 재발송 유도
  3. `offer ↔ fixed` 타입 전환 반복으로 관심 끌기
  4. `reserved` 직전/직후 가격 급변으로 상대 압박

후보 내부 지표:
- `priceChangeCount7d`
- `priceChangeNotificationCount24h`
- `suspiciousPriceFlipFlag`
- `priceVolatilityBand`

### 68.10 API 후보
#### 사용자용
- `POST /listings/{listingId}/price`
- `GET /listings/{listingId}/price-history`

#### 응답 예시
```json
{
  "listingId": "listing_123",
  "price": {
    "priceType": "negotiable",
    "priceAmount": 500000,
    "priceLabel": "50만 아데나",
    "lastPriceChangedAt": "2026-03-13T09:50:00+09:00",
    "change": {
      "direction": "down",
      "deltaAmount": 50000,
      "deltaLabel": "5만 아데나 인하"
    }
  },
  "availableActions": ["create_chat", "make_offer"],
  "policyHints": []
}
```

```json
{
  "items": [
    {
      "changedAt": "2026-03-13T09:50:00+09:00",
      "before": {"priceType": "negotiable", "priceAmount": 550000, "label": "55만 아데나"},
      "after": {"priceType": "negotiable", "priceAmount": 500000, "label": "50만 아데나"},
      "changeReasonCode": "seller_adjustment"
    }
  ]
}
```

설계 원칙:
- 일반 `PATCH /listings/{listingId}`에 가격 변경을 섞을 수도 있으나, 가격 이벤트/알림/운영 룰을 분리하기 위해 별도 endpoint가 더 명확하다.
- 최소한 서버 내부에서는 가격 변경을 독립 이벤트로 취급해야 한다.

### 68.11 분석 이벤트 후보
- `listing_price_change`
- `listing_price_drop_notified`
- `listing_price_history_view`
- `listing_price_change_blocked`
- `listing_price_change_reverted`

필수 속성 후보:
- `listingId`
- `listingType`
- `priceTypeBefore`
- `priceTypeAfter`
- `priceAmountBefore`
- `priceAmountAfter`
- `deltaAmount`
- `deltaPercentBand`
- `listingStatusAtChange`
- `changeReasonCode`

핵심 분석 포인트:
- 가격 인하 후 상세 재방문률 / 채팅 시작률 상승 여부
- 가격 변동성이 높은 매물의 완료 전환율과 신고율 상관관계
- `offer` 매물에서 공개 가격 전환 후 협상/완료 전환 개선 여부

### 68.12 오픈 질문
- 가격 이력은 비회원/일반 회원에게 어디까지 공개할 것인가, 아니면 작성자/찜 사용자에게만 제한할 것인가?
- `reserved` 상태에서의 가격 변경을 완전 금지할지, 우선 상대 동의 기반으로 허용할지?
- `buy` 매물의 희망가 변경이 찜/문의 사용자에게도 알림 가치가 있는지?
- 의미 있는 가격 변동 임계값을 절대값, 비율, 카테고리별 기준 중 무엇으로 둘지?
- `PATCH /listings/{listingId}`에 통합할지, 별도 가격 API로 분리할지?

## 69. 계정/기능 제한(Restriction) 정책 계약(초안)
### 69.1 목표
- 신고/안티어뷰즈/정책 위반/신규 계정 가드레일에서 발생하는 제한을 하나의 공통 도메인으로 정의해, 운영정책/백오피스/API/화면이 같은 기준을 사용하게 한다.
- `정지됨`, `일부 기능 제한`, `신뢰 제한`, `등록만 차단` 같은 상태를 단순 텍스트가 아닌 구조화된 restriction 객체로 다뤄, 사용자 경험과 운영 판단의 모호함을 줄인다.
- 계정 상태(`accountStatus`)와 개별 기능 제한(`restriction scope`)을 분리해 과잉 제재 없이 필요한 범위만 제한할 수 있게 한다.

### 69.2 기본 원칙
1. **최소 제한 우선**: 위험을 제어할 수 있다면 전체 정지보다 기능 제한을 우선한다.
2. **설명 가능성 유지**: 사용자에게는 이해 가능한 제한 사유와 해제 조건을 보여주고, 내부적으로는 상세 사유/증빙/룰 hit를 분리 보관한다.
3. **시간 제한 우선**: 영구 제한은 반복/중대 위반에 한정하고, 대부분의 제재는 기간과 재검토 시점을 가져야 한다.
4. **객체 단위와 계정 단위 분리**: 특정 매물/채팅만 잠그는 조치와 계정 전체 기능 제한을 구분한다.
5. **읽기와 쓰기 분리**: 제한 상태에서도 사용자가 내 거래 기록, 운영 통지, 이의제기 경로를 확인할 최소 읽기 권한은 유지한다.

### 69.3 restriction scope 분류
| scope | 의미 | 대표 사례 | 사용자 영향 |
|---|---|---|---|
| `warning_only` | 기능 차단 없는 경고 | 첫 경미 위반, 정책 안내 | 배너/알림만 노출 |
| `listing_only` | 매물 등록/수정/끌어올리기 제한 | 도배, 허위매물, 카탈로그 악용 | 채팅/기록 조회는 가능 |
| `chat_only` | 신규 채팅/메시지/예약 제한 | 욕설, 괴롭힘, 외부 연락처 반복 유도 | 매물 조회/관리 일부 가능 |
| `trade_execution_only` | 예약/완료/후기 등 거래 실행 단계 제한 | 노쇼 반복, 허위 완료 요청 | 탐색/문의는 가능할 수 있음 |
| `trust_limited` | 노출/랭킹/상대 신뢰 표시 제한 | 반복 reserved 악용, 가격 장난 | 검색 노출 감점, 배지 하향 |
| `read_only_account` | 대부분 쓰기 차단, 기록/이의제기만 허용 | 고위험 조사 중 임시 동결 | 읽기/증빙 제출만 가능 |
| `temporary_suspend` | 기간 정지 | 중대한 정책 위반 | 로그인 또는 쓰기 전면 제한 |
| `permanent_ban` | 영구 이용 제한 | 반복 악성 사기/중대 위반 | 복구는 이의제기 경로로만 |

원칙:
- 하나의 사용자에게 여러 scope가 동시에 적용될 수 있다.
- 클라이언트는 `accountStatus`만 보지 말고 활성 restriction 집합을 기준으로 버튼 노출/비활성/배너를 결정해야 한다.

### 69.4 restriction level과 계정 상태의 관계
| accountStatus | 의미 | 허용 restriction 예시 | 비고 |
|---|---|---|---|
| `active` | 기본 정상 상태 | `warning_only`, `listing_only`, `chat_only`, `trust_limited` | 부분 제한 가능 |
| `restricted` | 일부 기능 제한 중 | 대부분의 부분 제한 | 로그인 유지 가능 |
| `suspended` | 기간 정지 또는 조사 동결 | `read_only_account`, `temporary_suspend` | 쓰기 전면 제한 가능 |
| `banned` | 영구 이용 제한 | `permanent_ban` | 공개 프로필/활성 매물 종료 필요 |
| `withdrawn` | 자발 탈퇴 | restriction 비적용 또는 기록 보존용 | 제재 상태와 별도 축 |

설계 원칙:
- `accountStatus`는 계정의 전반 상태를 나타내고, 실제 기능 차단은 활성 restriction이 결정한다.
- 예: `active + trust_limited`, `restricted + chat_only`, `suspended + read_only_account` 같은 조합이 가능하다.

### 69.5 restriction 수명주기
| restrictionStatus | 의미 | 생성 주체 | 다음 상태 |
|---|---|---|---|
| `draft` | 운영 검토 중, 아직 미적용 | 운영자 | `active`, `dismissed` |
| `active` | 현재 적용 중 | 운영자/시스템 | `expired`, `lifted`, `escalated` |
| `expired` | 기간 만료로 종료 | 시스템 | 종료 |
| `lifted` | 운영 복구/조기 해제 | 운영자 | 종료 |
| `escalated` | 더 강한 restriction으로 승격 | 운영자/시스템 | 새 restriction 생성 |
| `dismissed` | 오탐/무근거로 미적용 | 운영자 | 종료 |

원칙:
- 제재 강도 변경은 기존 레코드 덮어쓰기보다 `escalated` + 신규 restriction 생성 방식이 감사 추적에 유리하다.
- 자동 만료와 운영 수동 해제를 모두 구분 저장해야 한다.

### 69.6 발동 트리거 유형
| triggerType | 설명 | 예시 |
|---|---|---|
| `report_based` | 신고 누적/운영 판단 기반 | 욕설 신고, 허위매물 신고 |
| `policy_violation` | 콘텐츠 정책 위반 | 전화번호/오픈채팅 링크 반복 입력 |
| `behavioral_risk` | 행동 패턴 기반 | reserved 악용, 가격 장난, 대량 문의 |
| `trust_bootstrap` | 신규 계정 보호성 제한 | 가입 직후 과다 채팅/매물 등록 |
| `manual_safety` | 운영자 수동 보호 조치 | 분쟁 조사 중 임시 동결 |
| `system_abuse` | 봇/자동화/스크래핑 의심 | 비정상 API 호출 패턴 |

원칙:
- 같은 restriction에도 `triggerType`과 `reasonCode`를 함께 남겨야 운영/분석/이의제기 처리 품질이 올라간다.

### 69.7 사용자 노출 원칙
사용자에게는 내부 탐지 룰 이름이나 점수를 직접 노출하지 않는다. 대신 아래 수준의 설명을 제공한다.

| scope | 사용자 노출 문구 예시 | 해제/행동 유도 |
|---|---|---|
| `warning_only` | `운영정책 위반 가능성이 있어 안내드려요.` | 정책 보기 |
| `listing_only` | `현재 매물 등록/수정 기능이 일시 제한되었어요.` | 만료 시각/이의제기 보기 |
| `chat_only` | `채팅 기능이 일시 제한되었어요.` | 남은 기간/고객지원 보기 |
| `trade_execution_only` | `예약/완료 같은 거래 실행 기능이 제한되었어요.` | 진행 중 거래 정리 안내 |
| `trust_limited` | `일부 노출과 신뢰 표시가 조정되었어요.` | 개선 조건은 내부 관리, 과도 노출 지양 |
| `read_only_account` | `계정 검토 중으로 읽기 전용 상태예요.` | 소명 제출/이의제기 |
| `temporary_suspend` | `계정 이용이 일시 정지되었어요.` | 종료 시각/이의제기 |
| `permanent_ban` | `운영정책 위반으로 이용이 제한되었어요.` | 이의제기 1회 경로 |

원칙:
- 사용자에게는 `무엇이 안 되는지`, `언제 풀리는지`, `어디서 이의제기하는지`를 우선 보여준다.
- 고위험 사기 탐지 등 민감 사유는 상세 근거 대신 일반화된 public reason code로 축약한다.

### 69.8 진행 중 거래 보호 규칙
- 제한이 걸리더라도 진행 중 `reservation_confirmed`, `trade_due_soon`, `completion_waiting_*`, `dispute_open` 거래는 별도 정리 정책이 필요하다.
- 기본안:
  1. `listing_only`: 진행 중 거래 유지, 신규 매물/수정만 차단
  2. `chat_only`: 신규 채팅/일반 메시지 차단 가능하나, 완료 확인/분쟁 소명/운영 요청 응답은 예외 허용 검토
  3. `trade_execution_only`: 새 예약/완료 요청 차단, 이미 열린 분쟁 소명은 허용
  4. `temporary_suspend` 이상: 진행 중 거래 상대에게 상태 안내 배너를 제공하고, 필요 시 운영자 개입으로 취소/잠금 처리
- 진행 중 거래를 갑자기 끊어 상대방이 피해를 보지 않도록 `grace action` 개념을 둘 수 있다.

### 69.9 API/응답 계약 시사점
읽기 API는 현재 사용자 기준 활성 restriction을 함께 반환하는 구조를 권장한다.

```json
{
  "viewerContext": {
    "accountStatus": "restricted"
  },
  "activeRestrictions": [
    {
      "restrictionId": "rst_123",
      "scope": "chat_only",
      "status": "active",
      "publicReasonCode": "ABUSIVE_CHAT_PATTERN",
      "startsAt": "2026-03-13T09:00:00+09:00",
      "endsAt": "2026-03-15T09:00:00+09:00",
      "allowedExceptionActions": ["submit_dispute_statement", "view_notifications"]
    }
  ],
  "availableActions": ["view_chat_history", "report_user"]
}
```

쓰기 API 공통 에러/힌트 후보:
- `ACCOUNT_RESTRICTED`
- `LISTING_ACTION_RESTRICTED`
- `CHAT_ACTION_RESTRICTED`
- `TRADE_EXECUTION_RESTRICTED`
- `READ_ONLY_ACCOUNT`

원칙:
- 같은 restriction이라도 화면별 허용 예외가 다를 수 있으므로 `allowedExceptionActions` 또는 최종 `availableActions`가 필요하다.
- Admin API는 restriction 생성/해제/승격 시 반드시 사유 코드, 증빙 링크, 검토자 메모를 요구해야 한다.

### 69.10 데이터 모델 후보
#### Restriction
| 컬럼 | 설명 |
|---|---|
| `restrictionId` | 제한 식별자 |
| `userId` | 대상 사용자 |
| `scope` | 제한 범위 |
| `restrictionStatus` | `draft` / `active` / `expired` / `lifted` / `escalated` / `dismissed` |
| `severityLevel` | 내부 강도 레벨 |
| `triggerType` | 발동 트리거 유형 |
| `reasonCode` | 내부 상세 사유 |
| `publicReasonCode` | 사용자 노출용 축약 사유 |
| `sourceReportId(optional)` | 연계 신고 |
| `sourceModerationActionId(optional)` | 연계 운영 조치 |
| `startsAt` | 시작 시각 |
| `endsAt(optional)` | 종료 시각 |
| `liftedAt(optional)` | 수동 해제 시각 |
| `liftedByAdminId(optional)` | 해제 운영자 |
| `appealEligibleUntil(optional)` | 이의제기 가능 시한 |
| `metadataJson(optional)` | 룰 hit, 대상 객체, 임계치 등 |
| `createdAt` / `createdBy` | 생성 감사 정보 |

#### RestrictionAffectedResource(optional)
- 특정 매물/채팅/거래에만 연결되는 제한을 위해 아래 구조를 검토한다.
- `restrictionId`
- `resourceType`: `listing` / `chat_room` / `trade_thread`
- `resourceId`
- `effectType`: `read_only` / `write_block` / `hidden_from_others`

### 69.11 백오피스 요구사항
운영 화면은 아래를 한 화면에서 볼 수 있어야 한다.
1. 현재 활성 restriction 목록
2. 생성 배경(신고, 자동탐지, 과거 제재, 대상 객체)
3. 남은 기간/자동 만료 예정 시각
4. 사용자 노출 문구 preview
5. 진행 중 거래 영향 요약
6. 해제/승격/연장 액션
7. 이의제기 접수 여부와 처리 상태

운영 액션 가드레일:
- `temporary_suspend`, `permanent_ban`은 2차 확인 또는 상위 권한 승인 권장
- 자동 탐지 기반 restriction은 수동 확정 전 `draft` 또는 낮은 범위 제한으로 시작하는 것이 안전
- 해제 시에도 해제 사유와 재발 방지 메모를 남겨야 한다

### 69.12 분석 이벤트 및 KPI 후보
- `restriction_created`
- `restriction_activated`
- `restriction_expired`
- `restriction_lifted`
- `restriction_escalated`
- `restriction_appeal_submitted`
- `restricted_action_blocked`

핵심 분석 포인트:
- scope별 재위반률
- `warning_only` 이후 정상 전환률 vs escalated 비율
- 신규 계정 보호성 restriction이 정상 사용자 전환을 얼마나 해치거나 돕는지
- restriction 적용 사용자군의 신고율, 완료율, 노쇼율 변화
- 오탐률(해제율)과 사용자 이탈률

### 69.13 오픈 질문
- `trust_limited`를 사용자에게 얼마나 명시적으로 보여줄지, 아니면 내부 랭킹/배지 조정으로만 처리할지?
- `chat_only` 제한 중 완료 확인/분쟁 소명 메시지를 예외 허용할지 별도 전용 액션으로 분리할지?
- 제한 해제 후 신뢰 배지/랭킹 복구를 즉시 할지, 일정 관찰 기간을 둘지?
- 신규 사용자 보호용 restriction과 위반 제재용 restriction을 같은 테이블/도메인으로 완전히 통합할지?

## 70. 도메인 이벤트 / 아웃박스 / 발행 계약(초안)
### 70.1 목표
- 쓰기 트랜잭션 결과가 read model, 알림, analytics, 운영 타임라인에 서로 다른 해석으로 복제되지 않도록 **공통 사건 단위**를 정의한다.
- API 성공 응답과 비동기 후속 처리(알림 fanout, projection 갱신, 운영 큐 적재)가 같은 원천 사실을 바라보게 한다.
- 향후 SSE/WebSocket, 배치 재처리, 감사 로그 재구성, 데이터 복구 작업이 모두 동일 이벤트 계약 위에서 동작할 수 있게 한다.

### 70.2 기본 원칙
1. 사용자/운영자/시스템이 상태를 바꾸는 모든 핵심 쓰기 액션은 **도메인 이벤트**를 남긴다.
2. DB 상태 변경과 이벤트 발행은 분리 호출이 아니라 **outbox 기반 원자적 기록**을 기본 원칙으로 한다.
3. 이벤트는 `무슨 사실이 확정되었는가`를 표현하고, UI 문구나 푸시 카피를 직접 담지 않는다.
4. read model/analytics/notification은 가능하면 동일 이벤트를 소비하되, 각 소비자는 독립적으로 멱등 처리해야 한다.
5. 이벤트는 append-only이며, 정정이 필요하면 기존 이벤트 수정이 아니라 보정 이벤트를 추가한다.

### 70.3 이벤트 분류 계층
| 계층 | 의미 | 예시 | 사용처 |
|---|---|---|---|
| Domain Event | 비즈니스 사실 | `listing_published`, `reservation_confirmed` | projection, notification, audit |
| Integration Event | 외부/타 서비스 전달용 정제 이벤트 | `notification_requested`, `search_reindex_requested` | fanout, 검색색인, 외부 연동 |
| Analytics Event | 제품 분석용 이벤트 | `trade_thread_primary_action_clicked` | BI/대시보드 |

원칙:
- PRD 기준 원천은 **Domain Event**다.
- Integration/Analytics는 Domain Event에서 파생되며, 원천 상태를 역으로 바꾸지 않는다.

### 70.4 공통 이벤트 envelope 후보
```json
{
  "eventId": "evt_01H...",
  "eventType": "reservation_confirmed",
  "aggregateType": "reservation",
  "aggregateId": "res_123",
  "aggregateVersion": 4,
  "occurredAt": "2026-03-13T13:01:00+09:00",
  "actor": {
    "actorType": "user",
    "actorId": "user_123"
  },
  "causationId": "req_abc",
  "correlationId": "tradeThread_tt_123",
  "payload": {
    "listingId": "listing_123",
    "chatRoomId": "chat_123",
    "scheduledAt": "2026-03-13T21:00:00+09:00"
  }
}
```

필드 원칙:
- `eventId`: 전역 유니크, 소비자 멱등 키
- `aggregateType` / `aggregateId`: 사건 소유 aggregate 식별
- `aggregateVersion`: 같은 aggregate 내 순서 보장/중복 방지 기준
- `causationId`: API requestId, jobId, moderationActionId 등 직접 원인
- `correlationId`: 하나의 거래 흐름/사건 묶음 추적용. 기본 후보는 `tradeThreadId`, `reportId`, `disputeId`
- `payload`는 소비자에게 필요한 최소 사실만 담고, 대용량 본문/이미지 원문은 참조 ID를 우선 사용

### 70.5 Aggregate별 필수 이벤트 후보
#### Listing
- `listing_created`
- `listing_published`
- `listing_updated`
- `listing_status_changed`
- `listing_price_changed`
- `listing_expired`
- `listing_bumped`
- `listing_hidden`
- `listing_restored`

#### Chat / Message
- `chat_room_created`
- `chat_message_sent`
- `chat_message_blocked_by_policy`
- `chat_locked`
- `chat_unlocked`
- `chat_read_cursor_advanced`

#### Reservation / Meeting
- `reservation_proposed`
- `reservation_updated`
- `reservation_confirmed`
- `reservation_cancelled`
- `reservation_expired`
- `reservation_meeting_changed`
- `reservation_location_acknowledged`

#### Completion / Dispute
- `trade_completion_requested`
- `trade_completion_confirmed`
- `trade_completion_auto_confirmed`
- `trade_completion_disputed`
- `dispute_opened`
- `dispute_statement_submitted`
- `dispute_resolved`

#### Review / Report / Restriction
- `review_submitted`
- `review_hidden`
- `report_submitted`
- `report_triaged`
- `moderation_action_executed`
- `restriction_activated`
- `restriction_lifted`

### 70.6 이벤트와 상태머신의 관계
- 상태 변경은 가능하면 `status_changed` 단일 이벤트만으로 끝내지 않고, 도메인 의미가 큰 경우 구체 이벤트를 우선 사용한다.
  - 예: `reservation_confirmed`는 `reservation_status_changed(to=confirmed)`보다 읽기 쉽고 downstream 설계가 단순하다.
- 단, 이력 조회와 범용 운영 화면을 위해 `fromStatus`, `toStatus`, `reasonCode`는 payload 또는 공통 metadata로 함께 남기는 것이 좋다.
- 하나의 command가 여러 aggregate를 바꾸면 **주 소유 aggregate 이벤트 + 파생 이벤트**를 구분한다.
  - 예: 예약 확정 시
    1. `reservation_confirmed` (주 이벤트)
    2. `listing_status_changed` (`available -> reserved`)
    3. `trade_thread_projection_refresh_requested`(integration 성격)

### 70.7 발행 순서 / 트랜잭션 계약
#### 단일 aggregate 쓰기
- 같은 aggregate 안에서는 `aggregateVersion`이 반드시 1씩 증가해야 한다.
- 소비자는 `aggregateId + aggregateVersion` 역전 현상을 감지할 수 있어야 한다.

#### 다중 aggregate 쓰기
- 하나의 DB 트랜잭션에서 여러 aggregate 상태가 함께 바뀌더라도, outbox에는 논리 순서를 보장해 저장한다.
- 기본 우선순위 권장:
  1. command 소유 aggregate 이벤트
  2. 직접 영향을 받은 상태 동기화 이벤트
  3. notification/search/rebuild 요청 등 integration 이벤트

예시: 완료 요청
1. `trade_completion_requested`
2. `listing_status_changed` (`pending_trade -> completed`)
3. `trade_thread_status_changed` 또는 projection refresh event
4. `notification_requested`(상대 확인 필요)

### 70.8 Outbox 저장 구조 후보
| 컬럼 | 설명 |
|---|---|
| `outboxId` | 내부 순차/ULID 식별자 |
| `eventId` | 외부 소비용 유니크 ID |
| `eventType` | 이벤트 타입 |
| `aggregateType` / `aggregateId` | 소유 경계 |
| `aggregateVersion` | 순서 제어용 버전 |
| `payloadJson` | 직렬화 payload |
| `occurredAt` | 비즈니스 사건 시각 |
| `persistedAt` | outbox 저장 시각 |
| `publishStatus` | `pending` / `publishing` / `published` / `dead_letter` |
| `publishAttempts` | 재시도 횟수 |
| `lastErrorCode(optional)` | 실패 사유 |
| `publishedAt(optional)` | 최종 발행 시각 |

원칙:
- outbox는 business table과 같은 트랜잭션에서 insert되어야 한다.
- 발행 실패는 API 실패로 되돌리지 않고, 재시도/운영 모니터링 대상으로 남긴다.
- `dead_letter` 전환 기준과 재처리 runbook이 필요하다.

### 70.9 소비자(Subscriber) 책임 분리
| 소비자 | 주 역할 | 멱등 키 후보 | 실패 영향 |
|---|---|---|---|
| Trade/Listing Projection | 목록/상세/내거래 projection 갱신 | `eventId` 또는 `aggregateVersion` | 화면 최신성 지연 |
| Notification Fanout | 인앱/푸시 알림 생성 | `eventId + recipientUserId` | 알림 누락/중복 |
| Audit Timeline Builder | 운영 타임라인 구성 | `eventId` | 감사 화면 누락 |
| Search/Ranking Updater | 검색색인/랭킹 갱신 | `eventId + indexTarget` | 검색 노출 지연 |
| Analytics Exporter | 제품 분석 적재 | `eventId` | 지표 누락 |

원칙:
- 한 소비자 실패가 다른 소비자 성공을 롤백하지 않는다.
- 사용자에게 직접 보이는 projection/notification은 우선순위가 높고, analytics는 늦어져도 복구 가능해야 한다.

### 70.10 알림/실시간 전달과의 연결 규칙
- 푸시/SSE/WebSocket 전송은 `chat_message_sent`, `reservation_confirmed`, `trade_completion_requested` 같은 Domain Event를 직접 또는 1단계 변환 후 소비한다.
- 실시간 전송 실패가 이벤트 자체를 소실시키면 안 되며, 항상 인앱 알림함/동기화 API로 복구 가능해야 한다.
- 메시지 전송처럼 체감 지연이 중요한 경우에도 **source of truth는 저장된 이벤트/메시지 레코드**여야 하며, 소켓 전송 성공 여부가 저장 성공을 대체하지 않는다.

### 70.11 운영/복구 시사점
- 특정 사용자가 `예약 확정 알림을 못 받았다`고 주장할 때, 운영자는 아래를 순서대로 추적할 수 있어야 한다.
  1. `reservation_confirmed` 도메인 이벤트 존재 여부
  2. outbox 발행 성공 여부
  3. notification fanout 생성 여부
  4. push suppressed/delivered/opened 여부
  5. projection 반영 여부
- projection 재빌드 시 원천 테이블 스냅샷 대신 이벤트 재생이 가능하면 이상적이지만, MVP에서는 현재 상태 테이블 + 상태 이력 + 핵심 outbox 보존 조합으로도 시작 가능하다.
- dead-letter 재처리는 aggregate 순서를 깨지 않도록 대상 범위를 좁혀 수행해야 한다.

### 70.12 API/DB/문서 파생 포인트
- DB 스키마 문서:
  - outbox table, publish status index, retention 정책
  - aggregate version 증가 규칙
- API 명세:
  - 동기 응답과 비동기 후속 반영 차이 설명
  - `requestId`, `correlationId` 노출 여부
- 운영정책:
  - 이벤트 유실/중복/역순 처리 대응 runbook
  - dead-letter 모니터링 임계치
- 화면명세:
  - optimistic UI 허용 범위
  - `동기 저장 완료`와 `후속 반영 중` 상태 분리 기준

### 70.13 오픈 질문
- MVP에서 projection 재생의 원천을 상태 테이블 + 이력 테이블로 둘지, outbox/event log 장기보존까지 포함할지?
- `chat_read_cursor_advanced` 같은 고빈도 이벤트를 full event log로 남길지, 요약 snapshot만 둘지?
- search index 갱신을 domain event 직접 구독으로 할지, 별도 integration event로 분리할지?
- 외부 webhook/파트너 연동이 생길 경우 어떤 이벤트만 공개 surface로 허용할지?

## 71. 검색 쿼리 / 필터 패싯 / Facet Count 계약(초안)
### 71.1 목표
- 거래소 목록, 홈 검색 진입, 딥링크, SEO 랜딩, 저장된 필터가 모두 같은 검색 해석을 사용해야 한다.
- 필터 UI에서 보여주는 패싯 개수와 실제 결과 집합이 어긋나지 않도록 `query snapshot + facet calculation` 계약을 명시한다.
- 검색은 단순 텍스트 매칭이 아니라 `실제 거래 가능한 매물`을 빨리 좁히는 도구여야 한다.

### 71.2 검색 요청의 표준 구성
`GET /listings` 또는 동등한 search endpoint는 아래 세 층을 분리해 해석한다.

| 층 | 예시 파라미터 | 역할 |
|---|---|---|
| query | `q`, `serverId`, `categoryId` | 결과 집합을 강하게 결정하는 핵심 조건 |
| filter | `listingType`, `status[]`, `priceMin`, `priceMax`, `tradeMethod[]`, `attributeFilters[]` | 사용자가 결과를 추가로 좁히는 조건 |
| sort/view | `sort`, `cursor`, `pageSize`, `includeFacets` | 결과 표시 방식과 부가정보 제어 |

원칙:
- `query`는 검색 의도, `filter`는 결과 정제, `sort`는 노출 순서를 담당한다.
- 같은 URL/딥링크를 재열 때 동일 결과를 복구할 수 있도록 모든 사용자 선택 상태는 직렬화 가능해야 한다.
- `includeFacets=false`일 때는 facet 계산을 생략해 비용을 줄일 수 있어야 한다.

### 71.3 기본 검색 컨텍스트
초기 진입 시 클라이언트가 암묵적으로 적용하는 기본값을 명문화한다.

| 항목 | 기본값 | 비고 |
|---|---|---|
| `status[]` | `available` | 사용자가 명시적으로 바꾸지 않으면 `reserved`, `pending_trade`, `completed`, `cancelled` 제외 |
| `visibility` | `public` | 숨김/차단 매물 제외 |
| `listingType` | 전체 또는 마지막 선택 복원 | 홈 탭/딥링크에 따라 달라질 수 있음 |
| `serverId` | 사용자의 최근 선택 또는 비선택 | 서버 미선택 시 전체 검색 허용 여부는 surface별 결정 |
| `sort` | `recommended` | 추천순 기본 |
| `includeFacets` | true | 목록 화면 기본 |

세부 원칙:
- `reserved 포함`은 사용자가 직접 토글할 때만 결과 집합에 포함하는 기본안을 유지한다.
- `pending_trade`, `completed`, `cancelled`는 공개 탐색용 facet에서 기본 제외한다.
- 작성자 본인이나 거래참여자의 기록 화면은 공개 탐색 검색과 다른 기본 컨텍스트를 사용해도 되지만, endpoint level에서는 `searchMode=public|owner|participant`처럼 의도를 분리하는 편이 안전하다.

### 71.4 우선 지원 패싯 정의
#### 71.4.1 단일 선택 패싯
- `serverId`
- `categoryId` (MVP는 단일 선택 우선)
- `listingType` (`sell` / `buy`)
- `priceType` (`fixed` / `negotiable` / `offer`)

#### 71.4.2 다중 선택 패싯
- `status[]` (`available`, `reserved` 중심)
- `tradeMethod[]` (`in_game`, `offline_pc_bang`, `either`)
- `attributeFilters[]` (카테고리별 구조화 속성)

#### 71.4.3 범위 패싯
- `priceMin`, `priceMax`
- `updatedWithinDays` 또는 `activityWithin`
- 추후 `enhancementLevelMin/Max` 같은 숫자형 속성도 attribute facet으로 흡수 가능

원칙:
- MVP에서는 패싯 종류를 과도하게 늘리기보다 **서버 / 카테고리 / 거래유형 / 가격 / 거래방식 / 상태**에 집중한다.
- 카테고리별 속성 패싯은 74장 속성 템플릿과 동일 키를 사용해야 한다.
- 자유입력 텍스트 옵션은 facet이 아니라 전문검색 필드로 다루고, 패싯 승격 여부는 사용 빈도 기반으로 판단한다.

### 71.5 facet count 계산 원칙
Facet count는 “현재 결과 집합에서 특정 필터를 추가/전환했을 때 예상되는 결과 수”로 정의한다.

| 방식 | 설명 | MVP 권장 여부 |
|---|---|---|
| fully constrained | 현재 적용 필터 전체를 기준으로 count 계산 | 권장 |
| self-excluded | 자기 facet 값만 제외하고 계산 | 일부 facet에 권장 |
| global total | 전체 public inventory 기준 count | 비권장 |

권장 기본안:
- 서버/카테고리/거래방식/속성 facet은 **현재 query + 다른 active filter를 유지한 상태에서 self-excluded count**를 사용한다.
- 상태 facet은 `available` 기본 필터 때문에 혼동이 크므로, 현재 visible result 기준 count와 toggle 후 예상 count를 일관되게 보여야 한다.
- 패싯 count는 사용자가 현재 볼 수 없는 hidden/blocked inventory를 포함하면 안 된다.

예시:
- 현재 `serverId=데포로쥬`, `listingType=sell`, `status=available`이면
  - `tradeMethod=in_game` count는 같은 조건에서 거래방식만 바꾼 예상치
  - `status=reserved` count는 `available + reserved` 조합 적용 시 몇 건인지 명확히 정의 필요

### 71.6 상태 패싯 특례
상태 패싯은 일반 패싯과 달리 공개 정책과 강하게 연결되므로 별도 규칙을 둔다.

| 상태 | 공개 탐색 facet 노출 | count 포함 기준 | 비고 |
|---|---|---|---|
| `available` | 기본 ON | public + available | 기본 결과 집합 |
| `reserved` | 선택 노출 | public + reserved | 기본 OFF, 사용자가 포함 가능 |
| `pending_trade` | 기본 미노출 | 포함하지 않음 | 참여자/작성자 화면 전용 |
| `completed` | 미노출 | 포함하지 않음 | 공개 탐색 제외 |
| `cancelled` | 미노출 | 포함하지 않음 | 공개 탐색 제외 |
| `expired/stale` 내부 상태 | 직접 노출 안 함 | visibility/ranking 정책으로 흡수 | 사용자 개념 단순화 |

원칙:
- 상태 패싯은 internal enum을 그대로 노출하지 않고, 사용자 의미 단위(`거래 가능`, `예약중`)로 단순화한다.
- `reserved` count는 기본 목록에 없는 재고를 보여주는 성격이므로, UI 카피도 “예약중 포함”처럼 명시적이어야 한다.

### 71.7 속성 패싯(attribute facet) 계약
- 속성 패싯은 카테고리별 속성 템플릿에서 `filterable=true`로 선언된 속성만 생성한다.
- 속성 타입별 지원 방식:

| 속성 타입 | facet 형태 | 예시 |
|---|---|---|
| enum | 다중 선택 | 희귀도, 속성 타입 |
| boolean | 토글 | 축복 여부 |
| numeric | 범위 또는 버킷 | 강화 수치 |
| text-normalized | suggestion 선택 | 아이템 별칭 |

원칙:
- facet key는 사람이 읽는 라벨이 아니라 안정적인 `attributeKey`를 사용한다.
- facet 응답에는 `displayLabel`, `value`, `count`, `selected`를 포함한다.
- 동일 카테고리라도 목록 카드에 없는 속성은 facet으로 드러날 수 있으므로, 상세/필터 칩/검색결과 empty 상태에 동일 라벨 규칙이 필요하다.

### 71.8 API 응답 계약 후보
```json
{
  "querySnapshot": {
    "q": "장검",
    "serverId": "deporoju",
    "listingType": "sell",
    "status": ["available"],
    "sort": "recommended"
  },
  "items": [
    {
      "listingId": "lst_123",
      "summaryText": "+7 장검 · 데포로쥬 · 50만",
      "displayStatusLabel": "거래 가능"
    }
  ],
  "facets": {
    "server": [
      { "value": "deporoju", "label": "데포로쥬", "count": 128, "selected": true }
    ],
    "status": [
      { "value": "available", "label": "거래 가능", "count": 128, "selected": true },
      { "value": "reserved", "label": "예약중 포함", "count": 17, "selected": false }
    ],
    "tradeMethod": [
      { "value": "in_game", "label": "인게임", "count": 97, "selected": false }
    ]
  },
  "sortOptions": [
    { "value": "recommended", "selected": true },
    { "value": "latest", "selected": false },
    { "value": "price_low", "selected": false }
  ],
  "cursor": {
    "next": "c_abc",
    "snapshotToken": "snp_123"
  }
}
```

원칙:
- facet count와 item 목록은 같은 `querySnapshot` 기준으로 계산되어야 한다.
- `snapshotToken`이 있으면 페이지네이션/새로고침 중 같은 결과 스냅샷을 유지할 수 있다.
- `facets`는 클라이언트가 조합 규칙을 추론하지 않아도 되도록 서버가 완성형으로 제공하는 것을 권장한다.

### 71.9 딥링크 / URL 직렬화 규칙
- 목록 URL/딥링크는 아래 수준까지 안정적으로 복원 가능해야 한다.
  - 검색어
  - 서버/카테고리/거래유형
  - 상태 포함 여부
  - 가격 범위
  - 거래방식
  - 정렬
- 예시:
  - `/listings?server=deporoju&type=sell&q=%EC%9E%A5%EA%B2%80&status=available,reserved&tradeMethod=in_game`
- URL에는 내부 facet key를 쓰되, 사람이 읽기 좋은 slug와 안정 키 중 하나를 일관되게 선택해야 한다.
- 잘못된 facet 조합은 400으로 실패시키기보다, 가능한 범위 내 정규화하고 `appliedCorrections`를 응답에 포함하는 방안을 검토한다.

### 71.10 zero-result / 복구 UX 규칙
검색 결과 0건은 실패가 아니라 다음 행동 유도 surface다.

| 상황 | 우선 복구 제안 |
|---|---|
| 서버 + 아이템 조합 0건 | 서버 전체 보기 또는 유사 서버 제안 |
| 가격 범위 과도 | 가격 필터 해제 |
| `available`만 보고 0건 | `reserved 포함` 토글 제안 |
| 속성 필터 과도 | 마지막 선택 속성 facet 제거 제안 |
| 오타/별칭 가능성 | 교정 검색어 제안 |

원칙:
- zero-result 상태에서도 facet count 또는 추천 완화 액션은 남겨야 한다.
- 단순 “검색 결과가 없어요”보다, 어떤 필터를 빼면 결과가 생기는지 제안해야 한다.
- 검색 교정 제안은 자동 치환보다 opt-in CTA를 우선한다.

### 71.11 인덱스 / 캐시 / 성능 시사점
- facet count는 목록 item 조회와 별도 집계 경로가 필요할 수 있으므로, read model/materialized view 또는 search index 집계를 우선 검토한다.
- 카운트가 다소 지연될 수는 있어도, item 목록과 상충하는 stale count는 최소화해야 한다.
- 권장 우선순위:
  1. item result 정확성
  2. cursor 안정성
  3. facet count 최신성
- `includeFacets=false` 요청, debounce 검색, 인기 질의 캐시가 필요하다.
- 홈/SEO 랜딩은 full facet보다 축약 facet만 제공해도 된다.

### 71.12 분석 / QA / 운영 포인트
분석 이벤트 후보:
- `search_result_view`
- `search_facet_impression`
- `search_facet_select`
- `search_facet_clear`
- `search_zero_result_recovery_click`
- `search_snapshot_refresh`

핵심 지표:
- facet 사용률
- facet 적용 후 상세 진입률/채팅 시작률
- `reserved 포함` 토글 사용률과 거래 전환율
- zero-result 발생률과 복구 액션 성공률

QA 체크포인트:
- 동일 `querySnapshot`에서 item count와 facet count가 논리적으로 일관되는가
- `available` 기본 필터가 딥링크 재진입 시 유지되는가
- `reserved 포함` 토글 on/off 시 목록/카피/배지가 함께 바뀌는가
- category 변경 시 attribute facet가 즉시 재계산되고 invalid selection이 정리되는가
- hidden/blocked inventory가 facet count에 섞이지 않는가

### 71.13 오픈 질문
- 검색 구현을 DB 기반 materialized view로 시작할지, 별도 search engine 도입을 초기부터 고려할지?
- `reserved` count를 `available`와 합산 토글형으로 보여줄지, 별도 상태 탭으로 분리할지?
- 다중 서버 선택을 MVP에 포함할지, 단일 서버 선택 + 빠른 전환으로 제한할지?
- category facet를 트리형(대/중분류)으로 시작할지, 단일 depth로 단순화할지?



## 67. 저장 검색(Saved Search) / 재알림 구독 정책(초안)
### 67.1 목표
- 사용자가 매번 같은 조건을 다시 입력하지 않아도 원하는 거래 기회를 놓치지 않게 한다.
- 단순 검색 편의 기능이 아니라, `검색 -> 관심 -> 신규 매물/가격 변화 감지 -> 즉시 채팅 진입`으로 이어지는 거래 전환 장치로 설계한다.
- 검색 품질, 푸시 소음, 랭킹/캐시 비용이 함께 관리되도록 저장 검색의 범위와 발송 규칙을 명시한다.

### 67.2 기본 개념
- **Saved Search**: 사용자가 저장한 검색 조건 스냅샷
- **Search Alert Subscription**: 저장 검색에 연결된 알림 규칙
- **Baseline Result Set**: 저장 시점 또는 마지막 알림 시점 기준으로 이미 본 결과 집합
- **Match Event**: 저장 검색 조건을 새로 만족하게 된 사건(신규 매물 등록, 가격 하락, `reserved -> available` 복귀 등)

원칙:
- 저장 검색은 “현재 URL 상태 저장”이 아니라, 서버가 해석 가능한 구조화된 조건 집합으로 저장한다.
- 알림은 단순 조회 결과 전체가 아니라 **새로운 가치가 생긴 변화**에 대해서만 발송한다.

### 67.3 저장 가능한 검색 조건 범위
MVP 기본안에서 저장 가능한 조건은 아래로 제한한다.

| 조건 | 저장 여부 | 비고 |
|---|---|---|
| `listingType` | 필수 | `sell` / `buy` |
| `serverId` | 권장 | 단일 서버 우선 |
| `categoryId` | 선택 | 카테고리 단위 |
| `itemQuery` | 선택 | 사용자가 입력한 검색어 원문 |
| `normalizedItemId` | 선택 | 검색 정규화 성공 시 함께 저장 |
| `priceMin`, `priceMax` | 선택 | 고정가/협의가 혼합 해석 필요 |
| `tradeMethod` | 선택 | `in_game` / `offline_pc_bang` / `either` |
| `statusFilter` | 제한적 | 기본은 `available`; `reserved 포함`은 opt-in |
| `attributeFilters` | 선택 | 카테고리별 구조화 속성 필터 |
| `sort` | 저장 가능하나 알림 판단에는 비핵심 | 기본은 추천순 |

제한 원칙:
- 페이지네이션 커서, 일시적 정렬 상태, 실험 플래그, 디버그 파라미터는 저장 검색 조건에 포함하지 않는다.
- 너무 넓은 질의(예: 서버/카테고리/검색어 없이 전체 매물)는 저장은 허용할 수 있어도 푸시 알림 대상에서는 제한할 수 있다.

### 67.4 저장 검색 생성/수정 UX 원칙
- 검색 결과 화면에서 현재 적용된 필터가 유효할 때 `이 검색 저장` CTA를 제공한다.
- 저장 시 사용자에게 아래를 최소 설정으로 보여준다.
  1. 저장 이름(자동 제안 가능)
  2. 알림 받기 여부
  3. 알림 유형(신규 매물 / 가격 하락 / 상태 복귀)
- 저장 이름 자동 제안 예시:
  - `데포로쥬 +7 장검 판매`
  - `기란 재료 구매글`
- 사용자는 저장 후에도 `필터 수정`, `알림 일시중지`, `삭제`를 쉽게 할 수 있어야 한다.

### 67.5 저장 검색 결과 기준선(Baseline) 원칙
알림 스팸을 막기 위해, 저장 검색에는 “무엇이 이미 알려진 결과인지”에 대한 기준선이 필요하다.

기본안:
- 저장 검색 생성 시점의 상위 결과 집합을 baseline으로 기록한다.
- 이후 아래 경우에만 새 match로 본다.
  1. 새 `listingId`가 조건을 만족하며 baseline에 없었다.
  2. 기존 매물이 `reserved -> available`로 복귀해 다시 거래 가능해졌다.
  3. 기존 매물의 `priceAmount`가 사용자가 관심 가질 정도로 하락했다.
  4. 운영 숨김/차단 해제 후 다시 공개되었고, 사용자에게 실질적 신규 기회가 생겼다.
- 단순 `updatedAt` 변경, 사소한 설명 수정, 이미지 추가는 기본적으로 알림 트리거로 보지 않는다.

설계 원칙:
- baseline은 전체 결과 스냅샷을 영구 저장하기보다, 상위 N개 result fingerprint 또는 마지막 알림 시점 이후 알려진 listing 집합을 TTL과 함께 관리하는 방식이 현실적이다.
- baseline 갱신 시점은 `알림 발송 성공 시점`을 우선 기준으로 둔다.

### 67.6 Match Event 유형 정의
| eventType | 의미 | 알림 기본값 | 비고 |
|---|---|---|---|
| `new_listing_match` | 새 매물이 조건 일치 | ON | 핵심 이벤트 |
| `listing_back_available` | `reserved/pending`에서 `available` 복귀 | ON | 거래 재기회 |
| `price_drop_match` | 관심 조건 내 가격 하락 | ON(선택형 가능) | 변동폭 기준 필요 |
| `high_quality_match` | 이미지/응답성/신뢰 조건이 좋은 신규 매물 | MVP 제외 | 랭킹 심화 이후 검토 |
| `bulk_digest_match` | 다수 신규 결과 요약 | ON | 너무 넓은 검색의 폴백 |

세부 원칙:
- `price_drop_match`는 절대값 또는 비율 기준(예: 5% 이상 또는 일정 금액 이상, 가정)이 있어야 한다.
- 같은 매물에 대해 `new_listing_match` 이후 곧바로 `price_drop_match`가 발생해도 짧은 기간 중복 발송은 억제해야 한다.

### 67.7 알림 발송/억제 정책
| 상황 | 기본 동작 | 억제 규칙 |
|---|---|---|
| 조건 일치 신규 매물 1~3개 | 개별 또는 소규모 묶음 알림 | 같은 저장 검색에 최근 발송 후 짧은 시간 내 중복 억제 |
| 조건 일치 신규 매물 다수 | digest형 묶음 알림 | 너무 넓은 질의는 즉시 다건 푸시 금지 |
| 가격 하락 | 같은 매물의 첫 의미 있는 하락만 알림 | 소폭 반복 하락은 묶음/무시 |
| 상태 복귀 | `available` 재진입 시 알림 | 직전 이미 사용자가 상세 진입/찜 중이면 중복 푸시 억제 가능 |
| 심야 시간 | 인앱 기록 우선, 푸시는 중요도 따라 지연 | 거래 필수 알림보다 우선순위 낮음 |

기본 권장안:
- 저장 검색 알림은 **거래 필수 알림보다 한 단계 낮은 중요도**로 본다.
- 푸시를 허용하더라도 하루 최대 발송 수, 검색당 최소 발송 간격(cooldown), 같은 매물 중복 알림 제한이 필요하다.
- 넓은 검색어는 실시간 푸시보다 하루 1~2회 digest를 우선하는 것이 안전하다.

### 67.8 저장 검색 상태/수명주기
| 상태 | 의미 | 사용자 액션 | 시스템 액션 |
|---|---|---|---|
| `active` | 저장 및 알림 활성 | 수정/일시중지/삭제 | match 평가 대상 |
| `paused` | 저장 유지, 알림만 중지 | 재개/삭제 | 평가 생략 가능 |
| `expired` | 장기 미사용/정책상 만료 | 재활성화/삭제 | 자동 중지 |
| `deleted` | 사용자 삭제 | 복구 불가 또는 짧은 복구 유예 | 평가 중단 |

기본안:
- 사용자가 60~90일 이상 저장 검색을 열람/활용하지 않으면 `expired` 후보로 전환 가능하다(가정).
- 만료 전 1회 인앱 안내를 주고, 사용자가 다시 검색을 열면 즉시 `active` 복귀 가능해야 한다.

### 67.9 Saved Search 데이터 모델 후보
#### SavedSearch
| 필드 | 설명 |
|---|---|
| `savedSearchId` | 저장 검색 식별자 |
| `userId` | 소유자 |
| `name` | 사용자 표시 이름 |
| `queryJson` | 구조화된 검색 조건 스냅샷 |
| `queryHash` | 중복 탐지용 정규화 해시 |
| `status` | `active` / `paused` / `expired` / `deleted` |
| `alertEnabled` | 알림 활성 여부 |
| `lastViewedAt` | 사용자가 마지막으로 이 저장 검색을 열람한 시각 |
| `lastEvaluatedAt` | 마지막 매칭 평가 시각 |
| `lastNotifiedAt` | 마지막 알림 발송 시각 |
| `createdAt` | 생성 시각 |
| `updatedAt` | 수정 시각 |

#### SavedSearchNotification
| 필드 | 설명 |
|---|---|
| `savedSearchNotificationId` | 알림 기록 식별자 |
| `savedSearchId` | 연결 저장 검색 |
| `eventType` | `new_listing_match` 등 |
| `matchedListingId` | 단건 매칭 시 대상 매물 |
| `matchedCount` | digest형일 때 결과 수 |
| `baselineVersion` | 어떤 기준선 대비 신규인지 추적 |
| `deliveryStatus` | `queued` / `sent` / `suppressed` / `failed` |
| `suppressedReasonCode` | 억제 사유 |
| `createdAt` | 생성 시각 |
| `sentAt` | 발송 시각 |

권장 원칙:
- `queryJson`은 화면 URL 파라미터 원문이 아니라 정규화된 API 필터 구조를 저장한다.
- `queryHash`로 동일 사용자의 중복 저장 검색을 감지하고, 완전 중복이면 병합/이름만 수정하도록 유도할 수 있다.

### 67.10 API 후보
#### 사용자용
- `GET /me/saved-searches`
- `POST /me/saved-searches`
- `PATCH /me/saved-searches/{savedSearchId}`
- `POST /me/saved-searches/{savedSearchId}/pause`
- `POST /me/saved-searches/{savedSearchId}/resume`
- `DELETE /me/saved-searches/{savedSearchId}`
- `GET /me/saved-searches/{savedSearchId}/matches` (확장 후보)

#### 내부/시스템
- 저장 검색 매칭 평가 job
- digest 생성 job
- baseline 재구성/정리 job

응답 예시 후보:
```json
{
  "savedSearchId": "ss_123",
  "name": "데포로쥬 +7 장검",
  "status": "active",
  "alertEnabled": true,
  "query": {
    "listingType": "sell",
    "serverId": "depo",
    "itemQuery": "+7 장검",
    "statusFilter": ["available"]
  },
  "alertRules": {
    "newListing": true,
    "priceDrop": true,
    "backAvailable": true,
    "digestMode": "auto"
  },
  "lastNotifiedAt": "2026-03-13T15:55:00+09:00"
}
```

### 67.11 운영/성능 가드레일
- 사용자당 저장 검색 개수 상한이 필요하다(예: 20~50개, 가정).
- 너무 넓은 질의, 너무 잦은 결과 변동이 있는 질의는 실시간 매칭이 아니라 배치 digest로 강등할 수 있어야 한다.
- 검색 인프라 비용을 고려해, 저장 검색 평가는 `listing changed event` 기반 증분 평가를 우선하고 전체 재검색 반복은 최소화한다.
- 차단/숨김/제재된 매물은 저장 검색 매칭에서 즉시 제외되어야 하며, 발송 직전 재검증이 필요하다.
- 운영자는 아래를 볼 수 있어야 한다.
  - 저장 검색 수/활성률
  - 알림 발송량/억제율
  - 넓은 질의 상위 목록
  - 오발송/중복 알림 사례

### 67.12 화면/IA 시사점
- 저장 검색 진입점 후보:
  1. 검색 결과 상단 `이 조건 저장`
  2. 마이페이지 `저장한 검색`
  3. zero-result 화면 `이 검색 알림 받기`
- `zero-result + 저장 검색` 조합은 특히 중요하다.
  - 지금 결과가 없어도, 나중에 매물이 생기면 즉시 재방문을 유도할 수 있다.
- 저장 검색 목록에는 아래가 보여야 한다.
  - 이름
  - 핵심 조건 요약
  - 최근 일치 결과 수 또는 마지막 알림 시각
  - 활성/일시중지 상태

### 67.13 분석 이벤트 / KPI 후보
- `saved_search_create`
- `saved_search_update`
- `saved_search_pause`
- `saved_search_delete`
- `saved_search_match_found`
- `saved_search_notification_sent`
- `saved_search_notification_open`
- `saved_search_result_click`
- `saved_search_zero_result_subscribe`

핵심 KPI 후보:
- 검색 사용자 대비 저장 검색 생성률
- zero-result 검색 대비 저장 검색 전환율
- 저장 검색 알림 오픈율
- 알림 수신 후 상세 진입률 / 채팅 시작률
- 저장 검색 기반 거래 완료 기여율
- 저장 검색 해지율 / 과다 알림 신고율

### 67.14 오픈 질문
- 저장 검색 알림을 MVP에 포함할지, 저장만 먼저 포함하고 알림은 Post-MVP로 둘지 확정 필요
- `reserved 포함` 검색을 저장한 사용자에게 `reserved -> available` 복귀 알림을 기본 ON으로 둘지 검토 필요
- 검색 결과가 매우 많은 질의에서 실시간 푸시 대신 digest만 허용하는 임계치를 몇 건으로 둘지 필요
- 가격 하락 알림의 최소 변동폭(절대값/비율)과 category별 예외가 필요한지 확인 필요

## 68. 찜(Favorite) / 관심목록 / 상태복귀 알림 계약
### 68.1 목표
- `찜`을 단순 북마크가 아니라 **재방문 탐색, 가격/상태 추적, 거래 재개 진입점**으로 정의한다.
- 저장 검색과 역할을 분리하되, 둘이 중복되지 않도록 해석 경계를 명확히 한다.
- 홈, 찜 목록, 상세, 알림함, 푸시, 검색 랭킹이 동일한 favorite 기준을 사용하도록 공통 계약을 둔다.

### 68.2 도메인 정의
| 개념 | 의미 | 대표 사용처 |
|---|---|---|
| `Favorite` | 사용자가 특정 매물을 다시 보고 싶다고 표시한 관계 | 상세, 목록 카드, 찜 목록 |
| `FavoriteListing` | 찜한 매물의 화면용 projection | 찜 목록, 홈 모듈 |
| `FavoriteState` | 찜 관계의 현재 유효 상태 | active / archived / removed_by_user |
| `FavoriteAlertPreference` | 해당 찜에 대해 어떤 변화를 알릴지 | 가격변경, 상태복귀, 종료 |

구분 원칙:
- `Favorite`는 **개별 매물 추적**이다.
- `SavedSearch`는 **조건 집합 추적**이다.
- 같은 사용자가 둘 다 사용할 수 있지만, UI에서는 각각의 목적이 다르게 보여야 한다.

### 68.3 사용자 가치와 해석 원칙
- 사용자는 찜을 통해 아래 중 하나를 기대한다.
  1. 나중에 다시 비교하기
  2. 상대 반응/상태를 기다리기
  3. 가격이 바뀌거나 거래 가능 상태로 돌아오면 다시 보기
  4. 채팅 진입 전 후보를 모아두기
- 따라서 찜은 단순 `좋아요`나 인기 신호라기보다 **개인 추적 리스트** 성격이 우선이다.
- 공개 화면에서 `favoriteCount`는 노출할 수 있지만, 개별 사용자의 찜 사실은 비공개로 유지한다.

### 68.4 Favorite 상태 모델
| favoriteState | 의미 | 사용자 기본 액션 | 시스템 처리 |
|---|---|---|---|
| `active` | 현재 사용자가 추적 중 | 찜 해제, 알림 설정 변경, 상세 진입 | 가격/상태 이벤트 구독 가능 |
| `archived` | 매물이 완료/취소/만료되어 기록성 보관 상태 | 기록 보기, 유사 매물 보기, 찜 정리 | 알림 기본 중단 |
| `suppressed` | 사용자는 찜 유지 중이나 알림만 억제됨 | 알림 재개 | 푸시/인앱 일부 억제 |
| `removed_by_user` | 사용자가 직접 찜 해제 | 없음 | 목록/알림에서 제거 |
| `deleted_by_policy` | 대상 매물 삭제/비공개/제재 등으로 추적 불가 | 없음 | soft landing 또는 제거 |

원칙:
- `Listing.status`가 바뀌어도 찜 관계는 즉시 삭제하지 않는다.
- 사용자가 완료/취소 매물을 나중에 참고할 수 있도록 기본은 `archived` 전환을 우선 검토한다.
- 운영 숨김/차단 매물은 사용자 목록에 남기더라도 상세 원문 대신 제한 메시지로 대체해야 한다.

### 68.5 찜 생성/해제 규칙
- 비회원은 찜 불가, 로그인 유도 필요
- 작성자는 본인 매물을 찜할 수 없도록 하는 기본안을 우선 권장한다.
- 동일 사용자-매물 조합은 최대 1개 active 관계만 허용한다.
- 이미 찜한 매물을 다시 누르면 `active -> removed_by_user` 토글로 처리한다.
- 네트워크 재시도/중복 탭을 고려해 찜 API는 upsert/멱등적으로 동작해야 한다.

### 68.6 상태 변화에 따른 찜 보존 규칙
| Listing 상태/가시성 변화 | Favorite 처리 | 사용자 목록 표시 |
|---|---|---|
| `available -> reserved` | `active` 유지 | `예약중` 배지 + 상태복귀 알림 옵션 유지 |
| `reserved -> available` | `active` 유지 | `다시 거래 가능` 배지/정렬 가산 가능 |
| `pending_trade -> completed` | `archived` 전환 후보 | 기록 카드 + 유사 매물 CTA |
| `available/reserved -> cancelled` | `archived` 전환 | `거래 취소됨` 표시 |
| `visibility=hidden`(작성자 숨김) | `archived` 또는 `suppressed` | 상세 원문 축소 |
| `visibility=blocked`(운영 차단) | `deleted_by_policy` 또는 제한 표시 | 정책상 비공개 메시지 |

세부 원칙:
- 사용자가 찜한 매물이 `reserved`가 되었다고 자동 해제하지 않는다.
- `reserved -> available` 복귀는 찜의 핵심 가치이므로 별도 이벤트/알림 트리거로 다룬다.
- 완료/취소 매물은 `찜 목록`의 기본 탭에서는 숨기고 `보관됨` 탭에서 접근 가능하게 하는 안이 적절하다.

### 68.7 찜 알림 기본 정책
기본 알림 후보:
1. 가격 변경
2. `reserved -> available` 상태 복귀
3. `hidden/blocked`가 아닌 범위에서 재노출/재게시
4. 거래 종료(선택적)

이벤트별 기본안:
| 이벤트 | 인앱 | 푸시 | 기본값 | 억제 조건 |
|---|---|---|---|---|
| 가격 하락 | O | 선택적 | ON 후보 | 짧은 시간 내 반복 변동, 변동폭 미달 |
| 가격 상승 | 선택 | 기본 OFF | OFF | 사용자 혼란 방지 |
| `reserved -> available` | O | O | ON | 최근 동일 복귀 알림 발송 |
| `cancelled/completed` | O | 선택 | OFF 또는 약한 ON | 사용자가 이미 상세/채팅에서 확인 |
| 운영 비공개 전환 | O | 중요도 따라 | ON | 세부 사유 비공개 원칙 |

원칙:
- 찜 알림은 마케팅이 아니라 거래 기회 추적 알림으로 취급한다.
- 동일 매물에서 가격이 여러 번 바뀌어도, 사용자 체감상 의미 있는 변화만 보내야 한다.
- 푸시 본문에는 민감 정보보다 `무엇이 달라졌는지`만 요약한다.

### 68.8 찜 목록 정보구조(IA)
필수 탭 기본안:
1. `활성`
2. `예약중 포함`
3. `보관됨`

카드 필수 요소:
- 썸네일
- 매물 요약명
- 현재 가격 라벨
- 서버/카테고리
- 현재 상태 배지
- 마지막 변화 문구(`가격이 내려갔어요`, `다시 거래 가능해졌어요`)
- 빠른 액션(`상세`, `채팅`, `찜 해제`)

정렬 기본안:
1. 최근 상태변화 있음
2. 액션 가능(`available`) 우선
3. 최근 찜한 순
4. 가격/최근 활동 보조 정렬

빈 상태 UX:
- 찜 없음: 탐색 CTA + 저장검색 소개
- 활성 없음/보관만 있음: `최근 종료된 매물만 있어요` + 유사 매물 추천
- 정책 비공개만 남음: 상세 원문 대신 제한 안내

### 68.9 홈/상세/검색과의 연결 규칙
#### 홈
- 홈에는 `찜한 매물 중 다시 거래 가능해진 항목`, `가격이 바뀐 항목`, `예약중으로 기다리는 항목`을 1~2개 모듈로 노출 가능
- 단순 전체 찜 수보다 **행동 필요성**이 높은 찜을 우선 노출한다.

#### 매물 상세
- 찜 여부를 명확히 보여주고, 찜한 상태에서 알림 옵션을 세부 조정할 수 있어야 한다.
- `reserved` 상태 상세에서는 `찜 유지 + 거래 가능해지면 알림` 문구가 핵심 가치가 된다.

#### 검색/목록
- 카드 단위 찜 토글 지원
- 찜한 항목은 검색 결과에서 아이콘/미세 배지로 표시 가능
- 검색 랭킹은 `내가 찜했다`는 사실을 개인화에 활용할 수 있으나, 공개 랭킹 신호와는 분리해야 한다.

### 68.10 데이터 모델 후보
#### `Favorite`
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `favoriteId` | 필수 | 찜 식별자 |
| `userId` | 필수 | 찜한 사용자 |
| `listingId` | 필수 | 대상 매물 |
| `favoriteState` | 필수 | active / archived / suppressed / removed_by_user / deleted_by_policy |
| `priceAlertEnabled` | 필수 | 가격변경 알림 여부 |
| `availabilityAlertEnabled` | 필수 | 상태복귀 알림 여부 |
| `lastNotifiedAt` | 선택 | 최근 알림 시각 |
| `lastSeenListingVersion` | 선택 | 사용자가 마지막으로 본 매물 변경 버전 |
| `createdAt` | 필수 | 찜 생성 시각 |
| `updatedAt` | 필수 | 마지막 갱신 시각 |
| `removedAt` | 선택 | 해제 시각 |

#### `FavoriteListingProjection`
- `userId`
- `listingId`
- `favoriteState`
- `listingStatus`
- `listingVisibility`
- `displayPriceLabel`
- `changeSummaryLabel`
- `thumbnailUrl`
- `sortPriority`
- `nextRecommendedAction`
- `lastMeaningfulChangeAt`

원칙:
- `favoriteCount` 집계와 개인 `Favorite` 관계를 분리 저장한다.
- 찜 목록 화면은 Listing 원본 조인보다 projection/read model을 우선 고려하는 것이 유리하다.

### 68.11 API 후보
#### 사용자용
- `POST /listings/{listingId}/favorite`
- `DELETE /listings/{listingId}/favorite`
- `GET /me/favorites`
- `PATCH /me/favorites/{favoriteId}`
- `POST /me/favorites/read`

#### 응답 projection 후보
```json
{
  "favoriteId": "fav_123",
  "favoriteState": "active",
  "listing": {
    "listingId": "listing_123",
    "title": "+7 장검 팝니다",
    "status": "reserved",
    "displayPriceLabel": "50만 아데나",
    "summaryText": "+7 장검 · 데포로쥬",
    "thumbnailUrl": "https://..."
  },
  "changeSummary": {
    "code": "returned_to_available",
    "label": "다시 거래 가능해졌어요",
    "changedAt": "2026-03-13T16:20:00+09:00"
  },
  "notificationPreference": {
    "priceAlertEnabled": true,
    "availabilityAlertEnabled": true
  },
  "availableActions": ["open_listing", "remove_favorite"]
}
```

API 원칙:
- 목록/상세 카드에서 즉시 토글 가능한 저지연 API여야 한다.
- `GET /me/favorites`는 커서 페이지네이션과 상태별 필터(`active`, `archived`)를 지원하는 것이 바람직하다.
- `PATCH /me/favorites/{favoriteId}`는 알림 설정 토글과 상태 전환(`archive`, `suppress`)을 함께 다룰 수 있다.

### 68.12 분석 이벤트 후보
- `favorite_add`
- `favorite_remove`
- `favorite_list_view`
- `favorite_item_open`
- `favorite_alert_toggle`
- `favorite_notification_sent`
- `favorite_notification_open`
- `favorite_return_to_available_click`

핵심 KPI 후보:
- 상세 조회 대비 찜 전환율
- 찜 목록 재방문률
- 찜 알림 오픈율
- 찜 기반 상세 재진입률 / 채팅 시작률
- `reserved -> available` 복귀 알림 후 거래 전환율
- 보관된 찜 비중 / 수동 정리율

### 68.13 운영/정책 시사점
- 운영 숨김/차단된 매물은 찜 목록에도 원문을 그대로 노출하지 않는다.
- 차단 관계가 생긴 상대의 매물은 찜 목록 유지 여부를 정책 결정해야 하나, 기본안은 **자동 제거 대신 제한 표시 + 재노출 금지**가 안전하다.
- 찜 수는 공개 인기 지표로 쓸 수 있지만, 조작/허위 인기 유도 탐지를 위해 내부 안티어뷰즈 모니터링이 필요하다.
- 동일 사용자의 대량 찜/해제 반복은 랭킹 조작이나 봇 시그널로 활용할 수 있으나, 일반 사용자의 비교 탐색 행동과 구분해야 한다.

### 68.14 오픈 질문
- MVP에서 찜 알림 세부 설정까지 포함할지, 기본 ON/OFF만 제공할지 확정 필요
- 완료/취소된 찜 매물을 얼마나 오래 `보관됨` 탭에 유지할지 결정 필요
- 작성자 본인 매물 찜을 금지할지, 허용하되 공개 집계에서 제외할지 확정 필요
- 차단된 상대 매물이 기존 찜 목록에 남아야 하는지, 즉시 숨겨야 하는지 정책 결정 필요

## 69. 핵심 enum / 코드셋 카탈로그 계약
### 69.1 목표
- PRD에 등장하는 상태값/사유코드/액션코드가 DB enum, OpenAPI schema, 프론트 타입, 분석 이벤트 속성에서 서로 다른 이름으로 분기되지 않도록 **정본 코드셋**을 정의한다.
- 사용자 노출 라벨과 내부 저장 코드를 분리해, 문구 변경이 있어도 DB/API 계약은 안정적으로 유지되게 한다.
- 이후 파생 문서(DDL, OpenAPI, 화면명세, 운영정책)가 “어떤 값이 public contract인지, 어떤 값이 internal-only인지”를 명확히 구분하도록 한다.

### 69.2 코드셋 설계 공통 원칙
1. **저장 코드는 영어 snake_case 또는 lower enum**으로 고정하고, 화면 라벨은 별도 i18n/label map에서 관리한다.
2. **공개 API enum**은 additive change만 허용한다. 기존 값의 의미 변경/삭제는 major 버전 업 없이 금지한다.
3. **내부 운영 enum**은 세분화 가능하되, public enum으로 노출될 때는 안정적인 상위 코드로 매핑한다.
4. `unknown`은 저장 enum 값으로 남발하지 않고, 클라이언트 fallback 표현용으로만 사용한다.
5. 상태 enum과 사유 코드는 분리한다. 예: `status=cancelled`, `reasonCode=no_show_confirmed`.
6. 같은 사건을 여러 객체가 참조할 때는 가능한 한 동일 코드셋을 재사용한다. 중복 정의는 피한다.

### 69.3 코드셋 분류
| 분류 | 용도 | 예시 | 외부 노출 여부 |
|---|---|---|---|
| 공개 상태 enum | 사용자 화면/API 계약 | `listingStatus`, `tradeThreadStatus` | 노출 |
| 내부 상태 enum | 운영/배치/이벤트 세분화 | `completionStage`, `agingState` | 제한 노출 |
| 사유 코드 | 취소/숨김/제재/억제 원인 | `cancellationReasonCode` | 일부만 노출 |
| 액션 코드 | 버튼/인가/감사로그 계약 | `create_chat`, `confirm_reservation` | 노출 |
| 정책 코드 | 경고/차단/배너/운영 탐지 | `RESERVED_WAITING_ONLY` | 노출/비노출 혼합 |
| 분석 속성 enum | 이벤트 세그먼트 | `notificationChannel=push` | 내부 중심 |

### 69.4 공개 도메인 enum 정본
#### 69.4.1 ListingStatus
```text
available
reserved
pending_trade
completed
cancelled
```
원칙:
- 검색/목록/상세/찜/홈 등 사용자 탐색 surface는 이 5개를 기준으로 동작한다.
- `completed_pending_confirmation` 같은 값은 공개 enum에 추가하지 않고 내부 completion stage로 흡수한다.

#### 69.4.2 ListingVisibility
```text
public
hidden
blocked
```
원칙:
- `status`와 독립 축이다.
- `blocked`는 운영/정책 차단 성격이며, 사용자 자가 숨김은 `hidden`을 우선 사용한다.

#### 69.4.3 ListingType
```text
sell
buy
```

#### 69.4.4 PriceType
```text
fixed
negotiable
offer
```

#### 69.4.5 TradeMethod
```text
in_game
offline_pc_bang
either
```

#### 69.4.6 ChatStatus
```text
open
reservation_proposed
reservation_confirmed
trade_due
deal_completed
deal_cancelled
report_locked
```
원칙:
- `deal_completed`는 채팅방 타임라인 관점의 종결 상태이고, 최종 거래 판정은 `TradeCompletion`을 기준으로 본다.

#### 69.4.7 ReservationStatus
```text
proposed
confirmed
expired
cancelled
fulfilled
no_show_reported
```

#### 69.4.8 ReviewVisibilityStatus
```text
visible
hidden
pending_moderation
```

#### 69.4.9 ReportStatus
```text
submitted
triaged
investigating
resolved
rejected
```

#### 69.4.10 NotificationType
```text
chat
reservation
status
review
report
system
```

### 69.5 내부 전용 또는 제한 노출 enum 정본
#### 69.5.1 CompletionStage
```text
none
requested
confirmed_by_counterparty
auto_confirmed
disputed
resolved_completed
resolved_not_completed
closed_without_review
```
원칙:
- DB와 운영 백오피스는 이 enum을 기준으로 완료/분쟁 처리 단계를 저장한다.
- 외부 앱에는 필요 최소한만 노출하고, 보통은 사용자용 배지(`completion_waiting_me`, `dispute_open`)로 변환한다.

#### 69.5.2 DisputeStatus
```text
open
waiting_statement
under_review
resolved
closed
```

#### 69.5.3 RestrictionScope
```text
warning
listing_only
chat_only
trust_limited
temporary_suspend
permanent_ban
```

#### 69.5.4 TradeThreadAgingState
```text
fresh
response_due
stale_waiting_counterparty
stale_waiting_owner
stale_mutual_inactive
auto_closed_candidate
```
원칙:
- 목록 노출용 enum이 아니라 SLA, 운영 큐, 리마인드 배치의 판단 입력으로 사용한다.

### 69.6 사용자 노출 가능한 액션 코드(Action Code) 정본
아래 액션 코드는 `availableActions`, 인가 정책, QA 케이스, 감사 로그에서 동일하게 쓴다.

| 객체 | 액션 코드 |
|---|---|
| Listing | `create_chat`, `favorite`, `unfavorite`, `edit_listing`, `change_listing_status`, `duplicate_listing`, `report_listing`, `bump_listing` |
| Chat | `send_message`, `propose_reservation`, `confirm_reservation`, `cancel_reservation`, `mark_completed`, `dispute_completion`, `report_chat`, `block_user`, `mute_chat` |
| Reservation | `edit_reservation`, `ack_location_change`, `mark_no_show`, `mark_fulfilled` |
| Review | `create_review`, `edit_review`, `report_review` |
| Notification | `mark_read`, `open_deep_link` |
| Admin | `triage_report`, `hide_listing`, `restore_listing`, `lock_chat`, `unlock_chat`, `restrict_user`, `restore_user`, `resolve_dispute` |

원칙:
- 액션 코드는 UI 문구가 아니라 **서버 명령 의미**를 표현해야 한다.
- `availableActions`에 없는 액션은 클라이언트가 가정으로 노출하지 않는다.
- Admin 액션도 같은 네이밍 규칙을 써 감사 로그와 권한 번들이 바로 연결되게 한다.

### 69.7 대표 사유 코드셋 카탈로그
#### 69.7.1 ListingCancellationReasonCode
```text
seller_cancelled
buyer_cancelled
reservation_expired
no_show_confirmed
item_unavailable
price_changed_off_platform
duplicate_listing
policy_violation
other
```

#### 69.7.2 ReservationCancellationReasonCode
```text
counterparty_declined
schedule_conflict
location_conflict
no_response
seller_cancelled
buyer_cancelled
policy_blocked
other
```

#### 69.7.3 CompletionDisputeReasonCode
```text
trade_not_completed
wrong_counterparty
partial_trade_disagreement
meeting_mismatch
no_show_claim
abusive_completion_request
other
```

#### 69.7.4 ModerationReasonCode 상위 분류
```text
fake_listing
suspected_scam
abuse_or_harassment
spam_or_flood
prohibited_content
sensitive_info_exposed
reserved_abuse
review_retaliation
ban_evasion
other
```
원칙:
- 운영 내부에서는 더 세분화된 하위 코드를 둘 수 있으나, PRD 기준 public/analytics 상위 코드는 위 수준으로 우선 통일한다.

### 69.8 공개 라벨과 코드 저장의 분리 규칙
- DB/API에는 `reserved`, `pending_trade`, `abuse_or_harassment` 같은 **코드값**만 저장한다.
- 화면 문구는 별도 라벨 맵으로 관리한다.

예시:
```json
{
  "status": "reserved",
  "displayStatusLabel": "예약중",
  "statusTone": "warning"
}
```

원칙:
- `displayStatusLabel`은 UX/언어 정책에 따라 바뀔 수 있지만 `status`는 바뀌지 않는다.
- 서버가 라벨을 내려주더라도, i18n 필요 시 클라이언트가 코드 기반 라벨 테이블을 병행 사용할 수 있다.

### 69.9 OpenAPI / DDL / 프론트 타입 파생 규칙
#### DDL
- Postgres enum 사용 여부는 구현 단계에서 결정하되, 최소한 PRD 코드셋과 1:1 매핑되는 CHECK 제약 또는 참조 테이블이 있어야 한다.
- 내부 세분화가 잦은 코드셋(`moderationReasonCode`, `policyHintCode`)은 enum보다 코드 테이블이 변경 친화적이다.

#### OpenAPI
- 공개 API에 노출되는 enum은 schema에 `enum`으로 명시한다.
- 내부 전용 enum은 admin/internal schema로 분리하거나 `x-internal: true` 성격 메타를 붙여 문서 범위를 구분한다.

#### 프론트 타입
- `type ListingStatus = 'available' | 'reserved' | ...` 식의 literal union을 생성한다.
- unknown 값 fallback 컴포넌트를 두되, 서버에 없는 가짜 enum 값을 추가하지 않는다.

### 69.10 코드셋 변경 관리 원칙
| 변경 유형 | 허용 여부 | 예시 | 처리 원칙 |
|---|---|---|---|
| 신규 값 추가 | 조건부 허용 | `notificationType=marketing` 추가 | fallback/UI 검토 후 additive 배포 |
| 기존 값 삭제 | 금지(동일 major) | `reserved` 삭제 | major 버전 또는 migration 필요 |
| 기존 의미 변경 | 금지 | `completed`를 미확정 완료로 재정의 | 새 enum 도입 |
| 라벨 변경 | 허용 | `예약중` → `우선 거래중` | 코드 변경 없이 라벨만 변경 |
| 내부 enum 세분화 | 허용 | `disputed` 하위 상태 추가 | public 매핑 유지 필요 |

### 69.11 QA / 운영 체크리스트 파생 포인트
- 같은 상태가 목록/상세/채팅/내 거래/알림에서 서로 다른 코드로 저장되거나 번역되지 않는가
- `availableActions`와 실제 쓰기 API 권한 판정이 같은 코드셋을 참조하는가
- 감사 로그가 UI 액션명 대신 액션 코드를 저장하는가
- 운영자가 사유 코드를 선택했을 때 사용자 노출 문구와 내부 로그 문구가 분리되는가
- enum 신규 값 추가 시 구버전 앱이 안전한 fallback 배지를 보여주는가

### 69.12 후속 파생 산출물
이 섹션을 기준으로 다음 문서를 직접 만들 수 있다.
1. `docs/schema/codebook.md` — 공개/내부 코드셋 표준 문서
2. OpenAPI enum schema (`ListingStatus`, `ChatStatus`, `RestrictionScope`, `ActionCode`)
3. Postgres enum 또는 lookup table DDL 초안
4. 프론트 타입 자동생성 규칙(TypeScript literal unions)
5. 운영자 사유 코드 선택 UI와 사용자 노출 문구 매핑표

## 70. MVP OpenAPI 산출물 팩 / Starter DDL 파생 기준(초안)
### 70.1 목표
- 현재 PRD를 기준으로 실제 구현 착수용 `openapi.yaml`과 초기 migration/DDL 팩을 바로 만들 수 있도록 최소 산출물 경계를 정의한다.
- 제품/디자인/서버/모바일/운영이 같은 endpoint, schema, enum, 제약 조건을 승인 대상으로 보게 한다.
- 이 섹션은 구현 세부 기술 선택이 아니라 **무엇이 반드시 명세되어야 build-now 상태라고 볼 수 있는지**를 정리한다.

### 70.2 OpenAPI MVP 산출물 팩 범위
MVP OpenAPI 초안은 아래 5개 endpoint group을 반드시 포함해야 한다.

| 그룹 | 필수 endpoint 범위 | 목적 |
|---|---|---|
| Auth / Me | `POST /auth/login`, `POST /auth/logout`, `GET /me`, `PATCH /me/profile` | 인증/프로필 최소 흐름 |
| Listings | `GET /listings`, `POST /listings`, `GET /listings/{listingId}`, `PATCH /listings/{listingId}`, `POST /listings/{listingId}/status`, favorite 관련 | 탐색/등록/상태 관리 |
| Chat / Reservation | `POST /listings/{listingId}/chats`, `GET /chats`, `GET /chats/{chatRoomId}`, `GET /chats/{chatRoomId}/messages`, `POST /chats/{chatRoomId}/messages`, 예약 생성/확정/취소 | 거래 협의와 예약 실행 |
| Completion / Review / Report | 완료 요청/확정/분쟁, 후기 생성, 신고 생성/조회 | 종결/신뢰/분쟁 |
| Notifications / Admin MVP | `GET /notifications`, `POST /notifications/read`, 최소 admin report/listing/user action API | 알림 복구와 운영 처리 |

원칙:
- catalog/saved search/search suggestion 같은 확장 surface는 별도 파일 또는 Post-MVP 섹션으로 분리 가능하다.
- MVP OpenAPI는 “모든 후보 endpoint”보다 “실제로 앱이 호출해야 하는 승인 범위”를 우선 명시한다.

### 70.3 OpenAPI 공통 컴포넌트 필수 목록
아래 schema/component는 endpoint보다 먼저 공통 정의돼야 한다.

#### 70.3.1 Envelope / Context
- `ViewerContext`
- `AvailableActionCode[]`
- `PolicyHint`
- `StateReason`
- `PagedCursorRequest`
- `PagedCursorResponseMeta`
- `IdempotencyKeyHeader` 명세

#### 70.3.2 Domain summary schemas
- `ListingCard`
- `ListingDetail`
- `AuthorProfileSummary`
- `ChatThreadSummary`
- `ChatMessage`
- `ReservationCard`
- `TradeThreadProjection`
- `TradeCompletionSummary`
- `ReviewSummary`
- `NotificationItem`
- `ReportSummary`

#### 70.3.3 Error / conflict schemas
- `ErrorResponse`
- `ValidationErrorResponse`
- `ConflictErrorResponse`
- `PolicyBlockedErrorResponse`
- `RateLimitedErrorResponse`

원칙:
- 모든 read endpoint는 필요 수준에 맞게 `viewerContext`, `availableActions`, `policyHints` 포함 여부를 명시해야 한다.
- 모든 write endpoint는 성공 응답 외에 최소 401/403/404/409/422 응답 계약을 가져야 한다.

### 70.4 Typed error / conflict 계약 최소 기준
OpenAPI 초안은 아래 항목을 endpoint별로 빠짐없이 적어야 한다.
1. **권한 실패**: 로그인 필요, 객체 접근 불가, 역할 부족
2. **상태 충돌**: 종결 매물 채팅 생성, 중복 예약 확정, 이미 확정된 완료 요청 재확인
3. **정책 차단**: 차단/정지/민감정보 차단/신고 잠금
4. **멱등 재호출**: 동일 요청 재시도 시 동일 성공 또는 replay 응답
5. **validation 세부 필드**: `field`, `reason`, `message`, `rejectedValue(optional)`

최소 예시:
```yaml
409:
  description: Listing status conflict
  content:
    application/json:
      schema:
        $ref: '#/components/schemas/ConflictErrorResponse'
```

### 70.5 Enum freeze 범위
실제 OpenAPI/DDL 작성 전 아래 공개 enum은 v1 major 안에서 의미 변경 없이 freeze 대상으로 본다.
- `ListingStatus`
- `ListingVisibility`
- `ChatStatus`
- `ReservationStatus`
- `TradeThreadStatus`
- `CompletionStage`(publicly surfaced subset only)
- `ReviewVisibilityStatus`
- `NotificationType`
- `RestrictionScope`
- `ActionCode`
- `PublicReasonCode`

원칙:
- 내부 세분화 enum은 admin/internal namespace로 분리 가능하나, public 응답에 매핑되는 subset은 먼저 고정한다.
- OpenAPI 초안에는 각 enum의 unknown fallback 원칙을 설명 주석 또는 description으로 남겨야 한다.

### 70.6 Starter DDL / migration 팩 최소 범위
초기 migration 팩은 아래 테이블/제약/인덱스를 최소 포함해야 한다.

#### 70.6.1 Core transactional tables
- `users`
- `user_profiles`
- `listings`
- `listing_images`
- `favorites`
- `chat_rooms`
- `chat_messages`
- `reservations`
- `trade_completions`
- `reviews`
- `reports`
- `moderation_actions`
- `notifications`
- `restrictions`
- `audit_logs`

#### 70.6.2 권장 보조/이력 테이블
- `listing_status_history`
- `reservation_status_history`
- `completion_stage_history`
- `report_status_history`
- `chat_participant_states`
- `chat_device_states`(MVP 포함 여부는 선택이지만 schema placeholder 검토)
- `outbox_events`

#### 70.6.3 반드시 명시할 핵심 제약
- 동일 `listingId` 기준 활성 예약 중복 방지
- `chat_rooms`의 `(listing_id, seller_user_id, buyer_user_id)` 유니크
- `reviews`의 `(completion_id, reviewer_user_id)` 유니크
- `favorites`의 `(user_id, listing_id)` 유니크
- `reserved`/`pending_trade` 상태에서 필요한 FK 존재 조건
- soft delete/visibility 컬럼 포함 시 partial index 전략

### 70.7 Starter index / query path 기준
Starter migration은 실제 화면 조회를 기준으로 다음 인덱스를 빠뜨리면 안 된다.

| 화면/조회 | 필수 인덱스 방향 |
|---|---|
| 거래소 목록 | `listings(status, visibility, server_id, category_id, bumped_at/last_activity_at desc)` |
| 내 매물 | `listings(author_user_id, status, updated_at desc)` |
| 채팅 목록 | `chat_rooms(seller_user_id, last_message_at desc)`, `chat_rooms(buyer_user_id, last_message_at desc)` |
| 메시지 타임라인 | `chat_messages(chat_room_id, sent_at asc/desc)` |
| 예약 임박 처리 | `reservations(reservation_status, scheduled_at)` |
| 운영 신고 큐 | `reports(report_status, priority, created_at)` |
| 알림함 | `notifications(user_id, read_at, created_at desc)` |

원칙:
- read model/projection을 후속 도입하더라도 starter DDL은 source-of-truth 조회 경로의 최소 인덱스를 보장해야 한다.
- 검색 전문색인/별도 검색엔진 전환은 후속 결정이어도, `item_name_raw`, `title` 기준 기본 검색 전략은 starter 단계에서 명시해야 한다.

### 70.8 OpenAPI ↔ DDL traceability 규칙
- 각 write endpoint는 어떤 aggregate/table을 소유하는지 명시해야 한다.
- 각 public enum은 어떤 DB 컬럼 또는 code table에 저장되는지 역추적 가능해야 한다.
- 각 화면 projection 필드는 원천 테이블 또는 projection 테이블 출처가 표기돼야 한다.

권장 표기 예시:
| API schema field | 저장 위치 | 비고 |
|---|---|---|
| `listing.status` | `listings.status` | public enum |
| `listing.availableActions[]` | persisted 아님 | server-calculated |
| `tradeThread.unreadCount` | projection or derived | source-of-truth는 chat participant state |
| `notification.deepLinkPayload` | `notifications.deep_link_payload_json` | JSON payload |

### 70.9 산출물 acceptance 기준
다음 조건을 만족해야 “PRD로부터 OpenAPI/DDL starter pack 도출 가능” 상태로 본다.
1. 모든 MVP endpoint에 request/response schema가 있다.
2. 모든 write endpoint에 typed error와 conflict case가 있다.
3. 모든 public enum이 codebook/DDL/OpenAPI에서 같은 이름을 쓴다.
4. 핵심 uniqueness / FK / partial index 규칙이 migration에 반영된다.
5. 목록/상세/채팅/내 거래/운영 큐의 주요 query path가 인덱스로 설명된다.
6. `availableActions`, `viewerContext`, `policyHints`의 계산성 필드 여부가 명확하다.
7. soft delete / retention / audit 대상 컬럼이 최소 starter schema에 포함된다.

### 70.10 다음 실제 산출물 순서
이 섹션을 기준으로 다음 문서를 순서대로 만드는 것을 권장한다.
1. `docs/openapi/openapi.yaml` 초안
2. `docs/schema/codebook.md` 정식판
3. `docs/schema/relational-model.md` 또는 ERD 초안
4. `db/migrations/0001_init.sql` starter migration
5. `docs/operations/admin-rbac-matrix.md` 및 action-code mapping

오픈 질문:
- `chat_device_states`를 starter migration에 실제 포함할지 placeholder로 둘지
- `CompletionStage` 전체를 DB enum으로 고정할지 code table로 열어둘지
- listings 검색을 starter 단계에서 Postgres FTS로 갈지 prefix + trigram 혼합으로 갈지


## 66. 증빙(Evidence) 접근등급 / 다운로드 / 외부반출 통제 정책(초안)
### 66.1 목표
- 신고/분쟁 처리에 필요한 증빙은 충분히 보존하되, 일반 채팅 첨부와 같은 수준으로 무제한 열람·다운로드되지 않게 통제한다.
- 사용자용 첨부, 운영 검토용 증빙, 법적/보안 이슈가 있는 민감 파일을 같은 저장소에 두더라도 **접근등급(access level)** 과 **반출정책(export policy)** 는 분리해야 한다.
- 백오피스, 스토리지, API, 감사로그, retention 정책이 동일한 evidence vocabulary를 사용하도록 한다.

### 66.2 증빙 객체 정의
본 섹션에서 `Evidence`는 아래 범위를 포함한다.
- 신고 접수 시 사용자가 첨부한 스크린샷/이미지/텍스트 보조자료
- 완료 분쟁/노쇼 분쟁에서 제출한 추가 소명 첨부
- 정책 위반 탐지(OCR/민감정보 탐지) 결과와 연결된 원본/마스킹본
- 운영자가 사건 검토 중 생성한 redacted copy 또는 case export bundle

일반 채팅 첨부(`chat attachment`)와의 차이:
- 일반 채팅 첨부는 거래 참여 UX의 일부
- Evidence는 운영 판정/이의제기/감사 대응의 일부
- 따라서 동일 파일 포맷을 써도 접근 주체, 다운로드 규칙, 노출 기간이 달라야 한다.

### 66.3 증빙 분류 체계
| 분류 축 | 값 후보 | 설명 |
|---|---|---|
| `evidenceKind` | `report_attachment` / `dispute_attachment` / `ocr_capture` / `moderation_export` | 생성 출처 |
| `accessLevel` | `participant_limited` / `staff_masked` / `staff_sensitive` / `legal_hold_only` | 기본 열람 권한 |
| `downloadDisposition` | `inline_preview_only` / `signed_download` / `staff_only_download` / `export_only` | 다운로드 방식 |
| `redactionState` | `raw` / `masked` / `redacted` / `quarantined` | 마스킹/격리 상태 |
| `retentionClass` | `chat_standard` / `dispute_standard` / `safety_sensitive` / `legal_hold` | 보존 기준 |
| `sourceVisibility` | `user_uploaded` / `system_generated` / `staff_generated` | 생성 주체 |

원칙:
- 같은 파일도 사건 단계에 따라 `accessLevel`, `downloadDisposition` 이 승격/강등될 수 있다.
- `redactionState` 변경은 원본 덮어쓰기가 아니라 파생본 추가를 기본 원칙으로 한다.

### 66.4 접근등급 기본안
| accessLevel | 사용자 당사자 | 거래 상대 | 운영자 | 민감도 설명 |
|---|---|---|---|---|
| `participant_limited` | 업로더 본인과 사건 상대에게 제한 미리보기 가능 | 가능(정책 범위 내) | 가능 | 일반 분쟁 스크린샷 등 |
| `staff_masked` | 업로더 본인만 원본 참조, 상대는 요약 또는 비노출 | 제한적 또는 불가 | 마스킹본 기본 열람 | 개인정보 포함 가능 |
| `staff_sensitive` | 사용자 상호 비공개 | 불가 | 권한 있는 운영자만 열람 | 계좌/연락처/신분 단서 등 고위험 |
| `legal_hold_only` | 불가 | 불가 | 지정 사건/법무 승인 하에만 | 법적 요청 또는 심각 사건 보존 |

기본 정책:
- 신고/분쟁 첨부는 기본 `participant_limited` 또는 `staff_masked`에서 시작한다.
- OCR/민감정보 탐지 결과가 있으면 자동으로 `staff_masked` 이상으로 상향 가능하다.
- 운영자가 수동 상향/하향할 수 있으나 반드시 사유 코드와 감사 로그를 남긴다.

### 66.5 미리보기 / 원본 / 다운로드 원칙
#### 사용자 앱/웹
- 기본은 **미리보기 중심**이며, 원본 파일 직접 URL을 사용자에게 장시간 노출하지 않는다.
- 사건 상대에게 증빙을 보여줄 필요가 있더라도, 원본보다 `masked preview` 또는 축약 이미지 우선 노출을 기본안으로 한다.
- 사용자 단 download는 사건 맥락상 꼭 필요한 경우에만 짧은 TTL의 signed URL로 허용한다.

#### 운영 백오피스
- 기본 화면은 썸네일/미리보기/마스킹본 우선
- 원본 다운로드는 `staff_only_download` 이상 권한과 사유 입력이 필요
- 동일 사건에서 여러 운영자가 파일을 반복 다운로드하지 않도록 인앱 preview 우선 UX를 제공한다.

#### 파일 제공 방식 기본안
| 상황 | 기본 제공 방식 | TTL/가드레일 |
|---|---|---|
| 사용자 사건 상세 미리보기 | API 게이트웨이 프록시 또는 단기 signed preview URL | 1~5분 수준 |
| 운영 미리보기 | 인증세션 기반 프록시 | 세션 권한 재검증 |
| 운영 원본 다운로드 | 1회성 signed download URL + 다운로드 사유 기록 | 1분 이하 + 단일 사용 후보 |
| 외부 반출 export | 비동기 bundle 생성 + 승인 기록 + 암호화 전달 후보 | 건별 승인 |

### 66.6 일반 채팅 첨부와 분쟁 증빙의 분리 원칙
| 항목 | 일반 채팅 첨부 | 분쟁/신고 증빙 |
|---|---|---|
| 기본 목적 | 대화/거래 협의 | 운영 판정/소명 |
| 기본 노출 | 채팅 참여자 | 사건 관계자 + 운영자 |
| 원본 다운로드 | 상대적으로 완화 가능 | 더 엄격, 기본 제한 |
| 마스킹/워터마크 | 선택 | 기본 적용 우선 검토 |
| 감사로그 | 업로드/삭제 중심 | 열람/미리보기/다운로드/반출 모두 필수 |
| retentionClass | `chat_standard` | `dispute_standard` 이상 |

권장 기본안:
- 저장 버킷이 같더라도 object prefix, key namespace, IAM policy, CDN rule은 분리한다.
- 분쟁/신고 evidence는 public CDN 캐시를 사용하지 않는다.

### 66.7 워터마크 / 마스킹 / 적색편집(redaction) 정책
- 운영자 미리보기에는 사건번호, 열람자 role, 열람시각 기반 동적 워터마크를 우선 검토한다.
- 사용자 상호 열람이 허용된 증빙도 필요 시 `case watermark`를 적용해 무단 재유포 억제 신호를 제공한다.
- 전화번호, 계좌번호, 메신저 ID, 좌석번호/정확 주소 등은 미리보기 단계에서 자동 마스킹 후보로 본다.
- `redacted` 파생본은 원본 훼손 없이 별도 object로 저장하며, 어느 redaction rule이 적용됐는지 메타데이터로 남긴다.
- 악성 파일/정책 위반 파일은 `quarantined` 상태로 전환해 일반 열람을 차단한다.

### 66.8 외부 반출(export) 정책
외부 반출은 아래 경우로 한정한다.
1. 사용자의 본인 사건 자료 요청
2. 운영 escalation / 내부 보안 리뷰
3. 법적 요청 또는 정책상 정당한 보존/제출 필요

기본 원칙:
- 운영자가 증빙 파일을 로컬로 무제한 저장하는 방식은 금지하고, 가급적 case export bundle을 통해 통제한다.
- export는 단일 파일보다 `사건 메타 + 파일 목록 + 생성 시각 + redaction 상태`를 포함한 bundle 단위가 바람직하다.
- export 실행 시 최소 필수 메타:
  - `exportId`
  - `caseId(reportId/disputeId)`
  - `requestedBy`
  - `approvedBy(optional)`
  - `exportReasonCode`
  - `includedEvidenceIds`
  - `createdAt`
  - `expiresAt`

### 66.9 감사 로그 / 관측성 원칙
Evidence 관련 아래 액션은 모두 감사 대상이다.
- 업로드
- 미리보기 열람
- 원본 열람
- 다운로드
- accessLevel 변경
- redaction 생성
- quarantine 전환/해제
- export 생성/다운로드/만료

감사 로그 최소 필드:
- `auditActionCode`
- `evidenceId`
- `caseType`
- `caseId`
- `actorType` (`user` / `staff` / `system`)
- `actorId`
- `result` (`allowed` / `denied` / `expired`)
- `reasonCode`
- `sourceIp/device(optional, policy permitting)`
- `createdAt`

운영 대시보드/이상징후 룰 후보:
- 특정 운영자가 단기간 대량 evidence download
- 동일 사건 export 반복 생성
- 권한 부족 사용자의 evidence 접근 시도 반복
- quarantined 파일 비율 급증

### 66.10 API / 백오피스 파생 기준
#### 사용자용 API 후보
- `POST /reports/{reportId}/attachments`
- `POST /disputes/{disputeId}/attachments`
- `GET /evidence/{evidenceId}/preview`
- `POST /evidence/{evidenceId}/download-token` (정책 허용 시에만)

#### 관리자용 API 후보
- `GET /admin/evidence/{evidenceId}`
- `POST /admin/evidence/{evidenceId}/access-level`
- `POST /admin/evidence/{evidenceId}/redactions`
- `POST /admin/evidence/{evidenceId}/quarantine`
- `POST /admin/cases/{caseId}/exports`

응답 필드 후보:
```json
{
  "evidenceId": "ev_123",
  "evidenceKind": "dispute_attachment",
  "accessLevel": "staff_masked",
  "redactionState": "masked",
  "downloadDisposition": "inline_preview_only",
  "availableActions": ["preview_masked", "request_export"],
  "auditRequired": true
}
```

### 66.11 데이터 모델 시사점
후보 엔티티:
- `Evidence`
- `EvidenceVariant` (raw/masked/redacted/thumbnail)
- `EvidenceAccessLog`
- `EvidenceExport`

`Evidence` 최소 필드 후보:
- `evidenceId`
- `ownerType` (`report` / `dispute` / `message` / `moderation_case`)
- `ownerId`
- `uploadedByUserId(optional)`
- `sourceVisibility`
- `accessLevel`
- `downloadDisposition`
- `retentionClass`
- `redactionState`
- `quarantineState`
- `storageKeyCurrent`
- `mimeType`
- `byteSize`
- `createdAt`
- `retentionUntil`

`EvidenceVariant` 최소 필드 후보:
- `variantId`
- `evidenceId`
- `variantType` (`raw` / `thumbnail` / `masked_preview` / `redacted_export`)
- `storageKey`
- `generatedBy`
- `generatedAt`
- `watermarkProfile`

### 66.12 retention / 삭제 / 복구 연계 원칙
- evidence는 원본과 파생본을 같은 파기 시점으로 묶되, `legal_hold`가 걸리면 전체 variant를 함께 보존한다.
- 사용자가 신고를 취소해도 사건이 이미 처리/조사 중이면 evidence 즉시 삭제를 보장하지 않는다.
- 복구는 삭제 복구라기보다 `accessLevel` 하향 또는 quarantine 해제 형태가 많으므로, 상태이력 저장이 필수다.
- 운영자가 redaction 오류를 수정해도 이전 variant는 감사 추적을 위해 일정 기간 유지하는 것이 바람직하다.

### 66.13 오픈 질문
- 사용자 본인 사건 자료 다운로드를 MVP에 열 것인지, 운영 문의 기반으로만 처리할 것인지?
- 운영 원본 다운로드를 Senior Moderator 이상으로 제한할지, 특정 action code 승인 모델로 갈지?
- evidence preview를 앱 내 전용 프록시로만 제공할지, 단기 signed URL을 혼용할지?
- 분쟁 증빙에 기본 워터마크를 항상 적용할지, 민감도 기반 선택 적용할지?

### 66.14 증빙 다운로드 / 반출 워크플로우 상태 계약(초안)
#### 66.14.1 목표
- `preview`, `download`, `export` 요청을 단순 URL 발급이 아니라 **권한 평가 → 승인 필요 여부 판단 → 짧은 수명 토큰 발급 → 사용/만료/회수 추적** 흐름으로 정의한다.
- 백오피스, API, 스토리지 signed URL, 감사로그가 같은 상태값을 사용하도록 해 구현 중 임의 분기를 줄인다.
- 특히 원본 다운로드와 외부 반출은 "누가 왜 요청했고 누가 승인했는지"를 사건 단위로 재구성 가능해야 한다.

#### 66.14.2 요청 유형 정의
| requestType | 의미 | 대표 호출 주체 | 기본 승인 정책 |
|---|---|---|---|
| `preview_masked` | 마스킹/워터마크된 미리보기 | 사용자/운영자 | 즉시 허용 우선 |
| `preview_raw` | 원본 미리보기 | 운영자 | 민감도 따라 승인 필요 |
| `download_masked` | 마스킹본 저장/다운로드 | 운영자, 제한적 사용자 | 사건/권한 따라 즉시 또는 승인대기 |
| `download_raw` | 원본 파일 다운로드 | 운영자 | 기본 승인 필요 |
| `export_bundle` | 사건 단위 묶음 반출 | 운영자/법무 플로우 | 승인 필수 |

원칙:
- `preview_masked`는 UX 복구용 기능이므로 가장 짧고 단순한 경로를 유지한다.
- `download_raw`, `export_bundle`은 링크 발급 전에 별도 Access Request 객체를 남기는 것을 기본안으로 한다.

#### 66.14.3 Access Request 상태 모델
| evidenceAccessRequestState | 의미 | 다음 가능 상태 |
|---|---|---|
| `requested` | 사용자가 다운로드/반출 요청 생성 | `granted`, `approval_required`, `denied`, `cancelled` |
| `approval_required` | 즉시 허용 불가, 승인 대기 | `granted`, `denied`, `expired` |
| `granted` | 토큰/세션 발급 가능 상태 | `consumed`, `expired`, `revoked` |
| `denied` | 정책/권한상 거절 | 종료 |
| `consumed` | 토큰이 실제 사용됨 | 종료 |
| `expired` | 승인 또는 토큰 유효시간 만료 | 종료 |
| `revoked` | 발급 후 회수됨 | 종료 |
| `cancelled` | 요청자가 사용 전 취소 | 종료 |

설계 원칙:
- `granted`와 실제 파일 다운로드 성공은 구분한다. 링크를 발급받았더라도 다운로드가 실패하거나 사용되지 않을 수 있기 때문이다.
- 승인형 요청은 `approval_required` 상태를 별도 저장해 운영 큐와 연동해야 한다.

#### 66.14.4 Export Bundle 상태 모델
| exportBundleState | 의미 | 비고 |
|---|---|---|
| `requested` | 반출 요청 접수 | 메타만 생성 |
| `approved` | 승인 완료, 번들 생성 대기 | 승인자/사유 고정 |
| `building` | 파일 수집/워터마크/redaction 적용 중 | 비동기 워커 처리 |
| `ready` | 다운로드 가능 | 짧은 TTL 권장 |
| `partially_failed` | 일부 파일 누락/실패 | 운영자 안내 필요 |
| `downloaded` | 최초 다운로드 완료 | 추가 다운로드 허용 여부 정책화 |
| `expired` | 만료되어 재생성 필요 | 원본 번들 폐기 |
| `revoked` | 승인 철회/사건 상태 변경으로 회수 | 즉시 접근 차단 |

권장 기본안:
- `export_bundle`은 동기 응답으로 파일을 바로 주지 않고 `requested -> building -> ready` 비동기 흐름을 기본으로 한다.
- `ready` 상태 번들은 재다운로드보다 재요청을 우선해 링크 장기 생존을 막는다.

#### 66.14.5 승인 정책 매트릭스
| requestType | participant_limited | staff_masked | staff_sensitive | legal_hold_only |
|---|---|---|---|---|
| `preview_masked` | 즉시 허용 | 즉시 허용 | 불가 | 불가 |
| `preview_raw` | 불가 | senior 이상 또는 승인 | admin 승인 | 불가 |
| `download_masked` | 사건 당사자 요청 시 제한 허용 후보 | senior 이상 즉시 또는 승인 | admin 승인 | 불가 |
| `download_raw` | 불가 | 승인 필요 | 승인 필요 + 사유 필수 | 불가 |
| `export_bundle` | 불가 | 승인 필요 | admin/법무 승인 | 지정 승인만 |

보강 원칙:
- `legal_hold_only`는 일반 제품 워크플로우 밖으로 두고, 실사용 API보다 별도 운영/법무 플로우에 가깝게 설계한다.
- 사용자 당사자에게도 `participant_limited` 원본 무제한 다운로드를 기본 허용하지 않는다.

#### 66.14.6 토큰/URL 수명주기 원칙
| grantScope | 설명 | 권장 TTL | 추가 제약 |
|---|---|---|---|
| `inline_preview_session` | 앱/웹 내 일시 미리보기 | 1~5분 | 화면/세션 종속 |
| `single_download` | 1회 다운로드 | 1분 이하 | 1회 사용 후 즉시 만료 |
| `bundle_download` | export bundle 다운로드 | 10~30분 | 재사용 횟수 제한 후보 |
| `staff_review_session` | 백오피스 검토 세션 | 세션 수명 내 단기 갱신 | 권한 재검증 |

후보 필드:
- `downloadGrantScope`
- `grantedAt`
- `grantExpiresAt`
- `maxConsumeCount`
- `consumedCount`
- `lastConsumedAt`

#### 66.14.7 회수(revocation) / 무효화 트리거
다음 상황에서는 발급된 토큰/번들을 회수 또는 무효화할 수 있어야 한다.
- 사건 상태가 `resolved` 또는 `appeal_locked`로 바뀌며 접근 범위가 축소된 경우
- evidence의 `accessLevel`이 상향 조정된 경우
- 운영자가 오승인/오반출을 발견한 경우
- 법적 보존/보안 사고 대응으로 즉시 차단이 필요한 경우
- 사용자가 탈퇴/제재되어 기존 grant 범위를 유지하면 안 되는 경우

후보 코드:
- `evidenceTokenRevocationReason = access_level_upgraded | case_closed_scope_reduced | approval_retracted | security_incident | actor_restricted`

#### 66.14.8 사용자/운영 UX 원칙
- 사용자 앱에서는 승인 필요한 다운로드 요청에 대해 `다운로드 준비 중`, `운영 확인 필요`, `만료되어 다시 요청 필요` 정도의 단순 상태만 노출한다.
- 운영 백오피스는 각 요청에 대해 `요청자`, `사건`, `사유`, `승인자`, `유효시간`, `실사용 여부`를 한 줄로 볼 수 있어야 한다.
- `export_bundle`은 ready 후에도 기본적으로 파일이 아니라 `번들 구성 요약 + 만료 시각 + 다운로드 버튼`을 먼저 노출하는 것이 안전하다.

#### 66.14.9 API / 데이터 모델 파생 기준
추가 API 후보:
- `POST /evidence/{evidenceId}/access-requests`
- `GET /evidence/{evidenceId}/access-requests/{requestId}`
- `POST /admin/evidence-access-requests/{requestId}/approve`
- `POST /admin/evidence-access-requests/{requestId}/deny`
- `POST /admin/evidence-access-requests/{requestId}/revoke`
- `GET /admin/case-exports/{exportId}`

추가 엔티티 후보:
- `EvidenceAccessRequest`
- `EvidenceGrant`
- `EvidenceExportBundle`

`EvidenceAccessRequest` 최소 필드 후보:
- `requestId`
- `evidenceId(optional)`
- `caseId`
- `requestType`
- `requestedByActorType`
- `requestedByActorId`
- `requestReasonCode`
- `requestState`
- `approvedByActorId(optional)`
- `decisionMemo(optional)`
- `createdAt`
- `resolvedAt(optional)`

#### 66.14.10 운영/관측성 파생 포인트
분석/운영 지표 후보:
- raw download 요청 대비 승인율
- grant 발급 후 실제 consume 비율
- export bundle 생성 평균 시간
- 만료 후 재요청률
- revoke 발생률과 사유 분포

이상징후 룰 후보:
- 동일 운영자가 짧은 시간에 다수 `download_raw` 승인
- 사건 종결 후에도 repeated grant 재발급
- denied 후 즉시 동일 request를 반복 생성하는 사용자/운영자 패턴
- partially_failed export가 특정 저장소/파일 유형에서 반복 발생

## 67. 운영 액션 코드 / Admin RBAC 매트릭스 / 승인 체인 계약(초안)
### 67.1 목표
- 백오피스 버튼, Admin API 인가, 운영 감사 로그, runbook이 동일한 `actionCode` vocabulary를 사용하도록 정리한다.
- 단순 역할명(`moderator`, `admin`)만으로 권한을 추론하지 않고, **행동 단위 허용/승인 정책**으로 운영 리스크를 낮춘다.
- 신고/분쟁/증빙/제재 처리에서 `누가 실행 가능하고`, `누가 승인해야 하며`, `어떤 로그를 남겨야 하는지`를 구현 가능한 수준으로 명문화한다.

### 67.2 설계 원칙
1. 역할(Role)과 액션(Action Code)을 분리한다.
2. UI 노출 여부와 실제 실행 가능 여부를 모두 서버가 최종 판단한다.
3. 민감 액션은 역할만으로 충분하지 않고 `approvalPolicy`, `caseAssignment`, `reasonCode`, `ticket/reference`를 함께 요구한다.
4. 자기 사건(self-case) 처리, 자기 제재 해제, 자기 액션 감사로그 수정은 금지한다.
5. emergency/break-glass 액션은 허용하되, 사후 검토와 별도 감사 태그를 강제한다.

### 67.3 운영 역할 재정의
| 역할 | 설명 | 기본 범위 |
|---|---|---|
| `cs_operator` | 접수/분류/기본 응답 담당 | 조회, triage, 비파괴 메모/할당 |
| `moderator` | 일반 정책 집행 담당 | 경고, 임시 숨김, 저위험 제한, 채팅 잠금 |
| `senior_moderator` | 고위험 사건/복구/민감 열람 담당 | 기간 정지, 복구 승인, 원본 열람 일부 |
| `admin` | 정책 owner 및 최고 권한 | 영구 제재, export 승인, break-glass 승인 |
| `system` | 배치/자동화 actor | 자동 만료, 자동확정, 알림/큐 처리 |

원칙:
- 역할은 사람의 직급이 아니라 **기본 권한 묶음**이다.
- 실제 실행 가능 여부는 `role + actionCode + case scope + approval requirement`의 교집합으로 판단한다.

### 67.4 actionCode 네이밍 원칙
- 형식: `{domain}.{object}.{verb}`
- 예시:
  - `report.case.assign`
  - `report.case.triage`
  - `listing.moderation.hide_temp`
  - `listing.moderation.restore`
  - `chat.moderation.lock`
  - `chat.moderation.unlock`
  - `user.restriction.warn`
  - `user.restriction.suspend_temp`
  - `user.restriction.ban_permanent`
  - `evidence.asset.view_raw`
  - `evidence.asset.export`
  - `dispute.case.resolve_completed`
  - `dispute.case.resolve_not_completed`

원칙:
- actionCode는 UI 문구가 아니라 **불변에 가까운 시스템 식별자**로 관리한다.
- 감사 로그/권한 테이블/OpenAPI/운영 매뉴얼에서 같은 문자열을 재사용한다.

### 67.5 운영 액션 그룹
| 액션 그룹 | 대표 actionCode | 설명 |
|---|---|---|
| 사건 접수/분류 | `report.case.assign`, `report.case.triage`, `report.case.request_info` | 사건 소유권/우선순위/추가자료 요청 |
| 콘텐츠 노출 제어 | `listing.moderation.hide_temp`, `listing.moderation.hide_policy`, `listing.moderation.restore` | 매물 숨김/복구 |
| 대화 제어 | `chat.moderation.lock`, `chat.moderation.unlock`, `chat.moderation.mask_message` | 채팅 잠금, 메시지 마스킹 |
| 사용자 제한 | `user.restriction.warn`, `user.restriction.limit_listing`, `user.restriction.limit_chat`, `user.restriction.suspend_temp`, `user.restriction.ban_permanent`, `user.restriction.restore` | 계정/기능 제한 |
| 분쟁 판정 | `dispute.case.resolve_completed`, `dispute.case.resolve_not_completed`, `dispute.case.close_no_fault` | 완료/불발/무조치 판정 |
| 증빙 접근 | `evidence.asset.view_masked`, `evidence.asset.view_raw`, `evidence.asset.download_raw`, `evidence.asset.export` | 미리보기/원본/외부반출 |
| 운영 예외 | `admin.break_glass.execute`, `policy.override.apply` | 긴급 우회/정책 override |

### 67.6 역할별 허용 액션 매트릭스(기본안)
| actionCode | cs_operator | moderator | senior_moderator | admin | 추가 조건 |
|---|---|---|---|---|---|
| `report.case.assign` | 가능 | 가능 | 가능 | 가능 | 자기 자신에게 재할당 허용 |
| `report.case.triage` | 가능 | 가능 | 가능 | 가능 | reasonCode 필수 |
| `report.case.request_info` | 가능 | 가능 | 가능 | 가능 | dueAt 권장 |
| `listing.moderation.hide_temp` | 불가 | 가능 | 가능 | 가능 | reportId 또는 caseRef 필요 |
| `listing.moderation.hide_policy` | 불가 | 가능 | 가능 | 가능 | policyCode 필수 |
| `listing.moderation.restore` | 불가 | 불가 | 가능 | 가능 | 복구 사유 필수 |
| `chat.moderation.lock` | 불가 | 가능 | 가능 | 가능 | 사건 연결 필수 |
| `chat.moderation.unlock` | 불가 | 불가 | 가능 | 가능 | 자기 잠금 해제 단독 처리 금지 |
| `chat.moderation.mask_message` | 불가 | 가능 | 가능 | 가능 | 원문 보관 전제 |
| `user.restriction.warn` | 불가 | 가능 | 가능 | 가능 | expiry 없음 |
| `user.restriction.limit_listing` | 불가 | 가능 | 가능 | 가능 | duration/사유 필수 |
| `user.restriction.limit_chat` | 불가 | 가능 | 가능 | 가능 | duration/사유 필수 |
| `user.restriction.suspend_temp` | 불가 | 불가 | 가능 | 가능 | 기간 제한 필수 |
| `user.restriction.ban_permanent` | 불가 | 불가 | 불가 | 가능 | 2인 승인 필수 |
| `user.restriction.restore` | 불가 | 불가 | 가능 | 가능 | 원조치 참조 필수 |
| `dispute.case.resolve_completed` | 불가 | 불가 | 가능 | 가능 | 증빙/판정 메모 필수 |
| `dispute.case.resolve_not_completed` | 불가 | 불가 | 가능 | 가능 | 재오픈/취소 후속액션 명시 |
| `dispute.case.close_no_fault` | 불가 | 불가 | 가능 | 가능 | 양측 통지 필요 |
| `evidence.asset.view_masked` | 가능 | 가능 | 가능 | 가능 | 사건 컨텍스트 내 |
| `evidence.asset.view_raw` | 불가 | 불가 | 가능 | 가능 | 열람 사유 + 감사 태그 |
| `evidence.asset.download_raw` | 불가 | 불가 | 가능 | 가능 | 시간제한 URL/승인 정책 |
| `evidence.asset.export` | 불가 | 불가 | 불가 | 가능 | 2인 승인 + export 목적 필수 |
| `admin.break_glass.execute` | 불가 | 불가 | 불가 | 가능 | 사후 리뷰 강제 |
| `policy.override.apply` | 불가 | 불가 | 불가 | 가능 | 만료시각/영향범위 필수 |

### 67.7 approvalPolicy 기본안
| actionCode 패밀리 | approvalPolicy | 설명 |
|---|---|---|
| `report.case.*` | `single_actor` | 일반 triage/할당은 단독 처리 가능 |
| `listing.moderation.hide_*` | `single_actor_with_reason` | 숨김은 빠른 조치 우선 |
| `listing.moderation.restore` | `single_actor_senior` | 오조치 복구는 senior 이상 |
| `chat.moderation.unlock` | `single_actor_senior` | 잠금 해제는 재노출 리스크 고려 |
| `user.restriction.warn` / `limit_*` | `single_actor_with_reason` | 저위험 제재 |
| `user.restriction.suspend_temp` | `dual_review_recommended` | 1인 실행 가능하나 P1/P2는 사후 승인 권장 |
| `user.restriction.ban_permanent` | `dual_review_required` | 최소 2인 승인 필수 |
| `evidence.asset.view_raw` | `single_actor_senior_with_justification` | 민감 원본 열람 |
| `evidence.asset.export` | `dual_review_required` | 외부 반출은 가장 엄격 |
| `admin.break_glass.execute` | `dual_review_posthoc_required` | 즉시 실행 후 사후 승인 필수 |

### 67.8 사건 소유권(case assignment) 규칙
- 모든 신고/분쟁/고위험 evidence 사건은 `caseOwnerUserId`, `caseQueue`, `assignedAt`를 가져야 한다.
- `cs_operator`는 사건 생성 직후 triage와 assign까지만 수행 가능하다.
- 사건이 `under_review`로 바뀌면 단일 owner를 기본으로 하되, `secondaryReviewerUserId(optional)`를 둘 수 있다.
- `ban_permanent`, `evidence export`, `break_glass` 사건은 secondary reviewer 지정이 필수다.
- unassigned 상태의 P1 사건은 SLA 위반으로 운영 대시보드에 별도 노출한다.

### 67.9 자기사건(self-case) / 이해상충 방지 규칙
- 운영자는 자신이 신고자/피신고자/거래당사자인 사건을 직접 판정할 수 없다.
- 자신의 이전 액션을 스스로 최종 승인하거나 스스로 영구 제재를 복구할 수 없다.
- 시스템은 아래 conflict flag를 계산할 수 있어야 한다.
  - `isReporterConflict`
  - `isTargetConflict`
  - `isPriorActorConflict`
  - `isSecondaryReviewerConflict`
- conflict 발생 시 UI에서 실행 버튼을 비활성화하고 다른 reviewer 할당을 요구한다.

### 67.10 break-glass / 긴급 우회 정책
- 대상: 대규모 개인정보 노출, 명백한 사기 파급, 법적/보안 긴급상황
- 허용 action 예시:
  - 즉시 전역 비노출
  - 즉시 강제 잠금
  - 원본 증빙 긴급 확보
- 필수 입력:
  - `emergencyReasonCode`
  - `impactScope`
  - `expectedRollbackPlan`
  - `postReviewDueAt`
- break-glass 실행 후 24시간 내 사후 검토 기록이 없으면 운영 경보를 발생시킨다.

### 67.11 Admin API / 백오피스 응답 계약 시사점
백오피스 상세 API는 최소한 아래를 함께 반환하는 것이 바람직하다.

```json
{
  "caseId": "rep_123",
  "viewerRole": "senior_moderator",
  "allowedAdminActions": [
    "chat.moderation.lock",
    "user.restriction.suspend_temp",
    "evidence.asset.view_masked"
  ],
  "approvalRequirements": {
    "user.restriction.suspend_temp": "dual_review_recommended",
    "evidence.asset.export": "not_allowed_for_viewer"
  },
  "conflictFlags": {
    "isReporterConflict": false,
    "isTargetConflict": false,
    "isPriorActorConflict": true
  }
}
```

원칙:
- 프론트는 역할명만 보고 버튼을 노출하지 않고, `allowedAdminActions`와 `approvalRequirements`를 우선 사용한다.
- 같은 역할이어도 사건 종류/소유 여부/conflict에 따라 허용 액션이 달라질 수 있어야 한다.

### 67.12 감사 로그 최소 필드
모든 admin action은 아래 필드를 남겨야 한다.

| 필드 | 설명 |
|---|---|
| `auditLogId` | 감사 로그 식별자 |
| `actionCode` | 실행 액션 코드 |
| `actorUserId` | 실행 운영자 |
| `actorRole` | 실행 시점 역할 |
| `targetType` / `targetId` | 대상 객체 |
| `caseRefType` / `caseRefId` | 연결 사건 |
| `reasonCode` | 구조화 사유 |
| `noteText` | 자유 메모 |
| `approvalPolicy` | 적용된 승인 정책 |
| `approvedByUserIds` | 승인자 목록 |
| `conflictCheckResult` | self-case 검사 결과 |
| `beforeSnapshot` / `afterSnapshot` | 핵심 diff |
| `createdAt` | 실행 시각 |
| `breakGlass` | 긴급 우회 여부 |

원칙:
- `noteText`는 자유 입력이지만, `reasonCode` 없이 note만으로 실행하는 액션은 금지한다.
- diff 전체 저장이 부담되면 최소한 핵심 상태 필드와 권한 영향 필드는 남겨야 한다.

### 67.13 DB / permission bundle 파생 기준
후속 DDL/OpenAPI에서 아래 구조를 직접 파생할 수 있어야 한다.
- `admin_role`
- `admin_action_code`
- `admin_permission_bundle`
- `admin_role_action_grant`
- `admin_case_assignment`
- `admin_action_approval`
- `admin_conflict_check_result`
- `audit_log`

권장 원칙:
- role과 action grant를 코드로 하드코딩하기보다 테이블/seed 기반으로 관리해 정책 변경을 쉽게 한다.
- 단, MVP에서는 seed 고정값으로 시작하되 runtime editable admin RBAC UI는 Post-MVP로 둔다.

### 67.14 오픈API / runbook 파생 포인트
- OpenAPI:
  - 각 Admin endpoint에 `x-action-code`, `x-approval-policy` 메타데이터를 붙일 수 있어야 한다.
  - 403 응답은 단순 `FORBIDDEN` 외에 `ADMIN_CONFLICT_OF_INTEREST`, `ADMIN_APPROVAL_REQUIRED`, `ADMIN_ROLE_NOT_GRANTED` 같은 typed error를 가질 수 있다.
- 운영 runbook:
  - P1 사기 의심 대응 runbook
  - evidence export 승인 runbook
  - permanent ban 2인 승인 runbook
  - break-glass 사후 검토 checklist
- QA:
  - 역할별 버튼 노출
  - conflict 플래그 발생 시 실행 차단
  - approval required 액션의 다단계 승인 흐름

### 67.15 오픈 질문
- `suspend_temp`를 senior 단독 허용으로 둘지, P1/P2에서만 dual approval 강제로 둘지?
- evidence raw download와 export를 완전히 분리할지, 같은 승인 체인으로 볼지?
- admin role/action grant를 DB seed 고정으로 시작할지, 초기부터 config 기반으로 둘지?
- 자기 사건 금지 예외를 둘 필요가 있는지, 있다면 어떤 break-glass 절차로 통제할지?

## 68. 알림 인박스(Notification Inbox) 상태 / 액션 / 보존 계약(초안)
### 68.1 목표
- 알림함을 단순 로그가 아니라 **사용자가 거래를 다시 이어가기 위한 작업 큐**로 정의한다.
- 같은 사건을 푸시/인앱/배지/딥링크가 서로 다르게 해석하지 않도록 공통 상태 모델을 둔다.
- 예약, 완료 확인, 분쟁 소명, 운영 경고처럼 `다음 행동`이 중요한 알림은 읽음 여부와 `행동 완료 여부`를 분리해 관리한다.

### 68.2 핵심 원칙
1. 알림은 `발송 사실`보다 `사용자에게 아직 남아 있는 해야 할 일`을 더 잘 표현해야 한다.
2. `읽음(read)`과 `행동 완료(acted)`는 다른 상태다. 사용자가 열어봤다고 해서 거래가 해결된 것은 아니다.
3. 같은 거래 스레드에서 발생한 연쇄 이벤트는 가능한 한 묶어 보여 주되, 중요한 최신 상태를 잃어버리면 안 된다.
4. 만료된 알림도 운영/분쟁 관점에서는 사건 이력으로 남길 수 있어야 하지만, 사용자 인박스는 과도하게 쌓이지 않게 관리해야 한다.

### 68.3 알림 단위 정의
| 개념 | 의미 | 대표 예시 |
|---|---|---|
| `NotificationEvent` | 도메인 이벤트 기반 원천 사건 | 예약 제안 도착, 완료 요청 도착 |
| `NotificationInboxItem` | 사용자 알림함에 보이는 아이템 | `오늘 21시 예약 응답 필요` |
| `NotificationDelivery` | 채널별 발송 기록 | 인앱 저장, 푸시 발송, 푸시 실패 |
| `NotificationAction` | 알림에서 이어지는 실제 행동 | 예약 수락, 분쟁 소명 제출 |

원칙:
- 하나의 `NotificationEvent`가 사용자별로 1개 이상의 `NotificationInboxItem`을 만들 수 있다.
- 반대로 여러 이벤트가 하나의 인박스 아이템으로 병합될 수 있다. 예: 같은 채팅방의 연속 메시지, 같은 예약의 변경 연쇄.

### 68.4 인박스 상태 모델
`notificationState`는 사용자가 체감하는 인박스 상태를 나타낸다.

| notificationState | 의미 | 사용자 노출 | 전환 조건 |
|---|---|---|---|
| `unread` | 아직 열어보지 않음 | 강조 배지/굵은 표시 | 신규 생성 직후 |
| `read` | 상세 또는 딥링크 진입으로 확인함 | 일반 표시 | 사용자가 열람 |
| `acted` | 알림이 요구한 행동이 충족됨 | 완료/정리됨 표기 가능 | 예약 수락, 완료 확인 등 성공 |
| `dismissed` | 사용자가 명시적으로 치움 | 인박스 기본 목록에서 숨김 가능 | 스와이프 닫기, 숨기기 |
| `expired` | 시간 경과로 알림 의미가 사라짐 | 만료 표기 또는 자동 정리 | 예약안 만료, 소명 기한 만료 |
| `superseded` | 더 최신 상태에 의해 대체됨 | 기본 목록에서는 최신 건만 유지 | 예약 변경, 상태 갱신, 분쟁 단계 변경 |

원칙:
- `read`는 가볍고, `acted`는 강한 종결 상태다.
- `dismissed`는 사용자 인박스 UI 행동이며, 도메인 사건 자체를 삭제하지 않는다.
- `superseded`는 사용자가 놓친 과거 상태를 기록으로 남기되 기본 인박스 혼잡을 줄이기 위한 내부 상태다.

### 68.5 액션 상태 모델
행동이 필요한 알림은 `actionState`를 별도로 가진다.

| actionState | 의미 | 대표 알림 |
|---|---|---|
| `none` | 정보성 알림 | 새 후기 도착, 상태 복귀 알림 |
| `pending` | 사용자의 응답/확인이 필요 | 예약 제안, 완료 확인 요청 |
| `in_progress` | 일부 행동은 했지만 아직 종결 전 | 분쟁 소명 시작, 장소 변경 재확인 대기 |
| `resolved` | 필요한 행동이 완료됨 | 예약 수락 완료, 완료 확인 완료 |
| `missed` | 행동 기한을 놓침 | 소명 기한 만료, 응답기한 경과 자동확정 |

원칙:
- `notificationState=read`이면서 `actionState=pending`일 수 있다.
- 사용자 인박스 정렬은 `notificationState`보다 `actionState`와 시한 임박도를 우선한다.

### 68.6 알림 유형별 액션 계약
| notificationType | 기본 목적 | actionType 후보 | actionState 기본값 |
|---|---|---|---|
| `chat_message` | 대화 복귀 | `open_chat` | `none` |
| `reservation_proposed` | 예약 응답 유도 | `accept_reservation`, `reject_reservation`, `counter_reservation` | `pending` |
| `reservation_changed` | 변경 사항 재확인 | `review_change`, `ack_change` | `pending` 또는 `none` |
| `reservation_reminder` | 시간 임박 안내 | `open_trade_detail`, `send_arrival_message` | `none` |
| `completion_requested` | 완료 여부 확정 | `confirm_completion`, `dispute_completion` | `pending` |
| `dispute_update` | 분쟁 후속 행동 | `submit_statement`, `open_dispute` | `pending` 또는 `in_progress` |
| `review_request` | 후기 작성 유도 | `write_review` | `pending` |
| `listing_status_reopened` | 거래 재개 알림 | `open_listing`, `open_chat` | `none` |
| `moderation_notice` | 정책 결과 통지 | `open_policy_detail`, `submit_appeal` | `pending` 또는 `none` |

원칙:
- `chat_message`는 미읽음은 중요하지만 대부분 정보성이다.
- `reservation_proposed`, `completion_requested`, `dispute_update`는 명확한 작업 큐 성격이므로 홈/내 거래/알림함 우선순위가 높아야 한다.

### 68.7 병합/대체 규칙
#### 68.7.1 병합 키(`threadMergeKey`)
같은 거래 맥락의 알림 묶음을 위해 아래 키를 검토한다.
- 채팅성: `chatRoomId`
- 거래성: `tradeThreadId` 또는 `listingId + counterpartyUserId`
- 예약성: `reservationId`
- 분쟁성: `disputeId`
- 운영성: `restrictionId` 또는 `reportId`

#### 68.7.2 병합 원칙
- 같은 채팅방의 연속 메시지는 1개 인박스 아이템으로 묶을 수 있다.
- 예약 변경은 최신 예약 상태만 기본 노출하고, 이전 알림은 `superseded` 처리한다.
- 완료 요청 이후 분쟁이 열리면 `completion_requested` 인박스 아이템은 `superseded`, `dispute_update`가 최신 작업 큐가 된다.
- 운영 경고와 기능 제한은 사용자가 놓치면 안 되므로 과도한 병합을 지양하고 개별 아이템 유지 원칙을 우선한다.

### 68.8 우선순위/정렬 계약
`inboxPriorityTier`는 인박스와 홈 요약 카드에서 공통 사용한다.

| tier | 의미 | 대표 예시 |
|---|---|---|
| `P0` | 즉시 응답 필요, 시간 민감도 매우 높음 | 완료 확인 요청, 분쟁 소명 마감 임박 |
| `P1` | 거래 성사에 직접 영향 | 예약 제안, 예약 변경 재확인 |
| `P2` | 대화 유지/재방문 유도 | 새 메시지, 새 문의 |
| `P3` | 사후 정리/리텐션 | 후기 요청, 상태 복귀, 찜 알림 |
| `P4` | 참고성/기록성 | 정책 공지, 일반 시스템 안내 |

정렬 기본안:
1. `actionState in (pending, in_progress)`
2. `inboxPriorityTier`
3. `actionDueAt` 임박 순
4. `notificationState=unread`
5. 최신 `createdAt`

### 68.9 읽음/행동/배지 계산 규칙
#### 68.9.1 읽음 처리
- 사용자가 알림 리스트만 스크롤해 봤다고 자동 읽음 처리하지 않는다.
- 아이템 상세 진입, 관련 화면 진입, 또는 명시적 읽음 액션 때 `readAt`을 기록한다.
- 단, 푸시 탭으로 관련 화면에 직접 진입한 경우는 해당 알림을 읽음 처리할 수 있다.

#### 68.9.2 행동 완료 처리
- `actedAt`은 사용자가 실제 후속 행동을 성공시킨 시점이다.
- 예시:
  - 예약 수락 성공 → 관련 `reservation_proposed` 알림 `acted`
  - 완료 확인 성공 → 관련 `completion_requested` 알림 `acted`
  - 분쟁 소명 제출 1회만으로 끝나지 않는 케이스는 `in_progress` 유지 가능

#### 68.9.3 배지 계산
- 전역 앱 배지는 `notificationState=unread`만이 아니라 `actionState=pending`을 더 강하게 반영해야 한다.
- 홈/내 거래 배지는 아래 분리 계산을 권장한다.
  - `unreadBadgeCount`: 읽지 않은 알림 수
  - `actionRequiredCount`: 아직 행동 필요한 알림 수
- MVP UI에서는 하나의 숫자 배지로 시작하더라도 내부적으로는 둘을 구분 저장하는 것이 바람직하다.

### 68.10 만료/자동 정리 정책
| 알림 유형 | 만료 조건 | 사용자 표시 | 보관 원칙 |
|---|---|---|---|
| 예약 제안 | 예약안 만료/취소/수락 | `만료됨` 또는 `처리됨` | 사건 이력 보관 |
| 완료 확인 요청 | 확인/분쟁/자동확정 | `처리됨` | 사건 이력 보관 |
| 분쟁 소명 요청 | 기한 경과 | `기한 만료` | 운영 감사 목적 보관 |
| 새 메시지 | 사용자가 채팅 읽음 + 최신 메시지 없음 | 일반 읽음 상태 | 장기 인박스에서는 자동 정리 가능 |
| 후기 요청 | 작성 완료 또는 마감 경과 | `작성 완료`/`마감됨` | 기록성 약하므로 짧은 보존 가능 |
| 운영 제재 통지 | 이의제기 만료 후 | 기록 유지 | 장기 보관 우선 |

원칙:
- 자동 정리는 인박스 기본 목록에서 숨기는 의미이지, 감사/분석 원천 삭제를 뜻하지 않는다.
- 분쟁/운영/제재 관련 알림은 일반 리텐션보다 길게 보관하는 별도 class가 필요하다.

### 68.11 화면 요구사항 파생 포인트
#### 알림함 목록
- 필수 그룹 후보: `지금 처리할 일`, `새 소식`, `지난 알림`
- 각 셀은 최소 아래를 포함한다.
  - 아이콘/유형
  - 제목
  - 요약 본문
  - 시간
  - 상태 배지(`읽지 않음`, `응답 필요`, `만료됨`)
  - 대표 CTA 1개

#### 알림 상세/딥링크
- 알림은 가능한 한 원문 상세가 아니라 **행동할 수 있는 원본 객체 화면**으로 보내야 한다.
- 딥링크 실패 시 fallback 우선순위:
  1. 관련 거래 상세
  2. 채팅방
  3. 매물 상세
  4. 알림함 복귀 + 만료 안내

#### 홈/내 거래 연계
- `P0~P1`이며 `actionState=pending`인 알림은 홈 상단 `액션 필요 카드`와 내 거래 상단 정렬에 재사용할 수 있어야 한다.
- 즉, 알림함 projection과 홈/내 거래 action module이 같은 원천 vocabulary를 써야 한다.

### 68.12 데이터 모델 후보
`NotificationInboxItem` 필드 후보:
- `notificationId`
- `userId`
- `notificationType`
- `notificationState`
- `actionState`
- `inboxPriorityTier`
- `threadMergeKey`
- `title`
- `body`
- `deepLink`
- `sourceEntityType`
- `sourceEntityId`
- `actionDueAt(optional)`
- `readAt(optional)`
- `actedAt(optional)`
- `dismissedAt(optional)`
- `expiredAt(optional)`
- `supersededByNotificationId(optional)`
- `createdAt`
- `updatedAt`

`NotificationDelivery` 필드 후보:
- `deliveryId`
- `notificationId`
- `channel` (`in_app` / `push`)
- `deliveryState` (`queued` / `sent` / `delivered` / `failed` / `suppressed`)
- `suppressedReasonCode(optional)`
- `deliveredAt(optional)`
- `providerMessageKey(optional)`

### 68.13 API 후보 정리
#### 사용자용
- `GET /notifications`
- `POST /notifications/read`
- `POST /notifications/{notificationId}/dismiss`
- `POST /notifications/{notificationId}/act` (권장하지 않음, 원본 도메인 API 호출 우선)

권장 원칙:
- 예약 수락/완료 확인 같은 실제 행동은 알림 API가 아니라 원본 도메인 API에서 수행한다.
- 알림 API는 읽음/숨김/목록 조회처럼 **인박스 상태 관리**에 집중한다.

`GET /notifications` 응답 후보 필드:
```json
{
  "items": [
    {
      "notificationId": "noti_123",
      "notificationType": "completion_requested",
      "notificationState": "unread",
      "actionState": "pending",
      "inboxPriorityTier": "P0",
      "title": "거래완료 확인이 필요해요",
      "body": "상대가 거래완료를 요청했습니다.",
      "deepLink": "/trades/thread_123",
      "actionDueAt": "2026-03-14T20:00:00+09:00",
      "availableActions": ["open_trade_detail", "confirm_completion", "dispute_completion"]
    }
  ],
  "summary": {
    "unreadBadgeCount": 7,
    "actionRequiredCount": 2
  }
}
```

### 68.14 운영/분석 시사점
운영에서 확인해야 할 항목:
- 중요한 알림이 `read`는 되었지만 `acted`되지 않고 방치되는 비율
- 예약 제안/완료 요청 알림의 action completion latency
- `superseded` 과다 발생으로 사용자가 중간 상태를 이해하지 못하는지 여부
- 푸시 억제로 인해 인박스만 남은 사건의 후속 전환율

분석 이벤트 후보:
- `notification_inbox_view`
- `notification_item_open`
- `notification_dismiss`
- `notification_action_completed`
- `notification_expired`
- `notification_superseded`

핵심 KPI 후보:
- `completion_requested -> confirm/dispute` 평균 소요시간
- `reservation_proposed -> acted` 전환율
- `actionRequiredCount > 0` 사용자 중 24시간 내 처리율
- `read_without_action` 비율

### 68.15 오픈 질문
- 알림함에서 스와이프 dismiss를 얼마나 적극 허용할 것인가? `pending` 알림도 dismiss 가능하게 둘지 여부
- 새 메시지 알림과 거래 액션 알림을 완전히 같은 피드에 둘지, 상단 세그먼트로 분리할지
- 장기 보관이 필요한 운영/분쟁 알림을 사용자 알림함에서 언제 `지난 알림`으로 이동시킬지
- 홈/내 거래/알림함이 같은 `actionRequiredCount`를 쓸 때, 화면별 강조 임계치를 다르게 둘지

## 67. 신원확인(Verification) / 신뢰 레벨 / 제한 해제 계약(초안)
### 67.1 목표
- 린클은 완전 익명 게시판이 아니라 `실제 거래 완료율`을 높이는 서비스이므로, 계정 생성 직후 모든 기능을 동일하게 열기보다 위험도에 맞춰 신뢰를 점진적으로 부여해야 한다.
- 다만 초기 진입 장벽이 과도하면 매물 등록/문의 전환이 떨어지므로, `필수 최소 진입`과 `고위험 상황 추가 확인`을 분리한다.
- 제품, 서버, 운영, 프로필 노출, 제한 정책이 서로 다른 신뢰 기준을 쓰지 않도록 공통 vocabulary를 정의한다.

### 67.2 기본 원칙
1. **MVP 기본안은 실명 인증 강제가 아니라 단계적 신뢰 부트스트랩**이다.
2. 기본 거래 진입은 낮은 마찰로 허용하되, 고위험 행동에는 추가 확인을 요구할 수 있어야 한다.
3. 신뢰 신호는 `공개 가능한 신호`와 `운영 전용 위험 신호`를 분리한다.
4. 신원확인은 `더 많은 기능을 여는 수단`이지, 공개 프로필에 과도한 개인정보를 노출하는 근거가 되어서는 안 된다.
5. 동일 사용자의 신뢰 레벨 상승/하락은 온보딩, 매물 노출, 채팅 제한, 운영 큐 우선순위에 일관되게 반영되어야 한다.

### 67.3 핵심 객체/코드셋 후보
| 코드셋 | 의미 | 대표 값 후보 |
|---|---|---|
| `verificationLevel` | 계정이 통과한 확인 수준 | `none` / `basic_login_verified` / `contact_verified` / `enhanced_verified` |
| `tradeEligibilityState` | 현재 거래 기능 허용 상태 | `read_only` / `limited_trade` / `trade_enabled` / `review_required` / `restricted` |
| `trustSignalLevel` | 사용자에게 제한 공개 가능한 신뢰 단계 | `new` / `basic` / `active` / `trusted` |
| `restrictionLiftPolicy` | 제한 해제 조건 | `time_based` / `activity_based` / `verification_required` / `manual_review_required` |
| `verificationTriggerReason` | 추가 확인 요구 사유 | `new_account_high_risk_listing` / `spam_pattern_detected` / `repeat_reported` / `appeal_required` |

원칙:
- `verificationLevel`은 사실(fact), `tradeEligibilityState`는 현재 정책 결정(state), `trustSignalLevel`은 공개 표현(label)로 분리한다.
- 같은 사용자가 `verificationLevel=contact_verified`여도 반복 신고로 인해 `tradeEligibilityState=restricted`가 될 수 있다.

### 67.4 계정 단계와 기능 개방 기본안
| 계정 단계 | 설명 | 허용 기능 기본안 | 제한/가드레일 |
|---|---|---|---|
| `guest` | 로그인 전 탐색자 | 목록/상세 제한 조회 | 찜/채팅/신고/예약 불가 |
| `member_basic` | 로그인만 완료 | 찜, 검색 저장, 프로필 작성 시작 | 고빈도 문의/고가 매물 등록 제한 가능 |
| `member_trade_enabled` | 기본 거래 가능 상태 | 매물 등록, 채팅, 예약, 후기, 신고 | 위험도 기반 rate limit 적용 |
| `member_limited_trade` | 조건부 거래 상태 | 문의/채팅 일부 가능, 매물 등록 제한 가능 | 추가 확인 또는 시간 경과 필요 |
| `member_restricted` | 정책/운영 제한 상태 | 읽기 위주, 일부 기록 열람 | 신규 쓰기 액션 차단 |

권장 MVP 기본안:
- 로그인 완료 후 즉시 `member_trade_enabled`로 진입 가능하되, 신규 계정은 더 강한 rate limit과 탐지 규칙을 적용한다.
- 운영상 위험 신호가 높을 때만 `member_limited_trade` 또는 `review_required`로 격하한다.
- 별도 본인확인 강제는 MVP 비차단 기본안으로 두고, 고위험 상황에서만 추가 요구하는 쪽이 제품 방향에 맞다.

### 67.5 추가 확인(Verification) 트리거 규칙
다음 조건은 `verificationTriggerReason`을 발생시킬 수 있다.

| 트리거 | 기본 처리 | 비고 |
|---|---|---|
| 신규 계정이 고가/고위험 매물 등록 시도 | 경고 또는 추가 확인 요청 | 즉시 차단보다 soft gate 우선 |
| 짧은 시간 내 다수 채팅 생성/반복 취소 | rate limit 강화 + review 후보 | 안티어뷰즈 연계 |
| 다수 사용자 차단/신고 집중 | `limited_trade` 또는 운영 큐 | 자동 영구 제재 금지 |
| 이의제기/복구 요청 시 신원 확인 필요 | enhanced verification 후보 | 운영 신뢰 보강 |
| 반복된 외부 연락처 유도/사기 패턴 | 거래 제한 + 추가 확인 또는 수동 검토 | 고위험 |

원칙:
- 추가 확인은 가능한 한 `즉시 영구 차단`이 아니라 `기능 축소 + 추가 확인 + 운영 검토` 순으로 적용한다.
- 사용자는 “왜 더 확인이 필요한지”를 일반화된 사유로 안내받아야 한다. 내부 리스크 점수는 비공개다.

### 67.6 공개 프로필과의 연결 규칙
| 신호 | 공개 여부 | 노출 방식 |
|---|---|---|
| `verificationLevel` 상세 원문 | 비공개 기본 | 내부/운영 전용 |
| 신뢰 레벨 배지(`trustSignalLevel`) | 공개 가능 | `신규`, `활동 중`, `거래 경험 있음`, `신뢰 우수` 등 |
| 추가 확인 진행 중 여부 | 원칙적 비공개 | 필요한 경우 본인에게만 안내 |
| 제재 해제 후 복구 상태 | 부분 공개 또는 비공개 | 운영정책 확정 필요 |

원칙:
- 사용자가 인증을 더 했다는 사실은 `거래 판단 보조 신호`로 제한적으로 표현할 수 있지만, 인증 수단 자체(전화/실명/문서 등)를 공개 프로필에 그대로 노출하지 않는다.
- `trustSignalLevel`은 완료 거래 수, 응답성, 제재 이력, 최근 활동성과 함께 계산되며 인증만으로 단독 결정하지 않는다.

### 67.7 데이터 모델 후보
#### UserVerificationSnapshot
| 필드 | 설명 |
|---|---|
| `userId` | 사용자 식별자 |
| `verificationLevel` | 현재 최고 확인 수준 |
| `verificationUpdatedAt` | 마지막 갱신 시각 |
| `verificationSource` | `self_service` / `support_review` / `system` |
| `verificationExpiresAt(optional)` | 재확인 필요 시각 |
| `verificationFlagsJson` | 세부 확인 플래그(공개 비노출) |

#### UserTradeEligibility
| 필드 | 설명 |
|---|---|
| `userId` | 사용자 식별자 |
| `tradeEligibilityState` | 현재 거래 허용 상태 |
| `effectiveFrom` | 적용 시작 시각 |
| `effectiveUntil(optional)` | 종료 예정 시각 |
| `triggerReasonCode` | 제한/격상 사유 |
| `restrictionLiftPolicy` | 해제 방식 |
| `reviewRequired` | 운영 검토 필요 여부 |

설계 원칙:
- 확인 사실과 정책 결정을 같은 컬럼에 섞지 않는다.
- 운영/시스템이 `verificationLevel`을 올렸다고 해서 자동으로 모든 제한이 해제되지는 않을 수 있다.

### 67.8 API/화면 계약 시사점
#### 읽기 응답 후보
```json
{
  "me": {
    "verificationLevel": "basic_login_verified",
    "tradeEligibilityState": "trade_enabled",
    "trustSignalLevel": "basic",
    "policyHints": [
      {
        "code": "NEW_ACCOUNT_HIGHER_LIMITS_APPLY",
        "severity": "info",
        "message": "신규 계정은 초기 거래 활동에 일부 속도 제한이 적용될 수 있어요."
      }
    ]
  }
}
```

#### API 후보
- `GET /me/trust-status`
- `POST /me/verification/request`
- `GET /me/verification/status`
- `POST /admin/users/{userId}/verification/approve`
- `POST /admin/users/{userId}/trade-eligibility`

원칙:
- MVP에서 별도 verification UI가 과하면 `/me` 또는 `/me/trust-status`에 최소 상태만 포함해도 된다.
- 클라이언트는 제한 이유를 `tradeEligibilityState`와 `policyHints`로 해석하고, 인증 내부 세부를 직접 추론하지 않는다.

### 67.9 운영/제재/복구 연결 규칙
- `member_limited_trade`는 제재라기보다 **위험 완화 상태**로 취급한다.
- 제재 해제와 별도로 `verification_required` 정책을 둘 수 있다.
- 반복 오탐 사례가 있는 룰은 `manual_review_required`로 낮추고 자동 제한 강도를 완화해야 한다.
- 운영자는 아래를 한 화면에서 봐야 한다.
  1. 현재 `verificationLevel`
  2. 현재 `tradeEligibilityState`
  3. 최근 trigger reason
  4. 제한 해제 조건과 예상 시각
  5. 과거 verification/eligibility 변경 이력

### 67.10 MVP / Post-MVP 경계
#### MVP에 포함할 것
- `tradeEligibilityState` 기본 모델
- 신규 계정/고위험 행동에 대한 추가 제한 훅
- 본인 대상 제한 사유/안내 문구
- 운영자가 수동으로 제한/해제를 관리할 수 있는 최소 도구

#### Post-MVP 후보
- 강화된 본인확인 플로우
- 증빙 제출/재확인 만료 주기
- 신뢰 배지 계산식 고도화
- 제한 해제 자동화(`activity_based`, `time_based`) 세분화

### 67.11 오픈 질문
- MVP에서 `contact_verified` 수준을 실제 제품 기능으로 열지, 내부 확장 슬롯으로만 둘지?
- 고가 매물의 기준선을 금액/카테고리/신고 이력 중 무엇으로 정의할지?
- 추가 확인 없이도 충분히 안전한 초기 속도 제한만으로 시작 가능한지?
- 공개 프로필에서 신뢰 배지에 인증 기여도를 얼마나 반영할지?


## 68. 신규 사용자 첫 거래 활성화(Activation) / 첫 행동 미션 / 전환 보호 계약(초안)
### 68.1 목표
- 회원가입과 거래 가능 상태만으로는 실제 성장을 설명하기 부족하므로, `첫 매물 등록`, `첫 문의 응답`, `첫 예약 확정`, `첫 거래 완료`까지의 **activation funnel**을 별도 관리한다.
- 신규 사용자가 어디에서 막히는지(`등록 전`, `문의는 왔지만 응답 못함`, `예약 직전 이탈`)를 제품/알림/운영이 같은 vocabulary로 해석하게 한다.
- 홈 화면, 온보딩 체크리스트, 알림 리마인드, 초기 제한 해제 정책이 모두 같은 단계 모델을 참조하도록 한다.

### 68.2 activationStage 정의
| activationStage | 의미 | 기본 진입 조건 | 다음 목표 |
|---|---|---|---|
| `joined` | 가입만 완료 | 회원가입/로그인 성공 | 프로필 최소 설정 |
| `profile_ready` | 기본 프로필/서버 설정 완료 | 닉네임, 기본 서버, 최소 소개 또는 기본값 확보 | 첫 매물 등록 또는 첫 검색 진입 |
| `listing_drafted` | 첫 매물 초안 작성 시작 | draft 생성 | 첫 매물 게시 |
| `listing_published` | 첫 매물 공개 완료 | 첫 매물 publish 성공 | 첫 문의/첫 채팅 시작 |
| `first_chat_started` | 첫 거래 대화 진입 | 채팅 생성 또는 첫 메시지 발송 | 첫 예약 제안/응답 |
| `first_reservation_confirmed` | 첫 예약 성립 | Reservation confirmed | 첫 거래 완료 |
| `first_completion_confirmed` | 첫 거래 완료 확정 | completion confirmed/auto_confirmed/resolved_completed | 후기 작성/반복 거래 |
| `retained_trader` | 재방문/반복 거래자 진입 | 일정 기간 내 2회 이상 완료 또는 반복 사용 기준 | 장기 리텐션 |

원칙:
- `tradeEligibilityState=trade_enabled`와 `activationStage`는 별도 축이다.
- 사용자가 거래는 가능해도 `joined` 단계에 오래 머무를 수 있으며, 이 상태는 제품 개선 대상이다.

### 68.3 첫 행동 미션(firstTradeMission) 모델
신규 사용자는 상황에 따라 다른 첫 미션을 본다.

| firstTradeMissionType | 대상 사용자 | 성공 조건 | 대표 진입 surface |
|---|---|---|---|
| `publish_first_listing` | 판매 의도 사용자 | 첫 매물 게시 | 홈, 온보딩, 내 매물 빈 상태 |
| `start_first_chat` | 구매 탐색 위주 사용자 | 첫 채팅 생성 | 검색 결과, 상세, 홈 추천 |
| `respond_to_first_inquiry` | 첫 문의를 받은 신규 판매자 | 일정 시간 내 첫 응답 | 알림함, 내 거래 |
| `confirm_first_reservation` | 채팅은 했지만 거래 전환 못한 사용자 | 첫 예약 확정 | 채팅, 내 거래 |
| `complete_first_trade` | 예약까지 갔지만 완료 경험 없는 사용자 | 첫 완료 확정 | 내 거래, 임박 거래 카드 |
| `write_first_review` | 첫 완료 직후 사용자 | 후기 작성 | 완료 화면, 알림함 |

설계 원칙:
- 한 시점에 사용자에게 강하게 노출되는 미션은 1개만 둔다.
- 같은 사용자가 판매자/구매자 양쪽 행동을 할 수 있으므로, 미션은 role 고정이 아니라 최근 의도/행동 기반으로 선택한다.

### 68.4 activationBlockerCode 후보
신규 사용자 이탈 이유를 구조화한다.

| activationBlockerCode | 의미 | 대표 대응 |
|---|---|---|
| `PROFILE_INCOMPLETE` | 기본 프로필/서버 설정 부족 | 최소 설정 유도 |
| `NO_LISTING_PUBLISHED` | 매물 작성 시작 못함 | 홈 체크리스트/샘플 템플릿 |
| `LISTING_POLICY_BLOCKED` | 첫 매물이 정책 차단됨 | 수정 가이드/사유 노출 |
| `NO_FIRST_RESPONSE` | 문의는 왔지만 응답 안 함 | 즉시 응답 리마인드 |
| `NO_RESERVATION_PROGRESS` | 채팅은 있으나 예약 진입 없음 | 예약 quick action 강조 |
| `NO_COMPLETION_FOLLOWUP` | 거래 후 완료 정리 안 함 | 완료 확인 리마인드 |
| `TRADE_ELIGIBILITY_LIMITED` | 초기 제한/추가 검토로 진행 막힘 | 제한 사유 + 해제 조건 안내 |
| `PUSH_DISABLED` | 중요 알림 누락 가능성 큼 | 거래 맥락형 푸시 opt-in |

원칙:
- blocker는 하나만 대표값으로 내려주되, 내부적으로는 복수 원인을 저장할 수 있다.
- 사용자 노출은 친절한 문구로, 내부 분석/운영은 code 기반으로 관리한다.

### 68.5 홈/온보딩 surface 계약
- 홈 상단 또는 첫 진입 카드에 `currentActivationStage`, `currentFirstTradeMission`, `activationProgressPercent(optional)`를 노출할 수 있어야 한다.
- 신규 사용자의 홈 기본 모듈 우선순위:
  1. 현재 미션 카드
  2. 빠른 서버 선택/검색
  3. 첫 매물 등록 CTA 또는 초안 복구 CTA
  4. 안전/거래 팁 1개
- 내 매물/내 거래 빈 상태는 단순 empty가 아니라 activation 목적 문구를 가져야 한다.

예시 응답 후보:
```json
{
  "activation": {
    "activationStage": "listing_published",
    "firstTradeMissionType": "respond_to_first_inquiry",
    "activationBlockerCode": null,
    "recommendedNextAction": "open_trade_thread",
    "nudgePriority": "high"
  }
}
```

### 68.6 activationNudgePolicy 기본안
| 상황 | 기본 채널 | 기본 강도 | 억제 규칙 |
|---|---|---|---|
| 가입 후 24시간 내 매물/채팅 없음 | 인앱 우선 | low | 같은 날 1회 초과 금지 |
| 첫 문의 수신 후 미응답 | 푸시 + 인앱 | high | 이미 채팅 열람 중이면 푸시 억제 |
| 첫 예약 제안 도착 후 미응답 | 푸시 + 인앱 | high | 예약 만료 직전 중복 푸시 제한 |
| 첫 완료 요청 대기 | 푸시 + 인앱 | high | 완료 확인/분쟁 액션 이후 중지 |
| 첫 후기 미작성 | 인앱 우선 | medium | 최대 2회 |

원칙:
- activation nudges는 마케팅 알림이 아니라 거래 성공 보조 알림으로 취급한다.
- 다만 첫 7일 동안도 과도한 압박이 되지 않도록 일별 총량 cap이 필요하다.

### 68.7 초기 제한 해제와 activation의 연결
- 일부 `tradeEligibilityState` 제한은 activation milestone 달성과 연결될 수 있다.
- 예시:
  - `member_limited_trade` 상태라도 첫 정상 응답/첫 정상 예약 후 rate limit 일부 완화
  - `manual_review_required`는 activation 미션 노출보다 해결 절차 안내를 우선
- 원칙:
  - activation 미션은 사용자를 속여 제한을 숨기면 안 된다.
  - 제한이 blocker인 경우, 미션 카드보다 `왜 막혔는지`와 `어떻게 해제되는지`가 먼저 보여야 한다.

### 68.8 데이터 모델 후보
#### UserActivationSnapshot
| 필드 | 설명 |
|---|---|
| `userId` | 사용자 식별자 |
| `activationStage` | 현재 activation 단계 |
| `firstTradeMissionType` | 현재 대표 미션 |
| `activationBlockerCode(optional)` | 대표 blocker |
| `stageEnteredAt` | 현재 단계 진입 시각 |
| `lastNudgedAt(optional)` | 최근 nudge 발송 시각 |
| `nudgeCount7d` | 최근 7일 nudge 횟수 |
| `firstListingPublishedAt(optional)` | 첫 매물 공개 시각 |
| `firstChatStartedAt(optional)` | 첫 채팅 시작 시각 |
| `firstReservationConfirmedAt(optional)` | 첫 예약 확정 시각 |
| `firstCompletionConfirmedAt(optional)` | 첫 완료 확정 시각 |

원칙:
- activation snapshot은 원천 로그를 대체하지 않고, 홈/알림/CRM성 판단을 위한 projection으로 본다.
- stage 전환은 이벤트 기반 갱신을 우선하고, daily reconciliation job으로 보정 가능해야 한다.

### 68.9 API 후보
- `GET /me/activation`
- `POST /me/activation/dismiss-nudge`
- `GET /home` 또는 `GET /me/home` 응답에 activation card 임베드
- `GET /admin/users/{userId}/activation-history`

응답 필드 후보:
```json
{
  "activationStage": "first_chat_started",
  "firstTradeMissionType": "confirm_first_reservation",
  "activationBlockerCode": "NO_RESERVATION_PROGRESS",
  "availableActions": ["open_chat", "propose_reservation"],
  "policyHints": [
    {
      "code": "FIRST_RESERVATION_IMPROVES_MATCH_SUCCESS",
      "severity": "info",
      "message": "시간과 장소를 먼저 확정하면 거래 완료율이 높아져요."
    }
  ]
}
```

### 68.10 분석/KPI 파생 포인트
필수 이벤트 후보:
- `activation_stage_entered`
- `first_trade_mission_impression`
- `first_trade_mission_click`
- `activation_nudge_sent`
- `activation_blocker_detected`
- `activation_recovered`

핵심 KPI 후보:
- 가입 후 24시간 내 `listing_published` 또는 `first_chat_started` 도달률
- 가입 후 7일 내 `first_reservation_confirmed` 도달률
- 가입 후 14일 내 `first_completion_confirmed` 도달률
- blocker별 회복률(`LISTING_POLICY_BLOCKED`, `NO_FIRST_RESPONSE` 등)
- activation nudge 오픈율과 실제 다음 행동 전환율

### 68.11 화면/운영/QA 시사점
- 화면명세:
  - 홈/내 매물/내 거래 빈 상태 카피가 activation stage별로 달라져야 한다.
  - first mission card는 dismiss 가능하되, 중요한 blocker가 있으면 dismiss보다 해결 CTA를 우선해야 한다.
- 운영정책:
  - 신규 사용자 대량 유입 시 activation nudge가 과도하게 발송되지 않도록 전역 cap/실험 플래그 필요
  - 운영자는 특정 사용자가 어디에서 막혔는지 activation history로 확인 가능해야 한다.
- QA 체크포인트:
  - 첫 매물 공개 직후 `activationStage`가 즉시 `listing_published`로 갱신되는가
  - 첫 문의 수신 후 미응답 시 올바른 blocker와 nudge가 생성되는가
  - 제한 상태 사용자가 misleading한 첫 미션 카드를 보지 않는가

### 68.12 오픈 질문
- MVP에서 activation card를 홈에만 둘지, 내 매물/내 거래 빈 상태에도 공통 컴포넌트로 둘지?
- `retained_trader` 진입 기준을 완료 수 기반으로 할지, 반복 주간 활동 기반으로 할지?
- 신규 사용자의 첫 미션 문구를 역할별(판매자/구매자)로 강하게 분기할지, 행동 기반으로만 둘지?
- activation nudge를 실시간 이벤트 트리거 위주로 갈지, 하루 1회 요약형까지 허용할지?

## 69. 거래 당일 실행 / 도착 확인 / 노쇼 판단 계약(초안)
### 69.1 목표
- 예약이 `confirmed` 된 뒤 실제 거래 현장에서 가장 많이 깨지는 구간인 `출발/도착/지연/노쇼 판단`을 구조화한다.
- 채팅 자유대화에만 의존하지 않고, 거래 당일 필요한 실행 신호를 별도 상태와 액션으로 다뤄 `완료 전환율`과 `분쟁 판단 가능성`을 동시에 높인다.
- 내 거래, 채팅, 알림, 운영 백오피스가 같은 `meeting execution` vocabulary를 사용하도록 통일한다.

### 69.2 핵심 원칙
1. 거래 당일 실행 단계는 예약 확정 이후의 별도 도메인 레이어로 본다.
2. `도착했어요`, `늦어요`, `상대가 안 보여요` 같은 행동은 일반 메시지보다 구조화 액션을 우선 제공한다.
3. 시스템은 노쇼를 자동 확정하지 않으며, `지연`과 `노쇼 의심`을 분리한다.
4. 거래 당일 상태는 분쟁/신뢰 지표에 영향을 줄 수 있으므로 append-only 이력과 최소 증빙 구조가 필요하다.

### 69.3 실행 상태 모델
#### A. 개인별 도착 상태(`arrivalState`)
| 값 | 의미 | 설정 주체 |
|---|---|---|
| `not_started` | 아직 거래 당일 액션 없음 | 기본값 |
| `heading` | 출발/이동 중 | 사용자 |
| `arrived` | 약속 장소 또는 인게임 접선 위치 도착 | 사용자 |
| `delayed` | 늦는다고 알림 | 사용자 |
| `unable_to_arrive` | 오늘 거래 불가 선언 | 사용자 |
| `unknown` | 상태 불명 / 오래된 상태 | 시스템 계산 후보 |

#### B. 거래 실행 상태(`meetingExecutionState`)
| 값 | 의미 | 계산 기준 |
|---|---|---|
| `scheduled` | 예약 확정, 아직 당일 실행 전 | Reservation=confirmed |
| `reminder_window` | 거래 임박, 체크인 유도 구간 | 예약 시각 근접 |
| `one_party_arrived` | 한쪽만 도착 표시 | arrivalState 불일치 |
| `both_arrived` | 양측 도착 표시 | 양측 `arrived` |
| `delay_negotiating` | 한쪽 이상이 지연 알림 | `delayed` 존재 |
| `meeting_at_risk` | 시각 경과 + 상태 불일치/무응답 | grace 진입 |
| `no_show_claimed` | 한쪽이 노쇼 주장 제출 | no-show 액션 발생 |
| `execution_closed` | 완료/취소/분쟁으로 종결 | 후속 상태 연결 |

#### C. 노쇼 판단 상태(`noShowDecisionState`)
| 값 | 의미 |
|---|---|
| `none` | 노쇼 판단 시작 전 |
| `self_reported_unavailable` | 한쪽이 거래 불가를 먼저 알림 |
| `late_but_active` | 지연 상태이나 소통 지속 |
| `claim_pending_counterparty` | 한쪽이 노쇼 주장을 제출했고 상대 응답 대기 |
| `mutual_miss` | 서로 상대 미도착/혼선 주장 |
| `under_review` | 운영 검토 중 |
| `upheld_no_show` | 운영 또는 정책상 노쇼 인정 |
| `rejected_no_show` | 노쇼 인정 불가 |
| `settled_without_fault` | 누구의 명백한 귀책 없이 종료 |

### 69.4 당일 타임라인 기본안
| 시점 | 시스템 동작 | 사용자 주요 CTA |
|---|---|---|
| 예약 시각 2시간 전 | 당일 거래 리마인드 | `장소 다시 확인`, `채팅 열기` |
| 예약 시각 30분 전 | 체크인 유도 | `출발해요`, `늦어요` |
| 예약 시각 도래 | 실행 상태 강조 | `도착했어요`, `지연 알리기`, `채팅 보기` |
| 예약 시각 +15분(가정) | grace 1단계 | `상대가 안 보여요`, `조금 더 기다릴게요` |
| 예약 시각 +30분(가정) | risk 상태 진입 | `노쇼 신고`, `거래 취소`, `시간 다시 잡기` |
| 예약 시각 +24시간 | 장기 미정리 정리 유도 | `완료`, `불발`, `신고` |

원칙:
- 위 시간값은 기본 가정이며 운영/카테고리별 보정 가능하다.
- `+15분`, `+30분`은 실제 노쇼 확정 시간이 아니라 UX/운영 트리거용 grace 기준이다.

### 69.5 사용자 액션 계약
| 액션 코드 후보 | 설명 | 상태 영향 |
|---|---|---|
| `mark_heading` | 지금 이동 중임을 알림 | `arrivalState=heading` |
| `mark_arrived` | 장소 도착 표시 | `arrivalState=arrived` |
| `mark_delayed` | 지연 사유/예상 시간 공유 | `arrivalState=delayed` |
| `mark_unavailable_today` | 오늘 거래 불가 통지 | `arrivalState=unable_to_arrive` |
| `claim_counterparty_no_show` | 상대 노쇼 주장 시작 | `noShowDecisionState=claim_pending_counterparty` |
| `ack_i_am_here` | 상대 노쇼 주장에 반박하며 현재 도착 상태 재확인 | 실행 상태 재평가 |
| `request_reschedule_from_due` | 거래 직전/직후 재일정 제안 | 예약 변경 플로우 연결 |
| `close_due_as_cancelled` | 당일 거래 불발 종료 | 취소/불발 기록 |

원칙:
- 위 액션은 채팅 자유 메시지와 별도로 시스템 카드/버튼으로 제공되어야 한다.
- 채팅 메시지에서 동일 의미 문구를 보내더라도, 핵심 상태 전환은 구조화 액션을 우선 기록하는 것이 좋다.

### 69.6 노쇼 주장 처리 규칙
- 한쪽이 `claim_counterparty_no_show`를 누르면 즉시 상대에게 아래를 요구한다.
  1. 현재 상태 확인(`도착`, `지연`, `오늘 불가`)
  2. 간단 소명 입력(optional)
  3. 재일정/취소 선택
- 상대가 일정 시간(예: 30분, 가정) 내 응답하면 자동으로 `upheld_no_show` 하지 않는다.
- 양측 모두 도착을 주장하거나 장소가 직전 변경된 이력이 있으면 `mutual_miss` 또는 `under_review` 후보로 우선 분류한다.
- 한쪽이 사전에 `unable_to_arrive`를 눌렀다면 기본적으로 `노쇼`보다 `사전 취소`로 해석한다.
- `no_show_claimed`는 신뢰 신호 입력값일 뿐, 즉시 공개 프로필 낙인으로 노출하지 않는다.

### 69.7 최소 증빙 구조(`arrivalEvidenceType`) 후보
| 값 | 설명 | MVP 포함 여부 |
|---|---|---|
| `status_tap` | 구조화 버튼 액션 자체 | 필수 |
| `system_timestamp` | 서버 기록 시각 | 필수 |
| `chat_context` | 직전/직후 채팅 발화 | 필수 |
| `location_ack_state` | 마지막 장소 재확인 여부 | 필수 |
| `attachment_optional` | 현장 스크린샷/사진 등 | 확장 |
| `device_presence_signal` | 앱 foreground/open 여부 | 확장 |

원칙:
- MVP에서는 버튼 액션, 시각, 직전 채팅, 장소 재확인 이력 정도면 충분하다.
- GPS 같은 강한 위치증명은 개인정보/복잡도 때문에 MVP 기본 범위에서 제외한다.

### 69.8 화면 요구사항
#### 채팅 화면
- 예약 시각이 임박하면 입력창 위에 `당일 거래 카드`를 고정 노출한다.
- 카드에는 최소한 `약속 시간`, `장소 요약`, `내 상태`, `상대 상태`, `다음 액션`을 보여줘야 한다.
- `mark_arrived` 이후에는 상대에게 `도착 알림` 시스템 메시지가 노출되어야 한다.

#### 내 거래 상세
- `trade_due_soon`, `meeting_at_risk`, `no_show_claimed` 상태에서 상단 hero 영역으로 승격한다.
- `arrivalState`는 사용자별 칩/배지로 보여주고 마지막 갱신 시각을 함께 표시한다.
- 거래 불발 종료와 재일정 제안 CTA를 완료 CTA와 혼동되지 않게 분리한다.

#### 알림
- `상대가 도착했어요`
- `상대가 늦는다고 알렸어요`
- `거래 시간이 지났어요. 완료/불발을 정리해 주세요`
- `상대가 노쇼를 주장했어요. 현재 상태를 알려 주세요`

### 69.9 API 후보
- `POST /reservations/{reservationId}/arrival-status`
- `POST /reservations/{reservationId}/no-show-claims`
- `POST /reservations/{reservationId}/reschedule-from-due`
- `GET /reservations/{reservationId}/execution-status`

예시 응답 후보:
```json
{
  "reservationId": "res_123",
  "meetingExecutionState": "one_party_arrived",
  "myArrivalState": "arrived",
  "counterpartyArrivalState": "heading",
  "noShowDecisionState": "none",
  "availableActions": ["mark_delayed", "claim_counterparty_no_show", "open_chat"],
  "graceEndsAt": "2026-03-14T21:30:00+09:00"
}
```

### 69.10 데이터 모델 시사점
후보 객체:
- `ReservationExecutionState`
- `ReservationExecutionEvent`
- `NoShowClaim`

후보 필드:
- `reservationId`
- `actorUserId`
- `arrivalState`
- `meetingExecutionState`
- `noShowDecisionState`
- `graceEndsAt`
- `claimedAt`
- `respondedAt`
- `resolvedAt`
- `resolutionReasonCode`

설계 원칙:
- 실행 상태는 Reservation 현재 레코드에 cache 형태로 둘 수 있으나, 사용자 액션 이벤트는 append-only 이력으로 남겨야 한다.
- `NoShowClaim`은 일반 Report와 연결될 수 있지만, 거래 당일 실행 문맥이 강하므로 `reservationId` 중심으로 조회 가능해야 한다.

### 69.11 운영/정책 시사점
- 운영자는 노쇼 판단 시 아래를 한 화면에서 재구성할 수 있어야 한다.
  1. 마지막 확정 장소/시간
  2. 양측 arrival 상태 변경 시각
  3. grace 기간 내 응답 여부
  4. 직전 장소 변경/재확인 여부
  5. 채팅 맥락과 기존 노쇼 이력
- `late_but_active`와 `upheld_no_show`를 분리해, 단순 지각 사용자가 과도하게 제재되지 않도록 해야 한다.
- 반복 `claim_counterparty_no_show` 제출자도 허위/남용 탐지 대상이므로 신고자 품질을 함께 본다.

### 69.12 분석/KPI 파생 포인트
필수 이벤트 후보:
- `trade_due_card_impression`
- `arrival_state_updated`
- `counterparty_arrived_notified`
- `delay_notice_sent`
- `no_show_claim_submitted`
- `no_show_claim_resolved`
- `due_trade_rescheduled`

핵심 KPI 후보:
- 예약 확정 대비 `mark_arrived` 실행률
- `trade_due_soon` 진입 건 중 완료/취소/노쇼 종결 비율
- 노쇼 주장 건 중 `late_but_active`로 반전되는 비율
- 당일 거래 카드 노출 대비 완료 전환율
- 반복 노쇼 인정 사용자 비율 vs 허위 주장 사용자 비율

### 69.13 오픈 질문
- `mark_heading`/`mark_arrived`를 MVP에 모두 넣을지, `도착했어요` 하나만 우선 둘지?
- 거래 방식이 `in_game`일 때와 `offline_pc_bang`일 때 grace 기준을 다르게 둘지?
- 노쇼 주장을 분쟁 객체로 즉시 승격할지, 예약 실행 레이어 안에서 먼저 자체 해소를 시도할지?
- 당일 실행 카드에서 첨부형 증빙(스크린샷 등)을 MVP에 포함할지?



## 70. 약관/정책 동의 버전 / 재동의 / 기능 게이트 계약
### 70.1 목표
- 린클의 거래 기능은 단순 회원가입만으로 열리는 것이 아니라, 서비스 이용약관/개인정보 처리/거래 안전 정책/콘텐츠 정책에 대한 최신 동의 상태와 연결되어야 한다.
- 정책 문구 변경, 금지 품목 범위 조정, 개인정보 처리 범위 변경 시 사용자의 재동의를 어떻게 요구하고 어떤 기능을 일시 제한할지 명확히 정의한다.
- 온보딩, 쓰기 API, 운영 백오피스, 감사 로그, 배치가 같은 `정책 동의 상태` vocabulary를 사용하도록 한다.

### 70.2 정책 문서 유형 정의
| policyDocumentType | 의미 | 기본 성격 | 동의 필요성 |
|---|---|---|---|
| `terms_of_service` | 서비스 이용약관 | 필수 | 회원가입/계속 이용 필수 |
| `privacy_policy` | 개인정보 처리방침 | 필수 | 회원가입/계속 이용 필수 |
| `trade_safety_policy` | 거래 안전 가이드/주의사항 | 준필수 | 거래 쓰기 기능 전 필수 후보 |
| `content_policy` | 거래 가능 품목/금지 표현 정책 | 필수 | 매물 등록/채팅 쓰기 전 필수 후보 |
| `marketing_opt_in` | 마케팅 수신 동의 | 선택 | 기본 OFF 가능 |
| `push_permission_education` | 푸시 권한 안내/프롬프트 기록 | 선택 | 시스템 권한과 별도 |

원칙:
- 법적 효력이 필요한 필수 문서와 UX 안내성 문서를 구분한다.
- `trade_safety_policy`, `content_policy`는 MVP에서 별도 체크박스 분리 여부가 미확정이어도, 서버는 최소한 독립 문서 버전으로 관리할 수 있어야 한다.

### 70.3 정책 문서 버전 객체 후보
| 필드 | 설명 |
|---|---|
| `policyDocumentId` | 문서 식별자 |
| `policyDocumentType` | 문서 유형 |
| `versionLabel` | 사용자 표시 버전(`2026.03`, `v1.0`) |
| `effectiveAt` | 시행 시각 |
| `requiresReConsent` | 재동의 필요 여부 |
| `summaryText` | 변경 요약 |
| `locale` | 언어/지역 |
| `status` | `draft` / `published` / `retired` |
| `contentChecksum` | 변경 추적용 해시 후보 |

원칙:
- 정책 전문 원문과 사용자에게 보여준 요약/체크박스 문구는 함께 보관하는 것이 바람직하다.
- `effectiveAt` 이전에는 공지 가능하지만 기능 게이트는 시행 시각 이후부터 적용한다.

### 70.4 사용자 동의 상태 모델
| 필드 | 설명 |
|---|---|
| `policyAcceptanceId` | 동의 이력 식별자 |
| `userId` | 사용자 |
| `policyDocumentId` | 동의한 문서 버전 |
| `policyDocumentType` | 빠른 조회용 캐시 |
| `acceptanceState` | `accepted` / `declined` / `superseded` / `pending_reacceptance` |
| `acceptedAt` | 동의 시각 |
| `acceptanceMethod` | `signup_checkbox` / `blocking_modal` / `settings_reconfirm` / `admin_recorded` |
| `ipAddressHash(optional)` | 감사/보안 참고 |
| `userAgentSnapshot(optional)` | 기기/세션 참고 |
| `evidenceSnapshotJson` | 당시 체크박스/문구/링크 스냅샷 후보 |

원칙:
- 현재 유효 동의 상태만 덮어쓰지 말고 append-only 이력으로 남겨야 한다.
- 사용자가 이전 버전에 동의했더라도 최신 필수 문서가 `requiresReConsent=true`이면 `pending_reacceptance`로 간주할 수 있어야 한다.

### 70.5 기능 게이트 범위 정의
| policyGateScope | 차단되는 기능 | 예시 |
|---|---|---|
| `auth_only` | 로그인만 허용, 나머지 제한 | 필수 약관 미동의 |
| `read_only` | 읽기만 허용, 쓰기 제한 | 재동의 대기 중 |
| `listing_write_blocked` | 매물 등록/수정 제한 | 콘텐츠 정책 미동의 |
| `chat_write_blocked` | 채팅 발송/예약 제안 제한 | 안전 정책 미동의 |
| `trade_action_blocked` | 예약 확정/완료/후기 제한 | 거래 안전 재동의 필요 |
| `marketing_only` | 마케팅 채널만 영향 | 선택 동의 철회 |

핵심 원칙:
- 필수 문서 미동의는 `회원 존재`와 `거래 가능` 상태를 분리하는 핵심 조건이다.
- 읽기 기능은 가능한 유지하되, 거래 리스크가 있는 쓰기 액션부터 단계적으로 막는 것이 UX/법적 안정성에 유리하다.

### 70.6 재동의 트리거 코드 후보
| reConsentTriggerCode | 설명 | 기본 영향 범위 |
|---|---|---|
| `TOS_MATERIAL_CHANGE` | 이용약관 핵심 조항 변경 | `read_only` 또는 `auth_only` 후보 |
| `PRIVACY_SCOPE_EXPANDED` | 개인정보 처리 범위 확대 | `read_only` 후보 |
| `CONTENT_POLICY_UPDATED` | 금지 품목/표현 정책 중대 변경 | `listing_write_blocked`, `chat_write_blocked` 후보 |
| `TRADE_SAFETY_RULE_UPDATED` | 노쇼/외부연락처/증빙 정책 변경 | `trade_action_blocked` 후보 |
| `LEGAL_NOTICE_REFRESH` | 운영/법적 고지 문구 개정 | 안내만 또는 약한 재확인 |
| `AGE_OR_REGION_REQUIREMENT_CHANGED` | 연령/지역 요건 변경 | `auth_only` 후보 |

원칙:
- 단순 오탈자/비핵심 문구 수정은 재동의 없이 공지로 처리 가능해야 한다.
- `material change` 여부는 운영자가 임의 판단하지 않도록 게시 시 `requiresReConsent`와 `reConsentTriggerCode`를 명시해야 한다.

### 70.7 온보딩 / 재고지 UX 원칙
#### 회원가입 시
- 필수 동의와 선택 동의를 시각적으로 분리한다.
- 거래 안전/콘텐츠 정책이 별도 체크박스가 아니더라도, 최소한 `거래 기능 사용 전 확인 필요` 메시지를 노출한다.
- 체크박스 문구, 버전, 링크 URL은 서버 주도 데이터로 내려받아 앱 하드코딩을 줄인다.

#### 재동의 필요 시
- 로그인 직후 전면 차단보다, 먼저 변경 요약과 영향 기능을 보여주는 blocking modal/full screen을 우선한다.
- 사용자가 재동의를 미루면 허용되는 범위를 명확히 안내한다. 예: `매물은 볼 수 있지만 새 채팅은 보낼 수 없어요.`
- 거래 진행 중인 사용자는 갑작스러운 전면 차단보다 `읽기 유지 + 핵심 후속 액션 허용 여부`를 별도 정책으로 정해야 한다.

### 70.8 진행 중 거래와 정책 게이트의 충돌 규칙
| 상황 | 기본 원칙 |
|---|---|
| 활성 채팅만 있고 예약 전 | 재동의 전까지 새 메시지/예약 제안 제한 가능 |
| 이미 예약 확정됨 | 최소한 일정 확인/취소/안전 고지 확인은 가능해야 함 |
| 완료 확인 대기 중 | 분쟁/확인/신고처럼 사용자 권리 보호 액션은 과도하게 막지 않음 |
| 제재/신고 대응 중 | 정책 미동의와 무관하게 사건 열람/소명은 허용 범위를 별도 정의 |

권장 기본안:
- 재동의가 필요해도 `신고`, `분쟁 소명`, `예약 취소`, `완료 이의 제기` 같은 보호성 액션은 가능한 남겨둔다.
- 반대로 `새 매물 등록`, `새 채팅 시작`, `외부 위험이 큰 거래 쓰기 액션`은 우선 차단한다.

### 70.9 API/응답 계약 시사점
주요 읽기 API와 `GET /me` 응답에는 아래 상태 후보를 포함한다.

```json
{
  "policyAcceptance": {
    "hasBlockingPendingPolicies": true,
    "requiredActions": ["reaccept_content_policy"],
    "gateScopes": ["listing_write_blocked", "chat_write_blocked"],
    "latestPendingDocuments": [
      {
        "policyDocumentType": "content_policy",
        "versionLabel": "2026.03",
        "summaryText": "금지 품목 및 외부 연락처 유도 정책이 업데이트되었어요."
      }
    ]
  }
}
```

후보 API:
- `GET /policies/current`
- `GET /me/policy-acceptances`
- `POST /me/policy-acceptances`
- `GET /me/policy-gates`
- `POST /admin/policies/{policyDocumentId}/publish`

### 70.10 운영/백오피스 요구사항
- 운영자는 사용자별 최신 동의 버전, 동의 시각, 재동의 대기 여부를 조회할 수 있어야 한다.
- 정책 게시 시 아래를 함께 결정해야 한다.
  1. 시행 시각
  2. 재동의 필요 여부
  3. 영향을 받는 gate scope
  4. 공지 대상 범위
  5. 예외 적용 여부(예: 진행 중 거래 보호)
- 운영자가 특정 사용자에게 수동 예외를 부여했다면 `admin_override_policy_gate` 이력과 만료 시각을 남겨야 한다.

### 70.11 데이터/감사 로그 시사점
- 정책 문서 publish, 수정, retire, 재동의 강제는 모두 감사 로그 대상이다.
- `PolicyAcceptance`는 법적/운영 증적 성격이 있으므로 soft delete가 아니라 append-only + supersede 방식이 적합하다.
- 사용자 동의 당시 노출된 요약/체크박스 문구를 보존해야 나중에 “무엇에 동의했는지” 재구성 가능하다.

### 70.12 분석/KPI 파생 포인트
필수 이벤트 후보:
- `policy_gate_impression`
- `policy_summary_open`
- `policy_accept_submit`
- `policy_accept_decline`
- `policy_gate_blocked_action`
- `policy_reaccept_completed`

핵심 KPI 후보:
- 재동의 대상 사용자 대비 완료율
- 정책 게이트 노출 후 이탈률
- gate scope별 거래 쓰기 차단 발생률
- 재동의 후 7일 내 거래 재개율
- 진행 중 거래 사용자에게서 발생한 정책 차단 민원 비율

### 70.13 오픈 질문
- `trade_safety_policy`, `content_policy`를 회원가입 필수 체크로 둘지, 첫 거래 전 동의로 둘지?
- 재동의 필요 시 `read_only`까지 막을지, 쓰기만 막을지?
- 진행 중 예약/분쟁 사용자를 정책 게이트에서 어느 범위까지 예외 처리할지?
- 법적 효력 관점에서 동의 증적에 IP/user-agent를 어디까지 저장할지?


## 71. 부분거래 / 잔여 수량 / 수량 할당 계약
### 71.1 목표
- 하나의 매물 수량이 한 번에 전부 거래되지 않고 일부만 먼저 소진되는 상황을 별도 계약으로 다룬다.
- `부분거래`가 단순 수량 수정으로 덮여서, 완료/후기/검색 노출/분쟁 추적이 깨지지 않게 한다.
- 화면, API, DB, 운영정책이 모두 같은 `allocation` vocabulary를 사용하게 한다.

### 71.2 적용 시나리오
- 판매자가 10개 묶음 매물을 올렸고, 한 구매자에게 4개만 먼저 판매한 경우
- 구매 매물의 희망 수량 20개 중 일부 공급자에게 5개만 먼저 확보한 경우
- 고가 묶음 매물에서 일부 구성품만 거래하고 잔여 구성품으로 계속 거래를 이어가려는 경우
- 실제로는 부분거래였는데 사용자가 매물 전체를 `completed`로 잘못 종결하려는 경우

적용 제외:
- 단일 개체성 아이템(사실상 수량 1, indivisible) 거래는 기본적으로 부분거래 대상이 아니다.
- 묶음 판매로 등록했더라도 정책상 분할 불가(`package_indivisible=true`)면 부분거래를 허용하지 않는다.

### 71.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `listingUnitState` | `single_unit` / `multi_unit_divisible` / `multi_unit_indivisible` | 매물이 수량 분할 가능한지 |
| `allocationState` | `none` / `proposed` / `locked` / `fulfilled_partial` / `fulfilled_full` / `released` / `disputed` | 특정 상대에게 수량이 배정된 상태 |
| `partialCompletionReasonCode` | `buyer_requested_partial` / `seller_stock_short` / `staged_delivery` / `bundle_split_exception` / `dispute_adjusted_quantity` | 왜 부분완료가 발생했는지 |
| `residualHandlingPolicy` | `remain_active` / `remain_reserved` / `manual_reconfirm_required` / `close_listing` | 부분거래 후 잔여 수량을 어떻게 다룰지 |
| `quantityDisclosureMode` | `exact_remaining` / `rounded_band` / `hidden_until_chat` | 남은 수량 공개 방식 |

원칙:
- 부분거래는 `수량 차감 + 완료 이벤트`가 같이 남아야 한다. 단순 `Listing.quantity` overwrite만으로 처리하면 안 된다.
- `Allocation`은 실제 거래 대상 상대에게 잠시 홀드된 수량이며, 최종 완료 전까지 재고와 구분해 추적해야 한다.

### 71.4 기본 정책
- MVP 기본안: `quantity > 1`이고 `listingUnitState=multi_unit_divisible`인 매물만 부분거래 허용 후보로 본다.
- `sell` 매물과 `buy` 매물 모두 부분거래를 허용할 수 있으나, UI copy는 다르게 해석해야 한다.
  - 판매 매물: 남은 수량 관리 중심
  - 구매 매물: 아직 충족되지 않은 목표 수량 관리 중심
- 부분거래 후에도 잔여 수량이 남으면 매물은 기본적으로 종결되지 않는다. 다만 잔여분이 다른 상대와 즉시 거래 불가능하면 `manual_reconfirm_required`를 적용할 수 있다.

### 71.5 불변식 / 정합성 규칙
- `originalQuantity >= fulfilledQuantity + lockedQuantity + availableQuantity` 를 항상 만족해야 한다.
- 동시에 여러 채팅 상대에게 수량을 배정할 수는 있으나, 전체 합이 원래 수량을 초과하면 안 된다.
- `fulfilled_full`이 되면 매물은 `completed` 또는 내부 종결 상태로 전환되어야 한다.
- `fulfilled_partial`만으로는 매물 전체 `completed` 전환이 불가하다.
- 잔여 수량이 0이 아니면 검색/상세/내 매물에서 `부분거래됨`과 `남은 수량`을 일관되게 보여야 한다.

### 71.6 상태 연동 규칙
#### Listing 상태
- 부분거래가 발생해도 잔여 수량이 남고 추가 문의가 가능하면 `available` 유지 가능
- 특정 상대에게 잔여 대부분을 홀드했으면 `reserved` 유지 가능. 이때 `reserved` 의미는 `전체 매물`이 아니라 `잔여 거래 우선권`이어야 한다.
- `pending_trade`는 특정 allocation이 실제 실행 직전일 때만 사용한다.
- 잔여 수량이 0이 되면 그 시점의 마지막 fulfillment가 `fulfilled_full`이 되고 매물은 `completed`로 종결한다.

#### Chat / Reservation 상태
- 채팅방/예약은 매물 전체가 아니라 `allocationQuantity`를 참조할 수 있어야 한다.
- 동일 매물에 여러 채팅방이 존재해도 각 채팅방이 홀드/거래 중인 수량을 분리해서 알아야 한다.
- 한 채팅방의 부분완료는 다른 채팅방을 자동 종결시키지 않는다. 다만 잔여 수량 감소 사실은 read model에 즉시 반영되어야 한다.

### 71.7 Allocation 객체 후보
부분거래를 명시적으로 다루기 위해 `TradeAllocation` 또는 동등 read/write 모델을 검토한다.

| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `allocationId` | 필수 | 수량 배정 식별자 |
| `listingId` | 필수 | 매물 참조 |
| `chatRoomId` | 필수 | 배정이 연결된 채팅방 |
| `reservationId` | 선택 | 예약 연결 |
| `allocatedQuantity` | 필수 | 잠금/제안한 수량 |
| `fulfilledQuantity` | 필수 | 실제 완료된 수량 |
| `allocationState` | 필수 | 상태 |
| `lockedAt` | 선택 | 홀드 시각 |
| `releasedAt` | 선택 | 해제 시각 |
| `fulfilledAt` | 선택 | 실제 부분/전체 이행 시각 |
| `partialCompletionReasonCode` | 선택 | 부분완료 사유 |
| `createdByUserId` | 필수 | 배정 시작 주체 |
| `createdAt` | 필수 | 생성 시각 |

원칙:
- `Reservation`만으로 수량 홀드를 표현하기 어렵다면 `Allocation`을 별도 aggregate 또는 하위 엔티티로 둔다.
- MVP에서 테이블 분리가 과하면, 최소한 fulfillment event log와 per-chat quantity snapshot은 필요하다.

### 71.8 화면 요구사항
#### 매물 상세
- `quantity > 1` 매물은 총 수량과 남은 수량을 분리 표기할 수 있어야 한다.
- 부분거래가 있었으면 `일부 거래 완료` 배지/타임라인 이벤트를 노출한다.
- 수량 공개 방식이 `hidden_until_chat`이면 공개 상세에서는 exact 수량 대신 범위/문구로 대체 가능하다.

#### 채팅 화면
- 예약/완료 카드에 `이번 거래 수량`이 포함되어야 한다.
- 완료 CTA는 `전체 완료`와 `일부만 완료`를 구분할 수 있어야 한다.
- 판매자가 같은 매물의 다른 상대에게 이미 일부 수량을 배정한 경우, 현재 남은 거래 가능 수량을 채팅 헤더/카드에 보여야 한다.

#### 내 매물 / 내 거래
- `원수량`, `잠금 수량`, `완료 수량`, `남은 수량`이 분리 표기되어야 한다.
- 부분거래가 반복된 매물은 `몇 건 완료 / 몇 개 남음` 요약이 필요하다.
- 작성자가 잔여 수량 정책을 정하지 않으면 `manual_reconfirm_required` 배너를 통해 다음 액션을 요구해야 한다.

### 71.9 완료 / 후기 정책 연동
- 후기 작성 자격은 `allocation` 또는 `completion` 단위 실제 거래 상대에게만 열린다.
- 한 매물에서 여러 상대와 부분거래가 일어나면, 각 fulfillment 단위로 후기/신뢰 집계가 생성될 수 있다.
- 단, 동일 상대와 아주 짧은 기간에 연속 부분거래가 여러 번 발생한 경우 리뷰 스팸을 막기 위해 `reviewWindowMergePolicy`를 둘 수 있다.
- 매물 전체가 아직 active여도, 이미 완료된 allocation 상대와의 거래는 완료/후기 흐름으로 종결 가능해야 한다.

### 71.10 API 후보
- `POST /listings/{listingId}/allocations`
- `POST /allocations/{allocationId}/lock`
- `POST /allocations/{allocationId}/release`
- `POST /allocations/{allocationId}/complete`
- `GET /listings/{listingId}/allocations`

부분완료 요청 예시:
```json
{
  "chatRoomId": "chat_123",
  "fulfilledQuantity": 4,
  "residualHandlingPolicy": "remain_active",
  "reasonCode": "buyer_requested_partial",
  "note": "10개 중 4개 먼저 거래 완료"
}
```

응답 예시:
```json
{
  "allocationId": "alloc_123",
  "allocationState": "fulfilled_partial",
  "listing": {
    "listingId": "listing_123",
    "status": "available",
    "originalQuantity": 10,
    "fulfilledQuantity": 4,
    "remainingQuantity": 6
  },
  "completion": {
    "completionId": "comp_123",
    "reviewEligible": true
  }
}
```

### 71.11 검색 / 랭킹 / 알림 규칙
- 검색 목록은 `remainingQuantity > 0`인 매물만 기본 노출 대상으로 본다.
- 부분거래 이력이 있더라도 남은 수량이 충분하고 최근 활동이 있으면 계속 노출 가능하다.
- 잔여 수량이 매우 적어 실질 거래 가능성이 낮으면 랭킹 감점 또는 `few_left` 배지 후보로 처리할 수 있다.
- 찜 사용자/대기 채팅 참여자에게는 `남은 수량 감소`를 무조건 알리지 않고, 가격/상태/재오픈만큼 중요한 이벤트인지 별도 threshold를 둔다.

### 71.12 운영 / 분쟁 / 안티어뷰즈 규칙
- 악용 패턴 후보:
  - 실제 재고보다 큰 수량 등록 후 부분거래를 반복하며 신뢰를 부풀리는 행위
  - 잔여 수량을 계속 조정하며 허위 희소성을 만드는 행위
  - 부분완료를 이용해 다수 후기만 수집하는 행위
- 운영자는 아래를 재구성할 수 있어야 한다.
  1. 원래 등록 수량
  2. 각 allocation lock/release/fulfillment 이력
  3. 잔여 수량 변경 시점
  4. 어떤 상대가 실제로 얼마를 거래했는지
- 분쟁 시 핵심 판단 질문:
  - 합의된 거래 수량은 얼마였는가?
  - 실제 완료로 인정할 수 있는 수량은 얼마인가?
  - 잔여 수량을 매물이 계속 판매 가능한 상태로 유지해야 하는가?

### 71.13 분석 이벤트 후보
- `allocation_created`
- `allocation_locked`
- `allocation_released`
- `allocation_completed_partial`
- `allocation_completed_full`
- `listing_remaining_quantity_changed`

핵심 KPI 후보:
- 부분거래 허용 매물 비율
- 부분거래 매물의 최종 완판율
- 부분거래 이후 후속 거래 성사율
- 부분거래 후 dispute 발생률
- 잔여 수량 1개 이하 구간의 체류 시간

### 71.14 오픈 질문 / 기본 가정
- 가정: MVP에서는 exact remaining quantity를 기본 공개하고, 희소성 과장 방지를 위해 수동 수정 이력은 강하게 남긴다.
- 가정: `buy` 매물의 부분충족도 동일 모델을 재사용하되, UI 문구는 `남은 필요 수량` 중심으로 분리한다.
- 결정 필요: allocation을 별도 write-model table로 둘지, Reservation/TradeCompletion 확장으로 흡수할지.
- 결정 필요: 동일 상대와 연속 부분거래를 후기 1건으로 합칠 기준 시간(window)을 둘지.
- 결정 필요: 부분거래 후 남은 수량이 너무 적을 때 자동 종결/자동 경고 기준을 둘지.

## 65. 노쇼 claim / 증빙 / 분쟁·신고 연결 계약
### 65.1 목표
- 예약 당일 실행 단계에서 발생하는 `안 왔어요`, `늦는다고만 하고 안 왔어요`, `장소를 바꿔서 못 만났어요` 같은 사건을 단순 채팅 감정싸움이 아니라 **구조화된 claim 객체**로 다룬다.
- `arrival`, `reschedule`, `meetingExecution`, `dispute`, `report`, `restriction`이 각자 다른 기준으로 같은 사건을 해석하지 않도록 연결 규칙을 명시한다.
- 노쇼는 곧바로 유죄 확정이 아니라, 사용자 self-report → 상대 응답 → 증빙 제출 → 운영 판정 또는 비판정 종료의 흐름으로 본다.

### 65.2 적용 범위
노쇼 claim은 아래 조건을 모두 만족할 때 생성 가능하다.
- 활성 또는 직전 종료된 예약이 존재한다.
- 마지막으로 유효한 약속 시각은 `원예약`이 아니라 `lastAcceptedReschedule` 기준으로 계산한다.
- 현재 시각이 약속 시각 이전이면 생성 불가. 단, 상대가 명시적으로 `cannot_attend`류 실행 액션을 보낸 경우 사전 불발 claim으로 전환 가능하다.
- 이미 `completed` 최종 확정된 거래에는 일반 no-show claim을 새로 열 수 없고, 필요 시 별도 dispute/report로 연결한다.

적용 제외:
- 아직 `reservationStatus=proposed`인 상태의 미응답은 노쇼가 아니라 예약 응답 실패로 본다.
- 채팅만 있었고 예약이 없던 케이스는 no-show claim 대상이 아니라 일반 신고 대상이다.
- 단순 지각으로 grace period 내 도착 확인이 된 경우는 노쇼 성립 대상이 아니다.

### 65.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `noShowClaimState` | `draft` / `submitted` / `countered` / `evidence_requested` / `under_review` / `resolved` / `withdrawn` / `expired` | claim 자체의 상태 |
| `noShowClaimRole` | `claimant` / `respondent` | 사건 내 상대적 역할 |
| `noShowClaimReasonCode` | `counterparty_absent` / `counterparty_unreachable` / `unagreed_location_change` / `late_beyond_grace` / `claimed_arrival_without_evidence` | 제기 사유 |
| `arrivalEvidenceDisposition` | `not_provided` / `provided_self_asserted` / `provided_geo_weak` / `provided_chat_corroborated` / `provided_staff_verified` | 도착 증빙 강도 |
| `noShowResolutionType` | `claim_upheld` / `claim_rejected` / `mutual_fault` / `insufficient_evidence` / `converted_to_reschedule` / `converted_to_general_dispute` | 최종 판정 타입 |
| `claimLinkagePolicy` | `report_optional` / `report_required_on_escalation` / `restriction_signal_only` | report/restriction 연결 방식 |

원칙:
- `noShowClaim`은 신고의 하위 메모가 아니라 예약 실행 실패에 특화된 별도 사건 객체로 다룬다.
- 사용자 노출 문구는 `노쇼 신고`, `도착 못함 신고`, `상대 미도착 이슈` 정도로 단순화하고, 내부 상태는 별도로 유지한다.

### 65.4 사건 생성 기준
#### 생성 가능 시점
- 약속 시각 도래 후 grace period가 시작되면 양측 모두 claim CTA를 볼 수 있다.
- grace period 경과 후에는 `no_show_claim` CTA 우선순위를 높인다.
- 한쪽이 `arrived`를 남겼고 상대가 `not_arrived` 상태이면 claimant 우선 claim 생성이 가능하다.

#### 생성 제한
- 동일 예약에 활성 no-show claim은 최대 1건만 허용한다.
- 이미 `under_review`인 claim이 있으면 새 claim 생성 대신 기존 사건으로 합류한다.
- claim 생성 후 상대는 별도 새 claim을 여는 대신 `counter_statement` 또는 `counter_claim_reason`으로 응답한다.
- claim 생성만으로 자동 제재/자동 후기 비노출은 하지 않는다.

### 65.5 원예약 / 재일정 / grace period 해석 규칙
- 마지막으로 `accepted`된 재일정만 노쇼 판단 기준 시각/장소를 덮어쓴다.
- `requested` 또는 `counter_proposed` 상태의 미합의 재일정은 기준 시각을 바꾸지 않는다.
- `arrivalState=arrived`를 남긴 시각이 있어도 기준 장소와 불일치하면 자동 인정하지 않고 claim 판단에 참고자료로만 사용한다.
- `grace period`는 기본적으로 마지막 유효 예약 기준으로 계산하며, `accepted same-day reschedule`이 있으면 재설정한다.
- 노쇼 claim 도중 운영자 또는 양측 합의로 재일정이 성립하면 `noShowResolutionType=converted_to_reschedule`로 종료 가능하다.

### 65.6 사용자 플로우
#### Claimant 플로우
1. `거래 상세/채팅/당일 실행 카드`에서 `상대가 오지 않았어요` 선택
2. reason code 선택
3. 도착 증빙/채팅 맥락/메모 제출
4. 사건 생성 후 상대 응답 대기
5. 필요 시 추가 소명 제출 또는 일반 신고로 확장

#### Respondent 플로우
1. no-show claim 알림 수신
2. `도착했다`, `지각 중이었다`, `장소 합의가 달랐다`, `실제론 거래 취소 합의했다` 중 응답 선택
3. 증빙/메모 제출
4. 사건이 종료되거나 운영 검토로 넘어갈 때까지 진행 상태 확인

#### 시스템/운영 플로우
- 증빙이 거의 없고 상호 진술만 충돌하는 경우 즉시 강한 제재보다 `restriction_signal_only`로 누적한다.
- 반복 패턴, accepted reschedule 무시, 다수 상대 동일 패턴이 결합될 때에만 운영 우선순위를 높인다.

### 65.7 화면 요구사항
#### 채팅/당일 실행 카드
- `도착했어요`, `5분 늦어요`, `재일정 요청`, `노쇼 신고` CTA는 서로 구분되어야 한다.
- no-show claim 시작 시 필수 노출:
  - 마지막 유효 약속 시간/장소 요약
  - grace period 기준
  - accepted reschedule 반영 여부
  - 상대 응답 가능 기한
  - 허위 claim도 운영 제재 대상이라는 경고

#### 내 거래 상세
- 상태 배지 후보:
  - `상대 응답 대기`
  - `증빙 보강 필요`
  - `운영 검토 중`
  - `노쇼 이슈 종료`
- 같은 화면에서 일반 신고와 no-show claim을 구분해 보여야 한다.

#### 운영 백오피스
- 한 화면에서 아래를 타임라인으로 재구성 가능해야 한다.
  1. 예약 원본
  2. 마지막 accepted reschedule
  3. arrival signal 양측 이력
  4. no-show claim 제출/응답/추가 소명
  5. linked report / dispute / restriction

### 65.8 데이터 모델 후보
#### NoShowClaim
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `noShowClaimId` | 필수 | 사건 식별자 |
| `reservationId` | 필수 | 기준 예약 |
| `listingId` | 필수 | 조회 편의용 |
| `chatRoomId` | 필수 | 채팅 연동 |
| `claimantUserId` | 필수 | 최초 제기자 |
| `respondentUserId` | 필수 | 상대방 |
| `lastAcceptedRescheduleId` | 선택 | 마지막 유효 재일정 |
| `claimState` | 필수 | 사건 상태 |
| `claimReasonCode` | 필수 | 주장 사유 |
| `claimSubmittedAt` | 필수 | 접수 시각 |
| `responseDueAt` | 선택 | 상대 응답 기한 |
| `resolvedAt` | 선택 | 종료 시각 |
| `resolutionType` | 선택 | 판정/종료 결과 |
| `linkedDisputeId` | 선택 | 분쟁 연결 |
| `linkedReportId` | 선택 | 신고 연결 |
| `riskSignalScore` | 선택 | 내부 운영 신호 |

#### NoShowClaimStatement
- `statementId`
- `noShowClaimId`
- `submittedByUserId`
- `statementType`: `initial_claim` / `counter_statement` / `additional_context` / `admin_requested_response`
- `bodyText`
- `reasonCode(optional)`
- `submittedAt`

#### NoShowEvidence
- `evidenceId`
- `noShowClaimId`
- `submittedByUserId`
- `evidenceType`: `chat_message_ref` / `arrival_checkin` / `image` / `location_note` / `system_timestamp` / `reservation_snapshot`
- `evidenceDisposition`
- `storageObjectId(optional)`
- `createdAt`

### 65.9 API 후보
#### 사용자용
- `POST /reservations/{reservationId}/no-show-claims`
- `GET /no-show-claims/{noShowClaimId}`
- `POST /no-show-claims/{noShowClaimId}/statements`
- `POST /no-show-claims/{noShowClaimId}/withdraw`
- `POST /no-show-claims/{noShowClaimId}/escalate-report`

#### 관리자용
- `GET /admin/no-show-claims`
- `GET /admin/no-show-claims/{noShowClaimId}`
- `POST /admin/no-show-claims/{noShowClaimId}/resolve`
- `POST /admin/no-show-claims/{noShowClaimId}/link-restriction`

Response 필드 후보:
```json
{
  "noShowClaimId": "nsc_123",
  "claimState": "submitted",
  "reservationId": "res_123",
  "lastEffectiveScheduledAt": "2026-03-14T21:30:00+09:00",
  "lastEffectiveMeetingSummary": "기란 마을 창고 앞",
  "linkedObjects": {
    "disputeId": null,
    "reportId": null
  },
  "availableActions": ["submit_counter_statement", "escalate_report"]
}
```

### 65.10 Report / Dispute / Restriction 연결 규칙
- 기본적으로 no-show claim은 **독립 사건**으로 시작하고, 모든 건을 즉시 일반 신고로 만들지 않는다.
- 아래 경우 `linkedReportId` 생성 또는 연결을 권장한다.
  1. 욕설/협박/사기 유도 등 추가 안전 이슈가 동반됨
  2. 동일 사용자의 반복 no-show 패턴이 누적됨
  3. 운영 개입/제재가 실제로 필요함
- 아래 경우 `linkedDisputeId` 연결을 우선 고려한다.
  1. no-show가 단순 미도착이 아니라 `장소 합의 해석 충돌`, `거래 완료 주장 충돌`로 번짐
  2. accepted reschedule 해석 자체가 분쟁 핵심인 경우
- restriction 연결 원칙:
  - 단건 claim만으로 자동 제재하지 않는다.
  - `claim_upheld` 반복, `mutual_fault` 반복, respondent 무응답 누적, 고위험 신규계정 패턴이 결합될 때 restriction signal로 승격한다.

### 65.11 후기 / 신뢰도 / 분석 반영 규칙
- unresolved no-show claim이 열린 동안 해당 거래의 후기 작성은 기본적으로 잠시 보류하는 보수안을 둔다.
- `claim_upheld` 시 respondent의 공개 신뢰도에 즉시 큰 패널티를 주기보다, 내부 no-show metric과 반복 패턴 신호를 먼저 누적한다.
- `claim_rejected` 또는 `insufficient_evidence`는 공개 낙인 없이 내부 품질 지표로만 남긴다.
- 분석 이벤트 후보:
  - `no_show_claim_submitted`
  - `no_show_claim_countered`
  - `no_show_claim_resolved`
  - `no_show_claim_escalated_to_report`
  - `no_show_claim_converted_to_reschedule`

### 65.12 운영 판정 원칙
- 운영은 아래 질문 순서로 판단한다.
  1. 마지막 상호 합의된 시간/장소가 무엇이었는가?
  2. grace period 안에 어떤 arrival signal이 있었는가?
  3. accepted reschedule이 있었는가, 아니면 일방 주장뿐이었는가?
  4. 외부 증빙 없이도 플랫폼 내부 로그만으로 어느 정도 일관된가?
  5. 동일 사용자의 반복 패턴이 있는가?
- 증빙이 약하면 `insufficient_evidence` 또는 `restriction_signal_only`로 종료할 수 있어야 한다.
- 운영 결과는 사건 종료와 계정 제재를 분리 기록해야 한다.

### 65.13 오픈 질문 / 기본 가정
- 가정: no-show claim 상대 응답 기한은 claim 생성 후 12시간 또는 다음날 오전 10시 중 더 이른 값으로 둔다. 실제 값은 운영 테스트로 보정 필요.
- 가정: unresolved no-show claim이 있으면 후기 공개를 일시 보류한다.
- 결정 필요: claimant/ respondent 모두에게 self-claim score를 얼마나 노출할지 여부.
- 결정 필요: geo/location 기반 도착 증빙을 MVP에서 아예 제외할지, 단순 체크인 timestamp만 사용할지.
- 결정 필요: no-show claim과 일반 `Report(reason=no_show)`를 최종적으로 하나의 운영 큐로 합칠지, 별도 큐로 유지할지.



## 65. 노쇼 claim adjudication / case linkage / 후기·제한 반영 계약
### 65.1 목표
- `no-show` 판단이 단순 신고 문구가 아니라 **사건 단위(case unit)** 로 수렴되도록 정의한다.
- 같은 약속에 대해 생성된 claim, counter-claim, report, dispute, restriction, review gating, analytics가 서로 다른 식별자/상태를 쓰지 않게 한다.
- 운영자가 `누가 먼저 claim 했는가`보다 `마지막 상호합의 예약`, `도착/지각/재일정`, `증빙`, `반복 패턴`, `최종 판정`을 기준으로 판단할 수 있게 한다.

### 65.2 사건 단위(case unit) 원칙
- 노쇼 사건의 canonical 단위는 `reservationId` 기반 `NoShowCase`다.
- 하나의 `reservationId`에는 동시에 1개의 활성 `NoShowCase`만 존재할 수 있다.
- 한쪽이 먼저 claim을 생성하면 `NoShowCase`가 열리고, 상대방의 반대 주장/소명은 동일 case에 append된다.
- 별도 `Report`를 추가 제출하더라도 운영 내부에서는 `linkedNoShowCaseId`를 통해 같은 사건으로 병합 조회할 수 있어야 한다.

### 65.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `caseLinkPolicy` | `case_first` / `report_first` / `dual_linked` | claim과 report/dispute 연결 우선 원칙 |
| `claimAdjudicationState` | `case_opened` / `waiting_counterparty` / `evidence_collecting` / `under_review` / `resolved` / `closed_without_action` | 노쇼 사건 운영 상태 |
| `caseResolutionType` | `no_show_upheld` / `mutual_fault` / `reschedule_accepted_so_no_no_show` / `insufficient_evidence` / `claim_abuse_detected` | 최종 판정 유형 |
| `reviewPublicationGate` | `allow_normal` / `hold_until_resolution` / `deny_for_this_trade` / `allow_with_note` | 후기 공개/작성 게이트 |
| `restrictionEscalationTier` | `none` / `warning_candidate` / `trust_penalty_candidate` / `temporary_limit_candidate` / `manual_restriction_review` | 판정 이후 제재 후보 강도 |
| `caseOutcomeSeverity` | `low` / `medium` / `high` / `critical_pattern` | 운영 우선순위 및 재범 분석 강도 |

기본 원칙:
- MVP 기본안은 `case_first`를 채택한다. 즉, 노쇼 관련 `Report`와 `Dispute`는 사건을 새로 만들지 않고 기존 `NoShowCase`에 연결되는 보조 entry로 본다.
- 후기/제한/신뢰도 반영은 개별 claim 기준이 아니라 **case resolution 기준** 으로만 수행한다.

### 65.4 객체 관계 원칙
| 객체 | 관계 | 비고 |
|---|---|---|
| `NoShowCase` | 사건 루트 aggregate | `reservationId`, `listingId`, `chatRoomId`, 양측 사용자 보유 |
| `NoShowClaim` | case 하위 append-only claim | 최초 claim, counter-claim, 추가 설명을 포함 |
| `ArrivalEvidence` | case 또는 claim에 연결 | 도착 인증/위치/메시지/시각 증빙 |
| `Report` | 선택 연결(`linkedNoShowCaseId`) | 욕설/괴롭힘 등 별도 정책 이슈가 있으면 유지 |
| `Dispute` | 선택 연결(`linkedNoShowCaseId`) | 완료분쟁과 결합된 경우만 사용 |
| `RestrictionReview` | case resolution 결과로 파생 | 자동 제재가 아니라 검토 후보 생성 |
| `Review` | case resolution 전/후 게이트 적용 | 공개 보류 가능 |

### 65.5 claim 생성/병합 규칙
- 한쪽이 `no_show` claim을 만들면 case가 없을 때만 새 `NoShowCase`를 생성한다.
- 이미 열린 case가 있으면 새 claim은 `counter_claim` 또는 `additional_statement`로 append한다.
- 같은 사용자가 동일 예약에 대해 짧은 시간 내 동일 취지 claim을 반복 제출하면 새 case/새 report를 만들지 않고 기존 claim에 merge하거나 dedup한다.
- `accepted` 재일정이 존재하면 case는 항상 마지막 `accepted`된 시각/장소를 기준 약속으로 삼는다.
- `requested` 상태의 미수락 재일정은 참고 증빙일 뿐 기준 약속을 덮어쓰지 않는다.

### 65.6 운영 adjudication 상태 머신
#### 기본 흐름
1. `case_opened`
2. 상대방 통지 및 `waiting_counterparty`
3. 증빙 제출 또는 자동 수집 메타데이터 정리 → `evidence_collecting`
4. 운영자 배정/검토 → `under_review`
5. 최종 판정 → `resolved`
6. 신고 철회/무효/중복 케이스 병합 시 → `closed_without_action`

#### 상태 해석 규칙
- 상대방이 응답하지 않아도 case는 자동 `no_show_upheld`로 곧바로 판정하지 않는다. 다만 evidence 부족이 아니고 반복 패턴이 강하면 우선순위를 높일 수 있다.
- 양측 모두 상대를 노쇼로 주장해도 case는 2개로 쪼개지지 않고 하나의 사건 안에서 `mutual_fault` 또는 `insufficient_evidence` 가능성을 검토한다.
- 완료 요청/완료 분쟁이 이미 열린 거래에서 노쇼 claim이 뒤늦게 들어오면, 별도 사실관계 충돌 여부를 보고 기존 `Dispute`에 연결하거나 `NoShowCase`를 독립 유지하되 `linkedDisputeId`를 저장한다.

### 65.7 판정 기준 우선순위
운영자 판단의 기본 우선순위는 아래와 같다.
1. 마지막으로 상호 수락된 예약 시각/장소/방식
2. 예약 직전/직후의 도착 신호(`arrived`, `running_late`, `location_ack`)
3. accepted reschedule 존재 여부와 시각
4. 채팅 타임라인상 응답 공백, 도착 예고, 취소 의사
5. 제출 증빙(스크린샷, 위치 설명, 도착 메시지, 시스템 timestamp)
6. 동일 사용자 반복 패턴(노쇼 누적, same-day reschedule 남용, 직전 취소 습관)

원칙:
- 외부 캡처보다 플랫폼 내부 timestamp와 상태 로그를 우선한다.
- `running_late`는 자동 면책 사유가 아니며, 상대방의 명시적 수락 또는 accepted reschedule이 없으면 기준 시간을 연장하지 않는다.
- 당일 직전 장소 대폭 변경은 노쇼 판정의 단독 근거는 아니지만 `mutual_fault` 또는 `insufficient_evidence` 판단에 영향을 줄 수 있다.

### 65.8 case resolution과 후속 반영 규칙
| `caseResolutionType` | 후기 처리 | 신뢰/제한 처리 | 사용자 안내 기본안 |
|---|---|---|---|
| `no_show_upheld` | `hold_until_resolution` 후 결과 반영 | 패턴 따라 `warning_candidate` 이상 | 상대의 노쇼 주장이 일부 인정됨 |
| `mutual_fault` | `allow_with_note` 또는 보류 | 즉시 강한 제재 지양 | 양측 조율 실패로 판단됨 |
| `reschedule_accepted_so_no_no_show` | `allow_normal` | 제재 후보 없음 | 재일정 합의가 인정되어 노쇼 불성립 |
| `insufficient_evidence` | `allow_normal` 또는 선택적 보류 해제 | 누적 반영 최소화 | 증빙 부족으로 단정 불가 |
| `claim_abuse_detected` | 허위/악용 작성자 후기에 제한 가능 | `temporary_limit_candidate` 또는 수동 검토 | 허위 claim 악용 정황이 확인됨 |

### 65.9 후기/리뷰 publication gating 규칙
- `NoShowCase`가 `resolved`되기 전까지 해당 거래의 후기 공개는 기본적으로 `hold_until_resolution`을 우선안으로 둔다.
- 단, 노쇼 claim이 거래 완료 이후 장시간 뒤늦게 접수되었고 이미 review가 공개된 경우에는 소급 비공개보다 `reviewUnderCaseReview=true` 배지 형태를 Post-MVP 후보로 남긴다. MVP에서는 공개 후 소급 변경을 최소화한다.
- `claim_abuse_detected`가 판정되면 허위 claim 작성자의 후기 작성권을 해당 거래 건에서 제한하거나 운영 검토 후 비노출 처리할 수 있다.
- `mutual_fault`는 한쪽만 일방적 가해자로 보지 않으므로 리뷰 공개는 허용하되, 추천/비추천 집계 반영 강도를 완화할지 여부는 운영 정책 확정이 필요하다.

### 65.10 restriction / trust 반영 규칙
- 단건 `no_show_upheld`만으로 자동 계정 정지를 실행하지 않는다.
- 대신 아래 누적 패턴에서 `RestrictionReview`를 생성한다.
  - 최근 30일 내 upheld no-show 2회 이상
  - same-day accepted reschedule 후 upheld no-show 2회 이상
  - 첫 거래 cohort에서 upheld no-show + 외부 연락처 유도 + 신고 결합
  - `claim_abuse_detected` 1회 이상
- `trustPenaltyCandidate`는 공개 배지 하향보다 내부 랭킹/주의 신호에 우선 반영한다.
- 노쇼 판정 이력은 공개 프로필에 사건 단위 수치를 직접 노출하지 않고, 운영/신뢰 엔진 입력으로만 우선 사용한다.

### 65.11 화면 요구사항
#### 내 거래 / 채팅
- 사용자는 현재 사건 상태를 `상대 응답 대기`, `운영 검토 중`, `판정 완료` 수준으로 이해할 수 있어야 한다.
- 같은 거래에 report/dispute/no-show case가 모두 얽혀 있어도 UI는 사건 단위를 1개의 `문제 해결 카드`로 묶어 보여주는 것이 바람직하다.
- 후기가 보류된 경우 `노쇼 관련 검토가 끝나면 후기 작성/공개가 열립니다` 같은 문구를 노출한다.

#### 운영 백오피스
- case 상세에는 아래가 한 타임라인에 보여야 한다.
  1. 원예약
  2. 마지막 accepted reschedule
  3. arrival/no-show 이벤트
  4. 양측 claim/counter-claim
  5. 연결된 report/dispute/restriction review
  6. 최종 resolution 및 후속 액션

### 65.12 API / 데이터 모델 후보
#### NoShowCase
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `noShowCaseId` | 필수 | 사건 식별자 |
| `reservationId` | 필수 | canonical 사건 기준 |
| `listingId` | 필수 | 조회 편의용 |
| `chatRoomId` | 필수 | 대화 스레드 연결 |
| `sellerUserId` | 필수 | 참여자 A |
| `buyerUserId` | 필수 | 참여자 B |
| `claimAdjudicationState` | 필수 | 운영 상태 |
| `caseResolutionType` | 선택 | 최종 판정 |
| `linkedDisputeId` | 선택 | 완료 분쟁 연결 |
| `linkedPrimaryReportId` | 선택 | 기본 report 연결 |
| `reviewPublicationGate` | 필수 | 후기 게이트 |
| `restrictionEscalationTier` | 필수 | 제한 검토 강도 |
| `openedAt` | 필수 | 사건 시작 시각 |
| `resolvedAt` | 선택 | 판정 완료 시각 |
| `resolvedByAdminId` | 선택 | 처리 운영자 |

#### NoShowClaim
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `noShowClaimId` | 필수 | claim 식별자 |
| `noShowCaseId` | 필수 | 사건 참조 |
| `claimantUserId` | 필수 | 주장자 |
| `claimType` | 필수 | `initial_claim` / `counter_claim` / `additional_statement` |
| `claimedAgainstUserId` | 선택 | 상대 지목 |
| `claimReasonCode` | 선택 | 무응답, 장소불일치, late-no-ack 등 |
| `statementText` | 선택 | 설명 텍스트 |
| `submittedAt` | 필수 | 제출 시각 |
| `dedupKey` | 선택 | 중복 병합용 |

#### API 후보
- `POST /reservations/{reservationId}/no-show-cases`
- `GET /reservations/{reservationId}/no-show-case`
- `POST /no-show-cases/{noShowCaseId}/claims`
- `POST /admin/no-show-cases/{noShowCaseId}/resolve`
- `POST /admin/no-show-cases/{noShowCaseId}/link-report`

### 65.13 알림 규칙
- case 개설 시 상대방에게 즉시 알림을 보내되, 푸시 본문은 비난 표현 없이 `약속 관련 이슈 제기가 접수되었습니다` 수준으로 제한한다.
- `waiting_counterparty` 상태에서 응답 기한 알림 1회를 허용한다.
- `resolved` 시 양측에 결과 요약 + 다음 행동(후기 가능 여부, 이의제기 가능 여부, 제한 여부)을 안내한다.
- 운영자가 linked report/dispute를 추가해도 사용자 알림은 사건 카드 기준으로 묶어 중복 전송을 피한다.

### 65.14 분석 이벤트 / KPI 연결
- `no_show_case_opened`
- `no_show_case_counter_claimed`
- `no_show_case_resolved`
- `no_show_case_linked_report`
- `no_show_case_linked_dispute`
- `review_publication_held_by_no_show_case`
- `restriction_review_created_from_no_show_case`

핵심 관찰 지표:
- 예약 대비 no-show case 생성률
- case 중 양측 counter-claim 비율
- resolution type 분포
- upheld no-show 이후 30일 재범률
- case 발생 거래의 후기 공개 지연시간
- no-show case와 first-trade churn 상관관계

### 65.15 evidence sufficiency / adjudication SLA / 사용자 결과 노출 계약
#### 65.15.1 목표
- 노쇼 사건이 열렸을 때 운영자가 `증빙이 부족하지만 느낌상 의심됨` 수준으로 오래 끌지 않고, 어떤 입력이 있어야 판정 가능하고 언제 `insufficient_evidence`로 닫아야 하는지 공통 기준을 둔다.
- 내 거래/채팅/후기/제한/운영 큐가 같은 판정 결과를 서로 다른 문구로 번역하지 않도록 사건 결과 노출 계약을 정리한다.
- MVP 기준에서 자동 판정은 최소화하되, **운영 SLA와 증빙 충족도**는 구조화해 이후 정책 엔진/큐 자동 우선순위화로 바로 파생 가능하게 한다.

#### 65.15.2 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `evidenceSufficiencyTier` | `none` / `self_statement_only` / `timeline_supported` / `mutual_conflict_with_logs` / `decision_ready` | 사건이 현재 어느 수준의 판정 입력을 갖췄는지 |
| `adjudicationSlaTier` | `urgent_same_day` / `standard_trade` / `batched_low_signal` | 운영 처리 목표 시간대 |
| `userOutcomeVisibility` | `status_only` / `status_plus_next_step` / `status_plus_reason_summary` | 사용자에게 얼마나 자세한 결과를 보여줄지 |
| `appealEligibilityState` | `not_available` / `available_once` / `closed_after_appeal` | 판정 후 이의제기 가능 상태 |
| `caseClosureReasonCode` | `resolved_with_decision` / `duplicate_merged` / `insufficient_evidence_timeout` / `withdrawn_by_claimant` / `converted_to_other_case_type` | 사건 종료 직접 사유 |

기본 원칙:
- `NoShowCase`는 항상 `evidenceSufficiencyTier`를 가진다. 운영자는 사건을 읽을 때 resolution보다 먼저 현재 증빙 충족도를 본다.
- `insufficient_evidence`는 운영 소극성의 euphemism이 아니라, **정해진 입력/시간 기준을 만족하지 못한 사건 종료 결과**로 정의한다.
- 사용자에게는 내부 증빙 점수나 rule hit를 노출하지 않고, 필요한 다음 행동 중심으로만 보여준다.

#### 65.15.3 증빙 충족도 tier 기준
| tier | 최소 조건 | 운영 해석 | 기본 다음 액션 |
|---|---|---|---|
| `none` | claim만 있고 설명/로그 포인트 없음 | 사실상 판정 불가 | 상대 응답 대기 또는 추가 설명 요청 |
| `self_statement_only` | 한쪽 서술만 있음, 시스템 로그는 약함 | 단독 주장 단계 | counter-claim 또는 증빙 요청 |
| `timeline_supported` | 예약/읽음/메시지/arrival log 중 일부가 주장과 연결됨 | 최소한 사건 뼈대는 있음 | 운영 triage 가능 |
| `mutual_conflict_with_logs` | 양측 주장이 충돌하고 시스템 로그도 일부 상반됨 | 사람이 판정해야 하는 케이스 | 수동 review 우선 |
| `decision_ready` | 마지막 accepted reservation, arrival/no-show 로그, 주요 응답 공백, 보조 증빙이 충분 | 판정 가능 | resolve 진행 |

추가 해석 규칙:
- `decision_ready`는 외부 첨부가 많다는 뜻이 아니라 **상호수락된 약속 + 핵심 시스템 로그 + 최소 설명**이 정렬되었다는 뜻이다.
- 외부 스크린샷만 많고 플랫폼 기준 타임라인이 비어 있으면 `decision_ready`로 올리지 않는다.
- same-day 재일정이 얽힌 사건은 `accepted reschedule` 확인 전에는 최대 `mutual_conflict_with_logs`까지만 자동 승격 가능하다.

#### 65.15.4 adjudication SLA 기본안
| `adjudicationSlaTier` | 대표 케이스 | 1차 triage 목표 | 최종 판정 목표 |
|---|---|---|---|
| `urgent_same_day` | 거래 당일, 양측 활발, 재일정/도착/노쇼가 직결된 케이스 | 2시간 이내 | 24시간 이내 |
| `standard_trade` | 일반 노쇼 claim, 증빙 존재 | 12시간 이내 | 72시간 이내 |
| `batched_low_signal` | 오래 지난 claim, 증빙 약함, 중복/경미 | 24시간 이내 | 5영업일 이내 |

우선순위 승격 조건:
- 첫 거래 cohort의 no-show case
- same-day accepted reschedule 이후 no-show claim
- 외부 연락처 유도/괴롭힘 report가 함께 연결된 case
- 양측 모두 현재 action-needed 상태로 후기/제한/완료 플로우가 막힌 case

강등 조건:
- duplicate merge 예정 case
- claimant가 스스로 철회한 case
- 핵심 로그가 거의 없고 예약 시각과 claim 접수 시각이 과도하게 벌어진 case

#### 65.15.5 추가 증빙 요청 / 자동 종료 규칙
- 운영자는 `timeline_supported` 이하 사건에만 추가 증빙 요청을 기본으로 보낸다. 이미 `decision_ready`인 사건은 불필요한 자료 수집으로 지연시키지 않는다.
- 추가 증빙 요청은 최대 1회 기본안을 둔다. 같은 당사자에게 무한 소명 요청을 반복하지 않는다.
- 아래 중 하나면 `insufficient_evidence_timeout` 종료 후보가 된다.
  - 상대방 무응답 + claimant 추가자료 없음 + 시스템 로그도 약함
  - 양측 모두 서술만 반복하고 새 증빙이 없음
  - claim 접수 시점이 거래 시점과 과도하게 멀어 핵심 로그만으로도 단정 불가
- 자동 종료가 아니라 운영 종료를 원칙으로 하되, 워크플로우 상 `closureSuggested=true` 같은 큐 힌트는 허용한다.

#### 65.15.6 사용자 결과 노출 계약
| 사건 결과 | `userOutcomeVisibility` 기본안 | 내 거래/채팅 문구 방향 | 후속 CTA |
|---|---|---|---|
| `no_show_upheld` | `status_plus_reason_summary` | 약속 불이행 주장이 일부 인정됨 | 후기/이의제기/고객센터 |
| `mutual_fault` | `status_plus_reason_summary` | 조율 실패로 판단됨 | 후기/다시 거래하지 않기 |
| `reschedule_accepted_so_no_no_show` | `status_plus_next_step` | 재일정 합의가 인정되어 노쇼로 보지 않음 | 거래 재개/후기 대기 |
| `insufficient_evidence` | `status_plus_next_step` | 증빙 부족으로 단정할 수 없음 | 필요 시 1회 이의제기 |
| `claim_abuse_detected` | `status_plus_reason_summary` | 허위 또는 악용성 claim 정황이 확인됨 | 제한 안내/이의제기 |

원칙:
- 사용자 문구는 상대방 비난을 확정형으로 쓰지 않고 `주장이 일부 인정됨`, `증빙 부족`, `조율 실패` 같은 운영 표현을 사용한다.
- 제한이 실제 실행되지 않았다면 `제한될 수 있습니다` 같은 예고 문구를 남발하지 않는다.
- 후기 게이트가 열리거나 닫히는 순간은 사건 결과 카드와 함께 설명돼야 한다.

#### 65.15.7 후기 / 제한 / 이의제기 연결 규칙
- `reviewPublicationGate=hold_until_resolution` 상태였다면, 사건 `resolvedAt` 기준으로 후기 CTA를 즉시 재계산한다.
- `restrictionEscalationTier`는 결과 저장 시점에 함께 freeze되어야 하며, 이후 실제 제재 생성 여부와 분리 저장한다. 즉 `candidate`와 `applied restriction`을 혼동하지 않는다.
- `appealEligibilityState=available_once`는 사건 판정 직후에만 열리고, appeal 제출 후 `closed_after_appeal`로 전환한다.
- appeal은 원 case를 뒤집는 새 claim이 아니라 `case resolution review request`로 저장하는 것이 바람직하다.

#### 65.15.8 백오피스 / analytics 파생 기준
운영 큐는 최소 아래 정렬 키를 가져야 한다.
1. `adjudicationSlaTier`
2. `evidenceSufficiencyTier`
3. same-day 여부
4. linked report/dispute 존재 여부
5. first-trade cohort 여부

분석 이벤트 후보 추가:
- `no_show_case_evidence_tier_changed`
- `no_show_case_additional_evidence_requested`
- `no_show_case_closed_insufficient_evidence`
- `no_show_case_appeal_submitted`
- `no_show_case_outcome_viewed`

핵심 관찰 지표 추가:
- evidence tier별 최종 resolution 분포
- 추가 증빙 요청 후 resolution lead time
- `insufficient_evidence_timeout` 비율
- no-show case 판정 후 appeal 제출률
- outcome별 후기 작성 재개율

### 65.16 오픈 질문 / 기본 가정
- 가정: 노쇼 사건의 상대방 기본 응답 기한은 case 개설 후 24시간으로 둔다.
- 가정: same-day 고위험 패턴은 일반 report 큐보다 no-show case 전용 큐에서 먼저 분류한다.
- 가정: 추가 증빙 요청은 기본 1회, 예외적으로 고위험 사건만 2회까지 허용한다.
- 결정 필요: `mutual_fault` 사례를 공개 후기 집계에 완전히 제외할지, 중립 사례로 둘지.
- 결정 필요: `insufficient_evidence` 종료 사건의 후기 공개를 즉시 열지, 짧은 cooling period 후 열지.
- 결정 필요: no-show case를 `Dispute`의 subtype으로 흡수할지, 별도 aggregate로 유지할지. 본 PRD는 별도 aggregate 유지안을 우선 지지한다.
- 결정 필요: claim abuse가 confirmed일 때 해당 사용자의 향후 no-show claim 생성 자체에 쿨다운을 둘지.

### 65.17 노쇼 판정 이후 appeal / 재심 / 신뢰도 재계산 계약
#### 65.17.1 목표
- no-show 사건이 1차 판정된 뒤에도 사용자가 예측 가능한 범위에서 1회 재심을 요청할 수 있게 하되, 운영 큐가 무한 재오픈되는 것을 막는다.
- 원판정이 유지/수정/철회될 때 후기 게이트, restriction candidate, trust aggregate, analytics가 서로 다른 사건 버전을 보지 않도록 `revision-first` 원칙을 둔다.
- 운영자 수동 판정과 배치성 신뢰 재집계가 같은 `case revision` vocabulary를 사용하게 한다.

#### 65.17.2 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `appealState` | `not_available` / `available_once` / `submitted` / `screening` / `under_reconsideration` / `resolved_keep` / `resolved_revise` / `resolved_void` / `closed_no_new_input` | 사건 판정 이후 appeal 자체의 상태 |
| `appealEligibilityReasonCode` | `new_evidence_required` / `time_window_expired` / `already_appealed` / `case_not_final` / `claim_abuse_locked` | appeal 가능/불가 사유 |
| `resolutionRevisionType` | `keep_original` / `soft_reword` / `reverse_fault_direction` / `downgrade_to_insufficient_evidence` / `void_case_outcome` | 재심 결과가 원판정을 어떻게 바꾸는지 |
| `recomputationScope` | `review_gate_only` / `restriction_candidate_only` / `trust_aggregate_only` / `all_downstream_effects` | 재심 결과가 재계산해야 하는 파급 범위 |
| `trustReconciliationMode` | `append_delta` / `full_case_replay` / `manual_override_with_audit` | 신뢰/집계 재반영 방식 |
| `caseRevisionNo` | 정수 증가 | 같은 사건의 판정 revision 번호 |

원칙:
- 사건 최종 결과는 단일 current snapshot으로 조회되더라도, 내부적으로는 `caseRevisionNo` 증가형 이력을 보존해야 한다.
- appeal은 새 no-show claim이 아니라 기존 case에 대한 `resolution review request`로 취급한다.

#### 65.17.3 appeal 생성 자격 / 입력 요건
- `appealEligibilityState=available_once`인 final case만 appeal 생성 가능하다.
- 기본안으로 appeal 제출 가능 기한은 원판정 결과 통지 후 7일 이내로 둔다.
- appeal은 아래 중 최소 1개를 만족해야 한다.
  1. 원판정 이후 새로 확보된 증빙 첨부
  2. 원판정에서 누락/오독된 플랫폼 로그 지점 명시
  3. 결과 문구 또는 제한 연결이 사실과 다르게 적용됐다는 구체 사유
- 단순 불만 표출, 동일 문장 반복, 상대 비난만 있는 appeal은 `closed_no_new_input` 후보다.
- `claim_abuse_detected`가 확정된 사용자는 동일 사건 appeal은 허용하되, 다른 신규 no-show claim 생성은 별도 cooldown 후보로 본다.

#### 65.17.4 appeal 상태 머신
- 사건 final resolution 후 `available_once`
- 사용자 제출 시 `submitted`
- 운영 1차 입력 검토 후 `screening`
- 재심 대상으로 채택되면 `under_reconsideration`
- 결과가 원판정 유지면 `resolved_keep`
- 결과 문구/귀책/후속조치가 수정되면 `resolved_revise`
- 사건 자체를 무효화해야 하면 `resolved_void`
- 새 입력이 없거나 기한이 지나면 `closed_no_new_input`

해석 규칙:
- `screening` 단계에서는 원판정 효력을 유지한다. 즉 appeal 제출만으로 후기 게이트나 제한을 자동 롤백하지 않는다.
- `under_reconsideration` 진입 시에만 `downstreamEffectsFrozen=true` 같은 내부 힌트를 둘 수 있다. 단 이미 집행된 제재는 별도 revision 결과가 나오기 전까지 유지 가능하다.
- 사건당 appeal은 기본 1회만 허용한다. 재심 결과에 대한 재재심은 신규 법적/정책 경로가 아닌 이상 열지 않는다.

#### 65.17.5 판정 수정 유형과 downstream effect 재계산 원칙
| `resolutionRevisionType` | 의미 | 기본 `recomputationScope` | 설명 |
|---|---|---|---|
| `keep_original` | 원판정 유지 | `review_gate_only` | 문구/근거만 보강될 수 있으나 결과는 유지 |
| `soft_reword` | 귀책 문구 완화/정밀화 | `review_gate_only` | 사용자 노출 copy만 바뀌고 제한/집계는 유지 가능 |
| `reverse_fault_direction` | 귀책 방향 반전 | `all_downstream_effects` | 후기 게이트, restriction candidate, trust 집계 전부 재평가 |
| `downgrade_to_insufficient_evidence` | 단정 판정을 증빙부족으로 완화 | `all_downstream_effects` | 강한 귀책 효과 제거, 후기/제한 정책 재계산 |
| `void_case_outcome` | 사건 자체 무효 | `all_downstream_effects` | case 기반 파생 효과를 원칙적으로 철회 |

원칙:
- `recomputationScope=all_downstream_effects`이면 review publication, restriction candidate, trust counter, analytics outcome bucket까지 같은 revision 번호로 갱신해야 한다.
- 이미 사용자에게 노출된 결과 카드가 있더라도 최신 revision이 생기면 `superseded_by_revision` 표기를 남길 수 있어야 한다.

#### 65.17.6 후기 게이트 / 제한 / 신뢰도 연결 규칙
- `reviewPublicationGate`는 사건 current revision 기준으로만 열린다. appeal 접수만으로 자동 open되지 않는다.
- 원판정에서 `restrictionEscalationTier=candidate_high`였더라도, 재심 결과가 `downgrade_to_insufficient_evidence` 또는 `void_case_outcome`이면 candidate를 `superseded` 처리해야 한다.
- `trustAggregate`는 사건별 immutable event 합이 아니라 `current effective resolution` 기준 read model을 별도 계산하는 것이 안전하다.
- 즉 trust 지표 반영은 아래 2층 구조를 권장한다.
  1. append-only case revision log
  2. current-effective trust projection
- 사용자 프로필의 no-show 관련 배지/주의 신호는 `current-effective trust projection`만 본다. 과거 뒤집힌 판정을 계속 노출하지 않는다.

#### 65.17.7 배치 / projection 재집계 기본안
- `trustReconciliationMode=append_delta`는 문구 보정이나 제한 없는 경미 수정에만 사용한다.
- 귀책 반전, 증빙부족 강등, 사건 무효화는 `full_case_replay`를 기본안으로 둔다.
- `full_case_replay`는 최소 아래 projection에 대해 idempotent 재계산 가능해야 한다.
  - `UserTrustProjection`
  - `ReviewPublicationProjection`
  - `RestrictionCandidateProjection`
  - `NoShowCaseSummaryProjection`
  - outcome 기반 analytics materialization
- 수동 예외 조정이 필요하면 `manual_override_with_audit`를 쓰되, override reason과 approver를 필수 저장한다.

#### 65.17.8 사용자 결과 화면 / 카피 원칙
- 사건 상세에는 아래 3층 정보를 구분해 보여야 한다.
  1. 현재 유효한 결과
  2. appeal 가능 여부/남은 기한
  3. revision이 있었다면 `이전 판정이 수정됨` 이력
- 사용자 문구 예시:
  - `재심 요청이 접수되었어요. 현재 판정 효력은 검토 전까지 유지됩니다.`
  - `재심 결과 일부 판정이 수정되었어요.`
  - `새로운 증빙이 부족해 재심이 접수되지 않았어요.`
- 상대방 실명형 비난 문구는 피하고, 사건 기준 언어를 유지한다.

#### 65.17.9 API / DB 파생 기준
API 후보:
- `POST /no-show-cases/{caseId}/appeals`
- `GET /no-show-cases/{caseId}/appeal`
- `POST /admin/no-show-cases/{caseId}/appeals/{appealId}/screen`
- `POST /admin/no-show-cases/{caseId}/appeals/{appealId}/resolve`
- `POST /internal/projections/trust/recompute` (internal job/admin-trigger candidate)

DB 후보 엔티티:
- `NoShowCaseAppeal`
- `NoShowCaseResolutionRevision`
- `TrustReconciliationJob`

`NoShowCaseAppeal` 최소 필드 후보:
- `appealId`
- `caseId`
- `submittedByUserId`
- `appealState`
- `appealReasonCode`
- `newEvidenceSummaryText`
- `submittedAt`
- `screenedAt`
- `resolvedAt`
- `resolvedByAdminId`
- `linkedRevisionNo`

#### 65.17.10 analytics / 운영 관찰 지표
분석 이벤트 후보 추가:
- `no_show_case_appeal_opened`
- `no_show_case_appeal_submitted`
- `no_show_case_appeal_screened_out`
- `no_show_case_resolution_revised`
- `trust_projection_reconciled_from_case_revision`

핵심 관찰 지표:
- 사건 결과별 appeal 제출률
- appeal screening reject 비율
- `resolved_revise` 비율
- revision 발생 후 trust projection 반영 lead time
- 판정 뒤집힘이 후기 작성률/제한 집행률에 미치는 영향

#### 65.17.11 오픈 질문 / 기본 가정
- 가정: appeal은 사건당 1회, 결과 통지 후 7일 이내만 허용한다.
- 가정: 귀책 반전 또는 사건 무효화는 `full_case_replay`를 기본으로 한다.
- 결정 필요: `soft_reword`만 발생한 경우 사용자에게 push를 다시 보낼지, 인앱 결과 카드 갱신만 할지.
- 결정 필요: 이미 공개된 후기 중 사건 판정 뒤집힘으로 인해 임시 hold가 필요한 범위를 어디까지 둘지.
- 결정 필요: appeal screen-out 결과를 moderator 단독 처리로 둘지, senior moderator 승인까지 요구할지.



### 65.18 거래 당일 실행 카드 CTA family / 상태카피 / action gating 계약
#### 65.18.1 목표
- 거래 당일 사용자가 가장 자주 누르는 `도착`, `늦음`, `재일정`, `노쇼`, `완료` 액션이 서로 다른 의미로 해석되도록 CTA family를 명시적으로 분리한다.
- 채팅 상단 카드, 내 거래 카드, 푸시 딥링크, 운영 타임라인이 같은 action vocabulary를 사용해 어떤 버튼이 왜 보였는지 재구성 가능하게 한다.
- 단순한 버튼 노출을 넘어서 `언제 disabled 되는지`, `어떤 경고 문구를 붙이는지`, `어떤 클릭이 실제 상태 변경인지`를 규정해 화면명세·API·QA 파생 기준을 만든다.

#### 65.18.2 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `dayOfTradeCardMode` | `pre_checkin` / `checkin_window` / `counterparty_arrived` / `running_late_risk` / `reschedule_pending` / `mutual_presence` / `result_needed` / `claim_open` / `closed` | 당일 실행 카드의 상위 표현 모드 |
| `executionCTAType` | `check_in_arrived` / `send_running_late` / `open_reschedule` / `view_meeting_details` / `nudge_counterparty` / `mark_trade_completed` / `start_no_show_claim` / `submit_counter_evidence` / `ack_outcome` | 카드에서 노출 가능한 실행 액션 종류 |
| `ctaAvailabilityReasonCode` | `within_window` / `already_arrived` / `awaiting_counterparty_response` / `reschedule_required_first` / `claim_already_open` / `reservation_not_final` / `policy_locked` / `completed_or_cancelled` | CTA 노출/비노출/disabled 사유 |
| `checkinConfidenceTier` | `self_report_only` / `self_report_plus_message` / `mutual_presence_signal` / `evidence_attached` | 도착/지연 신호의 강도 |
| `executionCopyTone` | `neutral_prompt` / `urgent_action_needed` / `risk_warning` / `post_outcome_summary` | 사용자-facing copy tone 계층 |

원칙:
- `CTA label`과 `executionCTAType`를 분리한다. 같은 `늦어요` 라벨이라도 단순 메시지와 실제 재일정 요청은 다른 타입이어야 한다.
- 카드 모드는 read model 계산값이며, 실제 상태 원천은 reservation / arrival / reschedule / no-show case 객체다.

#### 65.18.3 CTA family 정의
| family | 대표 CTA | 의미 | 상태 변경 여부 |
|---|---|---|---|
| arrival | `도착했어요` | 내가 약속 장소/상태에 도달했음을 self-report | `arrivalState` 변경 가능 |
| delay | `5분 늦어요` | 짧은 지연 고지, 원예약 유지 전제 | 기본은 상태 변경 아님, signal/event 저장 |
| reschedule | `시간 다시 제안` / `장소 바꾸기` | 원예약을 명시적으로 바꾸는 협상 시작 | `RescheduleRequest` 생성 |
| presence-nudge | `상대에게 다시 알리기` | 상대 확인/응답 촉구 | 상태 변경 아님, nudge event 생성 |
| outcome | `거래 완료` / `노쇼 신고` | 결과 판정 흐름 시작 | completion / claim 객체 생성 |
| evidence | `도착 증빙 추가` / `소명 보내기` | 이미 열린 사건에 증빙 추가 | case/evidence 객체 변경 |

핵심 구분:
- `delay`는 grace 범위 안의 짧은 지연 신호다. 원예약을 바꾸지 않는다.
- `reschedule`은 새 시간/장소에 대한 명시적 동의 흐름이다. accepted 되기 전까지 원예약 기준이 유지된다.
- `presence-nudge`는 상대 행동을 촉구할 뿐, 도착/노쇼 판단 자체를 바꾸지 않는다.

#### 65.18.4 카드 모드별 기본 CTA 매핑
| `dayOfTradeCardMode` | 주 CTA | 보조 CTA | 금지/주의 |
|---|---|---|---|
| `pre_checkin` | `view_meeting_details` | `open_reschedule` | 너무 이른 시점의 no-show CTA 금지 |
| `checkin_window` | `check_in_arrived` | `send_running_late`, `view_meeting_details` | 완료/노쇼 CTA는 예약 시각 전 과노출 금지 |
| `counterparty_arrived` | `check_in_arrived` | `nudge_counterparty`, `view_meeting_details` | 상대 도착만으로 자동 no-show 유도 금지 |
| `running_late_risk` | `send_running_late` | `open_reschedule`, `view_meeting_details` | 지연 고지와 재일정 CTA를 혼동시키지 않음 |
| `reschedule_pending` | `open_reschedule` 응답형(`수락/거절/대안`) | `view_meeting_details` | no-show CTA는 상대 응답 만료 전 후순위 |
| `mutual_presence` | `mark_trade_completed` | `view_meeting_details` | no-show CTA 기본 숨김 |
| `result_needed` | `mark_trade_completed` | `start_no_show_claim`, `submit_counter_evidence` | 둘 다 1차 CTA로 동시에 강조하지 않음 |
| `claim_open` | `submit_counter_evidence` 또는 `ack_outcome` | `view_meeting_details` | 중복 claim 생성 금지 |
| `closed` | `ack_outcome` | 기록 보기 | 실행형 CTA 비노출 |

#### 65.18.5 CTA 노출 우선순위 / gating 규칙
- 동일 카드에서 primary CTA는 1개만 강조한다.
- 우선순위 기본안:
  1. 상대 응답이 필요한 action (`reschedule_pending`, `claim_open`)
  2. 시간 민감한 self-report (`check_in_arrived`, `send_running_late`)
  3. 결과 종결 (`mark_trade_completed`, `start_no_show_claim`)
  4. 보조 확인 (`view_meeting_details`, `nudge_counterparty`)
- `mark_trade_completed`와 `start_no_show_claim`가 동시에 가능하더라도, copy와 배치로 의미를 분리해야 한다. 기본안은 `거래 완료`를 primary, `노쇼/문제 신고`를 destructive secondary로 둔다.
- `claim_already_open`, `policy_locked`, `completed_or_cancelled` 상태에서는 outcome CTA를 disabled 또는 숨김 처리한다.
- accepted 되지 않은 재일정이 열려 있으면 `ctaAvailabilityReasonCode=reschedule_required_first`로 no-show CTA를 후순위 또는 제한 처리할 수 있다. 단 원예약 cutoff 초과 후에는 다시 claim CTA를 활성화할 수 있어야 한다.

#### 65.18.6 카피 원칙
- CTA 라벨은 짧게, 보조 문구는 해석 가능하게 쓴다.
- 권장 카피 방향:
  - `도착했어요` + `상대에게 도착 알림이 전송돼요`
  - `5분 늦어요` + `원래 약속은 유지돼요`
  - `시간 다시 제안` + `상대 수락 전까지 원래 약속이 유지돼요`
  - `노쇼 신고` + `상호 기록과 증빙을 바탕으로 검토돼요`
- 위험 문구 금지:
  - 상대 귀책을 단정하는 감정형 문구
  - `신고하면 바로 제재됩니다` 같은 과장 표현
  - delay CTA에 재일정 효력이 있는 것처럼 보이는 문구
- `executionCopyTone=risk_warning`는 노쇼 claim, cutoff 초과, repeated reschedule 고위험 문맥에서만 사용한다.

#### 65.18.7 채팅 / 내 거래 / 알림 surface별 표현 원칙
##### 채팅 상단 카드
- 가장 풍부한 CTA surface로 간주한다.
- primary 1개 + secondary 최대 2개 + 보조 링크 1개를 기본 상한으로 둔다.
- 시스템 메시지 타임라인과 카드 CTA가 같은 사건 키를 공유해야 한다.

##### 내 거래 카드
- 요약 우선. `mode badge + primary CTA + 약속 요약` 3요소를 기본으로 한다.
- secondary CTA는 `상세 보기`로 접고, destructive action은 상세 진입 후 노출하는 보수안을 우선한다.

##### 알림 / 푸시
- 푸시에서는 action intent만 전달하고 다중 CTA 문구를 싣지 않는다.
- 예: `약속 시간이 다가와요. 도착 여부를 알려주세요.`
- 알림 오픈 후 landing screen은 해당 `executionCTAType`이 primary인 카드 상태로 복원되어야 한다.

#### 65.18.8 API / read model / DB 파생 기준
read model 후보 필드:
- `dayOfTradeCardMode`
- `primaryExecutionCTA`
- `secondaryExecutionCTAs[]`
- `ctaAvailabilityReasonCodes[]`
- `checkinConfidenceTier`
- `cardHeadline`
- `cardSupportText`

API 응답 예시:
```json
{
  "tradeThreadId": "tt_123",
  "dayOfTradeCard": {
    "mode": "checkin_window",
    "primaryCTA": "check_in_arrived",
    "secondaryCTAs": ["send_running_late", "view_meeting_details"],
    "ctaAvailabilityReasonCodes": ["within_window"],
    "headline": "약속 시간이 가까워졌어요",
    "supportText": "도착했거나 조금 늦는다면 상대에게 알려주세요."
  }
}
```

DB/이벤트 후보:
- `TradeThreadProjection.day_of_trade_card_mode`
- `TradeExecutionSignal` (`signalType=arrived|running_late|nudge_sent`)
- `ExecutionCTAClickEvent` (analytics/event log)

원칙:
- CTA 클릭 로그와 실제 상태 변경 이벤트를 분리 저장한다. 클릭했지만 취소/실패한 케이스를 분석해야 하기 때문이다.
- `send_running_late`는 기본적으로 lightweight signal row로 저장하고, accepted reschedule과 혼동되지 않게 별도 타입을 둔다.

#### 65.18.9 analytics / QA / 운영 시사점
분석 이벤트 후보:
- `day_of_trade_card_impression`
- `execution_cta_clicked`
- `execution_signal_sent`
- `execution_cta_blocked`
- `day_of_trade_mode_transitioned`

핵심 관찰 지표:
- mode별 CTA 클릭률
- `send_running_late` 이후 accepted reschedule 전환률
- `check_in_arrived` self-report 후 mutual completion 전환률
- CTA blocked 비율과 주된 `ctaAvailabilityReasonCode`
- claim CTA 노출 대비 실제 claim 생성률

QA 파생 체크포인트:
- 같은 상태에서도 surface별로 primary CTA가 일관적인가
- accepted reschedule 대기 중 no-show CTA가 과도하게 노출되지 않는가
- `도착했어요`를 눌러도 재일정/완료 상태로 오인되지 않는가
- closed 사건에서 실행형 CTA가 다시 살아나지 않는가

#### 65.18.10 오픈 질문 / 기본 가정
- 가정: 당일 실행 카드의 primary CTA는 항상 1개만 강하게 노출한다.
- 가정: `5분 늦어요`는 기본적으로 재일정이 아니라 lightweight delay signal이다.
- 결정 필요: `send_running_late` 누적 횟수를 신뢰/노쇼 판정의 약한 신호로만 볼지, repeated pattern에서 stronger signal로 승격할지.
- 결정 필요: `counterparty_arrived` 모드에서 상대 self-report를 얼마나 강하게 신뢰할지(`checkinConfidenceTier`)와 증빙 업로드 CTA를 MVP에 포함할지.


## 64. 노쇼 사건(No-show Case) canonical 객체 / 화면 표시 / API·DB 파생 기준
### 64.1 목표
- `노쇼 claim`, `counter-claim`, `증빙`, `운영 판정`, `appeal`이 서로 다른 임시 객체로 흩어지지 않도록 **단일 canonical case**를 정의한다.
- 채팅/내 거래/운영 큐/API/DB/analytics가 모두 같은 `noShowCaseId`를 기준으로 사건을 해석하도록 한다.
- 사용자가 보는 것은 단순한 "노쇼 신고"이더라도, 내부적으로는 사건 생성/병합/판정/재심까지 추적 가능한 구조를 가져야 한다.

### 64.2 canonical 사건 단위 정의
- **NoShowCase**: 특정 거래 실행 약속(`reservationId` 또는 동등한 day-of-trade execution context)에 대해 발생한 노쇼/도착 불일치 사건의 집계 루트
- 사건 생성 기준 기본안:
  - 동일 `reservationId`에 대해 동시에 열릴 수 있는 활성 no-show 사건은 최대 1개
  - 한쪽이 먼저 claim을 열고, 상대가 반박하면 **별도 사건 생성이 아니라 같은 사건의 counter-position**으로 흡수
  - accepted reschedule 이후 이전 예약 컨텍스트에 대한 늦은 claim은 원칙적으로 새 active case를 만들지 않고 `invalid_after_reschedule` 또는 참고 메모로만 남김

### 64.3 핵심 vocabulary
| 용어 | 의미 |
|---|---|
| `noShowCaseId` | 노쇼 사건 canonical 식별자 |
| `casePartyRole` | `claimant` / `respondent` / `counter-claimant` / `moderator` |
| `caseOutcomeType` | `claim_upheld`, `claim_rejected`, `shared_fault`, `insufficient_evidence`, `void_after_reschedule`, `withdrawn` |
| `caseTimelineEventType` | `case_opened`, `counter_position_added`, `evidence_added`, `sla_tier_assigned`, `moderator_requested_more_info`, `decision_recorded`, `appeal_opened`, `decision_revised`, `case_closed` |
| `evidenceBundleScope` | `claimant_only`, `respondent_only`, `shared_case`, `moderator_only` |
| `caseVisibilityLevel` | `user_summary`, `participant_detail`, `moderator_full` |

원칙:
- `Report`, `Dispute`, `Restriction`, `Review publication gate`는 사건을 참조할 수 있으나, **노쇼 사실 판단의 원본은 항상 NoShowCase**다.
- 사용자에게는 `노쇼 신고`, `상대가 반박함`, `운영 검토 중` 같은 쉬운 문구로 보여주되 내부 키는 위 vocabulary를 유지한다.

### 64.4 NoShowCase 필수 필드 후보
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `noShowCaseId` | 필수 | 사건 식별자 |
| `reservationId` | 필수 | 기준 예약 |
| `listingId` | 필수 | 연결 매물 |
| `chatRoomId` | 필수 | 실제 협의 채팅 |
| `executionWindowStartedAt` | 필수 | 당일 실행 window 시작 |
| `executionWindowEndedAt` | 선택 | grace 포함 window 종료 |
| `claimOpenedByUserId` | 필수 | 최초 claim 생성자 |
| `respondentUserId` | 필수 | 상대방 |
| `caseState` | 필수 | 내부 사건 진행 상태 |
| `caseOutcomeType` | 선택 | 최종 판정 결과 |
| `primaryClaimReasonCode` | 필수 | 최초 주장 사유 |
| `claimOpenedAt` | 필수 | 최초 접수 시각 |
| `counterPositionAt` | 선택 | 상대 반박 제출 시각 |
| `decisionAt` | 선택 | 1차 판정 시각 |
| `appealOpenedAt` | 선택 | 이의제기 시각 |
| `closedAt` | 선택 | 완전 종료 시각 |
| `caseLinkPolicy` | 필수 | report/dispute/review linkage 해석 |
| `adjudicationSlaTier` | 필수 | 처리 우선순위/SLA |
| `reviewGateResult` | 선택 | 후기 공개 게이트 결과 |
| `restrictionEscalationTier` | 선택 | 제한 반영 레벨 |

### 64.5 사건 상태 머신 기본안
| `caseState` | 의미 | 다음 가능 액션 |
|---|---|---|
| `opened` | 한쪽 claim 생성, 상대 답변 대기 | counter-position, evidence add, withdraw |
| `countered` | 상대 반박 또는 상충 진술 존재 | evidence add, moderator review |
| `awaiting_more_info` | 운영이 추가 자료 요청 | evidence add, timeout close |
| `under_review` | 운영 판정 중 | decision record |
| `decided` | 1차 판정 완료 | appeal open, side-effect apply |
| `appeal_pending` | 재심 입력 수신 | review appeal |
| `revised` | 1차 판정 수정/철회 완료 | close |
| `closed` | 더 이상 변경 없음 | read-only |

상태 규칙:
- `opened`에서 상대가 같은 예약에 대해 "나도 상대가 노쇼"를 주장하더라도 별도 사건을 만들지 않고 `countered`로 전환한다.
- `decided` 이전에는 후기 공개, trust 반영, restriction 반영이 `pending_case_result` 게이트에 묶일 수 있어야 한다.
- `closed` 이후에는 새 evidence를 직접 append하지 않고, 필요하면 `appeal_pending`을 통해 revision 흐름으로만 재개한다.

### 64.6 사건 하위 객체 후보
#### 64.6.1 NoShowCaseParty
- 사건 당사자별 진술/행동 상태를 보관
- 필드 후보:
  - `noShowCasePartyId`
  - `noShowCaseId`
  - `userId`
  - `partyRole`
  - `positionType` (`claimed_arrived`, `claimed_counterparty_absent`, `claimed_late_but_present`, `claimed_rescheduled`, `claimed_cancelled_in_time`)
  - `statementSummaryText`
  - `submittedAt`
  - `lastUpdatedAt`

#### 64.6.2 NoShowCaseEvidence
- 사건 증빙 메타데이터를 보관
- 필드 후보:
  - `evidenceId`
  - `noShowCaseId`
  - `submittedByUserId`
  - `evidenceType`
  - `evidenceBundleScope`
  - `capturedAt`
  - `submittedAt`
  - `visibilityLevel`
  - `verificationDisposition`
  - `storageObjectId`

#### 64.6.3 NoShowCaseDecision
- 판정 단위와 side-effect snapshot 보관
- 필드 후보:
  - `decisionId`
  - `noShowCaseId`
  - `decisionVersion`
  - `outcomeType`
  - `rationaleSummary`
  - `reviewGateResult`
  - `restrictionEscalationTier`
  - `trustReconciliationMode`
  - `decidedByAdminId`
  - `decidedAt`

#### 64.6.4 NoShowCaseTimelineEvent
- 사건 연표 projection 원천
- 필드 후보:
  - `timelineEventId`
  - `noShowCaseId`
  - `eventType`
  - `actorType`
  - `actorId`
  - `summaryLabel`
  - `metadataJson`
  - `createdAt`

### 64.7 화면 표시 계약
#### 64.7.1 채팅 화면
- 일반 채팅 타임라인에는 사건 전체 원문을 늘어놓지 않고 아래 수준의 시스템 카드만 노출한다.
  - `노쇼 신고가 접수됨`
  - `상대가 반박을 제출함`
  - `운영이 추가 자료를 요청함`
  - `판정 완료: 검토 결과 확인`
- 상세 근거/판정문 전문은 `case detail sheet` 또는 별도 사건 화면에서 확인한다.
- `report_locked` 또는 사건 검토 상태일 때 채팅 입력창은 일반 대화용이 아니라 `사건 관련 추가 자료 제출` CTA로 치환될 수 있어야 한다.

#### 64.7.2 내 거래 화면
- 거래 카드에는 `caseState`, `nextBestAction`, `decisionDueHint` 기반 요약만 노출한다.
- 예시:
  - `노쇼 검토 중 · 추가 자료 1건 필요`
  - `노쇼 판정 완료 · 결과 확인`
  - `재심 접수됨 · 운영 검토 중`
- 사용자가 해야 할 행동은 항상 1개 우선 CTA로 제한한다.

#### 64.7.3 운영 큐
- 운영 큐 기본 열:
  - `noShowCaseId`
  - `reservationId`
  - `claimant / respondent`
  - `caseState`
  - `evidenceSufficiencyTier`
  - `adjudicationSlaTier`
  - `latestTimelineEventAt`
  - `restrictionRiskPreview`
- 운영 상세에서는 채팅 원본, 예약/재일정 이력, 도착확인 이벤트, 증빙 번들, 과거 노쇼 관련 사건을 한 화면에서 연결 조회 가능해야 한다.

### 64.8 API candidate surface
#### 읽기 API 후보
- `GET /me/no-show-cases`
- `GET /no-show-cases/{noShowCaseId}`
- `GET /admin/no-show-cases`
- `GET /admin/no-show-cases/{noShowCaseId}`

#### 쓰기 API 후보
- `POST /reservations/{reservationId}/no-show-cases`
- `POST /no-show-cases/{noShowCaseId}/counter-position`
- `POST /no-show-cases/{noShowCaseId}/evidence`
- `POST /no-show-cases/{noShowCaseId}/withdraw`
- `POST /admin/no-show-cases/{noShowCaseId}/request-more-info`
- `POST /admin/no-show-cases/{noShowCaseId}/decisions`
- `POST /no-show-cases/{noShowCaseId}/appeals`

#### command 규칙
- 동일 `reservationId`에 active case가 있으면 `POST /reservations/{reservationId}/no-show-cases`는 새 사건 생성 대신 기존 사건 참조 + `ALREADY_EXISTS_ACTIVE_CASE` 응답 또는 idempotent success를 반환한다.
- `counter-position`은 상대방만 제출 가능하고, 이미 제출된 경우에는 overwrite가 아니라 revision 이력으로 저장한다.
- `decision` 작성은 운영 권한만 가능하며, 이전 decision이 있어도 새 버전 append 방식으로 저장한다.

### 64.9 DB / schema 파생 기준
- 최소 테이블 후보:
  - `no_show_cases`
  - `no_show_case_parties`
  - `no_show_case_evidences`
  - `no_show_case_decisions`
  - `no_show_case_timeline_events`
- 필수 제약:
  - `unique active case per reservation`
  - `one current party row per case + user`
  - `decision version unique per case`
- 인덱스 후보:
  - `(reservation_id, case_state)`
  - `(claim_opened_by_user_id, claim_opened_at desc)`
  - `(respondent_user_id, claim_opened_at desc)`
  - `(case_state, adjudication_sla_tier, claim_opened_at)`
  - `(no_show_case_id, created_at)` on timeline/evidence

### 64.10 review / trust / restriction side-effect 연결 원칙
- 사건 판정 전에는 후기 공개를 `pending_case_result`로 보류할 수 있어야 한다.
- `claim_upheld` 시:
  - upheld 대상 반대편 사용자에 대한 `trust aggregate` 하향 검토
  - 반복 패턴이면 `restrictionEscalationTier`에 따라 경고/제한 후보 생성
- `claim_rejected` 시:
  - 허위/경솔 claim 누적을 anti-abuse 신호로 반영 가능
- `shared_fault` 또는 `insufficient_evidence` 시:
  - 신뢰도 영향은 보수적으로 하며 자동 강한 제재는 피한다.
- side-effect는 사건 본문과 별도 projection으로 퍼질 수 있으나, 항상 `decisionId`를 reference key로 가져야 한다.

### 64.11 analytics 파생 기준
필수 이벤트 후보:
- `no_show_case_opened`
- `no_show_case_countered`
- `no_show_case_evidence_added`
- `no_show_case_decided`
- `no_show_case_appealed`
- `no_show_case_revised`

공통 속성 후보:
- `noShowCaseId`
- `reservationId`
- `listingId`
- `caseState`
- `caseOutcomeType`
- `evidenceSufficiencyTier`
- `adjudicationSlaTier`
- `decisionVersion`

핵심 분석 포인트:
- reservation 대비 no-show case 발생률
- case open → decision lead time
- insufficient evidence 비율
- appeal 발생률 / revision 비율
- upheld 대상의 재발률

### 64.12 오픈 질문 / 기본 가정
- 가정: active no-show case는 reservation 단위 1개로 제한하는 것이 MVP에서 가장 단순하다.
- 가정: 사건 상세는 참여자에게는 요약+본인 제출 자료 중심으로, 운영자에게만 full timeline을 노출한다.
- 오픈 질문:
  1. `shared_fault` 판정 시 후기 공개를 허용할지, 제한 공개로 둘지
  2. no-show case와 일반 report를 MVP에서 별도 큐로 둘지, 통합 moderation queue에서 subtype으로 다룰지
  3. appeal 제출 가능 횟수를 1회로 고정할지, 신규 증빙 기반 재개를 허용할지

## 65. 거래 스레드 종료 / 아카이브 / 재진입 계약
### 65.1 목표
- `채팅방은 남아 있는데 내 거래에서는 사라짐`, `완료 거래가 계속 액션 큐에 남음`, `종료된 상대와 다시 거래하고 싶은데 어디서 재진입하는지 모름` 같은 UX 혼선을 줄인다.
- 활성 거래 워크스페이스와 기록 보관함의 경계를 명확히 해 채팅목록, 내 거래, 알림, 프로필, 운영 타임라인이 같은 종료 기준을 사용하게 한다.
- 거래가 끝난 뒤에도 증빙/후기/재거래 맥락은 유지하되, 더 이상 실행이 필요 없는 스레드는 작업 큐에서 빠지도록 한다.

### 65.2 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `threadClosureReason` | `completed_confirmed` / `completed_auto_confirmed` / `cancelled_mutual` / `cancelled_owner` / `cancelled_system_expired` / `closed_by_restriction` / `closed_by_moderation` / `migrated_to_repeat_trade` | 스레드가 활성 큐에서 빠지는 직접 사유 |
| `threadArchiveTrigger` | `immediate_on_close` / `after_review_window` / `after_case_resolution` / `manual_hide_only` | 언제 archive view로 이동하는지 |
| `archiveVisibilityLevel` | `active_workspace` / `recently_closed` / `archived` / `staff_locked` | 사용자 화면에서 어느 레이어에 보이는지 |
| `reopenEligibilityState` | `not_allowed` / `duplicate_listing_required` / `same_listing_residual_allowed` / `repeat_trade_shortcut_allowed` | 종료 후 다시 거래를 시작할 수 있는 방식 |
| `threadResidency` | `chat_only` / `trade_workspace` / `both` | 채팅목록과 내 거래에 각각 남는 범위 |
| `archiveRetentionClass` | `standard_trade_record` / `dispute_extended` / `restriction_hold` | 보관/삭제 정책 파생 기준 |

원칙:
- `닫힘(closed)`은 곧바로 `삭제(deleted)`를 의미하지 않는다.
- 거래 실행 관점의 종료와 기록 보존 관점의 보관을 분리해 저장해야 한다.
- 사용자는 `왜 사라졌는지`, `어디에서 다시 찾을 수 있는지`, `다시 거래할 수 있는지`를 예측 가능하게 알아야 한다.

### 65.3 종료와 보관의 단계 구분
| 단계 | 의미 | 기본 노출 위치 | 대표 사용자 액션 |
|---|---|---|---|
| `active_workspace` | 현재 거래 진행 또는 액션 필요 | 내 거래 기본 탭 + 채팅목록 | 메시지, 예약, 완료, 신고 |
| `recently_closed` | 종결됐지만 후기/분쟁/증빙 확인 창이 열려 있음 | 내 거래 `최근 종료` / 채팅목록 하단 섹션 | 후기 작성, 기록 보기, 신고 |
| `archived` | 실무 액션은 끝났고 기록 보관 중심 | 보관함/검색/프로필 거래기록 | 재거래, 증빙 확인 |
| `staff_locked` | 정책/분쟁/제재로 사용자 재진입 제한 | 읽기 전용 상세, 운영 배너 | 제한 사유 확인, 이의제기 |

기본안:
- `completed_confirmed`, `completed_auto_confirmed`는 후기 마감 또는 dispute window 종료 전까지 `recently_closed`에 머문다.
- `cancelled_*`는 즉시 `recently_closed` 또는 `archived`로 갈 수 있으나, no-show/dispute 가능 창이 있으면 `recently_closed`를 거친다.
- 운영 잠금/제재 연계 스레드는 `staff_locked`가 `archived`보다 우선한다.

### 65.4 종료 트리거 규칙
#### 활성 스레드가 종료되는 조건
- 거래 완료 최종 확정
- 예약/거래가 취소되어 더 이상 액션 큐에 남길 이유가 없음
- 운영자 잠금/제재로 일반 사용자 액션이 중단됨
- 잔여 수량 없이 부분거래가 소진됨
- repeat trade shortcut이 새로운 thread를 생성하며 기존 thread를 기록 모드로 전환함

#### 종료되지 않는 조건
- 단순 미읽음 없음
- 예약 제안 만료 후 재협상 가능성이 높아 동일 스레드가 계속 쓰이는 경우
- `requested` completion, `disputed`, `no_show_case_opened`처럼 후속 결론이 아직 필요한 경우
- 한 매물에 잔여 수량이 남아 동일 상대와 계속 협의 중인 경우

### 65.5 채팅목록 vs 내 거래 residency 규칙
| 상황 | 채팅목록 | 내 거래 | 설명 |
|---|---|---|---|
| 활성 협의 중 | 유지 | 유지 | 기본 `both` |
| 완료 직후 후기 대기 | 유지 | `최근 종료`에 유지 | 사용자가 후기를 놓치지 않게 함 |
| 완전 종료 + 액션 없음 | 최근 대화 영역에서 제외, archive 검색 가능 | 기본 탭에서 제외, 보관함 이동 | 작업 큐 오염 방지 |
| 분쟁/노쇼 사건 진행 중 | 유지(읽기 또는 제한 입력) | 유지(사건 우선 배지) | 사건 해결 전 보관 금지 |
| 운영 잠금 | 유지 가능하나 `staff_locked` 표시 | 유지 가능하나 action disabled | 사용자가 맥락을 잃지 않게 함 |

원칙:
- 채팅목록은 `대화 기록 접근성`을, 내 거래는 `행동 우선순위`를 우선한다.
- 따라서 동일 thread라도 `threadResidency=chat_only` 상태가 될 수 있다.
- 사용자가 `내 거래`에서 사라진 thread를 `채팅 보관함` 또는 `거래 기록`에서 찾아갈 수 있어야 한다.

### 65.6 재진입 / 다시 거래하기 규칙
- 종료된 thread에서 허용 가능한 재진입은 아래 셋 중 하나로 제한한다.
  1. `duplicate_listing_required`: 새 매물/새 거래로 시작
  2. `same_listing_residual_allowed`: 잔여 수량이 남은 동일 매물에서 이어감
  3. `repeat_trade_shortcut_allowed`: 완료된 상대와 새 거래 thread를 빠르게 개설
- `completed` 확정 거래를 일반 메시지 몇 개로 다시 `active_workspace`로 되돌리는 것은 금지한다.
- 사용자가 종료된 대화에서 다시 메시지를 보내고 싶을 때, 시스템은 `새 거래 시작`, `다시 거래하기`, `남은 수량으로 이어가기` 중 허용된 CTA만 보여준다.
- `closed_by_restriction` 또는 `closed_by_moderation`은 기본적으로 `reopenEligibilityState=not_allowed`다.

### 65.7 화면 요구사항
#### 채팅목록
- 활성 스레드와 종료 스레드를 섞어 정렬하지 않고, 최소한 `진행 중` / `최근 종료` / `보관됨` 구분이 필요하다.
- 종료 카드에는 마지막 종료 사유와 재진입 가능 여부를 한 줄로 보여준다.
  - 예: `거래완료 · 후기 작성 가능`, `거래취소 · 기록 보관됨`, `운영 잠금 · 재진입 불가`

#### 내 거래
- 기본 탭에는 `active_workspace`와 `recently_closed` 중 액션이 남은 건만 노출한다.
- `archived`는 별도 탭 또는 필터로 분리해, 완료 기록을 찾을 수는 있지만 기본 액션 큐를 오염시키지 않게 한다.
- `staff_locked`는 숨기지 말고 제한 사유 배너와 함께 노출한다.

#### 거래 상세 / 스레드 상세
- 종료 시점, 종료 사유, 마지막 상태 전이, 후기/분쟁 가능 기간, 재거래 CTA를 한 모듈에서 보여줘야 한다.
- 종료 스레드에서 메시지 입력창을 제거하더라도, `기록 보기`, `증빙 다운로드`, `다시 거래하기` CTA는 남길 수 있다.

### 65.8 API candidate surface
#### 읽기
- `GET /me/trades?bucket=active|recently_closed|archived`
- `GET /chats?bucket=active|recently_closed|archived`
- `GET /trade-threads/{threadId}` 응답에 아래 필드 포함 후보:
```json
{
  "threadId": "tt_123",
  "tradeThreadState": "completed",
  "threadClosureReason": "completed_confirmed",
  "archiveVisibilityLevel": "recently_closed",
  "reopenEligibilityState": "repeat_trade_shortcut_allowed",
  "availableActions": ["create_review", "start_repeat_trade"],
  "threadResidency": "both"
}
```

#### 쓰기
- `POST /trade-threads/{threadId}/archive` (사용자 수동 보관, optional)
- `POST /trade-threads/{threadId}/unarchive` (운영 또는 사용자 개인 보관 해제, optional)
- `POST /trade-threads/{threadId}/start-repeat-trade`

원칙:
- archive는 공용 도메인 상태와 사용자 개인 정리 상태를 구분해야 한다. 예: `userHiddenAt`와 `archiveVisibilityLevel`은 별개일 수 있다.
- 사용자 개인 보관 해제는 상대방의 화면 상태를 바꾸지 않는다.

### 65.9 데이터 모델 파생 기준
#### TradeThread 또는 projection 필드 후보
| 필드 | 설명 |
|---|---|
| `closedAt` | 활성 큐에서 빠진 시각 |
| `threadClosureReason` | 종료 사유 |
| `archiveVisibilityLevel` | 현재 노출 레이어 |
| `archiveEnteredAt` | 보관함 진입 시각 |
| `reopenEligibilityState` | 재진입 방식 |
| `userHiddenByAAt` / `userHiddenByBAt` | 개인 보관/숨김 상태 |
| `lastActionableAt` | 마지막 행동 필요 시각 |
| `archiveRetentionClass` | retention 파생 키 |

원칙:
- `closedAt`은 기록상 불변에 가깝고, `archiveVisibilityLevel`은 사건/후기/운영 상태에 따라 재계산될 수 있다.
- read model은 `activeCount`, `recentlyClosedCount`, `archivedCount`를 빠르게 계산할 수 있어야 한다.

### 65.10 운영 / 정책 연계
- 운영자는 종료 스레드를 다시 활성화하지 않고, 필요한 경우 `새 사건 생성` 또는 `repeat trade shortcut 허용 해제`처럼 side-effect를 명시적으로 남겨야 한다.
- 제재/분쟁 hold가 걸린 스레드는 `archiveRetentionClass=restriction_hold` 또는 `dispute_extended`로 승격한다.
- 사용자가 스레드를 개인 보관 처리해도, 운영 보존·신고·감사 열람 가능성에는 영향을 주지 않는다.

### 65.11 analytics 파생 기준
필수 이벤트 후보:
- `trade_thread_closed`
- `trade_thread_archived`
- `trade_thread_unarchived`
- `trade_thread_reentered`
- `repeat_trade_started_from_archived_thread`

공통 속성 후보:
- `threadId`
- `listingId`
- `threadClosureReason`
- `archiveVisibilityLevel`
- `reopenEligibilityState`
- `threadResidency`
- `daysUntilArchive`

핵심 분석 포인트:
- active → recently_closed → archived 전환 소요시간
- 후기 작성률이 `recently_closed` 유지 기간과 어떤 상관이 있는지
- archived thread 기반 repeat trade 시작률
- 작업 큐(내 거래)에서 종료 thread 잔존이 재방문/응답성에 미치는 영향

### 65.12 오픈 질문 / 기본 가정
- 가정: 완료 거래는 후기 작성 가능 창이 끝나기 전까지 `recently_closed`에 남기는 것이 MVP에서 가장 직관적이다.
- 가정: 채팅목록과 내 거래는 같은 thread를 보더라도 서로 다른 bucket 계산을 허용한다.
- 오픈 질문:
  1. 사용자의 개인 `archive/hide` 상태를 상대방과 완전히 분리된 preference로 둘지
  2. repeat trade shortcut이 기존 thread 복제인지, 새 thread+과거 맥락 연결인지
  3. `recently_closed`를 별도 탭으로 둘지, `종료됨` 섹션 inline grouping으로 둘지


## 66. 재거래 시작(Start Repeat Trade) / 과거 맥락 승계 / 새 스레드 생성 계약
### 66.1 목표
- 완료되거나 취소된 거래 이후 같은 상대와 다시 거래할 때, 기존 thread를 억지로 되살리는 대신 **새 거래 스레드**를 예측 가능하게 생성한다.
- 사용자가 `이전 거래의 좋은 맥락은 이어가고 싶지만, 과거 분쟁/노쇼/민감 정보는 그대로 끌고 오고 싶지 않다`는 요구를 제품 구조로 해결한다.
- 프로필/보관함/채팅/내 거래/운영/분석이 모두 동일한 기준으로 `재거래`를 해석하게 한다.

### 66.2 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `repeatTradeSourceType` | `completed_trade` / `cancelled_trade` / `archived_thread` / `favorite_counterparty` / `profile_shortcut` | 재거래 시작의 출발점 |
| `repeatTradeSpawnMode` | `new_thread_same_listing_residual` / `new_thread_new_listing` / `new_thread_no_listing_yet` | 어떤 도메인 루트에서 새 thread를 만드는지 |
| `contextCarryForwardPolicy` | `none` / `counterparty_only` / `counterparty_and_last_meeting_prefill` / `counterparty_and_listing_snapshot_reference` | 이전 거래 맥락을 얼마나 승계하는지 |
| `carryForwardDisclosureLevel` | `hidden_internal` / `visible_summary` / `visible_full_to_participants` | 새 거래 화면에서 과거 맥락을 어느 수준으로 보여주는지 |
| `repeatTradeEligibilityState` | `allowed` / `allowed_with_warning` / `blocked_by_restriction` / `blocked_by_block_relation` / `blocked_by_open_case` | 재거래 시작 가능 여부 |
| `repeatTradeLinkageType` | `fresh_start` / `linked_history` / `risk_review_linked` | 새 thread와 과거 thread의 관계 |

원칙:
- 재거래는 **기존 thread reopen**이 아니라 **새 thread spawn**을 기본으로 한다.
- 과거 거래와 새 거래는 연결될 수 있지만, 상태 머신과 unread/action queue는 분리되어야 한다.
- 과거 거래의 민감 정보나 제재 사유를 새 상대 화면에 과도하게 노출하지 않는다.

### 66.3 새 thread 생성 기본 원칙
- `completed_confirmed`, `completed_auto_confirmed`, `cancelled_mutual`, `cancelled_owner`, `cancelled_system_expired` 상태의 종료 thread는 재거래 시작의 source가 될 수 있다.
- `closed_by_restriction`, `closed_by_moderation`, `staff_locked`, `open dispute`, `active no-show case`가 연결된 thread는 기본적으로 재거래 차단 또는 경고 후 차단한다.
- 재거래 시작 시 기존 thread의 미읽음/액션 필요 상태를 되살리지 않고, 새 `tradeThreadId` / `chatRoomId`를 발급한다.
- 사용자가 보관함에서 `다시 거래하기`를 눌러도 기존 메시지 입력창이 풀리는 것이 아니라, 새 거래 composer 또는 새 채팅 시작 confirmation sheet가 열려야 한다.

### 66.4 spawn mode 규칙
| 상황 | 권장 `repeatTradeSpawnMode` | 설명 |
|---|---|---|
| 완료된 단건 거래 후 같은 상대와 새 건 논의 | `new_thread_no_listing_yet` 또는 `new_thread_new_listing` | 새 매물 또는 새 문의에서 시작 |
| 부분거래 후 잔여 수량으로 이어가기 | `new_thread_same_listing_residual` | 기존 listing residual과 연결하되 thread는 새로 생성 |
| 취소됐지만 상대와 다시 시간 맞춰 거래 재개 | `new_thread_new_listing` 또는 `new_thread_no_listing_yet` | 실패 맥락은 남기되 실행은 새 흐름으로 시작 |
| 프로필에서 단골 상대에게 직접 재문의 | `new_thread_no_listing_yet` | listing 선택 전 대화 시작 가능 여부는 정책 결정 |

기본안:
- MVP에서는 `completed trade -> new_thread_new_listing` 또는 `new_thread_no_listing_yet`를 우선 지원하고, residual 기반 shortcut은 부분거래 기능과 같이 제한적으로 연동한다.
- `same listing residual`을 허용하더라도 completion history와 새 negotiation timeline은 분리 저장한다.

### 66.5 과거 맥락 승계 범위
#### 기본 승계 허용
- 상대방 식별자
- 최근 사용 거래 방식(`in_game` / `offline_pc_bang`)
- 최근 사용 서버/거래 선호
- 최근 meeting template 요약(사용자 개인 prefill 수준)
- repeat-trade affinity / 선호 신호

#### 기본 승계 금지
- 과거 정확 장소 원문
- 과거 dispute statement / evidence 원문
- 운영 내부 메모 / 제재 사유 상세
- 과거 no-show claim 세부 본문
- 상대가 원치 않을 수 있는 민감 free-text 메모

원칙:
- `prefill`은 허용하되 `auto-send`는 금지한다.
- 승계된 값은 새 거래 생성 전 사용자 확인을 거쳐야 하며, 상대방에게는 `이전 거래 기준으로 제안됨` 정도의 summary만 필요 시 노출한다.

### 66.6 eligibility / gating 규칙
| 조건 | `repeatTradeEligibilityState` | 처리 원칙 |
|---|---|---|
| 양측 모두 active, 상호 차단 없음, 열린 사건 없음 | `allowed` | 즉시 재거래 시작 가능 |
| 최근 no-show 판정 또는 제한 이력 있으나 차단은 아님 | `allowed_with_warning` | 경고 배너 + 안전 수칙 노출 |
| 상호 차단 관계 존재 | `blocked_by_block_relation` | CTA 비노출 또는 차단 해제 안내 |
| 활성 restriction / moderation hold 존재 | `blocked_by_restriction` | 재거래 생성 불가 |
| 완료 분쟁, no-show case, appeal 진행 중 | `blocked_by_open_case` | 사건 종료 전 재거래 금지 |

세부 원칙:
- `allowed_with_warning`은 단순 낙인 표시가 아니라, `기록은 남기고 약속/대금 확인을 더 신중히` 수준의 중립 가이드를 제공한다.
- 운영자 override가 있더라도 기존 사건을 덮어쓰지 않고 `risk_review_linked` 관계로 새 thread를 생성한다.

### 66.7 화면 요구사항
#### 보관함 / 종료 스레드 상세
- `다시 거래하기` CTA는 `repeatTradeEligibilityState`가 `allowed` 또는 `allowed_with_warning`일 때만 노출한다.
- CTA 노출 시 함께 보여야 할 요약:
  - 이전 거래 종료 상태
  - 같은 상대와의 완료 거래 수
  - 최근 거래 방식/서버 선호
  - 재거래 차단 또는 경고 사유

#### 프로필 화면
- `repeat trade shortcut`은 단골/반복 거래 맥락이 있을 때만 노출 가능하다.
- 프로필 CTA가 새 거래를 시작하더라도 이전 채팅 기록을 자동으로 펼치지 않는다.
- `allowed_with_warning`이면 안전 가드 문구를 먼저 보여준다.

#### 새 거래 시작 sheet / composer
- 사용자가 아래를 확인/선택할 수 있어야 한다.
  1. 기존 매물 기반인지 새 거래인지
  2. 최근 서버/장소/방식 prefill 사용 여부
  3. 새 거래 제목/첫 메시지 초안
  4. 과거 맥락 링크 여부(`linked_history` summary 표시)
- 첫 메시지 입력창에는 과거 요약을 텍스트로 자동 주입하지 않는다.

### 66.8 새 thread와 과거 thread의 linkage 규칙
- 새 thread에는 `sourceTradeThreadId(optional)`를 둘 수 있다.
- 과거 thread 상세에는 `repeatTradeSpawnedThreadId(optional)` 또는 count를 둘 수 있다.
- linkage는 아래 수준으로 나눈다.
  - `fresh_start`: 사용자 화면에 과거 연결 노출 없음, 내부 분석만 연결
  - `linked_history`: 참여자에게 `이전 거래에서 다시 시작` summary 노출
  - `risk_review_linked`: 운영자에게만 강한 연결 노출, 사용자에게는 제한 안내만 제공

원칙:
- linkage는 read model/analytics/운영 추적에는 중요하지만, 사용자 액션 큐에서는 독립 thread로 취급한다.
- 새 thread의 후기/분쟁/노쇼는 과거 thread에 append하지 않고 새 사건으로 생성한다.

### 66.9 API candidate surface
#### 읽기
- `GET /trade-threads/{threadId}/repeat-trade-eligibility`
- `GET /users/{userId}/repeat-trade-context`

예시 응답:
```json
{
  "threadId": "tt_123",
  "repeatTradeEligibilityState": "allowed_with_warning",
  "repeatTradeSourceType": "archived_thread",
  "recommendedSpawnMode": "new_thread_new_listing",
  "contextCarryForwardPolicy": "counterparty_and_last_meeting_prefill",
  "warningHints": [
    {
      "code": "RECENT_NO_SHOW_HISTORY_EXISTS",
      "message": "최근 약속 변경/노쇼 이력이 있어요. 새 약속 전 시간을 다시 확인하세요."
    }
  ],
  "availableActions": ["start_repeat_trade"]
}
```

#### 쓰기
- `POST /trade-threads/{threadId}/start-repeat-trade`
- `POST /users/{userId}/start-repeat-trade`

Request 예시:
```json
{
  "spawnMode": "new_thread_new_listing",
  "contextCarryForwardPolicy": "counterparty_and_last_meeting_prefill",
  "listingId": null,
  "draft": {
    "serverId": "ken-01",
    "tradeMethod": "in_game",
    "openingMessage": "지난번처럼 오늘 저녁 거래 가능하실까요?"
  }
}
```

Response 예시:
```json
{
  "tradeThreadId": "tt_456",
  "chatRoomId": "chat_456",
  "repeatTradeLinkageType": "linked_history",
  "sourceTradeThreadId": "tt_123",
  "tradeThreadState": "open",
  "availableActions": ["send_message", "propose_reservation"]
}
```

### 66.10 데이터 모델 파생 기준
#### TradeThread / ChatRoom 후보 필드
| 필드 | 설명 |
|---|---|
| `sourceTradeThreadId` | 재거래의 출발 thread |
| `repeatTradeSourceType` | source 종류 |
| `repeatTradeLinkageType` | 새/과거 관계 |
| `repeatTradeEligibilityState` | read model 계산값 또는 snapshot |
| `contextCarryForwardPolicy` | 생성 시 적용된 승계 정책 |
| `carryForwardDisclosureLevel` | 사용자 노출 수준 |
| `spawnedFromCompletionId` | 완료 거래 기반 재거래 추적 |

#### 사용자/상대 관계 후보 객체
- `CounterpartyAffinity`
  - `userId`
  - `counterpartyUserId`
  - `completedTradeCount`
  - `lastCompletedAt`
  - `lastTradeMethod`
  - `lastServerId`
  - `repeatTradeAffinityLevel`
  - `repeatTradeBlockedReasonCode(optional)`

원칙:
- affinity는 recommendation/shortcut/read model용 요약 객체이고, 제재 판정의 단독 근거로 쓰지 않는다.
- repeat trade linkage는 projection으로도 계산 가능하지만, source thread id는 원천 데이터에 남기는 편이 운영 추적에 유리하다.

### 66.11 운영 / 정책 연계
- 운영자는 반복 거래가 기존 사건 회피 수단이 되지 않도록 `open case -> repeat trade blocked`를 기본값으로 둔다.
- 고위험 사용자의 재거래 생성 시 운영 큐에 자동 적재하지는 않더라도, `risk_review_linked` metadata는 남겨야 한다.
- 사용자가 차단 관계를 해제한 직후 재거래를 시작하는 패턴, 분쟁 종료 직후 즉시 재거래하는 패턴은 이상징후 분석 대상으로 볼 수 있다.
- repeat trade shortcut은 신뢰 강화를 위한 편의 기능이지, 기존 후기/분쟁을 세탁하는 기능이 되어서는 안 된다.

### 66.12 analytics 파생 기준
필수 이벤트 후보:
- `repeat_trade_cta_impression`
- `repeat_trade_cta_click`
- `repeat_trade_started`
- `repeat_trade_blocked`
- `repeat_trade_prefill_used`

공통 속성 후보:
- `sourceTradeThreadId`
- `repeatTradeSourceType`
- `repeatTradeSpawnMode`
- `repeatTradeEligibilityState`
- `contextCarryForwardPolicy`
- `repeatTradeLinkageType`
- `daysSincePreviousTradeClosed`

핵심 분석 포인트:
- archived thread 기반 repeat trade 시작률
- repeat trade가 첫 거래 대비 예약 확정률/완료율을 얼마나 높이는지
- prefill 사용 시 예약 생성 시간 단축 여부
- blocked/allowed_with_warning 비율과 이후 신고율의 상관관계

### 66.13 오픈 질문 / 기본 가정
- 가정: MVP에서는 재거래를 `새 thread 생성`으로만 허용하고, 기존 thread reopen은 지원하지 않는다.
- 가정: 과거 exact 장소/민감 free-text는 carry forward 대상에서 제외하는 것이 안전하다.
- 오픈 질문:
  1. listing 없이 상대 프로필에서 바로 새 거래 thread를 여는 것을 MVP에 포함할지
  2. repeat trade shortcut 노출 기준을 `1회 완료만 있어도` 허용할지, `2회 이상 완료`부터 강화할지
  3. `allowed_with_warning` 상태에서 어떤 안전 문구가 과도한 낙인 없이 충분한 경고가 되는지
## 100. 운영 공지 / 시스템 배너 / 정책 메시지 노출 계약
### 100.1 목표
- 긴급 점검, 정책 변경, 거래 안전 경고, 기능 제한 안내를 각 화면이 제각각 문자열로 구현하지 않도록 **canonical notice model**을 정의한다.
- 단순 마케팅 배너가 아니라 거래 완료율·안전·운영정책 집행에 직접 영향을 주는 `system notice`를 구조화한다.
- 홈/목록/상세/채팅/내 거래/알림함/운영 백오피스/API가 같은 기준으로 `무엇을`, `누구에게`, `어디에`, `언제까지`, `반드시 확인시킬지` 해석하도록 한다.

### 100.2 적용 범위
본 섹션의 notice는 아래 범위를 다룬다.
- 긴급 시스템 점검, 부분 장애, degraded mode 안내
- 거래 필수 정책 변경(약관 재동의, 외부연락처 정책 강화, 후기 공개 규칙 변경)
- 계정/기능 제한 안내(채팅 제한, 매물 등록 제한, 추가 인증 요구)
- 거래 안전 경고(사기 다발 패턴, 특정 행동 금지, 노쇼 급증 시간대 주의)
- 개인 맞춤 과제성 안내(예약 응답 필요, 정책 확인 필요, 설정 갱신 필요)

적용 제외:
- 일반 마케팅/프로모션 배너
- 임시 A/B 테스트 카피
- 단일 객체 내부 설명 툴팁

### 100.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `noticeAudienceScope` | `all_users` / `guests` / `members` / `active_traders` / `listing_owners` / `chat_participants` / `restricted_users` / `admins_only` | 대상자 범위 |
| `noticeSurface` | `home_banner` / `listing_feed_inline` / `listing_detail_banner` / `chat_banner` / `trade_workspace_banner` / `notification_inbox` / `modal_blocking` / `settings_banner` | 노출 위치 |
| `noticeSeverity` | `info` / `attention` / `warning` / `critical` | 시각적·행동 우선순위 |
| `actionRequirementLevel` | `optional` / `recommended` / `required_before_next_action` / `blocking` | 사용자 행동 필요도 |
| `ackRequirementType` | `none` / `dismiss_once` / `ack_once_per_version` / `ack_every_session_until_resolved` | 확인 이력 요구 수준 |
| `noticeLifecycleState` | `draft` / `scheduled` / `active` / `expired` / `withdrawn` | 공지 생명주기 |
| `noticeTriggerType` | `manual_ops` / `policy_release` / `system_incident` / `risk_rule` / `user_state_change` | 생성 원인 |
| `noticeContextScope` | `global` / `server_specific` / `listing_specific` / `trade_thread_specific` / `account_specific` | 문맥 적용 범위 |
| `noticeSuppressionReasonCode` | `already_acknowledged` / `not_applicable_state` / `surface_capacity_limit` / `lower_priority_overridden` / `session_recently_seen` | 미노출 사유 |

원칙:
- 공지는 단순 문자열이 아니라 `targeting + severity + action gating + acknowledgement`를 가진 객체로 본다.
- 동일 메시지가 홈/채팅/알림함에서 재사용되더라도 source notice id는 같고 surface projection만 다르게 관리한다.

### 100.4 공지 유형 taxonomy
| noticeType | 목적 | 예시 | 기본 severity |
|---|---|---|---|
| `incident_notice` | 시스템 장애/점검 공지 | 채팅 지연, 예약 알림 지연 | `warning` 또는 `critical` |
| `policy_notice` | 서비스 정책 변경 안내 | 외부 연락처 정책 강화 | `attention` |
| `safety_notice` | 거래 안전 경고 | 선입금 유도 주의 | `warning` |
| `eligibility_notice` | 기능 사용 조건 안내 | 추가 인증 필요 | `attention` |
| `restriction_notice` | 기능 제한/제재 안내 | 채팅 기능 3일 제한 | `critical` |
| `task_notice` | 특정 행동 촉구 | 예약 응답 필요, 약관 재동의 필요 | `warning` |
| `education_notice` | 사용법/가이드 | 예약 전에 장소를 확정하세요 | `info` |

### 100.5 우선순위 및 동시 노출 규칙
- 하나의 surface에는 동시에 너무 많은 공지가 보이지 않도록 **최대 1개의 상단 sticky notice + 필요 시 1개의 inline notice**를 기본 원칙으로 한다.
- 우선순위는 아래 순서를 따른다.
  1. `blocking` restriction/policy notice
  2. `critical` incident notice
  3. 거래 진행 중 action-required task notice
  4. safety notice
  5. 일반 info/education notice
- 동일 사용자가 이미 수행 가능한 다음 행동을 알고 있고 방해가 큰 경우, `optional/info` notice는 suppress 가능하다.
- `trade_thread_specific` notice는 전역 공지보다 해당 채팅/내 거래 화면에서 우선한다.

### 100.6 action gating 규칙
| actionRequirementLevel | 의미 | 예시 |
|---|---|---|
| `optional` | 읽지 않아도 기능 진행 가능 | 신규 기능 소개 |
| `recommended` | 읽는 것이 좋지만 건너뛸 수 있음 | 사기 주의 가이드 |
| `required_before_next_action` | 특정 액션 전 확인 필요 | 첫 거래 전 안전 가이드 확인 |
| `blocking` | 기능 사용 전 필수 처리 | 약관 재동의, 제한 안내 확인 |

세부 원칙:
- `required_before_next_action`은 사용자의 현재 flow를 최대한 보존하면서 해당 행동 직전에만 개입해야 한다.
- `blocking`은 남용 금지. 거래 성사와 직접 관련된 정책/권리/계정 상태 변화에만 사용한다.
- `restriction_notice`가 `blocking`일 때는 왜 차단되었는지, 언제 해제되는지, 어디서 이의제기하는지 함께 제공해야 한다.

### 100.7 acknowledgement(확인 이력) 계약
- 공지별로 dismiss와 acknowledge를 구분한다.
  - `dismiss`: 현재 surface에서 닫기
  - `acknowledge`: 정책/제한/가이드 확인 완료 기록
- `ack_once_per_version`은 정책 문구가 바뀌면 다시 요구할 수 있어야 한다.
- `dismiss_once`는 같은 세션 또는 일정 기간 동안만 재노출을 억제한다.
- `incident_notice`는 상태가 해결되기 전까지 세션마다 다시 노출될 수 있다.
- 사용자의 acknowledgement 이력은 단순 프론트 로컬스토리지에만 두지 않고 서버 저장을 기본으로 한다.

### 100.8 surface별 요구사항
#### 홈
- 전역 점검/정책/안전 notice의 기본 surface다.
- 거래와 무관한 info notice는 홈에서만 소화하고, 채팅/내 거래에 과도하게 반복하지 않는다.

#### 목록/상세
- `listing_feed_inline`, `listing_detail_banner`는 서버 특화 공지, 카테고리 특화 안전 경고, high-risk listing 경고에 사용한다.
- 매물 상세에서는 listing/trust 상태와 직접 관련된 공지가 전역 공지보다 우선할 수 있다.

#### 채팅/내 거래
- 예약 응답 필요, 재일정 미확인, 정책 제한, 기능 degrade 같은 거래 실행 notice가 최우선이다.
- 채팅 화면 notice는 메시지 타임라인과 섞이지 않고 별도 banner로 표현한다.
- 거래 thread-specific notice는 해당 thread 해결 후 자동 만료되어야 한다.

#### 알림함
- notice 자체를 알림함 항목으로 투사할 수 있으나, 모든 notice를 알림함에 적재하지는 않는다.
- `blocking`, `critical`, `required_before_next_action` 수준 notice는 인박스 기록 보존을 우선한다.

#### 설정/마이페이지
- 정책 재동의, 제한 이유, 알림 설정 갱신, 본인 인증 요구 notice의 canonical landing surface로 사용한다.

### 100.9 notice와 기존 도메인 객체의 관계
- `restriction_notice`는 Restriction 객체를 대체하지 않는다. Restriction은 판정 원천이고, notice는 사용자 전달 projection이다.
- `policy_notice`는 Policy Acceptance/Re-consent 객체와 연결되며, acknowledgement가 acceptance를 곧바로 의미하지는 않는다.
- `incident_notice`는 시스템 상태 원천(ops incident)과 연결되되, 사용자별 노출 여부는 별도 projection으로 계산한다.
- `task_notice`는 TradeThread, Reservation, Reschedule, CompletionStage 등에서 파생될 수 있다.

### 100.10 데이터 모델 후보
#### SystemNotice
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `noticeId` | 필수 | 공지 식별자 |
| `noticeType` | 필수 | taxonomy 값 |
| `title` | 필수 | 제목 |
| `body` | 필수 | 본문 |
| `severity` | 필수 | `info` / `attention` / `warning` / `critical` |
| `actionRequirementLevel` | 필수 | 행동 필요도 |
| `ackRequirementType` | 필수 | 확인 요구 타입 |
| `audienceScope` | 필수 | 대상 범위 |
| `contextScope` | 필수 | 문맥 범위 |
| `contextRefType` | 선택 | listing / trade_thread / policy / restriction / incident |
| `contextRefId` | 선택 | 연결 객체 ID |
| `startsAt` | 필수 | 노출 시작 시각 |
| `endsAt` | 선택 | 노출 종료 시각 |
| `state` | 필수 | lifecycle state |
| `triggerType` | 필수 | 생성 원인 |
| `createdBy` | 필수 | system 또는 admin |
| `versionKey` | 선택 | 정책 버전/문구 버전 |

#### UserNoticeState
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `userNoticeStateId` | 필수 | 식별자 |
| `noticeId` | 필수 | 대상 notice |
| `userId` | 필수 | 사용자 |
| `deliveryState` | 필수 | `eligible` / `shown` / `suppressed` / `expired` |
| `firstShownAt` | 선택 | 첫 노출 시각 |
| `lastShownAt` | 선택 | 마지막 노출 시각 |
| `dismissedAt` | 선택 | 닫은 시각 |
| `acknowledgedAt` | 선택 | 확인 시각 |
| `suppressionReasonCode` | 선택 | 미노출 사유 |
| `surfaceLastSeen` | 선택 | 마지막 노출 surface |

### 100.11 API 후보
#### 사용자용
- `GET /system-notices`
- `POST /system-notices/{noticeId}/dismiss`
- `POST /system-notices/{noticeId}/acknowledge`

응답 예시:
```json
{
  "notices": [
    {
      "noticeId": "not_123",
      "noticeType": "task_notice",
      "severity": "warning",
      "surface": "trade_workspace_banner",
      "title": "예약 응답이 필요해요",
      "body": "상대가 시간 변경을 요청했어요. 수락 또는 거절해 주세요.",
      "actionRequirementLevel": "required_before_next_action",
      "ackRequirementType": "none",
      "relatedAction": {
        "actionCode": "open_reschedule_card",
        "deepLink": "/trades/thread_123"
      }
    }
  ]
}
```

#### 관리자용
- `POST /admin/system-notices`
- `PATCH /admin/system-notices/{noticeId}`
- `POST /admin/system-notices/{noticeId}/activate`
- `POST /admin/system-notices/{noticeId}/expire`
- `GET /admin/system-notices`

### 100.12 운영 가드레일
- 운영자는 동일 대상군에 중복/상충되는 notice를 과도하게 활성화할 수 없어야 한다.
- `critical` 또는 `blocking` notice는 승인 체인 또는 최소 role 기준을 둔다.
- 공지 수정은 덮어쓰기보다 versioned edit를 우선 검토한다. 특히 정책 공지/제한 안내는 사후 감사가 가능해야 한다.
- 해결된 incident notice는 즉시 숨기기보다 `resolved` 상태 공지로 짧게 전환할 수 있다.

### 100.13 analytics 이벤트 후보
- `notice_impression`
- `notice_dismiss`
- `notice_acknowledge`
- `notice_cta_click`
- `notice_suppressed`

공통 속성 후보:
- `noticeId`
- `noticeType`
- `severity`
- `surface`
- `actionRequirementLevel`
- `contextScope`
- `sourceObjectType`
- `sourceObjectId`

핵심 분석 포인트:
- blocking notice 노출 후 이탈률
- safety notice 확인 후 외부연락처 위반률 감소 여부
- incident notice 노출군의 CS 문의 감소 여부
- re-consent notice 노출 후 completion rate 하락/회복 추이

### 100.14 오픈 질문 / 가정
- 가정: MVP에서도 최소한 `restriction_notice`, `policy_notice`, `task_notice`, `incident_notice` 4종은 필요하다.
- 오픈 질문: 전역 공지를 알림함에 항상 적재할지, severity 기준으로만 적재할지 확정 필요
- 오픈 질문: 공지 acknowledgement를 약관 동의/정책 acceptance와 어디까지 분리할지 세부 정책 필요
- 오픈 질문: guest 대상 notice를 SEO 공개 페이지에도 렌더링할지 여부 확정 필요

## 101. 사용자 문의/지원 케이스(Support Case) / 운영 커뮤니케이션 계약
### 101.1 목표
- 사용자의 `문의하기`, 신고 후속 질문, 제재 이의제기, 분쟁 보완자료 제출, 정책 해석 요청이 서로 다른 임시 채널로 흩어지지 않도록 하나의 `Support Case` 단위로 수렴한다.
- 운영자가 `Report`, `Dispute`, `Restriction`, `Notice`, `Policy Acceptance`와 연계된 사용자 커뮤니케이션을 같은 사건 맥락에서 처리할 수 있게 한다.
- 제품 화면, 백오피스, 알림, 템플릿 메시지, SLA, 감사 로그가 동일한 케이스 상태 머신을 공유하게 한다.

### 101.2 적용 범위
Support Case는 아래 범위를 다룬다.
- 일반 도움 요청: 기능 사용법, 계정 접근 문제, 알림/채팅 오동작 문의
- 거래 관련 문의: 예약 상태 이해, 완료 처리 질문, 후기 공개 시점 문의
- 안전/운영 후속 문의: 신고 처리 현황 확인, 제재 사유 문의, 복구 요청
- 사건 연결형 문의: 기존 `reportId`, `disputeId`, `restrictionId`, `noticeId`에 대한 추가 설명/자료 제출

적용 제외:
- 실시간 거래 협의는 채팅/거래 thread에서 처리한다.
- 운영자 내부 협업 메모만 존재하는 항목은 Support Case를 만들지 않는다.
- 완전히 외부 채널(예: 법무/결제사)로 이관된 사안은 linked external reference만 남기고 내부 Support Case는 종결 또는 handoff 상태로 둔다.

### 101.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `supportCaseType` | `general_help` / `account_access` / `trade_question` / `report_followup` / `dispute_followup` / `restriction_appeal` / `policy_question` / `bug_report` | 문의 성격 |
| `caseOriginSurface` | `help_center` / `report_detail` / `dispute_detail` / `restriction_notice` / `chat_escalation` / `admin_outreach` | 케이스 생성 출처 |
| `supportCaseState` | `open` / `awaiting_user` / `awaiting_staff` / `in_review` / `resolved` / `closed` / `handoff_external` | 케이스 상태 |
| `userExpectationState` | `acknowledged` / `under_review` / `action_taken` / `needs_more_info` / `no_action_possible` | 사용자에게 보여주는 기대 상태 |
| `resolutionCommunicationMode` | `in_app_thread` / `notification_only` / `policy_notice_link` / `manual_admin_message` | 결과 전달 방식 |
| `caseSensitivityTier` | `standard` / `safety` / `privacy` / `legal_risk` | 접근 통제/우선순위 기준 |
| `caseLinkPolicy` | `standalone` / `report_primary` / `dispute_primary` / `restriction_primary` | 다른 사건과의 주종 관계 |

원칙:
- Support Case는 Report/Dispute를 대체하지 않고, **사용자 커뮤니케이션 워크스페이스** 역할을 맡는다.
- 사용자 화면에는 복잡한 내부 case graph를 노출하지 않고 `무엇을 기다리는지`, `언제 답이 오는지`, `추가 자료가 필요한지`만 명확히 보여준다.

### 101.4 Support Case 생성 규칙
- 아래 상황에서는 Support Case 자동 생성 또는 생성 제안을 기본안으로 둔다.
  1. 사용자가 신고 상세에서 `처리 현황 문의` 또는 `추가 설명 보내기`를 누름
  2. 사용자가 제재 알림 화면에서 `이의제기`를 시작함
  3. 운영자가 분쟁/신고/제재 처리 중 추가 자료를 요청함
  4. 사용자가 고객센터/도움말에서 문의를 시작함
- 하나의 원사건에 대해 동시 활성 Support Case는 1건을 기본으로 하며, 중복 생성 시 기존 케이스로 라우팅한다.
- 일반 bug report처럼 원사건이 없는 경우 `caseLinkPolicy=standalone`으로 생성한다.

### 101.5 상태 머신
#### 기본 흐름
- 사용자 또는 운영자 생성 → `open`
- 운영자가 확인 후 처리중 전환 → `in_review`
- 사용자 추가 자료가 필요하면 → `awaiting_user`
- 운영자 응답이 필요하면 → `awaiting_staff`
- 결론 또는 안내 완료 시 → `resolved`
- 일정 기간 무응답/사안 종료 후 → `closed`

#### 상태 해석 규칙
- `open`은 아직 triage 전 상태다.
- `awaiting_user`는 사용자의 추가 자료 제출 기한을 동반해야 한다.
- `resolved`는 답변 완료 상태이지, 반드시 사용자가 만족했음을 뜻하지 않는다.
- `closed`는 재오픈 불가가 아니라, 새 자료가 생기면 같은 케이스 재개 또는 새 케이스 spawn 정책을 둘 수 있다.
- `handoff_external`은 내부 답변만으로 처리할 수 없는 사안(법적 요청 등)에 한정한다.

### 101.6 원사건 링크 규칙
| 원사건 | Support Case 역할 | 링크 원칙 |
|---|---|---|
| `Report` | 신고 후속 문의/상태 확인/추가 증빙 제출 창구 | `linkedReportId` 필수 가능 |
| `Dispute` | 운영 요청 응답/판정 설명/추가 자료 커뮤니케이션 | `linkedDisputeId` 우선 |
| `Restriction` | 제재 사유 설명/이의제기/해제 조건 안내 | `linkedRestrictionId` 우선 |
| `Notice` | 정책 배너/재동의/기능제한에 대한 문의 | `linkedNoticeId` 선택 |
| 없음 | 일반 도움/버그 제보/정책 질문 | standalone |

원칙:
- Report/Dispute의 판단 상태와 Support Case의 대화 상태를 혼동하지 않는다.
- 예를 들어 `reportStatus=resolved` 이후에도 사용자가 후속 질문을 하면 Support Case는 열릴 수 있다.
- Support Case 해결은 원사건 상태를 자동 변경하지 않는다. 단, 운영자 명시 액션이 있을 때만 side-effect를 발생시킨다.

### 101.7 사용자 화면 요구사항
#### Help / 내 문의
- 사용자는 `내 문의` 목록에서 아래를 확인할 수 있어야 한다.
  - 문의 유형
  - 현재 상태
  - 마지막 운영 응답 시각
  - 추가 자료 필요 여부
  - 연결된 사건 요약(예: 신고 #123, 제재 알림)
- `awaiting_user` 케이스는 상단 우선 노출한다.
- 사용자는 새 문의 생성 시 카테고리 선택, 설명 텍스트, 관련 거래/신고 선택, 첨부(선택) 정도만 입력하면 되도록 단순화한다.

#### 신고/분쟁/제재 상세
- 원사건 상세에서 `운영에 문의`, `추가 자료 제출`, `이의제기` CTA가 Support Case 생성/재진입으로 연결되어야 한다.
- 사용자는 원사건의 판단 상태와 Support Case의 응답 상태를 동시에 구분해 볼 수 있어야 한다.
  - 예: `신고 처리 완료 · 후속 문의 답변 대기 중`
  - 예: `제재 유지 · 추가 자료 요청됨`

### 101.8 운영 백오피스 요구사항
- 운영자는 Support Case를 단독 큐로도 보고, Report/Dispute/Restriction 상세 안의 linked panel로도 볼 수 있어야 한다.
- 케이스 목록 기본 컬럼:
  - caseId
  - supportCaseType
  - originSurface
  - linkedObject summary
  - caseSensitivityTier
  - supportCaseState
  - userExpectationState
  - lastInboundAt / lastOutboundAt
  - assignee
  - SLA due
- 운영 액션:
  - 상태 변경
  - 사용자에게 템플릿 답변 발송
  - 추가 자료 요청
  - Report/Dispute/Restriction linked action 열기
  - 민감도 상향/하향
  - handoff_external 처리
- 안전/개인정보 민감 케이스는 최소 role 또는 2인 승인 정책과 연결할 수 있어야 한다.

### 101.9 메시지/템플릿 정책
- Support Case 내부 메시지는 일반 거래 채팅과 분리된 톤과 포맷을 사용한다.
- 템플릿은 최소 아래 family를 가진다.
  - `ack_received`
  - `needs_more_info`
  - `under_review`
  - `resolved_with_action`
  - `resolved_no_action`
  - `appeal_rejected`
  - `case_closed_timeout`
- 템플릿은 변수 주입을 지원하되 민감 내부 근거 전체를 자동 삽입하지 않는다.
- 운영자는 템플릿 + 자유 메모 조합을 쓸 수 있으나, 자유 메모도 감사 대상으로 저장한다.

### 101.10 SLA / 타이머 규칙
| case type | 1차 응답 목표 | 사용자 자료 대기 기본 기한 | 자동 종료 후보 |
|---|---|---|---|
| `general_help` | 24시간 | 72시간 | 7일 무응답 |
| `trade_question` | 12시간 | 48시간 | 5일 무응답 |
| `report_followup` | 12시간 | 72시간 | 원사건 종료 후 7일 |
| `dispute_followup` | 6시간 | 48시간 | 판정 후 5일 |
| `restriction_appeal` | 24시간 | 72시간 | 결과 통지 후 7일 |
| `privacy` / `legal_risk` tier | 별도 운영정책 | 별도 운영정책 | 자동종료 지양 |

원칙:
- `awaiting_user` 상태에서 기한이 지나면 1회 리마인드 후 `resolved_no_action` 또는 `closed` 후보로 전환할 수 있다.
- `safety`, `privacy`, `legal_risk` tier는 일반 템플릿 자동 종료보다 수동 검토를 우선한다.

### 101.11 데이터 모델 후보
#### SupportCase
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `supportCaseId` | 필수 | 케이스 식별자 |
| `supportCaseType` | 필수 | 문의 유형 |
| `originSurface` | 필수 | 생성 출처 |
| `supportCaseState` | 필수 | 상태 |
| `userExpectationState` | 필수 | 사용자 기대 상태 |
| `caseSensitivityTier` | 필수 | 민감도 |
| `caseLinkPolicy` | 필수 | 원사건 관계 |
| `linkedReportId` | 선택 | 연결 신고 |
| `linkedDisputeId` | 선택 | 연결 분쟁 |
| `linkedRestrictionId` | 선택 | 연결 제재 |
| `linkedNoticeId` | 선택 | 연결 공지 |
| `createdByUserId` | 선택 | 사용자 생성 시 |
| `createdByAdminId` | 선택 | 운영 생성 시 |
| `assigneeAdminId` | 선택 | 담당자 |
| `firstResponseDueAt` | 선택 | 1차 응답 SLA |
| `waitingForUserUntil` | 선택 | 사용자 자료 대기 만료 |
| `resolvedAt` | 선택 | 해결 시각 |
| `closedAt` | 선택 | 종료 시각 |
| `resolutionCode` | 선택 | 결론 코드 |
| `resolutionSummary` | 선택 | 사용자 노출 요약 |
| `createdAt` | 필수 | 생성 시각 |
| `updatedAt` | 필수 | 갱신 시각 |

#### SupportCaseMessage
- `supportCaseMessageId`
- `supportCaseId`
- `senderType`: `user` / `admin` / `system`
- `senderUserId or senderAdminId`
- `messageType`: `free_text` / `template` / `attachment` / `status_update`
- `templateCode(optional)`
- `bodyText`
- `attachmentsJson(optional)`
- `visibilityScope`: `user_visible` / `staff_only`
- `sentAt`

원칙:
- staff-only 메모와 user-visible 메시지를 동일 테이블 또는 동일 aggregate 아래에서 구분 저장할 수 있어야 한다.
- 사용자에게 보낸 결론 요약은 케이스 header에 projection 캐시로 유지하는 것이 목록 UX에 유리하다.

### 101.12 API 후보
#### 사용자용
- `POST /support-cases`
- `GET /me/support-cases`
- `GET /support-cases/{supportCaseId}`
- `POST /support-cases/{supportCaseId}/messages`
- `POST /support-cases/{supportCaseId}/close`

#### 관리자용
- `GET /admin/support-cases`
- `GET /admin/support-cases/{supportCaseId}`
- `POST /admin/support-cases/{supportCaseId}/assign`
- `POST /admin/support-cases/{supportCaseId}/request-more-info`
- `POST /admin/support-cases/{supportCaseId}/resolve`
- `POST /admin/support-cases/{supportCaseId}/reopen`

Response projection 후보:
```json
{
  "supportCaseId": "sc_123",
  "supportCaseType": "restriction_appeal",
  "supportCaseState": "awaiting_user",
  "userExpectationState": "needs_more_info",
  "linkedObject": {
    "type": "restriction",
    "id": "rst_123",
    "summary": "채팅 기능 7일 제한"
  },
  "availableActions": ["send_message", "upload_attachment"],
  "responseSla": {
    "firstResponseDueAt": "2026-03-14T18:00:00+09:00",
    "waitingForUserUntil": "2026-03-17T18:00:00+09:00"
  }
}
```

### 101.13 analytics 이벤트 후보
- `support_case_created`
- `support_case_opened`
- `support_case_message_sent`
- `support_case_resolved`
- `support_case_closed`
- `support_case_reopened`
- `support_case_sla_breached`

공통 속성 후보:
- `supportCaseType`
- `originSurface`
- `caseSensitivityTier`
- `linkedObjectType`
- `resolutionCode`
- `assigneeRole`

핵심 분석 포인트:
- 신고/분쟁/제재별 후속 문의율
- 1차 응답 SLA 준수율
- `needs_more_info` 이후 사용자 회신율
- 템플릿 답변만으로 해결된 비율 vs 수동 장문 답변 비율
- Support Case 재오픈률과 CS 품질 상관관계

### 101.14 오픈 질문 / 가정
- 가정: MVP에서도 최소한 `report_followup`, `restriction_appeal`, `general_help` 3종은 Support Case로 구조화하는 것이 운영 효율상 유리하다.
- 오픈 질문: Support Case를 인앱 전용으로 둘지, 이메일/외부 채널과의 단방향 sync를 둘지 결정 필요
- 오픈 질문: 사용자가 `resolved` 결과에 대해 만족/불만족 피드백을 남기게 할지 여부 결정 필요
- 오픈 질문: bug report를 Support Case와 별도 product issue intake로 분리할지 결정 필요

## 102. 상대 사용자 관계(User Relationship Edge) / 차단·단골·위험 플래그 read model 계약
### 102.1 목표
- 거래 상대 간 관계를 단순 `차단 여부`나 `이전 채팅 존재 여부`로만 보지 않고, 실제 제품 행동을 결정하는 **관계 요약 객체**로 구조화한다.
- 매물 상세, 채팅, 내 거래, 프로필, 반복 거래, 신고/제재, 운영 백오피스가 모두 같은 관계 vocabulary를 사용하도록 한다.
- 사용자가 상대를 다시 만났을 때 `안전한 재거래`, `주의 필요`, `접촉 금지`, `추가 확인 필요`를 일관되게 해석하도록 CTA gating 기준을 명문화한다.

### 102.2 관계 모델이 필요한 이유
기존 객체만으로는 아래 문제가 생긴다.
1. `BlockRelation`, `Repeat Trade`, `Review`, `Restriction`, `Report`, `Dispute`가 각각 따로 존재해도 현재 상대와의 실제 관계 상태를 화면에서 즉시 판단하기 어렵다.
2. 같은 상대를 다시 만났을 때 어떤 CTA를 노출해야 하는지 화면마다 다르게 해석할 위험이 있다.
3. 운영자는 특정 사용자 자체의 위험도뿐 아니라 **두 사용자 사이에서 누적된 마찰/신뢰**를 함께 봐야 한다.

따라서 사용자 쌍(user pair) 기준의 canonical read model을 둔다.

### 102.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `relationshipState` | `new` / `known` / `repeat_trade` / `muted` / `blocked` / `restricted_pair` | 두 사용자 사이의 현재 관계 요약 |
| `safetyInteractionState` | `normal` / `caution` / `heightened_risk` / `contact_locked` | 안전 관점에서의 상호작용 상태 |
| `repeatTradeAffinityState` | `none` / `positive_history` / `preferred_counterparty` / `mixed_history` / `avoid_retrade` | 반복 거래 선호도 요약 |
| `contactPolicyState` | `contact_allowed` / `contact_limited` / `contact_requires_ack` / `contact_disallowed` | 신규 채팅/재거래 시작 가능 여부 |
| `relationshipReasonCode` | enum/string | 관계 상태를 만든 주된 근거 코드 |
| `pairCaseBurdenTier` | `none` / `light` / `medium` / `heavy` | 이 사용자 쌍 사이에 열린 사건/마찰의 밀도 |
| `visibilityScope` | `viewer_only` / `both_participants` / `staff_only` | 어떤 관계 신호를 누구에게 노출할지 |

원칙:
- 이 객체는 원천 데이터(source of truth)가 아니라 여러 원천을 해석한 **읽기 모델/projection**이다.
- 사용자에게는 단순 라벨/배너만 보여주고, 내부 reason code와 계산 근거는 운영/analytics용으로 유지한다.

### 102.4 관계 집계 기준
관계 요약은 최소 아래 원천을 합성해 계산한다.
- 과거 완료 거래 수
- 최근 거래 완료/취소/노쇼/분쟁 비율
- 양측 후기 결과와 숨김 여부
- 상호 차단/차단 해제 이력
- 열린 신고/분쟁/지원 케이스 존재 여부
- 운영 제한 중 상대방과의 접촉에 직접 영향 주는 제재 여부
- 반복 거래 시작/수락/거절 패턴

해석 원칙:
- 특정 사용자의 전역 신뢰도와 별개로, **특정 상대와의 관계만 나쁠 수 있음**을 허용한다.
- 단일 악성 이벤트 1건만으로 `avoid_retrade`를 자동 확정하지 않고, severity와 운영 판정 여부를 함께 본다.
- 쌍방 차단, 노쇼 상호 주장, 반복 분쟁은 `pairCaseBurdenTier`를 높이는 핵심 신호다.

### 102.5 상태 판정 가이드
#### `relationshipState`
- `new`: 과거 완료 거래/활성 사건/차단 이력 없음
- `known`: 채팅 또는 예약 이력은 있으나 repeat trade 수준은 아님
- `repeat_trade`: 완료 거래 또는 재거래 시작 이력이 누적됨
- `muted`: 관계는 유지되나 알림/노출만 약화된 상태
- `blocked`: 한쪽 또는 양쪽 차단이 활성
- `restricted_pair`: 운영/정책상 두 사용자 간 접촉을 제한해야 하는 상태

#### `safetyInteractionState`
- `normal`: 별도 경고 없이 상호작용 가능
- `caution`: 과거 취소/불발/후기 불일치 등 경미한 마찰 존재
- `heightened_risk`: 최근 분쟁, 노쇼 판정, 반복 신고, 고위험 경고 존재
- `contact_locked`: 차단, 사건 잠금, pair-level restriction으로 접촉 금지

### 102.6 화면별 해석 계약
#### 매물 상세
- 상대가 과거 긍정 거래 상대면 `다시 거래하기 좋은 상대` 수준의 viewer-only 신호를 노출할 수 있다.
- 상대가 `heightened_risk`이면 전면 차단 대신 `주의 배너 + 플랫폼 내 기록 유지 권장`을 우선한다.
- `contact_disallowed`면 `채팅 시작` CTA 대신 비활성 문구와 정책 안내를 보여준다.

#### 채팅 화면
- 헤더/상태 배너는 관계 요약을 가장 축약된 형태로 보여야 한다.
- 예시:
  - `이전 거래 3회 완료`
  - `최근 이 상대와 분쟁 기록 있음`
  - `차단 상태로 읽기 전용입니다`
- pair-level 사건이 열려 있으면 일반 후기/재거래 CTA보다 사건 해결 CTA를 우선한다.

#### 내 거래 / 보관함
- 반복 거래 선호 상대는 `preferred_counterparty` badge를 내부 정렬 신호로 활용할 수 있다.
- `avoid_retrade` 또는 `contact_requires_ack` 상태면 재거래 CTA를 숨기거나 경고-confirm step을 요구한다.

#### 프로필
- 관계 요약은 **viewer-personalized** 정보이므로 공개 프로필 공통 정보와 분리해야 한다.
- 같은 프로필이라도 보는 사람에 따라 `이전에 거래함`, `차단한 사용자`, `분쟁 있었음` 배지가 다를 수 있다.

### 102.7 CTA gating 규칙
| `contactPolicyState` | 허용 동작 | 제한 동작 |
|---|---|---|
| `contact_allowed` | 채팅 시작, 재거래 시작, 예약 제안 | 없음 |
| `contact_limited` | 채팅 시작 가능, 단 경고 노출 | 즉시 재거래 추천/자동 제안 노출 제한 |
| `contact_requires_ack` | 재접촉 전 confirm modal 필요 | 원탭 채팅 시작 금지 |
| `contact_disallowed` | 기록 조회, 신고, 이의제기만 허용 | 신규 채팅, 재거래, 예약 생성 금지 |

세부 규칙:
- 차단이 활성인 경우 `contact_disallowed`가 최우선이다.
- pair-level restriction은 전역 account restriction보다 더 좁은 범위의 예외 제한으로 동작할 수 있다.
- `preferred_counterparty`는 단순 정렬 신호이지, 정책 차단을 우회하는 권한이 아니다.

### 102.8 운영/백오피스 시사점
운영 화면은 단일 사용자 히스토리뿐 아니라 **pair timeline**을 조회할 수 있어야 한다.

필수 요약 필드 후보:
- `userAId`, `userBId`
- `lastTradeAt`
- `completedTradeCountBetween`
- `cancelledTradeCountBetween`
- `openCaseCountBetween`
- `activeBlockDirection`
- `relationshipState`
- `safetyInteractionState`
- `contactPolicyState`
- `topRelationshipReasonCodes[]`

운영 액션 후보:
- `set_pair_contact_restriction`
- `clear_pair_contact_restriction`
- `mark_preferred_counterparty_override` (예외적으로 금지 권장)
- `add_pair_watch_note`

원칙:
- pair-level override는 남용 위험이 있으므로 강한 감사 로그가 필요하다.
- 운영자는 개별 사용자 위험도와 pair-level 마찰을 혼동하지 않도록 두 축을 분리 표시해야 한다.

### 102.9 데이터 모델 후보
#### UserRelationshipEdge
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `relationshipEdgeId` | 필수 | 식별자 |
| `userLowId` | 필수 | 순서 정규화된 사용자 A |
| `userHighId` | 필수 | 순서 정규화된 사용자 B |
| `relationshipState` | 필수 | 관계 요약 상태 |
| `safetyInteractionState` | 필수 | 안전 상호작용 상태 |
| `repeatTradeAffinityState` | 필수 | 재거래 선호 상태 |
| `contactPolicyState` | 필수 | 접촉 정책 상태 |
| `pairCaseBurdenTier` | 필수 | 사건 부담도 |
| `lastTradeAt` | 선택 | 마지막 완료 거래 시각 |
| `lastIncidentAt` | 선택 | 마지막 분쟁/신고/차단 시각 |
| `computedAt` | 필수 | projection 계산 시각 |
| `version` | 필수 | projection 버전 |

#### UserRelationshipReason
- `relationshipEdgeId`
- `reasonCode`
- `sourceObjectType`
- `sourceObjectId`
- `weight(optional)`
- `createdAt`

원칙:
- 원천 로그를 덮어쓰지 않고, 관계 projection은 재계산 가능해야 한다.
- user pair 저장은 `(min(userId), max(userId))` 정규화로 중복을 방지한다.

### 102.10 API/read model 후보
#### 사용자용 read model
- `GET /users/{userId}/relationship`
- `GET /me/relationship-edges?state=repeat_trade`
- `GET /chats/{chatRoomId}` 응답 내 embedded relationship summary

응답 예시:
```json
{
  "counterpartyUserId": "usr_123",
  "relationship": {
    "relationshipState": "repeat_trade",
    "safetyInteractionState": "caution",
    "repeatTradeAffinityState": "positive_history",
    "contactPolicyState": "contact_allowed",
    "summaryLabel": "이전 거래 2회 완료",
    "viewerHints": [
      {
        "code": "KEEP_CONVERSATION_ON_PLATFORM",
        "severity": "info",
        "message": "이전 거래 이력이 있지만, 중요한 조율은 앱 안에서 남기는 것이 안전합니다."
      }
    ]
  }
}
```

#### 관리자용
- `GET /admin/user-relationship-edges?userId=...`
- `GET /admin/user-relationship-edges/{relationshipEdgeId}`
- `POST /admin/user-relationship-edges/{relationshipEdgeId}/restrict-contact`

### 102.11 analytics 이벤트 후보
- `relationship_badge_impression`
- `relationship_warning_shown`
- `relationship_blocked_cta_attempt`
- `repeat_trade_preference_used`
- `pair_contact_restriction_applied`
- `pair_contact_restriction_cleared`

공통 속성 후보:
- `relationshipState`
- `safetyInteractionState`
- `contactPolicyState`
- `pairCaseBurdenTier`
- `sourceScreen`

핵심 분석 포인트:
- `positive_history` 관계에서의 재거래 전환율
- `caution` 배너 노출이 분쟁률/노쇼율에 미치는 영향
- 차단 직전/직후 재접촉 시도 빈도
- pair-level restriction 도입 후 신고 재발률 변화

### 102.12 오픈 질문 / 가정
- 가정: MVP에서는 `UserRelationshipEdge`를 강한 정책 엔진이 아니라 read model + CTA gating 보조 레이어로 도입하는 것이 현실적이다.
- 오픈 질문: `preferred_counterparty`를 사용자가 직접 지정할 수 있게 할지, 완료 거래 기반 자동 추론만 할지 결정 필요
- 오픈 질문: 과거 분쟁이 해소된 뒤 `avoid_retrade`를 언제/어떻게 완화할지 정책 결정 필요
- 오픈 질문: pair-level restriction을 운영 수동 조치만 허용할지, 특정 고위험 사건 판정 시 자동 후보로 올릴지 결정 필요

## 103. 거래 상대 식별 정보(Trade Counterparty Identity Handshake) / 캐릭터명·접선 식별자 공개/확인 계약
### 103.1 목표
- 실제 거래 직전에 "누구를 만나야 하는지"가 불분명해 발생하는 혼선, 사칭, 잘못된 거래 완료, 노쇼 오판을 줄인다.
- 전화번호/실명 같은 과도한 개인정보 없이도 인게임/오프라인 거래에서 상대를 식별할 수 있는 최소 정보 교환 규칙을 구조화한다.
- 매물/채팅/예약/당일 실행 카드/분쟁/운영도구가 동일한 식별 vocabulary를 공유하도록 한다.

### 103.2 적용 범위
이 계약은 아래 상황에 적용한다.
- 인게임 거래 시 캐릭터명/길드명/접속 채널 등 상대 식별 정보 교환
- 오프라인 거래 시 닉네임 기반 접선 확인 문구, 도착 확인 키워드, 착장/소지품처럼 민감하지 않은 식별 힌트 교환
- 예약 확정 이후 실제 실행 직전 `누가 상대방인지`를 상호 확인하는 단계
- 노쇼/사칭/장소 혼선 분쟁 시 마지막으로 합의된 식별 힌트 재구성

적용 제외:
- 실명, 전화번호, 계좌번호 등 직접 식별 개인정보는 본 계약의 기본 식별 수단으로 보지 않는다.
- 외부 메신저 ID 교환은 별도 외부 연락처 정책을 따른다.

### 103.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `counterpartyIdentifierType` | `character_name` / `character_name_with_server` / `guild_name` / `arrival_keyword` / `appearance_hint` / `pickup_phrase` | 실제 만남 시 상대를 인지하기 위한 식별 수단 |
| `identityDisclosureLevel` | `hidden` / `reservation_only` / `confirmed_only` / `day_of_trade_only` / `post_trade_masked` | 어느 단계에서 공개되는지 |
| `identifierVerificationState` | `not_shared` / `shared_by_one_side` / `shared_by_both_sides` / `acknowledged` / `mismatch_reported` / `withdrawn` | 식별 정보 공유/확인 상태 |
| `tradeRecognitionMethod` | `mutual_character_name_match` / `reservation_code_word` / `arrival_chat_confirmation` / `manual_visual_confirmation` | 어떤 방식으로 상대를 확인하는지 |
| `identifierSensitivityTier` | `low` / `medium` / `restricted` | 민감도에 따른 노출/보관 제어 |
| `identityMismatchReasonCode` | `name_changed` / `server_mismatch` / `wrong_person_contacted` / `identifier_not_visible` / `counterparty_denied` | 식별 불일치/혼선 사유 |

원칙:
- 식별 정보는 예약의 부속 텍스트가 아니라 구조화 가능한 `handshake` 단위로 다뤄야 한다.
- 동일 예약에서 마지막으로 상호 인지된 식별 정보와, 이후 변경 요청 이력을 분리 저장해야 한다.

### 103.4 공개 단계 원칙
| 단계 | 허용 식별 정보 | 노출 수준 |
|---|---|---|
| 매물 공개 | 없음 또는 선호 방식만 | `hidden` |
| 일반 채팅(open) | 필요 시 캐릭터명 공유 유도 가능하나 기본 비공개 | `reservation_only` 후보 |
| 예약 제안(proposed) | 인게임 캐릭터명 또는 오프라인 접선용 짧은 확인 문구 제안 가능 | `confirmed_only` |
| 예약 확정(confirmed) | 실제 사용할 식별 수단 1개 이상 확정 | `confirmed_only` |
| 당일 실행(trade_due) | 도착 확인용 키워드/캐릭터명/서버 조합 노출 | `day_of_trade_only` |
| 거래 종료 후 | 분쟁 대응 최소 정보만 마스킹 보관 | `post_trade_masked` |

세부 원칙:
- 매물 본문/공개 프로필에는 상대 식별용 캐릭터명 전체를 기본 노출하지 않는다.
- 인게임 거래는 `character_name_with_server`를 권장 기본안으로 둔다.
- 오프라인 거래는 실명/연락처 대신 `arrival_keyword` 또는 비민감 `appearance_hint`를 우선 권장한다.

### 103.5 식별 정보 handshake 생성 규칙
- `reservationStatus=confirmed` 이전에는 식별 정보 handshake를 선택 입력으로 두되, 확정 시점에는 최소 1개의 `tradeRecognitionMethod`가 결정되어야 한다.
- `meetingType=in_game`이면 아래 중 하나가 필수다.
  1. `character_name_with_server`
  2. `character_name` + `arrival_chat_confirmation`
- `meetingType=offline_pc_bang`이면 아래 중 하나가 필수다.
  1. `arrival_keyword`
  2. `appearance_hint`
  3. `arrival_chat_confirmation`
- 양측이 서로 다른 식별 수단을 제시해도 되지만, 최종적으로는 하나의 `acknowledged` 가능한 공통 확인 방식이 있어야 한다.

### 103.6 변경 / 철회 / 재확인 규칙
- 캐릭터명 변경, 접속 캐릭터 교체, 오프라인 접선 힌트 수정은 append-only 변경 이력으로 저장한다.
- 예약 시각이 임박한 상태에서 식별 정보가 바뀌면 `identifierVerificationState`를 다시 `shared_by_one_side` 또는 `shared_by_both_sides`로 낮추고 상대 재확인을 요구한다.
- 거래 시각 직전(가정: 30분 이내) 식별 정보 변경은 `identityMismatchRisk=true` 신호를 남겨 노쇼/사칭 분쟁 판단 참고값으로 사용한다.
- 사용자가 식별 정보를 철회하면(`withdrawn`) 기존 예약은 유지될 수 있으나 당일 실행 카드에는 `상대 확인 재필요` 경고를 띄워야 한다.

### 103.7 화면 요구사항
#### 채팅/예약 카드
- 예약 카드 또는 별도 handshake 카드에 아래 항목을 보여야 한다.
  - 현재 식별 방식
  - 마지막 업데이트 시각
  - 상대 확인 상태
  - 변경/재확인 필요 배지
- CTA 후보:
  - `캐릭터명 공유`
  - `도착 확인 문구 정하기`
  - `식별 정보 확인 완료`
  - `식별 정보 다시 보내기`
  - `상대가 달라요`

#### 당일 실행 카드
- `도착했어요` CTA와 별도로 `상대 확인됨`, `상대가 안 보여요`, `식별 정보 다시 보기` CTA가 필요하다.
- 인게임 거래는 `서버 + 캐릭터명`, 오프라인 거래는 `확인 키워드 + 위치 요약`이 한 줄로 보여야 한다.
- 사용자가 `상대가 달라요`를 누르면 일반 no-show가 아니라 `mismatch_reported` 경로로 분기할 수 있어야 한다.

#### 내 거래 화면
- 임박 거래 카드에 `식별 정보 확인 완료 여부`를 표시해야 한다.
- `acknowledged` 전 상태의 거래는 같은 시각의 일반 메시지 미읽음보다 우선순위를 높게 둘 수 있다.

### 103.8 노쇼 / 사칭 / 분쟁과의 관계
- `no_show_claimed`와 `mismatch_reported`는 같은 사건으로 자동 합치지 않고, 동일 예약에 연결된 별도 판단 축으로 저장하는 것이 바람직하다.
- 한쪽이 도착했지만 식별 불일치 때문에 거래가 무산된 경우, 운영은 아래를 함께 재구성할 수 있어야 한다.
  1. 마지막 `acknowledged` 식별 정보
  2. 직전 변경 시각
  3. 도착 확인 메시지/시간
  4. `identityMismatchReasonCode`
- `mismatch_reported`가 제출되면 day-of-trade 카드의 기본 CTA는 `노쇼 신고`보다 `상대 식별 불일치 신고`를 우선 노출하는 안을 권장한다.

### 103.9 데이터 모델 후보
#### `TradeIdentityHandshake`
| 컬럼 | 필수 여부 | 설명 |
|---|---|---|
| `handshakeId` | 필수 | 식별 handshake 식별자 |
| `reservationId` | 필수 | 연결 예약 |
| `listingId` | 필수 | 연결 매물 |
| `chatRoomId` | 필수 | 연결 채팅 |
| `ownerUserId` | 필수 | 해당 식별 정보를 제공한 사용자 |
| `counterpartyIdentifierType` | 필수 | 식별 유형 |
| `identifierValueEncrypted` | 조건부 필수 | 원문 또는 암호화 저장값 |
| `identifierDisplayMasked` | 필수 | 사용자 노출용 마스킹 텍스트 |
| `identityDisclosureLevel` | 필수 | 공개 레벨 |
| `identifierVerificationState` | 필수 | 현재 검증 상태 |
| `tradeRecognitionMethod` | 필수 | 확인 방식 |
| `identifierSensitivityTier` | 필수 | 민감도 등급 |
| `acknowledgedAt` | 선택 | 상대 확인 시각 |
| `withdrawnAt` | 선택 | 철회 시각 |
| `lastChangedAt` | 필수 | 마지막 변경 시각 |

#### `TradeIdentityMismatchEvent`
- `mismatchEventId`
- `reservationId`
- `reportedByUserId`
- `identityMismatchReasonCode`
- `detailsText(optional)`
- `reportedAt`
- `resolvedDisposition(optional)`

원칙:
- 원문 식별값은 검색 인덱스, 푸시 본문, 공개 프로필, analytics raw payload에 직접 노출하지 않는다.
- 오프라인 힌트는 최소 보관/마스킹 원칙을 적용한다.

### 103.10 API 후보
#### 사용자용
- `POST /reservations/{reservationId}/identity-handshakes`
- `POST /identity-handshakes/{handshakeId}/acknowledge`
- `PATCH /identity-handshakes/{handshakeId}`
- `POST /reservations/{reservationId}/identity-mismatches`
- `GET /reservations/{reservationId}/identity-handshakes`

응답 예시:
```json
{
  "reservationId": "res_123",
  "identityHandshake": {
    "counterpartyIdentifierType": "character_name_with_server",
    "identifierDisplayMasked": "켄트섭 린***",
    "identifierVerificationState": "acknowledged",
    "tradeRecognitionMethod": "mutual_character_name_match",
    "availableActions": ["view_identity_hint", "report_identity_mismatch"]
  }
}
```

#### 관리자용
- `GET /admin/reservations/{reservationId}/identity-handshakes`
- `GET /admin/identity-mismatches`
- `POST /admin/identity-mismatches/{mismatchEventId}/resolve`

### 103.11 analytics 이벤트 후보
- `identity_handshake_created`
- `identity_handshake_acknowledged`
- `identity_handshake_changed`
- `identity_hint_viewed`
- `identity_mismatch_reported`
- `identity_mismatch_resolved`

공통 속성 후보:
- `meetingType`
- `counterpartyIdentifierType`
- `identifierVerificationState`
- `tradeRecognitionMethod`
- `identityDisclosureLevel`
- `sourceScreen`

핵심 분석 포인트:
- `acknowledged` handshake가 있는 예약의 완료율 vs 없는 예약의 완료율
- 인게임/오프라인 거래별 식별 불일치 신고율
- 당일 직전 식별 변경이 노쇼/분쟁률에 미치는 영향

### 103.12 오픈 질문 / 가정
- 가정: MVP에서는 얼굴사진/실명 인증보다 `캐릭터명 + 도착 확인 문구` 같은 저마찰 handshake가 현실적이다.
- 오픈 질문: 오프라인 거래에서 `appearance_hint`를 어디까지 허용할지(의상/가방 수준 vs 더 상세 묘사 금지) 정책 결정 필요
- 오픈 질문: 거래 종료 후 식별 handshake 원문 보관 기간을 evidence retention과 완전히 동일하게 둘지, 더 짧게 둘지 결정 필요
- 오픈 질문: `mismatch_reported`를 no-show 사건과 자동 연결할지, 운영 수동 병합만 허용할지 결정 필요

## 104. 거래 조건 확정(Deal Terms Snapshot) / 최종 합의 조건 / 변경 이력 계약
### 104.1 목적
- 실제 거래는 매물 원문만으로 성사되지 않고, 가격 제안, 부분 수량, 예약 시간/장소, 상대 식별 정보, 대금 전달 방식이 마지막까지 조정되므로 `최종 합의 조건`을 별도 canonical 객체로 다뤄야 한다.
- 채팅/예약/당일 실행/완료/노쇼/분쟁/후기가 서로 다른 조건 세트를 참조하면 “무엇을 기준으로 약속이 성사됐는지” 해석이 갈리므로, 실행 직전 기준 스냅샷을 명문화한다.
- 매물 원문 수정, offer 수락, 재일정, 식별 handshake 변경, 지급 방식 경고가 모두 같은 `deal terms` 계층에 반영되도록 제품/백엔드/운영/분석 vocabulary를 통일한다.

### 104.2 canonical 개념 정의
- **Deal Terms Snapshot**: 특정 거래 스레드에서 현재 유효한 최종 합의 조건의 읽기/감사 기준 객체
- **Base Terms**: 매물 원문에서 파생된 최초 거래 조건(가격, 수량, 서버, 거래 방식 등)
- **Negotiated Terms**: offer, 채팅 합의, 예약 제안으로 조정된 조건
- **Locked Terms**: 당일 실행/완료/분쟁 판단의 기준이 되는 확정 조건
- **Terms Revision**: 최종 조건이 바뀔 때마다 남기는 변경 이력 단위

### 104.3 핵심 vocabulary
#### 상태/출처
- `dealTermsState`: `draft` | `proposed` | `partially_acknowledged` | `locked` | `superseded` | `voided`
- `termsSourceType`: `listing_base` | `offer_acceptance` | `reservation_confirmation` | `manual_chat_structured_update` | `admin_resolution`
- `termsChangeImpact`: `cosmetic_only` | `informational` | `requires_ack` | `execution_blocking`
- `executionReadinessState`: `not_ready` | `ready_with_pending_ack` | `ready` | `blocked`

#### 조건 범주
- `priceAgreementState`: `listing_price` | `offer_pending` | `offer_accepted` | `fixed_final` | `disputed`
- `quantityAgreementState`: `full_quantity` | `partial_quantity` | `residual_split_pending` | `quantity_disputed`
- `meetingAgreementState`: `not_set` | `proposed` | `confirmed` | `changed_after_confirmation`
- `counterpartyRecognitionState`: `not_shared` | `shared` | `acknowledged` | `mismatch_reported`
- `settlementAgreementState`: `not_set` | `cash_on_meet` | `in_game_exchange_only` | `bank_transfer_risky` | `mixed_or_custom`

### 104.4 Deal Terms Snapshot 필수 필드 후보
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `dealTermsSnapshotId` | 필수 | 스냅샷 식별자 |
| `tradeThreadId` | 필수 | 연결된 trade thread |
| `listingId` | 필수 | 원 매물 |
| `chatRoomId` | 필수 | 실제 협의 채팅방 |
| `activeReservationId` | 선택 | 예약 확정 시 연결 |
| `snapshotVersion` | 필수 | thread 내 증가 버전 |
| `dealTermsState` | 필수 | 현재 합의 상태 |
| `executionReadinessState` | 필수 | 실행 가능 readiness |
| `termsSourceType` | 필수 | 이번 스냅샷의 주 생성 출처 |
| `finalPriceAmount` | 선택 | 최종 합의 가격 |
| `priceCurrencyType` | 선택 | 가격 통화/단위 |
| `finalQuantity` | 선택 | 최종 거래 수량 |
| `residualQuantityAfterTrade` | 선택 | 부분거래 후 잔여 수량 |
| `meetingType` | 선택 | 실제 거래 방식 |
| `meetingSnapshotRefId` | 선택 | 장소/시간 snapshot 참조 |
| `counterpartyIdentifierSnapshotRefId` | 선택 | 식별 handshake 참조 |
| `settlementMethod` | 선택 | 대금 전달 방식 |
| `specialTermsText` | 선택 | 구조화 어려운 부가 조건 요약 |
| `termsAcknowledgedBySellerAt` | 선택 | 판매 역할 사용자 확인 시각 |
| `termsAcknowledgedByBuyerAt` | 선택 | 구매 역할 사용자 확인 시각 |
| `lockedAt` | 선택 | locked 전환 시각 |
| `supersededBySnapshotId` | 선택 | 후속 revision 연결 |
| `createdByActorType` | 필수 | `user` / `system` / `admin` |
| `createdByActorId` | 선택 | 생성 주체 |
| `createdAt` | 필수 | 생성 시각 |

원칙:
- `Deal Terms Snapshot`은 매물 원문을 덮어쓰지 않고 thread 문맥의 실행 조건을 보존하는 별도 계층이다.
- 동일 thread에는 `locked` 또는 `ready` 성격의 활성 snapshot은 최대 1개만 허용하고, 변경 시 이전 snapshot은 `superseded` 처리한다.

### 104.5 조건 확정 lifecycle
1. **listing_base 생성**: 채팅 시작 시점에 매물 원문 기반 base snapshot 생성 가능
2. **offer/채팅 조정**: 가격, 수량, 지급 방식 조정이 생기면 `proposed` revision 생성
3. **예약 확정 연동**: 시간/장소/방식이 확정되면 `meetingAgreementState=confirmed` 반영
4. **식별 handshake 연동**: 상대 캐릭터명/식별자 acknowledgement가 붙으면 `counterpartyRecognitionState=acknowledged`
5. **양측 확인**: 가격/수량/시간/장소/식별/지급방식이 모두 실행 가능한 수준이면 `executionReadinessState=ready`
6. **locked 전환**: 거래 직전 또는 완료 요청 직전 기준 snapshot을 `locked`로 고정
7. **완료/분쟁 참조**: 완료 확인, no-show, mismatch, payment risk, dispute 판단은 가장 최근 `locked` snapshot을 우선 참조

### 104.6 어떤 변경이 재확인(`requires_ack`)을 강제하는가
아래 변경은 `termsChangeImpact=requires_ack` 또는 `execution_blocking`으로 간주한다.

| 변경 항목 | 기본 impact | 재확인 필요 여부 | 비고 |
|---|---|---|---|
| 최종 가격 변경 | `requires_ack` | 필수 | offer 수락 후 재변경 포함 |
| 최종 수량 변경 | `requires_ack` | 필수 | 부분거래/묶음 분리 포함 |
| 예약 시간 변경 | `requires_ack` | 필수 | reschedule 규칙과 연결 |
| 예약 장소/방식 변경 | `requires_ack` | 필수 | 특히 당일 직전 변경 주의 |
| 캐릭터명/식별자 변경 | `execution_blocking` | 필수 | 사칭/혼선 리스크 큼 |
| settlementMethod 변경 | `execution_blocking` | 필수 | 계좌이체 유도 등 위험 신호 |
| specialTermsText 오탈자 수정 | `cosmetic_only` | 선택 아님 | 의미 불변일 때만 |
| 설명 보강(예: 도착 후 채팅 주세요) | `informational` | 조건부 | 실행 해석 바뀌면 requires_ack |

원칙:
- `execution_blocking` 변경이 발생하면 당일 실행 카드, 내 거래, 채팅 상단에 `재확인 필요` 배지를 강제 노출한다.
- 상대가 마지막 revision을 보지 못했거나 ack하지 않았다면 `executionReadinessState=ready_with_pending_ack` 또는 `blocked`로 유지한다.

### 104.7 화면 surface 요구사항
#### 채팅 화면
- 채팅 상단 또는 예약 카드 안에 `현재 최종 조건` 요약 카드가 보여야 한다.
- 필수 요약 항목: 가격, 수량, 시간, 장소 요약, 식별 방식, 지급 방식, 마지막 변경 시각, 양측 확인 상태
- 변경 diff가 있을 때는 `무엇이 바뀌었는지`를 직전 snapshot 대비 하이라이트해야 한다.
- 구조화된 조건 변경은 일반 텍스트 메시지로만 남기지 않고 시스템 카드/시스템 이벤트로도 남겨야 한다.

#### 내 거래 화면
- thread 카드에 `실행 준비 완료`, `상대 확인 대기`, `조건 다시 확인 필요`, `지급 방식 경고` 같은 상태 라벨을 노출해야 한다.
- 액션 필요 정렬에서 unread보다 `executionReadinessState`와 `termsChangeImpact`가 우선할 수 있어야 한다.

#### 당일 실행 카드
- `locked` snapshot 기준으로 `오늘의 거래 조건`을 1-screen 안에 확인 가능해야 한다.
- 사용자는 도착/지연/완료/노쇼 처리 전에 “내가 지금 어떤 조건으로 약속한 것인지”를 다시 볼 수 있어야 한다.
- `bank_transfer_risky`, `changed_after_confirmation`, `mismatch_reported` 조합이면 안전 경고 copy를 강화한다.

#### 완료/분쟁 화면
- 완료 확인 화면에는 `이 조건으로 거래가 끝났나요?` 문구와 함께 locked snapshot 요약을 보여야 한다.
- 분쟁/노쇼 화면에는 사용자가 주장하는 내용과 locked snapshot의 차이를 비교 가능하게 해야 한다.

### 104.8 API/read model 파생 기준
#### 사용자용 read model 후보
- `GET /me/trades/{tradeThreadId}` 응답에 `activeDealTermsSnapshot` 포함
- `GET /chats/{chatRoomId}` 응답에 `currentDealTermsSummary` 포함
- `GET /reservations/{reservationId}` 응답에 `dealTermsAckSummary` 포함

응답 예시:
```json
{
  "activeDealTermsSnapshot": {
    "dealTermsSnapshotId": "dts_123",
    "snapshotVersion": 4,
    "dealTermsState": "locked",
    "executionReadinessState": "ready",
    "finalPriceAmount": 120000,
    "finalQuantity": 3,
    "meetingAgreementState": "confirmed",
    "counterpartyRecognitionState": "acknowledged",
    "settlementAgreementState": "cash_on_meet",
    "lastChangedAt": "2026-03-14T05:58:00+09:00",
    "sellerAck": true,
    "buyerAck": true,
    "availableActions": ["view_terms_diff", "report_terms_mismatch"]
  }
}
```

#### command surface 후보
- `POST /trade-threads/{tradeThreadId}/deal-terms/propose`
- `POST /deal-terms/{dealTermsSnapshotId}/acknowledge`
- `POST /deal-terms/{dealTermsSnapshotId}/lock`
- `POST /deal-terms/{dealTermsSnapshotId}/report-mismatch`

원칙:
- offer 수락, reschedule confirm, identity handshake ack, settlement method 경고 수락이 발생하면 내부적으로는 동일 `deal terms revision` 계층을 갱신하거나 생성한다.
- `lock`은 예약 확정 직후 자동 생성될 수 있으나, 당일 직전 변경이 있으면 새 snapshot을 다시 lock해야 한다.

### 104.9 DB / 정합성 / 감사 기준
- `trade_thread_id + deal_terms_state in ('locked')` 활성 uniqueness 제약 후보
- `snapshotVersion`은 thread 단위 단조 증가
- `termsAcknowledgedBySellerAt`, `termsAcknowledgedByBuyerAt` 둘 다 비어 있고 `locked`인 상태는 허용하지 않는 기본안을 우선 검토
- `meetingSnapshotRefId`, `counterpartyIdentifierSnapshotRefId`, `activeReservationId`는 모두 같은 thread 경계 안에 있어야 한다.
- offer/예약/식별/지급 방식 변경이 기존 snapshot을 바꿀 때는 update-in-place보다 revision append를 우선한다.
- 완료/노쇼/분쟁 사건에는 `lockedDealTermsSnapshotId`를 FK로 남겨 사후 판정 기준을 고정한다.

### 104.10 운영 정책 / 분쟁 판단 기준
- 운영자는 `매물 원문`이 아니라 사건 시점의 `lockedDealTermsSnapshot`을 우선 기준으로 판정한다.
- 아래 사건은 deal terms mismatch로 구조화할 수 있어야 한다.
  - 가격을 다르게 요구함
  - 일부 수량만 주겠다고 현장 변경함
  - 장소/시간을 마지막에 바꾸고 상대 ack 없음
  - 약속한 캐릭터명/식별자가 달라 상대가 거래 상대를 찾지 못함
  - 지급 방식이 현장 직전에 위험 방식으로 변경됨
- `report_terms_mismatch`는 no-show, payment risk, identity mismatch, completion dispute와 연결될 수 있으나 동일 사실관계를 여러 사건으로 중복 집계하지 않도록 linkage 정책이 필요하다.

### 104.11 analytics 이벤트 후보
- `deal_terms_snapshot_created`
- `deal_terms_acknowledged`
- `deal_terms_locked`
- `deal_terms_changed_after_lock`
- `deal_terms_mismatch_reported`
- `execution_readiness_changed`

공통 속성 후보:
- `termsSourceType`
- `dealTermsState`
- `executionReadinessState`
- `priceAgreementState`
- `quantityAgreementState`
- `meetingAgreementState`
- `counterpartyRecognitionState`
- `settlementAgreementState`
- `snapshotVersion`

핵심 분석 포인트:
- `locked + 양측 ack` snapshot이 있는 거래와 없는 거래의 완료율 차이
- 조건 변경 횟수가 많은 거래의 노쇼/분쟁률
- settlement method 변경이 완료율·신고율에 미치는 영향
- 부분거래/가격 재협상 후 실제 완료까지 이어지는 비율

### 104.12 오픈 질문 / 가정
- 가정: MVP에서는 모든 텍스트 협의를 100% 구조화하지 못하므로, 가격/수량/시간/장소/식별/지급방식 같은 분쟁 핵심 항목만 우선 구조화한다.
- 오픈 질문: `lock`을 명시 CTA로 둘지, 예약 확정 + 양측 ack 시 자동 생성할지 결정 필요
- 오픈 질문: `specialTermsText`를 어디까지 허용할지(자유도 vs 운영 해석 가능성) 제한 필요
- 오픈 질문: 양측 ack 없이도 `locked` 생성이 가능한 예외(예: 시스템 자동 이전, 운영 강제 판정)를 어떤 범위까지 허용할지 결정 필요

## 105. 거래 인도/수령 확인(Exchange Confirmation) / 실제 물건 전달·수령 확인 계약
### 105.1 목적
- `거래완료`는 단순 버튼 클릭이 아니라, 직전까지 확정된 거래 조건이 실제로 `전달됨/받음/일부만 전달됨/받지 못함` 중 무엇으로 귀결됐는지를 확인하는 단계여야 한다.
- no-show, 조건 불일치, 지급 방식 변경, 부분거래는 `약속 장소 도착`만으로 해결되지 않으므로, 현장에서 실제 인도/수령 여부를 별도 canonical 레이어로 구조화해야 한다.
- 채팅/당일 실행 카드/완료 요청/분쟁/후기/운영이 같은 기준으로 “무엇이 건네졌고 무엇이 확인되지 않았는지”를 해석하도록 vocabulary를 통일한다.

### 105.2 canonical 개념 정의
- **Exchange Confirmation**: 특정 거래 스레드에서 실제 전달·수령 결과를 기록하는 canonical 객체
- **Handover**: 판매 역할 사용자가 아이템/재화를 전달했다고 주장하거나 확인한 행위
- **Receipt**: 구매 역할 사용자가 아이템/재화를 받았다고 주장하거나 확인한 행위
- **Exchange Mismatch**: 약속된 조건과 실제 전달 결과가 불일치하는 사건(수량 부족, 다른 아이템, 현장 조건 변경 등)
- **Completion Eligibility**: 완료 요청이 열릴 만큼 인도/수령 사실이 충분히 정리되었는지 여부

### 105.3 핵심 vocabulary
#### 전달/수령 상태
- `handoverState`: `not_started` | `seller_marked_ready` | `seller_marked_handed_over` | `partial_handover_claimed` | `handover_disputed`
- `receiptState`: `not_started` | `buyer_marked_ready` | `buyer_marked_received` | `partial_receipt_claimed` | `receipt_disputed`
- `exchangeConfirmationState`: `none` | `one_sided_claim` | `mutually_confirmed` | `partially_confirmed` | `mismatch_reported` | `resolved`
- `completionEligibilityState`: `not_eligible` | `eligible_with_warning` | `eligible` | `blocked_by_mismatch`

#### 불일치 유형
- `exchangeMismatchType`: `item_not_received` | `quantity_shortfall` | `wrong_item_or_option` | `price_changed_on_site` | `settlement_not_completed` | `counterparty_left_after_meet` | `other_execution_mismatch`
- `exchangeEvidenceType`: `chat_message` | `system_checkin` | `photo_attachment` | `structured_statement` | `admin_note`
- `exchangeResolutionType`: `completed_as_agreed` | `completed_partially` | `not_completed_after_meet` | `policy_violation_detected` | `insufficient_evidence`

### 105.4 Exchange Confirmation 필수 필드 후보
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `exchangeConfirmationId` | 필수 | 확인 객체 식별자 |
| `tradeThreadId` | 필수 | 연결된 trade thread |
| `listingId` | 필수 | 원 매물 |
| `chatRoomId` | 필수 | 실제 거래 채팅방 |
| `reservationId` | 선택 | 연결 예약 |
| `lockedDealTermsSnapshotId` | 필수 | 판단 기준이 되는 locked deal terms |
| `handoverState` | 필수 | 전달 상태 |
| `receiptState` | 필수 | 수령 상태 |
| `exchangeConfirmationState` | 필수 | 상호 확인 상태 |
| `completionEligibilityState` | 필수 | 완료 요청 가능 여부 |
| `sellerConfirmedQuantity` | 선택 | 판매자 기준 실제 전달 수량 |
| `buyerConfirmedQuantity` | 선택 | 구매자 기준 실제 수령 수량 |
| `sellerConfirmationAt` | 선택 | 판매자 현장 확인 시각 |
| `buyerConfirmationAt` | 선택 | 구매자 현장 확인 시각 |
| `exchangeMismatchType` | 선택 | 대표 불일치 유형 |
| `exchangeMismatchSummaryText` | 선택 | 사용자 요약 텍스트 |
| `resolutionType` | 선택 | 운영 또는 상호 정리 결과 |
| `resolvedAt` | 선택 | 해결 시각 |
| `createdAt` | 필수 | 생성 시각 |
| `updatedAt` | 필수 | 마지막 갱신 시각 |

원칙:
- `Exchange Confirmation`은 `TradeCompletion`을 대체하지 않고, 완료 요청의 사실 입력 계층으로 동작한다.
- 같은 thread에서 활성 `exchangeConfirmationState in ('one_sided_claim','partially_confirmed','mismatch_reported')` 객체는 최대 1건만 허용하는 기본안을 둔다.

### 105.5 lifecycle / 상태 전이
1. **당일 실행 진입**: reservation + locked deal terms가 준비되면 exchange confirmation 객체 생성 가능
2. **현장 준비 표시**: 판매자/구매자가 `ready` 또는 `도착`에 준하는 행동을 남김
3. **전달/수령 주장 생성**: 판매 역할이 `seller_marked_handed_over`, 구매 역할이 `buyer_marked_received`를 남길 수 있음
4. **상호 확인**: 양측이 동일 조건 기준으로 전달/수령을 인정하면 `mutually_confirmed`
5. **부분거래 반영**: 수량 일부만 주고받았으면 `partially_confirmed` + quantity 차이 기록
6. **불일치 신고**: 조건 불일치나 대금 미정산이면 `mismatch_reported`
7. **완료 요청 가능**: `mutually_confirmed` 또는 허용된 `partially_confirmed`에서만 completion CTA 활성화
8. **분쟁/운영 판정**: mismatch가 해소되면 `resolved`와 함께 완료/불발/부분완료 방향 결정

### 105.6 완료 요청과의 관계
| Exchange 상태 | 기본 완료 요청 허용 | 비고 |
|---|---|---|
| `none` | 불가 | 단순 도착만으로는 완료 불가 |
| `one_sided_claim` | 조건부 | 반대측 확인 전 경고 필요 |
| `mutually_confirmed` | 가능 | 기본 완료 요청 경로 |
| `partially_confirmed` | 가능(경고 포함) | partial completion 규칙 연결 |
| `mismatch_reported` | 불가 | 먼저 mismatch/dispute 정리 필요 |
| `resolved` + `completed_as_agreed` | 가능 또는 이미 완료 처리 | 운영 결과에 따름 |
| `resolved` + `not_completed_after_meet` | 불가 | 취소/분쟁 경로 |

원칙:
- `TradeCompletion` 생성 시 `exchangeConfirmationId`를 참조해, 단순 완료 클릭이 아니라 어떤 전달/수령 사실을 기준으로 했는지 고정한다.
- 양측 모두 현장 확인을 남기지 않았더라도, MVP에서는 한쪽 완료 요청을 완전히 막지 않을 수 있다. 다만 `eligible_with_warning`으로 분류해 상대 확인과 운영 추적성을 강화한다.

### 105.7 부분거래 / 현장 변경과의 관계
- `sellerConfirmedQuantity != buyerConfirmedQuantity`이면 기본적으로 `partially_confirmed` 또는 `mismatch_reported` 후보로 본다.
- locked deal terms의 `finalQuantity`보다 적은 수량이 상호 확인되면 `resolutionType=completed_partially`와 함께 residual quantity 정책을 호출해야 한다.
- 현장에서 가격/지급 방식이 바뀌면 `exchangeMismatchType=price_changed_on_site` 또는 `settlement_not_completed`를 우선 기록하고, 바로 완료로 닫지 않는 보수안을 기본으로 둔다.
- 식별 mismatch(`wrong character`, `wrong person`)는 104장 `deal terms mismatch`와 연결하되, 실제 접선 후 상대가 떠난 경우는 `counterparty_left_after_meet`로 별도 분류할 수 있어야 한다.

### 105.8 화면 surface 요구사항
#### 채팅 / 당일 실행 카드
- `도착했어요`, `전달했어요`, `받았어요`, `일부만 받았어요`, `약속과 달라요` CTA를 분리해야 한다.
- 사용자는 `완료`를 누르기 전에 “전달/수령 확인” 단계를 거친다는 것을 이해할 수 있어야 한다.
- 카드에는 최소 다음 정보가 필요하다.
  - locked deal terms 요약
  - 판매자 전달 상태
  - 구매자 수령 상태
  - 현재 불일치 여부
  - 완료 가능 여부와 막힌 이유

#### 내 거래 화면
- thread 카드에 `전달 확인 대기`, `수령 확인 대기`, `부분거래 정리 필요`, `현장 조건 불일치` 배지를 노출해야 한다.
- unread보다 `completionEligibilityState`와 `exchangeConfirmationState`가 더 높은 우선순위를 가질 수 있어야 한다.

#### 완료 / 분쟁 화면
- 완료 확인 모달에는 `무엇을 몇 개 실제로 주고받았는지`를 마지막으로 재확인하는 structured summary가 필요하다.
- mismatch 보고 화면에는 자유서술만이 아니라 `wrong_item_or_option`, `quantity_shortfall`, `settlement_not_completed` 같은 구조화 선택지를 우선 제공한다.

### 105.9 API / read model 파생 기준
#### 사용자용 read model 후보
- `GET /me/trades/{tradeThreadId}` 응답에 `activeExchangeConfirmation` 포함
- `GET /chats/{chatRoomId}` 응답에 `exchangeStatusSummary` 포함
- `GET /trade-completions/{completionId}` 응답에 `exchangeReference` 포함

#### command surface 후보
- `POST /trade-threads/{tradeThreadId}/exchange/seller-ready`
- `POST /trade-threads/{tradeThreadId}/exchange/handover`
- `POST /trade-threads/{tradeThreadId}/exchange/receipt`
- `POST /trade-threads/{tradeThreadId}/exchange/report-mismatch`
- `POST /trade-threads/{tradeThreadId}/exchange/resolve-partial`

응답 예시:
```json
{
  "activeExchangeConfirmation": {
    "exchangeConfirmationId": "exc_123",
    "handoverState": "seller_marked_handed_over",
    "receiptState": "buyer_marked_received",
    "exchangeConfirmationState": "mutually_confirmed",
    "completionEligibilityState": "eligible",
    "sellerConfirmedQuantity": 3,
    "buyerConfirmedQuantity": 3,
    "exchangeMismatchType": null,
    "availableActions": ["request_completion", "view_locked_terms"]
  }
}
```

### 105.10 DB / 정합성 / 감사 기준
- `lockedDealTermsSnapshotId`는 필수 FK로 두고, completion/dispute/no-show case에서 동일 기준 snapshot을 재사용할 수 있어야 한다.
- `sellerConfirmedQuantity`, `buyerConfirmedQuantity`는 음수/0 불가이며, `locked finalQuantity`를 초과할 수 없다.
- `exchangeConfirmationState='mutually_confirmed'`인데 양측 확인 시각이 모두 비어 있는 상태는 금지한다.
- `mismatch_reported` 이후 완료 요청이 발생하면 서버는 `blocked_by_mismatch` 또는 명시 승인 경로 없이는 거절해야 한다.
- 전달/수령 관련 사용자 행동은 일반 메시지와 별도로 시스템 이벤트/감사로그에 append-only로 남긴다.

### 105.11 운영 정책 / 분쟁 판단 기준
- 운영자는 `도착했다`는 주장과 `실제로 주고받았다`는 주장을 분리해 판단해야 한다.
- 아래 사건은 exchange mismatch로 우선 분류할 수 있어야 한다.
  - 만났지만 아이템을 안 줌
  - 수량을 일부만 줌
  - 다른 옵션/강화 상태의 아이템을 제시함
  - 현장에서 가격/지급 방식이 달라짐
  - 서로 받았다고/못 받았다고 엇갈리게 주장함
- `exchangeMismatchType`는 completion dispute, payment risk, no-show 이후 현장 만남 발생 케이스와 연결될 수 있으므로 동일 사실관계 중복 집계를 피하는 linkage 정책이 필요하다.

### 105.12 analytics 이벤트 후보 / 오픈 질문
이벤트 후보:
- `exchange_confirmation_created`
- `exchange_handover_marked`
- `exchange_receipt_marked`
- `exchange_mismatch_reported`
- `exchange_partially_confirmed`
- `completion_eligibility_changed`

핵심 분석 포인트:
- 상호 전달/수령 확인이 있는 거래와 없는 거래의 완료율·분쟁률 차이
- `one_sided_claim`에서 `mutually_confirmed`로 전환되는 비율
- 부분거래 이후 잔여 수량 재오픈 성공률
- 현장 불일치 유형별 신고/제재 전환율

오픈 질문 / 가정:
- 가정: MVP에서는 사진 기반 증빙 업로드를 필수로 만들지 않고, structured confirmation + 채팅/시스템 로그를 우선 사용한다.
- 오픈 질문: `seller_marked_handed_over`와 `buyer_marked_received`를 강한 CTA로 분리할지, 단일 `거래 진행` 체크리스트로 감쌀지 결정 필요
- 오픈 질문: `one_sided_claim` 상태에서도 완료 요청을 열지 여부는 전환율과 허위 완료 리스크를 함께 봐야 한다.

## 106. 최종 거래 결과(Final Trade Outcome) / 완료·취소·부분완료·분쟁 종결 canonical 계약
### 106.1 목표
- `Listing.status`, `TradeCompletion`, `No-show Case`, `Exchange Confirmation`, `Dispute`, `Review eligibility`가 각각 다른 “종료 해석”을 갖지 않도록, 실제 거래가 어떤 결과로 닫혔는지 표현하는 최종 단일 결과 객체를 정의한다.
- 사용자 화면에는 단순한 결과 카피를 제공하되, API/DB/운영/analytics는 동일한 outcome vocabulary로 완료·불발·부분완료·운영판정 종결을 해석할 수 있어야 한다.
- 후기 공개, 신뢰 집계, 잔여 수량 재오픈, 운영 통지, 사건 링크가 모두 같은 결과 객체를 기준으로 파생되도록 한다.

### 106.2 적용 범위와 원칙
- `Final Trade Outcome`은 활성 거래 스레드가 실질적으로 닫히는 시점에 생성되는 canonical 종결 레코드다.
- 공개 상태(`available`, `completed`, `cancelled`)와 내부 사건 객체는 유지하되, “최종적으로 무슨 일이 일어났는가”는 outcome 객체가 대표한다.
- 하나의 trade thread에는 동시에 활성 최종 결과가 1개만 존재해야 한다. 단, 운영 재심/재판정 시 revision 이력을 가진 superseded outcome을 남길 수 있다.
- outcome은 append-only 원칙을 따르며, 기존 결과를 덮어쓰기보다 `supersededByOutcomeId`로 교체 이력을 남긴다.

### 106.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `finalOutcomeType` | `completed_successfully` / `completed_partially` / `cancelled_mutual` / `cancelled_unilateral` / `cancelled_no_show` / `cancelled_mismatch` / `cancelled_policy_blocked` / `voided_admin` | 거래가 어떻게 닫혔는지 나타내는 최종 결과 타입 |
| `outcomeClosureMode` | `user_mutual` / `user_one_sided_auto` / `moderator_resolved` / `system_timeout` / `policy_enforced` | 어떤 메커니즘으로 결과가 확정됐는지 |
| `outcomeEvidenceTier` | `chat_only` / `terms_and_reservation` / `exchange_confirmed` / `exchange_and_evidence_bundle` / `moderator_determined` | 결과 판정의 근거 강도 |
| `reviewEligibilityOutcome` | `both_allowed` / `seller_only` / `buyer_only` / `blocked_until_review` / `none` | 후기 작성 가능 여부 |
| `trustImpactMode` | `standard_completion` / `partial_completion` / `neutral_cancel` / `no_show_penalty` / `mismatch_penalty` / `policy_penalty` / `manual_override` | 신뢰/제한 집계 반영 방식 |
| `residualDisposition` | `none` / `reopen_remaining_quantity` / `close_listing_no_reopen` / `spawn_followup_listing` | 잔여 수량 또는 후속 매물 처리 방식 |

### 106.4 결과 타입 판정 매트릭스
| 최종 결과 | 기본 조건 | 후기 | 신뢰 집계 | 매물/잔여 처리 |
|---|---|---|---|---|
| `completed_successfully` | 최종 조건과 실제 인도/수령이 실질적으로 일치하고 완료 확정됨 | 양측 허용 | 정상 완료 집계 | 매물 종결 |
| `completed_partially` | 일부 수량만 인도·수령 확인되었고 잔여 처리 방침이 확정됨 | 양측 허용 가능, 단 코멘트 가이드 강화 | 부분완료 집계 | 잔여 수량 재오픈 또는 후속 매물 |
| `cancelled_mutual` | 양측 모두 거래 불발/취소에 동의 | 기본 차단 또는 중립 후기 정책 | 중립 취소 | 매물 available 복귀 또는 종결 |
| `cancelled_unilateral` | 한쪽 종료 요청으로 닫혔으나 제재급 사유는 아님 | 제한적 또는 기본 차단 | 약한 음영만 남기거나 중립 | 사유 따라 복귀/종결 |
| `cancelled_no_show` | no-show case 판정 또는 명시적 no-show 인정 | 판정 후 정책에 따라 허용 | no-show penalty | 상대 교체/재오픈 후보 |
| `cancelled_mismatch` | 현장 조건 불일치가 핵심 원인으로 판정 | 판정 전 차단, 판정 후 제한적 허용 | mismatch penalty | 재오픈 여부는 실제 잔여 재고 기준 |
| `cancelled_policy_blocked` | 정책 위반/운영 개입으로 거래 중단 | 차단 | policy penalty | 매물/스레드 숨김 또는 종료 |
| `voided_admin` | 잘못된 완료/중복 사건/오판정 교정으로 무효화 | 차단 | manual override | 별도 운영 경로 |

### 106.5 주요 사건과의 연결 규칙
- `completed_successfully`, `completed_partially`는 기본적으로 `TradeCompletion`이 존재해야 하며, 가능하면 `Exchange Confirmation` 또는 동등한 실행 로그를 가진다.
- `cancelled_no_show`는 `No-show Case` 판정 결과를 source-of-truth로 삼되, outcome 객체가 후기/신뢰/잔여 처리 기준을 대표한다.
- `cancelled_mismatch`는 `exchangeMismatchType`과 연결되며, 단순 no-show와 혼동되면 안 된다.
- `cancelled_policy_blocked`, `voided_admin`는 운영 액션과 감사로그를 강하게 연결해야 한다.
- outcome은 `linkedCompletionId`, `linkedNoShowCaseId`, `linkedDisputeId`, `linkedExchangeConfirmationId`를 optional FK로 가질 수 있어야 한다.

### 106.6 생성 시점과 수정 제한
- outcome은 아래 중 하나를 만족할 때 생성 가능하다.
  1. 완료가 확정됨
  2. no-show / mismatch / policy 사건이 판정됨
  3. 양측 합의 취소가 확정됨
  4. 운영자가 thread 종결 결과를 명시적으로 판정함
- 일반 사용자는 outcome을 직접 임의 선택하기보다, 완료/취소/신고/노쇼/불일치 액션을 통해 간접적으로 outcome 후보를 만든다.
- 최종 outcome이 확정되면 동일 thread의 CTA는 결과에 맞게 축소되고, 재거래는 새 thread로만 시작한다.
- 결과 확정 후 `finalOutcomeType` 변경은 운영 재심 또는 시스템 정합성 복구에서만 허용한다.

### 106.7 Final Trade Outcome 필수 필드 후보
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `tradeOutcomeId` | 필수 | 결과 식별자 |
| `tradeThreadId` | 필수 | 종결 대상 thread |
| `listingId` | 필수 | 연결 매물 |
| `primaryCounterpartyPairId` | 필수 | 결과를 공유하는 거래 당사자 쌍 |
| `finalOutcomeType` | 필수 | 최종 결과 타입 |
| `outcomeClosureMode` | 필수 | 확정 방식 |
| `outcomeEvidenceTier` | 필수 | 근거 강도 |
| `reviewEligibilityOutcome` | 필수 | 후기 허용 상태 |
| `trustImpactMode` | 필수 | 신뢰도 반영 방식 |
| `residualDisposition` | 필수 | 잔여 처리 방침 |
| `resolvedAt` | 필수 | 결과 확정 시각 |
| `resolvedByActorType` | 필수 | `seller` / `buyer` / `system` / `moderator` |
| `summaryReasonCode` | 선택 | 결과 요약 사유 코드 |
| `supersededByOutcomeId` | 선택 | 재판정 시 후속 outcome |

### 106.8 화면 / 알림 / 운영 노출 기준
#### 내 거래 / 거래 상세
- thread 최상단에 `최종 결과 배지 + 한 줄 설명 + 후기 가능 여부 + 재거래 CTA`를 공통 형식으로 노출한다.
- `completed_partially`는 완료 배지와 함께 `잔여 수량 남음` 문구를 별도 표시해야 한다.
- `cancelled_no_show`, `cancelled_mismatch`는 감정적 문구보다 `운영 판정 결과` 또는 `불발 사유` 중심으로 요약한다.

#### 알림
- 결과 확정 알림은 단순 상태값이 아니라 outcome 중심으로 발송해야 한다.
- 예: `거래가 완료로 확정되었어요`, `노쇼 판정으로 거래가 취소되었어요`, `현장 조건 불일치로 거래가 종료되었어요`
- 후기 CTA는 `reviewEligibilityOutcome`을 그대로 반영해 잘못 열린 링크를 방지해야 한다.

#### 운영
- 운영 백오피스는 사건별 세부 객체 외에 `최종 결과 타임라인 카드`를 보여줘야 한다.
- 운영자는 결과 수정 시 기존 사건을 직접 파괴하지 않고 새 outcome revision을 생성해야 한다.

### 106.9 API 후보
#### 사용자용 read model
- `GET /me/trades/{tradeThreadId}/outcome`
- `GET /trade-outcomes/{tradeOutcomeId}`

#### 내부/운영용
- `POST /admin/trade-outcomes/{tradeThreadId}/resolve`
- `POST /admin/trade-outcomes/{tradeOutcomeId}/supersede`

응답 필드 예시:
```json
{
  "tradeOutcomeId": "outc_123",
  "tradeThreadId": "thread_123",
  "finalOutcomeType": "completed_partially",
  "outcomeClosureMode": "moderator_resolved",
  "reviewEligibilityOutcome": "both_allowed",
  "trustImpactMode": "partial_completion",
  "residualDisposition": "reopen_remaining_quantity",
  "resolvedAt": "2026-03-14T06:27:00+09:00"
}
```

### 106.10 DB / 정합성 / analytics 기준
- thread별 활성 outcome 유니크 제약이 필요하다. (`supersededByOutcomeId is null` 조건부 unique 후보)
- `completed_successfully` 또는 `completed_partially`인데 `resolvedAt` 또는 `reviewEligibilityOutcome`가 비어 있는 상태는 금지한다.
- `cancelled_no_show` outcome은 연결된 no-show 사건 또는 명시적 no-show 판정 근거 없이 생성되면 안 된다.
- `residualDisposition='reopen_remaining_quantity'`이면 listing 잔여 수량, allocation 상태, 검색 노출 재개 시점이 같이 정합성을 가져야 한다.
- analytics는 퍼널 종료 이벤트를 outcome 기준으로 집계해야 한다. 즉, 완료 요청/분쟁 접수보다 `최종 outcome 확정`이 진짜 종결 모수다.

이벤트 후보:
- `trade_outcome_resolved`
- `trade_outcome_superseded`
- `trade_outcome_review_eligibility_opened`
- `trade_outcome_residual_reopened`

오픈 질문 / 가정:
- 가정: MVP에서는 outcome을 별도 쓰기 surface로 사용자에게 노출하지 않고, 기존 완료/취소/노쇼/분쟁 액션의 결과 projection으로만 제공한다.
- 오픈 질문: `cancelled_mutual`에도 상호 후기(예: 약속은 잘 지켰지만 거래는 안 함)를 허용할지 여부는 악용 가능성과 신뢰도 효용을 함께 검토해야 한다.
- 오픈 질문: `voided_admin` outcome을 사용자에게 그대로 노출할지, 더 일반적인 `운영 조정됨` 문구로 추상화할지 결정 필요.

## 107. 재일정-노쇼-최종결과 연결 판정(Re-schedule to No-show Outcome Bridge) 계약
### 107.1 목표
- accepted/rejected/expired reschedule이 당일 실행, 노쇼 claim, dispute, final outcome으로 이어질 때 중간 해석이 화면마다 달라지지 않도록 canonical bridge 규칙을 둔다.
- `늦어요`, `장소 바꿔요`, `응답 없음`, `이미 도착함`, `서로 다른 장소에 있었음`이 모두 no-show로 뭉개지지 않게 판정 경계를 세분화한다.
- 채팅 CTA, no-show 사건 생성, outcome 확정, 후기 게이트, trust/analytics가 동일한 판단 단위를 바라보도록 한다.

### 107.2 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `executionDecisionState` | `keep_original_meeting` / `await_reschedule_response` / `rescheduled_binding` / `reschedule_failed_original_binding` / `reschedule_failed_no_show_window` / `mutual_abort_candidate` | 현재 약속 해석의 기준점 |
| `rescheduleNoShowBridgeType` | `none` / `accepted_resets_window` / `pending_does_not_reset` / `rejected_reverts_original` / `expired_reverts_original` / `late_change_requires_manual_review` | 재일정과 노쇼 grace를 어떻게 연결하는지 |
| `claimSuppressionReasonCode` | `accepted_reschedule_exists` / `counterparty_arrived_before_cutoff` / `mutual_abort_not_no_show` / `insufficient_location_ack` / `active_dispute_exists` | no-show claim 생성/집계를 억제하는 사유 |
| `outcomeCandidateSource` | `original_reservation` / `accepted_reschedule` / `arrival_evidence` / `manual_moderation` / `mutual_chat_agreement` | 최종 결과 판단의 주근거 |
| `scheduleBindingVersion` | integer | 원예약=1, accepted reschedule마다 +1 되는 유효 약속 버전 |

원칙:
- no-show는 항상 `현재 binding schedule version` 기준으로만 판단한다.
- accepted reschedule 없는 일방적 변경 요청은 binding schedule을 바꾸지 않는다.
- binding schedule이 바뀌면 이전 시간/장소 기준 도착 증빙은 참고자료일 뿐 자동 승패 근거가 아니다.

### 107.3 binding schedule 결정 규칙
| 조건 | binding schedule | 노쇼 기준 시각/장소 | 비고 |
|---|---|---|---|
| 재일정 없음 | original reservation | 원예약 | 기본 |
| reschedule `requested`만 존재 | original reservation | 원예약 | 일방 요청만으로 grace reset 금지 |
| reschedule `accepted` | latest accepted reschedule | 최신 accepted 스냅샷 | `scheduleBindingVersion` 증가 |
| reschedule `rejected` | original reservation 또는 직전 accepted | 기존 binding 유지 | reject 자체는 새 기준 생성 아님 |
| reschedule `expired` | original reservation 또는 직전 accepted | 기존 binding 유지 | 응답 없음은 원기준 유지 |
| 거래 직전 대폭 변경 + 읽음/확인 불충분 | manual review candidate | 자동 판정 억제 가능 | 장소 혼선/사칭 리스크 |

세부 규칙:
- `accepted`된 재일정이 존재하면 original reservation 기준 no-show claim은 `accepted_reschedule_exists`로 억제한다.
- 같은 예약에서 여러 accepted 재일정이 있었으면 가장 마지막 accepted만 binding schedule이 된다.
- `location_only` 재일정이라도 accepted이면 장소 기준이 바뀌므로, 이전 장소 도착 체크인은 단독 근거로 사용하지 않는다.

### 107.4 no-show claim 생성 허용/억제 매트릭스
| 상황 | claim 허용 여부 | 기본 처리 | 비고 |
|---|---|---|---|
| original schedule만 있고 상대 무도착 | 허용 | 일반 no-show 흐름 | |
| pending reschedule 요청만 있고 상대 무응답 | 허용 | original 기준 no-show 가능 | `pending_does_not_reset` |
| accepted reschedule 이후 새 시각 무도착 | 허용 | accepted 기준 no-show 가능 | |
| accepted reschedule 이후 옛 시각 기준 claim | 억제 | 안내 메시지 + 새 기준 제시 | |
| 양측이 채팅에서 `오늘 취소하자` 합의 | 억제 | `mutual_abort_candidate`로 outcome 후보화 | no-show 아님 |
| 장소가 직전 변경됐으나 상대 ack 없음 | 제한 허용 | manual review 우선 | 자동 trust 반영 금지 후보 |
| 이미 active no-show/dispute 존재 | 억제 | 기존 사건에 추가 진술 연결 | 중복 사건 생성 금지 |

### 107.5 당일 실행 카드 / CTA gating 파생 기준
#### 카드 상태 계산 원칙
- `executionDecisionState=await_reschedule_response`이면 기본 CTA는 `기존 약속 기준 유지`를 분명히 보여줘야 한다.
- `rescheduled_binding`이면 카드의 메인 시간/장소, ETA, no-show countdown 모두 최신 accepted 기준으로 재계산한다.
- `late_change_requires_manual_review`이면 사용자에게 단정적 `노쇼 신고`보다 `상황 기록 남기기` CTA를 우선 노출한다.

#### CTA family 후보
| 상태 | 우선 CTA | 보조 CTA |
|---|---|---|
| `keep_original_meeting` | `도착했어요` / `상대 안 보여요` | `시간/장소 변경 요청` |
| `await_reschedule_response` | `기존 약속대로 갈게요` / `변경안 확인하기` | `거절` |
| `rescheduled_binding` | `새 약속 확인` / `도착했어요` | `5분 늦어요` |
| `reschedule_failed_original_binding` | `원래 약속 기준으로 진행` | `노쇼 신고` |
| `late_change_requires_manual_review` | `상황 기록 남기기` | `운영 검토 요청` |

### 107.6 분쟁 / outcome 연결 규칙
- no-show case 생성 시 반드시 `evaluatedScheduleBindingVersion`과 `outcomeCandidateSource`를 저장해야 한다.
- `mutual_abort_candidate`는 기본적으로 `cancelled_mutual` outcome 후보이며 no-show strike 후보가 아니다.
- `late_change_requires_manual_review` 또는 `insufficient_location_ack`가 붙은 사건은 auto-adjudication 및 즉시 trust 반영을 금지하는 보수안을 기본으로 둔다.
- accepted reschedule 이후 발생한 no-show 판정은 `outcomeCandidateSource=accepted_reschedule`로 final outcome에 연결되어야 한다.
- original 기준 claim이 억제된 경우에도 사용자의 도착/evidence는 사건 메모로 보존해 운영자가 변경 이력 전체를 복기할 수 있어야 한다.

### 107.7 후기 / trust / restriction 반영 기준
| bridge 상황 | review 기본안 | trust 반영 | restriction 후보 |
|---|---|---|---|
| accepted reschedule 후 명확한 no-show | 후기 비허용 또는 판정 후 제한 허용 | no-show 반영 가능 | 반복 시 가능 |
| pending request만 있었고 original 기준 no-show 성립 | 후기 비허용 | original 기준 no-show 반영 | 가능 |
| mutual abort confirmed | 후기 기본 비허용 | no-show 반영 금지 | 금지 |
| location ack 불충분 / manual review | 판정 전 비허용 | 보류 | 보류 |
| rejected/expired reschedule 후 original 기준 무도착 | no-show 판정 후 정책 따라 처리 | 반영 가능 | 가능 |

원칙:
- `재일정 논의가 있었다`는 사실만으로 no-show 책임을 완화하지 않는다. 실제 binding schedule 변경 여부가 핵심이다.
- mutual abort, evidence-insufficient, location-confusion은 no-show strike 모수에서 분리 집계해야 한다.

### 107.8 API / DB / analytics 파생 기준
#### API 후보 필드
- `Reservation`/`TradeThreadProjection` read model에 아래 필드 후보를 둔다.
  - `currentScheduleBindingVersion`
  - `executionDecisionState`
  - `rescheduleNoShowBridgeType`
  - `noShowClaimAllowed`
  - `claimSuppressionReasonCode`

응답 예시:
```json
{
  "tradeThreadId": "thread_123",
  "currentScheduleBindingVersion": 2,
  "executionDecisionState": "rescheduled_binding",
  "rescheduleNoShowBridgeType": "accepted_resets_window",
  "noShowClaimAllowed": true,
  "claimSuppressionReasonCode": null,
  "bindingSchedule": {
    "source": "accepted_reschedule",
    "scheduledAt": "2026-03-14T21:30:00+09:00",
    "meetingSummary": "기란 마을 창고 앞"
  }
}
```

#### DB / 정합성
- accepted reschedule가 생성될 때 `scheduleBindingVersion`을 증가시키고 이전 binding schedule snapshot을 immutable하게 보존해야 한다.
- no-show case는 `tradeThreadId + evaluatedScheduleBindingVersion + claimantUserId` 기준 중복 제한을 검토한다.
- `claimSuppressionReasonCode`가 있는 억제 이벤트도 analytics/운영 추적을 위해 별도 로그 또는 timeline event로 남겨야 한다.

#### analytics 이벤트 후보
- `reschedule_binding_shifted`
- `no_show_claim_suppressed`
- `no_show_claim_opened_on_binding_version`
- `mutual_abort_detected`
- `late_location_change_manual_review_required`

핵심 분석 포인트:
- accepted reschedule 이후 no-show율 vs original schedule no-show율
- pending request만 남긴 사용자의 no-show 관여 비율
- mutual abort와 true no-show를 분리했을 때 strike/제한 정책의 오탐 감소 여부
- location ack 불충분 케이스의 운영 개입률과 판정 뒤집힘 비율

### 107.9 오픈 질문 / 가정
- 가정: MVP에서는 `mutual abort confirmed`를 명시 버튼보다 채팅 합의 + 취소 액션 조합으로 먼저 판정하고, 명시 전용 CTA는 추후 도입한다.
- 가정: same-day accepted reschedule는 최대 2회까지만 UX 상 적극 지원하고, 그 이상은 `too_many_changes` 경고를 강화한다.
- 오픈 질문: accepted reschedule 이후 상대가 old place에 도착한 경우를 완전 사용자 책임으로 볼지, 첫 1회는 완화 규칙을 둘지 운영 데이터 기반 보정이 필요하다.
- 오픈 질문: location ack 부족 상황에서 사용자에게 `노쇼 신고` 버튼을 완전히 숨길지, `운영 검토 요청`으로 대체할지 세부 copy 결정을 남긴다.

## 108. 콘텐츠 검수 케이스(Content Moderation Review Case) / 매물·채팅·후기·프로필 정책판정 canonical 계약
### 108.1 목표
- Listing, ChatMessage, Review, UserProfile, Reservation note에서 발생하는 정책 검수를 **개별 객체의 부가 상태**가 아니라 공통 사건 단위로 다룬다.
- 자동 탐지, 사용자 수정 유도, 운영 검토, 임시숨김, 제재, 복구가 객체별로 서로 다른 용어를 쓰지 않도록 canonical read/write 모델을 정의한다.
- 등록 화면, 채팅 입력 UX, 운영 백오피스, API, 감사 로그, analytics가 동일한 `contentModerationCase`를 기준으로 동작하게 한다.

### 108.2 적용 범위
본 섹션은 아래 콘텐츠 객체에 적용한다.
- `listing_body`: 매물 제목/설명/속성/이미지/OCR 결과
- `chat_message`: 일반 채팅 메시지/첨부/OCR 결과
- `review_content`: 후기 본문/평가 코멘트
- `profile_content`: 닉네임/소개/아바타
- `reservation_note`: 예약 메모/장소 설명/상대 안내 텍스트

적용 제외:
- 거래 결과 판정 자체는 `Dispute`, `No-show Case`, `Final Trade Outcome`가 소유한다.
- 계정 수준 제한 결정은 `Restriction`이 소유하되, 본 섹션의 검수 케이스가 입력 신호가 될 수 있다.

### 108.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `moderationObjectType` | `listing_body` / `chat_message` / `review_content` / `profile_content` / `reservation_note` | 검수 대상 객체 유형 |
| `moderationTriggerType` | `pre_submit_scan` / `post_publish_scan` / `user_report` / `appeal_recheck` / `bulk_policy_migration` / `manual_sampling` | 검수 시작 계기 |
| `moderationCaseState` | `detected` / `user_fix_required` / `under_review` / `temporarily_limited` / `resolved_allow` / `resolved_restrict` / `resolved_remove` / `reopened` / `dismissed` | 사건 상태 |
| `moderationDecisionType` | `allow` / `warn_only` / `limit_visibility` / `block_submit` / `remove_content` / `mask_partial` / `escalate_restriction` | 최종 또는 중간 판정 종류 |
| `policyEvidenceSource` | `text_rule` / `ocr_rule` / `image_classifier` / `user_report_bundle` / `manual_observation` / `historical_pattern` | 판정 근거 소스 |
| `decisionVisibilityLevel` | `user_facing` / `staff_only` / `mixed` | 사유/근거의 외부 노출 수준 |
| `caseClosureReasonCode` | `user_fixed` / `false_positive` / `policy_violation_confirmed` / `insufficient_evidence` / `superseded_by_new_case` / `appeal_upheld` / `appeal_rejected` | 종료 사유 |

원칙:
- 하나의 콘텐츠 객체는 시간에 따라 여러 moderation case를 가질 수 있다.
- 한 시점에 **활성 케이스(active case)** 는 최대 1개를 기본으로 하며, 새 케이스가 열리면 이전 active case는 `superseded` 성격 종료로 닫는다.

### 108.4 canonical 객체 정의
`contentModerationCase`는 아래 정보를 최소 포함해야 한다.

```json
{
  "moderationCaseId": "cmc_123",
  "moderationObjectType": "listing_body",
  "objectId": "listing_123",
  "objectVersion": 7,
  "moderationTriggerType": "pre_submit_scan",
  "moderationCaseState": "user_fix_required",
  "moderationDecisionType": "warn_only",
  "primaryPolicyCode": "EXTERNAL_CONTACT_IN_BODY",
  "secondaryPolicyCodes": ["SENSITIVE_INFO_DETECTED"],
  "decisionVisibilityLevel": "user_facing",
  "riskTier": "high",
  "detectedAt": "2026-03-14T08:03:00+09:00",
  "resolvedAt": null,
  "linkedReportId": null,
  "linkedRestrictionId": null,
  "active": true
}
```

핵심 필드 원칙:
- `objectVersion`은 어떤 원문/수정본에 대한 판정인지 식별하는 immutable 참조다.
- `primaryPolicyCode`는 사용자/운영/API가 공통으로 바라보는 대표 코드이며, 세부 rule hit는 별도 evidence bundle에 저장한다.
- `linkedRestrictionId`는 콘텐츠 위반이 계정 제재로 이어졌을 때만 연결한다.

### 108.5 상태 전이 규칙
#### 기본 흐름
- 자동/수동 탐지 발생 → `detected`
- 사용자가 즉시 고칠 수 있는 입력 단계 위반 → `user_fix_required`
- 운영 판단이 필요한 케이스 → `under_review`
- 객체 노출/전송을 임시 제한한 상태 → `temporarily_limited`
- 문제 없음 확정 → `resolved_allow` 또는 `dismissed`
- 정책 위반 확정, 제출 차단/비노출/마스킹 → `resolved_restrict` 또는 `resolved_remove`
- 이의제기/재검토로 재개봉 → `reopened`

#### 객체별 해석 규칙
- `listing_body`: 등록 전 차단이면 `block_submit`, 등록 후 숨김이면 `remove_content` 또는 `limit_visibility`
- `chat_message`: 전송 전 차단이면 `block_submit`, 전송 후 일부 마스킹이면 `mask_partial`
- `review_content`: 공개 전 보류면 `temporarily_limited`, 공개 후 비노출이면 `resolved_remove`
- `profile_content`: 수정 요구 중심이면 `user_fix_required`, 반복 위반이면 `escalate_restriction`

### 108.6 사용자 경험 / 화면 계약
#### 입력 단계
- 사용자가 즉시 수정 가능한 경우는 `user_fix_required`를 우선 사용하고, 곧바로 운영 큐로 보내지 않는다.
- 입력 UI는 내부 rule 이름 대신 `수정 필요 이유`, `다시 시도 가능 여부`, `차단된 부분`만 보여준다.
- 예: `연락처나 외부 링크는 본문에 직접 적을 수 없어요` / `욕설·광고성 표현은 후기에 작성할 수 없어요`

#### 게시 후/전송 후
- 이미 게시된 매물/후기가 제한된 경우, 소유자에게는 `검토 중`, `비노출됨`, `수정 후 복구 가능` 중 하나로 단순 노출한다.
- 상대 사용자에게는 과도한 낙인 없이 `일시적으로 볼 수 없음`, `운영정책에 따라 숨김 처리됨` 정도의 문구만 제공한다.
- 채팅 메시지 마스킹 시 상대에게는 원문 대신 `운영정책에 따라 일부 내용이 가려졌어요` 수준의 placeholder를 노출한다.

### 108.7 운영 큐 / 우선순위 계약
운영 큐는 `Report`, `Dispute`, `SupportCase`와 별도로 분리할 수 있으나, 동일한 case feed에서 교차 탐색 가능해야 한다.

| 큐 우선순위 | 조건 예시 | 기대 처리 |
|---|---|---|
| `CM1` | 사기 유도, 외부 연락처 강한 유도, 개인정보 노출, 불법/심각한 괴롭힘 | 즉시 제한 또는 숨김 |
| `CM2` | 반복 광고, 후기 비방, 이미지 OCR 민감정보, 예약 직전 위험 장소 문구 | 당일 검토 |
| `CM3` | 설명 부족, 오탐 가능성 높은 문구, 프로필 소개 경미 위반 | 순차 검토 |
| `CM4` | 표본 점검, 정책 마이그레이션 재스캔 | 배치/비실시간 |

운영 화면 필수 요소:
- 원문/마스킹 원문 비교
- policy code / evidence source / confidence tier
- 동일 객체 이전 케이스 이력
- 동일 사용자 최근 위반 묶음
- `수정 후 재제출`, `비노출 유지`, `복구`, `restriction 연결` 액션

### 108.8 evidence bundle / 감사 추적 원칙
각 moderation case는 `contentModerationEvidenceBundle`을 가져야 한다.

최소 포함 후보:
- 원문 스냅샷 또는 해시
- 탐지된 span/토큰/이미지 영역 정보
- OCR/분류기/규칙 hit 목록
- 사용자 신고 설명/첨부(있는 경우)
- 운영자 메모와 판정 diff
- 판정 시점의 객체 공개 상태/가시성/버전

원칙:
- 증빙은 append-only로 보관하고, 이후 원문이 수정되어도 당시 판정 근거는 변하지 않아야 한다.
- 사용자에게는 필요한 범위만 요약 노출하고, rule 내부 세부값/탐지 임계치는 공개하지 않는다.

### 108.9 다른 도메인과의 연결 규칙
- `Report`가 원인일 때 `linkedReportId`를 연결하되, Report 종료가 moderation case 종료를 자동 보장하지 않는다.
- 같은 정책 위반이 반복되면 `policyStrikeCount`를 증가시키고, `Restriction` 생성 여부는 별도 정책 엔진이 결정한다.
- `SupportCase`는 사용자의 이의제기/문의 창구로 연결되며, 원 canonical 판정 객체는 moderation case다.
- `Analytics`는 단순 차단 횟수보다 `detect → fix → allow`, `detect → review → remove`, `appeal → reopen → restore` 전환을 추적해야 한다.

### 108.10 API / DB / analytics 파생 기준
#### API 후보
- `GET /me/moderation-cases`
- `POST /moderation-cases/{moderationCaseId}/appeal`
- `GET /admin/moderation-cases`
- `GET /admin/moderation-cases/{moderationCaseId}`
- `POST /admin/moderation-cases/{moderationCaseId}/resolve`
- `POST /admin/moderation-cases/{moderationCaseId}/restore-content`

#### DB 후보 엔티티
- `ContentModerationCase`
- `ContentModerationEvidenceBundle`
- `ContentModerationDecisionLog`
- `ContentPolicyStrikeLedger`

#### analytics 이벤트 후보
- `content_moderation_case_opened`
- `content_fix_requested`
- `content_resubmitted_after_fix`
- `content_moderation_resolved_allow`
- `content_moderation_resolved_remove`
- `content_moderation_appeal_submitted`
- `content_moderation_restored`

핵심 분석 포인트:
- object type별 오탐률(`resolved_allow` / 전체 탐지)
- `user_fix_required` 후 재제출 성공률
- 운영 검토 평균 리드타임
- 콘텐츠 위반 → restriction 전환율
- appeal 뒤 복구 비율과 policy code별 편차

### 108.11 오픈 질문 / 가정
- 가정: MVP에서는 `contentModerationCase`를 우선 listing/review/profile에 적용하고, chat_message는 고위험 케이스부터 단계 도입한다.
- 가정: 입력 단계 차단은 가능한 한 `user_fix_required`로 시작하고, 명백한 금지 패턴만 `block_submit`을 사용한다.
- 오픈 질문: 후기/프로필의 경미 위반을 운영 큐 없이 자동 마스킹으로 닫을지 여부는 오탐 비용과 CS 비용을 함께 보고 결정한다.
- 오픈 질문: 채팅 메시지 마스킹 placeholder를 양측 모두 동일하게 볼지, 발신자에게만 더 상세 수정 사유를 노출할지 copy/정책 확정이 필요하다.

## 109. 제한 해제 / 재활성화 / probation(관찰기간) canonical 계약
### 109.1 목표
- 경고, 기능 제한, 일시 정지, appeal 승인 이후 사용자가 **어떤 조건으로 어떤 권한을 되찾는지**를 구조화한다.
- `Restriction`, `TradeTrustSnapshot`, `SupportCase`, 알림, 프로필 배지, 운영 백오피스가 서로 다른 해석을 쓰지 않도록 해제 단위를 canonical하게 정의한다.
- "제재 중"만이 아니라 "해제됐지만 관찰 중", "조건 미충족으로 아직 복귀 불가", "자동 만료 해제"를 분리해 제품/운영/DB/API 기준을 맞춘다.

### 109.2 적용 범위
이 섹션은 아래 제한 계열에 적용한다.
- 채팅 제한(`chat_only`, `messaging_cooldown`)
- 매물 등록 제한(`listing_only`, `listing_publish_blocked`)
- 신뢰 제한(`trust_limited`, `review_gated`)
- 거래 실행 제한(`reservation_blocked`, `completion_confirmation_blocked`)
- 계정 수준 일시 정지(`temporary_suspend`)

적용 제외:
- 영구 밴(`permanent_ban`)의 완전 복구는 기본 플로우가 아니라 exceptional appeal path로만 다룬다.
- 콘텐츠 객체 단위 숨김/복구는 108장 `contentModerationCase`가 primary source of truth이고, 이 섹션은 계정/기능 제한 복귀를 다룬다.

### 109.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `restrictionRecoveryState` | `not_eligible` / `eligible_now` / `pending_requirement` / `pending_review` / `lift_scheduled` / `lifted` / `lifted_with_probation` / `reinstated_restricted` | 제한 해제 관점의 현재 상태 |
| `liftDecisionType` | `auto_expire` / `manual_lift` / `appeal_upheld` / `requirement_completed` / `partial_lift` / `emergency_restore` | 어떤 방식으로 해제됐는지 |
| `probationTier` | `none` / `light` / `standard` / `strict` | 해제 후 관찰 강도 |
| `recoveryRequirementType` | `acknowledge_policy` / `profile_fix` / `content_cleanup` / `identity_reverify` / `support_reply_required` / `cooldown_wait` | 복귀 전 사용자 충족 요건 |
| `recoveryFailureReasonCode` | `repeat_violation` / `appeal_rejected` / `requirement_expired` / `new_report_opened` / `risk_not_reduced` | 복귀 실패/보류 사유 |
| `reinstatementScope` | `full_account` / `chat_only` / `listing_only` / `trust_surface_only` | 복구 범위 |

원칙:
- 제한(`restrictionState`)과 복귀(`restrictionRecoveryState`)는 별도 축이다.
- 같은 제한이 종료돼도 즉시 `정상 사용자`가 되는 것이 아니라 probation이 남을 수 있다.
- UI에는 내부 용어 그대로 노출하지 않고 `제한 해제 예정`, `정책 확인 후 재사용 가능`, `일부 기능만 복구됨`, `관찰 중` 정도의 카피로 매핑한다.

### 109.4 복귀 상태 모델
| 상태 | 의미 | 사용자 노출 예시 | 주요 허용 액션 |
|---|---|---|---|
| `not_eligible` | 아직 해제 대상 아님 | 이용 제한 중 | 제한 이유 확인, 문의 |
| `pending_requirement` | 사용자 조치 필요 | 정책 확인 후 다시 이용 가능 | 약관/정책 확인, 프로필 수정, 소명 제출 |
| `pending_review` | 운영 검토 대기 | 검토 후 안내 예정 | 상태 확인, 추가 자료 제출 |
| `lift_scheduled` | 자동 종료 시각이 정해짐 | 03/16 21:00 이후 해제 예정 | 대기 |
| `eligible_now` | 내부 요건 충족, 즉시 해제 가능 | 다시 사용할 수 있어요 | 재활성화 확인 |
| `lifted` | 정상 복구 완료 | 제한 해제됨 | 일반 사용 |
| `lifted_with_probation` | 복구됐지만 관찰 중 | 일부 안전 제한이 유지돼요 | 제한 범위 내 사용 |
| `reinstated_restricted` | 해제 후 재위반으로 제한 재적용 | 다시 이용 제한됨 | 문의/appeal |

### 109.5 제한 해제 방식별 규칙
#### A. 자동 만료 해제(`auto_expire`)
- 기간형 제한이 `endsAt` 도달로 끝난 경우.
- 별도 사용자 확인이 필요 없는 제한에만 적용한다.
- 만료 즉시 모든 제한을 풀지 않고, 리스크가 남는 경우 `lifted_with_probation`으로 전환할 수 있다.

#### B. 수동 해제(`manual_lift`)
- 운영자가 오탐, 상황 종료, 내부 재검토를 근거로 제한을 철회하는 경우.
- 원 제한 로그를 덮어쓰지 않고 `liftDecision`를 append-only로 남긴다.
- 사용자에게는 "운영 검토 후 복구" 성격의 결과 메시지를 보낸다.

#### C. appeal 승인(`appeal_upheld`)
- 사용자가 이의제기했고 운영이 원 제재를 취소/축소한 경우.
- 원 제한의 정당성 평가와 별도로, 신규 정보 반영 여부(`new_evidence_added=true`)를 남겨야 한다.
- 오탐/정책경계/조건부허용 케이스를 구분할 수 있어야 한다.

#### D. 요건 충족 후 해제(`requirement_completed`)
- 정책 재확인, 프로필 수정, 금지 콘텐츠 삭제, 추가 인증, support reply 등 특정 조치 완료 후 해제되는 경우.
- 해제 가능한 최소 scope부터 부분 복구를 허용할 수 있다.
- 예: listing 제한은 해제됐지만 chat 제한은 probation 동안 유지.

### 109.6 probation(관찰기간) 계약
probation은 제재의 연장이 아니라 **정상 복귀 직후 재위반 리스크를 낮추기 위한 완충 상태**다.

| probationTier | 기간 기본안(가정) | 제약 예시 |
|---|---|---|
| `light` | 3일 | 대량행동 rate limit 강화 |
| `standard` | 7일 | 예약/후기/외부연락처 관련 정책 경고 강화 |
| `strict` | 14일 | 일부 고위험 action 추가 검토 또는 한도 축소 |

원칙:
- probation 중에도 사용자는 핵심 목적 행동을 할 수 있어야 하며, 사실상 영구 반쯤 제한 상태가 되어서는 안 된다.
- 단, 동일 또는 인접 policy code 재위반 시 escalation이 더 빠르게 적용될 수 있다.
- 공개 프로필에는 probation 자체를 직접 노출하지 않는다. 내부 risk/read model과 운영 도구, 사용자 본인 상태화면에만 사용한다.

### 109.7 부분 복구 / 단계적 복구 규칙
| 원 제한 | 복구 기본안 | 비고 |
|---|---|---|
| `listing_only` | listing 작성/수정부터 복구 | bump/대량등록은 probation 동안 제한 가능 |
| `chat_only` | 메시지 전송 복구, 대량 DM rate limit 강화 | 신고 다발 사용자에 유용 |
| `trust_limited` | 거래는 가능하되 공개 배지/후기 게이트는 지연 해제 가능 | trust snapshot 반영 필요 |
| `temporary_suspend` | 로그인/읽기→쓰기 순으로 단계 복귀 가능 | 고위험 케이스는 partial lift 허용 |
| `review_gated` | 후기 작성 또는 공개만 별도 복구 | 보복성 후기 대응 |

부분 복구 원칙:
- 사용자가 무엇이 풀렸고 무엇이 아직 제한인지 한 화면에서 이해할 수 있어야 한다.
- API `availableActions`는 partial lift 이후 즉시 최신 상태를 반영해야 한다.

### 109.8 재위반 / 재에스컬레이션 규칙
- probation 중 동일 계열 위반이 발생하면 `reinstated_restricted`로 전환 가능하다.
- 재에스컬레이션은 아래를 함께 본다.
  1. 동일 policy/restriction family인지
  2. probation 남은 기간
  3. 사용자 충족 요건이 형식적 완료였는지 실질적 개선인지
  4. 신규 support/appeal 케이스 존재 여부
- 단순 1회 오탐 또는 경미한 경계 위반은 즉시 더 높은 레벨로 올리지 않고 운영 검토를 우선한다.

### 109.9 사용자 화면 / 운영 화면 요구사항
#### 사용자 측
- 계정 상태/도움말 화면에 아래를 구조화해 보여야 한다.
  - 현재 제한 요약
  - 자동 해제 예정 시각 또는 필요한 조치
  - 일부 기능만 복구되었는지 여부
  - appeal 또는 support 문의 가능 여부
- 채팅/매물/후기 화면에서는 전면 차단 문구 대신 **행동 단위 제한 사유**를 보여준다.
  - 예: `현재는 새 매물 등록만 제한되어 있어요`
  - 예: `정책 확인을 완료하면 다시 채팅할 수 있어요`

#### 운영 측
- 원 제한과 해제 결정을 한 타임라인에서 봐야 한다.
- 운영자는 아래를 확인 가능해야 한다.
  - 제한 근거 사건
  - recovery requirement 충족 여부
  - probation 종료 예정 시각
  - 재위반 여부 및 escalation history
  - 사용자에게 실제 노출된 안내 문구 버전

### 109.10 `TradeTrustSnapshot` / `SupportCase` / 알림 연동 규칙
- `TradeTrustSnapshot`은 active restriction뿐 아니라 `lifted_with_probation` 여부를 내부 risk 계산에 반영할 수 있어야 한다.
- `SupportCase`는 restriction 해제 요청/appeal/조건 충족 확인의 사용자 커뮤니케이션 창구가 되며, restriction 자체를 대체하지 않는다.
- 알림은 아래 시점에 발송한다.
  1. 제한 시작
  2. 자동 해제 예정 1회
  3. 해제 완료
  4. 추가 조치 필요
  5. appeal 결과
- 푸시 카피에는 제한 상세 사유 전체를 노출하지 않고, 앱 진입 후 상세 화면에서 확인하도록 한다.

### 109.11 API / DB / analytics 파생 기준
#### API 후보
- `GET /me/restrictions`
- `POST /me/restrictions/{restrictionId}/acknowledge`
- `POST /me/restrictions/{restrictionId}/complete-requirement`
- `POST /me/restrictions/{restrictionId}/appeal`
- `GET /admin/restrictions`
- `GET /admin/restrictions/{restrictionId}`
- `POST /admin/restrictions/{restrictionId}/lift`
- `POST /admin/restrictions/{restrictionId}/set-probation`

#### DB 후보 엔티티 / 필드
- `RestrictionRecoveryDecision`
- `RestrictionRecoveryRequirement`
- `RestrictionProbationState`
- `Restriction.recoveryState`
- `Restriction.probationTier`
- `Restriction.nextEligibleLiftAt`
- `Restriction.lastLiftDecisionAt`
- `Restriction.reinstatementCount`

#### analytics 이벤트 후보
- `restriction_lift_scheduled`
- `restriction_requirement_completed`
- `restriction_lifted`
- `restriction_lifted_with_probation`
- `restriction_reinstated`
- `restriction_appeal_upheld`
- `restriction_appeal_rejected`

핵심 분석 포인트:
- 제한 유형별 복귀율 / 재위반율
- `requirement_completed -> lifted` 전환 시간
- probation tier별 재에스컬레이션 편차
- 오탐성 해제(`manual_lift`, `appeal_upheld`) 비중
- 사용자 안내 확인(`acknowledge`) 이후 개선율

### 109.12 오픈 질문 / 가정
- 가정: MVP에서는 `temporary_suspend`, `chat_only`, `listing_only`, `trust_limited`에 우선 적용하고, 세부 requirement workflow는 단일 체크리스트형으로 시작한다.
- 가정: probation은 사용자 본인과 운영자에게만 보이고 공개 프로필에는 직접 노출하지 않는다.
- 오픈 질문: policy re-acknowledge만으로 충분한 케이스와 identity re-verify가 필요한 케이스의 경계는 anti-abuse 정책과 함께 별도 세분화가 필요하다.
- 오픈 질문: 자동 만료 해제 직전 사용자 확인(step-up confirmation)을 둘지 여부는 전환율 저하와 재위반 감소 효과를 보고 결정한다.

## 110. 후기 자격 스냅샷(Review Eligibility Snapshot) / 작성 가능 창 / 봉인(freeze) canonical 계약
### 110.1 목표
- 후기 작성 가능 여부를 단순히 `completed면 가능`으로 보지 않고, `최종 거래 결과`, `노쇼 판정`, `분쟁/appeal`, `제재/콘텐츠 검수`, `후기 공개 지연`을 하나의 canonical gate로 통합한다.
- 내 거래/프로필/후기 작성 화면/운영도구/API/analytics가 서로 다른 조건으로 후기 CTA를 노출하지 않도록 공통 vocabulary를 정의한다.
- 후기의 **작성 자격(eligibility)**, **작성 가능 기간(window)**, **공개/집계 봉인(freeze)**, **수정 가능 상태(mutability)** 를 분리해 모델링한다.

### 110.2 적용 범위
이 섹션은 아래 객체/화면에 공통 적용된다.
- `tradeCaseSummary`
- `finalOutcome`
- `tradeTrustSnapshot`
- 후기 작성/수정/신고/비노출 흐름
- 운영자의 no-show/dispute/restriction/review moderation 판정

적용 제외:
- 일반 프로필 소개, 매물 본문, 채팅 메시지에 대한 콘텐츠 검수 자체는 108장 기준을 따른다.
- 후기 노출 문구/SEO 범위는 기존 공개 프로필/후기 정책 섹션을 따르되, “지금 이 후기를 써도 되는가/보여도 되는가”의 canonical 판단은 본 섹션을 우선한다.

### 110.3 핵심 vocabulary
| 필드 | 후보 값 | 설명 |
|---|---|---|
| `reviewEligibilityState` | `not_eligible` / `eligible_to_write` / `write_blocked_pending_outcome` / `write_blocked_by_policy` / `write_closed_expired` / `already_written` | 현재 사용자가 상대 후기를 작성할 수 있는지 |
| `reviewWindowState` | `not_opened` / `open` / `closing_soon` / `closed` / `reopened_exceptionally` | 후기 작성 가능 기간 상태 |
| `publicationFreezeReasonCode` | `pending_counterparty_confirmation` / `final_outcome_under_review` / `no_show_case_open` / `appeal_open` / `content_moderation_pending` / `restriction_lock` / `legal_hold` | 후기 공개/집계 봉인 사유 |
| `reviewMutabilityState` | `not_created` / `editable_grace_period` / `locked_visible` / `locked_hidden` / `reopen_edit_allowed_by_staff` | 후기 수정 가능 상태 |
| `reviewCounterpartyPolicy` | `mutual` / `single_sided_allowed` / `staff_overridden` | 후기 생성이 상호형인지 단독 허용인지 |
| `trustAggregationState` | `excluded` / `included_provisional` / `included_final` / `removed_after_revision` | trust 집계 반영 상태 |

원칙:
- `reviewEligibilityState`는 “작성 가능 여부”를 뜻하고, `publicationFreezeReasonCode`는 “작성은 됐지만 아직 공개/집계 보류”까지 포함하는 별도 축이다.
- UI는 내부 vocabulary 전체를 노출하지 않고 `후기 작성 가능`, `운영 검토 후 공개`, `분쟁 종료 후 작성 가능`, `작성 기간 종료` 같은 사용자 카피로 매핑한다.

### 110.4 후기 자격 생성 조건
후기 자격은 아래 조건을 모두 만족할 때 생성된다.
1. 해당 사용자가 `finalOutcome`의 당사자다.
2. `reviewEligibilityOutcome`이 후기 허용 결과로 판정되었다.
3. 동일 outcome/tradeCase에 대해 해당 사용자의 기존 후기가 없다.
4. 계정/기능 제한이 후기 작성 자체를 막는 상태가 아니다.
5. 분쟁/appeal/open no-show case가 “작성 보류”로 판정되지 않았다.

기본 생성 트리거:
- `finalOutcomeType=completed` 이고 결과가 확정되면 양측 후기 자격 생성
- `finalOutcomeType=partial_completed` 는 실제 상대평가가 의미 있는 범위일 때만 생성
- `finalOutcomeType=cancelled` 는 기본적으로 후기 자격 미생성
- `finalOutcomeType=no_show_confirmed` 는 정책상 후기 허용 여부를 별도 판정하며 기본안은 **제한적 허용**이다

### 110.5 결과별 기본 정책 매트릭스
| final outcome / 사건 결과 | 후기 작성 자격 | 공개/집계 기본안 | 비고 |
|---|---|---|---|
| `completed` | 양측 허용 | 일반 공개/집계 | 정상 케이스 |
| `partial_completed` | 양측 허용 후보 | 일반 공개/집계 | 잔여 수량 분쟁 없을 때 |
| `cancelled_mutual` | 기본 비허용 | 없음 | 거래 성립 전 상호 취소 |
| `cancelled_due_to_no_response` | 기본 비허용 | 없음 | 단순 무응답은 후기보다 trust/log 우선 |
| `no_show_confirmed` | 정책상 허용 가능 | 공개는 제한적/문구 가이드 필요 | 보복성 서술 강화 방지 필요 |
| `dispute_resolved_completed` | 양측 허용 | 필요 시 공개 지연 후 집계 | 운영 판정 반영 |
| `dispute_resolved_not_completed` | 기본 비허용 | 없음 | 사실상 거래 미성사 |
| `insufficient_evidence` | 기본 비허용 | 없음 | 신뢰 집계는 운영사건 기준만 반영 |

기본안:
- 실제 인도/수령이 있었던 결과만 후기 허용을 우선한다.
- `no_show_confirmed`는 커뮤니티 신뢰상 의미가 있으므로 완전 금지보다, **템플릿/제한형 후기** 또는 운영 가이드가 붙은 자유서술 후기 허용을 우선 검토한다.

### 110.6 후기 작성 가능 창(window) 기본안
- 후기 window는 `reviewEligibleAt`에 열린다.
- 기본 기간은 **7일**이며, `closing_soon`은 만료 24시간 전부터 적용한다.
- `appeal_open`, `content_moderation_pending`, `restriction_lock` 상태에서는 window가 열려 있어도 `write_blocked_*`로 전환될 수 있다.
- 운영 예외로 reopening 시 `reviewWindowState=reopened_exceptionally`를 사용한다.

후보 필드:
- `reviewEligibleAt`
- `reviewWindowClosesAt`
- `reviewWindowExtendedAt`
- `reviewWindowExtensionReasonCode`

### 110.7 봉인(freeze)과 작성 차단의 차이
후기 흐름은 아래 두 축을 분리한다.

1. **write gate**
   - 지금 사용자가 후기를 새로 쓸 수 있는가
2. **publication/trust freeze gate**
   - 이미 써진 후기를 지금 공개/집계할 수 있는가

예시:
- 양측 모두 후기 작성 완료, 그러나 one-side appeal open → `already_written` + `publicationFreezeReasonCode=appeal_open`
- 후기는 작성 가능하지만 상대 no-show case adjudication 대기 → `write_blocked_pending_outcome`
- 작성은 끝났고 공개는 되지만 trust 합산은 provisional → `trustAggregationState=included_provisional`

### 110.8 후기 수정/잠금 규칙
- 후기 생성 직후 `editable_grace_period`를 부여하고 기본안은 **15분 1회 수정 허용**이다.
- grace period 종료 후에는 `locked_visible` 또는 `locked_hidden`으로 전환된다.
- 콘텐츠 검수/신고로 비노출되면 공개상태와 무관하게 수정 자동 허용으로 바꾸지 않는다. 기본은 `locked_hidden`이다.
- 운영자가 재작성/정정 기회를 주는 경우에만 `reopen_edit_allowed_by_staff`를 사용한다.

금지 규칙:
- 분쟁 결과가 뒤집혔다고 해서 기존 후기 원문을 조용히 편집해 덮어쓰지 않는다. 필요한 경우 숨김 + 재작성 window 재오픈으로 처리한다.
- 상대 후기 열람 후 보복 수정이 가능해지지 않도록, grace period 이후 방향 전환은 운영 예외 없이는 금지한다.

### 110.9 no-show / dispute / restriction과의 연결 규칙
- `noShowCase`가 열려 있으면 기본적으로 `publicationFreezeReasonCode=no_show_case_open`을 부여한다.
- 최종 판정이 `no_show_confirmed`로 끝나고 정책상 후기 허용이면, 양식 가이드 또는 reason-tag 기반 후기를 우선 허용한다.
- `appealState`가 열려 있는 동안에는 이미 작성된 후기의 trust 집계를 `included_provisional` 또는 `excluded`로 전환할 수 있어야 한다.
- 후기 작성자가 `temporary_suspend`, `chat_only`, `review_write_blocked` 등 제한 상태면 `write_blocked_by_policy`를 사용한다.
- 후기 자체가 108장 콘텐츠 검수 케이스로 들어가면 review object는 유지하되 `publicationFreezeReasonCode=content_moderation_pending`을 부여한다.

### 110.10 canonical read model: `reviewEligibilitySnapshot`
후기 관련 화면/API는 아래 read model을 기준으로 같은 해석을 써야 한다.

```json
{
  "tradeCaseId": "tc_123",
  "viewerUserId": "user_123",
  "targetUserId": "user_456",
  "reviewEligibilityState": "eligible_to_write",
  "reviewWindowState": "open",
  "reviewCounterpartyPolicy": "mutual",
  "publicationFreezeReasonCodes": [],
  "reviewMutabilityState": "not_created",
  "trustAggregationState": "excluded",
  "reviewEligibleAt": "2026-03-14T08:00:00+09:00",
  "reviewWindowClosesAt": "2026-03-21T08:00:00+09:00",
  "nextBestAction": "write_review",
  "userFacingHintCode": "REVIEW_AVAILABLE_NOW"
}
```

필수 성질:
- 동일 tradeCase + viewer 조합당 1개 snapshot
- 프로필, 내 거래, 후기 작성 진입, 운영 상세가 같은 snapshot을 바라봄
- outcome/no-show/appeal/restriction/review moderation 변경 시 재계산 가능

### 110.11 화면 파생 요구사항
#### 내 거래 / 거래 상세
- 후기 CTA는 `reviewEligibilityState` 기준으로만 노출한다.
- “작성 가능”과 “작성했지만 공개 보류”를 구분해 배지/문구를 다르게 표시한다.
- 만료 임박 시 `closing_soon` 배지와 종료 시각을 함께 노출한다.

#### 프로필 / 후기 탭
- 공개 후기 수는 `trustAggregationState in (included_provisional, included_final)` 범위와 별도 계산 가능해야 한다.
- appeal/open moderation 상태의 후기는 즉시 카운트에 넣지 않거나 provisional 구분이 가능해야 한다.

#### 운영 백오피스
- 어떤 사건 때문에 후기 write/publication/trust가 막혔는지 한 눈에 보여야 한다.
- 운영자가 reopening 예외를 줄 때 변경 전/후 snapshot diff를 기록해야 한다.

### 110.12 API / DB 파생 기준
API 후보:
- `GET /trade-cases/{tradeCaseId}/review-eligibility`
- `POST /trade-completions/{completionId}/reviews`
- `PATCH /reviews/{reviewId}` (grace period 또는 staff override 범위 내)
- `POST /admin/reviews/{reviewId}/reopen-window`

DB/엔티티 후보:
- `ReviewEligibilitySnapshot` materialized read model
- `review_eligibility_state`
- `review_window_closes_at`
- `publication_freeze_reason_code[]` 또는 linkage table
- `review_mutability_state`
- `trust_aggregation_state`
- `window_reopened_by_admin_id` / `window_reopened_reason_code`

정합성 규칙:
- `already_written` 상태인데 review row가 없으면 안 된다.
- `write_closed_expired`는 `reviewWindowClosesAt < now`를 만족해야 한다.
- `publicationFreezeReasonCode`가 비어 있지 않아도 write gate는 열려 있을 수 있다.

### 110.13 analytics / audit 기준
이벤트 후보:
- `review_eligibility_opened`
- `review_window_closing_soon_shown`
- `review_write_blocked`
- `review_publication_frozen`
- `review_window_reopened_by_staff`
- `review_trust_aggregation_revised`

핵심 분석 포인트:
- outcome type별 후기 작성률
- 분쟁/노쇼 케이스의 후기 window 차단 비율
- grace period 내 수정률과 moderation rate 상관관계
- 공개 보류 후 최종 공개/비노출 전환 비율

### 110.14 오픈 질문 / 기본 가정
- `no_show_confirmed` 결과에 자유서술 후기를 허용할지, tag/template형 후기만 허용할지 최종 정책 확정 필요
- provisional trust 반영을 사용자 프로필에 노출할지 내부 집계에만 둘지 결정 필요
- 후기 window 재오픈을 사용자 self-serve로 허용할지, staff override only로 제한할지 결정 필요

기본 가정:
- MVP는 `completed`, `partial_completed`, 일부 `no_show_confirmed`에 한해 후기 허용
- 후기 수정은 15분/1회 grace period만 허용
- appeal 또는 moderation open 상태에서는 공개/집계 봉인을 우선 적용

## 111. 거래 작업 기한 / 응답 마감 / Action Deadline Snapshot canonical 계약
### 111.1 목표
- 문의 응답, 예약 수락, 장소 재확인, 거래완료 확인, 노쇼 소명, 운영 추가자료 제출처럼 `언제까지 무엇을 해야 하는지`가 중요한 액션을 개별 객체마다 제각각 계산하지 않도록 하나의 canonical deadline 모델로 묶는다.
- 홈, 내 거래, 채팅, 알림함, 운영 큐, 배치 잡, analytics가 모두 동일한 `마감 시각`, `현재 긴급도`, `기한 경과 후 결과`를 해석하도록 한다.
- 사용자가 “지금 가장 급한 일”을 한눈에 볼 수 있게 하고, 운영자는 SLA와 사용자 액션 지연을 같은 구조로 관찰할 수 있게 한다.

### 111.2 적용 범위
`ActionDeadlineSnapshot`은 아래 deadline family에 적용한다.
- 신규 문의 첫 응답 마감
- 가격 제안/역제안 응답 마감
- 예약 제안 수락/거절 마감
- 예약 변경안(reschedule/location change) 재확인 마감
- 거래 직전 도착 확인/지연 알림 권장 시점
- 거래완료 요청 확인/이의제기 마감
- 노쇼 claim 반박/소명 제출 마감
- 분쟁 추가자료 제출 마감
- 운영 제한 해제 조건 제출/재동의 마감
- 후기 작성 가능 창의 마감(읽기용 secondary deadline)

원칙:
- 모든 deadline이 별도 테이블을 가져야 하는 것은 아니며, 원천 aggregate에서 계산한 값을 `ActionDeadlineSnapshot` read model로 수렴시킨다.
- 한 trade case에 복수 deadline이 동시에 존재할 수 있으므로, `currentPrimaryDeadline`과 `secondaryDeadlines[]`를 분리한다.

### 111.3 핵심 vocabulary
#### `deadlineType`
- `lead_first_response_due`
- `offer_response_due`
- `reservation_response_due`
- `reschedule_ack_due`
- `location_ack_due`
- `arrival_checkin_recommended`
- `completion_response_due`
- `no_show_statement_due`
- `dispute_evidence_due`
- `policy_reconsent_due`
- `restriction_recovery_due`
- `review_window_close`

#### `deadlineState`
- `scheduled`: 미래 마감이 설정됨
- `due_now`: 현재 액션이 즉시 필요함
- `approaching`: 임박했으나 아직 grace 내
- `overdue_in_grace`: 공식 기한은 지났지만 grace/보호 구간 안
- `expired`: 기한 종료, 후속 자동전이 또는 action closed
- `resolved`: 사용자가 기한 내 필요한 액션을 완료함
- `suppressed`: 상위 상태 변화로 더 이상 의미 없는 deadline
- `replaced`: 더 최신 deadline에 의해 대체됨

#### `deadlineUrgencyTier`
- `critical`
- `high`
- `medium`
- `low`
- `informational`

#### `expiryConsequenceType`
- `none`
- `soft_nudge_only`
- `auto_expire_proposal`
- `auto_confirm_completion`
- `close_write_window`
- `mark_thread_stale`
- `escalate_to_review_queue`
- `downgrade_priority`
- `lift_offer_priority`
- `freeze_outcome`

#### `autoTransitionPolicy`
- `manual_only`
- `auto_after_expiry`
- `auto_after_expiry_with_grace`
- `auto_after_multi_reminders`
- `auto_blocked_when_case_locked`

### 111.4 canonical object 정의
```json
{
  "actionDeadlineId": "adl_123",
  "tradeCaseId": "tc_123",
  "threadId": "tt_123",
  "deadlineType": "completion_response_due",
  "ownerPartyRole": "buyer",
  "counterpartyRole": "seller",
  "sourceAggregateType": "trade_completion",
  "sourceAggregateId": "comp_123",
  "deadlineState": "approaching",
  "deadlineUrgencyTier": "high",
  "dueAt": "2026-03-15T21:00:00+09:00",
  "graceEndsAt": "2026-03-15T21:30:00+09:00",
  "expiryConsequenceType": "auto_confirm_completion",
  "autoTransitionPolicy": "auto_after_expiry_with_grace",
  "resolvedAt": null,
  "suppressedReasonCode": null,
  "nextBestAction": "confirm_completion",
  "userFacingHintCode": "CONFIRM_TRADE_RESULT_TODAY",
  "reminderPlan": {
    "lastReminderSentAt": "2026-03-15T19:00:00+09:00",
    "reminderCount": 1,
    "maxReminderCount": 2
  }
}
```

필수 성질:
- 동일 source aggregate + deadlineType + ownerPartyRole 조합당 동시 active snapshot은 1개만 허용
- `resolved`, `expired`, `suppressed`, `replaced`는 terminal read state로 취급 가능
- 홈/내 거래/알림은 snapshot projection을 직접 사용하고, 원천 객체별 계산 로직을 중복 구현하지 않는다.

### 111.5 생성 규칙
| source event | 생성되는 deadline | owner | 기본 due rule | expiry consequence |
|---|---|---|---|---|
| 새 문의 생성 | `lead_first_response_due` | 매물 작성자 | 생성 후 N분/시간 | `mark_thread_stale` 또는 우선순위 하락 |
| offer/counter-offer 생성 | `offer_response_due` | 상대 제안 수신자 | 제안 만료시각 | `auto_expire_proposal` |
| 예약 제안 생성 | `reservation_response_due` | 상대방 | min(제안+12h, 예약-2h) | `auto_expire_proposal` |
| reschedule/location change 제안 | `reschedule_ack_due` 또는 `location_ack_due` | 상대방 | 제안 후 지정 시간 | `escalate_to_review_queue` optional |
| 완료 요청 생성 | `completion_response_due` | 상대방 | 완료 요청 + 24h(가정) | `auto_confirm_completion` |
| no-show claim 생성 | `no_show_statement_due` | 피신고 당사자 | claim + 24h/48h | `freeze_outcome` 또는 claim 단독 심사 |
| 운영 추가자료 요청 | `dispute_evidence_due` | 요청받은 당사자 | 운영 설정 기한 | `escalate_to_review_queue` / `freeze_outcome` |
| 후기 창 오픈 | `review_window_close` | 후기 작성 가능 당사자 | reviewWindowClosesAt | `close_write_window` |

원칙:
- `review_window_close`는 행동 유도에 쓰이지만, thread primary CTA보다 낮은 secondary deadline으로 분류한다.
- 같은 사건에서 더 강한 새 deadline이 생기면 이전 deadline은 `replaced` 또는 `suppressed`로 종료한다.

### 111.6 상태 전이 규칙
기본 전이:
- `scheduled -> approaching -> due_now -> overdue_in_grace -> expired`
- 사용자가 제시간에 액션 완료 시 어떤 단계에서든 `resolved`
- 상위 사건이 닫히거나 다른 상태로 넘어가 더 이상 deadline 의미가 없어지면 `suppressed`
- 새 제안/재전송/재스케줄로 기한이 갱신되면 기존 snapshot은 `replaced`, 새 snapshot 생성

세부 규칙:
- `approaching` 진입 시점은 deadlineType마다 다를 수 있으나, 사용자 노출용 urgency tier 계산은 공통 함수로 관리한다.
- `due_now`는 단순 시각 비교만이 아니라 홈/내 거래에서 최상단으로 올려야 하는 상태다.
- `overdue_in_grace`는 사용자에게 `아직 처리 가능하지만 곧 자동 처리됨`을 설명하는 구간이다.
- `expired` 이후 self-serve action이 막히더라도 운영자는 staff override로 reopening 가능하다.

### 111.7 urgency 계산 원칙
`deadlineUrgencyTier`는 절대 시각보다 `도메인 중요도 x 남은 시간 x 결과 강도`의 조합으로 계산한다.

우선순위 기본안:
1. `completion_response_due`, `no_show_statement_due`, `dispute_evidence_due`
2. `reservation_response_due`, `location_ack_due`, `reschedule_ack_due`
3. `lead_first_response_due`, `offer_response_due`
4. `review_window_close`
5. `arrival_checkin_recommended`

세부 원칙:
- 동일 시간대라면 `auto_confirm_completion`이 `soft_nudge_only`보다 우선
- `review_window_close`는 임박해도 거래 실행 deadline보다 우선하지 않음
- 운영 큐에서는 P1/P2 사건의 `dispute_evidence_due`가 홈 UX보다 더 높은 urgency로 재정렬될 수 있음

### 111.8 화면 파생 요구사항
#### 홈 / 내 거래
- 각 trade thread는 `currentPrimaryDeadline` 1개와 `secondaryCount`를 가져야 한다.
- 카드에는 `남은 시간`, `기한 후 결과`, `주 액션`을 함께 보여준다.
- 예시: `오늘 21:00까지 완료 확인 필요 · 미응답 시 자동 확정`
- `due_now`/`critical`은 unread보다 우선 정렬한다.

#### 채팅 / 거래 상세
- 타임라인 상단 sticky action bar는 active deadline 기준으로만 CTA를 노출한다.
- deadline이 `overdue_in_grace`면 경고색 + grace 종료 시각을 같이 보여준다.
- `suppressed` 또는 `resolved`된 deadline은 히스토리 이벤트로만 남기고 CTA에서는 제거한다.

#### 알림함
- 알림 객체 자체가 아니라 `deadline reminder`와 `state change`를 구분한다.
- 같은 deadline에 대한 반복 리마인드는 groupingKey를 공유해 묶음 표시한다.

#### 운영 백오피스
- 사용자 지연과 운영 지연을 분리해서 보여야 한다.
- `dispute_evidence_due`가 만료된 사건은 운영 큐에서 자동 필터링/강조할 수 있어야 한다.

### 111.9 API / DB 파생 기준
API 후보:
- `GET /me/trade-deadlines`
- `GET /trade-threads/{threadId}/deadlines`
- `POST /deadlines/{actionDeadlineId}/dismiss` (informational only)
- `POST /admin/deadlines/{actionDeadlineId}/reopen`
- `POST /admin/deadlines/{actionDeadlineId}/extend`

응답 필드 후보:
- `currentPrimaryDeadline`
- `secondaryDeadlines[]`
- `deadlineUrgencyTier`
- `dueAt`
- `graceEndsAt`
- `expiryConsequenceType`
- `nextBestAction`
- `userFacingHintCode`

DB/read model 후보:
- `ActionDeadlineSnapshot`
- `action_deadline_state`
- `deadline_type`
- `deadline_owner_party_role`
- `due_at`
- `grace_ends_at`
- `expiry_consequence_type`
- `auto_transition_policy`
- `replaced_by_deadline_id`
- `suppressed_reason_code`

정합성 규칙:
- `expired`인데 `dueAt > now`이면 안 된다.
- `resolvedAt`가 있으면 `deadlineState`는 `resolved`여야 한다.
- `replaced` 상태는 `replacedByDeadlineId`를 가져야 한다.
- `completion_response_due` active snapshot은 동일 completion에 대해 1개 초과 생성되면 안 된다.

### 111.10 notification / automation 기준
- reminder는 snapshot 기반으로만 발송하고, 원천 aggregate별 개별 스케줄러가 임의로 사용자 알림을 중복 생성하지 않는다.
- `maxReminderCount`, `minimumReminderInterval`, `quietHourBypassPolicy`는 deadlineType별 정책 테이블로 분리 가능해야 한다.
- `completion_response_due`와 `reservation_response_due`는 거래 필수 알림으로 분류하고 quiet hours 우회 가능 여부를 별도 정책으로 둔다.
- `review_window_close`는 quiet hours를 우회하지 않는 기본안을 사용한다.

### 111.11 analytics / audit 기준
이벤트 후보:
- `deadline_created`
- `deadline_entered_approaching`
- `deadline_reminder_sent`
- `deadline_resolved_on_time`
- `deadline_resolved_in_grace`
- `deadline_expired`
- `deadline_auto_transition_executed`
- `deadline_reopened_by_staff`
- `deadline_extended_by_staff`

핵심 분석 포인트:
- deadlineType별 on-time resolution rate
- `approaching` 진입 후 실제 행동 전환율
- grace 구간 의존 비율과 자동전이 비율
- reminder 수와 거래 완료율/피로도 상관관계
- 서버/카테고리/신뢰레벨별 응답 지연 편차

### 111.12 오픈 질문 / 기본 가정
오픈 질문:
- `lead_first_response_due`를 seller/buyer 경험 보호를 위해 얼마나 강하게 랭킹/노출 패널티와 연결할지 확정 필요
- `location_ack_due` 만료 시 즉시 no-show 분쟁 후보로 올릴지, 단순 stale 상태로 둘지 결정 필요
- 사용자가 deadline 연장을 self-serve로 요청할 수 있는 범위를 어디까지 허용할지 정책 확정 필요

기본 가정:
- MVP는 `lead_first_response_due`, `reservation_response_due`, `completion_response_due`, `review_window_close`를 우선 지원
- `no_show_statement_due`, `dispute_evidence_due`는 운영/분쟁 기능 rollout과 함께 활성화
- 홈/내 거래는 `currentPrimaryDeadline` 기반 정렬을 canonical 동작으로 사용


## 112. 거래 가능 시간 Slot / 예약 제안 가능 범위 / availability-to-booking canonical 계약
### 112.1 목적
- 사용자 프로필의 `가능 시간`, 매물 상세의 `거래 가능 시간`, 예약 제안 시점의 `실제 예약 가능 여부`, 당일 실행 카드의 `지연/재일정`, 노쇼 분쟁의 `늦었지만 허용 범위였는지`를 하나의 시간 모델로 연결한다.
- 단순 자유서술 `availableTimeText`만으로는 예약 가능 판단, 알림, 자동 검증, 운영 판정이 일관될 수 없으므로 canonical slot 계약을 정의한다.
- 홈/목록/상세/채팅/예약/API/DB/analytics가 동일한 시간 vocabulary를 사용해 `언제 거래 가능하며, 지금 예약을 제안해도 되는가`를 같은 방식으로 해석하도록 한다.

### 112.2 핵심 개념
| 용어 | 의미 |
|---|---|
| `bookingAvailabilitySnapshot` | 특정 사용자/매물/거래 스레드 기준 예약 가능 시간 상태를 요약한 canonical read model |
| `availabilitySourceType` | 가용 시간의 원천 (`profile_default` / `listing_override` / `thread_negotiated` / `reservation_locked`) |
| `availabilitySlot` | 예약 제안에 사용할 수 있는 개별 시간 구간 |
| `slotDisclosureLevel` | 슬롯 공개 수준 (`coarse` / `bookable_window` / `exact_time`) |
| `bookabilityState` | 현재 시점에서 예약 제안 가능한지 (`bookable` / `bookable_with_warning` / `requires_confirmation` / `temporarily_unavailable` / `blocked`) |
| `bookingConflictDisposition` | 슬롯 충돌 처리 방식 (`soft_warn` / `reject` / `counter_suggest` / `manual_review`) |
| `leadTimePolicy` | 지금부터 몇 분/시간 이후 슬롯부터 제안 가능한지 규정하는 정책 |
| `bufferPolicy` | 슬롯 전후 준비시간/이동시간/정리시간 버퍼 규칙 |
| `blackoutReasonCode` | 특정 시간대 예약 불가 사유 코드 |
| `executionWindowPolicy` | 예약 확정 후 실제 거래 수행을 허용하는 시간 오차 규칙 |

원칙:
- `availableTimeText`는 보조 표시용일 뿐, 예약 가능 판정의 canonical source가 되어서는 안 된다.
- 예약 가능 여부는 `프로필 기본 가용시간`, `매물별 override`, `이미 잡힌 예약`, `제한/정책`, `당일 변경/재확인 필요 상태`를 함께 반영해야 한다.
- 화면에는 단순 문구로 보이더라도 내부적으로는 slot 기반 read model을 사용해 예약/재일정/노쇼 판단을 일관되게 만든다.

### 112.3 시간 정보 계층
| 계층 | 의미 | 예시 | 주요 사용처 |
|---|---|---|---|
| `profile_default` | 사용자의 일반 거래 가능 패턴 | `평일 20:00~24:00` | 프로필, 홈 추천, 기본 제안 |
| `listing_override` | 특정 매물에만 적용되는 거래 가능 시간 | `오늘만 21:00 이후 가능` | 상세, 목록 라벨, 예약 후보 |
| `thread_negotiated` | 특정 상대와 채팅 중 협의된 잠정 가능 시간 | `내일 22시쯤 가능` | 채팅 quick slot 제안 |
| `reservation_locked` | 실제 예약 확정으로 잠긴 시간 | `3/15 21:00~21:30` | 내 거래, 충돌검사, 알림 |

우선순위 원칙:
1. `reservation_locked`
2. `thread_negotiated`
3. `listing_override`
4. `profile_default`

즉, 사용자의 기본 가용시간이 있더라도 이미 예약이 확정된 시간과 충돌하면 해당 slot은 `bookable`로 취급하지 않는다.

### 112.4 availability slot 구조
| 필드 | 필수 여부 | 설명 |
|---|---|---|
| `slotId` | 필수 | 슬롯 식별자 또는 read model 내 안정 ID |
| `sourceType` | 필수 | `availabilitySourceType` |
| `startAt` | 필수 | 슬롯 시작 시각(시간대 포함) |
| `endAt` | 필수 | 슬롯 종료 시각 |
| `isRecurring` | 필수 | 반복 슬롯 여부 |
| `recurrenceRule` | 선택 | 반복 규칙(요일/주기) |
| `slotDisclosureLevel` | 필수 | 공개 수준 |
| `bookabilityState` | 필수 | 예약 가능 상태 |
| `leadTimeMinMinutes` | 선택 | 최소 선행 제안 시간 |
| `bufferBeforeMinutes` | 선택 | 슬롯 시작 전 준비 버퍼 |
| `bufferAfterMinutes` | 선택 | 슬롯 종료 후 정리 버퍼 |
| `blackoutReasonCode` | 선택 | 예약 불가 사유 |
| `requiresConfirmation` | 필수 | 현재 시점에 추가 확인이 필요한지 |
| `slotConfidenceTier` | 필수 | `declared` / `recently_confirmed` / `historically_reliable` / `stale` |

정합성 규칙:
- `endAt <= startAt` 인 slot은 허용되지 않는다.
- `bookabilityState=blocked`이면 `blackoutReasonCode` 또는 정책 근거가 있어야 한다.
- `requiresConfirmation=true`이면 exact 예약 확정 전 사용자 명시 확인 UX가 필요하다.

### 112.5 공개 수준 정의
| `slotDisclosureLevel` | 의미 | 공개 surface | 예시 |
|---|---|---|---|
| `coarse` | 대략적 시간대만 공개 | 목록/상세/프로필 | `평일 저녁 가능` |
| `bookable_window` | 예약 가능한 window 공개 | 상세/채팅 quick pick | `오늘 20:00~23:00 가능` |
| `exact_time` | 실제 제안/확정용 exact slot | 채팅/예약/내 거래 | `3/15 21:00~21:30` |

원칙:
- 공개 프로필과 공개 매물에서는 기본적으로 `coarse` 또는 `bookable_window`까지만 사용한다.
- `exact_time`은 채팅 진입 이후 또는 예약 제안/확정 단계에서만 노출하는 것을 기본안으로 한다.
- 푸시/잠금화면 알림에는 `exact_time`이 있더라도 상황에 따라 coarse화된 문구를 우선 사용한다.

### 112.6 예약 가능 판정(`bookabilityState`) 정의
| 상태 | 의미 | 대표 원인 | 사용자 노출 예시 |
|---|---|---|---|
| `bookable` | 바로 제안 가능 | 슬롯 활성, 충돌 없음 | `예약 제안 가능` |
| `bookable_with_warning` | 제안 가능하지만 주의 필요 | 리드타임 짧음, 최근 변동 많음 | `가능하지만 확인이 필요해요` |
| `requires_confirmation` | 상대 또는 작성자 재확인 후 가능 | 오래된 가용시간, 직전 재일정 | `상대 확인 후 예약하세요` |
| `temporarily_unavailable` | 현재는 제안 불가 | 다른 예약과 충돌, blackout | `이 시간은 불가` |
| `blocked` | 정책/제한으로 차단 | 제재, 거래중지, 계정 상태 | `현재 예약 제안 불가` |

판정 우선순위:
1. 정책 차단/제한
2. 이미 확정된 예약과의 충돌
3. 최소 리드타임 위반
4. 버퍼 침범 여부
5. freshness/confirmation requirement

### 112.7 리드타임 / 버퍼 / 충돌 규칙
기본 가정:
- 예약 제안 최소 리드타임은 `10분` 이상을 기본안으로 둔다.
- 오프라인 접선은 이동 필요성을 고려해 `in_game`보다 더 긴 기본 리드타임을 가질 수 있다.
- 같은 사용자의 두 예약은 `bufferBeforeMinutes`, `bufferAfterMinutes`를 포함한 시간 구간이 겹치면 충돌로 본다.

충돌 처리 규칙:
- `bookingConflictDisposition=reject`면 슬롯 제안 자체를 차단한다.
- `soft_warn`면 사용자는 경고를 보고 계속 제안할 수 있으나 상대 확인이 필요해질 수 있다.
- `counter_suggest`면 시스템이 인접 가능한 slot을 추천한다.
- `manual_review`는 일반 사용자 예약에서는 거의 쓰지 않고 운영/특수 제약용으로만 둔다.

### 112.8 예약 제안 / 재일정과의 연결 규칙
- 예약 제안 카드는 canonical `availabilitySlot`을 기반으로 생성해야 하며, 자유텍스트만으로 예약 제안을 확정해서는 안 된다.
- 재일정은 기존 `reservation_locked`를 대체하는 새 slot 제안으로 해석한다.
- accepted reschedule이 확정되면 이전 slot은 `replaced` 성격으로 남기고, 새 slot이 canonical 실행 slot이 된다.
- 상대가 `정확 시간 미확정` 상태로만 응답했다면 `thread_negotiated`는 갱신되지만 `reservation_locked`는 생성되지 않는다.

### 112.9 당일 실행 / 지연 / 노쇼 판정과의 연결
- 당일 실행 카드의 `도착했어요`, `5분 늦어요`, `재일정 요청`, `노쇼 신고`는 마지막 `reservation_locked` slot과 `executionWindowPolicy`를 기준으로 판단해야 한다.
- 예를 들어 `scheduledAt=21:00`, `grace=15분`이면 21:12 도착은 지연이지만 즉시 no-show로 해석하지 않는다.
- `bookable_with_warning` 상태에서 생성된 예약은 노쇼 분쟁 시 운영 참고 신호(`slotConfidenceTier 낮음`)로 남길 수 있다.
- 운영자는 분쟁 시 아래를 함께 재구성해야 한다.
  1. 마지막 확정 slot
  2. 직전 재일정 횟수
  3. slot 변경 시각과 상대 읽음/동의 여부
  4. grace/buffer 정책

### 112.10 화면 파생 요구사항
#### 홈/목록/상세
- exact slot이 아니라 `오늘 저녁 가능`, `주말 오후 가능`, `지금 거래 가능` 같은 coarse 가용 신호를 사용한다.
- 매물 카드에는 가격/상태보다 약한 우선순위지만 거래 전환에 도움이 되는 가용 시간 배지를 노출할 수 있다.

#### 채팅/예약
- quick slot picker는 canonical slot만 선택하게 하거나, 자유입력 시에도 slot 구조화 과정을 거쳐야 한다.
- 사용자가 선택한 slot이 `requires_confirmation`이면 CTA 카피를 `예약 제안`이 아니라 `확인 요청 보내기`로 바꿀 수 있어야 한다.
- 예약 카드에는 `시간`, `지속시간`, `가용 출처`, `재확인 필요 여부`를 포함하는 것이 바람직하다.

#### 내 거래
- 임박 거래 요약은 `정확 시간` 기준으로 정렬하되, 불확실한 건은 `확인 필요` 배지로 구분한다.
- 시간이 지났지만 grace 내인 건과 이미 실질적으로 expired인 건을 시각적으로 다르게 보여야 한다.

### 112.11 API / DB 파생 기준
API 후보:
- `GET /users/{userId}/availability`
- `GET /listings/{listingId}/availability`
- `POST /chats/{chatRoomId}/availability-slots`
- `POST /reservations/{reservationId}/reschedule-slot`
- `GET /me/bookable-slots?listingId=...`

응답 필드 후보:
```json
{
  "availability": {
    "bookabilityState": "bookable_with_warning",
    "sourceType": "listing_override",
    "slots": [
      {
        "slotId": "slot_123",
        "startAt": "2026-03-15T21:00:00+09:00",
        "endAt": "2026-03-15T21:30:00+09:00",
        "slotDisclosureLevel": "exact_time",
        "slotConfidenceTier": "recently_confirmed",
        "requiresConfirmation": false
      }
    ],
    "leadTimePolicy": {
      "minimumMinutes": 10
    },
    "bufferPolicy": {
      "beforeMinutes": 5,
      "afterMinutes": 10
    }
  }
}
```

DB/read model 후보:
- `UserAvailabilityProfile`
- `ListingAvailabilitySnapshot`
- `AvailabilitySlot`
- `ThreadNegotiatedAvailability`
- `ReservationExecutionWindow`

정합성 규칙:
- 하나의 확정 예약은 정확히 하나의 canonical execution slot을 가져야 한다.
- `reservation_locked` slot이 active면 동일 참여자 기준 중복 exact slot 충돌 검사가 필요하다.
- `listing_override`가 만료되었으면 홈/상세에서 stale slot로 표시하거나 제외해야 한다.

### 112.12 analytics / 운영 기준
이벤트 후보:
- `availability_slot_viewed`
- `availability_slot_selected`
- `booking_conflict_warned`
- `booking_slot_confirmation_requested`
- `booking_slot_confirmed`
- `booking_slot_replaced_by_reschedule`
- `availability_stale_detected`

핵심 분석 포인트:
- coarse availability 노출이 채팅 시작률/예약 전환율에 미치는 영향
- `requires_confirmation` 상태 비율과 실제 예약 성사율
- slot 충돌/리드타임 위반으로 인한 이탈률
- 인게임/오프라인 방식별 평균 리드타임과 노쇼율 상관관계

### 112.13 오픈 질문 / 기본 가정
오픈 질문:
- `profile_default`를 얼마나 구조화할지(요일 반복만 우선인지, 날짜별 예외까지 MVP 포함인지) 확정 필요
- 오프라인 거래의 이동시간 버퍼를 사용자가 직접 설정하게 할지, meetingType별 기본값으로 둘지 정책 결정 필요
- exact slot 없이 `오늘 밤 가능`만으로 예약 제안하는 느슨한 모드를 MVP에서 허용할지 결정 필요

기본 가정:
- MVP는 `coarse` + `exact_time` 2계층을 우선 지원하고, `bookable_window`는 read model상 열어두되 점진 도입 가능
- 사용자 기본 가용시간은 주간 반복 패턴 중심으로 시작하고, 날짜별 예외/휴무는 Post-MVP 확장 후보
- 예약 확정 시점에는 반드시 exact execution slot이 canonical source가 되어야 한다

## 113. 거래 실행 준비도(Execution Readiness Snapshot) / 예약 성립과 실제 실행 가능성 구분 canonical 계약

### 113.1 목표
- `예약이 있다`와 `실제로 지금 거래를 진행해도 된다`를 분리해 해석한다.
- 시간 슬롯, 장소, 상대 식별 정보, 최종 조건, 정산 방식, 취소 의사, 수용량(capacity) 상태가 각 화면/운영도구에서 서로 다른 기준으로 섞여 해석되지 않도록 하나의 canonical snapshot으로 수렴시킨다.
- 채팅/내 거래/당일 실행 카드/API/DB/운영이 동일한 준비도 판정 기준을 사용해 CTA, 리마인드, 완료 가능 여부, no-show 억제, 분쟁 우선순위를 결정하도록 한다.

### 113.2 핵심 원칙
- `Reservation.confirmed`는 필요조건일 뿐 충분조건이 아니다.
- 실제 실행 준비도는 아래 선행조건이 모두 최신 스냅샷 기준으로 충족되었을 때에만 `ready_to_execute`로 본다.
  1. exact execution slot 확정
  2. 장소/접선 정보 상호 인지
  3. 상대 식별 정보 또는 recognition method 확정
  4. 최종 거래 조건(deal terms) 확정
  5. 취소/재일정/capacity overcommit blocker 부재
- 준비도 snapshot은 write command의 원천이 아니라 read-model 성격의 canonical summary로 두되, 완료/노쇼/취소/재일정 액션의 server-side guard에 재사용 가능해야 한다.

### 113.3 핵심 vocabulary
| 필드 | 설명 |
|---|---|
| `executionReadinessState` | 실제 실행 가능 준비도 최종 상태 |
| `blockingPrerequisiteType` | 아직 충족되지 않은 선행조건 유형 |
| `readinessDriftReasonCode` | 한 번 준비되었던 거래가 다시 불안정/미확정 상태로 돌아간 이유 |
| `finalGoSignalSource` | 마지막 실행 가능 판정을 만든 근거 출처 |
| `counterpartyAcknowledgementState` | 상대가 최신 실행 정보(시간/장소/조건)를 인지했는지 여부 |
| `executionRiskTier` | 실행 직전 사용자/운영에 보여줄 위험 수준 |
| `completionEligibilityState` | 완료 처리 버튼을 열 수 있는지에 대한 판정 |
| `noShowSuppressionState` | 준비도 미달 때문에 노쇼 판단을 억제해야 하는지 여부 |

### 113.4 `executionReadinessState` enum 후보
| 값 | 의미 | 사용자 노출 예시 |
|---|---|---|
| `not_started` | 예약 또는 실행 준비 흐름이 아직 본격 시작되지 않음 | 거래 준비 전 |
| `needs_confirmation` | 주요 선행조건 일부가 미확정/미인지 상태 | 확인 필요 |
| `ready_soft` | 실행 가능 정보는 있으나 직전 변경/ack 지연 등으로 리스크가 남음 | 거의 준비됨 |
| `ready_to_execute` | 실제 거래 진행에 필요한 필수 정보가 최신 상태로 정렬됨 | 거래 진행 가능 |
| `at_risk` | 시간 임박, ack 누락, 장소 변경, capacity 충돌 등으로 실행 실패 위험 높음 | 실행 위험 높음 |
| `blocked` | 취소 의사, 상태 충돌, 필수 정보 누락 등으로 실행 불가 | 거래 진행 불가 |
| `superseded` | 재일정/취소/다른 확정 흐름으로 대체됨 | 이전 약속 만료 |

원칙:
- `ready_soft`와 `ready_to_execute`를 분리해, 예약은 성립했지만 최신 장소 변경 ack가 누락된 상태 등을 표현한다.
- `at_risk`는 금지 상태가 아니라 강한 재확인/리마인드가 필요한 상태다.
- `blocked`는 완료/노쇼/도착확인 같은 action gating에 직접 반영된다.

### 113.5 선행조건 묶음과 `blockingPrerequisiteType`
| prerequisite | 충족 기준 | 미충족 시 code 후보 |
|---|---|---|
| 시간 슬롯 | exact execution slot 존재, lead time/buffer 위반 없음 | `slot_missing`, `slot_conflicted`, `slot_stale` |
| 장소 | 마지막 확정 장소 snapshot 존재, 심각한 ambiguity 없음 | `location_missing`, `location_ack_missing`, `location_changed_late` |
| 식별 정보 | 캐릭터명/recognition method 등 필요한 식별수단 정렬 | `identity_missing`, `identity_unverified` |
| 거래 조건 | 가격/수량/단위/잔여 처리/결제 방식이 최신 terms snapshot으로 정렬 | `terms_unconfirmed`, `unit_mismatch`, `settlement_unconfirmed` |
| 상태/의사 | cancel intent 없음, reschedule pending 없음 | `cancel_pending`, `reschedule_pending`, `reservation_not_confirmed` |
| capacity | active commitment 상한/중복 시간 충돌 없음 | `capacity_conflict`, `overcommit_risk_blocked` |
| 정책/제한 | restriction, report lock, moderation hold 없음 | `policy_blocked`, `case_locked` |

설계 원칙:
- 단일 blocker만 갖는 것이 아니라 `blockingPrerequisiteTypes[]` 배열로 여러 원인을 함께 보존한다.
- 사용자 화면에는 가장 중요한 1~2개 이유만 축약 노출하고, 운영/API/read model에서는 전체 배열을 유지한다.

### 113.6 `finalGoSignalSource`와 판정 우선순위
| source | 의미 | 우선순위 |
|---|---|---|
| `system_computed` | 서버가 snapshot 조합으로 자동 산출 | 기본 |
| `counterparty_mutual_ack` | 양측이 최신 시간/장소/조건을 명시 확인 | 높음 |
| `seller_last_update` | 판매자(또는 작성자) 마지막 변경이 최신 기준 | 중간 |
| `buyer_last_update` | 구매자 마지막 변경이 최신 기준 | 중간 |
| `moderator_override` | 운영자가 예외적으로 상태 정렬/잠금 해제 | 최상위 예외 |

판정 규칙:
- `counterparty_mutual_ack`가 있으면 같은 snapshot revision 내에서는 `ready_soft -> ready_to_execute` 승격 우선 후보가 된다.
- 단, capacity conflict, cancel intent, policy block이 있으면 mutual ack가 있어도 `blocked` 또는 `at_risk`를 유지한다.
- 운영자 override는 audit log와 만료성 메모를 남기는 예외 경로로만 허용한다.

### 113.7 준비도 drift / downgrade 규칙
준비도는 단방향이 아니라 아래 상황에서 다시 낮아질 수 있다.

| trigger | `readinessDriftReasonCode` 예시 | 결과 |
|---|---|---|
| 예약 시간/장소 늦은 변경 | `late_location_change`, `late_slot_change` | `ready_to_execute -> ready_soft` 또는 `at_risk` |
| 상대 ack 미완료 상태로 실행 시점 임박 | `ack_timeout_near_execution` | `ready_soft -> at_risk` |
| 취소 의사 제기 | `cancel_intent_opened` | `blocked` |
| 재일정 제안 중 | `reschedule_reopened` | `needs_confirmation` |
| 활성 commitment 과다 | `capacity_overcommit_detected` | `at_risk` 또는 `blocked` |
| 운영 잠금/분쟁 생성 | `case_lock_created` | `blocked` |
| deal terms 변경 | `terms_snapshot_replaced` | `needs_confirmation` |

원칙:
- drift는 이전 상태를 덮어쓰는 것이 아니라 timeline/audit 상에서 downgrade event로 보존해야 한다.
- 당일 실행 카드, 알림, 운영 큐는 현재 상태뿐 아니라 최근 downgrade 발생 여부를 함께 해석해야 한다.

### 113.8 `completionEligibilityState` / `noShowSuppressionState`
| 필드 | 값 후보 | 의미 |
|---|---|---|
| `completionEligibilityState` | `eligible`, `eligible_with_warning`, `ineligible` | 완료 처리 CTA를 열 수 있는지 판정 |
| `noShowSuppressionState` | `allow_claim`, `warn_claim`, `suppress_claim` | 노쇼 claim을 허용/경고/억제할지 판정 |

기본 규칙:
- `executionReadinessState=ready_to_execute`면 기본적으로 `completionEligibilityState=eligible` 후보다.
- `ready_soft`, `at_risk` 상태에서는 완료는 `eligible_with_warning` 가능하나, late change나 ack 누락이 있으면 운영 참고 배너를 함께 노출한다.
- `blocked`, `needs_confirmation`이면 기본적으로 `completionEligibilityState=ineligible`다.
- `noShowSuppressionState=suppress_claim` 조건 예시:
  - exact 장소가 마지막 시점까지 상호 확인되지 않음
  - 당일 직전 장소 변경 후 상대 ack 없음
  - reschedule pending 상태가 정리되지 않음
  - cancel intent가 열려 있음
- 이 규칙은 노쇼 남용을 줄이기 위한 가드레일이며, 운영자는 예외적 수동 심사를 할 수 있다.

### 113.9 화면/UX 파생 기준
#### 채팅 / 당일 실행 카드
- 필수 노출 필드:
  - `executionReadinessState`
  - 현재 blocker 1순위
  - 마지막 mutual ack 시각
  - latest terms/location/slot revision 요약
  - 다음 추천 액션(`confirm_location`, `confirm_terms`, `mark_arrived`, `resolve_reschedule`, `cancel_trade` 등)
- CTA 우선순위:
  1. blocker 해소 액션
  2. 상호 확인 액션
  3. 도착/지연/완료 액션
- `blocked` 상태에서는 완료/노쇼 버튼보다 `문제 해결`/`상대 확인 요청` CTA를 우선 노출한다.

#### 내 거래 / 목록 카드
- 카드 요약 라벨 예시:
  - `거래 가능`
  - `장소 확인 필요`
  - `조건 다시 확인 필요`
  - `취소 검토 중`
  - `실행 위험 높음`
- 단순 예약 여부보다 준비도/기한/액션 필요도를 정렬 우선순위에 더 강하게 반영한다.

#### 운영 큐
- 운영은 개별 객체 원문보다 먼저 `tradeCaseSummary + executionReadinessSnapshot` 요약을 본다.
- `at_risk` 장기 유지, 잦은 downgrade, suppress된 no-show claim 시도는 별도 모니터링 후보가 된다.

### 113.10 API / read model 후보
응답 예시:
```json
{
  "executionReadiness": {
    "executionReadinessState": "ready_soft",
    "blockingPrerequisiteTypes": ["location_ack_missing"],
    "counterpartyAcknowledgementState": "awaiting_counterparty_ack",
    "finalGoSignalSource": "seller_last_update",
    "executionRiskTier": "medium",
    "completionEligibilityState": "eligible_with_warning",
    "noShowSuppressionState": "warn_claim",
    "lastMutualAckAt": "2026-03-14T08:20:00+09:00",
    "snapshotRevision": 7
  }
}
```

후보 read model:
- `ExecutionReadinessSnapshot`
- `ExecutionReadinessTimelineEvent`
- `TradeExecutionGuardProjection`

후보 API surface:
- `GET /me/trades` 응답에 `executionReadiness` 포함
- `GET /chats/{chatRoomId}` 응답에 latest readiness summary 포함
- `POST /trade-execution/{tradeThreadId}/ack-readiness`
- `POST /trade-execution/{tradeThreadId}/refresh-readiness` (internal/system 성격 우선)

### 113.11 DB / 이벤트 파생 기준
저장 후보 필드:
- `execution_readiness_state`
- `blocking_prerequisite_types[]`
- `final_go_signal_source`
- `last_mutual_ack_at`
- `readiness_drift_reason_code`
- `completion_eligibility_state`
- `no_show_suppression_state`
- `snapshot_revision`

이벤트 후보:
- `execution_readiness_recomputed`
- `execution_readiness_downgraded`
- `execution_readiness_upgraded`
- `execution_mutual_ack_recorded`
- `execution_blocker_resolved`
- `execution_no_show_claim_suppressed`

정합성 원칙:
- readiness snapshot은 listing/chat/reservation/completion/dispute aggregate를 직접 대체하지 않는다.
- 하지만 완료/노쇼/취소/재일정 write API는 최신 readiness snapshot 또는 equivalent server recomputation을 guard 조건으로 참조할 수 있어야 한다.
- snapshot revision은 최신 terms/location/slot/identity revision과 trace 가능해야 한다.

### 113.12 analytics / 운영 기준
이벤트 후보:
- `execution_readiness_viewed`
- `execution_blocker_seen`
- `execution_mutual_ack_clicked`
- `execution_ready_reached`
- `execution_ready_lost`
- `execution_complete_attempt_blocked`
- `no_show_claim_suppressed`

핵심 분석 포인트:
- `ready_to_execute` 도달 비율과 실제 완료율 상관관계
- `location_ack_missing`, `terms_unconfirmed`, `capacity_conflict` 같은 blocker 분포
- `ready_soft -> at_risk -> completed` 케이스와 노쇼율 관계
- mutual ack 수행 여부가 분쟁률/노쇼율 감소에 미치는 영향

### 113.13 오픈 질문 / 기본 가정
오픈 질문:
- `mutual_ack`를 버튼 클릭 기반으로만 볼지, 특정 write action(도착/확인 응답)으로 암묵 인정할지 확정 필요
- `ready_soft`를 MVP에 포함할지, 초기에는 `needs_confirmation`과 `ready_to_execute` 2단계로 단순화할지 결정 필요
- 운영자 override를 어떤 role/action code로 제한할지 RBAC 문서와 정렬 필요

기본 가정:
- MVP는 `needs_confirmation / ready_to_execute / at_risk / blocked` 4개 상태를 우선 노출하고, `not_started / ready_soft / superseded`는 내부 enum으로 먼저 도입 가능
- no-show claim 허용 여부는 readiness snapshot을 단독 진실원으로 쓰지 않고, claim 입력 시점에 server recomputation으로 재검증한다
- 실행 준비도 snapshot은 홈/내 거래/채팅/운영 큐가 공유하는 canonical read model로 승격하는 것이 바람직하다


## 114. 거래 실행 상호확인(Mutual Execution Ack) / 명시 인정·암묵 인정·최종 출발 신호 canonical 계약
### 114.1 목표
- 예약 성립과 실제 거래 직전의 `서로 같은 약속을 마지막으로 인지했는가`를 분리해 다룬다.
- 사용자가 “예약은 잡았는데 상대가 마지막 변경을 봤는지”, “도착 메시지를 보냈는데 왜 아직 ready가 아닌지”를 예측 가능하게 이해하도록 한다.
- 채팅, 내 거래, 당일 실행 카드, no-show claim, completion gate, 운영 판정이 같은 상호확인 모델을 바라보게 한다.

### 114.2 왜 별도 계약이 필요한가
다음 상황은 reservation/terms/location만으로는 충분히 설명되지 않는다.
- 시간·장소가 확정되어도 상대가 마지막 변경사항을 아직 읽지 못한 경우
- 한쪽은 `도착했어요`를 눌렀지만 상대는 아직 출발 전이라 실제 execution readiness가 비대칭인 경우
- 재일정/장소 변경 이후 상대가 명시 확인 없이 곧바로 `completed` 또는 `no_show`를 주장하는 경우
- 운영이 “양측 모두 같은 약속을 인지했다고 볼 수 있는 최소 근거”를 사건 단위로 재구성해야 하는 경우

따라서 `Execution Readiness Snapshot`과 별도로, 혹은 그 하위 구성요소로서 `Mutual Execution Ack` 계약을 둔다.

### 114.3 핵심 vocabulary
| 필드 | 의미 |
|---|---|
| `mutualExecutionAckState` | 양측의 마지막 실행 조건 인지 상태 요약 |
| `partyAckState` | 각 당사자별 ack 상태 |
| `ackEvidenceType` | ack 판단 근거 타입 |
| `ackTriggerMode` | 명시/암묵/운영대체 중 어떤 방식으로 ack가 성립했는지 |
| `ackScope` | 무엇을 확인한 ack인지 범위(시간/장소/식별/전체) |
| `ackFreshnessState` | 마지막 ack가 여전히 최신 조건에 유효한지 |
| `departureReadinessState` | 단순 인지 수준을 넘어 실제 출발/도착 준비까지 포함한 실행 직전 상태 |
| `finalGoSignalSource` | 실제 실행 가능하다고 볼 최종 신호의 출처 |
| `ackBlockingReasonCode` | 상호확인이 부족해 다음 액션을 막는 사유 |

### 114.4 상태 정의
#### 114.4.1 `partyAckState`
| 값 | 의미 |
|---|---|
| `unknown` | 아직 상대 조건을 읽었는지/인지했는지 근거 없음 |
| `viewed_not_acknowledged` | 최신 카드/변경사항 열람 흔적은 있으나 확인 신호 없음 |
| `implicitly_acknowledged` | 특정 행동으로 최신 조건을 사실상 인정한 상태 |
| `explicitly_acknowledged` | 버튼/응답 등 명시 확인을 남긴 상태 |
| `superseded` | 이전 ack가 있었지만 이후 조건 변경으로 효력 상실 |
| `revoked` | 사용자가 취소 의사/재확인 필요를 밝혀 ack 효력 철회 |

#### 114.4.2 `mutualExecutionAckState`
| 값 | 의미 | 사용자 문구 예시 |
|---|---|---|
| `none` | 양측 모두 유효 ack 없음 | 서로 마지막 확인이 아직 안 됨 |
| `one_sided` | 한쪽만 유효 ack 보유 | 상대 확인 대기 |
| `mutual_soft` | 양측 모두 암묵 ack 수준 | 서로 본 것으로 보이지만 최종 확인 권장 |
| `mutual_confirmed` | 양측 모두 유효 ack 보유, 최소 실행 조건 충족 | 서로 확인 완료 |
| `stale_after_change` | 양측 ack가 있었으나 변경으로 다시 확인 필요 | 변경사항 재확인 필요 |
| `suppressed_by_risk` | ack는 있으나 restriction/risk/overcommit 등으로 실행 신호 차단 | 확인은 됐지만 실행 보류 |

#### 114.4.3 `departureReadinessState`
| 값 | 의미 |
|---|---|
| `not_ready` | 출발/도착 전, 실행 직전 카드 진입 전 |
| `preparing` | 약속은 확인됐으나 출발/접속 준비 중 |
| `en_route` | 이동 중/접속 중을 명시 또는 추정 가능 |
| `arrived_waiting` | 한쪽 또는 양측이 도착 상태 |
| `contacting_on_site` | 현장 또는 인게임 위치에서 상호 찾는 중 |
| `ready_for_handover` | 식별/장소/조건 모두 맞아 실제 인도 직전 |
| `lost_sync` | 도착/지연/장소 변경으로 준비 상태 동기화가 깨짐 |

### 114.5 ack 성립 방식
#### 114.5.1 `ackTriggerMode`
| 값 | 의미 | 예시 |
|---|---|---|
| `explicit_button` | 전용 CTA로 확인 | `약속 확인`, `변경사항 확인`, `도착 확인` |
| `explicit_message_action` | 구조화 메시지 응답으로 확인 | 예약 카드의 `확인했어요` quick action |
| `implicit_followup_action` | 후속 write action 자체가 최신 조건 인지로 해석 | 최신 장소 기준 `도착했어요` 전송 |
| `implicit_counterproposal_acceptance` | 역제안/재일정 수락이 조건 ack를 포함 | 수정된 시간 제안 수락 |
| `system_inferred_read_plus_action` | 읽음 + 관련 실행 행동 조합으로 암묵 인정 | 최신 카드 읽음 후 `출발해요` |
| `operator_override` | 운영 판단으로 대체 인정 | 분쟁 중 명백한 로그 근거 기반 |

#### 114.5.2 `ackEvidenceType`
- `reservation_confirm_click`
- `reschedule_accept_click`
- `location_change_ack_click`
- `terms_change_ack_click`
- `arrival_checkin`
- `delay_notice`
- `on_site_message_after_latest_revision`
- `identity_confirmation_action`
- `operator_case_decision`

원칙:
- 자유 텍스트만으로도 운영 참고 증빙은 될 수 있지만, `canonical ack` 계산의 1순위 근거는 구조화 액션이다.
- 텍스트 메시지는 `ackEvidenceType=free_text_contextual` 같은 보조 근거로만 저장하고, product surface는 구조화 액션 유도를 우선한다.

### 114.6 ack 범위(`ackScope`)와 효력
| 범위 | 의미 | 언제 필요 |
|---|---|---|
| `time_only` | 시간만 확인 | 같은 장소/조건 유지 상태에서 시간만 재확인 |
| `location_only` | 장소만 확인 | 장소 변경 직후 |
| `identity_only` | 캐릭터명/식별자만 확인 | 직전 식별 정보 갱신 시 |
| `terms_only` | 가격/수량/정산만 확인 | deal terms 변경 직후 |
| `full_execution_bundle` | 시간/장소/식별/조건 전체 확인 | 최초 확정, 대폭 변경 후 재확인 |

효력 원칙:
- 더 좁은 scope의 ack는 해당 범위 변경에만 소멸한다.
- `full_execution_bundle` ack 이후 장소만 바뀌면 `ackFreshnessState=stale_for_location`이 되고 전체를 다시 처음부터 깨기보다 `location_only` 재ack로 회복 가능하게 설계하는 것이 바람직하다.

### 114.7 ack freshness / 무효화 규칙
#### 114.7.1 `ackFreshnessState`
| 값 | 의미 |
|---|---|
| `fresh` | 최신 revision에 대해 유효 |
| `stale_for_time` | 시간 변경으로 무효화 |
| `stale_for_location` | 장소 변경으로 무효화 |
| `stale_for_identity` | 식별 정보 변경으로 무효화 |
| `stale_for_terms` | 가격/수량/정산 변경으로 무효화 |
| `expired_by_deadline` | 당일 컷오프 또는 액션 기한 경과 |

무효화 원칙:
- 시간/장소/식별/조건 snapshot revision이 증가하면 관련 scope ack는 자동 `superseded` 처리한다.
- `도착했어요` 같은 실행형 액션은 최신 revision 기준일 때만 암묵 ack로 승격할 수 있다.
- 예전 revision 기준으로 남긴 도착/지연 메시지는 presence evidence로는 유지되지만 최신 ack 효력은 주지 않는다.

### 114.8 화면 / UX 계약
#### 114.8.1 채팅
- 예약/재일정/장소변경/조건변경 카드에는 항상 `상대 확인 상태`를 표시한다.
- 최신 revision에 대한 ack 부족 시 자유 텍스트 입력창 위에 `변경사항 확인 요청` 또는 `도착 전 확인 필요` 배너를 노출한다.
- 사용자가 최신 조건을 본 뒤 `도착했어요`, `출발해요`, `지연돼요`를 누르면 필요 시 해당 액션이 암묵 ack를 함께 생성할 수 있음을 UI copy로 설명한다.

#### 114.8.2 내 거래 / 당일 실행 카드
- 카드 요약 필드 후보:
  - `mutualExecutionAckState`
  - `myAckState`
  - `counterpartyAckState`
  - `lastAckedRevisionLabel`
  - `ackBlockingReasonCode[]`
  - `finalGoSignalSource`
- CTA 우선순위 기본안:
  1. `변경사항 확인`
  2. `상대 확인 요청`
  3. `출발해요`
  4. `도착했어요`
  5. `지연 알리기`
  6. `노쇼 판단` / `완료 진행`
- `mutual_confirmed` 이전에는 `완료`보다 `확인/도착/지연` CTA를 먼저 노출한다.

#### 114.8.3 상태 카피 기본안
| 상태 | 카피 예시 |
|---|---|
| `none` | 아직 마지막 약속 확인이 안 됐어요 |
| `one_sided` | 내가 확인했어요 · 상대 확인 대기 |
| `mutual_soft` | 서로 확인한 것으로 보여요 · 필요하면 최종 확인 남기기 |
| `mutual_confirmed` | 서로 마지막 약속 확인 완료 |
| `stale_after_change` | 변경사항이 생겨 다시 확인이 필요해요 |
| `suppressed_by_risk` | 확인은 됐지만 현재 실행 가드가 걸려 있어요 |

### 114.9 no-show / 완료 / 재일정과의 연동 규칙
- no-show claim 기본안:
  - `mutual_confirmed` 또는 운영상 동등한 근거가 있으면 일반 claim 허용
  - `one_sided` 이하거나 `stale_after_change`면 claim은 허용하되 `claimConfidenceTier`를 낮추고 `claimSuppressionReasonCode` 후보를 함께 기록
  - 최신 변경 후 상대 ack가 없는 상태에서 곧바로 no-show를 주장하면 기본적으로 `suppressed_or_review_required`로 처리하는 것이 바람직하다.
- completion gate 기본안:
  - `ready_for_handover` 또는 equivalent execution evidence가 없으면 `complete` CTA를 막기보다 `confirm conditions first` 경고를 우선 제공
  - 단, 운영 override 또는 handover evidence가 있으면 ack 부족만으로 완료를 절대 차단하지 않는다.
- reschedule linkage:
  - accepted reschedule은 새로운 revision을 생성하고 이전 ack를 `superseded`
  - rejected/expired reschedule은 기존 최신 confirmed bundle을 유지할 수 있으나, 사용자 혼선이 크면 `mutual_soft`로 강등하는 정책을 검토한다.

### 114.10 API / DB 파생 기준
#### 114.10.1 read model 필드 후보
```json
{
  "mutualExecutionAck": {
    "mutualExecutionAckState": "mutual_confirmed",
    "myAckState": "explicitly_acknowledged",
    "counterpartyAckState": "implicitly_acknowledged",
    "ackFreshnessState": "fresh",
    "ackScope": "full_execution_bundle",
    "finalGoSignalSource": "mutual_ack_plus_arrival",
    "ackBlockingReasonCodes": []
  }
}
```

#### 114.10.2 command/API 후보
- `POST /reservations/{reservationId}/ack-latest`
- `POST /reservations/{reservationId}/ack-location-change`
- `POST /reservations/{reservationId}/ack-terms-change`
- `POST /trade-threads/{threadId}/departure-status`
- `POST /trade-threads/{threadId}/request-counterparty-ack`

#### 114.10.3 저장 모델 후보
- `TradeExecutionAckSnapshot`
- `TradeExecutionAckPartyState`
- `TradeExecutionAckEvent`

필드 후보:
- `ackSnapshotId`
- `tradeThreadId`
- `latestExecutionRevision`
- `partyRole`
- `partyAckState`
- `ackTriggerMode`
- `ackEvidenceType`
- `ackScope`
- `ackAt`
- `supersededAt`
- `supersededReasonCode`

### 114.11 운영 / 감사 / analytics 기준
운영자가 봐야 할 최소 정보:
- 최신 execution revision과 각 변경 유형(time/location/identity/terms)
- 양측 ack 타임라인
- 어떤 ack가 explicit이고 어떤 ack가 implicit인지
- no-show / completion / dispute 발생 시점의 effective ack state

이벤트 후보:
- `execution_ack_requested`
- `execution_ack_recorded`
- `execution_ack_superseded`
- `execution_ack_missing_warning_shown`
- `departure_status_set`
- `final_go_signal_reached`
- `final_go_signal_lost`

핵심 분석 포인트:
- explicit ack 비율 vs implicit ack 비율
- `stale_after_change` 상태에서의 no-show/dispute 발생률
- `도착` 액션이 ack 없이 먼저 발생하는 비정상 순서 비율
- ack request 후 상대 응답 시간 분포

### 114.12 오픈 질문 / 기본 가정
오픈 질문:
- `출발해요`를 항상 암묵 ack로 승격할지, 최신 장소/시간이 fresh할 때만 승격할지 최종 확정 필요
- `mutual_soft`를 사용자에게 그대로 보여줄지, MVP에서는 `상호 확인됨`과 `재확인 필요` 2단계로 단순화할지 결정 필요
- `complete` CTA를 ack 부족만으로 얼마나 강하게 막을지 UX/운영 trade-off 확정 필요

기본 가정:
- MVP는 explicit CTA를 우선하되, `도착했어요`/`지연돼요`/`재일정 수락` 같은 구조화 행동은 최신 revision 기준에서 암묵 ack로 인정한다.
- 자유 텍스트는 운영 참고 증빙일 뿐 canonical ack 계산의 주 근거로는 사용하지 않는다.
- `Mutual Execution Ack`는 `Execution Readiness Snapshot`의 하위 구성요소로 쓰되, 화면/API에서는 독립 read model 블록으로 노출하는 것이 구현·QA·운영 측면에서 유리하다.

## 64. Starter DDL / 관계형 테이블 책임 / FK·유니크·삭제 정책(초안)
### 64.1 목적
- 현재 PRD의 상태머신/권한/운영 정책을 실제 초기 RDB 스키마로 옮길 때, 어떤 테이블이 어떤 불변식을 소유하는지 명시한다.
- `문서상 엔티티 후보`와 `실제 1차 migration 대상 테이블` 사이의 간극을 줄인다.
- soft delete, FK 삭제 규칙, 유니크 제약, read model 분리 원칙을 정해 서버 구현과 DB 설계가 다른 해석을 하지 않게 한다.

### 64.2 설계 원칙
1. **write model 우선, read model 분리**: 거래 정합성이 필요한 원천 테이블과 목록/홈/내거래 projection 테이블을 분리한다.
2. **hard delete 최소화**: 거래/신고/감사 관련 원천 테이블은 soft delete 또는 비식별화를 기본으로 한다.
3. **상태머신의 단일 소유권**: 하나의 상태는 하나의 원천 테이블이 소유한다. 예를 들어 Listing 공개 상태는 `listings`, 예약 상태는 `reservations`, 완료 내부 단계는 `trade_completions`가 소유한다.
4. **스냅샷과 이력 분리**: 현재값 컬럼과 별도 상태 이력/event 테이블을 함께 둔다.
5. **운영 복구 가능성 우선**: 운영 숨김/잠금/복구는 update overwrite가 아니라 이력형 액션을 남긴다.

### 64.3 1차 write table 후보
| 테이블 | 역할 | 소유 aggregate | 비고 |
|---|---|---|---|
| `users` | 계정 원천 | User | 인증 공급자, 계정 상태 |
| `user_profiles` | 공개/앱 표시 프로필 | User | 닉네임, 소개, 활동/신뢰 캐시 |
| `user_verifications` | 인증/신뢰 부트스트랩 이력 | User | verification level 이력 분리 |
| `user_restrictions` | 계정/기능 제한 원천 | Restriction | 기간 제한, 해제, probation |
| `user_relationships` | 차단/단골/위험 플래그 | RelationshipEdge | block/favorite-counterparty 등 |
| `listings` | 매물 원천 | Listing | 공개 상태, 핵심 거래 조건 |
| `listing_images` | 매물 이미지 메타 | Listing | 대표 이미지 순서 포함 |
| `listing_attribute_values` | 카테고리별 구조화 속성 | Listing | attribute template 기반 |
| `listing_price_history` | 가격 변경 이력 | Listing | 검색/알림/운영 추적 |
| `listing_status_history` | 공개 상태 변경 이력 | Listing | actor/reason/memo 포함 |
| `favorites` | 찜 관계 | Favorite | 사용자-매물 N:N |
| `chat_rooms` | 매물 기준 1:1 협의 원천 | TradeThread | listingId+pair unique |
| `chat_participants` | 읽음/뮤트/알림 상태 | TradeThread | user별 cursor/notification state |
| `chat_messages` | 메시지 원문 | TradeThread | 일반/system/attachment/event |
| `chat_message_attachments` | 메시지 첨부 메타 | TradeThread | 스토리지 키/권한 등 |
| `reservations` | 예약 원천 | Reservation | 제안/확정/취소/만료 |
| `reservation_meeting_snapshots` | 장소/시간/식별 정보 스냅샷 | Reservation | 변경 diff와 ack 근거 |
| `reservation_status_history` | 예약 상태 이력 | Reservation | no-show/expire 포함 |
| `trade_completions` | 완료 내부 단계 원천 | Completion | requested/confirmed/disputed |
| `exchange_confirmations` | 인도/수령/불일치 기록 | Completion | handover/receipt evidence |
| `final_trade_outcomes` | 최종 종결 결과 | TradeCase | completed/cancelled/no_show/disputed |
| `reviews` | 후기 원천 | Review | 공개 상태/수정/숨김 |
| `reports` | 신고 접수 원천 | Report | 대상/사유/우선순위 |
| `report_attachments` | 신고 증빙 메타 | Report | evidence access policy 연동 |
| `support_cases` | 문의/이의제기/지원 케이스 | SupportCase | 신고/제재/분쟁과 링크 |
| `moderation_cases` | 콘텐츠/행동 정책 검토 케이스 | ModerationCase | auto/manual decision |
| `moderation_actions` | 실제 조치 이력 | ModerationAction | hide/restrict/restore |
| `notifications` | 인앱 알림 원천 | Notification | read/action state |
| `notification_deliveries` | 채널별 발송 시도 | Notification | push suppression/retry |
| `audit_logs` | 운영/시스템 감사 로그 | Audit | break-glass 포함 |
| `domain_events` | outbox/event 원천 | Event | projection/analytics fanout |

### 64.4 projection/read model 테이블 후보
초기 migration에서 projection을 전부 만들 필요는 없지만, 아래는 별도 read model 영역으로 분리하는 것을 권장한다.

| 테이블 | 소비 화면/용도 | 원천 |
|---|---|---|
| `listing_search_documents` | 거래소 목록/검색 | listings, listing attributes, trust snapshot |
| `trade_thread_summaries` | 채팅 목록, 내 거래 | chat_rooms, reservations, completions |
| `home_feed_modules` | 홈 진입 모듈 | listings, notifications, saved search |
| `notification_inbox_projection` | 알림함 | notifications, deliveries |
| `moderation_queue_projection` | 운영 큐 | reports, moderation cases/actions |
| `user_trust_snapshots` | 프로필/상세 배지 | reviews, outcomes, restrictions |

원칙:
- projection은 rebuild 가능해야 하며, 원천 무결성 판단에 사용하지 않는다.
- projection PK는 원천 aggregate PK를 그대로 재사용하거나 별도 natural key를 두되, source version 추적이 가능해야 한다.

### 64.5 핵심 FK / on delete 기본안
| 자식 테이블 | 부모 테이블 | on delete 기본안 | 이유 |
|---|---|---|---|
| `user_profiles.user_id` | `users.id` | `RESTRICT` 또는 soft delete only | 탈퇴해도 거래 참조 유지 필요 |
| `listings.author_user_id` | `users.id` | `RESTRICT` | 매물 이력 보존 |
| `listing_images.listing_id` | `listings.id` | `RESTRICT` | soft delete 매물 이미지 추적 |
| `chat_rooms.listing_id` | `listings.id` | `RESTRICT` | 종결 후에도 대화/분쟁 추적 |
| `chat_participants.chat_room_id` | `chat_rooms.id` | `CASCADE` 허용 가능 | 채팅방 hard delete를 안 하면 영향 적음 |
| `chat_messages.chat_room_id` | `chat_rooms.id` | `RESTRICT` | 대화 원문 보존 우선 |
| `reservations.chat_room_id` | `chat_rooms.id` | `RESTRICT` | 예약 맥락 유지 |
| `reservations.listing_id` | `listings.id` | `RESTRICT` | 매물 삭제와 무관하게 보존 |
| `reservation_meeting_snapshots.reservation_id` | `reservations.id` | `CASCADE` 가능 | reservation 내부 부속 snapshot |
| `trade_completions.listing_id` | `listings.id` | `RESTRICT` | 완료 기록 보존 |
| `trade_completions.chat_room_id` | `chat_rooms.id` | `RESTRICT` | 실제 완료 채널 고정 |
| `reviews.completion_id` | `trade_completions.id` | `RESTRICT` | 후기 근거 유지 |
| `reports.reporter_user_id` | `users.id` | `RESTRICT` | 신고 이력 유지 |
| `moderation_actions.report_id` | `reports.id` | `SET NULL` 허용 | 단독 운영조치도 존재 가능 |
| `notifications.user_id` | `users.id` | `RESTRICT` | 사용자 탈퇴 후도 알림 감사 가능 |

원칙:
- `CASCADE`는 truly-owned child에만 제한적으로 사용한다.
- 거래/신고/감사 관련 FK는 기본 `RESTRICT` 또는 `SET NULL` 우선이다.

### 64.6 soft delete / visibility / anonymization 분리 원칙
같은 `보이지 않음`이라도 아래는 다른 상태로 취급해야 한다.

| 개념 | 예시 컬럼 | 의미 |
|---|---|---|
| soft delete | `deleted_at` | 사용자가 UI에서 삭제/철회한 상태 |
| policy hidden | `hidden_at`, `visibility_reason_code` | 운영/정책상 숨김 |
| anonymized | `anonymized_at` | 개인정보 표시값 치환 |
| archived | `archived_at` | 활성 워크스페이스에서 제외된 기록 |
| purged | `purged_at` | 보존 만료 후 원문 파기 완료 |

원칙:
- 하나의 테이블에 `status` 하나로 모든 비노출 사유를 뭉개지 않는다.
- 검색/목록 노출은 `status + visibility + deleted_at + hidden_at` 조합으로 계산한다.

### 64.7 유니크 제약 기본안
| 테이블 | 유니크 키 | 목적 |
|---|---|---|
| `users` | `(login_provider, login_provider_user_key)` | 외부 계정 중복 방지 |
| `user_profiles` | `(nickname)` partial unique on active/non-withdrawn | 닉네임 충돌 방지 |
| `favorites` | `(user_id, listing_id)` | 중복 찜 방지 |
| `chat_rooms` | `(listing_id, seller_user_id, buyer_user_id)` | 동일 쌍 중복 스레드 방지 |
| `chat_participants` | `(chat_room_id, user_id)` | 참여자 중복 방지 |
| `reviews` | `(completion_id, reviewer_user_id)` | 동일 거래 중복 후기 방지 |
| `user_relationships` | `(subject_user_id, related_user_id, relationship_type)` | block/favorite-counterparty 중복 방지 |
| `notification_deliveries` | `(notification_id, channel, target_device_key)` | 채널별 중복 fanout 제어 |

부분 유니크(partial unique) 후보:
- `reservations(listing_id)` where `reservation_status in ('proposed','confirmed','reschedule_pending')` and `is_active_commitment=true`
- `trade_completions(listing_id)` where `completion_stage in ('requested','confirmed_by_counterparty','auto_confirmed','disputed','resolved_completed')`
- `listings(slug)` where `deleted_at is null`

### 64.8 체크 제약 / enum freeze 포인트
초기 migration에는 아래 성격의 check constraint를 적극 검토한다.
- `price_amount >= 0`
- `quantity > 0`
- `listing_type in ('sell','buy')`
- `price_type='offer'` 인 경우 `price_amount is null or optional policy`
- `meeting_type='in_game'` 인 경우 `server_id is not null`
- `completed_at >= created_at`, `confirmed_at >= proposed_at` 같은 시계열 역전 방지
- `seller_user_id <> buyer_user_id`

원칙:
- 비즈니스 규칙 전부를 DB constraint로 넣지는 않되, **명백히 깨지면 복구 비용이 큰 불변식**은 DB에서도 한 번 더 막는다.

### 64.9 상태 이력 / event / 감사 로그 테이블 구분
| 계층 | 테이블 예시 | 목적 |
|---|---|---|
| 상태 이력 | `listing_status_history`, `reservation_status_history` | 사용자가 이해할 수 있는 상태 전이 추적 |
| 도메인 이벤트 | `domain_events` | outbox, projection fanout, analytics 원천 |
| 감사 로그 | `audit_logs` | 운영/보안/민감열람 추적 |

구분 원칙:
- 상태 이력은 “무슨 상태가 바뀌었나” 중심.
- 도메인 이벤트는 “다른 시스템이 반응해야 하는 사실” 중심.
- 감사 로그는 “누가 어떤 권한으로 무엇을 봤고 바꿨나” 중심.
- 세 가지를 하나의 범용 로그 테이블로 합치지 않는다.

### 64.10 파티셔닝/보관 우선순위
초기 MVP에서 즉시 파티셔닝이 필수는 아니지만, 아래는 장기적으로 분리 가능하도록 설계한다.
- `chat_messages`: 월 단위 range partition 후보
- `domain_events`: 생성일 기준 partition 후보
- `audit_logs`: 생성일 기준 partition + cold storage export 후보
- `notification_deliveries`: TTL/보관기간 짧게 가져가는 별도 전략 후보

원칙:
- `listings`, `chat_rooms`, `reservations`, `trade_completions`, `reports` 같은 핵심 집계 루트는 지나친 조기 파티셔닝보다 단순성과 FK 무결성을 우선한다.

### 64.11 초기 migration 묶음 제안
**Migration 001 — identity / user / restriction**
- `users`
- `user_profiles`
- `user_verifications`
- `user_restrictions`
- `user_relationships`

**Migration 002 — listing core**
- `listings`
- `listing_images`
- `listing_attribute_values`
- `listing_price_history`
- `listing_status_history`
- `favorites`

**Migration 003 — chat / reservation / trade execution**
- `chat_rooms`
- `chat_participants`
- `chat_messages`
- `chat_message_attachments`
- `reservations`
- `reservation_meeting_snapshots`
- `reservation_status_history`
- `trade_completions`
- `exchange_confirmations`
- `final_trade_outcomes`

**Migration 004 — review / report / moderation / support**
- `reviews`
- `reports`
- `report_attachments`
- `support_cases`
- `moderation_cases`
- `moderation_actions`

**Migration 005 — notification / audit / outbox**
- `notifications`
- `notification_deliveries`
- `audit_logs`
- `domain_events`

이 순서는 초기 구현이 `인증 → 매물 → 채팅/예약 → 신고/운영 → 알림/이벤트`로 자연스럽게 내려가도록 하기 위한 권장안이다.

### 64.12 오픈 질문
- `reservation_meeting_snapshots`를 별도 테이블로 분리할지, `reservations` JSON snapshot으로 둘지?
- `final_trade_outcomes`를 `trade_completions`의 closure 필드로 흡수할지, 별도 종결 aggregate로 둘지?
- `moderation_cases`와 `reports`를 분리 유지할지, MVP에서는 reports 하나로 단순화할지?
- search/home/trade summary projection을 DB materialized table로 둘지, search index/Redis 혼합으로 둘지?
- partial unique를 DB에서 직접 강제할지, 일부는 application lock으로 보완할지?



### 64.13 재일정 승인 상태 / 당일 실행 재확인 / 노쇼·분쟁 action bridge 계약
이 섹션의 목적은 `예약 변경 제안`이 단순 채팅 이벤트로 흩어지지 않고, 당일 실행 카드·`executionReadinessSnapshot`·`mutualExecutionAckState`·`noShowCase`·`Dispute`가 같은 판정 언어를 사용하도록 고정하는 것이다. 특히 “재일정 제안을 보냈는데 상대가 못 봤다”, “변경안에 암묵 동의한 것으로 볼 수 있는가”, “이 상태에서 노쇼 claim을 바로 열 수 있는가”를 implementation-ready 수준으로 정리한다.

#### 64.13.1 canonical vocabulary
| 필드/코드 | 의미 | 비고 |
|---|---|---|
| `rescheduleDecisionState` | 최신 재일정 제안의 승인 판정 상태 | `none` / `pending_counterparty` / `accepted_explicit` / `accepted_implicit` / `rejected` / `expired` / `superseded` |
| `executionDecisionState` | 당일 실행 가능성에 대한 현재 결론 | `go_original_plan` / `go_rescheduled_plan` / `hold_waiting_counterparty` / `hold_waiting_ack` / `cannot_execute_today` / `under_review` |
| `executionActionCode` | UI/API/운영에서 공통으로 쓰는 실행 액션 코드 | 아래 64.13.4 표 참조 |
| `claimSuppressionReasonCode` | no-show claim 생성 억제 사유 | `active_reschedule_pending`, `location_change_unread`, `ack_required_not_met`, `grace_period_not_elapsed`, `counterparty_arrival_unverified`, `operator_hold` |
| `rescheduleAckRequirement` | 재일정안에 대해 명시 재확인이 필요한 정도 | `none`, `light`, `strict` |
| `executionPlanSource` | 현재 실행 기준이 된 계획의 출처 | `original_reservation`, `latest_accepted_reschedule`, `operator_override` |
| `executionDecisionChangedAt` | 현재 실행 결론이 마지막으로 바뀐 시각 | 당일 카드/알림 우선순위 기준 |

원칙:
- `rescheduleState`는 변경 제안 객체 자체의 수명주기이고, `rescheduleDecisionState`는 **현재 거래 실행 관점에서 그 변경안이 어떤 판정을 받았는지**를 뜻한다.
- `executionDecisionState`는 UI에서 사용자가 체감하는 “오늘 이 거래를 그대로 진행해도 되는가”를 표현하는 상위 판정값이다.
- no-show claim 가능 여부는 예약 원문이 아니라 **현재 유효한 execution plan + suppression code** 조합으로 판단한다.

#### 64.13.2 재일정안 승인 판정 규칙
| 상황 | `rescheduleDecisionState` | `executionDecisionState` | 기본 해석 |
|---|---|---|---|
| 활성 재일정안 없음 | `none` | `go_original_plan` 또는 기존 readiness 기준 | 원래 예약 기준 유지 |
| 재일정안 발송 후 상대 미열람 | `pending_counterparty` | `hold_waiting_counterparty` | 원래 계획으로 강행하지 않음이 기본 |
| 상대가 명시 수락 | `accepted_explicit` | `go_rescheduled_plan` | 새 시간/장소가 canonical |
| 상대가 변경 diff를 열람했고 `확인했어요/좋아요` 등 인정 | `accepted_explicit` | `go_rescheduled_plan` | 명시 인정으로 간주 |
| 상대가 변경안 이후 도착/지연/길찾기 등 새 계획 기준 행동 수행 | `accepted_implicit` | `go_rescheduled_plan` | 단, strict ack 필요 상태에서는 암묵 인정 불충분 |
| 상대가 거절 | `rejected` | `go_original_plan` 또는 `cannot_execute_today` | 원래 계획 복귀 여부는 시점/조건 따라 다름 |
| 응답 없이 변경안 만료 | `expired` | `go_original_plan` 또는 `cannot_execute_today` | 예약 시각이 너무 임박하면 `cannot_execute_today` |
| 더 최신 재일정안이 생성됨 | `superseded` | 최신안 기준으로 재평가 | 이전안은 판정 소스 아님 |

추가 규칙:
- 재일정안이 시간/장소를 모두 크게 바꾸는 경우 `rescheduleAckRequirement=strict`를 기본으로 두고, `accepted_implicit`만으로는 최종 출발 신호를 만들지 않는다.
- 단순 5~10분 지연 또는 coarse 동일 장소 내 미세 변경은 `light`로 보고 암묵 인정 가능성을 열어둔다.
- `pending_counterparty` 동안은 한쪽이 일방적으로 “상대가 안 왔다”고 주장해도 기본적으로 `claimSuppressionReasonCode=active_reschedule_pending`을 우선 부여한다.

#### 64.13.3 재일정 변경 강도와 ack requirement
| 변경 유형 | 예시 | `rescheduleAckRequirement` | no-show 영향 |
|---|---|---|---|
| `minor_time_shift` | 21:00 → 21:10 | `light` | grace period 재기산 가능 |
| `major_time_shift` | 21:00 → 22:30 / 내일로 변경 | `strict` | 명시 수락 전 no-show 억제 |
| `minor_location_shift` | 같은 마을 내 창고 앞 → 텔포사 앞 | `light` | 암묵 인정 가능 |
| `major_location_shift` | 기란 → 은기사 / 강남역 PC방 → 선릉역 PC방 | `strict` | 명시 수락 전 기존 계획도 hold 후보 |
| `identity_or_method_shift` | 인게임 → 오프라인, 캐릭터명 변경 | `strict` | 식별 혼선으로 claim 억제 우선 |
| `condition_shift` | 수량/가격/지급방식 동반 변경 | `strict` | 사실상 새 deal terms 재확인 필요 |

원칙:
- `strict`가 붙는 순간 `mutualExecutionAckState=ready`로 승격되려면 상대의 명시 수락 또는 운영 override가 필요하다.
- `light` 변경은 상대가 새 카드 열람 + 관련 행동(`on_my_way`, `arrived`, `delay_ack`)을 수행하면 `accepted_implicit`로 승격 가능하다.

#### 64.13.4 당일 실행 카드 CTA / command surface canonical map
| `executionActionCode` | 사용자 의미 | 대표 노출 조건 | API/command 후보 |
|---|---|---|---|
| `confirm_original_plan` | 원래 약속대로 진행 확인 | 재일정안이 reject/expire됨, 원래 계획 유지 가능 | `POST /reservations/{reservationId}/confirm-original-plan` |
| `accept_reschedule` | 변경안 수락 | `pending_counterparty`이며 수신자 본인 | `POST /reschedules/{rescheduleId}/accept` |
| `reject_reschedule` | 변경안 거절 | `pending_counterparty`이며 수신자 본인 | `POST /reschedules/{rescheduleId}/reject` |
| `ack_reschedule_seen` | 변경안 읽고 인지 표시 | `light` 또는 `strict` 변경안 수신 후 | `POST /reschedules/{rescheduleId}/ack` |
| `share_delay` | 몇 분 늦음 공유 | 당일 실행 중 지연 발생 | `POST /trade-execution/{tradeCaseId}/delay` |
| `mark_on_the_way` | 출발/이동 시작 | readiness 충족 후 | `POST /trade-execution/{tradeCaseId}/depart` |
| `mark_arrived` | 도착 알림 | meetup 직전/직후 | `POST /trade-execution/{tradeCaseId}/arrive` |
| `request_reconfirm` | 상대 재확인 요청 | `hold_waiting_ack` 상태 | `POST /trade-execution/{tradeCaseId}/request-reconfirm` |
| `open_no_show_claim` | 노쇼 claim 시작 | suppression 없음 + grace 경과 | `POST /trade-execution/{tradeCaseId}/no-show-claims` |
| `open_execution_dispute` | 실행/조건 불일치 분쟁 시작 | 장소 혼선/조건 충돌/상호 주장 상반 | `POST /trade-execution/{tradeCaseId}/disputes` |

설계 원칙:
- 화면 버튼 텍스트는 자유롭게 현지화하되, analytics/audit/API는 위 `executionActionCode`를 canonical key로 사용한다.
- `open_no_show_claim`과 `open_execution_dispute`는 동시에 열려선 안 되며, 먼저 열린 사건이 `tradeCaseSummary.activeCaseType`을 점유한다.

#### 64.13.5 no-show claim 생성 게이트
`open_no_show_claim`은 아래 조건을 모두 만족해야 한다.
1. 현재 `executionDecisionState`가 `go_original_plan` 또는 `go_rescheduled_plan`
2. 유효 execution plan의 scheduled time 기준 grace period 경과
3. 내 쪽 `arrivalState` 또는 이에 준하는 도착/근접 증빙이 최소 1개 존재
4. `claimSuppressionReasonCode`가 비어 있음
5. 같은 execution plan 기준 열린 unresolved no-show case 없음

자동 억제 규칙:
- 활성 재일정안이 미응답 상태면 `active_reschedule_pending`
- 큰 장소 변경 후 상대 ack 미완료면 `ack_required_not_met`
- 최신 장소 카드가 상대에게 unread면 `location_change_unread`
- 내 도착 증빙이 없고 단순 채팅 주장만 있으면 `counterparty_arrival_unverified`
- 운영자가 이미 execution dispute 검토를 열었으면 `operator_hold`

#### 64.13.6 no-show vs execution dispute 분기 규칙
| 상황 | 기본 사건 유형 | 비고 |
|---|---|---|
| 한쪽은 도착, 상대는 무응답/미도착 주장 | `no_show_case` | 전형적 no-show |
| 시간/장소/캐릭터 변경 인지 여부가 쟁점 | `execution_dispute` | no-show보다 우선 |
| 둘 다 도착 주장하나 서로 못 찾음 | `execution_dispute` | identity/location mismatch |
| 가격/수량/지급방식 변경으로 현장 결렬 | `execution_dispute` | no-show 아님 |
| 재일정안이 만료된 상태에서 누구 계획이 유효한지 불명확 | `execution_dispute` | 운영 판정 필요 |

원칙:
- 실행 계획의 유효성 자체가 쟁점이면 no-show보다 dispute를 우선한다.
- 운영자는 dispute를 판정하며 필요 시 결과를 `confirmed_no_show`, `mutual_miscoordination`, `terms_mismatch_cancelled` 등 final outcome으로 연결한다.

#### 64.13.7 DB / read model 파생 기준
최소 저장 후보:
- `reschedules.reschedule_decision_state`
- `reschedules.ack_requirement`
- `trade_execution_snapshots.execution_decision_state`
- `trade_execution_snapshots.execution_plan_source`
- `trade_execution_snapshots.claim_suppression_reason_code`
- `trade_execution_action_history.execution_action_code`
- `trade_execution_action_history.action_result_state`
- `trade_execution_action_history.plan_version`

read model 필수 필드:
- `currentPlanSummaryLabel`
- `requiresExplicitAck`
- `canOpenNoShowClaim`
- `disabledExecutionActionCodes[]`
- `disabledReasonByActionCode{}`
- `latestRescheduleDecisionState`
- `activeCaseType`

#### 64.13.8 화면 요구사항 파생 기준
- **채팅/당일 실행 카드**는 “현재 유효한 계획” 1개만 primary로 보여주고, superseded plan은 접힌 이력으로만 노출한다.
- **내 거래 리스트**는 `hold_waiting_counterparty`, `hold_waiting_ack`를 단순 미읽음보다 높은 urgency로 정렬한다.
- **노쇼 claim 진입 화면**은 시작 전에 suppression reason이 있으면 claim form 대신 “왜 지금 바로 노쇼 접수가 안 되는지”를 설명해야 한다.
- **분쟁 진입 화면**은 `시간/장소/식별/조건 중 무엇이 엇갈렸는지`를 구조화 선택하게 해야 한다.

#### 64.13.9 운영 adjudication 기준
운영자는 아래 순서로 판단한다.
1. 마지막 유효 execution plan이 original인지 rescheduled인지
2. 그 plan에 strict ack가 필요했는지
3. 상대가 그 plan을 명시/암묵적으로 수락했는지
4. 각 당사자의 출발/도착/지연 증빙 존재 여부
5. no-show인지, 단순 miscoordination인지, terms mismatch인지

기본 판정 원칙:
- strict ack가 필요한 큰 변경안은 명시 수락 없이는 상대 귀책 no-show로 쉽게 보지 않는다.
- 양측 모두 불명확하거나 마지막 유효 plan 자체가 모호하면 `mutual_miscoordination` 또는 `insufficient_evidence` 우선
- 반복적으로 strict 변경을 직전에 던지는 사용자는 no-show와 별도로 trust/moderation risk signal로 반영 가능

#### 64.13.10 analytics / KPI 파생 기준
추가 이벤트 후보:
- `reschedule_ack_requested`
- `reschedule_ack_completed`
- `execution_plan_switched_to_rescheduled`
- `execution_action_blocked`
- `no_show_claim_suppressed`
- `execution_dispute_opened`

추가 KPI 후보:
- reschedule 제안 대비 explicit acceptance 비율
- strict reschedule 이후 거래 완료율
- claim suppression 발생률 및 주요 suppression reason 분포
- no-show로 열렸다가 execution dispute로 재분류되는 비율
- miscoordination 판정 비율(시간/장소/식별 문제 분해 포함)

#### 64.13.11 거래 실행 projection 응답 계약 / 내 거래·거래상세·당일 실행 surface canonical payload
이 하위 섹션의 목적은 앞서 정의한 `executionReadinessSnapshot`, `mutualExecutionAck`, `rescheduleDecisionState`, `executionDecisionState`, `claimSuppressionReasonCode`가 화면별로 제각각 재조합되지 않도록, `GET /me/trades`, `GET /me/trades/{tradeThreadId}`, 당일 실행 카드가 공유해야 할 canonical response block을 고정하는 것이다. 구현 관점에서는 projection table/schema, OpenAPI response, mobile view model, QA fixture의 공통 shape로 사용한다.

##### 64.13.11.1 projection 블록 분리 원칙
- `tradeThreadSummary`는 리스트 정렬/배지/CTA 노출에 필요한 최소 집약 블록이다.
- `executionControl`은 당일 실행 의사결정에 필요한 readiness/ack/reschedule/no-show gating 블록이다.
- `activeCaseSummary`는 no-show/dispute/support case가 열려 있을 때 실행 액션을 억제하거나 대체 CTA로 치환하는 블록이다.
- 화면은 원칙적으로 원문 aggregate(`reservation`, `reschedule`, `completion`, `report`)를 직접 해석하지 않고 위 projection block을 우선 사용한다.

##### 64.13.11.2 `GET /me/trades` 리스트 item 필수 필드
| 필드 | 타입/예시 | 의미 |
|---|---|---|
| `tradeThreadId` | string | 내 거래 스레드 canonical ID |
| `tradeCaseId` | string | 실행/분쟁/action history를 묶는 case ID |
| `listingId` | string | 원 매물 ID |
| `listingTitle` | string | 리스트 카드 제목 |
| `listingType` | `sell` / `buy` | 역할 해석 기준 |
| `counterpartySummary` | object | 상대 닉네임/신뢰 배지/차단 여부 요약 |
| `threadState` | enum | 기존 `tradeThreadState` 계약 재사용 |
| `threadUrgencyTier` | enum | 정렬 우선순위 |
| `nextBestAction` | enum | 카드 primary CTA 후보 |
| `scheduledExecutionAt` | datetime/null | 현재 유효 plan 기준 실행 시각 |
| `executionDecisionState` | enum | `go_original_plan` 등 |
| `rescheduleDecisionState` | enum | 최신 재일정 승인 판정 |
| `executionReadinessState` | enum | 최신 readiness snapshot 요약 |
| `mutualExecutionAckState` | enum | 상호 실행 인지 수준 |
| `claimSuppressionReasonCode` | enum/null | no-show claim 억제 사유 |
| `dayOfTradeCardMode` | enum/null | 당일 카드 모드 |
| `availableExecutionActionCodes` | string[] | 카드에서 즉시 노출 가능한 실행 action |
| `disabledReasonByActionCode` | object | action별 disabled 사유 |
| `activeCaseSummary` | object/null | no-show/dispute/support active case |
| `updatedAt` | datetime | projection freshness |

리스트 규칙:
- `GET /me/trades`의 기본 정렬은 `threadUrgencyTier desc`, 그 다음 `scheduledExecutionAt asc nulls last`, 그 다음 `updatedAt desc`를 권장한다.
- `claimSuppressionReasonCode`가 존재하면 리스트 카드의 `노쇼 신고` CTA는 숨기지 말고 disabled 또는 대체 설명형 CTA로 노출한다.
- `rescheduleDecisionState=pending_counterparty` 또는 `executionDecisionState=hold_waiting_ack`는 단순 미읽음보다 높은 urgency를 가져야 한다.

##### 64.13.11.3 `executionControl` 공통 블록
`GET /me/trades` item과 `GET /me/trades/{tradeThreadId}` 상세는 아래 블록을 같은 이름으로 재사용하는 것을 권장한다.

```json
{
  "executionControl": {
    "planVersion": 7,
    "executionPlanSource": "latest_accepted_reschedule",
    "scheduledExecutionAt": "2026-03-14T21:10:00+09:00",
    "meetingPointLabel": "기란 마을 창고 앞",
    "meetingType": "in_game",
    "executionDecisionState": "go_rescheduled_plan",
    "executionReadinessState": "ready",
    "mutualExecutionAckState": "ready",
    "rescheduleDecisionState": "accepted_explicit",
    "requiresExplicitAck": false,
    "claimSuppressionReasonCode": null,
    "canOpenNoShowClaim": true,
    "dayOfTradeCardMode": "result_needed",
    "availableExecutionActionCodes": ["mark_arrived", "open_no_show_claim"],
    "disabledReasonByActionCode": {
      "request_reconfirm": "already_ready"
    }
  }
}
```

필드 원칙:
- `planVersion`은 action history, SSE reconciliation, idempotent command guard에 공통 사용한다.
- `requiresExplicitAck`는 UI 편의 필드지만 서버 recomputation 결과와 동일 의미를 가져야 한다.
- `canOpenNoShowClaim`은 단순 boolean convenience field이고, command 실행 시 서버는 `claimSuppressionReasonCode`와 grace 조건을 다시 검증해야 한다.
- `meetingPointLabel`은 공개/마스킹 정책이 반영된 표시 문자열이며 원문 장소 상세를 그대로 의미하지 않는다.

##### 64.13.11.4 `GET /me/trades/{tradeThreadId}` 상세 추가 블록
상세 endpoint는 리스트 item보다 아래를 추가로 포함해야 한다.
- `currentPlanSummary`: 현재 유효 실행 계획의 시간/장소/식별/조건 diff 요약
- `latestRescheduleProposal`: 최신 재일정 카드 원문 요약 및 superseded 여부
- `mutualExecutionAck`: ack actor, ack evidence type, ack updated at, strict/light 여부
- `readinessBlockers[]`: `blockingPrerequisiteType`, `blockingReasonCode`, `clearActionCode`
- `executionActionHistory[]`: 최근 action 10~20개, `executionActionCode`, actor, result, planVersion
- `claimEligibilitySummary`: `canOpenNoShowClaim`, suppression 사유, 필요 추가 증빙
- `activeCaseSummary`: 사건 유형, 상태, next operator/user action, case deep link

상세 규칙:
- 상세는 superseded reschedule plan을 히스토리 탭에서 볼 수 있게 하되, primary decision 영역에는 최신 유효 plan 1개만 보여준다.
- `executionActionHistory`는 채팅 타임라인을 대체하지 않지만, 당일 실행 판단에 필요한 action만 빠르게 읽을 수 있는 운영/사용자용 보조 타임라인이다.

##### 64.13.11.5 당일 실행 카드 최소 렌더링 계약
당일 실행 카드 surface는 아래만 받아도 1차 렌더링이 가능해야 한다.
- `dayOfTradeCardMode`
- `executionDecisionState`
- `executionReadinessState`
- `mutualExecutionAckState`
- `scheduledExecutionAt`
- `meetingPointLabel`
- `availableExecutionActionCodes[]`
- `disabledReasonByActionCode{}`
- `claimSuppressionReasonCode`
- `activeCaseSummary.caseType`

렌더링 원칙:
- 카드 상단 상태카피는 `dayOfTradeCardMode` 우선, 없으면 `executionDecisionState`, 그마저 없으면 `executionReadinessState` 순으로 fallback 한다.
- CTA 노출은 `availableExecutionActionCodes`를 1차 truth source로 사용하고, 클라이언트가 enum 조합을 재추론하지 않는다.
- `claimSuppressionReasonCode`가 있으면 노쇼 CTA를 숨기지 말고 `왜 지금은 안 되는지`를 설명하는 secondary sheet/bottom sheet를 열 수 있어야 한다.

##### 64.13.11.6 projection 저장/동기화 파생 기준
권장 projection 컬럼/문서 필드:
- `trade_thread_projection.execution_control_json`
- `trade_thread_projection.day_of_trade_card_mode`
- `trade_thread_projection.next_best_action`
- `trade_thread_projection.thread_urgency_tier`
- `trade_thread_projection.active_case_type`
- `trade_thread_projection.active_case_state`
- `trade_thread_projection.claim_suppression_reason_code`
- `trade_thread_projection.plan_version`

동기화 원칙:
- `reservation_confirmed`, `reschedule_accepted`, `reschedule_rejected`, `execution_ack_recorded`, `arrival_recorded`, `no_show_case_opened`, `execution_dispute_opened`, `final_outcome_closed` 이벤트는 모두 projection 재계산 트리거다.
- SSE/polling delta payload는 최소 `tradeThreadId`, `updatedAt`, `planVersion`, `executionDecisionState`, `dayOfTradeCardMode`, `availableExecutionActionCodes`를 포함하는 것을 권장한다.

##### 64.13.11.7 OpenAPI / QA fixture 파생 기준
OpenAPI에서는 아래 schema 후보로 바로 파생 가능해야 한다.
- `TradeThreadListItem`
- `ExecutionControlBlock`
- `TradeThreadDetailExecutionSection`
- `ClaimEligibilitySummary`
- `ExecutionActionHistoryItem`

QA fixture 최소 세트:
1. 원래 계획 유지(`go_original_plan`) 정상 거래
2. light reschedule + implicit accept
3. strict reschedule + explicit accept 필요
4. pending reschedule로 no-show claim suppressed
5. execution dispute로 no-show 대신 분쟁 진입
6. active no-show case로 일반 execution CTA 일부 잠김

검증 포인트:
- 같은 거래가 목록/상세/당일 실행 카드에서 서로 다른 계획 시각·장소를 보여주지 않는가
- disabled CTA 사유가 API 응답과 화면 문구에서 일치하는가
- `planVersion`이 증가한 뒤 stale command가 서버에서 거절되는가


##### 64.13.12 당일 실행 알림 cadence / suppression / escalation canonical 계약

목표:
- 예약 확정 이후 실제 거래가 일어나는 마지막 24시간 동안, 사용자가 필요한 행동만 적시에 받도록 하고 중복/소음/오판 유발 알림을 줄인다.
- `executionReadinessSnapshot`, `mutualExecutionAck`, `dayOfTradeCardMode`, `noShowCase`, `Dispute`가 서로 다른 알림 규칙을 쓰지 않도록 canonical notification vocabulary를 정의한다.
- 푸시/인앱/배지/scheduler/운영 runbook이 같은 reminder 단계와 suppression reason을 공유하도록 한다.

###### 64.13.12.1 핵심 execution notification event family
| event family | 의미 | 대표 트리거 | 기본 수신자 | 기본 채널 |
|---|---|---|---|---|
| `execution_window_opened` | 거래 당일 준비 창 진입 | 예약 시각이 execution reminder window에 진입 | 양측 | 인앱 + 선택적 푸시 |
| `execution_reconfirm_requested` | 계획 재확인 필요 | strict/light ack 필요 상태 진입 | ack 미완료 당사자 | 인앱 + 푸시 |
| `execution_plan_changed` | 시간/장소/식별/조건 변경 | 유효 planVersion 증가 | 양측 | 인앱 + 푸시 |
| `execution_arrival_prompt` | 도착/출발/지연 입력 유도 | day-of-trade card가 arrival 단계로 전환 | 양측 | 인앱 + 푸시 |
| `execution_counterparty_arrived` | 상대 도착 알림 | counterparty arrival recorded | 상대방 | 인앱 + 푸시 |
| `execution_delay_declared` | 상대 지연 통지 | late/ delay action 실행 | 상대방 | 인앱 + 푸시 |
| `execution_no_show_eligibility_opened` | 노쇼 claim 가능 창 진입 | grace 종료 + suppression 해제 | eligible party | 인앱 우선 + 푸시 |
| `execution_case_opened` | no-show/dispute/support case 생성 | case created | 당사자 + 운영 큐 | 인앱 + 중요 푸시 |
| `execution_case_action_required` | 소명/응답/검토 필요 | case state가 user_action_required/operator_action_required | 해당 당사자/담당 운영자 | 인앱 + 푸시 |
| `execution_case_resolved` | 사건 판정 종료 | case outcome finalized | 사건 당사자 | 인앱 + 푸시 |

원칙:
- 같은 도메인 사건이라도 UI 문구는 `event family + actor role + urgency tier` 조합으로 파생하되, underlying event family는 projection/analytics/scheduler에서 동일해야 한다.
- execution 관련 알림은 marketing/re-engagement 알림과 별도 큐/우선순위를 가져야 한다.

###### 64.13.12.2 reminder stage vocabulary
| stage | 기본 시간창(가정) | 목적 | 대표 CTA |
|---|---|---|---|
| `t_minus_24h` | scheduledExecutionAt 24h 전후 | 내일 거래 인지, plan drift 조기 발견 | 일정 보기, 다시 확인 |
| `t_minus_3h` | 3h 전후 | 장소/시간/식별 final check | 확인했어요, 재일정 제안 |
| `t_minus_30m` | 30m 전후 | 출발/준비 유도 | 출발할게요, 조금 늦어요 |
| `t_minus_5m_to_t_plus_15m` | 직전~초기 grace | 도착/현장 상태 입력 유도 | 도착했어요, 못 가요, 지연 알림 |
| `t_plus_grace` | grace 종료 시점 | no-show 여부 판단 가능성 열기 | 노쇼 신고, 분쟁 열기 |
| `post_case_open` | case 생성 후 | 사건 소명/판정 진행 | 소명 제출, 상태 보기 |

가정/정책 원칙:
- 정확한 시간값은 config로 옮길 수 있으나, PRD 기준 vocabulary는 고정한다.
- `t_minus_24h`와 `t_minus_3h`는 예약 시각/카테고리/거래방식에 따라 일부 생략 가능하지만, `t_minus_30m`와 `t_plus_grace`는 execution-critical stage로 본다.

###### 64.13.12.3 notification cadence decision rules
1. `executionReadinessState=ready`이고 `mutualExecutionAckState=aligned`이면 reminder는 최소 cadence만 보낸다.
2. `requiresExplicitAck=true` 또는 `readinessBlockers[]` 존재 시 `execution_reconfirm_requested`를 일반 reminder보다 우선한다.
3. `planVersion`이 증가한 직후 기존 예약 reminder는 무효화하고 최신 plan 기준 reminder schedule을 다시 계산한다.
4. `cancelIntentState`가 활성화되거나 `finalOutcomeType`이 확정되면 future execution reminder는 전부 취소한다.
5. active `noShowCase` 또는 `Dispute`가 열리면 일반 execution reminder 대신 case-specific notification family만 허용한다.

###### 64.13.12.4 suppression reason code family
execution notification은 아래 suppression reason을 공통 사용해야 한다.

| suppression reason | 의미 | 예시 |
|---|---|---|
| `already_aligned_recently` | 최근 ack/확인으로 동일 목적 알림 불필요 | 10분 전 명시 재확인 완료 |
| `plan_superseded` | 더 최신 plan 존재 | 재일정 승인 직후 이전 reminder 취소 |
| `outcome_closed` | 거래가 이미 완료/취소/분쟁 종결 | completed 후 arrival prompt 억제 |
| `case_controls_surface` | 사건 UI가 primary source | no-show case 중 일반 reminder 금지 |
| `user_currently_active_in_surface` | 사용자가 same trade 화면 활성 사용 중 | 채팅/당일 실행 카드 열람 중 푸시 억제 |
| `quiet_hours_non_critical` | 심야 비중요 알림 | 후기성 리마인드/저우선 재확인 |
| `counterparty_action_already_recorded` | 상대 액션으로 목적 달성 | 상대 도착 알림 후 추가 arrival prompt 불필요 |
| `policy_delivery_restricted` | 제재/뮤트/채널 차단 | push opt-out, thread mute |

원칙:
- suppression은 “발송 안 함”만이 아니라 “푸시만 억제하고 인앱 기록은 남김”을 포함한다.
- `suppressionReasonCode`는 NotificationDispatch/analytics/runbook에서 그대로 추적 가능해야 한다.

###### 64.13.12.5 urgency tier / delivery contract
| urgency tier | 설명 | 채널 원칙 | 예시 |
|---|---|---|---|
| `critical_execution` | 지금 놓치면 거래 실패 가능성 높음 | 인앱 + 푸시 기본 | arrival prompt, no-show eligibility |
| `high_execution` | 짧은 시간 내 재확인 필요 | 인앱 + 푸시 권장 | plan changed, reconfirm requested |
| `normal_execution` | 준비/인지 목적 | 인앱 기본, 푸시 선택 | t-24h reminder |
| `case_critical` | 소명/판정/제한에 직접 영향 | 인앱 + 중요 푸시 | case action required |

전달 원칙:
- `critical_execution`은 quiet hours라도 억제하지 않는 기본안을 둔다. 단, exact location 등 민감 세부는 푸시에 넣지 않는다.
- `normal_execution`은 same-thread active usage 시 push를 생략하고 notification inbox만 남겨도 된다.
- `case_critical`은 알림함, trade detail badge, support/case surface badge가 동시에 일관되게 반영되어야 한다.

###### 64.13.12.6 payload contract for execution notifications
모든 execution notification은 최소 아래 payload를 공유하는 것이 권장된다.

```json
{
  "notificationType": "execution",
  "eventFamily": "execution_reconfirm_requested",
  "tradeThreadId": "trade_123",
  "listingId": "listing_123",
  "urgencyTier": "high_execution",
  "planVersion": 7,
  "scheduledExecutionAt": "2026-03-14T21:00:00+09:00",
  "meetingPointLabel": "기란 마을 창고 앞",
  "deepLink": "/me/trades/trade_123?focus=execution",
  "availableExecutionActionCodes": ["ack_plan", "request_reschedule"],
  "suppressionReasonCode": null
}
```

payload 원칙:
- `meetingPointLabel`은 공개/마스킹 정책이 반영된 표출용 문자열이다.
- `availableExecutionActionCodes`는 알림 열기 후 즉시 landing CTA를 구성하는 1차 source로 사용한다.
- exact address, 상대 실명 추정 정보, 내부 판정 메모는 payload에 포함하지 않는다.

###### 64.13.12.7 scheduler / re-schedule / cancellation 연동 규칙
- notification scheduler는 `scheduledExecutionAt`, `graceEndsAt`, `planVersion`, `finalOutcomeType`, `activeCaseType`를 기준으로 due job을 계산해야 한다.
- `reschedule_accepted`가 발생하면 기존 reminder job은 `plan_superseded`로 취소하고 새 시간 기준으로 재생성한다.
- `reschedule_rejected`, `cancel_intent_resolved=cancelled`, `trade_completed_confirmed`, `no_show_case_opened`, `execution_dispute_opened`는 모두 future execution reminder purge trigger다.
- scheduler는 “이미 보낸 알림”보다 “현재 유효한 execution plan”을 truth source로 삼아야 하며, catch-up batch가 있어도 stale reminder를 뒤늦게 보내면 안 된다.

###### 64.13.12.8 operator escalation / queue linkage
execution notification은 운영 큐와 아래처럼 연결되어야 한다.

| 상황 | 운영 큐 적재 조건 | 목적 |
|---|---|---|
| repeated no response to reconfirm | strict ack required 상태에서 deadline 초과 반복 | 거래 실패 사전 탐지, 사용자 가이드 |
| repeated late/no-show signals | 동일 사용자 최근 누적 패턴 | abuse/reliability review |
| case opened by execution failure | no-show/dispute/support case 생성 | operator adjudication |
| scheduler failure on critical reminder | critical_execution due job 실패 | ops/investigation |

원칙:
- 운영 큐 적재는 사용자 알림 자체와 분리되며, “알림을 못 받았기 때문에 자동 제재”가 되지 않도록 한다.
- 다만 `critical_execution` reminder가 여러 번 실패했거나 사용자가 반복적으로 last-minute drift를 유발하면 support/moderation 관점의 review candidate가 될 수 있다.

###### 64.13.12.9 analytics / KPI 파생 기준
추가 추적 이벤트 후보:
- `execution_reminder_scheduled`
- `execution_reminder_sent`
- `execution_reminder_suppressed`
- `execution_reminder_opened`
- `execution_reconfirm_completed_from_notification`
- `execution_arrival_recorded_from_notification`
- `execution_case_opened_after_notification`

핵심 지표:
- reminder stage별 open rate / action completion rate
- `t_minus_30m` reminder 후 arrival ack 전환율
- `execution_reconfirm_requested` 발송 후 strict ack 완료율
- stale/superseded reminder 발송률(목표: 0에 수렴)
- no-show case 중 “사전 reconfirm 실패” 선행 비율

###### 64.13.12.10 화면 / API / DB / runbook 파생 기준
화면 요구사항:
- 알림함에서 execution notification은 일반 메시지 알림과 다른 iconography/priority treatment를 가져야 한다.
- 알림 탭 진입 시 `eventFamily`, `urgencyTier`, `scheduledExecutionAt`, `meetingPointLabel`, `nextBestAction`이 일관되게 보이도록 한다.

API/OpenAPI 파생 기준:
- `NotificationItem` schema는 execution family 전용 payload block을 지원해야 한다.
- `GET /notifications`와 `GET /me/trades`가 같은 `availableExecutionActionCodes` 의미를 공유해야 한다.

DB/worker 파생 기준:
- `notification_dispatch` 또는 동등 테이블에 `event_family`, `urgency_tier`, `plan_version`, `suppression_reason_code`, `case_link_id` 저장을 권장한다.
- critical reminder due job은 idempotent re-run 가능해야 하고, stale planVersion이면 skip 처리해야 한다.

운영 runbook 파생 기준:
- “거래 당일 알림 장애 시 수동 공지/배지 fallback”
- “reconfirm reminder 반복 실패 사용자의 케이스 해석 가이드”
- “plan superseded 이후 stale push 발송 사고 대응”

## 63. 가시성 패널티(Visibility Penalty) / 검색 랭킹 제한 / 조용한 제한(shadow moderation) canonical 계약
### 63.1 목적
- 즉시 차단/정지로 가기 전 단계에서 `노출 제한`, `추천 제외`, `알림 억제`, `검토 강화`를 구조화해 운영 정책과 랭킹 시스템이 같은 언어를 쓰도록 한다.
- 허위매물 의심, 반복 예약 불이행, 외부연락처 유도, 스팸성 등록처럼 사용자 피해 가능성은 있으나 즉시 영구 제재까지는 아닌 상황을 일관되게 다룬다.
- 사용자에게는 과도한 낙인 없이 필요한 수준의 안내만 제공하고, 내부적으로는 검색/추천/운영큐/API/analytics가 같은 패널티 객체를 바라보게 한다.

### 63.2 핵심 vocabulary
| 필드 | 의미 | 후보 값 |
|---|---|---|
| `visibilityPenaltyState` | 현재 가시성 패널티 상태 | `none` / `watch` / `limited_exposure` / `review_only` / `hidden_until_review` |
| `penaltyReasonFamily` | 패널티가 발생한 이유의 상위 분류 | `spam_listing` / `reserved_abuse` / `high_risk_new_account` / `fraud_signal` / `contact_evasion` / `review_abuse` / `quality_low` / `policy_repeat_offense` |
| `exposureSurface` | 패널티가 영향을 주는 노출면 | `search` / `home_recommendation` / `saved_search_notification` / `public_profile` / `review_feed` / `chat_initiation` |
| `userDisclosureLevel` | 사용자에게 어떤 수준으로 안내할지 | `silent` / `soft_hint` / `explicit_notice` |
| `penaltyScopeType` | 어떤 객체 범위에 적용되는지 | `listing` / `user` / `review` / `search_term_cluster` |
| `penaltyTriggerSource` | 패널티를 만든 근거 | `rule` / `risk_model` / `manual_staff` / `appeal_revision` |
| `penaltyExpiryMode` | 어떻게 종료되는지 | `time_based` / `event_based` / `manual_release` / `recompute_daily` |
| `penaltyImpactTier` | 랭킹/노출/행동 제한 강도 | `P1_low` / `P2_moderate` / `P3_high` / `P4_hold` |

### 63.3 설계 원칙
1. **명시 제재와 분리**: visibility penalty는 `restriction`이나 `ban`과 다른 객체/상태로 관리한다.
2. **사용자 보호 우선**: 사기/괴롭힘/개인정보 노출 의심이 강하면 조용한 제한보다 명시 차단/숨김이 우선한다.
3. **설명 가능성 확보**: 내부적으로는 왜 노출이 줄었는지 설명 가능한 reason/evidence를 반드시 남긴다.
4. **무기한 shadow 처리 금지**: 모든 패널티는 만료/재계산/재심 규칙이 필요하다.
5. **노출면별 영향 분리**: 검색 후순위와 신규 채팅 제한, saved-search 알림 억제는 같은 패널티라도 독립적으로 설정 가능해야 한다.

### 63.4 패널티 상태 정의
| 상태 | 의미 | 사용자 영향 | 운영 처리 기본안 |
|---|---|---|---|
| `none` | 패널티 없음 | 정상 노출 | 없음 |
| `watch` | 내부 관찰만 수행 | 직접 체감 없음 | 운영 모니터링/지표 추적 |
| `limited_exposure` | 검색/추천 후순위, 일부 알림 제외 | 노출 감소 가능 | 자동 또는 수동 적용 가능 |
| `review_only` | 공개 노출은 유지하되 추천/랭킹 승격 제외, 신규 고위험 행동은 검토 강화 | 일부 행동에서 지연/검토 안내 가능 | 운영 큐와 연동 |
| `hidden_until_review` | 공개 노출 중단, 직접 링크/작성자/운영만 접근 | 사실상 임시 비공개 | 빠른 운영 검토 필수 |

세부 원칙:
- `hidden_until_review`는 사실상 `temporary hide`에 가까우므로 장시간 유지 시 명시적 운영 상태로 승격하거나 해제해야 한다.
- `watch`와 `limited_exposure`는 사용자 체감이 약할 수 있으므로 오탐 모니터링과 해제 기준이 특히 중요하다.

### 63.5 reason family 기본안
| family | 설명 | 대표 트리거 예시 | 기본 영향 |
|---|---|---|---|
| `spam_listing` | 도배성/유사 매물 반복 등록 | 짧은 시간 유사 제목·가격·아이템 반복 | search/home 후순위 |
| `reserved_abuse` | 허위 희소성, 예약 상태 남용 | `reserved -> available` 반복, 장기 reserved 방치 | search 후순위, reserved 배지 신뢰도 하향 |
| `high_risk_new_account` | 신규 계정 + 고가/고위험 패턴 | 가입 직후 고가 매물 다건, 응답 회피 | recommendation 제외, review_only |
| `fraud_signal` | 사기 의심 정황 누적 | 신고 다발, 증빙 일치, 외부 이동 강요 | hidden_until_review 또는 manual queue |
| `contact_evasion` | 플랫폼 밖 이동 지속 유도 | 연락처 패턴 반복 차단, 외부 메신저 강한 유도 | chat initiation 제한 강화, recommendation 제외 |
| `review_abuse` | 후기 조작/보복 패턴 | 맞비추천 반복, 동일 문구 다발 | review feed 비노출, trust 집계 hold |
| `quality_low` | 품질 저하로 거래 가능성 낮음 | 설명 부족, 속성 미완성, stale 지속 | search/home 후순위 |
| `policy_repeat_offense` | 같은 정책 위반 반복 | 금칙어/연락처/광고성 문구 재범 | listing/user 단위 단계 상향 |

### 63.6 노출면(surface)별 영향 계약
| surface | 가능한 영향 | 비고 |
|---|---|---|
| `search` | 기본 정렬 점수 감점, facet 기본 결과 제외, reserved 포함 필터에서만 노출 | 검색 품질 보호 목적 |
| `home_recommendation` | 개인화 추천/인기 슬롯 제외 | 직접 검색은 허용 가능 |
| `saved_search_notification` | 신규매물/가격변경/상태복귀 알림 발송 제외 | 알림 소음/피해 방지 |
| `public_profile` | 프로필 신뢰 배지 보수적 표현, 공개 프로필 임시 축소 | 민감 사유 직접 노출 금지 |
| `review_feed` | 최근 후기 목록 비노출, 집계 hold | review abuse 대응 |
| `chat_initiation` | 신규 채팅 시작 전 추가 검토/경고/쿨다운 | 강한 패널티일수록 명시 제한 전환 고려 |

원칙:
- 같은 객체라도 `search`는 제한하고 `direct_link` 접근은 허용할 수 있다.
- `saved_search_notification` 억제는 노출보다 사용자 보호 우선 목적이 강하므로 독립 토글이 필요하다.

### 63.7 사용자 고지 수준
| `userDisclosureLevel` | 사용자 문구 예시 | 사용 시점 |
|---|---|---|
| `silent` | 직접 문구 없음 | 내부 watch, 미세 랭킹 조정 |
| `soft_hint` | `노출 품질 향상을 위해 일부 정보를 보완해 주세요` | quality_low, metadata 부족 |
| `explicit_notice` | `현재 매물은 검토 후 다시 노출될 수 있어요` | hidden_until_review, review_only |

원칙:
- `fraud_signal`, `contact_evasion`처럼 조사 회피를 유도할 수 있는 사유는 상세 근거를 과도하게 노출하지 않는다.
- `quality_low`나 `metadata 부족`은 사용자에게 수정 기회를 주는 편이 제품 품질에 유리하므로 `soft_hint`를 기본으로 한다.

### 63.8 패널티 적용 단위
#### listing 단위
- 특정 매물만 노출 제한
- 가장 흔한 적용 단위
- 허위매물 의심, 품질 저하, spam_listing, reserved_abuse에 적합

#### user 단위
- 계정 전체의 추천/알림/신규 채팅 노출 정책 조정
- 다계정/반복 사기 의심/정책 재범에 적합
- 과도하게 넓은 영향이므로 review와 만료 규칙이 더 엄격해야 함

#### review 단위
- 특정 후기의 공개/집계 hold
- review_abuse, 개인정보 포함, 보복성 후기 대응에 적합

#### search_term_cluster 단위
- 특정 검색어/카테고리에서만 노출 감점
- 스팸이 특정 아이템명 군에서 집중될 때 사용 가능
- Post-MVP 우선 검토

### 63.9 트리거와 기본 action matrix
| 트리거 | 기본 상태 | 영향 surface | 해제 기본안 |
|---|---|---|---|
| 설명/속성 미완성 + stale 장기화 | `limited_exposure` | `search`, `home_recommendation` | 사용자가 수정하거나 freshness 회복 시 |
| reserved 장기 유지 + 반복 복귀 | `limited_exposure` | `search`, `saved_search_notification` | 7일 관찰 후 재계산 |
| 신규계정 고가매물 + 다량 문의 유도 | `review_only` | `home_recommendation`, `chat_initiation` | 인증/첫 정상 거래 후 완화 |
| 외부 연락처 반복 유도 | `review_only` 또는 `hidden_until_review` | `chat_initiation`, `search` | 운영 검토/수정 후 |
| 허위매물 강한 의심 + 신고 다발 | `hidden_until_review` | 전체 공개 surface | 운영 판정 후 |
| 후기 보복 패턴 | `review_only` | `review_feed`, trust aggregate | 운영 재심/만료 |

### 63.10 운영 큐 연동
- `watch`는 기본적으로 자동 관찰 대상이지만, 아래 조건이면 운영 큐로 승격한다.
  1. 같은 `penaltyReasonFamily`가 짧은 기간 반복될 때
  2. `limited_exposure`가 2회 이상 재적용될 때
  3. `hidden_until_review`가 30분 이상 지속될 때
  4. 신고/분쟁/제재 객체와 링크될 때
- 운영 큐에는 최소 아래 정보가 보여야 한다.
  - penalty scope/type
  - active surfaces
  - trigger summary
  - linked report/restriction ids
  - first applied at / latest recomputed at
  - current disclosure level
  - suggested next action (`keep`, `release`, `escalate_restriction`, `request_profile_completion`)

### 63.11 검색/추천 랭킹 파생 규칙
- search ranker는 `visibilityPenaltyState`를 독립 feature로 사용한다.
- `limited_exposure`는 hard filter보다 score 감점을 우선 적용한다.
- `review_only`는 추천/인기 슬롯에서는 제외하되 exact match 직접검색에서는 완전 제거하지 않는 안을 기본으로 한다.
- `hidden_until_review`는 검색, 홈, saved search, SEO 공개면에서 제외한다.
- 랭킹 문서 파생용 최소 필드:
  - `penaltyImpactTier`
  - `penaltyReasonFamily`
  - `penaltyAppliedAt`
  - `penaltyExpiresAt`
  - `manualOverrideYn`

### 63.12 알림/성장 surface 파생 규칙
- `saved_search_notification`이 억제된 매물은 신규 매물 알림, 가격변경 알림, 상태복귀 알림 fanout 대상에서 제외한다.
- `home_recommendation` 제외 매물은 cold-start 추천, 인기 매물, 재방문 추천에도 포함하지 않는다.
- growth CRM/리마인드 시스템은 visibility penalty가 높은 계정에 대해 프로모션 알림을 보내지 않는다.
- 운영/정책 알림은 억제하지 않되, 사용자에게 필요한 수정/검토 안내는 유지한다.

### 63.13 DB / read model / API 파생 기준
#### write model 후보
- `VisibilityPenalty`
  - `penaltyId`
  - `scopeType`
  - `scopeId`
  - `visibilityPenaltyState`
  - `penaltyReasonFamily`
  - `impactTier`
  - `activeSurfacesJson`
  - `userDisclosureLevel`
  - `triggerSource`
  - `evidenceSummaryJson`
  - `appliedAt`
  - `expiresAt`
  - `releasedAt`
  - `releaseReasonCode`
  - `createdByActorType`
  - `createdByActorId`

#### read model 후보
- `ListingVisibilitySnapshot`
- `UserExposureSnapshot`
- `ReviewPublicationSnapshot`

#### API/응답 필드 후보
- 작성자/운영자 전용 응답에는 아래를 포함 가능
```json
{
  "visibilitySnapshot": {
    "visibilityPenaltyState": "limited_exposure",
    "penaltyReasonFamily": "reserved_abuse",
    "activeSurfaces": ["search", "saved_search_notification"],
    "userDisclosureLevel": "soft_hint",
    "expiresAt": "2026-03-21T00:00:00+09:00"
  }
}
```
- 일반 사용자/비회원 API에는 내부 reason family를 그대로 노출하지 않고, 필요 시 `policyHints`만 제공한다.

### 63.14 이벤트 / analytics
필수 이벤트 후보:
- `visibility_penalty_applied`
- `visibility_penalty_recomputed`
- `visibility_penalty_released`
- `visibility_penalty_escalated_to_restriction`
- `visibility_penalty_user_hint_shown`
- `visibility_penalty_user_fix_completed`

공통 속성 후보:
- `scopeType`
- `penaltyReasonFamily`
- `visibilityPenaltyState`
- `impactTier`
- `triggerSource`
- `surfaceCount`
- `linkedReportYn`
- `timeSinceAccountCreationBucket`

핵심 KPI 후보:
- 패널티 적용 후 정상 수정/회복률
- 패널티 적용 후 신고율 변화
- 오탐 복구율
- `limited_exposure -> completed trade` 전환율
- `hidden_until_review` 평균 체류 시간

### 63.15 해제 / 재심 / 에스컬레이션 원칙
- 패널티는 아래 중 하나로 종료된다.
  1. **자동 만료**: time_based expiry 도래
  2. **행동 회복**: 정상 완료 거래, 프로필/매물 수정, 인증 완료
  3. **운영 해제**: manual release
  4. **명시 제재 승격**: restriction/ban으로 이관
- 같은 reason family가 반복되면 `watch -> limited_exposure -> review_only -> hidden_until_review -> restriction candidate` 순으로 상향 가능하다.
- 사용자가 수정 가능한 사안(`quality_low`)은 수정 완료 시 즉시 재계산을 시도한다.
- 사용자가 직접 이의를 제기할 수 있는 것은 `explicit_notice` 수준부터를 기본안으로 한다.

### 63.16 오남용 방지 원칙
- shadow moderation은 운영 편의성만으로 남용하지 않는다.
- 검색 품질 개선용 패널티와 안전/정책용 패널티를 혼동하지 않는다.
- 사용자 피해가 큰 사안은 조용한 제한으로 오래 끌지 말고 명시 검토/제재로 전환한다.
- 운영자는 패널티가 실제 성사율 개선에 기여하는지, 단순 전환 저하만 만드는지 주기적으로 검토해야 한다.

### 63.17 오픈 질문
- `limited_exposure` 상태의 매물을 작성자에게 얼마나 구체적으로 알려줄 것인가?
- review abuse/quality_low 같은 제품 품질 패널티와 fraud/contact_evasion 같은 안전 패널티를 별도 객체로 분리할 것인가?
- 신규 계정 고가매물에 대한 review_only를 인증/첫 완료 거래와 어떻게 연결할 것인가?
- 추천 제외만 적용하고 검색은 유지하는 정책이 실제 사용자 피해를 줄이는지 런칭 후 관찰이 필요하다.


## 64. 노출 설명(Exposure Explanation) / 사용자 복구 액션 / 운영 해제 UX canonical 계약
### 64.1 목적
- `Visibility Penalty`가 내부 랭킹 플래그에만 머물지 않고, 실제 사용자 화면·운영 툴·알림·API에서 일관된 설명과 복구 흐름으로 이어지게 한다.
- 과도한 shadow moderation으로 인해 작성자가 “왜 노출이 안 되는지” 전혀 모르는 상태를 줄이되, 사기 탐지·우회 시도에 악용될 만큼 상세한 내부 근거는 노출하지 않는다.
- 제품/디자인/운영이 같은 기준으로 `힌트 제공`, `수정 유도`, `검토 대기`, `해제`, `제재 승격`을 표현하도록 canonical vocabulary를 정의한다.

### 64.2 핵심 vocabulary
| 필드 | 의미 | 후보 값 |
|---|---|---|
| `exposureDecisionState` | 현재 사용자에게 어떤 설명/조치를 제공하는지 | `normal` / `hinted_recoverable` / `awaiting_user_fix` / `awaiting_staff_review` / `released` / `escalated_restriction` |
| `userRecoverabilityTier` | 사용자가 스스로 회복 가능한 정도 | `self_fixable` / `self_fixable_after_ack` / `staff_only` / `non_recoverable_user_side` |
| `exposureHintMode` | 사용자에게 보이는 설명 강도 | `none` / `banner` / `inline_field_hint` / `blocking_sheet` / `status_chip` |
| `recoveryActionType` | 회복을 위해 유도되는 행동 종류 | `complete_listing_fields` / `refresh_listing_content` / `verify_account` / `ack_policy_notice` / `request_review` / `wait_for_recompute` |
| `recoveryProgressState` | 사용자의 복구 진행 상태 | `not_started` / `in_progress` / `submitted` / `recompute_pending` / `resolved` / `failed_reappeared` |
| `operatorDispositionType` | 운영자가 내릴 수 있는 노출 관련 판정 | `keep_penalty` / `release_penalty` / `narrow_scope` / `extend_expiry` / `request_user_fix` / `escalate_to_restriction` |
| `exposureSurfaceVisibility` | 특정 surface에서 사용자에게 현재 상태를 보이는지 | `hidden_from_user` / `owner_visible` / `participant_visible` / `staff_visible_only` |
| `transparencyGuardLevel` | 설명 과정에서 내부 탐지 규칙을 얼마나 감출지 | `low_risk_explainable` / `guarded_summary` / `safety_sensitive` |

### 64.3 설계 원칙
1. **복구 가능한 문제는 복구 경로를 준다**: `quality_low`, metadata 부족, stale 상태는 무조건 조용히 깎기보다 수정 CTA를 제공한다.
2. **안전 민감 사유는 요약만 노출한다**: `fraud_signal`, `contact_evasion` 계열은 구체 탐지 근거를 노출하지 않는다.
3. **내부 패널티와 사용자 경험을 분리하지 않는다**: penalty가 있으면 owner surface에는 최소한의 상태/가이드가 있어야 한다.
4. **무한 검토 대기 금지**: `awaiting_staff_review`는 SLA와 예상 다음 단계를 함께 가져야 한다.
5. **재발 추적 가능성 확보**: 사용자가 수정 후 다시 같은 penalty가 붙는 경우 `failed_reappeared`로 추적한다.

### 64.4 노출 설명 상태 정의
| 상태 | 의미 | owner 화면 | 운영 처리 기본안 |
|---|---|---|---|
| `normal` | 설명할 패널티 없음 | 노출 상태 표시 없음 | 없음 |
| `hinted_recoverable` | 경미한 노출 저하가 있으나 수정 가능 | 부드러운 배너/필드 힌트 | 자동 재계산 우선 |
| `awaiting_user_fix` | 사용자의 수정/확인 액션이 필요 | 명시 CTA + 남은 작업 표시 | 수정 기한/재계산 큐 연결 |
| `awaiting_staff_review` | 사용자가 더 해도 자동 회복 불가, 운영 검토 필요 | 검토 대기 배지 + 예상 처리 수준 안내 | SLA 기반 큐 처리 |
| `released` | 패널티 해제 완료 | 회복 완료 토스트/이력 | 종료 이력 저장 |
| `escalated_restriction` | 노출 문제를 넘어 명시 제재로 승격 | 제재/제한 화면으로 전환 | restriction 객체로 이관 |

### 64.5 reason family별 기본 transparency / recoverability 매핑
| `penaltyReasonFamily` | 기본 hint mode | recoverability | transparency guard | 기본 CTA |
|---|---|---|---|---|
| `quality_low` | `banner` + `inline_field_hint` | `self_fixable` | `low_risk_explainable` | `complete_listing_fields`, `refresh_listing_content` |
| `reserved_abuse` | `banner` | `self_fixable_after_ack` | `guarded_summary` | `ack_policy_notice`, `wait_for_recompute` |
| `spam_listing` | `banner` | `self_fixable` | `guarded_summary` | `refresh_listing_content` |
| `high_risk_new_account` | `status_chip` | `self_fixable_after_ack` | `guarded_summary` | `verify_account` |
| `contact_evasion` | `blocking_sheet` 또는 `banner` | `staff_only` | `safety_sensitive` | `ack_policy_notice`, `request_review` |
| `fraud_signal` | `status_chip` | `staff_only` | `safety_sensitive` | `request_review` 또는 없음 |
| `review_abuse` | `status_chip` | `staff_only` | `guarded_summary` | `request_review` |
| `policy_repeat_offense` | `blocking_sheet` | `non_recoverable_user_side` | `safety_sensitive` | 없음 또는 support 진입 |

원칙:
- 같은 family라도 `triggerSource=manual_staff`이거나 linked report가 있으면 설명 강도를 보수적으로 조정한다.
- `self_fixable` 계열은 수정 후 재계산 ETA를 가급적 사용자에게 보여준다.

### 64.6 owner-facing surface 계약
#### 내 매물 목록
- 카드 레벨에는 아래만 노출한다.
  - `노출 낮음` / `검토 중` / `수정 필요` 같은 status chip
  - 회복 가능 여부(`수정하면 회복될 수 있어요`)
  - 즉시 CTA 1개 (`정보 보완`, `본인확인`, `검토 요청`)
- 카드에서 내부 risk score, report count, 탐지 룰명은 노출하지 않는다.

#### 매물 상세/수정 화면
- `quality_low` 계열은 필드 단위 힌트를 표시할 수 있어야 한다.
  - 예: 이미지 없음, 설명 부족, 속성 누락, 가격 단위 모호
- `awaiting_user_fix` 상태에서는 상단 배너에 다음을 포함한다.
  - 현재 상태 요약
  - 필요한 액션 목록
  - 완료 후 재계산/반영 예상 방식
- `awaiting_staff_review`는 “검토 후 다시 노출될 수 있음” 수준으로만 안내하고 세부 근거는 숨긴다.

#### 홈/알림
- 노출 패널티는 홈 피드 자체에서 일반 사용자에게 드러내지 않는다.
- owner에게만 `내 매물 점검 필요` 유형 알림을 보낼 수 있다.
- 알림 카피는 문제 지적보다 회복 행동 중심으로 작성한다.

### 64.7 사용자 복구 action family
| action type | UX 의미 | 성공 조건 | side effect |
|---|---|---|---|
| `complete_listing_fields` | 누락 필드/속성/이미지 보완 | required checklist 충족 | 재계산 큐 적재 |
| `refresh_listing_content` | 제목/설명/가격/수량을 더 명확히 수정 | spam/quality 기준 개선 | freshness 갱신, 재계산 |
| `verify_account` | 계정 또는 거래 가능 상태 보강 | verification level 상승 | 일부 surface 즉시 회복 가능 |
| `ack_policy_notice` | 특정 정책을 읽고 재확인 | ack 기록 저장 | 같은 사유 재발 시 더 강한 단계 가능 |
| `request_review` | 운영 재검토 요청 | request submit 성공 | support/moderation queue 연결 |
| `wait_for_recompute` | 별도 수정 없이 시간/행동 회복 대기 | expiry 또는 rule recompute | user가 직접 바꿀 수 없음 |

원칙:
- 동시에 여러 action이 필요해도 primary CTA는 1개만 강하게 보여준다.
- `request_review`는 `explicit_notice` 또는 `awaiting_staff_review` 수준에서만 노출하는 것을 기본안으로 한다.

### 64.8 운영자 해제 / 축소 / 승격 workflow
1. **사건 확인**: linked report, evidence summary, recent changes, prior penalties 확인
2. **recoverability 판단**: 사용자가 고칠 수 있는 문제인지, staff 개입이 필요한지 결정
3. **operator disposition 선택**
   - `keep_penalty`
   - `release_penalty`
   - `narrow_scope`
   - `extend_expiry`
   - `request_user_fix`
   - `escalate_to_restriction`
4. **사용자 커뮤니케이션 생성**: soft hint / explicit notice / no disclosure
5. **재계산 또는 승격 실행**: 랭킹·알림·프로필 surface 반영

운영 원칙:
- `release_penalty`는 reason/evidence가 사라졌거나 오탐이 확인된 경우에만 사용한다.
- `narrow_scope`는 예를 들어 `search + home` 제한을 `home만 제한`으로 줄이는 식의 부분 완화다.
- `escalate_to_restriction`은 visibility penalty만으로 사용자 피해를 막기 어렵다고 판단될 때만 사용한다.

### 64.9 SLA / 자동 전이 기준
| 상태 | 기본 SLA/타이머 | 자동 처리 |
|---|---|---|
| `hinted_recoverable` | 72시간 내 사용자 수정 관찰 | 수정 없으면 유지 또는 약한 감점 지속 |
| `awaiting_user_fix` | 7일 내 수정/ack 기대 | 미수정 시 penalty 유지 또는 scope 확장 가능 |
| `awaiting_staff_review` | P2: 24h, P3: 72h 기본안 | SLA 초과 시 queue escalation |
| `released` | 즉시 반영 또는 다음 recompute cycle | release event 발행 |
| `escalated_restriction` | restriction policy SLA 따름 | penalty 객체는 종료 상태로 전환 |

### 64.10 API / read model 파생 기준
#### read model 후보
- `OwnerListingExposureSnapshot`
- `ExposureRecoveryTask`
- `OperatorExposureReviewItem`

#### owner-facing API 응답 후보
```json
{
  "exposureExplanation": {
    "exposureDecisionState": "awaiting_user_fix",
    "hintMode": "banner",
    "summaryLabel": "노출 품질을 높이려면 정보 보완이 필요해요",
    "recoverabilityTier": "self_fixable",
    "primaryRecoveryAction": "complete_listing_fields",
    "recoveryChecklist": [
      "대표 이미지 추가",
      "강화/옵션 정보 보완"
    ],
    "recoveryProgressState": "not_started",
    "nextRecomputeAt": "2026-03-15T10:00:00+09:00"
  }
}
```

#### operator API 필드 후보
- `linkedPenaltyId`
- `operatorDispositionType`
- `transparencyGuardLevel`
- `userVisibleSummaryTemplateKey`
- `recoveryEligibility`
- `suggestedScopeChange`

### 64.11 analytics / KPI
필수 이벤트 후보:
- `exposure_hint_impression`
- `exposure_recovery_cta_click`
- `exposure_recovery_submitted`
- `exposure_recovery_recompute_passed`
- `exposure_recovery_recompute_failed`
- `exposure_operator_release`
- `exposure_escalated_to_restriction`

핵심 관찰 지표 후보:
- hint 노출 후 24시간 내 수정률
- 수정 후 penalty 해제율
- `request_review` 대비 실제 release 비율
- 오탐으로 판정된 `awaiting_staff_review` 비율
- recoverable reason family별 평균 회복 시간
- recovery 시도 후 재발(`failed_reappeared`) 비율

### 64.12 화면/운영/정책 파생 포인트
- **화면명세**: 내 매물 카드 chip, 상세 배너, 필드별 힌트, 검토중 empty/error copy
- **DB 스키마**: penalty와 별도인 `recovery task` 저장 필요 여부, ack 이력, user-visible summary template key
- **API 스펙**: owner 응답과 public 응답의 필드 분리, unknown hint mode fallback 규칙
- **운영정책**: recoverable vs non-recoverable 사유 분류표, 해제/재심/승격 템플릿
- **QA**: 같은 penalty라도 owner/public/staff surface에서 서로 다른 설명 수준이 정확히 유지되는지 검증

### 64.13 오픈 질문
- `quality_low` 회복 checklist를 category별로 얼마나 세분화할 것인가?
- `reserved_abuse`처럼 행동 패턴 계열 사유를 사용자에게 어디까지 self-fixable로 보여줄 것인가?
- `request_review`를 앱 안 단일 CTA로 둘지, SupportCase를 생성하는 흐름으로만 둘지 결정이 필요하다.
- 해제 직후 search/home/saved-search exposure가 즉시 반영되어야 하는지, 다음 recompute cycle까지 지연 허용할지 확정이 필요하다.

## 65. 사기 의심 리스크 신호(Fraud Suspicion Signal) / 거래 중단 보호 / 수동 검토 승격 canonical 계약
### 65.1 목적
- 허위매물, 외부연락처 유도, 선입금 강요, 조건 급변, 다계정/반복 패턴 같은 `사기 의심 정황`을 단순 신고 사유나 개별 moderation memo로 흩어두지 않고 하나의 canonical risk read model로 정의한다.
- 채팅/예약/당일 실행/내 거래/운영 큐/API/analytics가 모두 같은 기준으로 `지금 이 거래를 계속 진행해도 되는가`, `어떤 보호 조치를 즉시 켜야 하는가`, `언제 수동 검토나 명시 제재로 승격해야 하는가`를 해석하도록 한다.
- 이 섹션은 법적 사실 단정이 아니라 제품/운영 관점의 `사기 의심 위험 관리` 계약이며, 실제 확정 제재는 별도 moderation evidence와 운영 판정을 따른다.

### 65.2 핵심 개념
| 용어 | 의미 |
|---|---|
| `fraudSuspicionSnapshot` | 특정 listing/tradeThread/user pair 기준의 사기 의심 상태를 요약한 canonical read model |
| `fraudSuspicionState` | 현재 위험 단계 (`none` / `watch` / `user_warned` / `protected_flow` / `staff_review_required` / `restriction_escalated`) |
| `fraudSignalFamily` | 위험 정황 분류 (`off_platform_diversion` / `advance_payment_pressure` / `identity_mismatch` / `terms_instability` / `repeat_pattern_abuse` / `evidence_inconsistency`) |
| `signalSeverityTier` | 개별 신호 강도 (`low` / `medium` / `high` / `critical`) |
| `tradeProtectionMode` | 사용자 보호 조치 수준 (`notice_only` / `send_guard` / `action_gated` / `temporary_hold`) |
| `counterpartyDisclosureLevel` | 상대방에게 보여주는 경고 수준 (`none` / `generic_caution` / `action_specific_warning` / `transaction_pause`) |
| `fraudReviewDisposition` | 운영 검토 결과 (`clear` / `monitor` / `sustain_protection` / `escalate_restriction`) |

원칙:
- `fraudSuspicionState`는 확정 제재와 다르며, 거래 보호를 위한 선제적 read model이다.
- 동일한 신호라도 `listing-only`, `tradeThread-only`, `account-wide` scope가 다를 수 있다.
- 사용자에게는 낙인 표현 대신 `주의 필요`, `플랫폼 내 기록을 유지하세요`, `이 거래는 잠시 검토 중이에요`처럼 행동 중심 문구를 사용한다.

### 65.3 신호 패밀리 정의
#### 65.3.1 `off_platform_diversion`
- 외부 메신저, 오픈채팅, 개인 연락처, 외부 결제 링크로 대화를 이탈시키려는 패턴
- 본문/이미지/OCR/링크/반복 문구 탐지와 연결
- 기본 보호: `send_guard` 이상

#### 65.3.2 `advance_payment_pressure`
- 선입금/선전달/계좌 선공유를 반복 요구하거나 플랫폼 밖 정산을 강하게 압박하는 패턴
- 예약 확정 전/상호확인 전 요구일수록 severity 가중
- 기본 보호: `action_specific_warning`, 필요 시 `temporary_hold`

#### 65.3.3 `identity_mismatch`
- 직전 예약/식별정보와 다른 캐릭터명/연락처/접선 주체가 반복 제시되거나, 상호 확인된 식별자와 충돌하는 패턴
- 거래 직전 급격한 변경, 다수 상대에게 상이한 식별정보 제시 포함
- 기본 보호: `send_guard` 또는 `action_gated`

#### 65.3.4 `terms_instability`
- 가격/수량/장소/시간/정산 방식이 예약 직전 반복적으로 바뀌고 상대 재확인이 누락되는 패턴
- 단순 재일정이 아니라 거래 실행을 불안정하게 만드는 급변을 의미
- 기본 보호: `notice_only`에서 시작해 반복 시 `protected_flow` 승격

#### 65.3.5 `repeat_pattern_abuse`
- 신규/저신뢰 계정의 고가 반복 등록, 다수 상대 동시 유도, 반복 no-show/dispute 결합, 다계정 의심 패턴
- 단일 메시지보다 계정/기기/기간 관점의 패턴 신호
- 기본 보호: `staff_review_required` 후보

#### 65.3.6 `evidence_inconsistency`
- 완료 주장, 도착 주장, 정산 주장과 실제 플랫폼 로그/예약/식별정보/증빙이 상충하는 패턴
- 분쟁, no-show, completion dispute, chargeback 유사 소명 판단의 입력 신호
- 기본 보호: `transaction_pause` 또는 운영 검토 승격

### 65.4 신호 생성 소스
| 소스 | 예시 | 기본 신뢰도 |
|---|---|---|
| message/content detector | 연락처/오픈채팅/선입금 표현 | 중간 |
| reservation/execution drift | 직전 장소/시간/식별정보 급변 | 중간 |
| trust/restriction history | 반복 no-show, prior restriction, multi-account risk | 중간~높음 |
| user report | 사기 의심 신고, 상대 경고 | 낮음~중간 |
| staff observation | 운영 수동 메모, case linkage | 높음 |
| evidence comparison | 채팅/예약/증빙 간 불일치 | 높음 |

원칙:
- 단일 저신뢰 신호만으로 `restriction_escalated`로 바로 이동하지 않는다.
- 반대로 `critical` 신호가 platform safety를 직접 위협하면 운영 검토 전이라도 `temporary_hold`를 허용한다.

### 65.5 상태 전이 기본안
| 현재 상태 | 트리거 | 다음 상태 |
|---|---|---|
| `none` | low/medium signal 최초 감지 | `watch` |
| `watch` | 사용자 경고 필요 판단 | `user_warned` |
| `watch`/`user_warned` | 메시지 전송/예약/완료 액션 보호 필요 | `protected_flow` |
| `protected_flow` | high/critical signal 추가 또는 반복 재발 | `staff_review_required` |
| `staff_review_required` | 운영이 제한 필요 판정 | `restriction_escalated` |
| `staff_review_required` | 운영이 오탐 또는 경미 판정 | `watch` 또는 `none` |
| `restriction_escalated` | 제재 종료/오탐 해제 | 별도 restriction recovery 정책에 따름 |

상태 규칙:
- `watch`는 내부 모니터링 중심이며 사용자에게 반드시 노출하지 않아도 된다.
- `user_warned`는 최소 1개 surface에서 사용자 행동 가이드를 제공해야 한다.
- `protected_flow`는 거래를 완전히 막지 않더라도 메시지 전송 경고, 예약 재확인, 완료 확정 hold 같은 보호 장치를 켜야 한다.

### 65.6 보호 조치(`tradeProtectionMode`) 정의
| mode | 의미 | 대표 동작 |
|---|---|---|
| `notice_only` | 일반 주의만 제공 | 채팅 상단 안전 배너, 예약 직전 caution copy |
| `send_guard` | 특정 메시지/행동 전 경고 | 연락처/계좌/외부링크 전송 전 soft block |
| `action_gated` | 특정 액션에 추가 확인 필요 | 장소 급변 후 상대 ack 없으면 완료/no-show claim 제한 |
| `temporary_hold` | 일시 중단 후 검토 | 거래완료 확정, 예약 확정, listing 노출 일부 hold |

세부 원칙:
- `temporary_hold`는 거래 중단 비용이 높으므로 `critical` 또는 다중 high signal 결합일 때만 사용한다.
- 보호 조치는 가능하면 scope-specific으로 적용한다. 예: 특정 tradeThread만 hold, 계정 전역은 아님.
- 보호 조치의 이유와 복구 경로는 owner/counterparty/staff마다 다른 설명 수준을 가져야 한다.

### 65.7 surface별 해석 규칙
#### 채팅
- `notice_only`: 상단 안전 배너 + 관련 quick tip 노출
- `send_guard`: 외부연락처/계좌/선입금 표현 전송 전 재확인 모달
- `action_gated`: 예약 확정/완료 요청 전 추가 확인 step 요구
- `temporary_hold`: 일부 전송/완료 액션 disabled + 운영 검토중 배너

#### 예약/당일 실행 카드
- `terms_instability`, `identity_mismatch`가 있으면 `상대 재확인 필요` 배지를 우선 노출
- `advance_payment_pressure`가 있으면 결제/정산 관련 주의 문구와 off-platform caution을 강화
- `temporary_hold` 중에는 `거래완료`, `노쇼 확정`, `자동확정` 같은 irreversible action을 막을 수 있어야 한다.

#### 내 거래/상세
- risk가 높을수록 `nextBestAction`은 `계속 대화`보다 `정보 재확인`, `운영 문의`, `증빙 보관` 쪽으로 바뀌어야 한다.
- generic warning만 필요한 경우 상대에게 낙인처럼 보이는 라벨은 쓰지 않는다.

#### 운영 큐
- `staff_review_required`는 report/support/no-show/dispute와 별개로 anti-fraud queue projection에 나타나야 한다.
- 큐에는 `signal families`, `latest trigger`, `affected surface`, `recommended protection mode`, `linked cases`가 보여야 한다.

### 65.8 운영 액션 계약
| 액션 | 설명 | 제약 |
|---|---|---|
| `keep_monitoring` | watch 유지, 추가 자동화만 허용 | 사용자 노출 없음 가능 |
| `warn_user` | 주의 문구/정책 링크 전달 | 확정 제재처럼 보이면 안 됨 |
| `enable_protection_mode` | tradeThread/listing/account에 보호 모드 적용 | scope 기록 필수 |
| `request_more_evidence` | 양측/신고자에게 추가자료 요청 | SLA와 제출기한 필요 |
| `pause_irreversible_actions` | 완료 확정, 예약 확정, 일부 노출 hold | high 이상 근거 필요 |
| `escalate_to_restriction` | 명시 제재/노출제한으로 승격 | moderation/restriction policy 연계 |
| `clear_signal` | 오탐 또는 해소 판정 | recompute 및 user-facing warning 해제 |

운영 원칙:
- anti-fraud 보호와 policy restriction은 별도 객체로 저장하되 linkage는 유지한다.
- `pause_irreversible_actions`는 거래를 영구 차단하는 게 아니라 `잘못된 완료 확정/노쇼 확정`을 막는 임시 안전장치다.
- clear 이후에도 동일 family 재발 여부를 위해 historical signal은 보존해야 한다.

### 65.9 API / read model 파생 기준
#### read model 후보
- `FraudSuspicionSnapshot`
- `TradeProtectionState`
- `FraudReviewQueueItem`

#### participant-facing 응답 후보
```json
{
  "fraudProtection": {
    "fraudSuspicionState": "protected_flow",
    "counterpartyDisclosureLevel": "generic_caution",
    "tradeProtectionMode": "action_gated",
    "warningLabel": "거래 전 마지막 조건을 다시 확인해 주세요",
    "blockedActions": ["auto_complete_confirm"],
    "requiredChecks": [
      "meeting_identity_reconfirm",
      "settlement_method_reconfirm"
    ],
    "supportEntryEnabled": true
  }
}
```

#### operator API 필드 후보
- `fraudSignalFamilies[]`
- `highestSeverityTier`
- `signalEvidenceSummary`
- `affectedScopes[]`
- `activeProtectionMode`
- `recommendedDisposition`
- `linkedRestrictionId`
- `recomputeEligibleAt`

### 65.10 DB / event / audit 시사점
후보 저장 단위:
- `fraud_signal_event`
- `fraud_protection_state`
- `fraud_review_case_link`

필수 필드 후보:
- `signalId`
- `subjectType` (`listing` / `trade_thread` / `user_pair` / `account`)
- `subjectId`
- `fraudSignalFamily`
- `severityTier`
- `signalSourceType`
- `evidenceRef`
- `triggeredAt`
- `clearedAt`
- `activeYn`

이벤트 후보:
- `fraud_signal_detected`
- `fraud_warning_shown`
- `trade_protection_enabled`
- `trade_protection_blocked_action`
- `fraud_review_escalated`
- `fraud_signal_cleared`

### 65.11 analytics / KPI
필수 이벤트 후보:
- `fraud_guard_impression`
- `fraud_guard_send_abandoned`
- `fraud_guard_override_confirmed`
- `trade_protection_action_blocked`
- `fraud_review_case_created`
- `fraud_review_case_escalated_to_restriction`

핵심 관찰 지표 후보:
- signal family별 guard 노출 대비 실제 신고/분쟁 전환율
- `protected_flow` 적용 거래의 completion rate vs dispute rate
- `temporary_hold` 후 clear 비율(오탐 추정)
- 외부연락처/선입금 guard 이후 이탈률과 신고율 변화
- `restriction_escalated` 전 단계에서 조기 경고로 종결된 비율

### 65.12 오픈 질문
- `temporary_hold`가 예약 확정까지 막을지, 완료/노쇼 같은 irreversible action만 막을지 MVP 범위를 확정해야 한다.
- fraud signal을 상대방에게 얼마나 직접적으로 설명할지(`generic_caution` vs `transaction_pause`) 정책 선택이 필요하다.
- 다계정/기기 기반 신호를 MVP에서 어디까지 사용할지, false positive 책임 주체를 운영정책에 어떻게 둘지 결정이 필요하다.
- anti-fraud queue를 기존 moderation queue의 subtype으로 둘지, 별도 queue/workflow로 분리할지 운영 설계 확정이 필요하다.

## 66. 보호조치 이의제기(Protection Appeal) / 추가 증빙 제출 / 해제 결과 반영 canonical 계약
### 66.1 목표
- fraud suspicion, visibility penalty, restriction, temporary hold 같은 보호조치가 발동된 뒤 사용자가 왜 막혔는지 이해하고, 어떤 경우 재심을 요청할 수 있으며, 해제 결과가 어디에 어떻게 반영되는지 일관되게 정의한다.
- 사용자 안내 UX, support case, moderation queue, admin action, audit log, trust/restriction recalculation이 서로 다른 사건 키와 상태 모델을 쓰지 않도록 한다.
- 오탐(false positive) 복구를 제품의 일부로 다루어, “막는 것”만이 아니라 “정당한 사용자 복구”도 implementation-ready 수준으로 명문화한다.

### 66.2 적용 범위
본 섹션은 아래 보호조치 계열에 공통 적용된다.
- `visibility penalty` 계열: 검색 후순위, 추천 제외, 검토 대기, publish hold
- `fraud protection` 계열: caution banner, action friction, temporary hold, manual review required
- `restriction` 계열: listing/chat/trade/review/support 일부 기능 제한
- `content moderation` 계열: 매물/채팅/후기/pro필 임시 비노출 후 재심 가능 상태

적용 원칙:
- 모든 보호조치가 appeal 가능해야 하는 것은 아니다.
- irreversible enforcement(영구 정지, 법적 이슈, 명백한 금지행위)는 이의제기 가능 범위를 더 좁게 둘 수 있다.
- 단, 사용자 표면에는 `appealable` 여부와 현재 가능한 다음 행동을 명시해야 한다.

### 66.3 핵심 vocabulary
| 용어 | 의미 |
|---|---|
| `protectionCase` | 하나의 보호조치 사건 단위. visibility/fraud/restriction/content review를 묶는 canonical case |
| `protectionOriginType` | 사건 발생 원천. `auto_rule`, `user_report`, `moderator_action`, `appeal_reopen`, `scheduled_recheck` |
| `protectionAppealState` | 사건 기준 이의제기 상태 |
| `appealSubmissionWindowState` | 재심 접수 가능 창의 현재 상태 |
| `appealEligibilityDisposition` | 접수 가능/불가와 그 사유 판정 |
| `appealEvidencePack` | 사용자가 제출한 텍스트/첨부/참조 링크/추가 설명 묶음 |
| `evidenceReviewDisposition` | 제출 증빙의 충분성/추가 요청 필요 여부 |
| `protectionLiftDecisionType` | 유지/부분완화/전면해제/더강한제재 전환 결정 |
| `protectionLiftEffectScope` | 어떤 surface와 capability에 해제/완화가 반영되는지 범위 |
| `appealOutcomeDisclosureLevel` | 사용자에게 결과를 어느 정도까지 설명할지 수준 |

### 66.4 상태 모델
#### 66.4.1 `protectionAppealState`
- `not_opened`: 아직 이의제기되지 않음
- `eligible_to_submit`: 접수 가능 상태
- `submitted`: 접수 완료, triage 대기
- `needs_more_evidence`: 추가 증빙 요청됨
- `under_review`: 운영/전문가 검토 중
- `decided_lifted`: 전면 해제 결정
- `decided_partially_relaxed`: 일부 완화 결정
- `decided_upheld`: 원조치 유지 결정
- `decided_escalated`: 재심 과정에서 더 강한 제재로 전환
- `closed_expired`: 제출 가능 기간 만료 또는 추가 제출 기한 만료로 종료
- `closed_non_appealable`: 애초에 재심 대상 아님

#### 66.4.2 `appealSubmissionWindowState`
- `open_now`
- `closing_soon`
- `expired`
- `temporarily_paused`
- `reopened_by_operator`

#### 66.4.3 `evidenceReviewDisposition`
- `sufficient_for_decision`
- `needs_clarification`
- `needs_attachment_or_screenshot`
- `duplicate_of_existing_record`
- `out_of_scope_submission`
- `malicious_or_abusive_submission`

### 66.5 사건 단위 / linkage 원칙
- appeal은 독립 문서가 아니라 기존 `protectionCase`에 연결된 하위 workflow다.
- 같은 보호조치에 대해 여러 차례 제출이 있더라도 canonical case id는 유지하고, 제출본은 `appealAttempt` 또는 `appealSubmission` 단위로 append-only 저장한다.
- `protectionCase`는 아래 객체들과 직접 연결 가능해야 한다.
  - `supportCase`
  - `report`
  - `moderationAction`
  - `restrictionAction`
  - `visibilityPenaltyRecord`
  - `fraudSignal` / `tradeProtectionAction`
- 사용자 앱/알림에서는 세부 내부 객체를 다 노출하지 않고 `protectionCase` 중심으로 설명한다.

### 66.6 이의제기 가능성 판정 규칙
#### 66.6.1 기본 원칙
- 시스템은 보호조치 생성 시점에 `appealEligibilityDisposition`과 `appealSubmissionDeadlineAt`를 함께 계산/저장해야 한다.
- 사용자는 보호조치 상세 또는 support entrypoint에서 현재 판정을 확인할 수 있어야 한다.

#### 66.6.2 `appealEligibilityDisposition` 후보
- `eligible_standard_window`
- `eligible_once_only`
- `eligible_after_cooldown`
- `ineligible_policy_final`
- `ineligible_legal_or_safety_lock`
- `ineligible_duplicate_pending`
- `ineligible_missing_user_action_first`

#### 66.6.3 기본 규칙
- `temporary_hold`, `visibility penalty`, `limited restriction`, `content hidden pending review`는 기본적으로 1회 이상 appeal 가능하다.
- 이미 `under_review` 또는 `needs_more_evidence` 상태인 동일 사건에는 중복 appeal 생성 대신 기존 사건으로 merge한다.
- 사용자에게 먼저 가능한 self-recovery action(예: 금칙어 수정, 연락처 제거, 프로필 재작성, 인증 완료)이 남아 있으면 appeal 전 self-recovery를 우선 유도할 수 있다.
- `permanent_ban`, `legal_hold`, `child_safety`, `severe fraud confirmed` 계열은 `closed_non_appealable` 또는 별도 오프라인 채널 유도로 제한할 수 있다.

### 66.7 사용자 surface requirements
#### 66.7.1 보호조치 배너/상세
사용자 화면은 최소 아래를 설명해야 한다.
- 어떤 기능/노출이 영향을 받는지
- 임시 조치인지 확정 조치인지
- 스스로 바로 해결 가능한 액션이 있는지
- appeal 가능 여부와 제출 기한
- 현재 검토 상태와 예상 다음 단계

필수 필드 후보:
- `caseId`
- `protectionType`
- `headline`
- `userFacingReasonCode`
- `impactSummary`
- `appealEligibilityDisposition`
- `appealSubmissionWindowState`
- `appealDeadlineAt`
- `selfRecoveryAction[]`
- `nextExpectedReviewAt`

#### 66.7.2 제출 UX
- 텍스트 설명 입력창
- 스크린샷/증빙 첨부
- 기존 채팅/매물/거래 thread 참조 선택
- `무엇을 검토해 달라는지`를 명시하는 reason selector
- 허위/욕설/중복 제출에 대한 경고

제출 UX 원칙:
- 사용자는 일반 채팅에 소명하지 않고 canonical appeal form으로 제출한다.
- 기존 support case가 있더라도 appeal submission과 자유 문의 메시지는 구분 저장한다.
- 제출 후에는 수정이 아니라 추가 제출(addendum) 방식이 기본이다.

### 66.8 운영 queue / 처리 규칙
#### 66.8.1 queue taxonomy
appeal queue는 아래 subtype으로 분류 가능해야 한다.
- `visibility_recovery`
- `fraud_hold_review`
- `restriction_appeal`
- `content_restore_request`
- `identity_or_verification_recheck`

#### 66.8.2 우선순위 기본안
- 거래 실행을 직접 막는 `temporary_hold`, `trade action blocked`는 높은 우선순위
- 검색 후순위/추천 제외만 있는 visibility appeal은 상대적으로 낮은 우선순위
- 동일 사용자의 진행 중 예약/당일 실행 건이 걸린 사건은 urgency 가산

#### 66.8.3 운영 액션 코드 후보
- `appeal_accept_full_lift`
- `appeal_accept_partial_relief`
- `appeal_reject_uphold`
- `appeal_request_more_evidence`
- `appeal_merge_duplicate`
- `appeal_reopen_case`
- `appeal_escalate_restriction`

### 66.9 결과 반영 / recalculation 원칙
#### 66.9.1 `protectionLiftDecisionType`
- `full_lift`
- `partial_relief`
- `uphold`
- `escalate`
- `expire_without_submission`
- `expire_without_response`

#### 66.9.2 `protectionLiftEffectScope`
- `search_visibility_only`
- `listing_publish_capability`
- `chat_send_capability`
- `trade_execute_capability`
- `profile_public_visibility`
- `review_publication`
- `support_contact_capability`
- `all_attached_surfaces`

#### 66.9.3 반영 규칙
- `full_lift` 시 관련 visibility penalty, fraud hold, temporary restriction의 active flag가 동일 case linkage 기준으로 함께 재계산될 수 있어야 한다.
- `partial_relief`는 일부 surface만 풀고 일부는 유지하는 결정이 가능해야 한다. 예: 검색 노출은 복구하되 예약 확정은 추가 검토 전까지 유지.
- `uphold`는 현재 조치를 유지하되, 동일 사유 재제출 cooldown을 함께 설정할 수 있다.
- `escalate`는 appeal 제출 내용에서 명백한 추가 위반/협박/허위가 드러난 경우에만 허용하며 별도 approval policy를 권장한다.

### 66.10 알림 / SLA / 만료 규칙
- appeal 접수 시 `submitted` 확인 알림을 발송한다.
- 추가 증빙 요청 시 `needs_more_evidence` 전환과 제출 기한을 함께 고지한다.
- 기한 내 미제출이면 `closed_expired`로 닫고, 기존 보호조치는 유지된다.
- 결과 통지에는 최소 `유지/부분완화/해제`와 영향 범위를 포함해야 한다.

기본 SLA 가정:
- 거래 실행 차단형 appeal: 24h 내 1차 검토
- visibility/content restore appeal: 72h 내 1차 검토
- 추가 자료 요청 후 사용자 응답 기한: 72h

### 66.11 API / DB 파생 기준
#### 66.11.1 API 후보
- `GET /me/protection-cases`
- `GET /me/protection-cases/{caseId}`
- `POST /me/protection-cases/{caseId}/appeals`
- `POST /me/protection-cases/{caseId}/appeals/{appealId}/attachments`
- `POST /admin/protection-cases/{caseId}/request-more-evidence`
- `POST /admin/protection-cases/{caseId}/decide`
- `POST /admin/protection-cases/{caseId}/reopen`

#### 66.11.2 테이블/객체 후보
- `ProtectionCase`
- `ProtectionImpactSnapshot`
- `ProtectionAppeal`
- `ProtectionAppealEvidence`
- `ProtectionDecisionHistory`

필수 필드 후보:
- `caseId`
- `subjectType`
- `subjectId`
- `protectionType`
- `originType`
- `appealState`
- `appealEligibilityDisposition`
- `appealDeadlineAt`
- `lastSubmittedAt`
- `decisionType`
- `decisionScope`
- `decidedAt`
- `decidedByAdminId`

### 66.12 analytics / KPI
필수 이벤트 후보:
- `protection_appeal_viewed`
- `protection_appeal_started`
- `protection_appeal_submitted`
- `protection_appeal_needs_more_evidence`
- `protection_appeal_decided`
- `protection_relief_applied`

핵심 관찰 지표 후보:
- 보호조치 유형별 appeal submission rate
- `submitted -> full_lift/partial_relief/uphold` 비율
- false positive 추정치(`full_lift` + 빠른 해제 비율)
- 추가 증빙 요청 후 응답률
- appeal 처리 전후 거래 실행/재신고/재위반 변화

### 66.13 오픈 질문
- `ProtectionCase`를 SupportCase와 완전히 분리할지, support는 conversation wrapper로만 둘지 정보구조 결정을 확정해야 한다.
- visibility penalty appeal을 사용자에게 proactive하게 열어둘지, 일정 수준 이상의 penalty에만 expose할지 정책 결정이 필요하다.
- `escalate` 결정을 재심 처리자 단독으로 허용할지, 2인 승인으로 둘지 운영 승인체계를 확정해야 한다.
- 추가 증빙 첨부 저장소를 dispute evidence와 통합할지, protection-case 전용 버킷/권한모델로 분리할지 아키텍처 선택이 필요하다.

## 67. 매물 등록 폼 schema / 단계형 입력 / validation severity canonical 계약
### 67.1 목적
- 매물 등록을 단순 `POST /listings` 폼 제출이 아니라, 모바일 기준으로 빠르게 완성 가능한 `구조화된 입력 플로우`로 정의한다.
- 화면/클라이언트 validation/API write DTO/draft 저장/운영 검수 큐가 서로 다른 기준을 쓰지 않도록 등록 폼 vocabulary를 고정한다.
- 같은 Listing schema라도 `필수 누락`, `경고성 품질 부족`, `정책 검토 필요`, `게시 불가`를 분리해 전환율과 안전성을 함께 관리한다.

### 67.2 설계 원칙
1. **게시 가능성 우선**: 사용자가 지금 바로 올릴 수 있는지(`publishReadinessState`)를 명확히 보여줘야 한다.
2. **검증 강도 분리**: 모든 문제를 blocking error로 다루지 않고 `info/warn/block/review_required`를 구분한다.
3. **draft 우선 저장**: 입력 도중 이탈해도 핵심 필드가 유실되지 않도록 step 단위 draft 저장을 기본으로 한다.
4. **카테고리 확장 가능성**: MVP는 공통 필드 중심이지만, attribute template이 붙어도 같은 field schema를 재사용할 수 있어야 한다.
5. **게시 직전 요약**: 사용자는 최종 게시 전 `검색 카드에 어떻게 보이는지`, `누락된 신뢰 신호가 무엇인지`를 확인할 수 있어야 한다.

### 67.3 핵심 vocabulary
#### 67.3.1 `stepId`
- `basics`: 거래 유형, 서버, 카테고리, 아이템명
- `pricing_quantity`: 가격 정책, 가격, 수량, 거래 단위
- `details`: 강화/옵션/설명/이미지
- `trade_preferences`: 거래 방식, 가능 시간, 희망 장소, 캐릭터명 공개 범위
- `review_publish`: 최종 요약, 정책 점검, 게시

원칙:
- MVP 화면은 3~5 step으로 구현 가능하되, backend/draft/event는 위 canonical `stepId`를 기준으로 저장한다.
- 단일 스크롤 폼을 채택하더라도 내부적으로는 step progression을 유지한다.

#### 67.3.2 `fieldGroup`
- `identity`
- `classification`
- `pricing`
- `quantity`
- `quality_signal`
- `execution_preference`
- `safety_policy`
- `publish_gate`

#### 67.3.3 `validationSeverity`
- `info`: 게시는 가능하나 품질 개선 힌트 제공
- `warn`: 게시는 가능하지만 검색/전환/신뢰에 불리할 수 있음
- `block`: 게시 불가, 수정 필수
- `review_required`: 게시 요청은 가능하나 즉시 공개 대신 검토/제한공개 후보

#### 67.3.4 `publishReadinessState`
- `not_started`
- `draft_incomplete`
- `ready_with_warnings`
- `ready_to_publish`
- `ready_but_review_required`
- `publish_blocked`

#### 67.3.5 `draftPromotionRule`
- `local_only`: 로컬 임시저장만 허용
- `server_draft_optional`: 서버 draft 저장 권장, 게시 전 동기화 가능
- `server_draft_required`: 멀티디바이스/검토 추적을 위해 서버 저장 필수
- `publish_atomic`: 최종 게시 시점에만 write aggregate 생성

### 67.4 등록 폼 canonical schema
```json
{
  "listingFormSchemaVersion": "2026-03-14",
  "steps": [
    {
      "stepId": "basics",
      "title": "무엇을 거래하나요?",
      "requiredFieldKeys": ["listingType", "serverId", "categoryId", "itemNameRaw"],
      "fieldGroups": ["identity", "classification"]
    },
    {
      "stepId": "pricing_quantity",
      "title": "가격과 수량을 정해주세요",
      "requiredFieldKeys": ["priceType", "quantity"],
      "fieldGroups": ["pricing", "quantity"]
    },
    {
      "stepId": "details",
      "title": "상세 정보를 보강하세요",
      "requiredFieldKeys": ["description"],
      "fieldGroups": ["quality_signal"]
    },
    {
      "stepId": "trade_preferences",
      "title": "거래 방식을 정해주세요",
      "requiredFieldKeys": ["tradeMethod"],
      "fieldGroups": ["execution_preference", "safety_policy"]
    },
    {
      "stepId": "review_publish",
      "title": "게시 전 확인",
      "requiredFieldKeys": [],
      "fieldGroups": ["publish_gate"]
    }
  ]
}
```

### 67.5 필드별 계약
| fieldKey | 필수 시점 | validation severity 기본안 | 화면 힌트 | DB/API 시사점 |
|---|---|---|---|---|
| `listingType` | `basics` step 완료 전 필수 | `block` | 판매/구매 선택이 전체 카피와 필드 의미를 바꾼다 | write model 필수 enum |
| `serverId` | `basics` step 완료 전 필수 | `block` | 최근 사용 서버 quick pick 제공 | 표준 카탈로그 FK 또는 code |
| `categoryId` | `basics` step 완료 전 필수 | `block` | 카테고리별 attribute template 로드 기준 | template selection key |
| `itemNameRaw` | `basics` step 완료 전 필수 | `block` | 자동완성 + 자유입력 병행 | raw + normalized 후보 저장 |
| `priceType` | `pricing_quantity` step 필수 | `block` | `fixed/negotiable/offer`별 입력 필드 변화 | DTO branching 필요 |
| `priceAmount` | `fixed/negotiable`일 때 필수 | `block` 또는 과도값 `warn` | offer일 때 숨김/null 허용 | money/decimal 정책 필요 |
| `quantity` | `pricing_quantity` step 필수 | `block` | stackable 여부 따라 단위 힌트 제공 | trade unit model 연결 |
| `tradeUnitSnapshot` | category/template 따라 조건부 필수 | `warn` 또는 `block` | 묶음/개당/총액 표시 | unit basis projection 필요 |
| `description` | `details` step 필수 | 길이 부족 `block`, 정보 부족 `warn` | 검색 카드/상세 미리보기 제공 | completeness metric 산정 입력 |
| `enhancementLevel` | 카테고리별 조건부 | 누락 시 `info` 또는 template 기준 `warn` | 장비류에서 강조 | structured attribute 저장 |
| `optionsText` | 선택 | 품질 낮음/과도한 반복문자 `warn` | 옵션 템플릿 chip 보조 가능 | raw text + moderation scan |
| `images[]` | 선택 | 없음 `warn`, 정책 위반 `block/review_required` | 대표 이미지 미리보기 | upload lifecycle 필요 |
| `tradeMethod` | `trade_preferences` step 필수 | `block` | 인게임/오프라인/either에 따라 후속 필드 표시 | execution flow gating |
| `availabilityInput` | 선택이나 권장 | 없음 `warn` | 시간대 미입력 시 응답 저하 경고 | availability snapshot 생성 |
| `meetingPreferenceInput` | `offline` 포함 시 권장 | 모호함 `warn` | 실제 주소 대신 지역 수준 가이드 | meeting point disclosure policy |
| `characterDisclosurePreference` | 거래 방식 따라 선택 | 과노출 위험 `warn` | 공개/예약 후 공개/미공개 중 선택 | identity handshake policy 연동 |

### 67.6 게시 가능성 판정 규칙
| 상태 | 조건 가이드 | 사용자 의미 |
|---|---|---|
| `not_started` | 핵심 step 미진입 | 아직 등록 시작 전 |
| `draft_incomplete` | block 이슈 존재 또는 필수 step 미완료 | 더 입력해야 게시 가능 |
| `ready_with_warnings` | block 없음, warn 존재 | 게시 가능하지만 품질 개선 권장 |
| `ready_to_publish` | block/warn 거의 없음, policy clear | 즉시 공개 가능 |
| `ready_but_review_required` | block 없음, review_required 존재 | 게시 가능하나 검토/제한공개 가능 |
| `publish_blocked` | 정책/권한/제재로 게시 불가 | 수정 또는 지원 필요 |

판정 원칙:
- `warn`는 게시를 막지 않지만 listing quality score와 초기 노출에 반영될 수 있다.
- `review_required`는 사용자 경험상 `게시 실패`가 아니라 `검토 대기` 또는 `제한적 공개` 후보로 안내한다.
- `publish_blocked`는 폼 입력 오류뿐 아니라 계정 restriction/policy gate도 포함한다.

### 67.7 step별 UX 계약
#### 67.7.1 `basics`
- 첫 화면에서 거래 유형(`sell/buy`)과 서버 선택을 빠르게 끝내야 한다.
- `itemNameRaw` 입력 시 자동완성 후보와 최근 입력값을 함께 보여주되, 자유입력을 막지 않는다.
- 이 step 완료 전에는 draft 제목/미리보기 카드 생성이 제한적일 수 있다.

#### 67.7.2 `pricing_quantity`
- `priceType` 선택에 따라 `priceAmount` 입력 UI가 즉시 바뀌어야 한다.
- `offer` 타입이면 가격 필드 대신 협상 안내 카피와 검색 노출 정책 힌트를 제공한다.
- 수량 입력은 단위(`개`, `묶음`, `총액 기준`)와 함께 보여야 한다.

#### 67.7.3 `details`
- 사용자는 description만 쓰는 것이 아니라, 검색/신뢰에 도움이 되는 구조화 힌트를 받아야 한다.
- 이미지가 없어도 게시는 가능하되, `warn`과 검색 클릭률 저하 힌트를 제공한다.
- 정책 위반 감지는 입력 즉시 soft warning, 게시 시점 hard check 두 레이어로 나눈다.

#### 67.7.4 `trade_preferences`
- 거래 방식에 따라 장소/가용시간/캐릭터 공개 설정이 문맥적으로 바뀌어야 한다.
- `offline_pc_bang` 또는 `either` 선택 시 안전 가이드와 공개 범위 경고를 반드시 보여준다.
- availability는 자유입력과 구조화 slot 선택을 함께 둘 수 있으나, 저장 시 canonical snapshot으로 정규화되어야 한다.

#### 67.7.5 `review_publish`
- 사용자는 게시 직전 아래 4가지를 확인해야 한다.
  1. 목록 카드 미리보기
  2. 누락/경고 항목
  3. 정책 검토 가능성
  4. 게시 후 수정 제한 예상 포인트
- 게시 버튼 주변에는 `즉시 공개`, `검토 후 공개`, `게시 불가` 중 하나가 명확히 표기되어야 한다.

### 67.8 draft / autosave 계약
- `basics` step의 필수값 중 2개 이상 입력되면 local draft 생성 가능
- `basics` 완료 시점부터 `server_draft_optional`을 기본으로 하며, 로그인 사용자라면 서버 draft upsert를 권장
- 이미지 업로드/정책 경고/검토 필요 상태가 생기면 `server_draft_required`로 승격 가능
- 최종 publish는 `publish_atomic` 원칙을 따르며, draft와 public listing aggregate를 명시적으로 구분한다

필수 draft 필드 후보:
- `draftId`
- `draftVersion`
- `currentStepId`
- `lastEditedAt`
- `lastAutosavedAt`
- `publishReadinessState`
- `validationIssueSummaryJson`
- `isRecoverable`

### 67.9 validation issue canonical 객체
```json
{
  "fieldKey": "description",
  "validationCode": "DESCRIPTION_TOO_SHORT",
  "severity": "block",
  "stepId": "details",
  "userMessage": "설명을 조금 더 적어주세요.",
  "adminOnlyContext": null,
  "suggestedAction": "expand_description",
  "isDismissible": false
}
```

원칙:
- 동일 이슈 객체를 클라이언트 inline error, publish summary sheet, analytics, 운영 검토 참고에 재사용한다.
- `adminOnlyContext`는 자동 탐지 룰 정보처럼 사용자에게 숨겨야 하는 맥락을 위한 필드다.

### 67.10 API / DB / analytics 파생 기준
#### 67.10.1 API surface 후보
- `GET /listing-form/schema`
- `POST /listing-drafts`
- `PATCH /listing-drafts/{draftId}`
- `POST /listing-drafts/{draftId}/validate`
- `POST /listing-drafts/{draftId}/publish`
- 또는 MVP 단순안으로 `POST /listings/validate`, `POST /listings` 조합을 사용하되, 응답에는 반드시 `publishReadinessState`와 `validationIssues`를 포함한다.

#### 67.10.2 DB/aggregate 후보
- `ListingDraft`
- `ListingDraftStepState`
- `ListingValidationIssue`
- `ListingPublishAttempt`

필수 필드 후보:
- `draftId`
- `authorUserId`
- `schemaVersion`
- `stepStateJson`
- `currentPayloadJson`
- `publishReadinessState`
- `reviewRequiredYn`
- `blockedReasonCode`
- `lastValidatedAt`
- `publishedListingId`

#### 67.10.3 analytics 이벤트 후보
- `listing_form_started`
- `listing_form_step_completed`
- `listing_form_validation_shown`
- `listing_form_warning_dismissed`
- `listing_draft_autosaved`
- `listing_publish_attempted`
- `listing_publish_blocked`
- `listing_publish_review_required`
- `listing_published`

핵심 관찰 지표:
- step별 이탈률
- `warn` 상태에서 게시 완료율
- `review_required` 비율과 실제 승인율
- 이미지 첨부 여부에 따른 채팅 시작 전환율
- `pricing_quantity` step 실패 사유 상위 분포

### 67.11 오픈 질문
- MVP에서 실제로 `ListingDraft` 서버 저장을 포함할지, 로컬 draft만으로 시작할지 구현 범위를 확정해야 한다.
- 카테고리별 attribute template를 등록 초기에 어디까지 강제할지(`warn` vs `block`) 정책 결정이 필요하다.
- 이미지 없는 매물을 기본 허용하되 랭킹 감점으로 처리할지, 특정 카테고리에서는 최소 1장 필수로 둘지 결정해야 한다.
- `review_required` 상태를 `게시 성공 + 검토 대기`로 보일지, `즉시 비공개 저장`로 보일지 출시 범위를 확정해야 한다.


## 59. 화면 명세 패키지(Screen Spec Package) / routeId·moduleId·state contract
### 59.1 목적
- PRD에서 바로 화면설계서, 프론트엔드 backlog, QA 시나리오, 딥링크 문서로 파생할 수 있도록 `화면 단위 canonical package`를 정의한다.
- 현재 문서 곳곳에 흩어진 IA, CTA 규칙, API 의존성, 공통 상태를 `routeId + moduleId + state variant` 관점으로 다시 묶어 실제 구현 단위로 정렬한다.
- 동일 화면을 iOS/Android/Web이 각자 다른 이름과 분해 방식으로 구현하지 않도록 최소 공통 vocabulary를 고정한다.

### 59.2 설계 원칙
1. **화면은 route 기준, 기능은 module 기준으로 분리**한다. 라우트는 진입점이고, 모듈은 재사용 가능한 화면 블록이다.
2. **상태는 화면 전체 상태와 모듈 상태를 분리**한다. 예: 화면은 `ready`지만 추천 슬롯만 `partial`일 수 있다.
3. **액션은 CTA copy가 아니라 command surface로 정의**한다. 버튼 문구는 변해도 command 의미는 고정한다.
4. **모든 핵심 화면은 deep link, analytics, empty/error 복구 경로를 가져야 한다.**
5. **Screen spec은 서버 projection과 1:1로 맞추지 않는다.** 하나의 화면은 여러 projection을 조합할 수 있으나, primary source를 명시해야 한다.

### 59.3 핵심 vocabulary
#### 59.3.1 `screenFamily`
- `home`
- `listing_market`
- `listing_detail`
- `listing_editor`
- `chat_trade`
- `trade_workspace`
- `notification_inbox`
- `profile_trust`
- `admin_ops`

#### 59.3.2 `routeId`
- 앱 라우트/딥링크/푸시가 공유하는 canonical 식별자
- 예시:
  - `home.main`
  - `market.list`
  - `listing.detail`
  - `listing.editor.create`
  - `listing.editor.edit`
  - `chat.thread`
  - `trade.list`
  - `trade.detail`
  - `notifications.list`
  - `profile.user`
  - `admin.report.queue`
  - `admin.report.detail`

#### 59.3.3 `moduleId`
- 화면 내부 재사용 블록 식별자
- 예시:
  - `search_entry_bar`
  - `server_context_chip_row`
  - `listing_card_feed`
  - `listing_summary_header`
  - `trust_summary_block`
  - `availability_slot_block`
  - `trade_action_panel`
  - `message_timeline`
  - `reservation_card`
  - `day_of_trade_card`
  - `report_action_sheet`
  - `admin_case_timeline`

#### 59.3.4 `screenStateVariant`
- `loading_initial`
- `loading_refresh`
- `ready_normal`
- `ready_action_required`
- `empty_cold_start`
- `empty_filtered`
- `partial_degraded`
- `error_retryable`
- `error_terminal`
- `policy_blocked`

#### 59.3.5 `primaryCommandSurface`
- 화면에서 가장 중요한 1차 액션 그룹
- 예시:
  - `search_and_filter`
  - `start_chat`
  - `save_or_publish_listing`
  - `propose_or_confirm_reservation`
  - `mark_arrival_or_delay`
  - `confirm_or_dispute_completion`
  - `report_or_block`
  - `triage_or_restrict`

### 59.4 screen spec 산출물 최소 템플릿
각 화면명세서는 아래 항목을 반드시 포함해야 한다.
1. `screenFamily`
2. `routeId`
3. 화면 목적 / primary user intent
4. 진입 조건 / deep link parameter / fallback route
5. 주요 module 목록과 우선순위
6. `screenStateVariant`별 UI 차이
7. primary read model / primary write command surface
8. 정책 배너 / restriction / disclosure 규칙
9. analytics event mapping
10. QA acceptance checklist

권장 JSON skeleton:
```json
{
  "screenFamily": "listing_detail",
  "routeId": "listing.detail",
  "params": ["listingId"],
  "primaryCommandSurface": ["start_chat", "favorite", "report_or_block"],
  "modules": [
    "listing_summary_header",
    "trust_summary_block",
    "listing_attribute_block",
    "listing_action_footer"
  ],
  "states": [
    "loading_initial",
    "ready_normal",
    "policy_blocked",
    "error_terminal"
  ]
}
```

### 59.5 핵심 사용자 화면 패키지
#### 59.5.1 홈 `routeId=home.main`
- 목적: 탐색 시작 + 거래 재개 + 개인화된 next action 진입
- primary command surface: `search_and_filter`, `resume_trade`, `create_listing`
- 필수 module:
  - `search_entry_bar`
  - `server_context_chip_row`
  - `action_required_trade_strip`
  - `recommended_listing_feed`
  - `recent_activity_summary`
  - `safety_notice_banner`
- primary read sources:
  - Home surface projection
  - unread notification/trade counts
- 핵심 상태:
  - `empty_cold_start`: 신규 사용자 onboarding mission 노출
  - `ready_action_required`: 예약 응답/완료 확인 대기 카드 상단 고정
  - `partial_degraded`: 추천 모듈 실패 시 거래 재개 모듈은 유지
- 핵심 analytics:
  - `home_view`
  - `home_module_impression`
  - `home_resume_trade_click`
  - `home_create_listing_click`

#### 59.5.2 거래소 목록 `routeId=market.list`
- 목적: 조건 기반 후보 좁히기 + 빠른 채팅 진입
- primary command surface: `search_and_filter`, `start_chat`, `favorite`
- 필수 module:
  - `market_filter_header`
  - `facet_sheet`
  - `listing_card_feed`
  - `saved_search_entry`
  - `empty_result_recovery_block`
- primary read sources:
  - `GET /listings`
  - facet/search suggestion projections
- 핵심 상태:
  - `empty_filtered`: 추천 필터 해제 CTA 필수
  - `partial_degraded`: facet count 실패 시 목록은 계속 노출
- 리스트 카드 최소 포함 요소:
  - 가격/거래유형/서버/상태
  - author trust summary minimal
  - available action set

#### 59.5.3 매물 상세 `routeId=listing.detail`
- 목적: 거래 판단 + 즉시 행동 + 신뢰 확인
- primary command surface: `start_chat`, `favorite`, `report_or_block`
- 필수 module:
  - `listing_summary_header`
  - `listing_price_unit_block`
  - `listing_attribute_block`
  - `trust_summary_block`
  - `availability_slot_block`
  - `meeting_preference_block`
  - `listing_action_footer`
- primary read sources:
  - `GET /listings/{listingId}`
  - user public trust projection
- 상태 파생 규칙:
  - `policy_blocked`: 제재/차단/본인 매물 등으로 start_chat 대체 CTA 표시
  - `partial_degraded`: 일부 신뢰 데이터 지연 시 listing core는 계속 노출
  - `error_terminal`: 비공개/삭제/종결 soft landing 포함

#### 59.5.4 매물 등록/수정 `routeId=listing.editor.create|edit`
- 목적: 빠르고 일관된 listing write flow
- primary command surface: `save_draft`, `publish_listing`, `preview_listing`
- 필수 module:
  - `listing_type_step`
  - `server_category_step`
  - `attribute_form_block`
  - `price_quantity_block`
  - `availability_meeting_block`
  - `image_upload_block`
  - `policy_validation_banner`
  - `publish_footer_bar`
- primary read/write sources:
  - listing draft DTO
  - listing write API
  - form validation contract
- 상태 파생 규칙:
  - `policy_blocked`: 금지 입력 존재, 게시 불가
  - `ready_action_required`: 경고성 validation만 남은 게시 가능 상태
  - `partial_degraded`: 이미지 업로드 실패와 텍스트 draft 저장 실패를 분리

#### 59.5.5 채팅 거래 화면 `routeId=chat.thread`
- 목적: 협상, 예약, 당일 실행, 완료/신고까지 하나의 거래 워크스페이스에서 처리
- primary command surface: `send_message`, `propose_or_confirm_reservation`, `mark_arrival_or_delay`, `confirm_or_dispute_completion`
- 필수 module:
  - `counterparty_trade_header`
  - `trade_status_pill_row`
  - `message_timeline`
  - `reservation_card`
  - `day_of_trade_card`
  - `trade_action_panel`
  - `report_action_sheet`
- primary read sources:
  - chat thread projection
  - trade case summary
  - execution readiness snapshot
- 상태 파생 규칙:
  - `ready_action_required`: 예약 응답/당일 확인/완료 확인 대기 시 액션 패널 sticky
  - `policy_blocked`: report_locked, restriction, protection hold
  - `partial_degraded`: 타임라인은 열리지만 일부 첨부/증빙 프리뷰 실패

#### 59.5.6 내 거래 목록/상세 `routeId=trade.list|trade.detail`
- 목적: 사용자의 액션 필요 거래를 우선 정리
- primary command surface: `resume_trade`, `propose_or_confirm_reservation`, `confirm_or_dispute_completion`
- 필수 module:
  - `trade_priority_filter_bar`
  - `trade_thread_summary_list`
  - `execution_readiness_block`
  - `deadline_warning_block`
  - `outcome_history_block`
- primary read sources:
  - `GET /me/trades`
  - `GET /me/trades/{tradeThreadId}`
- 상태 파생 규칙:
  - `ready_action_required`가 기본 우선 surface
  - `empty_cold_start`: 아직 거래 없음 + 첫 문의/첫 매물 미션 노출

#### 59.5.7 알림함 `routeId=notifications.list`
- 목적: 놓친 거래 이벤트 복구
- primary command surface: `open_notification_target`, `mark_read`, `bulk_clear_non_actionable`
- 필수 module:
  - `notification_filter_tabs`
  - `notification_feed_list`
  - `notification_action_chip_row`
- primary read sources:
  - notification inbox projection
- 상태 규칙:
  - action-required 알림과 informational 알림은 시각적으로 분리
  - expired/deep link invalid 상태에 fallback CTA 필수

#### 59.5.8 프로필/신뢰 `routeId=profile.user`
- 목적: 상대의 거래 맥락 판단 + 반복 거래 전환
- primary command surface: `view_reviews`, `start_repeat_trade`, `report_or_block`
- 필수 module:
  - `profile_identity_header`
  - `trust_summary_block`
  - `review_summary_block`
  - `repeat_trade_history_block`
  - `activity_availability_block`
- primary read sources:
  - user profile projection
  - user public trust projection
  - review summary projection

### 59.6 운영 화면 패키지
#### 59.6.1 신고 큐 `routeId=admin.report.queue`
- primary command surface: `triage_or_assign`
- 필수 module:
  - `queue_filter_bar`
  - `queue_sla_summary`
  - `case_list_table`
  - `bulk_action_bar`
- 상태 규칙:
  - P1/P2는 bulk action 비활성 기본
  - SLA breach 위험 건은 고정 상단

#### 59.6.2 신고 상세 `routeId=admin.report.detail`
- primary command surface: `restrict`, `hide_restore`, `request_evidence`, `close_case`
- 필수 module:
  - `case_header_summary`
  - `evidence_bundle_panel`
  - `linked_object_panel`
  - `admin_case_timeline`
  - `decision_form_panel`
  - `audit_meta_footer`
- primary read sources:
  - report detail projection
  - support/moderation case linkage
- 상태 규칙:
  - evidence insufficiency와 imminent risk를 같은 화면에서 구분
  - irreversible action은 approval gate module 필수

### 59.7 모듈 재사용 기준
아래 모듈은 공통 재사용 대상으로 고정한다.

| moduleId | 재사용 화면 | 책임 |
|---|---|---|
| `trust_summary_block` | 목록 카드, 상세, 채팅 상단, 프로필 | 공개 신뢰 요약 렌더링 |
| `reservation_card` | 채팅, 내 거래 상세, 알림 딥링크 상세 | 예약안/확정안 표시와 명령 실행 |
| `day_of_trade_card` | 채팅, 내 거래 상세 | 도착/지연/노쇼/완료 CTA |
| `policy_validation_banner` | 등록, 채팅 입력, 후기 입력 | 정책 차단/경고 설명 |
| `deadline_warning_block` | 내 거래, 채팅, 운영 상세 | 응답 마감/소명 기한 시각화 |
| `report_action_sheet` | 상세, 채팅, 프로필 | 신고/차단/안전 액션 |

원칙:
- 동일 moduleId는 surface별 스타일 차이는 허용하지만, 입력 props와 command semantics는 최대한 공유해야 한다.
- 모듈별 API 호출을 분산시키기보다 상위 route container가 projection을 조합해 주입하는 구조를 우선 권장한다.

### 59.8 screen state와 API/command 연결 규칙
- `loading_initial`: 최소 skeleton 명세 필요, analytics impression 중복 방지 규칙 포함
- `loading_refresh`: 기존 데이터 유지 + 상단/풀다운 refresh affordance
- `ready_action_required`: sticky CTA 또는 primary panel 강조 필수
- `partial_degraded`: 어떤 모듈이 실패했는지 사용자에게 과도하게 노출하지 않되 복구 CTA 제공
- `policy_blocked`: 단순 disable이 아니라 reason code + recovery action 제공
- `error_terminal`: 홈 이동, 뒤로가기, 유사 매물/기록 보기 같은 escape hatch 포함

### 59.9 API 명세 파생 기준
screen spec 문서는 각 route마다 아래를 명시해야 한다.
- primary read endpoint
- secondary read endpoint
- optimistic update 가능 command
- blocking command
- idempotency required command
- stale-while-revalidate 허용 모듈

예시:
- `chat.thread`
  - primary read: `GET /chats/{chatRoomId}` + `GET /chats/{chatRoomId}/messages`
  - blocking command: `POST /reservations/{reservationId}/confirm`, `POST /trade-completions/{completionId}/confirm`
  - optimistic update candidate: `POST /chats/{chatRoomId}/messages`
- `listing.editor.create`
  - blocking command: `POST /listings`
  - optimistic local state: draft autosave
  - async subtask: image upload

### 59.10 analytics / QA 파생 기준
각 route는 최소 3종 이벤트를 가져야 한다.
1. `screen_view`
2. `primary_command_click`
3. `primary_outcome`

예시 이벤트 naming:
- `screen_view_home_main`
- `screen_view_listing_detail`
- `command_click_start_chat`
- `command_outcome_publish_listing_success`
- `command_outcome_confirm_reservation_failure`

QA 체크리스트 최소 기준:
- route 진입 파라미터 누락/만료 케이스
- `screenStateVariant` 전부 재현 가능 여부
- primaryCommandSurface 각 action의 visible/disabled/hidden 조건
- push deep link 진입 후 복귀 경로
- partial_degraded 상태에서 핵심 거래 플로우가 계속 가능한지 여부

### 59.11 오픈 질문
- `trade.list`와 `chat.list`를 MVP에서 완전히 분리할지, 하나의 inbox에서 탭으로 처리할지 확정 필요
- 모바일 하단 탭 IA에서 `내 매물`과 `내 거래`를 독립 탭으로 둘지, `내 활동` 허브로 통합할지 결정 필요
- 프로필 화면에서 repeat trade CTA를 상세보다 앞에 둘지, 후기/신뢰를 먼저 둘지 사용자 테스트 필요
- 운영 백오피스에서 queue/list/detail split view가 필요한지, 모바일 대응 범위를 어디까지 볼지 확정 필요
- design system 차원에서 `moduleId`를 컴포넌트 명과 동일시할지, 상위 개념으로 둘지 합의 필요

## 60. 거래 결과 증빙 패키지(Trade Outcome Evidence Pack) / 완료·노쇼·분쟁 판정 입력 계약
### 60.1 목적
- 완료 확인, 노쇼 claim, 조건 불일치 dispute, protection appeal이 각각 다른 첨부/메모 규칙을 쓰지 않도록 `거래 결과 판정용 증빙 패키지`를 canonical 객체로 정의한다.
- 사용자 제출 UX, 운영 큐, 저장소, API 응답, 다운로드 통제, audit trail이 같은 증빙 단위를 바라보게 해 향후 support/moderation policy 문서로 직접 파생 가능하게 한다.
- 증빙은 `많이 받는 것`보다 `판정 가능한 최소 구성`을 더 중요하게 보며, 과도한 개인정보 수집·노출을 피하는 방향을 기본 원칙으로 둔다.

### 60.2 설계 원칙
1. **사건 중심 제출**: 증빙은 listing/chat/reservation raw attachment가 아니라 `trade outcome case`에 묶여야 한다.
2. **판정 가능성 우선**: 운영자가 어떤 결론을 내리기에 충분한지 `decisionReadinessTier`로 평가한다.
3. **원본과 공개본 분리**: 사용자 상호 노출본, 운영 검토본, export/download본을 분리한다.
4. **최소 노출**: 상대에게는 필요한 설명만, 운영에는 판정 가능한 근거만 노출한다.
5. **연쇄 재사용**: 같은 증빙은 no-show, completion dispute, appeal, restriction review에서 재링크 가능해야 한다.

### 60.3 핵심 vocabulary
#### 60.3.1 `outcomeEvidencePack`
특정 거래 결과 사건을 판정하기 위해 묶인 증빙 패키지.

핵심 필드 후보:
- `evidencePackId`
- `tradeCaseId`
- `sourceCaseType`: `completion_confirmation` | `no_show_claim` | `trade_dispute` | `appeal`
- `submissionWindowState`
- `decisionReadinessTier`
- `evidenceDisclosureScope`
- `primaryNarrative`
- `artifactCount`
- `submittedByPartyRole`
- `counterpartyAccessMode`
- `reviewerSummaryState`
- `asOf`

#### 60.3.2 `evidenceArtifactType`
- `chat_message_reference`
- `reservation_snapshot`
- `arrival_checkin_log`
- `location_proof_text`
- `image_attachment`
- `system_state_snapshot`
- `payment_note`
- `counterparty_identifier_note`
- `operator_requested_statement`
- `appeal_supporting_document`

원칙:
- MVP에서는 외부 문서 업로드보다 플랫폼 내부 기록 reference와 이미지/텍스트 제출을 우선한다.
- `payment_note`는 결제 보증이 아니라 분쟁 맥락 설명용 메모/증빙 참조이며, 민감정보 직접 노출 금지 원칙을 따른다.

#### 60.3.3 `decisionReadinessTier`
- `insufficient`: 현재 정보로 판정 불가
- `basic`: 단일 방향 결론은 가능하나 불확실성 높음
- `substantiated`: 합리적 운영 판정 가능
- `strongly_substantiated`: 여러 독립 증빙이 정합적으로 일치

#### 60.3.4 `evidenceDisclosureScope`
- `operator_only`: 운영만 열람
- `party_summary_only`: 당사자에게는 요약만 노출
- `party_redacted_preview`: 마스킹된 일부 증빙 미리보기 허용
- `export_restricted`: 외부 다운로드 금지, 화면 열람만 허용

#### 60.3.5 `reviewerSummaryState`
- `not_started`
- `needs_more_context`
- `ready_for_decision`
- `decision_recorded`
- `reopened_for_appeal`

### 60.4 canonical 객체 예시
```json
{
  "evidencePackId": "ep_123",
  "tradeCaseId": "case_456",
  "sourceCaseType": "no_show_claim",
  "submissionWindowState": "open",
  "decisionReadinessTier": "basic",
  "evidenceDisclosureScope": "party_summary_only",
  "submittedByPartyRole": "claimant",
  "primaryNarrative": {
    "summary": "상대가 약속 시각 이후 20분 이상 응답하지 않았다고 주장함",
    "narrativeCompleteness": "complete"
  },
  "artifacts": [
    {
      "artifactId": "ea_1",
      "type": "reservation_snapshot",
      "artifactState": "accepted",
      "priority": 1
    },
    {
      "artifactId": "ea_2",
      "type": "arrival_checkin_log",
      "artifactState": "accepted",
      "priority": 2
    }
  ],
  "reviewerSummaryState": "needs_more_context",
  "asOf": "2026-03-14T10:55:00+09:00"
}
```

### 60.5 사건 유형별 최소 증빙 기준
| 사건 유형 | 최소 필수 증빙 | 권장 보강 증빙 | 판정 포인트 |
|---|---|---|---|
| 완료 확인 dispute | 거래 조건 snapshot + completion request log | 채팅 합의 메시지, 인도/수령 확인 로그 | 실제 거래가 성립했는지 |
| no-show claim | confirmed reservation + 당일 check-in/응답 이력 | 도착 메시지, 지연/재일정 요청 이력 | 약속 불이행인지, 상호 오해인지 |
| 조건 불일치 dispute | deal terms snapshot | 변경 제안 메시지, 현장 불일치 이미지/메모 | 누가 어떤 조건을 마지막으로 동의했는지 |
| appeal | 원판정 요약 + appeal statement | 신규 이미지, 누락 로그 reference | 원판정을 뒤집을 새 근거가 있는지 |

원칙:
- 외부 캡처보다 플랫폼 내부 상태 snapshot과 message reference를 우선 증빙으로 간주한다.
- `insufficient` 판정은 증빙이 적어서라기보다, 사건의 핵심 쟁점에 연결되지 않을 때 발생한다.

### 60.6 artifact 상태 수명주기
#### 60.6.1 `artifactState`
- `submitted`
- `accepted`
- `redacted`
- `rejected_irrelevant`
- `rejected_policy_blocked`
- `expired_access`

#### 60.6.2 처리 원칙
- `submitted`: 사용자 업로드/선택 직후
- `accepted`: 사건과 관련 있고 보존 가능
- `redacted`: 관련성은 있으나 민감정보 마스킹 필요
- `rejected_irrelevant`: 사건 쟁점과 무관
- `rejected_policy_blocked`: 과도한 개인정보, 외부 계정정보, 제3자 정보 포함
- `expired_access`: 보존은 유지하되 사용자 다운로드/미리보기 창 만료

### 60.7 사용자 제출 UX 계약
#### 60.7.1 제출 entry point
- `거래 완료 이의제기`
- `노쇼 신고`
- `운영이 추가 증빙 요청`
- `보호조치 appeal`

#### 60.7.2 제출 폼 최소 구성
- 사건 요약 1문장(`primaryNarrative.summary`)
- 시간축 보강 항목(언제, 무엇이, 누구 기준으로 문제였는지)
- 내부 기록 reference 선택(예약, 메시지, 체크인 로그)
- 선택 첨부(이미지/추가 메모)
- 상대에게 공개 가능한지 여부가 아니라, 시스템이 자동 결정한 disclosure summary 미리보기

#### 60.7.3 UX 가드레일
- 무제한 첨부보다 `핵심 3개 증빙 먼저` 구조를 우선한다.
- 전화번호/계좌번호/실명 등 직접 식별 정보가 보이면 전송 전 마스킹 경고를 제공한다.
- 동일 사건에 대한 중복 제출은 새 pack 생성보다 기존 pack append를 기본으로 한다.

### 60.8 운영 판정 UX / 큐 계약
운영 상세 화면은 최소 아래 블록을 가져야 한다.
- `case_outcome_summary`
- `evidence_pack_overview`
- `artifact_relevance_panel`
- `redaction_review_panel`
- `decision_readiness_badge`
- `missing_evidence_request_cta`
- `decision_record_form`

판정 흐름:
1. 사건 요약 확인
2. 핵심 쟁점별 증빙 매핑 확인
3. `decisionReadinessTier` 판정
4. 필요 시 추가 자료 요청
5. 최종 결론 기록
6. 당사자 공개용 summary 생성

### 60.9 당사자 공개/비공개 기준
| 정보 | 제출자 본인 | 상대방 | 운영자 |
|---|---|---|---|
| 원본 이미지 | 가능(정책 범위 내) | 기본 비공개 또는 redacted preview | 가능 |
| 내부 로그 reference id | 가능 | 비공개 | 가능 |
| 운영 판정 요약 | 가능 | 가능 | 가능 |
| 운영 메모/신뢰 리스크 | 비공개 | 비공개 | 가능 |
| 민감정보 마스킹 전 원본 | 제한적 | 비공개 | 권한자만 가능 |

원칙:
- 상대방에게는 `무엇이 제출되었는지`보다 `운영이 어떤 근거군을 검토했는지` 수준의 summary를 우선 제공한다.
- 증빙 공개가 보복/개인정보 노출을 유발할 수 있으면 `party_summary_only`를 기본값으로 한다.

### 60.10 DB / 스토리지 파생 기준
신규/보강 후보 테이블:
- `trade_outcome_evidence_pack`
- `trade_outcome_evidence_artifact`
- `trade_outcome_evidence_link`
- `trade_outcome_decision_summary`

주요 컬럼 후보:
- `artifactType`
- `artifactState`
- `disclosureScope`
- `submittedByUserId`
- `relatedMessageId`
- `relatedReservationId`
- `storageObjectKey`
- `redactionState`
- `relevanceNote`
- `acceptedAt`
- `retentionUntil`

스토리지 원칙:
- 일반 채팅 첨부와 분쟁/판정 증빙 첨부는 접근등급을 분리한다.
- redaction 파생본과 원본을 구분 저장하고, export/download 권한도 별도로 기록한다.

### 60.11 API candidate surface
- `GET /me/trades/{tradeThreadId}/evidence-pack`
- `POST /me/trades/{tradeThreadId}/evidence-pack/artifacts`
- `POST /me/trades/{tradeThreadId}/evidence-pack/submit`
- `GET /admin/trade-cases/{tradeCaseId}/evidence-pack`
- `POST /admin/trade-cases/{tradeCaseId}/evidence-pack/request-more`
- `POST /admin/trade-cases/{tradeCaseId}/decision`

응답에는 최소 아래가 포함되어야 한다.
- `decisionReadinessTier`
- `missingEvidenceHints`
- `viewerDisclosureMode`
- `artifactCountByState`
- `latestDecisionSummary`

### 60.12 analytics / 운영 KPI 파생 기준
이벤트 후보:
- `evidence_pack_started`
- `evidence_artifact_added`
- `evidence_artifact_redacted`
- `evidence_pack_submitted`
- `evidence_more_requested`
- `evidence_decision_recorded`

관찰 KPI 후보:
- 사건 유형별 `substantiated` 도달률
- 추가 증빙 요청 후 제출 완료율
- `insufficient` 종료 비율
- redaction 발생률
- 판정 리드타임 대비 evidence completeness 상관관계

### 60.13 오픈 질문
- MVP에서 이미지 첨부를 증빙 패키지에 바로 포함할지, 텍스트 + 내부 로그 reference만 우선 지원할지 확정 필요
- no-show 사건에서 위치/도착 증빙을 어느 수준까지 허용할지(자기진술 vs 체크인 로그 vs 이미지) 정책 결정 필요
- 사용자 간 redacted preview 허용 범위를 어디까지 둘지, 전면 summary-only로 갈지 결정 필요
- appeal 단계에서 기존 pack append만 허용할지, 별도 appeal pack을 생성할지 확정 필요


## 61. 운영 정책 문서 파생 팩(Operational Policy Document Pack) / policyDocId·policyRuleId·enforcementSurface canonical 계약
### 61.1 목적
- 현재 PRD에는 moderation, restriction, visibility penalty, fraud protection, evidence, appeal 규칙이 풍부하지만, 실제 운영정책 문서로 파생할 때 어떤 규칙이 어떤 문서 단위로 묶여야 하는지가 아직 명시적으로 고정되어 있지 않다.
- 본 섹션은 PRD를 `운영정책서`, `백오피스 처리 가이드`, `사용자 고지 템플릿`, `QA 정책 테스트 세트`, `감사/컴플라이언스 참고문서`로 직접 파생하기 위한 canonical policy packaging 기준을 정의한다.
- 목표는 같은 정책을 문서마다 다르게 해석하는 일을 줄이고, 운영자 액션·사용자 안내·API reason code·analytics·감사로그가 동일한 policy identity를 공유하도록 만드는 것이다.

### 61.2 설계 원칙
1. **정책 식별자 고정**: 운영 정책은 자유 서술 문단이 아니라 `policyDocId`, `policyRuleId`를 가진 식별 가능한 규칙 단위여야 한다.
2. **도메인과 집행 분리**: 같은 도메인 정책이라도 `등록 차단`, `노출 제한`, `경고`, `수동 검토`, `제한`, `복구`는 별도 enforcement layer로 구분한다.
3. **내부/외부 설명 분리**: 내부 운영 판정 근거와 사용자 고지 문구는 동일하지 않다. 정책 원천은 같되 communication template은 audience별로 분리한다.
4. **예외와 재심 포함**: 정책 규칙은 위반 판정만이 아니라 예외 허용 조건, 증빙 부족 시 fallback, appeal/review 가능 여부까지 포함해야 한다.
5. **문서 파생 가능성 우선**: 각 정책 rule은 운영 runbook, FAQ, admin UI, API schema, QA 케이스로 바로 파생 가능한 최소 필드를 가져야 한다.

### 61.3 핵심 vocabulary
#### 61.3.1 `policyDocId`
- `listing_content_policy`
- `chat_safety_policy`
- `trade_execution_policy`
- `review_publication_policy`
- `report_and_restriction_policy`
- `evidence_and_appeal_policy`
- `search_exposure_policy`
- `account_trust_and_verification_policy`

원칙:
- `policyDocId`는 실제 운영 문서 묶음의 최상위 문서 키다.
- 하나의 PRD 섹션이 여러 policyDoc에 재사용될 수 있으나, 정책 문서 파생 시 반드시 대표 문서 키를 가져야 한다.

#### 61.3.2 `policyRuleId`
형식 권장:
`{policyDocId}.{domain}.{short_rule_name}`

예시:
- `listing_content_policy.listing.external_contact_block`
- `chat_safety_policy.chat.harassment_escalation`
- `trade_execution_policy.trade.no_show_claim_window`
- `report_and_restriction_policy.restriction.repeat_abuse_ladder`
- `search_exposure_policy.exposure.shadow_penalty_recovery`

원칙:
- rule id는 enum/API reason code와 직접 1:1 대응될 필요는 없지만, 최소한 역추적 가능해야 한다.
- 하나의 운영 action이 어떤 rule에 근거했는지 감사로그에 남길 수 있어야 한다.

#### 61.3.3 `policyDomain`
- `listing`
- `chat`
- `reservation`
- `execution`
- `completion`
- `review`
- `report`
- `restriction`
- `trust`
- `exposure`
- `evidence`
- `appeal`
- `admin_ops`

#### 61.3.4 `policyAudience`
- `user_public`: 비회원/회원에게 공개되는 정책 설명
- `user_participant`: 거래 당사자에게만 보이는 정책 설명
- `operator_internal`: 운영자 실행 가이드
- `senior_reviewer`: 고권한 검토자 전용 세부 기준
- `system_integrator`: API/배치/정책엔진 구현용 문서
- `qa_internal`: 테스트케이스/검증용 문서

#### 61.3.5 `policyEnforcementMode`
- `inform_only`: 안내만, 기능 제한 없음
- `warn_before_action`: 실행 전 경고 후 수정 유도
- `block_on_write`: 쓰기/전송/등록 차단
- `limit_visibility`: 검색/추천/알림 노출 제한
- `queue_for_review`: 운영 검토 큐 적재
- `temporary_restrict`: 일정 기간 기능 제한
- `hard_restrict`: 강한 제한/정지/영구제한
- `recover_or_restore`: 복구/해제/재활성화 단계

#### 61.3.6 `policyOutcomeClass`
- `allowed`
- `allowed_with_warning`
- `review_required`
- `restricted`
- `restored`
- `rejected`
- `insufficient_basis`
- `escalated`

#### 61.3.7 `enforcementSurface`
- `listing_write`
- `listing_read`
- `chat_send`
- `chat_read`
- `reservation_flow`
- `day_of_trade`
- `completion_flow`
- `review_publication`
- `search_ranking`
- `notification_dispatch`
- `admin_casework`
- `appeal_center`
- `audit_log`

#### 61.3.8 `policyCommunicationTemplate`
- `pre_action_warning`
- `write_block_notice`
- `visibility_limit_notice`
- `restriction_notice`
- `case_update_notice`
- `appeal_result_notice`
- `operator_macro`
- `faq_summary`

### 61.4 운영 정책 문서 canonical 객체
```json
{
  "policyDocId": "chat_safety_policy",
  "policyRuleId": "chat_safety_policy.chat.external_contact_after_reservation",
  "policyDomain": "chat",
  "summary": "예약 확정 전 외부 연락처 교환을 제한하고, 예약 확정 후에도 일부 위험 패턴은 경고 또는 차단한다.",
  "audiences": ["user_public", "operator_internal", "system_integrator"],
  "enforcementSurfaces": ["chat_send", "admin_casework", "audit_log"],
  "enforcementMode": "warn_before_action",
  "outcomeClasses": ["allowed_with_warning", "restricted", "escalated"],
  "relatedReasonCodes": ["SENSITIVE_INFO_BLOCKED", "EXTERNAL_CONTACT_RISK"],
  "evidenceRequirements": ["message_pattern", "conversation_stage", "repeat_violation_count"],
  "appealAllowed": true,
  "ownerTeam": "trust_and_safety",
  "versionTag": "draft"
}
```

### 61.5 운영 문서 파생 단위
| 산출물 | 주 audience | 필수 포함 요소 | 원천 키 |
|---|---|---|---|
| 운영 정책서 | `operator_internal` | policy scope, decision matrix, enforcement ladder, exception, SLA | `policyDocId` |
| 사용자 정책/FAQ | `user_public` | 금지/허용 예시, 자주 묻는 질문, appeal 가능 여부 | `policyDocId` + `policyCommunicationTemplate` |
| 백오피스 액션 가이드 | `operator_internal` | action precondition, evidence check, required note, escalation path | `policyRuleId` |
| API/정책엔진 구현 문서 | `system_integrator` | trigger, input signal, output action, reason code mapping | `policyRuleId` + `enforcementSurface` |
| QA 정책 테스트 세트 | `qa_internal` | pass/fail fixture, boundary case, copy expectation | `policyRuleId` |
| 감사/사후검토 문서 | `senior_reviewer` | change history, override record, false-positive notes | `policyRuleId` |

원칙:
- 정책 문서는 prose 중심이어도 되지만, 원천 추적은 반드시 `policyDocId/policyRuleId`를 기준으로 한다.
- 동일 rule의 사용자용 문구와 운영자용 문구는 별도 template를 갖되 동일 rule id에 매달려야 한다.

### 61.6 주요 policyDoc 묶음 제안
#### 61.6.1 `listing_content_policy`
포함 rule 범위:
- 금지 품목/표현
- 외부 연락처/계좌/링크 차단
- 이미지/OCR 검수
- 등록 품질 미달 경고
- review required 전환 기준

직접 파생 문서:
- 매물 등록 정책 문서
- 등록 차단/수정 유도 카피 표
- 콘텐츠 검수 운영 가이드

#### 61.6.2 `trade_execution_policy`
포함 rule 범위:
- 예약 제안/확정/만료
- 당일 실행 ack/check-in/no-show claim window
- completion/dispute bridge
- cancel intent / reschedule / readiness drift

직접 파생 문서:
- 거래 당일 운영 runbook
- 내 거래/채팅 CTA 정책서
- no-show 판정 운영 기준

#### 61.6.3 `report_and_restriction_policy`
포함 rule 범위:
- 신고 분류 및 SLA
- restriction ladder
- visibility penalty 승격 조건
- probation/lift/recovery
- operator approval chain

직접 파생 문서:
- 제재 정책 문서
- 운영 권한표 / 승인 체인 가이드
- 사용자 제한 고지 템플릿

#### 61.6.4 `evidence_and_appeal_policy`
포함 rule 범위:
- evidence pack 요구 수준
- disclosure scope
- appeal window / submission completeness
- decision revision / restoration

직접 파생 문서:
- 이의제기 센터 운영 매뉴얼
- 증빙 제출 UX 카피
- 복구/유지/부분완화 결과문 템플릿

### 61.7 rule-level 명세 최소 필드
각 `policyRuleId`는 최소 아래 항목으로 내려갈 수 있어야 한다.

| 필드 | 설명 |
|---|---|
| `ruleSummary` | 한 줄 요약 |
| `triggerCondition` | 어떤 입력/상황에서 rule이 평가되는지 |
| `decisionInputs` | 필요한 로그/증빙/상태/카운터 |
| `defaultOutcomeClass` | 기본 판정 |
| `allowedExceptions` | 예외 허용 케이스 |
| `enforcementMode` | 경고/차단/검토/제한 등 |
| `requiredOperatorNote` | 운영자 메모 필수 여부 |
| `userFacingTemplateId` | 사용자 안내 템플릿 키 |
| `appealPolicy` | appeal 허용/기한/추가자료 필요 여부 |
| `analyticsEventHooks` | 정책 이벤트 발행 포인트 |

### 61.8 정책-액션-문구 연결 규칙
정책은 `rule -> action -> message` 3단 연결로 해석되어야 한다.

| 층위 | 예시 | 원칙 |
|---|---|---|
| Rule | `chat.external_contact_after_reservation` | 정책 근거 |
| Action | `warn_message_send`, `block_message_send`, `queue_content_review` | 시스템/운영 실행 |
| Message | `외부 연락처 공유는 이 단계에서 제한돼요` | 사용자/운영 고지 |

원칙:
- 하나의 rule이 상황에 따라 여러 action으로 이어질 수 있다.
- message template는 action이 아니라 rule과 outcome class를 기준으로 선택해야 copy drift가 줄어든다.
- 감사로그에는 최소 `ruleId`, `actionCode`, `outcomeClass`가 함께 남아야 한다.

### 61.9 운영 override / 예외 / 2인 승인 기준
| 정책 상황 | 기본 처리 | override 허용 여부 | 추가 요구 |
|---|---|---|---|
| 명백한 금지 연락처/계좌 노출 | `block_on_write` | 제한적 허용 없음 | 자동 로그 필수 |
| 애매한 사기 의심 표현 | `queue_for_review` | 가능 | senior review 권고 |
| visibility penalty 해제 | `recover_or_restore` | 가능 | 해제 사유 기록 |
| hard restriction 집행 | `hard_restrict` | 가능 | 2인 승인 또는 사후 승인 로그 |
| appeal 승인에 따른 복구 | `restored` | 가능 | recomputation scope 명시 |

원칙:
- override는 rule을 무시하는 행위가 아니라, `overrideReasonCode`를 가진 예외 decision으로 저장돼야 한다.
- override가 발생하면 후속 false-positive 분석과 정책 개선 backlog로 연결할 수 있어야 한다.

### 61.10 API / DB / 감사로그 파생 기준
신규/보강 후보 객체:
- `policy_document`
- `policy_rule`
- `policy_rule_version`
- `policy_enforcement_log`
- `policy_communication_template`
- `policy_override_record`

주요 컬럼 후보:
- `policyDocId`
- `policyRuleId`
- `policyDomain`
- `enforcementMode`
- `outcomeClass`
- `audience`
- `templateId`
- `overrideReasonCode`
- `effectiveFrom`
- `effectiveTo`
- `supersededBy`

API candidate surface:
- `GET /admin/policies`
- `GET /admin/policies/{policyDocId}`
- `GET /admin/policies/rules/{policyRuleId}`
- `GET /admin/policies/templates/{templateId}`
- `POST /admin/policies/enforcement-preview`
- `GET /admin/cases/{caseId}/policy-trace`

원칙:
- MVP에서 full policy CMS를 만들 필요는 없지만, 최소한 감사로그와 운영 액션 결과가 어떤 rule 근거인지 역추적 가능해야 한다.
- `policy-trace`는 운영 디버깅/재심/QA에서 동일한 판단 경로를 재현하는 데 유용하다.

### 61.11 screen / backoffice 파생 기준
운영 화면은 각 사건에 대해 최소 아래 정보를 보여줘야 한다.
- `적용된 policyRuleId 목록`
- `primary outcome class`
- `operator-facing macro copy`
- `user-facing notice preview`
- `appeal eligibility`
- `override history`

사용자 화면/센터는 아래 수준만 노출한다.
- 무엇이 제한/보류/복구되었는지
- 다음에 무엇을 할 수 있는지
- appeal 가능 여부와 기한
- 자세한 내부 리스크/탐지 로직은 비노출

### 61.12 analytics / 운영 KPI 파생 기준
이벤트 후보:
- `policy_rule_evaluated`
- `policy_outcome_applied`
- `policy_override_used`
- `policy_notice_sent`
- `policy_appeal_opened`
- `policy_appeal_resolved`

관찰 KPI 후보:
- rule별 발동 빈도
- outcome class별 사용자 전환 손실률
- override 발생률
- appeal 인용률
- rule별 false-positive 추정치(override/restore 비율 기반)
- 정책 문구 발송 후 self-recovery 비율

### 61.13 오픈 질문
- MVP에서 `policy_rule` 자체를 DB 관리 대상으로 둘지, 코드/문서 기준으로 시작할지 결정 필요
- 사용자용 정책센터를 별도 화면으로 둘지, 각 제한/차단/신고 surface 안에 분산 노출할지 결정 필요
- `visibility penalty`처럼 완전 공개가 어려운 정책을 사용자에게 어느 수준까지 설명할지 최종 정책 확정 필요
- 운영자 macro/template를 제품 내장형으로 둘지 외부 운영문서로 둘지 결정 필요


## 62. 정책 집행 결과 통지(Actionable Policy Notice) / 사용자 액션 카드 / acknowledgment canonical 계약
### 62.1 목적
- 정책 엔진/운영 액션의 결과가 단순 상태 변경으로 끝나지 않고, 사용자에게 `무슨 일이 일어났고`, `지금 무엇을 해야 하며`, `하지 않으면 어떤 결과가 있는지`를 일관되게 전달해야 한다.
- 제재, 검토 대기, 추가 증빙 요청, 복구 완료, 재동의 필요, 노출 제한 같은 사건이 화면마다 제각각 문구/CTA를 갖지 않도록 `정책 집행 결과 통지`를 canonical 객체로 정의한다.
- 앱 surface, 푸시/인박스, 지원센터, 운영 백오피스, 감사로그, QA fixture가 같은 통지 단위를 기준으로 해석하도록 한다.

### 62.2 설계 원칙
1. **행동 가능성 우선**: 통지는 설명만 하는 배너가 아니라, 가능한 다음 행동과 마감 시점이 있는 action card여야 한다.
2. **정책과 카피 분리**: 내부 `policyRuleId`/`policyOutcomeClass`와 사용자용 문구는 분리하되, 서로 역추적 가능해야 한다.
3. **surface 일관성**: 홈/내 거래/내 매물/알림/지원센터가 같은 사건을 서로 다른 심각도나 다른 CTA로 보여주지 않는다.
4. **ack와 action 분리**: 단순 확인(`읽음/이해함`)이 필요한 경우와 실제 조치(`재동의/수정/증빙제출`)가 필요한 경우를 구분한다.
5. **재진입 가능성**: 사용자가 앱을 닫았다가 돌아와도 같은 notice card를 찾고 이어서 처리할 수 있어야 한다.

### 62.3 핵심 vocabulary
#### 62.3.1 `policyNoticeState`
- `queued`: 정책 결과는 확정됐으나 사용자 노출 전 fanout 대기
- `active`: 사용자에게 보여줘야 하는 현재 활성 통지
- `acknowledged`: 사용자가 확인했으나 후속 액션은 남아 있음
- `action_submitted`: 사용자가 요구된 액션을 제출함
- `under_review`: 제출된 액션/증빙을 운영 또는 시스템이 검토 중
- `resolved`: 요구사항이 해소되어 notice가 닫힘
- `expired`: 기한 경과로 자동 종결 또는 더 강한 상태로 승격됨
- `suppressed`: 상위/중복 notice에 흡수되어 독립 노출 안 함

#### 62.3.2 `requiredUserActionType`
- `none`
- `acknowledge_notice`
- `re_accept_policy`
- `submit_evidence`
- `edit_content`
- `confirm_identity_signal`
- `contact_support`
- `wait_for_review`
- `start_appeal`

#### 62.3.3 `ackRequirementMode`
- `none`: 읽음 추적만 하고 명시 확인 필요 없음
- `soft_ack`: 닫기/확인 버튼으로 acknowledgment 남기면 충분
- `hard_ack`: 확인하지 않으면 일부 기능/화면 진입을 막음
- `blocking_ack`: 확인 + 후속 액션 전까지 관련 기능 전체 차단

#### 62.3.4 `recoveryChecklistItem`
- `remove_external_contact`
- `rewrite_listing_for_clarity`
- `upload_additional_evidence`
- `complete_profile_basics`
- `reconfirm_trade_terms`
- `accept_updated_policy`
- `wait_until_restriction_end`
- `appeal_with_new_evidence`

#### 62.3.5 `noticeEscalationMode`
- `none`
- `remind_only`
- `limit_visibility`
- `block_write_action`
- `lock_surface_access`
- `open_support_case`
- `promote_to_moderation_case`

### 62.4 canonical 객체
```json
{
  "policyNoticeId": "pn_123",
  "sourceCaseType": "restriction",
  "sourceCaseId": "rst_123",
  "policyDocId": "safety-policy-v1",
  "policyRuleId": "contact-exchange-001",
  "policyOutcomeClass": "warn_and_recover",
  "policyNoticeState": "active",
  "audienceUserId": "user_123",
  "surfaceScope": ["home", "listing_editor", "notifications", "support_center"],
  "severity": "high",
  "headline": "외부 연락처는 매물 본문에 직접 적을 수 없어요",
  "body": "본문의 연락처를 제거하면 다시 게시할 수 있어요.",
  "requiredUserActionType": "edit_content",
  "ackRequirementMode": "hard_ack",
  "primaryCta": {
    "actionCode": "edit_listing",
    "label": "매물 수정하기"
  },
  "secondaryCta": {
    "actionCode": "open_policy_help",
    "label": "정책 보기"
  },
  "recoveryChecklist": [
    {
      "item": "remove_external_contact",
      "status": "pending"
    }
  ],
  "expiresAt": "2026-03-16T11:00:00+09:00",
  "escalationModeOnExpiry": "block_write_action",
  "acknowledgedAt": null,
  "resolvedAt": null,
  "createdAt": "2026-03-14T11:02:00+09:00"
}
```

### 62.5 notice 유형 family
| notice family | 대표 상황 | 기본 action | 기본 종료 조건 | 운영 연결 |
|---|---|---|---|---|
| `policy_reaccept` | 약관/정책 업데이트 | 재동의 | 재동의 완료 | account/policy gate |
| `content_fix_required` | 매물/후기/프로필 수정 필요 | 수정 제출 | 검수 통과 | content moderation case |
| `evidence_request` | 추가 증빙 요청 | 증빙 제출 | review 완료 | dispute/no-show/support |
| `restriction_notice` | 기능 제한/제재 통지 | 확인 또는 대기 | 제한 종료/해제 | restriction case |
| `visibility_notice` | 노출 제한/검토 대기 | 확인/수정 | 해제 또는 만료 | exposure decision |
| `trust_recovery_notice` | 프로필/거래 신뢰 회복 안내 | 체크리스트 수행 | 회복 조건 충족 | trust/protection |
| `appeal_result_notice` | 이의제기 결과 | 확인/후속 없음 | ack 또는 만료 | appeal/support |
| `execution_safety_notice` | 거래 당일 주의/재확인 | 재확인/안전행동 | 거래 종료 | trade execution |

### 62.6 사용자 surface별 노출 계약
#### 62.6.1 홈
- 글로벌 계정/정책 notice 1개 + 거래 관련 notice 1개까지 요약 카드로 노출
- `blocking_ack` 또는 `block_write_action` notice는 홈 진입 직후 가장 높은 우선순위로 노출
- 홈 카드는 항상 `왜`, `해야 할 일`, `기한` 3요소를 포함해야 한다.

#### 62.6.2 매물 등록/수정
- `content_fix_required`, `policy_reaccept`, `visibility_notice` 중심으로 노출
- 폼 상단 배너와 필드 인라인 오류/힌트를 연결해야 하며, notice와 field-level validation이 서로 모순되면 안 된다.
- 사용자가 수정 제출 후 notice는 `action_submitted` 또는 `under_review`로 전환된다.

#### 62.6.3 채팅/내 거래
- `execution_safety_notice`, `evidence_request`, `restriction_notice`를 거래 맥락 안에서 표시
- 거래 당일 notice는 day-of-trade CTA보다 위에 오지 않되, 노쇼/분쟁/보호조치와 직접 관련되면 sticky card로 승격 가능
- 추가 증빙 제출 notice는 해당 case 상세 deep link를 반드시 제공해야 한다.

#### 62.6.4 알림함 / 푸시
- notice는 일반 notification과 별도 grouping key를 가져야 한다.
- 푸시 카피는 민감 사유를 축약하고, 앱 진입 후 full card에서 상세 설명/CTA를 보여준다.
- 같은 notice family는 최신 상태 기준 1개로 병합 가능하나, 다른 sourceCaseId는 합치지 않는다.

#### 62.6.5 지원센터 / 내 문의
- 사용자는 활성/종결 notice를 사건 타임라인과 함께 다시 확인 가능해야 한다.
- `정책 보기`, `이의제기 가능 여부`, `제출한 증빙`, `운영 응답`이 하나의 흐름으로 이어져야 한다.

### 62.7 우선순위 / 병합 / 억제 규칙
1. `blocking_ack` > `hard_ack` > `soft_ack` > `none` 순으로 우선한다.
2. 동일 `sourceCaseId + policyRuleId + requiredUserActionType` notice는 중복 생성하지 않고 최신 revision으로 upsert한다.
3. 상위 case가 닫히면 하위 notice도 함께 `resolved` 또는 `suppressed` 처리한다.
4. 동일 사용자의 활성 notice가 과도하게 많을 경우, 홈에서는 최대 2개만 강조하고 나머지는 알림함/지원센터로 보낸다.
5. 거래 당일 notice는 deadline/action notice보다 아래로 밀리지 않도록 별도 우선순위 lane을 가진다.

### 62.8 상태 전이 규칙
- `queued -> active`: fanout 성공 후 사용자 surface에 노출 가능해짐
- `active -> acknowledged`: 사용자가 확인 버튼 또는 진입 기반 soft ack 수행
- `active/acknowledged -> action_submitted`: required action 제출
- `action_submitted -> under_review`: 운영/시스템 검토 필요 시
- `under_review -> resolved`: 승인/해제/통과
- `active/acknowledged -> expired`: 기한 경과 + recovery 미수행
- `expired -> active`: 더 강한 제재/상위 notice로 재발행되는 경우 새 noticeId 권장

원칙:
- `resolved`된 notice는 immutable history로 남긴다.
- 사용자에게는 최신 상태만 강하게 보이되, 운영/지원에서는 revision history를 추적 가능해야 한다.

### 62.9 API 후보 surface
- `GET /me/policy-notices`
- `GET /me/policy-notices/{policyNoticeId}`
- `POST /me/policy-notices/{policyNoticeId}/acknowledge`
- `POST /me/policy-notices/{policyNoticeId}/actions/reaccept-policy`
- `POST /me/policy-notices/{policyNoticeId}/actions/submit-evidence`
- `POST /me/policy-notices/{policyNoticeId}/actions/request-appeal`
- `GET /admin/policy-notices`
- `POST /admin/policy-notices/{policyNoticeId}/resolve`
- `POST /admin/policy-notices/{policyNoticeId}/reissue`

응답 최소 필드:
- `policyNoticeId`, `sourceCaseType`, `sourceCaseId`
- `policyNoticeState`, `severity`, `headline`, `body`
- `requiredUserActionType`, `ackRequirementMode`
- `primaryCta`, `secondaryCta`, `recoveryChecklist`
- `expiresAt`, `availableActions`, `policyHints`

### 62.10 DB / projection / audit 파생 기준
write 모델 후보:
- `PolicyNotice`
- `PolicyNoticeRevision`
- `PolicyNoticeAck`
- `PolicyNoticeActionSubmission`

projection 후보:
- `ActivePolicyNoticeProjection`
- `PolicyNoticeHistoryProjection`
- `PolicySurfaceBadgeProjection`

감사로그 필수 필드:
- `policyNoticeId`
- `sourceCaseId`
- `policyRuleId`
- `previousState`, `newState`
- `actorType` (`user` / `staff` / `system`)
- `triggerActionCode`
- `templateVersion`

### 62.11 운영 runbook / QA 파생 기준
운영은 아래 질문에 즉시 답할 수 있어야 한다.
- 사용자는 어떤 통지를 언제 받았는가?
- 어떤 CTA를 눌렀고, 어떤 요구사항이 아직 남았는가?
- 기한이 지나면 어떤 escalation이 예정되어 있는가?
- 동일 사안 notice가 중복/누락 없이 묶였는가?

must-pass QA 시나리오 후보:
1. 매물 본문 연락처 탐지 → `content_fix_required` notice 생성 → 수정 제출 → resolved
2. 정책 재동의 필요 → blocking gate → 재동의 완료 후 write action 해제
3. 분쟁 추가 증빙 요청 → push/inbox/support center 동시 일관 노출 → 제출 후 under_review 전환
4. 제한 종료 notice → 자동 resolved 또는 recovery notice 전환
5. appeal 결과 통지 → user ack 후 history 보존

### 62.12 analytics / KPI 파생 기준
이벤트 후보:
- `policy_notice_created`
- `policy_notice_impression`
- `policy_notice_acknowledged`
- `policy_notice_primary_cta_clicked`
- `policy_notice_action_submitted`
- `policy_notice_resolved`
- `policy_notice_expired`

관찰 KPI 후보:
- notice family별 ack률
- required action completion rate
- expiry 후 escalation 비율
- notice 발송 후 self-recovery율
- notice-to-support deflection rate
- policy template별 confusion rate(지원 문의/재오픈율 기반)

### 62.13 오픈 질문
- MVP에서 `정책센터(policy center)`를 독립 탭으로 둘지, 홈/알림/지원센터 안에 분산 노출할지 결정 필요
- `soft_ack`를 화면 진입만으로 처리할지, 명시 버튼 클릭만 인정할지 surface별 세분화 필요
- notice template의 다국어/카피 버전 관리를 코드 기반으로 시작할지 DB 기반으로 시작할지 결정 필요
- 특정 notice를 푸시 없이 인앱 전용으로 둘 기준(민감도/소음/오해 가능성) 최종 확정 필요
