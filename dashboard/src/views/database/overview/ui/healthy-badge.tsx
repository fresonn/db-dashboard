import { Typography } from '@/components/ui/typography'
import { BadgeCheck } from 'lucide-react'

/*
if (
  isReachable &&
  errorRate < threshold &&
  connectionsOk &&
  replicationOk
) {
  status = "healthy"
}
*/

export function HealthyBadge() {
  return (
    <div className="bg-theme-color/10 border-theme-color/50 text-theme-color flex items-center rounded-lg border px-5 py-0.5">
      <BadgeCheck strokeWidth={1.5} size={20} className="mr-1" />
      <Typography variant="small">Healthy</Typography>
    </div>
  )
}
