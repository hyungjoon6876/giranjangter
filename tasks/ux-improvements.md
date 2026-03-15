# 기란장터 UX Improvement Areas

Research-backed UX improvements focused on user flow efficiency, interaction patterns, and usability -- NOT visual design (colors, fonts, spacing).

---

## 1. Search & Discovery

### 1.1 Inline Search (Critical)
**Current**: Search opens a modal dialog, requiring multiple taps and losing browsing context.
**Recommendation**: Replace the dialog with an inline search bar at the top of the listing list. The search field should be always visible (or expandable in the AppBar) and filter results in real-time with debounced API calls (300ms). This matches the pattern used by 아이템매니아 and 아이템베이 where search is the primary discovery path.

### 1.2 Search Suggestions & Recent Searches
**Current**: No search history or suggestions.
**Recommendation**: Store recent search queries locally (SharedPreferences). Show them as chips below the search field. For future: implement server-side popular search terms and autocomplete on item names.

### 1.3 Category Filtering
**Current**: Only server filter chips exist. No category filtering on the listing list.
**Recommendation**: Add a second row of filter chips for categories, or a combined filter panel. Game item trading platforms typically let users filter by: item category (weapon, armor, accessory, etc.), price range, listing type (sell/buy), and enhancement level. Consider a filter sheet that slides up on mobile.

