import { test, expect } from "@playwright/test";

test.describe("네비게이션 + 접근성", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/");
    await page.evaluate(() => {
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
    });
    await page.reload();
    await page.waitForLoadState("networkidle");
  });

  test("skip-to-content 링크가 Tab으로 접근 가능하다", async ({ page }) => {
    await page.keyboard.press("Tab");
    const skipLink = page.getByRole("link", { name: "본문으로 건너뛰기" });
    await expect(skipLink).toBeFocused();
  });

  test("헤더 네비게이션에 aria-current가 설정된다", async ({ page }) => {
    const nav = page.getByRole("navigation", { name: "메인 메뉴" });
    const activeLink = nav.getByRole("link", { name: "마켓" });
    await expect(activeLink).toHaveAttribute("aria-current", "page");
  });

  test("하단 탭바에 aria-current가 설정된다", async ({ page }) => {
    // 모바일 뷰포트로 변경
    await page.setViewportSize({ width: 375, height: 812 });
    await page.reload();
    await page.waitForLoadState("networkidle");

    const bottomNav = page.getByRole("navigation", { name: "하단 메뉴" });
    await expect(bottomNav).toBeVisible();

    const activeTab = bottomNav.getByRole("link", { name: "마켓" });
    await expect(activeTab).toHaveAttribute("aria-current", "page");
  });

  test("로그인 페이지로 이동할 수 있다", async ({ page }) => {
    await page.getByRole("banner").getByRole("link", { name: "로그인" }).click();
    await page.waitForURL("**/login");
    await expect(page.getByRole("heading", { name: "기란장터" })).toBeVisible();
  });

  test("모든 주요 페이지가 크래시 없이 로드된다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    const routes = ["/", "/login", "/notifications", "/profile"];
    for (const route of routes) {
      await page.goto(route);
      await page.waitForLoadState("networkidle");
    }

    const jsErrors = errors.filter(
      (e) => e.includes("TypeError") || e.includes("ReferenceError") || e.includes("Cannot read")
    );
    expect(jsErrors).toHaveLength(0);
  });
});
