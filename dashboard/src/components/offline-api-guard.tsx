import { Button } from './ui/button'
import { TriangleAlert } from 'lucide-react'
import { Typography } from '@/components/ui/typography'
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader
} from '@/components/ui/shadcn/alert-dialog'

export function OfflineApiGuard({
  open,
  loading,
  onRetry
}: {
  open: boolean
  onRetry(): void
  loading: boolean
}) {
  return (
    <AlertDialog open={open}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <div className="mb-5 flex justify-center text-red-500">
            <div className="flex size-16 items-center justify-center rounded-full bg-red-500/20">
              <TriangleAlert className="size-10" />
            </div>
          </div>
          <Typography variant="h3">No connection with api</Typography>
          <Typography variant="muted">
            Connection to dashboard api server failed or lost. Check the logs or the correctness of
            the launch data
          </Typography>
        </AlertDialogHeader>
        <AlertDialogFooter className="pt-4">
          <Button size="lg" fullWidth loading={loading} onClick={onRetry}>
            Refresh
          </Button>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
