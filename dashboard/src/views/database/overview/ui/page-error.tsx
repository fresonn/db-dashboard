import { TriangleAlert } from 'lucide-react'
import { Typography } from '@/components/ui/typography'

export function PageError() {
  return (
    <div className="h-screen-sub-header relative">
      <div className="absolute top-5/12 left-1/2 -translate-x-1/2 -translate-y-1/2">
        <div className="text-destructive flex flex-col items-center">
          <div className="border-destructive/50 bg-destructive/10 mb-2 rounded-xl border p-5 text-red-500">
            <TriangleAlert size={45} strokeWidth={1.5} />
          </div>
          <Typography className="mb-2 text-4xl uppercase">Failed to fetch database</Typography>
          <Typography variant="small" className="text-destructive/80">
            Try to find out what is happening in the logs.
          </Typography>
        </div>
      </div>
    </div>
  )
}
