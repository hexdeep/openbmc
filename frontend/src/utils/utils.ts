
export interface List<T> {
  data: T[],
  total: number,
}

export function formatSize(size: number) {
  if (size < 1024) return size + ' B';
  if (size < 1024 * 1024) return (size / 1024).toFixed(1) + ' KB';
  if (size < 1024 * 1024 * 1024) return (size / 1024 / 1024).toFixed(1) + ' MB';
  return (size / 1024 / 1024 / 1024).toFixed(1) + ' GB';
}

export function delay(ms: number) {
  return () =>
    new Promise<void>((resolve) => {
      setTimeout(resolve, ms);
    });
}

