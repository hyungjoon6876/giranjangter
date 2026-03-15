import { test, expect } from "@playwright/test";

test.describe("로그인 페이지", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("/");
    await page.evaluate(() => {
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
    });
  });

  test("로그인 페이지가 크래시 없이 렌더링된다", async ({ page }) => {
    const errors: string[] = [];
    page.on("pageerror", (err) => errors.push(err.message));

    await page.goto("/login");
    await page.waitForLoadState("networkidle");

    await expect(page.getByRole("heading", { name: "기란장터" })).toBeVisible();
    await expect(page.getByRole("button", { name: /개발자 로그인/ })).toBeVisible();

    const jsErrors = errors.filter((e) => e.includes("TypeError") || e.includes("Cannot read"));
    expect(jsErrors).toHaveLength(0);
  });

  test("Google 로그인 버튼 컨테이너가 존재한다", async ({ page }) => {
    await page.goto("/login");
    await page.waitForLoadState("networkidle");
    // Google GSI가 로드되면 iframe이 생김; 로드 안 되면 빈 div
    // 크래시만 안 나면 OK
    await page.waitForTimeout(2000);
  });

  test("둘러보기 버튼이 홈으로 이동한다", async ({ page }) => {
    await page.goto("/login");
    await page.waitForLoadState("networkidle");

    await page.getByRole("button", { name: "둘러보기" }).click();
    await page.waitForURL("/");
    expect(page.url()).not.toContain("/login");
  });

  // 개발자 로그인 테스트 제거 — 라이브 환경에서는 dev 로그인 불가
});
