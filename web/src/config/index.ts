export const Config = {
  REST_ENDPOINT: import.meta.env.VITE_REST_ENDPOINT || 'http://localhost:8080',
  DEFAULT_REST_TIMEOUT_MS: 5000,
  DETAULT_REFETCH_INTERVAL: 15_000,
};
