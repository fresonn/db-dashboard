import { useState } from 'react'

import {
  getCoreRowModel,
  useReactTable,
  flexRender,
  getFilteredRowModel,
  getFacetedMinMaxValues,
  getPaginationRowModel
} from '@tanstack/react-table'

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/components/ui/shadcn/table'

import type { Role } from './types'
import { columns } from './columns'
import { Typography } from '@/components/ui/typography'
import { TablePagination } from './pagination'

export type RolesTableProps = {
  data: Role[]
}

export function RolesTable({ data }: RolesTableProps) {
  const isFetching = false

  const [pagination, setPagination] = useState({
    pageIndex: 0, //initial page index
    pageSize: 3 //default page size
  })

  const table = useReactTable({
    columns,
    data,
    state: {
      pagination
    },
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    onPaginationChange: setPagination,
    getFilteredRowModel: getFilteredRowModel(),
    getFacetedMinMaxValues: getFacetedMinMaxValues()
  })

  return (
    <div>
      <div className="relative">
        <Table wrapperClassName="">
          <TableHeader className="bg-neutral-900">
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
                  className="cursor-pointer"
                  //   onClick={() => {
                  //     if (row.getValue('allowConnections')) {
                  //       handleDatabaseClick(row.getValue('name'))
                  //     }
                  //   }}
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
        <div className="absolute -top-10 right-0">
          <TablePagination table={table} />
        </div>
      </div>
    </div>
  )
}
