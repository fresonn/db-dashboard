import { Button } from '@/components/ui/button'
import { useDatabasesDetailed } from '@/lib/api/gen'
import { useQueryParams } from '@/views/overview/databases/params'
import { createFileRoute } from '@tanstack/react-router'
import type { SortingState } from '@tanstack/react-table'
import { useState } from 'react'

export const Route = createFileRoute('/_authenticated/apps')({
  component: RouteComponent
})

function RouteComponent() {
  const [sorting, setSorting] = useState<SortingState>([])

  const params = useQueryParams(sorting)

  const { data } = useDatabasesDetailed(params)

  return (
    <div>
      <div className="mb-3">
        <Button onClick={() => setSorting([{ id: 'sizeBytes', desc: false }])}>Set ASC</Button>
        <Button onClick={() => setSorting([{ id: 'sizeBytes', desc: true }])}>Set Desc</Button>
        <Button onClick={() => setSorting([])}>Clear</Button>
      </div>
      {data?.map((db) => (
        <div key={db.name}>{JSON.stringify({ name: db.name, size: db.sizePretty })}</div>
      ))}
    </div>
  )
}
