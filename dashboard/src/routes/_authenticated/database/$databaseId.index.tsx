import { createFileRoute } from '@tanstack/react-router'
import { DatabaseOverviewPage } from '@/views/database/overview'

export const Route = createFileRoute('/_authenticated/database/$databaseId/')({
  component: Page
})

function Page() {
  const { databaseId } = Route.useParams()

  return <DatabaseOverviewPage databaseId={databaseId} />
}
