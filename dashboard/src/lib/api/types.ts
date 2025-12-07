export function hasErrorField<T extends string>(obj: unknown, field: T): obj is Record<T, unknown> {
  return typeof obj === 'object' && obj !== null && field in obj
}
