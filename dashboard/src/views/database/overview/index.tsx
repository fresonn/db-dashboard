import { Header } from '@/components/layout/header/header'
import { type Database } from '@/lib/api/gen'

export function DatabaseOverviewPage({ database }: { database: Database }) {
  return (
    <div>
      <Header title={`Database — ${database.name}`} />

      <p>{database.id}</p>
    </div>
  )
}
