import { useState } from 'react'
import type { SortingState } from '@tanstack/react-table'
import { useQueryParams } from './params'
import { DatabasesTable } from './table/table'
import { useDatabasesDetailed } from '@/lib/api/gen'
import { DatabasesTableSkeleton } from './table/skeleton-loader'

export function AvailableDatabases() {
  const [sorting, setSorting] = useState<SortingState>([])

  const queryParams = useQueryParams(sorting)

  const { data, isLoading } = useDatabasesDetailed(queryParams)
  if (isLoading) {
    return <DatabasesTableSkeleton rows={11} />
  }

  return <DatabasesTable data={data ?? []} sorting={sorting} onSortingChange={setSorting} />
}
