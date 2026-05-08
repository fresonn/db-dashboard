import { createFileRoute } from '@tanstack/react-router'
import { DatabaseOverviewPage } from '@/views/database/overview'
import { databaseQueryOptions, type Database400, type Database404 } from '@/lib/api/gen'

export const Route = createFileRoute('/_authenticated/database/$databaseId/')({
  loader: async ({ params, context }) => {
    return await context.queryClient.ensureQueryData(
      databaseQueryOptions(Number(params.databaseId))
    )
  },
  pendingComponent: () => <p>Loading</p>,
  component: () => {
    const data = Route.useLoaderData()

    return <DatabaseOverviewPage database={data} />
  },
  errorComponent: ({ error }) => {
    const e = error as {
      status?: number
      body?: Database400 | Database404
    }

    if (e.status === 404) {
      // return <NotFound />
      return <p>Database not found</p>
    }

    if (e.status === 400) {
      // return <BadRequest />
      return <p>Bad request</p>
    }

    // return <Error />
    return <p>Some Error</p>
  },
  notFoundComponent: () => {
    return <h1>Page not found</h1>
  }
})
