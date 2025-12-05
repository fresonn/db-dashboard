import postgresLogoDark from '@/assets/postgresql-logo-black.svg?url'
import postgresLogo from '@/assets/postgresql-logo.svg?url'
import { Button } from '@/components/ui/button'
import { Label } from '@/components/ui/label'
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput
} from '@/components/ui/shadcn/input-group'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/shadcn/popover'
import { ThemeToggle } from '@/components/ui/theme-toggle'
import { TypographyH1, TypographySmall } from '@/components/ui/typography'
import { useTheme } from '@/hooks/use-theme'
import { useClusterConnect } from '@/lib/api/gen'
import { valibotResolver } from '@hookform/resolvers/valibot'
import { Database, EthernetPort, Info, KeyRound, Server, UserRound } from 'lucide-react'
import { useForm, type SubmitHandler } from 'react-hook-form'
import styles from './index.module.css'
import { validationSchema, type FormSubmitData } from './validation-schema'
import { toast } from 'sonner'
import { capitalize } from '@/lib/utils'
import { hasErrorField } from '@/lib/api/types'
import { useNavigate } from '@tanstack/react-router'

export function ConnectView() {
  const navigate = useNavigate()
  const [theme] = useTheme()

  const { isPending, mutate } = useClusterConnect({
    mutation: {
      onError: (error) => {
        console.log(error)

        toast.error(capitalize(error.message), {
          description: hasErrorField(error, 'reason') && error.reason,
          duration: 3000
        })
      },
      onSuccess: () => {
        navigate({
          to: '/overview'
        })
      }
    }
  })

  const { register, handleSubmit } = useForm({
    mode: 'onSubmit',
    resolver: valibotResolver(validationSchema),
    shouldUnregister: true
  })

  const onSubmit: SubmitHandler<FormSubmitData> = (values) => {
    if (values.database === '') {
      values.database = undefined
    }

    mutate({ data: { ...values } })
  }

  return (
    <div className="grid h-screen grid-cols-[2fr_1fr]">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="relative flex items-center justify-center dark:bg-neutral-900"
      >
        <div className="absolute top-8 right-10">
          <ThemeToggle />
        </div>
        <div className="w-full max-w-[700px] px-12 py-10">
          <div className="flex items-center justify-center">
            <img
              src={theme === 'light' ? postgresLogoDark : postgresLogo}
              className="mr-2 size-16"
              alt="postgresql logo"
            />
            <TypographyH1>Simple Dashboard</TypographyH1>
          </div>
          <div className="mb-10 pt-10">
            <ul className="*:mb-4">
              <li>
                <Label htmlFor="database-input">Database (optional)</Label>
                <InputGroup>
                  <InputGroupAddon>
                    <Database />
                  </InputGroupAddon>
                  <InputGroupInput
                    id="database-input"
                    placeholder="postgres"
                    className="pl-2!"
                    {...register('database')}
                  />
                  <InputGroupAddon align="inline-end">
                    <Popover>
                      <PopoverTrigger asChild>
                        <InputGroupButton className="cursor-pointer rounded-full" size="icon-xs">
                          <Info />
                        </InputGroupButton>
                      </PopoverTrigger>
                      <PopoverContent className="text-sm leading-normal">
                        By default it tries to connect to{' '}
                        <TypographySmall>postgres</TypographySmall> database
                      </PopoverContent>
                    </Popover>
                  </InputGroupAddon>
                </InputGroup>
              </li>
              <li className="flex">
                <div className="mr-8">
                  <Label htmlFor="host-input">Host</Label>
                  <InputGroup>
                    <InputGroupInput
                      id="host-input"
                      placeholder="127.0.0.1"
                      {...register('host')}
                    />
                    <InputGroupAddon>
                      <Server />
                    </InputGroupAddon>
                  </InputGroup>
                </div>
                <div>
                  <Label htmlFor="port-input">Port</Label>
                  <InputGroup>
                    <InputGroupInput
                      id="port-input"
                      type="number"
                      placeholder="5432"
                      {...register('port')}
                    />
                    <InputGroupAddon>
                      <EthernetPort />
                    </InputGroupAddon>
                  </InputGroup>
                </div>
              </li>
              <li>
                <Label htmlFor="user-input">User</Label>
                <InputGroup>
                  <InputGroupInput id="user-input" placeholder="postgres" {...register('user')} />
                  <InputGroupAddon>
                    <UserRound />
                  </InputGroupAddon>
                </InputGroup>
              </li>
              <li>
                <Label htmlFor="password-input">Password</Label>
                <InputGroup>
                  <InputGroupInput id="password-input" type="password" {...register('password')} />
                  <InputGroupAddon>
                    <KeyRound />
                  </InputGroupAddon>
                </InputGroup>
              </li>
            </ul>
          </div>
          <Button type="submit" loading={isPending} fullWidth size="lg">
            Connect
          </Button>
        </div>
      </form>
      <div className="flex items-center justify-center bg-white/90">
        <div className={styles.view} />
      </div>
    </div>
  )
}
