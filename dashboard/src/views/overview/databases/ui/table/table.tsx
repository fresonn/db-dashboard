import { useNavigate } from '@tanstack/react-router'
import { Typography } from '@/components/ui/typography'
import { TableToolbar } from './toolbar'
import { columns } from './columns'
import { toast } from 'sonner'
import { capitalize } from '@/lib/utils'
import { useQueryClient } from '@tanstack/react-query'
import type { Database } from './types'
import {
  getCoreRowModel,
  useReactTable,
  flexRender,
  getFacetedRowModel,
  getFacetedUniqueValues,
  getFilteredRowModel,
  getFacetedMinMaxValues,
  type SortingState,
  type OnChangeFn
} from '@tanstack/react-table'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/components/ui/shadcn/table'
import { getStatusQueryKey, useSwitchCurrentDatabase } from '@/lib/api/gen'

export type DatabasesTableProps = {
  data: Database[]
  sorting: SortingState
  isFetching: boolean
  onSortingChange: OnChangeFn<SortingState>
}

export function DatabasesTable({
  data,
  sorting,
  isFetching,
  onSortingChange
}: DatabasesTableProps) {
  const table = useReactTable({
    columns,
    data,
    manualSorting: true,
    state: {
      sorting
    },
    onSortingChange,
    getCoreRowModel: getCoreRowModel(),
    getFacetedRowModel: getFacetedRowModel(),
    getFacetedUniqueValues: getFacetedUniqueValues(),
    getFilteredRowModel: getFilteredRowModel(),
    getFacetedMinMaxValues: getFacetedMinMaxValues()
  })

  const navigate = useNavigate()
  const queryClient = useQueryClient()

  const { mutate } = useSwitchCurrentDatabase()

  const handleDatabaseClick = (databaseId: string, databaseName: string) => {
    mutate(
      {
        data: {
          databaseId: Number(databaseId)
        }
      },
      {
        onSuccess: () => {
          queryClient.invalidateQueries({
            queryKey: getStatusQueryKey()
          })
          navigate({
            to: '/database/$databaseId',
            params: {
              databaseId
            }
          })
        },
        onError: (error) => {
          toast.error(capitalize(error.message), {
            description: `Database name: ${databaseName}`,
            duration: 3000
          })
        }
      }
    )
  }

  return (
    <div>
      <TableToolbar table={table} />
      <div className="border">
        <Table wrapperClassName="h-[450px]">
          <TableHeader className="sticky top-0 bg-neutral-900">
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => (
                  <TableHead key={header.id} className="text-foreground">
                    {header.isPlaceholder
                      ? null
                      : flexRender(header.column.columnDef.header, {
                          ...header.getContext(),
                          isFetching
                        })}
                  </TableHead>
                ))}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody loading={isFetching}>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && 'selected'}
                  className={
                    row.original.allowConnections
                      ? 'cursor-pointer'
                      : 'cursor-not-allowed opacity-50'
                  }
                  onClick={() => {
                    if (row.getValue('allowConnections')) {
                      handleDatabaseClick(row.original.id, row.original.name)
                    }
                  }}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow className="absolute inset-0 flex items-center justify-center hover:bg-transparent">
                <TableCell>
                  <Typography variant="h3" className="text-foreground/70">
                    No databases found
                  </Typography>
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}
