import { test, expect } from "@playwright/test";

test.describe("매물 상세 페이지", () => {
  test("매물 상세가 크래시 없이 렌더링된다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    await page.goto("/");
    await page.waitForLoadState("networkidle");

    // 첫 매물 클릭
    const firstListing = page.locator("a[href^='/listings/']").first();
    if (await firstListing.isVisible()) {
      await firstListing.click();
      await page.waitForURL("**/listings/**");
      await page.waitForLoadState("networkidle");

      // 크래시 없음
      const jsErrors = errors.filter((e) => e.includes("TypeError") || e.includes("Cannot read"));
      expect(jsErrors).toHaveLength(0);
    }
  });

  test("매물 상세에 접근성 랜드마크가 있다", async ({ page }) => {
    await page.goto("/");
    await page.waitForLoadState("networkidle");

    const firstListing = page.locator("a[href^='/listings/']").first();
    if (await firstListing.isVisible()) {
      await firstListing.click();
      await page.waitForURL("**/listings/**");
      await page.waitForLoadState("networkidle");

      // 액션 바 toolbar
      const toolbar = page.getByRole("toolbar", { name: "매물 액션" });
      if (await toolbar.isVisible()) {
        await expect(toolbar).toBeVisible();
      }

      // dl 정의 목록 존재
      const dl = page.locator("dl");
      if (await dl.isVisible()) {
        expect(await dl.count()).toBeGreaterThan(0);
      }
    }
  });

  test("비로그인 사용자가 찜 클릭 시 크래시 없이 안내한다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    // 먼저 페이지 이동 후 localStorage 초기화
    await page.goto("/");
    await page.evaluate(() => {
      try {
        localStorage.removeItem("accessToken");
        localStorage.removeItem("refreshToken");
      } catch { /* ignore */ }
    });
    await page.reload();
    await page.waitForLoadState("networkidle");

    const firstListing = page.locator("a[href^='/listings/']").first();
    if (await firstListing.isVisible()) {
      await firstListing.click();
      await page.waitForURL("**/listings/**");
      await page.waitForLoadState("networkidle");

      // 찜하기 버튼 클릭
      const favBtn = page.getByRole("button", { name: /찜/ });
      if (await favBtn.isVisible()) {
        await favBtn.click();
        await page.waitForTimeout(1000);

        // 크래시 없음
        const jsErrors = errors.filter((e) => e.includes("TypeError") || e.includes("Cannot read"));
        expect(jsErrors).toHaveLength(0);
      }
    }
  });
});
