import { CircleOff, Infinity as InfinityIcon } from 'lucide-react'
import { Typography } from './typography'

// pg-limit
// -1 — unlimited
//  0 — connections is not allowed
// >0 — max connections

export function PgLimitBadge({ state }: { state: number }) {
  switch (state) {
    case -1:
      return (
        <div className="flex items-center">
          <InfinityIcon strokeWidth={2} className="dark:text-theme-color mr-1 text-green-500" />
          <Typography as="span" variant="small">
            Unlimited
          </Typography>
        </div>
      )
    case 0:
      return (
        <div className="flex items-center">
          <CircleOff size={19} className="mr-1 text-red-500 dark:text-red-400" />
          <Typography as="span" variant="small">
            Not allowed
          </Typography>
        </div>
      )
    default:
      return (
        <div>
          <Typography variant="small">
            {state} {state === 1 ? 'connection' : 'connections'}
          </Typography>
        </div>
      )
  }
}
