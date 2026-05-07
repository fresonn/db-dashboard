import { Header } from '@/components/layout/header/header'

export function DatabaseOverviewPage({ databaseId }: { databaseId: string }) {
  return (
    <div>
      <Header title="Database Overview" />
      <p>{databaseId}</p>
    </div>
  )
}
