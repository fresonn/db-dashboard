import { useMemo } from 'react'
import { databasesDetailedQueryParamsOrderEnum } from '@/lib/api/gen'

const sortFieldMap: Record<string, 'size' | 'connection'> = {
  sizeBytes: 'size',
  totalConnections: 'connection'
}

export function useQueryParams(sorting: { id: string; desc: boolean }[]) {
  return useMemo(() => {
    if (!sorting.length) return undefined

    const { id, desc } = sorting[0]
    return {
      sort: sortFieldMap[id],
      order: desc
        ? databasesDetailedQueryParamsOrderEnum.desc
        : databasesDetailedQueryParamsOrderEnum.asc
    } as const
  }, [sorting])
}
