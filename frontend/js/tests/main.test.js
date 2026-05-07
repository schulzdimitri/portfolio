import { describe, it, expect, vi, beforeEach } from 'vitest';
import { loadRuntimeConfig, resolveApiBase } from '../main.js';

describe('resolveApiBase', () => {
    it('uses runtime config when provided', () => {
        expect(resolveApiBase('https://api.example.com', false)).toBe('https://api.example.com');
    });

    it('falls back to localhost during local development', () => {
        expect(resolveApiBase('', true)).toBe('http://localhost:8080');
    });

    it('returns empty string for production when config is missing', () => {
        expect(resolveApiBase('', false)).toBe('');
    });
});

describe('loadRuntimeConfig', () => {
    beforeEach(() => {
        vi.restoreAllMocks();
    });

    it('reads apiBase from config.json', async () => {
        global.fetch = vi.fn().mockResolvedValue({
            ok: true,
            json: async () => ({ apiBase: 'https://api.example.com' }),
        });

        await expect(loadRuntimeConfig()).resolves.toEqual({ apiBase: 'https://api.example.com' });
    });

    it('ignores masked or invalid values', async () => {
        global.fetch = vi.fn().mockResolvedValue({
            ok: true,
            json: async () => ({ apiBase: '***' }),
        });

        await expect(loadRuntimeConfig()).resolves.toEqual({ apiBase: '' });
    });

    it('returns empty config when fetch fails', async () => {
        global.fetch = vi.fn().mockRejectedValue(new Error('network error'));

        await expect(loadRuntimeConfig()).resolves.toEqual({ apiBase: '' });
    });
});