import { DatabaseHeader } from './ui/header'
import { useDatabase } from '@/lib/api/gen'
import { NotFound } from '@/components/layout/not-found'
import { PageError } from './ui/page-error'

export function DatabaseOverviewPage({ databaseId }: { databaseId: string }) {
  const { data, isLoading, isError, error } = useDatabase(Number(databaseId), {
    query: {
      retry: false,
      refetchOnWindowFocus: false,
      refetchOnReconnect: false
    }
  })

  if (isLoading) return <p>Loading</p>

  if (isError) {
    const e = error as { status?: number }

    if (e.status === 404) return <NotFound resource="database" />

    return <PageError />
  }

  if (!data) return <NotFound resource="database" />

  return (
    <div>
      <DatabaseHeader
        name={data?.name}
        owner={data?.owner}
        encoding={data?.encoding}
        ctype={data.ctype}
        tablespace={data?.tablespace}
        isTemplate={data?.isTemplate}
        description={data?.description}
      />
    </div>
  )
}
