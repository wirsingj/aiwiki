export type InfoboxRow = {
  label: string;
  value: string;
};

export type Infobox = {
  heading: string;
  rows: InfoboxRow[];
};

export type Section = {
  heading: string;
  paragraphs: string[];
  links: string[];
};

export type Reference = {
  label: string;
  note: string;
};

export type Article = {
  slug: string;
  title: string;
  summary: string;
  infobox?: Infobox;
  sections: Section[];
  references: Reference[];
  seeAlso: string[];
  generatedAt: string;
  cached: boolean;
};

export const MIN_TOPIC_LENGTH = 1;
export const MAX_TOPIC_LENGTH = 120;

declare global {
  interface Window {
    __AIWIKI_CONFIG__?: {
      API_BASE_URL?: string;
    };
  }
}

const API_BASE_URL = normalizeBaseURL(
  window.__AIWIKI_CONFIG__?.API_BASE_URL ?? import.meta.env.VITE_API_BASE_URL ?? '',
);

export function slugify(input: string): string {
  const lower = input.trim().toLowerCase();
  let slug = '';
  let lastDash = false;

  for (const char of lower) {
    if (/[\p{Letter}\p{Number}]/u.test(char)) {
      slug += char;
      lastDash = false;
    } else if (!lastDash && slug.length > 0) {
      slug += '-';
      lastDash = true;
    }
  }

  slug = slug.replace(/^-+|-+$/g, '');
  return slug || 'untitled';
}

export function normalizeTopicInput(input: string): string {
  return input.replace(/\s+/gu, ' ').trim();
}

export function validateTopicInput(input: string): string | null {
  const normalized = normalizeTopicInput(input);
  if (normalized.length < MIN_TOPIC_LENGTH) {
    return 'Enter a topic to search.';
  }
  if ([...normalized].length > MAX_TOPIC_LENGTH) {
    return `Topics must be ${MAX_TOPIC_LENGTH} characters or fewer.`;
  }
  if (/[\u0000-\u0008\u000B\u000C\u000E-\u001F\u007F]/u.test(input)) {
    return 'Topics cannot contain control characters.';
  }
  return null;
}

export function safeTopicSlug(input: string): string {
  return slugify(normalizeTopicInput(input));
}

export async function fetchArticle(slug: string, refresh = false): Promise<Article> {
  const query = refresh ? '?refresh=1' : '';
  return request<Article>(`${API_BASE_URL}/api/article/${encodeURIComponent(slug)}${query}`);
}

export async function createArticle(topic: string, refresh = false): Promise<Article> {
  return request<Article>(`${API_BASE_URL}/api/article`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ topic, refresh }),
  });
}

export async function fetchRandomArticle(): Promise<Article> {
  return request<Article>(`${API_BASE_URL}/api/random`);
}

async function request<T>(url: string, init?: RequestInit): Promise<T> {
  const response = await fetch(url, init);
  const data = await response.json().catch(() => null);

  if (!response.ok) {
    const message = data?.error ?? `Request failed with status ${response.status}`;
    throw new Error(message);
  }

  return data as T;
}

function normalizeBaseURL(value: string): string {
  return value.trim().replace(/\/+$/, '');
}
