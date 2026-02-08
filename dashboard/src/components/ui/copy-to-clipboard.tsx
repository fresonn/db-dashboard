import { cva, type VariantProps } from 'cva'
import { Check, Copy } from 'lucide-react'
import { toast } from 'sonner'
import { useCopyToClipboard } from '@/hooks/use-copy-to-clipboard'

const containerVariant = cva('shrink-0 scale-120 cursor-pointer', {
  variants: {
    size: {
      sm: 'text-sm',
      base: 'text-base',
      lg: 'text-lg',
      xl: 'text-xl'
    }
  }
})

export function CopyToClipboard({
  what,
  size = 'sm'
}: {
  what: string
} & VariantProps<typeof containerVariant>) {
  const [copiedText, copy] = useCopyToClipboard()

  const handleClick = () => {
    copy(what)
      .then(() => {
        toast.success('Copied to clipboard!', {
          duration: 2000
        })
      })
      .catch(() => {
        toast.error('Failed to copy', {
          duration: 3000,
          description: 'Clipboard disabled or not supported, check your browser preferences.'
        })
      })
  }

  return (
    <div onClick={handleClick} className={containerVariant({ size })}>
      {copiedText ? (
        <Check size="1em" color="currentColor" />
      ) : (
        <Copy size="1em" color="currentColor" />
      )}
    </div>
  )
}
