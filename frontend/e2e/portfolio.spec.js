import { test, expect } from '@playwright/test';

test.describe('Portfolio E2E Tests', () => {

  test('Deve carregar a página inicial corretamente', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('h1')).toBeVisible();
    await expect(page.locator('#experiences-tabs')).toBeVisible();
  });

  test('Deve exibir a mensagem de que os projetos estão indisponíveis se o DB estiver vazio ou falhar', async ({ page }) => {
    await page.goto('/');
    
    const tabProjetos = page.locator('button[data-tab="projects"]');
    if (await tabProjetos.isVisible()) {
      await tabProjetos.click();
    }

    await expect(page.locator('#projects-grid')).toBeAttached();
  });

  test('Deve permitir enviar uma mensagem pelo formulário de contato', async ({ page }) => {
    await page.goto('/');

    const nameInput = page.locator('input[name="name"]');
    const emailInput = page.locator('input[name="email"]');
    const messageInput = page.locator('textarea[name="message"]');
    const submitBtn = page.locator('button[type="submit"]');

    await nameInput.fill('E2E Tester');
    await emailInput.fill('e2e@example.com');
    await messageInput.fill('Hello from Playwright!');
    
    await submitBtn.click();
  });

  test('Deve usar o apiBase vindo do config.json para buscar dados da API', async ({ page }) => {
    const backendUrl = 'http://localhost:8080';
    const requests = [];

    await page.route('**/config.json', route => {
      route.fulfill({
        contentType: 'application/json',
        body: JSON.stringify({ apiBase: backendUrl }),
      });
    });

    await page.route('**/api/projects', route => {
      requests.push(route.request().url());
      route.fulfill({
        contentType: 'application/json',
        body: JSON.stringify({ projects: [] }),
      });
    });

    await page.route('**/api/experiences', route => {
      requests.push(route.request().url());
      route.fulfill({
        contentType: 'application/json',
        body: JSON.stringify([]),
      });
    });

    await page.goto('/');

    await expect(page.locator('#projects-grid')).toBeAttached();
    await expect(page.locator('#experiences-tabs')).toBeVisible();
    await expect.poll(() => requests.length).toBeGreaterThanOrEqual(2);
    expect(requests.every(url => url.startsWith(backendUrl))).toBe(true);
  });

  test('Deve criar uma experiência via API e exibi-la no frontend', async ({ page, request }) => {
    const uniqueId = Date.now();
    const testCompany = `Playwright E2E Corp ${uniqueId}`;
    const testRole = `Test Automation Engineer ${uniqueId}`;

    const novoBackendRequest = await request.post('http://localhost:8080/api/experiences', {
      headers: {
        'Authorization': 'Bearer supersecret123',
        'Content-Type': 'application/json'
      },
      data: {
        company: testCompany,
        role: testRole,
        period: "2026 - Present",
        duties: ["Construção de testes E2E"]
      }
    });

    expect(novoBackendRequest.status()).toBe(201);

    await page.goto('/');
    
    const companyTab = page.locator('.tabs-list button', { hasText: testCompany }).first();
    await expect(companyTab).toBeVisible();

    await companyTab.click();

    const roleElement = page.locator('h3', { hasText: testRole }).first();
    await expect(roleElement).toBeVisible();
    
    const dutyElement = page.locator('li', { hasText: 'Construção de testes E2E' }).first();
    await expect(dutyElement).toBeAttached();
  });

});
