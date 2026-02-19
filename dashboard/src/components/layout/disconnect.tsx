import React from 'react'
import { Button } from '@/components/ui/button'
import { Unplug } from 'lucide-react'
import { useClusterDisconnect } from '@/lib/api/gen'
import { toast } from 'sonner'
import { capitalize } from '@/lib/utils'
import { useNavigate } from '@tanstack/react-router'
import { SidebarMenuButton } from '../ui/shadcn/sidebar'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogMedia,
  AlertDialogTitle,
  AlertDialogTrigger
} from '@/components/ui/shadcn/alert-dialog'

export function DisconnectDialog() {
  const navigate = useNavigate()

  const { isPending, mutateAsync } = useClusterDisconnect({
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

  const handleDisconnect: React.MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await mutateAsync()
  }

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <SidebarMenuButton asChild className="cursor-pointer" tooltip="Disconnect">
          <div className="text-red-500">
            <Unplug className="size-3" />
            Disconnect
          </div>
        </SidebarMenuButton>
      </AlertDialogTrigger>
      <AlertDialogContent size="sm">
        <AlertDialogHeader>
          <AlertDialogMedia className="bg-destructive/10 text-destructive dark:bg-destructive/20">
            <Unplug className="size-7" />
          </AlertDialogMedia>
          <AlertDialogTitle>Disconnect from cluster?</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to disconnect? You will need to connect again to access dashboard.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel asChild>
            <Button variant="outline">Cancel</Button>
          </AlertDialogCancel>
          <AlertDialogAction asChild>
            <Button loading={isPending} variant="destructive" onClick={handleDisconnect}>
              Disconnect
            </Button>
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
