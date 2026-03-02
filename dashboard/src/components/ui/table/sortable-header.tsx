import type { Column } from '@tanstack/react-table'
import { ChevronsUpDown, ChevronUp } from 'lucide-react'

type SortableHeaderProps<TData> = {
  column: Column<TData, unknown>
  title: string
}

export function SortableHeader<TData>({ column, title }: SortableHeaderProps<TData>) {
  const isSorted = column.getIsSorted()

  return (
    <div
      onClick={column.getToggleSortingHandler()}
      className="flex cursor-pointer items-center gap-0.5 select-none"
    >
      {title}

      {isSorted === 'asc' && <ChevronUp className="size-5" />}
      {isSorted === 'desc' && <ChevronUp className="size-5 rotate-180" />}
      {!isSorted && <ChevronsUpDown strokeWidth={3} className="size-4 opacity-50" />}
    </div>
  )
}
