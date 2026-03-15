# PRD Chat Sync Notes

이 문서는 PRD.md의 채팅 동기화/읽음/전송 보장 보강 작업을 위한 보조 메모다.
최종 기준은 항상 ../PRD.md를 따른다.

핵심 포인트:
- 채팅은 단순 message list가 아니라 event timeline이며, 모바일 재접속/오프라인 복구를 고려한 cursor 기반 동기화가 필요하다.
- unread 계산은 메시지 개수보다 `마지막 읽은 이벤트` 기준이 일관적이다.
- 실시간 수신 실패 시에도 polling/backfill로 복구 가능해야 한다.
- 네트워크 재시도 시 message dedup 키가 필요하다.
