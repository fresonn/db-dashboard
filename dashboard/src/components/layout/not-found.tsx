import { Search } from 'lucide-react'
import { Typography } from '../ui/typography'

export function NotFound({ resource }: { resource?: string }) {
  return (
    <div className="h-screen-sub-header relative">
      <div className="absolute top-5/12 left-1/2 -translate-x-1/2 -translate-y-1/2">
        <div className="flex flex-col items-center">
          <div className="bg-section-box mb-2 rounded-xl border p-5">
            <Search size={45} />
          </div>
          <Typography className="mb-2 text-4xl uppercase">{resource} not found</Typography>
          <Typography variant="small" className="text-neutral-500">
            Make sure the resource ID is correct or check the logs
          </Typography>
        </div>
      </div>
    </div>
  )
}
