import { Header } from '@/components/layout/header/header'
import { WALAnalytics } from './wal'

export function DashboardPage() {
  return (
    <div>
      <Header title="Dashboard" />
      <WALAnalytics />
    </div>
  )
}
