import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/_authenticated/database/$databaseId')({
  component: RouteComponent
})

function RouteComponent() {
  return (
    <div>
      <Outlet />
    </div>
  )
}
