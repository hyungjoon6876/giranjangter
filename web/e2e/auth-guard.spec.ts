import { test, expect } from "@playwright/test";

test.describe("인증 가드 — 비로그인 사용자 보호", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/");
    await page.evaluate(() => {
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
    });
    await page.reload();
  });

  test("매물 등록 페이지 접근 시 로그인으로 리다이렉트", async ({ page }) => {
    await page.goto("/create");
    await page.waitForURL("**/login**", { timeout: 10000 });
    expect(page.url()).toContain("/login");
  });

  test("채팅 목록 접근 시 로그인으로 리다이렉트", async ({ page }) => {
    await page.goto("/chats");
    await page.waitForURL("**/login", { timeout: 5000 });
    expect(page.url()).toContain("/login");
  });

  test("매물 상세는 비로그인도 볼 수 있다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    // 홈에서 첫 매물 클릭
    await page.goto("/");
    await page.waitForLoadState("networkidle");

    const firstListing = page.locator("a[href^='/listings/']").first();
    if (await firstListing.isVisible()) {
      await firstListing.click();
      await page.waitForURL("**/listings/**");
      await page.waitForLoadState("networkidle");

      // 크래시 없이 로드되어야 함
      const jsErrors = errors.filter((e) => e.includes("TypeError") || e.includes("Cannot read"));
      expect(jsErrors).toHaveLength(0);

      // 상세 페이지 URL 유지 (리다이렉트 안 됨)
      expect(page.url()).toContain("/listings/");
    }
  });

  test("프로필 페이지는 로그인 안내를 표시한다", async ({ page }) => {
    await page.goto("/profile");
    await page.waitForLoadState("networkidle");
    await expect(page.getByText("로그인이 필요합니다")).toBeVisible();
  });
});

// 로그인 사용자 테스트는 Google OAuth가 필요하므로 E2E에서는 제외
// 실제 로그인 후 동작은 수동 테스트 또는 Google OAuth 테스트 환경에서 검증