### 1.4 Sort Controls
**Current**: No visible sort options (API supports `sort` parameter but UI doesn't expose it).
**Recommendation**: Add a sort dropdown or chip row: "최신순" (recent), "가격 낮은순" (price low), "가격 높은순" (price high), "인기순" (popular/views). This is standard in every marketplace platform.

### 1.5 Active Filter Visibility
**Current**: Server filter is visible, but if a search query is active there's no visual indicator.
**Recommendation**: Show active filters as removable pills/chips above the listing list. Display result count ("23개의 매물"). Let users clear individual filters or "전체 초기화" to reset all.

---

## 2. Listing List (Information Architecture)

### 2.1 Pagination / Infinite Scroll
**Current**: All listings load at once with no pagination. The API supports cursor-based pagination but the UI doesn't use it.
**Recommendation**: Implement "Load More" button pattern (not infinite scroll -- better for backtracking and position awareness in a trading context). Show the cursor-based pagination the API already supports. Display a count like "20개 매물 / 더보기" at the bottom.

### 2.2 Empty State with Context
**Current**: Empty state says "매물이 없습니다 / 첫 매물을 등록해보세요!" regardless of context.
**Recommendation**: Make empty states contextual:
- No results for search: "'{query}'에 대한 매물이 없습니다. 다른 검색어를 시도해보세요."
- No results for filter: "선택한 서버에 매물이 없습니다. 다른 서버를 선택해보세요."
- Truly empty marketplace: "아직 매물이 없습니다. 첫 매물을 등록해보세요!"
- Optionally offer "알림 설정" to be notified when matching listings appear.

### 2.3 Listing Card - Relative Time Precision
**Current**: Uses `formatTimeAgo` but the display is pushed to a corner.
**Recommendation**: In game item trading, freshness is critical. Make the timestamp more prominent. Consider highlighting listings posted within the last hour with a "NEW" badge. Show "방금 전", "5분 전", "1시간 전" clearly.

### 2.4 Grid vs List View Toggle
**Current**: List view only.
**Recommendation**: Offer a grid/list toggle. Grid view is useful on desktop (wider viewport). List view is better for mobile. Default should be responsive: grid on desktop, list on mobile.

---

## 3. Responsive Layout (PC vs Mobile Web)

### 3.1 Desktop Layout Optimization
**Current**: Single-column layout designed for mobile. On a desktop browser, content stretches awkwardly or is too narrow.
**Recommendation**: Implement responsive breakpoints:
- **< 600px (mobile)**: Current single-column layout, bottom navigation bar
- **600-1024px (tablet)**: 2-column grid for listings, side panel for filters
- **> 1024px (desktop)**: 3-column grid for listings, persistent sidebar filters, replace bottom nav with side rail or top navigation, max content width ~1200px

### 3.2 Navigation Pattern by Viewport
**Current**: Bottom NavigationBar always shown (mobile pattern).
**Recommendation**: On desktop viewports, switch to a NavigationRail (side rail) or top app bar with tabs. Bottom navigation on desktop wastes the most valuable screen real estate. Flutter's `LayoutBuilder` or `MediaQuery` can drive this.

### 3.3 Detail Page Layout
**Current**: Single-column scroll on detail page.
**Recommendation**: On desktop, use a two-column layout: left column for item info/description, right column for seller info and action buttons (sticky). This reduces scrolling and keeps the CTA always visible.

### 3.4 Chat Layout on Desktop
**Current**: Full-screen chat detail.
**Recommendation**: On desktop, use a master-detail pattern: chat list on the left (1/3 width), chat messages on the right (2/3 width). This is the standard pattern for messaging on wider viewports (similar to Telegram Web, Slack).

---

## 4. Loading, Error & Empty States

### 4.1 Skeleton Screens
**Current**: `CircularProgressIndicator` spinner for all loading states.
**Recommendation**: Replace spinners with skeleton screens that match the layout of the content being loaded. For the listing list, show 3-5 skeleton cards with shimmer animation. For detail pages, show skeleton blocks matching the header, info card, and description card shapes. Research shows skeleton screens reduce perceived loading time by ~20% vs spinners.

### 4.2 Error State Recovery
**Current**: Most error handlers silently swallow errors (`catch (e) { setState(() => _loading = false); }`). No error UI in listing list, listing detail, or chat screens.
**Recommendation**: Every data-fetching screen needs an explicit error state with:
- Clear error message (not raw exception text)
- "다시 시도" (retry) button
- Option to go back
The notifications screen already does this well -- replicate that pattern everywhere.

### 4.3 Optimistic Updates
**Current**: Favorite toggle reloads the entire listing detail page.
**Recommendation**: Implement optimistic UI updates for:
- **Favorite toggle**: Immediately update the icon/count, revert on failure
- **Send message**: Add message to list immediately, show "sending..." indicator, confirm or show retry
- **Mark notifications read**: Update UI immediately, fire API in background

### 4.4 Network State Awareness
**Current**: No handling for offline/poor connectivity.
**Recommendation**: Show a banner when the device is offline or SSE connection drops. Queue outgoing actions (messages, favorites) and retry when reconnected. This is especially important for Flutter Web where users might have intermittent connectivity.

---

## 5. Chat UX

### 5.1 Real-time Message Delivery
**Current**: Messages are loaded once via `_loadMessages()` in `initState`. No real-time updates. Users must manually refresh or re-enter the screen to see new messages.
**Recommendation**: Connect to the SSE event stream (`event.Broker` already exists on the backend). Listen for `new_message` events and append messages in real-time. Show incoming messages immediately without manual refresh.

### 5.2 Listing Context in Chat
**Current**: Chat detail screen shows no context about which listing the conversation is about.
**Recommendation**: Add a compact listing card/banner at the top of the chat (item name, price, status, small icon). This is critical for game item trading where a user may have multiple concurrent negotiations. Link it so users can tap to view the full listing. Korean trading platforms like 아이템매니아 always show the item being discussed.

### 5.3 Message Timestamps
**Current**: Messages show no timestamps at all.
**Recommendation**: Show timestamps for messages. Group messages by date with date separators ("오늘", "어제", "2024.03.15"). Show time on each message or at intervals (every 5-minute gap). This is essential for negotiation history tracking.

### 5.4 Typing Indicators & Read Receipts
**Current**: No typing indicators or read receipts.
**Recommendation**: Future enhancement. For now, at minimum show "마지막 접속: X분 전" for the counterparty so the user knows if they can expect a quick response. The SSE broker's `IsOnline` method can power a simple online/offline indicator.

### 5.5 Chat List - Last Message Preview & Time
**Current**: Chat list shows counterparty name and listing title, but no last message or timestamp.
**Recommendation**: Show the last message preview (truncated to 1 line) and relative timestamp ("3분 전"). Show unread message count badge. Sort by most recent message. This is the standard for every messaging interface.

### 5.6 Message Input Improvements
**Current**: Basic text input with send button.
**Recommendation**:
- Enter key to send (already implemented via `onSubmitted`), but also support Shift+Enter for newline on desktop
- Show character count or limit for very long messages
- Disable send button when input is empty (visual feedback)
- Auto-focus the input field when entering chat

---

## 6. Form Usability (Listing Creation)

### 6.1 Form Validation Timing
**Current**: Validation only runs on submit (`_formKey.currentState!.validate()`).
**Recommendation**: Implement `AutovalidateMode.onUserInteraction` so fields validate as users type/leave them. This provides immediate feedback and prevents the frustration of filling a long form only to see errors at the end.

### 6.2 Field Ordering & Grouping
**Current**: All fields are in a single flat list.
**Recommendation**: Group related fields visually with section headers:
- **기본 정보**: Type toggle, server, category, item name
- **상세 정보**: Title, description, enhancement level
- **거래 조건**: Price type, price amount, trade method
This reduces cognitive load. Consider a multi-step wizard for mobile (3 steps matching the groups above) with a progress indicator.

### 6.3 Price Input UX
**Current**: Raw number input with no formatting.
**Recommendation**:
- Show formatted price as user types (10,000 -> 10,000원)
- Add quick-set buttons for common amounts (만 / 10만 / 100만)
- When "제안받음" (offer) is selected, clearly disable the price field with an explanation
- Show the price in a larger, prominent style as the user types it

### 6.4 Draft Saving
**Current**: Leaving the form loses all entered data.
**Recommendation**: Auto-save form state to local storage. When the user returns to the create listing screen, offer to restore the draft. This is important because form abandonment is high when users accidentally navigate away.

### 6.5 Success & Next Steps
**Current**: Shows a snackbar "매물이 등록되었습니다!" and pops back.
**Recommendation**: After successful creation, show a brief success screen or dialog with options:
- "매물 보기" (view my listing)
- "또 등록하기" (create another)
- "목록으로" (go to list)
This keeps the user in flow and encourages continued engagement.

### 6.6 Image Upload Support
**Current**: No image upload capability.
**Recommendation**: Allow users to attach item screenshots. In game item trading, images are crucial for verifying item attributes, enhancement level, and options. Even one image significantly increases trust and conversion.

---

## 7. Navigation & Information Architecture

### 7.1 Deep Linking & URL Structure
**Current**: Routes exist (`/listings/:id`, `/chats/:chatId`) but no mechanism for sharing or bookmarking.
**Recommendation**: Ensure URLs are meaningful and shareable. When a user copies a listing URL, it should work when pasted in another browser/tab. This is a key advantage of Flutter Web over native apps.

### 7.2 Back Navigation Preservation
**Current**: Navigating to a detail page and pressing back reloads the listing list (and loses scroll position).
**Recommendation**: Preserve the listing list scroll position and filter state when returning from a detail page. The `StatefulShellRoute` should help, but verify it works with filter/search state. Consider caching the last listing list response.

### 7.3 Notification Deep Links
**Current**: Notifications screen shows items but doesn't navigate anywhere on tap.
**Recommendation**: Each notification should navigate to the relevant screen when tapped:
- Chat notification -> Chat detail
- Reservation notification -> Chat detail with reservation info
- Trade notification -> Trade detail
- Review notification -> Review detail

### 7.4 Tab State Preservation
**Current**: `StatefulShellRoute.indexedStack` is used (good -- preserves tab state).
**Recommendation**: Verify that when switching between tabs and back, the previous tab's scroll position and data are preserved. Also verify that the chat list refreshes when switching TO the chat tab (to show new messages received while browsing listings).

---

## 8. Keyboard & Accessibility

### 8.1 Focus Management
**Current**: `GestureDetector` used for filter chips and type toggle -- these are not keyboard-accessible.
**Recommendation**: Replace `GestureDetector` with proper Material widgets (`InkWell`, `ChoiceChip`, `FilterChip`, `ToggleButtons`) that handle keyboard focus, hover states, and screen reader semantics. Every interactive element must be reachable and operable via Tab/Enter/Space keys.

### 8.2 Semantic Labels
**Current**: Icons and badges have no semantic labels for screen readers.
**Recommendation**: Add `Semantics` widgets or `tooltip`/`semanticLabel` properties to:
- Filter chips (e.g., "서버: 바츠, 선택됨")
- Status badges (e.g., "상태: 판매중")
- Meta chips (e.g., "조회수 42, 찜 3, 채팅 1")
- Icon buttons (e.g., "검색", "신고하기")
- The favorite button (e.g., "찜하기" or "찜 취소")

### 8.3 Keyboard Shortcuts (Desktop)
**Current**: No keyboard shortcuts.
**Recommendation**: For desktop web users, add keyboard shortcuts:
- `/` or `Ctrl+K`: Focus search
- `N`: New listing (when on listing list)
- `Esc`: Close modals/sheets
- Arrow keys: Navigate between listings in the list
These are standard patterns for web applications and significantly improve power user efficiency.

### 8.4 Tab Order
**Current**: Not explicitly managed.
**Recommendation**: Ensure logical tab order on all screens. Form fields should be tabbable in visual order. Modal/sheet content should trap focus within the modal. When a modal closes, focus should return to the element that triggered it.

---

## 9. Performance Perception

### 9.1 Flutter Web Initial Load
**Current**: Flutter Web's inherent initial load time (downloading/initializing the rendering engine).
**Recommendation**:
- Add a simple HTML loading indicator in `index.html` that shows before Flutter initializes
- Use tree-shaking and deferred loading for non-critical features
- Enable WASM compilation for better performance (Flutter 3.11+ supports this)
- Consider code splitting to load the listing list screen first, then lazy-load chat/profile

### 9.2 Image Loading
**Current**: Item icons use `Image.network` with no caching or placeholder strategy. URLs are hardcoded to `localhost:8080`.
**Recommendation**:
- Use `CachedNetworkImage` (or Flutter's built-in `fadeInDuration`) for icon loading
- Show a placeholder shimmer while images load
- Use proper base URL from config (not hardcoded localhost)
- Consider using WebP format and responsive image sizes

### 9.3 Data Caching
**Current**: Every screen fetches fresh data on every visit.
**Recommendation**: Cache API responses locally (in-memory or SharedPreferences):
- Server list and category list: Cache for 1 hour (rarely changes)
- Listing list: Cache for 30 seconds (show cached, refresh in background)
- User profile: Cache until explicit refresh
Use stale-while-revalidate pattern: show cached data immediately, then update when fresh data arrives.

---

## 10. Real-Time & Notification UX

### 10.1 SSE Connection Management
**Current**: SSE broker exists on the backend but the frontend doesn't connect to it.
**Recommendation**: Establish an SSE connection when the user is authenticated. Handle:
- Automatic reconnection with exponential backoff
- Connection status indicator (subtle, e.g., a small dot in the nav bar)
- Queue events received while the app was backgrounded

### 10.2 In-App Notification Badges
**Current**: No unread count badges on the navigation bar.
**Recommendation**: Show badge counts on navigation items:
- Chat tab: Unread message count
- Profile/notifications: Unread notification count
These should update in real-time via SSE events. Use `Badge` widget on `NavigationDestination` icons.

### 10.3 Push Notification Strategy (Web)
**Current**: No web push notifications.
**Recommendation**: Implement Web Push API for notifying users when the tab is not active:
- New chat messages
- Reservation proposals/responses
- Trade completion prompts
- Request permission at the right moment (after first successful trade, not on first visit)

### 10.4 Live Listing Updates
**Current**: Listing list is static after initial load.
**Recommendation**: Use SSE to push updates when:
- A new listing matching the user's current filter is posted
- A listing the user is viewing changes status (available -> reserved)
- Show a non-intrusive "새로운 매물이 있습니다" banner that the user can tap to refresh

---

## 11. Trust & Safety UX

### 11.1 Seller Reputation Visibility
**Current**: Author card in listing detail shows nickname and trade count. Trust badge appears at 5+ trades.
**Recommendation**: Make trust signals more prominent and graduated:
- Show star rating from reviews (not just count)
- Show member-since date
- Show response time average ("보통 10분 내 응답")
- Make the trust badge tappable to see full trade history/reviews

### 11.2 Transaction Safety Flow
**Current**: The reservation -> trade completion -> review flow exists but may not be obvious to users.
**Recommendation**: Add a progress tracker in the chat/listing detail showing where in the transaction flow the user is:
1. 채팅 중 (Chatting)
2. 예약 제안 (Reservation proposed)
3. 예약 확정 (Reservation confirmed)
4. 거래 완료 (Trade completed)
5. 후기 작성 (Review written)
This guides users through the expected flow and builds confidence.

---

## 12. Flutter Web-Specific Issues

### 12.1 Text Selection
**Current**: Flutter Web renders with CanvasKit/WASM, which means standard browser text selection may not work.
**Recommendation**: Wrap key text content (listing descriptions, chat messages) with `SelectableText` or `SelectionArea` so users can copy text. This is critical for sharing item details, prices, and in-game meeting coordinates.

### 12.2 Browser Back/Forward
**Current**: GoRouter handles routing, but browser back/forward behavior needs verification.
**Recommendation**: Test and ensure that browser back/forward buttons work correctly through the entire navigation flow. Users on desktop web expect browser navigation to work. Ensure filter/search state is preserved in URL parameters for proper back/forward behavior.

### 12.3 Right-Click Context Menu
**Current**: Flutter Web disables the default browser context menu.
**Recommendation**: For key links (listing URLs, user profiles), consider adding `BrowserContextMenu.enable()` or implementing a custom context menu that includes "새 탭에서 열기" (open in new tab) and "링크 복사" (copy link).

### 12.4 Input Focus Issues
**Current**: Flutter Web has known issues with TextField losing focus when scrolling (flutter/flutter#122501).
**Recommendation**: Test form inputs (especially the listing create form) on mobile browsers where the keyboard can cause scroll/focus issues. Apply workarounds like `SingleChildScrollView` with proper `Scrollable.ensureVisible` calls.

---

## Priority Ranking

### P0 (Must-Have for MVP quality)
1. Error states on all screens (4.2)
2. Skeleton loading screens (4.1)
3. Real-time chat messages via SSE (5.1)
4. Listing context in chat (5.2)
5. Chat list - last message & time (5.5)
6. Inline search (1.1)
7. Active filter visibility (1.5)
8. Sort controls (1.4)
9. Notification deep links (7.3)
10. Unread badges on nav bar (10.2)

### P1 (Important for good UX)
11. Responsive desktop layout (3.1, 3.2, 3.3)
12. Message timestamps (5.3)
13. Pagination / Load More (2.1)
14. Category filtering (1.3)
15. Optimistic updates (4.3)
16. Form validation timing (6.1)
17. Focus management / keyboard accessibility (8.1)
18. Back navigation preservation (7.2)
19. Contextual empty states (2.2)
20. Image caching & proper base URL (9.2)

### P2 (Nice-to-Have for polish)
21. Desktop chat master-detail layout (3.4)
22. Form field grouping (6.2)
23. Price input UX (6.3)
24. Grid/list toggle (2.4)
25. Search history (1.2)
26. Draft saving (6.4)
27. Keyboard shortcuts (8.3)
28. Semantic labels (8.2)
29. SSE connection management (10.1)
30. Text selection support (12.1)
31. Seller reputation enhancement (11.1)
32. Transaction progress tracker (11.2)
33. Network state awareness (4.4)
34. Web push notifications (10.3)
35. Image upload (6.6)
36. Data caching strategy (9.3)
37. Flutter Web initial load optimization (9.1)
38. Browser back/forward verification (12.2)

---

## Sources

- [Rigby: Marketplace UX Design Feature-by-Feature Guide](https://www.rigbyjs.com/blog/marketplace-ux)
- [Excited Agency: 9 Marketplace UX Design Best Practices](https://excited.agency/blog/marketplace-ux-design)
- [LogRocket: Filtering UX/UI Design Patterns](https://blog.logrocket.com/ux-design/filtering-ux-ui-design-patterns-best-practices/)
- [Pencil & Paper: Filter UX Design Patterns](https://www.pencilandpaper.io/articles/ux-pattern-analysis-enterprise-filtering)
- [Algolia: Best Marketplace UX for Search](https://www.algolia.com/blog/ecommerce/best-marketplace-ux-practices-for-search)
- [NN/g: Skeleton Screens 101](https://www.nngroup.com/articles/skeleton-screens/)
- [NN/g: Infinite Scrolling Tips](https://www.nngroup.com/articles/infinite-scrolling-tips/)
- [Smashing Magazine: Pagination vs Infinite Scrolling](https://www.smashingmagazine.com/2016/03/pagination-infinite-scrolling-load-more-buttons/)
- [Milan Meurrens: Flutter for Web 2025 Guide](https://www.milanmeurrens.com/guides/when-to-use-flutter-for-web-in-2025-a-comprehensive-guide)
- [Flutter Docs: Web Accessibility](https://docs.flutter.dev/ui/accessibility/web-accessibility)
- [Flutter Docs: Keyboard Focus System](https://docs.flutter.dev/ui/interactivity/focus)
- [Sendbird: Marketplace Buyer-Seller Chat Playbook](https://sendbird.com/resources/marketplace-playbook-a-guide-to-implementing-buyer-to-seller-chat)
- [CometChat: Improving Marketplace UX with In-App Chat](https://www.cometchat.com/blog/improving-marketplace-user-experience-with-in-app-chat)
- [UX Design: When to Use Loaders & Empty States](https://uxdesign.cc/when-to-use-loaders-empty-states-ebd23cecc7d6)
- [UX Design: Notification Center Design](https://uxdesign.cc/notification-center-7ec3d41efb10)
- [UXPin: WCAG 2.1.1 Keyboard Accessibility](https://www.uxpin.com/studio/blog/wcag-211-keyboard-accessibility-explained/)
- [UsableNet: E-commerce Accessibility Guide](https://blog.usablenet.com/ecommerce-website-accessibility-guide)
- [UXCam: Mobile UX Design Ultimate Guide 2026](https://uxcam.com/blog/mobile-ux/)
- [Digipixel: UI/UX 2025 Mobile vs Desktop](https://digipixel.sg/ui-ux-design-in-2025-mobile-vs-desktop-what-designers-absolutely-must-know/)
- [Downgraf: UX Lessons from Game Item Marketplaces](https://www.downgraf.com/all-articles/ux-lessons-from-game-item-marketplaces/)
- [Smashing Magazine: Optimistic UI](https://www.smashingmagazine.com/2016/11/true-lies-of-optimistic-user-interfaces/)
- [W3C: Progressive Enhancement vs Graceful Degradation](https://www.w3.org/wiki/Graceful_degradation_versus_progressive_enhancement)
