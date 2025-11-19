import { lazy, Suspense } from 'react'
import { createRootRoute, Link, Outlet } from '@tanstack/react-router'

const TanStackRouterDevtools = import.meta.env.PROD
  ? () => null
  : lazy(() =>
      import('@tanstack/react-router-devtools').then((res) => ({
        default: res.TanStackRouterDevtools
        // For Embedded Mode
        // default: res.TanStackRouterDevtoolsPanel
      }))
    )

const RootLayout = () => (
  <>
    <div className="flex gap-2 p-2">
      <Link to="/" className="[&.active]:font-bold">
        Home
      </Link>{' '}
      <Link to="/connect" className="[&.active]:font-bold">
        Connect
      </Link>
    </div>
    <hr />
    <Outlet />
    <Suspense>
      <TanStackRouterDevtools position="bottom-right" />
    </Suspense>
  </>
)

export const Route = createRootRoute({ component: RootLayout })
