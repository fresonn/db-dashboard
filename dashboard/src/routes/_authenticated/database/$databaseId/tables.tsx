import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_authenticated/database/$databaseId/tables')({
  component: RouteComponent
})

function RouteComponent() {
  return <div>Hello "/_authenticated/database/$databaseId/tables"!</div>
}
