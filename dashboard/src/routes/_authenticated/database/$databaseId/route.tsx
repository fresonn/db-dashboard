import { createFileRoute, Outlet } from '@tanstack/react-router'
import { Header } from '@/components/layout/header/header'

export const Route = createFileRoute('/_authenticated/database/$databaseId')({
  component: Page
})

function Page() {
  return (
    <div>
      <Header title="Database Overview" />
      <Outlet />
    </div>
  )
}
