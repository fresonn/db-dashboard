import { lazy, Suspense } from 'react'
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import type { QueryClient } from '@tanstack/react-query'
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
    try {
      const clusterStatus = await context.queryClient.ensureQueryData(clusterStatusQuery)
      return {
        clusterStatus,
        connectionError: null
      }
    } catch (error) {
      return {
        connectionError: error + ''
      }
    }
  },
  component: () => {
    const context = Route.useRouteContext()

    return (
      <>
        <OfflineApiGuard open={!!context.connectionError} />
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
