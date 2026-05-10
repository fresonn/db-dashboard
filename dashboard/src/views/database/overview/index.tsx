import { DatabaseHeader } from './ui/header'
import { useDatabase } from '@/lib/api/gen'

export function DatabaseOverviewPage({ databaseId }: { databaseId: string }) {
  const { data, isLoading, isError, error } = useDatabase(Number(databaseId))

  if (isLoading || !data) {
    return <p>Loading</p>
  }

  if (isError) {
    const e = error as { status?: number }

    if (e.status === 404) {
      return <p>Not Found</p>
    }

    return <p>{error.message}</p>
  }

  return (
    <div>
      <DatabaseHeader
        name={data?.name}
        owner={data?.owner}
        encoding={data?.encoding}
        ctype={'C'} // todo: add property
        tablespace={data?.tablespace}
        isTemplate={data?.isTemplate}
        description={data?.description}
      />
    </div>
  )
}
