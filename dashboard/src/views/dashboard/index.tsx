import { Header } from '@/components/layout/header/header'
import { ChartWAL } from './wal'

export function DashboardPage() {
  return (
    <div>
      <Header title="Dashboard" />
      <ChartWAL />
    </div>
  )
}
