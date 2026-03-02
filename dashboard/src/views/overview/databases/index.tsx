import { useState } from 'react'
import type { SortingState } from '@tanstack/react-table'
import { useQueryParams } from './params'
import { DatabasesTable } from './table/table'
import { useDatabasesDetailed } from '@/lib/api/gen'

export function AvailableDatabases() {
  const [sorting, setSorting] = useState<SortingState>([])

  const queryParams = useQueryParams(sorting)

  const { data } = useDatabasesDetailed(queryParams)

  return <DatabasesTable data={data ?? []} sorting={sorting} onSortingChange={setSorting} />
}
