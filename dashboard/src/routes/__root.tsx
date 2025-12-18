import { lazy, Suspense } from 'react'
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { useQuery, useQueryClient, type QueryClient } from '@tanstack/react-query'
import { clusterStatusQuery } from '@/lib/api/cluster-status'
import { OfflineApiGuard } from '@/components/offline-api-guard'
import { Toaster } from 'sonner'

export interface RouterContext {
  queryClient: QueryClient
}

const TanStackRouterDevtools = import.meta.env.PROD
  ? () => null
  : lazy(() =>
      import('@tanstack/react-router-devtools').then((res) => ({
        default: res.TanStackRouterDevtools
      }))
    )

export const Route = createRootRouteWithContext<RouterContext>()({
  async beforeLoad({ context }) {
    await context.queryClient.prefetchQuery(clusterStatusQuery)

    const clusterStatus = context.queryClient.getQueryData(clusterStatusQuery.queryKey)

    return {
      clusterStatus: clusterStatus
    }
  },
  component: () => {
    const { clusterStatus } = Route.useRouteContext()

    const queryClient = useQueryClient()

    const { isError, isFetching } = useQuery({
      ...clusterStatusQuery,
      placeholderData: clusterStatus,
      staleTime: Infinity,
      gcTime: Infinity,
      refetchOnWindowFocus: false,
      refetchOnReconnect: false,
      refetchOnMount: false,
      retry: false
    })

    const handleRetry = () => {
      queryClient.refetchQueries({
        queryKey: clusterStatusQuery.queryKey
      })
    }

    return (
      <>
        <OfflineApiGuard open={isError} onRetry={handleRetry} loading={isFetching} />
        <Outlet />
        <Toaster
          position="bottom-right"
          visibleToasts={4}
          toastOptions={{
            classNames: {
              toast: '!border-border dark:!bg-neutral-900 !bg-background !text-foreground',
              description: '!text-muted-foreground',
              success: '!border-green-500/30 !text-green-500',
              error: '!border-red-500/30 dark:!text-red-400 !text-red-600'
            }
          }}
        />
        <Suspense>
          <TanStackRouterDevtools position="bottom-right" />
        </Suspense>
      </>
    )
  }
})
