import { Button } from '@/components/ui/button'
import { LogOut } from 'lucide-react'
import { useClusterDisconnect } from '@/lib/api/gen'
import { toast } from 'sonner'
import { capitalize } from '@/lib/utils'
import { useNavigate } from '@tanstack/react-router'
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from '@/components/ui/shadcn/alert-dialog'

export function DisconnectDialog() {
  const navigate = useNavigate()

  const { isPending, mutate } = useClusterDisconnect({
    mutation: {
      onError: (error) => {
        toast.error('Disconnect failed', {
          duration: 3000,
          description: capitalize(error.message)
        })
      },
      onSuccess: () => {
        toast.success('Disconnected successfully', { duration: 5000 })
        navigate({
          to: '/connect'
        })
      }
    }
  })

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button className="text-sm text-red-500" variant="ghost">
          <LogOut />
          Disconnect
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Disconnect from cluster?</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to disconnect? You will need to connect again to access dashboard.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <Button loading={isPending} variant="destructive" onClick={() => mutate()}>
            Disconnect
          </Button>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
